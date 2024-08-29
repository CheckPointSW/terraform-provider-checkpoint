package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementInterface() *schema.Resource {
	return &schema.Resource{
		Create: createManagementInterface,
		Read:   readManagementInterface,
		Update: updateManagementInterface,
		Delete: deleteManagementInterface,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Network interface name.",
			},
			"gateway_uid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Gateway or cluster object uid that the interface belongs to. <font color=\"red\">Required only if</font> name was specified.",
			},

			"anti_spoofing": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable anti-spoofing.",
				Default:     false,
			},
			"anti_spoofing_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Anti Spoofing Settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).",
							Default:     "prevent",
						},
						"exclude_packets": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Don't check packets from excluded network.",
						},
						"excluded_network_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Excluded network name.",
						},
						"excluded_network_uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Excluded network UID.",
						},
						"spoof_tracking": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Spoof tracking.",
							Default:     "log",
						},
					},
				},
			},
			"cluster_members": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Network interface settings for cluster members.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Cluster member network interface name.",
						},
						"member_uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cluster member object uid.",
						},
						"member_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Cluster member object name.",
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
						"network_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly.",
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
					},
				},
			},
			"cluster_network_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster interface type.",
			},
			"dynamic_ip": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable dynamic interface.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4 network address.",
			},
			"ipv4_mask_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "IPv4 mask length.",
			},
			"ipv4_network_mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4 network mask.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv6 address.",
			},
			"ipv6_mask_length": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "IPv6 mask length.",
			},
			"ipv6_network_mask": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv6 network mask.",
			},
			"monitored_by_cluster": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When Private is selected as the Cluster interface type, cluster can monitor or not monitor the interface.",
				Default:     false,
			},
			"network_interface_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Network Interface Type.",
			},
			"security_zone_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Security Zone Settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_calculated": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Security Zone is calculated according to where the interface leads to.",
						},
						"specific_zone": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Security Zone specified manually.",
						},
						"specific_zone_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Security Zone specified manually.",
						},
						"auto_calculated_zone": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"auto_calculated_zone_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"specific_security_zone_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "N/A",
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
			"topology": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Topology configuration.",
			},
			"topology_automatic": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Topology configuration automatically calculated by get-interfaces command.",
			},
			"topology_manual": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Topology configuration manually defined.",
			},
			"topology_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Topology Settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_leads_to_dmz": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether this interface leads to demilitarized zone (perimeter network).",
						},
						"ip_address_behind_this_interface": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Network Settings behind this interface.",
							Default:     "not defined",
						},
						"specific_network": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Network behind this interface.",
						},
						"specific_network_uid": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
					},
				},
			},
			"topology_settings_automatic": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Topology Settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_leads_to_dmz": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this interface leads to demilitarized zone (perimeter network).",
						},
						"ip_address_behind_this_interface": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network Settings behind this interface.",
						},
						"specific_network": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network behind this interface.",
						},
						"specific_network_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
					},
				},
			},
			"topology_settings_manual": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Topology Settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interface_leads_to_dmz": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether this interface leads to demilitarized zone (perimeter network).",
						},
						"ip_address_behind_this_interface": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network Settings behind this interface.",
						},
						"specific_network": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Network behind this interface.",
						},
						"specific_network_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
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
	}
}

