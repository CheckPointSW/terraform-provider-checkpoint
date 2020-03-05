package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementServiceDceRpc() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementServiceDceRpc,
        Read:   readManagementServiceDceRpc,
        Update: updateManagementServiceDceRpc,
        Delete: deleteManagementServiceDceRpc,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "interface_uuid": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Network interface UUID.",
            },
            "keep_connections_open_after_policy_installation": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
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

func createManagementServiceDceRpc(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    serviceDceRpc := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        serviceDceRpc["name"] = v.(string)
    }

    if v, ok := d.GetOk("interface_uuid"); ok {
        serviceDceRpc["interface-uuid"] = v.(string)
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
        serviceDceRpc["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if v, ok := d.GetOk("tags"); ok {
        serviceDceRpc["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        serviceDceRpc["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        serviceDceRpc["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        serviceDceRpc["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        serviceDceRpc["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        serviceDceRpc["ignore-errors"] = v.(bool)
    }

    log.Println("Create ServiceDceRpc - Map = ", serviceDceRpc)

    addServiceDceRpcRes, err := client.ApiCall("add-service-dce-rpc", serviceDceRpc, client.GetSessionID(), true, false)
    if err != nil || !addServiceDceRpcRes.Success {
        if addServiceDceRpcRes.ErrorMsg != "" {
            return fmt.Errorf(addServiceDceRpcRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addServiceDceRpcRes.GetData()["uid"].(string))

    return readManagementServiceDceRpc(d, m)
}

func readManagementServiceDceRpc(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showServiceDceRpcRes, err := client.ApiCall("show-service-dce-rpc", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showServiceDceRpcRes.Success {
		if objectNotFound(showServiceDceRpcRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showServiceDceRpcRes.ErrorMsg)
    }

    serviceDceRpc := showServiceDceRpcRes.GetData()

    log.Println("Read ServiceDceRpc - Show JSON = ", serviceDceRpc)

	if v := serviceDceRpc["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceDceRpc["interface-uuid"]; v != nil {
		_ = d.Set("interface_uuid", v)
	}

	if v := serviceDceRpc["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

    if serviceDceRpc["tags"] != nil {
        tagsJson, ok := serviceDceRpc["tags"].([]interface{})
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

	if v := serviceDceRpc["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceDceRpc["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if serviceDceRpc["groups"] != nil {
        groupsJson, ok := serviceDceRpc["groups"].([]interface{})
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

	if v := serviceDceRpc["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := serviceDceRpc["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementServiceDceRpc(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    serviceDceRpc := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        serviceDceRpc["name"] = oldName
        serviceDceRpc["new-name"] = newName
    } else {
        serviceDceRpc["name"] = d.Get("name")
    }

    if ok := d.HasChange("interface_uuid"); ok {
	       serviceDceRpc["interface-uuid"] = d.Get("interface_uuid")
    }

    if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
	       serviceDceRpc["keep-connections-open-after-policy-installation"] = v.(bool)
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            serviceDceRpc["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           serviceDceRpc["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       serviceDceRpc["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       serviceDceRpc["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            serviceDceRpc["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           serviceDceRpc["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       serviceDceRpc["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       serviceDceRpc["ignore-errors"] = v.(bool)
    }

    log.Println("Update ServiceDceRpc - Map = ", serviceDceRpc)

    updateServiceDceRpcRes, err := client.ApiCall("set-service-dce-rpc", serviceDceRpc, client.GetSessionID(), true, false)
    if err != nil || !updateServiceDceRpcRes.Success {
        if updateServiceDceRpcRes.ErrorMsg != "" {
            return fmt.Errorf(updateServiceDceRpcRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementServiceDceRpc(d, m)
}

func deleteManagementServiceDceRpc(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    serviceDceRpcPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete ServiceDceRpc")

    deleteServiceDceRpcRes, err := client.ApiCall("delete-service-dce-rpc", serviceDceRpcPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteServiceDceRpcRes.Success {
        if deleteServiceDceRpcRes.ErrorMsg != "" {
            return fmt.Errorf(deleteServiceDceRpcRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

