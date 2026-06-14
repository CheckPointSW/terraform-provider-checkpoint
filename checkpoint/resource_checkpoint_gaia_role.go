package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaRole() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaRole,
        Read:   readGaiaRole,
        Update: updateGaiaRole,
        Delete: deleteGaiaRole,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Role name`,
            },
            "features": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `Specifies which features will be assigned to the role.`,
                Set: func(v interface{}) int { return schema.HashString(v.(map[string]interface{})["name"].(string)) },
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Feature name. Valid values: feature name as shown in show-features API output or 'all' to specify all features. `,
                        },
                        "permission": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Feature permission. Valid values: read-write ,read-only. `,
                        },
                    },
                },
            },
            "extended_commands": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `Specifies which extended commands will be assigned to the role.Valid values: extended commands as shown in show-extended-commands API output or 'all' to specify all extended-commands. `,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "users": {
                Type:        schema.TypeSet,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
        },
    }
}

func createGaiaRole(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v := d.Get("features"); len(v.(*schema.Set).List()) > 0 {
        payload["features"] = v.(*schema.Set).List()
    }

    if v := d.Get("extended_commands"); len(v.(*schema.Set).List()) > 0 {
        payload["extended-commands"] = v.(*schema.Set).List()
    }

    log.Println("Create Role - Map = ", payload)

    addRoleRes, err := client.ApiCallSimple("add-role", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addRoleRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addRoleRes.Success {
            errMsg = addRoleRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addRoleRes.GetData()
        }

        debugLogOperation(
            "role",        // resource type
            "create",                       // operation
            "add-role",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add role: %v", err)
    }
    if !addRoleRes.Success {
        if addRoleRes.ErrorMsg != "" {
            return fmt.Errorf(addRoleRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("role-" + acctest.RandString(10)))
    return readGaiaRole(d, m)
}

func readGaiaRole(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showRoleRes, err := client.ApiCallSimple("show-role", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showRoleRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showRoleRes.Success {
            errMsg = showRoleRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showRoleRes.GetData()
        }

        debugLogOperation(
            "role",        // resource type
            "read",                       // operation
            "show-role",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show role: %v", err)
    }
    if !showRoleRes.Success {
        if data := showRoleRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showRoleRes.ErrorMsg)
    }

    role := showRoleRes.GetData()

    log.Println("Read Role - Show JSON = ", role)

    if v, exists := role["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := role["features"]; exists {
        d.Set("features", v.([]interface{}))
    }
    if v, exists := role["extended-commands"]; exists {
        d.Set("extended_commands", v.([]interface{}))
    }
    if v, exists := role["users"]; exists {
        d.Set("users", v.([]interface{}))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaRole(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v := d.Get("features"); len(v.(*schema.Set).List()) > 0 {
        payload["features"] = v.(*schema.Set).List()
    }

    if v := d.Get("extended_commands"); len(v.(*schema.Set).List()) > 0 {
        payload["extended-commands"] = v.(*schema.Set).List()
    }

    setRoleRes, err := client.ApiCallSimple("set-role", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setRoleRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setRoleRes.Success {
            errMsg = setRoleRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setRoleRes.GetData()
        }

        debugLogOperation(
            "role",        // resource type
            "update",                       // operation
            "set-role",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set role: %v", err)
    }
    if !setRoleRes.Success {
        return fmt.Errorf(setRoleRes.ErrorMsg)
    }

    return readGaiaRole(d, m)
}

func deleteGaiaRole(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteRoleRes, err := client.ApiCallSimple("delete-role", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteRoleRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteRoleRes.Success {
            errMsg = deleteRoleRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteRoleRes.GetData()
        }

        debugLogOperation(
            "role",        // resource type
            "delete",                       // operation
            "delete-role",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete role: %v", err)
    }
    if !deleteRoleRes.Success {
        return fmt.Errorf(deleteRoleRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

