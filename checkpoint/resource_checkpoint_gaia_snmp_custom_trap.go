package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaSnmpCustomTrap() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSnmpCustomTrap,
        Read:   readGaiaSnmpCustomTrap,
        Update: updateGaiaSnmpCustomTrap,
        Delete: deleteGaiaSnmpCustomTrap,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Custom trap name`,
            },
            "oid": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `OID (object identifier)`,
            },
            "operator": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Comparison operator`,
            },
            "threshold": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The value you want to compare to`,
            },
            "frequency": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `Polling interval in seconds`,
            },
            "message": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Custom trap message`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaSnmpCustomTrap(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("oid"); ok {
        payload["oid"] = v.(string)
    }

    if v, ok := d.GetOk("operator"); ok {
        payload["operator"] = v.(string)
    }

    if v, ok := d.GetOk("threshold"); ok {
        payload["threshold"] = v.(string)
    }

    if v, ok := d.GetOk("frequency"); ok {
        payload["frequency"] = v.(int)
    }

    if v, ok := d.GetOk("message"); ok {
        payload["message"] = v.(string)
    }

    log.Println("Create SnmpCustomTrap - Map = ", payload)

    addSnmpCustomTrapRes, err := client.ApiCallSimple("add-snmp-custom-trap", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addSnmpCustomTrapRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addSnmpCustomTrapRes.Success {
            errMsg = addSnmpCustomTrapRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addSnmpCustomTrapRes.GetData()
        }

        debugLogOperation(
            "snmp-custom-trap",        // resource type
            "create",                       // operation
            "add-snmp-custom-trap",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add snmp-custom-trap: %v", err)
    }
    if !addSnmpCustomTrapRes.Success {
        if addSnmpCustomTrapRes.ErrorMsg != "" {
            return fmt.Errorf(addSnmpCustomTrapRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("snmp-custom-trap-" + acctest.RandString(10)))
    return readGaiaSnmpCustomTrap(d, m)
}

func readGaiaSnmpCustomTrap(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showSnmpCustomTrapRes, err := client.ApiCallSimple("show-snmp-custom-trap", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showSnmpCustomTrapRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showSnmpCustomTrapRes.Success {
            errMsg = showSnmpCustomTrapRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showSnmpCustomTrapRes.GetData()
        }

        debugLogOperation(
            "snmp-custom-trap",        // resource type
            "read",                       // operation
            "show-snmp-custom-trap",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show snmp-custom-trap: %v", err)
    }
    if !showSnmpCustomTrapRes.Success {
        if data := showSnmpCustomTrapRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showSnmpCustomTrapRes.ErrorMsg)
    }

    snmpCustomTrap := showSnmpCustomTrapRes.GetData()

    log.Println("Read SnmpCustomTrap - Show JSON = ", snmpCustomTrap)

    if v, exists := snmpCustomTrap["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := snmpCustomTrap["oid"]; exists {
        d.Set("oid", fmt.Sprintf("%v", v))
    }
    if v, exists := snmpCustomTrap["operator"]; exists {
        d.Set("operator", fmt.Sprintf("%v", v))
    }
    if v, exists := snmpCustomTrap["threshold"]; exists {
        d.Set("threshold", fmt.Sprintf("%v", v))
    }
    if v, exists := snmpCustomTrap["frequency"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("frequency", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("frequency", _n)
            }
        }
    }
    if v, exists := snmpCustomTrap["message"]; exists {
        d.Set("message", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaSnmpCustomTrap(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("oid"); ok {
        payload["oid"] = v.(string)
    }

    if v, ok := d.GetOk("operator"); ok {
        payload["operator"] = v.(string)
    }

    if v, ok := d.GetOk("threshold"); ok {
        payload["threshold"] = v.(string)
    }

    if v, ok := d.GetOk("frequency"); ok {
        payload["frequency"] = v.(int)
    }

    if v, ok := d.GetOk("message"); ok {
        payload["message"] = v.(string)
    }

    setSnmpCustomTrapRes, err := client.ApiCallSimple("set-snmp-custom-trap", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setSnmpCustomTrapRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setSnmpCustomTrapRes.Success {
            errMsg = setSnmpCustomTrapRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setSnmpCustomTrapRes.GetData()
        }

        debugLogOperation(
            "snmp-custom-trap",        // resource type
            "update",                       // operation
            "set-snmp-custom-trap",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set snmp-custom-trap: %v", err)
    }
    if !setSnmpCustomTrapRes.Success {
        return fmt.Errorf(setSnmpCustomTrapRes.ErrorMsg)
    }

    return readGaiaSnmpCustomTrap(d, m)
}

func deleteGaiaSnmpCustomTrap(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteSnmpCustomTrapRes, err := client.ApiCallSimple("delete-snmp-custom-trap", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteSnmpCustomTrapRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteSnmpCustomTrapRes.Success {
            errMsg = deleteSnmpCustomTrapRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteSnmpCustomTrapRes.GetData()
        }

        debugLogOperation(
            "snmp-custom-trap",        // resource type
            "delete",                       // operation
            "delete-snmp-custom-trap",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete snmp-custom-trap: %v", err)
    }
    if !deleteSnmpCustomTrapRes.Success {
        return fmt.Errorf(deleteSnmpCustomTrapRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

