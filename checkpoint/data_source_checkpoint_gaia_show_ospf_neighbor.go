package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowOspfNeighbor() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowOspfNeighbor,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "protocol_instance": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Existing OSPFv2 Instance`,
            },
            "neighbor": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Existing OSPFv2 Neighbor`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "address": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "area_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "dead": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "errors": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "interface": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "priority": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "router_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "state": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowOspfNeighbor(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("protocol_instance"); ok {
        payload["protocol-instance"] = v.(string)
    }

    if v, ok := d.GetOk("neighbor"); ok {
        payload["neighbor"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-ospf-neighbor - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-ospf-neighbor", payload)
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
            "ospf-neighbor",        // resource type
            "read",                       // operation
            "show-ospf-neighbor",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-ospf-neighbor: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["address"]; exists {
        d.Set("address", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["area-id"]; exists {
        d.Set("area_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["dead"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("dead", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["errors"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("errors", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["interface"]; exists {
        d.Set("interface", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["priority"]; exists {
        if _f, _ok := v.(float64); _ok {
            d.Set("priority", int(_f))
        }
    }
    if v, exists := commandRes.GetData()["protocol-instance"]; exists {
        d.Set("protocol_instance", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["router-id"]; exists {
        d.Set("router_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["state"]; exists {
        d.Set("state", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-ospf-neighbor-" + acctest.RandString(10)))
    return nil
}

