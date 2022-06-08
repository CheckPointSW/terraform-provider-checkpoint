package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementProvisioningProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementProvisioningProfileRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster member unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "N/A",
			},
			"type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object type.",
			},
			"configuration_script": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Manage settings mode: locally on the device or centrally from this application.",
						},
						"override_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Override settings mode: allowed, mandatory or denied. Relevant only when settings are managed centrally.",
						},
						"configuration_script_base64": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Configuration script in base64.",
						},
					},
				},
			},
			"dns": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "DNS Settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Manage settings mode: locally on the device or centrally from this application.",
						},
						"override_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Override settings mode: allowed, mandatory or denied. Relevant only when settings are managed centrally.",
						},
						"dns_proxy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "DNS proxy enabled.",
						},
						"primary_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Primary DNS Server.",
						},
						"secondary_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Secondary DNS Server.",
						},
						"servers_configuration_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Servers configuration mode. Auto- dns configuration provided by the active internet connection. Manual- set dns servers configuration manually.",
						},
						"tertiary_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Tertiary DNS Server.",
						},
					},
				},
			},
			"domain_name": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Domain Name Settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Manage settings mode: locally on the device or centrally from this application.",
						},
						"override_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Override settings mode: allowed, mandatory or denied. Relevant only when settings are managed centrally.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Domain Name.",
						},
					},
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
			"hosts": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Hosts Settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Manage settings mode: locally on the device or centrally from this application.",
						},
						"override_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Override settings mode: allowed, mandatory or denied. Relevant only when settings are managed centrally.",
						},
						"hosts": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Hosts Settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"host_ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Host IP-Address.",
									},
									"host_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Host Name.",
									},
								},
							},
						},
					},
				},
			},
			"hotspot": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Hotspot Settings. Relevant only for Gaia Embedded (SMB) profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Manage settings mode: locally on the device or centrally from this application.",
						},
						"override_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Override settings mode: allowed, mandatory or denied. Relevant only when settings are managed centrally.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Hotspot enabled on device.",
						},
						"portal_title": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Portal title.",
						},
						"portal_message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Portal message.",
						},
						"display_terms_of_use": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Use terms of use.",
						},
						"terms_of_use": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Terms of use.",
						},
						"require_authentication": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Require authentication.",
						},
						"allow_users_from_specific_group": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Allow users from specific group.",
						},
						"allowed_users_groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Allowed users groups.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"radius": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "RADIUS Servers Settings. Relevant only for Gaia Embedded (SMB) profile.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Manage settings mode: locally on the device or centrally from this application.",
						},
						"override_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Override settings mode: allowed, mandatory or denied. Relevant only when settings are managed centrally.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Hotspot enabled on device.",
						},
						"radius_servers": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "RADIUS Servers.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"radius_server_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Radius server Name.",
									},
								},
							},
						},
						"allow_administrators_from_specific_radius_group_only": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Allow administrators from specific radius group only.",
						},
						"allowed_radius_groups": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Allowed radius groups.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementProvisioningProfileRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showProvisioningProfileRes, err := client.ApiCall("show-provisioning-profile", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showProvisioningProfileRes.Success {
		if objectNotFound(showProvisioningProfileRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showProvisioningProfileRes.ErrorMsg)
	}

	provisioningProfile := showProvisioningProfileRes.GetData()

	log.Println("Read ProvisioningProfile - Show JSON = ", provisioningProfile)

	if v := provisioningProfile["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := provisioningProfile["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := provisioningProfile["type"]; v != nil {
		_ = d.Set("type", v)
	}

	if provisioningProfile["configuration-script"] != nil {

		configurationScriptMap := provisioningProfile["configuration-script"].(map[string]interface{})

		configurationScriptMapToReturn := make(map[string]interface{})

		if v, _ := configurationScriptMap["manage-settings"]; v != nil {
			configurationScriptMapToReturn["manage_settings"] = v
		}

		if v, _ := configurationScriptMap["override-settings"]; v != nil {
			configurationScriptMapToReturn["override_settings"] = v
		}

		if v, _ := configurationScriptMap["configuration-script-base64"]; v != nil {
			configurationScriptMapToReturn["configuration_script_base64"] = v
		}

		_ = d.Set("configuration_script", configurationScriptMapToReturn)
	} else {
		_ = d.Set("configuration_script", nil)
	}

	if provisioningProfile["dns"] != nil {

		dnsMap := provisioningProfile["dns"].(map[string]interface{})

		dnsMapToReturn := make(map[string]interface{})

		if v, _ := dnsMap["manage-settings"]; v != nil {
			dnsMapToReturn["manage_settings"] = v
		}

		if v, _ := dnsMap["override-settings"]; v != nil {
			dnsMapToReturn["override_settings"] = v
		}

		if v, _ := dnsMap["dns-proxy"]; v != nil {
			dnsMapToReturn["dns_proxy"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := dnsMap["primary-server"]; v != nil {
			dnsMapToReturn["primary_server"] = v
		}

		if v, _ := dnsMap["secondary-server"]; v != nil {
			dnsMapToReturn["secondary_server"] = v
		}

		if v, _ := dnsMap["servers-configuration-mode"]; v != nil {
			dnsMapToReturn["servers_configuration_mode"] = v
		}

		if v, _ := dnsMap["tertiary-server"]; v != nil {
			dnsMapToReturn["tertiary_server"] = v
		}
		_ = d.Set("dns", dnsMapToReturn)
	} else {
		_ = d.Set("dns", nil)
	}

	if provisioningProfile["domain-name"] != nil {

		domainNameMap := provisioningProfile["domain-name"].(map[string]interface{})

		domainNameMapToReturn := make(map[string]interface{})

		if v, _ := domainNameMap["manage-settings"]; v != nil {
			domainNameMapToReturn["manage_settings"] = v
		}

		if v, _ := domainNameMap["override-settings"]; v != nil {
			domainNameMapToReturn["override_settings"] = v
		}

		if v, _ := domainNameMap["name"]; v != nil {
			domainNameMapToReturn["name"] = v
		}

		_ = d.Set("domain_name", domainNameMapToReturn)
	} else {
		_ = d.Set("domain_name", nil)
	}

	if provisioningProfile["groups"] != nil {

		interfacesList, ok := provisioningProfile["groups"].([]interface{})

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
	if provisioningProfile["hosts"] != nil {

		hostsMap := provisioningProfile["hosts"].(map[string]interface{})

		hostsMapToReturn := make(map[string]interface{})

		if v, _ := hostsMap["manage-settings"]; v != nil {
			hostsMapToReturn["manage_settings"] = v
		}

		if v, _ := hostsMap["override-settings"]; v != nil {
			hostsMapToReturn["override_settings"] = v
		}

		if hostsMap["hosts"] != nil {

			hostsList, ok := hostsMap["hosts"].([]interface{})

			var hostsListToReturn []map[string]interface{}

			if ok {

				if len(hostsList) > 0 {

					for i := range hostsList {

						hostsObjMap := hostsList[i].(map[string]interface{})

						hostsMapToAdd := make(map[string]interface{})

						if v, _ := hostsObjMap["host-ip-address"]; v != nil {
							hostsMapToAdd["host_ip_address"] = v
						}

						if v, _ := hostsObjMap["host-name"]; v != nil {
							hostsMapToAdd["host_name"] = v
						}

						hostsListToReturn = append(hostsListToReturn, hostsMapToAdd)
					}
				}
			}
			log.Println("hosts hosts: ", hostsListToReturn)
			hostsMapToReturn["hosts"] = hostsListToReturn
		}

		_ = d.Set("hosts", []interface{}{hostsMapToReturn})
	} else {
		_ = d.Set("hosts", []interface{}{})
	}

	if provisioningProfile["hotspot"] != nil {

		hotspotMap := provisioningProfile["hotspot"].(map[string]interface{})

		hotspotMapToReturn := make(map[string]interface{})

		if v, _ := hotspotMap["manage-settings"]; v != nil {
			hotspotMapToReturn["manage_settings"] = v
		}

		if v, _ := hotspotMap["override-settings"]; v != nil {
			hotspotMapToReturn["override_settings"] = v
		}

		if v, _ := hotspotMap["enabled"]; v != nil {
			hotspotMapToReturn["enabled"] = v.(bool)
		}

		if v, _ := hotspotMap["portal-title"]; v != nil {
			hotspotMapToReturn["portal_title"] = v
		}

		if v, _ := hotspotMap["portal-message"]; v != nil {
			hotspotMapToReturn["portal_message"] = v
		}

		if v, _ := hotspotMap["display-terms-of-use"]; v != nil {
			hotspotMapToReturn["display_terms_of_use"] = v.(bool)
		}

		if v, _ := hotspotMap["terms-of-use"]; v != nil {
			hotspotMapToReturn["terms_of_use"] = v
		}

		if v, _ := hotspotMap["require-authentication"]; v != nil {
			hotspotMapToReturn["require_authentication"] = v.(bool)
		}

		if v, _ := hotspotMap["allow-users-from-specific-group"]; v != nil {
			hotspotMapToReturn["allow_users_from_specific_group"] = v.(bool)
		}

		if v, _ := hotspotMap["allowed-users-groups"]; v != nil {
			allowedUsersJson, ok := hotspotMap["allowed-users-groups"].([]interface{})
			if ok {
				usersIds := make([]string, 0)
				if len(allowedUsersJson) > 0 {
					for _, user := range allowedUsersJson {
						usersIds = append(usersIds, user.(string))
					}
				}
				hotspotMapToReturn["allowed_users_groups"] = usersIds
			}
		} else {
			hotspotMapToReturn["allowed_users_groups"] = nil
		}

		err = d.Set("hotspot", []interface{}{hotspotMapToReturn})
		log.Println("hotspot gg: ", err)

	} else {
		_ = d.Set("hotspot", []interface{}{})
	}

	if provisioningProfile["radius"] != nil {

		radiusMap := provisioningProfile["radius"].(map[string]interface{})

		radiusMapToReturn := make(map[string]interface{})

		if v, _ := radiusMap["manage-settings"]; v != nil {
			radiusMapToReturn["manage_settings"] = v
		}

		if v, _ := radiusMap["override-settings"]; v != nil {
			radiusMapToReturn["override_settings"] = v
		}

		if v, _ := radiusMap["enabled"]; v != nil {
			radiusMapToReturn["enabled"] = v.(bool)
		}

		if radiusMap["radius-server"] != nil {

			serversList, ok := radiusMap["radius-server"].([]interface{})

			var serversListToReturn []map[string]interface{}

			if ok {

				if len(serversList) > 0 {

					for i := range serversList {

						serversMap := serversList[i].(map[string]interface{})

						serversMapToAdd := make(map[string]interface{})

						if v, _ := serversMap["radius-server-name"]; v != nil {
							serversMapToAdd["radius_server_name"] = v
						}

						serversListToReturn = append(serversListToReturn, serversMapToAdd)
					}
				}
			}
			radiusMapToReturn["radius_server"] = serversListToReturn
		}

		if v, _ := radiusMap["allow-administrators-from-specific-radius-group-only"]; v != nil {
			radiusMapToReturn["allow_administrators_from_specific_radius_group_only"] = v.(bool)
		}

		if v, _ := radiusMap["allowed-radius-groups"]; v != nil {
			allowedUsersJson, ok := radiusMap["allowed-radius-groups"].([]interface{})
			if ok {
				usersIds := make([]string, 0)
				if len(allowedUsersJson) > 0 {
					for _, user := range allowedUsersJson {
						usersIds = append(usersIds, user.(string))
					}
				}
				radiusMapToReturn["allowed_radius_groups"] = usersIds
			}
		} else {
			radiusMapToReturn["allowed_radius_groups"] = nil
		}

		_ = d.Set("radius", []interface{}{radiusMapToReturn})
	} else {
		_ = d.Set("radius", []interface{}{})
	}

	if provisioningProfile["tags"] != nil {
		tagsJson, ok := provisioningProfile["tags"].([]interface{})
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

	return nil

}
