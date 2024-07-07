package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementDataTypeWeightedKeywords() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceReadManagementDataTypeWeightedKeywords,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
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
			"weighted_keywords": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of keywords or phrases.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"keyword": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "keyword or regular expression to be weighted.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Weight of the expression.",
						},
						"max_weight": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Max weight of the expression.",
						},
						"regex": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determine whether to consider the expression as a regular expression.",
						},
					},
				},
			},
			"sum_of_weights_threshold": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Define the number of appearances, by weight, of all the keywords that, beyond this threshold,  the data containing this list of words or phrases will be recognized as data to be protected.",
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

func dataSourceReadManagementDataTypeWeightedKeywords(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showDataTypeWeightedKeywordsRes, err := client.ApiCall("show-data-type-weighted-keywords", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataTypeWeightedKeywordsRes.Success {
		if objectNotFound(showDataTypeWeightedKeywordsRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDataTypeWeightedKeywordsRes.ErrorMsg)
	}

	dataTypeWeightedKeywords := showDataTypeWeightedKeywordsRes.GetData()

	log.Println("Read DataTypeWeightedKeywords - Show JSON = ", dataTypeWeightedKeywords)

	if v := dataTypeWeightedKeywords["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := dataTypeWeightedKeywords["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dataTypeWeightedKeywords["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if dataTypeWeightedKeywords["weighted-keywords"] != nil {

		weightedKeywordsList, ok := dataTypeWeightedKeywords["weighted-keywords"].([]interface{})

		if ok {

			if len(weightedKeywordsList) > 0 {

				var weightedKeywordsListToReturn []map[string]interface{}

				for i := range weightedKeywordsList {

					weightedKeywordsMap := weightedKeywordsList[i].(map[string]interface{})

					weightedKeywordsMapToAdd := make(map[string]interface{})

					if v, _ := weightedKeywordsMap["keyword"]; v != nil {
						weightedKeywordsMapToAdd["keyword"] = v
					}
					if v, _ := weightedKeywordsMap["weight"]; v != nil {
						weightedKeywordsMapToAdd["weight"] = v
					}
					if v, _ := weightedKeywordsMap["max-weight"]; v != nil {
						weightedKeywordsMapToAdd["max_weight"] = v
					}
					if v, _ := weightedKeywordsMap["regex"]; v != nil {
						weightedKeywordsMapToAdd["regex"] = v
					}
					weightedKeywordsListToReturn = append(weightedKeywordsListToReturn, weightedKeywordsMapToAdd)

				}
				_ = d.Set("weighted_keywords", weightedKeywordsListToReturn)
			}
		}
	}

	if v := dataTypeWeightedKeywords["sum-of-weights-threshold"]; v != nil {
		_ = d.Set("sum_of_weights_threshold", v)
	}

	if dataTypeWeightedKeywords["tags"] != nil {
		tagsJson, ok := dataTypeWeightedKeywords["tags"].([]interface{})
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

	if v := dataTypeWeightedKeywords["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dataTypeWeightedKeywords["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
