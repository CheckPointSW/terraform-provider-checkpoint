package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckpointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
)


func resourceManagementRunIpsUpdate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRunIpsUpdate,
		Read:   readManagementRunIpsUpdate,
		Delete: deleteManagementRunIpsUpdate,
		Schema: map[string]*schema.Schema{
			"package_path": {
				Type: schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: "Offline update package path.",
			},
		},
	}
}

func createManagementRunIpsUpdate(d *schema.ResourceData, m interface{}) error {
	return readManagementPublish(d, m)
}

func readManagementRunIpsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = make(map[string]interface{})

	if v, ok := d.GetOk("package_path"); ok {
		payload["package-path"] = v.(string)
	}

	runIpsUpdateRes, _ := client.ApiCall("run-ips-update", payload, client.GetSessionID(),true,false)
	if !runIpsUpdateRes.Success {
		return fmt.Errorf(runIpsUpdateRes.ErrorMsg)
	}
	// Set Schema UID = Session UID
	d.SetId(runIpsUpdateRes.GetData()["task-id"].(string))
	return nil
}

func deleteManagementRunIpsUpdate(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
