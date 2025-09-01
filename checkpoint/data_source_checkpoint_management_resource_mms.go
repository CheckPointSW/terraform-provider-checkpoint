package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementResourceMms() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementResourceMmsRead,
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
			"track": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Logs the activity when a packet matches on a Firewall Rule with the Resource.",
			},
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Accepts or Drops traffic that matches a Firewall Rule using the Resource.",
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

func dataSourceManagementResourceMmsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showResourceMmsRes, err := client.ApiCallSimple("show-resource-mms", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceMmsRes.Success {
		return fmt.Errorf(showResourceMmsRes.ErrorMsg)
	}

	resourceMms := showResourceMmsRes.GetData()

	log.Println("Read ResourceMms - Show JSON = ", resourceMms)

	if v := resourceMms["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := resourceMms["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceMms["track"]; v != nil {
		_ = d.Set("track", v)
	}

	if v := resourceMms["action"]; v != nil {
		_ = d.Set("action", v)
	}

	if resourceMms["tags"] != nil {
		tagsJson, ok := resourceMms["tags"].([]interface{})
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

	if v := resourceMms["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceMms["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
