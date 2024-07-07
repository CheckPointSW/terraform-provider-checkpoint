package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementDataTypeFileAttributes() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementDataTypeFileAttributesRead,

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
			"match_by_file_type": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determine whether to consider file type.",
			},
			"file_groups_list": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "The file must be one of the types specified in the list. Identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"match_by_file_name": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determine whether to consider file name.",
			},
			"file_name_contains": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "File name should contain the expression.",
			},
			"match_by_file_size": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Determine whether to consider file size.",
			},
			"file_size": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Min File size in KB.",
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

func dataSourceManagementDataTypeFileAttributesRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := dataTypeFileAttributes["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

	return nil

}
