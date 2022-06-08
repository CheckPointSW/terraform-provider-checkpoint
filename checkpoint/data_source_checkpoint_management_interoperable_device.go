package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"math"
	"strconv"
)

func dataSourceManagementInteroperableDevice() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementInteroperableDeviceRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 address of the Interoperable Device.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 address of the Interoperable Device.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address.",
						},
						"ipv4_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network address.",
						},
						"ipv6_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network address.",
						},
						"ipv4_mask_length": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network mask length.",
						},
						"ipv6_mask_length": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network mask length.",
						},
						"tags": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Collection of tag identifiers.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"topology": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topology configuration.",
						},
						"topology_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Internal topology settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address_behind_this_interface": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network settings behind this interface.",
									},
									"interface_leads_to_dmz": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether this interface leads to demilitarized zone (perimeter network).",
									},
									"specific_network": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network behind this interface.",
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
						"domains_to_process": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ignore_warnings": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Apply changes ignoring warnings.",
						},
						"ignore_errors": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
						},
					},
				},
			},
			"vpn_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "VPN domain properties for the Interoperable Device.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpn_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network group representing the customized encryption domain. Must be set when vpn-domain-type is set to 'manual' option.",
						},
						"vpn_domain_exclude_external_ip_addresses": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Exclude the external IP addresses from the VPN domain of this Interoperable device.",
						},
						"vpn_domain_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates the encryption domain.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"domains_to_process": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"groups": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
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
						"color": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Color of the object. Should be one of existing colors.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementInteroperableDeviceRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showInteroperableDeviceRes, err := client.ApiCall("show-interoperable-device", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showInteroperableDeviceRes.Success {
		if objectNotFound(showInteroperableDeviceRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showInteroperableDeviceRes.ErrorMsg)
	}

	interoperableDevice := showInteroperableDeviceRes.GetData()

	log.Println("Read InteroperableDevice - Show JSON = ", interoperableDevice)

	if v := interoperableDevice["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := interoperableDevice["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := interoperableDevice["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := interoperableDevice["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if interoperableDevice["interfaces"] != nil {

		interfacesList, ok := interoperableDevice["interfaces"].([]interface{})

		if ok {

			if len(interfacesList) > 0 {

				var interfacesListToReturn []map[string]interface{}

				for i := range interfacesList {

					interfacesMap := interfacesList[i].(map[string]interface{})

					interfacesMapToReturn := make(map[string]interface{})

					if v := interfacesMap["name"]; v != nil {
						interfacesMapToReturn["name"] = v
					}
					if v := interfacesMap["ipv4-address"]; v != nil {
						interfacesMapToReturn["ipv4_address"] = v
					}
					if v := interfacesMap["ipv6-address"]; v != nil {
						interfacesMapToReturn["ipv6_address"] = v
					}
					if v := interfacesMap["ipv4-network-mask"]; v != nil {
						interfacesMapToReturn["ipv4_network_mask"] = v
					}
					if v := interfacesMap["ipv6-network-mask"]; v != nil {
						interfacesMapToReturn["ipv6_network_mask"] = v
					}
					if v := interfacesMap["ipv4-mask-length"]; v != nil {
						interfacesMapToReturn["ipv4_mask_length"] = strconv.Itoa(int(math.Round(v.(float64))))
					}
					if v := interfacesMap["ipv6-mask-length"]; v != nil {
						interfacesMapToReturn["ipv6_mask_length"] = strconv.Itoa(int(math.Round(v.(float64))))
					}
					if v := interfacesMap["tags"]; v != nil {
						interfacesMapToReturn["tags"] = v
					}
					if v := interfacesMap["topology"]; v != nil {
						interfacesMapToReturn["topology"] = v
					}
					if v, ok := interfacesMap["topology-settings"]; ok {

						topologySettingsMap, ok := v.(map[string]interface{})
						if ok {
							topologySettingsMapToReturn := make(map[string]interface{})

							if v, _ := topologySettingsMap["ip-address-behind-this-interface"]; v != nil {
								topologySettingsMapToReturn["ip_address_behind_this_interface"] = v
							}
							if v, _ := topologySettingsMap["interface-leads-to-dmz"]; v != nil {
								topologySettingsMapToReturn["interface_leads_to_dmz"] = v
							}
							if v, _ := topologySettingsMap["specific-network"]; v != nil {
								topologySettingsMapToReturn["specific_network"] = v
							}
							interfacesMapToReturn["topology_settings"] = []interface{}{topologySettingsMapToReturn}
						}
					}
					if v := interfacesMap["color"]; v != nil {
						interfacesMapToReturn["color"] = v
					}
					if v := interfacesMap["comments"]; v != nil {
						interfacesMapToReturn["comments"] = v
					}
					if v := interfacesMap["domains-to-process"]; v != nil {
						interfacesMapToReturn["domains_to_process"] = v
					}
					interfacesListToReturn = append(interfacesListToReturn, interfacesMapToReturn)
				}
				_ = d.Set("interfaces", interfacesListToReturn)

			}
		}
	} else {
		_ = d.Set("interfaces", nil)
	}

	if interoperableDevice["vpn-settings"] != nil {

		vpnSettingsMap := interoperableDevice["vpn-settings"].(map[string]interface{})

		vpnSettingsMapToReturn := make(map[string]interface{})
		if v, _ := vpnSettingsMap["vpn-domain"]; v != nil {
			vpnSettingsMapToReturn["vpn_domain"] = v
		}
		if v := vpnSettingsMap["vpn-domain-exclude-external-ip-addresses"]; v != nil {
			vpnSettingsMapToReturn["vpn_domain_exclude_external_ip_addresses"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := vpnSettingsMap["vpn-domain-type"]; v != nil {
			vpnSettingsMapToReturn["vpn_domain_type"] = v.(string)
		}
		_ = d.Set("vpn_settings", vpnSettingsMapToReturn)
	} else {
		_ = d.Set("vpn_settings", nil)
	}

	if interoperableDevice["tags"] != nil {
		tagsJson, ok := interoperableDevice["tags"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("tags", tagsIds)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if v := interoperableDevice["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := interoperableDevice["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if interoperableDevice["domains_to_process"] != nil {
		domainsToProcessJson, ok := interoperableDevice["domains_to_process"].([]interface{})
		if ok {
			domainsToProcessIds := make([]string, 0)
			if len(domainsToProcessJson) > 0 {
				for _, domains_to_process := range domainsToProcessJson {
					domains_to_process := domains_to_process.(map[string]interface{})
					domainsToProcessIds = append(domainsToProcessIds, domains_to_process["name"].(string))
				}
			}
			_ = d.Set("domains_to_process", domainsToProcessIds)
		}
	} else {
		_ = d.Set("domains_to_process", nil)
	}

	if v := interoperableDevice["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if interoperableDevice["groups"] != nil {

		interfacesList, ok := interoperableDevice["groups"].([]interface{})

		var interfacesListToReturn []map[string]interface{}

		if ok {

			if len(interfacesList) > 0 {

				for i := range interfacesList {

					interfacesMap := interfacesList[i].(map[string]interface{})

					interfacesMapToAdd := make(map[string]interface{})

					if v, _ := interfacesMap["name"]; v != nil {
						interfacesMapToAdd["name"] = v
					}
					if v, _ := interfacesMap["uid"]; v != nil {
						interfacesMapToAdd["uid"] = v
					}
					if v, _ := interfacesMap["type"]; v != nil {
						interfacesMapToAdd["type"] = v
					}
					if v, _ := interfacesMap["color"]; v != nil {
						interfacesMapToAdd["color"] = v
					}
					interfacesListToReturn = append(interfacesListToReturn, interfacesMapToAdd)
				}
			}
		}
		_ = d.Set("groups", interfacesListToReturn)
	}

	if v := interoperableDevice["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil
}
