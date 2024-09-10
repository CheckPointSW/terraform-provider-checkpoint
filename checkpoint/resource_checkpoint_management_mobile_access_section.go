package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementMobileAccessSection() *schema.Resource {
	return &schema.Resource{
		Create: createManagementMobileAccessSection,
		Read:   readManagementMobileAccessSection,
		Update: updateManagementMobileAccessSection,
		Delete: deleteManagementMobileAccessSection,
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

func createManagementMobileAccessSection(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	mobileAccessSection := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		mobileAccessSection["name"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		mobileAccessSection["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mobileAccessSection["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mobileAccessSection["ignore-errors"] = v.(bool)
	}

	if _, ok := d.GetOk("position"); ok {

		if v, ok := d.GetOk("position.top"); ok {
			if v.(string) == "top" {
				mobileAccessSection["position"] = "top" // entire rule-base
			} else {
				mobileAccessSection["position"] = map[string]interface{}{"top": v.(string)} // section-name
			}
		}

		if v, ok := d.GetOk("position.above"); ok {
			mobileAccessSection["position"] = map[string]interface{}{"above": v.(string)}
		}

		if v, ok := d.GetOk("position.below"); ok {
			mobileAccessSection["position"] = map[string]interface{}{"below": v.(string)}
		}

		if v, ok := d.GetOk("position.bottom"); ok {
			if v.(string) == "bottom" {
				mobileAccessSection["position"] = "bottom" // entire rule-base
			} else {
				mobileAccessSection["position"] = map[string]interface{}{"bottom": v.(string)} // section-name
			}
		}
	}

	log.Println("Create MobileAccessSection - Map = ", mobileAccessSection)

	addMobileAccessSectionRes, err := client.ApiCall("add-mobile-access-section", mobileAccessSection, client.GetSessionID(), true, false)
	if err != nil || !addMobileAccessSectionRes.Success {
		if addMobileAccessSectionRes.ErrorMsg != "" {
			return fmt.Errorf(addMobileAccessSectionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addMobileAccessSectionRes.GetData()["uid"].(string))

	return readManagementMobileAccessSection(d, m)
}

func readManagementMobileAccessSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showMobileAccessSectionRes, err := client.ApiCall("show-mobile-access-section", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMobileAccessSectionRes.Success {
		if objectNotFound(showMobileAccessSectionRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showMobileAccessSectionRes.ErrorMsg)
	}

	mobileAccessSection := showMobileAccessSectionRes.GetData()

	log.Println("Read MobileAccessSection - Show JSON = ", mobileAccessSection)

	if v := mobileAccessSection["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if mobileAccessSection["tags"] != nil {
		tagsJson, ok := mobileAccessSection["tags"].([]interface{})
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

func updateManagementMobileAccessSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	mobileAccessSection := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		mobileAccessSection["name"] = oldName
		mobileAccessSection["new-name"] = newName
	} else {
		mobileAccessSection["name"] = d.Get("name")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			mobileAccessSection["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			mobileAccessSection["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mobileAccessSection["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mobileAccessSection["ignore-errors"] = v.(bool)
	}

	if d.HasChange("position") {
		if _, ok := d.GetOk("position"); ok {

			if v, ok := d.GetOk("position.top"); ok {
				if v.(string) == "top" {
					mobileAccessSection["new-position"] = "top" // entire rule-base
				} else {
					mobileAccessSection["new-position"] = map[string]interface{}{"top": v.(string)} // specific section-name
				}
			}

			if v, ok := d.GetOk("position.above"); ok {
				mobileAccessSection["new-position"] = map[string]interface{}{"above": v.(string)}
			}

			if v, ok := d.GetOk("position.below"); ok {
				mobileAccessSection["new-position"] = map[string]interface{}{"below": v.(string)}
			}

			if v, ok := d.GetOk("position.bottom"); ok {
				if v.(string) == "bottom" {
					mobileAccessSection["new-position"] = "bottom" // entire rule-base
				} else {
					mobileAccessSection["new-position"] = map[string]interface{}{"bottom": v.(string)} // specific section-name
				}
			}
		}
	}

	log.Println("Update MobileAccessSection - Map = ", mobileAccessSection)

	updateMobileAccessSectionRes, err := client.ApiCall("set-mobile-access-section", mobileAccessSection, client.GetSessionID(), true, false)
	if err != nil || !updateMobileAccessSectionRes.Success {
		if updateMobileAccessSectionRes.ErrorMsg != "" {
			return fmt.Errorf(updateMobileAccessSectionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementMobileAccessSection(d, m)
}

func deleteManagementMobileAccessSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	mobileAccessSectionPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete MobileAccessSection")

	deleteMobileAccessSectionRes, err := client.ApiCall("delete-mobile-access-section", mobileAccessSectionPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteMobileAccessSectionRes.Success {
		if deleteMobileAccessSectionRes.ErrorMsg != "" {
			return fmt.Errorf(deleteMobileAccessSectionRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
