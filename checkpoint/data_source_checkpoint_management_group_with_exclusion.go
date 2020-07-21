package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementGroupWithExclusion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementGroupWithExclusionRead,
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
			"except": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name or UID of an object which the group excludes.",
			},
			"include": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name or UID of an object which the group includes.",
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

func dataSourceManagementGroupWithExclusionRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showGroupWithExclusionRes, err := client.ApiCall("show-group-with-exclusion", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGroupWithExclusionRes.Success {
		return fmt.Errorf(showGroupWithExclusionRes.ErrorMsg)
	}

	groupWithExclusion := showGroupWithExclusionRes.GetData()

	log.Println("Read GroupWithExclusion - Show JSON = ", groupWithExclusion)

	if v := groupWithExclusion["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := groupWithExclusion["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := groupWithExclusion["except"]; v != nil {
		_ = d.Set("except", v)
	}

	if v := groupWithExclusion["include"]; v != nil {
		_ = d.Set("include", v)
	}

	if groupWithExclusion["tags"] != nil {
		tagsJson, ok := groupWithExclusion["tags"].([]interface{})
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

	if v := groupWithExclusion["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := groupWithExclusion["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if groupWithExclusion["groups"] != nil {
		groupsJson, ok := groupWithExclusion["groups"].([]interface{})
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
