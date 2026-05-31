package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaSnmpPreDefinedTraps() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSnmpPreDefinedTraps,
        Read:   readGaiaSnmpPreDefinedTraps,
        Update: updateGaiaSnmpPreDefinedTraps,
        Delete: deleteGaiaSnmpPreDefinedTraps,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "authorizationerror": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `authorizationError Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "biosfailure": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `biosFailure Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "configurationchange": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `configurationChange Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "configurationsave": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `configurationSave Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "fanfailure": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `fanFailure Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "highvoltage": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `highVoltage Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "linkuplinkdown": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `linkUpLinkDown Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "clusterxlfailover": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `clusterXLFailover Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "lowvoltage": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `lowVoltage Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "overtemperature": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `overTemperature Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "powersupplyfailure": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `powerSupplyFailure Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "raidvolumestate": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `raidVolumeState Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "vrrpv2authfailure": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `vrrpv2AuthFailure Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "vrrpv2newmaster": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `vrrpv2NewMaster Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "vrrpv3newmaster": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `vrrpv3NewMaster Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "vrrpv3protoerror": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `vrrpv3ProtoError Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                    },
                },
            },
            "coldstart": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `ColdStart Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
                        },
                        "threshold": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `coldStart threshold (seconds), prevents sending coldStart trap when system up-time is greater than the threshold`,
                        },
                        "reboot_only": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `ColdStart reboot only, allows sending ColdStart trap only on reboot`,
                        },
                    },
                },
            },
            "lowdiskspaceallpartitions": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `lowDiskSpaceAllPartitions Trap`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Pre-defined trap state`,
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

