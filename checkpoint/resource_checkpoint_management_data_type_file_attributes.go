package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementDataTypeFileAttributes() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDataTypeFileAttributes,
		Read:   readManagementDataTypeFileAttributes,
		Update: updateManagementDataTypeFileAttributes,
		Delete: deleteManagementDataTypeFileAttributes,
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
			"match_by_file_type": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determine whether to consider file type.",
				Default:     false,
			},
			"file_groups_list": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The file must be one of the types specified in the list. Identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"match_by_file_name": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determine whether to consider file name.",
				Default:     false,
			},
			"file_name_contains": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File name should contain the expression.",
			},
			"match_by_file_size": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Determine whether to consider file size.",
				Default:     false,
			},
			"file_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Min File size in KB.",
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

func createManagementDataTypeFileAttributes(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dataTypeFileAttributes := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		dataTypeFileAttributes["name"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		dataTypeFileAttributes["description"] = v.(string)
	}

	if v, ok := d.GetOkExists("match_by_file_type"); ok {
		dataTypeFileAttributes["match-by-file-type"] = v.(bool)
	}

	if v, ok := d.GetOk("file_groups_list"); ok {
		dataTypeFileAttributes["file-groups-list"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("match_by_file_name"); ok {
		dataTypeFileAttributes["match-by-file-name"] = v.(bool)
	}

	if v, ok := d.GetOk("file_name_contains"); ok {
		dataTypeFileAttributes["file-name-contains"] = v.(string)
	}

	if v, ok := d.GetOkExists("match_by_file_size"); ok {
		dataTypeFileAttributes["match-by-file-size"] = v.(bool)
	}

	if v, ok := d.GetOk("file_size"); ok {
		dataTypeFileAttributes["file-size"] = v.(int)
	}

	if v, ok := d.GetOk("tags"); ok {
		dataTypeFileAttributes["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dataTypeFileAttributes["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dataTypeFileAttributes["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeFileAttributes["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeFileAttributes["ignore-errors"] = v.(bool)
	}

	log.Println("Create DataTypeFileAttributes - Map = ", dataTypeFileAttributes)

	addDataTypeFileAttributesRes, err := client.ApiCall("add-data-type-file-attributes", dataTypeFileAttributes, client.GetSessionID(), true, false)
	if err != nil || !addDataTypeFileAttributesRes.Success {
		if addDataTypeFileAttributesRes.ErrorMsg != "" {
			return fmt.Errorf(addDataTypeFileAttributesRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDataTypeFileAttributesRes.GetData()["uid"].(string))

	return readManagementDataTypeFileAttributes(d, m)
}

func readManagementDataTypeFileAttributes(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showDataTypeFileAttributesRes, err := client.ApiCall("show-data-type-file-attributes", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataTypeFileAttributesRes.Success {
		if objectNotFound(showDataTypeFileAttributesRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDataTypeFileAttributesRes.ErrorMsg)
	}

	dataTypeFileAttributes := showDataTypeFileAttributesRes.GetData()

	log.Println("Read DataTypeFileAttributes - Show JSON = ", dataTypeFileAttributes)

	if v := dataTypeFileAttributes["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dataTypeFileAttributes["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := dataTypeFileAttributes["match-by-file-type"]; v != nil {
		_ = d.Set("match_by_file_type", v)
	}

	if v := dataTypeFileAttributes["file-groups-list"]; v != nil {

		var fileListToReturn []interface{}
		fileList := v.([]interface{})
		for i := range fileList {
			fileName := fileList[i].(map[string]interface{})
			fileListToReturn = append(fileListToReturn, fileName["name"])

		}
		_ = d.Set("file_groups_list", fileListToReturn)

	} else {
		_ = d.Set("file_groups_list", nil)
	}

	if v := dataTypeFileAttributes["match-by-file-name"]; v != nil {
		_ = d.Set("match_by_file_name", v)
	}

	if v := dataTypeFileAttributes["file-name-contains"]; v != nil {
		_ = d.Set("file_name_contains", v)
	}

	if v := dataTypeFileAttributes["match-by-file-size"]; v != nil {
		_ = d.Set("match_by_file_size", v)
	}

	if v := dataTypeFileAttributes["file-size"]; v != nil {
		_ = d.Set("file_size", v)
	}

	if dataTypeFileAttributes["tags"] != nil {
		tagsJson, ok := dataTypeFileAttributes["tags"].([]interface{})
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

	if v := dataTypeFileAttributes["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dataTypeFileAttributes["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dataTypeFileAttributes["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dataTypeFileAttributes["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDataTypeFileAttributes(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dataTypeFileAttributes := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		dataTypeFileAttributes["name"] = oldName
		dataTypeFileAttributes["new-name"] = newName
	} else {
		dataTypeFileAttributes["name"] = d.Get("name")
	}

	if ok := d.HasChange("description"); ok {
		dataTypeFileAttributes["description"] = d.Get("description")
	}

	if v, ok := d.GetOkExists("match_by_file_type"); ok {
		dataTypeFileAttributes["match-by-file-type"] = v.(bool)
	}

	if d.HasChange("file_groups_list") {
		if v, ok := d.GetOk("file_groups_list"); ok {
			dataTypeFileAttributes["file-groups-list"] = v.(*schema.Set).List()
		} else {
			oldFile_Groups_List, _ := d.GetChange("file_groups_list")
			dataTypeFileAttributes["file_groups_list"] = map[string]interface{}{"remove": oldFile_Groups_List.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("match_by_file_name"); ok {
		dataTypeFileAttributes["match-by-file-name"] = v.(bool)
	}

	if ok := d.HasChange("file_name_contains"); ok {
		dataTypeFileAttributes["file-name-contains"] = d.Get("file_name_contains")
	}

	if v, ok := d.GetOkExists("match_by_file_size"); ok {
		dataTypeFileAttributes["match-by-file-size"] = v.(bool)
	}

	if ok := d.HasChange("file_size"); ok {
		dataTypeFileAttributes["file-size"] = d.Get("file_size")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dataTypeFileAttributes["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dataTypeFileAttributes["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dataTypeFileAttributes["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dataTypeFileAttributes["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeFileAttributes["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeFileAttributes["ignore-errors"] = v.(bool)
	}

	log.Println("Update DataTypeFileAttributes - Map = ", dataTypeFileAttributes)

	updateDataTypeFileAttributesRes, err := client.ApiCall("set-data-type-file-attributes", dataTypeFileAttributes, client.GetSessionID(), true, false)
	if err != nil || !updateDataTypeFileAttributesRes.Success {
		if updateDataTypeFileAttributesRes.ErrorMsg != "" {
			return fmt.Errorf(updateDataTypeFileAttributesRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDataTypeFileAttributes(d, m)
}

func deleteManagementDataTypeFileAttributes(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataTypeFileAttributesPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DataTypeFileAttributes")

	deleteDataTypeFileAttributesRes, err := client.ApiCall("delete-data-type-file-attributes", dataTypeFileAttributesPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteDataTypeFileAttributesRes.Success {
		if deleteDataTypeFileAttributesRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDataTypeFileAttributesRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
