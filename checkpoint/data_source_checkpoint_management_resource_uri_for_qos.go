package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementResourceUriForQos() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementResourceUriForQosRead,
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
			"search_for_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL string that will be matched to an HTTP connection.",
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

func dataSourceManagementResourceUriForQosRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showResourceUriForQosRes, err := client.ApiCallSimple("show-resource-uri-for-qos", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceUriForQosRes.Success {
		return fmt.Errorf(showResourceUriForQosRes.ErrorMsg)
	}

	resourceUriForQos := showResourceUriForQosRes.GetData()

	log.Println("Read ResourceUriForQos - Show JSON = ", resourceUriForQos)

	if v := resourceUriForQos["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := resourceUriForQos["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceUriForQos["search-for-url"]; v != nil {
		_ = d.Set("search_for_url", v)
	}

	if resourceUriForQos["tags"] != nil {
		tagsJson, ok := resourceUriForQos["tags"].([]interface{})
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

	if v := resourceUriForQos["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceUriForQos["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
