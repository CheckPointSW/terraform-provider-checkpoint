package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaSshServerSettings() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSshServerSettings,
        Read:   readGaiaSshServerSettings,
        Update: updateGaiaSshServerSettings,
        Delete: deleteGaiaSshServerSettings,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled_ciphers": {
                Type:        schema.TypeList,
                Optional:    true,
                MaxItems:    1,
                Description: `Specifies the SSH ciphers that are enabled. Ciphers are encryption algorithms used to secure SSH connections.`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "add": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `List of algorithms to enable.`,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                        "remove": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `List of algorithms to disable.`,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },
            "enabled_mac_algorithms": {
                Type:        schema.TypeList,
                Optional:    true,
                MaxItems:    1,
                Description: `Specifies the SSH MAC (Message Authentication Code) algorithms that are enabled. These algorithms ensure data integrity and authenticity during SSH communication.`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "add": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `List of algorithms to enable.`,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                        "remove": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `List of algorithms to disable.`,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },
            "enabled_kex_algorithms": {
                Type:        schema.TypeList,
                Optional:    true,
                MaxItems:    1,
                Description: `Specifies the SSH key exchange (KEX) algorithms that are enabled. These algorithms are used to securely exchange cryptographic keys between the client and server.`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "add": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `List of algorithms to enable.`,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                        "remove": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `List of algorithms to disable.`,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },
            "enabled_public_key_algorithms": {
                Type:        schema.TypeList,
                Optional:    true,
                MaxItems:    1,
                Description: `Specifies the SSH public key algorithms that are enabled. These algorithms are used for authenticating the client to the server using public key cryptography.`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "add": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `List of algorithms to enable.`,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                        "remove": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `List of algorithms to disable.`,
                            Elem: &schema.Schema{Type: schema.TypeString},
                        },
                    },
                },
            },
            "password_authentication": {
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Sensitive:   true,
                Description: `Enables or disables password authentication. When enabled, users can authenticate using a password.`,
            },
            "permit_root_login": {
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: `Enables or disables root login. When enabled, the root user is allowed to log in directly.`,
            },
            "use_dns": {
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: `Enables or disables DNS usage. When enabled, the server performs a reverse DNS lookup to resolve the client's IP to a hostname.`,
            },
            "client_alive_interval": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: `Sets the interval (in seconds) for sending alive messages to the client. This helps in keeping the connection active and detecting unresponsive clients.`,
            },
            "login_grace_time": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: `Sets the time (in seconds) allowed for a user to successfully log in. If the user fails to log in within this time, the server disconnects the session.`,
            },
            "include_disabled_values": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Include disabled algorithms`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaSshServerSettings(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("enabled_ciphers"); len(v.([]interface{})) > 0 {
        block := v.([]interface{})[0].(map[string]interface{})
        algos := make(map[string]interface{})
        if add, ok := block["add"].([]interface{}); ok && len(add) > 0 {
            algos["add"] = add
        }
        if rem, ok := block["remove"].([]interface{}); ok && len(rem) > 0 {
            algos["remove"] = rem
        }
        if len(algos) > 0 {
            payload["enabled-ciphers"] = algos
        }
    }
    if v := d.Get("enabled_mac_algorithms"); len(v.([]interface{})) > 0 {
        block := v.([]interface{})[0].(map[string]interface{})
        algos := make(map[string]interface{})
        if add, ok := block["add"].([]interface{}); ok && len(add) > 0 {
            algos["add"] = add
        }
        if rem, ok := block["remove"].([]interface{}); ok && len(rem) > 0 {
            algos["remove"] = rem
        }
        if len(algos) > 0 {
            payload["enabled-mac-algorithms"] = algos
        }
    }
    if v := d.Get("enabled_kex_algorithms"); len(v.([]interface{})) > 0 {
        block := v.([]interface{})[0].(map[string]interface{})
        algos := make(map[string]interface{})
        if add, ok := block["add"].([]interface{}); ok && len(add) > 0 {
            algos["add"] = add
        }
        if rem, ok := block["remove"].([]interface{}); ok && len(rem) > 0 {
            algos["remove"] = rem
        }
        if len(algos) > 0 {
            payload["enabled-kex-algorithms"] = algos
        }
    }
    if v := d.Get("enabled_public_key_algorithms"); len(v.([]interface{})) > 0 {
        block := v.([]interface{})[0].(map[string]interface{})
        algos := make(map[string]interface{})
        if add, ok := block["add"].([]interface{}); ok && len(add) > 0 {
            algos["add"] = add
        }
        if rem, ok := block["remove"].([]interface{}); ok && len(rem) > 0 {
            algos["remove"] = rem
        }
        if len(algos) > 0 {
            payload["enabled-public-key-algorithms"] = algos
        }
    }
    if v, ok := d.GetOkExists("password_authentication"); ok {
        payload["password-authentication"] = v.(bool)
    }

    if v, ok := d.GetOkExists("permit_root_login"); ok {
        payload["permit-root-login"] = v.(bool)
    }

    if v, ok := d.GetOkExists("use_dns"); ok {
        payload["use-dns"] = v.(bool)
    }

    if v, ok := d.GetOk("client_alive_interval"); ok {
        payload["client-alive-interval"] = v.(int)
    }

    if v, ok := d.GetOk("login_grace_time"); ok {
        payload["login-grace-time"] = v.(int)
    }

    log.Println("Create SshServerSettings - Map = ", payload)

    addSshServerSettingsRes, err := client.ApiCallSimple("set-ssh-server-settings", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addSshServerSettingsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addSshServerSettingsRes.Success {
            errMsg = addSshServerSettingsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addSshServerSettingsRes.GetData()
        }

        debugLogOperation(
            "ssh-server-settings",        // resource type
            "create",                       // operation
            "set-ssh-server-settings",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add ssh-server-settings: %v", err)
    }
    if !addSshServerSettingsRes.Success {
        if addSshServerSettingsRes.ErrorMsg != "" {
            return fmt.Errorf(addSshServerSettingsRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("ssh-server-settings-" + acctest.RandString(10)))
    return readGaiaSshServerSettings(d, m)
}

func readGaiaSshServerSettings(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("include_disabled_values"); ok {
        payload["include-disabled-values"] = v.(bool)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showSshServerSettingsRes, err := client.ApiCallSimple("show-ssh-server-settings", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showSshServerSettingsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showSshServerSettingsRes.Success {
            errMsg = showSshServerSettingsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showSshServerSettingsRes.GetData()
        }

        debugLogOperation(
            "ssh-server-settings",        // resource type
            "read",                       // operation
            "show-ssh-server-settings",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show ssh-server-settings: %v", err)
    }
    if !showSshServerSettingsRes.Success {
        if data := showSshServerSettingsRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showSshServerSettingsRes.ErrorMsg)
    }

    sshServerSettings := showSshServerSettingsRes.GetData()

    log.Println("Read SshServerSettings - Show JSON = ", sshServerSettings)

    if v, exists := sshServerSettings["password-authentication"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("password_authentication", b)
        } else if s, ok := v.(string); ok {
            d.Set("password_authentication", s == "true")
        }
    }
    if v, exists := sshServerSettings["permit-root-login"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("permit_root_login", b)
        } else if s, ok := v.(string); ok {
            d.Set("permit_root_login", s == "true")
        }
    }
    if v, exists := sshServerSettings["use-dns"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("use_dns", b)
        } else if s, ok := v.(string); ok {
            d.Set("use_dns", s == "true")
        }
    }
    if v, exists := sshServerSettings["client-alive-interval"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("client_alive_interval", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("client_alive_interval", _n)
            }
        }
    }
    if v, exists := sshServerSettings["login-grace-time"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("login_grace_time", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("login_grace_time", _n)
            }
        }
    }
    if v, exists := sshServerSettings["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaSshServerSettings(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("enabled_ciphers"); len(v.([]interface{})) > 0 {
        block := v.([]interface{})[0].(map[string]interface{})
        algos := make(map[string]interface{})
        if add, ok := block["add"].([]interface{}); ok && len(add) > 0 {
            algos["add"] = add
        }
        if rem, ok := block["remove"].([]interface{}); ok && len(rem) > 0 {
            algos["remove"] = rem
        }
        if len(algos) > 0 {
            payload["enabled-ciphers"] = algos
        }
    }
    if v := d.Get("enabled_mac_algorithms"); len(v.([]interface{})) > 0 {
        block := v.([]interface{})[0].(map[string]interface{})
        algos := make(map[string]interface{})
        if add, ok := block["add"].([]interface{}); ok && len(add) > 0 {
            algos["add"] = add
        }
        if rem, ok := block["remove"].([]interface{}); ok && len(rem) > 0 {
            algos["remove"] = rem
        }
        if len(algos) > 0 {
            payload["enabled-mac-algorithms"] = algos
        }
    }
    if v := d.Get("enabled_kex_algorithms"); len(v.([]interface{})) > 0 {
        block := v.([]interface{})[0].(map[string]interface{})
        algos := make(map[string]interface{})
        if add, ok := block["add"].([]interface{}); ok && len(add) > 0 {
            algos["add"] = add
        }
        if rem, ok := block["remove"].([]interface{}); ok && len(rem) > 0 {
            algos["remove"] = rem
        }
        if len(algos) > 0 {
            payload["enabled-kex-algorithms"] = algos
        }
    }
    if v := d.Get("enabled_public_key_algorithms"); len(v.([]interface{})) > 0 {
        block := v.([]interface{})[0].(map[string]interface{})
        algos := make(map[string]interface{})
        if add, ok := block["add"].([]interface{}); ok && len(add) > 0 {
            algos["add"] = add
        }
        if rem, ok := block["remove"].([]interface{}); ok && len(rem) > 0 {
            algos["remove"] = rem
        }
        if len(algos) > 0 {
            payload["enabled-public-key-algorithms"] = algos
        }
    }
    if v, ok := d.GetOkExists("password_authentication"); ok {
        payload["password-authentication"] = v.(bool)
    }

    if v, ok := d.GetOkExists("permit_root_login"); ok {
        payload["permit-root-login"] = v.(bool)
    }

    if v, ok := d.GetOkExists("use_dns"); ok {
        payload["use-dns"] = v.(bool)
    }

    if v, ok := d.GetOk("client_alive_interval"); ok {
        payload["client-alive-interval"] = v.(int)
    }

    if v, ok := d.GetOk("login_grace_time"); ok {
        payload["login-grace-time"] = v.(int)
    }

    setSshServerSettingsRes, err := client.ApiCallSimple("set-ssh-server-settings", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setSshServerSettingsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setSshServerSettingsRes.Success {
            errMsg = setSshServerSettingsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setSshServerSettingsRes.GetData()
        }

        debugLogOperation(
            "ssh-server-settings",        // resource type
            "update",                       // operation
            "set-ssh-server-settings",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set ssh-server-settings: %v", err)
    }
    if !setSshServerSettingsRes.Success {
        return fmt.Errorf(setSshServerSettingsRes.ErrorMsg)
    }

    return readGaiaSshServerSettings(d, m)
}

func deleteGaiaSshServerSettings(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    