package main

import (
	chkp "api_go_sdk/APIFiles"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
)


func resourceNetwork() *schema.Resource {
	return &schema.Resource{
		Create: createNetwork,
		Read:   readNetwork,
		Update: updateNetwork,
		Delete: deleteNetwork,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				Description: "Object name. Should be unique in the domain.",
			},
			"subnet": {
				Type: schema.TypeString,
				Optional: true,
				Description: "IPv4 or IPv6 network address. If both addresses are required use subnet4 and subnet6 fields explicitly.",
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
			"mask_length": {
				Type: schema.TypeInt,
				Optional: true,
				Description: "IPv4 or IPv6 network mask length. If both masks are required use mask-length4 and mask-length6 fields explicitly. Instead of IPv4 mask length it is possible to specify IPv4 mask itself in subnet-mask field.",
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
			"subnet_mask": {
				Type: schema.TypeString,
				Optional: true,
				Description: "IPv4 network mask.",
			},
			"nat_settings" : {
				Type: schema.TypeSet,
				Optional: true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:     schema.TypeBool,
							Optional: true,
							Description: "Whether to add automatic address translation rules.",
						},
						"ip_address": {
							Type:     schema.TypeString,
							Optional: true,
							Description: "IPv4 or IPv6 address. If both addresses are required use ipv4-address and ipv6-address fields explicitly. This parameter is not required in case \"method\" parameter is \"hide\" and \"hide-behind\" parameter is \"gateway\"",
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
			"tags": {
				Type: schema.TypeList,
				Optional: true,
				Description: "Collection of tag name.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"broadcast": {
				Type: schema.TypeString,
				Optional: true,
				Description: "Allow broadcast address inclusion.",
			},
			"color": {
				Type:         schema.TypeString,
				Optional:     true,
				Description:  "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"details_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The level of detail for some of the fields in the response can vary from showing only the UID value of the object to a fully detailed representation of the object.",
			},
			"groups": {
				Type: schema.TypeList,
				Optional: true,
				Description: "Collection of group name.",
				Elem: &schema.Schema {
					Type: schema.TypeString,
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
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func parseSchemaToMap(d *schema.ResourceData, createResource bool) map[string]interface{} {
	networkMap := make(map[string]interface{})

	// Parsing only attributes that were set by client or default value attributes.
	if val, ok := d.GetOk("name"); ok {
		networkMap["name"] = val.(string)
	}
	// subnet handling
	if val, ok := d.GetOk("subnet"); ok {
		networkMap["subnet"] = val.(string)
	}
	if val, ok := d.GetOk("subnet4"); ok {
		networkMap["subnet4"] = val.(string)
	}
	if val, ok := d.GetOk("subnet6"); ok {
		networkMap["subnet6"] = val.(string)
	}

	// mask handling
	if val, ok := d.GetOk("mask_length"); ok {
		networkMap["mask-length"] = val.(int)
	}
	if val, ok := d.GetOk("mask_length4"); ok {
		networkMap["mask-length4"] = val.(int)
	}
	if val, ok := d.GetOk("mask_length6"); ok {
		networkMap["mask-length6"] = val.(int)
	}
	if val, ok := d.GetOk("subnet_mask"); ok {
		networkMap["subnet-mask"] = val.(string)
	}

	// nat settings handling
	if val, ok := d.GetOk("nat_settings"); ok {
		natSettingsSchemaMap := val.(*schema.Set).List()[0].(map[string]interface{}) // v[0] = Nat settings item cast as map
		natSettingsConf := make(map[string]interface{})

		natSettingsConf["auto-rule"] = natSettingsSchemaMap["auto_rule"].(bool)
		if v := natSettingsSchemaMap["ip_address"].(string); v != "" {
				natSettingsConf["ip-address"] = v
		}
		if v := natSettingsSchemaMap["ipv4_address"].(string); v != "" {
				natSettingsConf["ipv4-address"] = v
		}
		if v := natSettingsSchemaMap["ipv6_address"].(string); v != ""{
				natSettingsConf["ipv6-address"] = v
		}
		if v := natSettingsSchemaMap["hide_behind"].(string); v != "" {
				natSettingsConf["hide-behind"] = v
		}
		if v := natSettingsSchemaMap["install_on"].(string); v != "" {
				natSettingsConf["install-on"] = v
		}
		if v := natSettingsSchemaMap["method"].(string); v != "" {
				natSettingsConf["method"] = v
		}
		networkMap["nat-settings"] = natSettingsConf
	}
	if val, ok := d.GetOk("tags"); ok {
		networkMap["tags"] = val.([]interface{})
	}
	if val, ok := d.GetOk("groups"); ok {
		networkMap["groups"] = val.([]interface{})
	}
	if val, ok := d.GetOk("broadcast"); ok {
		networkMap["broadcast"] = val.(string)
	}
	if val, ok := d.GetOk("comments"); ok {
		networkMap["comments"] = val.(string)
	}
	if val, ok := d.GetOk("set_if_exists"); ok {
		networkMap["set-if-exists"] = val.(bool)
	}
	if val, ok := d.GetOk("color"); ok {
		networkMap["color"] = val.(string)
	}
	if val, ok := d.GetOk("details_level"); ok {
		networkMap["details-level"] = val.(bool)
	}
	if val, ok := d.GetOk("ignore_errors"); ok {
		networkMap["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOk("ignore_warnings"); ok {
		networkMap["ignore-warnings"] = val.(bool)
	}

	if !createResource {
		// Call from updateNetwork
		// Remove attributes that cannot be set from map - Schema contain ADD + SET attr.
		delete(networkMap,"set-if-exists")
	}
	return networkMap
}

func createNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*chkp.ApiClient)
	payload := parseSchemaToMap(d, true)
	log.Println(payload)
	addNetworkRes, _ := client.ApiCall("add-network",payload,client.GetSessionID(),true,false)
	if !addNetworkRes.Success {
		return fmt.Errorf(addNetworkRes.ErrorMsg)
	}

	// Set Schema UID = Object UID
	d.SetId(addNetworkRes.GetData()["uid"].(string))
	return readNetwork(d, m)
}

func readNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*chkp.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	showNetworkRes, _ := client.ApiCall("show-network",payload,client.GetSessionID(),true,false)
	if !showNetworkRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showNetworkRes.GetData()["code"].(string)) {
			d.SetId("") // Destroy resource
			return nil
		}
		return fmt.Errorf(showNetworkRes.ErrorMsg)
	}
	networkJson := showNetworkRes.GetData()
	log.Println(networkJson)
	// Handle non-nested attributes
	if _, ok := d.GetOk("name"); ok {
		_ = d.Set("name", networkJson["name"].(string))
	}
	if _, ok := d.GetOk("subnet4"); ok {
		_ = d.Set("subnet4", networkJson["subnet4"].(string))
	}
	if _, ok := d.GetOk("subnet6"); ok {
		_ = d.Set("subnet6", networkJson["subnet6"].(string))
	}
	if _, ok := d.GetOk("mask_length4"); ok {
		_ = d.Set("mask_length4", int(networkJson["mask-length4"].(float64)))
	}
	if _, ok := d.GetOk("mask_length6"); ok {
		_ = d.Set("mask_length6", int(networkJson["mask-length6"].(float64)))
	}
	if _, ok := d.GetOk("subnet_mask"); ok {
		_ = d.Set("subnet_mask", networkJson["subnet-mask"].(string))
	}
	if _, ok := d.GetOk("broadcast"); ok {
		_ = d.Set("broadcast", networkJson["broadcast"].(string))
	}
	if _, ok := d.GetOk("color"); ok {
		_ = d.Set("color", networkJson["color"].(string))
	}
	if _, ok := d.GetOk("comments"); ok {
		_ = d.Set("comments", networkJson["comments"].(string))
	}

	// Handle nat settings
	if _, ok := d.GetOk("nat_settings"); ok {
		var natSettingsConf []interface{}
		natSettingsConf = append(natSettingsConf, networkJson["nat-settings"].(map[string]interface{}))
		_ = d.Set("nat_settings", natSettingsConf)
	}

	// Handle groups. Creates slice of groups uid
	if _, ok := d.GetOk("groups"); ok {
		var groupsIds []interface{}
		groupsJson := networkJson["groups"].([]interface{})
		if len(groupsJson) != 0 {
			for _, group := range groupsJson {
				group := group.(map[string]interface{})
				groupsIds = append(groupsIds, group["name"].(string))
			}
		}
		_ = d.Set("groups", groupsIds)
	}

	// Handle tags. Creates slice of tags uid
	if _, ok := d.GetOk("tags"); ok {
		var tagsIds []interface{}
		tagsJson := networkJson["tags"].([]interface{})
		if len(tagsJson) != 0 {
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	}
	return nil
}

func updateNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*chkp.ApiClient)
	payload := parseSchemaToMap(d, false)
	if d.HasChange("name") {
		oldName , newName := d.GetChange("name")
		payload["name"] = oldName
		payload["new-name"] = newName
	}
	log.Println(payload)
	setNetworkRes, _ := client.ApiCall("set-network",payload,client.GetSessionID(),true,false)
	if !setNetworkRes.Success {
		return fmt.Errorf(setNetworkRes.ErrorMsg)
	}
	return readNetwork(d, m)
}

func deleteNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*chkp.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	deleteNetworkRes, _ := client.ApiCall("delete-network",payload,client.GetSessionID(),true,false)
	if !deleteNetworkRes.Success {
		return fmt.Errorf(deleteNetworkRes.ErrorMsg)
	}
	d.SetId("") // Destroy resource
	return nil
}
