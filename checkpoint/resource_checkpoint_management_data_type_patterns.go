package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementDataTypePatterns() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDataTypePatterns,
		Read:   readManagementDataTypePatterns,
		Update: updateManagementDataTypePatterns,
		Delete: deleteManagementDataTypePatterns,
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
			"patterns": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Regular expressions to be evaluated.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"number_of_occurrences": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Define how many times the patterns must appear to be considered data to be protected.",
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

func createManagementDataTypePatterns(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dataTypePatterns := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		dataTypePatterns["name"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		dataTypePatterns["description"] = v.(string)
	}

	if v, ok := d.GetOk("patterns"); ok {
		dataTypePatterns["patterns"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("number_of_occurrences"); ok {
		dataTypePatterns["number-of-occurrences"] = v.(int)
	}

	if v, ok := d.GetOk("tags"); ok {
		dataTypePatterns["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dataTypePatterns["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dataTypePatterns["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypePatterns["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypePatterns["ignore-errors"] = v.(bool)
	}

	log.Println("Create DataTypePatterns - Map = ", dataTypePatterns)

	addDataTypePatternsRes, err := client.ApiCall("add-data-type-patterns", dataTypePatterns, client.GetSessionID(), true, false)
	if err != nil || !addDataTypePatternsRes.Success {
		if addDataTypePatternsRes.ErrorMsg != "" {
			return fmt.Errorf(addDataTypePatternsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDataTypePatternsRes.GetData()["uid"].(string))

	return readManagementDataTypePatterns(d, m)
}

func readManagementDataTypePatterns(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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

	if v := dataTypePatterns["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dataTypePatterns["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDataTypePatterns(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dataTypePatterns := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		dataTypePatterns["name"] = oldName
		dataTypePatterns["new-name"] = newName
	} else {
		dataTypePatterns["name"] = d.Get("name")
	}

	if ok := d.HasChange("description"); ok {
		dataTypePatterns["description"] = d.Get("description")
	}

	if d.HasChange("patterns") {
		if v, ok := d.GetOk("patterns"); ok {
			dataTypePatterns["patterns"] = v.(*schema.Set).List()
		} else {
			oldPatterns, _ := d.GetChange("patterns")
			dataTypePatterns["patterns"] = map[string]interface{}{"remove": oldPatterns.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("number_of_occurrences"); ok {
		dataTypePatterns["number-of-occurrences"] = d.Get("number_of_occurrences")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dataTypePatterns["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dataTypePatterns["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dataTypePatterns["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dataTypePatterns["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypePatterns["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypePatterns["ignore-errors"] = v.(bool)
	}

	log.Println("Update DataTypePatterns - Map = ", dataTypePatterns)

	updateDataTypePatternsRes, err := client.ApiCall("set-data-type-patterns", dataTypePatterns, client.GetSessionID(), true, false)
	if err != nil || !updateDataTypePatternsRes.Success {
		if updateDataTypePatternsRes.ErrorMsg != "" {
			return fmt.Errorf(updateDataTypePatternsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDataTypePatterns(d, m)
}

func deleteManagementDataTypePatterns(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataTypePatternsPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DataTypePatterns")

	deleteDataTypePatternsRes, err := client.ApiCall("delete-data-type-patterns", dataTypePatternsPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteDataTypePatternsRes.Success {
		if deleteDataTypePatternsRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDataTypePatternsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
