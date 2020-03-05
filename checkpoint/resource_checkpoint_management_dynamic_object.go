package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementDynamicObject() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementDynamicObject,
        Read:   readManagementDynamicObject,
        Update: updateManagementDynamicObject,
        Delete: deleteManagementDynamicObject,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
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

func createManagementDynamicObject(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    dynamicObject := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        dynamicObject["name"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        dynamicObject["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        dynamicObject["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        dynamicObject["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        dynamicObject["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        dynamicObject["ignore-errors"] = v.(bool)
    }

    log.Println("Create DynamicObject - Map = ", dynamicObject)

    addDynamicObjectRes, err := client.ApiCall("add-dynamic-object", dynamicObject, client.GetSessionID(), true, false)
    if err != nil || !addDynamicObjectRes.Success {
        if addDynamicObjectRes.ErrorMsg != "" {
            return fmt.Errorf(addDynamicObjectRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addDynamicObjectRes.GetData()["uid"].(string))

    return readManagementDynamicObject(d, m)
}

func readManagementDynamicObject(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showDynamicObjectRes, err := client.ApiCall("show-dynamic-object", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showDynamicObjectRes.Success {
		if objectNotFound(showDynamicObjectRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showDynamicObjectRes.ErrorMsg)
    }

    dynamicObject := showDynamicObjectRes.GetData()

    log.Println("Read DynamicObject - Show JSON = ", dynamicObject)

	if v := dynamicObject["name"]; v != nil {
		_ = d.Set("name", v)
	}

    if dynamicObject["tags"] != nil {
        tagsJson, ok := dynamicObject["tags"].([]interface{})
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

	if v := dynamicObject["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dynamicObject["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dynamicObject["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dynamicObject["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDynamicObject(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    dynamicObject := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        dynamicObject["name"] = oldName
        dynamicObject["new-name"] = newName
    } else {
        dynamicObject["name"] = d.Get("name")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            dynamicObject["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           dynamicObject["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       dynamicObject["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       dynamicObject["comments"] = d.Get("comments")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       dynamicObject["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       dynamicObject["ignore-errors"] = v.(bool)
    }

    log.Println("Update DynamicObject - Map = ", dynamicObject)

    updateDynamicObjectRes, err := client.ApiCall("set-dynamic-object", dynamicObject, client.GetSessionID(), true, false)
    if err != nil || !updateDynamicObjectRes.Success {
        if updateDynamicObjectRes.ErrorMsg != "" {
            return fmt.Errorf(updateDynamicObjectRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementDynamicObject(d, m)
}

func deleteManagementDynamicObject(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    dynamicObjectPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete DynamicObject")

    deleteDynamicObjectRes, err := client.ApiCall("delete-dynamic-object", dynamicObjectPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteDynamicObjectRes.Success {
        if deleteDynamicObjectRes.ErrorMsg != "" {
            return fmt.Errorf(deleteDynamicObjectRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

