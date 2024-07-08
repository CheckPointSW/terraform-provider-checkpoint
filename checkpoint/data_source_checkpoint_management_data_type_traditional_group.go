package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementDataTypeTraditionalGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDataTypeTraditionalGroupRead,
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
			"data_types": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of data-types. If data matches any of the data types in the group, the data type group is matched. Identified by name or UID.",
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
func dataSourceManagementDataTypeTraditionalGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showDataTypeTraditionalGroupRes, err := client.ApiCall("show-data-type-traditional-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataTypeTraditionalGroupRes.Success {
		if objectNotFound(showDataTypeTraditionalGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDataTypeTraditionalGroupRes.ErrorMsg)
	}

	dataTypeTraditionalGroup := showDataTypeTraditionalGroupRes.GetData()

	log.Println("Read DataTypeTraditionalGroup - Show JSON = ", dataTypeTraditionalGroup)

	if v := dataTypeTraditionalGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := dataTypeTraditionalGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dataTypeTraditionalGroup["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if dataTypeTraditionalGroup["data-types"] != nil {
		dataTypesJson, ok := dataTypeTraditionalGroup["data-types"].([]interface{})
		if ok {
			dataTypesIds := make([]string, 0)
			if len(dataTypesJson) > 0 {
				for _, data_types := range dataTypesJson {
					data_types := data_types.(map[string]interface{})
					dataTypesIds = append(dataTypesIds, data_types["name"].(string))
				}
			}
			_ = d.Set("data_types", dataTypesIds)
		}
	} else {
		_ = d.Set("data_types", nil)
	}

	if dataTypeTraditionalGroup["tags"] != nil {
		tagsJson, ok := dataTypeTraditionalGroup["tags"].([]interface{})
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

	if v := dataTypeTraditionalGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dataTypeTraditionalGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
