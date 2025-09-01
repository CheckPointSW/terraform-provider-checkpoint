package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementHosts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementHostsRead,
		Schema: map[string]*schema.Schema{
			"filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Search expression to filter objects by.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximal number of returned results.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of the results to initially skip.",
			},
			"order": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"asc": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sorts results by the given field in ascending order.",
						},
						"desc": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Sorts results by the given field in descending order.",
						},
					},
				},
			},
			"fetch_all": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If true, fetches all results.",
				Default:     false,
			},
			"from": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "From which element number the query was done.",
			},
			"to": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "To which element number the query was done.",
			},
			"total": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of elements returned by the query.",
			},
			"objects": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Objects list.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name.",
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
							Type:        schema.TypeList,
							Computed:    true,
							Description: "NAT settings.",
							MaxItems:    1,
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
						"domain": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Information about the domain that holds the Object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name.",
									},
									"uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object unique identifier.",
									},
									"domain_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Domain type.",
									},
								},
							},
						},
						"icon": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object icon.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementHostsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	showHostsRes := checkpoint.APIResponse{}
	var err error
	fetchAll, _ := d.GetOkExists("fetch_all")

	if fetchAll.(bool) {
		showHostsRes, err = client.ApiQuery("show-hosts", "full", "objects", true, map[string]interface{}{})
	} else {
		payload := make(map[string]interface{})

		if val, ok := d.GetOk("filter"); ok {
			payload["filter"] = val.(string)
		}

		if val, ok := d.GetOk("limit"); ok {
			payload["limit"] = val.(int)
		}

		if v, ok := d.GetOk("order"); ok {

			orderList := v.([]interface{})
			if len(orderList) > 0 {

				var orderPayload []map[string]interface{}

				for i := range orderList {

					payload := make(map[string]interface{})

					if v, ok := d.GetOk("order." + strconv.Itoa(i) + ".asc"); ok {
						payload["ASC"] = v.(string)
					}
					if v, ok := d.GetOk("order." + strconv.Itoa(i) + ".desc"); ok {
						payload["DESC"] = v.(string)
					}

					orderPayload = append(orderPayload, payload)
				}

				payload["order"] = orderPayload
			}
		}

		if val, ok := d.GetOk("offset"); ok {
			payload["offset"] = val.(int)
		}

		payload["details-level"] = "full"
		showHostsRes, err = client.ApiCallSimple("show-hosts", payload)
	}

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showHostsRes.Success {
		return fmt.Errorf(showHostsRes.ErrorMsg)
	}

	hosts := showHostsRes.GetData()

	log.Println("Read Hosts - Show JSON = ", hosts)

	d.SetId("show-hosts-" + acctest.RandString(10))

	if v := hosts["from"]; v != nil {
		_ = d.Set("from", v)
	}

	if v := hosts["to"]; v != nil {
		_ = d.Set("to", v)
	}

	if v := hosts["total"]; v != nil {
		_ = d.Set("total", v)
	}

	if v := hosts["objects"]; v != nil {
		objectsList := v.([]interface{})
		if len(objectsList) > 0 {
			var objectsListState []map[string]interface{}
			for i := range objectsList {
				objectMap := objectsList[i].(map[string]interface{})
				objectMapToAdd := make(map[string]interface{})

				if v := objectMap["name"]; v != nil {
					objectMapToAdd["name"] = v
				}

				if v := objectMap["uid"]; v != nil {
					objectMapToAdd["uid"] = v
				}

				if v := objectMap["type"]; v != nil {
					objectMapToAdd["type"] = v
				}

				if v := objectMap["ipv4-address"]; v != nil {
					objectMapToAdd["ipv4_address"] = v
				}

				if v := objectMap["ipv6-address"]; v != nil {
					objectMapToAdd["ipv6_address"] = v
				}

				if v := objectMap["comments"]; v != nil {
					objectMapToAdd["comments"] = v
				}

				if objectMap["interfaces"] != nil {

					interfacesList := objectMap["interfaces"].([]interface{})

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

						objectMapToAdd["interfaces"] = interfacesListToReturn
					} else {
						objectMapToAdd["interfaces"] = interfacesList
					}
				} else {
					objectMapToAdd["interfaces"] = nil
				}

				if objectMap["nat-settings"] != nil {

					natSettingsMap := objectMap["nat-settings"].(map[string]interface{})

					natSettingsMapToReturn := make(map[string]interface{})

					if v, _ := natSettingsMap["auto-rule"]; v != nil {
						natSettingsMapToReturn["auto_rule"] = v
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

					objectMapToAdd["nat_settings"] = []interface{}{natSettingsMapToReturn}

				} else {
					objectMapToAdd["nat_settings"] = nil
				}

				if objectMap["host-servers"] != nil {

					objectMapServersMap := objectMap["host-servers"].(map[string]interface{})

					objectMapServersMapToReturn := make(map[string]interface{})

					if v, _ := objectMapServersMap["dns-server"]; v != nil {
						objectMapServersMapToReturn["dns_server"] = v
					}
					if v, _ := objectMapServersMap["mail-server"]; v != nil {
						objectMapServersMapToReturn["mail_server"] = v
					}
					if v, _ := objectMapServersMap["web-server"]; v != nil {
						objectMapServersMapToReturn["web_server"] = v
					}
					if v, ok := objectMapServersMap["web-server-config"]; ok {

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
							showProtectedByRes, err := client.ApiCall("show-object", payload, client.GetSessionID(), true, client.IsProxyUsed())
							if err != nil || !showProtectedByRes.Success {
								if showProtectedByRes.ErrorMsg != "" {
									return fmt.Errorf(showProtectedByRes.ErrorMsg)
								}
								return fmt.Errorf(err.Error())
							}

							webServerConfigMapToReturn["protected_by"] = showProtectedByRes.GetData()["object"].(map[string]interface{})["name"]
						}
						objectMapServersMapToReturn["web_server_config"] = []interface{}{webServerConfigMapToReturn}
					}

					objectMapToAdd["host_servers"] = []interface{}{objectMapServersMapToReturn}

				} else {
					objectMapToAdd["host_servers"] = nil
				}

				if objectMap["groups"] != nil {
					groupsJson := objectMap["groups"].([]interface{})
					groupsIds := make([]string, 0)
					if len(groupsJson) > 0 {
						// Create slice of group names
						for _, group := range groupsJson {
							group := group.(map[string]interface{})
							groupsIds = append(groupsIds, group["name"].(string))
						}
					}
					objectMapToAdd["groups"] = groupsIds
				} else {
					objectMapToAdd["groups"] = nil
				}

				if objectMap["tags"] != nil {
					tagsJson := objectMap["tags"].([]interface{})
					var tagsIds = make([]string, 0)
					if len(tagsJson) > 0 {
						// Create slice of tag names
						for _, tag := range tagsJson {
							tag := tag.(map[string]interface{})
							tagsIds = append(tagsIds, tag["name"].(string))
						}
					}
					objectMapToAdd["tags"] = tagsIds
				} else {
					objectMapToAdd["tags"] = nil
				}

				if v := objectMap["domain"]; v != nil {
					domainMap := v.(map[string]interface{})
					domainMapToAdd := make(map[string]interface{})

					if v := domainMap["name"]; v != nil {
						domainMapToAdd["name"] = v
					}

					if v := domainMap["uid"]; v != nil {
						domainMapToAdd["uid"] = v
					}

					if v := domainMap["domain-type"]; v != nil {
						domainMapToAdd["domain_type"] = v
					}
					objectMapToAdd["domain"] = domainMapToAdd
				}

				if v := objectMap["color"]; v != nil {
					objectMapToAdd["color"] = v
				}

				if v := objectMap["icon"]; v != nil {
					objectMapToAdd["icon"] = v
				}

				objectsListState = append(objectsListState, objectMapToAdd)
			}
			_ = d.Set("objects", objectsListState)
		} else {
			_ = d.Set("objects", objectsList)
		}
	} else {
		_ = d.Set("objects", nil)
	}

	return nil
}
