package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementMobileAccessProfileSection() *schema.Resource {
	return &schema.Resource{
		Create: createManagementMobileAccessProfileSection,
		Read:   readManagementMobileAccessProfileSection,
		Update: updateManagementMobileAccessProfileSection,
		Delete: deleteManagementMobileAccessProfileSection,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createManagementMobileAccessProfileSection(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	mobileAccessProfileSection := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		mobileAccessProfileSection["name"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		mobileAccessProfileSection["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mobileAccessProfileSection["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mobileAccessProfileSection["ignore-errors"] = v.(bool)
	}

	if _, ok := d.GetOk("position"); ok {

		if v, ok := d.GetOk("position.top"); ok {
			if v.(string) == "top" {
				mobileAccessProfileSection["position"] = "top" // entire rule-base
			} else {
				mobileAccessProfileSection["position"] = map[string]interface{}{"top": v.(string)} // section-name
			}
		}

		if v, ok := d.GetOk("position.above"); ok {
			mobileAccessProfileSection["position"] = map[string]interface{}{"above": v.(string)}
		}

		if v, ok := d.GetOk("position.below"); ok {
			mobileAccessProfileSection["position"] = map[string]interface{}{"below": v.(string)}
		}

		if v, ok := d.GetOk("position.bottom"); ok {
			if v.(string) == "bottom" {
				mobileAccessProfileSection["position"] = "bottom" // entire rule-base
			} else {
				mobileAccessProfileSection["position"] = map[string]interface{}{"bottom": v.(string)} // section-name
			}
		}
	}
	log.Println("Create MobileAccessProfileSection - Map = ", mobileAccessProfileSection)

	addMobileAccessProfileSectionRes, err := client.ApiCall("add-mobile-access-profile-section", mobileAccessProfileSection, client.GetSessionID(), true, false)
	if err != nil || !addMobileAccessProfileSectionRes.Success {
		if addMobileAccessProfileSectionRes.ErrorMsg != "" {
			return fmt.Errorf(addMobileAccessProfileSectionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addMobileAccessProfileSectionRes.GetData()["uid"].(string))

	return readManagementMobileAccessProfileSection(d, m)
}

func readManagementMobileAccessProfileSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showMobileAccessProfileSectionRes, err := client.ApiCall("show-mobile-access-profile-section", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMobileAccessProfileSectionRes.Success {
		if objectNotFound(showMobileAccessProfileSectionRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showMobileAccessProfileSectionRes.ErrorMsg)
	}

	mobileAccessProfileSection := showMobileAccessProfileSectionRes.GetData()

	log.Println("Read MobileAccessProfileSection - Show JSON = ", mobileAccessProfileSection)

	if v := mobileAccessProfileSection["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if mobileAccessProfileSection["tags"] != nil {
		tagsJson, ok := mobileAccessProfileSection["tags"].([]interface{})
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

	return nil

}

func updateManagementMobileAccessProfileSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	mobileAccessProfileSection := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		mobileAccessProfileSection["name"] = oldName
		mobileAccessProfileSection["new-name"] = newName
	} else {
		mobileAccessProfileSection["name"] = d.Get("name")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			mobileAccessProfileSection["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			mobileAccessProfileSection["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mobileAccessProfileSection["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mobileAccessProfileSection["ignore-errors"] = v.(bool)
	}

	if d.HasChange("position") {
		if _, ok := d.GetOk("position"); ok {

			if v, ok := d.GetOk("position.top"); ok {
				if v.(string) == "top" {
					mobileAccessProfileSection["new-position"] = "top" // entire rule-base
				} else {
					mobileAccessProfileSection["new-position"] = map[string]interface{}{"top": v.(string)} // specific section-name
				}
			}

			if v, ok := d.GetOk("position.above"); ok {
				mobileAccessProfileSection["new-position"] = map[string]interface{}{"above": v.(string)}
			}

			if v, ok := d.GetOk("position.below"); ok {
				mobileAccessProfileSection["new-position"] = map[string]interface{}{"below": v.(string)}
			}

			if v, ok := d.GetOk("position.bottom"); ok {
				if v.(string) == "bottom" {
					mobileAccessProfileSection["new-position"] = "bottom" // entire rule-base
				} else {
					mobileAccessProfileSection["new-position"] = map[string]interface{}{"bottom": v.(string)} // specific section-name
				}
			}
		}
	}

	log.Println("Update MobileAccessProfileSection - Map = ", mobileAccessProfileSection)

	updateMobileAccessProfileSectionRes, err := client.ApiCall("set-mobile-access-profile-section", mobileAccessProfileSection, client.GetSessionID(), true, false)
	if err != nil || !updateMobileAccessProfileSectionRes.Success {
		if updateMobileAccessProfileSectionRes.ErrorMsg != "" {
			return fmt.Errorf(updateMobileAccessProfileSectionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementMobileAccessProfileSection(d, m)
}

func deleteManagementMobileAccessProfileSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	mobileAccessProfileSectionPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete MobileAccessProfileSection")

	deleteMobileAccessProfileSectionRes, err := client.ApiCall("delete-mobile-access-profile-section", mobileAccessProfileSectionPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteMobileAccessProfileSectionRes.Success {
		if deleteMobileAccessProfileSectionRes.ErrorMsg != "" {
			return fmt.Errorf(deleteMobileAccessProfileSectionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
