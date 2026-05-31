package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSetBgpExternal() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetBgpExternal,
        Read:   readGaiaSetBgpExternal,
        Delete: deleteGaiaSetBgpExternal,
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
                Description: `Enable/disable the peer group for the specified AS.`,
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
            "inject_routemap_list": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure conditional route injection for a routemap`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The name of the inject routemap`,
                        },
                        "preference": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Preference for the routemap. Routemaps are evaluated in order of increasing preference value.`,
                        },
                        "any_pass_routemap": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The name of the any-pass-routemap that will be the condition for injection`,
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
            "local_address": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configures the address to be used on the local end of the TCP connection.`,
            },
            "outdelay": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Specifies the length of time (seconds) that a route must be present in the routing database before it is redistributed to BGP.`,
            },
            "remote_as": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `The Autonomous System number of the peer group to configure.The value can be one of the following:<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
            },
        },
    }
}

func createGaiaSetBgpExternal(d *schema.ResourceData, m interface{}) error {
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

    if v := d.Get("inject_routemap_list"); len(v.([]interface{})) > 0 {
        injectroutemaplistList := v.([]interface{})
        injectroutemaplistArray := make([]interface{}, 0, len(injectroutemaplistList))
        for i := range injectroutemaplistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.preference", i)); ok {
                itemMap["preference"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.any_pass_routemap", i)); ok {
                itemMap["any-pass-routemap"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("inject_routemap_list.%d.family", i)); ok {
                itemMap["family"] = v.(string)
            }
            if len(itemMap) > 0 {
                injectroutemaplistArray = append(injectroutemaplistArray, itemMap)
            }
        }
        if len(injectroutemaplistArray) > 0 {
            payload["inject-routemap-list"] = injectroutemaplistArray
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

    if v, ok := d.GetOk("local_address"); ok {
        payload["local-address"] = v.(string)
    }

    if v, ok := d.GetOk("outdelay"); ok {
        payload["outdelay"] = v.(string)
    }

    if v, ok := d.GetOk("remote_as"); ok {
        payload["remote-as"] = v.(string)
    }

    log.Println("Execute set-bgp-external - Payload = ", payload)

    GaiaSetBgpExternalRes, err := client.ApiCallSimple("set-bgp-external", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetBgpExternalRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetBgpExternalRes.Success {
            errMsg = GaiaSetBgpExternalRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetBgpExternalRes.GetData()
        }

        debugLogOperation(
            "set-bgp-external",        // resource type
            "command",                       // operation
            "set-bgp-external",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-bgp-external: %v", err)
    }
    if !GaiaSetBgpExternalRes.Success {
        if GaiaSetBgpExternalRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetBgpExternalRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }



    d.SetId(fmt.Sprintf("set-bgp-external-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetBgpExternal(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSetBgpExternal(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

