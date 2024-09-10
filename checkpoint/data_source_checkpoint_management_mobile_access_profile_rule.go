package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementMobileAccessProfileRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementMobileAccessProfileRuleRead,
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
			"mobile_profile": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Profile configuration for User groups - identified by the name or UID.",
			},
			"user_groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "User groups that will be configured with the profile object - identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable/Disable the rule.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}
func dataSourceManagementMobileAccessProfileRuleRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showMobileAccessProfileRuleRes, err := client.ApiCall("show-mobile-access-profile-rule", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMobileAccessProfileRuleRes.Success {
		if objectNotFound(showMobileAccessProfileRuleRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showMobileAccessProfileRuleRes.ErrorMsg)
	}

	mobileAccessProfileRule := showMobileAccessProfileRuleRes.GetData()

	log.Println("Read MobileAccessProfileRule - Show JSON = ", mobileAccessProfileRule)

	if v := mobileAccessProfileRule["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := mobileAccessProfileRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := mobileAccessProfileRule["mobile-profile"]; v != nil {
		profileObj := v.(map[string]interface{})
		if v := profileObj["name"]; v != nil {
			_ = d.Set("mobile_profile", v)
		}
	}

	if mobileAccessProfileRule["user-groups"] != nil {
		userGroupsJson, ok := mobileAccessProfileRule["user-groups"].([]interface{})
		if ok {
			userGroupsNames := make([]string, 0)
			if len(userGroupsJson) > 0 {
				for _, user_groups := range userGroupsJson {
					Obj := user_groups.(map[string]interface{})
					name := Obj["name"].(string)
					userGroupsNames = append(userGroupsNames, name)
				}
			}
			_ = d.Set("user_groups", userGroupsNames)
		}
	} else {
		_ = d.Set("user_groups", nil)
	}

	if v := mobileAccessProfileRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if mobileAccessProfileRule["tags"] != nil {
		tagsJson, ok := mobileAccessProfileRule["tags"].([]interface{})
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

	if v := mobileAccessProfileRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
