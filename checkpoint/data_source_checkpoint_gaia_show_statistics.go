package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowStatistics() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowStatistics,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "new_query": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Create new query.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "filter": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `Filter the results by a list of labels and stat-IDs`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "show_info": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Data response contains additional information.`,
                        },
                        "ignore_warnings": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Ignore all warnings.`,
                        },
                        "historical_data": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Data response should contains historical data.. Supported starting from Gaia version R82.10`,
                        },
                        "from_date": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The starting date (posix/iso-8601) for the query.. Supported starting from Gaia version R82.10`,
                        },
                        "to_date": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The ending date (posix/iso-8601) for the query.. Supported starting from Gaia version R82.10`,
                        },
                        "records_limit": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Limit the amount of records in the result.. Supported starting from Gaia version R82.10`,
                        },
                        "override_db_name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Override the historical database name (Notice you can only provide a name - You must upload the DB as <Name>.dat).. Supported starting from Gaia version R82.10`,
                        },
                    },
                },
            },
            "use_cursor": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Request for data on existing query.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "cursor_id": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The cursor ID for an existing query. For cursor based pagination.`,
                        },
                    },
                },
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "data": {
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
                        "description": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "records": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "timestamp": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "posix": {
                                                    Type:        schema.TypeInt,
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
                                    "value": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "stat_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "table_columns": {
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
                                    "description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "column_stat_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "type": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "next_cursor": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowStatistics(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("new_query"); len(v.([]interface{})) > 0 {
        _ = v
        newqueryMap := make(map[string]interface{})
        if v := d.Get("new_query.0.filter"); len(v.(*schema.Set).List()) > 0 {
            newqueryMap["filter"] = v.(*schema.Set).List()
        }
        if v, ok := d.GetOkExists("new_query.0.show_info"); ok && v.(bool) {
            newqueryMap["show-info"] = v.(bool)
        }
        if v, ok := d.GetOkExists("new_query.0.ignore_warnings"); ok && v.(bool) {
            newqueryMap["ignore-warnings"] = v.(bool)
        }
        if v, ok := d.GetOkExists("new_query.0.historical_data"); ok && v.(bool) {
            newqueryMap["historical-data"] = v.(bool)
        }
        if v, ok := d.GetOk("new_query.0.from_date"); ok {
            newqueryMap["from-date"] = v.(string)
        }
        if v, ok := d.GetOk("new_query.0.to_date"); ok {
            newqueryMap["to-date"] = v.(string)
        }
        if v, ok := d.GetOk("new_query.0.records_limit"); ok {
            newqueryMap["records-limit"] = v.(int)
        }
        if v, ok := d.GetOk("new_query.0.override_db_name"); ok {
            newqueryMap["override-db-name"] = v.(string)
        }
        if len(newqueryMap) > 0 {
            payload["new-query"] = newqueryMap
        }
    }

    if v := d.Get("use_cursor"); len(v.([]interface{})) > 0 {
        _ = v
        usecursorMap := make(map[string]interface{})
        if v, ok := d.GetOk("use_cursor.0.cursor_id"); ok {
            usecursorMap["cursor-id"] = v.(string)
        }
        if len(usecursorMap) > 0 {
            payload["use-cursor"] = usecursorMap
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-statistics - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-statistics", payload)
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
            "statistics",        // resource type
            "read",                       // operation
            "show-statistics",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-statistics: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["data"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "description": func() string { if _v, _ok := m["description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "records": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["records"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "timestamp": func() []interface{} {
                                                if _dobj, _ok := _sgm["timestamp"].(map[string]interface{}); _ok {
                                                    return []interface{}{map[string]interface{}{
                                                        "posix": func() int { if f, ok := _dobj["posix"].(float64); ok { return int(f) }; return 0 }(),
                                                        "iso_8601": func() string { if _v, _ok := _dobj["iso-8601"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                                    }}
                                                }
                                                return nil
                                            }(),
                                            "value": func() string { if _v, _ok := _sgm["value"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "stat_id": func() string { if _v, _ok := m["stat-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "table_columns": func() []interface{} {
                            var _sgOut []interface{}
                            if _arr, _ok := m["table-columns"].([]interface{}); _ok {
                                for _, _sg := range _arr {
                                    if _sgm, _ok := _sg.(map[string]interface{}); _ok {
                                        _sgOut = append(_sgOut, map[string]interface{}{
                                            "name": func() string { if _v, _ok := _sgm["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "description": func() string { if _v, _ok := _sgm["description"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "column_stat_id": func() string { if _v, _ok := _sgm["column-stat-id"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            "type": func() string { if _v, _ok := _sgm["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                        })
                                    }
                                }
                            }
                            return _sgOut
                        }(),
                        "type": func() string { if _v, _ok := m["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("data", mapped)
        }
    } else {
        d.Set("data", []interface{}{})
    }
    if v, exists := commandRes.GetData()["next-cursor"]; exists {
        d.Set("next_cursor", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-statistics-" + acctest.RandString(10)))
    return nil
}

