package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementDataTypeCompoundGroup() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementDataTypeCompoundGroupRead,

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
			"matched_groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Each one of these data types must be matched - Select existing data types to add. Traffic must match all the data types of this group to match a rule. Identified by name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"unmatched_groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Each one of these data types must not be matched - Select existing data types to add to the definition. Traffic that does not contain any data matching the types in this list will match this compound data type. Identified by name or UID.",
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
func dataSourceManagementDataTypeCompoundGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := dataTypeCompoundGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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
