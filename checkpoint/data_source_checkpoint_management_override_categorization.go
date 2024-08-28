package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementOverrideCategorization() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementOverrideCategorizationRead,

		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The URL for which we want to update the category and risk definitions, the URL and the object name are the same for Override Categorization.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"url_defined_as_regular_expression": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "States whether the URL is defined as a Regular Expression or not.",
			},
			"new_primary_category": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Uid or name of the primary category based on its most defining aspect.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"risk": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "States the override categorization risk.",
			},
			"additional_categories": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Uid or name of the categories to override in the Application and URL Filtering or Threat Prevention.",
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
func dataSourceManagementOverrideCategorizationRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("url"); ok {
		payload["url"] = v
	} else {
		payload["uid"] = d.Get("uid")
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
