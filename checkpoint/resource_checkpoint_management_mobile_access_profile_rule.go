package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementMobileAccessProfileRule() *schema.Resource {
	return &schema.Resource{
		Create: createManagementMobileAccessProfileRule,
		Read:   readManagementMobileAccessProfileRule,
		Update: updateManagementMobileAccessProfileRule,
		Delete: deleteManagementMobileAccessProfileRule,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"mobile_profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Profile configuration for User groups - identified by the name or UID.",
				Default:     "Default_Profile",
			},
			"user_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "User groups that will be configured with the profile object - identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable/Disable the rule.",
				Default:     true,
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

func createManagementMobileAccessProfileRule(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	mobileAccessProfileRule := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		mobileAccessProfileRule["name"] = v.(string)
	}

	if v, ok := d.GetOk("mobile_profile"); ok {
		mobileAccessProfileRule["mobile-profile"] = v.(string)
	}

	if v, ok := d.GetOk("user_groups"); ok {
		mobileAccessProfileRule["user-groups"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		mobileAccessProfileRule["enabled"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		mobileAccessProfileRule["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("comments"); ok {
		mobileAccessProfileRule["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mobileAccessProfileRule["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mobileAccessProfileRule["ignore-errors"] = v.(bool)
	}

	if _, ok := d.GetOk("position"); ok {

		if v, ok := d.GetOk("position.top"); ok {
			if v.(string) == "top" {
				mobileAccessProfileRule["position"] = "top" // entire rule-base
			} else {
				mobileAccessProfileRule["position"] = map[string]interface{}{"top": v.(string)} // section-name
			}
		}

		if v, ok := d.GetOk("position.above"); ok {
			mobileAccessProfileRule["position"] = map[string]interface{}{"above": v.(string)}
		}

		if v, ok := d.GetOk("position.below"); ok {
			mobileAccessProfileRule["position"] = map[string]interface{}{"below": v.(string)}
		}

		if v, ok := d.GetOk("position.bottom"); ok {
			if v.(string) == "bottom" {
				mobileAccessProfileRule["position"] = "bottom" // entire rule-base
			} else {
				mobileAccessProfileRule["position"] = map[string]interface{}{"bottom": v.(string)} // section-name
			}
		}
	}

	log.Println("Create MobileAccessProfileRule - Map = ", mobileAccessProfileRule)

	addMobileAccessProfileRuleRes, err := client.ApiCall("add-mobile-access-profile-rule", mobileAccessProfileRule, client.GetSessionID(), true, false)
	if err != nil || !addMobileAccessProfileRuleRes.Success {
		if addMobileAccessProfileRuleRes.ErrorMsg != "" {
			return fmt.Errorf(addMobileAccessProfileRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addMobileAccessProfileRuleRes.GetData()["uid"].(string))

	return readManagementMobileAccessProfileRule(d, m)
}

func readManagementMobileAccessProfileRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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

	if v := mobileAccessProfileRule["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := mobileAccessProfileRule["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementMobileAccessProfileRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	mobileAccessProfileRule := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		mobileAccessProfileRule["name"] = oldName
		mobileAccessProfileRule["new-name"] = newName
	} else {
		mobileAccessProfileRule["name"] = d.Get("name")
	}

	if ok := d.HasChange("mobile_profile"); ok {
		mobileAccessProfileRule["mobile-profile"] = d.Get("mobile_profile")
	}

	if d.HasChange("user_groups") {
		if v, ok := d.GetOk("user_groups"); ok {
			mobileAccessProfileRule["user-groups"] = v.(*schema.Set).List()
		} else {
			oldUser_Groups, _ := d.GetChange("user_groups")
			mobileAccessProfileRule["user-groups"] = map[string]interface{}{"remove": oldUser_Groups.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		mobileAccessProfileRule["enabled"] = v.(bool)
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			mobileAccessProfileRule["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			mobileAccessProfileRule["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("comments"); ok {
		mobileAccessProfileRule["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mobileAccessProfileRule["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mobileAccessProfileRule["ignore-errors"] = v.(bool)
	}

	if d.HasChange("position") {
		if _, ok := d.GetOk("position"); ok {

			if v, ok := d.GetOk("position.top"); ok {
				if v.(string) == "top" {
					mobileAccessProfileRule["new-position"] = "top" // entire rule-base
				} else {
					mobileAccessProfileRule["new-position"] = map[string]interface{}{"top": v.(string)} // specific section-name
				}
			}

			if v, ok := d.GetOk("position.above"); ok {
				mobileAccessProfileRule["new-position"] = map[string]interface{}{"above": v.(string)}
			}

			if v, ok := d.GetOk("position.below"); ok {
				mobileAccessProfileRule["new-position"] = map[string]interface{}{"below": v.(string)}
			}

			if v, ok := d.GetOk("position.bottom"); ok {
				if v.(string) == "bottom" {
					mobileAccessProfileRule["new-position"] = "bottom" // entire rule-base
				} else {
					mobileAccessProfileRule["new-position"] = map[string]interface{}{"bottom": v.(string)} // specific section-name
				}
			}
		}
	}

	log.Println("Update MobileAccessProfileRule - Map = ", mobileAccessProfileRule)

	updateMobileAccessProfileRuleRes, err := client.ApiCall("set-mobile-access-profile-rule", mobileAccessProfileRule, client.GetSessionID(), true, false)
	if err != nil || !updateMobileAccessProfileRuleRes.Success {
		if updateMobileAccessProfileRuleRes.ErrorMsg != "" {
			return fmt.Errorf(updateMobileAccessProfileRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementMobileAccessProfileRule(d, m)
}

func deleteManagementMobileAccessProfileRule(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	mobileAccessProfileRulePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete MobileAccessProfileRule")

	deleteMobileAccessProfileRuleRes, err := client.ApiCall("delete-mobile-access-profile-rule", mobileAccessProfileRulePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteMobileAccessProfileRuleRes.Success {
		if deleteMobileAccessProfileRuleRes.ErrorMsg != "" {
			return fmt.Errorf(deleteMobileAccessProfileRuleRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
