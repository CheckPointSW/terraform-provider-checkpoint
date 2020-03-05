package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementWildcard() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementWildcard,
        Read:   readManagementWildcard,
        Update: updateManagementWildcard,
        Delete: deleteManagementWildcard,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "ipv4_address": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "IPv4 address.",
            },
            "ipv4_mask_wildcard": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "IPv4 mask wildcard.",
            },
            "ipv6_address": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "IPv6 address.",
            },
            "ipv6_mask_wildcard": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "IPv6 mask wildcard.",
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

func createManagementWildcard(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    wildcard := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        wildcard["name"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_address"); ok {
        wildcard["ipv4-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_mask_wildcard"); ok {
        wildcard["ipv4-mask-wildcard"] = v.(string)
    }

    if v, ok := d.GetOk("ipv6_address"); ok {
        wildcard["ipv6-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv6_mask_wildcard"); ok {
        wildcard["ipv6-mask-wildcard"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        wildcard["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        wildcard["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        wildcard["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        wildcard["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        wildcard["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        wildcard["ignore-errors"] = v.(bool)
    }

    log.Println("Create Wildcard - Map = ", wildcard)

    addWildcardRes, err := client.ApiCall("add-wildcard", wildcard, client.GetSessionID(), true, false)
    if err != nil || !addWildcardRes.Success {
        if addWildcardRes.ErrorMsg != "" {
            return fmt.Errorf(addWildcardRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addWildcardRes.GetData()["uid"].(string))

    return readManagementWildcard(d, m)
}

func readManagementWildcard(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showWildcardRes, err := client.ApiCall("show-wildcard", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showWildcardRes.Success {
		if objectNotFound(showWildcardRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showWildcardRes.ErrorMsg)
    }

    wildcard := showWildcardRes.GetData()

    log.Println("Read Wildcard - Show JSON = ", wildcard)

	if v := wildcard["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := wildcard["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := wildcard["ipv4-mask-wildcard"]; v != nil {
		_ = d.Set("ipv4_mask_wildcard", v)
	}

	if v := wildcard["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := wildcard["ipv6-mask-wildcard"]; v != nil {
		_ = d.Set("ipv6_mask_wildcard", v)
	}

    if wildcard["tags"] != nil {
        tagsJson, ok := wildcard["tags"].([]interface{})
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

	if v := wildcard["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := wildcard["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if wildcard["groups"] != nil {
        groupsJson, ok := wildcard["groups"].([]interface{})
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

	if v := wildcard["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := wildcard["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementWildcard(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    wildcard := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        wildcard["name"] = oldName
        wildcard["new-name"] = newName
    } else {
        wildcard["name"] = d.Get("name")
    }

    if ok := d.HasChange("ipv4_address"); ok {
	       wildcard["ipv4-address"] = d.Get("ipv4_address")
    }

    if ok := d.HasChange("ipv4_mask_wildcard"); ok {
	       wildcard["ipv4-mask-wildcard"] = d.Get("ipv4_mask_wildcard")
    }

    if ok := d.HasChange("ipv6_address"); ok {
	       wildcard["ipv6-address"] = d.Get("ipv6_address")
    }

    if ok := d.HasChange("ipv6_mask_wildcard"); ok {
	       wildcard["ipv6-mask-wildcard"] = d.Get("ipv6_mask_wildcard")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            wildcard["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           wildcard["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       wildcard["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       wildcard["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            wildcard["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           wildcard["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       wildcard["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       wildcard["ignore-errors"] = v.(bool)
    }

    log.Println("Update Wildcard - Map = ", wildcard)

    updateWildcardRes, err := client.ApiCall("set-wildcard", wildcard, client.GetSessionID(), true, false)
    if err != nil || !updateWildcardRes.Success {
        if updateWildcardRes.ErrorMsg != "" {
            return fmt.Errorf(updateWildcardRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementWildcard(d, m)
}

func deleteManagementWildcard(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    wildcardPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete Wildcard")

    deleteWildcardRes, err := client.ApiCall("delete-wildcard", wildcardPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteWildcardRes.Success {
        if deleteWildcardRes.ErrorMsg != "" {
            return fmt.Errorf(deleteWildcardRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

