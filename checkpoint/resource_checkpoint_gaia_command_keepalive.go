package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaKeepalive() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaKeepalive,
        Read:   readGaiaKeepalive,
        Delete: deleteGaiaKeepalive,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "message": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaKeepalive(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    log.Println("Execute keepalive - Payload = ", payload)

    GaiaKeepaliveRes, err := client.ApiCallSimple("keepalive", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaKeepaliveRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaKeepaliveRes.Success {
            errMsg = GaiaKeepaliveRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaKeepaliveRes.GetData()
        }

        debugLogOperation(
            "keepalive",        // resource type
            "command",                       // operation
            "keepalive",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute keepalive: %v", err)
    }
    if !GaiaKeepaliveRes.Success {
        if GaiaKeepaliveRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaKeepaliveRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaKeepaliveRes.GetData()
    if v, exists := _respData["message"]; exists {
        d.Set("message", toString(v))
    }


    d.SetId(fmt.Sprintf("keepalive-" + acctest.RandString(10)))
    return nil
}

func readGaiaKeepalive(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    log.Println("Execute keepalive - Payload = ", payload)

    GaiaKeepaliveRes, err := client.ApiCallSimple("keepalive", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaKeepaliveRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaKeepaliveRes.Success {
            errMsg = GaiaKeepaliveRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaKeepaliveRes.GetData()
        }

        debugLogOperation(
            "keepalive",        // resource type
            "read",                       // operation
            "keepalive",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute keepalive: %v", err)
    }
    if !GaiaKeepaliveRes.Success {
        if GaiaKeepaliveRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaKeepaliveRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaKeepaliveRes.GetData()
    if v, exists := _respData["message"]; exists {
        d.Set("message", toString(v))
    }
    return nil
}

func deleteGaiaKeepalive(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

