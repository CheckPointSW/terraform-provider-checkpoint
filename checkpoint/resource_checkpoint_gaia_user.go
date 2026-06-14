package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaUser() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaUser,
        Read:   readGaiaUser,
        Update: updateGaiaUser,
        Delete: deleteGaiaUser,
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
            "uid": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `Specifies a numeric user ID used to identify permissions of a user, duplicate UIDs are not allowed`,
            },
            "homedir": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Specifies the user's home directory as the full UNIX path name where the user is placed on login. If the directory doesn't exist, it is created. Range: Must be under '/home' and must not contain colon (:). `,
            },
            "primary_system_group_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `GID. Numeric ID which is used in identifying the primary group to which this user belongs. `,
            },
            "secondary_system_groups": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `This operation assigns groups to the user. Valid values: must be names of known groups. `,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "password": {
                Type:        schema.TypeString,
                Optional:    true,
                Sensitive:   true,
                Description: `Specifies new password on command line. Check Point recommends that a password be at least eight characters long. A password must contain at least six characters. Enforcement level can be modified via 'password control' feature. `,
            },
            "password_hash": {
                Type:        schema.TypeString,
                Optional:    true,
                Sensitive:   true,
                Description: `An encrypted representation of the password. `,
            },
            "real_name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Specifies a string describing a user; conventionally it's the user's full name. `,
            },
            "shell": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Specifies the user's default command-line interpreter during login. `,
            },
            "allow_access_using": {
                Type:        schema.TypeSet,
                Optional:    true,
                Computed:    true,
                Description: `Modify the access-mechanisms available for a user. Valid values: CLI, Web-UI, Gaia-API (supported from R81.10). `,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "must_change_password": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Forcing password change is relevant only when a password is set. When set to 'True': Force the user to change their password the next time they log in. If they don't log in within the time limit configured in 'set password-controls expiration-lockout-days' they may not be able to log in at all. When set to 'False': If the user was being forced to change their password, cancel that. If the user was locked out due to failure to change their password within the time limit configured in 'set password-controls expiration-lockout-days' they will no longer be locked out. `,
            },
            "roles": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "requires_two_factor_authentication": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Force Two-Factor Authentication for this user. Upon their next login, if Two-Factor Authentication is not already set up, the user will be required to generate the authentication keys.`,
            },
            "unlock": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `If the user has been locked out, cancel that. True: cancel lock-out. False: do nothing. `,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "locked": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaUser(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("uid"); ok {
        payload["uid"] = v.(int)
    }

    if v, ok := d.GetOk("homedir"); ok {
        payload["homedir"] = v.(string)
    }

    if v, ok := d.GetOk("primary_system_group_id"); ok {
        payload["primary-system-group-id"] = v.(int)
    }

    if v := d.Get("secondary_system_groups"); len(v.(*schema.Set).List()) > 0 {
        payload["secondary-system-groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v, ok := d.GetOk("password_hash"); ok {
        payload["password-hash"] = v.(string)
    }

    if v, ok := d.GetOk("real_name"); ok {
        payload["real-name"] = v.(string)
    }

    if v, ok := d.GetOk("shell"); ok {
        payload["shell"] = v.(string)
    }

    if v := d.Get("allow_access_using"); len(v.(*schema.Set).List()) > 0 {
        payload["allow-access-using"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("must_change_password"); ok {
        payload["must-change-password"] = v.(bool)
    }

    if v := d.Get("roles"); len(v.(*schema.Set).List()) > 0 {
        payload["roles"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("requires_two_factor_authentication"); ok {
        payload["requires-two-factor-authentication"] = v.(bool)
    }

    log.Println("Create User - Map = ", payload)

    addUserRes, err := client.ApiCallSimple("add-user", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addUserRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addUserRes.Success {
            errMsg = addUserRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addUserRes.GetData()
        }

        debugLogOperation(
            "user",        // resource type
            "create",                       // operation
            "add-user",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add user: %v", err)
    }
    if !addUserRes.Success {
        if addUserRes.ErrorMsg != "" {
            return fmt.Errorf(addUserRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("user-" + acctest.RandString(10)))
    return readGaiaUser(d, m)
}

func readGaiaUser(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showUserRes, err := client.ApiCallSimple("show-user", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showUserRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showUserRes.Success {
            errMsg = showUserRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showUserRes.GetData()
        }

        debugLogOperation(
            "user",        // resource type
            "read",                       // operation
            "show-user",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show user: %v", err)
    }
    if !showUserRes.Success {
        if data := showUserRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showUserRes.ErrorMsg)
    }

    user := showUserRes.GetData()

    log.Println("Read User - Show JSON = ", user)

    if v, exists := user["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := user["uid"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("uid", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("uid", _n)
            }
        }
    }
    if v, exists := user["homedir"]; exists {
        d.Set("homedir", fmt.Sprintf("%v", v))
    }
    if v, exists := user["primary-system-group-id"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("primary_system_group_id", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("primary_system_group_id", _n)
            }
        }
    }
    if v, exists := user["secondary-system-groups"]; exists {
        d.Set("secondary_system_groups", v.([]interface{}))
    }
    if v, exists := user["real-name"]; exists {
        d.Set("real_name", fmt.Sprintf("%v", v))
    }
    if v, exists := user["shell"]; exists {
        d.Set("shell", fmt.Sprintf("%v", v))
    }
    if v, exists := user["allow-access-using"]; exists {
        d.Set("allow_access_using", v.([]interface{}))
    }
    if v, exists := user["roles"]; exists {
        d.Set("roles", v.([]interface{}))
    }
    if v, exists := user["locked"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("locked", b)
        } else if s, ok := v.(string); ok {
            d.Set("locked", s == "yes" || s == "true")
        }
    }
    if v, exists := user["requires-two-factor-authentication"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("requires_two_factor_authentication", b)
        } else if s, ok := v.(string); ok {
            d.Set("requires_two_factor_authentication", s == "true")
        }
    }
    if v, exists := user["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaUser(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("uid"); ok {
        payload["uid"] = v.(int)
    }

    if v, ok := d.GetOk("homedir"); ok {
        payload["homedir"] = v.(string)
    }

    if v, ok := d.GetOk("primary_system_group_id"); ok {
        payload["primary-system-group-id"] = v.(int)
    }

    if v := d.Get("secondary_system_groups"); len(v.(*schema.Set).List()) > 0 {
        payload["secondary-system-groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v, ok := d.GetOk("password_hash"); ok {
        payload["password-hash"] = v.(string)
    }

    if v, ok := d.GetOk("real_name"); ok {
        payload["real-name"] = v.(string)
    }

    if v, ok := d.GetOk("shell"); ok {
        payload["shell"] = v.(string)
    }

    if v := d.Get("allow_access_using"); len(v.(*schema.Set).List()) > 0 {
        payload["allow-access-using"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("must_change_password"); ok {
        payload["must-change-password"] = v.(bool)
    }

    if v := d.Get("roles"); len(v.(*schema.Set).List()) > 0 {
        payload["roles"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("requires_two_factor_authentication"); ok {
        payload["requires-two-factor-authentication"] = v.(bool)
    }

    if v, ok := d.GetOkExists("unlock"); ok {
        payload["unlock"] = v.(bool)
    }

    setUserRes, err := client.ApiCallSimple("set-user", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setUserRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setUserRes.Success {
            errMsg = setUserRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setUserRes.GetData()
        }

        debugLogOperation(
            "user",        // resource type
            "update",                       // operation
            "set-user",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set user: %v", err)
    }
    if !setUserRes.Success {
        return fmt.Errorf(setUserRes.ErrorMsg)
    }

    return readGaiaUser(d, m)
}

func deleteGaiaUser(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteUserRes, err := client.ApiCallSimple("delete-user", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteUserRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteUserRes.Success {
            errMsg = deleteUserRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteUserRes.GetData()
        }

        debugLogOperation(
            "user",        // resource type
            "delete",                       // operation
            "delete-user",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete user: %v", err)
    }
    if !deleteUserRes.Success {
        return fmt.Errorf(deleteUserRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

