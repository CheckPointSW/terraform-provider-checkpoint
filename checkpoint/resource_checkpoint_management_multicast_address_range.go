package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementMulticastAddressRange() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementMulticastAddressRange,
        Read:   readManagementMulticastAddressRange,
        Update: updateManagementMulticastAddressRange,
        Delete: deleteManagementMulticastAddressRange,
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
            "ipv6_address": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "IPv6 address.",
            },
            "ipv4_address_first": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "First IPv4 address in the range.",
            },
            "ipv6_address_first": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "First IPv6 address in the range.",
            },
            "ipv4_address_last": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Last IPv4 address in the range.",
            },
            "ipv6_address_last": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "Last IPv6 address in the range.",
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

func createManagementMulticastAddressRange(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    multicastAddressRange := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        multicastAddressRange["name"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_address"); ok {
        multicastAddressRange["ipv4-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv6_address"); ok {
        multicastAddressRange["ipv6-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_address_first"); ok {
        multicastAddressRange["ipv4-address-first"] = v.(string)
    }

    if v, ok := d.GetOk("ipv6_address_first"); ok {
        multicastAddressRange["ipv6-address-first"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_address_last"); ok {
        multicastAddressRange["ipv4-address-last"] = v.(string)
    }

    if v, ok := d.GetOk("ipv6_address_last"); ok {
        multicastAddressRange["ipv6-address-last"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        multicastAddressRange["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        multicastAddressRange["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        multicastAddressRange["comments"] = v.(string)
    }

    if v, ok := d.GetOk("groups"); ok {
        multicastAddressRange["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        multicastAddressRange["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        multicastAddressRange["ignore-errors"] = v.(bool)
    }

    log.Println("Create MulticastAddressRange - Map = ", multicastAddressRange)

    addMulticastAddressRangeRes, err := client.ApiCall("add-multicast-address-range", multicastAddressRange, client.GetSessionID(), true, false)
    if err != nil || !addMulticastAddressRangeRes.Success {
        if addMulticastAddressRangeRes.ErrorMsg != "" {
            return fmt.Errorf(addMulticastAddressRangeRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addMulticastAddressRangeRes.GetData()["uid"].(string))

    return readManagementMulticastAddressRange(d, m)
}

func readManagementMulticastAddressRange(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showMulticastAddressRangeRes, err := client.ApiCall("show-multicast-address-range", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showMulticastAddressRangeRes.Success {
		if objectNotFound(showMulticastAddressRangeRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showMulticastAddressRangeRes.ErrorMsg)
    }

    multicastAddressRange := showMulticastAddressRangeRes.GetData()

    log.Println("Read MulticastAddressRange - Show JSON = ", multicastAddressRange)

	if v := multicastAddressRange["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := multicastAddressRange["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := multicastAddressRange["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := multicastAddressRange["ipv4-address-first"]; v != nil {
		_ = d.Set("ipv4_address_first", v)
	}

	if v := multicastAddressRange["ipv6-address-first"]; v != nil {
		_ = d.Set("ipv6_address_first", v)
	}

	if v := multicastAddressRange["ipv4-address-last"]; v != nil {
		_ = d.Set("ipv4_address_last", v)
	}

	if v := multicastAddressRange["ipv6-address-last"]; v != nil {
		_ = d.Set("ipv6_address_last", v)
	}

    if multicastAddressRange["tags"] != nil {
        tagsJson, ok := multicastAddressRange["tags"].([]interface{})
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

	if v := multicastAddressRange["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := multicastAddressRange["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

    if multicastAddressRange["groups"] != nil {
        groupsJson, ok := multicastAddressRange["groups"].([]interface{})
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

	if v := multicastAddressRange["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := multicastAddressRange["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementMulticastAddressRange(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    multicastAddressRange := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        multicastAddressRange["name"] = oldName
        multicastAddressRange["new-name"] = newName
    } else {
        multicastAddressRange["name"] = d.Get("name")
    }

    if ok := d.HasChange("ipv4_address"); ok {
	       multicastAddressRange["ipv4-address"] = d.Get("ipv4_address")
    }

    if ok := d.HasChange("ipv6_address"); ok {
	       multicastAddressRange["ipv6-address"] = d.Get("ipv6_address")
    }

    if ok := d.HasChange("ipv4_address_first"); ok {
	       multicastAddressRange["ipv4-address-first"] = d.Get("ipv4_address_first")
    }

    if ok := d.HasChange("ipv6_address_first"); ok {
	       multicastAddressRange["ipv6-address-first"] = d.Get("ipv6_address_first")
    }

    if ok := d.HasChange("ipv4_address_last"); ok {
	       multicastAddressRange["ipv4-address-last"] = d.Get("ipv4_address_last")
    }

    if ok := d.HasChange("ipv6_address_last"); ok {
	       multicastAddressRange["ipv6-address-last"] = d.Get("ipv6_address_last")
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            multicastAddressRange["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           multicastAddressRange["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       multicastAddressRange["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       multicastAddressRange["comments"] = d.Get("comments")
    }

    if d.HasChange("groups") {
        if v, ok := d.GetOk("groups"); ok {
            multicastAddressRange["groups"] = v.(*schema.Set).List()
        } else {
            oldGroups, _ := d.GetChange("groups")
	           multicastAddressRange["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
        }
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       multicastAddressRange["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       multicastAddressRange["ignore-errors"] = v.(bool)
    }

    log.Println("Update MulticastAddressRange - Map = ", multicastAddressRange)

    updateMulticastAddressRangeRes, err := client.ApiCall("set-multicast-address-range", multicastAddressRange, client.GetSessionID(), true, false)
    if err != nil || !updateMulticastAddressRangeRes.Success {
        if updateMulticastAddressRangeRes.ErrorMsg != "" {
            return fmt.Errorf(updateMulticastAddressRangeRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementMulticastAddressRange(d, m)
}

func deleteManagementMulticastAddressRange(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    multicastAddressRangePayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete MulticastAddressRange")

    deleteMulticastAddressRangeRes, err := client.ApiCall("delete-multicast-address-range", multicastAddressRangePayload , client.GetSessionID(), true, false)
    if err != nil || !deleteMulticastAddressRangeRes.Success {
        if deleteMulticastAddressRangeRes.ErrorMsg != "" {
            return fmt.Errorf(deleteMulticastAddressRangeRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

