package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementResourceTcp() *schema.Resource {
	return &schema.Resource{
		Create: createManagementResourceTcp,
		Read:   readManagementResourceTcp,
		Update: updateManagementResourceTcp,
		Delete: deleteManagementResourceTcp,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the TCP resource.",
				Default:     "ufp",
			},
			"exception_track": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Configures how to track connections that match this rule but fail the content security checks.",
				Default:     "None",
			},
			"ufp_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "UFP settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "UFP server identified by name or UID.",
						},
						"caching_control": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies if and how caching is to be enabled.",
							Default:     "no_caching",
						},
						"ignore_ufp_server_after_failure": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The UFP server will be ignored after numerous UFP server connections were unsuccessful.",
							Default:     false,
						},
						"number_of_failures_before_ignore": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Signifies at what point the UFP server should be ignored, Applicable only if 'ignore after fail' is enabled.",
							Default:     0,
						},
						"timeout_before_reconnecting": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The amount of time, in seconds, that must pass before a UFP server connection should be attempted, Applicable only if 'ignore after fail' is enabled.",
							Default:     0,
						},
					},
				},
			},
			"cvp_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "CVP settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "CVP server identified by name or UID. The CVP server must already be defined as an OPSEC Application.",
						},
						"allowed_to_modify_content": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Configures the CVP server to inspect but not modify content.",
							Default:     true,
						},
						"reply_order": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Designates when the CVP server returns data to the Security Gateway security server.",
							Default:     "return_data_after_content_is_approved",
						},
					},
				},
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
		},
	}
}

