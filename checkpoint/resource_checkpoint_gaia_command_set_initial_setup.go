package checkpoint

import (
        "context"
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSetInitialSetup() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetInitialSetup,
        Read:   readGaiaSetInitialSetup,
        Delete: deleteGaiaSetInitialSetup,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "password": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Sensitive:   true,
                Description: `Password of user admin. Required in case default initial password has not been changed before`,
            },
            "security_gateway": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Install Security Gateway.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "activation_key": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Secure Internal Communication key`,
                        },
                        "dynamically_assigned_ip": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Enable DAIP (Dynamic IP) gateway. Should be false if cluster-member or security-management enabled`,
                        },
                        "cluster_member": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Enable/Disable ClusterXL.`,
                        },
                        "vsnext": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Enable/Disable VSNext. To use VSNext, elastic-xl must be true`,
                        },
                        "elastic_xl": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Enable/Disable ElasticXL. Cannot be enabled in combination with cluster-member`,
                        },
                    },
                },
            },
            "security_management": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Install Security Management or Multi-Domain Server`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Type of security management or Multi-Domain Server`,
                        },
                        "multi_domain": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Install Security Multi-Domain Server, it can be primary or secondary or Log Server according to type parameter`,
                        },
                        "gui_clients": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Choose which GUI clients can log into the Security Management. Fill one of the parameters (range/network/Single IP), for Multi-Domain Server it can be only Single IP or can keep the default value`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "range": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Range of IPs allowed to connect to management`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "first_ipv4_range": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `First IP in range`,
                                                },
                                                "last_ipv4_range": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Last IP in range`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPs from specific network allowed to connect to management`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `IPv4 address of network`,
                                                },
                                                "mask_length": {
                                                    Type:        schema.TypeInt,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Mask length of network`,
                                                },
                                            },
                                        },
                                    },
                                    "single_ip": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `In case of a single IP which allowed to connect to management`,
                                    },
                                },
                            },
                        },
                        "activation_key": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Secure Internal Communication key, relevant in case of secondary or Log Server`,
                        },
                        "leading_interface": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Leading Multi-Domain Server interface, relevant in case of Multi-Domain Server enabled`,
                        },
                    },
                },
            },
            "grub_password": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Sensitive:   true,
                Description: `Password of the GRUB maintenence. Required in case default initial GRUB password has not been changed before`,
            },
            "task_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaSetInitialSetup(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v := d.Get("security_gateway"); len(v.([]interface{})) > 0 {
        _ = v
        securitygatewayMap := make(map[string]interface{})
        if v, ok := d.GetOk("security_gateway.0.activation_key"); ok {
            securitygatewayMap["activation-key"] = v.(string)
        }
        if v, ok := d.GetOkExists("security_gateway.0.dynamically_assigned_ip"); ok && v.(bool) {
            securitygatewayMap["dynamically-assigned-ip"] = v.(bool)
        }
        if v, ok := d.GetOkExists("security_gateway.0.cluster_member"); ok && v.(bool) {
            securitygatewayMap["cluster-member"] = v.(bool)
        }
        if v, ok := d.GetOkExists("security_gateway.0.vsnext"); ok && v.(bool) {
            securitygatewayMap["vsnext"] = v.(bool)
        }
        if v, ok := d.GetOkExists("security_gateway.0.elastic_xl"); ok && v.(bool) {
            securitygatewayMap["elastic-xl"] = v.(bool)
        }
        if len(securitygatewayMap) > 0 {
            payload["security-gateway"] = securitygatewayMap
        }
    }

    if v := d.Get("security_management"); len(v.([]interface{})) > 0 {
        _ = v
        securitymanagementMap := make(map[string]interface{})
        if v, ok := d.GetOk("security_management.0.type"); ok {
            securitymanagementMap["type"] = v.(string)
        }
        if v, ok := d.GetOkExists("security_management.0.multi_domain"); ok && v.(bool) {
            securitymanagementMap["multi-domain"] = v.(bool)
        }
        if v, ok := d.GetOk("security_management.0.gui_clients"); ok {
            _ = v
            guiclientsMap := make(map[string]interface{})
            if v, ok := d.GetOk("security_management.0.gui_clients.0.range"); ok {
                _ = v
                rangeMap := make(map[string]interface{})
                if v, ok := d.GetOk("security_management.0.gui_clients.0.range.0.first_ipv4_range"); ok {
                    rangeMap["first-IPv4-range"] = v.(string)
                }
                if v, ok := d.GetOk("security_management.0.gui_clients.0.range.0.last_ipv4_range"); ok {
                    rangeMap["last-IPv4-range"] = v.(string)
                }
                if len(rangeMap) > 0 {
                    guiclientsMap["range"] = rangeMap
                }
            }
            if v, ok := d.GetOk("security_management.0.gui_clients.0.network"); ok {
                _ = v
                networkMap := make(map[string]interface{})
                if v, ok := d.GetOk("security_management.0.gui_clients.0.network.0.address"); ok {
                    networkMap["address"] = v.(string)
                }
                if v, ok := d.GetOk("security_management.0.gui_clients.0.network.0.mask_length"); ok {
                    networkMap["mask-length"] = v.(int)
                }
                if len(networkMap) > 0 {
                    guiclientsMap["network"] = networkMap
                }
            }
            if v, ok := d.GetOk("security_management.0.gui_clients.0.single_ip"); ok {
                guiclientsMap["single-ip"] = v.(string)
            }
            if len(guiclientsMap) > 0 {
                securitymanagementMap["gui-clients"] = guiclientsMap
            }
        }
        if v, ok := d.GetOk("security_management.0.activation_key"); ok {
            securitymanagementMap["activation-key"] = v.(string)
        }
        if v, ok := d.GetOk("security_management.0.leading_interface"); ok {
            securitymanagementMap["leading-interface"] = v.(string)
        }
        if len(securitymanagementMap) > 0 {
            payload["security-management"] = securitymanagementMap
        }
    }

    if v, ok := d.GetOk("grub_password"); ok {
        payload["grub-password"] = v.(string)
    }

    log.Println("Execute set-initial-setup - Payload = ", payload)

    GaiaSetInitialSetupRes, err := client.ApiCallSimple("set-initial-setup", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetInitialSetupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetInitialSetupRes.Success {
            errMsg = GaiaSetInitialSetupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetInitialSetupRes.GetData()
        }

        debugLogOperation(
            "set-initial-setup",        // resource type
            "command",                       // operation
            "set-initial-setup",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-initial-setup: %v", err)
    }
    if !GaiaSetInitialSetupRes.Success {
        if GaiaSetInitialSetupRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetInitialSetupRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-initial-setup", GaiaSetInitialSetupRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-initial-setup task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        msg := taskRes.Message
        if msg == "" {
            msg = fmt.Sprintf("set-initial-setup task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(msg)
    }

    _respData := GaiaSetInitialSetupRes.GetData()
    if v, exists := _respData["task-id"]; exists {
        d.Set("task_id", toString(v))
    }


    d.SetId(fmt.Sprintf("set-initial-setup-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetInitialSetup(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSetInitialSetup(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

