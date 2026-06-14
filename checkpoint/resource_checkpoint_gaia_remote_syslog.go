package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaRemoteSyslog() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaRemoteSyslog,
        Read:   readGaiaRemoteSyslog,
        Update: updateGaiaRemoteSyslog,
        Delete: deleteGaiaRemoteSyslog,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "server_ip": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Remote server address, IPv6 and Hostname supported from R82.`,
            },
            "level": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `N/A`,
            },
            "port": {
                Type:        schema.TypeInt,
                Optional:    true,
                Computed:    true,
                Description: `Log port. Supported starting from Gaia version R81.20`,
            },
            "protocol": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Log protocol. Supported starting from Gaia version R81.20`,
            },
            "queuing_mechanism": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Log queuing mechanism state. Supported starting from Gaia version R82`,
            },
            "tls_encryption": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `TLS encryption status. Supported starting from Gaia version R82`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `TLS Encryption state. Supported starting from Gaia version R82`,
                        },
                        "auth_mode": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Mode used for TLS authentication. supported modes:           name - certificate validation and subject name authentication. Most secure mode.           fingerprint - fingerprint of the server's certificate (which can be self-signed).           certvalid - server's certificate validation only.             It validates the remote peers certificate, but does not check the subject name.             It is recommended NOT to use this mode.           anon - anonymous authentication, this mode is vulnerable to man in the middle attacks as well as unauthorized access.             It is recommended NOT to use this mode.         Note:           When setting anon or certvalid modes, permitted-peers of the remote host will removed.           When setting name or fingerprint modes, it is mandatory to set permitted-peers for the remote host. Supported starting from Gaia version R82`,
                        },
                        "permitted_peers": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Computed:    true,
                            Description: `Common name CN or fingerprint of the permitted peer.          In case of fingerprint, Accepted SHA1.          To specify multiple remote peers separate each one with a comma.          This parameter is mandatory when using auth-mode of 'name' or 'fingerprint'.           Note that usually a single remote peer should be all you need.          Support for multiple peers is primarily included in support of load balancing scenarios.          If the connection goes to a specific server, only one specific peer is ever expected. Supported starting from Gaia version R82`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
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

