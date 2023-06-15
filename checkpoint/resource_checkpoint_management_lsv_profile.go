package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
)

func resourceManagementLsvProfile() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLsvProfile,
		Read:   readManagementLsvProfile,
		Update: updateManagementLsvProfile,
		Delete: deleteManagementLsvProfile,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"certificate_authority": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Trusted Certificate authority for establishing trust between VPN peers, identified by name or UID.",
			},
			"allowed_ip_addresses": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of network objects identified by name or UID that represent IP addresses allowed in profile's VPN domain.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"restrict_allowed_addresses": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Indicate whether the IP addresses allowed in the VPN Domain will be restricted or not, according to allowed-ip-addresses field.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpn_domain": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "peers' VPN Domain properties.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limit_peer_domain_size": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Use this parameter to limit the number of IP addresses in the VPN Domain of each peer according to the value in the max-allowed-addresses field.",
						},
						"max_allowed_addresses": {
							Type:        schema.TypeInt,
							Optional:    true,
							Default:     256,
							Description: "Maximum number of IP addresses in the VPN Domain of each peer. This value will be enforced only when limit-peer-domain-size field is set to true. Select a value between 1 and 256. Default value is 256.",
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "black",
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementLsvProfile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	lsvProfile := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		lsvProfile["name"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_authority"); ok {
		lsvProfile["certificate-authority"] = v.(string)
	}

	if v, ok := d.GetOk("allowed_ip_addresses"); ok {
		lsvProfile["allowed-ip-addresses"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("restrict_allowed_addresses"); ok {
		lsvProfile["restrict-allowed-addresses"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		lsvProfile["tags"] = v.(*schema.Set).List()
	}

	if _, ok := d.GetOk("vpn_domain"); ok {
		res := make(map[string]interface{})

		if v, ok := d.GetOk("vpn_domain.limit_peer_domain_size"); ok {
			res["limit-peer-domain-size"] = v
		}
		if v, ok := d.GetOk("vpn_domain.max_allowed_addresses"); ok {
			res["max-allowed-addresses"] = v
		}

		lsvProfile["vpn-domain"] = res
	}

	if v, ok := d.GetOk("color"); ok {
		lsvProfile["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		lsvProfile["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		lsvProfile["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		lsvProfile["ignore-errors"] = v.(bool)
	}

	log.Println("Create Lsv Profile - Map = ", lsvProfile)

	addLsvProfileRes, err := client.ApiCall("add-lsv-profile", lsvProfile, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addLsvProfileRes.Success {
		if addLsvProfileRes.ErrorMsg != "" {
			return fmt.Errorf(addLsvProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addLsvProfileRes.GetData()["uid"].(string))

	return readManagementLsvProfile(d, m)
}

func readManagementLsvProfile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showLsvProfileRes, err := client.ApiCall("show-lsv-profile", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLsvProfileRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showLsvProfileRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLsvProfileRes.ErrorMsg)
	}

	lsvProfile := showLsvProfileRes.GetData()

	log.Println("Read Lsv Profile - Show JSON = ", lsvProfile)

	if v := lsvProfile["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := lsvProfile["certificate-authority"]; v != nil {
		_ = d.Set("certificate_authority", v)
	}

	if lsvProfile["allowed-ip-addresses"] != nil {
		allowedIpAddressesJson := lsvProfile["allowed-ip-addresses"].([]interface{})
		var allowedIpAddressesIds = make([]string, 0)
		if len(allowedIpAddressesJson) > 0 {
			for _, allowedIpAddress := range allowedIpAddressesJson {
				allowedIpAddress := allowedIpAddress.(map[string]interface{})
				allowedIpAddressesIds = append(allowedIpAddressesIds, allowedIpAddress["name"].(string))
			}
		}
		_ = d.Set("allowed_ip_addresses", allowedIpAddressesIds)
	} else {
		_ = d.Set("allowed_ip_addresses", nil)
	}

	if v := lsvProfile["restrict-allowed-addresses"]; v != nil {
		_ = d.Set("restrict_allowed_addresses", v)
	}

	if lsvProfile["tags"] != nil {
		tagsJson := lsvProfile["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	if lsvProfile["vpn-domain"] != nil {
		vpnDomainMap := lsvProfile["vpn-domain"].(map[string]interface{})

		vpnDomainMapToReturn := make(map[string]interface{})

		if v, _ := vpnDomainMap["limit-peer-domain-size"]; v != nil {
			vpnDomainMapToReturn["limit_peer_domain_size"] = v
		}
		if v, _ := vpnDomainMap["max-allowed-addresses"]; v != nil {
			vpnDomainMapToReturn["max_allowed_addresses"] = v
		}

		_, vpnDomainInConf := d.GetOk("vpn_domain")
		defaultVpnDomain := map[string]interface{}{"limit_peer_domain_size": "false", "max_allowed_addresses": "256"}
		if reflect.DeepEqual(defaultVpnDomain, vpnDomainMapToReturn) && !vpnDomainInConf {
			_ = d.Set("vpn_domain", map[string]interface{}{})
		} else {
			_ = d.Set("vpn_domain", vpnDomainMapToReturn)
		}

	} else {
		_ = d.Set("vpn_domain", nil)
	}

	if v := lsvProfile["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := lsvProfile["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementLsvProfile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	lsvProfile := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		lsvProfile["name"] = oldName
		lsvProfile["new-name"] = newName
	} else {
		lsvProfile["name"] = d.Get("name")
	}

	if ok := d.HasChange("certificate_authority"); ok {
		lsvProfile["certificate-authority"] = d.Get("certificate_authority")
	}

	if ok := d.HasChange("allowed_ip_addresses"); ok {
		if v, ok := d.GetOk("allowed_ip_addresses"); ok {
			lsvProfile["allowed-ip-addresses"] = v.(*schema.Set).List()
		} else {
			old, _ := d.GetChange("allowed_ip_addresses")
			lsvProfile["allowed-ip-addresses"] = map[string]interface{}{"remove": old.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("restrict_allowed_addresses"); ok {
		lsvProfile["restrict-allowed-addresses"] = v.(bool)
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			lsvProfile["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			lsvProfile["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("vpn_domain"); ok {

		if _, ok := d.GetOk("vpn_domain"); ok {
			res := make(map[string]interface{})

			if v, ok := d.GetOk("vpn_domain.limit_peer_domain_size"); ok {
				res["limit-peer-domain-size"] = v
			}
			if v, ok := d.GetOk("vpn_domain.max_allowed_addresses"); ok {
				res["max-allowed-addresses"] = v
			}

			lsvProfile["vpn-domain"] = res
		} else {
			lsvProfile["vpn-domain"] = map[string]interface{}{"limit-peer-domain_size": "false", "max-allowed-addresses": "256"}
		}
	}

	if ok := d.HasChange("color"); ok {
		lsvProfile["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		lsvProfile["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		lsvProfile["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		lsvProfile["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Lsv Profile - Map = ", lsvProfile)

	updateLsvProfileRes, err := client.ApiCall("set-lsv-profile", lsvProfile, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateLsvProfileRes.Success {
		if updateLsvProfileRes.ErrorMsg != "" {
			return fmt.Errorf(updateLsvProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementLsvProfile(d, m)
}

func deleteManagementLsvProfile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	lsvProfilePayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_errors"); ok {
		lsvProfilePayload["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		lsvProfilePayload["ignore-warnings"] = v.(bool)
	}
	deleteLsvProfileRes, err := client.ApiCall("delete-lsv-profile", lsvProfilePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteLsvProfileRes.Success {
		if deleteLsvProfileRes.ErrorMsg != "" {
			return fmt.Errorf(deleteLsvProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId("")
	return nil
}
