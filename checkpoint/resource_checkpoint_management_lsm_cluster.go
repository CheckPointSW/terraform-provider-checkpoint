package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementLsmCluster() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLsmCluster,
		Read:   readManagementLsmCluster,
		Update: updateManagementLsmCluster,
		Delete: deleteManagementLsmCluster,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},

			"main_ip_address": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Main IP address.",
			},
			"security_profile": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "LSM profile.",
			},
			"dynamic_objects": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Dynamic Objects.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UID",
						},
						"comments": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Comments.",
						},
						"resolved_ip_addresses": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Single IP-address or a range of addresses.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv4 Address.",
									},
									"ipv4_address_range": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "IPv4 Address range.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"from_ipv4_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "First IPv4 address of the IP address range.",
												},
												"to_ipv4_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Last IPv4 address of the IP address range.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
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
							Optional:    true,
							Description: "New name. Overrides the interface name on profile.",
						},
						"ip_address_override": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IP address override. Net mask is defined by the attached LSM profile.",
						}, /**
						"cluster_ip_address_override": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address override. Net mask is defined by the attached LSM profile.",
						},*/
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
				Optional:    true,
				Description: "Cluster members.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Member Name. Consists of the member name in the LSM profile and the name or prefix or suffix of the cluster.",
						},
						"member_uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Member UID",
						},
						"device_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Device ID.",
						},
						"provisioning_state": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Provisioning state. This field is relevant just for SMB clusters. By default the state is 'manual'- enable provisioning but not attach to profile.If 'using-profile' state is provided a provisioning profile must be provided in provisioning-settings.",
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
							Optional:    true,
							MaxItems:    1,
							Description: "Provisioning settings. This field is relevant just for SMB clusters.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"provisioning_profile": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Provisioning profile.",
										Default:     "No Provisioning Profile",
									},
								},
							},
						},

						"sic": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "Secure Internal Communication.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IP address. When IP address is provided- initiate trusted communication immediately using this IP address.",
									},
									"one_time_password": {
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
										Description: "One-time password. When one-time password is provided without ip-address- trusted communication is automatically initiated  when the gateway connects to the Security Management server for the first time.",
									},
								},
							},
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
			"os_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Device platform operating system.",
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
				Optional:    true,
				Description: "Topology.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manual_vpn_domain": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A list of IP-addresses ranges, defined the VPN community network.This field is relevant only when manual option of vpn-domain is checked.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"comments": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Comments string.",
									},
									"from_ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "First IPv4 address of the IP address range.",
									},
									"to_ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Last IPv4 address of the IP address range.",
									},
								},
							},
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Optional:    true,
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

