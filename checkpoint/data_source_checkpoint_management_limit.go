package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementLimit() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementLimitRead,

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
			"enable_download": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable throughput limit for downloads from the internet to the organization.",
			},
			"download_rate": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The Rate for the maximum permitted bandwidth.",
			},
			"download_unit": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Unit for the maximum permitted bandwidth.",
			},
			"enable_upload": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable throughput limit for uploads from the organization to the internet.",
			},
			"upload_rate": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The Rate for the maximum permitted bandwidth.",
			},
			"upload_unit": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Unit for the maximum permitted bandwidth.",
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

func dataSourceManagementLimitRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showLimitRes, err := client.ApiCall("show-limit", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLimitRes.Success {
		if objectNotFound(showLimitRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLimitRes.ErrorMsg)
	}

	limit := showLimitRes.GetData()

	log.Println("Read Limit - Show JSON = ", limit)

	if v := limit["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := limit["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := limit["enable-download"]; v != nil {
		_ = d.Set("enable_download", v)
	}

	if v := limit["download-rate"]; v != nil {
		_ = d.Set("download_rate", v)
	}

	if v := limit["download-unit"]; v != nil {
		_ = d.Set("download_unit", v)
	}

	if v := limit["enable-upload"]; v != nil {
		_ = d.Set("enable_upload", v)
	}

	if v := limit["upload-rate"]; v != nil {
		_ = d.Set("upload_rate", v)
	}

	if v := limit["upload-unit"]; v != nil {
		_ = d.Set("upload_unit", v)
	}

	if limit["tags"] != nil {
		tagsJson, ok := limit["tags"].([]interface{})
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

	if v := limit["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := limit["comments"]; v != nil {
		_ = d.Set("comments", v)
	}
	return nil

}
