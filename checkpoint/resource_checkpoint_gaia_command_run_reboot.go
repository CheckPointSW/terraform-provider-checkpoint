package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaRunReboot() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaRunReboot,
        Read:   readGaiaRunReboot,
        Delete: deleteGaiaRunReboot,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "task_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaRunReboot(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    log.Println("Execute run-reboot - Payload = ", payload)

    GaiaRunRebootRes, err := client.ApiCallSimple("run-reboot", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaRunRebootRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaRunRebootRes.Success {
            errMsg = GaiaRunRebootRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaRunRebootRes.GetData()
        }

        debugLogOperation(
            "run-reboot",        // resource type
            "command",                       // operation
            "run-reboot",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute run-reboot: %v", err)
    }
    if !GaiaRunRebootRes.Success {
        if GaiaRunRebootRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaRunRebootRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaRunRebootRes.GetData()
    if v, exists := _respData["task-id"]; exists {
        d.Set("task_id", toString(v))
    }


    d.SetId(fmt.Sprintf("run-reboot-" + acctest.RandString(10)))
    return nil
}

func readGaiaRunReboot(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaRunReboot(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

