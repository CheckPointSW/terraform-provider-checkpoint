package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementApplicationSiteGroup() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementApplicationSiteGroup,
        Read:   readManagementApplicationSiteGroup,
        Update: updateManagementApplicationSiteGroup,
        Delete: deleteManagementApplicationSiteGroup,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "members": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: "Collection of application and URL filtering objects identified by the name or UID.",
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

func createManagementApplicationSiteGroup(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    applicationSiteGroup := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        applicationSiteGroup["name"] = v.(string)
    }

    if v, ok := d.GetOk("members"); ok {
        applicationSiteGroup["members"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("tags"); ok {
        applicationSiteGroup["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        applicationSiteGroup["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        applicationSiteGroup["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        applicationSiteGroup["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        applicationSiteGroup["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        applicationSiteGroup["ignore-errors"] = v.(bool)
    }

    log.Println("Create ApplicationSiteGroup - Map = ", applicationSiteGroup)

    addApplicationSiteGroupRes, err := client.ApiCall("add-application-site-group", applicationSiteGroup, client.GetSessionID(), true, false)
    if err != nil || !addApplicationSiteGroupRes.Success {
        if addApplicationSiteGroupRes.ErrorMsg != "" {
            return fmt.Errorf(addApplicationSiteGroupRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addApplicationSiteGroupRes.GetData()["uid"].(string))

    return readManagementApplicationSiteGroup(d, m)
}

func readManagementApplicationSiteGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showApplicationSiteGroupRes, err := client.ApiCall("show-application-site-group", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showApplicationSiteGroupRes.Success {
		if objectNotFound(showApplicationSiteGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showApplicationSiteGroupRes.ErrorMsg)
    }

    applicationSiteGroup := showApplicationSiteGroupRes.GetData()

    log.Println("Read ApplicationSiteGroup - Show JSON = ", applicationSiteGroup)

	if v := applicationSiteGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

    if applicationSiteGroup["members"] != nil {
        membersJson, ok := applicationSiteGroup["members"].([]interface{})
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

    if applicationSiteGroup["tags"] != nil {
        tagsJson, ok := applicationSiteGroup["tags"].([]interface{})
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

	if v := applicationSiteGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := applicationSiteGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if applicationSiteGroup["groups"] != nil {
        groupsJson, ok := applicationSiteGroup["groups"].([]interface{})
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

	if v := applicationSiteGroup["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := applicationSiteGroup["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementApplicationSiteGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    applicationSiteGroup := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        applicationSiteGroup["name"] = oldName
        applicationSiteGroup["new-name"] = newName
    } else {
        applicationSiteGroup["name"] = d.Get("name")
    }

    if d.HasChange("members") {
        if v, ok := d.GetOk("members"); ok {
            applicationSiteGroup["members"] = v.(*schema.Set).List()
        } else {
            oldMembers, _ := d.GetChange("members")
	           applicationSiteGroup["members"] = map[string]interface{}{"remove": oldMembers.(*schema.Set).List()}
        }
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            applicationSiteGroup["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           applicationSiteGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       applicationSiteGroup["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       applicationSiteGroup["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            applicationSiteGroup["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           applicationSiteGroup["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       applicationSiteGroup["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       applicationSiteGroup["ignore-errors"] = v.(bool)
    }

    log.Println("Update ApplicationSiteGroup - Map = ", applicationSiteGroup)

    updateApplicationSiteGroupRes, err := client.ApiCall("set-application-site-group", applicationSiteGroup, client.GetSessionID(), true, false)
    if err != nil || !updateApplicationSiteGroupRes.Success {
        if updateApplicationSiteGroupRes.ErrorMsg != "" {
            return fmt.Errorf(updateApplicationSiteGroupRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementApplicationSiteGroup(d, m)
}

func deleteManagementApplicationSiteGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    applicationSiteGroupPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete ApplicationSiteGroup")

    deleteApplicationSiteGroupRes, err := client.ApiCall("delete-application-site-group", applicationSiteGroupPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteApplicationSiteGroupRes.Success {
        if deleteApplicationSiteGroupRes.ErrorMsg != "" {
            return fmt.Errorf(deleteApplicationSiteGroupRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

