package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaRadius() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaRadius,
        Read:   readGaiaRadius,
        Update: updateGaiaRadius,
        Delete: deleteGaiaRadius,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "nas_ip": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The NAS-IP for the RADIUS client`,
            },
            "default_shell": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Default shell when login`,
            },
            "super_user_uid": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `The UID that will be given to a super user`,
            },
            "servers": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `RADIUS servers list`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "priority": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Server priority (lower values comes first)`,
                        },
                        "address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Server address`,
                        },
                        "port": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `UDP port to contact on the RADIUS server`,
                        },
                        "timeout": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "secret": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: `RADIUS authentication state`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaRadius(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("nas_ip"); ok {
        payload["nas-ip"] = v.(string)
    }

    if v, ok := d.GetOk("default_shell"); ok {
        payload["default-shell"] = v.(string)
    }

    if v, ok := d.GetOk("super_user_uid"); ok {
        payload["super-user-uid"] = v.(string)
    }

    if v := d.Get("servers"); len(v.([]interface{})) > 0 {
        serversList := v.([]interface{})
        serversArray := make([]interface{}, 0, len(serversList))
        for i := range serversList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.priority", i)); ok {
                itemMap["priority"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.address", i)); ok {
                itemMap["address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.port", i)); ok {
                itemMap["port"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.timeout", i)); ok {
                itemMap["timeout"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.secret", i)); ok {
                itemMap["secret"] = v.(string)
            }
            if len(itemMap) > 0 {
                serversArray = append(serversArray, itemMap)
            }
        }
        if len(serversArray) > 0 {
            payload["servers"] = serversArray
        }
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    log.Println("Create Radius - Map = ", payload)

    addRadiusRes, err := client.ApiCallSimple("set-radius", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addRadiusRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addRadiusRes.Success {
            errMsg = addRadiusRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addRadiusRes.GetData()
        }

        debugLogOperation(
            "radius",        // resource type
            "create",                       // operation
            "set-radius",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add radius: %v", err)
    }
    if !addRadiusRes.Success {
        if addRadiusRes.ErrorMsg != "" {
            return fmt.Errorf(addRadiusRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("radius-" + acctest.RandString(10)))
    return readGaiaRadius(d, m)
}

func readGaiaRadius(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showRadiusRes, err := client.ApiCallSimple("show-radius", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showRadiusRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showRadiusRes.Success {
            errMsg = showRadiusRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showRadiusRes.GetData()
        }

        debugLogOperation(
            "radius",        // resource type
            "read",                       // operation
            "show-radius",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show radius: %v", err)
    }
    if !showRadiusRes.Success {
        if data := showRadiusRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showRadiusRes.ErrorMsg)
    }

    radius := showRadiusRes.GetData()

    log.Println("Read Radius - Show JSON = ", radius)

    if v, exists := radius["nas-ip"]; exists {
        d.Set("nas_ip", fmt.Sprintf("%v", v))
    }
    if v, exists := radius["default-shell"]; exists {
        d.Set("default_shell", fmt.Sprintf("%v", v))
    }
    if v, exists := radius["super-user-uid"]; exists {
        d.Set("super_user_uid", fmt.Sprintf("%v", v))
    }
    if v, exists := radius["servers"]; exists {
        if rawList, ok := v.([]interface{}); ok {
            for i, item := range rawList {
                if m, ok := item.(map[string]interface{}); ok {
                    if existing, ok := d.GetOk(fmt.Sprintf("servers.%d.secret", i)); ok {
                        m["secret"] = existing.(string)
                    }
                    rawList[i] = m
                }
            }
            d.Set("servers", rawList)
        }
    }
    if v, exists := radius["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaRadius(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("nas_ip"); ok {
        payload["nas-ip"] = v.(string)
    }

    if v, ok := d.GetOk("default_shell"); ok {
        payload["default-shell"] = v.(string)
    }

    if v, ok := d.GetOk("super_user_uid"); ok {
        payload["super-user-uid"] = v.(string)
    }

    if v := d.Get("servers"); len(v.([]interface{})) > 0 {
        serversList := v.([]interface{})
        serversArray := make([]interface{}, 0, len(serversList))
        for i := range serversList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.priority", i)); ok {
                itemMap["priority"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.address", i)); ok {
                itemMap["address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.port", i)); ok {
                itemMap["port"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.timeout", i)); ok {
                itemMap["timeout"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("servers.%d.secret", i)); ok {
                itemMap["secret"] = v.(string)
            }
            if len(itemMap) > 0 {
                serversArray = append(serversArray, itemMap)
            }
        }
        if len(serversArray) > 0 {
            payload["servers"] = serversArray
        }
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    setRadiusRes, err := client.ApiCallSimple("set-radius", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setRadiusRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setRadiusRes.Success {
            errMsg = setRadiusRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setRadiusRes.GetData()
        }

        debugLogOperation(
            "radius",        // resource type
            "update",                       // operation
            "set-radius",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set radius: %v", err)
    }
    if !setRadiusRes.Success {
        return fmt.Errorf(setRadiusRes.ErrorMsg)
    }

    return readGaiaRadius(d, m)
}

func deleteGaiaRadius(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)
    payload := map[string]interface{}{"servers": []interface{}{}}
    res, err := client.ApiCallSimple("set-radius", payload)
    if err != nil {
        return fmt.Errorf("Failed to clear radius: %v", err)
    }
    if !res.Success {
        return fmt.Errorf(res.ErrorMsg)
    }
    d.SetId("")
    return nil
}

