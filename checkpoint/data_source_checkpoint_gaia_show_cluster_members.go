package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowClusterMembers() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowClusterMembers,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "members": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hostname": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "serial_number": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "request_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "site_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "member_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "model": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "version": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "member_status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "site_status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "installed_jumbo_take": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "pending_gateways": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hostname": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "serial_number": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "request_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "site_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "member_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "model": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "version": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "member_status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "site_status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "installed_jumbo_take": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
        },
    }
}

func readGaiaShowClusterMembers(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    log.Println("Execute show-cluster-members - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-cluster-members", payload)
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
            "cluster-members",        // resource type
            "read",                       // operation
            "show-cluster-members",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-cluster-members: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["members"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "hostname": func() string { if _v, _ok := m["hostname"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "serial_number": func() string { if _v, _ok := m["serial-number"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "request_id": func() string { if _v, _ok := m["request-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "site_id": func() int { if f, ok := m["site-id"].(float64); ok { return int(f) }; return 0 }(),
                        "member_id": func() int { if f, ok := m["member-id"].(float64); ok { return int(f) }; return 0 }(),
                        "model": func() string { if _v, _ok := m["model"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "version": func() string { if _v, _ok := m["version"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "member_status": func() string { if _v, _ok := m["member-status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "site_status": func() string { if _v, _ok := m["site-status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "state": func() string { if _v, _ok := m["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "installed_jumbo_take": func() int { if f, ok := m["installed-jumbo-take"].(float64); ok { return int(f) }; return 0 }(),
                    }
                }
            }
            d.Set("members", mapped)
        }
    } else {
        d.Set("members", []interface{}{})
    }
    if v, exists := commandRes.GetData()["pending-gateways"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "hostname": func() string { if _v, _ok := m["hostname"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "serial_number": func() string { if _v, _ok := m["serial-number"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "request_id": func() string { if _v, _ok := m["request-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "site_id": func() int { if f, ok := m["site-id"].(float64); ok { return int(f) }; return 0 }(),
                        "member_id": func() int { if f, ok := m["member-id"].(float64); ok { return int(f) }; return 0 }(),
                        "model": func() string { if _v, _ok := m["model"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "version": func() string { if _v, _ok := m["version"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "member_status": func() string { if _v, _ok := m["member-status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "site_status": func() string { if _v, _ok := m["site-status"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "state": func() string { if _v, _ok := m["state"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "installed_jumbo_take": func() int { if f, ok := m["installed-jumbo-take"].(float64); ok { return int(f) }; return 0 }(),
                    }
                }
            }
            d.Set("pending_gateways", mapped)
        }
    } else {
        d.Set("pending_gateways", []interface{}{})
    }
    d.SetId(fmt.Sprintf("show-cluster-members-" + acctest.RandString(10)))
    return nil
}

