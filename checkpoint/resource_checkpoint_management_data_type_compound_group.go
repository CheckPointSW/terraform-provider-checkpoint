package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementDataTypeCompoundGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDataTypeCompoundGroup,
		Read:   readManagementDataTypeCompoundGroup,
		Update: updateManagementDataTypeCompoundGroup,
		Delete: deleteManagementDataTypeCompoundGroup,
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
			"matched_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Each one of these data types must be matched - Select existing data types to add. Traffic must match all the data types of this group to match a rule. Identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"unmatched_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Each one of these data types must not be matched - Select existing data types to add to the definition. Traffic that does not contain any data matching the types in this list will match this compound data type. Identified by name or UID.",
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

func createManagementDataTypeCompoundGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dataTypeCompoundGroup := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		dataTypeCompoundGroup["name"] = v.(string)
	}

	if v, ok := d.GetOk("description"); ok {
		dataTypeCompoundGroup["description"] = v.(string)
	}

	if v, ok := d.GetOk("matched_groups"); ok {
		dataTypeCompoundGroup["matched-groups"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("unmatched_groups"); ok {
		dataTypeCompoundGroup["unmatched-groups"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		dataTypeCompoundGroup["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dataTypeCompoundGroup["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dataTypeCompoundGroup["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeCompoundGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeCompoundGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Create DataTypeCompoundGroup - Map = ", dataTypeCompoundGroup)

	addDataTypeCompoundGroupRes, err := client.ApiCall("add-data-type-compound-group", dataTypeCompoundGroup, client.GetSessionID(), true, false)
	if err != nil || !addDataTypeCompoundGroupRes.Success {
		if addDataTypeCompoundGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addDataTypeCompoundGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDataTypeCompoundGroupRes.GetData()["uid"].(string))

	return readManagementDataTypeCompoundGroup(d, m)
}

func readManagementDataTypeCompoundGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showDataTypeCompoundGroupRes, err := client.ApiCall("show-data-type-compound-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDataTypeCompoundGroupRes.Success {
		if objectNotFound(showDataTypeCompoundGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDataTypeCompoundGroupRes.ErrorMsg)
	}

	dataTypeCompoundGroup := showDataTypeCompoundGroupRes.GetData()

	log.Println("Read DataTypeCompoundGroup - Show JSON = ", dataTypeCompoundGroup)

	if v := dataTypeCompoundGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dataTypeCompoundGroup["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if dataTypeCompoundGroup["matched-groups"] != nil {
		matchedGroupsJson, ok := dataTypeCompoundGroup["matched-groups"].([]interface{})
		if ok {
			matchedGroupsIds := make([]string, 0)
			if len(matchedGroupsJson) > 0 {
				for _, matched_groups := range matchedGroupsJson {
					matched_groups := matched_groups.(map[string]interface{})
					matchedGroupsIds = append(matchedGroupsIds, matched_groups["name"].(string))
				}
			}
			_ = d.Set("matched_groups", matchedGroupsIds)
		}
	} else {
		_ = d.Set("matched_groups", nil)
	}

	if dataTypeCompoundGroup["unmatched-groups"] != nil {
		unmatchedGroupsJson, ok := dataTypeCompoundGroup["unmatched-groups"].([]interface{})
		if ok {
			unmatchedGroupsIds := make([]string, 0)
			if len(unmatchedGroupsJson) > 0 {
				for _, unmatched_groups := range unmatchedGroupsJson {
					unmatched_groups := unmatched_groups.(map[string]interface{})
					unmatchedGroupsIds = append(unmatchedGroupsIds, unmatched_groups["name"].(string))
				}
			}
			_ = d.Set("unmatched_groups", unmatchedGroupsIds)
		}
	} else {
		_ = d.Set("unmatched_groups", nil)
	}

	if dataTypeCompoundGroup["tags"] != nil {
		tagsJson, ok := dataTypeCompoundGroup["tags"].([]interface{})
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

	if v := dataTypeCompoundGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dataTypeCompoundGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dataTypeCompoundGroup["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dataTypeCompoundGroup["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDataTypeCompoundGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dataTypeCompoundGroup := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		dataTypeCompoundGroup["name"] = oldName
		dataTypeCompoundGroup["new-name"] = newName
	} else {
		dataTypeCompoundGroup["name"] = d.Get("name")
	}

	if ok := d.HasChange("description"); ok {
		dataTypeCompoundGroup["description"] = d.Get("description")
	}

	if d.HasChange("matched_groups") {
		if v, ok := d.GetOk("matched_groups"); ok {
			dataTypeCompoundGroup["matched-groups"] = v.(*schema.Set).List()
		} else {
			oldMatched_Groups, _ := d.GetChange("matched_groups")
			dataTypeCompoundGroup["matched-groups"] = map[string]interface{}{"remove": oldMatched_Groups.(*schema.Set).List()}
		}
	}

	if d.HasChange("unmatched_groups") {
		if v, ok := d.GetOk("unmatched_groups"); ok {
			dataTypeCompoundGroup["unmatched-groups"] = v.(*schema.Set).List()
		} else {
			oldUnmatched_Groups, _ := d.GetChange("unmatched_groups")
			dataTypeCompoundGroup["unmatched-groups"] = map[string]interface{}{"remove": oldUnmatched_Groups.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dataTypeCompoundGroup["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dataTypeCompoundGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dataTypeCompoundGroup["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dataTypeCompoundGroup["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dataTypeCompoundGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dataTypeCompoundGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Update DataTypeCompoundGroup - Map = ", dataTypeCompoundGroup)

	updateDataTypeCompoundGroupRes, err := client.ApiCall("set-data-type-compound-group", dataTypeCompoundGroup, client.GetSessionID(), true, false)
	if err != nil || !updateDataTypeCompoundGroupRes.Success {
		if updateDataTypeCompoundGroupRes.ErrorMsg != "" {
			return fmt.Errorf(updateDataTypeCompoundGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDataTypeCompoundGroup(d, m)
}

func deleteManagementDataTypeCompoundGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dataTypeCompoundGroupPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DataTypeCompoundGroup")

	deleteDataTypeCompoundGroupRes, err := client.ApiCall("delete-data-type-compound-group", dataTypeCompoundGroupPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteDataTypeCompoundGroupRes.Success {
		if deleteDataTypeCompoundGroupRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDataTypeCompoundGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
