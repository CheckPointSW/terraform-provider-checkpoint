package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func dataSourceManagementHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementHostRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
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
			"interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Host interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"subnet4": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network address.",
						},
						"subnet6": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network address.",
						},
						"mask_length4": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPv4 network mask length.",
						},
						"mask_length6": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPv6 network mask length.",
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
				},
			},
			"nat_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to add automatic address translation rules.",
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
						"hide_behind": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"host_servers": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
				Description: "Servers Configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Gets True if this server is a DNS Server.",
						},
						"mail_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Gets True if this server is a Mail Server.",
						},
						"web_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Gets True if this server is a Web Server.",
						},
						"web_server_config": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Computed:    true,
							Description: "Web Server configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_ports": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Server additional ports.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"application_engines": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Application engines of this web server.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"listen_standard_port": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether server listens to standard port.",
									},
									"operating_system": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Operating System.",
									},
									"protected_by": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network object which protects this server identified by the name or UID.",
									},
								},
							},
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
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
		},
	}
}

func dataSourceManagementHostRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showHostRes, err := client.ApiCall("show-host", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showHostRes.Success {
		return fmt.Errorf(showHostRes.ErrorMsg)
	}

	host := showHostRes.GetData()

	log.Println("Read Host - Show JSON = ", host)

	if v := host["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := host["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := host["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := host["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := host["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := host["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if host["interfaces"] != nil {

		interfacesList := host["interfaces"].([]interface{})

		if len(interfacesList) > 0 {

			var interfacesListToReturn []map[string]interface{}

			for i := range interfacesList {

				interfaceMap := interfacesList[i].(map[string]interface{})

				interfaceMapToAdd := make(map[string]interface{})

				if v, _ := interfaceMap["name"]; v != nil {
					interfaceMapToAdd["name"] = v
				}
				if v, _ := interfaceMap["subnet4"]; v != nil {
					interfaceMapToAdd["subnet4"] = v
				}
				if v, _ := interfaceMap["subnet6"]; v != nil {
					interfaceMapToAdd["subnet6"] = v
				}
				if v, _ := interfaceMap["mask-length4"]; v != nil {
					interfaceMapToAdd["mask_length4"] = v
				}
				if v, _ := interfaceMap["mask-length6"]; v != nil {
					interfaceMapToAdd["mask_length6"] = v
				}
				if v, _ := interfaceMap["color"]; v != nil {
					interfaceMapToAdd["color"] = v
				}
				if v, _ := interfaceMap["comments"]; v != nil {
					interfaceMapToAdd["comments"] = v
				}
				interfacesListToReturn = append(interfacesListToReturn, interfaceMapToAdd)
			}

			_ = d.Set("interfaces", interfacesListToReturn)
		} else {
			_ = d.Set("interfaces", interfacesList)
		}
	} else {
		_ = d.Set("interfaces", nil)
	}

	if host["nat-settings"] != nil {

		natSettingsMap := host["nat-settings"].(map[string]interface{})

		natSettingsMapToReturn := make(map[string]interface{})

		if v, _ := natSettingsMap["auto-rule"]; v != nil {
			natSettingsMapToReturn["auto_rule"] = strconv.FormatBool(v.(bool))
		}

		if v, _ := natSettingsMap["ipv4-address"]; v != "" && v != nil {
			natSettingsMapToReturn["ipv4_address"] = v
		}

		if v, _ := natSettingsMap["ipv6-address"]; v != "" && v != nil {
			natSettingsMapToReturn["ipv6_address"] = v
		}

		if v, _ := natSettingsMap["hide-behind"]; v != nil {
			natSettingsMapToReturn["hide_behind"] = v
		}

		if v, _ := natSettingsMap["install-on"]; v != nil {
			natSettingsMapToReturn["install_on"] = v
		}

		if v, _ := natSettingsMap["method"]; v != nil {
			natSettingsMapToReturn["method"] = v
		}

		_, natSettingInConf := d.GetOk("nat_settings")
		defaultNatSettings := map[string]interface{}{"auto_rule": "false"}
		if reflect.DeepEqual(defaultNatSettings, natSettingsMapToReturn) && !natSettingInConf {
			_ = d.Set("nat_settings", map[string]interface{}{})
		} else {
			_ = d.Set("nat_settings", natSettingsMapToReturn)
		}

	} else {
		_ = d.Set("nat_settings", nil)
	}

	if host["host-servers"] != nil {

		hostServersMap := host["host-servers"].(map[string]interface{})

		hostServersMapToReturn := make(map[string]interface{})

		if v, _ := hostServersMap["dns-server"]; v != nil {
			hostServersMapToReturn["dns_server"] = v
		}
		if v, _ := hostServersMap["mail-server"]; v != nil {
			hostServersMapToReturn["mail_server"] = v
		}
		if v, _ := hostServersMap["web-server"]; v != nil {
			hostServersMapToReturn["web_server"] = v
		}
		if v, ok := hostServersMap["web-server-config"]; ok {

			webServerConfigMap := v.(map[string]interface{})
			webServerConfigMapToReturn := make(map[string]interface{})

			if v, _ := webServerConfigMap["additional-ports"]; v != nil {
				webServerConfigMapToReturn["additional_ports"] = v
			}
			if v, _ := webServerConfigMap["application-engines"]; v != nil {
				webServerConfigMapToReturn["application_engines"] = v
			}
			if v, _ := webServerConfigMap["listen-standard-port"]; v != nil {
				webServerConfigMapToReturn["listen_standard_port"] = v
			}
			if v, _ := webServerConfigMap["operating-system"]; v != nil {
				webServerConfigMapToReturn["operating_system"] = v
			}
			if v, _ := webServerConfigMap["protected-by"]; v != nil {

				//show returned the uid, we want to set the name.
				payload := map[string]interface{}{
					"uid": v,
				}
				showProtectedByRes, err := client.ApiCall("show-object", payload, client.GetSessionID(), true, false)
				if err != nil || !showProtectedByRes.Success {
					if showProtectedByRes.ErrorMsg != "" {
						return fmt.Errorf(showProtectedByRes.ErrorMsg)
					}
					return fmt.Errorf(err.Error())
				}

				webServerConfigMapToReturn["protected_by"] = showProtectedByRes.GetData()["object"].(map[string]interface{})["name"]
			}
			hostServersMapToReturn["web_server_config"] = []interface{}{webServerConfigMapToReturn}
		}

		_ = d.Set("host_servers", []interface{}{hostServersMapToReturn})

	} else {
		_ = d.Set("host_servers", nil)
	}

	if host["groups"] != nil {
		groupsJson := host["groups"].([]interface{})
		groupsIds := make([]string, 0)
		if len(groupsJson) > 0 {
			// Create slice of group names
			for _, group := range groupsJson {
				group := group.(map[string]interface{})
				groupsIds = append(groupsIds, group["name"].(string))
			}
		}
		_ = d.Set("groups", groupsIds)
	} else {
		_ = d.Set("groups", nil)
	}

	if host["tags"] != nil {
		tagsJson := host["tags"].([]interface{})
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

	return nil
}
