package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementHost() *schema.Resource {
	return &schema.Resource{
		Create: createManagementHost,
		Read:   readManagementHost,
		Update: updateManagementHost,
		Delete: deleteManagementHost,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"ipv4_address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv6 address.",
			},
			"interfaces": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Host interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"subnet4": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 network address.",
						},
						"subnet6": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 network address.",
						},
						"mask_length4": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "IPv4 network mask length.",
						},
						"mask_length6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "IPv6 network mask length.",
						},
						"ignore_warnings": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Apply changes ignoring warnings.",
						},
						"ignore_errors": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
						},
						"color": {
							Type:        schema.TypeString,
							Optional:    true,
							Default:     "black",
							Description: "Color of the object. Should be one of existing colors.",
						},
						"comments": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Comments string.",
						},
					},
				},
			},
			"nat_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "NAT settings.",
				//Default: map[string]interface{}{"auto_rule":false},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether to add automatic address translation rules.",
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
						"hide_behind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"host_servers": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Servers Configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dns_server": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Gets True if this server is a DNS Server.",
						},
						"mail_server": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Gets True if this server is a Mail Server.",
						},
						"web_server": &schema.Schema{
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Gets True if this server is a Web Server.",
						},
						"web_server_config": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Web Server configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"additional_ports": &schema.Schema{
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Server additional ports.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"application_engines": &schema.Schema{
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Application engines of this web server.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
									"listen_standard_port": &schema.Schema{
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
										Description: "Whether server listens to standard port.",
									},
									"operating_system": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "other",
										Description: "Operating System.",
									},
									"protected_by": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "97aeb368-9aea-11d5-bd16-0090272ccb30",
										Description: "Network object which protects this server identified by the name or UID.",
									},
								},
							},
						},
					},
				},
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
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
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

				if v, ok := d.GetOk("host_servers.0.web_server_config.0.additional_ports"); ok {
					webServerConfigPayLoad["additional-ports"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("host_servers.0.web_server_config.0.application_engines"); ok {
					webServerConfigPayLoad["application-engines"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOkExists("host_servers.0.web_server_config.0.listen_standard_port"); ok {
					webServerConfigPayLoad["listen-standard-port"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("host_servers.0.web_server_config.0.operating_system"); ok {
					webServerConfigPayLoad["operating-system"] = v.(string)
				}
				if v, ok := d.GetOk("host_servers.0.web_server_config.0.protected_by"); ok {
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

	if val, ok := d.GetOk("color"); ok {
		host["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		host["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		host["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Host - Map = ", host)

	addHostRes, err := client.ApiCall("add-host", host, client.GetSessionID(), true, false)
	if err != nil || !addHostRes.Success {
		if addHostRes.ErrorMsg != "" {
			return fmt.Errorf(addHostRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addHostRes.GetData()["uid"].(string))

	return readManagementHost(d, m)
}

func readManagementHost(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showHostRes, err := client.ApiCall("show-host", payload, client.GetSessionID(), true, false)
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

	//we are compromising here since we cant represent map inside map
	//see also  https://github.com/hashicorp/terraform-plugin-sdk/issues/155
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
				interfaceMapToAdd["ignore_errors"] = false
				interfaceMapToAdd["ignore_warnings"] = false

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

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		host["name"] = oldName
		host["new-name"] = newName
	} else {
		host["name"] = d.Get("name")
	}

	if d.HasChange("interfaces") {

		if v, ok := d.GetOk("interfaces"); ok {

			interfacesList := v.([]interface{})

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
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".ignore_warnings") {
					payload["ignore-warnings"] = d.Get("interfaces." + strconv.Itoa(i) + ".ignore_warnings")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".color") {
					payload["color"] = d.Get("interfaces." + strconv.Itoa(i) + ".color")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".ignore_errors") {
					payload["ignore-errors"] = d.Get("interfaces." + strconv.Itoa(i) + ".ignore_errors")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".comments") {
					payload["comments"] = d.Get("interfaces." + strconv.Itoa(i) + ".comments")
				}
				interfacesPayload = append(interfacesPayload, payload)
			}

			host["interfaces"] = interfacesPayload

		} else { //delete all of the list
			oldInterfaces, _ := d.GetChange("interfaces")
			var interfacesToDelete []interface{}
			for _, inter := range oldInterfaces.([]interface{}) {
				interfacesToDelete = append(interfacesToDelete, inter.(map[string]interface{})["name"].(string))
			}
			host["interfaces"] = map[string]interface{}{"remove": interfacesToDelete}
		}
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		host["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		host["ignore-warnings"] = v.(bool)
	}

	if ok := d.HasChange("comments"); ok {
		host["comments"] = d.Get("comments")
	}

	if ok := d.HasChange("color"); ok {
		host["color"] = d.Get("color")
	}

	if ok := d.HasChange("ipv4_address"); ok {
		host["ipv4-address"] = d.Get("ipv4_address")
	}

	if ok := d.HasChange("ipv6_address"); ok {
		host["ipv6-address"] = d.Get("ipv6_address")
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
				res["hide-behind"] = d.Get("nat_settings.hide_behind")
			}
			if d.HasChange("nat_settings.install_on") {
				res["install-on"] = d.Get("nat_settings.install_on")
			}
			if d.HasChange("nat_settings.method") {
				res["method"] = d.Get("nat_settings.method")
			}

			host["nat-settings"] = res
		} else { //argument deleted - go back to defaults
			host["nat-settings"] = map[string]interface{}{"auto-rule": "false"}
		}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			host["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			host["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
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
				if d.HasChange("host_servers.0.web_server") {
					hostServersPayload["web-server"] = d.Get("host_servers.0.web_server").(bool)
				}

				if d.HasChange("host_servers.0.web_server_config") {

					hostServersPayload["web-server"] = d.Get("host_servers.0.web_server").(bool)
					webServerConfigPayLoad := make(map[string]interface{})

					if d.HasChange("host_servers.0.web_server_config.0.additional_ports") {
						webServerConfigPayLoad["additional-ports"] = d.Get("host_servers.0.web_server_config.0.additional_ports").(*schema.Set).List()
					}
					if d.HasChange("host_servers.0.web_server_config.0.application_engines") {
						webServerConfigPayLoad["application-engines"] = d.Get("host_servers.0.web_server_config.0.application_engines").(*schema.Set).List()
					}
					//boolean nested field - isn't recognized by diff on the first time created
					webServerConfigPayLoad["listen-standard-port"] = d.Get("host_servers.0.web_server_config.0.listen_standard_port")

					if d.HasChange("host_servers.0.web_server_config.0.operating_system") {
						webServerConfigPayLoad["operating-system"] = d.Get("host_servers.0.web_server_config.0.operating_system").(string)
					}
					if d.HasChange("host_servers.0.web_server_config.0.protected_by") {
						webServerConfigPayLoad["protected-by"] = d.Get("host_servers.0.web_server_config.0.protected_by").(string)
					}
					hostServersPayload["web-server-config"] = webServerConfigPayLoad
				}
				host["host-servers"] = hostServersPayload
			}
		} else { // argument deleted - go back to defaults
			host["host-servers"] = map[string]interface{}{"dns-server": false, "mail-server": false, "web-server": false}
		}
	}

	log.Println("Update Host - Map = ", host)
	updateHostRes, err := client.ApiCall("set-host", host, client.GetSessionID(), true, false)
	if err != nil || !updateHostRes.Success {
		if updateHostRes.ErrorMsg != "" {
			return fmt.Errorf(updateHostRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementHost(d, m)
}

func deleteManagementHost(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	hostPayload := map[string]interface{}{
		"uid": d.Id(),
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