func createGaiaSnmpPreDefinedTraps(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("authorizationerror"); len(v.([]interface{})) > 0 {
        _ = v
        authorizationErrorMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("authorizationerror.0.enabled"); ok && v.(bool) {
            authorizationErrorMap["enabled"] = v.(bool)
        }
        if len(authorizationErrorMap) > 0 {
            payload["authorizationError"] = authorizationErrorMap
        }
    }

    if v := d.Get("biosfailure"); len(v.([]interface{})) > 0 {
        _ = v
        biosFailureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("biosfailure.0.enabled"); ok && v.(bool) {
            biosFailureMap["enabled"] = v.(bool)
        }
        if len(biosFailureMap) > 0 {
            payload["biosFailure"] = biosFailureMap
        }
    }

    if v := d.Get("configurationchange"); len(v.([]interface{})) > 0 {
        _ = v
        configurationChangeMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("configurationchange.0.enabled"); ok && v.(bool) {
            configurationChangeMap["enabled"] = v.(bool)
        }
        if len(configurationChangeMap) > 0 {
            payload["configurationChange"] = configurationChangeMap
        }
    }

    if v := d.Get("configurationsave"); len(v.([]interface{})) > 0 {
        _ = v
        configurationSaveMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("configurationsave.0.enabled"); ok && v.(bool) {
            configurationSaveMap["enabled"] = v.(bool)
        }
        if len(configurationSaveMap) > 0 {
            payload["configurationSave"] = configurationSaveMap
        }
    }

    if v := d.Get("fanfailure"); len(v.([]interface{})) > 0 {
        _ = v
        fanFailureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("fanfailure.0.enabled"); ok && v.(bool) {
            fanFailureMap["enabled"] = v.(bool)
        }
        if len(fanFailureMap) > 0 {
            payload["fanFailure"] = fanFailureMap
        }
    }

    if v := d.Get("highvoltage"); len(v.([]interface{})) > 0 {
        _ = v
        highVoltageMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("highvoltage.0.enabled"); ok && v.(bool) {
            highVoltageMap["enabled"] = v.(bool)
        }
        if len(highVoltageMap) > 0 {
            payload["highVoltage"] = highVoltageMap
        }
    }

    if v := d.Get("linkuplinkdown"); len(v.([]interface{})) > 0 {
        _ = v
        linkUpLinkDownMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("linkuplinkdown.0.enabled"); ok && v.(bool) {
            linkUpLinkDownMap["enabled"] = v.(bool)
        }
        if len(linkUpLinkDownMap) > 0 {
            payload["linkUpLinkDown"] = linkUpLinkDownMap
        }
    }

    if v := d.Get("clusterxlfailover"); len(v.([]interface{})) > 0 {
        _ = v
        clusterXLFailoverMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("clusterxlfailover.0.enabled"); ok && v.(bool) {
            clusterXLFailoverMap["enabled"] = v.(bool)
        }
        if len(clusterXLFailoverMap) > 0 {
            payload["clusterXLFailover"] = clusterXLFailoverMap
        }
    }

    if v := d.Get("lowvoltage"); len(v.([]interface{})) > 0 {
        _ = v
        lowVoltageMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("lowvoltage.0.enabled"); ok && v.(bool) {
            lowVoltageMap["enabled"] = v.(bool)
        }
        if len(lowVoltageMap) > 0 {
            payload["lowVoltage"] = lowVoltageMap
        }
    }

    if v := d.Get("overtemperature"); len(v.([]interface{})) > 0 {
        _ = v
        overTemperatureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("overtemperature.0.enabled"); ok && v.(bool) {
            overTemperatureMap["enabled"] = v.(bool)
        }
        if len(overTemperatureMap) > 0 {
            payload["overTemperature"] = overTemperatureMap
        }
    }

    if v := d.Get("powersupplyfailure"); len(v.([]interface{})) > 0 {
        _ = v
        powerSupplyFailureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("powersupplyfailure.0.enabled"); ok && v.(bool) {
            powerSupplyFailureMap["enabled"] = v.(bool)
        }
        if len(powerSupplyFailureMap) > 0 {
            payload["powerSupplyFailure"] = powerSupplyFailureMap
        }
    }

    if v := d.Get("raidvolumestate"); len(v.([]interface{})) > 0 {
        _ = v
        raidVolumeStateMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("raidvolumestate.0.enabled"); ok && v.(bool) {
            raidVolumeStateMap["enabled"] = v.(bool)
        }
        if len(raidVolumeStateMap) > 0 {
            payload["raidVolumeState"] = raidVolumeStateMap
        }
    }

    if v := d.Get("vrrpv2authfailure"); len(v.([]interface{})) > 0 {
        _ = v
        vrrpv2AuthFailureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vrrpv2authfailure.0.enabled"); ok && v.(bool) {
            vrrpv2AuthFailureMap["enabled"] = v.(bool)
        }
        if len(vrrpv2AuthFailureMap) > 0 {
            payload["vrrpv2AuthFailure"] = vrrpv2AuthFailureMap
        }
    }

    if v := d.Get("vrrpv2newmaster"); len(v.([]interface{})) > 0 {
        _ = v
        vrrpv2NewMasterMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vrrpv2newmaster.0.enabled"); ok && v.(bool) {
            vrrpv2NewMasterMap["enabled"] = v.(bool)
        }
        if len(vrrpv2NewMasterMap) > 0 {
            payload["vrrpv2NewMaster"] = vrrpv2NewMasterMap
        }
    }

    if v := d.Get("vrrpv3newmaster"); len(v.([]interface{})) > 0 {
        _ = v
        vrrpv3NewMasterMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vrrpv3newmaster.0.enabled"); ok && v.(bool) {
            vrrpv3NewMasterMap["enabled"] = v.(bool)
        }
        if len(vrrpv3NewMasterMap) > 0 {
            payload["vrrpv3NewMaster"] = vrrpv3NewMasterMap
        }
    }

    if v := d.Get("vrrpv3protoerror"); len(v.([]interface{})) > 0 {
        _ = v
        vrrpv3ProtoErrorMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vrrpv3protoerror.0.enabled"); ok && v.(bool) {
            vrrpv3ProtoErrorMap["enabled"] = v.(bool)
        }
        if len(vrrpv3ProtoErrorMap) > 0 {
            payload["vrrpv3ProtoError"] = vrrpv3ProtoErrorMap
        }
    }

    if v := d.Get("coldstart"); len(v.([]interface{})) > 0 {
        _ = v
        coldStartMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("coldstart.0.enabled"); ok && v.(bool) {
            coldStartMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("coldstart.0.threshold"); ok {
            coldStartMap["threshold"] = v.(int)
        }
        if v, ok := d.GetOkExists("coldstart.0.reboot_only"); ok && v.(bool) {
            coldStartMap["reboot-only"] = v.(bool)
        }
        if len(coldStartMap) > 0 {
            payload["coldStart"] = coldStartMap
        }
    }

    if v := d.Get("lowdiskspaceallpartitions"); len(v.([]interface{})) > 0 {
        _ = v
        lowDiskSpaceAllPartitionsMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("lowdiskspaceallpartitions.0.enabled"); ok && v.(bool) {
            lowDiskSpaceAllPartitionsMap["enabled"] = v.(bool)
        }
        if len(lowDiskSpaceAllPartitionsMap) > 0 {
            payload["lowDiskSpaceAllPartitions"] = lowDiskSpaceAllPartitionsMap
        }
    }

    log.Println("Create SnmpPreDefinedTraps - Map = ", payload)

    addSnmpPreDefinedTrapsRes, err := client.ApiCallSimple("set-snmp-pre-defined-traps", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addSnmpPreDefinedTrapsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addSnmpPreDefinedTrapsRes.Success {
            errMsg = addSnmpPreDefinedTrapsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addSnmpPreDefinedTrapsRes.GetData()
        }

        debugLogOperation(
            "snmp-pre-defined-traps",        // resource type
            "create",                       // operation
            "set-snmp-pre-defined-traps",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add snmp-pre-defined-traps: %v", err)
    }
    if !addSnmpPreDefinedTrapsRes.Success {
        if addSnmpPreDefinedTrapsRes.ErrorMsg != "" {
            return fmt.Errorf(addSnmpPreDefinedTrapsRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("snmp-pre-defined-traps-" + acctest.RandString(10)))
    return readGaiaSnmpPreDefinedTraps(d, m)
}

func readGaiaSnmpPreDefinedTraps(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showSnmpPreDefinedTrapsRes, err := client.ApiCallSimple("show-snmp-pre-defined-traps", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showSnmpPreDefinedTrapsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showSnmpPreDefinedTrapsRes.Success {
            errMsg = showSnmpPreDefinedTrapsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showSnmpPreDefinedTrapsRes.GetData()
        }

        debugLogOperation(
            "snmp-pre-defined-traps",        // resource type
            "read",                       // operation
            "show-snmp-pre-defined-traps",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show snmp-pre-defined-traps: %v", err)
    }
    if !showSnmpPreDefinedTrapsRes.Success {
        if data := showSnmpPreDefinedTrapsRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showSnmpPreDefinedTrapsRes.ErrorMsg)
    }

    snmpPreDefinedTraps := showSnmpPreDefinedTrapsRes.GetData()

    log.Println("Read SnmpPreDefinedTraps - Show JSON = ", snmpPreDefinedTraps)

    if v, exists := snmpPreDefinedTraps["authorizationError"]; exists {
        d.Set("authorizationerror", v)
    }
    if v, exists := snmpPreDefinedTraps["biosFailure"]; exists {
        d.Set("biosfailure", v)
    }
    if v, exists := snmpPreDefinedTraps["configurationChange"]; exists {
        d.Set("configurationchange", v)
    }
    if v, exists := snmpPreDefinedTraps["configurationSave"]; exists {
        d.Set("configurationsave", v)
    }
    if v, exists := snmpPreDefinedTraps["fanFailure"]; exists {
        d.Set("fanfailure", v)
    }
    if v, exists := snmpPreDefinedTraps["highVoltage"]; exists {
        d.Set("highvoltage", v)
    }
    if v, exists := snmpPreDefinedTraps["linkUpLinkDown"]; exists {
        d.Set("linkuplinkdown", v)
    }
    if v, exists := snmpPreDefinedTraps["lowDiskSpace"]; exists {
        d.Set("lowdiskspace", v)
    }
    if v, exists := snmpPreDefinedTraps["lowVoltage"]; exists {
        d.Set("lowvoltage", v)
    }
    if v, exists := snmpPreDefinedTraps["overTemperature"]; exists {
        d.Set("overtemperature", v)
    }
    if v, exists := snmpPreDefinedTraps["powerSupplyFailure"]; exists {
        d.Set("powersupplyfailure", v)
    }
    if v, exists := snmpPreDefinedTraps["raidVolumeState"]; exists {
        d.Set("raidvolumestate", v)
    }
    if v, exists := snmpPreDefinedTraps["vrrpv2AuthFailure"]; exists {
        d.Set("vrrpv2authfailure", v)
    }
    if v, exists := snmpPreDefinedTraps["vrrpv2NewMaster"]; exists {
        d.Set("vrrpv2newmaster", v)
    }
    if v, exists := snmpPreDefinedTraps["vrrpv3NewMaster"]; exists {
        d.Set("vrrpv3newmaster", v)
    }
    if v, exists := snmpPreDefinedTraps["vrrpv3ProtoError"]; exists {
        d.Set("vrrpv3protoerror", v)
    }
    if v, exists := snmpPreDefinedTraps["coldStart"]; exists {
        d.Set("coldstart", v)
    }
    if v, exists := snmpPreDefinedTraps["clusterXLFailover"]; exists {
        d.Set("clusterxlfailover", v)
    }
    if v, exists := snmpPreDefinedTraps["lowDiskSpaceAllPartitions"]; exists {
        d.Set("lowdiskspaceallpartitions", v)
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaSnmpPreDefinedTraps(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("authorizationerror"); len(v.([]interface{})) > 0 {
        _ = v
        authorizationErrorMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("authorizationerror.0.enabled"); ok && v.(bool) {
            authorizationErrorMap["enabled"] = v.(bool)
        }
        if len(authorizationErrorMap) > 0 {
            payload["authorizationError"] = authorizationErrorMap
        }
    }

    if v := d.Get("biosfailure"); len(v.([]interface{})) > 0 {
        _ = v
        biosFailureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("biosfailure.0.enabled"); ok && v.(bool) {
            biosFailureMap["enabled"] = v.(bool)
        }
        if len(biosFailureMap) > 0 {
            payload["biosFailure"] = biosFailureMap
        }
    }

    if v := d.Get("configurationchange"); len(v.([]interface{})) > 0 {
        _ = v
        configurationChangeMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("configurationchange.0.enabled"); ok && v.(bool) {
            configurationChangeMap["enabled"] = v.(bool)
        }
        if len(configurationChangeMap) > 0 {
            payload["configurationChange"] = configurationChangeMap
        }
    }

    if v := d.Get("configurationsave"); len(v.([]interface{})) > 0 {
        _ = v
        configurationSaveMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("configurationsave.0.enabled"); ok && v.(bool) {
            configurationSaveMap["enabled"] = v.(bool)
        }
        if len(configurationSaveMap) > 0 {
            payload["configurationSave"] = configurationSaveMap
        }
    }

    if v := d.Get("fanfailure"); len(v.([]interface{})) > 0 {
        _ = v
        fanFailureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("fanfailure.0.enabled"); ok && v.(bool) {
            fanFailureMap["enabled"] = v.(bool)
        }
        if len(fanFailureMap) > 0 {
            payload["fanFailure"] = fanFailureMap
        }
    }

    if v := d.Get("highvoltage"); len(v.([]interface{})) > 0 {
        _ = v
        highVoltageMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("highvoltage.0.enabled"); ok && v.(bool) {
            highVoltageMap["enabled"] = v.(bool)
        }
        if len(highVoltageMap) > 0 {
            payload["highVoltage"] = highVoltageMap
        }
    }

    if v := d.Get("linkuplinkdown"); len(v.([]interface{})) > 0 {
        _ = v
        linkUpLinkDownMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("linkuplinkdown.0.enabled"); ok && v.(bool) {
            linkUpLinkDownMap["enabled"] = v.(bool)
        }
        if len(linkUpLinkDownMap) > 0 {
            payload["linkUpLinkDown"] = linkUpLinkDownMap
        }
    }

    if v := d.Get("clusterxlfailover"); len(v.([]interface{})) > 0 {
        _ = v
        clusterXLFailoverMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("clusterxlfailover.0.enabled"); ok && v.(bool) {
            clusterXLFailoverMap["enabled"] = v.(bool)
        }
        if len(clusterXLFailoverMap) > 0 {
            payload["clusterXLFailover"] = clusterXLFailoverMap
        }
    }

    if v := d.Get("lowvoltage"); len(v.([]interface{})) > 0 {
        _ = v
        lowVoltageMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("lowvoltage.0.enabled"); ok && v.(bool) {
            lowVoltageMap["enabled"] = v.(bool)
        }
        if len(lowVoltageMap) > 0 {
            payload["lowVoltage"] = lowVoltageMap
        }
    }

    if v := d.Get("overtemperature"); len(v.([]interface{})) > 0 {
        _ = v
        overTemperatureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("overtemperature.0.enabled"); ok && v.(bool) {
            overTemperatureMap["enabled"] = v.(bool)
        }
        if len(overTemperatureMap) > 0 {
            payload["overTemperature"] = overTemperatureMap
        }
    }

    if v := d.Get("powersupplyfailure"); len(v.([]interface{})) > 0 {
        _ = v
        powerSupplyFailureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("powersupplyfailure.0.enabled"); ok && v.(bool) {
            powerSupplyFailureMap["enabled"] = v.(bool)
        }
        if len(powerSupplyFailureMap) > 0 {
            payload["powerSupplyFailure"] = powerSupplyFailureMap
        }
    }

    if v := d.Get("raidvolumestate"); len(v.([]interface{})) > 0 {
        _ = v
        raidVolumeStateMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("raidvolumestate.0.enabled"); ok && v.(bool) {
            raidVolumeStateMap["enabled"] = v.(bool)
        }
        if len(raidVolumeStateMap) > 0 {
            payload["raidVolumeState"] = raidVolumeStateMap
        }
    }

    if v := d.Get("vrrpv2authfailure"); len(v.([]interface{})) > 0 {
        _ = v
        vrrpv2AuthFailureMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vrrpv2authfailure.0.enabled"); ok && v.(bool) {
            vrrpv2AuthFailureMap["enabled"] = v.(bool)
        }
        if len(vrrpv2AuthFailureMap) > 0 {
            payload["vrrpv2AuthFailure"] = vrrpv2AuthFailureMap
        }
    }

    if v := d.Get("vrrpv2newmaster"); len(v.([]interface{})) > 0 {
        _ = v
        vrrpv2NewMasterMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vrrpv2newmaster.0.enabled"); ok && v.(bool) {
            vrrpv2NewMasterMap["enabled"] = v.(bool)
        }
        if len(vrrpv2NewMasterMap) > 0 {
            payload["vrrpv2NewMaster"] = vrrpv2NewMasterMap
        }
    }

    if v := d.Get("vrrpv3newmaster"); len(v.([]interface{})) > 0 {
        _ = v
        vrrpv3NewMasterMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vrrpv3newmaster.0.enabled"); ok && v.(bool) {
            vrrpv3NewMasterMap["enabled"] = v.(bool)
        }
        if len(vrrpv3NewMasterMap) > 0 {
            payload["vrrpv3NewMaster"] = vrrpv3NewMasterMap
        }
    }

    if v := d.Get("vrrpv3protoerror"); len(v.([]interface{})) > 0 {
        _ = v
        vrrpv3ProtoErrorMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("vrrpv3protoerror.0.enabled"); ok && v.(bool) {
            vrrpv3ProtoErrorMap["enabled"] = v.(bool)
        }
        if len(vrrpv3ProtoErrorMap) > 0 {
            payload["vrrpv3ProtoError"] = vrrpv3ProtoErrorMap
        }
    }

    if v := d.Get("coldstart"); len(v.([]interface{})) > 0 {
        _ = v
        coldStartMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("coldstart.0.enabled"); ok && v.(bool) {
            coldStartMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("coldstart.0.threshold"); ok {
            coldStartMap["threshold"] = v.(int)
        }
        if v, ok := d.GetOkExists("coldstart.0.reboot_only"); ok && v.(bool) {
            coldStartMap["reboot-only"] = v.(bool)
        }
        if len(coldStartMap) > 0 {
            payload["coldStart"] = coldStartMap
        }
    }

    if v := d.Get("lowdiskspaceallpartitions"); len(v.([]interface{})) > 0 {
        _ = v
        lowDiskSpaceAllPartitionsMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("lowdiskspaceallpartitions.0.enabled"); ok && v.(bool) {
            lowDiskSpaceAllPartitionsMap["enabled"] = v.(bool)
        }
        if len(lowDiskSpaceAllPartitionsMap) > 0 {
            payload["lowDiskSpaceAllPartitions"] = lowDiskSpaceAllPartitionsMap
        }
    }

    setSnmpPreDefinedTrapsRes, err := client.ApiCallSimple("set-snmp-pre-defined-traps", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setSnmpPreDefinedTrapsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setSnmpPreDefinedTrapsRes.Success {
            errMsg = setSnmpPreDefinedTrapsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setSnmpPreDefinedTrapsRes.GetData()
        }

        debugLogOperation(
            "snmp-pre-defined-traps",        // resource type
            "update",                       // operation
            "set-snmp-pre-defined-traps",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set snmp-pre-defined-traps: %v", err)
    }
    if !setSnmpPreDefinedTrapsRes.Success {
        return fmt.Errorf(setSnmpPreDefinedTrapsRes.ErrorMsg)
    }

    return readGaiaSnmpPreDefinedTraps(d, m)
}

func deleteGaiaSnmpPreDefinedTraps(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    