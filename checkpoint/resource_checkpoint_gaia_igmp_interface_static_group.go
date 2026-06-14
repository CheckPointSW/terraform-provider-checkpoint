package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaIgmpInterfaceStaticGroup() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaIgmpInterfaceStaticGroup,
        Read:   readGaiaIgmpInterfaceStaticGroup,
        Update: updateGaiaIgmpInterfaceStaticGroup,
        Delete: deleteGaiaIgmpInterfaceStaticGroup,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "interface": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The name of the IGMP interface`,
            },
            "static_group": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The statically configured group address that this IGMP interface receives multicast data for`,
            },
            "group_count": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The number of adjacent static groups`,
            },
            "group_increment": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `The increment between IGMP static groups (default: 0.0.0.1)`,
            },
            "sources": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `The list of IPv4 sources from which to receive traffic for this static group`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "source": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The IPv4 source from which to receive traffic for this static group`,
                        },
                        "source_count": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The number of adjacent static group sources`,
                        },
                        "source_increment": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The increment between IGMP static group sources (default: 0.0.0.1)`,
                        },
                    },
                },
            },
            "source_all_off": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Remove all sources of a static group`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaIgmpInterfaceStaticGroup(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("static_group"); ok {
        payload["static-group"] = v.(string)
    }

    if v, ok := d.GetOk("group_count"); ok {
        payload["group-count"] = v.(string)
    }

    if v, ok := d.GetOk("group_increment"); ok {
        payload["group-increment"] = v.(string)
    }

    if v := d.Get("sources"); len(v.([]interface{})) > 0 {
        sourcesList := v.([]interface{})
        sourcesArray := make([]interface{}, 0, len(sourcesList))
        for i := range sourcesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("sources.%d.source", i)); ok {
                itemMap["source"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("sources.%d.source_count", i)); ok {
                itemMap["source-count"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("sources.%d.source_increment", i)); ok {
                itemMap["source-increment"] = v.(string)
            }
            if len(itemMap) > 0 {
                sourcesArray = append(sourcesArray, itemMap)
            }
        }
        if len(sourcesArray) > 0 {
            payload["sources"] = sourcesArray
        }
    }

    log.Println("Create IgmpInterfaceStaticGroup - Map = ", payload)

    addIgmpInterfaceStaticGroupRes, err := client.ApiCallSimple("add-igmp-interface-static-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addIgmpInterfaceStaticGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addIgmpInterfaceStaticGroupRes.Success {
            errMsg = addIgmpInterfaceStaticGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addIgmpInterfaceStaticGroupRes.GetData()
        }

        debugLogOperation(
            "igmp-interface-static-group",        // resource type
            "create",                       // operation
            "add-igmp-interface-static-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add igmp-interface-static-group: %v", err)
    }
    if !addIgmpInterfaceStaticGroupRes.Success {
        if addIgmpInterfaceStaticGroupRes.ErrorMsg != "" {
            return fmt.Errorf(addIgmpInterfaceStaticGroupRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    // Two-phase creation: Apply update-only fields if present
    hasUpdateOnlyFields := false
    updatePayload := map[string]interface{}{
    "interface": payload["interface"],
    "static-group": payload["static-group"],
    }

    if v, ok := d.GetOk("source_all_off"); ok {
        updatePayload["source-all-off"] = v.(string)
        hasUpdateOnlyFields = true
    }

    if hasUpdateOnlyFields {
        log.Println("Two-phase creation: Applying update-only fields - Map = ", updatePayload)
        
        setIgmpInterfaceStaticGroupRes, err := client.ApiCallSimple("set-igmp-interface-static-group", updatePayload)
        if err != nil {
            return fmt.Errorf("Failed to apply update-only fields for igmp-interface-static-group: %v", err)
        }
        if !setIgmpInterfaceStaticGroupRes.Success {
            return fmt.Errorf("Failed to apply update-only fields: %s", setIgmpInterfaceStaticGroupRes.ErrorMsg)
        }
    }

    d.SetId(fmt.Sprintf("igmp-interface-static-group-" + acctest.RandString(10)))
    return readGaiaIgmpInterfaceStaticGroup(d, m)
}

func readGaiaIgmpInterfaceStaticGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("static_group"); ok {
        payload["static-group"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showIgmpInterfaceStaticGroupRes, err := client.ApiCallSimple("show-igmp-interface-static-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showIgmpInterfaceStaticGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showIgmpInterfaceStaticGroupRes.Success {
            errMsg = showIgmpInterfaceStaticGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showIgmpInterfaceStaticGroupRes.GetData()
        }

        debugLogOperation(
            "igmp-interface-static-group",        // resource type
            "read",                       // operation
            "show-igmp-interface-static-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show igmp-interface-static-group: %v", err)
    }
    if !showIgmpInterfaceStaticGroupRes.Success {
        if data := showIgmpInterfaceStaticGroupRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showIgmpInterfaceStaticGroupRes.ErrorMsg)
    }

    igmpInterfaceStaticGroup := showIgmpInterfaceStaticGroupRes.GetData()

    log.Println("Read IgmpInterfaceStaticGroup - Show JSON = ", igmpInterfaceStaticGroup)

    if v, exists := igmpInterfaceStaticGroup["interface"]; exists {
        d.Set("interface", fmt.Sprintf("%v", v))
    }
    if v, exists := igmpInterfaceStaticGroup["static-group"]; exists {
        d.Set("static_group", fmt.Sprintf("%v", v))
    }
    if v, exists := igmpInterfaceStaticGroup["group-count"]; exists {
        d.Set("group_count", fmt.Sprintf("%v", v))
    }
    if v, exists := igmpInterfaceStaticGroup["group-increment"]; exists {
        d.Set("group_increment", fmt.Sprintf("%v", v))
    }
    if v, exists := igmpInterfaceStaticGroup["sources"]; exists {
        d.Set("sources", v.([]interface{}))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaIgmpInterfaceStaticGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("static_group"); ok {
        payload["static-group"] = v.(string)
    }

    if v, ok := d.GetOk("group_count"); ok {
        payload["group-count"] = v.(string)
    }

    if v, ok := d.GetOk("group_increment"); ok {
        payload["group-increment"] = v.(string)
    }

    if v := d.Get("sources"); len(v.([]interface{})) > 0 {
        sourcesList := v.([]interface{})
        sourcesArray := make([]interface{}, 0, len(sourcesList))
        for i := range sourcesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("sources.%d.source", i)); ok {
                itemMap["source"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("sources.%d.source_count", i)); ok {
                itemMap["source-count"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("sources.%d.source_increment", i)); ok {
                itemMap["source-increment"] = v.(string)
            }
            if len(itemMap) > 0 {
                sourcesArray = append(sourcesArray, itemMap)
            }
        }
        if len(sourcesArray) > 0 {
            payload["sources"] = sourcesArray
        }
    }

    if v, ok := d.GetOk("source_all_off"); ok {
        payload["source-all-off"] = v.(string)
    }

    setIgmpInterfaceStaticGroupRes, err := client.ApiCallSimple("set-igmp-interface-static-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setIgmpInterfaceStaticGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setIgmpInterfaceStaticGroupRes.Success {
            errMsg = setIgmpInterfaceStaticGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setIgmpInterfaceStaticGroupRes.GetData()
        }

        debugLogOperation(
            "igmp-interface-static-group",        // resource type
            "update",                       // operation
            "set-igmp-interface-static-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set igmp-interface-static-group: %v", err)
    }
    if !setIgmpInterfaceStaticGroupRes.Success {
        return fmt.Errorf(setIgmpInterfaceStaticGroupRes.ErrorMsg)
    }

    return readGaiaIgmpInterfaceStaticGroup(d, m)
}

func deleteGaiaIgmpInterfaceStaticGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("static_group"); ok {
        payload["static-group"] = v.(string)
    }

    deleteIgmpInterfaceStaticGroupRes, err := client.ApiCallSimple("delete-igmp-interface-static-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteIgmpInterfaceStaticGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteIgmpInterfaceStaticGroupRes.Success {
            errMsg = deleteIgmpInterfaceStaticGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteIgmpInterfaceStaticGroupRes.GetData()
        }

        debugLogOperation(
            "igmp-interface-static-group",        // resource type
            "delete",                       // operation
            "delete-igmp-interface-static-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete igmp-interface-static-group: %v", err)
    }
    if !deleteIgmpInterfaceStaticGroupRes.Success {
        return fmt.Errorf(deleteIgmpInterfaceStaticGroupRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

