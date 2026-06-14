package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaMaestroGateway() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaMaestroGateway,
        Read:   readGaiaMaestroGateway,
        Update: updateGaiaMaestroGateway,
        Delete: deleteGaiaMaestroGateway,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "resource_id": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `ID of Gateway to modify`,
            },
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `New Gateway description`,
            },
            "security_group": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `ID of a Security Group. If specified, the Gateway will be assigned to this Security Group,regardless of it's current assignment status. In case you want to unassign Gateway from Security Group, use 0`,
            },
            "include_pending_changes": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `If true, show pending topology. If false, show deployed topology`,
            },
            "site": {
                Type:     schema.TypeInt,
                Computed: true,
            },
            "member_id": {
                Type:     schema.TypeInt,
                Computed: true,
            },
            "model": {
                Type:     schema.TypeString,
                Computed: true,
            },
            "state": {
                Type:     schema.TypeString,
                Computed: true,
            },
            "weight": {
                Type:     schema.TypeInt,
                Computed: true,
            },
            "version": {
                Type:     schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "major": {
                            Type:     schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },
            "downlink_ports": {
                Type:     schema.TypeList,
                Computed: true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "orchestrator_id": {
                            Type:     schema.TypeString,
                            Computed: true,
                        },
                        "port": {
                            Type:     schema.TypeString,
                            Computed: true,
                        },
                    },
                },
            },
        },
    }
}

func createGaiaMaestroGateway(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    if v, ok := d.GetOk("description"); ok {
        payload["description"] = v.(string)
    }

    if v, ok := d.GetOkExists("security_group"); ok {
        payload["security-group"] = v.(int)
    }

    log.Println("Create MaestroGateway - Map = ", payload)

    addMaestroGatewayRes, err := client.ApiCallSimple("set-maestro-gateway", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addMaestroGatewayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addMaestroGatewayRes.Success {
            errMsg = addMaestroGatewayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addMaestroGatewayRes.GetData()
        }

        debugLogOperation(
            "maestro-gateway",        // resource type
            "create",                       // operation
            "set-maestro-gateway",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add maestro-gateway: %v", err)
    }
    if !addMaestroGatewayRes.Success {
        if addMaestroGatewayRes.ErrorMsg != "" {
            return fmt.Errorf(addMaestroGatewayRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("maestro-gateway-" + acctest.RandString(10)))
    return readGaiaMaestroGateway(d, m)
}

func readGaiaMaestroGateway(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    if v, ok := d.GetOkExists("include_pending_changes"); ok {
        payload["include-pending-changes"] = v.(bool)
    }

    showMaestroGatewayRes, err := client.ApiCallSimple("show-maestro-gateway", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showMaestroGatewayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showMaestroGatewayRes.Success {
            errMsg = showMaestroGatewayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showMaestroGatewayRes.GetData()
        }

        debugLogOperation(
            "maestro-gateway",        // resource type
            "read",                       // operation
            "show-maestro-gateway",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show maestro-gateway: %v", err)
    }
    if !showMaestroGatewayRes.Success {
        if data := showMaestroGatewayRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showMaestroGatewayRes.ErrorMsg)
    }

    maestroGateway := showMaestroGatewayRes.GetData()

    log.Println("Read MaestroGateway - Show JSON = ", maestroGateway)

    if v, exists := maestroGateway["id"]; exists {
        d.Set("resource_id", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroGateway["site"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("site", int(f))
        }
    }
    if v, exists := maestroGateway["security-group"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("security_group", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("security_group", _n)
            }
        }
    }
    if v, exists := maestroGateway["member-id"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("member_id", int(f))
        }
    }
    if v, exists := maestroGateway["model"]; exists {
        d.Set("model", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroGateway["version"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            d.Set("version", []interface{}{map[string]interface{}{
                "major": fmt.Sprintf("%v", m["major"]),
            }})
        }
    }
    if v, exists := maestroGateway["downlink-ports"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    out = append(out, map[string]interface{}{
                        "orchestrator_id": fmt.Sprintf("%v", m["orchestrator-id"]),
                        "port":            fmt.Sprintf("%v", m["port"]),
                    })
                }
            }
            d.Set("downlink_ports", out)
        }
    }
    if v, exists := maestroGateway["description"]; exists {
        d.Set("description", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroGateway["state"]; exists {
        d.Set("state", fmt.Sprintf("%v", v))
    }
    if v, exists := maestroGateway["weight"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("weight", int(f))
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaMaestroGateway(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    if v, ok := d.GetOk("description"); ok {
        payload["description"] = v.(string)
    }

    if v, ok := d.GetOkExists("security_group"); ok {
        payload["security-group"] = v.(int)
    }

    setMaestroGatewayRes, err := client.ApiCallSimple("set-maestro-gateway", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setMaestroGatewayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setMaestroGatewayRes.Success {
            errMsg = setMaestroGatewayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setMaestroGatewayRes.GetData()
        }

        debugLogOperation(
            "maestro-gateway",        // resource type
            "update",                       // operation
            "set-maestro-gateway",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set maestro-gateway: %v", err)
    }
    if !setMaestroGatewayRes.Success {
        return fmt.Errorf(setMaestroGatewayRes.ErrorMsg)
    }

    return readGaiaMaestroGateway(d, m)
}

func deleteGaiaMaestroGateway(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    