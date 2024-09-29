package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementMobileAccessRule() *schema.Resource {
	return &schema.Resource{

		Read: dataManagementMobileAccessRuleRead,

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
			"user_groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "User groups that will be associated with the apps - identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"applications": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Available apps that will be associated with the user groups - identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable/Disable the rule.",
			},
			"install_on": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
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
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}
func dataManagementMobileAccessRuleRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showMobileAccessRuleRes, err := client.ApiCall("show-mobile-access-rule", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMobileAccessRuleRes.Success {
		if objectNotFound(showMobileAccessRuleRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showMobileAccessRuleRes.ErrorMsg)
	}

	mobileAccessRule := showMobileAccessRuleRes.GetData()

	log.Println("Read MobileAccessRule - Show JSON = ", mobileAccessRule)

	if v := mobileAccessRule["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := mobileAccessRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if mobileAccessRule["user-groups"] != nil {
		userGroupsJson, ok := mobileAccessRule["user-groups"].([]interface{})
		if ok {
			userGroupsIds := make([]string, 0)
			if len(userGroupsJson) > 0 {
				for _, user_groups := range userGroupsJson {
					user_groups := user_groups.(map[string]interface{})
					userGroupsIds = append(userGroupsIds, user_groups["name"].(string))
				}
			}
			_ = d.Set("user_groups", userGroupsIds)
		}
	} else {
		_ = d.Set("user_groups", nil)
	}

	if mobileAccessRule["applications"] != nil {
		applicationsJson, ok := mobileAccessRule["applications"].([]interface{})
		if ok {
			applicationsIds := make([]string, 0)
			if len(applicationsJson) > 0 {
				for _, applications := range applicationsJson {
					applications := applications.(map[string]interface{})
					applicationsIds = append(applicationsIds, applications["name"].(string))
				}
			}
			_ = d.Set("applications", applicationsIds)
		}
	} else {
		_ = d.Set("applications", nil)
	}

	if v := mobileAccessRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if mobileAccessRule["install-on"] != nil {
		installOnJson, ok := mobileAccessRule["install-on"].([]interface{})
		if ok {
			installOnIds := make([]string, 0)
			if len(installOnJson) > 0 {
				for _, install_on := range installOnJson {
					install_on := install_on.(map[string]interface{})
					installOnIds = append(installOnIds, install_on["name"].(string))
				}
			}
			_ = d.Set("install_on", installOnIds)
		}
	} else {
		_ = d.Set("install_on", nil)
	}

	if mobileAccessRule["tags"] != nil {
		tagsJson, ok := mobileAccessRule["tags"].([]interface{})
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

	if v := mobileAccessRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
