package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCMEGWConfigurationsGCP() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEGWConfigurationsGCPRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the configuration.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The GW version.",
			},
			"sic_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The configuration sic key.",
			},
			"policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configuration policy.",
			},
			"related_account": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Related account name (aws/azure/gcp accounts)",
			},
			"section_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of a rule section in the Access and NAT layers in the policy, where to insert the automatically generated rules.",
			},
			"x_forwarded_for": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable XFF headers in HTTP / HTTPS requests.",
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the gateways objects in SmartConsole.",
			},
			"communication_with_servers_behind_nat": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway behind NAT communications settings with the Check Point Servers" +
					"(Management, Multi-Domain, Log Servers).",
			},
			"blades": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Dictionary of activated/deactivated blades on the GW.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ips": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "IPS blade",
						},
						"identity_awareness": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Identity Awareness blade",
						},
						"content_awareness": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Content Awareness blade",
						},
						"https_inspection": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "HTTPS Inspection blade",
						},
						"application_control": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Application Control blade",
						},
						"url_filtering": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "URL Filtering blade",
						},
						"anti_bot": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Anti-Bot blade",
						},
						"anti_virus": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Anti-Virus blade",
						},
						"threat_emulation": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Threat Emulation blade",
						},
						"ipsec_vpn": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "IPsec VPN blade",
						},
						"vpn": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "VPN blade",
						},
						"autonomous_threat_prevention": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Autonomous Threat Prevention blade.",
						},
					},
				},
			},
			"identity_awareness_settings": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Computed: true,
				Description: "Dictionary of identity awareness settings that can be configured on the gateway: " +
					"enable_cloudguard_controller (enabling IDA Web API) and receive_identities_from (list of PDP gateway to" +
					"receive identities from through identity sharing feature)",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cloudguard_controller": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enabling Web API identity source for CloudGuard Controller",
						},
						"receive_identities_from": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of PDP gateways names to receive identities from through identity sharing",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"repository_gateway_scripts": {
				Type:     schema.TypeList,
				Computed: true,
				Description: "List of objects that each contains name/UID of a script that exists in the scripts repository" +
					" on the Management server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Script name",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Script uid",
						},
						"parameters": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Script parameters (separated by space)",
						},
					},
				},
			},
			"send_logs_to_server": {
				Type:     schema.TypeList,
				Computed: true,
				Description: "Primary Log Servers names to which logs are sent. Defined Log Server will act as Log and" +
					" Alert Servers. Must be defined as part of Log Servers parameters.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Backup Log Servers names to which logs are sent in case Primary Log Servers are unavailable.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_alerts_to_server": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Alert Log Servers names to which alerts are sent.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementCMEGWConfigurationsGCPRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	log.Println("Read cme GCP GW configuration - name = ", name)

	url := CmeApiPath + "/gwConfigurations/" + name

	GCPGWConfigurationRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	gwConfiguration := GCPGWConfigurationRes.GetData()
	if checkIfRequestFailed(gwConfiguration) {
		errMessage := buildErrorMessage(gwConfiguration)
		return fmt.Errorf(errMessage)
	}

	d.SetId("cme-gcp-gw-configuration-" + name + "-" + acctest.RandString(10))

	GCPGWConfiguration := gwConfiguration["result"].(map[string]interface{})

	_ = d.Set("name", GCPGWConfiguration["name"])

	_ = d.Set("version", GCPGWConfiguration["version"])

	_ = d.Set("sic_key", GCPGWConfiguration["sic_key"])

	_ = d.Set("policy", GCPGWConfiguration["policy"])

	_ = d.Set("related_account", GCPGWConfiguration["related_account"])

	var bladesListToReturn []map[string]interface{}
	bladesMapToAdd := make(map[string]interface{})
	if GCPGWConfiguration["blades"] != nil {
		bladesMap := GCPGWConfiguration["blades"].(map[string]interface{})
		bladesMapToAdd["ips"] = bladesMap["ips"]
		bladesMapToAdd["identity_awareness"] = bladesMap["identity-awareness"]
		bladesMapToAdd["content_awareness"] = bladesMap["content-awareness"]
		bladesMapToAdd["https_inspection"] = bladesMap["https-inspection"]
		bladesMapToAdd["application_control"] = bladesMap["application-control"]
		bladesMapToAdd["url_filtering"] = bladesMap["url-filtering"]
		bladesMapToAdd["anti_bot"] = bladesMap["anti-bot"]
		bladesMapToAdd["anti_virus"] = bladesMap["anti-virus"]
		bladesMapToAdd["threat_emulation"] = bladesMap["threat-emulation"]
		bladesMapToAdd["ipsec_vpn"] = bladesMap["ipsec-vpn"]
		bladesMapToAdd["vpn"] = bladesMap["vpn"]
		bladesMapToAdd["autonomous_threat_prevention"] = bladesMap["autonomous-threat-prevention"]
	} else {
		bladesMapToAdd["ips"] = false
		bladesMapToAdd["identity_awareness"] = false
		bladesMapToAdd["content_awareness"] = false
		bladesMapToAdd["https_inspection"] = false
		bladesMapToAdd["application_control"] = false
		bladesMapToAdd["url_filtering"] = false
		bladesMapToAdd["anti_bot"] = false
		bladesMapToAdd["anti_virus"] = false
		bladesMapToAdd["threat_emulation"] = false
		bladesMapToAdd["ipsec_vpn"] = false
		bladesMapToAdd["vpn"] = false
		bladesMapToAdd["autonomous_threat_prevention"] = false
	}
	bladesListToReturn = append(bladesListToReturn, bladesMapToAdd)
	_ = d.Set("blades", bladesListToReturn)

	var IDASettingsListToReturn []map[string]interface{}
	IDASettingsMapToAdd := make(map[string]interface{})
	if GCPGWConfiguration["identity-awareness-settings"] != nil {
		IDASettingsMap := GCPGWConfiguration["identity-awareness-settings"].(map[string]interface{})
		IDASettingsMapToAdd["enable_cloudguard_controller"] = IDASettingsMap["enable-cloudguard-controller"]
		IDASettingsMapToAdd["receive_identities_from"] = IDASettingsMap["receive-identities-from"]
		IDASettingsListToReturn = append(IDASettingsListToReturn, IDASettingsMapToAdd)
		_ = d.Set("identity_awareness_settings", IDASettingsListToReturn)
	} else {
		_ = d.Set("identity_awareness_settings", nil)
	}
	
	if GCPGWConfiguration["repository-gateway-scripts"] != nil {
		scriptsList := GCPGWConfiguration["repository-gateway-scripts"].([]interface{})
		if len(scriptsList) > 0 {
			var scriptsListToReturn []map[string]interface{}
			for i := range scriptsList {
				scriptMap := scriptsList[i].(map[string]interface{})
				scriptMapToAdd := make(map[string]interface{})
				scriptMapToAdd["name"] = scriptMap["name"]
				scriptMapToAdd["uid"] = scriptMap["uid"]
				scriptMapToAdd["parameters"] = scriptMap["parameters"]
				scriptsListToReturn = append(scriptsListToReturn, scriptMapToAdd)
			}
			_ = d.Set("repository_gateway_scripts", scriptsListToReturn)
		} else {
			_ = d.Set("repository_gateway_scripts", scriptsList)
		}
	} else {
		_ = d.Set("repository_gateway_scripts", nil)
	}

	_ = d.Set("send_logs_to_server", GCPGWConfiguration["send-logs-to-server"])

	_ = d.Set("send_logs_to_backup_server", GCPGWConfiguration["send-logs-to-backup-server"])

	_ = d.Set("send_alerts_to_server", GCPGWConfiguration["send-alerts-to-server"])

	_ = d.Set("section_name", GCPGWConfiguration["section_name"])

	_ = d.Set("x_forwarded_for", GCPGWConfiguration["x_forwarded_for"])

	_ = d.Set("color", GCPGWConfiguration["color"])

	_ = d.Set("communication_with_servers_behind_nat", GCPGWConfiguration["communication-with-servers-behind-nat"])

	return nil
}
