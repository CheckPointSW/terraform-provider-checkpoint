package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowRemoteSyslogs() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowRemoteSyslogs,
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
            "objects": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "server_ip": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "level": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "port": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "protocol": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "tls_encryption": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "enabled": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "auth_mode": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "permitted_peers": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "queuing_mechanism": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
        },
    }
}

func readGaiaShowRemoteSyslogs(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-remote-syslogs - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-remote-syslogs", payload)
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
            "remote-syslogs",        // resource type
            "read",                       // operation
            "show-remote-syslogs",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-remote-syslogs: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["objects"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "server_ip": func() string { if _v, _ok := m["server-ip"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "level": func() string { if _v, _ok := m["level"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "port": func() int { if f, ok := m["port"].(float64); ok { return int(f) }; return 0 }(),
                        "protocol": func() string { if _v, _ok := m["protocol"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "tls_encryption": func() []interface{} {
                            if _obj, _ok := m["tls-encryption"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "enabled": func() bool { if b, ok := _obj["enabled"].(bool); ok { return b }; if s, ok := _obj["enabled"].(string); ok { return s == "true" }; return false }(),
                                    "auth_mode": func() string { if _v, _ok := _obj["auth-mode"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "permitted_peers": func() []interface{} {
                                        if _sl, _ok := _obj["permitted-peers"].([]interface{}); _ok {
                                            return _sl
                                        }
                                        return nil
                                    }(),
                                }}
                            }
                            return nil
                        }(),
                        "queuing_mechanism": func() bool { if b, ok := m["queuing-mechanism"].(bool); ok { return b }; if s, ok := m["queuing-mechanism"].(string); ok { return s == "true" }; return false }(),
                    }
                }
            }
            d.Set("objects", mapped)
        }
    } else {
        d.Set("objects", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-remote-syslogs-" + acctest.RandString(10)))
    return nil
}

