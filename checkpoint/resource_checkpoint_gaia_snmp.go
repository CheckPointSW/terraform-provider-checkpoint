package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaSnmp() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSnmp,
        Read:   readGaiaSnmp,
        Update: updateGaiaSnmp,
        Delete: deleteGaiaSnmp,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Enables/Disables the SNMP Agent`,
            },
            "version": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configures the supported SNMP version:<br>all - support SNMP v1, v2 and v3<br>v3-Only - support SNMP v3 only`,
            },
            "trap_usm": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The user which will generate the SNMP traps, should be existed usm user`,
            },
            "contact": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `SNMP contact string`,
            },
            "location": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `SNMP location string: Specifies a string that contains the location for the device`,
            },
            "read_only_community": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `SNMP read-only community password, Where:<br>* read-only: lets you only read the values of SNMP objects`,
            },
            "read_write_community": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `SNMP read-write community password, Where:<br>* read-write: read and set the values as well`,
            },
            "interfaces": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `Adds a local interface to the list of local interfaces, on which the SNMP daemon listens`,
            },
            "pre_defined_traps_settings": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Pre-defined traps settings`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "polling_frequency": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Polling interval in seconds`,
                        },
                    },
                },
            },
            "custom_traps_settings": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Custom traps settings`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "clear_trap_interval": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Interval in second between clear traps`,
                        },
                        "clear_trap_amount": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Number of clear traps that is sent after custom trap termination`,
                        },
                    },
                },
            },
            "vsx_settings": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `VSX settings`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `True if SNMP is in vsx mode`,
                        },
                        "vs_access": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `SNMP vs-access type direct/indirect queries on Virtual-Devices<br>direct: SNMP direct queries on Virtual-Devices<br>indirect: SNMP direct queries via VS0`,
                        },
                        "sysname": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `This command is relevant only for VSX with SNMP VS mode, Where:<br>False (default) = the sysname OID for all Virtual Devices will return the same result: VS0 hostname<br>True =<br>* VS0 sysname OID returns the VSX hostname<br>* Virtual Device sysname OID returns the Check Point object name of the Virtual Device`,
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

