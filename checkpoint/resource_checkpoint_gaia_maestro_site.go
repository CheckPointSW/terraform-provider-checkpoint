package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaMaestroSite() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaMaestroSite,
        Read:   readGaiaMaestroSite,
        Update: updateGaiaMaestroSite,
        Delete: deleteGaiaMaestroSite,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "site_id": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `N/A`,
            },
            "descriptions": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Provide optional site description per Security Group`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "security_group": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `The Site Security Group`,
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Site description`,
                        },
                    },
                },
            },
            "include_pending_changes": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `N/A`,
            },
            "gateways": {
                Type:     schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:     schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

func createGaiaMaestroSite(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("site_id"); ok {
        payload["site-id"] = v.(int)
    }

    if v := d.Get("descriptions"); len(v.([]interface{})) > 0 {
        descriptionsList := v.([]interface{})
        descriptionsArray := make([]interface{}, 0, len(descriptionsList))
        for i := range descriptionsList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("descriptions.%d.security_group", i)); ok {
                itemMap["security-group"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("descriptions.%d.description", i)); ok {
                itemMap["description"] = v.(string)
            }
            if len(itemMap) > 0 {
                descriptionsArray = append(descriptionsArray, itemMap)
            }
        }
        if len(descriptionsArray) > 0 {
            payload["descriptions"] = descriptionsArray
        }
    }

    log.Println("Create MaestroSite - Map = ", payload)

    addMaestroSiteRes, err := client.ApiCallSimple("set-maestro-site", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addMaestroSiteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addMaestroSiteRes.Success {
            errMsg = addMaestroSiteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addMaestroSiteRes.GetData()
        }

        debugLogOperation(
            "maestro-site",        // resource type
            "create",                       // operation
            "set-maestro-site",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add maestro-site: %v", err)
    }
    if !addMaestroSiteRes.Success {
        if addMaestroSiteRes.ErrorMsg != "" {
            return fmt.Errorf(addMaestroSiteRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("maestro-site-" + acctest.RandString(10)))
    return readGaiaMaestroSite(d, m)
}

func readGaiaMaestroSite(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("site_id"); ok {
        payload["site-id"] = v.(int)
    }

    if v, ok := d.GetOkExists("include_pending_changes"); ok {
        payload["include-pending-changes"] = v.(bool)
    }

    showMaestroSiteRes, err := client.ApiCallSimple("show-maestro-site", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showMaestroSiteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showMaestroSiteRes.Success {
            errMsg = showMaestroSiteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showMaestroSiteRes.GetData()
        }

        debugLogOperation(
            "maestro-site",        // resource type
            "read",                       // operation
            "show-maestro-site",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show maestro-site: %v", err)
    }
    if !showMaestroSiteRes.Success {
        if data := showMaestroSiteRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showMaestroSiteRes.ErrorMsg)
    }

    maestroSite := showMaestroSiteRes.GetData()

    log.Println("Read MaestroSite - Show JSON = ", maestroSite)

    if v, exists := maestroSite["site-id"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("site_id", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("site_id", _n)
            }
        }
    }
    if v, exists := maestroSite["descriptions"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    sg := 0
                    if f, ok := m["security-group"].(float64); ok {
                        sg = int(f)
                    }
                    out = append(out, map[string]interface{}{
                        "security_group": sg,
                        "description":    fmt.Sprintf("%v", m["description"]),
                    })
                }
            }
            d.Set("descriptions", out)
        }
    }
    if v, exists := maestroSite["gateways"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    out = append(out, map[string]interface{}{
                        "resource_id": fmt.Sprintf("%v", m["id"]),
                    })
                }
            }
            d.Set("gateways", out)
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaMaestroSite(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("site_id"); ok {
        payload["site-id"] = v.(int)
    }

    if v := d.Get("descriptions"); len(v.([]interface{})) > 0 {
        descriptionsList := v.([]interface{})
        descriptionsArray := make([]interface{}, 0, len(descriptionsList))
        for i := range descriptionsList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("descriptions.%d.security_group", i)); ok {
                itemMap["security-group"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("descriptions.%d.description", i)); ok {
                itemMap["description"] = v.(string)
            }
            if len(itemMap) > 0 {
                descriptionsArray = append(descriptionsArray, itemMap)
            }
        }
        if len(descriptionsArray) > 0 {
            payload["descriptions"] = descriptionsArray
        }
    }

    setMaestroSiteRes, err := client.ApiCallSimple("set-maestro-site", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setMaestroSiteRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setMaestroSiteRes.Success {
            errMsg = setMaestroSiteRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setMaestroSiteRes.GetData()
        }

        debugLogOperation(
            "maestro-site",        // resource type
            "update",                       // operation
            "set-maestro-site",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set maestro-site: %v", err)
    }
    if !setMaestroSiteRes.Success {
        return fmt.Errorf(setMaestroSiteRes.ErrorMsg)
    }

    return readGaiaMaestroSite(d, m)
}

func deleteGaiaMaestroSite(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    