package checkpoint

import (
        "context"
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowConnections() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowConnections,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "max_results": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Max number of connection to display, or \"all\" (10 by default)`,
            },
            "use_preset": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `A preset represents a pre-configured request with default values. But any value can be overridden by being explicitly provided in the request.`,
            },
            "preset": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `The names of the presets to be used (API call \"show-connections-presets\" returns the list of available presets)`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "filter": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Connection filtering options`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "instance_id": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `Query a specific instance / set of instances (all by default)`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "source": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Return connections with the given source IP address`,
                        },
                        "destination": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Return connections with the given destination IP address`,
                        },
                        "ip_protocol": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Return connections with the given IP protocol number`,
                        },
                        "source_port": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Return connections with the given source port`,
                        },
                        "destination_port": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Return connections with the given destination port`,
                        },
                        "ip_version": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Return connections of the given IP version (v4 / v6 / any (default))`,
                        },
                    },
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "results": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "data": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowConnections(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("max_results"); ok {
        payload["max-results"] = v.(int)
    }

    if v, ok := d.GetOkExists("use_preset"); ok {
        payload["use-preset"] = v.(bool)
    }

    if v := d.Get("preset"); len(v.(*schema.Set).List()) > 0 {
        payload["preset"] = v.(*schema.Set).List()
    }

    if v := d.Get("filter"); len(v.([]interface{})) > 0 {
        _ = v
        filterMap := make(map[string]interface{})
        if v := d.Get("filter.0.instance_id"); len(v.(*schema.Set).List()) > 0 {
            filterMap["instance-id"] = v.(*schema.Set).List()
        }
        if v, ok := d.GetOk("filter.0.source"); ok {
            filterMap["source"] = v.(string)
        }
        if v, ok := d.GetOk("filter.0.destination"); ok {
            filterMap["destination"] = v.(string)
        }
        if v, ok := d.GetOk("filter.0.ip_protocol"); ok {
            filterMap["ip-protocol"] = v.(int)
        }
        if v, ok := d.GetOk("filter.0.source_port"); ok {
            filterMap["source-port"] = v.(int)
        }
        if v, ok := d.GetOk("filter.0.destination_port"); ok {
            filterMap["destination-port"] = v.(int)
        }
        if v, ok := d.GetOk("filter.0.ip_version"); ok {
            filterMap["ip-version"] = v.(string)
        }
        if len(filterMap) > 0 {
            payload["filter"] = filterMap
        }
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-connections - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-connections", payload)
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
            "connections",        // resource type
            "read",                       // operation
            "show-connections",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-connections: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "show-connections", commandRes, true, 0)
    if err != nil {
        return fmt.Errorf("show-connections task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        msg := taskRes.Message
        if msg == "" {
            msg = fmt.Sprintf("show-connections task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(msg)
    }

    _taskDetailsRes, _tdErr := client.ApiCallSimple("show-task", map[string]interface{}{"task-id": taskRes.TaskID})
    var _asyncRespData map[string]interface{}
    if _tdErr == nil && _taskDetailsRes.Success {
        _td := _taskDetailsRes.GetData()
        if _tasks, _ok := _td["tasks"].([]interface{}); _ok && len(_tasks) > 0 {
            if _task0, _ok := _tasks[0].(map[string]interface{}); _ok {
                if _details, _ok := _task0["task-details"].([]interface{}); _ok && len(_details) > 0 {
                    if _detail0, _ok := _details[0].(map[string]interface{}); _ok {
                        _asyncRespData = _detail0
                    }
                }
            }
        }
    }
    if _asyncRespData == nil {
        _asyncRespData = commandRes.GetData()
    }

    if v, exists := _asyncRespData["use-preset"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("use_preset", b)
        } else if s, ok := v.(string); ok {
            d.Set("use_preset", s == "true")
        }
    }
    if v, exists := _asyncRespData["preset"]; exists {
        d.Set("preset", v.([]interface{}))
    } else {
        d.Set("preset", []interface{}{})
    }
    if v, exists := _asyncRespData["filter"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("filter", []interface{}{map[string]interface{}{
                "instance_id": func() string { if _v, _ok := _m["instance-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "source": func() string { if _v, _ok := _m["source"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "destination": func() string { if _v, _ok := _m["destination"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "ip_protocol": func() int { if f, ok := _m["ip-protocol"].(float64); ok { return int(f) }; return 0 }(),
                "source_port": func() int { if f, ok := _m["source-port"].(float64); ok { return int(f) }; return 0 }(),
                "destination_port": func() int { if f, ok := _m["destination-port"].(float64); ok { return int(f) }; return 0 }(),
                "ip_version": func() string { if _v, _ok := _m["ip-version"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
            }})
        }
    }
    if v, exists := _asyncRespData["results"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("results", int(_f))
        }
    }
    if v, exists := _asyncRespData["data"]; exists {
        d.Set("data", fmt.Sprintf("%v", v))
    }
    if v, exists := _asyncRespData["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-connections-" + acctest.RandString(10)))
    return nil
}