func createManagementResourceTcp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	resourceTcp := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		resourceTcp["name"] = v.(string)
	}

	if v, ok := d.GetOk("resource_type"); ok {
		resourceTcp["resource-type"] = v.(string)
	}

	if v, ok := d.GetOk("exception_track"); ok {
		resourceTcp["exception-track"] = v.(string)
	}

	if v, ok := d.GetOk("ufp_settings"); ok {

		res := make(map[string]interface{})

		v := v.([]interface{})

		ufpSettingsMap := v[0].(map[string]interface{})

		if v := ufpSettingsMap["server"]; v != nil {
			res["server"] = v
		}
		if v := ufpSettingsMap["caching_control"]; v != nil {
			res["caching-control"] = v
		}
		if v := ufpSettingsMap["ignore_ufp_server_after_failure"]; v != nil {
			res["ignore-ufp-server-after-failure"] = v
		}
		if v := ufpSettingsMap["number_of_failures_before_ignore"]; v != nil {
			res["number-of-failures-before-ignore"] = v
		}
		if v := ufpSettingsMap["timeout_before_reconnecting"]; v != nil {
			res["timeout-before-reconnecting"] = v
		}

		resourceTcp["ufp-settings"] = res
	}

	if v, ok := d.GetOk("cvp_settings"); ok {

		res := make(map[string]interface{})

		v := v.([]interface{})

		cvpSettingsMap := v[0].(map[string]interface{})

		if v := cvpSettingsMap["server"]; v != nil {
			if len(v.(string)) > 0 {
				res["server"] = v
			}

		}
		if v := cvpSettingsMap["allowed_to_modify_content"]; v != nil {
			res["allowed-to-modify-content"] = v
		}
		if v := cvpSettingsMap["reply_order"]; v != nil {
			res["reply-order"] = v
		}

		resourceTcp["cvp-settings"] = res
	}

	if v, ok := d.GetOk("tags"); ok {
		resourceTcp["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		resourceTcp["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		resourceTcp["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceTcp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceTcp["ignore-errors"] = v.(bool)
	}

	log.Println("Create ResourceTcp - Map = ", resourceTcp)

	addResourceTcpRes, err := client.ApiCallSimple("add-resource-tcp", resourceTcp)
	if err != nil || !addResourceTcpRes.Success {
		if addResourceTcpRes.ErrorMsg != "" {
			return fmt.Errorf(addResourceTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addResourceTcpRes.GetData()["uid"].(string))

	return readManagementResourceTcp(d, m)
}

func readManagementResourceTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showResourceTcpRes, err := client.ApiCallSimple("show-resource-tcp", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceTcpRes.Success {
		if objectNotFound(showResourceTcpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showResourceTcpRes.ErrorMsg)
	}

	resourceTcp := showResourceTcpRes.GetData()

	log.Println("Read ResourceTcp - Show JSON = ", resourceTcp)

	if v := resourceTcp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceTcp["resource-type"]; v != nil {
		_ = d.Set("resource_type", v)
	}

	if v := resourceTcp["exception-track"]; v != nil {
		_ = d.Set("exception_track", v)
	}

	if resourceTcp["ufp-settings"] != nil {

		ufpSettingsMap := resourceTcp["ufp-settings"].(map[string]interface{})

		ufpSettingsMapToReturn := make(map[string]interface{})

		if v, _ := ufpSettingsMap["server"]; v != nil {
			ufpSettingsMapToReturn["server"] = v
		}
		if v, _ := ufpSettingsMap["caching-control"]; v != nil {
			ufpSettingsMapToReturn["caching_control"] = v
		}
		if v, _ := ufpSettingsMap["ignore-ufp-server-after-failure"]; v != nil {
			ufpSettingsMapToReturn["ignore_ufp_server_after_failure"] = v
		}
		if v, _ := ufpSettingsMap["number-of-failures-before-ignore"]; v != nil {
			ufpSettingsMapToReturn["number_of_failures_before_ignore"] = v
		}
		if v, _ := ufpSettingsMap["timeout-before-reconnecting"]; v != nil {
			ufpSettingsMapToReturn["timeout_before_reconnecting"] = v
		}

		_ = d.Set("ufp_settings", []interface{}{ufpSettingsMapToReturn})
	} else {
		_ = d.Set("ufp_settings", nil)
	}

	if resourceTcp["cvp-settings"] != nil {

		cvpMap := resourceTcp["cvp-settings"].(map[string]interface{})

		cvpMapToReturn := make(map[string]interface{})

		if v, _ := cvpMap["server"]; v != nil {
			objMap := v.(map[string]interface{})
			if v := objMap["server"]; v != nil {
				cvpMapToReturn["server"] = v
			}
		}
		if v, _ := cvpMap["cvp-server-is-allowed-to-modify-content"]; v != nil {
			cvpMapToReturn["allowed_to_modify_content"] = v
		}
		if v, _ := cvpMap["reply-order"]; v != nil {
			cvpMapToReturn["reply_order"] = v
		}
		_ = d.Set("cvp_settings", []interface{}{cvpMapToReturn})
	} else {
		_ = d.Set("cvp_settings", nil)
	}

	if resourceTcp["tags"] != nil {
		tagsJson, ok := resourceTcp["tags"].([]interface{})
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

	if v := resourceTcp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceTcp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := resourceTcp["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := resourceTcp["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementResourceTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	resourceTcp := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		resourceTcp["name"] = oldName
		resourceTcp["new-name"] = newName
	} else {
		resourceTcp["name"] = d.Get("name")
	}

	if ok := d.HasChange("resource_type"); ok {
		resourceTcp["resource-type"] = d.Get("resource_type")
	}

	if ok := d.HasChange("exception_track"); ok {
		resourceTcp["exception-track"] = d.Get("exception_track")
	}

	if d.HasChange("ufp_settings") {

		if v, ok := d.GetOk("ufp_settings"); ok {

			res := make(map[string]interface{})

			v := v.([]interface{})

			ufpSettingsMap := v[0].(map[string]interface{})

			if v := ufpSettingsMap["server"]; v != nil {
				if len(v.(string)) > 0 {
					res["server"] = v
				}
			}
			if v := ufpSettingsMap["caching_control"]; v != nil {
				res["caching-control"] = v
			}
			if v := ufpSettingsMap["ignore_ufp_server_after_failure"]; v != nil {
				res["ignore-ufp-server-after-failure"] = v
			}
			if v := ufpSettingsMap["number_of_failures_before_ignore"]; v != nil {
				res["number-of-failures-before-ignore"] = v
			}
			if v := ufpSettingsMap["timeout_before_reconnecting"]; v != nil {
				res["timeout-before-reconnecting"] = v
			}

			resourceTcp["ufp-settings"] = res
		}
	}

	if d.HasChange("cvp_settings") {

		if v, ok := d.GetOk("cvp_settings"); ok {

			res := make(map[string]interface{})

			v := v.([]interface{})

			cvpSettingsMap := v[0].(map[string]interface{})

			if v := cvpSettingsMap["server"]; v != nil {
				if len(v.(string)) > 0 {
					res["server"] = v
				}
			}
			if v := cvpSettingsMap["allowed_to_modify_content"]; v != nil {
				res["allowed-to-modify-content"] = v
			}
			if v := cvpSettingsMap["reply_order"]; v != nil {
				res["reply-order"] = v
			}

			resourceTcp["cvp-settings"] = res
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			resourceTcp["tags"] = v.(*schema.Set).List()
		}
	}

	if ok := d.HasChange("color"); ok {
		resourceTcp["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		resourceTcp["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceTcp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceTcp["ignore-errors"] = v.(bool)
	}

	log.Println("Update ResourceTcp - Map = ", resourceTcp)

	updateResourceTcpRes, err := client.ApiCallSimple("set-resource-tcp", resourceTcp)
	if err != nil || !updateResourceTcpRes.Success {
		if updateResourceTcpRes.ErrorMsg != "" {
			return fmt.Errorf(updateResourceTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementResourceTcp(d, m)
}

func deleteManagementResourceTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	resourceTcpPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ResourceTcp")

	deleteResourceTcpRes, err := client.ApiCallSimple("delete-resource-tcp", resourceTcpPayload)
	if err != nil || !deleteResourceTcpRes.Success {
		if deleteResourceTcpRes.ErrorMsg != "" {
			return fmt.Errorf(deleteResourceTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
