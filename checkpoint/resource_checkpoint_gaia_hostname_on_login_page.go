package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaHostnameOnLoginPage() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaHostnameOnLoginPage,
        Read:   readGaiaHostnameOnLoginPage,
        Update: updateGaiaHostnameOnLoginPage,
        Delete: deleteGaiaHostnameOnLoginPage,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Hostname on Gaia Portal login page enabled (true/false)`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaHostnameOnLoginPage(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    log.Println("Create HostnameOnLoginPage - Map = ", payload)

    addHostnameOnLoginPageRes, err := client.ApiCallSimple("set-hostname-on-login-page", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addHostnameOnLoginPageRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addHostnameOnLoginPageRes.Success {
            errMsg = addHostnameOnLoginPageRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addHostnameOnLoginPageRes.GetData()
        }

        debugLogOperation(
            "hostname-on-login-page",        // resource type
            "create",                       // operation
            "set-hostname-on-login-page",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add hostname-on-login-page: %v", err)
    }
    if !addHostnameOnLoginPageRes.Success {
        if addHostnameOnLoginPageRes.ErrorMsg != "" {
            return fmt.Errorf(addHostnameOnLoginPageRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("hostname-on-login-page-" + acctest.RandString(10)))
    return readGaiaHostnameOnLoginPage(d, m)
}

func readGaiaHostnameOnLoginPage(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showHostnameOnLoginPageRes, err := client.ApiCallSimple("show-hostname-on-login-page", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showHostnameOnLoginPageRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showHostnameOnLoginPageRes.Success {
            errMsg = showHostnameOnLoginPageRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showHostnameOnLoginPageRes.GetData()
        }

        debugLogOperation(
            "hostname-on-login-page",        // resource type
            "read",                       // operation
            "show-hostname-on-login-page",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show hostname-on-login-page: %v", err)
    }
    if !showHostnameOnLoginPageRes.Success {
        if data := showHostnameOnLoginPageRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showHostnameOnLoginPageRes.ErrorMsg)
    }

    hostnameOnLoginPage := showHostnameOnLoginPageRes.GetData()

    log.Println("Read HostnameOnLoginPage - Show JSON = ", hostnameOnLoginPage)

    if v, exists := hostnameOnLoginPage["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaHostnameOnLoginPage(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    setHostnameOnLoginPageRes, err := client.ApiCallSimple("set-hostname-on-login-page", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setHostnameOnLoginPageRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setHostnameOnLoginPageRes.Success {
            errMsg = setHostnameOnLoginPageRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setHostnameOnLoginPageRes.GetData()
        }

        debugLogOperation(
            "hostname-on-login-page",        // resource type
            "update",                       // operation
            "set-hostname-on-login-page",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set hostname-on-login-page: %v", err)
    }
    if !setHostnameOnLoginPageRes.Success {
        return fmt.Errorf(setHostnameOnLoginPageRes.ErrorMsg)
    }

    return readGaiaHostnameOnLoginPage(d, m)
}

func deleteGaiaHostnameOnLoginPage(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    