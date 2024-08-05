package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func dataSourceManagementCMEGWConfigurationsAWS() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEGWConfigurationsAWSRead,
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
				Type:     schema.TypeString,
				Computed: true,
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
			"vpn_domain": {
				Type:     schema.TypeString,
				Computed: true,
				Description: "The group object to be set as the VPN domain for the VPN gateway." +
					" An empty string will automatically set an empty group as the encryption domain." +
					" Always empty string for 'TGW' deployment type.",
			},
			"vpn_community": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A star community in which to place the VPN gateway as center.",
			},
			"deployment_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The deployment type of the CloudGuard Security Gateways.",
			},
			"tgw_static_routes": {
				Type:     schema.TypeList,
				Computed: true,
				Description: "Comma separated list of cidrs, for each cidr a static route will be created on each" +
					" gateway of the TGW auto scaling group.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tgw_spoke_routes": {
				Type:     schema.TypeList,
				Computed: true,
				Description: "Comma separated list of spoke cidrs, each spoke cidr that was learned from the TGW over" +
					" bgp will be re-advertised by the gateways of the TGW auto scaling group to the AWS TGW.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func dataSourceManagementCMEGWConfigurationsAWSRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	log.Println("Read cme AWS GW configuration - name = ", name)

	url := CmeApiPath + "/gwConfigurations/" + name

	AWSGWConfigurationRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	gwConfiguration := AWSGWConfigurationRes.GetData()
	if checkIfRequestFailed(gwConfiguration) {
		errMessage := buildErrorMessage(gwConfiguration)
		return fmt.Errorf(errMessage)
	}

	d.SetId("cme-aws-gw-configuration-" + name + "-" + acctest.RandString(10))

	AWSGWConfiguration := gwConfiguration["result"].(map[string]interface{})

	_ = d.Set("name", AWSGWConfiguration["name"])

	_ = d.Set("version", AWSGWConfiguration["version"])

	_ = d.Set("sic_key", AWSGWConfiguration["sic_key"])

	_ = d.Set("policy", AWSGWConfiguration["policy"])

	_ = d.Set("related_account", AWSGWConfiguration["related_account"])

	var bladesListToReturn []map[string]interface{}
	bladesMapToAdd := make(map[string]interface{})
	if AWSGWConfiguration["blades"] != nil {
		bladesMap := AWSGWConfiguration["blades"].(map[string]interface{})
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

	if AWSGWConfiguration["repository-gateway-scripts"] != nil {
		scriptsList := AWSGWConfiguration["repository-gateway-scripts"].([]interface{})
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
	_ = d.Set("vpn_domain", AWSGWConfiguration["vpn_domain"])

	_ = d.Set("vpn_community", AWSGWConfiguration["vpn_community"])

	_ = d.Set("deployment_type", AWSGWConfiguration["deployment_type"])

	if tgwStaticRoutes, ok := AWSGWConfiguration["tgw_static_routes"].(string); ok {
		_ = d.Set("tgw_static_routes", strings.Split(tgwStaticRoutes, ","))
	}

	if tgwSpokeRoutes, ok := AWSGWConfiguration["tgw_spoke_routes"].(string); ok {
		_ = d.Set("tgw_spoke_routes", strings.Split(tgwSpokeRoutes, ","))
	}

	_ = d.Set("send_logs_to_server", AWSGWConfiguration["send-logs-to-server"])

	_ = d.Set("send_logs_to_backup_server", AWSGWConfiguration["send-logs-to-backup-server"])

	_ = d.Set("send_alerts_to_server", AWSGWConfiguration["send-alerts-to-server"])

	_ = d.Set("section_name", AWSGWConfiguration["section_name"])

	_ = d.Set("x_forwarded_for", AWSGWConfiguration["x_forwarded_for"])

	_ = d.Set("color", AWSGWConfiguration["color"])

	_ = d.Set("communication_with_servers_behind_nat", AWSGWConfiguration["communication-with-servers-behind-nat"])

	return nil
}
