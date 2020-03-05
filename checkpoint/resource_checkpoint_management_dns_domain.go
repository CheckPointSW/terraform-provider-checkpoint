package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementDnsDomain() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementDnsDomain,
        Read:   readManagementDnsDomain,
        Update: updateManagementDnsDomain,
        Delete: deleteManagementDnsDomain,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "is_sub_domain": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Whether to match sub-domains in addition to the domain itself.",
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

func createManagementDnsDomain(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    dnsDomain := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        dnsDomain["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("is_sub_domain"); ok {
        dnsDomain["is-sub-domain"] = v.(bool)
    }

    if v, ok := d.GetOk("tags"); ok {
        dnsDomain["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        dnsDomain["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        dnsDomain["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        dnsDomain["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        dnsDomain["ignore-errors"] = v.(bool)
    }

    log.Println("Create DnsDomain - Map = ", dnsDomain)

    addDnsDomainRes, err := client.ApiCall("add-dns-domain", dnsDomain, client.GetSessionID(), true, false)
    if err != nil || !addDnsDomainRes.Success {
        if addDnsDomainRes.ErrorMsg != "" {
            return fmt.Errorf(addDnsDomainRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addDnsDomainRes.GetData()["uid"].(string))

    return readManagementDnsDomain(d, m)
}

func readManagementDnsDomain(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showDnsDomainRes, err := client.ApiCall("show-dns-domain", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showDnsDomainRes.Success {
		if objectNotFound(showDnsDomainRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showDnsDomainRes.ErrorMsg)
    }

    dnsDomain := showDnsDomainRes.GetData()

    log.Println("Read DnsDomain - Show JSON = ", dnsDomain)

	if v := dnsDomain["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dnsDomain["is-sub-domain"]; v != nil {
		_ = d.Set("is_sub_domain", v)
	}

    if dnsDomain["tags"] != nil {
        tagsJson, ok := dnsDomain["tags"].([]interface{})
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

	if v := dnsDomain["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dnsDomain["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dnsDomain["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dnsDomain["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDnsDomain(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    dnsDomain := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        dnsDomain["name"] = oldName
        dnsDomain["new-name"] = newName
    } else {
        dnsDomain["name"] = d.Get("name")
    }

    if v, ok := d.GetOkExists("is_sub_domain"); ok {
	       dnsDomain["is-sub-domain"] = v.(bool)
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            dnsDomain["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           dnsDomain["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       dnsDomain["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       dnsDomain["comments"] = d.Get("comments")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       dnsDomain["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       dnsDomain["ignore-errors"] = v.(bool)
    }

    log.Println("Update DnsDomain - Map = ", dnsDomain)

    updateDnsDomainRes, err := client.ApiCall("set-dns-domain", dnsDomain, client.GetSessionID(), true, false)
    if err != nil || !updateDnsDomainRes.Success {
        if updateDnsDomainRes.ErrorMsg != "" {
            return fmt.Errorf(updateDnsDomainRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementDnsDomain(d, m)
}

func deleteManagementDnsDomain(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    dnsDomainPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete DnsDomain")

    deleteDnsDomainRes, err := client.ApiCall("delete-dns-domain", dnsDomainPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteDnsDomainRes.Success {
        if deleteDnsDomainRes.ErrorMsg != "" {
            return fmt.Errorf(deleteDnsDomainRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

