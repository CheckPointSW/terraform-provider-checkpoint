package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementResourceTcp() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementResourceTcpRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object uid.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"resource_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the TCP resource.",
			},
			"exception_track": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configures how to track connections that match this rule but fail the content security checks. Identified by name or UID,",
			},
			"ufp_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "UFP settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UFP server identified by name or UID.",
						},
						"caching_control": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies if and how caching is to be enabled.",
						},
						"ignore_ufp_server_after_failure": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The UFP server will be ignored after numerous UFP server connections were unsuccessful.",
						},
						"number_of_failures_before_ignore": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Signifies at what point the UFP server should be ignored, Applicable only if 'ignore after fail' is enabled.",
						},
						"timeout_before_reconnecting": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The amount of time, in seconds, that must pass before a UFP server connection should be attempted, Applicable only if 'ignore after fail' is enabled.",
						},
					},
				},
			},
			"cvp_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Configure CVP inspection on mail messages.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CVP server identified by name or UID. The CVP server must already be defined as an OPSEC Application.",
						},
						"allowed_to_modify_content": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Configures the CVP server to inspect but not modify content.",
						},
						"reply_order": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Designates when the CVP server returns data to the Security Gateway security server.",
						},
					},
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
	}
}

func dataSourceManagementResourceTcpRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showResourceTcpRes, err := client.ApiCallSimple("show-resource-tcp", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceTcpRes.Success {
		return fmt.Errorf(showResourceTcpRes.ErrorMsg)
	}

	resourceTcp := showResourceTcpRes.GetData()

	log.Println("Read ResourceTcp - Show JSON = ", resourceTcp)

	if v := resourceTcp["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := resourceTcp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceTcp["resource-type"]; v != nil {
		_ = d.Set("resource_type", v)
	}

	if v := resourceTcp["exception-track"]; v != nil {
		_ = d.Set("exception_track", v.(map[string]interface{})["name"].(string))
	}

	if resourceTcp["ufp-settings"] != nil {

		ufpSettingsMap := resourceTcp["ufp-settings"].(map[string]interface{})

		ufpSettingsMapToReturn := make(map[string]interface{})

		if v, _ := ufpSettingsMap["server"]; v != nil {
			ufpSettingsMapToReturn["server"] = v.(map[string]interface{})["uid"].(string)
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

	return nil

}
