package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementRunIpsUpdate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRunIpsUpdate,
		Read:   readManagementRunIpsUpdate,
		Delete: deleteManagementRunIpsUpdate,
		Schema: map[string]*schema.Schema{
			"package_path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Offline update package path.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementRunIpsUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = make(map[string]interface{})

	if v, ok := d.GetOk("package_path"); ok {
		payload["package-path"] = v.(string)
	}

	runIpsUpdateRes, _ := client.ApiCall("run-ips-update", payload, client.GetSessionID(), true, false)
	if !runIpsUpdateRes.Success {
		return fmt.Errorf(runIpsUpdateRes.ErrorMsg)
	}

	d.SetId("run-ips-update-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(runIpsUpdateRes.GetData()))

	return readManagementPublish(d, m)
}

func readManagementRunIpsUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementRunIpsUpdate(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