func createManagementLsmCluster(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	lsmCluster := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		lsmCluster["name"] = v.(string)
	}
	if v, ok := d.GetOk("main_ip_address"); ok {
		lsmCluster["main-ip-address"] = v.(string)
	}

	if v, ok := d.GetOk("security_profile"); ok {
		lsmCluster["security-profile"] = v.(string)
	}
	if v, ok := d.GetOk("dynamic_objects"); ok {

		dynamicObjectsList := v.([]interface{})

		if len(dynamicObjectsList) > 0 {

			var dynamicObjectsPayload []map[string]interface{}

			for i := range dynamicObjectsList {

				dynamicObject := dynamicObjectsList[i].(map[string]interface{})

				objectPayload := make(map[string]interface{})

				if v, ok := dynamicObject["name"]; ok {
					objectPayload["name"] = v.(string)
				}
				if v, ok := dynamicObject["resolved_ip_addresses"]; ok {
					resolvedIpAddressesList := v.([]interface{})

					var resolvedIpAddressesPayload []map[string]interface{}

					if len(resolvedIpAddressesList) > 0 {

						for j := range resolvedIpAddressesList {

							payLoad := make(map[string]interface{})
							resolvedIpAddressObj := resolvedIpAddressesList[j].(map[string]interface{})

							if v := resolvedIpAddressObj["ipv4_address"]; v != nil {
								if len(v.(string)) > 0 {
									payLoad["ipv4-address"] = v
								}

							}
							if v := resolvedIpAddressObj["ipv4_address_range"]; v != nil {

								list := v.([]interface{})
								if list != nil && len(list) > 0 {

									innerMap := list[0].(map[string]interface{})

									mapToReturn := make(map[string]interface{})

									if v := innerMap["from_ipv4_address"]; v != nil {
										mapToReturn["from-ipv4-address"] = v
									}
									if v := innerMap["to_ipv4_address"]; v != nil {
										mapToReturn["to-ipv4-address"] = v
									}
									payLoad["ipv4-address-range"] = mapToReturn
								}
							}
							resolvedIpAddressesPayload = append(resolvedIpAddressesPayload, payLoad)
						}

					}
					objectPayload["resolved-ip-addresses"] = resolvedIpAddressesPayload
				}

				dynamicObjectsPayload = append(dynamicObjectsPayload, objectPayload)
			}

			lsmCluster["dynamic-objects"] = dynamicObjectsPayload
		}
	}
	if v, ok := d.GetOk("interfaces"); ok {

		interfacesList := v.([]interface{})

		if len(interfacesList) > 0 {

			var interfacesPayload []map[string]interface{}

			for i := range interfacesList {

				Payload := make(map[string]interface{})
				intObj := interfacesList[i].(map[string]interface{})
				if v := intObj["name"]; v != nil {
					Payload["name"] = v.(string)
				}
				if v := intObj["new_name"]; v != nil {
					Payload["new-name"] = v.(string)
				}
				if v := intObj["ip_address_override"]; v != nil {
					Payload["ip-address-override"] = v.(string)
				}
				if v := intObj["member_address_override"]; v != nil {
					Payload["member-network-override"] = v.(string)
				}
				interfacesPayload = append(interfacesPayload, Payload)
			}
			lsmCluster["interfaces"] = interfacesPayload
		}
	}

	if v, ok := d.GetOk("members"); ok {

		membersList := v.([]interface{})

		var membersPayLoadList []map[string]interface{}

		if len(membersList) > 0 {

			for i := range membersList {

				memberPayload := make(map[string]interface{})

				memberObj := membersList[i].(map[string]interface{})

				if v := memberObj["name"]; v != nil {
					memberPayload["name"] = v.(string)
				}
				if v := memberObj["device_id"]; v != nil {
					memberPayload["device-id"] = v.(string)
				}
				if v := memberObj["provisioning_state"]; v != nil {
					memberPayload["provisioning-state"] = v.(string)
				}
				if v := memberObj["provisioning_settings"]; v != nil {

					provisioningSettingsPayload := make(map[string]interface{})
					if (len(v.([]interface{}))) > 0 {
						provisioningMap := v.([]interface{})[0].(map[string]interface{})

						if v := provisioningMap["provisioning_profile"]; v != nil {
							provisioningSettingsPayload["provisioning-profile"] = v.(string)
						}
						memberPayload["provisioning-settings"] = provisioningSettingsPayload
					}

				}

				if v := memberObj["sic"]; v != nil {

					sicPayload := make(map[string]interface{})
					if (len(v.([]interface{}))) > 0 {
						sicMap := v.([]interface{})[0].(map[string]interface{})
						if v := sicMap["ip_address"]; v != nil {
							if len(v.(string)) > 0 {
								sicPayload["ip-address"] = v.(string)
							}
						}
						if v := sicMap["one_time_password"]; v != nil {
							sicPayload["one-time-password"] = v.(string)
						}
					}

					memberPayload["sic"] = sicPayload
				}
				if v := memberObj["ignore_warnings"]; v != nil {
					memberPayload["ignore-warnings"] = v.(bool)
				}
				if v := memberObj["ignore_errors"]; v != nil {
					memberPayload["ignore-errors"] = v.(bool)
				}
				membersPayLoadList = append(membersPayLoadList, memberPayload)

			}

			lsmCluster["members"] = membersPayLoadList
		}
	}
	if v, ok := d.GetOk("tags"); ok {
		lsmCluster["tags"] = v.(*schema.Set).List()
	}

	if _, ok := d.GetOk("topology"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("topology.0.vpn_domain"); ok {
			res["vpn-domain"] = v.(string)
		}
		if v, ok := d.GetOk("topology.0.manual_vpn_domain"); ok {

			manualVpnDomainsList := v.([]interface{})

			var manualVpnDomainsPayload []map[string]interface{}

			if len(manualVpnDomainsList) > 0 {

				for i := range manualVpnDomainsList {

					localMap := manualVpnDomainsList[i].(map[string]interface{})
					payload := make(map[string]interface{})

					if v := localMap["commens"]; v != nil {
						payload["comments"] = v.(string)
					}
					if v := localMap["from_ipv4_address"]; v != nil {
						payload["from-ipv4-address"] = v.(string)
					}
					if v := localMap["to_ipv4_address"]; v != nil {
						payload["to-ipv4-address"] = v.(string)
					}
					manualVpnDomainsPayload = append(manualVpnDomainsPayload, payload)
				}

			}

			res["manual-vpn-domain"] = manualVpnDomainsPayload
		}

		lsmCluster["topology"] = res

	}
	if v, ok := d.GetOk("color"); ok {
		lsmCluster["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		lsmCluster["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		lsmCluster["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		lsmCluster["ignore-errors"] = v.(bool)
	}

	log.Println("Create LsmCluster - Map = ", lsmCluster)

	addLsmClusterRes, err := client.ApiCall("add-lsm-cluster", lsmCluster, client.GetSessionID(), true, false)
	if err != nil || !addLsmClusterRes.Success {
		if addLsmClusterRes.ErrorMsg != "" {
			return fmt.Errorf(addLsmClusterRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addLsmClusterRes.GetData()["uid"].(string))

	return readManagementLsmCluster(d, m)
}

func readManagementLsmCluster(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showLsmClusterRes, err := client.ApiCall("show-lsm-cluster", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLsmClusterRes.Success {
		if objectNotFound(showLsmClusterRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLsmClusterRes.ErrorMsg)
	}

	lsmCluster := showLsmClusterRes.GetData()

	log.Println("Read LsmCluster - Show JSON = ", lsmCluster)

	if v := lsmCluster["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if lsmCluster["dynamic-objects"] != nil {

		dynamicObjectsList := lsmCluster["dynamic-objects"].([]interface{})

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

	} else {
		_ = d.Set("dynamic_objects", nil)
	}

	if lsmCluster["interfaces"] != nil {

		interfacesList, ok := lsmCluster["interfaces"].([]interface{})

		if ok {

			if len(interfacesList) > 0 {

				var interfacesListToReturn []map[string]interface{}

				for i := range interfacesList {

					interfacesMap := interfacesList[i].(map[string]interface{})

					interfacesMapToAdd := make(map[string]interface{})

					if v, _ := interfacesMap["name"]; v != nil {
						interfacesMapToAdd["name"] = v
					}
					if v, _ := interfacesMap["member-network-override"]; v != nil {
						interfacesMapToAdd["member_network_override"] = v
					}
					if v, _ := interfacesMap["cluster-ip-address-override"]; v != nil {
						interfacesMapToAdd["ip_address_override"] = v
					}
					interfacesListToReturn = append(interfacesListToReturn, interfacesMapToAdd)
				}
				_ = d.Set("interfaces", interfacesListToReturn)
			}
		}
	}

	if v := lsmCluster["main-ip-address"]; v != nil {
		_ = d.Set("main_ip_address", v)
	}

	if lsmCluster["members"] != nil {

		var listOfMembersObject []map[string]interface{}

		listOfMembers := lsmCluster["members"].([]interface{})

		for i := range listOfMembers {

			memeberObj := listOfMembers[i].(map[string]interface{})

			memberMap := make(map[string]interface{})

			if v := memeberObj["member-name"]; v != nil {
				memberMap["name"] = v.(string)
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

			if confMmebersList, ok := d.GetOk("members"); ok {

				for i := range confMmebersList.([]interface{}) {
					if uid, ok := d.GetOk("members." + strconv.Itoa(i) + ".member_uid"); ok {
						if memberMap["member_uid"] == uid {
							sicMapToReturn := make(map[string]interface{})
							if sicObj, ok := d.GetOk("members." + strconv.Itoa(i) + ".sic.0"); ok {
								sicMap := sicObj.(map[string]interface{})
								if v := sicMap["ip_address"]; v != nil {
									sicMapToReturn["ip_address"] = v
								}
								if v := sicMap["one_time_password"]; v != nil {
									sicMapToReturn["one_time_password"] = v
								}
								memberMap["sic"] = []interface{}{sicMapToReturn}
								break
							}

						}
					}

				}
				if memberMap["sic"] == nil {
					memberMap["sic"] = []interface{}{}
				}
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
	if v := lsmCluster["os-name"]; v != nil {
		_ = d.Set("os_name", v)
	}
	if v := lsmCluster["security-profile"]; v != nil {
		_ = d.Set("security_profile", v)
	}
	if v := lsmCluster["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if lsmCluster["tags"] != nil {
		tagsJson, ok := lsmCluster["tags"].([]interface{})
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

	if lsmCluster["topology"] != nil {

		topologyMap, ok := lsmCluster["topology"].(map[string]interface{})

		if ok {
			topologyMapToReturn := make(map[string]interface{})

			if v, ok := topologyMap["manual-vpn-domain"]; ok {

				manualVpnDomainList := v.([]interface{})

				if len(manualVpnDomainList) > 0 {

					var manualVpnDomainMapToReturn []map[string]interface{}

					for i := range manualVpnDomainList {

						manualVpnDomain := manualVpnDomainList[i].(map[string]interface{})

						mapToReturn := make(map[string]interface{})

						if v, _ := manualVpnDomain["comments"]; v != nil {
							mapToReturn["comments"] = v
						}
						if v, _ := manualVpnDomain["from-ipv4-address"]; v != nil {
							mapToReturn["from_ipv4_address"] = v
						}
						if v, _ := manualVpnDomain["to-ipv4-address"]; v != nil {
							mapToReturn["to_ipv4_address"] = v
						}

						manualVpnDomainMapToReturn = append(manualVpnDomainMapToReturn, mapToReturn)
					}
					topologyMapToReturn["manual_vpn_domain"] = []interface{}{manualVpnDomainMapToReturn}
				}

			}
			if v := topologyMap["vpn-domain"]; v != nil {
				topologyMapToReturn["vpn_domain"] = v
			}
			_ = d.Set("topology", topologyMapToReturn)
		}
	} else {
		_ = d.Set("topology", nil)
	}

	if v := lsmCluster["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := lsmCluster["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := lsmCluster["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := lsmCluster["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := lsmCluster["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementLsmCluster(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	lsmCluster := make(map[string]interface{})

	if name, ok := d.GetOk("name"); ok {

		lsmCluster["name"] = name

	} else {
		lsmCluster["uid"] = d.Id()
	}

	if d.HasChange("dynamic_objects") {

		if v, ok := d.GetOk("dynamic_objects"); ok {

			dynamicObjectsList := v.([]interface{})

			var dynamicObjectsToReturn []map[string]interface{}

			if len(dynamicObjectsList) > 0 {

				for j := range dynamicObjectsList {

					objectToReturn := make(map[string]interface{})

					dynamicObject := dynamicObjectsList[j].(map[string]interface{})

					if v := dynamicObject["name"]; v != nil {
						if len(v.(string)) > 0 {
							objectToReturn["name"] = v.(string)
						}

					} else {
						if v := dynamicObject["uid"]; v != nil {
							if len(v.(string)) > 0 {
								objectToReturn["uid"] = v.(string)
							}
						}
					}

					if v := dynamicObject["resolved_ip_addresses"]; v != nil {

						resolvedIpAddressesList := v.([]interface{})

						var resolvedIpAddressesToReturn []map[string]interface{}

						if len(resolvedIpAddressesList) > 0 {

							for j := range resolvedIpAddressesList {

								innerMapToAdd := make(map[string]interface{})

								resolvedIpAddress := resolvedIpAddressesList[j].(map[string]interface{})

								if v := resolvedIpAddress["ipv4_address"]; v != nil {
									if len(v.(string)) > 0 {
										innerMapToAdd["ipv4-address"] = v.(string)
									}

								}
								if v := resolvedIpAddress["ipv4_address_range"]; v != nil {

									list := v.([]interface{})
									if len(list) > 0 {
										rangeValues := list[0].(map[string]interface{})
										rangeMapToReturn := make(map[string]interface{})

										if v := rangeValues["from_ipv4_address"]; v != nil {
											rangeMapToReturn["from-ipv4-address"] = v
										}

										if v := rangeValues["to_ipv4_address"]; v != nil {
											rangeMapToReturn["to-ipv4-address"] = v
										}
										innerMapToAdd["ipv4-address-range"] = rangeMapToReturn
									}
								}

								resolvedIpAddressesToReturn = append(resolvedIpAddressesToReturn, innerMapToAdd)
							}

						}
						objectToReturn["resolved-ip-addresses"] = resolvedIpAddressesToReturn

						dynamicObjectsToReturn = append(dynamicObjectsToReturn, objectToReturn)
					}

				}

				lsmCluster["dynamic-objects"] = dynamicObjectsToReturn
			}

		}

	}

	if d.HasChange("interfaces") {

		if v, ok := d.GetOk("interfaces"); ok {

			interfacesList := v.([]interface{})

			var interfacesPayload []map[string]interface{}

			for i := range interfacesList {

				Payload := make(map[string]interface{})

				if _, ok = d.GetOk("interfaces." + strconv.Itoa(i) + ".name"); ok {
					Payload["name"] = d.Get("interfaces." + strconv.Itoa(i) + ".name")
				}
				if _, ok = d.GetOk("interfaces." + strconv.Itoa(i) + ".new_name"); ok {
					Payload["new-name"] = d.Get("interfaces." + strconv.Itoa(i) + ".new_name")
				}
				if _, ok = d.GetOk("interfaces." + strconv.Itoa(i) + ".ip_address_override"); ok {
					Payload["ip-address-override"] = d.Get("interfaces." + strconv.Itoa(i) + ".ip_address_override")
				}
				if _, ok = d.GetOk("interfaces." + strconv.Itoa(i) + ".member_network_override"); ok {
					Payload["member-network-override"] = d.Get("interfaces." + strconv.Itoa(i) + ".member_network_override")
				}
				interfacesPayload = append(interfacesPayload, Payload)
			}
			lsmCluster["interfaces"] = interfacesPayload
		} else {
			oldinterfaces, _ := d.GetChange("interfaces")
			var interfacesToDelete []interface{}
			for _, i := range oldinterfaces.([]interface{}) {
				interfacesToDelete = append(interfacesToDelete, i.(map[string]interface{})["name"].(string))
			}
			lsmCluster["interfaces"] = map[string]interface{}{"remove": interfacesToDelete}
		}
	}

	if d.HasChange("members") {

		if v, ok := d.GetOk("members"); ok {

			membersList := v.([]interface{})

			if len(membersList) > 0 {

				var membersPayLoadList []map[string]interface{}

				for i := range membersList {

					memberObj := membersList[i].(map[string]interface{})
					memberMap := make(map[string]interface{})

					if v := memberObj["name"]; v != nil {
						memberMap["name"] = v
					}
					if v := memberObj["device_id"]; v != nil {
						memberMap["device-id"] = v
					}
					if v := memberObj["provisioning_state"]; v != nil {
						memberMap["provisioning-state"] = v
					}
					if v := memberObj["provisioning_settings"]; v != nil {

						provisioningSettingsPayload := make(map[string]interface{})
						provisioningObj := v.([]interface{})[0].(map[string]interface{})
						if v := provisioningObj["provisioning_profile"]; v != nil {
							provisioningSettingsPayload["provisioning-profile"] = v
						}

						memberMap["provisioning-settings"] = provisioningSettingsPayload
					}
					if v := memberObj["sic"]; v != nil {
						sicMap := make(map[string]interface{})
						sicObj := v.([]interface{})[0].(map[string]interface{})
						if v := sicObj["one_time_password"]; v != nil {
							sicMap["one-time-password"] = v
						}
						if v := sicObj["ip_address"]; v != nil {
							if len(v.(string)) > 0 {
								sicMap["ip-address"] = v
							}

						}
						memberMap["sic"] = sicMap
					}
					membersPayLoadList = append(membersPayLoadList, memberMap)
				}

				lsmCluster["members"] = membersPayLoadList
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			lsmCluster["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			lsmCluster["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if d.HasChange("topology") {

		if _, ok := d.GetOk("topology"); ok {

			topologyPayload := make(map[string]interface{})

			if v := d.Get("topology.0.vpn_domain"); v != nil {

				topologyPayload["vpn-domain"] = v.(string)
			}
			if v := d.Get("topology.0.manual_vpn_domain"); v != nil {

				var manualVpnDomainsToReturn []map[string]interface{}

				manualVpnDomainsList := v.([]interface{})

				if len(manualVpnDomainsList) > 0 {

					for i := range manualVpnDomainsList {

						manualVpnDomainPayload := make(map[string]interface{})

						manualVpnDomain := manualVpnDomainsList[i].(map[string]interface{})

						if v := manualVpnDomain["comments"]; v != nil {
							manualVpnDomainPayload["comments"] = v
						}
						if v := manualVpnDomain["from_ipv4_address"]; v != nil {
							manualVpnDomainPayload["from-ipv4-address"] = v
						}
						if v := manualVpnDomain["to_ipv4_address"]; v != nil {
							manualVpnDomainPayload["to-ipv4-address"] = v
						}
						manualVpnDomainsToReturn = append(manualVpnDomainsToReturn, manualVpnDomainPayload)
					}
				}

				topologyPayload["manual-vpn-domain"] = manualVpnDomainsToReturn
			}

			lsmCluster["topology"] = topologyPayload

		}
	}

	if ok := d.HasChange("color"); ok {
		lsmCluster["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		lsmCluster["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		lsmCluster["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		lsmCluster["ignore-errors"] = v.(bool)
	}

	log.Println("Update LsmCluster - Map = ", lsmCluster)

	updateLsmClusterRes, err := client.ApiCall("set-lsm-cluster", lsmCluster, client.GetSessionID(), true, false)
	if err != nil || !updateLsmClusterRes.Success {
		if updateLsmClusterRes.ErrorMsg != "" {
			return fmt.Errorf(updateLsmClusterRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementLsmCluster(d, m)
}

func deleteManagementLsmCluster(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	lsmClusterPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete LsmCluster")

	deleteLsmClusterRes, err := client.ApiCall("delete-lsm-cluster", lsmClusterPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteLsmClusterRes.Success {
		if deleteLsmClusterRes.ErrorMsg != "" {
			return fmt.Errorf(deleteLsmClusterRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
