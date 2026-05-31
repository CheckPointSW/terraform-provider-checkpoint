package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSetBgpInternal() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetBgpInternal,
        Read:   readGaiaSetBgpInternal,
        Delete: deleteGaiaSetBgpInternal,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `Enable Internal BGP (IBGP) on this router.`,
            },
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Adds a brief description of the peer group.`,
            },
            "export_routemap_list": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure export policy for the given BGP peer group or peer.`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Name of the routemap`,
                        },
                        "preference": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Preference for the routemap. Routemaps are evaluated in order of increasing preference value.`,
                        },
                        "family": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Describes which family of routes this routemap will be applied to.`,
                        },
                        "conditional_routemap": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Condition to apply to the routemap`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `The name of the routemap condition`,
                                    },
                                    "condition": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `The condition can be any-pass or no-pass`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "import_routemap_list": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure import policy for the given BGP peer group.`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Name of the routemap`,
                        },
                        "preference": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Preference for the routemap. Routemaps are evaluated in order of increasing preference value.`,
                        },
                        "family": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Describes which family of routes this routemap will be applied to.`,
                        },
                    },
                },
            },
            "interface_list": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Specifies the interfaces for which third-party next hops may be used. By default, all interfaces are enabled.`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "local_address": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configures the address to be used on the local end of the TCP connection.`,
            },
            "med": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Defines the Multi-Exit Discriminator (MED) metric used when advertising routes to all peers in this group.`,
            },
            "enable_nexthop_self": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `When this option is enabled, the router sends its own IP address as the BGP next hop.`,
            },
            "outdelay": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Specifies the length of time (seconds) that a route must be present in the routing database before it is redistributed to BGP.`,
            },
            "protocol_list": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Enables specific routing protocols to use as an Interior Gateway Protocol. The possible values that can be used are: all, bgp, direct, rip, static, ospf, ospfase, ospf3, ospf3ase and ripng. By default, all protocols are enabled.`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaSetBgpInternal(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("description"); ok {
        payload["description"] = v.(string)
    }

    if v := d.Get("export_routemap_list"); len(v.([]interface{})) > 0 {
        exportroutemaplistList := v.([]interface{})
        exportroutemaplistArray := make([]interface{}, 0, len(exportroutemaplistList))
        for i := range exportroutemaplistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.preference", i)); ok {
                itemMap["preference"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.family", i)); ok {
                itemMap["family"] = v.(string)
            }
            if sv, ok := d.GetOk(fmt.Sprintf("export_routemap_list.%d.conditional_routemap", i)); ok {
                if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                    rawDict := ivList[0].(map[string]interface{})
                    conditional_routemapMap := make(map[string]interface{})
                    if sv, ok := rawDict["name"]; ok && sv.(string) != "" {
                        conditional_routemapMap["name"] = sv.(string)
                    }
                    if sv, ok := rawDict["condition"]; ok && sv.(string) != "" {
                        conditional_routemapMap["condition"] = sv.(string)
                    }
                    if len(conditional_routemapMap) > 0 {
                        itemMap["conditional-routemap"] = conditional_routemapMap
                    }
                }
            }
            if len(itemMap) > 0 {
                exportroutemaplistArray = append(exportroutemaplistArray, itemMap)
            }
        }
        if len(exportroutemaplistArray) > 0 {
            payload["export-routemap-list"] = exportroutemaplistArray
        }
    }

    if v := d.Get("import_routemap_list"); len(v.([]interface{})) > 0 {
        importroutemaplistList := v.([]interface{})
        importroutemaplistArray := make([]interface{}, 0, len(importroutemaplistList))
        for i := range importroutemaplistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("import_routemap_list.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("import_routemap_list.%d.preference", i)); ok {
                itemMap["preference"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("import_routemap_list.%d.family", i)); ok {
                itemMap["family"] = v.(string)
            }
            if len(itemMap) > 0 {
                importroutemaplistArray = append(importroutemaplistArray, itemMap)
            }
        }
        if len(importroutemaplistArray) > 0 {
            payload["import-routemap-list"] = importroutemaplistArray
        }
    }

    if v := d.Get("interface_list"); len(v.([]interface{})) > 0 {
        interfacelistList := v.([]interface{})
        interfacelistArray := make([]interface{}, 0, len(interfacelistList))
        for _, item := range interfacelistList {
            if s, ok := item.(string); ok && s != "" {
                interfacelistArray = append(interfacelistArray, s)
            }
        }
        if len(interfacelistArray) > 0 {
            payload["interface-list"] = interfacelistArray
        }
    }

    if v, ok := d.GetOk("local_address"); ok {
        payload["local-address"] = v.(string)
    }

    if v, ok := d.GetOk("med"); ok {
        payload["med"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_nexthop_self"); ok {
        payload["enable-nexthop-self"] = v.(bool)
    }

    if v, ok := d.GetOk("outdelay"); ok {
        payload["outdelay"] = v.(string)
    }

    if v := d.Get("protocol_list"); len(v.([]interface{})) > 0 {
        protocollistList := v.([]interface{})
        protocollistArray := make([]interface{}, 0, len(protocollistList))
        for _, item := range protocollistList {
            if s, ok := item.(string); ok && s != "" {
                protocollistArray = append(protocollistArray, s)
            }
        }
        if len(protocollistArray) > 0 {
            payload["protocol-list"] = protocollistArray
        }
    }

    log.Println("Execute set-bgp-internal - Payload = ", payload)

    GaiaSetBgpInternalRes, err := client.ApiCallSimple("set-bgp-internal", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetBgpInternalRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetBgpInternalRes.Success {
            errMsg = GaiaSetBgpInternalRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetBgpInternalRes.GetData()
        }

        debugLogOperation(
            "set-bgp-internal",        // resource type
            "command",                       // operation
            "set-bgp-internal",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-bgp-internal: %v", err)
    }
    if !GaiaSetBgpInternalRes.Success {
        if GaiaSetBgpInternalRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetBgpInternalRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaSetBgpInternalRes.GetData()
    if v, exists := _respData["member-id"]; exists {
        d.Set("member_id", toString(v))
    }


    d.SetId(fmt.Sprintf("set-bgp-internal-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetBgpInternal(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSetBgpInternal(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

