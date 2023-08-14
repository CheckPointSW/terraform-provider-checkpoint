package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementLsmCluster() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceManagementLsmClusterRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object uid.",
			},
			"dynamic_objects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Dynamic Objects.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comments.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"resolved_ip_addresses": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Single IP-address or a range of addresses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv4 Address.",
									},
									"ipv4_address_range": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "IPv4 Address range.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"from_ipv4_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "First IPv4 address of the IP address range.",
												},
												"to_ipv4_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Last IPv4 address of the IP address range.",
												},
											},
										},
									},
								},
							},
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UID",
						},
					},
				},
			},

			"main_ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Main IP address.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Interface name.",
						},
						"new_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "New name. Overrides the interface name on profile.",
						},
						"ip_address_override": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IP address override. Net mask is defined by the attached LSM profile.",
						},
						"cluster_ip_address_override": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address override. Net mask is defined by the attached LSM profile.",
						},
						"member_network_override": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Member network override. Net mask is defined by the attached LSM profile.",
						},
					},
				},
			},
			"members": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cluster members.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Member Name. Consists of the member name in the LSM profile and the name or prefix or suffix of the cluster.",
						},
						"member_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Member UID",
						},
						"device_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Device ID.",
						},
						"interfaces": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Interfaces",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IP address.",
									},
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Interface name",
									},
								},
							},
						},
						"main_ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Main ip address",
						},
						"provisioning_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Provisioning settings. This field is relevant just for SMB clusters.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"provisioning_profile": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Provisioning profile.",
									},
								},
							},
						},
						"provisioning_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Provisioning state. This field is relevant just for SMB clusters. By default the state is 'manual'- enable provisioning but not attach to profile.If 'using-profile' state is provided a provisioning profile must be provided in provisioning-settings.",
						},
						"sic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secure Internal Communication name",
						},
						"sic_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secure Internal Communication state",
						},
					},
				},
			},
			"os_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Device platform operating system.",
			},
			"security_profile": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "LSM profile.",
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
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Topology.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manual_vpn_domain": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A list of IP-addresses ranges, defined the VPN community network.This field is relevant only when manual option of vpn-domain is checked.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"comments": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Comments string.",
									},
									"from_ipv4_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "First IPv4 address of the IP address range.",
									},
									"to_ipv4_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Last IPv4 address of the IP address range.",
									},
								},
							},
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN Domain type.  'external-interfaces-only' is relevnt only for Gaia devices. 'hide-behind-gateway-external-ip-address' is relevant only for SMB devices.",
						},
					},
				},
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Device platform version.",
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

func dataSourceManagementLsmClusterRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showLsmClusterRes, err := client.ApiCall("show-lsm-cluster", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLsmClusterRes.Success {
		return fmt.Errorf(showLsmClusterRes.ErrorMsg)
	}

	LsmCluster := showLsmClusterRes.GetData()

	log.Println("Read Lsm Cluster - Show JSON = ", LsmCluster)

	if v := LsmCluster["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}
	if v := LsmCluster["name"]; v != nil {
		_ = d.Set("name", v)
	}
	if LsmCluster["dynamic-objects"] != nil {

		dynamicObjectsList := LsmCluster["dynamic-objects"].([]interface{})

		if len(dynamicObjectsList) > 0 {
			var dynamicObjectsToReturn []map[string]interface{}

			for i := range dynamicObjectsList {

				mapToAdd := make(map[string]interface{})

				dynamicObject := dynamicObjectsList[i].(map[string]interface{})

				if v := dynamicObject["name"]; v != nil {
					mapToAdd["name"] = v
				}
				if v := dynamicObject["uid"]; v != nil {
					mapToAdd["uid"] = v
				}

				if v := dynamicObject["comments"]; v != nil {
					mapToAdd["comments"] = v
				}

				if v := dynamicObject["resolved-ip-addresses"]; v != nil {

					resolvedIpAddressesList := v.([]interface{})

					var resolvedIpAddressesToReturn []map[string]interface{}

					if len(resolvedIpAddressesList) > 0 {

						for j := range resolvedIpAddressesList {

							innerMapToAdd := make(map[string]interface{})

							resolvedIpAddress := resolvedIpAddressesList[j].(map[string]interface{})

							if v := resolvedIpAddress["ipv4-address"]; v != nil {
								innerMapToAdd["ipv4_address"] = v
							}
							if v := resolvedIpAddress["ipv4-address-range"]; v != nil {

								rangeValues := v.(map[string]interface{})

								rangeMapToReturn := make(map[string]interface{})

								if v := rangeValues["from-ipv4-address"]; v != nil {
									rangeMapToReturn["from_ipv4_address"] = v
								}

								if v := rangeValues["to-ipv4-address"]; v != nil {
									rangeMapToReturn["to_ipv4_address"] = v
								}
								innerMapToAdd["ipv4_address_range"] = []interface{}{rangeMapToReturn}
							} else {
								innerMapToAdd["ipv4_address_range"] = nil
							}
							resolvedIpAddressesToReturn = append(resolvedIpAddressesToReturn, innerMapToAdd)
						}

					}

					mapToAdd["resolved_ip_addresses"] = resolvedIpAddressesToReturn

				}

				dynamicObjectsToReturn = append(dynamicObjectsToReturn, mapToAdd)
			}

			_ = d.Set("dynamic_objects", dynamicObjectsToReturn)
		}
	}
	if LsmCluster["interfaces"] != nil {

		interfacesList := LsmCluster["interfaces"].([]interface{})

		if len(interfacesList) > 0 {

			var interfacesListToReturn []map[string]interface{}

			for i := range interfacesList {

				interfacesMap := interfacesList[i].(map[string]interface{})

				interfacesMapToAdd := make(map[string]interface{})

				if v, _ := interfacesMap["name"]; v != nil {
					interfacesMapToAdd["name"] = v.(string)
				}
				if v, _ := interfacesMap["member-network-override"]; v != nil {
					interfacesMapToAdd["member_network_override"] = v.(string)
				}
				if v, _ := interfacesMap["cluster-ip-address-override"]; v != nil {
					interfacesMapToAdd["ip_address_override"] = v.(string)
				}
				interfacesListToReturn = append(interfacesListToReturn, interfacesMapToAdd)
			}

			_ = d.Set("interfaces", interfacesListToReturn)
		}
	} else {
		_ = d.Set("interfaces", nil)
	}
	if v := LsmCluster["main-ip-address"]; v != nil {
		_ = d.Set("main_ip_address", v)
	}
	if LsmCluster["members"] != nil {

		var listOfMembersObject []map[string]interface{}

		listOfMembers := LsmCluster["members"].([]interface{})

		for i := range listOfMembers {

			memeberObj := listOfMembers[i].(map[string]interface{})

			memberMap := make(map[string]interface{})

			if v := memeberObj["member-name"]; v != nil {
				memberMap["member_name"] = v.(string)
			}

			if v := memeberObj["member-uid"]; v != nil {
				memberMap["member_uid"] = v.(string)
			}

			if v := memeberObj["device-id"]; v != nil {
				memberMap["device_id"] = v.(string)
			}
			if v := memeberObj["main-ip-address"]; v != nil {
				memberMap["main_ip_address"] = v.(string)
			}
			if v := memeberObj["provisioning-state"]; v != nil {
				memberMap["provisioning_state"] = v.(string)
			}
			if v := memeberObj["sic-name"]; v != nil {
				memberMap["sic_name"] = v.(string)
			}
			if v := memeberObj["sic-state"]; v != nil {
				memberMap["sic_state"] = v.(string)
			}
			if v, ok := memeberObj["provisioning-settings"]; ok {

				provisioningSettingsMap, ok := v.(map[string]interface{})
				if ok {
					provisioningSettingsMapToReturn := make(map[string]interface{})

					if v, _ := provisioningSettingsMap["provisioning-profile"]; v != nil {
						provisioningSettingsMapToReturn["provisioning_profile"] = v.(string)
					}
					memberMap["provisioning_settings"] = []interface{}{provisioningSettingsMapToReturn}
				}
			} else {
				memberMap["provisioning_settings"] = nil
			}

			if v := memberMap["interfaces"]; v != nil {

				var interfacesList []map[string]interface{}
				list := v.([]interface{})
				if len(list) > 0 {

					for j := range list {

						localMap := list[j].(map[string]interface{})
						mapToReturn := make(map[string]interface{})

						if v := localMap["name"]; v != nil {
							mapToReturn["name"] = v.(string)
						}
						if v := localMap["ip-address"]; v != nil {
							mapToReturn["ip_address"] = v.(string)
						}

						interfacesList = append(interfacesList, mapToReturn)
					}
				}

				memberMap["interfaces"] = interfacesList
			} else {
				memberMap["interfaces"] = nil
			}

			listOfMembersObject = append(listOfMembersObject, memberMap)
		}

		_ = d.Set("members", listOfMembersObject)

	} else {
		_ = d.Set("members", nil)
	}
	if v := LsmCluster["os-name"]; v != nil {
		_ = d.Set("os_name", v)
	}
	if v := LsmCluster["security-profile"]; v != nil {
		_ = d.Set("security_profile", v)
	}
	if LsmCluster["tags"] != nil {
		tagsJson := LsmCluster["tags"].([]interface{})
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

	if LsmCluster["topology"] != nil {

		topologyMap := LsmCluster["topology"].(map[string]interface{})

		topologyMapToReturn := make(map[string]interface{})

		if v, _ := topologyMap["vpn-domain"]; v != nil {
			topologyMapToReturn["vpn_domain"] = v.(string)
		}
		if v, _ := topologyMap["manual-vpn-domain"]; v != nil {

			manualVpnDomainsList, ok := topologyMap["manual-vpn-domain"].([]interface{})

			if ok {
				var ManualVpnDomainsListToReturn []map[string]interface{}

				if len(manualVpnDomainsList) > 0 {

					for i := range manualVpnDomainsList {

						manualVpnDomainMap := manualVpnDomainsList[i].(map[string]interface{})

						manualVpnDomainMapToAdd := make(map[string]interface{})

						if v, _ := manualVpnDomainMap["comments"]; v != nil {
							manualVpnDomainMapToAdd["comments"] = v
						}
						if v, _ := manualVpnDomainMap["from-ipv4-address"]; v != nil {
							manualVpnDomainMapToAdd["from_ipv4_address"] = v
						}
						if v, _ := manualVpnDomainMap["to-ipv4-address"]; v != nil {
							manualVpnDomainMapToAdd["to_ipv4_address"] = v
						}
						ManualVpnDomainsListToReturn = append(ManualVpnDomainsListToReturn, manualVpnDomainMapToAdd)
					}

				}

				topologyMapToReturn["manual_vpn_domain"] = ManualVpnDomainsListToReturn
			}

		}

		_ = d.Set("topology", []interface{}{topologyMapToReturn})

	} else {
		_ = d.Set("topology", nil)
	}
	if v := LsmCluster["version"]; v != nil {
		_ = d.Set("version", v)
	}
	if v := LsmCluster["comments"]; v != nil {
		_ = d.Set("comments", v)
	}
	if v := LsmCluster["color"]; v != nil {
		_ = d.Set("color", v)
	}
	return nil
}
