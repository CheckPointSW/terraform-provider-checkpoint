package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaAuthenticationOrder() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaAuthenticationOrder,
        Read:   readGaiaAuthenticationOrder,
        Update: updateGaiaAuthenticationOrder,
        Delete: deleteGaiaAuthenticationOrder,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "radius": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Server type`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "priority": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Authentication priority`,
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Server state`,
                        },
                    },
                },
            },
            "tacacs": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Server type`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "priority": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Authentication priority`,
                        },
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Server state`,
                        },
                    },
                },
            },
            "local": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Server type`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "priority": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Authentication priority`,
                        },
                    },
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaAuthenticationOrder(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("radius"); len(v.([]interface{})) > 0 {
        _ = v
        radiusMap := make(map[string]interface{})
        if v, ok := d.GetOk("radius.0.priority"); ok {
            radiusMap["priority"] = v.(int)
        }
        if v, ok := d.GetOkExists("radius.0.enabled"); ok && v.(bool) {
            radiusMap["enabled"] = v.(bool)
        }
        if len(radiusMap) > 0 {
            payload["radius"] = radiusMap
        }
    }

    if v := d.Get("tacacs"); len(v.([]interface{})) > 0 {
        _ = v
        tacacsMap := make(map[string]interface{})
        if v, ok := d.GetOk("tacacs.0.priority"); ok {
            tacacsMap["priority"] = v.(int)
        }
        if v, ok := d.GetOkExists("tacacs.0.enabled"); ok && v.(bool) {
            tacacsMap["enabled"] = v.(bool)
        }
        if len(tacacsMap) > 0 {
            payload["tacacs"] = tacacsMap
        }
    }

    if v := d.Get("local"); len(v.([]interface{})) > 0 {
        _ = v
        localMap := make(map[string]interface{})
        if v, ok := d.GetOk("local.0.priority"); ok {
            localMap["priority"] = v.(int)
        }
        if len(localMap) > 0 {
            payload["local"] = localMap
        }
    }

    log.Println("Create AuthenticationOrder - Map = ", payload)

    addAuthenticationOrderRes, err := client.ApiCallSimple("set-authentication-order", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addAuthenticationOrderRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addAuthenticationOrderRes.Success {
            errMsg = addAuthenticationOrderRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addAuthenticationOrderRes.GetData()
        }

        debugLogOperation(
            "authentication-order",        // resource type
            "create",                       // operation
            "set-authentication-order",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add authentication-order: %v", err)
    }
    if !addAuthenticationOrderRes.Success {
        if addAuthenticationOrderRes.ErrorMsg != "" {
            return fmt.Errorf(addAuthenticationOrderRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("authentication-order-" + acctest.RandString(10)))
    return readGaiaAuthenticationOrder(d, m)
}

func readGaiaAuthenticationOrder(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showAuthenticationOrderRes, err := client.ApiCallSimple("show-authentication-order", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showAuthenticationOrderRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showAuthenticationOrderRes.Success {
            errMsg = showAuthenticationOrderRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showAuthenticationOrderRes.GetData()
        }

        debugLogOperation(
            "authentication-order",        // resource type
            "read",                       // operation
            "show-authentication-order",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show authentication-order: %v", err)
    }
    if !showAuthenticationOrderRes.Success {
        if data := showAuthenticationOrderRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showAuthenticationOrderRes.ErrorMsg)
    }

    authenticationOrder := showAuthenticationOrderRes.GetData()

    log.Println("Read AuthenticationOrder - Show JSON = ", authenticationOrder)

    if v, exists := authenticationOrder["radius"]; exists {
        d.Set("radius", v)
    }
    if v, exists := authenticationOrder["tacacs"]; exists {
        d.Set("tacacs", v)
    }
    if v, exists := authenticationOrder["local"]; exists {
        d.Set("local", v)
    }
    if v, exists := authenticationOrder["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaAuthenticationOrder(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("radius"); len(v.([]interface{})) > 0 {
        _ = v
        radiusMap := make(map[string]interface{})
        if v, ok := d.GetOk("radius.0.priority"); ok {
            radiusMap["priority"] = v.(int)
        }
        if v, ok := d.GetOkExists("radius.0.enabled"); ok && v.(bool) {
            radiusMap["enabled"] = v.(bool)
        }
        if len(radiusMap) > 0 {
            payload["radius"] = radiusMap
        }
    }

    if v := d.Get("tacacs"); len(v.([]interface{})) > 0 {
        _ = v
        tacacsMap := make(map[string]interface{})
        if v, ok := d.GetOk("tacacs.0.priority"); ok {
            tacacsMap["priority"] = v.(int)
        }
        if v, ok := d.GetOkExists("tacacs.0.enabled"); ok && v.(bool) {
            tacacsMap["enabled"] = v.(bool)
        }
        if len(tacacsMap) > 0 {
            payload["tacacs"] = tacacsMap
        }
    }

    if v := d.Get("local"); len(v.([]interface{})) > 0 {
        _ = v
        localMap := make(map[string]interface{})
        if v, ok := d.GetOk("local.0.priority"); ok {
            localMap["priority"] = v.(int)
        }
        if len(localMap) > 0 {
            payload["local"] = localMap
        }
    }

    setAuthenticationOrderRes, err := client.ApiCallSimple("set-authentication-order", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setAuthenticationOrderRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setAuthenticationOrderRes.Success {
            errMsg = setAuthenticationOrderRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setAuthenticationOrderRes.GetData()
        }

        debugLogOperation(
            "authentication-order",        // resource type
            "update",                       // operation
            "set-authentication-order",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set authentication-order: %v", err)
    }
    if !setAuthenticationOrderRes.Success {
        return fmt.Errorf(setAuthenticationOrderRes.ErrorMsg)
    }

    return readGaiaAuthenticationOrder(d, m)
}

func deleteGaiaAuthenticationOrder(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    