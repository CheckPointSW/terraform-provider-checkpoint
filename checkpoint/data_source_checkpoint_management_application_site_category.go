package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementApplicationSiteCategory() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementApplicationSiteCategoryRead,
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
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "N/A",
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
			"groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of group identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementApplicationSiteCategoryRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showApplicationSiteCategoryRes, err := client.ApiCall("show-application-site-category", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showApplicationSiteCategoryRes.Success {
		return fmt.Errorf(showApplicationSiteCategoryRes.ErrorMsg)
	}

	applicationSiteCategory := showApplicationSiteCategoryRes.GetData()

	log.Println("Read ApplicationSiteCategory - Show JSON = ", applicationSiteCategory)

	if v := applicationSiteCategory["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := applicationSiteCategory["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := applicationSiteCategory["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if applicationSiteCategory["tags"] != nil {
		tagsJson, ok := applicationSiteCategory["tags"].([]interface{})
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

	if v := applicationSiteCategory["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := applicationSiteCategory["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if applicationSiteCategory["groups"] != nil {
		groupsJson, ok := applicationSiteCategory["groups"].([]interface{})
		if ok {
			groupsIds := make([]string, 0)
			if len(groupsJson) > 0 {
				for _, groups := range groupsJson {
					groups := groups.(map[string]interface{})
					groupsIds = append(groupsIds, groups["name"].(string))
				}
			}
			_ = d.Set("groups", groupsIds)
		}
	} else {
		_ = d.Set("groups", nil)
	}

	return nil
}
