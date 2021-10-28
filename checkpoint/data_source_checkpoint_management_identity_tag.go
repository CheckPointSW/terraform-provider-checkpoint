package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementIdentityTag() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementIdentityTagRead,
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
			"external_identifier": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "External identifier. For example: Cisco ISE security group tag.",
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

func dataSourceManagementIdentityTagRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showIdentityTagRes, err := client.ApiCall("show-identity-tag", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIdentityTagRes.Success {
		return fmt.Errorf(showIdentityTagRes.ErrorMsg)
	}

	identityTag := showIdentityTagRes.GetData()

	log.Println("Read IdentityTag - Show JSON = ", identityTag)

	if v := identityTag["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := identityTag["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := identityTag["external-identifier"]; v != nil {
		_ = d.Set("external_identifier", v)
	}

	if identityTag["tags"] != nil {
		tagsJson, ok := identityTag["tags"].([]interface{})
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

	if v := identityTag["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := identityTag["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
