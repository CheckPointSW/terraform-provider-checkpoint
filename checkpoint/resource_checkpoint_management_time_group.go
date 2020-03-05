package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementTimeGroup() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementTimeGroup,
        Read:   readManagementTimeGroup,
        Update: updateManagementTimeGroup,
        Delete: deleteManagementTimeGroup,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "members": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Collection of Time Group objects identified by the name or UID.",
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

func createManagementTimeGroup(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    timeGroup := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        timeGroup["name"] = v.(string)
    }

    if v, ok := d.GetOk("members"); ok {
        timeGroup["members"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("tags"); ok {
        timeGroup["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        timeGroup["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        timeGroup["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        timeGroup["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        timeGroup["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        timeGroup["ignore-errors"] = v.(bool)
    }

    log.Println("Create TimeGroup - Map = ", timeGroup)

    addTimeGroupRes, err := client.ApiCall("add-time-group", timeGroup, client.GetSessionID(), true, false)
    if err != nil || !addTimeGroupRes.Success {
        if addTimeGroupRes.ErrorMsg != "" {
            return fmt.Errorf(addTimeGroupRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addTimeGroupRes.GetData()["uid"].(string))

    return readManagementTimeGroup(d, m)
}

func readManagementTimeGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showTimeGroupRes, err := client.ApiCall("show-time-group", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showTimeGroupRes.Success {
		if objectNotFound(showTimeGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showTimeGroupRes.ErrorMsg)
    }

    timeGroup := showTimeGroupRes.GetData()

    log.Println("Read TimeGroup - Show JSON = ", timeGroup)

	if v := timeGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

    if timeGroup["members"] != nil {
        membersJson, ok := timeGroup["members"].([]interface{})
        if ok {
            membersIds := make([]string, 0)
            if len(membersJson) > 0 {
                for _, members := range membersJson {
                    members := members.(map[string]interface{})
                    membersIds = append(membersIds, members["name"].(string))
                }
            }
        _ = d.Set("members", membersIds)
        }
    } else {
        _ = d.Set("members", nil)
    }

    if timeGroup["tags"] != nil {
        tagsJson, ok := timeGroup["tags"].([]interface{})
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

	if v := timeGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := timeGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if timeGroup["groups"] != nil {
        groupsJson, ok := timeGroup["groups"].([]interface{})
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

	if v := timeGroup["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := timeGroup["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementTimeGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    timeGroup := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        timeGroup["name"] = oldName
        timeGroup["new-name"] = newName
    } else {
        timeGroup["name"] = d.Get("name")
    }

    if d.HasChange("members") {
        if v, ok := d.GetOk("members"); ok {
            timeGroup["members"] = v.(*schema.Set).List()
        } else {
            oldMembers, _ := d.GetChange("members")
	           timeGroup["members"] = map[string]interface{}{"remove": oldMembers.(*schema.Set).List()}
        }
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            timeGroup["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           timeGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       timeGroup["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       timeGroup["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            timeGroup["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           timeGroup["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       timeGroup["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       timeGroup["ignore-errors"] = v.(bool)
    }

    log.Println("Update TimeGroup - Map = ", timeGroup)

    updateTimeGroupRes, err := client.ApiCall("set-time-group", timeGroup, client.GetSessionID(), true, false)
    if err != nil || !updateTimeGroupRes.Success {
        if updateTimeGroupRes.ErrorMsg != "" {
            return fmt.Errorf(updateTimeGroupRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementTimeGroup(d, m)
}

func deleteManagementTimeGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    timeGroupPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete TimeGroup")

    deleteTimeGroupRes, err := client.ApiCall("delete-time-group", timeGroupPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteTimeGroupRes.Success {
        if deleteTimeGroupRes.ErrorMsg != "" {
            return fmt.Errorf(deleteTimeGroupRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

