package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCMEGWConfigurations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEGWConfigurationsRead,
		Schema: map[string]*schema.Schema{
			"result": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Response data - contains all GW configurations",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
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
						"related_account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Related account name (aws/azure/gcp accounts)",
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
							Description: "Dictionary of identity awareness settings that can be configured by CME: " +
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
				},
			},
		},
	}
}

func dataSourceManagementCMEGWConfigurationsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	log.Println("Read cme GW configurations")

	url := CmeApiPath + "/gwConfigurations"

	cmeGWConfigurationsRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	gwConfigurations := cmeGWConfigurationsRes.GetData()
	if checkIfRequestFailed(gwConfigurations) {
		errMessage := buildErrorMessage(gwConfigurations)
		return fmt.Errorf(errMessage)
	}

	d.SetId("cme-gw-configurations-" + acctest.RandString(10))

	gwConfigurationsList := gwConfigurations["result"].([]interface{})
	var gwConfigurationsListToReturn []map[string]interface{}
	if len(gwConfigurationsList) > 0 {
		for i := range gwConfigurationsList {
			singleGWConfiguration := gwConfigurationsList[i].(map[string]interface{})
			tempObject := make(map[string]interface{})
			tempObject["name"] = singleGWConfiguration["name"]
			tempObject["version"] = singleGWConfiguration["version"]
			tempObject["sic_key"] = singleGWConfiguration["sic_key"]
			tempObject["policy"] = singleGWConfiguration["policy"]
			tempObject["related_account"] = singleGWConfiguration["related_account"]
			tempObject["section_name"] = singleGWConfiguration["section_name"]
			tempObject["x_forwarded_for"] = singleGWConfiguration["x_forwarded_for"]
			tempObject["color"] = singleGWConfiguration["color"]
			tempObject["communication_with_servers_behind_nat"] = singleGWConfiguration["communication-with-servers-behind-nat"]

			var bladesListToReturn []map[string]interface{}
			bladesMapToAdd := make(map[string]interface{})
			if singleGWConfiguration["blades"] != nil {
				bladesMap := singleGWConfiguration["blades"].(map[string]interface{})
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
			tempObject["blades"] = bladesListToReturn

			var IDASettingsListToReturn []map[string]interface{}
			IDASettingsMapToAdd := make(map[string]interface{})
			if singleGWConfiguration["identity-awareness-settings"] != nil {
				IDASettingsMap := singleGWConfiguration["identity-awareness-settings"].(map[string]interface{})
				IDASettingsMapToAdd["enable_cloudguard_controller"] = IDASettingsMap["enable-cloudguard-controller"]
				IDASettingsMapToAdd["receive_identities_from"] = IDASettingsMap["receive-identities-from"]
				IDASettingsListToReturn = append(IDASettingsListToReturn, IDASettingsMapToAdd)
			}
			tempObject["identity_awareness_settings"] = IDASettingsListToReturn

			if singleGWConfiguration["repository-gateway-scripts"] != nil {
				scriptsList := singleGWConfiguration["repository-gateway-scripts"].([]interface{})
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
					tempObject["repository_gateway_scripts"] = scriptsListToReturn
				} else {
					tempObject["repository_gateway_scripts"] = scriptsList
				}
			}
			tempObject["send_logs_to_server"] = singleGWConfiguration["send-logs-to-server"]
			tempObject["send_logs_to_backup_server"] = singleGWConfiguration["send-logs-to-backup-server"]
			tempObject["send_alerts_to_server"] = singleGWConfiguration["send-alerts-to-server"]

			gwConfigurationsListToReturn = append(gwConfigurationsListToReturn, tempObject)
		}
		_ = d.Set("result", gwConfigurationsListToReturn)
	} else {
		_ = d.Set("result", []interface{}{})
	}
	return nil
}