func createGaiaSnmp(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("version"); ok {
        payload["version"] = v.(string)
    }

    if v, ok := d.GetOk("trap_usm"); ok {
        payload["trap-usm"] = v.(string)
    }

    if v, ok := d.GetOk("contact"); ok {
        payload["contact"] = v.(string)
    }

    if v, ok := d.GetOk("location"); ok {
        payload["location"] = v.(string)
    }

    if v, ok := d.GetOk("read_only_community"); ok {
        payload["read-only-community"] = v.(string)
    }

    if v, ok := d.GetOk("read_write_community"); ok {
        payload["read-write-community"] = v.(string)
    }

    if v, ok := d.GetOk("interfaces"); ok {
        payload["interfaces"] = v.(string)
    }

    if v := d.Get("pre_defined_traps_settings"); len(v.([]interface{})) > 0 {
        _ = v
        predefinedtrapssettingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("pre_defined_traps_settings.0.polling_frequency"); ok {
            predefinedtrapssettingsMap["polling-frequency"] = v.(int)
        }
        if len(predefinedtrapssettingsMap) > 0 {
            payload["pre-defined-traps-settings"] = predefinedtrapssettingsMap
        }
    }

    if v := d.Get("custom_traps_settings"); len(v.([]interface{})) > 0 {
        _ = v
        customtrapssettingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("custom_traps_settings.0.clear_trap_interval"); ok {
            customtrapssettingsMap["clear-trap-interval"] = v.(int)
        }
        if v, ok := d.GetOk("custom_traps_settings.0.clear_trap_amount"); ok {
            customtrapssettingsMap["clear-trap-amount"] = v.(int)
        }
        if len(customtrapssettingsMap) > 0 {
            payload["custom-traps-settings"] = customtrapssettingsMap
        }
    }

    if v := d.Get("vsx_settings"); len(v.([]interface{})) > 0 {
        _ = v
        vsxsettingsMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vsx_settings.0.enabled"); ok && v.(bool) {
            vsxsettingsMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("vsx_settings.0.vs_access"); ok {
            vsxsettingsMap["vs-access"] = v.(string)
        }
        if v, ok := d.GetOkExists("vsx_settings.0.sysname"); ok && v.(bool) {
            vsxsettingsMap["sysname"] = v.(bool)
        }
        if len(vsxsettingsMap) > 0 {
            payload["vsx-settings"] = vsxsettingsMap
        }
    }

    log.Println("Create Snmp - Map = ", payload)

    addSnmpRes, err := client.ApiCallSimple("set-snmp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addSnmpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addSnmpRes.Success {
            errMsg = addSnmpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addSnmpRes.GetData()
        }

        debugLogOperation(
            "snmp",        // resource type
            "create",                       // operation
            "set-snmp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add snmp: %v", err)
    }
    if !addSnmpRes.Success {
        if addSnmpRes.ErrorMsg != "" {
            return fmt.Errorf(addSnmpRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("snmp-" + acctest.RandString(10)))
    return readGaiaSnmp(d, m)
}

func readGaiaSnmp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showSnmpRes, err := client.ApiCallSimple("show-snmp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showSnmpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showSnmpRes.Success {
            errMsg = showSnmpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showSnmpRes.GetData()
        }

        debugLogOperation(
            "snmp",        // resource type
            "read",                       // operation
            "show-snmp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show snmp: %v", err)
    }
    if !showSnmpRes.Success {
        if data := showSnmpRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showSnmpRes.ErrorMsg)
    }

    snmp := showSnmpRes.GetData()

    log.Println("Read Snmp - Show JSON = ", snmp)

    if v, exists := snmp["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := snmp["version"]; exists {
        d.Set("version", fmt.Sprintf("%v", v))
    }
    if v, exists := snmp["trap-usm"]; exists {
        d.Set("trap_usm", fmt.Sprintf("%v", v))
    }
    if v, exists := snmp["contact"]; exists {
        d.Set("contact", fmt.Sprintf("%v", v))
    }
    if v, exists := snmp["location"]; exists {
        d.Set("location", fmt.Sprintf("%v", v))
    }
    if v, exists := snmp["read-only-community"]; exists {
        d.Set("read_only_community", fmt.Sprintf("%v", v))
    }
    if v, exists := snmp["read-write-community"]; exists {
        d.Set("read_write_community", fmt.Sprintf("%v", v))
    }
    if v, exists := snmp["interfaces"]; exists {
        d.Set("interfaces", fmt.Sprintf("%v", v))
    }
    if v, exists := snmp["pre-defined-traps-settings"]; exists {
        d.Set("pre_defined_traps_settings", v)
    }
    if v, exists := snmp["custom-traps-settings"]; exists {
        d.Set("custom_traps_settings", v)
    }
    if v, exists := snmp["vsx-settings"]; exists {
        d.Set("vsx_settings", v)
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaSnmp(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("version"); ok {
        payload["version"] = v.(string)
    }

    if v, ok := d.GetOk("trap_usm"); ok {
        payload["trap-usm"] = v.(string)
    }

    if v, ok := d.GetOk("contact"); ok {
        payload["contact"] = v.(string)
    }

    if v, ok := d.GetOk("location"); ok {
        payload["location"] = v.(string)
    }

    if v, ok := d.GetOk("read_only_community"); ok {
        payload["read-only-community"] = v.(string)
    }

    if v, ok := d.GetOk("read_write_community"); ok {
        payload["read-write-community"] = v.(string)
    }

    if v, ok := d.GetOk("interfaces"); ok {
        payload["interfaces"] = v.(string)
    }

    if v := d.Get("pre_defined_traps_settings"); len(v.([]interface{})) > 0 {
        _ = v
        predefinedtrapssettingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("pre_defined_traps_settings.0.polling_frequency"); ok {
            predefinedtrapssettingsMap["polling-frequency"] = v.(int)
        }
        if len(predefinedtrapssettingsMap) > 0 {
            payload["pre-defined-traps-settings"] = predefinedtrapssettingsMap
        }
    }

    if v := d.Get("custom_traps_settings"); len(v.([]interface{})) > 0 {
        _ = v
        customtrapssettingsMap := make(map[string]interface{})
        if v, ok := d.GetOk("custom_traps_settings.0.clear_trap_interval"); ok {
            customtrapssettingsMap["clear-trap-interval"] = v.(int)
        }
        if v, ok := d.GetOk("custom_traps_settings.0.clear_trap_amount"); ok {
            customtrapssettingsMap["clear-trap-amount"] = v.(int)
        }
        if len(customtrapssettingsMap) > 0 {
            payload["custom-traps-settings"] = customtrapssettingsMap
        }
    }

    if v := d.Get("vsx_settings"); len(v.([]interface{})) > 0 {
        _ = v
        vsxsettingsMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vsx_settings.0.enabled"); ok && v.(bool) {
            vsxsettingsMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("vsx_settings.0.vs_access"); ok {
            vsxsettingsMap["vs-access"] = v.(string)
        }
        if v, ok := d.GetOkExists("vsx_settings.0.sysname"); ok && v.(bool) {
            vsxsettingsMap["sysname"] = v.(bool)
        }
        if len(vsxsettingsMap) > 0 {
            payload["vsx-settings"] = vsxsettingsMap
        }
    }

    setSnmpRes, err := client.ApiCallSimple("set-snmp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setSnmpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setSnmpRes.Success {
            errMsg = setSnmpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setSnmpRes.GetData()
        }

        debugLogOperation(
            "snmp",        // resource type
            "update",                       // operation
            "set-snmp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set snmp: %v", err)
    }
    if !setSnmpRes.Success {
        return fmt.Errorf(setSnmpRes.ErrorMsg)
    }

    return readGaiaSnmp(d, m)
}

func deleteGaiaSnmp(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    