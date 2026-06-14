package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSetSnmpSession() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetSnmpSession,
        Read:   readGaiaSetSnmpSession,
        Delete: deleteGaiaSetSnmpSession,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "community_string": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `SNMP v2 community password.<br>                 <b>required for SNMP v1/v2</b>`,
            },
            "v3_object": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `SNMPv3 USM (User-based Security Model) details<br>                       <b>required for SNMP v3</b><br>                       <b>Preferred</b>`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `SNMPv3 USM user`,
                        },
                        "authentication": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Authentication details`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "protocol": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Authentication protocol, MD5 and SHA1 are not supported starting from R81`,
                                    },
                                    "password": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Sensitive:   true,
                                        Description: `Authentication Password - (8 or more printable characters)<br>Each SNMPv3 USM user must have an authentication pass phrase.<br>This will be used by the SNMPv3 agent to verify the identity of the user before granting access.`,
                                    },
                                },
                            },
                        },
                        "data_privacy": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Related to AutoPriv/AutnNoPriv in SecurityLevel in the RFC.<br>True: AutoPriv<br>False: AuthNoPriv`,
                        },
                    },
                },
            },
            "session_timeout": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `Session expiration timeout in seconds`,
            },
            "snmp_sid": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaSetSnmpSession(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("community_string"); ok {
        payload["community-string"] = v.(string)
    }

    if v := d.Get("v3_object"); len(v.([]interface{})) > 0 {
        _ = v
        v3objectMap := make(map[string]interface{})
        if v, ok := d.GetOk("v3_object.0.name"); ok {
            v3objectMap["name"] = v.(string)
        }
        if v, ok := d.GetOk("v3_object.0.authentication"); ok {
            _ = v
            authenticationMap := make(map[string]interface{})
            if v, ok := d.GetOk("v3_object.0.authentication.0.protocol"); ok {
                authenticationMap["protocol"] = v.(string)
            }
            if v, ok := d.GetOk("v3_object.0.authentication.0.password"); ok {
                authenticationMap["password"] = v.(string)
            }
            if len(authenticationMap) > 0 {
                v3objectMap["authentication"] = authenticationMap
            }
        }
        if v, ok := d.GetOkExists("v3_object.0.data_privacy"); ok && v.(bool) {
            v3objectMap["data-privacy"] = v.(bool)
        }
        if len(v3objectMap) > 0 {
            payload["v3-object"] = v3objectMap
        }
    }

    if v, ok := d.GetOk("session_timeout"); ok {
        payload["session-timeout"] = v.(int)
    }

    log.Println("Execute set-snmp-session - Payload = ", payload)

    GaiaSetSnmpSessionRes, err := client.ApiCallSimple("set-snmp-session", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetSnmpSessionRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetSnmpSessionRes.Success {
            errMsg = GaiaSetSnmpSessionRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetSnmpSessionRes.GetData()
        }

        debugLogOperation(
            "set-snmp-session",        // resource type
            "command",                       // operation
            "set-snmp-session",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-snmp-session: %v", err)
    }
    if !GaiaSetSnmpSessionRes.Success {
        if GaiaSetSnmpSessionRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetSnmpSessionRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaSetSnmpSessionRes.GetData()
    if v, exists := _respData["snmp-sid"]; exists {
        d.Set("snmp_sid", toString(v))
    }


    d.SetId(fmt.Sprintf("set-snmp-session-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetSnmpSession(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSetSnmpSession(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

