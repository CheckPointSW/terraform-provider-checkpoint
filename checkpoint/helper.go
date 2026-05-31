package checkpoint

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

type TaskResult struct {
	TaskID    string
	Status    string // "succeeded","failed","timeout","inline","queued","unknown"
	Message   string // best-effort message or pretty JSON from task-details
	Completed bool
	Polled    bool
}

// Returns a pretty JSON string of task-details if present.
func taskDetailsJSON(data map[string]interface{}) (string, bool) {
	// common path: data["tasks"][0]["task-details"]
	if ts, ok := data["tasks"].([]interface{}); ok && len(ts) > 0 {
		if t0, ok := ts[0].(map[string]interface{}); ok {
			if td, ok := t0["task-details"]; ok {
				if b, err := json.MarshalIndent(td, "", "  "); err == nil {
					return string(b), true
				}
			}
		}
	}
	// sometimes present at top-level
	if td, ok := data["task-details"]; ok {
		if b, err := json.MarshalIndent(td, "", "  "); err == nil {
			return string(b), true
		}
	}
	return "", false
}

// IsSuccess is a convenience for callers.
func (r TaskResult) IsSuccess() bool { return r.Status == "succeeded" || r.Status == "inline" }

// HandleTaskCreate runs a create-like command that may return a task and (optionally) polls it.
// Returns TaskResult describing the outcome; error is only for transport/SDK issues.
func HandleTaskCreate(
	ctx context.Context,
	client *checkpoint.ApiClient,
	command string,
	apiRes checkpoint.APIResponse,
	poll bool,
	timeout time.Duration,
) (TaskResult, error) {

	data := normalizeData(apiRes.GetData())

	// If a task-id is present, always treat as async and poll —
	// even if task-details is also present in the submission response.
	tid := getTaskID(data)

	if tid == "" {
		// No task-id: check if the result is inline (synchronous object returned directly)
		if _, ok := objectFromTasks(ctx, data); ok {
			inlineTid := fmt.Sprintf("%s-%d", command, time.Now().UnixNano())
			return TaskResult{
				TaskID:    inlineTid,
				Status:    "inline",
				Message:   "operation completed inline",
				Completed: true,
				Polled:    false,
			}, nil
		}
		msg := fmt.Sprintf("%s returned neither 'tasks' nor 'task-id'", command)
		return TaskResult{Status: "unknown", Message: msg}, fmt.Errorf(msg)
	}

	if !poll {
		return TaskResult{
			TaskID:  tid,
			Status:  "queued",
			Message: "task submitted; polling disabled",
			Polled:  false,
		}, nil
	}

	if timeout <= 0 {
		timeout = 3 * time.Minute
	}
	deadline := time.Now().Add(timeout)
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		if time.Now().After(deadline) {
			log.Printf("[INFO] task %s did not complete within %s", tid, timeout)
			return TaskResult{
				TaskID:  tid,
				Status:  "timeout",
				Message: "task did not complete before timeout",
				Polled:  true,
			}, nil
		}

		select {
		case <-ctx.Done():
			return TaskResult{
				TaskID:  tid,
				Status:  "timeout",
				Message: "context canceled or deadline exceeded",
				Polled:  true,
			}, ctx.Err()

		case <-ticker.C:
			showRes, err := client.ApiCall("show-task",
				map[string]interface{}{"task-id": tid},
				client.GetSessionID(), false, client.IsProxyUsed())
			if err != nil {
				return TaskResult{
					TaskID:  tid,
					Status:  "unknown",
					Message: fmt.Sprintf("show-task transport error: %v", err),
					Polled:  true,
				}, err
			}
			if !showRes.Success {
				return TaskResult{
					TaskID:  tid,
					Status:  "failed",
					Message: "show-task returned success=false",
					Polled:  true,
				}, nil
			}

			showData := normalizeData(showRes.GetData())
			status := TaskStatusFromAny(showData)

			// Try to surface rich details if available
			pretty, hasPretty := taskDetailsJSON(showData)
			msg := TaskMessageFromAny(showData)
			if hasPretty && strings.TrimSpace(pretty) != "" {
				msg = pretty
			}

			switch {
			case status == "succeeded" || status == "finished":
				return TaskResult{
					TaskID:    tid,
					Status:    "succeeded",
					Message:   msg,
					Completed: true,
					Polled:    true,
				}, nil

			case isFailedStatus(status):
				if strings.TrimSpace(msg) == "" {
					msg = "task failed"
				}
				return TaskResult{
					TaskID:    tid,
					Status:    "failed",
					Message:   msg,
					Completed: true,
					Polled:    true,
				}, nil
			}
		}
	}
}

func normalizeData(m map[string]interface{}) map[string]interface{} {
	cur := m
	for {
		next, ok := cur["data"].(map[string]interface{})
		if !ok || next == nil {
			return cur
		}
		cur = next
	}
}

func toString(v interface{}) string {
	switch x := v.(type) {
	case nil:
		return ""
	case string:
		return x
	case *string:
		if x == nil {
			return ""
		}
		return *x
	case []byte:
		return string(x)
	case fmt.Stringer:
		return x.String()
	default:
		// for maps/structs/slices, fall back to JSON
		b, err := json.Marshal(v)
		if err != nil {
			return fmt.Sprintf("%v", v)
		}
		return string(b)
	}
}

