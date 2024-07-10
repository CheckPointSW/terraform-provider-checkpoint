package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementContentAwarenessAdvancedSettings() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementSetContentAwarenessAdvancedSettingsRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"internal_error_fail_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "In case of internal system error, allow or block all connections.",
			},
			"supported_services": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Specify the services that Content Awareness inspects.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"httpi_non_standard_ports": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Servers usually send HTTP traffic on TCP port 80. Some servers send HTTP traffic on other ports also. By default, this setting is enabled and Content Awareness inspects HTTP traffic on non-standard ports. You can disable this setting and configure Content Awareness to inspect HTTP traffic only on port 80.",
			},
			"inspect_archives": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Examine the content of archive files. For example, files with the extension .zip, .gz, .tgz, .tar.Z, .tar, .lzma, .tlz, 7z, .rar.",
			},
		},
	}
}

func dataSourceManagementSetContentAwarenessAdvancedSettingsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	contentAwarenessAdvancedSettingsRes, _ := client.ApiCall("show-content-awareness-advanced-settings", payload, client.GetSessionID(), true, false)
	if !contentAwarenessAdvancedSettingsRes.Success {
		return fmt.Errorf(contentAwarenessAdvancedSettingsRes.ErrorMsg)
	}
	contentAwarenessAdvancedSettingsData := contentAwarenessAdvancedSettingsRes.GetData()

	if v := contentAwarenessAdvancedSettingsData["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := contentAwarenessAdvancedSettingsData["internal-error-fail-mode"]; v != nil {
		_ = d.Set("internal_error_fail_mode", v)
	}

	if v := contentAwarenessAdvancedSettingsData["supported-services"]; v != nil {
		servicesJson, ok := v.([]interface{})
		if ok {
			log.Println("service jason is ", servicesJson)
			servicesNames := make([]string, 0)
			if len(servicesJson) > 0 {
				for _, svc := range servicesJson {
					services := svc.(map[string]interface{})
					servicesNames = append(servicesNames, services["name"].(string))
				}

			}
			_ = d.Set("supported_services", servicesNames)
		}

	}

	if v := contentAwarenessAdvancedSettingsData["httpi-non-standard-ports"]; v != nil {
		_ = d.Set("httpi_non_standard_ports", v)
	}

	if v := contentAwarenessAdvancedSettingsData["inspect-archives"]; v != nil {
		_ = d.Set("inspect_archives", v)
	}

	return nil
}
