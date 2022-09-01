package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementThreatAdvancedSettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatAdvancedSettingsRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type.",
			},
			"feed_retrieving_interval": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Feed retrieving intervals of External Feed, in the form of HH:MM.",
			},
			"httpi_non_standard_ports": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable HTTP Inspection on non standard ports for Threat Prevention blades.",
			},
			"internal_error_fail_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "In case of internal system error, allow or block all connections.",
			},
			"log_unification_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Session unification timeout for logs (minutes).",
			},
			"resource_classification": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Allow (Background) or Block (Hold) requests until categorization is complete.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_settings": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Custom resources classification per service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"anti_bot": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Custom Settings for Anti Bot Blade.",
									},
									"anti_virus": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Custom Settings for Anti Virus Blade.",
									},
									"zero_phishing": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Custom Settings for Zero Phishing Blade.",
									},
								},
							},
						},
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Set all services to the same mode or choose a custom mode.",
						},
						"web_service_fail_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Block connections when the web service is unavailable.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementThreatAdvancedSettingsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showThreatAdvancedSettingsRes, err := client.ApiCall("show-threat-advanced-settings", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if !showThreatAdvancedSettingsRes.Success {
		fmt.Errorf(showThreatAdvancedSettingsRes.ErrorMsg)
	}

	threatAdvancedSettings := showThreatAdvancedSettingsRes.GetData()

	log.Println("Read Threat Advanced Settings - Show JSON = ", threatAdvancedSettings)

	d.SetId("show-threat-advanced-settings-" + acctest.RandString(10))

	if v := threatAdvancedSettings["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	if v := threatAdvancedSettings["type"]; v != nil {
		_ = d.Set("type", v)
	}

	if v := threatAdvancedSettings["feed-retrieving-interval"]; v != nil {
		_ = d.Set("feed_retrieving_interval", v)
	}

	if v := threatAdvancedSettings["httpi-non-standard-ports"]; v != nil {
		_ = d.Set("httpi_non_standard_ports", v)
	}

	if v := threatAdvancedSettings["internal-error-fail-mode"]; v != nil {
		_ = d.Set("internal_error_fail_mode", v)
	}

	if v := threatAdvancedSettings["log-unification-timeout"]; v != nil {
		_ = d.Set("log_unification_timeout", v)
	}

	if threatAdvancedSettings["resource-classification"] != nil {
		resourceClassificationMap := threatAdvancedSettings["resource-classification"].(map[string]interface{})

		resourceClassificationMapToReturn := make(map[string]interface{})

		if resourceClassificationMap["custom-settings"] != nil {
			customSettingsMap := resourceClassificationMap["custom-settings"].(map[string]interface{})

			customSettingsMapToReturn := make(map[string]interface{})

			if v, _ := customSettingsMap["anti-bot"]; v != nil {
				customSettingsMapToReturn["anti_bot"] = v
			}
			if v, _ := customSettingsMap["anti-virus"]; v != nil {
				customSettingsMapToReturn["anti_virus"] = v
			}
			if v, _ := customSettingsMap["zero-phishing"]; v != nil {
				customSettingsMapToReturn["zero_phishing"] = v
			}

			resourceClassificationMapToReturn["custom_settings"] = customSettingsMapToReturn
		}

		if v, _ := resourceClassificationMap["mode"]; v != nil {
			resourceClassificationMapToReturn["mode"] = v
		}
		if v, _ := resourceClassificationMap["web-service-fail-mode"]; v != nil {
			resourceClassificationMapToReturn["web_service_fail_mode"] = v
		}

		_ = d.Set("resource_classification", resourceClassificationMapToReturn)
	} else {
		_ = d.Set("resource_classification", nil)
	}

	return nil
}
