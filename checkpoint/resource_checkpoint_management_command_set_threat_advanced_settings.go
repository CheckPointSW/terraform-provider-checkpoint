package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetThreatAdvancedSettings() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetThreatAdvancedSettings,
		Read:   readManagementSetThreatAdvancedSettings,
		Delete: deleteManagementSetThreatAdvancedSettings,
		Schema: map[string]*schema.Schema{
			"feed_retrieving_interval": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Feed retrieving intervals of External Feed, in the form of HH:MM.",
			},
			"httpi_non_standard_ports": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Enable HTTP Inspection on non standard ports for Threat Prevention blades.",
			},
			"internal_error_fail_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "In case of internal system error, allow or block all connections.",
			},
			"log_unification_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Session unification timeout for logs (minutes).",
			},
			"resource_classification": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Allow (Background) or Block (Hold) requests until categorization is complete.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"custom_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							ForceNew:    true,
							Description: "On Custom mode, custom resources classification per service.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"anti_bot": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Custom Settings for Anti Bot Blade.",
									},
									"anti_virus": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Custom Settings for Anti Virus Blade.",
									},
									"zero_phishing": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Custom Settings for Zero Phishing Blade.",
									},
								},
							},
						},
						"mode": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Set all services to the same mode or choose a custom mode.",
						},
						"web_service_fail_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Block connections when the web service is unavailable.",
						},
					},
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementSetThreatAdvancedSettings(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("feed_retrieving_interval"); ok {
		payload["feed-retrieving-interval"] = v.(string)
	}

	if v, ok := d.GetOkExists("httpi_non_standard_ports"); ok {
		payload["httpi-non-standard-ports"] = v.(bool)
	}

	if v, ok := d.GetOk("internal_error_fail_mode"); ok {
		payload["internal-error-fail-mode"] = v.(string)
	}

	if v, ok := d.GetOk("log_unification_timeout"); ok {
		payload["log-unification-timeout"] = v.(int)
	}

	if _, ok := d.GetOk("resource_classification"); ok {

		res := make(map[string]interface{})

		//if v, ok := d.GetOk("resource_classification.custom_settings"); ok {
		//    res["custom-settings"] = v
		//}
		if _, ok := d.GetOk("resource_classification.custom_settings"); ok {
			customSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("resource_classification.custom_settings.anti_bot"); ok {
				customSettingsPayload["anti-bot"] = v.(string)
			}
			if v, ok := d.GetOk("resource_classification.custom_settings.anti_virus"); ok {
				customSettingsPayload["anti-virus"] = v.(string)
			}
			if v, ok := d.GetOk("resource_classification.custom_settings.zero_phishing"); ok {
				customSettingsPayload["zero-phishing"] = v.(string)
			}
			payload["custom-settings"] = customSettingsPayload
		}

		if v, ok := d.GetOk("resource_classification.mode"); ok {
			res["mode"] = v.(string)
		}
		if v, ok := d.GetOk("resource_classification.web_service_fail_mode"); ok {
			res["web-service-fail-mode"] = v.(string)
		}
		payload["resource-classification"] = res
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	SetThreatAdvancedSettingsRes, _ := client.ApiCall("set-threat-advanced-settings", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !SetThreatAdvancedSettingsRes.Success {
		return fmt.Errorf(SetThreatAdvancedSettingsRes.ErrorMsg)
	}

	d.SetId("set-threat-advanced-settings-" + acctest.RandString(10))
	return readManagementSetThreatAdvancedSettings(d, m)
}

func readManagementSetThreatAdvancedSettings(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetThreatAdvancedSettings(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
