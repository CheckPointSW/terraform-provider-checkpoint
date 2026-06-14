package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "context"
    "strings"

)
func resourceGaiaVirtualGateway() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaVirtualGateway,
        Read:   readGaiaVirtualGateway,
        Update: updateGaiaVirtualGateway,
        Delete: deleteGaiaVirtualGateway,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "resource_id": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `Virtual gateway identifier can be an integer or the next avaliable id (auto)`,
            },
            "one_time_password": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Sensitive:   true,
                Description: `One time password, used for the secure internal communication with the Management object`,
            },
            "interfaces": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `Network interface(s) to be attached`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "virtual_switches": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `Virtual switche(s) to be connected, mgmt-switch (id 500) is set as default`,
                Elem: &schema.Schema{
                    Type: schema.TypeInt,
                },
            },
            "resources": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Additional resources`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "firewall_ipv4_instances": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `CoreXL IPv4 instances amount. Must be between 1 and the greater of 32 and the number of CPU cores.`,
                        },
                        "firewall_ipv6_instances": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `CoreXL IPv6 instances amount. Must be between 0 and the greater of 32 and the number of CPU cores. Must not exceed the number of IPv4 CoreXL instances.`,
                        },
                    },
                },
            },
            "mgmt_connection": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Management connection configurations`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "mgmt_connection_identifier": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Management connection identifier according to the connection type (interface or virtual-switch (id or name))`,
                        },
                        "mgmt_connection_type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Management connection type - interface or virtual link connected to virtual-switch (identified by name or id)`,
                        },
                        "mgmt_ipv4_configuration": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Management IPv4 configuration`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ipv4_address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `IPv4 address`,
                                    },
                                    "ipv4_mask": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `IPv4 mask`,
                                    },
                                    "ipv4_default_gateway": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `IPv4 default gateway`,
                                    },
                                },
                            },
                        },
                        "mgmt_ipv6_configuration": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Management IPv6 configuration`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ipv6_address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `IPv6 address`,
                                    },
                                    "ipv6_mask": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `IPv6 mask length`,
                                    },
                                    "ipv6_default_gateway": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `IPv6 default gateway`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "set_if_exist": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `If another virtual gateway with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original virtual gateway's fields will be overwritten by the fields provided in the request payload!`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "action": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "status": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "message": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "vsxd_task_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "vs_id": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaVirtualGateway(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v, ok := d.GetOk("one_time_password"); ok {
        payload["one-time-password"] = v.(string)
    }

    if v := d.Get("interfaces"); len(v.(*schema.Set).List()) > 0 {
        payload["interfaces"] = v.(*schema.Set).List()
    }

    if v := d.Get("virtual_switches"); len(v.(*schema.Set).List()) > 0 {
        payload["virtual-switches"] = v.(*schema.Set).List()
    }

    if v := d.Get("resources"); len(v.([]interface{})) > 0 {
        _ = v
        resourcesMap := make(map[string]interface{})
        if v, ok := d.GetOk("resources.0.firewall_ipv4_instances"); ok {
            resourcesMap["firewall-ipv4-instances"] = v.(int)
        }
        if v, ok := d.GetOk("resources.0.firewall_ipv6_instances"); ok {
            resourcesMap["firewall-ipv6-instances"] = v.(int)
        }
        if len(resourcesMap) > 0 {
            payload["resources"] = resourcesMap
        }
    }

    if v := d.Get("mgmt_connection"); len(v.([]interface{})) > 0 {
        _ = v
        mgmtconnectionMap := make(map[string]interface{})
        if v, ok := d.GetOk("mgmt_connection.0.mgmt_connection_identifier"); ok {
            mgmtconnectionMap["mgmt-connection-identifier"] = v.(string)
        }
        if v, ok := d.GetOk("mgmt_connection.0.mgmt_connection_type"); ok {
            mgmtconnectionMap["mgmt-connection-type"] = v.(string)
        }
        if v, ok := d.GetOk("mgmt_connection.0.mgmt_ipv4_configuration"); ok {
            _ = v
            mgmtipv4configurationMap := make(map[string]interface{})
            if v, ok := d.GetOk("mgmt_connection.0.mgmt_ipv4_configuration.0.ipv4_address"); ok {
                mgmtipv4configurationMap["ipv4-address"] = v.(string)
            }
            if v, ok := d.GetOk("mgmt_connection.0.mgmt_ipv4_configuration.0.ipv4_mask"); ok {
                mgmtipv4configurationMap["ipv4-mask"] = v.(int)
            }
            if v, ok := d.GetOk("mgmt_connection.0.mgmt_ipv4_configuration.0.ipv4_default_gateway"); ok {
                mgmtipv4configurationMap["ipv4-default-gateway"] = v.(string)
            }
            if len(mgmtipv4configurationMap) > 0 {
                mgmtconnectionMap["mgmt-ipv4-configuration"] = mgmtipv4configurationMap
            }
        }
        if v, ok := d.GetOk("mgmt_connection.0.mgmt_ipv6_configuration"); ok {
            _ = v
            mgmtipv6configurationMap := make(map[string]interface{})
            if v, ok := d.GetOk("mgmt_connection.0.mgmt_ipv6_configuration.0.ipv6_address"); ok {
                mgmtipv6configurationMap["ipv6-address"] = v.(string)
            }
            if v, ok := d.GetOk("mgmt_connection.0.mgmt_ipv6_configuration.0.ipv6_mask"); ok {
                mgmtipv6configurationMap["ipv6-mask"] = v.(int)
            }
            if v, ok := d.GetOk("mgmt_connection.0.mgmt_ipv6_configuration.0.ipv6_default_gateway"); ok {
                mgmtipv6configurationMap["ipv6-default-gateway"] = v.(string)
            }
            if len(mgmtipv6configurationMap) > 0 {
                mgmtconnectionMap["mgmt-ipv6-configuration"] = mgmtipv6configurationMap
            }
        }
        if len(mgmtconnectionMap) > 0 {
            payload["mgmt-connection"] = mgmtconnectionMap
        }
    }

    if v, ok := d.GetOkExists("set_if_exist"); ok {
        payload["set-if-exist"] = v.(bool)
    }

    log.Println("Create VirtualGateway - Map = ", payload)

    addVirtualGatewayRes, err := client.ApiCallSimple("add-virtual-gateway", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addVirtualGatewayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addVirtualGatewayRes.Success {
            errMsg = addVirtualGatewayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addVirtualGatewayRes.GetData()
        }

        debugLogOperation(
            "virtual-gateway",        // resource type
            "create",                       // operation
            "add-virtual-gateway",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add virtual-gateway: %v", err)
    }
    if !addVirtualGatewayRes.Success {
        if addVirtualGatewayRes.ErrorMsg != "" {
            return fmt.Errorf(addVirtualGatewayRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "add-virtual-gateway", addVirtualGatewayRes, true, 0)
    if err != nil {
        return fmt.Errorf("add-virtual-gateway task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("add-virtual-gateway task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    d.SetId(fmt.Sprintf("virtual-gateway-" + acctest.RandString(10)))
    return readGaiaVirtualGateway(d, m)
}

func readGaiaVirtualGateway(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showVirtualGatewayRes, err := client.ApiCallSimple("show-virtual-gateway", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showVirtualGatewayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showVirtualGatewayRes.Success {
            errMsg = showVirtualGatewayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showVirtualGatewayRes.GetData()
        }

        debugLogOperation(
            "virtual-gateway",        // resource type
            "read",                       // operation
            "show-virtual-gateway",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show virtual-gateway: %v", err)
    }
    if !showVirtualGatewayRes.Success {
        if data := showVirtualGatewayRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showVirtualGatewayRes.ErrorMsg)
    }

    virtualGateway := showVirtualGatewayRes.GetData()

    log.Println("Read VirtualGateway - Show JSON = ", virtualGateway)

    if v, exists := virtualGateway["id"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("resource_id", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("resource_id", _n)
            }
        }
    }
    if v, exists := virtualGateway["interfaces"]; exists {
        d.Set("interfaces", v.([]interface{}))
    }
    if v, exists := virtualGateway["virtual-switches"]; exists {
        if vsList, ok := v.([]interface{}); ok {
            ids := make([]interface{}, 0, len(vsList))
            for _, item := range vsList {
                switch val := item.(type) {
                case map[string]interface{}:
                    if id, ok := val["id"].(float64); ok {
                        ids = append(ids, int(id))
                    }
                case float64:
                    ids = append(ids, int(val))
                }
            }
            d.Set("virtual_switches", ids)
        }
    }
    if v, exists := virtualGateway["resources"]; exists {
        if rm, ok := v.(map[string]interface{}); ok {
            resourcesMap := map[string]interface{}{}
            if val, ok := rm["firewall-ipv4-instances"]; ok {
                if f, ok := val.(float64); ok {
                    resourcesMap["firewall_ipv4_instances"] = int(f)
                }
            }
            if val, ok := rm["firewall-ipv6-instances"]; ok {
                if f, ok := val.(float64); ok {
                    resourcesMap["firewall_ipv6_instances"] = int(f)
                }
            }
            d.Set("resources", []interface{}{resourcesMap})
        }
    }
    if v, exists := virtualGateway["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaVirtualGateway(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v := d.Get("interfaces"); len(v.(*schema.Set).List()) > 0 {
        payload["interfaces"] = v.(*schema.Set).List()
    }

    if v := d.Get("virtual_switches"); len(v.(*schema.Set).List()) > 0 {
        payload["virtual-switches"] = v.(*schema.Set).List()
    }

    if v := d.Get("resources"); len(v.([]interface{})) > 0 {
        _ = v
        resourcesMap := make(map[string]interface{})
        if v, ok := d.GetOk("resources.0.firewall_ipv4_instances"); ok {
            resourcesMap["firewall-ipv4-instances"] = v.(int)
        }
        if v, ok := d.GetOk("resources.0.firewall_ipv6_instances"); ok {
            resourcesMap["firewall-ipv6-instances"] = v.(int)
        }
        if len(resourcesMap) > 0 {
            payload["resources"] = resourcesMap
        }
    }

    setVirtualGatewayRes, err := client.ApiCallSimple("set-virtual-gateway", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setVirtualGatewayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setVirtualGatewayRes.Success {
            errMsg = setVirtualGatewayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setVirtualGatewayRes.GetData()
        }

        debugLogOperation(
            "virtual-gateway",        // resource type
            "update",                       // operation
            "set-virtual-gateway",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set virtual-gateway: %v", err)
    }
    if !setVirtualGatewayRes.Success {
        return fmt.Errorf(setVirtualGatewayRes.ErrorMsg)
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-virtual-gateway", setVirtualGatewayRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-virtual-gateway task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("set-virtual-gateway task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    return readGaiaVirtualGateway(d, m)
}

func deleteGaiaVirtualGateway(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    deleteVirtualGatewayRes, err := client.ApiCallSimple("delete-virtual-gateway", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteVirtualGatewayRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteVirtualGatewayRes.Success {
            errMsg = deleteVirtualGatewayRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteVirtualGatewayRes.GetData()
        }

        debugLogOperation(
            "virtual-gateway",        // resource type
            "delete",                       // operation
            "delete-virtual-gateway",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete virtual-gateway: %v", err)
    }
    if !deleteVirtualGatewayRes.Success {
        return fmt.Errorf(deleteVirtualGatewayRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

