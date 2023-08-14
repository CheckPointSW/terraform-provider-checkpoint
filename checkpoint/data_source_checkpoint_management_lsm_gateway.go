package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementLsmGateway() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceManagementLsmGatewayRead,
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
			"security_profile": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "LSM profile.",
			},
			"device_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Device ID.",
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
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UID.",
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
				Computed:    true,
				Description: "Provisioning settings.",
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
				Description: "Provisioning state. By default the state is 'manual'- enable provisioning but not attach to profile. If 'using-profile' state is provided a provisioning profile must be provided in provisioning-settings.",
			},

			"sic": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Secure Internal Communication.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"one_time_password": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "One-time password. When one-time password is provided without ip-address- trusted communication is automatically initiated  when the gateway connects to the Security Management server for the first time.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
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
				Computed:    true,
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
							Description: "A list of IP-addresses ranges, defined the VPN community network.This field is relevant only when 'manual' option of vpn-domain is checked.",
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
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
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
	}

}
func dataSourceManagementLsmGatewayRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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
	if v := lsmGateway["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
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
