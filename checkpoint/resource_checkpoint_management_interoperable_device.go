package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"math"

	"strconv"
)

func resourceManagementInteroperableDevice() *schema.Resource {
	return &schema.Resource{
		Create: createManagementInteroperableDevice,
		Read:   readManagementInteroperableDevice,
		Update: updateManagementInteroperableDevice,
		Delete: deleteManagementInteroperableDevice,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4 address of the Interoperable Device.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv6 address of the Interoperable Device.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 address.",
						},
						"ipv4_network_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 network address.",
						},
						"ipv6_network_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 network address.",
						},
						"ipv4_mask_length": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 network mask length.",
						},
						"ipv6_mask_length": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 network mask length.",
						},
						"tags": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Collection of tag identifiers.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"topology": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Topology configuration.",
							Default:     "internal",
						},
						"topology_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Internal topology settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address_behind_this_interface": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Network settings behind this interface.",
									},
									"interface_leads_to_dmz": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether this interface leads to demilitarized zone (perimeter network).",
									},
									"specific_network": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Network behind this interface.",
									},
								},
							},
						},
						"color": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Color of the object. Should be one of existing colors.",
							Default:     "black",
						},
						"comments": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Comments string.",
						},
						"domains_to_process": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ignore_warnings": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Apply changes ignoring warnings.",
							Default:     false,
						},
						"ignore_errors": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
							Default:     false,
						},
					},
				},
			},
			"vpn_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "VPN domain properties for the Interoperable Device.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vpn_domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Network group representing the customized encryption domain. Must be set when vpn-domain-type is set to 'manual' option.",
						},
						"vpn_domain_exclude_external_ip_addresses": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Exclude the external IP addresses from the VPN domain of this Interoperable device.",
						},
						"vpn_domain_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates the encryption domain.",
							Default:     "addresses_behind_gw",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"domains_to_process": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default:     false,
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
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
		},
	}
}

