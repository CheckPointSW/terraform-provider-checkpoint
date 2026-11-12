package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementUpdateCloudLicense() *schema.Resource {
	return &schema.Resource{
		Create: createManagementUpdateCloudLicense,
		Read:   readManagementUpdateCloudLicense,
		Delete: deleteManagementUpdateCloudLicense,
		Schema: map[string]*schema.Schema{
			"license": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The updated license string received from the User Center - without 'cplic put'.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "update-cloud-license task UID. Use show-task command to check the progress of the task.",
			},
		},
	}
}

func createManagementUpdateCloudLicense(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("license"); ok {
		payload["license"] = v.(string)
	}

	UpdateCloudLicenseRes, err := client.ApiCall("update-cloud-license", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !UpdateCloudLicenseRes.Success {
		return fmt.Errorf(UpdateCloudLicenseRes.ErrorMsg)
	}

	d.SetId("update-cloud-license-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(UpdateCloudLicenseRes.GetData()))
	return nil
}

func readManagementUpdateCloudLicense(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementUpdateCloudLicense(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
