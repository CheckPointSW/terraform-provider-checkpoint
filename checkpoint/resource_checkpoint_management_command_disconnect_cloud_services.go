package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementDisconnectCloudServices() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDisconnectCloudServices,
		Read:   readManagementDisconnectCloudServices,
		Delete: deleteManagementDisconnectCloudServices,
		Schema: map[string]*schema.Schema{
			"force": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Disconnect the Management Server from Check Point Infinity Portal, and reset the connection locally, regardless of the result in the Infinity Portal. This flag can be used if the disconnect-cloud-services command failed. Since with this flag this command affects only the local configuration, make sure to disconnect the Management Server in the Infinity Portal as well.",
			},
		},
	}
}

func createManagementDisconnectCloudServices(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOkExists("force"); ok {
		payload["force"] = v.(bool)
	}

	disconnectCloudServicesRes, err := client.ApiCall("disconnect-cloud-services", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !disconnectCloudServicesRes.Success {
		return fmt.Errorf(disconnectCloudServicesRes.ErrorMsg)
	}

	d.SetId("disconnect-cloud-services" + acctest.RandString(5))
	return readManagementDisconnectCloudServices(d, m)
}

func readManagementDisconnectCloudServices(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementDisconnectCloudServices(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
