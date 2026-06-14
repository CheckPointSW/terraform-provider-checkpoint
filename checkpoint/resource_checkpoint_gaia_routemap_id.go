package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "encoding/json"
    "strings"

)
func resourceGaiaRoutemapId() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaRoutemapId,
        Read:   readGaiaRoutemapId,
        Update: updateGaiaRoutemapId,
        Delete: deleteGaiaRoutemapId,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Configures Routemap name. If the Routemap with the name doesn't exist, it will be created.<br><br>The Routemap name will be used to configure export and import policy for various dynamic routing protocols.`,
            },
            "resource_id": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `Configures Routemap ID for Routemap name.<br><br>The Routemap ID has match condition and actions for the import and export policies.`,
            },
            "state": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Configures state for Routemap ID. Any of the following values are permissible: allow, restrict, inactive.<br><br>Allow - If route is matched, it will be accepted by import or export policy.<br>Restrict - If route is matched, it will be rejected by the import or export policy.<br>Inactive - If route is matched, it will not be taken into consideration by the import and export policy.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "ids": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "match": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "action": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
        },
    }
}

func createGaiaRoutemapId(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v, ok := d.GetOk("state"); ok {
        payload["state"] = v.(string)
    }

    log.Println("Create RoutemapId - Map = ", payload)

    addRoutemapIdRes, err := client.ApiCallSimple("add-routemap-id", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addRoutemapIdRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addRoutemapIdRes.Success {
            errMsg = addRoutemapIdRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addRoutemapIdRes.GetData()
        }

        debugLogOperation(
            "routemap-id",        // resource type
            "create",                       // operation
            "add-routemap-id",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add routemap-id: %v", err)
    }
    if !addRoutemapIdRes.Success {
        if addRoutemapIdRes.ErrorMsg != "" {
            return fmt.Errorf(addRoutemapIdRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("routemap-id-" + acctest.RandString(10)))
    return readGaiaRoutemapId(d, m)
}

func readGaiaRoutemapId(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showRoutemapIdRes, err := client.ApiCallSimple("show-routemap", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showRoutemapIdRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showRoutemapIdRes.Success {
            errMsg = showRoutemapIdRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showRoutemapIdRes.GetData()
        }

        debugLogOperation(
            "routemap-id",        // resource type
            "read",                       // operation
            "show-routemap",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show routemap-id: %v", err)
    }
    if !showRoutemapIdRes.Success {
        if data := showRoutemapIdRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showRoutemapIdRes.ErrorMsg)
    }

    routemapId := showRoutemapIdRes.GetData()

    log.Println("Read RoutemapId - Show JSON = ", routemapId)

    if v, exists := routemapId["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := routemapId["ids"]; exists {
        if rawList, ok := v.([]interface{}); ok {
            normalized := make([]interface{}, 0, len(rawList))
            for _, elem := range rawList {
                if m, ok := elem.(map[string]interface{}); ok {
                    entry := make(map[string]interface{})
                    for k, val := range m {
                        stateKey := k
                        if k == "id" { stateKey = "resource_id" }
                        if _, isMap := val.(map[string]interface{}); isMap {
                            if b, err := json.Marshal(val); err == nil {
                                entry[stateKey] = string(b)
                            } else {
                                entry[stateKey] = fmt.Sprintf("%v", val)
                            }
                        } else {
                            entry[stateKey] = val
                        }
                    }
                    normalized = append(normalized, entry)
                } else {
                    normalized = append(normalized, elem)
                }
            }
            d.Set("ids", normalized)

            targetId := 0
            if id, ok := d.GetOk("resource_id"); ok {
                targetId = id.(int)
            }
            for _, elem := range rawList {
                if m, ok := elem.(map[string]interface{}); ok {
                    if idVal, ok := m["id"].(float64); ok && int(idVal) == targetId {
                        if stateVal, ok := m["state"].(string); ok {
                            d.Set("state", stateVal)
                        }
                        break
                    }
                }
            }
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaRoutemapId(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v, ok := d.GetOk("state"); ok {
        payload["state"] = v.(string)
    }

    setRoutemapIdRes, err := client.ApiCallSimple("set-routemap-id", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setRoutemapIdRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setRoutemapIdRes.Success {
            errMsg = setRoutemapIdRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setRoutemapIdRes.GetData()
        }

        debugLogOperation(
            "routemap-id",        // resource type
            "update",                       // operation
            "set-routemap-id",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set routemap-id: %v", err)
    }
    if !setRoutemapIdRes.Success {
        return fmt.Errorf(setRoutemapIdRes.ErrorMsg)
    }

    return readGaiaRoutemapId(d, m)
}

func deleteGaiaRoutemapId(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    deleteRoutemapIdRes, err := client.ApiCallSimple("delete-routemap-id", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteRoutemapIdRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteRoutemapIdRes.Success {
            errMsg = deleteRoutemapIdRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteRoutemapIdRes.GetData()
        }

        debugLogOperation(
            "routemap-id",        // resource type
            "delete",                       // operation
            "delete-routemap-id",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete routemap-id: %v", err)
    }
    if !deleteRoutemapIdRes.Success {
        return fmt.Errorf(deleteRoutemapIdRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

