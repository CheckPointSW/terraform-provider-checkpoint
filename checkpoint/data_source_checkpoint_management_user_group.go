package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementUserGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementUserGroupRead,
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
			"email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Email Address.",
			},
			"members": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of User Group objects identified by the name or UID.",
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

func dataSourceManagementUserGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showUserGroupRes, err := client.ApiCall("show-user-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showUserGroupRes.Success {
		return fmt.Errorf(showUserGroupRes.ErrorMsg)
	}

	userGroup := showUserGroupRes.GetData()

	log.Println("Read UserGroup - Show JSON = ", userGroup)

	if v := userGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := userGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := userGroup["email"]; v != nil {
		_ = d.Set("email", v)
	}

	if userGroup["members"] != nil {
		membersJson, ok := userGroup["members"].([]interface{})
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

	if userGroup["tags"] != nil {
		tagsJson, ok := userGroup["tags"].([]interface{})
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

	if v := userGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := userGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
