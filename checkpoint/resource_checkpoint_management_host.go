package checkpoint

import (
	"fmt"
	checkpoint "github.com/Checkpoint/api_go_sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"sort"
	"strconv"
)

func resourceManagementHost() *schema.Resource {
	return &schema.Resource{
		Create: createManagementHost,
		Read:   readManagementHost,
		Update: updateManagementHost,
		Delete: deleteManagementHost,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "Object name. Should be unique in the domain.",
			},
			"ipv4_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Description: "IPv4 address.",
			},
			"ipv6_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Description: "IPv6 address.",
			},
			"interfaces": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Description: "Host interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type: schema.TypeString,
							Required: true,
							Description: "Object name. Should be unique in the domain.",
						},
						"subnet4": {
							Type: schema.TypeString,
							Optional: true,
							Description: "IPv4 network address.",
						},
						"subnet6": {
							Type: schema.TypeString,
							Optional: true,
							Description: "IPv6 network address.",
						},
						"mask_length4": {
							Type: schema.TypeInt,
							Optional: true,
							Description: "IPv4 network mask length.",
						},
						"mask_length6": {
							Type: schema.TypeInt,
							Optional: true,
							Description: "IPv6 network mask length.",
						},
						"ignore_warnings": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Apply changes ignoring warnings.",
							Default: false,
						},
						"ignore_errors": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
							Default: false,
						},
						"color": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Color of the object. Should be one of existing colors.",
							Default: "black",
						},
						"comments": &schema.Schema{
							Type:	schema.TypeString,
							Optional: true,
							Description: "Comments string.",
						},
						"details_level": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.",
							Default: "standard",
						},
					},
				},
			},
			"nat_settings" : {
				Type: schema.TypeMap,
				Optional: true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:     schema.TypeBool,
							Optional: true,
							Description: "Whether to add automatic address translation rules.",
						},
						"ipv4_address": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "IPv6 address.",
						},
						"hide_behind": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"host_servers": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_server": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Description: "Gets True if this server is a DNS Server.",
						},
						"mail_server": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Description: "Gets True if this server is a Mail Server.",
						},
						"web_server": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Description: "Gets True if this server is a Web Server.",
						},
						"web_server_config": &schema.Schema{
							Type:     schema.TypeMap,
							Optional: true,
							Description: "Web Server configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_ports": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Description: "Server additional ports.",
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"application_engines": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Description: "Application engines of this web server.",
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"listen_standard_port": &schema.Schema{
										Type:     schema.TypeBool,
										Optional: true,
										Description: "Whether server listens to standard port.",
									},
									"operating_system": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Description: "Operating System.",
									},
									"protected_by": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"set_if_exists": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If another object with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original object's fields will be overwritten by the fields provided in the request payload!",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default: false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default: false,
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default: "black",
			},
			"comments": &schema.Schema{
				Type:	schema.TypeString,
				Optional: true,
				Description: "Comments string.",
			},
			"details_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.",
				Default: "standard",
			},
			"groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Type: schema.TypeSet,
				Optional: true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createManagementHost(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	host := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		host["name"] = v.(string)
	}
	if v, ok := d.GetOk("ipv4_address"); ok {
		host["ipv4-address"] = v.(string)
	}
	if v, ok := d.GetOk("ipv6_address"); ok {
		host["ipv6-address"] = v.(string)
	}

	//list of objects
	if v, ok := d.GetOk("interfaces"); ok {

		interfacesList := v.([]interface{})
		if len(interfacesList) > 0 {

			var interfacesPayload []map[string]interface{}

			for i := range interfacesList {

				payload := make(map[string]interface{})

				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".name"); ok {
					payload["name"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".subnet4"); ok {
					payload["subnet4"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".subnet6"); ok {
					payload["subnet6"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".mask_length4"); ok {
					payload["mask-length4"] = v.(int)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".mask_length6"); ok {
					payload["mask-length6"] = v.(int)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ignore_warnings"); ok {
					payload["ignore-warnings"] = v.(bool)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ignore_errors"); ok {
					payload["ignore-errors"] = v.(bool)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".color"); ok {
					payload["color"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".comments"); ok {
					payload["comments"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".details_level"); ok {
					payload["details-level"] = v.(string)
				}
				interfacesPayload = append(interfacesPayload, payload)
			}
			host["interfaces"] = interfacesPayload
		}
	}

	if _, ok := d.GetOk("nat_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("nat_settings.auto_rule"); ok {
			res["auto-rule"] = v
		}
		if v, ok := d.GetOk("nat_settings.ipv4_address"); ok {
			res["ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.ipv6_address"); ok {
			res["ipv6-address"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.hide_behind"); ok {
			res["hide-behind"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.install_on"); ok {
			res["install-on"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.method"); ok {
			res["method"] = v.(string)
		}
		host["nat-settings"] = res
	}

	//2 level values - host-servers
	//can be only object
	if v, ok := d.GetOk("host_servers"); ok {

		hostServersList := v.([]interface{})
		if len(hostServersList) > 0 {

			hostServersPayload := make(map[string]interface{})

			if v, ok := d.GetOk("host_servers.0.dns_server"); ok {
				hostServersPayload["dns-server"] = v.(bool)
			}
			if v, ok := d.GetOk("host_servers.0.mail_server"); ok {
				hostServersPayload["mail-server"] = v.(bool)
			}
			if v, ok := d.GetOk("host_servers.0.web_server"); ok {
				hostServersPayload["web-server"] = v.(bool)
			}
			if _, ok := d.GetOk("host_servers.0.web_server_config"); ok {

				webServerConfigPayLoad := make(map[string]interface{})

				if v, ok := d.GetOk("host_servers.0.web_server_config.additional_ports"); ok {
					webServerConfigPayLoad["additional-ports"] = v.([]interface{})
				}
				if v, ok := d.GetOk("host_servers.0.web_server_config.application_engines"); ok {
					webServerConfigPayLoad["application-engines"] = v.([]interface{})
				}
				if v, ok := d.GetOk("host_servers.0.web_server_config.listen_standard_port"); ok {
					webServerConfigPayLoad["listen-standard-port"] = v
				}
				if v, ok := d.GetOk("host_servers.0.web_server_config.operating_system"); ok {
					webServerConfigPayLoad["operating-system"] = v.(string)
				}
				if v, ok := d.GetOk("host_servers.0.web_server_config.protected_by"); ok {
					webServerConfigPayLoad["protected-by"] = v.(string)
				}
				hostServersPayload["web-server-config"] = webServerConfigPayLoad
			}
			host["host-servers"] = hostServersPayload
		}
	}

	if val, ok := d.GetOk("comments"); ok {
		host["comments"] = val.(string)
	}
	if val, ok := d.GetOk("tags"); ok {
		host["tags"] = val.(*schema.Set).List()
	}
	if val, ok := d.GetOk("groups"); ok {
		host["groups"] = val.(*schema.Set).List()
	}
	if val, ok := d.GetOk("set_if_exists"); ok {
		host["set-if-exists"] = val.(bool)
	}
	if val, ok := d.GetOk("color"); ok {
		host["color"] = val.(string)
	}
	if val, ok := d.GetOk("details_level"); ok {
		host["details-level"] = val.(string)
	}
	if val, ok := d.GetOk("ignore_errors"); ok {
		host["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOk("ignore_warnings"); ok {
		host["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Host - Map = ", host)

	addHostRes, err := client.ApiCall("add-host", host, client.GetSessionID(), false, false)
	if err != nil || !addHostRes.Success {
		if addHostRes.ErrorMsg != "" {
			return fmt.Errorf(addHostRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addHostRes.GetData()["uid"].(string))

	return readManagementHost(d, m)
}

func readManagementHost(d *schema.ResourceData, m interface{}) error{

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showHostRes, err := client.ApiCall("show-host",payload,client.GetSessionID(),true,false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showHostRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showHostRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showHostRes.ErrorMsg)
	}

	host := showHostRes.GetData()

	log.Println("Read Host - Show JSON = ", host)


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

			//we would like to sort the interfaces as they are in the .tf file
			interfacesListToOrder := d.Get("interfaces").([]interface{})
			sort.Slice(interfacesList, func(i, j int) bool { return interfacesListToOrder[i].(map[string]interface{})["name"].(string) > interfacesListToOrder[j].(map[string]interface{})["name"].(string) })

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
				interfaceMapToAdd["details_level"] = "standard"
				interfaceMapToAdd["ignore_errors"] = false
				interfaceMapToAdd["ignore_warnings"] = false

				interfacesListToReturn = append(interfacesListToReturn, interfaceMapToAdd)
			}
			_ = d.Set("interfaces", interfacesListToReturn)
		}
	} else {
		_ = d.Set("interfaces", nil)
	}


	if host["nat-settings"] != nil {

		natSettingsMap := host["nat-settings"].(map[string]interface{})

		natSettingsMapToReturn := make(map[string]interface{})

		if v, _ := natSettingsMap["auto-rule"]; v != nil {
			natSettingsMapToReturn["auto_rule"] = v
		}
		if v, _ := natSettingsMap["ipv4-address"]; v != nil {
			natSettingsMapToReturn["ipv4_address"] = v
		}
		if v, _ := natSettingsMap["ipv6-address"]; v != nil {
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

		var natSettingsList []interface{}
		natSettingsList = append(natSettingsList, natSettingsMapToReturn)
		_ = d.Set("nat_settings", natSettingsList)
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
				webServerConfigMapToReturn["protected_by"] = v
			}
			if v, _ := webServerConfigMap["standard-port-number"]; v != nil {
				webServerConfigMapToReturn["standard_port_number"] = v
			}
			hostServersMapToReturn["web_server_config"] = webServerConfigMapToReturn
		}
		_ = d.Set("host-servers", hostServersMap)
	} else {
		_ = d.Set("host-servers", nil)
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

func updateManagementHost(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	host := make(map[string]interface{})
	apiCall := false

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		host["name"] = oldName
			host["new-name"] = newName
		apiCall = true
	} else {
		host["name"] = d.Get("name")
	}

	if d.HasChange("interfaces") {
		if v, ok := d.GetOk("interfaces"); ok {

			interfacesList := v.([]interface{})

			if len(interfacesList) > 0 {

				var interfacesPayload []map[string]interface{}

				for i := range interfacesList {

					payload := make(map[string]interface{})

					//name, subnets, mask lengths are required to request
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".name"); ok {
						payload["name"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".subnet4"); ok {
						payload["subnet4"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".subnet6"); ok {
						payload["subnet6"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".mask_length4"); ok {
						payload["mask-length4"] = v.(int)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".mask_length6"); ok {
						payload["mask-length6"] = v.(int)
					}
					if d.HasChange("interfaces." + strconv.Itoa(i) + ".ignore_warnings"){
						if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ignore_warnings"); ok {
							payload["ignore-warnings"] = v.(bool)
						}
					}
					if d.HasChange("interfaces." + strconv.Itoa(i) + ".color"){
						if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".color"); ok {
							payload["color"] = v.(string)
						}
					}
					if d.HasChange("interfaces." + strconv.Itoa(i) + ".ignore_errors"){
						if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ignore_errors"); ok {
							payload["ignore-errors"] = v.(bool)
						}
					}
					if d.HasChange("interfaces." + strconv.Itoa(i) + ".comments"){
						if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".comments"); ok {
							payload["comments"] = v.(string)
						}
					}
					if d.HasChange("interfaces." + strconv.Itoa(i) + ".details_level"){
						if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".details_level"); ok {
							payload["details-level"] = v.(string)
						}
					}
					interfacesPayload = append(interfacesPayload, payload)
				}
				host["interfaces"] = interfacesPayload
				apiCall = true
			}

		}
	}

	if v, ok := d.GetOk("details_level"); ok {
		host["details-level"] = v.(string)
		apiCall = true
	}
	if v, ok := d.GetOk("ignore_errors"); ok {
		host["ignore-errors"] = v.(bool)
		apiCall = true
	}
	if v, ok := d.GetOk("ignore_warnings"); ok {
		host["ignore-warnings"] = v.(bool)
		apiCall = true
	}

	if ok := d.HasChange("comments"); ok {
		if v, ok := d.GetOk("comments"); ok {
			host["comments"] = v.(string)
			apiCall = true
		}
	}
	if ok := d.HasChange("color"); ok {
		if v, ok := d.GetOk("color"); ok {
			host["color"] = v.(string)
			apiCall = true
		}
	}

	if ok := d.HasChange("ipv4_address"); ok {
		if v, ok := d.GetOk("ipv4_address"); ok {
			host["ipv4-address"] = v.(string)
			apiCall = true
		}
	}
	if ok := d.HasChange("ipv6_address"); ok {
		if v, ok := d.GetOk("ipv6_address"); ok {
			host["ipv6-address"] = v.(string)
			apiCall = true
		}
	}
	if ok := d.HasChange("nat_settings"); ok {

		if _, ok := d.GetOk("nat_settings"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOk("nat_settings.auto_rule"); ok {
				res["auto-rule"] = v
			}
			if v, ok := d.GetOk("nat_settings.ipv4_address"); ok {
				res["ipv4-address"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.ipv6_address"); ok {
				res["ipv6-address"] = v.(string)
			}
			if d.HasChange("nat_settings.hide_behind") {
				if v, ok := d.GetOk("nat_settings.hide_behind"); ok {
					res["hide-behind"] = v.(string)
				}
			}
			if d.HasChange("nat_settings.install_on"){
				if v, ok := d.GetOk("nat_settings.install_on"); ok {
					res["install-on"] = v.(string)
				}
			}
			if d.HasChange("nat_settings.method") {
				if v, ok := d.GetOk("nat_settings.method"); ok {
					res["method"] = v.(string)
				}
			}
			host["nat-settings"] = res
			apiCall = true
		}
	}
	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			host["tags"] = v.(*schema.Set).List()
			apiCall = true
		}
	}
	if ok := d.HasChange("groups"); ok {
		if v, ok := d.GetOk("groups"); ok {
			host["groups"] = v.(*schema.Set).List()
			apiCall = true
		}
	}
	if ok := d.HasChange("host_servers"); ok {
		if v, ok := d.GetOk("host_servers"); ok {
			hostServersList := v.([]interface{})
			if len(hostServersList) > 0 {
				hostServersPayload := make(map[string]interface{})

				if d.HasChange("host_servers.0.dns_server") {
					hostServersPayload["dns-server"] = d.Get("host_servers.0.dns_server").(bool)
				}
				if d.HasChange("host_servers.0.mail_server") {
					hostServersPayload["mail-server"] = d.Get("host_servers.0.mail_server").(bool)
				}
				if d.HasChange("host_servers.0.web_server"){
					hostServersPayload["web-server"] = d.Get("host_servers.0.web_server").(bool)
				}
				if d.HasChange("host_servers.0.web_server_config") {

					hostServersPayload["web-server"] = d.Get("host_servers.0.web_server").(bool)
					webServerConfigPayLoad := make(map[string]interface{})

					if d.HasChange("host_servers.0.web_server_config.additional_ports") {
						webServerConfigPayLoad["additional-ports"] = d.Get("host_servers.0.web_server_config.additional_ports").([]interface{})
					}
					if d.HasChange("host_servers.0.web_server_config.application_engines") {
						webServerConfigPayLoad["application-engines"] = d.Get("host_servers.0.web_server_config.application_engines").([]interface{})
					}
					if d.HasChange("host_servers.0.web_server_config.listen_standard_port") {
						webServerConfigPayLoad["listen-standard-port"] = d.Get("host_servers.0.web_server_config.listen_standard_port")
					}
					if d.HasChange("host_servers.0.web_server_config.operating_system") {
						webServerConfigPayLoad["operating-system"] = d.Get("host_servers.0.web_server_config.operating_system").(string)
					}
					if d.HasChange("host_servers.0.web_server_config.protected_by") {
						webServerConfigPayLoad["protected-by"] = d.Get("host_servers.0.web_server_config.protected_by").(string)
					}

					hostServersPayload["web-server-config"] = webServerConfigPayLoad
				}
				host["host-servers"] = hostServersPayload

			}
			apiCall = true
		}
	}


	if apiCall{
		log.Println("Update Host - Map = ", host)
		updateHostRes, err := client.ApiCall("set-host", host, client.GetSessionID(), false, false)
		if err != nil || !updateHostRes.Success {
			if updateHostRes.ErrorMsg != "" {
				return fmt.Errorf(updateHostRes.ErrorMsg)
			}
			return fmt.Errorf(err.Error())
		}
	}

	return readManagementHost(d, m)
}

func deleteManagementHost(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	hostPayload := map[string]interface{}{
		"uid" : d.Id(),
	}

	deleteHostRes, err := client.ApiCall("delete-host", hostPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteHostRes.Success {
		if deleteHostRes.ErrorMsg != "" {
			return fmt.Errorf(deleteHostRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")
	return nil
}


