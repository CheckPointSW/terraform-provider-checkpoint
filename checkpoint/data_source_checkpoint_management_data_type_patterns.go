package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementDataTypePatterns() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementDataTypePatternsRead,
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
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "For built-in data types, the description explains the purpose of this type of data representation. For custom-made data types, you can use this field to provide more details.",
			},
			"patterns": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Regular expressions to be evaluated.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"number_of_occurrences": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Define how many times the patterns must appear to be considered data to be protected.",
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
		},
	}
}

func dataSourceManagementDataTypePatternsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showDataTypePatternsRes, err := client.ApiCall("show-data-type-patterns", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataTypePatternsRes.Success {
		if objectNotFound(showDataTypePatternsRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDataTypePatternsRes.ErrorMsg)
	}

	dataTypePatterns := showDataTypePatternsRes.GetData()

	log.Println("Read DataTypePatterns - Show JSON = ", dataTypePatterns)

	if v := dataTypePatterns["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := dataTypePatterns["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dataTypePatterns["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := dataTypePatterns["patterns"]; v != nil {
		_ = d.Set("patterns", v.([]interface{}))
	} else {
		_ = d.Set("patterns", nil)
	}

	if v := dataTypePatterns["number-of-occurrences"]; v != nil {
		_ = d.Set("number_of_occurrences", v)
	}

	if dataTypePatterns["tags"] != nil {
		tagsJson, ok := dataTypePatterns["tags"].([]interface{})
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

	if v := dataTypePatterns["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dataTypePatterns["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
