package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementLsvProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementLsvProfileRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"allowed_ip_addresses": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of network objects identified by name or UID that represent IP addresses allowed in profile's VPN domain.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"certificate_authority": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trusted Certificate authority for establishing trust between VPN peers, identified by name or UID.",
			},
			"restrict_allowed_addresses": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicate whether the IP addresses allowed in the VPN Domain will be restricted or not, according to allowed-ip-addresses field.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpn_domain": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "peers' VPN Domain properties.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"limit_peer_domain_size": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use this parameter to limit the number of IP addresses in the VPN Domain of each peer according to the value in the max-allowed-addresses field.",
						},
						"max_allowed_addresses": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum number of IP addresses in the VPN Domain of each peer. This value will be enforced only when limit-peer-domain-size field is set to true. Select a value between 1 and 256. Default value is 256.",
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementLsvProfileRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showLsvProfileRes, err := client.ApiCall("show-lsv-profile", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLsvProfileRes.Success {
		return fmt.Errorf(showLsvProfileRes.ErrorMsg)
	}

	lsvProfile := showLsvProfileRes.GetData()

	log.Println("Read Lsv Profile - Show JSON = ", lsvProfile)

	if v := lsvProfile["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := lsvProfile["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if lsvProfile["allowed-ip-addresses"] != nil {
		allowedIpAddressesJson := lsvProfile["allowed-ip-addresses"].([]interface{})
		var allowedIpAddressesIds = make([]string, 0)
		if len(allowedIpAddressesJson) > 0 {
			// Create slice of tag names
			for _, allowedIpAddress := range allowedIpAddressesJson {
				allowedIpAddress := allowedIpAddress.(map[string]interface{})
				allowedIpAddressesIds = append(allowedIpAddressesIds, allowedIpAddress["name"].(string))
			}
		}
		_ = d.Set("allowed_ip_addresses", allowedIpAddressesIds)
	} else {
		_ = d.Set("allowed_ip_addresses", nil)
	}

	if v := lsvProfile["certificate-authority"]; v != nil {
		_ = d.Set("certificate_authority", v)
	}

	if v := lsvProfile["restrict_allowed_addresses"]; v != nil {
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

		_ = d.Set("vpn_domain", vpnDomainMapToReturn)
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
