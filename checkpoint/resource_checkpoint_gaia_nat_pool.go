package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaNatPool() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaNatPool,
        Read:   readGaiaNatPool,
        Update: updateGaiaNatPool,
        Delete: deleteGaiaNatPool,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "prefix": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Specifies the IPv4 or IPv6 destination prefix of a NAT pool to be configured.<br><br>Note: A prefix cannot be of type IPv6, if IPv6 capabilities are not enabled`,
            },
            "comment": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Specifies a comment on a NAT pool. If the empty string is given, no comments will be added to the NAT pool.<br><br>Note: The length of the comment cannot exceed 100 characters`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaNatPool(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("prefix"); ok {
        payload["prefix"] = v.(string)
    }

    if v, ok := d.GetOk("comment"); ok {
        payload["comment"] = v.(string)
    }

    log.Println("Create NatPool - Map = ", payload)

    addNatPoolRes, err := client.ApiCallSimple("add-nat-pool", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addNatPoolRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addNatPoolRes.Success {
            errMsg = addNatPoolRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addNatPoolRes.GetData()
        }

        debugLogOperation(
            "nat-pool",        // resource type
            "create",                       // operation
            "add-nat-pool",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add nat-pool: %v", err)
    }
    if !addNatPoolRes.Success {
        if addNatPoolRes.ErrorMsg != "" {
            return fmt.Errorf(addNatPoolRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("nat-pool-" + acctest.RandString(10)))
    return readGaiaNatPool(d, m)
}

func readGaiaNatPool(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("prefix"); ok {
        payload["prefix"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showNatPoolRes, err := client.ApiCallSimple("show-nat-pool", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showNatPoolRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showNatPoolRes.Success {
            errMsg = showNatPoolRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showNatPoolRes.GetData()
        }

        debugLogOperation(
            "nat-pool",        // resource type
            "read",                       // operation
            "show-nat-pool",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show nat-pool: %v", err)
    }
    if !showNatPoolRes.Success {
        if data := showNatPoolRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showNatPoolRes.ErrorMsg)
    }

    natPool := showNatPoolRes.GetData()

    log.Println("Read NatPool - Show JSON = ", natPool)

    if v, exists := natPool["prefix"]; exists {
        d.Set("prefix", fmt.Sprintf("%v", v))
    }
    if v, exists := natPool["comment"]; exists {
        d.Set("comment", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaNatPool(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("prefix"); ok {
        payload["prefix"] = v.(string)
    }

    if v, ok := d.GetOk("comment"); ok {
        payload["comment"] = v.(string)
    }

    setNatPoolRes, err := client.ApiCallSimple("set-nat-pool", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setNatPoolRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setNatPoolRes.Success {
            errMsg = setNatPoolRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setNatPoolRes.GetData()
        }

        debugLogOperation(
            "nat-pool",        // resource type
            "update",                       // operation
            "set-nat-pool",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set nat-pool: %v", err)
    }
    if !setNatPoolRes.Success {
        return fmt.Errorf(setNatPoolRes.ErrorMsg)
    }

    return readGaiaNatPool(d, m)
}

func deleteGaiaNatPool(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("prefix"); ok {
        payload["prefix"] = v.(string)
    }

    deleteNatPoolRes, err := client.ApiCallSimple("delete-nat-pool", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteNatPoolRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteNatPoolRes.Success {
            errMsg = deleteNatPoolRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteNatPoolRes.GetData()
        }

        debugLogOperation(
            "nat-pool",        // resource type
            "delete",                       // operation
            "delete-nat-pool",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete nat-pool: %v", err)
    }
    if !deleteNatPoolRes.Success {
        return fmt.Errorf(deleteNatPoolRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

