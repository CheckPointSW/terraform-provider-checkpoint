package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementGroupWithExclusion() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementGroupWithExclusion,
        Read:   readManagementGroupWithExclusion,
        Update: updateManagementGroupWithExclusion,
        Delete: deleteManagementGroupWithExclusion,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "except": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Name or UID of an object which the group excludes.",
            },
            "include": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Name or UID of an object which the group includes.",
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
            "groups": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Collection of group identifiers.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
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

func createManagementGroupWithExclusion(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    groupWithExclusion := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        groupWithExclusion["name"] = v.(string)
    }

    if v, ok := d.GetOk("except"); ok {
        groupWithExclusion["except"] = v.(string)
    }

    if v, ok := d.GetOk("include"); ok {
        groupWithExclusion["include"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        groupWithExclusion["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        groupWithExclusion["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        groupWithExclusion["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        groupWithExclusion["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        groupWithExclusion["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        groupWithExclusion["ignore-errors"] = v.(bool)
    }

    log.Println("Create GroupWithExclusion - Map = ", groupWithExclusion)

    addGroupWithExclusionRes, err := client.ApiCall("add-group-with-exclusion", groupWithExclusion, client.GetSessionID(), true, false)
    if err != nil || !addGroupWithExclusionRes.Success {
        if addGroupWithExclusionRes.ErrorMsg != "" {
            return fmt.Errorf(addGroupWithExclusionRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addGroupWithExclusionRes.GetData()["uid"].(string))

    return readManagementGroupWithExclusion(d, m)
}

func readManagementGroupWithExclusion(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showGroupWithExclusionRes, err := client.ApiCall("show-group-with-exclusion", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showGroupWithExclusionRes.Success {
		if objectNotFound(showGroupWithExclusionRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showGroupWithExclusionRes.ErrorMsg)
    }

    groupWithExclusion := showGroupWithExclusionRes.GetData()

    log.Println("Read GroupWithExclusion - Show JSON = ", groupWithExclusion)

	if v := groupWithExclusion["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := groupWithExclusion["except"]; v != nil {
		_ = d.Set("except", v)
	}

	if v := groupWithExclusion["include"]; v != nil {
		_ = d.Set("include", v)
	}

    if groupWithExclusion["tags"] != nil {
        tagsJson, ok := groupWithExclusion["tags"].([]interface{})
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

	if v := groupWithExclusion["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := groupWithExclusion["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if groupWithExclusion["groups"] != nil {
        groupsJson, ok := groupWithExclusion["groups"].([]interface{})
        if ok {
            groupsIds := make([]string, 0)
            if len(groupsJson) > 0 {
                for _, groups := range groupsJson {
                    groups := groups.(map[string]interface{})
                    groupsIds = append(groupsIds, groups["name"].(string))
                }
            }
        _ = d.Set("groups", groupsIds)
        }
    } else {
        _ = d.Set("groups", nil)
    }

	if v := groupWithExclusion["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := groupWithExclusion["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementGroupWithExclusion(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    groupWithExclusion := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        groupWithExclusion["name"] = oldName
        groupWithExclusion["new-name"] = newName
    } else {
        groupWithExclusion["name"] = d.Get("name")
    }

    if ok := d.HasChange("except"); ok {
	       groupWithExclusion["except"] = d.Get("except")
    }

    if ok := d.HasChange("include"); ok {
	       groupWithExclusion["include"] = d.Get("include")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            groupWithExclusion["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           groupWithExclusion["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       groupWithExclusion["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       groupWithExclusion["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            groupWithExclusion["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           groupWithExclusion["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       groupWithExclusion["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       groupWithExclusion["ignore-errors"] = v.(bool)
    }

    log.Println("Update GroupWithExclusion - Map = ", groupWithExclusion)

    updateGroupWithExclusionRes, err := client.ApiCall("set-group-with-exclusion", groupWithExclusion, client.GetSessionID(), true, false)
    if err != nil || !updateGroupWithExclusionRes.Success {
        if updateGroupWithExclusionRes.ErrorMsg != "" {
            return fmt.Errorf(updateGroupWithExclusionRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementGroupWithExclusion(d, m)
}

func deleteManagementGroupWithExclusion(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    groupWithExclusionPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete GroupWithExclusion")

    deleteGroupWithExclusionRes, err := client.ApiCall("delete-group-with-exclusion", groupWithExclusionPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteGroupWithExclusionRes.Success {
        if deleteGroupWithExclusionRes.ErrorMsg != "" {
            return fmt.Errorf(deleteGroupWithExclusionRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

