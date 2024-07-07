package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementDataTypeTraditionalGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDataTypeTraditionalGroup,
		Read:   readManagementDataTypeTraditionalGroup,
		Update: updateManagementDataTypeTraditionalGroup,
		Delete: deleteManagementDataTypeTraditionalGroup,
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
			"data_types": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of data-types. If data matches any of the data types in the group, the data type group is matched. Identified by name or UID.",
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

func createManagementDataTypeTraditionalGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dataTypeTraditionalGroup := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		dataTypeTraditionalGroup["name"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		dataTypeTraditionalGroup["description"] = v.(string)
	}

	if v, ok := d.GetOk("data_types"); ok {
		dataTypeTraditionalGroup["data-types"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		dataTypeTraditionalGroup["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dataTypeTraditionalGroup["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dataTypeTraditionalGroup["comments"] = v.(string)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeTraditionalGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeTraditionalGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Create DataTypeTraditionalGroup - Map = ", dataTypeTraditionalGroup)

	addDataTypeTraditionalGroupRes, err := client.ApiCall("add-data-type-traditional-group", dataTypeTraditionalGroup, client.GetSessionID(), true, false)
	if err != nil || !addDataTypeTraditionalGroupRes.Success {
		if addDataTypeTraditionalGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addDataTypeTraditionalGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDataTypeTraditionalGroupRes.GetData()["uid"].(string))

	return readManagementDataTypeTraditionalGroup(d, m)
}

func readManagementDataTypeTraditionalGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
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
	if v := dataTypeTraditionalGroup["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dataTypeTraditionalGroup["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}
	return nil

}

func updateManagementDataTypeTraditionalGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dataTypeTraditionalGroup := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		dataTypeTraditionalGroup["name"] = oldName
		dataTypeTraditionalGroup["new-name"] = newName
	} else {
		dataTypeTraditionalGroup["name"] = d.Get("name")
	}

	if ok := d.HasChange("description"); ok {
		dataTypeTraditionalGroup["description"] = d.Get("description")
	}

	if d.HasChange("data_types") {
		if v, ok := d.GetOk("data_types"); ok {
			dataTypeTraditionalGroup["data-types"] = v.(*schema.Set).List()
		} else {
			oldData_Types, _ := d.GetChange("data_types")
			dataTypeTraditionalGroup["data-types"] = map[string]interface{}{"remove": oldData_Types.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dataTypeTraditionalGroup["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dataTypeTraditionalGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dataTypeTraditionalGroup["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dataTypeTraditionalGroup["comments"] = d.Get("comments")
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeTraditionalGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeTraditionalGroup["ignore-errors"] = v.(bool)
	}
	log.Println("Update DataTypeTraditionalGroup - Map = ", dataTypeTraditionalGroup)

	updateDataTypeTraditionalGroupRes, err := client.ApiCall("set-data-type-traditional-group", dataTypeTraditionalGroup, client.GetSessionID(), true, false)
	if err != nil || !updateDataTypeTraditionalGroupRes.Success {
		if updateDataTypeTraditionalGroupRes.ErrorMsg != "" {
			return fmt.Errorf(updateDataTypeTraditionalGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDataTypeTraditionalGroup(d, m)
}

func deleteManagementDataTypeTraditionalGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataTypeTraditionalGroupPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DataTypeTraditionalGroup")

	deleteDataTypeTraditionalGroupRes, err := client.ApiCall("delete-data-type-traditional-group", dataTypeTraditionalGroupPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteDataTypeTraditionalGroupRes.Success {
		if deleteDataTypeTraditionalGroupRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDataTypeTraditionalGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
