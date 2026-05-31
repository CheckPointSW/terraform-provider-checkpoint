package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowSnmpOid() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowSnmpOid,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "oid": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `OID (object identifier)`,
            },
            "snmp_sid": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `SNMP session id returned from set-snmp-session`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "value": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "oid_next": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowSnmpOid(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("oid"); ok {
        payload["oid"] = v.(string)
    }

    if v, ok := d.GetOk("snmp_sid"); ok {
        payload["snmp-sid"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-snmp-oid - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-snmp-oid", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && commandRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !commandRes.Success {
            errMsg = commandRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = commandRes.GetData()
        }

        debugLogOperation(
            "snmp-oid",        // resource type
            "read",                       // operation
            "show-snmp-oid",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-snmp-oid: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["value"]; exists {
        d.Set("value", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["oid-next"]; exists {
        d.Set("oid_next", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-snmp-oid-" + acctest.RandString(10)))
    return nil
}

