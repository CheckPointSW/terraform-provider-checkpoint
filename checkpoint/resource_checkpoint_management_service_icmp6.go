package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementServiceIcmp6() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementServiceIcmp6,
        Read:   readManagementServiceIcmp6,
        Update: updateManagementServiceIcmp6,
        Delete: deleteManagementServiceIcmp6,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "icmp_code": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: "As listed in: <a href=\"http://www.iana.org/assignments/icmp-parameters\" target=\"_blank\">RFC 792</a>.",
                Default:     0,
            },
            "icmp_type": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: "As listed in: <a href=\"http://www.iana.org/assignments/icmp-parameters\" target=\"_blank\">RFC 792</a>.",
                Default:     0,
            },
            "keep_connections_open_after_policy_installation": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
                Default:     false,
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

func createManagementServiceIcmp6(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    serviceIcmp6 := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        serviceIcmp6["name"] = v.(string)
    }

    if v, ok := d.GetOk("icmp_code"); ok {
        serviceIcmp6["icmp-code"] = v.(int)
    }

    if v, ok := d.GetOk("icmp_type"); ok {
        serviceIcmp6["icmp-type"] = v.(int)
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
        serviceIcmp6["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if v, ok := d.GetOk("tags"); ok {
        serviceIcmp6["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        serviceIcmp6["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        serviceIcmp6["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        serviceIcmp6["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        serviceIcmp6["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        serviceIcmp6["ignore-errors"] = v.(bool)
    }

    log.Println("Create ServiceIcmp6 - Map = ", serviceIcmp6)

    addServiceIcmp6Res, err := client.ApiCall("add-service-icmp6", serviceIcmp6, client.GetSessionID(), true, false)
    if err != nil || !addServiceIcmp6Res.Success {
        if addServiceIcmp6Res.ErrorMsg != "" {
            return fmt.Errorf(addServiceIcmp6Res.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addServiceIcmp6Res.GetData()["uid"].(string))

    return readManagementServiceIcmp6(d, m)
}

func readManagementServiceIcmp6(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showServiceIcmp6Res, err := client.ApiCall("show-service-icmp6", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showServiceIcmp6Res.Success {
		if objectNotFound(showServiceIcmp6Res.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showServiceIcmp6Res.ErrorMsg)
    }

    serviceIcmp6 := showServiceIcmp6Res.GetData()

    log.Println("Read ServiceIcmp6 - Show JSON = ", serviceIcmp6)

	if v := serviceIcmp6["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceIcmp6["icmp-code"]; v != nil {
		_ = d.Set("icmp_code", v)
	}

	if v := serviceIcmp6["icmp-type"]; v != nil {
		_ = d.Set("icmp_type", v)
	}

	if v := serviceIcmp6["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

    if serviceIcmp6["tags"] != nil {
        tagsJson, ok := serviceIcmp6["tags"].([]interface{})
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

	if v := serviceIcmp6["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceIcmp6["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if serviceIcmp6["groups"] != nil {
        groupsJson, ok := serviceIcmp6["groups"].([]interface{})
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

	if v := serviceIcmp6["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := serviceIcmp6["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementServiceIcmp6(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    serviceIcmp6 := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        serviceIcmp6["name"] = oldName
        serviceIcmp6["new-name"] = newName
    } else {
        serviceIcmp6["name"] = d.Get("name")
    }

    if ok := d.HasChange("icmp_code"); ok {
	       serviceIcmp6["icmp-code"] = d.Get("icmp_code")
    }

    if ok := d.HasChange("icmp_type"); ok {
	       serviceIcmp6["icmp-type"] = d.Get("icmp_type")
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
	       serviceIcmp6["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            serviceIcmp6["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           serviceIcmp6["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       serviceIcmp6["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       serviceIcmp6["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            serviceIcmp6["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           serviceIcmp6["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       serviceIcmp6["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       serviceIcmp6["ignore-errors"] = v.(bool)
    }

    log.Println("Update ServiceIcmp6 - Map = ", serviceIcmp6)

    updateServiceIcmp6Res, err := client.ApiCall("set-service-icmp6", serviceIcmp6, client.GetSessionID(), true, false)
    if err != nil || !updateServiceIcmp6Res.Success {
        if updateServiceIcmp6Res.ErrorMsg != "" {
            return fmt.Errorf(updateServiceIcmp6Res.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementServiceIcmp6(d, m)
}

func deleteManagementServiceIcmp6(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    serviceIcmp6Payload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete ServiceIcmp6")

    deleteServiceIcmp6Res, err := client.ApiCall("delete-service-icmp6", serviceIcmp6Payload , client.GetSessionID(), true, false)
    if err != nil || !deleteServiceIcmp6Res.Success {
        if deleteServiceIcmp6Res.ErrorMsg != "" {
            return fmt.Errorf(deleteServiceIcmp6Res.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

