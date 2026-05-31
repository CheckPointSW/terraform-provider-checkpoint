package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaCustomIntelligenceInterval() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaCustomIntelligenceInterval,
        Read:   readGaiaCustomIntelligenceInterval,
        Update: updateGaiaCustomIntelligenceInterval,
        Delete: deleteGaiaCustomIntelligenceInterval,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "interval": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `Check for updates frequency`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaCustomIntelligenceInterval(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("interval"); ok {
        payload["interval"] = v.(int)
    }

    log.Println("Create CustomIntelligenceInterval - Map = ", payload)

    addCustomIntelligenceIntervalRes, err := client.ApiCallSimple("set-custom-intelligence-interval", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addCustomIntelligenceIntervalRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addCustomIntelligenceIntervalRes.Success {
            errMsg = addCustomIntelligenceIntervalRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addCustomIntelligenceIntervalRes.GetData()
        }

        debugLogOperation(
            "custom-intelligence-interval",        // resource type
            "create",                       // operation
            "set-custom-intelligence-interval",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add custom-intelligence-interval: %v", err)
    }
    if !addCustomIntelligenceIntervalRes.Success {
        if addCustomIntelligenceIntervalRes.ErrorMsg != "" {
            return fmt.Errorf(addCustomIntelligenceIntervalRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("custom-intelligence-interval-" + acctest.RandString(10)))
    return readGaiaCustomIntelligenceInterval(d, m)
}

func readGaiaCustomIntelligenceInterval(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showCustomIntelligenceIntervalRes, err := client.ApiCallSimple("show-custom-intelligence-interval", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showCustomIntelligenceIntervalRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showCustomIntelligenceIntervalRes.Success {
            errMsg = showCustomIntelligenceIntervalRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showCustomIntelligenceIntervalRes.GetData()
        }

        debugLogOperation(
            "custom-intelligence-interval",        // resource type
            "read",                       // operation
            "show-custom-intelligence-interval",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show custom-intelligence-interval: %v", err)
    }
    if !showCustomIntelligenceIntervalRes.Success {
        if data := showCustomIntelligenceIntervalRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showCustomIntelligenceIntervalRes.ErrorMsg)
    }

    customIntelligenceInterval := showCustomIntelligenceIntervalRes.GetData()

    log.Println("Read CustomIntelligenceInterval - Show JSON = ", customIntelligenceInterval)

    if v, exists := customIntelligenceInterval["interval"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("interval", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("interval", _n)
            }
        }
    }
    if v, exists := customIntelligenceInterval["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaCustomIntelligenceInterval(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("interval"); ok {
        payload["interval"] = v.(int)
    }

    setCustomIntelligenceIntervalRes, err := client.ApiCallSimple("set-custom-intelligence-interval", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setCustomIntelligenceIntervalRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setCustomIntelligenceIntervalRes.Success {
            errMsg = setCustomIntelligenceIntervalRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setCustomIntelligenceIntervalRes.GetData()
        }

        debugLogOperation(
            "custom-intelligence-interval",        // resource type
            "update",                       // operation
            "set-custom-intelligence-interval",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set custom-intelligence-interval: %v", err)
    }
    if !setCustomIntelligenceIntervalRes.Success {
        return fmt.Errorf(setCustomIntelligenceIntervalRes.ErrorMsg)
    }

    return readGaiaCustomIntelligenceInterval(d, m)
}

func deleteGaiaCustomIntelligenceInterval(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    