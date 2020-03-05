package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementServiceRpc() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementServiceRpc,
        Read:   readManagementServiceRpc,
        Update: updateManagementServiceRpc,
        Delete: deleteManagementServiceRpc,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "keep_connections_open_after_policy_installation": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
            },
            "program_number": {
                Type:        schema.TypeInt,
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

func createManagementServiceRpc(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    serviceRpc := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        serviceRpc["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
        serviceRpc["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if v, ok := d.GetOk("program_number"); ok {
        serviceRpc["program-number"] = v.(int)
    }

    if v, ok := d.GetOk("tags"); ok {
        serviceRpc["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        serviceRpc["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        serviceRpc["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        serviceRpc["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        serviceRpc["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        serviceRpc["ignore-errors"] = v.(bool)
    }

    log.Println("Create ServiceRpc - Map = ", serviceRpc)

    addServiceRpcRes, err := client.ApiCall("add-service-rpc", serviceRpc, client.GetSessionID(), true, false)
    if err != nil || !addServiceRpcRes.Success {
        if addServiceRpcRes.ErrorMsg != "" {
            return fmt.Errorf(addServiceRpcRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addServiceRpcRes.GetData()["uid"].(string))

    return readManagementServiceRpc(d, m)
}

func readManagementServiceRpc(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showServiceRpcRes, err := client.ApiCall("show-service-rpc", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showServiceRpcRes.Success {
		if objectNotFound(showServiceRpcRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showServiceRpcRes.ErrorMsg)
    }

    serviceRpc := showServiceRpcRes.GetData()

    log.Println("Read ServiceRpc - Show JSON = ", serviceRpc)

	if v := serviceRpc["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceRpc["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if v := serviceRpc["program-number"]; v != nil {
		_ = d.Set("program_number", v)
	}

    if serviceRpc["tags"] != nil {
        tagsJson, ok := serviceRpc["tags"].([]interface{})
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

	if v := serviceRpc["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceRpc["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if serviceRpc["groups"] != nil {
        groupsJson, ok := serviceRpc["groups"].([]interface{})
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

	if v := serviceRpc["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := serviceRpc["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementServiceRpc(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    serviceRpc := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        serviceRpc["name"] = oldName
        serviceRpc["new-name"] = newName
    } else {
        serviceRpc["name"] = d.Get("name")
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
	       serviceRpc["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if ok := d.HasChange("program_number"); ok {
	       serviceRpc["program-number"] = d.Get("program_number")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            serviceRpc["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           serviceRpc["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       serviceRpc["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       serviceRpc["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            serviceRpc["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           serviceRpc["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       serviceRpc["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       serviceRpc["ignore-errors"] = v.(bool)
    }

    log.Println("Update ServiceRpc - Map = ", serviceRpc)

    updateServiceRpcRes, err := client.ApiCall("set-service-rpc", serviceRpc, client.GetSessionID(), true, false)
    if err != nil || !updateServiceRpcRes.Success {
        if updateServiceRpcRes.ErrorMsg != "" {
            return fmt.Errorf(updateServiceRpcRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementServiceRpc(d, m)
}

func deleteManagementServiceRpc(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    serviceRpcPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete ServiceRpc")

    deleteServiceRpcRes, err := client.ApiCall("delete-service-rpc", serviceRpcPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteServiceRpcRes.Success {
        if deleteServiceRpcRes.ErrorMsg != "" {
            return fmt.Errorf(deleteServiceRpcRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

