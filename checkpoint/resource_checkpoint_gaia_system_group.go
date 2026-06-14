package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaSystemGroup() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSystemGroup,
        Read:   readGaiaSystemGroup,
        Update: updateGaiaSystemGroup,
        Delete: deleteGaiaSystemGroup,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `N/A`,
            },
            "gid": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `Numeric ID which is used in identifying a group; it must be unique`,
            },
            "users": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `New users to be added to a group. Users, as well as the group, must exist.`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
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

func createGaiaSystemGroup(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("gid"); ok {
        payload["gid"] = v.(int)
    }

    if v := d.Get("users"); len(v.(*schema.Set).List()) > 0 {
        payload["users"] = v.(*schema.Set).List()
    }

    log.Println("Create SystemGroup - Map = ", payload)

    addSystemGroupRes, err := client.ApiCallSimple("add-system-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addSystemGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addSystemGroupRes.Success {
            errMsg = addSystemGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addSystemGroupRes.GetData()
        }

        debugLogOperation(
            "system-group",        // resource type
            "create",                       // operation
            "add-system-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add system-group: %v", err)
    }
    if !addSystemGroupRes.Success {
        if addSystemGroupRes.ErrorMsg != "" {
            return fmt.Errorf(addSystemGroupRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("system-group-" + acctest.RandString(10)))
    return readGaiaSystemGroup(d, m)
}

func readGaiaSystemGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showSystemGroupRes, err := client.ApiCallSimple("show-system-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showSystemGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showSystemGroupRes.Success {
            errMsg = showSystemGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showSystemGroupRes.GetData()
        }

        debugLogOperation(
            "system-group",        // resource type
            "read",                       // operation
            "show-system-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show system-group: %v", err)
    }
    if !showSystemGroupRes.Success {
        if data := showSystemGroupRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showSystemGroupRes.ErrorMsg)
    }

    systemGroup := showSystemGroupRes.GetData()

    log.Println("Read SystemGroup - Show JSON = ", systemGroup)

    if v, exists := systemGroup["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := systemGroup["gid"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("gid", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("gid", _n)
            }
        }
    }
    if v, exists := systemGroup["users"]; exists {
        d.Set("users", v.([]interface{}))
    }
    if v, exists := systemGroup["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaSystemGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("gid"); ok {
        payload["gid"] = v.(int)
    }

    if v := d.Get("users"); len(v.(*schema.Set).List()) > 0 {
        payload["users"] = v.(*schema.Set).List()
    }

    setSystemGroupRes, err := client.ApiCallSimple("set-system-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setSystemGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setSystemGroupRes.Success {
            errMsg = setSystemGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setSystemGroupRes.GetData()
        }

        debugLogOperation(
            "system-group",        // resource type
            "update",                       // operation
            "set-system-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set system-group: %v", err)
    }
    if !setSystemGroupRes.Success {
        return fmt.Errorf(setSystemGroupRes.ErrorMsg)
    }

    return readGaiaSystemGroup(d, m)
}

func deleteGaiaSystemGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteSystemGroupRes, err := client.ApiCallSimple("delete-system-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteSystemGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteSystemGroupRes.Success {
            errMsg = deleteSystemGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteSystemGroupRes.GetData()
        }

        debugLogOperation(
            "system-group",        // resource type
            "delete",                       // operation
            "delete-system-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete system-group: %v", err)
    }
    if !deleteSystemGroupRes.Success {
        return fmt.Errorf(deleteSystemGroupRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

