package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaMdps() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaMdps,
        Read:   readGaiaMdps,
        Update: updateGaiaMdps,
        Delete: deleteGaiaMdps,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "separation_interfaces": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Management and/or Sync interface for the separation options`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "management": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Management interface`,
                        },
                        "sync": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Sync interface (used in ClusterXL)`,
                        },
                    },
                },
            },
            "routing_separation": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Routing Separation`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Routing separation state`,
                        },
                    },
                },
            },
            "resource_separation": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Resource Separation`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Resource separation state`,
                        },
                        "allocated_cpus": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Number of CPU's for resource separation`,
                        },
                    },
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaMdps(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("separation_interfaces"); len(v.([]interface{})) > 0 {
        _ = v
        separationinterfacesMap := make(map[string]interface{})
        if v, ok := d.GetOk("separation_interfaces.0.management"); ok {
            separationinterfacesMap["management"] = v.(string)
        }
        if v, ok := d.GetOk("separation_interfaces.0.sync"); ok {
            separationinterfacesMap["sync"] = v.(string)
        }
        if len(separationinterfacesMap) > 0 {
            payload["separation-interfaces"] = separationinterfacesMap
        }
    }

    if v, ok := d.GetOkExists("routing_separation.0.enabled"); ok {
        routingseparationMap := map[string]interface{}{}
        routingseparationMap["enabled"] = v.(bool)
        payload["routing-separation"] = routingseparationMap
    }

    if v, ok := d.GetOkExists("resource_separation.0.enabled"); ok {
        resourceseparationMap := map[string]interface{}{}
        resourceseparationMap["enabled"] = v.(bool)
        if cpus, ok2 := d.GetOk("resource_separation.0.allocated_cpus"); ok2 {
            resourceseparationMap["allocated-cpus"] = cpus.(int)
        }
        payload["resource-separation"] = resourceseparationMap
    }

    log.Println("Create Mdps - Map = ", payload)

    addMdpsRes, err := client.ApiCallSimple("set-mdps", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addMdpsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addMdpsRes.Success {
            errMsg = addMdpsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addMdpsRes.GetData()
        }

        debugLogOperation(
            "mdps",        // resource type
            "create",                       // operation
            "set-mdps",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add mdps: %v", err)
    }
    if !addMdpsRes.Success {
        if addMdpsRes.ErrorMsg != "" {
            return fmt.Errorf(addMdpsRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("mdps-" + acctest.RandString(10)))
    return readGaiaMdps(d, m)
}

func readGaiaMdps(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showMdpsRes, err := client.ApiCallSimple("show-mdps", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showMdpsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showMdpsRes.Success {
            errMsg = showMdpsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showMdpsRes.GetData()
        }

        debugLogOperation(
            "mdps",        // resource type
            "read",                       // operation
            "show-mdps",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show mdps: %v", err)
    }
    if !showMdpsRes.Success {
        if data := showMdpsRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showMdpsRes.ErrorMsg)
    }

    mdps := showMdpsRes.GetData()

    log.Println("Read Mdps - Show JSON = ", mdps)

    if v, exists := mdps["separation-interfaces"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            d.Set("separation_interfaces", []interface{}{map[string]interface{}{
                "management": fmt.Sprintf("%v", m["management"]),
                "sync":       fmt.Sprintf("%v", m["sync"]),
            }})
        }
    }
    if v, exists := mdps["routing-separation"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            d.Set("routing_separation", []interface{}{map[string]interface{}{
                "enabled": m["enabled"],
            }})
        }
    }
    if v, exists := mdps["resource-separation"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            cpus := 0
            if f, ok := m["allocated-cpus"].(float64); ok {
                cpus = int(f)
            }
            d.Set("resource_separation", []interface{}{map[string]interface{}{
                "enabled":        m["enabled"],
                "allocated_cpus": cpus,
            }})
        }
    }
    if v, exists := mdps["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaMdps(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("separation_interfaces"); len(v.([]interface{})) > 0 {
        _ = v
        separationinterfacesMap := make(map[string]interface{})
        if v, ok := d.GetOk("separation_interfaces.0.management"); ok {
            separationinterfacesMap["management"] = v.(string)
        }
        if v, ok := d.GetOk("separation_interfaces.0.sync"); ok {
            separationinterfacesMap["sync"] = v.(string)
        }
        if len(separationinterfacesMap) > 0 {
            payload["separation-interfaces"] = separationinterfacesMap
        }
    }

    if v, ok := d.GetOkExists("routing_separation.0.enabled"); ok {
        routingseparationMap := map[string]interface{}{}
        routingseparationMap["enabled"] = v.(bool)
        payload["routing-separation"] = routingseparationMap
    }

    if v, ok := d.GetOkExists("resource_separation.0.enabled"); ok {
        resourceseparationMap := map[string]interface{}{}
        resourceseparationMap["enabled"] = v.(bool)
        if cpus, ok2 := d.GetOk("resource_separation.0.allocated_cpus"); ok2 {
            resourceseparationMap["allocated-cpus"] = cpus.(int)
        }
        payload["resource-separation"] = resourceseparationMap
    }

    setMdpsRes, err := client.ApiCallSimple("set-mdps", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setMdpsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setMdpsRes.Success {
            errMsg = setMdpsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setMdpsRes.GetData()
        }

        debugLogOperation(
            "mdps",        // resource type
            "update",                       // operation
            "set-mdps",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set mdps: %v", err)
    }
    if !setMdpsRes.Success {
        return fmt.Errorf(setMdpsRes.ErrorMsg)
    }

    return readGaiaMdps(d, m)
}

func deleteGaiaMdps(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    