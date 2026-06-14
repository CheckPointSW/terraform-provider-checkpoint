package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaMessageOfTheDay() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaMessageOfTheDay,
        Read:   readGaiaMessageOfTheDay,
        Update: updateGaiaMessageOfTheDay,
        Delete: deleteGaiaMessageOfTheDay,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "message": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Message of the day for web, ssh and serial login. Empty string returns to default`,
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Message of the day enabled (true/false)`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaMessageOfTheDay(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("message"); ok {
        payload["message"] = v.(string)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    log.Println("Create MessageOfTheDay - Map = ", payload)

    addMessageOfTheDayRes, err := client.ApiCallSimple("set-message-of-the-day", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addMessageOfTheDayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addMessageOfTheDayRes.Success {
            errMsg = addMessageOfTheDayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addMessageOfTheDayRes.GetData()
        }

        debugLogOperation(
            "message-of-the-day",        // resource type
            "create",                       // operation
            "set-message-of-the-day",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add message-of-the-day: %v", err)
    }
    if !addMessageOfTheDayRes.Success {
        if addMessageOfTheDayRes.ErrorMsg != "" {
            return fmt.Errorf(addMessageOfTheDayRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("message-of-the-day-" + acctest.RandString(10)))
    return readGaiaMessageOfTheDay(d, m)
}

func readGaiaMessageOfTheDay(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showMessageOfTheDayRes, err := client.ApiCallSimple("show-message-of-the-day", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showMessageOfTheDayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showMessageOfTheDayRes.Success {
            errMsg = showMessageOfTheDayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showMessageOfTheDayRes.GetData()
        }

        debugLogOperation(
            "message-of-the-day",        // resource type
            "read",                       // operation
            "show-message-of-the-day",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show message-of-the-day: %v", err)
    }
    if !showMessageOfTheDayRes.Success {
        if data := showMessageOfTheDayRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showMessageOfTheDayRes.ErrorMsg)
    }

    messageOfTheDay := showMessageOfTheDayRes.GetData()

    log.Println("Read MessageOfTheDay - Show JSON = ", messageOfTheDay)

    if v, exists := messageOfTheDay["message"]; exists {
        d.Set("message", fmt.Sprintf("%v", v))
    }
    if v, exists := messageOfTheDay["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := messageOfTheDay["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaMessageOfTheDay(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("message"); ok {
        payload["message"] = v.(string)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    setMessageOfTheDayRes, err := client.ApiCallSimple("set-message-of-the-day", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setMessageOfTheDayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setMessageOfTheDayRes.Success {
            errMsg = setMessageOfTheDayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setMessageOfTheDayRes.GetData()
        }

        debugLogOperation(
            "message-of-the-day",        // resource type
            "update",                       // operation
            "set-message-of-the-day",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set message-of-the-day: %v", err)
    }
    if !setMessageOfTheDayRes.Success {
        return fmt.Errorf(setMessageOfTheDayRes.ErrorMsg)
    }

    return readGaiaMessageOfTheDay(d, m)
}

func deleteGaiaMessageOfTheDay(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    