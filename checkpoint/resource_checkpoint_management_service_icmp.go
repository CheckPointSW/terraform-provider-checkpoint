package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementServiceIcmp() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementServiceIcmp,
        Read:   readManagementServiceIcmp,
        Update: updateManagementServiceIcmp,
        Delete: deleteManagementServiceIcmp,
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

func createManagementServiceIcmp(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    serviceIcmp := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        serviceIcmp["name"] = v.(string)
    }

    if v, ok := d.GetOk("icmp_code"); ok {
        serviceIcmp["icmp-code"] = v.(int)
    }

    if v, ok := d.GetOk("icmp_type"); ok {
        serviceIcmp["icmp-type"] = v.(int)
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
        serviceIcmp["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if v, ok := d.GetOk("tags"); ok {
        serviceIcmp["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        serviceIcmp["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        serviceIcmp["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        serviceIcmp["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        serviceIcmp["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        serviceIcmp["ignore-errors"] = v.(bool)
    }

    log.Println("Create ServiceIcmp - Map = ", serviceIcmp)

    addServiceIcmpRes, err := client.ApiCall("add-service-icmp", serviceIcmp, client.GetSessionID(), true, false)
    if err != nil || !addServiceIcmpRes.Success {
        if addServiceIcmpRes.ErrorMsg != "" {
            return fmt.Errorf(addServiceIcmpRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addServiceIcmpRes.GetData()["uid"].(string))

    return readManagementServiceIcmp(d, m)
}

func readManagementServiceIcmp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showServiceIcmpRes, err := client.ApiCall("show-service-icmp", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showServiceIcmpRes.Success {
		if objectNotFound(showServiceIcmpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showServiceIcmpRes.ErrorMsg)
    }

    serviceIcmp := showServiceIcmpRes.GetData()

    log.Println("Read ServiceIcmp - Show JSON = ", serviceIcmp)

	if v := serviceIcmp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceIcmp["icmp-code"]; v != nil {
		_ = d.Set("icmp_code", v)
	}

	if v := serviceIcmp["icmp-type"]; v != nil {
		_ = d.Set("icmp_type", v)
	}

	if v := serviceIcmp["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

    if serviceIcmp["tags"] != nil {
        tagsJson, ok := serviceIcmp["tags"].([]interface{})
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

	if v := serviceIcmp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceIcmp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if serviceIcmp["groups"] != nil {
        groupsJson, ok := serviceIcmp["groups"].([]interface{})
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

	if v := serviceIcmp["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := serviceIcmp["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementServiceIcmp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    serviceIcmp := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        serviceIcmp["name"] = oldName
        serviceIcmp["new-name"] = newName
    } else {
        serviceIcmp["name"] = d.Get("name")
    }

    if ok := d.HasChange("icmp_code"); ok {
	       serviceIcmp["icmp-code"] = d.Get("icmp_code")
    }

    if ok := d.HasChange("icmp_type"); ok {
	       serviceIcmp["icmp-type"] = d.Get("icmp_type")
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
	       serviceIcmp["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            serviceIcmp["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           serviceIcmp["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       serviceIcmp["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       serviceIcmp["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            serviceIcmp["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           serviceIcmp["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       serviceIcmp["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       serviceIcmp["ignore-errors"] = v.(bool)
    }

    log.Println("Update ServiceIcmp - Map = ", serviceIcmp)

    updateServiceIcmpRes, err := client.ApiCall("set-service-icmp", serviceIcmp, client.GetSessionID(), true, false)
    if err != nil || !updateServiceIcmpRes.Success {
        if updateServiceIcmpRes.ErrorMsg != "" {
            return fmt.Errorf(updateServiceIcmpRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementServiceIcmp(d, m)
}

func deleteManagementServiceIcmp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    serviceIcmpPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete ServiceIcmp")

    deleteServiceIcmpRes, err := client.ApiCall("delete-service-icmp", serviceIcmpPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteServiceIcmpRes.Success {
        if deleteServiceIcmpRes.ErrorMsg != "" {
            return fmt.Errorf(deleteServiceIcmpRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

