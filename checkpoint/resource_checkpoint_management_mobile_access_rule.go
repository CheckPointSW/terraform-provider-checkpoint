package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementMobileAccessRule() *schema.Resource {
	return &schema.Resource{
		Create: createManagementMobileAccessRule,
		Read:   readManagementMobileAccessRule,
		Update: updateManagementMobileAccessRule,
		Delete: deleteManagementMobileAccessRule,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"user_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "User groups that will be associated with the apps - identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"applications": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Available apps that will be associated with the user groups - identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable/Disable the rule.",
			},
			"install_on": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
			"position": &schema.Schema{
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Position in the rulebase.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"top": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
						"above": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
						"below": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
						"bottom": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
					},
				},
			},
		},
	}
}

func createManagementMobileAccessRule(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	mobileAccessRule := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		mobileAccessRule["name"] = v.(string)
	}

	if v, ok := d.GetOk("user_groups"); ok {
		mobileAccessRule["user-groups"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("applications"); ok {
		mobileAccessRule["applications"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		mobileAccessRule["enabled"] = v.(bool)
	}

	if v, ok := d.GetOk("install_on"); ok {
		mobileAccessRule["install-on"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		mobileAccessRule["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("comments"); ok {
		mobileAccessRule["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mobileAccessRule["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mobileAccessRule["ignore-errors"] = v.(bool)
	}

	if _, ok := d.GetOk("position"); ok {

		if v, ok := d.GetOk("position.top"); ok {
			if v.(string) == "top" {
				mobileAccessRule["position"] = "top" // entire rule-base
			} else {
				mobileAccessRule["position"] = map[string]interface{}{"top": v.(string)} // section-name
			}
		}

		if v, ok := d.GetOk("position.above"); ok {
			mobileAccessRule["position"] = map[string]interface{}{"above": v.(string)}
		}

		if v, ok := d.GetOk("position.below"); ok {
			mobileAccessRule["position"] = map[string]interface{}{"below": v.(string)}
		}

		if v, ok := d.GetOk("position.bottom"); ok {
			if v.(string) == "bottom" {
				mobileAccessRule["position"] = "bottom" // entire rule-base
			} else {
				mobileAccessRule["position"] = map[string]interface{}{"bottom": v.(string)} // section-name
			}
		}
	}
	log.Println("Create MobileAccessRule - Map = ", mobileAccessRule)

	addMobileAccessRuleRes, err := client.ApiCall("add-mobile-access-rule", mobileAccessRule, client.GetSessionID(), true, false)
	if err != nil || !addMobileAccessRuleRes.Success {
		if addMobileAccessRuleRes.ErrorMsg != "" {
			return fmt.Errorf(addMobileAccessRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addMobileAccessRuleRes.GetData()["uid"].(string))

	return readManagementMobileAccessRule(d, m)
}

func readManagementMobileAccessRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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

func updateManagementMobileAccessRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	mobileAccessRule := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		mobileAccessRule["name"] = oldName
		mobileAccessRule["new-name"] = newName
	} else {
		mobileAccessRule["name"] = d.Get("name")
	}

	if d.HasChange("user_groups") {
		if v, ok := d.GetOk("user_groups"); ok {
			mobileAccessRule["user-groups"] = v.(*schema.Set).List()
		} else {
			oldUser_Groups, _ := d.GetChange("user_groups")
			mobileAccessRule["user-groups"] = map[string]interface{}{"remove": oldUser_Groups.(*schema.Set).List()}
		}
	}

	if d.HasChange("applications") {
		if v, ok := d.GetOk("applications"); ok {
			mobileAccessRule["applications"] = v.(*schema.Set).List()
		} else {
			oldApplications, _ := d.GetChange("applications")
			mobileAccessRule["applications"] = map[string]interface{}{"remove": oldApplications.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		mobileAccessRule["enabled"] = v.(bool)
	}

	if d.HasChange("install_on") {
		if v, ok := d.GetOk("install_on"); ok {
			mobileAccessRule["install-on"] = v.(*schema.Set).List()
		} else {
			oldInstall_On, _ := d.GetChange("install_on")
			mobileAccessRule["install-on"] = map[string]interface{}{"remove": oldInstall_On.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			mobileAccessRule["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			mobileAccessRule["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("comments"); ok {
		mobileAccessRule["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mobileAccessRule["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mobileAccessRule["ignore-errors"] = v.(bool)
	}

	if d.HasChange("position") {
		if _, ok := d.GetOk("position"); ok {

			if v, ok := d.GetOk("position.top"); ok {
				if v.(string) == "top" {
					mobileAccessRule["new-position"] = "top" // entire rule-base
				} else {
					mobileAccessRule["new-position"] = map[string]interface{}{"top": v.(string)} // specific section-name
				}
			}

			if v, ok := d.GetOk("position.above"); ok {
				mobileAccessRule["new-position"] = map[string]interface{}{"above": v.(string)}
			}

			if v, ok := d.GetOk("position.below"); ok {
				mobileAccessRule["new-position"] = map[string]interface{}{"below": v.(string)}
			}

			if v, ok := d.GetOk("position.bottom"); ok {
				if v.(string) == "bottom" {
					mobileAccessRule["new-position"] = "bottom" // entire rule-base
				} else {
					mobileAccessRule["new-position"] = map[string]interface{}{"bottom": v.(string)} // specific section-name
				}
			}
		}
	}
	log.Println("Update MobileAccessRule - Map = ", mobileAccessRule)

	updateMobileAccessRuleRes, err := client.ApiCall("set-mobile-access-rule", mobileAccessRule, client.GetSessionID(), true, false)
	if err != nil || !updateMobileAccessRuleRes.Success {
		if updateMobileAccessRuleRes.ErrorMsg != "" {
			return fmt.Errorf(updateMobileAccessRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementMobileAccessRule(d, m)
}

func deleteManagementMobileAccessRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	mobileAccessRulePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete MobileAccessRule")

	deleteMobileAccessRuleRes, err := client.ApiCall("delete-mobile-access-rule", mobileAccessRulePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteMobileAccessRuleRes.Success {
		if deleteMobileAccessRuleRes.ErrorMsg != "" {
			return fmt.Errorf(deleteMobileAccessRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
