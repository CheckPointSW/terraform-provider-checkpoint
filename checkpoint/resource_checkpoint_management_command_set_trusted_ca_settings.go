package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetTrustedCaSettings() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetTrustedCaSettings,
		Read:   readManagementSetTrustedCaSettings,
		Delete: deleteManagementSetTrustedCaSettings,
		Schema: map[string]*schema.Schema{
			"automatic_update": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether the trusted CAs package should be updated automatically.",
			},
		},
	}
}

func createManagementSetTrustedCaSettings(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("automatic_update"); ok {
		payload["automatic-update"] = v.(bool)
	}

	SetTrustedCaSettingsRes, _ := client.ApiCall("set-trusted-ca-settings", payload, client.GetSessionID(), true, false)
	if !SetTrustedCaSettingsRes.Success {
		return fmt.Errorf(SetTrustedCaSettingsRes.ErrorMsg)
	}
	d.SetId("set-trusted-ca-settings" + acctest.RandString(10))
	return readManagementSetTrustedCaSettings(d, m)
}

func readManagementSetTrustedCaSettings(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetTrustedCaSettings(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
