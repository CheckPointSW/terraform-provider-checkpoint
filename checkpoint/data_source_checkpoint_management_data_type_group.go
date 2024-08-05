package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataResourceManagementDataTypeGroup() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementDataTypeGroupRead,

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
			"file_type": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of data-types-file-attributes objects. Identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"file_content": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of Data Types. At least one must be matched. Identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func dataSourceManagementDataTypeGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := dataTypeGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

	return nil
}
