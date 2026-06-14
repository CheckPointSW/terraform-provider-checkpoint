package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaSyslog() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSyslog,
        Read:   readGaiaSyslog,
        Update: updateGaiaSyslog,
        Delete: deleteGaiaSyslog,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "audit_log": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `syslog auditlog permanent`,
            },
            "cp_logs": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `syslog auditlog permanent`,
            },
            "send_to_mgmt": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `sending logs to Management server`,
            },
            "filename": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `syslog output filename`,
            },
            "tls_configuration": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `system TLS configuration in order to enable sending encrtyped syslog messages to remote host, Supported starting from R82`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ca_certification": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Certificate file path of the certification authority CA. Supported starting from Gaia version R82`,
                        },
                        "public_key": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Public key file path signed by the CA. Supported starting from Gaia version R82`,
                        },
                        "private_key": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Private key file path. Supported starting from Gaia version R82`,
                        },
                    },
                },
            },
            "forwarded_logs_files": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Custom log files List. Supported starting from Gaia version R82.10`,
                Computed:    true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "path": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Path to the forwarded log file.`,
                        },
                        "tag": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Computed:    true,
                            Description: `Tag for the forwarded log file.`,
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

func createGaiaSyslog(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("audit_log"); ok {
        payload["audit-log"] = v.(bool)
    }

    if v, ok := d.GetOkExists("cp_logs"); ok {
        payload["cp-logs"] = v.(bool)
    }

    if v, ok := d.GetOkExists("send_to_mgmt"); ok {
        payload["send-to-mgmt"] = v.(bool)
    }

    if v, ok := d.GetOk("filename"); ok {
        payload["filename"] = v.(string)
    }

    if v := d.Get("tls_configuration"); len(v.([]interface{})) > 0 {
        _ = v
        tlsconfigurationMap := make(map[string]interface{})
        if v, ok := d.GetOk("tls_configuration.0.ca_certification"); ok {
            tlsconfigurationMap["ca_certification"] = v.(string)
        }
        if v, ok := d.GetOk("tls_configuration.0.public_key"); ok {
            tlsconfigurationMap["public_key"] = v.(string)
        }
        if v, ok := d.GetOk("tls_configuration.0.private_key"); ok {
            tlsconfigurationMap["private_key"] = v.(string)
        }
        if len(tlsconfigurationMap) > 0 {
            payload["tls-configuration"] = tlsconfigurationMap
        }
    }

    if v := d.Get("forwarded_logs_files"); len(v.([]interface{})) > 0 {
        forwardedlogsfilesList := v.([]interface{})
        forwardedlogsfilesArray := make([]interface{}, 0, len(forwardedlogsfilesList))
        for i := range forwardedlogsfilesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("forwarded_logs_files.%d.path", i)); ok {
                itemMap["path"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("forwarded_logs_files.%d.tag", i)); ok {
                itemMap["tag"] = v.(string)
            }
            if len(itemMap) > 0 {
                forwardedlogsfilesArray = append(forwardedlogsfilesArray, itemMap)
            }
        }
        if len(forwardedlogsfilesArray) > 0 {
            payload["forwarded-logs-files"] = forwardedlogsfilesArray
        }
    }

    log.Println("Create Syslog - Map = ", payload)

    addSyslogRes, err := client.ApiCallSimple("set-syslog", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addSyslogRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addSyslogRes.Success {
            errMsg = addSyslogRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addSyslogRes.GetData()
        }

        debugLogOperation(
            "syslog",        // resource type
            "create",                       // operation
            "set-syslog",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add syslog: %v", err)
    }
    if !addSyslogRes.Success {
        if addSyslogRes.ErrorMsg != "" {
            return fmt.Errorf(addSyslogRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("syslog-" + acctest.RandString(10)))
    return readGaiaSyslog(d, m)
}

func readGaiaSyslog(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showSyslogRes, err := client.ApiCallSimple("show-syslog", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showSyslogRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showSyslogRes.Success {
            errMsg = showSyslogRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showSyslogRes.GetData()
        }

        debugLogOperation(
            "syslog",        // resource type
            "read",                       // operation
            "show-syslog",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show syslog: %v", err)
    }
    if !showSyslogRes.Success {
        if data := showSyslogRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showSyslogRes.ErrorMsg)
    }

    syslog := showSyslogRes.GetData()

    log.Println("Read Syslog - Show JSON = ", syslog)

    if v, exists := syslog["audit-log"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("audit_log", b)
        } else if s, ok := v.(string); ok {
            d.Set("audit_log", s == "true")
        }
    }
    if v, exists := syslog["cp-logs"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("cp_logs", b)
        } else if s, ok := v.(string); ok {
            d.Set("cp_logs", s == "true")
        }
    }
    if v, exists := syslog["send-to-mgmt"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("send_to_mgmt", b)
        } else if s, ok := v.(string); ok {
            d.Set("send_to_mgmt", s == "true")
        }
    }
    if v, exists := syslog["filename"]; exists {
        d.Set("filename", fmt.Sprintf("%v", v))
    }
    if v, exists := syslog["tls-configuration"]; exists {
        d.Set("tls_configuration", v)
    }
    if v, exists := syslog["forwarded-logs-files"]; exists {
        if rawList, ok := v.([]interface{}); ok {
            knownPaths := map[string]bool{}
            if cur, ok := d.GetOk("forwarded_logs_files"); ok {
                for _, item := range cur.([]interface{}) {
                    if m, ok := item.(map[string]interface{}); ok {
                        if p, ok := m["path"].(string); ok && p != "" {
                            knownPaths[p] = true
                        }
                    }
                }
            }
            filtered := make([]interface{}, 0, len(rawList))
            for _, item := range rawList {
                if m, ok := item.(map[string]interface{}); ok {
                    entry := map[string]interface{}{}
                    if val, ok := m["path"]; ok { entry["path"] = fmt.Sprintf("%v", val) }
                    if val, ok := m["tag"]; ok { entry["tag"] = fmt.Sprintf("%v", val) }
                    if path, ok := entry["path"].(string); ok {
                        if len(knownPaths) == 0 || knownPaths[path] {
                            filtered = append(filtered, entry)
                        }
                    }
                }
            }
            d.Set("forwarded_logs_files", filtered)
        }
    }
    if v, exists := syslog["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaSyslog(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("audit_log"); ok {
        payload["audit-log"] = v.(bool)
    }

    if v, ok := d.GetOkExists("cp_logs"); ok {
        payload["cp-logs"] = v.(bool)
    }

    if v, ok := d.GetOkExists("send_to_mgmt"); ok {
        payload["send-to-mgmt"] = v.(bool)
    }

    if v, ok := d.GetOk("filename"); ok {
        payload["filename"] = v.(string)
    }

    if v := d.Get("tls_configuration"); len(v.([]interface{})) > 0 {
        _ = v
        tlsconfigurationMap := make(map[string]interface{})
        if v, ok := d.GetOk("tls_configuration.0.ca_certification"); ok {
            tlsconfigurationMap["ca_certification"] = v.(string)
        }
        if v, ok := d.GetOk("tls_configuration.0.public_key"); ok {
            tlsconfigurationMap["public_key"] = v.(string)
        }
        if v, ok := d.GetOk("tls_configuration.0.private_key"); ok {
            tlsconfigurationMap["private_key"] = v.(string)
        }
        if len(tlsconfigurationMap) > 0 {
            payload["tls-configuration"] = tlsconfigurationMap
        }
    }

    if v := d.Get("forwarded_logs_files"); len(v.([]interface{})) > 0 {
        forwardedlogsfilesList := v.([]interface{})
        forwardedlogsfilesArray := make([]interface{}, 0, len(forwardedlogsfilesList))
        for i := range forwardedlogsfilesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("forwarded_logs_files.%d.path", i)); ok {
                itemMap["path"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("forwarded_logs_files.%d.tag", i)); ok {
                itemMap["tag"] = v.(string)
            }
            if len(itemMap) > 0 {
                forwardedlogsfilesArray = append(forwardedlogsfilesArray, itemMap)
            }
        }
        if len(forwardedlogsfilesArray) > 0 {
            payload["forwarded-logs-files"] = forwardedlogsfilesArray
        }
    }

    // Pre-clear "forwarded-logs-files": device appends entries rather than replacing them.
    // Send an empty list first so the subsequent update achieves replace semantics.
    clearPayload := map[string]interface{}{"forwarded-logs-files": []interface{}{}}
    if v, ok := d.GetOkExists("audit_log"); ok {
        clearPayload["audit-log"] = v.(bool)
    }
    if v, ok := d.GetOkExists("cp_logs"); ok {
        clearPayload["cp-logs"] = v.(bool)
    }
    if v, ok := d.GetOkExists("send_to_mgmt"); ok {
        clearPayload["send-to-mgmt"] = v.(bool)
    }
    clearRes, err := client.ApiCallSimple("set-syslog", clearPayload)
    if err != nil {
        return fmt.Errorf("Failed to clear forwarded-logs-files: %v", err)
    }
    if !clearRes.Success {
        return fmt.Errorf(clearRes.ErrorMsg)
    }

    setSyslogRes, err := client.ApiCallSimple("set-syslog", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setSyslogRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setSyslogRes.Success {
            errMsg = setSyslogRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setSyslogRes.GetData()
        }

        debugLogOperation(
            "syslog",        // resource type
            "update",                       // operation
            "set-syslog",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set syslog: %v", err)
    }
    if !setSyslogRes.Success {
        return fmt.Errorf(setSyslogRes.ErrorMsg)
    }

    return readGaiaSyslog(d, m)
}

func deleteGaiaSyslog(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    