package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"strconv"
)

func resourceManagementSetHttpsAdvancedSettings() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetHttpsAdvancedSettings,
		Read:   readManagementSetHttpsAdvancedSettings,
		Delete: deleteManagementSetHttpsAdvancedSettings,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"bypass_on_client_failure": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether all requests should be bypassed or blocked-in case of client errors (Client closes the connection due to authentication issues during handshake)<br><ul style=\"list-style-type:square\"><li>true - Fail-open (bypass all requests).</li><li>false - Fail-close (block all requests.</li></ul><br>The default value is true.",
			},
			"bypass_on_failure": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether all requests should be bypassed or blocked-in case of server errors (for example validation error during GW-Server authentication)<br><ul style=\"list-style-type:square\"><li>true - Fail-open (bypass all requests).</li><li>false - Fail-close (block all requests.</li></ul><br>The default value is true.",
			},
			"bypass_under_load": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Bypass the HTTPS Inspection temporarily to improve connectivity during a heavy load on the Security Gateway. The HTTPS Inspection would resume as soon as the load decreases.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"track": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Whether to log and send a notification for the bypass under load:<ul style=\"list-style-type:square\"><li>None - Does not record the event.</li><li>Log - Records the event details. Use SmartConsole or SmartView to see the logs.</li><li>Alert - Logs the event and executes a command you configured.</li><li>Mail - Sends an email to the administrator.</li><li>SNMP Trap - Sends an SNMP alert to the configured SNMP Management Server.</li><li>User Defined Alert - Sends a custom alert.</li></ul>.",
							Default:     "Alert",
						},
					},
				},
			},
			"site_categorization_allow_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether all requests should be allowed or blocked until categorization is complete.<br><ul style=\"list-style-type:square\"><li>Background - to allow requests until categorization is complete.</li><li>Hold- to block requests until categorization is complete.</li></ul><br>The default value is hold.",
			},
			"server_certificate_validation_actions": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "When a Security Gateway receives an untrusted certificate from a website server, define when to drop the connection and how to track it.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"block_expired": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Set to be true in order to drop traffic from servers with expired server certificate.",
							Default:     false,
						},
						"block_revoked": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).",
							Default:     true,
						},
						"block_untrusted": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Set to be true in order to drop traffic from servers with untrusted server certificate.",
							Default:     false,
						},
						"track_errors": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Whether to log and send a notification for the server validation errors:<br><ul style=\"list-style-type:square\"><li>None - Does not record the event.</li><li>Log - Records the event details in SmartView.</li><li>Alert - Logs the event and executes a command.</li><li>Mail - Sends an email to the administrator.</li><li>SNMP Trap - Sends an SNMP alert to the SNMP GU.</li><li>User Defined Alert - Sends customized alerts.</li></ul>.",
						},
					},
				},
			},
			"retrieve_intermediate_ca_certificates": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Configure the value \"true\" to use the \"Certificate Authority Information Access\" extension to retrieve certificates that are missing from the certificate chain.<br>The default value is true.",
			},
			"blocked_certificates": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Collection of certificates objects identified by serial number.<br>Drop traffic from servers using the blocked certificate.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Describes the name, cannot be overridden.",
						},
						"cert_serial_number": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Certificate Serial Number (unique) in hexadecimal format HH:HH.",
						},
						"comments": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Describes the certificate by default, can be overridden by any text.",
						},
					},
				},
			},
			"blocked_certificate_tracking": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Controls whether to log and send a notification for dropped traffic.<br><ul style=\"list-style-type:square\"><li>None - Does not record the event.</li><li>Log - Records the event details in SmartView.</li><li>Alert - Logs the event and executes a command.</li><li>Mail - Sends an email to the administrator.</li><li>SNMP Trap - Sends an SNMP alert to the SNMP GU.</li><li>User Defined Alert - Sends customized alerts.</li></ul>.",
			},
			"bypass_update_services": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Configure the value \"true\" to bypass traffic to well-known software update services.<br>The default value is true.",
			},
			"certificate_pinned_apps_action": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Configure the value \"bypass\" to bypass traffic from certificate-pinned applications approved by Check Point.<br>HTTPS Inspection cannot inspect connections initiated by certificate-pinned applications.<br>Configure the value \"detect\" to send logs for traffic from certificate-pinned applications approved by Check Point.<br>The default value is bypass.",
			},
			"log_sessions": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "The value \"true\" configures the Security Gateway to send HTTPS Inspection session logs.<br>The default value is true.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementSetHttpsAdvancedSettings(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("bypass_on_client_failure"); ok {
		payload["bypass-on-client-failure"] = v.(bool)
	}

	if v, ok := d.GetOkExists("bypass_on_failure"); ok {
		payload["bypass-on-failure"] = v.(bool)
	}

	if _, ok := d.GetOk("bypass_under_load"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("bypass_under_load.track"); ok {
			res["track"] = v.(string)
		}
		payload["bypass-under-load"] = res
	}

	if v, ok := d.GetOk("site_categorization_allow_mode"); ok {
		payload["site-categorization-allow-mode"] = v.(string)
	}

	if _, ok := d.GetOk("server_certificate_validation_actions"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("server_certificate_validation_actions.block_expired"); ok {
			res["block-expired"] = v
		}
		if v, ok := d.GetOk("server_certificate_validation_actions.block_revoked"); ok {
			res["block-revoked"] = v
		}
		if v, ok := d.GetOk("server_certificate_validation_actions.block_untrusted"); ok {
			res["block-untrusted"] = v
		}
		if v, ok := d.GetOk("server_certificate_validation_actions.track_errors"); ok {
			res["track-errors"] = v
		}
		payload["server-certificate-validation-actions"] = res
	}

	if v, ok := d.GetOkExists("retrieve_intermediate_ca_certificates"); ok {
		payload["retrieve-intermediate-ca-certificates"] = v.(bool)
	}

	if v, ok := d.GetOk("blocked_certificates"); ok {

		blockedCertificatesList := v.([]interface{})

		if len(blockedCertificatesList) > 0 {

			var blockedCertificatesPayload []map[string]interface{}

			for i := range blockedCertificatesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("blocked_certificates." + strconv.Itoa(i) + ".name"); ok {
					Payload["name"] = v.(string)
				}
				if v, ok := d.GetOk("blocked_certificates." + strconv.Itoa(i) + ".cert_serial_number"); ok {
					Payload["cert-serial-number"] = v.(string)
				}
				if v, ok := d.GetOk("blocked_certificates." + strconv.Itoa(i) + ".comments"); ok {
					Payload["comments"] = v.(string)
				}
				blockedCertificatesPayload = append(blockedCertificatesPayload, Payload)
			}
			payload["blocked-certificates"] = blockedCertificatesPayload
		}
	}

	if v, ok := d.GetOk("blocked_certificate_tracking"); ok {
		payload["blocked-certificate-tracking"] = v.(string)
	}

	if v, ok := d.GetOkExists("bypass_update_services"); ok {
		payload["bypass-update-services"] = v.(bool)
	}

	if v, ok := d.GetOk("certificate_pinned_apps_action"); ok {
		payload["certificate-pinned-apps-action"] = v.(string)
	}

	if v, ok := d.GetOkExists("log_sessions"); ok {
		payload["log-sessions"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	SetHttpsAdvancedSettingsRes, _ := client.ApiCall("set-https-advanced-settings", payload, client.GetSessionID(), true, false)
	if !SetHttpsAdvancedSettingsRes.Success {
		return fmt.Errorf(SetHttpsAdvancedSettingsRes.ErrorMsg)
	}

	setHttpsAdvancedSettingsResData := SetHttpsAdvancedSettingsRes.GetData()
	if v := setHttpsAdvancedSettingsResData["uid"]; v != nil {
		d.Set("uid", v)
		d.SetId(v.(string))
	}

	return readManagementSetHttpsAdvancedSettings(d, m)
}

func readManagementSetHttpsAdvancedSettings(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetHttpsAdvancedSettings(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
