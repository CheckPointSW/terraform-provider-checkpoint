package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "context"
    "strings"

)
func resourceGaiaCustomIntelligenceFeed() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaCustomIntelligenceFeed,
        Read:   readGaiaCustomIntelligenceFeed,
        Update: updateGaiaCustomIntelligenceFeed,
        Delete: deleteGaiaCustomIntelligenceFeed,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "protocol": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `N/A`,
            },
            "url": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Set the feed URL`,
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `N/A`,
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `N/A`,
            },
            "action": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Set feed action`,
            },
            "account_name": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `N/A`,
            },
            "account_password": {
                Type:        schema.TypeString,
                Optional:    true,
                Sensitive:   true,
                Description: `N/A`,
            },
            "custom_csv_settings": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Define custom csv settings - CSV structure, Delimiter and rows to skip`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "csv_delimiter": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "csv_lines_to_be_skipped": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "csv_observable_name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set integer for index location in CSV file, or fixed value for the entire feed.`,
                        },
                        "csv_observable_value": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "csv_observable_type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set integer for index location in CSV file, or fixed value for the entire feed.`,
                        },
                        "csv_observable_description": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Set integer for index location in CSV file, or fixed value for the entire feed.`,
                        },
                        "csv_observable_confidence": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "csv_observable_severity": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "csv_observable_product": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "format": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `STIX: https://stixproject.github.io/. For more info see sk132193`,
            },
            "https_sha256_fingerprint": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Specify HTTPS SHA-256 fingerprint`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaCustomIntelligenceFeed(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("protocol"); ok {
        payload["protocol"] = v.(string)
    }

    if v, ok := d.GetOk("url"); ok {
        payload["url"] = v.(string)
    }

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("action"); ok {
        payload["action"] = v.(string)
    }

    if v, ok := d.GetOk("account_name"); ok {
        payload["account-name"] = v.(string)
    }

    if v, ok := d.GetOk("account_password"); ok {
        payload["account-password"] = v.(string)
    }

    if v := d.Get("custom_csv_settings"); len(v.([]interface{})) > 0 {
        _ = v
        customcsvsettingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("custom_csv_settings.0.csv_delimiter"); ok {
            customcsvsettingsMap["csv-delimiter"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_lines_to_be_skipped"); ok {
            customcsvsettingsMap["csv-lines-to-be-skipped"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_name"); ok {
            customcsvsettingsMap["csv-observable-name"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_value"); ok {
            customcsvsettingsMap["csv-observable-value"] = v.(int)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_type"); ok {
            customcsvsettingsMap["csv-observable-type"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_description"); ok {
            customcsvsettingsMap["csv-observable-description"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_confidence"); ok {
            customcsvsettingsMap["csv-observable-confidence"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_severity"); ok {
            customcsvsettingsMap["csv-observable-severity"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_product"); ok {
            customcsvsettingsMap["csv-observable-product"] = v.(string)
        }
        if len(customcsvsettingsMap) > 0 {
            payload["custom-csv-settings"] = customcsvsettingsMap
        }
    }

    if v, ok := d.GetOk("format"); ok {
        payload["format"] = v.(string)
    }

    if v, ok := d.GetOk("https_sha256_fingerprint"); ok {
        payload["https-sha256-fingerprint"] = v.(string)
    }

    log.Println("Create CustomIntelligenceFeed - Map = ", payload)

    addCustomIntelligenceFeedRes, err := client.ApiCallSimple("add-custom-intelligence-feed", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addCustomIntelligenceFeedRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addCustomIntelligenceFeedRes.Success {
            errMsg = addCustomIntelligenceFeedRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addCustomIntelligenceFeedRes.GetData()
        }

        debugLogOperation(
            "custom-intelligence-feed",        // resource type
            "create",                       // operation
            "add-custom-intelligence-feed",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add custom-intelligence-feed: %v", err)
    }
    if !addCustomIntelligenceFeedRes.Success {
        if addCustomIntelligenceFeedRes.ErrorMsg != "" {
            return fmt.Errorf(addCustomIntelligenceFeedRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "add-custom-intelligence-feed", addCustomIntelligenceFeedRes, true, 0)
    if err != nil {
        return fmt.Errorf("add-custom-intelligence-feed task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("add-custom-intelligence-feed task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    d.SetId(fmt.Sprintf("custom-intelligence-feed-" + acctest.RandString(10)))
    return readGaiaCustomIntelligenceFeed(d, m)
}

func readGaiaCustomIntelligenceFeed(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showCustomIntelligenceFeedRes, err := client.ApiCallSimple("show-custom-intelligence-feed", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showCustomIntelligenceFeedRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showCustomIntelligenceFeedRes.Success {
            errMsg = showCustomIntelligenceFeedRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showCustomIntelligenceFeedRes.GetData()
        }

        debugLogOperation(
            "custom-intelligence-feed",        // resource type
            "read",                       // operation
            "show-custom-intelligence-feed",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show custom-intelligence-feed: %v", err)
    }
    if !showCustomIntelligenceFeedRes.Success {
        if data := showCustomIntelligenceFeedRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showCustomIntelligenceFeedRes.ErrorMsg)
    }

    customIntelligenceFeed := showCustomIntelligenceFeedRes.GetData()

    log.Println("Read CustomIntelligenceFeed - Show JSON = ", customIntelligenceFeed)

    if v, exists := customIntelligenceFeed["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := customIntelligenceFeed["protocol"]; exists {
        d.Set("protocol", fmt.Sprintf("%v", v))
    }
    if v, exists := customIntelligenceFeed["url"]; exists {
        d.Set("url", fmt.Sprintf("%v", v))
    }
    if v, exists := customIntelligenceFeed["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := customIntelligenceFeed["action"]; exists {
        d.Set("action", fmt.Sprintf("%v", v))
    }
    if v, exists := customIntelligenceFeed["account-name"]; exists {
        d.Set("account_name", fmt.Sprintf("%v", v))
    }
    if v, exists := customIntelligenceFeed["account-password"]; exists {
        d.Set("account_password", fmt.Sprintf("%v", v))
    }
    if v, exists := customIntelligenceFeed["format"]; exists {
        d.Set("format", fmt.Sprintf("%v", v))
    }
    if v, exists := customIntelligenceFeed["https-sha256-fingerprint"]; exists {
        d.Set("https_sha256_fingerprint", fmt.Sprintf("%v", v))
    }
    if v, exists := customIntelligenceFeed["custom-csv-settings"]; exists {
        if raw, ok := v.(map[string]interface{}); ok {
            csvMap := map[string]interface{}{
                "csv_delimiter":              fmt.Sprintf("%v", raw["csv-delimiter"]),
                "csv_lines_to_be_skipped":   fmt.Sprintf("%v", raw["csv-lines-to-be-skipped"]),
                "csv_observable_name":        fmt.Sprintf("%v", raw["csv-observable-name"]),
                "csv_observable_value":       func() int { if f, ok := raw["csv-observable-value"].(float64); ok { return int(f) }; return 0 }(),
                "csv_observable_type":        fmt.Sprintf("%v", raw["csv-observable-type"]),
                "csv_observable_description": fmt.Sprintf("%v", raw["csv-observable-description"]),
                "csv_observable_confidence":  fmt.Sprintf("%v", raw["csv-observable-confidence"]),
                "csv_observable_severity":    fmt.Sprintf("%v", raw["csv-observable-severity"]),
                "csv_observable_product":     fmt.Sprintf("%v", raw["csv-observable-product"]),
            }
            d.Set("custom_csv_settings", []interface{}{csvMap})
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaCustomIntelligenceFeed(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol"); ok {
        payload["protocol"] = v.(string)
    }

    if v, ok := d.GetOk("url"); ok {
        payload["url"] = v.(string)
    }

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("action"); ok {
        payload["action"] = v.(string)
    }

    if v, ok := d.GetOk("account_name"); ok {
        payload["account-name"] = v.(string)
    }

    if v, ok := d.GetOk("account_password"); ok {
        payload["account-password"] = v.(string)
    }

    if v := d.Get("custom_csv_settings"); len(v.([]interface{})) > 0 {
        _ = v
        customcsvsettingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("custom_csv_settings.0.csv_delimiter"); ok {
            customcsvsettingsMap["csv-delimiter"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_lines_to_be_skipped"); ok {
            customcsvsettingsMap["csv-lines-to-be-skipped"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_name"); ok {
            customcsvsettingsMap["csv-observable-name"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_value"); ok {
            customcsvsettingsMap["csv-observable-value"] = v.(int)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_type"); ok {
            customcsvsettingsMap["csv-observable-type"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_description"); ok {
            customcsvsettingsMap["csv-observable-description"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_confidence"); ok {
            customcsvsettingsMap["csv-observable-confidence"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_severity"); ok {
            customcsvsettingsMap["csv-observable-severity"] = v.(string)
        }
        if v, ok := d.GetOk("custom_csv_settings.0.csv_observable_product"); ok {
            customcsvsettingsMap["csv-observable-product"] = v.(string)
        }
        if len(customcsvsettingsMap) > 0 {
            payload["custom-csv-settings"] = customcsvsettingsMap
        }
    }

    if v, ok := d.GetOk("format"); ok {
        payload["format"] = v.(string)
    }

    if v, ok := d.GetOk("https_sha256_fingerprint"); ok {
        payload["https-sha256-fingerprint"] = v.(string)
    }

    setCustomIntelligenceFeedRes, err := client.ApiCallSimple("set-custom-intelligence-feed", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setCustomIntelligenceFeedRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setCustomIntelligenceFeedRes.Success {
            errMsg = setCustomIntelligenceFeedRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setCustomIntelligenceFeedRes.GetData()
        }

        debugLogOperation(
            "custom-intelligence-feed",        // resource type
            "update",                       // operation
            "set-custom-intelligence-feed",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set custom-intelligence-feed: %v", err)
    }
    if !setCustomIntelligenceFeedRes.Success {
        return fmt.Errorf(setCustomIntelligenceFeedRes.ErrorMsg)
    }

    return readGaiaCustomIntelligenceFeed(d, m)
}

func deleteGaiaCustomIntelligenceFeed(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteCustomIntelligenceFeedRes, err := client.ApiCallSimple("delete-custom-intelligence-feed", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteCustomIntelligenceFeedRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteCustomIntelligenceFeedRes.Success {
            errMsg = deleteCustomIntelligenceFeedRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteCustomIntelligenceFeedRes.GetData()
        }

        debugLogOperation(
            "custom-intelligence-feed",        // resource type
            "delete",                       // operation
            "delete-custom-intelligence-feed",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete custom-intelligence-feed: %v", err)
    }
    if !deleteCustomIntelligenceFeedRes.Success {
        return fmt.Errorf(deleteCustomIntelligenceFeedRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

