package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementGenericDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGenericDataCenterServerRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the JSON feed (e.g. https://example.com/file.json).",
			},
			"interval": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Update interval of the feed in seconds.",
			},
			"custom_header": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "When set to false, The admin is not using Key and Value for a Custom Header in order to connect to the feed server.\n\nWhen set to true, The admin is using Key and Value for a Custom Header in order to connect to the feed server.",
			},
			"custom_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Key for the Custom Header, relevant and required only when custom_header set to true.",
			},
			"custom_value": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Value for the Custom Header, relevant and required only when custom_header set to true.",
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Apply changes ignoring warnings. By Setting this parameter to 'true' test connection failure will be ignored.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func dataSourceGenericDataCenterServerRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	var name string
	var uid string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("uid"); ok {
		uid = v.(string)
	}
	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showGenericDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGenericDataCenterServerRes.Success {
		if objectNotFound(showGenericDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showGenericDataCenterServerRes.ErrorMsg)
	}
	genericDataCenterServer := showGenericDataCenterServerRes.GetData()

	if v := genericDataCenterServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := genericDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := genericDataCenterServer["url"]; v != nil {
		_ = d.Set("url", v)
	}

	if v := genericDataCenterServer["interval"]; v != nil {
		_ = d.Set("interval", v)
	}

	if v := genericDataCenterServer["custom_header"]; v != nil {
		_ = d.Set("custom_header", v.(bool))
	}
	if v := genericDataCenterServer["custom_key"]; v != nil {
		_ = d.Set("custom_key", v)
	}

	if v := genericDataCenterServer["custom_value"]; v != nil {
		_ = d.Set("custom_value", v)
	}

	if genericDataCenterServer["tags"] != nil {
		tagsJson, ok := genericDataCenterServer["tags"].([]interface{})
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

	if v := genericDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := genericDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := genericDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := genericDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}
