package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementHttpsLayer() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementHttpsLayer,
        Read:   readManagementHttpsLayer,
        Update: updateManagementHttpsLayer,
        Delete: deleteManagementHttpsLayer,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "shared": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Define the Layer as Shared (TRUE/FALSE).",
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
        },
    }
}

func createManagementHttpsLayer(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    httpsLayer := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        httpsLayer["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("shared"); ok {
        httpsLayer["shared"] = v.(bool)
    }

    if v, ok := d.GetOk("tags"); ok {
        httpsLayer["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        httpsLayer["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        httpsLayer["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        httpsLayer["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        httpsLayer["ignore-errors"] = v.(bool)
    }

    log.Println("Create HttpsLayer - Map = ", httpsLayer)

    addHttpsLayerRes, err := client.ApiCall("add-https-layer", httpsLayer, client.GetSessionID(), true, false)
    if err != nil || !addHttpsLayerRes.Success {
        if addHttpsLayerRes.ErrorMsg != "" {
            return fmt.Errorf(addHttpsLayerRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addHttpsLayerRes.GetData()["uid"].(string))

    return readManagementHttpsLayer(d, m)
}

func readManagementHttpsLayer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
    }

    showHttpsLayerRes, err := client.ApiCall("show-https-layer", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showHttpsLayerRes.Success {
		if objectNotFound(showHttpsLayerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showHttpsLayerRes.ErrorMsg)
    }

    httpsLayer := showHttpsLayerRes.GetData()

    log.Println("Read HttpsLayer - Show JSON = ", httpsLayer)

	if v := httpsLayer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := httpsLayer["shared"]; v != nil {
		_ = d.Set("shared", v)
	}

    if httpsLayer["tags"] != nil {
        tagsJson, ok := httpsLayer["tags"].([]interface{})
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

	if v := httpsLayer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := httpsLayer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := httpsLayer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := httpsLayer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementHttpsLayer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    httpsLayer := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        httpsLayer["name"] = oldName
        httpsLayer["new-name"] = newName
    } else {
        httpsLayer["name"] = d.Get("name")
    }

    if v, ok := d.GetOkExists("shared"); ok {
	       httpsLayer["shared"] = v.(bool)
    }

    if d.HasChange("tags") {
        if v, ok := d.GetOk("tags"); ok {
            httpsLayer["tags"] = v.(*schema.Set).List()
        } else {
            oldTags, _ := d.GetChange("tags")
	           httpsLayer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
        }
    }

    if ok := d.HasChange("color"); ok {
	       httpsLayer["color"] = d.Get("color")
    }

    if ok := d.HasChange("comments"); ok {
	       httpsLayer["comments"] = d.Get("comments")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       httpsLayer["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       httpsLayer["ignore-errors"] = v.(bool)
    }

    log.Println("Update HttpsLayer - Map = ", httpsLayer)

    updateHttpsLayerRes, err := client.ApiCall("set-https-layer", httpsLayer, client.GetSessionID(), true, false)
    if err != nil || !updateHttpsLayerRes.Success {
        if updateHttpsLayerRes.ErrorMsg != "" {
            return fmt.Errorf(updateHttpsLayerRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementHttpsLayer(d, m)
}

func deleteManagementHttpsLayer(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    httpsLayerPayload := map[string]interface{}{
        "uid": d.Id(),
    }

    log.Println("Delete HttpsLayer")

    deleteHttpsLayerRes, err := client.ApiCall("delete-https-layer", httpsLayerPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteHttpsLayerRes.Success {
        if deleteHttpsLayerRes.ErrorMsg != "" {
            return fmt.Errorf(deleteHttpsLayerRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

