package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementAccessLayer() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementAccessLayer,
        Read:   readManagementAccessLayer,
        Update: updateManagementAccessLayer,
        Delete: deleteManagementAccessLayer,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "applications_and_url_filtering": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Whether to enable Applications & URL Filtering blade on the layer.",
                Default:     false,
            },
            "content_awareness": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Whether to enable Content Awareness blade on the layer.",
                Default:     false,
            },
            "detect_using_x_forward_for": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Whether to use X-Forward-For HTTP header, which is added by the  proxy server to keep track of the original source IP.",
                Default:     false,
            },
            "firewall": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Whether to enable Firewall blade on the layer.",
                Default:     true,
            },
            "implicit_cleanup_action": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: "The default \"catch-all\" action for traffic that does not match any explicit or implied rules in the layer.",
                Default:     "drop",
            },
            "mobile_access": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Whether to enable Mobile Access blade on the layer.",
                Default:     false,
            },
            "shared": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Whether this layer is shared.",
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
            "add_default_rule": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Indicates whether to include a cleanup rule in the new layer.",
                Default:     true,
            },
        },
    }
}

func createManagementAccessLayer(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    accessLayer := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        accessLayer["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("applications_and_url_filtering"); ok {
        accessLayer["applications-and-url-filtering"] = v.(bool)
    }

    if v, ok := d.GetOkExists("content_awareness"); ok {
        accessLayer["content-awareness"] = v.(bool)
    }

    if v, ok := d.GetOkExists("detect_using_x_forward_for"); ok {
        accessLayer["detect-using-x-forward-for"] = v.(bool)
    }

    if v, ok := d.GetOkExists("firewall"); ok {
        accessLayer["firewall"] = v.(bool)
    }

    if v, ok := d.GetOk("implicit_cleanup_action"); ok {
        accessLayer["implicit-cleanup-action"] = v.(string)
    }

    if v, ok := d.GetOkExists("mobile_access"); ok {
        accessLayer["mobile-access"] = v.(bool)
    }

    if v, ok := d.GetOkExists("shared"); ok {
        accessLayer["shared"] = v.(bool)
    }

    if v, ok := d.GetOk("tags"); ok {
        accessLayer["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        accessLayer["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        accessLayer["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        accessLayer["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        accessLayer["ignore-errors"] = v.(bool)
    }

    if v, ok := d.GetOkExists("add_default_rule"); ok {
        accessLayer["add-default-rule"] = v.(bool)
    }

    log.Println("Create AccessLayer - Map = ", accessLayer)

    addAccessLayerRes, err := client.ApiCall("add-access-layer", accessLayer, client.GetSessionID(), true, false)
    if err != nil || !addAccessLayerRes.Success {
        if addAccessLayerRes.ErrorMsg != "" {
            return fmt.Errorf(addAccessLayerRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addAccessLayerRes.GetData()["uid"].(string))

    return readManagementAccessLayer(d, m)
}

func readManagementAccessLayer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showAccessLayerRes, err := client.ApiCall("show-access-layer", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showAccessLayerRes.Success {
		if objectNotFound(showAccessLayerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showAccessLayerRes.ErrorMsg)
    }

    accessLayer := showAccessLayerRes.GetData()

    log.Println("Read AccessLayer - Show JSON = ", accessLayer)

	if v := accessLayer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := accessLayer["applications-and-url-filtering"]; v != nil {
		_ = d.Set("applications_and_url_filtering", v)
	}

	if v := accessLayer["content-awareness"]; v != nil {
		_ = d.Set("content_awareness", v)
	}

	if v := accessLayer["detect-using-x-forward-for"]; v != nil {
		_ = d.Set("detect_using_x_forward_for", v)
	}

	if v := accessLayer["firewall"]; v != nil {
		_ = d.Set("firewall", v)
	}

	if v := accessLayer["implicit-cleanup-action"]; v != nil {
		_ = d.Set("implicit_cleanup_action", v)
	}

	if v := accessLayer["mobile-access"]; v != nil {
		_ = d.Set("mobile_access", v)
	}

	if v := accessLayer["shared"]; v != nil {
		_ = d.Set("shared", v)
	}

    if accessLayer["tags"] != nil {
        tagsJson, ok := accessLayer["tags"].([]interface{})
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

	if v := accessLayer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := accessLayer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := accessLayer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := accessLayer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	if v := accessLayer["add-default-rule"]; v != nil {
		_ = d.Set("add_default_rule", v)
	}

	return nil

}

func updateManagementAccessLayer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    accessLayer := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        accessLayer["name"] = oldName
        accessLayer["new-name"] = newName
    } else {
        accessLayer["name"] = d.Get("name")
    }

    if v, ok := d.GetOkExists("applications_and_url_filtering"); ok {
	       accessLayer["applications-and-url-filtering"] = v.(bool)
    }

    if v, ok := d.GetOkExists("content_awareness"); ok {
	       accessLayer["content-awareness"] = v.(bool)
    }

    if v, ok := d.GetOkExists("detect_using_x_forward_for"); ok {
	       accessLayer["detect-using-x-forward-for"] = v.(bool)
    }

    if v, ok := d.GetOkExists("firewall"); ok {
	       accessLayer["firewall"] = v.(bool)
    }

    if ok := d.HasChange("implicit_cleanup_action"); ok {
	       accessLayer["implicit-cleanup-action"] = d.Get("implicit_cleanup_action")
    }

    if v, ok := d.GetOkExists("mobile_access"); ok {
	       accessLayer["mobile-access"] = v.(bool)
    }

    if v, ok := d.GetOkExists("shared"); ok {
	       accessLayer["shared"] = v.(bool)
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            accessLayer["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           accessLayer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       accessLayer["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       accessLayer["comments"] = d.Get("comments")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       accessLayer["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       accessLayer["ignore-errors"] = v.(bool)
    }

    if v, ok := d.GetOkExists("add_default_rule"); ok {
	       accessLayer["add-default-rule"] = v.(bool)
    }

    log.Println("Update AccessLayer - Map = ", accessLayer)

    updateAccessLayerRes, err := client.ApiCall("set-access-layer", accessLayer, client.GetSessionID(), true, false)
    if err != nil || !updateAccessLayerRes.Success {
        if updateAccessLayerRes.ErrorMsg != "" {
            return fmt.Errorf(updateAccessLayerRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementAccessLayer(d, m)
}

func deleteManagementAccessLayer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    accessLayerPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete AccessLayer")

    deleteAccessLayerRes, err := client.ApiCall("delete-access-layer", accessLayerPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteAccessLayerRes.Success {
        if deleteAccessLayerRes.ErrorMsg != "" {
            return fmt.Errorf(deleteAccessLayerRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

