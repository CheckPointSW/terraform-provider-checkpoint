package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaIpv6() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaIpv6,
        Read:   readGaiaIpv6,
        Update: updateGaiaIpv6,
        Delete: deleteGaiaIpv6,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Required:    true,
                Description: `N/A`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaIpv6(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    log.Println("Create Ipv6 - Map = ", payload)

    addIpv6Res, err := client.ApiCallSimple("set-ipv6", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addIpv6Res.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addIpv6Res.Success {
            errMsg = addIpv6Res.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addIpv6Res.GetData()
        }

        debugLogOperation(
            "ipv6",        // resource type
            "create",                       // operation
            "set-ipv6",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add ipv6: %v", err)
    }
    if !addIpv6Res.Success {
        if addIpv6Res.ErrorMsg != "" {
            return fmt.Errorf(addIpv6Res.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("ipv6-" + acctest.RandString(10)))
    return readGaiaIpv6(d, m)
}

func readGaiaIpv6(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showIpv6Res, err := client.ApiCallSimple("show-ipv6", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showIpv6Res.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showIpv6Res.Success {
            errMsg = showIpv6Res.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showIpv6Res.GetData()
        }

        debugLogOperation(
            "ipv6",        // resource type
            "read",                       // operation
            "show-ipv6",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show ipv6: %v", err)
    }
    if !showIpv6Res.Success {
        if data := showIpv6Res.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showIpv6Res.ErrorMsg)
    }

    ipv6 := showIpv6Res.GetData()

    log.Println("Read Ipv6 - Show JSON = ", ipv6)

    if v, exists := ipv6["reboot-required"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("reboot_required", b)
        } else if s, ok := v.(string); ok {
            d.Set("reboot_required", s == "true")
        }
    }
    if v, exists := ipv6["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaIpv6(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    setIpv6Res, err := client.ApiCallSimple("set-ipv6", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setIpv6Res.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setIpv6Res.Success {
            errMsg = setIpv6Res.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setIpv6Res.GetData()
        }

        debugLogOperation(
            "ipv6",        // resource type
            "update",                       // operation
            "set-ipv6",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set ipv6: %v", err)
    }
    if !setIpv6Res.Success {
        return fmt.Errorf(setIpv6Res.ErrorMsg)
    }

    return readGaiaIpv6(d, m)
}

func deleteGaiaIpv6(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    