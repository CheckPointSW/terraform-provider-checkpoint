package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaLogin() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaLogin,
        Read:   readGaiaLogin,
        Delete: deleteGaiaLogin,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "user": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Administrator user name`,
            },
            "password": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Sensitive:   true,
                Description: `Administrator password`,
            },
            "session_timeout": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `Session expiration timeout in seconds`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "verification_code": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Verification code, if Two-Factor Authentication is enabled for this user. This field must be a string comprised solely of digits.`,
            },
            "sid": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "api_server_version": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "last_login_was_at": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "posix": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "iso_8601": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "read_only": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "url": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaLogin(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("user"); ok {
        payload["user"] = v.(string)
    }

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v, ok := d.GetOk("session_timeout"); ok {
        payload["session-timeout"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    if v, ok := d.GetOk("verification_code"); ok {
        payload["verification-code"] = v.(string)
    }

    log.Println("Execute login - Payload = ", payload)

    GaiaLoginRes, err := client.ApiCallSimple("login", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaLoginRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaLoginRes.Success {
            errMsg = GaiaLoginRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaLoginRes.GetData()
        }

        debugLogOperation(
            "login",        // resource type
            "command",                       // operation
            "login",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute login: %v", err)
    }
    if !GaiaLoginRes.Success {
        if GaiaLoginRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaLoginRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaLoginRes.GetData()
    if v, exists := _respData["sid"]; exists {
        d.Set("sid", toString(v))
    }
    if v, exists := _respData["api-server-version"]; exists {
        d.Set("api_server_version", toString(v))
    }
    if v, exists := _respData["last-login-was-at"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            lastloginwasatMap := map[string]interface{}{
                "posix": m["posix"],
                "iso_8601": m["iso-8601"],
            }
            d.Set("last_login_was_at", []interface{}{lastloginwasatMap})
        }
    }
    if v, exists := _respData["read-only"]; exists {
        d.Set("read_only", v)
    }
    if v, exists := _respData["url"]; exists {
        d.Set("url", toString(v))
    }


    d.SetId(fmt.Sprintf("login-" + acctest.RandString(10)))
    return nil
}

func readGaiaLogin(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaLogin(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

