package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementCMEGWConfigurationsGCP() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEGWConfigurationsGCP,
		Update: updateManagementCMEGWConfigurationsGCP,
		Read:   readManagementCMEGWConfigurationsGCP,
		Delete: deleteManagementCMEGWConfigurationsGCP,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GW configuration name.",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The GW version.",
			},
			"base64_sic_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Base64 key for trusted communication between management and GW.",
			},
			"policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Policy name to be installed on the GW.",
			},
			"related_account": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CME account to associate with the GW Configuration.",
			},
			"blades": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "Dictionary of activated/deactivated blades on the GW.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ips": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "IPS blade",
						},
						"identity_awareness": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Identity Awareness blade",
						},
						"content_awareness": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Content Awareness blade",
						},
						"https_inspection": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "HTTPS Inspection blade",
						},
						"application_control": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Application Control blade",
						},
						"url_filtering": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "URL Filtering blade",
						},
						"anti_bot": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Anti-Bot blade",
						},
						"anti_virus": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Anti-Virus blade",
						},
						"threat_emulation": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Threat Emulation blade",
						},
						"ipsec_vpn": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "IPsec VPN blade",
						},
						"vpn": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "VPN blade",
						},
						"autonomous_threat_prevention": {
							Type:        schema.TypeBool,
							Required:    true,
							Description: "Autonomous Threat Prevention blade.",
						},
					},
				},
			},
			"repository_gateway_scripts": {
				Type:     schema.TypeList,
				Optional: true,
				Description: "List of objects that each contains name/UID of a script that exists in the scripts repository" +
					" on the Management server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Script name",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Script uid",
						},
						"parameters": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Script parameters (separated by space)",
						},
					},
				},
			},
			"send_logs_to_server": {
				Type:     schema.TypeList,
				Optional: true,
				Description: "Primary Log Servers names to which logs are sent. Defined Log Server will act as Log and" +
					" Alert Servers. Must be defined as part of Log Servers parameters.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Backup Log Servers names to which logs are sent in case Primary Log Servers are unavailable.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_alerts_to_server": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Alert Log Servers names to which alerts are sent.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func readManagementCMEGWConfigurationsGCP(d *schema.ResourceData, m interface{}) error {
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
		if cmeObjectNotFound(gwConfiguration) {
			d.SetId("")
			return nil
		}
		errMessage := buildErrorMessage(gwConfiguration)
		return fmt.Errorf(errMessage)
	}

	GCPGWConfiguration := gwConfiguration["result"].(map[string]interface{})

	_ = d.Set("name", GCPGWConfiguration["name"])

	_ = d.Set("version", GCPGWConfiguration["version"])

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

	return nil

}

