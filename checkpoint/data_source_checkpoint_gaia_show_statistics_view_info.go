package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowStatisticsViewInfo() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowStatisticsViewInfo,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "filter": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `Filter the results by a list of view-IDs`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "views": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "view_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "parent": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "children": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "stat_ids": {
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
        },
    }
}

func readGaiaShowStatisticsViewInfo(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("filter"); len(v.(*schema.Set).List()) > 0 {
        payload["filter"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-statistics-view-info - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-statistics-view-info", payload)
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
            "statistics-view-info",        // resource type
            "read",                       // operation
            "show-statistics-view-info",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-statistics-view-info: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["views"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "view_id": func() string { if _v, _ok := m["view-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "description": func() string { if _v, _ok := m["description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "parent": func() string { if _v, _ok := m["parent"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "children": func() []interface{} {
                            switch _ev := m["children"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "stat_ids": func() []interface{} {
                            switch _ev := m["stat-ids"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                    }
                }
            }
            d.Set("views", mapped)
        }
    } else {
        d.Set("views", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-statistics-view-info-" + acctest.RandString(10)))
    return nil
}

