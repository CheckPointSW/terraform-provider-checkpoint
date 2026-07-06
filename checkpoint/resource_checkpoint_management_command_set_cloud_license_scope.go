package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementSetCloudLicenseScope() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetCloudLicenseScope,
		Read:   readManagementSetCloudLicenseScope,
		Delete: deleteManagementSetCloudLicenseScope,
		Schema: map[string]*schema.Schema{
			"mode": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Set cloud license scope mode.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "set-cloud-license-scope task UID. Use show-task command to check the progress of the task.",
			},
		},
	}
}

func createManagementSetCloudLicenseScope(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("mode"); ok {
		payload["mode"] = v.(string)
	}

	SetCloudLicenseScopeRes, err := client.ApiCall("set-cloud-license-scope", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetCloudLicenseScopeRes.Success {
		return fmt.Errorf(SetCloudLicenseScopeRes.ErrorMsg)
	}

	d.SetId("set-cloud-license-scope-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(SetCloudLicenseScopeRes.GetData()))
	return nil
}

func readManagementSetCloudLicenseScope(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetCloudLicenseScope(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
