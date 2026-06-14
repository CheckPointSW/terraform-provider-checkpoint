package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaBanner() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaBanner,
        Read:   readGaiaBanner,
        Update: updateGaiaBanner,
        Delete: deleteGaiaBanner,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "message": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Banner message for the web, ssh and serial login. Empty string returns to default`,
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Banner message enabled (true/false)`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaBanner(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("message"); ok {
        payload["message"] = v.(string)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    log.Println("Create Banner - Map = ", payload)

    addBannerRes, err := client.ApiCallSimple("set-banner", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addBannerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addBannerRes.Success {
            errMsg = addBannerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addBannerRes.GetData()
        }

        debugLogOperation(
            "banner",        // resource type
            "create",                       // operation
            "set-banner",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add banner: %v", err)
    }
    if !addBannerRes.Success {
        if addBannerRes.ErrorMsg != "" {
            return fmt.Errorf(addBannerRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("banner-" + acctest.RandString(10)))
    return readGaiaBanner(d, m)
}

func readGaiaBanner(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showBannerRes, err := client.ApiCallSimple("show-banner", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showBannerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showBannerRes.Success {
            errMsg = showBannerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showBannerRes.GetData()
        }

        debugLogOperation(
            "banner",        // resource type
            "read",                       // operation
            "show-banner",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show banner: %v", err)
    }
    if !showBannerRes.Success {
        if data := showBannerRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showBannerRes.ErrorMsg)
    }

    banner := showBannerRes.GetData()

    log.Println("Read Banner - Show JSON = ", banner)

    if v, exists := banner["message"]; exists {
        d.Set("message", fmt.Sprintf("%v", v))
    }
    if v, exists := banner["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := banner["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaBanner(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("message"); ok {
        payload["message"] = v.(string)
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    setBannerRes, err := client.ApiCallSimple("set-banner", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setBannerRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setBannerRes.Success {
            errMsg = setBannerRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setBannerRes.GetData()
        }

        debugLogOperation(
            "banner",        // resource type
            "update",                       // operation
            "set-banner",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set banner: %v", err)
    }
    if !setBannerRes.Success {
        return fmt.Errorf(setBannerRes.ErrorMsg)
    }

    return readGaiaBanner(d, m)
}

func deleteGaiaBanner(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    