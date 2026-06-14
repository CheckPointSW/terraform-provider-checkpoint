package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaApplyMaestroSecurityGroupsChanges() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaApplyMaestroSecurityGroupsChanges,
        Read:   readGaiaApplyMaestroSecurityGroupsChanges,
        Delete: deleteGaiaApplyMaestroSecurityGroupsChanges,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "security_groups": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "interfaces": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "resource_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "name": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "vlans": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "gateways": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "resource_id": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "site": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "security_group": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "member_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "model": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "version": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "major": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "downlink_ports": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "orchestrator_id": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "port": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "state": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "weight": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "sites": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "resource_id": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "description": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "ftw_configuration": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "hostname": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "is_vsx": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "one_time_password": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Sensitive:   true,
                                        Description: `N/A`,
                                    },
                                    "admin_password": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Sensitive:   true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "mgmt_connectivity": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "ipv4_address": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_address": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv4_mask_length": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_mask_length": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "default_gateway": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_default_gateway": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "mgmt_interface_settings": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "create_mgmt_as_bond": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "bond_mode": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
        },
    }
}

func createGaiaApplyMaestroSecurityGroupsChanges(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    log.Println("Execute apply-maestro-security-groups-changes - Payload = ", payload)

    GaiaApplyMaestroSecurityGroupsChangesRes, err := client.ApiCallSimple("apply-maestro-security-groups-changes", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaApplyMaestroSecurityGroupsChangesRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaApplyMaestroSecurityGroupsChangesRes.Success {
            errMsg = GaiaApplyMaestroSecurityGroupsChangesRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaApplyMaestroSecurityGroupsChangesRes.GetData()
        }

        debugLogOperation(
            "apply-maestro-security-groups-changes",        // resource type
            "command",                       // operation
            "apply-maestro-security-groups-changes",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute apply-maestro-security-groups-changes: %v", err)
    }
    if !GaiaApplyMaestroSecurityGroupsChangesRes.Success {
        if GaiaApplyMaestroSecurityGroupsChangesRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaApplyMaestroSecurityGroupsChangesRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaApplyMaestroSecurityGroupsChangesRes.GetData()
    if v, exists := _respData["security-groups"]; exists {
        if rawGroups, ok := v.([]interface{}); ok {
            groups := make([]interface{}, 0, len(rawGroups))
            for _, g := range rawGroups {
                if gm, ok := g.(map[string]interface{}); ok {
                    group := map[string]interface{}{
                        "resource_id":  func() int { if f, ok := gm["resource-id"].(float64); ok { return int(f) }; return 0 }(),
                        "description":  fmt.Sprintf("%v", gm["description"]),
                    }
                    if ifaces, ok := gm["interfaces"].([]interface{}); ok {
                        ifaceList := make([]interface{}, 0, len(ifaces))
                        for _, iface := range ifaces {
                            if im, ok := iface.(map[string]interface{}); ok {
                                entry := map[string]interface{}{
                                    "resource_id": fmt.Sprintf("%v", im["resource-id"]),
                                    "name":        fmt.Sprintf("%v", im["name"]),
                                    "description": fmt.Sprintf("%v", im["description"]),
                                }
                                if vlans, ok := im["vlans"].([]interface{}); ok {
                                    entry["vlans"] = vlans
                                }
                                ifaceList = append(ifaceList, entry)
                            }
                        }
                        group["interfaces"] = ifaceList
                    }
                    if gws, ok := gm["gateways"].([]interface{}); ok {
                        gwList := make([]interface{}, 0, len(gws))
                        for _, gw := range gws {
                            if gwm, ok := gw.(map[string]interface{}); ok {
                                gwEntry := map[string]interface{}{
                                    "resource_id":    fmt.Sprintf("%v", gwm["resource-id"]),
                                    "site":           func() int { if f, ok := gwm["site"].(float64); ok { return int(f) }; return 0 }(),
                                    "security_group": func() int { if f, ok := gwm["security-group"].(float64); ok { return int(f) }; return 0 }(),
                                    "member_id":      func() int { if f, ok := gwm["member-id"].(float64); ok { return int(f) }; return 0 }(),
                                    "model":          fmt.Sprintf("%v", gwm["model"]),
                                    "description":    fmt.Sprintf("%v", gwm["description"]),
                                    "state":          fmt.Sprintf("%v", gwm["state"]),
                                    "weight":         func() int { if f, ok := gwm["weight"].(float64); ok { return int(f) }; return 0 }(),
                                }
                                if ver, ok := gwm["version"].(map[string]interface{}); ok {
                                    gwEntry["version"] = []interface{}{map[string]interface{}{
                                        "major": fmt.Sprintf("%v", ver["major"]),
                                    }}
                                }
                                if ports, ok := gwm["downlink-ports"].([]interface{}); ok {
                                    portList := make([]interface{}, 0, len(ports))
                                    for _, port := range ports {
                                        if pm, ok := port.(map[string]interface{}); ok {
                                            portList = append(portList, map[string]interface{}{
                                                "orchestrator_id": fmt.Sprintf("%v", pm["orchestrator-id"]),
                                                "port":            fmt.Sprintf("%v", pm["port"]),
                                            })
                                        }
                                    }
                                    gwEntry["downlink_ports"] = portList
                                }
                                gwList = append(gwList, gwEntry)
                            }
                        }
                        group["gateways"] = gwList
                    }
                    if sites, ok := gm["sites"].([]interface{}); ok {
                        siteList := make([]interface{}, 0, len(sites))
                        for _, site := range sites {
                            if sm, ok := site.(map[string]interface{}); ok {
                                siteList = append(siteList, map[string]interface{}{
                                    "resource_id": func() int { if f, ok := sm["resource-id"].(float64); ok { return int(f) }; return 0 }(),
                                    "description": fmt.Sprintf("%v", sm["description"]),
                                })
                            }
                        }
                        group["sites"] = siteList
                    }
                    if ftw, ok := gm["ftw-configuration"].(map[string]interface{}); ok {
                        group["ftw_configuration"] = []interface{}{map[string]interface{}{
                            "hostname":          fmt.Sprintf("%v", ftw["hostname"]),
                            "is_vsx":            func() bool { b, _ := ftw["is-vsx"].(bool); return b }(),
                            "one_time_password": fmt.Sprintf("%v", ftw["one-time-password"]),
                            "admin_password":    fmt.Sprintf("%v", ftw["admin-password"]),
                        }}
                    }
                    if mc, ok := gm["mgmt-connectivity"].(map[string]interface{}); ok {
                        group["mgmt_connectivity"] = []interface{}{map[string]interface{}{
                            "ipv4_address":        fmt.Sprintf("%v", mc["ipv4-address"]),
                            "ipv6_address":        fmt.Sprintf("%v", mc["ipv6-address"]),
                            "ipv4_mask_length":    func() int { if f, ok := mc["ipv4-mask-length"].(float64); ok { return int(f) }; return 0 }(),
                            "ipv6_mask_length":    func() int { if f, ok := mc["ipv6-mask-length"].(float64); ok { return int(f) }; return 0 }(),
                            "default_gateway":     fmt.Sprintf("%v", mc["default-gateway"]),
                            "ipv6_default_gateway": fmt.Sprintf("%v", mc["ipv6-default-gateway"]),
                        }}
                    }
                    if mis, ok := gm["mgmt-interface-settings"].(map[string]interface{}); ok {
                        group["mgmt_interface_settings"] = []interface{}{map[string]interface{}{
                            "create_mgmt_as_bond": func() bool { b, _ := mis["create-mgmt-as-bond"].(bool); return b }(),
                            "bond_mode":          fmt.Sprintf("%v", mis["bond-mode"]),
                        }}
                    }
                    groups = append(groups, group)
                }
            }
            d.Set("security_groups", groups)
        }
    }


    d.SetId(fmt.Sprintf("apply-maestro-security-groups-changes-" + acctest.RandString(10)))
    return nil
}

func readGaiaApplyMaestroSecurityGroupsChanges(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaApplyMaestroSecurityGroupsChanges(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

