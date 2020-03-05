package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	
	
)

func resourceManagementAccessSection() *schema.Resource {   
    return &schema.Resource{
        Create: createManagementAccessSection,
        Read:   readManagementAccessSection,
        Update: updateManagementAccessSection,
        Delete: deleteManagementAccessSection,
        Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Object name.",
            },
            "layer": {
                Type:        schema.TypeString,
                Required:    true,
                Description: "Layer that the rule belongs to identified by the name or UID.",
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

func createManagementAccessSection(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)

    accessSection := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        accessSection["name"] = v.(string)
    }

    if v, ok := d.GetOk("layer"); ok {
        accessSection["layer"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        accessSection["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        accessSection["ignore-errors"] = v.(bool)
    }

    if _, ok := d.GetOk("position"); ok {
		if _, ok := d.GetOk("position.top"); ok {
            accessSection["position"] = "top"
        }
        if v, ok := d.GetOk("position.above"); ok {
            accessSection["position"] = map[string]interface{}{"above": v.(string)}
        }
        if v, ok := d.GetOk("position.bottom"); ok {
            accessSection["position"] = map[string]interface{}{"bottom": v.(string)}
        }
        if _, ok := d.GetOk("position.bottom"); ok {
            accessSection["position"] = "bottom"
        }
    }
    log.Println("Create AccessSection - Map = ", accessSection)

    addAccessSectionRes, err := client.ApiCall("add-access-section", accessSection, client.GetSessionID(), true, false)
    if err != nil || !addAccessSectionRes.Success {
        if addAccessSectionRes.ErrorMsg != "" {
            return fmt.Errorf(addAccessSectionRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    d.SetId(addAccessSectionRes.GetData()["uid"].(string))

    return readManagementAccessSection(d, m)
}

func readManagementAccessSection(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    payload := map[string]interface{}{
        "uid": d.Id(),
		"layer": d.Get("layer"),
    }

    showAccessSectionRes, err := client.ApiCall("show-access-section", payload, client.GetSessionID(), true, false)
    if err != nil {
		return fmt.Errorf(err.Error())
	}
    if !showAccessSectionRes.Success {
		if objectNotFound(showAccessSectionRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
        return fmt.Errorf(showAccessSectionRes.ErrorMsg)
    }

    accessSection := showAccessSectionRes.GetData()

    log.Println("Read AccessSection - Show JSON = ", accessSection)

	if v := accessSection["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := accessSection["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := accessSection["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementAccessSection(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
    accessSection := make(map[string]interface{})

    if ok := d.HasChange("name"); ok {
        oldName, newName := d.GetChange("name")
        accessSection["name"] = oldName
        accessSection["new-name"] = newName
    } else {
        accessSection["name"] = d.Get("name")
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
	       accessSection["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
	       accessSection["ignore-errors"] = v.(bool)
    }

    if ok := d.HasChange("position"); ok {
		if _, ok := d.GetOk("position"); ok {
			if _, ok := d.GetOk("position.top"); ok {
                accessSection["new-position"] = "top"
            }
            if v, ok := d.GetOk("position.above"); ok {
                accessSection["new-position"] = map[string]interface{}{"above": v.(string)}
            }
            if v, ok := d.GetOk("position.below"); ok {
                accessSection["new-position"] = map[string]interface{}{"below": v.(string)}
            }
            if _, ok := d.GetOk("position.bottom"); ok {
                accessSection["new-position"] = "bottom"
            }
        }
    }

    log.Println("Update AccessSection - Map = ", accessSection)

    updateAccessSectionRes, err := client.ApiCall("set-access-section", accessSection, client.GetSessionID(), true, false)
    if err != nil || !updateAccessSectionRes.Success {
        if updateAccessSectionRes.ErrorMsg != "" {
            return fmt.Errorf(updateAccessSectionRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }

    return readManagementAccessSection(d, m)
}

func deleteManagementAccessSection(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    accessSectionPayload := map[string]interface{}{
        "uid": d.Id(),
        "layer": d.Get("layer"),
    }

    log.Println("Delete AccessSection")

    deleteAccessSectionRes, err := client.ApiCall("delete-access-section", accessSectionPayload , client.GetSessionID(), true, false)
    if err != nil || !deleteAccessSectionRes.Success {
        if deleteAccessSectionRes.ErrorMsg != "" {
            return fmt.Errorf(deleteAccessSectionRes.ErrorMsg)
        }
        return fmt.Errorf(err.Error())
    }
    d.SetId("")

    return nil
}

