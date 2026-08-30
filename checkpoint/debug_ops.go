package checkpoint

import (
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type debugRecord struct {
	Timestamp      time.Time              `json:"timestamp"`
	Server         string                 `json:"server"`        // GAiA IP/hostname (best effort)
	ResourceType   string                 `json:"resource_type"`
	Operation      string                 `json:"operation"`
	ApiCall        string                 `json:"api_call"`
	RequestPayload map[string]interface{} `json:"request_payload"`
	ResponseData   map[string]interface{} `json:"response_data"`
	Success        bool                   `json:"success"`
	ErrorMsg       string                 `json:"error_msg"`
	Classification string                 `json:"classification"`
}

var (
	debugServer   string
	serverInitOnce sync.Once
)

func resourceDebugEnabled(d *schema.ResourceData) bool {
	if os.Getenv("TF_CP_DEBUG") != "" {
		return true
	}
	if v, ok := d.GetOk("debug"); ok {
		return v.(bool)
	}
	return false
}

// Called from resources with the API client; we try to extract server/host once via reflection.
func ensureDebugServerFromClient(client interface{}) {
	serverInitOnce.Do(func() {
		if client == nil {
			return
		}
		v := reflect.ValueOf(client)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() != reflect.Struct {
			return
		}
		t := v.Type()

		// field names we consider as "server"
		candidates := []string{
			"Server",
			"Host",
			"Hostname",
			"ManagementServer",
			"MgmtServer",
			"Address",
			"IP",
			"Gateway",
			"Url",
			"URL",
		}

		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i)
			fv := v.Field(i)
			if fv.Kind() != reflect.String {
				continue
			}
			name := f.Name
			for _, cand := range candidates {
				if strings.EqualFold(name, cand) {
					debugServer = fv.String()
					if debugServer != "" {
						return
					}
				}
			}
		}
	})
}



func debugDir() string {
	if v := os.Getenv("TF_CP_DEBUG_DIR"); v != "" {
		return v
	}
	return "/tmp/tf-cp-debug"
}

// Generic logger for any resource
// resourceType: "gaia_snmp_custom_trap"
// operation:    "create"/"read"/"update"/"delete"
// apiCall:      "add-snmp-custom-trap", "show-snmp-custom-trap", etc.
func debugLogOperation(
	resourceType, operation, apiCall string,
	request map[string]interface{},
	response map[string]interface{},
	success bool,
	errMsg string,
) {

	rec := debugRecord{
		Timestamp:      time.Now(),
		Server:         debugServer, // may be empty if we couldn't detect it
		ResourceType:   resourceType,
		Operation:      operation,
		ApiCall:        apiCall,
		RequestPayload: safeCopyMap(request),
		ResponseData:   safeCopyMap(response),
		Success:        success,
		ErrorMsg:       errMsg,
		Classification: classifyError(success, errMsg, response),
	}

	dir := debugDir()
	if err := os.MkdirAll(dir, 0o700); err != nil {
		return
	}
	// refuse symlinks: Chmod/OpenFile would follow one out of our own directory
	if fi, err := os.Lstat(dir); err != nil || fi.Mode()&os.ModeSymlink != 0 {
		return
	}
	// re-tighten: modes above apply only on creation, not to pre-existing paths
	_ = os.Chmod(dir, 0o700)
	path := filepath.Join(dir, "terraform-gw-requests.jsonl")
	if fi, err := os.Lstat(path); err == nil && fi.Mode()&os.ModeSymlink != 0 {
		return
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o600)
	if err != nil {
		// never break Terraform because debug logging failed
		return
	}
	defer f.Close()
	_ = f.Chmod(0o600)

	enc := json.NewEncoder(f)
	_ = enc.Encode(&rec)
}

// credential-bearing keys that carry no "password"/"secret" token in their name
var sensitiveKeys = map[string]bool{
	"token":                true,
	"header-bearer-token":  true,
	"community-string":     true,
	"read-only-community":  true,
	"read-write-community": true,
	"sid":                  true,
	"session-id":           true,
	"activation-key":       true,
	"verification-code":    true,
	"private-key":          true,
}

// isSensitiveKey reports whether a payload key holds a credential that must not
// be written to the debug log. Matches the explicit list above plus any key
// whose name embeds a credential token. Over-matching is safe: this only
// affects the logged copy, never the payload sent to the API.
func isSensitiveKey(k string) bool {
	// normalise separators: payload keys mix private_key and private-key spellings
	lk := strings.ReplaceAll(strings.ToLower(k), "_", "-")
	if sensitiveKeys[lk] {
		return true
	}
	return strings.Contains(lk, "password") ||
		strings.Contains(lk, "secret") ||
		strings.Contains(lk, "passphrase") ||
		strings.Contains(lk, "psk")
}

