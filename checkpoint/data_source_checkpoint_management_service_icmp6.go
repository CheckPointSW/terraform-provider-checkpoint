package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementServiceIcmp6() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementServiceIcmp6Read,
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
			"icmp_code": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "As listed in: <a href=\"http://www.iana.org/assignments/icmp-parameters\" target=\"_blank\">RFC 792</a>.",
			},
			"icmp_type": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "As listed in: <a href=\"http://www.iana.org/assignments/icmp-parameters\" target=\"_blank\">RFC 792</a>.",
			},
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
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

func dataSourceManagementServiceIcmp6Read(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showServiceIcmp6Res, err := client.ApiCall("show-service-icmp6", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceIcmp6Res.Success {
		return fmt.Errorf(showServiceIcmp6Res.ErrorMsg)
	}

	serviceIcmp6 := showServiceIcmp6Res.GetData()

	log.Println("Read ServiceIcmp6 - Show JSON = ", serviceIcmp6)

	if v := serviceIcmp6["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serviceIcmp6["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceIcmp6["icmp-code"]; v != nil {
		_ = d.Set("icmp_code", v)
	}

	if v := serviceIcmp6["icmp-type"]; v != nil {
		_ = d.Set("icmp_type", v)
	}

	if v := serviceIcmp6["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if serviceIcmp6["tags"] != nil {
		tagsJson, ok := serviceIcmp6["tags"].([]interface{})
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

	if v := serviceIcmp6["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceIcmp6["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if serviceIcmp6["groups"] != nil {
		groupsJson, ok := serviceIcmp6["groups"].([]interface{})
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
