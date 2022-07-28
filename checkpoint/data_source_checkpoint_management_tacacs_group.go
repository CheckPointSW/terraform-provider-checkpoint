package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementTacacsGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementTacacsGroupRead,
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
			"members": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tacacs servers identified by the name or UID.",
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

func dataSourceManagementTacacsGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showTacacsGroupRes, err := client.ApiCall("show-tacacs-group", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if !showTacacsGroupRes.Success {
		fmt.Errorf(err.Error())
	}

	tacacsGroup := showTacacsGroupRes.GetData()

	log.Println("Read TacacsGroup - Show JSON = ", tacacsGroup)

	if v := tacacsGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := tacacsGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if tacacsGroup["members"] != nil {
		membersJson, ok := tacacsGroup["members"].([]interface{})
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

	if tacacsGroup["tags"] != nil {
		tagsJson, ok := tacacsGroup["tags"].([]interface{})
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

	if v := tacacsGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := tacacsGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
