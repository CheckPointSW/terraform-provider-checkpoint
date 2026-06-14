package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaDiscardMaestroSecurityGroupsChanges() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaDiscardMaestroSecurityGroupsChanges,
        Read:   readGaiaDiscardMaestroSecurityGroupsChanges,
        Delete: deleteGaiaDiscardMaestroSecurityGroupsChanges,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "virtual_system_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaDiscardMaestroSecurityGroupsChanges(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    log.Println("Execute discard-maestro-security-groups-changes - Payload = ", payload)

    GaiaDiscardMaestroSecurityGroupsChangesRes, err := client.ApiCallSimple("discard-maestro-security-groups-changes", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaDiscardMaestroSecurityGroupsChangesRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaDiscardMaestroSecurityGroupsChangesRes.Success {
            errMsg = GaiaDiscardMaestroSecurityGroupsChangesRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaDiscardMaestroSecurityGroupsChangesRes.GetData()
        }

        debugLogOperation(
            "discard-maestro-security-groups-changes",        // resource type
            "command",                       // operation
            "discard-maestro-security-groups-changes",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute discard-maestro-security-groups-changes: %v", err)
    }
    if !GaiaDiscardMaestroSecurityGroupsChangesRes.Success {
        if GaiaDiscardMaestroSecurityGroupsChangesRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaDiscardMaestroSecurityGroupsChangesRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaDiscardMaestroSecurityGroupsChangesRes.GetData()
    if v, exists := _respData["virtual-system-id"]; exists {
        d.Set("virtual_system_id", toString(v))
    }


    d.SetId(fmt.Sprintf("discard-maestro-security-groups-changes-" + acctest.RandString(10)))
    return nil
}

func readGaiaDiscardMaestroSecurityGroupsChanges(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaDiscardMaestroSecurityGroupsChanges(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