func stringify(obj interface{}) string {
	if s, ok := obj.(string); ok {
		return s
	}
	b, err := json.Marshal(obj) // compact JSON
	if err != nil {
		return toString(obj)
	}
	return string(b)
}

// prefer tasks[0].task-details.object or task-details.lightshot (JSON string);
// gracefully handle task-details as map OR []interface{}
func objectFromTasks(ctx context.Context, data map[string]interface{}) (string, bool) {
	// tasks[0]
	arr, ok := data["tasks"].([]interface{})
	if !ok || len(arr) == 0 {
		return "", false
	}
	t0, ok := arr[0].(map[string]interface{})
	if !ok {
		return "", false
	}

	// task-details (hyphen or underscore)
	detRaw, ok := t0["task-details"]
	if !ok {
		detRaw, ok = t0["task_details"]
	}
	if !ok || detRaw == nil {
		return "", false
	}

	switch det := detRaw.(type) {
	case map[string]interface{}:
		// map shape
		if obj, ok := det["object"]; ok {
			out := stringify(obj)
			return out, true
		}
		if ls, ok := det["lightshot"]; ok {
			out := stringify(ls)
			return out, true
		}
		out := stringify(det)
		return out, true

	case []interface{}:
		// list shape (your API example)
		for _, el := range det {
			m, ok := el.(map[string]interface{})
			if !ok {
				continue
			}
			if obj, ok := m["object"]; ok {
				out := stringify(obj)
				return out, true
			}
			if ls, ok := m["lightshot"]; ok {
				out := stringify(ls)
				return out, true
			}
			// single-key convenience
			if len(m) == 1 {
				for _, v := range m {
					out := stringify(v)
					return out, true
				}
			}
		}
		// stringify the whole list as a last resort
		out := stringify(det)
		return out, true

	default:
		// unexpected type
		out := stringify(det)
		return out, true
	}
}

// TaskStatusFromAny extracts tasks[0].status (lowercased), or "" if missing.
func TaskStatusFromAny(data map[string]interface{}) string {
	if arr, ok := data["tasks"].([]interface{}); ok && len(arr) > 0 {
		if t0, ok := arr[0].(map[string]interface{}); ok {
			if s, ok := t0["status"].(string); ok {
				return strings.ToLower(s)
			}
		}
	}
	if s, ok := data["status"].(string); ok {
		return strings.ToLower(s)
	}
	return ""
}

func isFailedStatus(s string) bool {
	switch strings.ToLower(s) {
	case "failed", "error", "canceled", "cancelled":
		return true
	default:
		return false
	}
}

// TaskMessageFromAny pulls a human message if present.
func TaskMessageFromAny(data map[string]interface{}) string {
	// Common keys seen in CP tasks
	keys := []string{"statusDescription", "status-description", "message", "progress-description", "error", "error-message"}
	for _, k := range keys {
		if v, ok := data[k].(string); ok && strings.TrimSpace(v) != "" {
			return v
		}
		// sometimes nested under tasks[0]
	}
	if arr, ok := data["tasks"].([]interface{}); ok && len(arr) > 0 {
		if t0, ok := arr[0].(map[string]interface{}); ok {
			for _, k := range keys {
				if v, ok := t0[k].(string); ok && strings.TrimSpace(v) != "" {
					return v
				}
			}
			// Gaia API surfaces errors under task-details[0]["errors"], not "fault-message".
			if det, ok := t0["task-details"].([]interface{}); ok && len(det) > 0 {
				if d0, ok := det[0].(map[string]interface{}); ok {
					if v, ok := d0["errors"].(string); ok && strings.TrimSpace(v) != "" {
						return v
					}
				}
			}
		}
	}
	return ""
}

// Only the "task-id" key is checked (top-level or tasks[0]).
func getTaskID(data map[string]interface{}) string {
	if v, ok := data["task-id"].(string); ok && v != "" {
		return v
	}
	if arr, ok := data["tasks"].([]interface{}); ok && len(arr) > 0 {
		if t0, ok := arr[0].(map[string]interface{}); ok {
			if v, ok := t0["task-id"].(string); ok && v != "" {
				return v
			}
		}
	}
	return ""
}

func getString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok && v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}
func expandExtendedCommands(s *schema.Set) interface{} {
	if s == nil {
		return nil
	}

	items := s.List()
	if len(items) == 1 {
		if str, ok := items[0].(string); ok && strings.EqualFold(str, "all") {
			return "all"
		}
	}

	return items
}

func flattenExtendedCommands(v interface{}) []interface{} {
	switch val := v.(type) {
	case nil:
		return []interface{}{}
	case string:
		trimmed := strings.TrimSpace(val)
		if trimmed == "" {
			return []interface{}{}
		}
		return []interface{}{trimmed}
	case []string:
		out := make([]interface{}, 0, len(val))
		for _, item := range val {
			if strings.TrimSpace(item) != "" {
				out = append(out, item)
			}
		}
		return out
	case []interface{}:
		return val
	default:
		if str, ok := val.(fmt.Stringer); ok {
			s := strings.TrimSpace(str.String())
			if s == "" {
				return []interface{}{}
			}
			return []interface{}{s}
		}
		return []interface{}{}
	}
}
