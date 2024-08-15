package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSetTrustedCaSettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSetTrustedCaSettingsRead,
		Schema: map[string]*schema.Schema{
			"automatic_update": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the trusted CAs package should be updated automatically.",
			},
		},
	}
}

func dataSourceManagementSetTrustedCaSettingsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showTrustedCaSettingsRes, err := client.ApiCall("show-trusted-ca-settings", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTrustedCaSettingsRes.Success {
		return fmt.Errorf(showTrustedCaSettingsRes.ErrorMsg)
	}

	trustedCaSettings := showTrustedCaSettingsRes.GetData()

	log.Println("Read Trusted CA Settings - Show JSON = ", trustedCaSettings)

	d.SetId("set-trusted-ca-settings" + acctest.RandString(10))

	if v := trustedCaSettings["automatic-update"]; v != nil {

		_ = d.Set("automatic_update", v.(bool))
	}

	return nil
}
