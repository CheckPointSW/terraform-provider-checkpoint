package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementSetCloudLicenseGateway() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetCloudLicenseGateway,
		Read:   readManagementSetCloudLicenseGateway,
		Delete: deleteManagementSetCloudLicenseGateway,
		Schema: map[string]*schema.Schema{
			"gateway": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Security gateway name or UID to set.",
			},
			"enable_auto_distribution": {
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Enable or disable auto distribution of cloud licenses for the specified gateway.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Domain name or UID for the gateway. Required when running from MDS context.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "set-cloud-license-gateway task UID. Use show-task command to check the progress of the task.",
			},
		},
	}
}

func createManagementSetCloudLicenseGateway(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("gateway"); ok {
		payload["gateway"] = v.(string)
	}

	if v, ok := d.GetOkExists("enable_auto_distribution"); ok {
		payload["enable-auto-distribution"] = v.(bool)
	}

	if v, ok := d.GetOk("domain"); ok {
		payload["domain"] = v.(string)
	}

	SetCloudLicenseGatewayRes, err := client.ApiCall("set-cloud-license-gateway", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetCloudLicenseGatewayRes.Success {
		return fmt.Errorf(SetCloudLicenseGatewayRes.ErrorMsg)
	}

	d.SetId("set-cloud-license-gateway-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(SetCloudLicenseGatewayRes.GetData()))
	return nil
}

func readManagementSetCloudLicenseGateway(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetCloudLicenseGateway(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
