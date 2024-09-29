package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSetHttpsAdvancedSettings() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSetHttpsAdvancedSettingsRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"bypass_on_client_failure": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether all requests should be bypassed or blocked-in case of client errors (Client closes the connection due to authentication issues during handshake)<br><ul style=\"list-style-type:square\"><li>true - Fail-open (bypass all requests).</li><li>false - Fail-close (block all requests.</li></ul><br>The default value is true.",
			},
			"bypass_on_failure": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether all requests should be bypassed or blocked-in case of server errors (for example validation error during GW-Server authentication)<br><ul style=\"list-style-type:square\"><li>true - Fail-open (bypass all requests).</li><li>false - Fail-close (block all requests.</li></ul><br>The default value is true.",
			},
			"bypass_under_load": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Bypass the HTTPS Inspection temporarily to improve connectivity during a heavy load on the Security Gateway. The HTTPS Inspection would resume as soon as the load decreases.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"track": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to log and send a notification for the bypass under load:<ul style=\"list-style-type:square\"><li>None - Does not record the event.</li><li>Log - Records the event details. Use SmartConsole or SmartView to see the logs.</li><li>Alert - Logs the event and executes a command you configured.</li><li>Mail - Sends an email to the administrator.</li><li>SNMP Trap - Sends an SNMP alert to the configured SNMP Management Server.</li><li>User Defined Alert - Sends a custom alert.</li></ul>.",
						},
					},
				},
			},
			"site_categorization_allow_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Whether all requests should be allowed or blocked until categorization is complete.<br><ul style=\"list-style-type:square\"><li>Background - to allow requests until categorization is complete.</li><li>Hold- to block requests until categorization is complete.</li></ul><br>The default value is hold.",
			},
			"server_certificate_validation_actions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "When a Security Gateway receives an untrusted certificate from a website server, define when to drop the connection and how to track it.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_expired": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with expired server certificate.",
						},
						"block_revoked": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).",
						},
						"block_untrusted": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with untrusted server certificate.",
						},
						"track_errors": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to log and send a notification for the server validation errors:<br><ul style=\"list-style-type:square\"><li>None - Does not record the event.</li><li>Log - Records the event details in SmartView.</li><li>Alert - Logs the event and executes a command.</li><li>Mail - Sends an email to the administrator.</li><li>SNMP Trap - Sends an SNMP alert to the SNMP GU.</li><li>User Defined Alert - Sends customized alerts.</li></ul>.",
						},
					},
				},
			},
			"retrieve_intermediate_ca_certificates": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Configure the value \"true\" to use the \"Certificate Authority Information Access\" extension to retrieve certificates that are missing from the certificate chain.<br>The default value is true.",
			},
			"blocked_certificates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of certificates objects identified by serial number.<br>Drop traffic from servers using the blocked certificate.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Describes the name, cannot be overridden.",
						},
						"cert_serial_number": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Certificate Serial Number (unique) in hexadecimal format HH:HH.",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Describes the certificate by default, can be overridden by any text.",
						},
					},
				},
			},
			"blocked_certificate_tracking": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Controls whether to log and send a notification for dropped traffic.<br><ul style=\"list-style-type:square\"><li>None - Does not record the event.</li><li>Log - Records the event details in SmartView.</li><li>Alert - Logs the event and executes a command.</li><li>Mail - Sends an email to the administrator.</li><li>SNMP Trap - Sends an SNMP alert to the SNMP GU.</li><li>User Defined Alert - Sends customized alerts.</li></ul>.",
			},
			"bypass_update_services": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Configure the value \"true\" to bypass traffic to well-known software update services.<br>The default value is true.",
			},
			"certificate_pinned_apps_action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configure the value \"bypass\" to bypass traffic from certificate-pinned applications approved by Check Point.<br>HTTPS Inspection cannot inspect connections initiated by certificate-pinned applications.<br>Configure the value \"detect\" to send logs for traffic from certificate-pinned applications approved by Check Point.<br>The default value is bypass.",
			},
			"log_sessions": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The value \"true\" configures the Security Gateway to send HTTPS Inspection session logs.<br>The default value is true.",
			},
		},
	}
}

func dataSourceManagementSetHttpsAdvancedSettingsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showHttpsAdvancedSettingsRes, err := client.ApiCall("show-https-advanced-settings", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showHttpsAdvancedSettingsRes.Success {
		return fmt.Errorf(showHttpsAdvancedSettingsRes.ErrorMsg)
	}

	httpsAdvancedSettings := showHttpsAdvancedSettingsRes.GetData()

	log.Println("Read Https Advanced Settings - Show JSON = ", httpsAdvancedSettings)

	if v := httpsAdvancedSettings["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := httpsAdvancedSettings["bypass-on-client-failure"]; v != nil {
		d.Set("bypass_on_client_failure", v)
	}
	if v := httpsAdvancedSettings["bypass-on-failure"]; v != nil {
		d.Set("bypass_on_failure", v)
	}
	if v := httpsAdvancedSettings["bypass-under-load"]; v != nil {
		mapToReturn := make(map[string]interface{})
		v := v.(map[string]interface{})
		if k := v["track"]; k != nil {
			mapToReturn["track"] = k
		}

		d.Set("bypass_under_load", mapToReturn)
	}

	if v := httpsAdvancedSettings["site-categorization-allow-mode"]; v != nil {
		d.Set("site_categorization_allow_mode", v)
	}
	if v := httpsAdvancedSettings["server-certificate-validation-actions"]; v != nil {

		mapToReturn := make(map[string]interface{})
		innerMap := v.(map[string]interface{})

		if v := innerMap["block-expired"]; v != nil {
			mapToReturn["block_expired"] = v
		}
		if v := innerMap["block-revoked"]; v != nil {
			mapToReturn["block_revoked"] = v
		}
		if v := innerMap["block-untrusted"]; v != nil {
			mapToReturn["block_untrusted"] = v
		}
		if v := innerMap["track-errors"]; v != nil {
			mapToReturn["track_errors"] = v
		}

		d.Set("server_certificate_validation_actions", []interface{}{mapToReturn})
	}

	if v := httpsAdvancedSettings["retrieve-intermediate-ca-certificates"]; v != nil {
		d.Set("retrieve_intermediate_ca_certificates", v)
	}
	if v := httpsAdvancedSettings["blocked-certificates"]; v != nil {

		var blockedCertificates []map[string]interface{}

		blockedCertificatesList := v.([]interface{})

		for i := range blockedCertificatesList {

			innerMap := blockedCertificatesList[i].(map[string]interface{})

			mapToReturn := make(map[string]interface{})

			if v := innerMap["name"]; v != nil {
				mapToReturn["name"] = v
			}
			if v := innerMap["cert-serial-number"]; v != nil {
				mapToReturn["cert_serial_number"] = v
			}
			if v := innerMap["comments"]; v != nil {
				mapToReturn["comments"] = v
			}
			blockedCertificates = append(blockedCertificates, mapToReturn)
		}

		d.Set("blocked_certificates", blockedCertificates)
	}

	if v := httpsAdvancedSettings["blocked-certificate-tracking"]; v != nil {
		d.Set("blocked_certificate_tracking", v)
	}
	if v := httpsAdvancedSettings["bypass-update-services"]; v != nil {
		d.Set("bypass_update_services", v)
	}
	if v := httpsAdvancedSettings["certificate-pinned-apps-action"]; v != nil {
		d.Set("certificate_pinned_apps_action", v)
	}
	if v := httpsAdvancedSettings["log-sessions"]; v != nil {
		d.Set("log_sessions", v)
	}

	return nil
}
