package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"strconv"
)

func resourceManagementDataTypeWeightedKeywords() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDataTypeWeightedKeywords,
		Read:   readManagementDataTypeWeightedKeywords,
		Update: updateManagementDataTypeWeightedKeywords,
		Delete: deleteManagementDataTypeWeightedKeywords,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "For built-in data types, the description explains the purpose of this type of data representation. For custom-made data types, you can use this field to provide more details.",
			},
			"weighted_keywords": {
				Type:        schema.TypeList,
				Required:    true,
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
							Optional:    true,
							Description: "Weight of the expression.",
							Default:     1,
						},
						"max_weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Max weight of the expression.",
							Default:     0,
						},
						"regex": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Determine whether to consider the expression as a regular expression.",
							Default:     false,
						},
					},
				},
			},
			"sum_of_weights_threshold": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Define the number of appearances, by weight, of all the keywords that, beyond this threshold,  the data containing this list of words or phrases will be recognized as data to be protected.",
				Default:     1,
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
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

func createManagementDataTypeWeightedKeywords(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dataTypeWeightedKeywords := make(map[string]interface{})
	log.Println("in create")
	if v, ok := d.GetOk("name"); ok {
		dataTypeWeightedKeywords["name"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		dataTypeWeightedKeywords["description"] = v.(string)
	}

	if v, ok := d.GetOk("weighted_keywords"); ok {

		weightedKeywordsList := v.([]interface{})
		if len(weightedKeywordsList) > 0 {

			var weightedKeywordsPayload []map[string]interface{}

			for i := range weightedKeywordsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("weighted_keywords." + strconv.Itoa(i) + ".keyword"); ok {
					Payload["keyword"] = v.(string)
				}
				if v, ok := d.GetOk("weighted_keywords." + strconv.Itoa(i) + ".weight"); ok {
					Payload["weight"] = v.(int)
				}
				if v, ok := d.GetOk("weighted_keywords." + strconv.Itoa(i) + ".max_weight"); ok {
					Payload["max-weight"] = v.(int)
				}
				if v, ok := d.GetOk("weighted_keywords." + strconv.Itoa(i) + ".regex"); ok {
					Payload["regex"] = v.(bool)
				}
				weightedKeywordsPayload = append(weightedKeywordsPayload, Payload)
			}
			dataTypeWeightedKeywords["weighted-keywords"] = weightedKeywordsPayload
		}
	}

	if v, ok := d.GetOk("sum_of_weights_threshold"); ok {
		dataTypeWeightedKeywords["sum-of-weights-threshold"] = v.(int)
	}

	if v, ok := d.GetOk("tags"); ok {
		dataTypeWeightedKeywords["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dataTypeWeightedKeywords["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dataTypeWeightedKeywords["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeWeightedKeywords["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeWeightedKeywords["ignore-errors"] = v.(bool)
	}

	log.Println("Create DataTypeWeightedKeywords - Map = ", dataTypeWeightedKeywords)

	addDataTypeWeightedKeywordsRes, err := client.ApiCall("add-data-type-weighted-keywords", dataTypeWeightedKeywords, client.GetSessionID(), true, false)
	if err != nil || !addDataTypeWeightedKeywordsRes.Success {
		if addDataTypeWeightedKeywordsRes.ErrorMsg != "" {
			return fmt.Errorf(addDataTypeWeightedKeywordsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDataTypeWeightedKeywordsRes.GetData()["uid"].(string))

	return readManagementDataTypeWeightedKeywords(d, m)
}

func readManagementDataTypeWeightedKeywords(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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

	if v := dataTypeWeightedKeywords["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dataTypeWeightedKeywords["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDataTypeWeightedKeywords(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dataTypeWeightedKeywords := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		dataTypeWeightedKeywords["name"] = oldName
		dataTypeWeightedKeywords["new-name"] = newName
	} else {
		dataTypeWeightedKeywords["name"] = d.Get("name")
	}

	if ok := d.HasChange("description"); ok {
		dataTypeWeightedKeywords["description"] = d.Get("description")
	}

	if d.HasChange("weighted_keywords") {

		if v, ok := d.GetOk("weighted_keywords"); ok {

			weightedKeywordsList := v.([]interface{})

			var weightedKeywordsPayload []map[string]interface{}

			for i := range weightedKeywordsList {

				Payload := make(map[string]interface{})

				weightedKeytword := weightedKeywordsList[i].(map[string]interface{})

				if v := weightedKeytword["keyword"]; v != nil {
					Payload["keyword"] = v
				}
				if v := weightedKeytword["weight"]; v != nil {
					Payload["weight"] = v
				}
				if v := weightedKeytword["max_weight"]; v != nil {
					Payload["max-weight"] = v
				}
				if v := weightedKeytword["regex"]; v != nil {
					Payload["regex"] = v
				}
				log.Println("payload is ", Payload)
				weightedKeywordsPayload = append(weightedKeywordsPayload, Payload)
			}
			dataTypeWeightedKeywords["weighted-keywords"] = weightedKeywordsPayload
		} else {
			oldweightedKeywords, _ := d.GetChange("weighted_keywords")
			var weightedKeywordsToDelete []interface{}
			for _, i := range oldweightedKeywords.([]interface{}) {
				weightedKeywordsToDelete = append(weightedKeywordsToDelete, i.(map[string]interface{})["name"].(string))
			}
			dataTypeWeightedKeywords["weighted-keywords"] = map[string]interface{}{"remove": weightedKeywordsToDelete}
		}
	}

	if ok := d.HasChange("sum_of_weights_threshold"); ok {
		dataTypeWeightedKeywords["sum-of-weights-threshold"] = d.Get("sum_of_weights_threshold")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dataTypeWeightedKeywords["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dataTypeWeightedKeywords["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dataTypeWeightedKeywords["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dataTypeWeightedKeywords["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeWeightedKeywords["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeWeightedKeywords["ignore-errors"] = v.(bool)
	}

	log.Println("Update DataTypeWeightedKeywords - Map = ", dataTypeWeightedKeywords)

	updateDataTypeWeightedKeywordsRes, err := client.ApiCall("set-data-type-weighted-keywords", dataTypeWeightedKeywords, client.GetSessionID(), true, false)
	if err != nil || !updateDataTypeWeightedKeywordsRes.Success {
		if updateDataTypeWeightedKeywordsRes.ErrorMsg != "" {
			return fmt.Errorf(updateDataTypeWeightedKeywordsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDataTypeWeightedKeywords(d, m)
}

func deleteManagementDataTypeWeightedKeywords(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataTypeWeightedKeywordsPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DataTypeWeightedKeywords")

	deleteDataTypeWeightedKeywordsRes, err := client.ApiCall("delete-data-type-weighted-keywords", dataTypeWeightedKeywordsPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteDataTypeWeightedKeywordsRes.Success {
		if deleteDataTypeWeightedKeywordsRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDataTypeWeightedKeywordsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
