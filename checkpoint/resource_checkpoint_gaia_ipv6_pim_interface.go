package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaIpv6PimInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaIpv6PimInterface,
        Read:   readGaiaIpv6PimInterface,
        Update: updateGaiaIpv6PimInterface,
        Delete: deleteGaiaIpv6PimInterface,
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
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "ip_reachability_detection": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "neighbor_filter": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
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

func createGaiaIpv6PimInterface(d *schema.ResourceData, m interface{}) error {
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

    log.Println("Create Ipv6PimInterface - Map = ", payload)

    addIpv6PimInterfaceRes, err := client.ApiCallSimple("add-ipv6-pim-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addIpv6PimInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addIpv6PimInterfaceRes.Success {
            errMsg = addIpv6PimInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addIpv6PimInterfaceRes.GetData()
        }

        debugLogOperation(
            "ipv6-pim-interface",        // resource type
            "create",                       // operation
            "add-ipv6-pim-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add ipv6-pim-interface: %v", err)
    }
    if !addIpv6PimInterfaceRes.Success {
        if addIpv6PimInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addIpv6PimInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    // Extract API-assigned fields from Create response before calling Read.
    if data := addIpv6PimInterfaceRes.GetData(); data != nil {
        if v, exists := data["ip-reachability-detection"]; exists {
            d.Set("ip_reachability_detection", v)
        }
        if v, exists := data["neighbor-filter"]; exists {
            d.Set("neighbor_filter", v)
        }
    }

    d.SetId(fmt.Sprintf("ipv6-pim-interface-" + acctest.RandString(10)))
    return readGaiaIpv6PimInterface(d, m)
}

func readGaiaIpv6PimInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showIpv6PimInterfaceRes, err := client.ApiCallSimple("show-ipv6-pim-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showIpv6PimInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showIpv6PimInterfaceRes.Success {
            errMsg = showIpv6PimInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showIpv6PimInterfaceRes.GetData()
        }

        debugLogOperation(
            "ipv6-pim-interface",        // resource type
            "read",                       // operation
            "show-ipv6-pim-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show ipv6-pim-interface: %v", err)
    }
    if !showIpv6PimInterfaceRes.Success {
        if data := showIpv6PimInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showIpv6PimInterfaceRes.ErrorMsg)
    }

    ipv6PimInterface := showIpv6PimInterfaceRes.GetData()

    log.Println("Read Ipv6PimInterface - Show JSON = ", ipv6PimInterface)

    if v, exists := ipv6PimInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := ipv6PimInterface["dr-priority"]; exists {
        d.Set("dr_priority", fmt.Sprintf("%v", v))
    }
    if v, exists := ipv6PimInterface["dr-address"]; exists {
        d.Set("dr_address", fmt.Sprintf("%v", v))
    }
    if v, exists := ipv6PimInterface["mode"]; exists {
        d.Set("mode", fmt.Sprintf("%v", v))
    }
    if v, exists := ipv6PimInterface["neighbor-amount"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("neighbor_amount", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("neighbor_amount", _n)
            }
        }
    }
    if v, exists := ipv6PimInterface["state"]; exists {
        d.Set("state", fmt.Sprintf("%v", v))
    }
    if v, exists := ipv6PimInterface["status"]; exists {
        d.Set("status", fmt.Sprintf("%v", v))
    }
    if v, exists := ipv6PimInterface["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    // enable-virtual-address is not returned by show-ipv6-pim-interface;
    // preserve the value configured by the user to avoid a perpetual diff.
    d.Set("enable_virtual_address", d.Get("enable_virtual_address"))
    d.SetId(d.Id())
    return nil
}

func updateGaiaIpv6PimInterface(d *schema.ResourceData, m interface{}) error {

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

    setIpv6PimInterfaceRes, err := client.ApiCallSimple("set-ipv6-pim-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setIpv6PimInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setIpv6PimInterfaceRes.Success {
            errMsg = setIpv6PimInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setIpv6PimInterfaceRes.GetData()
        }

        debugLogOperation(
            "ipv6-pim-interface",        // resource type
            "update",                       // operation
            "set-ipv6-pim-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set ipv6-pim-interface: %v", err)
    }
    if !setIpv6PimInterfaceRes.Success {
        return fmt.Errorf(setIpv6PimInterfaceRes.ErrorMsg)
    }

    // Capture fields returned only by the set API (not by show) before calling Read.
    if data := setIpv6PimInterfaceRes.GetData(); data != nil {
        if v, exists := data["ip-reachability-detection"]; exists {
            d.Set("ip_reachability_detection", v)
        }
        if v, exists := data["neighbor-filter"]; exists {
            d.Set("neighbor_filter", v)
        }
    }

    return readGaiaIpv6PimInterface(d, m)
}

func deleteGaiaIpv6PimInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteIpv6PimInterfaceRes, err := client.ApiCallSimple("delete-ipv6-pim-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteIpv6PimInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteIpv6PimInterfaceRes.Success {
            errMsg = deleteIpv6PimInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteIpv6PimInterfaceRes.GetData()
        }

        debugLogOperation(
            "ipv6-pim-interface",        // resource type
            "delete",                       // operation
            "delete-ipv6-pim-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete ipv6-pim-interface: %v", err)
    }
    if !deleteIpv6PimInterfaceRes.Success {
        return fmt.Errorf(deleteIpv6PimInterfaceRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

