package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementHttpsSection() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementHttpsSection,
        Read:   readManagementHttpsSection,
        Update: updateManagementHttpsSection,
        Delete: deleteManagementHttpsSection,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "layer": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Layer that holds the Object. Identified by the Name or UID.",
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
       			"position": &schema.Schema{
       				Type:        schema.TypeMap,
       				Required:    true,
       				Description: "Position in the rulebase.",
       				Elem: &schema.Resource{
       					Schema: map[string]*schema.Schema{
       						"top": {
       							Type:        schema.TypeString,
       							Optional:    true,
       							Description: "N/A",
       						},
       						"above": {
       							Type:        schema.TypeString,
       							Optional:    true,
       							Description: "N/A",
       						},
       						"below": {
       							Type:        schema.TypeString,
       							Optional:    true,
       							Description: "N/A",
       						},
       						"bottom": {
       							Type:        schema.TypeString,
       							Optional:    true,
       							Description: "N/A",
       						},
       					},
       				},
       			},
        },
    }
}

func createManagementHttpsSection(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    httpsSection := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        httpsSection["name"] = v.(string)
    }

    if v, ok := d.GetOk("layer"); ok {
        httpsSection["layer"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        httpsSection["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        httpsSection["ignore-errors"] = v.(bool)
    }

    if _, ok := d.GetOk("position"); ok {
		if _, ok := d.GetOk("position.top"); ok {
            httpsSection["position"] = "top"
        }
        if v, ok := d.GetOk("position.above"); ok {
            httpsSection["position"] = map[string]interface{}{"above": v.(string)}
        }
        if v, ok := d.GetOk("position.bottom"); ok {
            httpsSection["position"] = map[string]interface{}{"bottom": v.(string)}
        }
        if _, ok := d.GetOk("position.bottom"); ok {
            httpsSection["position"] = "bottom"
        }
    }
    log.Println("Create HttpsSection - Map = ", httpsSection)

    addHttpsSectionRes, err := client.ApiCall("add-https-section", httpsSection, client.GetSessionID(), true, false)
    if err != nil || !addHttpsSectionRes.Success {
        if addHttpsSectionRes.ErrorMsg != "" {
            return fmt.Errorf(addHttpsSectionRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addHttpsSectionRes.GetData()["uid"].(string))

    return readManagementHttpsSection(d, m)
}

func readManagementHttpsSection(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
		"layer": d.Get("layer"),
    }

    showHttpsSectionRes, err := client.ApiCall("show-https-section", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showHttpsSectionRes.Success {
		if objectNotFound(showHttpsSectionRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showHttpsSectionRes.ErrorMsg)
    }

    httpsSection := showHttpsSectionRes.GetData()

    log.Println("Read HttpsSection - Show JSON = ", httpsSection)

	if v := httpsSection["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := httpsSection["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := httpsSection["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementHttpsSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    httpsSection := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        httpsSection["name"] = oldName
        httpsSection["new-name"] = newName
    } else {
        httpsSection["name"] = d.Get("name")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       httpsSection["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       httpsSection["ignore-errors"] = v.(bool)
    }

    log.Println("Update HttpsSection - Map = ", httpsSection)

    updateHttpsSectionRes, err := client.ApiCall("set-https-section", httpsSection, client.GetSessionID(), true, false)
    if err != nil || !updateHttpsSectionRes.Success {
        if updateHttpsSectionRes.ErrorMsg != "" {
            return fmt.Errorf(updateHttpsSectionRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementHttpsSection(d, m)
}

func deleteManagementHttpsSection(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    httpsSectionPayload := map[string]interface{}{
        "uid": d.Id(),
        "layer": d.Get("layer"),
    }

    log.Println("Delete HttpsSection")

    deleteHttpsSectionRes, err := client.ApiCall("delete-https-section", httpsSectionPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteHttpsSectionRes.Success {
        if deleteHttpsSectionRes.ErrorMsg != "" {
            return fmt.Errorf(deleteHttpsSectionRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