func createManagementCMEGWConfigurationsGCP(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})

	if v, ok := d.GetOk("version"); ok {
		payload["version"] = v.(string)
	}
	if v, ok := d.GetOk("base64_sic_key"); ok {
		payload["base64_sic_key"] = v.(string)
	}
	if v, ok := d.GetOk("policy"); ok {
		payload["policy"] = v.(string)
	}
	if v, ok := d.GetOk("related_account"); ok {
		payload["related_account"] = v.(string)
	}
	if v, ok := d.GetOk("repository_gateway_scripts"); ok {
		scriptsList := v.([]interface{})
		if len(scriptsList) > 0 {
			var scriptsPayload []map[string]interface{}
			for i := range scriptsList {
				tempObject := make(map[string]interface{})

				if v, ok := d.GetOk("repository_gateway_scripts." + strconv.Itoa(i) + ".name"); ok {
					tempObject["name"] = v.(string)
				}
				if v, ok := d.GetOk("repository_gateway_scripts." + strconv.Itoa(i) + ".uid"); ok {
					tempObject["uid"] = v.(string)
				}
				if v, ok := d.GetOk("repository_gateway_scripts." + strconv.Itoa(i) + ".parameters"); ok {
					tempObject["parameters"] = v.(string)
				}
				scriptsPayload = append(scriptsPayload, tempObject)
			}
			payload["repository_gateway_scripts"] = scriptsPayload
		} else {
			payload["repository_gateway_scripts"] = scriptsList
		}
	}

	if v, ok := d.GetOk("send_logs_to_server"); ok {
		payload["send_logs_to_server"] = v.([]interface{})
	}
	if v, ok := d.GetOk("send_logs_to_backup_server"); ok {
		payload["send_logs_to_backup_server"] = v.([]interface{})
	}
	if v, ok := d.GetOk("send_alerts_to_server"); ok {
		payload["send_alerts_to_server"] = v.([]interface{})
	}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}
	if _, ok := d.GetOk("blades"); ok {
		tempObject := make(map[string]interface{})
		if v, ok := d.GetOk("blades.0.ips"); ok {
			tempObject["ips"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.identity_awareness"); ok {
			tempObject["identity-awareness"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.content_awareness"); ok {
			tempObject["content-awareness"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.https_inspection"); ok {
			tempObject["https-inspection"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.application_control"); ok {
			tempObject["application-control"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.url_filtering"); ok {
			tempObject["url-filtering"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.anti_bot"); ok {
			tempObject["anti-bot"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.anti_virus"); ok {
			tempObject["anti-virus"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.threat_emulation"); ok {
			tempObject["threat-emulation"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.ipsec_vpn"); ok {
			tempObject["ipsec-vpn"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.vpn"); ok {
			tempObject["vpn"] = v.(bool)
		}
		if v, ok := d.GetOk("blades.0.autonomous_threat_prevention"); ok {
			tempObject["autonomous-threat-prevention"] = v.(bool)
		}
		payload["blades"] = tempObject
	}

	log.Println("Create cme GCP GW configuration - name = ", payload["name"])

	url := CmeApiPath + "/gwConfigurations/gcp"

	cmeGWConfigurationRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeGWConfigurationRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	d.SetId("cme-gcp-gw-configuration-" + d.Get("name").(string) + "-" + acctest.RandString(10))

	return readManagementCMEGWConfigurationsGCP(d, m)
}

func updateManagementCMEGWConfigurationsGCP(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})

	if d.HasChange("version") {
		payload["version"] = d.Get("version")
	}
	if d.HasChange("base64_sic_key") {
		payload["base64_sic_key"] = d.Get("base64_sic_key")
	}
	if d.HasChange("policy") {
		payload["policy"] = d.Get("policy")
	}
	if d.HasChange("related_account") {
		payload["related_account"] = d.Get("related_account")
	}
	if d.HasChange("repository_gateway_scripts") {
		if v, ok := d.GetOk("repository_gateway_scripts"); ok {
			scriptsList := v.([]interface{})
			if len(scriptsList) > 0 {
				var scriptsPayload []map[string]interface{}
				for i := range scriptsList {
					tempObject := make(map[string]interface{})

					if v, ok := d.GetOk("repository_gateway_scripts." + strconv.Itoa(i) + ".name"); ok {
						tempObject["name"] = v.(string)
					}
					if v, ok := d.GetOk("repository_gateway_scripts." + strconv.Itoa(i) + ".uid"); ok {
						tempObject["uid"] = v.(string)
					}
					if v, ok := d.GetOk("repository_gateway_scripts." + strconv.Itoa(i) + ".parameters"); ok {
						tempObject["parameters"] = v.(string)
					}
					scriptsPayload = append(scriptsPayload, tempObject)
				}
				payload["repository_gateway_scripts"] = scriptsPayload
			} else {
				payload["repository_gateway_scripts"] = scriptsList
			}
		} else {
			payload["repository_gateway_scripts"] = v.([]interface{})
		}
	}
	if d.HasChange("send_logs_to_server") {
		payload["send_logs_to_server"] = d.Get("send_logs_to_server")
	}
	if d.HasChange("send_logs_to_backup_server") {
		payload["send_logs_to_backup_server"] = d.Get("send_logs_to_backup_server")
	}
	if d.HasChange("send_alerts_to_server") {
		payload["send_alerts_to_server"] = d.Get("send_alerts_to_server")
	}
	if d.HasChange("blades") {
		tempObject := make(map[string]interface{})
		if d.HasChange("blades.0.ips") {
			tempObject["ips"] = d.Get("blades.0.ips")
		}
		if d.HasChange("blades.0.identity_awareness") {
			tempObject["identity-awareness"] = d.Get("blades.0.identity_awareness")
		}
		if d.HasChange("blades.0.content_awareness") {
			tempObject["content-awareness"] = d.Get("blades.0.content_awareness")
		}
		if d.HasChange("blades.0.https_inspection") {
			tempObject["https-inspection"] = d.Get("blades.0.https_inspection")
		}
		if d.HasChange("blades.0.application_control") {
			tempObject["application-control"] = d.Get("blades.0.application_control")
		}
		if d.HasChange("blades.0.url_filtering") {
			tempObject["url-filtering"] = d.Get("blades.0.url_filtering")
		}
		if d.HasChange("blades.0.anti_bot") {
			tempObject["anti-bot"] = d.Get("blades.0.anti_bot")
		}
		if d.HasChange("blades.0.anti_virus") {
			tempObject["anti-virus"] = d.Get("blades.0.anti_virus")
		}
		if d.HasChange("blades.0.threat_emulation") {
			tempObject["threat-emulation"] = d.Get("blades.0.threat_emulation")
		}
		if d.HasChange("blades.0.ipsec_vpn") {
			tempObject["ipsec-vpn"] = d.Get("blades.0.ipsec_vpn")
		}
		if d.HasChange("blades.0.vpn") {
			tempObject["vpn"] = d.Get("blades.0.vpn")
		}
		if d.HasChange("blades.0.autonomous_threat_prevention") {
			tempObject["autonomous-threat-prevention"] = d.Get("blades.0.autonomous_threat_prevention")
		}
		payload["blades"] = tempObject
	}
	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	log.Println("Set cme GCP GW configuration - name = ", name)

	url := CmeApiPath + "/gwConfigurations/gcp/" + name
	cmeGWConfigurationRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed(), "PUT")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeGWConfigurationRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	return readManagementCMEGWConfigurationsGCP(d, m)

}

func deleteManagementCMEGWConfigurationsGCP(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	url := CmeApiPath + "/gwConfigurations/" + name

	log.Println("Delete cme GCP GW configuration - name = ", name)
	res, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "DELETE")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := res.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	d.SetId("")
	return nil
}
