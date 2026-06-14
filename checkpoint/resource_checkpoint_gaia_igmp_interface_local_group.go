package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"

)
func resourceGaiaIgmpInterfaceLocalGroup() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaIgmpInterfaceLocalGroup,
        Read:   readGaiaIgmpInterfaceLocalGroup,
        Delete: deleteGaiaIgmpInterfaceLocalGroup,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Enable debug logging for this resource.",
            },
            "interface": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `The name of the IGMP interface`,
            },
            "local_group": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `The locally configured group address that this IGMP interface receives multicast data for`,
            },
        },
    }
}

func createGaiaIgmpInterfaceLocalGroup(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("local_group"); ok {
        payload["local-group"] = v.(string)
    }

    log.Println("Create IgmpInterfaceLocalGroup - Map = ", payload)

    addIgmpInterfaceLocalGroupRes, err := client.ApiCallSimple("add-igmp-interface-local-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addIgmpInterfaceLocalGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addIgmpInterfaceLocalGroupRes.Success {
            errMsg = addIgmpInterfaceLocalGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addIgmpInterfaceLocalGroupRes.GetData()
        }

        debugLogOperation(
            "igmp-interface-local-group",        // resource type
            "create",                       // operation
            "add-igmp-interface-local-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add igmp-interface-local-group: %v", err)
    }
    if !addIgmpInterfaceLocalGroupRes.Success {
        if addIgmpInterfaceLocalGroupRes.ErrorMsg != "" {
            return fmt.Errorf(addIgmpInterfaceLocalGroupRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("igmp-interface-local-group-" + acctest.RandString(10)))
    return readGaiaIgmpInterfaceLocalGroup(d, m)
}

func readGaiaIgmpInterfaceLocalGroup(d *schema.ResourceData, m interface{}) error {


        // No API call - just preserve the ID to indicate resource still exists
        // This assumes the resource exists as long as it's in state
        return nil
}
func deleteGaiaIgmpInterfaceLocalGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("local_group"); ok {
        payload["local-group"] = v.(string)
    }

    deleteIgmpInterfaceLocalGroupRes, err := client.ApiCallSimple("delete-igmp-interface-local-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteIgmpInterfaceLocalGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteIgmpInterfaceLocalGroupRes.Success {
            errMsg = deleteIgmpInterfaceLocalGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteIgmpInterfaceLocalGroupRes.GetData()
        }

        debugLogOperation(
            "igmp-interface-local-group",        // resource type
            "delete",                       // operation
            "delete-igmp-interface-local-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete igmp-interface-local-group: %v", err)
    }
    if !deleteIgmpInterfaceLocalGroupRes.Success {
        return fmt.Errorf(deleteIgmpInterfaceLocalGroupRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

