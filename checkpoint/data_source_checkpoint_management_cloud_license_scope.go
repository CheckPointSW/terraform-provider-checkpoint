package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementCloudLicenseScope() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCloudLicenseScopeRead,
		Schema: map[string]*schema.Schema{
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "update-cloud-license task UID. Use show-task command to check the progress of the task.",
			},
		},
	}
}

func dataSourceManagementCloudLicenseScopeRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	showCloudLicenseScopeRes, err := client.ApiCall("show-cloud-license-scope", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showCloudLicenseScopeRes.Success {
		return fmt.Errorf(showCloudLicenseScopeRes.ErrorMsg)
	}

	cloudLicenseScope := showCloudLicenseScopeRes.GetData()

	log.Println("Read CloudLicenseScope - Show JSON = ", cloudLicenseScope)

	if v := resolveTaskId(cloudLicenseScope); v != nil {
		d.SetId(v.(string))
		_ = d.Set("task_id", v)
	} else {
		d.SetId("cloud-license-scope")
	}

	return nil
}