func createManagementInterface(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	interfaceMap := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		interfaceMap["name"] = v.(string)
	}

	if v, ok := d.GetOk("gateway_uid"); ok {
		interfaceMap["gateway-uid"] = v.(string)
	}

	if v, ok := d.GetOkExists("anti_spoofing"); ok {
		interfaceMap["anti-spoofing"] = v.(bool)
	}

	if _, ok := d.GetOk("anti_spoofing_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("anti_spoofing_settings.0.action"); ok {
			res["action"] = v
		}
		if v, ok := d.GetOk("anti_spoofing_settings.0.exclude_packets"); ok {
			res["exclude-packets"] = v
		}
		if v, ok := d.GetOk("anti_spoofing_settings.0.excluded_network_name"); ok {
			res["excluded-network-name"] = v
		}
		if v, ok := d.GetOk("anti_spoofing_settings.0.excluded_network_uid"); ok {
			res["excluded-network-uid"] = v
		}
		if v, ok := d.GetOk("anti_spoofing_settings.0.spoof_tracking"); ok {
			res["spoof-tracking"] = v
		}
		interfaceMap["anti-spoofing-settings"] = res
	}

	if v, ok := d.GetOk("cluster_members"); ok {

		clusterMembersList := v.([]interface{})

		if len(clusterMembersList) > 0 {

			var clusterMembersPayload []map[string]interface{}

			for i := range clusterMembersList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".name"); ok {
					Payload["name"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".member_uid"); ok {
					Payload["member-uid"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".member_name"); ok {
					Payload["member-name"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv4_address"); ok {
					Payload["ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv6_address"); ok {
					Payload["ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".network_mask"); ok {
					Payload["network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv4_network_mask"); ok {
					Payload["ipv4-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv6_network_mask"); ok {
					Payload["ipv6-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv4_mask_length"); ok {
					Payload["ipv4-mask-length"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv6_mask_length"); ok {
					Payload["ipv6-mask-length"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".tags"); ok {
					Payload["tags"] = v
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".color"); ok {
					Payload["color"] = v.(string)
				}
				if v, ok := d.GetOk("cluster_members." + strconv.Itoa(i) + ".comments"); ok {
					Payload["comments"] = v.(string)
				}
				clusterMembersPayload = append(clusterMembersPayload, Payload)
			}
			interfaceMap["cluster-members"] = clusterMembersPayload
		}
	}

	if v, ok := d.GetOk("cluster_network_type"); ok {
		interfaceMap["cluster-network-type"] = v.(string)
	}

	if v, ok := d.GetOkExists("dynamic_ip"); ok {
		interfaceMap["dynamic-ip"] = v.(bool)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		interfaceMap["ipv4-address"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_mask_length"); ok {
		interfaceMap["ipv4-mask-length"] = v.(int)
	}

	if v, ok := d.GetOk("ipv4_network_mask"); ok {
		interfaceMap["ipv4-network-mask"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address"); ok {
		interfaceMap["ipv6-address"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_mask_length"); ok {
		interfaceMap["ipv6-mask-length"] = v.(int)
	}

	if v, ok := d.GetOk("ipv6_network_mask"); ok {
		interfaceMap["ipv6-network-mask"] = v.(string)
	}

	if v, ok := d.GetOkExists("monitored_by_cluster"); ok {
		interfaceMap["monitored-by-cluster"] = v.(bool)
	}

	if v, ok := d.GetOk("network_interface_type"); ok {
		interfaceMap["network-interface-type"] = v.(string)
	}

	if _, ok := d.GetOk("security_zone_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("security_zone_settings.0.auto_calculated"); ok {
			res["auto-calculated"] = v
		}
		if v, ok := d.GetOk("security_zone_settings.0.specific_zone"); ok {
			res["specific-zone"] = v
		}
		interfaceMap["security-zone-settings"] = res
	}

	if v, ok := d.GetOk("tags"); ok {
		interfaceMap["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("topology"); ok {
		interfaceMap["topology"] = v.(string)
	}

	if _, ok := d.GetOk("topology_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("topology_settings.0.interface_leads_to_dmz"); ok {
			res["interface-leads-to-dmz"] = v
		}
		if v, ok := d.GetOk("topology_settings.0.ip_address_behind_this_interface"); ok {
			res["ip-address-behind-this-interface"] = v
		}
		if v, ok := d.GetOk("topology_settings.0.specific_network"); ok {
			res["specific-network"] = v
		}
		if v, ok := d.GetOk("topology_settings.0.specific_network_uid"); ok {
			res["specific-network-uid"] = v
		}
		interfaceMap["topology-settings"] = res
	}

	if v, ok := d.GetOk("color"); ok {
		interfaceMap["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		interfaceMap["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		interfaceMap["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		interfaceMap["ignore-errors"] = v.(bool)
	}

	log.Println("Create Interface - Map = ", interfaceMap)

	addInterfaceRes, err := client.ApiCall("add-interface", interfaceMap, client.GetSessionID(), true, false)
	if err != nil || !addInterfaceRes.Success {
		if addInterfaceRes.ErrorMsg != "" {
			return fmt.Errorf(addInterfaceRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addInterfaceRes.GetData()["uid"].(string))

	return readManagementInterface(d, m)
}

func readManagementInterface(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showInterfaceRes, err := client.ApiCall("show-interface", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showInterfaceRes.Success {
		if objectNotFound(showInterfaceRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showInterfaceRes.ErrorMsg)
	}

	interfaceMap := showInterfaceRes.GetData()

	log.Println("Read Interface - Show JSON = ", interfaceMap)

	if v := interfaceMap["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := interfaceMap["gateway"]; v != nil {

		gwObj := v.(map[string]interface{})
		if v := gwObj["uid"]; v != nil {
			_ = d.Set("gateway_uid", v)
		}

	}

	if v := interfaceMap["anti-spoofing"]; v != nil {
		_ = d.Set("anti_spoofing", v)
	}

	if interfaceMap["anti-spoofing-settings"] != nil {

		antiSpoofingSettingsMap := interfaceMap["anti-spoofing-settings"].(map[string]interface{})

		antiSpoofingSettingsMapToReturn := make(map[string]interface{})

		if v, _ := antiSpoofingSettingsMap["action"]; v != nil {
			antiSpoofingSettingsMapToReturn["action"] = v
		}
		if v, _ := antiSpoofingSettingsMap["exclude-packets"]; v != nil {
			antiSpoofingSettingsMapToReturn["exclude_packets"] = v
		}
		if v, _ := antiSpoofingSettingsMap["excluded-network-name"]; v != nil {
			antiSpoofingSettingsMapToReturn["excluded_network_name"] = v
		}
		if v, _ := antiSpoofingSettingsMap["excluded-network-uid"]; v != nil {
			antiSpoofingSettingsMapToReturn["excluded_network_uid"] = v
		}
		if v, _ := antiSpoofingSettingsMap["spoof-tracking"]; v != nil {
			antiSpoofingSettingsMapToReturn["spoof_tracking"] = v
		}
		_ = d.Set("anti_spoofing_settings", []interface{}{antiSpoofingSettingsMapToReturn})
	} else {
		_ = d.Set("anti_spoofing_settings", nil)
	}

	if interfaceMap["cluster-members"] != nil {

		clusterMembersList, ok := interfaceMap["cluster-members"].([]interface{})

		if ok {

			if len(clusterMembersList) > 0 {

				var clusterMembersListToReturn []map[string]interface{}

				for i := range clusterMembersList {

					clusterMembersMap := clusterMembersList[i].(map[string]interface{})

					clusterMembersMapToAdd := make(map[string]interface{})

					if v, _ := clusterMembersMap["name"]; v != nil {
						clusterMembersMapToAdd["name"] = v
					}
					if v, _ := clusterMembersMap["member-uid"]; v != nil {
						clusterMembersMapToAdd["member_uid"] = v
					}
					if v, _ := clusterMembersMap["member-name"]; v != nil {
						clusterMembersMapToAdd["member_name"] = v
					}
					if v, _ := clusterMembersMap["ipv4-address"]; v != nil {
						clusterMembersMapToAdd["ipv4_address"] = v
					}
					if v, _ := clusterMembersMap["ipv6-address"]; v != nil {
						clusterMembersMapToAdd["ipv6_address"] = v
					}
					if v, _ := clusterMembersMap["network-mask"]; v != nil {
						clusterMembersMapToAdd["network_mask"] = v
					}
					if v, _ := clusterMembersMap["ipv4-network-mask"]; v != nil {
						clusterMembersMapToAdd["ipv4_network_mask"] = v
					}
					if v, _ := clusterMembersMap["ipv6-network-mask"]; v != nil {
						clusterMembersMapToAdd["ipv6_network_mask"] = v
					}
					if v, _ := clusterMembersMap["ipv4-mask-length"]; v != nil {
						clusterMembersMapToAdd["ipv4_mask_length"] = v
					}
					if v, _ := clusterMembersMap["ipv6-mask-length"]; v != nil {
						clusterMembersMapToAdd["ipv6_mask_length"] = v
					}
					clusterMembersListToReturn = append(clusterMembersListToReturn, clusterMembersMapToAdd)
				}
				_ = d.Set("cluster_members", clusterMembersListToReturn)
			}
		}
	}

	if v := interfaceMap["cluster-network-type"]; v != nil {
		_ = d.Set("cluster_network_type", v)
	}

	if v := interfaceMap["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := interfaceMap["ipv4-mask-length"]; v != nil {
		_ = d.Set("ipv4_mask_length", v)
	}

	if v := interfaceMap["ipv4-network-mask"]; v != nil {
		_ = d.Set("ipv4_network_mask", v)
	}

	if v := interfaceMap["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := interfaceMap["ipv6-mask-length"]; v != nil {
		_ = d.Set("ipv6_mask_length", v)
	}

	if v := interfaceMap["ipv6-network-mask"]; v != nil {
		_ = d.Set("ipv6_network_mask", v)
	}

	if v := interfaceMap["monitored-by-cluster"]; v != nil {
		_ = d.Set("monitored_by_cluster", v)
	}

	if v := interfaceMap["network-interface-type"]; v != nil {
		_ = d.Set("network_interface_type", v)
	}

	if interfaceMap["security-zone-settings"] != nil {

		securityZoneSettingsMap := interfaceMap["security-zone-settings"].(map[string]interface{})

		securityZoneSettingsMapToReturn := make(map[string]interface{})

		if v, _ := securityZoneSettingsMap["auto-calculated"]; v != nil {
			securityZoneSettingsMapToReturn["auto_calculated"] = v
		}
		if v, _ := securityZoneSettingsMap["specific-zone"]; v != nil {
			securityZoneSettingsMapToReturn["specific_zone"] = v
		}
		if v, _ := securityZoneSettingsMap["specific-zone-uid"]; v != nil {
			securityZoneSettingsMapToReturn["specific_zone_uid"] = v
		}
		if v, _ := securityZoneSettingsMap["auto-calculated-zone"]; v != nil {
			securityZoneSettingsMapToReturn["auto_calculated_zone"] = v
		}
		if v, _ := securityZoneSettingsMap["auto-calculated-zone-uid"]; v != nil {
			securityZoneSettingsMapToReturn["auto_calculated_zone_uid"] = v
		}
		if v, _ := securityZoneSettingsMap["specific-security-zone-enabled"]; v != nil {
			securityZoneSettingsMapToReturn["specific_security_zone_enabled"] = v
		}
		_ = d.Set("security_zone_settings", []interface{}{securityZoneSettingsMapToReturn})
	} else {
		_ = d.Set("security_zone_settings", nil)
	}

	if interfaceMap["tags"] != nil {
		tagsJson, ok := interfaceMap["tags"].([]interface{})
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

	if v := interfaceMap["topology"]; v != nil {
		_ = d.Set("topology", v)
	}
	if v := interfaceMap["topology-automatic"]; v != nil {
		_ = d.Set("topology_automatic", v)
	}
	if v := interfaceMap["topology-settings"]; v != nil {
		_ = d.Set("topology_settings", v)
	}

	if interfaceMap["topology-settings"] != nil {

		topologySettingsMap := interfaceMap["topology-settings"].(map[string]interface{})

		topologySettingsMapToReturn := make(map[string]interface{})

		if v, _ := topologySettingsMap["interface-leads-to-dmz"]; v != nil {
			topologySettingsMapToReturn["interface_leads_to_dmz"] = v
		}
		if v, _ := topologySettingsMap["ip-address-behind-this-interface"]; v != nil {
			topologySettingsMapToReturn["ip_address_behind_this_interface"] = v
		}
		if v, _ := topologySettingsMap["specific-network"]; v != nil {
			topologySettingsMapToReturn["specific_network"] = v
		}
		if v, _ := topologySettingsMap["specific-network-uid"]; v != nil {
			topologySettingsMapToReturn["specific_network_uid"] = v
		}
		_ = d.Set("topology_settings", []interface{}{topologySettingsMapToReturn})
	} else {
		_ = d.Set("topology_settings", nil)
	}

	if interfaceMap["topology-settings-automatic"] != nil {

		topologySettingsMap := interfaceMap["topology-settings-automatic"].(map[string]interface{})

		topologySettingsMapToReturn := make(map[string]interface{})

		if v, _ := topologySettingsMap["interface-leads-to-dmz"]; v != nil {
			topologySettingsMapToReturn["interface_leads_to_dmz"] = v
		}
		if v, _ := topologySettingsMap["ip-address-behind-this-interface"]; v != nil {
			topologySettingsMapToReturn["ip_address_behind_this_interface"] = v
		}
		if v, _ := topologySettingsMap["specific-network"]; v != nil {
			topologySettingsMapToReturn["specific_network"] = v
		}
		if v, _ := topologySettingsMap["specific-network-uid"]; v != nil {
			topologySettingsMapToReturn["specific_network_uid"] = v
		}
		_ = d.Set("topology_settings_automatic", []interface{}{topologySettingsMapToReturn})
	} else {
		_ = d.Set("topology_settings_automatic", nil)
	}

	if interfaceMap["topology-settings-manual"] != nil {

		topologySettingsMap := interfaceMap["topology-settings-manual"].(map[string]interface{})

		topologySettingsMapToReturn := make(map[string]interface{})

		if v, _ := topologySettingsMap["interface-leads-to-dmz"]; v != nil {
			topologySettingsMapToReturn["interface_leads_to_dmz"] = v
		}
		if v, _ := topologySettingsMap["ip-address-behind-this-interface"]; v != nil {
			topologySettingsMapToReturn["ip_address_behind_this_interface"] = v
		}
		if v, _ := topologySettingsMap["specific-network"]; v != nil {
			topologySettingsMapToReturn["specific_network"] = v
		}
		if v, _ := topologySettingsMap["specific-network-uid"]; v != nil {
			topologySettingsMapToReturn["specific_network_uid"] = v
		}
		_ = d.Set("topology_settings_manual", []interface{}{topologySettingsMapToReturn})
	} else {
		_ = d.Set("topology_settings_manual", nil)
	}

	if v := interfaceMap["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := interfaceMap["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := interfaceMap["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := interfaceMap["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementInterface(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	interfaceMap := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		interfaceMap["name"] = oldName
		interfaceMap["new-name"] = newName
	} else {
		interfaceMap["name"] = d.Get("name")
	}

	if _, ok := d.GetOk("gateway_uid"); ok {
		interfaceMap["gateway-uid"] = d.Get("gateway_uid")
	}

	if v, ok := d.GetOkExists("anti_spoofing"); ok {
		interfaceMap["anti-spoofing"] = v.(bool)
	}

	if d.HasChange("anti_spoofing_settings") {

		if _, ok := d.GetOk("anti_spoofing_settings"); ok {

			res := make(map[string]interface{})

			if d.GetOk("anti_spoofing_settings.0.action"); ok {
				res["action"] = d.Get("anti_spoofing_settings.0.action")
			}
			if d.GetOk("anti_spoofing_settings.0.exclude_packets"); ok {
				res["exclude-packets"] = d.Get("anti_spoofing_settings.0.exclude_packets")
			}
			if d.GetOk("anti_spoofing_settings.0.excluded_network_name"); ok {
				res["excluded-network-name"] = d.Get("anti_spoofing_settings.0.excluded_network_name")
			}
			if d.GetOk("anti_spoofing_settings.0.excluded_network_uid"); ok {
				res["excluded-network-uid"] = d.Get("anti_spoofing_settings.0.excluded_network_uid")
			}
			if d.GetOk("anti_spoofing_settings.0.spoof_tracking"); ok {
				res["spoof-tracking"] = d.Get("anti_spoofing_settings.0.spoof_tracking")
			}
			interfaceMap["anti-spoofing-settings"] = res
		}
	}

	if d.HasChange("cluster_members") {

		if v, ok := d.GetOk("cluster_members"); ok {

			clusterMembersList := v.([]interface{})

			var clusterMembersPayload []map[string]interface{}

			for i := range clusterMembersList {

				Payload := make(map[string]interface{})

				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".name"); ok {
					Payload["name"] = d.Get("cluster_members." + strconv.Itoa(i) + ".name")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".member_uid"); ok {
					Payload["member-uid"] = d.Get("cluster_members." + strconv.Itoa(i) + ".member_uid")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".member_name"); ok {
					Payload["member-name"] = d.Get("cluster_members." + strconv.Itoa(i) + ".member_name")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv4_address"); ok {
					Payload["ipv4-address"] = d.Get("cluster_members." + strconv.Itoa(i) + ".ipv4_address")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv6_address"); ok {
					Payload["ipv6-address"] = d.Get("cluster_members." + strconv.Itoa(i) + ".ipv6_address")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".network_mask"); ok {
					Payload["network-mask"] = d.Get("cluster_members." + strconv.Itoa(i) + ".network_mask")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv4_network_mask"); ok {
					Payload["ipv4-network-mask"] = d.Get("cluster_members." + strconv.Itoa(i) + ".ipv4_network_mask")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv6_network_mask"); ok {
					Payload["ipv6-network-mask"] = d.Get("cluster_members." + strconv.Itoa(i) + ".ipv6_network_mask")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv4_mask_length"); ok {
					Payload["ipv4-mask-length"] = d.Get("cluster_members." + strconv.Itoa(i) + ".ipv4_mask_length")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".ipv6_mask_length"); ok {
					Payload["ipv6-mask-length"] = d.Get("cluster_members." + strconv.Itoa(i) + ".ipv6_mask_length")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".tags"); ok {
					Payload["tags"] = d.Get("cluster_members." + strconv.Itoa(i) + ".tags")
				}
				if d.GetOk("cluster_members." + strconv.Itoa(i) + ".color"); ok {
					Payload["color"] = d.Get("cluster_members." + strconv.Itoa(i) + ".color")
				}
				if d.HasChange("cluster_members." + strconv.Itoa(i) + ".comments") {
					Payload["comments"] = d.Get("cluster_members." + strconv.Itoa(i) + ".comments")
				}

				clusterMembersPayload = append(clusterMembersPayload, Payload)
			}
			interfaceMap["cluster-members"] = clusterMembersPayload
		} else {
			oldclusterMembers, _ := d.GetChange("cluster_members")
			var clusterMembersToDelete []interface{}
			for _, i := range oldclusterMembers.([]interface{}) {
				clusterMembersToDelete = append(clusterMembersToDelete, i.(map[string]interface{})["name"].(string))
			}
			interfaceMap["cluster-members"] = map[string]interface{}{"remove": clusterMembersToDelete}
		}
	}

	if ok := d.HasChange("cluster_network_type"); ok {
		interfaceMap["cluster-network-type"] = d.Get("cluster_network_type")
	}

	if v, ok := d.GetOkExists("dynamic_ip"); ok {
		interfaceMap["dynamic-ip"] = v.(bool)
	}

	if ok := d.HasChange("ipv4_address"); ok {
		interfaceMap["ipv4-address"] = d.Get("ipv4_address")
	}

	if ok := d.HasChange("ipv4_mask_length"); ok {
		interfaceMap["ipv4-mask-length"] = d.Get("ipv4_mask_length")
	}

	if ok := d.HasChange("ipv4_network_mask"); ok {
		interfaceMap["ipv4-network-mask"] = d.Get("ipv4_network_mask")
	}

	if ok := d.HasChange("ipv6_address"); ok {
		interfaceMap["ipv6-address"] = d.Get("ipv6_address")
	}

	if ok := d.HasChange("ipv6_mask_length"); ok {
		interfaceMap["ipv6-mask-length"] = d.Get("ipv6_mask_length")
	}

	if ok := d.HasChange("ipv6_network_mask"); ok {
		interfaceMap["ipv6-network-mask"] = d.Get("ipv6_network_mask")
	}

	if v, ok := d.GetOkExists("monitored_by_cluster"); ok {
		interfaceMap["monitored-by-cluster"] = v.(bool)
	}

	if ok := d.HasChange("network_interface_type"); ok {
		interfaceMap["network-interface-type"] = d.Get("network_interface_type")
	}

	if d.HasChange("security_zone_settings") {

		if _, ok := d.GetOk("security_zone_settings"); ok {

			res := make(map[string]interface{})

			if d.GetOk("security_zone_settings.0.auto_calculated"); ok {
				res["auto-calculated"] = d.Get("security_zone_settings.0.auto_calculated")
			}
			if d.GetOk("security_zone_settings.0.specific_zone"); ok {
				res["specific-zone"] = d.Get("security_zone_settings.0.specific_zone")
			}

			interfaceMap["security-zone-settings"] = res
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			interfaceMap["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			interfaceMap["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("topology"); ok {
		interfaceMap["topology"] = d.Get("topology")
	}

	if d.HasChange("topology_settings") {

		if _, ok := d.GetOk("topology_settings"); ok {

			res := make(map[string]interface{})

			if d.GetOk("topology_settings.0.interface_leads_to_dmz"); ok {
				res["interface-leads-to-dmz"] = d.Get("topology_settings.0.interface_leads_to_dmz")
			}
			if d.GetOk("topology_settings.0.ip_address_behind_this_interface"); ok {
				res["ip-address-behind-this-interface"] = d.Get("topology_settings.0.ip_address_behind_this_interface")
			}
			if d.GetOk("topology_settings.0.specific_network"); ok {
				res["specific-network"] = d.Get("topology_settings.0.specific_network")
			}
			if d.GetOk("topology_settings.0.specific_network_uid"); ok {
				res["specific-network-uid"] = d.Get("topology_settings.0.specific_network_uid")
			}
			interfaceMap["topology-settings"] = res
		}
	}

	if ok := d.HasChange("color"); ok {
		interfaceMap["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		interfaceMap["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		interfaceMap["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		interfaceMap["ignore-errors"] = v.(bool)
	}

	log.Println("Update Interface - Map = ", interfaceMap)

	updateInterfaceRes, err := client.ApiCall("set-interface", interfaceMap, client.GetSessionID(), true, false)
	if err != nil || !updateInterfaceRes.Success {
		if updateInterfaceRes.ErrorMsg != "" {
			return fmt.Errorf(updateInterfaceRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementInterface(d, m)
}

func deleteManagementInterface(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	interfacePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete Interface")

	deleteInterfaceRes, err := client.ApiCall("delete-interface", interfacePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteInterfaceRes.Success {
		if deleteInterfaceRes.ErrorMsg != "" {
			return fmt.Errorf(deleteInterfaceRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
