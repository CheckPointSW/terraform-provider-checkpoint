package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementDataTypeKeywords() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDataTypeKeywordsRead,
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
			"keywords": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Specify keywords or phrases to search for. cannot be empty",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"data_match_threshold": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If set to all-keywords - the data will be matched to the rule only if all the words in the list appear in the data contents. When set to min-keywords any number of the words may appear according to configuration.",
			},
			"min_number_of_keywords": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Define how many of the words in the list must appear in the contents of the data to match the rule. min value is 1.",
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
func dataSourceManagementDataTypeKeywordsRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showDataTypeKeywordsRes, err := client.ApiCall("show-data-type-keywords", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataTypeKeywordsRes.Success {
		if objectNotFound(showDataTypeKeywordsRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDataTypeKeywordsRes.ErrorMsg)
	}

	dataTypeKeywords := showDataTypeKeywordsRes.GetData()

	log.Println("Read DataTypeKeywords - Show JSON = ", dataTypeKeywords)

	if v := dataTypeKeywords["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := dataTypeKeywords["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dataTypeKeywords["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if dataTypeKeywords["keywords"] != nil {
		_ = d.Set("keywords", dataTypeKeywords["keywords"])
	} else {
		_ = d.Set("keywords", nil)
	}

	if v := dataTypeKeywords["data-match-threshold"]; v != nil {
		_ = d.Set("data_match_threshold", v)
	}

	if v := dataTypeKeywords["min-number-of-keywords"]; v != nil {
		_ = d.Set("min_number_of_keywords", v)
	}

	if dataTypeKeywords["tags"] != nil {
		tagsJson, ok := dataTypeKeywords["tags"].([]interface{})
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

	if v := dataTypeKeywords["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dataTypeKeywords["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
