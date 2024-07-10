package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetContentAwarenessAdvancedSettings() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetContentAwarenessAdvancedSettings,
		Read:   readManagementSetContentAwarenessAdvancedSettings,
		Delete: deleteManagementSetContentAwarenessAdvancedSettings,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"internal_error_fail_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "In case of internal system error, allow or block all connections.",
			},
			"supported_services": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Specify the services that Content Awareness inspects.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"httpi_non_standard_ports": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Servers usually send HTTP traffic on TCP port 80. Some servers send HTTP traffic on other ports also. By default, this setting is enabled and Content Awareness inspects HTTP traffic on non-standard ports. You can disable this setting and configure Content Awareness to inspect HTTP traffic only on port 80.",
			},
			"inspect_archives": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Examine the content of archive files. For example, files with the extension .zip, .gz, .tgz, .tar.Z, .tar, .lzma, .tlz, 7z, .rar.",
			},
		},
	}
}

func createManagementSetContentAwarenessAdvancedSettings(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("internal_error_fail_mode"); ok {
		payload["internal-error-fail-mode"] = v.(string)
	}

	if v, ok := d.GetOk("supported_services"); ok {
		payload["supported-services"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("httpi_non_standard_ports"); ok {
		payload["httpi-non-standard-ports"] = v.(bool)
	}

	if v, ok := d.GetOkExists("inspect_archives"); ok {
		payload["inspect-archives"] = v.(bool)
	}

	SetContentAwarenessAdvancedSettingsRes, _ := client.ApiCall("set-content-awareness-advanced-settings", payload, client.GetSessionID(), true, false)
	if !SetContentAwarenessAdvancedSettingsRes.Success {
		return fmt.Errorf(SetContentAwarenessAdvancedSettingsRes.ErrorMsg)
	}

	res := SetContentAwarenessAdvancedSettingsRes.GetData()

	_ = d.Set("uid", res["uid"])
	d.SetId(res["uid"].(string))

	return readManagementSetThreatAdvancedSettings(d, m)
}

func readManagementSetContentAwarenessAdvancedSettings(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetContentAwarenessAdvancedSettings(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