func createManagementInteroperableDevice(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	interoperableDevice := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		interoperableDevice["name"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		interoperableDevice["ipv4-address"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address"); ok {
		interoperableDevice["ipv6-address"] = v.(string)
	}

	if v, ok := d.GetOk("interfaces"); ok {

		interfacesList := v.([]interface{})

		if len(interfacesList) > 0 {
			var interfacesListToReturn []map[string]interface{}

			for i := range interfacesList {

				interfacesMap := interfacesList[i].(map[string]interface{})
				interfacesPayload := make(map[string]interface{})

				if v, ok := interfacesMap["name"]; ok {
					interfacesPayload["name"] = v.(string)
				}
				if v, ok := interfacesMap["ipv4_address"]; ok {
					interfacesPayload["ipv4-address"] = v.(string)
				}
				if v, ok := interfacesMap["ipv6_address"]; ok {
					interfacesPayload["ipv6-address"] = v.(string)
				}
				if v, ok := interfacesMap["ipv4_network_mask"]; ok {
					interfacesPayload["ipv4-network-mask"] = v.(string)
				}
				if v, ok := interfacesMap["ipv6_network_mask"]; ok {
					interfacesPayload["ipv6-network-mask"] = v.(string)
				}
				if v, ok := interfacesMap["ipv4_mask_length"]; ok {
					interfacesPayload["ipv4-mask-length"] = v.(string)
				}
				if v, ok := interfacesMap["ipv6_mask_length"]; ok {
					interfacesPayload["ipv6-mask-length"] = v.(string)
				}
				if v, ok := interfacesMap["tags"]; ok {
					interfacesPayload["tags"] = v.(*schema.Set).List()
				}
				if v, ok := interfacesMap["topology"]; ok {
					interfacesPayload["topology"] = v.(string)
				}
				if v, ok := interfacesMap["topology_settings"]; ok {

					topologySettingsList := v.([]interface{})
					if len(topologySettingsList) > 0 {
						topologySettingsMap := topologySettingsList[0].(map[string]interface{})
						topologySettingsPayload := make(map[string]interface{})

						if v, ok := topologySettingsMap["ip_address_behind_this_interface"]; ok && v != "" {
							topologySettingsPayload["ip-address-behind-this-interface"] = v.(string)
						}
						if v, ok := topologySettingsMap["interface_leads_to_dmz"]; ok {
							topologySettingsPayload["interface-leads-to-dmz"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := topologySettingsMap["specific_network"]; ok && v != "" {
							topologySettingsPayload["specific-network"] = v.(string)
						}
						interfacesPayload["topology-settings"] = topologySettingsPayload
					}
				}
				if v, ok := interfacesMap["color"]; ok {
					interfacesPayload["color"] = v.(string)
				}
				if v, ok := interfacesMap["comments"]; ok {
					interfacesPayload["comments"] = v.(string)
				}
				if v, ok := interfacesMap["domains_to_process"]; ok {
					interfacesPayload["domains-to-process"] = v.(*schema.Set).List()
				}
				if v, ok := interfacesMap["ignore_warnings"]; ok {
					interfacesPayload["ignore-warnings"] = v.(bool)
				}
				if v, ok := interfacesMap["ignore_errors"]; ok {
					interfacesPayload["ignore-errors"] = v.(bool)
				}
				interfacesListToReturn = append(interfacesListToReturn, interfacesPayload)
			}
			interoperableDevice["interfaces"] = interfacesListToReturn
		}
	}
	if _, ok := d.GetOk("vpn_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("vpn_settings.vpn_domain"); ok {
			res["vpn-domain"] = v.(string)
		}
		if v, ok := d.GetOk("vpn_settings.vpn_domain_exclude_external_ip_addresses"); ok {
			res["vpn-domain-exclude-external-ip-addresses"] = v
		}
		if v, ok := d.GetOk("vpn_settings.vpn_domain_type"); ok {
			res["vpn-domain-type"] = v.(string)
		}
		interoperableDevice["vpn-settings"] = res
	}

	if v, ok := d.GetOk("tags"); ok {
		interoperableDevice["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		interoperableDevice["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		interoperableDevice["comments"] = v.(string)
	}

	if v, ok := d.GetOk("domains_to_process"); ok {
		interoperableDevice["domains-to-process"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		interoperableDevice["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		interoperableDevice["ignore-errors"] = v.(bool)
	}

	log.Println("Create InteroperableDevice - Map = ", interoperableDevice)

	addInteroperableDeviceRes, err := client.ApiCall("add-interoperable-device", interoperableDevice, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addInteroperableDeviceRes.Success {
		if addInteroperableDeviceRes.ErrorMsg != "" {
			return fmt.Errorf(addInteroperableDeviceRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addInteroperableDeviceRes.GetData()["uid"].(string))

	return readManagementInteroperableDevice(d, m)
}

func readManagementInteroperableDevice(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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
							defaultVpnSettings := map[string]interface{}{
								"ip-address-behind-this-interface": "not defined",
								"interface-leads-to-dmz":           "false",
							}
							topologySettingsMapToReturn := make(map[string]interface{})

							if v, _ := topologySettingsMap["ip-address-behind-this-interface"]; v != nil && isArgDefault(v.(string), d, "interfaces."+strconv.Itoa(i)+".topology_settings.ip_address_behind_this_interface", defaultVpnSettings["ip-address-behind-this-interface"].(string)) {
								topologySettingsMapToReturn["ip_address_behind_this_interface"] = v
							}
							if v, _ := topologySettingsMap["interface-leads-to-dmz"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "interfaces."+strconv.Itoa(i)+".topology_settings.interface_leads_to_dmz", defaultVpnSettings["interface-leads-to-dmz"].(string)) {
								topologySettingsMapToReturn["interface_leads_to_dmz"] = v
							}
							if v, _ := topologySettingsMap["specific-network"]; v != nil {
								topologySettingsMapToReturn["specific_network"] = v
							}
							if len(topologySettingsMapToReturn) != 0 {
								interfacesMapToReturn["topology_settings"] = []interface{}{topologySettingsMapToReturn}
							} else {
								interfacesMapToReturn["topology_settings"] = []interface{}{}
							}
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
		defaultVpnSettings := map[string]interface{}{
			"vpn-domain-exclude-external-ip-addresses": "false",
			"vpn-domain-type":                          "addresses_behind_gw",
		}

		if v, _ := vpnSettingsMap["vpn-domain"]; v != nil {
			vpnSettingsMapToReturn["vpn_domain"] = v
		}
		if v := vpnSettingsMap["vpn-domain-exclude-external-ip-addresses"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "vpn_settings.vpn_domain_exclude_external_ip_addresses", defaultVpnSettings["vpn-domain-exclude-external-ip-addresses"].(string)) {
			vpnSettingsMapToReturn["vpn_domain_exclude_external_ip_addresses"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := vpnSettingsMap["vpn-domain-type"]; v != nil && isArgDefault(v.(string), d, "vpn_settings.vpn_domain_type", defaultVpnSettings["vpn-domain-type"].(string)) {
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
					domainsToProcess := domains_to_process.(map[string]interface{})
					domainsToProcessIds = append(domainsToProcessIds, domainsToProcess["name"].(string))
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

func updateManagementInteroperableDevice(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	interoperableDevice := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		interoperableDevice["name"] = oldName
		interoperableDevice["new-name"] = newName
	} else {
		interoperableDevice["name"] = d.Get("name")
	}

	if ok := d.HasChange("ipv4_address"); ok {
		interoperableDevice["ipv4-address"] = d.Get("ipv4_address")
	}

	if ok := d.HasChange("ipv6_address"); ok {
		interoperableDevice["ipv6-address"] = d.Get("ipv6_address")
	}

	if d.HasChange("interfaces") {
		log.Println("here0")
		if v, ok := d.GetOk("interfaces"); ok {
			log.Println("here1")

			interfacesList := v.([]interface{})

			if len(interfacesList) > 0 {

				var interfacesListToReturn []map[string]interface{}

				for i := range interfacesList {

					interfacesPayload := make(map[string]interface{})

					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".name"); ok {
						interfacesPayload["name"] = d.Get("interfaces." + strconv.Itoa(i) + ".name").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_address"); ok {
						interfacesPayload["ipv4-address"] = d.Get("interfaces." + strconv.Itoa(i) + ".ipv4_address").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_address"); ok {
						interfacesPayload["ipv6-address"] = d.Get("interfaces." + strconv.Itoa(i) + ".ipv6_address").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_network_mask"); ok {
						interfacesPayload["ipv4-network-mask"] = d.Get("interfaces." + strconv.Itoa(i) + ".ipv4_network_mask").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_network_mask"); ok {
						interfacesPayload["ipv6-network-mask"] = d.Get("interfaces." + strconv.Itoa(i) + ".ipv6_network_mask").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_mask_length"); ok {
						interfacesPayload["ipv4-mask-length"] = d.Get("interfaces." + strconv.Itoa(i) + ".ipv4_mask_length").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_mask_length"); ok {
						interfacesPayload["ipv6-mask-length"] = d.Get("interfaces." + strconv.Itoa(i) + ".ipv6_mask_length").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".tags"); ok {
						interfacesPayload["tags"] = d.Get("interfaces." + strconv.Itoa(i) + ".tags").(*schema.Set).List()
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology"); ok {
						interfacesPayload["topology"] = d.Get("interfaces." + strconv.Itoa(i) + ".topology").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings"); ok {

						topologySettingsPayload := make(map[string]interface{})

						if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.0.ip_address_behind_this_interface"); ok {
							topologySettingsPayload["ip-address-behind-this-interface"] = d.Get("interfaces." + strconv.Itoa(i) + ".topology_settings.0.ip_address_behind_this_interface")
						}
						if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.0.interface_leads_to_dmz"); ok {
							topologySettingsPayload["interface-leads-to-dmz"] = d.Get("interfaces." + strconv.Itoa(i) + ".topology_settings.0.interface_leads_to_dmz")
						}
						if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.0.specific_network"); ok {
							topologySettingsPayload["specific-network"] = d.Get("interfaces." + strconv.Itoa(i) + ".topology_settings.0.specific_network").(string)
						}
						interfacesPayload["topology-settings"] = topologySettingsPayload
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".color"); ok {
						interfacesPayload["color"] = d.Get("interfaces." + strconv.Itoa(i) + ".color").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".comments"); ok {
						interfacesPayload["comments"] = d.Get("interfaces." + strconv.Itoa(i) + ".comments").(string)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".domains_to_process"); ok {
						interfacesPayload["domains-to-process"] = d.Get("interfaces." + strconv.Itoa(i) + ".domains_to_process").(*schema.Set).List()
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ignore_warnings"); ok {
						interfacesPayload["ignore-warnings"] = d.Get("interfaces." + strconv.Itoa(i) + ".ignore_warnings").(bool)
					}
					if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ignore_errors"); ok {
						interfacesPayload["ignore-errors"] = d.Get("interfaces." + strconv.Itoa(i) + ".ignore_errors").(bool)
					}
					interfacesListToReturn = append(interfacesListToReturn, interfacesPayload)
				}
				interoperableDevice["interfaces"] = interfacesListToReturn
			}
		}
	}

	if d.HasChange("vpn_settings") {

		if _, ok := d.GetOk("vpn_settings"); ok {

			res := make(map[string]interface{})

			if d.HasChange("vpn_settings.vpn_domain") {
				res["vpn-domain"] = d.Get("vpn_settings.vpn_domain")
			}
			if d.HasChange("vpn_settings.vpn_domain_exclude_external_ip_addresses") {
				res["vpn-domain-exclude-external-ip-addresses"] = d.Get("vpn_settings.vpn_domain_exclude_external_ip_addresses")
			}
			if d.HasChange("vpn_settings.vpn_domain_type") {
				res["vpn-domain-type"] = d.Get("vpn_settings.vpn_domain_type")
			}
			interoperableDevice["vpn-settings"] = res
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			interoperableDevice["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			interoperableDevice["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		interoperableDevice["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		interoperableDevice["comments"] = d.Get("comments")
	}

	if d.HasChange("domains_to_process") {
		if v, ok := d.GetOk("domains_to_process"); ok {
			interoperableDevice["domains_to_process"] = v.(*schema.Set).List()
		} else {
			oldDomains_To_Process, _ := d.GetChange("domains_to_process")
			interoperableDevice["domains_to_process"] = map[string]interface{}{"remove": oldDomains_To_Process.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		interoperableDevice["ignore-warnings"] = v.(bool)
	}

	if d.HasChange("groups") {
		if v, ok := d.GetOk("groups"); ok {
			interoperableDevice["groups"] = v.(*schema.Set).List()
		} else {
			oldGroups, _ := d.GetChange("groups")
			interoperableDevice["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		interoperableDevice["ignore-errors"] = v.(bool)
	}

	log.Println("Update InteroperableDevice - Map = ", interoperableDevice)

	updateInteroperableDeviceRes, err := client.ApiCall("set-interoperable-device", interoperableDevice, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateInteroperableDeviceRes.Success {
		if updateInteroperableDeviceRes.ErrorMsg != "" {
			return fmt.Errorf(updateInteroperableDeviceRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementInteroperableDevice(d, m)
}

func deleteManagementInteroperableDevice(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	interoperableDevicePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		interoperableDevicePayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		interoperableDevicePayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete InteroperableDevice")

	deleteInteroperableDeviceRes, err := client.ApiCall("delete-interoperable-device", interoperableDevicePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteInteroperableDeviceRes.Success {
		if deleteInteroperableDeviceRes.ErrorMsg != "" {
			return fmt.Errorf(deleteInteroperableDeviceRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
