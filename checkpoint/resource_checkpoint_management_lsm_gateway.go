package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementLsmGateway() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLsmGateway,
		Read:   readManagementLsmGateway,
		Update: updateManagementLsmGateway,
		Delete: deleteManagementLsmGateway,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"security_profile": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "LSM profile.",
			},
			"device_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Device ID.",
			},
			"dynamic_objects": {
				Type:        schema.TypeList,
				Optional:    true,
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
							Optional:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UID.",
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
			"ip_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IP address",
			},
			"os_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Device platform operating system.",
			},
			"provisioning_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Provisioning settings.",
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
			"provisioning_state": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Provisioning state. By default the state is 'manual'- enable provisioning but not attach to profile. If 'using-profile' state is provided a provisioning profile must be provided in provisioning-settings.",
			},

			"sic": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Secure Internal Communication.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"one_time_password": {
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							Description: "One-time password. When one-time password is provided without ip-address- trusted communication is automatically initiated  when the gateway connects to the Security Management server for the first time.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IP address. When IP address is provided- initiate trusted communication immediately using this IP address",
						},
					},
				},
			},
			"sic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication name.",
			},
			"sic_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication state.",
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
							Description: "A list of IP-addresses ranges, defined the VPN community network.This field is relevant only when 'manual' option of vpn-domain is checked.",
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
				Optional:    true,
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

