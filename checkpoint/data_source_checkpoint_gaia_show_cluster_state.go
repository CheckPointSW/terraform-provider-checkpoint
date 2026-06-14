package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowClusterState() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowClusterState,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "mode": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "message": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "this_cluster_member": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "peerid": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "load": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "other_cluster_members": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "peerid": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "load": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "cluster_status": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "additional_info": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowClusterState(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-cluster-state - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-cluster-state", payload)
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
            "cluster-state",        // resource type
            "read",                       // operation
            "show-cluster-state",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-cluster-state: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["mode"]; exists {
        d.Set("mode", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["message"]; exists {
        d.Set("message", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["this-cluster-member"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("this_cluster_member", []interface{}{map[string]interface{}{
                "peerid": func() int { if f, ok := _m["peerid"].(float64); ok { return int(f) }; return 0 }(),
                "status": func() string { if _v, _ok := _m["status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "name": func() string { if _v, _ok := _m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "load": func() int { if f, ok := _m["load"].(float64); ok { return int(f) }; return 0 }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["other-cluster-members"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "peerid": func() int { if f, ok := m["peerid"].(float64); ok { return int(f) }; return 0 }(),
                        "status": func() string { if _v, _ok := m["status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "load": func() int { if f, ok := m["load"].(float64); ok { return int(f) }; return 0 }(),
                    }
                }
            }
            d.Set("other_cluster_members", mapped)
        }
    } else {
        d.Set("other_cluster_members", []interface{}{})
    }
    if v, exists := commandRes.GetData()["cluster-status"]; exists {
        d.Set("cluster_status", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["additional-info"]; exists {
        d.Set("additional_info", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-cluster-state-" + acctest.RandString(10)))
    return nil
}

