package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementInterface() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementInterfaceRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Network interface name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"gateway_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Gateway or cluster object uid that the interface belongs to. <font color=\"red\">Required only if</font> name was specified.",
			},

			"anti_spoofing": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable anti-spoofing.",
			},
			"anti_spoofing_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Anti Spoofing Settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).",
						},
						"exclude_packets": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Don't check packets from excluded network.",
						},
						"excluded_network_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Excluded network name.",
						},
						"excluded_network_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Excluded network UID.",
						},
						"spoof_tracking": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Spoof tracking.",
						},
					},
				},
			},
			"cluster_members": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Network interface settings for cluster members.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster member network interface name.",
						},
						"member_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster member object uid.",
						},
						"member_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster member object name.",
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
						"network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 or IPv6 network mask. If both masks are required use ipv4-network-mask and ipv6-network-mask fields explicitly. Instead of providing mask itself it is possible to specify IPv4 or IPv6 mask length in mask-length field. If both masks length are required use ipv4-mask-length and  ipv6-mask-length fields explicitly.",
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
					},
				},
			},
			"cluster_network_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster interface type.",
			},
			"dynamic_ip": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable dynamic interface.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 network address.",
			},
			"ipv4_mask_length": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IPv4 mask length.",
			},
			"ipv4_network_mask": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 network mask.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 address.",
			},
			"ipv6_mask_length": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "IPv6 mask length.",
			},
			"ipv6_network_mask": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 network mask.",
			},
			"monitored_by_cluster": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When Private is selected as the Cluster interface type, cluster can monitor or not monitor the interface.",
			},
			"network_interface_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Network Interface Type.",
			},
			"security_zone_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Security Zone Settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_calculated": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Security Zone is calculated according to where the interface leads to.",
						},
						"specific_zone": {
							Type:        schema.TypeString,
							Computed:    true,
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
		},
	}
}

func dataSourceManagementInterfaceRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	if v, ok := d.GetOk("gateway_uid"); ok {
		payload["gateway-uid"] = v
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

	if v := interfaceMap["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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
