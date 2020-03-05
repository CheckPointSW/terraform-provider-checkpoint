package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementSecurityZone() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementSecurityZone,
        Read:   readManagementSecurityZone,
        Update: updateManagementSecurityZone,
        Delete: deleteManagementSecurityZone,
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

func createManagementSecurityZone(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    securityZone := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        securityZone["name"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        securityZone["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        securityZone["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        securityZone["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        securityZone["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        securityZone["ignore-errors"] = v.(bool)
    }

    log.Println("Create SecurityZone - Map = ", securityZone)

    addSecurityZoneRes, err := client.ApiCall("add-security-zone", securityZone, client.GetSessionID(), true, false)
    if err != nil || !addSecurityZoneRes.Success {
        if addSecurityZoneRes.ErrorMsg != "" {
            return fmt.Errorf(addSecurityZoneRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addSecurityZoneRes.GetData()["uid"].(string))

    return readManagementSecurityZone(d, m)
}

func readManagementSecurityZone(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showSecurityZoneRes, err := client.ApiCall("show-security-zone", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showSecurityZoneRes.Success {
		if objectNotFound(showSecurityZoneRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showSecurityZoneRes.ErrorMsg)
    }

    securityZone := showSecurityZoneRes.GetData()

    log.Println("Read SecurityZone - Show JSON = ", securityZone)

	if v := securityZone["name"]; v != nil {
		_ = d.Set("name", v)
	}

    if securityZone["tags"] != nil {
        tagsJson, ok := securityZone["tags"].([]interface{})
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

	if v := securityZone["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := securityZone["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := securityZone["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := securityZone["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementSecurityZone(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    securityZone := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        securityZone["name"] = oldName
        securityZone["new-name"] = newName
    } else {
        securityZone["name"] = d.Get("name")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            securityZone["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           securityZone["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       securityZone["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       securityZone["comments"] = d.Get("comments")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       securityZone["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       securityZone["ignore-errors"] = v.(bool)
    }

    log.Println("Update SecurityZone - Map = ", securityZone)

    updateSecurityZoneRes, err := client.ApiCall("set-security-zone", securityZone, client.GetSessionID(), true, false)
    if err != nil || !updateSecurityZoneRes.Success {
        if updateSecurityZoneRes.ErrorMsg != "" {
            return fmt.Errorf(updateSecurityZoneRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementSecurityZone(d, m)
}

func deleteManagementSecurityZone(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    securityZonePayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete SecurityZone")

    deleteSecurityZoneRes, err := client.ApiCall("delete-security-zone", securityZonePayload , client.GetSessionID(), true, false)
    if err != nil || !deleteSecurityZoneRes.Success {
        if deleteSecurityZoneRes.ErrorMsg != "" {
            return fmt.Errorf(deleteSecurityZoneRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

