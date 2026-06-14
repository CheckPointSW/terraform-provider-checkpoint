package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaOpenTelemetry() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaOpenTelemetry,
        Read:   readGaiaOpenTelemetry,
        Update: updateGaiaOpenTelemetry,
        Delete: deleteGaiaOpenTelemetry,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "export_targets": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `Settings of OpenTelemetry export targets`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the type of this OpenTelemetry Exporter.`,
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Specifies the state of this OpenTelemetry Exporter.`,
                        },
                        "url": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Specifies the URL of this OpenTelemetry Exporter.`,
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Optional. Specifies the user-defined name for this OpenTelemetry Exporter.   You can use this name parameter later to remove or to change settings of this OpenTelemetry Exporter.`,
                        },
                        "client_auth": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Specifies the client authentication method for this OpenTelemetry Exporter.`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "basic": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Specifies the client authentication credentials for this OpenTelemetry Exporter.   If this OpenTelemetry Exporter does not require client authentication, then enter the value \"N/A\" for the username and for the password.`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "username": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `The username used for basic HTTP/HTTPS basic authentication on this OpenTelemetry Exporter.`,
                                                },
                                                "password": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Sensitive:   true,
                                                    Description: `The password used for basic HTTP/HTTPS basic authentication on this OpenTelemetry Exporter.   We recommend to use a bcrypt hash on the password.`,
                                                },
                                            },
                                        },
                                    },
                                    "token": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Sensitive:   true,
                                        Description: `Specifies the client authentication bearer token for this OpenTelemetry Exporter.   If this OpenTelemetry Exporter does not require client authentication, then enter the value \"N/A\" for the bearer token.`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "header_bearer_token": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Sensitive:   true,
                                                    Description: `Specifies the JWT or other similar protocol bearer token.`,
                                                },
                                                "custom_header": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Description: `A custom header for authentication purpose.`,
                                                    MaxItems:    1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "key": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Description: `The key name of the custom header used for authentication.`,
                                                            },
                                                            "value": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                Description: `The value of the custom header used for authentication.`,
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "server_auth": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Specifies the server authentication method.`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ca_public_key": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Specifies the Public key and its type for the Certificate Authority (CA) that this OpenTelemetry Exporter uses.`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `The CA Public Key Type for this OpenTelemetry Exporter - PEM-X509 or the default CA bundle.`,
                                                },
                                                "value": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `The CA Public Key string for this OpenTelemetry Exporter. If you specified the type 'Default', then enter the value 'N/A'.`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "metrics": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Settings to include or exclude which metrics are exporterd from this machine.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "include": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Determines the default behavior for metric inclusion.`,
                        },
                        "except": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `Defines specific metrics to be either excluded (if include is 'all') or included (if include is 'none').`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },
            "enabled": {
                Type:        schema.TypeBool,
                Required:    true,
                Description: `State of OpenTelemetry (enabled or disabled).`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "warnings": {
                Type:     schema.TypeList,
                Computed: true,
                Elem:     &schema.Schema{Type: schema.TypeString},
            },
        },
    }
}

func createGaiaOpenTelemetry(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("export_targets"); len(v.([]interface{})) > 0 {
        exporttargetsList := v.([]interface{})
        exporttargetsArray := make([]interface{}, 0, len(exporttargetsList))
        for i, item := range exporttargetsList {
            if itemMap, ok := item.(map[string]interface{}); ok {
                entry := map[string]interface{}{}
                if val, ok2 := itemMap["name"].(string); ok2 && val != "" {
                    entry["name"] = val
                }
                if val, ok2 := itemMap["url"].(string); ok2 && val != "" {
                    entry["url"] = val
                }
                if v2, ok2 := d.GetOkExists(fmt.Sprintf("export_targets.%d.enabled", i)); ok2 {
                    entry["enabled"] = v2.(bool)
                }
                if caList, ok2 := d.GetOk(fmt.Sprintf("export_targets.%d.client_auth", i)); ok2 {
                    if caItems, ok3 := caList.([]interface{}); ok3 && len(caItems) > 0 {
                        if caMap, ok4 := caItems[0].(map[string]interface{}); ok4 {
                            caEntry := map[string]interface{}{}
                            if basicList, ok5 := caMap["basic"].([]interface{}); ok5 && len(basicList) > 0 {
                                if bm, ok6 := basicList[0].(map[string]interface{}); ok6 {
                                    caEntry["basic"] = map[string]interface{}{
                                        "username": bm["username"],
                                        "password": bm["password"],
                                    }
                                }
                            }
                            if tokenList, ok5 := caMap["token"].([]interface{}); ok5 && len(tokenList) > 0 {
                                if tm, ok6 := tokenList[0].(map[string]interface{}); ok6 {
                                    tokenEntry := map[string]interface{}{}
                                    if hbt, ok7 := tm["header_bearer_token"].(string); ok7 && hbt != "" {
                                        tokenEntry["header-bearer-token"] = hbt
                                    }
                                    if chList, ok7 := tm["custom_header"].([]interface{}); ok7 && len(chList) > 0 {
                                        if chm, ok8 := chList[0].(map[string]interface{}); ok8 {
                                            tokenEntry["custom-header"] = map[string]interface{}{
                                                "key":   chm["key"],
                                                "value": chm["value"],
                                            }
                                        }
                                    }
                                    caEntry["token"] = tokenEntry
                                }
                            }
                            entry["client-auth"] = caEntry
                        }
                    }
                }
                if saList, ok2 := d.GetOk(fmt.Sprintf("export_targets.%d.server_auth", i)); ok2 {
                    if saItems, ok3 := saList.([]interface{}); ok3 && len(saItems) > 0 {
                        if saMap, ok4 := saItems[0].(map[string]interface{}); ok4 {
                            saEntry := map[string]interface{}{}
                            if cpkList, ok5 := saMap["ca_public_key"].([]interface{}); ok5 && len(cpkList) > 0 {
                                if cpkm, ok6 := cpkList[0].(map[string]interface{}); ok6 {
                                    saEntry["ca-public-key"] = map[string]interface{}{
                                        "type":  cpkm["type"],
                                        "value": cpkm["value"],
                                    }
                                }
                            }
                            entry["server-auth"] = saEntry
                        }
                    }
                }
                exporttargetsArray = append(exporttargetsArray, entry)
            }
        }
        payload["export-targets"] = exporttargetsArray
    }

    if v := d.Get("metrics"); len(v.([]interface{})) > 0 {
        _ = v
        metricsMap := make(map[string]interface{})
        if v, ok := d.GetOk("metrics.0.include"); ok {
            metricsMap["include"] = v.(string)
        }
        if v := d.Get("metrics.0.except"); len(v.(*schema.Set).List()) > 0 {
            metricsMap["except"] = v.(*schema.Set).List()
        }
        if len(metricsMap) > 0 {
            payload["metrics"] = metricsMap
        }
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    log.Println("Create OpenTelemetry - Map = ", payload)

    addOpenTelemetryRes, err := client.ApiCallSimple("set-open-telemetry", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addOpenTelemetryRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addOpenTelemetryRes.Success {
            errMsg = addOpenTelemetryRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addOpenTelemetryRes.GetData()
        }

        debugLogOperation(
            "open-telemetry",        // resource type
            "create",                       // operation
            "set-open-telemetry",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add open-telemetry: %v", err)
    }
    if !addOpenTelemetryRes.Success {
        if addOpenTelemetryRes.ErrorMsg != "" {
            return fmt.Errorf(addOpenTelemetryRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("open-telemetry-" + acctest.RandString(10)))
    return readGaiaOpenTelemetry(d, m)
}

func readGaiaOpenTelemetry(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showOpenTelemetryRes, err := client.ApiCallSimple("show-open-telemetry", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showOpenTelemetryRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showOpenTelemetryRes.Success {
            errMsg = showOpenTelemetryRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showOpenTelemetryRes.GetData()
        }

        debugLogOperation(
            "open-telemetry",        // resource type
            "read",                       // operation
            "show-open-telemetry",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show open-telemetry: %v", err)
    }
    if !showOpenTelemetryRes.Success {
        if data := showOpenTelemetryRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showOpenTelemetryRes.ErrorMsg)
    }

    openTelemetry := showOpenTelemetryRes.GetData()

    log.Println("Read OpenTelemetry - Show JSON = ", openTelemetry)

    if v, exists := openTelemetry["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := openTelemetry["export-targets"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    entry := map[string]interface{}{
                        "name":    fmt.Sprintf("%v", m["name"]),
                        "enabled": func() bool { b, _ := m["enabled"].(bool); return b }(),
                        "url":     fmt.Sprintf("%v", m["url"]),
                    }
                    if ca, ok := m["client-auth"].(map[string]interface{}); ok {
                        caEntry := map[string]interface{}{}
                        if basic, ok := ca["basic"].(map[string]interface{}); ok {
                            caEntry["basic"] = []interface{}{map[string]interface{}{
                                "username": fmt.Sprintf("%v", basic["username"]),
                                "password": fmt.Sprintf("%v", basic["password"]),
                            }}
                        }
                        if token, ok := ca["token"].(map[string]interface{}); ok {
                            tokenEntry := map[string]interface{}{
                                "header_bearer_token": fmt.Sprintf("%v", token["header-bearer-token"]),
                            }
                            if ch, ok := token["custom-header"].(map[string]interface{}); ok {
                                tokenEntry["custom_header"] = []interface{}{map[string]interface{}{
                                    "key":   fmt.Sprintf("%v", ch["key"]),
                                    "value": fmt.Sprintf("%v", ch["value"]),
                                }}
                            }
                            caEntry["token"] = []interface{}{tokenEntry}
                        }
                        entry["client_auth"] = []interface{}{caEntry}
                    }
                    if sa, ok := m["server-auth"].(map[string]interface{}); ok {
                        saEntry := map[string]interface{}{}
                        if cpk, ok := sa["ca-public-key"].(map[string]interface{}); ok {
                            saEntry["ca_public_key"] = []interface{}{map[string]interface{}{
                                "type":  fmt.Sprintf("%v", cpk["type"]),
                                "value": fmt.Sprintf("%v", cpk["value"]),
                            }}
                        }
                        entry["server_auth"] = []interface{}{saEntry}
                    }
                    out = append(out, entry)
                }
            }
            d.Set("export_targets", out)
        }
    }
    if v, exists := openTelemetry["warnings"]; exists {
        d.Set("warnings", v.([]interface{}))
    }
    if v, exists := openTelemetry["metrics"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            exceptList := []interface{}{}
            if ex, ok := m["except"].([]interface{}); ok {
                exceptList = ex
            }
            d.Set("metrics", []interface{}{map[string]interface{}{
                "include": fmt.Sprintf("%v", m["include"]),
                "except":  exceptList,
            }})
        }
    }
    if v, exists := openTelemetry["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaOpenTelemetry(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("export_targets"); len(v.([]interface{})) > 0 {
        exporttargetsList := v.([]interface{})
        exporttargetsArray := make([]interface{}, 0, len(exporttargetsList))
        for i, item := range exporttargetsList {
            if itemMap, ok := item.(map[string]interface{}); ok {
                entry := map[string]interface{}{}
                if val, ok2 := itemMap["name"].(string); ok2 && val != "" {
                    entry["name"] = val
                }
                if val, ok2 := itemMap["url"].(string); ok2 && val != "" {
                    entry["url"] = val
                }
                if v2, ok2 := d.GetOkExists(fmt.Sprintf("export_targets.%d.enabled", i)); ok2 {
                    entry["enabled"] = v2.(bool)
                }
                if caList, ok2 := d.GetOk(fmt.Sprintf("export_targets.%d.client_auth", i)); ok2 {
                    if caItems, ok3 := caList.([]interface{}); ok3 && len(caItems) > 0 {
                        if caMap, ok4 := caItems[0].(map[string]interface{}); ok4 {
                            caEntry := map[string]interface{}{}
                            if basicList, ok5 := caMap["basic"].([]interface{}); ok5 && len(basicList) > 0 {
                                if bm, ok6 := basicList[0].(map[string]interface{}); ok6 {
                                    caEntry["basic"] = map[string]interface{}{
                                        "username": bm["username"],
                                        "password": bm["password"],
                                    }
                                }
                            }
                            if tokenList, ok5 := caMap["token"].([]interface{}); ok5 && len(tokenList) > 0 {
                                if tm, ok6 := tokenList[0].(map[string]interface{}); ok6 {
                                    tokenEntry := map[string]interface{}{}
                                    if hbt, ok7 := tm["header_bearer_token"].(string); ok7 && hbt != "" {
                                        tokenEntry["header-bearer-token"] = hbt
                                    }
                                    if chList, ok7 := tm["custom_header"].([]interface{}); ok7 && len(chList) > 0 {
                                        if chm, ok8 := chList[0].(map[string]interface{}); ok8 {
                                            tokenEntry["custom-header"] = map[string]interface{}{
                                                "key":   chm["key"],
                                                "value": chm["value"],
                                            }
                                        }
                                    }
                                    caEntry["token"] = tokenEntry
                                }
                            }
                            entry["client-auth"] = caEntry
                        }
                    }
                }
                if saList, ok2 := d.GetOk(fmt.Sprintf("export_targets.%d.server_auth", i)); ok2 {
                    if saItems, ok3 := saList.([]interface{}); ok3 && len(saItems) > 0 {
                        if saMap, ok4 := saItems[0].(map[string]interface{}); ok4 {
                            saEntry := map[string]interface{}{}
                            if cpkList, ok5 := saMap["ca_public_key"].([]interface{}); ok5 && len(cpkList) > 0 {
                                if cpkm, ok6 := cpkList[0].(map[string]interface{}); ok6 {
                                    saEntry["ca-public-key"] = map[string]interface{}{
                                        "type":  cpkm["type"],
                                        "value": cpkm["value"],
                                    }
                                }
                            }
                            entry["server-auth"] = saEntry
                        }
                    }
                }
                exporttargetsArray = append(exporttargetsArray, entry)
            }
        }
        payload["export-targets"] = exporttargetsArray
    }

    if v := d.Get("metrics"); len(v.([]interface{})) > 0 {
        _ = v
        metricsMap := make(map[string]interface{})
        if v, ok := d.GetOk("metrics.0.include"); ok {
            metricsMap["include"] = v.(string)
        }
        if v := d.Get("metrics.0.except"); len(v.(*schema.Set).List()) > 0 {
            metricsMap["except"] = v.(*schema.Set).List()
        }
        if len(metricsMap) > 0 {
            payload["metrics"] = metricsMap
        }
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    setOpenTelemetryRes, err := client.ApiCallSimple("set-open-telemetry", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setOpenTelemetryRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setOpenTelemetryRes.Success {
            errMsg = setOpenTelemetryRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setOpenTelemetryRes.GetData()
        }

        debugLogOperation(
            "open-telemetry",        // resource type
            "update",                       // operation
            "set-open-telemetry",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set open-telemetry: %v", err)
    }
    if !setOpenTelemetryRes.Success {
        return fmt.Errorf(setOpenTelemetryRes.ErrorMsg)
    }

    return readGaiaOpenTelemetry(d, m)
}

func deleteGaiaOpenTelemetry(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    