package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaMldInterfaceStaticGroup() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaMldInterfaceStaticGroup,
        Read:   readGaiaMldInterfaceStaticGroup,
        Update: updateGaiaMldInterfaceStaticGroup,
        Delete: deleteGaiaMldInterfaceStaticGroup,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "interface": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The name of the MLD interface`,
            },
            "static_group": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The statically configured group address that this MLD interface receives multicast data for`,
            },
            "group_count": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The number of adjacent static groups`,
            },
            "group_increment": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The increment between MLD static groups (default: ::1)`,
                DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
                normalize := func(s string) string {
                    if s == "default" { return "::1" }
                    return s
                }
                return normalize(old) == normalize(new)
            },
            },
            "sources": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `The list of IPv6 sources from which to receive traffic for this static group`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "source": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The IPv6 source from which to receive traffic for this static group`,
                        },
                        "source_count": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The number of adjacent static group sources`,
                        },
                        "source_increment": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The increment between MLD static group sources (default: ::1)`,
                            DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
                normalize := func(s string) string {
                    if s == "default" { return "::1" }
                    return s
                }
                return normalize(old) == normalize(new)
            },
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

func createGaiaMldInterfaceStaticGroup(d *schema.ResourceData, m interface{}) error {
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
        val := v.(string)
        if val == "default" { val = "::1" }
        payload["group-increment"] = val
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

    log.Println("Create MldInterfaceStaticGroup - Map = ", payload)

    addMldInterfaceStaticGroupRes, err := client.ApiCallSimple("add-mld-interface-static-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addMldInterfaceStaticGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addMldInterfaceStaticGroupRes.Success {
            errMsg = addMldInterfaceStaticGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addMldInterfaceStaticGroupRes.GetData()
        }

        debugLogOperation(
            "mld-interface-static-group",        // resource type
            "create",                       // operation
            "add-mld-interface-static-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add mld-interface-static-group: %v", err)
    }
    if !addMldInterfaceStaticGroupRes.Success {
        if addMldInterfaceStaticGroupRes.ErrorMsg != "" {
            return fmt.Errorf(addMldInterfaceStaticGroupRes.ErrorMsg)
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
        
        setMldInterfaceStaticGroupRes, err := client.ApiCallSimple("set-mld-interface-static-group", updatePayload)
        if err != nil {
            return fmt.Errorf("Failed to apply update-only fields for mld-interface-static-group: %v", err)
        }
        if !setMldInterfaceStaticGroupRes.Success {
            return fmt.Errorf("Failed to apply update-only fields: %s", setMldInterfaceStaticGroupRes.ErrorMsg)
        }
    }

    d.SetId(fmt.Sprintf("mld-interface-static-group-" + acctest.RandString(10)))
    return readGaiaMldInterfaceStaticGroup(d, m)
}

func readGaiaMldInterfaceStaticGroup(d *schema.ResourceData, m interface{}) error {

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

    showMldInterfaceStaticGroupRes, err := client.ApiCallSimple("show-mld-interface-static-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showMldInterfaceStaticGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showMldInterfaceStaticGroupRes.Success {
            errMsg = showMldInterfaceStaticGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showMldInterfaceStaticGroupRes.GetData()
        }

        debugLogOperation(
            "mld-interface-static-group",        // resource type
            "read",                       // operation
            "show-mld-interface-static-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show mld-interface-static-group: %v", err)
    }
    if !showMldInterfaceStaticGroupRes.Success {
        if data := showMldInterfaceStaticGroupRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showMldInterfaceStaticGroupRes.ErrorMsg)
    }

    mldInterfaceStaticGroup := showMldInterfaceStaticGroupRes.GetData()

    log.Println("Read MldInterfaceStaticGroup - Show JSON = ", mldInterfaceStaticGroup)

    if v, exists := mldInterfaceStaticGroup["interface"]; exists {
        d.Set("interface", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterfaceStaticGroup["static-group"]; exists {
        d.Set("static_group", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterfaceStaticGroup["group-count"]; exists {
        d.Set("group_count", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterfaceStaticGroup["group-increment"]; exists {
        s := fmt.Sprintf("%v", v)
        if s == "::1" { s = "default" }
        d.Set("group_increment", s)
    }
    if v, exists := mldInterfaceStaticGroup["sources"]; exists {
        d.Set("sources", v.([]interface{}))
    }
    if v, exists := mldInterfaceStaticGroup["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaMldInterfaceStaticGroup(d *schema.ResourceData, m interface{}) error {

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
        val := v.(string)
        if val == "default" { val = "::1" }
        payload["group-increment"] = val
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

    setMldInterfaceStaticGroupRes, err := client.ApiCallSimple("set-mld-interface-static-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setMldInterfaceStaticGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setMldInterfaceStaticGroupRes.Success {
            errMsg = setMldInterfaceStaticGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setMldInterfaceStaticGroupRes.GetData()
        }

        debugLogOperation(
            "mld-interface-static-group",        // resource type
            "update",                       // operation
            "set-mld-interface-static-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set mld-interface-static-group: %v", err)
    }
    if !setMldInterfaceStaticGroupRes.Success {
        return fmt.Errorf(setMldInterfaceStaticGroupRes.ErrorMsg)
    }

    return readGaiaMldInterfaceStaticGroup(d, m)
}

func deleteGaiaMldInterfaceStaticGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOk("static_group"); ok {
        payload["static-group"] = v.(string)
    }

    deleteMldInterfaceStaticGroupRes, err := client.ApiCallSimple("delete-mld-interface-static-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteMldInterfaceStaticGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteMldInterfaceStaticGroupRes.Success {
            errMsg = deleteMldInterfaceStaticGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteMldInterfaceStaticGroupRes.GetData()
        }

        debugLogOperation(
            "mld-interface-static-group",        // resource type
            "delete",                       // operation
            "delete-mld-interface-static-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete mld-interface-static-group: %v", err)
    }
    if !deleteMldInterfaceStaticGroupRes.Success {
        return fmt.Errorf(deleteMldInterfaceStaticGroupRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

