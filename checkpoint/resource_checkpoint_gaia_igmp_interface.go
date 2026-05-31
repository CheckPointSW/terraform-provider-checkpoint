package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaIgmpInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaIgmpInterface,
        Read:   readGaiaIgmpInterface,
        Update: updateGaiaIgmpInterface,
        Delete: deleteGaiaIgmpInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The name of the IGMP interface`,
            },
            "last_member_query_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The number of seconds between queries that an IGMP router sends after it receives a \"Leave Group\" message from a host`,
            },
            "loss_robustness": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The loss-robustness value`,
            },
            "query_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The number of seconds between IGMP general queries`,
            },
            "query_response_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The maximum delay time in seconds for hosts to respond to an IGMP membership query`,
            },
            "router_alert": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Configure the router-alert option for this IGMP interface`,
            },
            "igmp_version": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The IGMP version running`,
            },
            "reset": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Reset all attributes of this interface to default values`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaIgmpInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("last_member_query_interval"); ok {
        payload["last-member-query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("loss_robustness"); ok {
        payload["loss-robustness"] = v.(string)
    }

    if v, ok := d.GetOk("query_interval"); ok {
        payload["query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("query_response_interval"); ok {
        payload["query-response-interval"] = v.(string)
    }

    if v, ok := d.GetOkExists("router_alert"); ok {
        payload["router-alert"] = v.(bool)
    }

    if v, ok := d.GetOk("igmp_version"); ok {
        payload["igmp-version"] = v.(string)
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    log.Println("Create IgmpInterface - Map = ", payload)

    addIgmpInterfaceRes, err := client.ApiCallSimple("set-igmp-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addIgmpInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addIgmpInterfaceRes.Success {
            errMsg = addIgmpInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addIgmpInterfaceRes.GetData()
        }

        debugLogOperation(
            "igmp-interface",        // resource type
            "create",                       // operation
            "set-igmp-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add igmp-interface: %v", err)
    }
    if !addIgmpInterfaceRes.Success {
        if addIgmpInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addIgmpInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("igmp-interface-" + acctest.RandString(10)))
    return readGaiaIgmpInterface(d, m)
}

func readGaiaIgmpInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showIgmpInterfaceRes, err := client.ApiCallSimple("show-igmp-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showIgmpInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showIgmpInterfaceRes.Success {
            errMsg = showIgmpInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showIgmpInterfaceRes.GetData()
        }

        debugLogOperation(
            "igmp-interface",        // resource type
            "read",                       // operation
            "show-igmp-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show igmp-interface: %v", err)
    }
    if !showIgmpInterfaceRes.Success {
        if data := showIgmpInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        if strings.Contains(strings.ToLower(showIgmpInterfaceRes.ErrorMsg), "igmp is not enabled") {
            // Daemon does not reflect config-db state yet; preserve existing state.
            return nil
        }
        return fmt.Errorf(showIgmpInterfaceRes.ErrorMsg)
    }

    igmpInterface := showIgmpInterfaceRes.GetData()

    log.Println("Read IgmpInterface - Show JSON = ", igmpInterface)

    if v, exists := igmpInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := igmpInterface["flags"]; exists {
        d.Set("flags", fmt.Sprintf("%v", v))
    }
    if v, exists := igmpInterface["igmp-version"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("igmp_version", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("igmp_version", _n)
            }
        }
    }
    if v, exists := igmpInterface["loss-robustness"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("loss_robustness", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("loss_robustness", _n)
            }
        }
    }
    if v, exists := igmpInterface["query-interval"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("query_interval", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("query_interval", _n)
            }
        }
    }
    if v, exists := igmpInterface["query-response-interval"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("query_response_interval", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("query_response_interval", _n)
            }
        }
    }
    if v, exists := igmpInterface["last-member-query-interval"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("last_member_query_interval", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("last_member_query_interval", _n)
            }
        }
    }
    if v, exists := igmpInterface["num-group-joins"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("num_group_joins", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("num_group_joins", _n)
            }
        }
    }
    if v, exists := igmpInterface["num-group-active"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("num_group_active", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("num_group_active", _n)
            }
        }
    }
    if v, exists := igmpInterface["next-membership-query"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("next_membership_query", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("next_membership_query", _n)
            }
        }
    }
    if v, exists := igmpInterface["querier-up-time"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("querier_up_time", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("querier_up_time", _n)
            }
        }
    }
    if v, exists := igmpInterface["advertise-address"]; exists {
        d.Set("advertise_address", fmt.Sprintf("%v", v))
    }
    if v, exists := igmpInterface["multicast-protocol"]; exists {
        d.Set("multicast_protocol", fmt.Sprintf("%v", v))
    }
    if v, exists := igmpInterface["router-alert"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("router_alert", b)
        } else if s, ok := v.(string); ok {
            d.Set("router_alert", s == "true")
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaIgmpInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("last_member_query_interval"); ok {
        payload["last-member-query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("loss_robustness"); ok {
        payload["loss-robustness"] = v.(string)
    }

    if v, ok := d.GetOk("query_interval"); ok {
        payload["query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("query_response_interval"); ok {
        payload["query-response-interval"] = v.(string)
    }

    if v, ok := d.GetOkExists("router_alert"); ok {
        payload["router-alert"] = v.(bool)
    }

    if v, ok := d.GetOk("igmp_version"); ok {
        payload["igmp-version"] = v.(string)
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    setIgmpInterfaceRes, err := client.ApiCallSimple("set-igmp-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setIgmpInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setIgmpInterfaceRes.Success {
            errMsg = setIgmpInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setIgmpInterfaceRes.GetData()
        }

        debugLogOperation(
            "igmp-interface",        // resource type
            "update",                       // operation
            "set-igmp-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set igmp-interface: %v", err)
    }
    if !setIgmpInterfaceRes.Success {
        return fmt.Errorf(setIgmpInterfaceRes.ErrorMsg)
    }

    return readGaiaIgmpInterface(d, m)
}

func deleteGaiaIgmpInterface(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    