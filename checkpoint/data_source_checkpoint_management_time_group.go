package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementTimeGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementTimeGroupRead,
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
			"members": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Time Group objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func dataSourceManagementTimeGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showTimeGroupRes, err := client.ApiCall("show-time-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTimeGroupRes.Success {
		return fmt.Errorf(showTimeGroupRes.ErrorMsg)
	}

	timeGroup := showTimeGroupRes.GetData()

	log.Println("Read TimeGroup - Show JSON = ", timeGroup)

	if v := timeGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := timeGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if timeGroup["members"] != nil {
		membersJson, ok := timeGroup["members"].([]interface{})
		if ok {
			membersIds := make([]string, 0)
			if len(membersJson) > 0 {
				for _, members := range membersJson {
					members := members.(map[string]interface{})
					membersIds = append(membersIds, members["name"].(string))
				}
			}
			_ = d.Set("members", membersIds)
		}
	} else {
		_ = d.Set("members", nil)
	}

	if timeGroup["tags"] != nil {
		tagsJson, ok := timeGroup["tags"].([]interface{})
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

	if v := timeGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := timeGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if timeGroup["groups"] != nil {
		groupsJson, ok := timeGroup["groups"].([]interface{})
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
