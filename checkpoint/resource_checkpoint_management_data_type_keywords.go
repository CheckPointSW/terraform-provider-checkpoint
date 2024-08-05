package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementDataTypeKeywords() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDataTypeKeywords,
		Read:   readManagementDataTypeKeywords,
		Update: updateManagementDataTypeKeywords,
		Delete: deleteManagementDataTypeKeywords,
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
			"keywords": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Specify keywords or phrases to search for. cannot be empty",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"data_match_threshold": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "If set to all-keywords - the data will be matched to the rule only if all the words in the list appear in the data contents. When set to min-keywords any number of the words may appear according to configuration.",
				Default:     "min-keywords",
			},
			"min_number_of_keywords": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Define how many of the words in the list must appear in the contents of the data to match the rule. min value is 1.",
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

func createManagementDataTypeKeywords(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dataTypeKeywords := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		dataTypeKeywords["name"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		dataTypeKeywords["description"] = v.(string)
	}

	if v, ok := d.GetOk("keywords"); ok {
		dataTypeKeywords["keywords"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("data_match_threshold"); ok {
		dataTypeKeywords["data-match-threshold"] = v.(string)
	}

	if v, ok := d.GetOk("min_number_of_keywords"); ok {
		dataTypeKeywords["min-number-of-keywords"] = v.(int)
	}

	if v, ok := d.GetOk("tags"); ok {
		dataTypeKeywords["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dataTypeKeywords["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dataTypeKeywords["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeKeywords["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeKeywords["ignore-errors"] = v.(bool)
	}

	log.Println("Create DataTypeKeywords - Map = ", dataTypeKeywords)

	addDataTypeKeywordsRes, err := client.ApiCall("add-data-type-keywords", dataTypeKeywords, client.GetSessionID(), true, false)
	if err != nil || !addDataTypeKeywordsRes.Success {
		if addDataTypeKeywordsRes.ErrorMsg != "" {
			return fmt.Errorf(addDataTypeKeywordsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDataTypeKeywordsRes.GetData()["uid"].(string))

	return readManagementDataTypeKeywords(d, m)
}

func readManagementDataTypeKeywords(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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

	if v := dataTypeKeywords["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dataTypeKeywords["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDataTypeKeywords(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dataTypeKeywords := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		dataTypeKeywords["name"] = oldName
		dataTypeKeywords["new-name"] = newName
	} else {
		dataTypeKeywords["name"] = d.Get("name")
	}

	if ok := d.HasChange("description"); ok {
		dataTypeKeywords["description"] = d.Get("description")
	}

	if d.HasChange("keywords") {
		if v, ok := d.GetOk("keywords"); ok {
			dataTypeKeywords["keywords"] = v.(*schema.Set).List()
		} else {
			oldKeywords, _ := d.GetChange("keywords")
			dataTypeKeywords["keywords"] = map[string]interface{}{"remove": oldKeywords.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("data_match_threshold"); ok {
		dataTypeKeywords["data-match-threshold"] = d.Get("data_match_threshold")
	}

	if ok := d.HasChange("min_number_of_keywords"); ok {
		dataTypeKeywords["min-number-of-keywords"] = d.Get("min_number_of_keywords")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dataTypeKeywords["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dataTypeKeywords["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dataTypeKeywords["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dataTypeKeywords["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeKeywords["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeKeywords["ignore-errors"] = v.(bool)
	}

	log.Println("Update DataTypeKeywords - Map = ", dataTypeKeywords)

	updateDataTypeKeywordsRes, err := client.ApiCall("set-data-type-keywords", dataTypeKeywords, client.GetSessionID(), true, false)
	if err != nil || !updateDataTypeKeywordsRes.Success {
		if updateDataTypeKeywordsRes.ErrorMsg != "" {
			return fmt.Errorf(updateDataTypeKeywordsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDataTypeKeywords(d, m)
}

func deleteManagementDataTypeKeywords(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataTypeKeywordsPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DataTypeKeywords")

	deleteDataTypeKeywordsRes, err := client.ApiCall("delete-data-type-keywords", dataTypeKeywordsPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteDataTypeKeywordsRes.Success {
		if deleteDataTypeKeywordsRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDataTypeKeywordsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
