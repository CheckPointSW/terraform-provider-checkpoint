package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaSnmpTrapReceiver() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSnmpTrapReceiver,
        Read:   readGaiaSnmpTrapReceiver,
        Update: updateGaiaSnmpTrapReceiver,
        Delete: deleteGaiaSnmpTrapReceiver,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "address": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Receiver address`,
            },
            "version": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Receiver version`,
            },
            "community_string": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Receiver community - Required only in case of v1/v2 versions<br>Trap Community String used by the trap receiver to determine which traps are accepted from a device.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaSnmpTrapReceiver(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("version"); ok {
        payload["version"] = v.(string)
    }

    if v, ok := d.GetOk("community_string"); ok {
        payload["community-string"] = v.(string)
    }

    log.Println("Create SnmpTrapReceiver - Map = ", payload)

    addSnmpTrapReceiverRes, err := client.ApiCallSimple("add-snmp-trap-receiver", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addSnmpTrapReceiverRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addSnmpTrapReceiverRes.Success {
            errMsg = addSnmpTrapReceiverRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addSnmpTrapReceiverRes.GetData()
        }

        debugLogOperation(
            "snmp-trap-receiver",        // resource type
            "create",                       // operation
            "add-snmp-trap-receiver",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add snmp-trap-receiver: %v", err)
    }
    if !addSnmpTrapReceiverRes.Success {
        if addSnmpTrapReceiverRes.ErrorMsg != "" {
            return fmt.Errorf(addSnmpTrapReceiverRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("snmp-trap-receiver-" + acctest.RandString(10)))
    return readGaiaSnmpTrapReceiver(d, m)
}

func readGaiaSnmpTrapReceiver(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showSnmpTrapReceiverRes, err := client.ApiCallSimple("show-snmp-trap-receiver", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showSnmpTrapReceiverRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showSnmpTrapReceiverRes.Success {
            errMsg = showSnmpTrapReceiverRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showSnmpTrapReceiverRes.GetData()
        }

        debugLogOperation(
            "snmp-trap-receiver",        // resource type
            "read",                       // operation
            "show-snmp-trap-receiver",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show snmp-trap-receiver: %v", err)
    }
    if !showSnmpTrapReceiverRes.Success {
        if data := showSnmpTrapReceiverRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showSnmpTrapReceiverRes.ErrorMsg)
    }

    snmpTrapReceiver := showSnmpTrapReceiverRes.GetData()

    log.Println("Read SnmpTrapReceiver - Show JSON = ", snmpTrapReceiver)

    if v, exists := snmpTrapReceiver["address"]; exists {
        d.Set("address", fmt.Sprintf("%v", v))
    }
    if v, exists := snmpTrapReceiver["version"]; exists {
        d.Set("version", fmt.Sprintf("%v", v))
    }
    if v, exists := snmpTrapReceiver["community-string"]; exists {
        d.Set("community_string", fmt.Sprintf("%v", v))
    }
    if v, exists := snmpTrapReceiver["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaSnmpTrapReceiver(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    if v, ok := d.GetOk("version"); ok {
        payload["version"] = v.(string)
    }

    if v, ok := d.GetOk("community_string"); ok {
        payload["community-string"] = v.(string)
    }

    setSnmpTrapReceiverRes, err := client.ApiCallSimple("set-snmp-trap-receiver", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setSnmpTrapReceiverRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setSnmpTrapReceiverRes.Success {
            errMsg = setSnmpTrapReceiverRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setSnmpTrapReceiverRes.GetData()
        }

        debugLogOperation(
            "snmp-trap-receiver",        // resource type
            "update",                       // operation
            "set-snmp-trap-receiver",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set snmp-trap-receiver: %v", err)
    }
    if !setSnmpTrapReceiverRes.Success {
        return fmt.Errorf(setSnmpTrapReceiverRes.ErrorMsg)
    }

    return readGaiaSnmpTrapReceiver(d, m)
}

func deleteGaiaSnmpTrapReceiver(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("address"); ok {
        payload["address"] = v.(string)
    }

    deleteSnmpTrapReceiverRes, err := client.ApiCallSimple("delete-snmp-trap-receiver", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteSnmpTrapReceiverRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteSnmpTrapReceiverRes.Success {
            errMsg = deleteSnmpTrapReceiverRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteSnmpTrapReceiverRes.GetData()
        }

        debugLogOperation(
            "snmp-trap-receiver",        // resource type
            "delete",                       // operation
            "delete-snmp-trap-receiver",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete snmp-trap-receiver: %v", err)
    }
    if !deleteSnmpTrapReceiverRes.Success {
        return fmt.Errorf(deleteSnmpTrapReceiverRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

