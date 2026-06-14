package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaLogout() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaLogout,
        Read:   readGaiaLogout,
        Delete: deleteGaiaLogout,
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

func createGaiaLogout(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    log.Println("Execute logout - Payload = ", payload)

    GaiaLogoutRes, err := client.ApiCallSimple("logout", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaLogoutRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaLogoutRes.Success {
            errMsg = GaiaLogoutRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaLogoutRes.GetData()
        }

        debugLogOperation(
            "logout",        // resource type
            "command",                       // operation
            "logout",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute logout: %v", err)
    }
    if !GaiaLogoutRes.Success {
        if GaiaLogoutRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaLogoutRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaLogoutRes.GetData()
    if v, exists := _respData["message"]; exists {
        d.Set("message", toString(v))
    }


    d.SetId(fmt.Sprintf("logout-" + acctest.RandString(10)))
    return nil
}

func readGaiaLogout(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaLogout(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

