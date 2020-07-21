package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementServiceGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServiceGroupRead,
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
				Description: "Collection of Network objects identified by the name or UID.",
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
				Description: "Collection of group name.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementServiceGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceGroupRes, err := client.ApiCall("show-service-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceGroupRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showServiceGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showServiceGroupRes.ErrorMsg)
	}

	serviceGroup := showServiceGroupRes.GetData()

	if v := serviceGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serviceGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := serviceGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if serviceGroup["members"] != nil {
		membersJson := serviceGroup["members"].([]interface{})
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

	if serviceGroup["groups"] != nil {
		groupsJson := serviceGroup["groups"].([]interface{})
		groupsIds := make([]string, 0)
		if len(groupsJson) > 0 {
			// Create slice of group names
			for _, group_ := range groupsJson {
				group_ := group_.(map[string]interface{})
				groupsIds = append(groupsIds, group_["name"].(string))
			}
		}
		_ = d.Set("groups", groupsIds)
	} else {
		_ = d.Set("groups", nil)
	}

	if serviceGroup["tags"] != nil {
		tagsJson := serviceGroup["tags"].([]interface{})
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
