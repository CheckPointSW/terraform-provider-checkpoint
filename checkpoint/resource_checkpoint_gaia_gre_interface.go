package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaGreInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaGreInterface,
        Read:   readGaiaGreInterface,
        Update: updateGaiaGreInterface,
        Delete: deleteGaiaGreInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "gre_id": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `ID number represents the tunnel ID.`,
            },
            "local_ip_address": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `IP address of the underlying local interface on this gateway.`,
            },
            "remote_ip_address": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `IP address of the underlying remote interface on the router on the other end of the tunnel.`,
            },
            "ttl": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `N/A`,
            },
            "ipv4_address": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Assigned IP of the GRE.`,
            },
            "ipv4_mask_length": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `N/A`,
            },
            "peer_address": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `IP address of the remote peer.`,
            },
            "comments": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `N/A`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "mtu": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `N/A`,
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                Computed:    true,
                Description: `N/A`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "link_state": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "speed": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "duplex": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_bytes": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "tx_packets": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_bytes": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "rx_packets": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_autoconfig": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_mask_length": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "ipv6_local_link_address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaGreInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("gre_id"); ok {
        payload["gre-id"] = v.(int)
    }

    if v, ok := d.GetOk("local_ip_address"); ok {
        payload["local-ip-address"] = v.(string)
    }

    if v, ok := d.GetOk("remote_ip_address"); ok {
        payload["remote-ip-address"] = v.(string)
    }

    if v, ok := d.GetOk("ttl"); ok {
        payload["ttl"] = v.(int)
    }

    if v, ok := d.GetOk("ipv4_address"); ok {
        payload["ipv4-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_mask_length"); ok {
        payload["ipv4-mask-length"] = v.(int)
    }

    if v, ok := d.GetOk("peer_address"); ok {
        payload["peer-address"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create GreInterface - Map = ", payload)

    addGreInterfaceRes, err := client.ApiCallSimple("add-gre-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addGreInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addGreInterfaceRes.Success {
            errMsg = addGreInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addGreInterfaceRes.GetData()
        }

        debugLogOperation(
            "gre-interface",        // resource type
            "create",                       // operation
            "add-gre-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add gre-interface: %v", err)
    }
    if !addGreInterfaceRes.Success {
        if addGreInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addGreInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    // Extract API-assigned fields from Create response before calling Read.
    if data := addGreInterfaceRes.GetData(); data != nil {
        if v, exists := data["name"]; exists {
            d.Set("name", v)
        }
    }

    // Two-phase creation: Apply update-only fields if present
    hasUpdateOnlyFields := false
    updatePayload := map[string]interface{}{
    }

    if v, ok := d.GetOk("mtu"); ok {
        updatePayload["mtu"] = v.(int)
        hasUpdateOnlyFields = true
    }

    if v, ok := d.GetOk("name"); ok {
        updatePayload["name"] = v.(string)
        hasUpdateOnlyFields = true
    }

    if v, ok := d.GetOkExists("enabled"); ok {
        updatePayload["enabled"] = v.(bool)
        hasUpdateOnlyFields = true
    }

    if hasUpdateOnlyFields {
        log.Println("Two-phase creation: Applying update-only fields - Map = ", updatePayload)
        
        setGreInterfaceRes, err := client.ApiCallSimple("set-gre-interface", updatePayload)
        if err != nil {
            return fmt.Errorf("Failed to apply update-only fields for gre-interface: %v", err)
        }
        if !setGreInterfaceRes.Success {
            return fmt.Errorf("Failed to apply update-only fields: %s", setGreInterfaceRes.ErrorMsg)
        }
    }

    d.SetId(fmt.Sprintf("gre-interface-" + acctest.RandString(10)))
    return readGaiaGreInterface(d, m)
}

func readGaiaGreInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showGreInterfaceRes, err := client.ApiCallSimple("show-gre-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showGreInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showGreInterfaceRes.Success {
            errMsg = showGreInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showGreInterfaceRes.GetData()
        }

        debugLogOperation(
            "gre-interface",        // resource type
            "read",                       // operation
            "show-gre-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show gre-interface: %v", err)
    }
    if !showGreInterfaceRes.Success {
        if data := showGreInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showGreInterfaceRes.ErrorMsg)
    }

    greInterface := showGreInterfaceRes.GetData()

    log.Println("Read GreInterface - Show JSON = ", greInterface)

    if v, exists := greInterface["link-state"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("link_state", b)
        } else if s, ok := v.(string); ok {
            d.Set("link_state", s == "true")
        }
    }
    if v, exists := greInterface["speed"]; exists {
        d.Set("speed", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["duplex"]; exists {
        d.Set("duplex", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["tx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_bytes", _n)
            }
        }
    }
    if v, exists := greInterface["tx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("tx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("tx_packets", _n)
            }
        }
    }
    if v, exists := greInterface["rx-bytes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_bytes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_bytes", _n)
            }
        }
    }
    if v, exists := greInterface["rx-packets"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("rx_packets", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("rx_packets", _n)
            }
        }
    }
    if v, exists := greInterface["gre-id"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("gre_id", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("gre_id", _n)
            }
        }
    }
    if v, exists := greInterface["local-ip-address"]; exists {
        d.Set("local_ip_address", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["remote-ip-address"]; exists {
        d.Set("remote_ip_address", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["ttl"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("ttl", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("ttl", _n)
            }
        }
    }
    if v, exists := greInterface["ipv4-address"]; exists {
        d.Set("ipv4_address", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["peer-address"]; exists {
        d.Set("peer_address", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := greInterface["mtu"]; exists {
        d.Set("mtu", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["ipv4-mask-length"]; exists {
        d.Set("ipv4_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["ipv6-autoconfig"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("ipv6_autoconfig", b)
        } else if s, ok := v.(string); ok {
            d.Set("ipv6_autoconfig", s == "true")
        }
    }
    if v, exists := greInterface["comments"]; exists {
        d.Set("comments", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["ipv6-address"]; exists {
        d.Set("ipv6_address", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["ipv6-mask-length"]; exists {
        d.Set("ipv6_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["ipv6-local-link-address"]; exists {
        d.Set("ipv6_local_link_address", fmt.Sprintf("%v", v))
    }
    if v, exists := greInterface["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaGreInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("mtu"); ok {
        payload["mtu"] = v.(int)
    }

   payload["name"] = d.Get("name")
    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    setGreInterfaceRes, err := client.ApiCallSimple("set-gre-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setGreInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setGreInterfaceRes.Success {
            errMsg = setGreInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setGreInterfaceRes.GetData()
        }

        debugLogOperation(
            "gre-interface",        // resource type
            "update",                       // operation
            "set-gre-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set gre-interface: %v", err)
    }
    if !setGreInterfaceRes.Success {
        return fmt.Errorf(setGreInterfaceRes.ErrorMsg)
    }

    return readGaiaGreInterface(d, m)
}

func deleteGaiaGreInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    deleteGreInterfaceRes, err := client.ApiCallSimple("delete-gre-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteGreInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteGreInterfaceRes.Success {
            errMsg = deleteGreInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteGreInterfaceRes.GetData()
        }

        debugLogOperation(
            "gre-interface",        // resource type
            "delete",                       // operation
            "delete-gre-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete gre-interface: %v", err)
    }
    if !deleteGreInterfaceRes.Success {
        return fmt.Errorf(deleteGreInterfaceRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

