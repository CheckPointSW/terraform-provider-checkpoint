package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementOverrideCategorization() *schema.Resource {
	return &schema.Resource{
		Create: createManagementOverrideCategorization,
		Read:   readManagementOverrideCategorization,
		Update: updateManagementOverrideCategorization,
		Delete: deleteManagementOverrideCategorization,
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL for which we want to update the category and risk definitions, the URL and the object name are the same for Override Categorization.",
			},
			"url_defined_as_regular_expression": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "States whether the URL is defined as a Regular Expression or not.",
				Default:     false,
			},
			"new_primary_category": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Uid or name of the primary category based on its most defining aspect.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"risk": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "States the override categorization risk.",
			},
			"additional_categories": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Uid or name of the categories to override in the Application and URL Filtering or Threat Prevention.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
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
		},
	}
}

func createManagementOverrideCategorization(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	overrideCategorization := make(map[string]interface{})

	if v, ok := d.GetOk("url"); ok {
		overrideCategorization["url"] = v.(string)
	}

	if v, ok := d.GetOkExists("url_defined_as_regular_expression"); ok {
		overrideCategorization["url-defined-as-regular-expression"] = v.(bool)
	}

	if v, ok := d.GetOk("new_primary_category"); ok {
		overrideCategorization["new-primary-category"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		overrideCategorization["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("risk"); ok {
		overrideCategorization["risk"] = v.(string)
	}

	if v, ok := d.GetOk("additional_categories"); ok {
		overrideCategorization["additional-categories"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		overrideCategorization["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		overrideCategorization["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		overrideCategorization["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		overrideCategorization["ignore-errors"] = v.(bool)
	}

	log.Println("Create OverrideCategorization - Map = ", overrideCategorization)

	addOverrideCategorizationRes, err := client.ApiCall("add-override-categorization", overrideCategorization, client.GetSessionID(), true, false)
	if err != nil || !addOverrideCategorizationRes.Success {
		if addOverrideCategorizationRes.ErrorMsg != "" {
			return fmt.Errorf(addOverrideCategorizationRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addOverrideCategorizationRes.GetData()["uid"].(string))

	return readManagementOverrideCategorization(d, m)
}

func readManagementOverrideCategorization(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("url"); ok {
		payload["url"] = v
	} else {
		payload["uid"] = d.Id()
	}

	showOverrideCategorizationRes, err := client.ApiCall("show-override-categorization", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showOverrideCategorizationRes.Success {
		if objectNotFound(showOverrideCategorizationRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showOverrideCategorizationRes.ErrorMsg)
	}

	overrideCategorization := showOverrideCategorizationRes.GetData()

	log.Println("Read OverrideCategorization - Show JSON = ", overrideCategorization)

	if v := overrideCategorization["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}
	if v := overrideCategorization["url"]; v != nil {
		_ = d.Set("url", v)
	}

	if v := overrideCategorization["url-defined-as-regular-expression"]; v != nil {
		_ = d.Set("url_defined_as_regular_expression", v)
	}

	if v := overrideCategorization["new-primary-category"]; v != nil {
		objMap := v.(map[string]interface{})
		if v := objMap["name"]; v != nil {
			_ = d.Set("new_primary_category", v)
		}
	}

	if overrideCategorization["tags"] != nil {
		tagsJson, ok := overrideCategorization["tags"].([]interface{})
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

	if v := overrideCategorization["risk"]; v != nil {
		_ = d.Set("risk", v)
	}

	if overrideCategorization["additional_categories"] != nil {
		additionalCategoriesJson, ok := overrideCategorization["additional_categories"].([]interface{})
		if ok {
			additionalCategoriesIds := make([]string, 0)
			if len(additionalCategoriesJson) > 0 {
				for _, additional_categories := range additionalCategoriesJson {
					additional_categories := additional_categories.(map[string]interface{})
					additionalCategoriesIds = append(additionalCategoriesIds, additional_categories["name"].(string))
				}
			}
			_ = d.Set("additional_categories", additionalCategoriesIds)
		}
	} else {
		_ = d.Set("additional_categories", nil)
	}

	if v := overrideCategorization["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := overrideCategorization["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}

func updateManagementOverrideCategorization(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	overrideCategorization := make(map[string]interface{})

	if ok := d.HasChange("url"); ok {
		oldName, newName := d.GetChange("url")
		overrideCategorization["url"] = oldName
		overrideCategorization["new-url"] = newName
	} else {
		overrideCategorization["url"] = d.Get("url")
	}

	if v, ok := d.GetOkExists("url_defined_as_regular_expression"); ok {
		overrideCategorization["url-defined-as-regular-expression"] = v.(bool)
	}

	if ok := d.HasChange("new_primary_category"); ok {
		overrideCategorization["new-primary-category"] = d.Get("new_primary_category")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			overrideCategorization["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			overrideCategorization["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("risk"); ok {
		overrideCategorization["risk"] = d.Get("risk")
	}

	if d.HasChange("additional_categories") {
		if v, ok := d.GetOk("additional_categories"); ok {
			overrideCategorization["additional-categories"] = v.(*schema.Set).List()
		} else {
			oldAdditional_Categories, _ := d.GetChange("additional_categories")
			overrideCategorization["additional-categories"] = map[string]interface{}{"remove": oldAdditional_Categories.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		overrideCategorization["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		overrideCategorization["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		overrideCategorization["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		overrideCategorization["ignore-errors"] = v.(bool)
	}

	log.Println("Update OverrideCategorization - Map = ", overrideCategorization)

	updateOverrideCategorizationRes, err := client.ApiCall("set-override-categorization", overrideCategorization, client.GetSessionID(), true, false)
	if err != nil || !updateOverrideCategorizationRes.Success {
		if updateOverrideCategorizationRes.ErrorMsg != "" {
			return fmt.Errorf(updateOverrideCategorizationRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementOverrideCategorization(d, m)
}

func deleteManagementOverrideCategorization(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	overrideCategorizationPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete OverrideCategorization")

	deleteOverrideCategorizationRes, err := client.ApiCall("delete-override-categorization", overrideCategorizationPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteOverrideCategorizationRes.Success {
		if deleteOverrideCategorizationRes.ErrorMsg != "" {
			return fmt.Errorf(deleteOverrideCategorizationRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
