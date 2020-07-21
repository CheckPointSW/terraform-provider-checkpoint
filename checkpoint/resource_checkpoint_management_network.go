package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementNetwork() *schema.Resource {
	return &schema.Resource{
		Create: createManagementNetwork,
		Read:   readManagementNetwork,
		Update: updateManagementNetwork,
		Delete: deleteManagementNetwork,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
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
			"nat_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Optional:    true,
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"broadcast": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Allow broadcast address inclusion.",
				Default:     "allow",
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

func createManagementNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	network := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		network["name"] = v.(string)
	}
	if val, ok := d.GetOk("subnet4"); ok {
		network["subnet4"] = val.(string)
	}
	if val, ok := d.GetOk("subnet6"); ok {
		network["subnet6"] = val.(string)
	}
	if val, ok := d.GetOk("mask_length4"); ok {
		network["mask-length4"] = val.(int)
	}
	if val, ok := d.GetOk("mask_length6"); ok {
		network["mask-length6"] = val.(int)
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
		network["nat-settings"] = res
	}

	if val, ok := d.GetOk("tags"); ok {
		network["tags"] = val.(*schema.Set).List()
	}

	if val, ok := d.GetOk("broadcast"); ok {
		network["broadcast"] = val.(string)
	}
	if val, ok := d.GetOk("comments"); ok {
		network["comments"] = val.(string)
	}
	if val, ok := d.GetOk("color"); ok {
		network["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		network["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		network["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Network - Map = ", network)

	addNetworkRes, err := client.ApiCall("add-network", network, client.GetSessionID(), true, false)
	if err != nil || !addNetworkRes.Success {
		if addNetworkRes.ErrorMsg != "" {
			return fmt.Errorf(addNetworkRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addNetworkRes.GetData()["uid"].(string))

	return readManagementNetwork(d, m)
}

func readManagementNetwork(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showNetworkRes, err := client.ApiCall("show-network", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNetworkRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showNetworkRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showNetworkRes.ErrorMsg)
	}

	network := showNetworkRes.GetData()

	if v := network["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := network["subnet4"]; v != nil {
		_ = d.Set("subnet4", v)
	}

	if v := network["subnet6"]; v != nil {
		_ = d.Set("subnet6", v)
	}

	if v := network["mask-length4"]; v != nil {
		_ = d.Set("mask_length4", v)
	}

	if v := network["mask-length6"]; v != nil {
		_ = d.Set("mask_length6", v)
	}

	if v := network["subnet-mask"]; v != nil {
		_ = d.Set("subnet_mask", v)
	}

	if v := network["broadcast"]; v != nil {
		_ = d.Set("broadcast", v)
	}

	if v := network["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := network["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if network["nat-settings"] != nil {

		natSettingsMap := network["nat-settings"].(map[string]interface{})

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

	if network["tags"] != nil {
		tagsJson := network["tags"].([]interface{})
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

func updateManagementNetwork(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	network := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		network["name"] = oldName.(string)
		network["new-name"] = newName.(string)
	} else {
		network["name"] = d.Get("name")
	}

	if ok := d.HasChange("subnet4"); ok {
		network["subnet4"] = d.Get("subnet4")
	}
	if ok := d.HasChange("subnet6"); ok {
		network["subnet6"] = d.Get("subnet6")
	}
	if ok := d.HasChange("mask_length4"); ok {
		network["mask-length4"] = d.Get("mask_length4")
	}
	if ok := d.HasChange("mask_length6"); ok {
		network["mask-length6"] = d.Get("mask_length6")
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

			network["nat-settings"] = res
		} else { //argument deleted - go back to defaults
			network["nat-settings"] = map[string]interface{}{"auto-rule": "false"}
		}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			network["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			network["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("broadcast"); ok {
		network["broadcast"] = d.Get("broadcast")
	}
	if ok := d.HasChange("comments"); ok {
		network["comments"] = d.Get("comments")
	}
	if ok := d.HasChange("color"); ok {
		network["color"] = d.Get("color")
	}
	if v, ok := d.GetOkExists("ignore_errors"); ok {
		network["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		network["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Network - Map = ", network)
	setNetworkRes, _ := client.ApiCall("set-network", network, client.GetSessionID(), true, false)
	if !setNetworkRes.Success {
		return fmt.Errorf(setNetworkRes.ErrorMsg)
	}

	return readManagementNetwork(d, m)
}

func deleteManagementNetwork(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	deleteNetworkRes, _ := client.ApiCall("delete-network", payload, client.GetSessionID(), true, false)
	if !deleteNetworkRes.Success {
		return fmt.Errorf(deleteNetworkRes.ErrorMsg)
	}
	d.SetId("")

	return nil
}
