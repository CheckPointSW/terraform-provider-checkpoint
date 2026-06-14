package checkpoint

import (
	"fmt"
	"log"
	"math/rand"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGaiaMldInterfaceLocalGroup() *schema.Resource {
	return &schema.Resource{
		Create: createGaiaMldInterfaceLocalGroup,
		Read:   readGaiaMldInterfaceLocalGroup,
		Update: updateGaiaMldInterfaceLocalGroup,
		Delete: deleteGaiaMldInterfaceLocalGroup,
		Schema: map[string]*schema.Schema{
			"debug": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Enable debugging for this resource only.",
			},
			"interface": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The name of the MLD interface",
			},
			"local_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The locally configured group address that this MLD interface receives multicast data for",
			},
		},
	}
}

func createGaiaMldInterfaceLocalGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	ensureDebugServerFromClient(client)

	payload := make(map[string]interface{})

	if v, ok := d.GetOk("interface"); ok {
		payload["interface"] = v.(string)
	}

	if v, ok := d.GetOk("local_group"); ok {
		payload["local-group"] = v.(string)
	}

	log.Println("Create MldInterfaceLocalGroup - Map = ", payload)

	addMldInterfaceLocalGroupRes, err := client.ApiCallSimple("add-mld-interface-local-group", payload)
	// DEBUG: generic logger
	if resourceDebugEnabled(d) {
		success := err == nil && addMldInterfaceLocalGroupRes.Success
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		} else if !addMldInterfaceLocalGroupRes.Success {
			errMsg = addMldInterfaceLocalGroupRes.ErrorMsg
		}

		var respData map[string]interface{}
		if err == nil {
			respData = addMldInterfaceLocalGroupRes.GetData()
		}

		debugLogOperation(
			"mld-interface-local-group",        // resource type
			"create",                       // operation
			"add-mld-interface-local-group",         // API call name
			payload,                        // request payload
			respData,                       // response data (if any)
			success,
			errMsg,
		)
	}
	if err != nil {
		return fmt.Errorf("Failed to add mld-interface-local-group: %v", err)
	}
	if !addMldInterfaceLocalGroupRes.Success {
		if addMldInterfaceLocalGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addMldInterfaceLocalGroupRes.ErrorMsg)
		}
		return fmt.Errorf("Unknown error occurred")
	}

	
	d.SetId(fmt.Sprintf("mld-interface-local-group.%d", rand.Intn(100000)))
	return readGaiaMldInterfaceLocalGroup(d, m)
}

func readGaiaMldInterfaceLocalGroup(d *schema.ResourceData, m interface{}) error {

        // No API call - just preserve the ID to indicate resource still exists
        // This assumes the resource exists as long as it's in state
        return nil
}

func updateGaiaMldInterfaceLocalGroup(d *schema.ResourceData, m interface{}) error {

	return readGaiaMldInterfaceLocalGroup(d, m)
}

func deleteGaiaMldInterfaceLocalGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	ensureDebugServerFromClient(client)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("interface"); ok {
		payload["interface"] = v.(string)
	}

	if v, ok := d.GetOk("local_group"); ok {
		payload["local-group"] = v.(string)
	}

	deleteMldInterfaceLocalGroupRes, err := client.ApiCallSimple("delete-mld-interface-local-group", payload)
	// DEBUG: generic logger
	if resourceDebugEnabled(d) {
		success := err == nil && deleteMldInterfaceLocalGroupRes.Success
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		} else if !deleteMldInterfaceLocalGroupRes.Success {
			errMsg = deleteMldInterfaceLocalGroupRes.ErrorMsg
		}

		var respData map[string]interface{}
		if err == nil {
			respData = deleteMldInterfaceLocalGroupRes.GetData()
		}

		debugLogOperation(
			"mld-interface-local-group",        // resource type
			"delete",                       // operation
			"delete-mld-interface-local-group",         // API call name
			payload,                        // request payload
			respData,                       // response data (if any)
			success,
			errMsg,
		)
	}
	if err != nil {
		return fmt.Errorf("Failed to delete mld-interface-local-group: %v", err)
	}
	if !deleteMldInterfaceLocalGroupRes.Success {
		return fmt.Errorf(deleteMldInterfaceLocalGroupRes.ErrorMsg)
	}

	d.SetId("")
	return nil
}
