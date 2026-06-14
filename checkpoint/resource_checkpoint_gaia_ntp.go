package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaNtp() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaNtp,
        Read:   readGaiaNtp,
        Update: updateGaiaNtp,
        Delete: deleteGaiaNtp,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `NTP status`,
            },
            "servers": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `Add, set or remove NTP server/pool`,
                Set: func(v interface{}) int { return schema.HashString(v.(map[string]interface{})["address"].(string)) },
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Address type. Should be server or pool (a dynamic collection of servers). Relevant only from R82 (V1.8). primary and secondary options are to support backward compatibility`,
                        },
                        "version": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "preferred": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Preferred address. Specify a particular server as preferred above others of similar statistical quality`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaNtp(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v := d.Get("servers"); len(v.(*schema.Set).List()) > 0 {
        payload["servers"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("preferred"); ok {
        payload["preferred"] = v.(string)
    }

    log.Println("Create Ntp - Map = ", payload)

    addNtpRes, err := client.ApiCallSimple("set-ntp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addNtpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addNtpRes.Success {
            errMsg = addNtpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addNtpRes.GetData()
        }

        debugLogOperation(
            "ntp",        // resource type
            "create",                       // operation
            "set-ntp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add ntp: %v", err)
    }
    if !addNtpRes.Success {
        if addNtpRes.ErrorMsg != "" {
            return fmt.Errorf(addNtpRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("ntp-" + acctest.RandString(10)))
    return readGaiaNtp(d, m)
}

func readGaiaNtp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showNtpRes, err := client.ApiCallSimple("show-ntp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showNtpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showNtpRes.Success {
            errMsg = showNtpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showNtpRes.GetData()
        }

        debugLogOperation(
            "ntp",        // resource type
            "read",                       // operation
            "show-ntp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show ntp: %v", err)
    }
    if !showNtpRes.Success {
        if data := showNtpRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showNtpRes.ErrorMsg)
    }

    ntp := showNtpRes.GetData()

    log.Println("Read Ntp - Show JSON = ", ntp)

    if v, exists := ntp["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := ntp["servers"]; exists {
        d.Set("servers", v.([]interface{}))
    }
    if v, exists := ntp["current"]; exists {
        d.Set("current", fmt.Sprintf("%v", v))
    }
    if v, exists := ntp["preferred"]; exists {
        d.Set("preferred", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaNtp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v := d.Get("servers"); len(v.(*schema.Set).List()) > 0 {
        payload["servers"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("preferred"); ok {
        payload["preferred"] = v.(string)
    }

    setNtpRes, err := client.ApiCallSimple("set-ntp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setNtpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setNtpRes.Success {
            errMsg = setNtpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setNtpRes.GetData()
        }

        debugLogOperation(
            "ntp",        // resource type
            "update",                       // operation
            "set-ntp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set ntp: %v", err)
    }
    if !setNtpRes.Success {
        return fmt.Errorf(setNtpRes.ErrorMsg)
    }

    return readGaiaNtp(d, m)
}

func deleteGaiaNtp(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    