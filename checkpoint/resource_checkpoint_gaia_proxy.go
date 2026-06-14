package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaProxy() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaProxy,
        Read:   readGaiaProxy,
        Update: updateGaiaProxy,
        Delete: deleteGaiaProxy,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "address": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `N/A`,
            },
            "port": {
                Type:        schema.TypeInt,
                Optional:    true,
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

func createGaiaProxy(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("port"); ok {
        payload["port"] = v.(int)
    }

    log.Println("Create Proxy - Map = ", payload)

    addProxyRes, err := client.ApiCallSimple("set-proxy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addProxyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addProxyRes.Success {
            errMsg = addProxyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addProxyRes.GetData()
        }

        debugLogOperation(
            "proxy",        // resource type
            "create",                       // operation
            "set-proxy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add proxy: %v", err)
    }
    if !addProxyRes.Success {
        if addProxyRes.ErrorMsg != "" {
            return fmt.Errorf(addProxyRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("proxy-" + acctest.RandString(10)))
    return readGaiaProxy(d, m)
}

func readGaiaProxy(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showProxyRes, err := client.ApiCallSimple("show-proxy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showProxyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showProxyRes.Success {
            errMsg = showProxyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showProxyRes.GetData()
        }

        debugLogOperation(
            "proxy",        // resource type
            "read",                       // operation
            "show-proxy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show proxy: %v", err)
    }
    if !showProxyRes.Success {
        if data := showProxyRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showProxyRes.ErrorMsg)
    }

    proxy := showProxyRes.GetData()

    log.Println("Read Proxy - Show JSON = ", proxy)

    if v, exists := proxy["address"]; exists {
        d.Set("address", fmt.Sprintf("%v", v))
    }
    if v, exists := proxy["port"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("port", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("port", _n)
            }
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaProxy(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("port"); ok {
        payload["port"] = v.(int)
    }

    setProxyRes, err := client.ApiCallSimple("set-proxy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setProxyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setProxyRes.Success {
            errMsg = setProxyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setProxyRes.GetData()
        }

        debugLogOperation(
            "proxy",        // resource type
            "update",                       // operation
            "set-proxy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set proxy: %v", err)
    }
    if !setProxyRes.Success {
        return fmt.Errorf(setProxyRes.ErrorMsg)
    }

    return readGaiaProxy(d, m)
}

func deleteGaiaProxy(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    deleteProxyRes, err := client.ApiCallSimple("delete-proxy", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteProxyRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteProxyRes.Success {
            errMsg = deleteProxyRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteProxyRes.GetData()
        }

        debugLogOperation(
            "proxy",        // resource type
            "delete",                       // operation
            "delete-proxy",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete proxy: %v", err)
    }
    if !deleteProxyRes.Success {
        return fmt.Errorf(deleteProxyRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

