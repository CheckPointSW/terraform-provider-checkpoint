package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementDataTypeGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDataTypeGroup,
		Read:   readManagementDataTypeGroup,
		Update: updateManagementDataTypeGroup,
		Delete: deleteManagementDataTypeGroup,
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
			"file_type": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of data-types-file-attributes objects. Identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"file_content": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of Data Types. At least one must be matched. Identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createManagementDataTypeGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dataTypeGroup := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		dataTypeGroup["name"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		dataTypeGroup["description"] = v.(string)
	}

	if v, ok := d.GetOk("file_type"); ok {
		dataTypeGroup["file-type"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("file_content"); ok {
		dataTypeGroup["file-content"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		dataTypeGroup["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dataTypeGroup["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dataTypeGroup["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Create DataTypeGroup - Map = ", dataTypeGroup)

	addDataTypeGroupRes, err := client.ApiCall("add-data-type-group", dataTypeGroup, client.GetSessionID(), true, false)
	if err != nil || !addDataTypeGroupRes.Success {
		if addDataTypeGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addDataTypeGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDataTypeGroupRes.GetData()["uid"].(string))

	return readManagementDataTypeGroup(d, m)
}

func readManagementDataTypeGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showDataTypeGroupRes, err := client.ApiCall("show-data-type-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataTypeGroupRes.Success {
		if objectNotFound(showDataTypeGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDataTypeGroupRes.ErrorMsg)
	}

	dataTypeGroup := showDataTypeGroupRes.GetData()

	log.Println("Read DataTypeGroup - Show JSON = ", dataTypeGroup)

	if v := dataTypeGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dataTypeGroup["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if dataTypeGroup["file-type"] != nil {
		fileTypeJson, ok := dataTypeGroup["file-type"].([]interface{})
		if ok {
			fileTypeIds := make([]string, 0)
			if len(fileTypeJson) > 0 {
				for _, file_type := range fileTypeJson {
					file_type := file_type.(map[string]interface{})
					fileTypeIds = append(fileTypeIds, file_type["name"].(string))
				}
			}
			_ = d.Set("file_type", fileTypeIds)
		}
	} else {
		_ = d.Set("file_type", nil)
	}

	if dataTypeGroup["file-content"] != nil {
		fileContentJson, ok := dataTypeGroup["file-content"].([]interface{})
		if ok {
			fileContentIds := make([]string, 0)
			if len(fileContentJson) > 0 {
				for _, file_content := range fileContentJson {
					file_content := file_content.(map[string]interface{})
					fileContentIds = append(fileContentIds, file_content["name"].(string))
				}
			}
			_ = d.Set("file_content", fileContentIds)
		}
	} else {
		_ = d.Set("file_content", nil)
	}

	if dataTypeGroup["tags"] != nil {
		tagsJson, ok := dataTypeGroup["tags"].([]interface{})
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

	if v := dataTypeGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dataTypeGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dataTypeGroup["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dataTypeGroup["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDataTypeGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dataTypeGroup := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		dataTypeGroup["name"] = oldName
		dataTypeGroup["new-name"] = newName
	} else {
		dataTypeGroup["name"] = d.Get("name")
	}

	if ok := d.HasChange("description"); ok {
		dataTypeGroup["description"] = d.Get("description")
	}

	if d.HasChange("file_type") {
		if v, ok := d.GetOk("file_type"); ok {
			dataTypeGroup["file-type"] = v.(*schema.Set).List()
		} else {
			oldData_Types, _ := d.GetChange("file_type")
			dataTypeGroup["file-type"] = map[string]interface{}{"remove": oldData_Types.(*schema.Set).List()}
		}
	}

	if d.HasChange("file_content") {
		if v, ok := d.GetOk("file_content"); ok {
			dataTypeGroup["file-content"] = v.(*schema.Set).List()
		} else {
			oldData_Types, _ := d.GetChange("file_content")
			dataTypeGroup["file-content"] = map[string]interface{}{"remove": oldData_Types.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dataTypeGroup["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dataTypeGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dataTypeGroup["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dataTypeGroup["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Update DataTypeGroup - Map = ", dataTypeGroup)

	updateDataTypeGroupRes, err := client.ApiCall("set-data-type-group", dataTypeGroup, client.GetSessionID(), true, false)
	if err != nil || !updateDataTypeGroupRes.Success {
		if updateDataTypeGroupRes.ErrorMsg != "" {
			return fmt.Errorf(updateDataTypeGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDataTypeGroup(d, m)
}

func deleteManagementDataTypeGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataTypeGroupPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DataTypeGroup")

	deleteDataTypeGroupRes, err := client.ApiCall("delete-data-type-group", dataTypeGroupPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteDataTypeGroupRes.Success {
		if deleteDataTypeGroupRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDataTypeGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
