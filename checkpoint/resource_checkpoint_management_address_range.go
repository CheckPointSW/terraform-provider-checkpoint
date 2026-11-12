package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/v3/upgraders"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceManagementAddressRange() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAddressRange,
		Read:   readManagementAddressRange,
		Update: updateManagementAddressRange,
		Delete: deleteManagementAddressRange,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementAddressRangeV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementAddressRangeStateUpgradeV0,
				Version: 0,
			},
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"ipv4_address_first": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "First IPv4 address in the range.",
			},
			"ipv6_address_first": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "First IPv6 address in the range.",
			},
			"ipv4_address_last": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Last IPv4 address in the range.",
			},
			"ipv6_address_last": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Last IPv6 address in the range.",
			},
			"nat_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
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

func createManagementAddressRange(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	addressRange := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		addressRange["name"] = v.(string)
	}
	if v, ok := d.GetOk("ipv4_address_first"); ok {
		addressRange["ipv4-address-first"] = v.(string)
	}
	if v, ok := d.GetOk("ipv6_address_first"); ok {
		addressRange["ipv6-address-first"] = v.(string)
	}
	if v, ok := d.GetOk("ipv4_address_last"); ok {
		addressRange["ipv4-address-last"] = v.(string)
	}
	if v, ok := d.GetOk("ipv6_address_last"); ok {
		addressRange["ipv6-address-last"] = v.(string)
	}

	if v, ok := d.GetOk("nat_settings"); ok {

		natSettingsList := v.([]interface{})

		if len(natSettingsList) > 0 {

			natSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOkExists("nat_settings.0.auto_rule"); ok {
				natSettingsPayload["auto-rule"] = v.(bool)
			}
			if v, ok := d.GetOk("nat_settings.0.ipv4_address"); ok {
				natSettingsPayload["ipv4-address"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.0.ipv6_address"); ok {
				natSettingsPayload["ipv6-address"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.0.hide_behind"); ok {
				natSettingsPayload["hide-behind"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.0.install_on"); ok {
				natSettingsPayload["install-on"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.0.method"); ok {
				natSettingsPayload["method"] = v.(string)
			}
			addressRange["nat-settings"] = natSettingsPayload
		}
	}

	if val, ok := d.GetOk("comments"); ok {
		addressRange["comments"] = val.(string)
	}
	if val, ok := d.GetOk("tags"); ok {
		addressRange["tags"] = val.(*schema.Set).List()
	}

	if val, ok := d.GetOk("color"); ok {
		addressRange["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		addressRange["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		addressRange["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Address Range - Map = ", addressRange)

	addAddressRangeRes, err := client.ApiCall("add-address-range", addressRange, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addAddressRangeRes.Success {
		if addAddressRangeRes.ErrorMsg != "" {
			return fmt.Errorf("%s", addAddressRangeRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}

	d.SetId(addAddressRangeRes.GetData()["uid"].(string))

	return readManagementAddressRange(d, m)
}

func readManagementAddressRange(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showAddressRangeRes, err := client.ApiCall("show-address-range", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showAddressRangeRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showAddressRangeRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showAddressRangeRes.ErrorMsg)
	}

	addressRange := showAddressRangeRes.GetData()

	log.Println("Read Address Range - Show JSON = ", addressRange)

	if v := addressRange["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := addressRange["ipv4-address-first"]; v != nil {
		_ = d.Set("ipv4_address_first", v)
	}

	if v := addressRange["ipv6-address-first"]; v != nil {
		_ = d.Set("ipv6_address_first", v)
	}

	if v := addressRange["ipv4-address-last"]; v != nil {
		_ = d.Set("ipv4_address_last", v)
	}

	if v := addressRange["ipv6-address-last"]; v != nil {
		_ = d.Set("ipv6_address_last", v)
	}

	if v := addressRange["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := addressRange["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if addressRange["nat-settings"] != nil {

		natSettingsMap := addressRange["nat-settings"].(map[string]interface{})

		natSettingsMapToReturn := make(map[string]interface{})

		if v := natSettingsMap["auto-rule"]; v != nil {
			natSettingsMapToReturn["auto_rule"] = v
		}
		if v := natSettingsMap["ipv4-address"]; v != nil {
			natSettingsMapToReturn["ipv4_address"] = v
		}
		if v := natSettingsMap["ipv6-address"]; v != nil {
			natSettingsMapToReturn["ipv6_address"] = v
		}
		if v := natSettingsMap["hide-behind"]; v != nil {
			natSettingsMapToReturn["hide_behind"] = v
		}
		if v := natSettingsMap["install-on"]; v != nil {
			natSettingsMapToReturn["install_on"] = v
		}
		if v := natSettingsMap["method"]; v != nil {
			natSettingsMapToReturn["method"] = v
		}

		_ = d.Set("nat_settings", []interface{}{natSettingsMapToReturn})

	} else {
		_ = d.Set("nat_settings", nil)
	}

	if addressRange["tags"] != nil {
		tagsJson := addressRange["tags"].([]interface{})
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

func updateManagementAddressRange(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	addressRange := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		addressRange["name"] = oldName
		addressRange["new-name"] = newName
	} else {
		addressRange["name"] = d.Get("name")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		addressRange["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		addressRange["ignore-warnings"] = v.(bool)
	}

	if ok := d.HasChange("comments"); ok {
		addressRange["comments"] = d.Get("comments")
	}
	if ok := d.HasChange("color"); ok {
		addressRange["color"] = d.Get("color")
	}

	if ok := d.HasChange("ipv4_address_first"); ok {
		addressRange["ipv4-address-first"] = d.Get("ipv4_address_first")
	}
	if ok := d.HasChange("ipv6_address_first"); ok {
		addressRange["ipv6-address-first"] = d.Get("ipv6_address_first")
	}
	if ok := d.HasChange("ipv4_address_last"); ok {
		addressRange["ipv4-address-last"] = d.Get("ipv4_address_last")
	}
	if ok := d.HasChange("ipv6_address_last"); ok {
		addressRange["ipv6-address-last"] = d.Get("ipv6_address_last")
	}

	if ok := d.HasChange("nat_settings"); ok {

		if v, ok := d.GetOk("nat_settings"); ok {

			natSettingsList := v.([]interface{})

			if len(natSettingsList) > 0 {

				natSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("nat_settings.0.auto_rule"); ok {
					natSettingsPayload["auto-rule"] = v.(bool)
				}
				if v, ok := d.GetOk("nat_settings.0.ipv4_address"); ok {
					natSettingsPayload["ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("nat_settings.0.ipv6_address"); ok {
					natSettingsPayload["ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("nat_settings.0.hide_behind"); ok {
					natSettingsPayload["hide-behind"] = v.(string)
				}
				if v, ok := d.GetOk("nat_settings.0.install_on"); ok {
					natSettingsPayload["install-on"] = v.(string)
				}
				if v, ok := d.GetOk("nat_settings.0.method"); ok {
					natSettingsPayload["method"] = v.(string)
				}
				addressRange["nat-settings"] = natSettingsPayload
			}
		}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			addressRange["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			addressRange["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	log.Println("Update Address Range - Map = ", addressRange)
	updateAddressRangeRes, err := client.ApiCall("set-address-range", addressRange, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateAddressRangeRes.Success {
		if updateAddressRangeRes.ErrorMsg != "" {
			return fmt.Errorf("%s", updateAddressRangeRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}

	return readManagementAddressRange(d, m)
}

func deleteManagementAddressRange(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	addressRangePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		addressRangePayload["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		addressRangePayload["ignore-warnings"] = v.(bool)
	}

	deleteAddressRangeRes, err := client.ApiCall("delete-address-range", addressRangePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteAddressRangeRes.Success {
		if deleteAddressRangeRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteAddressRangeRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
