package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementRadiusGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementRadiusGroupRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"members": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of radius servers identified by the name or UID.",
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
		},
	}
}

func dataSourceManagementRadiusGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showRadiusGroupRes, err := client.ApiCall("show-radius-group", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRadiusGroupRes.Success {
		return fmt.Errorf(showRadiusGroupRes.ErrorMsg)
	}

	radiusGroup := showRadiusGroupRes.GetData()

	if v := radiusGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := radiusGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := radiusGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := radiusGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if radiusGroup["members"] != nil {
		membersJson := radiusGroup["members"].([]interface{})
		membersIds := make([]string, 0)
		if len(membersJson) > 0 {
			// Create slice of members names
			for _, member := range membersJson {
				member := member.(map[string]interface{})
				membersIds = append(membersIds, member["name"].(string))
			}
		}
		_ = d.Set("members", membersIds)
	} else {
		_ = d.Set("members", nil)
	}

	if radiusGroup["tags"] != nil {
		tagsJson := radiusGroup["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	return nil
}
