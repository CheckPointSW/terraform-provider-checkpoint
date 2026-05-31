package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaRouterId() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaRouterId,
        Read:   readGaiaRouterId,
        Update: updateGaiaRouterId,
        Delete: deleteGaiaRouterId,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "router_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configures the Router ID used by BGP and OSPF. It is usually the IPv4 address of one of the local interfaces, and should uniquely identify the router within the local Autonomous System. It is generally recommended that a non-127.0.0.1 loopback address be used.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaRouterId(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("router_id"); ok {
        payload["router-id"] = v.(string)
    }

    log.Println("Create RouterId - Map = ", payload)

    addRouterIdRes, err := client.ApiCallSimple("set-router-id", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addRouterIdRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addRouterIdRes.Success {
            errMsg = addRouterIdRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addRouterIdRes.GetData()
        }

        debugLogOperation(
            "router-id",        // resource type
            "create",                       // operation
            "set-router-id",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add router-id: %v", err)
    }
    if !addRouterIdRes.Success {
        if addRouterIdRes.ErrorMsg != "" {
            return fmt.Errorf(addRouterIdRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("router-id-" + acctest.RandString(10)))
    return readGaiaRouterId(d, m)
}

func readGaiaRouterId(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showRouterIdRes, err := client.ApiCallSimple("show-router-id", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showRouterIdRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showRouterIdRes.Success {
            errMsg = showRouterIdRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showRouterIdRes.GetData()
        }

        debugLogOperation(
            "router-id",        // resource type
            "read",                       // operation
            "show-router-id",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show router-id: %v", err)
    }
    if !showRouterIdRes.Success {
        if data := showRouterIdRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showRouterIdRes.ErrorMsg)
    }

    routerId := showRouterIdRes.GetData()

    log.Println("Read RouterId - Show JSON = ", routerId)

    if v, exists := routerId["router-id"]; exists {
        if configuredVal, ok := d.GetOk("router_id"); ok && configuredVal.(string) == "default" {
            d.Set("router_id", "default")
        } else {
            d.Set("router_id", fmt.Sprintf("%v", v))
        }
    }
    if v, exists := routerId["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaRouterId(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("router_id"); ok {
        payload["router-id"] = v.(string)
    }

    setRouterIdRes, err := client.ApiCallSimple("set-router-id", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setRouterIdRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setRouterIdRes.Success {
            errMsg = setRouterIdRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setRouterIdRes.GetData()
        }

        debugLogOperation(
            "router-id",        // resource type
            "update",                       // operation
            "set-router-id",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set router-id: %v", err)
    }
    if !setRouterIdRes.Success {
        return fmt.Errorf(setRouterIdRes.ErrorMsg)
    }

    return readGaiaRouterId(d, m)
}

func deleteGaiaRouterId(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    