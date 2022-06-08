package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementUpdateProvisionedSatellites() *schema.Resource {
	return &schema.Resource{
		Create: createManagementUpdateProvisionedSatellites,
		Read:   readManagementUpdateProvisionedSatellites,
		Delete: deleteManagementUpdateProvisionedSatellites,
		Schema: map[string]*schema.Schema{
			"vpn_center_gateways": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "On what targets to execute this command. Targets may be identified by their name, or object unique identifier. The targets should be a corporate gateways.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementUpdateProvisionedSatellites(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("vpn_center_gateways"); ok {
		payload["vpn-center-gateways"] = v.(*schema.Set).List()
	}

	UpdateProvisionedSatellitesRes, _ := client.ApiCall("update-provisioned-satellites", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !UpdateProvisionedSatellitesRes.Success {
		return fmt.Errorf(UpdateProvisionedSatellitesRes.ErrorMsg)
	}

	d.SetId("update-provisioned-satellites" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(UpdateProvisionedSatellitesRes.GetData()))
	return readManagementUpdateProvisionedSatellites(d, m)
}

func readManagementUpdateProvisionedSatellites(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementUpdateProvisionedSatellites(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
