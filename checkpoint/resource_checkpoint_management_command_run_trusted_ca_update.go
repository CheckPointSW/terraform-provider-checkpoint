package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementRunTrustedCaUpdate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRunTrustedCaUpdate,
		Read:   readManagementRunTrustedCaUpdate,
		Delete: deleteManagementRunTrustedCaUpdate,
		Schema: map[string]*schema.Schema{
			"package_path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Path on the management server for offline Trusted CAs package update.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementRunTrustedCaUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = make(map[string]interface{})
	if v, ok := d.GetOk("package_path"); ok {
		payload["package-path"] = v.(string)
	}

	RunTrustedCaUpdateRes, _ := client.ApiCall("run-trusted-ca-update", payload, client.GetSessionID(), true, false)
	if !RunTrustedCaUpdateRes.Success {
		return fmt.Errorf(RunTrustedCaUpdateRes.ErrorMsg)
	}

	d.SetId("run-trusted-ca-update-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(RunTrustedCaUpdateRes.GetData()))
	return readManagementRunTrustedCaUpdate(d, m)
}

func readManagementRunTrustedCaUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementRunTrustedCaUpdate(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