func createManagementLsmGateway(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	lsmGateway := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		lsmGateway["name"] = v.(string)
	}

	if v, ok := d.GetOk("device_id"); ok {
		lsmGateway["device-id"] = v.(string)
	}
	if v, ok := d.GetOk("dynamic_objects"); ok {

		dynamicObjectsList := v.([]interface{})

		if len(dynamicObjectsList) > 0 {

			var dynamicObjectsMapToReturn []map[string]interface{}

			for j := range dynamicObjectsList {

				dynamicObject := dynamicObjectsList[j].(map[string]interface{})

				objectPayload := make(map[string]interface{})

				if v, ok := dynamicObject["name"]; ok {

					objectPayload["name"] = v.(string)
				}
				if v, ok := dynamicObject["resolved_ip_addresses"]; ok {

					resolvedIpAddressList := v.([]interface{})

					var resolvedIpAddressesPayload []map[string]interface{}

					if len(resolvedIpAddressList) > 0 {

						for i := range resolvedIpAddressList {

							payLoad := make(map[string]interface{})
							resolvedIpAddressObj := resolvedIpAddressList[i].(map[string]interface{})

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
				dynamicObjectsMapToReturn = append(dynamicObjectsMapToReturn, objectPayload)
			}
			lsmGateway["dynamic-objects"] = dynamicObjectsMapToReturn
		}
	}

	if _, ok := d.GetOk("provisioning_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("provisioning_settings.provisioning_profile"); ok {
			res["provisioning-profile"] = v.(string)
		}
		lsmGateway["provisioning-settings"] = res
	}

	if v, ok := d.GetOk("provisioning_state"); ok {
		lsmGateway["provisioning-state"] = v.(string)
	}

	if v, ok := d.GetOk("security_profile"); ok {
		lsmGateway["security-profile"] = v.(string)
	}

	if _, ok := d.GetOk("sic"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("sic.one_time_password"); ok {
			res["one-time-password"] = v.(string)
		}
		if v, ok := d.GetOk("ip_address"); ok {
			res["ip-address"] = v.(string)
		}
		lsmGateway["sic"] = res
	}

	if v, ok := d.GetOk("tags"); ok {
		lsmGateway["tags"] = v.(*schema.Set).List()
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

		lsmGateway["topology"] = res

	}
	if v, ok := d.GetOk("version"); ok {
		lsmGateway["version"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		lsmGateway["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		lsmGateway["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		lsmGateway["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		lsmGateway["ignore-errors"] = v.(bool)
	}

	log.Println("Create LsmGateway - Map = ", lsmGateway)

	addLsmGatewayRes, err := client.ApiCall("add-lsm-gateway", lsmGateway, client.GetSessionID(), true, false)
	if err != nil || !addLsmGatewayRes.Success {
		if addLsmGatewayRes.ErrorMsg != "" {
			return fmt.Errorf(addLsmGatewayRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addLsmGatewayRes.GetData()["uid"].(string))

	return readManagementLsmGateway(d, m)
}

func readManagementLsmGateway(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showLsmGatewayRes, err := client.ApiCall("show-lsm-gateway", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLsmGatewayRes.Success {
		if objectNotFound(showLsmGatewayRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLsmGatewayRes.ErrorMsg)
	}

	lsmGateway := showLsmGatewayRes.GetData()

	log.Println("Read LsmGateway - Show JSON = ", lsmGateway)

	if v := lsmGateway["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := lsmGateway["device-id"]; v != nil {
		_ = d.Set("device_id", v)
	}

	if lsmGateway["dynamic-objects"] != nil {

		dynamicObjectsList := lsmGateway["dynamic-objects"].([]interface{})

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
	if v := lsmGateway["ip-address"]; v != nil {
		_ = d.Set("ip_address", v)
	}
	if v := lsmGateway["os-name"]; v != nil {
		_ = d.Set("os_name", v)
	}
	if lsmGateway["provisioning-settings"] != nil {

		provisioningSettingsMap := lsmGateway["provisioning-settings"].(map[string]interface{})

		provisioningSettingsMapToReturn := make(map[string]interface{})

		if v, _ := provisioningSettingsMap["provisioning-profile"]; v != nil {
			provisioningSettingsMapToReturn["provisioning_profile"] = v
		}
		_ = d.Set("provisioning_settings", provisioningSettingsMapToReturn)
	} else {
		_ = d.Set("provisioning_settings", nil)
	}

	if v := lsmGateway["provisioning-state"]; v != nil {
		_ = d.Set("provisioning_state", v)
	}

	if v := lsmGateway["security-profile"]; v != nil {
		_ = d.Set("security_profile", v)
	}
	if v, ok := d.GetOk("sic"); ok {
		sicObj := v.(map[string]interface{})
		sicMap := make(map[string]interface{})
		if v := sicObj["one_time_password"]; v != nil {
			sicMap["one_time_password"] = v
		}
		if v := sicObj["ip_address"]; v != nil {
			sicMap["ip_address"] = v
		}
		_ = d.Set("sic", sicMap)
	}
	if v := lsmGateway["sic-name"]; v != nil {
		_ = d.Set("sic_name", v)
	}
	if v := lsmGateway["sic-state"]; v != nil {
		_ = d.Set("sic_state", v)
	}

	if lsmGateway["tags"] != nil {
		tagsJson, ok := lsmGateway["tags"].([]interface{})
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

	if lsmGateway["topology"] != nil {

		topologyMap, ok := lsmGateway["topology"].(map[string]interface{})

		if ok {
			topologyMapToReturn := make(map[string]interface{})

			if v, ok := topologyMap["manual-vpn-domain"]; ok {

				manualVpnDomainList := v.([]interface{})

				if len(manualVpnDomainList) > 0 {

					var manualVpnDomainMapToReturn []map[string]interface{}

					for i := range manualVpnDomainList {

						manualVpnDomain := manualVpnDomainList[i].(map[string]interface{})

						mapToReturn := make(map[string]interface{})

						if v := manualVpnDomain["comments"]; v != nil {
							mapToReturn["comments"] = v
						}
						if v := manualVpnDomain["from-ipv4-address"]; v != nil {
							mapToReturn["from_ipv4_address"] = v
						}
						if v := manualVpnDomain["to-ipv4-address"]; v != nil {
							mapToReturn["to_ipv4_address"] = v
						}

						manualVpnDomainMapToReturn = append(manualVpnDomainMapToReturn, mapToReturn)
					}
					topologyMapToReturn["manual_vpn_domain"] = manualVpnDomainMapToReturn
				}

			}
			if v := topologyMap["vpn-domain"]; v != nil {
				topologyMapToReturn["vpn_domain"] = v
			}

			_ = d.Set("topology", []interface{}{topologyMapToReturn})
		}
	} else {
		_ = d.Set("topology", nil)
	}
	if v := lsmGateway["version"]; v != nil {
		_ = d.Set("version", v)
	}
	if v := lsmGateway["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := lsmGateway["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := lsmGateway["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := lsmGateway["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementLsmGateway(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	lsmGateway := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		lsmGateway["name"] = oldName
		lsmGateway["new-name"] = newName
	} else {
		lsmGateway["name"] = d.Get("name")
	}

	if ok := d.HasChange("device_id"); ok {
		lsmGateway["device-id"] = d.Get("device_id")
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

				lsmGateway["dynamic-objects"] = dynamicObjectsToReturn
			}

		}

	}

	if d.HasChange("provisioning_settings") {

		if _, ok := d.GetOk("provisioning_settings"); ok {

			res := make(map[string]interface{})

			if d.HasChange("provisioning_settings.provisioning_profile") {
				res["provisioning-profile"] = d.Get("provisioning_settings.provisioning_profile")
			}
			lsmGateway["provisioning-settings"] = res
		}
	}

	if ok := d.HasChange("provisioning_state"); ok {
		lsmGateway["provisioning-state"] = d.Get("provisioning_state")
	}

	if d.HasChange("sic") {

		if _, ok := d.GetOk("sic"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOk("sic.one_time_password"); ok {
				res["one-time-password"] = v.(string)
			}
			lsmGateway["sic"] = res
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			lsmGateway["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			lsmGateway["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
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

			lsmGateway["topology"] = topologyPayload

		}
	}

	if ok := d.HasChange("color"); ok {
		lsmGateway["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		lsmGateway["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		lsmGateway["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		lsmGateway["ignore-errors"] = v.(bool)
	}

	log.Println("Update LsmGateway - Map = ", lsmGateway)

	updateLsmGatewayRes, err := client.ApiCall("set-lsm-gateway", lsmGateway, client.GetSessionID(), true, false)
	if err != nil || !updateLsmGatewayRes.Success {
		if updateLsmGatewayRes.ErrorMsg != "" {
			return fmt.Errorf(updateLsmGatewayRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementLsmGateway(d, m)
}

func deleteManagementLsmGateway(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	lsmGatewayPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete LsmGateway")

	deleteLsmGatewayRes, err := client.ApiCall("delete-lsm-gateway", lsmGatewayPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteLsmGatewayRes.Success {
		if deleteLsmGatewayRes.ErrorMsg != "" {
			return fmt.Errorf(deleteLsmGatewayRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
