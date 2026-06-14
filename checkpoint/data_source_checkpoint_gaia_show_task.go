package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowTask() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowTask,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "task_id": {
                Type:        schema.TypeSet,
                Required:    true,
                Description: `task id to show. expiration default time for task id is 1 day, after this time the task id will not be available`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "tasks": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "task_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "last_update_time": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "posix": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "iso_8601": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "progress_description": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "progress_percentage": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "start_time": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "posix": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "iso_8601": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "status_code": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "task_name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "task_details": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "execution_time": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "time_spent_in_queue": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
        },
    }
}

func readGaiaShowTask(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("task_id"); len(v.(*schema.Set).List()) > 0 {
        payload["task-id"] = v.(*schema.Set).List()
    }

    log.Println("Execute show-task - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-task", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && commandRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !commandRes.Success {
            errMsg = commandRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = commandRes.GetData()
        }

        debugLogOperation(
            "task",        // resource type
            "read",                       // operation
            "show-task",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-task: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["tasks"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "task_id": func() string { if _v, _ok := m["task-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "last_update_time": func() []interface{} {
                            if _obj, _ok := m["last-update-time"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "posix": func() string { if _v, _ok := _obj["posix"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "iso_8601": func() string { if _v, _ok := _obj["iso-8601"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "progress_description": func() string { if _v, _ok := m["progress-description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "progress_percentage": func() int { if f, ok := m["progress-percentage"].(float64); ok { return int(f) }; return 0 }(),
                        "start_time": func() []interface{} {
                            if _obj, _ok := m["start-time"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "posix": func() string { if _v, _ok := _obj["posix"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "iso_8601": func() string { if _v, _ok := _obj["iso-8601"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "status_code": func() int { if f, ok := m["status-code"].(float64); ok { return int(f) }; return 0 }(),
                        "task_name": func() string { if _v, _ok := m["task-name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "status": func() string { if _v, _ok := m["status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "task_details": func() []interface{} {
                            switch _ev := m["task-details"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "execution_time": func() string { if _v, _ok := m["execution-time"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "time_spent_in_queue": func() string { if _v, _ok := m["time-spent-in-queue"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("tasks", mapped)
        }
    } else {
        d.Set("tasks", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-task-" + acctest.RandString(10)))
    return nil
}

