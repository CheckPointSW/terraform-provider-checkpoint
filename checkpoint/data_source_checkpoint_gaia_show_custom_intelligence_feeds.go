package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowCustomIntelligenceFeeds() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowCustomIntelligenceFeeds,
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
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "protocol": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "url": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "action": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "account_name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "account_password": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Sensitive:   true,
                            Description: `N/A`,
                        },
                        "format": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "custom_csv_settings": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "csv_delimiter": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "csv_lines_to_be_skipped": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "csv_observable_name": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "csv_observable_value": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "csv_observable_type": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "csv_observable_description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "csv_observable_confidence": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "csv_observable_severity": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "csv_observable_product": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
    }
}

func readGaiaShowCustomIntelligenceFeeds(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-custom-intelligence-feeds - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-custom-intelligence-feeds", payload)
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
            "custom-intelligence-feeds",        // resource type
            "read",                       // operation
            "show-custom-intelligence-feeds",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-custom-intelligence-feeds: %v", err)
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
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "protocol": func() string { if _v, _ok := m["protocol"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "url": func() string { if _v, _ok := m["url"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "enabled": func() bool { if b, ok := m["enabled"].(bool); ok { return b }; if s, ok := m["enabled"].(string); ok { return s == "true" }; return false }(),
                        "action": func() string { if _v, _ok := m["action"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "account_name": func() string { if _v, _ok := m["account-name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "account_password": func() string { if _v, _ok := m["account-password"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "format": func() string { if _v, _ok := m["format"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "custom_csv_settings": func() []interface{} {
                            if _obj, _ok := m["custom-csv-settings"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "csv_delimiter": func() string { if _v, _ok := _obj["csv-delimiter"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "csv_lines_to_be_skipped": func() string { if _v, _ok := _obj["csv-lines-to-be-skipped"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "csv_observable_name": func() string { if _v, _ok := _obj["csv-observable-name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "csv_observable_value": func() int { if f, ok := _obj["csv-observable-value"].(float64); ok { return int(f) }; return 0 }(),
                                    "csv_observable_type": func() string { if _v, _ok := _obj["csv-observable-type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "csv_observable_description": func() string { if _v, _ok := _obj["csv-observable-description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "csv_observable_confidence": func() string { if _v, _ok := _obj["csv-observable-confidence"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "csv_observable_severity": func() string { if _v, _ok := _obj["csv-observable-severity"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "csv_observable_product": func() string { if _v, _ok := _obj["csv-observable-product"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
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
    d.SetId(fmt.Sprintf("show-custom-intelligence-feeds-" + acctest.RandString(10)))
    return nil
}

