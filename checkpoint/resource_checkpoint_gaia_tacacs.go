package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaTacacs() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaTacacs,
        Read:   readGaiaTacacs,
        Update: updateGaiaTacacs,
        Delete: deleteGaiaTacacs,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `TACACS+ authentication state`,
            },
            "super_user_uid": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The UID that will be given to a TACACS+ user`,
            },
            "servers": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `TACACS+ servers list`,
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
                            Description: `The server address`,
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
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaTacacs(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
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

    log.Println("Create Tacacs - Map = ", payload)

    addTacacsRes, err := client.ApiCallSimple("set-tacacs", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addTacacsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addTacacsRes.Success {
            errMsg = addTacacsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addTacacsRes.GetData()
        }

        debugLogOperation(
            "tacacs",        // resource type
            "create",                       // operation
            "set-tacacs",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add tacacs: %v", err)
    }
    if !addTacacsRes.Success {
        if addTacacsRes.ErrorMsg != "" {
            return fmt.Errorf(addTacacsRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("tacacs-" + acctest.RandString(10)))
    return readGaiaTacacs(d, m)
}

func readGaiaTacacs(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showTacacsRes, err := client.ApiCallSimple("show-tacacs", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showTacacsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showTacacsRes.Success {
            errMsg = showTacacsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showTacacsRes.GetData()
        }

        debugLogOperation(
            "tacacs",        // resource type
            "read",                       // operation
            "show-tacacs",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show tacacs: %v", err)
    }
    if !showTacacsRes.Success {
        if data := showTacacsRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showTacacsRes.ErrorMsg)
    }

    tacacs := showTacacsRes.GetData()

    log.Println("Read Tacacs - Show JSON = ", tacacs)

    if v, exists := tacacs["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := tacacs["super-user-uid"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("super_user_uid", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("super_user_uid", _n)
            }
        }
    }
    if v, exists := tacacs["servers"]; exists {
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
    d.SetId(d.Id())
    return nil
}

func updateGaiaTacacs(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
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

    setTacacsRes, err := client.ApiCallSimple("set-tacacs", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setTacacsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setTacacsRes.Success {
            errMsg = setTacacsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setTacacsRes.GetData()
        }

        debugLogOperation(
            "tacacs",        // resource type
            "update",                       // operation
            "set-tacacs",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set tacacs: %v", err)
    }
    if !setTacacsRes.Success {
        return fmt.Errorf(setTacacsRes.ErrorMsg)
    }

    return readGaiaTacacs(d, m)
}

func deleteGaiaTacacs(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    