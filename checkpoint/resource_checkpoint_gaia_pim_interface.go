package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaPimInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaPimInterface,
        Read:   readGaiaPimInterface,
        Update: updateGaiaPimInterface,
        Delete: deleteGaiaPimInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The interface name.`,
            },
            "dr_priority": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Used to determine the relative preference when electing a Designated Router (DR).`,
            },
            "enable_virtual_address": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Configures VRRP mode for the given interface.`,
            },
            "neighbor_filter": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure Neighbor Filter`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The multicast group prefix/mask, in CIDR notation.`,
                        },
                    },
                },
            },
            "ip_reachability_detection": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Configure BFD IP-Reachability Detection`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "dr_address": {
                Type:     schema.TypeString,
                Computed: true,
            },
            "mode": {
                Type:     schema.TypeString,
                Computed: true,
            },
            "neighbor_amount": {
                Type:     schema.TypeInt,
                Computed: true,
            },
            "state": {
                Type:     schema.TypeString,
                Computed: true,
            },
            "status": {
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}

func createGaiaPimInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("dr_priority"); ok {
        payload["dr-priority"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_virtual_address"); ok {
        payload["enable-virtual-address"] = v.(bool)
    }

    if v := d.Get("neighbor_filter"); len(v.([]interface{})) > 0 {
        neighborfilterList := v.([]interface{})
        neighborfilterArray := make([]interface{}, 0, len(neighborfilterList))
        for i := range neighborfilterList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("neighbor_filter.%d.address", i)); ok {
                itemMap["address"] = v.(string)
            }
            if len(itemMap) > 0 {
                neighborfilterArray = append(neighborfilterArray, itemMap)
            }
        }
        if len(neighborfilterArray) > 0 {
            payload["neighbor-filter"] = neighborfilterArray
        }
    }

    if v, ok := d.GetOkExists("ip_reachability_detection"); ok {
        payload["ip-reachability-detection"] = v.(bool)
    }

    log.Println("Create PimInterface - Map = ", payload)

    addPimInterfaceRes, err := client.ApiCallSimple("add-pim-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addPimInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addPimInterfaceRes.Success {
            errMsg = addPimInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addPimInterfaceRes.GetData()
        }

        debugLogOperation(
            "pim-interface",        // resource type
            "create",                       // operation
            "add-pim-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add pim-interface: %v", err)
    }
    if !addPimInterfaceRes.Success {
        if addPimInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addPimInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("pim-interface-" + acctest.RandString(10)))
    return readGaiaPimInterface(d, m)
}

func readGaiaPimInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showPimInterfaceRes, err := client.ApiCallSimple("show-pim-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showPimInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showPimInterfaceRes.Success {
            errMsg = showPimInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showPimInterfaceRes.GetData()
        }

        debugLogOperation(
            "pim-interface",        // resource type
            "read",                       // operation
            "show-pim-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show pim-interface: %v", err)
    }
    if !showPimInterfaceRes.Success {
        if data := showPimInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showPimInterfaceRes.ErrorMsg)
    }

    pimInterface := showPimInterfaceRes.GetData()

    log.Println("Read PimInterface - Show JSON = ", pimInterface)

    if v, exists := pimInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := pimInterface["dr-priority"]; exists {
        d.Set("dr_priority", fmt.Sprintf("%v", v))
    }
    if v, exists := pimInterface["dr-address"]; exists {
        d.Set("dr_address", fmt.Sprintf("%v", v))
    }
    if v, exists := pimInterface["mode"]; exists {
        d.Set("mode", fmt.Sprintf("%v", v))
    }
    if v, exists := pimInterface["neighbor-amount"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("neighbor_amount", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("neighbor_amount", _n)
            }
        }
    }
    if v, exists := pimInterface["state"]; exists {
        d.Set("state", fmt.Sprintf("%v", v))
    }
    if v, exists := pimInterface["status"]; exists {
        d.Set("status", fmt.Sprintf("%v", v))
    }
    if v, exists := pimInterface["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := pimInterface["enable-virtual-address"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enable_virtual_address", b)
        }
    }
    if v, exists := pimInterface["ip-reachability-detection"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ip_reachability_detection", b)
        }
    }
    if v, exists := pimInterface["neighbor-filter"]; exists {
        if nfList, ok := v.([]interface{}); ok {
            result := make([]interface{}, 0, len(nfList))
            for _, item := range nfList {
                if itemMap, ok := item.(map[string]interface{}); ok {
                    result = append(result, map[string]interface{}{
                        "address": itemMap["address"],
                    })
                }
            }
            d.Set("neighbor_filter", result)
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaPimInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("dr_priority"); ok {
        payload["dr-priority"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_virtual_address"); ok {
        payload["enable-virtual-address"] = v.(bool)
    }

    if v := d.Get("neighbor_filter"); len(v.([]interface{})) > 0 {
        neighborfilterList := v.([]interface{})
        neighborfilterArray := make([]interface{}, 0, len(neighborfilterList))
        for i := range neighborfilterList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("neighbor_filter.%d.address", i)); ok {
                itemMap["address"] = v.(string)
            }
            if len(itemMap) > 0 {
                neighborfilterArray = append(neighborfilterArray, itemMap)
            }
        }
        if len(neighborfilterArray) > 0 {
            payload["neighbor-filter"] = neighborfilterArray
        }
    }

    if v, ok := d.GetOkExists("ip_reachability_detection"); ok {
        payload["ip-reachability-detection"] = v.(bool)
    }

    setPimInterfaceRes, err := client.ApiCallSimple("set-pim-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setPimInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setPimInterfaceRes.Success {
            errMsg = setPimInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setPimInterfaceRes.GetData()
        }

        debugLogOperation(
            "pim-interface",        // resource type
            "update",                       // operation
            "set-pim-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set pim-interface: %v", err)
    }
    if !setPimInterfaceRes.Success {
        return fmt.Errorf(setPimInterfaceRes.ErrorMsg)
    }

    return readGaiaPimInterface(d, m)
}

func deleteGaiaPimInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deletePimInterfaceRes, err := client.ApiCallSimple("delete-pim-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deletePimInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deletePimInterfaceRes.Success {
            errMsg = deletePimInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deletePimInterfaceRes.GetData()
        }

        debugLogOperation(
            "pim-interface",        // resource type
            "delete",                       // operation
            "delete-pim-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete pim-interface: %v", err)
    }
    if !deletePimInterfaceRes.Success {
        return fmt.Errorf(deletePimInterfaceRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

