package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementApplicationSiteCategory() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementApplicationSiteCategory,
        Read:   readManagementApplicationSiteCategory,
        Update: updateManagementApplicationSiteCategory,
        Delete: deleteManagementApplicationSiteCategory,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "N/A",
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

func createManagementApplicationSiteCategory(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    applicationSiteCategory := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        applicationSiteCategory["name"] = v.(string)
    }

    if v, ok := d.GetOk("description"); ok {
        applicationSiteCategory["description"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        applicationSiteCategory["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        applicationSiteCategory["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        applicationSiteCategory["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        applicationSiteCategory["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        applicationSiteCategory["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        applicationSiteCategory["ignore-errors"] = v.(bool)
    }

    log.Println("Create ApplicationSiteCategory - Map = ", applicationSiteCategory)

    addApplicationSiteCategoryRes, err := client.ApiCall("add-application-site-category", applicationSiteCategory, client.GetSessionID(), true, false)
    if err != nil || !addApplicationSiteCategoryRes.Success {
        if addApplicationSiteCategoryRes.ErrorMsg != "" {
            return fmt.Errorf(addApplicationSiteCategoryRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addApplicationSiteCategoryRes.GetData()["uid"].(string))

    return readManagementApplicationSiteCategory(d, m)
}

func readManagementApplicationSiteCategory(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showApplicationSiteCategoryRes, err := client.ApiCall("show-application-site-category", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showApplicationSiteCategoryRes.Success {
		if objectNotFound(showApplicationSiteCategoryRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showApplicationSiteCategoryRes.ErrorMsg)
    }

    applicationSiteCategory := showApplicationSiteCategoryRes.GetData()

    log.Println("Read ApplicationSiteCategory - Show JSON = ", applicationSiteCategory)

	if v := applicationSiteCategory["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := applicationSiteCategory["description"]; v != nil {
		_ = d.Set("description", v)
	}

    if applicationSiteCategory["tags"] != nil {
        tagsJson, ok := applicationSiteCategory["tags"].([]interface{})
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

	if v := applicationSiteCategory["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := applicationSiteCategory["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if applicationSiteCategory["groups"] != nil {
        groupsJson, ok := applicationSiteCategory["groups"].([]interface{})
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

	if v := applicationSiteCategory["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := applicationSiteCategory["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementApplicationSiteCategory(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    applicationSiteCategory := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        applicationSiteCategory["name"] = oldName
        applicationSiteCategory["new-name"] = newName
    } else {
        applicationSiteCategory["name"] = d.Get("name")
    }

    if ok := d.HasChange("description"); ok {
	       applicationSiteCategory["description"] = d.Get("description")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            applicationSiteCategory["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           applicationSiteCategory["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       applicationSiteCategory["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       applicationSiteCategory["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            applicationSiteCategory["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           applicationSiteCategory["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       applicationSiteCategory["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       applicationSiteCategory["ignore-errors"] = v.(bool)
    }

    log.Println("Update ApplicationSiteCategory - Map = ", applicationSiteCategory)

    updateApplicationSiteCategoryRes, err := client.ApiCall("set-application-site-category", applicationSiteCategory, client.GetSessionID(), true, false)
    if err != nil || !updateApplicationSiteCategoryRes.Success {
        if updateApplicationSiteCategoryRes.ErrorMsg != "" {
            return fmt.Errorf(updateApplicationSiteCategoryRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementApplicationSiteCategory(d, m)
}

func deleteManagementApplicationSiteCategory(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    applicationSiteCategoryPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete ApplicationSiteCategory")

    deleteApplicationSiteCategoryRes, err := client.ApiCall("delete-application-site-category", applicationSiteCategoryPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteApplicationSiteCategoryRes.Success {
        if deleteApplicationSiteCategoryRes.ErrorMsg != "" {
            return fmt.Errorf(deleteApplicationSiteCategoryRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

