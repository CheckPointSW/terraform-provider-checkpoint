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
	_ = os.MkdirAll(dir, 0o755)
	path := filepath.Join(dir, "terraform-gw-requests.jsonl")

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0o644)
	if err != nil {
		// never break Terraform because debug logging failed
		return
	}
	defer f.Close()

	enc := json.NewEncoder(f)
	_ = enc.Encode(&rec)
}

func safeCopyMap(in map[string]interface{}) map[string]interface{} {
	if in == nil {
		return nil
	}
	out := make(map[string]interface{}, len(in))
	for k, v := range in {
		out[k] = v
	}
	return out
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
