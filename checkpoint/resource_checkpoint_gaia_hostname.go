package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaHostname() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaHostname,
        Read:   readGaiaHostname,
        Update: updateGaiaHostname,
        Delete: deleteGaiaHostname,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Hostname can be a combination of letters and numbers, it cannot be in IP format or start/end with characters such as '.' And '-' `,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaHostname(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    log.Println("Create Hostname - Map = ", payload)

    addHostnameRes, err := client.ApiCallSimple("set-hostname", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addHostnameRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addHostnameRes.Success {
            errMsg = addHostnameRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addHostnameRes.GetData()
        }

        debugLogOperation(
            "hostname",        // resource type
            "create",                       // operation
            "set-hostname",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add hostname: %v", err)
    }
    if !addHostnameRes.Success {
        if addHostnameRes.ErrorMsg != "" {
            return fmt.Errorf(addHostnameRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("hostname-" + acctest.RandString(10)))
    return readGaiaHostname(d, m)
}

func readGaiaHostname(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showHostnameRes, err := client.ApiCallSimple("show-hostname", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showHostnameRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showHostnameRes.Success {
            errMsg = showHostnameRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showHostnameRes.GetData()
        }

        debugLogOperation(
            "hostname",        // resource type
            "read",                       // operation
            "show-hostname",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show hostname: %v", err)
    }
    if !showHostnameRes.Success {
        if data := showHostnameRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showHostnameRes.ErrorMsg)
    }

    hostname := showHostnameRes.GetData()

    log.Println("Read Hostname - Show JSON = ", hostname)

    if v, exists := hostname["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaHostname(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    setHostnameRes, err := client.ApiCallSimple("set-hostname", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setHostnameRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setHostnameRes.Success {
            errMsg = setHostnameRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setHostnameRes.GetData()
        }

        debugLogOperation(
            "hostname",        // resource type
            "update",                       // operation
            "set-hostname",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set hostname: %v", err)
    }
    if !setHostnameRes.Success {
        return fmt.Errorf(setHostnameRes.ErrorMsg)
    }

    return readGaiaHostname(d, m)
}

func deleteGaiaHostname(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    