func createGaiaRemoteSyslog(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("server_ip"); ok {
        payload["server-ip"] = v.(string)
    }

    if v, ok := d.GetOk("level"); ok {
        payload["level"] = v.(string)
    }

    if v, ok := d.GetOk("port"); ok {
        payload["port"] = v.(int)
    }

    if v, ok := d.GetOk("protocol"); ok {
        payload["protocol"] = v.(string)
    }

    if v, ok := d.GetOkExists("queuing_mechanism"); ok {
        payload["queuing-mechanism"] = v.(bool)
    }

    if v := d.Get("tls_encryption"); len(v.([]interface{})) > 0 {
        _ = v
        tlsencryptionMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("tls_encryption.0.enabled"); ok && v.(bool) {
            tlsencryptionMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("tls_encryption.0.auth_mode"); ok {
            tlsencryptionMap["auth-mode"] = v.(string)
        }
        if v := d.Get("tls_encryption.0.permitted_peers"); len(v.(*schema.Set).List()) > 0 {
            tlsencryptionMap["permitted-peers"] = v.(*schema.Set).List()
        }
        if len(tlsencryptionMap) > 0 {
            payload["tls-encryption"] = tlsencryptionMap
        }
    }

    log.Println("Create RemoteSyslog - Map = ", payload)

    addRemoteSyslogRes, err := client.ApiCallSimple("add-remote-syslog", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addRemoteSyslogRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addRemoteSyslogRes.Success {
            errMsg = addRemoteSyslogRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addRemoteSyslogRes.GetData()
        }

        debugLogOperation(
            "remote-syslog",        // resource type
            "create",                       // operation
            "add-remote-syslog",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add remote-syslog: %v", err)
    }
    if !addRemoteSyslogRes.Success {
        if addRemoteSyslogRes.ErrorMsg != "" {
            return fmt.Errorf(addRemoteSyslogRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("remote-syslog-" + acctest.RandString(10)))
    return readGaiaRemoteSyslog(d, m)
}

func readGaiaRemoteSyslog(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("server_ip"); ok {
        payload["server-ip"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showRemoteSyslogRes, err := client.ApiCallSimple("show-remote-syslog", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showRemoteSyslogRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showRemoteSyslogRes.Success {
            errMsg = showRemoteSyslogRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showRemoteSyslogRes.GetData()
        }

        debugLogOperation(
            "remote-syslog",        // resource type
            "read",                       // operation
            "show-remote-syslog",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show remote-syslog: %v", err)
    }
    if !showRemoteSyslogRes.Success {
        if data := showRemoteSyslogRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showRemoteSyslogRes.ErrorMsg)
    }

    remoteSyslog := showRemoteSyslogRes.GetData()

    log.Println("Read RemoteSyslog - Show JSON = ", remoteSyslog)

    if v, exists := remoteSyslog["server-ip"]; exists {
        d.Set("server_ip", fmt.Sprintf("%v", v))
    }
    if v, exists := remoteSyslog["level"]; exists {
        d.Set("level", fmt.Sprintf("%v", v))
    }
    if v, exists := remoteSyslog["port"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("port", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("port", _n)
            }
        }
    }
    if v, exists := remoteSyslog["protocol"]; exists {
        d.Set("protocol", fmt.Sprintf("%v", v))
    }
    if v, exists := remoteSyslog["tls-encryption"]; exists {
        d.Set("tls_encryption", v)
    }
    if v, exists := remoteSyslog["queuing-mechanism"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("queuing_mechanism", b)
        } else if s, ok := v.(string); ok {
            d.Set("queuing_mechanism", s == "true")
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaRemoteSyslog(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("server_ip"); ok {
        payload["server-ip"] = v.(string)
    }

    if v, ok := d.GetOk("level"); ok {
        payload["level"] = v.(string)
    }

    if v, ok := d.GetOk("port"); ok {
        payload["port"] = v.(int)
    }

    if v, ok := d.GetOk("protocol"); ok {
        payload["protocol"] = v.(string)
    }

    if v, ok := d.GetOkExists("queuing_mechanism"); ok {
        payload["queuing-mechanism"] = v.(bool)
    }

    if v := d.Get("tls_encryption"); len(v.([]interface{})) > 0 {
        _ = v
        tlsencryptionMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("tls_encryption.0.enabled"); ok && v.(bool) {
            tlsencryptionMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("tls_encryption.0.auth_mode"); ok {
            tlsencryptionMap["auth-mode"] = v.(string)
        }
        if v := d.Get("tls_encryption.0.permitted_peers"); len(v.(*schema.Set).List()) > 0 {
            tlsencryptionMap["permitted-peers"] = v.(*schema.Set).List()
        }
        if len(tlsencryptionMap) > 0 {
            payload["tls-encryption"] = tlsencryptionMap
        }
    }

    setRemoteSyslogRes, err := client.ApiCallSimple("set-remote-syslog", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setRemoteSyslogRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setRemoteSyslogRes.Success {
            errMsg = setRemoteSyslogRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setRemoteSyslogRes.GetData()
        }

        debugLogOperation(
            "remote-syslog",        // resource type
            "update",                       // operation
            "set-remote-syslog",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set remote-syslog: %v", err)
    }
    if !setRemoteSyslogRes.Success {
        return fmt.Errorf(setRemoteSyslogRes.ErrorMsg)
    }

    return readGaiaRemoteSyslog(d, m)
}

func deleteGaiaRemoteSyslog(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("server_ip"); ok {
        payload["server-ip"] = v.(string)
    }

    deleteRemoteSyslogRes, err := client.ApiCallSimple("delete-remote-syslog", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteRemoteSyslogRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteRemoteSyslogRes.Success {
            errMsg = deleteRemoteSyslogRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteRemoteSyslogRes.GetData()
        }

        debugLogOperation(
            "remote-syslog",        // resource type
            "delete",                       // operation
            "delete-remote-syslog",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete remote-syslog: %v", err)
    }
    if !deleteRemoteSyslogRes.Success {
        return fmt.Errorf(deleteRemoteSyslogRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

