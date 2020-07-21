package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementWildcard() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementWildcardRead,
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
			"ipv4_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 address.",
			},
			"ipv4_mask_wildcard": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 mask wildcard.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 address.",
			},
			"ipv6_mask_wildcard": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 mask wildcard.",
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

func dataSourceManagementWildcardRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showWildcardRes, err := client.ApiCall("show-wildcard", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showWildcardRes.Success {
		return fmt.Errorf(showWildcardRes.ErrorMsg)
	}

	wildcard := showWildcardRes.GetData()

	log.Println("Read Wildcard - Show JSON = ", wildcard)

	if v := wildcard["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := wildcard["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := wildcard["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := wildcard["ipv4-mask-wildcard"]; v != nil {
		_ = d.Set("ipv4_mask_wildcard", v)
	}

	if v := wildcard["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := wildcard["ipv6-mask-wildcard"]; v != nil {
		_ = d.Set("ipv6_mask_wildcard", v)
	}

	if wildcard["tags"] != nil {
		tagsJson, ok := wildcard["tags"].([]interface{})
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

	if v := wildcard["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := wildcard["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if wildcard["groups"] != nil {
		groupsJson, ok := wildcard["groups"].([]interface{})
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