// bounded so a malformed or self-referential payload can never panic Terraform
const maxRedactDepth = 32

func safeCopyMap(in map[string]interface{}) map[string]interface{} {
	return redactMap(in, 0)
}

func redactMap(in map[string]interface{}, depth int) map[string]interface{} {
	if in == nil {
		return nil
	}
	out := make(map[string]interface{}, len(in))
	for k, v := range in {
		if isSensitiveKey(k) {
			out[k] = redactSensitive(v)
			continue
		}
		out[k] = redactValue(v, depth+1)
	}
	return out
}

// redactSensitive masks a value stored under a sensitive key. Numbers and
// booleans cannot carry a credential, so policy knobs such as
// password-expiration-days stay readable; strings and sub-objects are masked.
func redactSensitive(v interface{}) interface{} {
	switch v.(type) {
	case nil, bool, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		return v
	default:
		return "****"
	}
}

// redactValue copies v, redacting sensitive keys at any nesting depth. Secrets
// nest inside BGP authtype, RADIUS/TACACS servers and ISIS authentication, and
// arrive as typed collections as well as plain JSON containers.
func redactValue(v interface{}, depth int) interface{} {
	if depth > maxRedactDepth {
		return "****"
	}

	switch vv := v.(type) {
	case map[string]interface{}:
		return redactMap(vv, depth)
	case []interface{}:
		out := make([]interface{}, len(vv))
		for i, child := range vv {
			out[i] = redactValue(child, depth+1)
		}
		return out
	}

	// typed containers such as expandList's []map[string]interface{} match
	// neither case above, so walk them reflectively rather than logging as-is
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map:
		if rv.Type().Key().Kind() != reflect.String {
			return "****"
		}
		out := make(map[string]interface{}, rv.Len())
		iter := rv.MapRange()
		for iter.Next() {
			k := iter.Key().String()
			if isSensitiveKey(k) {
				out[k] = redactSensitive(iter.Value().Interface())
				continue
			}
			out[k] = redactValue(iter.Value().Interface(), depth+1)
		}
		return out
	case reflect.Slice, reflect.Array:
		if rv.Type().Elem().Kind() == reflect.Uint8 {
			return v // []byte: keep as-is rather than exploding into elements
		}
		out := make([]interface{}, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			out[i] = redactValue(rv.Index(i).Interface(), depth+1)
		}
		return out
	case reflect.Ptr, reflect.Interface:
		if rv.IsNil() {
			return v
		}
		return redactValue(rv.Elem().Interface(), depth+1)
	default:
		return v
	}
}

// Heuristic classifier: gateway_issue / provider_bug / schema_or_user_misuse / ok
func classifyError(success bool, errMsg string, resp map[string]interface{}) string {
	if success {
		return "ok"
	}

	lower := strings.ToLower(errMsg)

	// 5xx-ish / internal errors => gateway
	if strings.Contains(lower, "internal server error") ||
		strings.Contains(lower, "internal error") ||
		strings.Contains(lower, "gateway") {
		return "gateway_issue"
	}

	// GAiA-style response fields if present
	if resp != nil {
		if code, ok := resp["code"].(string); ok {
			lc := strings.ToLower(code)
			if strings.Contains(lc, "not_found") || strings.Contains(lc, "object_not_found") {
				return "schema_or_user_misuse"
			}
			if strings.Contains(lc, "internal_error") {
				return "gateway_issue"
			}
		}
		if msg, ok := resp["message"].(string); ok {
			lm := strings.ToLower(msg)
			if strings.Contains(lm, "already exists") ||
				strings.Contains(lm, "does not exist") ||
				strings.Contains(lm, "mandatory") ||
				strings.Contains(lm, "invalid value") {
				return "schema_or_user_misuse"
			}
			if strings.Contains(lm, "internal error") {
				return "gateway_issue"
			}
		}
	}

	// smells like provider mapping / JSON issue
	if strings.Contains(lower, "unsupported parameter") ||
		strings.Contains(lower, "unknown parameter") ||
		strings.Contains(lower, "json") {
		return "provider_bug"
	}

	// default: blame usage
	return "schema_or_user_misuse"
}
