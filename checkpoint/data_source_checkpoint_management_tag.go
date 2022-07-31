package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementTag() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementTagRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Must be unique in the domain",
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string",
			},
		},
	}
}

func dataSourceManagementTagRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showTag, err := client.ApiCall("show-tag", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if !showTag.Success {
		fmt.Errorf(err.Error())
	}

	tag := showTag.GetData()

	log.Println("Read Tag - Show JSON = ", tag)

	if v := tag["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := tag["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if tag["tags"] != nil {
		tagsJson, ok := tag["tags"].([]interface{})
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

	if v := tag["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := tag["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}