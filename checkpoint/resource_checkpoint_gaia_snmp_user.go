package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaSnmpUser() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSnmpUser,
        Read:   readGaiaSnmpUser,
        Update: updateGaiaSnmpUser,
        Delete: deleteGaiaSnmpUser,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `SNMPv3 USM user`,
            },
            "permission": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `User permission`,
            },
            "allowed_virtual_systems": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `Configured Virtual Devices allowed for the USM user - vsid range: 0-512`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "authentication": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `Authentication details`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "protocol": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Authentication protocol, MD5 and SHA1 are not supported starting from R81`,
                        },
                        "password": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `Authentication Password - (8 or more printable characters)<br>Each SNMPv3 USM user must have an authentication pass phrase.<br>This will be used by the SNMPv3 agent to verify the identity of the user before granting access.`,
                        },
                    },
                },
            },
            "privacy": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Privacy details. If provided, data privacy (encryption) is enabled`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "protocol": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Privacy protocol`,
                        },
                        "password": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `Privacy Password - (8 or more printable characters)<br>An SNMPv3 USM user with a privacy security level must have a privacy pass phrase.<br>This will be used by the SNMPv3 agent to keep other parties from eavesdropping on the SNMP interaction.<br>`,
                        },
                    },
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "data_privacy": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaSnmpUser(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("permission"); ok {
        payload["permission"] = v.(string)
    }

    if v := d.Get("allowed_virtual_systems"); len(v.(*schema.Set).List()) > 0 {
        payload["allowed-virtual-systems"] = v.(*schema.Set).List()
    }

    if v := d.Get("authentication"); len(v.([]interface{})) > 0 {
        _ = v
        authenticationMap := make(map[string]interface{})
        if v, ok := d.GetOk("authentication.0.protocol"); ok {
            authenticationMap["protocol"] = v.(string)
        }
        if v, ok := d.GetOk("authentication.0.password"); ok {
            authenticationMap["password"] = v.(string)
        }
        if len(authenticationMap) > 0 {
            payload["authentication"] = authenticationMap
        }
    }

    if v := d.Get("privacy"); len(v.([]interface{})) > 0 {
        _ = v
        privacyMap := make(map[string]interface{})
        if v, ok := d.GetOk("privacy.0.protocol"); ok {
            privacyMap["protocol"] = v.(string)
        }
        if v, ok := d.GetOk("privacy.0.password"); ok {
            privacyMap["password"] = v.(string)
        }
        if len(privacyMap) > 0 {
            payload["privacy"] = privacyMap
        }
    }

    log.Println("Create SnmpUser - Map = ", payload)

    addSnmpUserRes, err := client.ApiCallSimple("add-snmp-user", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addSnmpUserRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addSnmpUserRes.Success {
            errMsg = addSnmpUserRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addSnmpUserRes.GetData()
        }

        debugLogOperation(
            "snmp-user",        // resource type
            "create",                       // operation
            "add-snmp-user",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add snmp-user: %v", err)
    }
    if !addSnmpUserRes.Success {
        if addSnmpUserRes.ErrorMsg != "" {
            return fmt.Errorf(addSnmpUserRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("snmp-user-" + acctest.RandString(10)))
    return readGaiaSnmpUser(d, m)
}

func readGaiaSnmpUser(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showSnmpUserRes, err := client.ApiCallSimple("show-snmp-user", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showSnmpUserRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showSnmpUserRes.Success {
            errMsg = showSnmpUserRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showSnmpUserRes.GetData()
        }

        debugLogOperation(
            "snmp-user",        // resource type
            "read",                       // operation
            "show-snmp-user",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show snmp-user: %v", err)
    }
    if !showSnmpUserRes.Success {
        if data := showSnmpUserRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showSnmpUserRes.ErrorMsg)
    }

    snmpUser := showSnmpUserRes.GetData()

    log.Println("Read SnmpUser - Show JSON = ", snmpUser)

    if v, exists := snmpUser["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := snmpUser["permission"]; exists {
        d.Set("permission", fmt.Sprintf("%v", v))
    }
    if v, exists := snmpUser["allowed-virtual-systems"]; exists {
        switch val := v.(type) {
        case []interface{}:
            d.Set("allowed_virtual_systems", val)
        case string:
            if val != "" && val != "None" {
                d.Set("allowed_virtual_systems", []interface{}{val})
            }
        }
    }
    if v, exists := snmpUser["authentication"]; exists {
        if am, ok := v.(map[string]interface{}); ok {
            d.Set("authentication", []interface{}{map[string]interface{}{
                "protocol": fmt.Sprintf("%v", am["protocol"]),
                "password": d.Get("authentication.0.password"),
            }})
        }
    }
    if v, exists := snmpUser["privacy"]; exists {
        if pm, ok := v.(map[string]interface{}); ok {
            if proto := fmt.Sprintf("%v", pm["protocol"]); proto != "n/a" {
                d.Set("privacy", []interface{}{map[string]interface{}{
                    "protocol": proto,
                    "password": d.Get("privacy.0.password"),
                }})
            }
        }
    }
    if v, exists := snmpUser["data-privacy"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("data_privacy", b)
        } else if s, ok := v.(string); ok {
            d.Set("data_privacy", s == "true")
        }
    }
    if v, exists := snmpUser["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaSnmpUser(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("permission"); ok {
        payload["permission"] = v.(string)
    }

    if v := d.Get("allowed_virtual_systems"); len(v.(*schema.Set).List()) > 0 {
        payload["allowed-virtual-systems"] = v.(*schema.Set).List()
    }

    if v := d.Get("authentication"); len(v.([]interface{})) > 0 {
        _ = v
        authenticationMap := make(map[string]interface{})
        if v, ok := d.GetOk("authentication.0.protocol"); ok {
            authenticationMap["protocol"] = v.(string)
        }
        if v, ok := d.GetOk("authentication.0.password"); ok {
            authenticationMap["password"] = v.(string)
        }
        if len(authenticationMap) > 0 {
            payload["authentication"] = authenticationMap
        }
    }

    if v := d.Get("privacy"); len(v.([]interface{})) > 0 {
        _ = v
        privacyMap := make(map[string]interface{})
        if v, ok := d.GetOk("privacy.0.protocol"); ok {
            privacyMap["protocol"] = v.(string)
        }
        if v, ok := d.GetOk("privacy.0.password"); ok {
            privacyMap["password"] = v.(string)
        }
        if len(privacyMap) > 0 {
            payload["privacy"] = privacyMap
        }
    }

    setSnmpUserRes, err := client.ApiCallSimple("set-snmp-user", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setSnmpUserRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setSnmpUserRes.Success {
            errMsg = setSnmpUserRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setSnmpUserRes.GetData()
        }

        debugLogOperation(
            "snmp-user",        // resource type
            "update",                       // operation
            "set-snmp-user",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set snmp-user: %v", err)
    }
    if !setSnmpUserRes.Success {
        return fmt.Errorf(setSnmpUserRes.ErrorMsg)
    }

    return readGaiaSnmpUser(d, m)
}

func deleteGaiaSnmpUser(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteSnmpUserRes, err := client.ApiCallSimple("delete-snmp-user", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteSnmpUserRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteSnmpUserRes.Success {
            errMsg = deleteSnmpUserRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteSnmpUserRes.GetData()
        }

        debugLogOperation(
            "snmp-user",        // resource type
            "delete",                       // operation
            "delete-snmp-user",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete snmp-user: %v", err)
    }
    if !deleteSnmpUserRes.Success {
        return fmt.Errorf(deleteSnmpUserRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

