package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSetIsis() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetIsis,
        Read:   readGaiaSetIsis,
        Delete: deleteGaiaSetIsis,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "adjacency_check": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `Enable or disable strict protocol checks with neighbors`,
            },
            "area_list": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Add or remove an IS-IS area`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "default_metric": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Set IS-IS default metric`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "level": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the default metric for level 1`,
                        },
                        "metric": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the default metric for level 2`,
                        },
                    },
                },
            },
            "is_type": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure which level this IS-IS router should operate on`,
            },
            "max_areas": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Set the maximum number of areas`,
            },
            "overload_bit": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `Set whether this IS-IS router should include an overload bit`,
            },
            "system_id": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure IS System ID. Note that this can not be done when IS-IS is configured and running`,
            },
            "ignore_attached_bit": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `Set to ignore attached bits set by level 2 connected routers`,
            },
            "dynamic_hostname": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `Enable or disable dyanmic hostname mapping for system IDs`,
            },
            "hello": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure how IS-IS sends and receives hello messages`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "interface_point_to_point": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set hello type for point to point interfaces`,
                        },
                        "interface_broadcast": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set hello type for broadcast interfacec`,
                        },
                    },
                },
            },
            "metric_type": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure how IS-IS sends metric messages`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set metric type`,
                        },
                        "level": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set level for which this metric type applies`,
                        },
                    },
                },
            },
            "lsp": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure IS-IS LSP`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "lifetime": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the lifetime of the LSP`,
                        },
                        "mtu": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the mtu of the LSP`,
                        },
                        "refresh_interval": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the refresh interval of the LSP`,
                        },
                        "gen_interval": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the gen interval `,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "level": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the level for this gen interval`,
                                    },
                                    "max_interval": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the level 1 max gen interval `,
                                    },
                                    "initial_offset": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the level 1 intial offset`,
                                    },
                                    "second_offset": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the level 1 second offset`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "spf": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure IS-IS SPF`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "level": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The is type for spf`,
                        },
                        "max_interval": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the max interval `,
                        },
                        "initial_offset": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the intial offset`,
                        },
                        "second_offset": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the second offset`,
                        },
                    },
                },
            },
            "prc_interval": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure IS-IS PRC Interval`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "level": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The is type for the prc interval`,
                        },
                        "max_interval": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the level 1 max interval `,
                        },
                        "initial_offset": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the level 1 intial offset`,
                        },
                        "second_offset": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the level 1 second offset`,
                        },
                    },
                },
            },
            "authentication_ignore": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Ignore settings for authentication`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ignore_all": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Ignore all setting`,
                        },
                        "ignore_csnp": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Ignore csnp setting`,
                        },
                        "ignore_hello": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Ignore hello setting`,
                        },
                        "ignore_lsp": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Ignore lsp setting`,
                        },
                        "ignore_psnp": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Ignore psnp setting`,
                        },
                        "ignore_none": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Ignore none setting`,
                        },
                        "level": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The is type for the auth interval`,
                        },
                    },
                },
            },
            "authentication": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure IS-IS authentication`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Authentication type`,
                        },
                        "level": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `The is type for the auth interval`,
                        },
                        "encrypted_secret": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Sensitive:   true,
                            Description: `Encrypted secret`,
                        },
                        "secret": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Sensitive:   true,
                            Description: `Not encrypted secret`,
                        },
                        "keys": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Authentication key`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "encrypted_secret": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Sensitive:   true,
                                        Description: `Encrypted secret`,
                                    },
                                    "secret": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Sensitive:   true,
                                        Description: `Not encrypted secret`,
                                    },
                                    "resource_id": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Authentication key`,
                                    },
                                    "algorithm": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Authentication algorithm`,
                                    },
                                },
                            },
                        },
                        "active_key": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Active key`,
                        },
                    },
                },
            },
            "ipv6": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure IS-IS ipv6 options`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "overload_bit": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set whether this IS-IS router should include an overload bit`,
                        },
                        "ignore_attached_bit": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set to ignore attached bits set by level 2 connected routers`,
                        },
                        "multi_topology": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure IS-IS ipv6 multi topology`,
                        },
                        "prc_interval": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure IS-IS PRC Interval`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "level": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `The is type for the prc interval`,
                                    },
                                    "max_interval": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the level 1 max interval `,
                                    },
                                    "initial_offset": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the level 1 intial offset`,
                                    },
                                    "second_offset": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the level 1 second offset`,
                                    },
                                },
                            },
                        },
                        "spf": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Configure IS-IS SPF`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "level": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `The is type for spf`,
                                    },
                                    "max_interval": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the max interval `,
                                    },
                                    "initial_offset": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the intial offset`,
                                    },
                                    "second_offset": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the second offset`,
                                    },
                                },
                            },
                        },
                        "default_metric": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set IS-IS default metric`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "level": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the default metric for level 1`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Set the default metric for level 2`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
        },
    }
}

func createGaiaSetIsis(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("adjacency_check"); ok {
        payload["adjacency-check"] = v.(bool)
    }

    if v := d.Get("area_list"); len(v.([]interface{})) > 0 {
        arealistList := v.([]interface{})
        arealistArray := make([]interface{}, 0, len(arealistList))
        for _, item := range arealistList {
            if s, ok := item.(string); ok && s != "" {
                arealistArray = append(arealistArray, s)
            }
        }
        if len(arealistArray) > 0 {
            payload["area-list"] = arealistArray
        }
    }

    if v := d.Get("default_metric"); len(v.([]interface{})) > 0 {
        defaultmetricList := v.([]interface{})
        defaultmetricArray := make([]interface{}, 0, len(defaultmetricList))
        for i := range defaultmetricList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("default_metric.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("default_metric.%d.metric", i)); ok {
                itemMap["metric"] = v.(string)
            }
            if len(itemMap) > 0 {
                defaultmetricArray = append(defaultmetricArray, itemMap)
            }
        }
        if len(defaultmetricArray) > 0 {
            payload["default-metric"] = defaultmetricArray
        }
    }

    if v, ok := d.GetOk("is_type"); ok {
        payload["is-type"] = v.(string)
    }

    if v, ok := d.GetOk("max_areas"); ok {
        payload["max-areas"] = v.(string)
    }

    if v, ok := d.GetOkExists("overload_bit"); ok {
        payload["overload-bit"] = v.(bool)
    }

    if v, ok := d.GetOk("system_id"); ok {
        payload["system-id"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_attached_bit"); ok {
        payload["ignore-attached-bit"] = v.(bool)
    }

    if v, ok := d.GetOkExists("dynamic_hostname"); ok {
        payload["dynamic-hostname"] = v.(bool)
    }

    if v := d.Get("hello"); len(v.([]interface{})) > 0 {
        _ = v
        helloMap := make(map[string]interface{})
        if v, ok := d.GetOk("hello.0.interface_point_to_point"); ok {
            helloMap["interface-point-to-point"] = v.(string)
        }
        if v, ok := d.GetOk("hello.0.interface_broadcast"); ok {
            helloMap["interface-broadcast"] = v.(string)
        }
        if len(helloMap) > 0 {
            payload["hello"] = helloMap
        }
    }

    if v := d.Get("metric_type"); len(v.([]interface{})) > 0 {
        metrictypeList := v.([]interface{})
        metrictypeArray := make([]interface{}, 0, len(metrictypeList))
        for i := range metrictypeList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("metric_type.%d.type", i)); ok {
                itemMap["type"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("metric_type.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if len(itemMap) > 0 {
                metrictypeArray = append(metrictypeArray, itemMap)
            }
        }
        if len(metrictypeArray) > 0 {
            payload["metric-type"] = metrictypeArray
        }
    }

    if v := d.Get("lsp"); len(v.([]interface{})) > 0 {
        _ = v
        lspMap := make(map[string]interface{})
        if v, ok := d.GetOk("lsp.0.lifetime"); ok {
            lspMap["lifetime"] = v.(string)
        }
        if v, ok := d.GetOk("lsp.0.mtu"); ok {
            lspMap["mtu"] = v.(string)
        }
        if v, ok := d.GetOk("lsp.0.refresh_interval"); ok {
            lspMap["refresh-interval"] = v.(string)
        }
        if v, ok := d.GetOk("lsp.0.gen_interval"); ok {
            genintervalList := v.([]interface{})
            genintervalArray := make([]interface{}, 0, len(genintervalList))
            for i := range genintervalList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("lsp.0.gen_interval.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("lsp.0.gen_interval.%d.max_interval", i)); ok {
                    itemMap["max-interval"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("lsp.0.gen_interval.%d.initial_offset", i)); ok {
                    itemMap["initial-offset"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("lsp.0.gen_interval.%d.second_offset", i)); ok {
                    itemMap["second-offset"] = v.(string)
                }
                if len(itemMap) > 0 {
                    genintervalArray = append(genintervalArray, itemMap)
                }
            }
            if len(genintervalArray) > 0 {
                lspMap["gen-interval"] = genintervalArray
            }
        }
        if len(lspMap) > 0 {
            payload["lsp"] = lspMap
        }
    }

    if v := d.Get("spf"); len(v.([]interface{})) > 0 {
        spfList := v.([]interface{})
        spfArray := make([]interface{}, 0, len(spfList))
        for i := range spfList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("spf.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("spf.%d.max_interval", i)); ok {
                itemMap["max-interval"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("spf.%d.initial_offset", i)); ok {
                itemMap["initial-offset"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("spf.%d.second_offset", i)); ok {
                itemMap["second-offset"] = v.(string)
            }
            if len(itemMap) > 0 {
                spfArray = append(spfArray, itemMap)
            }
        }
        if len(spfArray) > 0 {
            payload["spf"] = spfArray
        }
    }

    if v := d.Get("prc_interval"); len(v.([]interface{})) > 0 {
        prcintervalList := v.([]interface{})
        prcintervalArray := make([]interface{}, 0, len(prcintervalList))
        for i := range prcintervalList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("prc_interval.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("prc_interval.%d.max_interval", i)); ok {
                itemMap["max-interval"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("prc_interval.%d.initial_offset", i)); ok {
                itemMap["initial-offset"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("prc_interval.%d.second_offset", i)); ok {
                itemMap["second-offset"] = v.(string)
            }
            if len(itemMap) > 0 {
                prcintervalArray = append(prcintervalArray, itemMap)
            }
        }
        if len(prcintervalArray) > 0 {
            payload["prc-interval"] = prcintervalArray
        }
    }

    if v := d.Get("authentication_ignore"); len(v.([]interface{})) > 0 {
        authenticationignoreList := v.([]interface{})
        authenticationignoreArray := make([]interface{}, 0, len(authenticationignoreList))
        for i := range authenticationignoreList {
            itemMap := make(map[string]interface{})
            if v := d.Get(fmt.Sprintf("authentication_ignore.%d.ignore_all", i)).(bool); v {
                itemMap["ignore-all"] = v
            }
            if v := d.Get(fmt.Sprintf("authentication_ignore.%d.ignore_csnp", i)).(bool); v {
                itemMap["ignore-csnp"] = v
            }
            if v := d.Get(fmt.Sprintf("authentication_ignore.%d.ignore_hello", i)).(bool); v {
                itemMap["ignore-hello"] = v
            }
            if v := d.Get(fmt.Sprintf("authentication_ignore.%d.ignore_lsp", i)).(bool); v {
                itemMap["ignore-lsp"] = v
            }
            if v := d.Get(fmt.Sprintf("authentication_ignore.%d.ignore_psnp", i)).(bool); v {
                itemMap["ignore-psnp"] = v
            }
            if v := d.Get(fmt.Sprintf("authentication_ignore.%d.ignore_none", i)).(bool); v {
                itemMap["ignore-none"] = v
            }
            if v, ok := d.GetOk(fmt.Sprintf("authentication_ignore.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if len(itemMap) > 0 {
                authenticationignoreArray = append(authenticationignoreArray, itemMap)
            }
        }
        if len(authenticationignoreArray) > 0 {
            payload["authentication-ignore"] = authenticationignoreArray
        }
    }

    if v := d.Get("authentication"); len(v.([]interface{})) > 0 {
        authenticationList := v.([]interface{})
        authenticationArray := make([]interface{}, 0, len(authenticationList))
        for i := range authenticationList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("authentication.%d.type", i)); ok {
                itemMap["type"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("authentication.%d.level", i)); ok {
                itemMap["level"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("authentication.%d.encrypted_secret", i)); ok {
                itemMap["encrypted-secret"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("authentication.%d.secret", i)); ok {
                itemMap["secret"] = v.(string)
            }
            if sv := d.Get(fmt.Sprintf("authentication.%d.keys", i)); len(sv.([]interface{})) > 0 {
                keysList := sv.([]interface{})
                keysArr := make([]interface{}, 0, len(keysList))
                for j := range keysList {
                    innerMap := make(map[string]interface{})
                    if iv, ok := d.GetOk(fmt.Sprintf("authentication.%d.keys.%d.encrypted_secret", i, j)); ok {
                        innerMap["encrypted-secret"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("authentication.%d.keys.%d.secret", i, j)); ok {
                        innerMap["secret"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("authentication.%d.keys.%d.resource_id", i, j)); ok {
                        innerMap["id"] = iv.(int)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("authentication.%d.keys.%d.algorithm", i, j)); ok {
                        innerMap["algorithm"] = iv.(string)
                    }
                    if len(innerMap) > 0 {
                        keysArr = append(keysArr, innerMap)
                    }
                }
                if len(keysArr) > 0 {
                    itemMap["keys"] = keysArr
                }
            }
            if v, ok := d.GetOk(fmt.Sprintf("authentication.%d.active_key", i)); ok {
                itemMap["active-key"] = v.(string)
            }
            if len(itemMap) > 0 {
                authenticationArray = append(authenticationArray, itemMap)
            }
        }
        if len(authenticationArray) > 0 {
            payload["authentication"] = authenticationArray
        }
    }

    if v := d.Get("ipv6"); len(v.([]interface{})) > 0 {
        _ = v
        ipv6Map := make(map[string]interface{})
        if v, ok := d.GetOkExists("ipv6.0.overload_bit"); ok && v.(bool) {
            ipv6Map["overload-bit"] = v.(bool)
        }
        if v, ok := d.GetOkExists("ipv6.0.ignore_attached_bit"); ok && v.(bool) {
            ipv6Map["ignore-attached-bit"] = v.(bool)
        }
        if v, ok := d.GetOk("ipv6.0.multi_topology"); ok {
            ipv6Map["multi-topology"] = v.(string)
        }
        if v, ok := d.GetOk("ipv6.0.prc_interval"); ok {
            prcintervalList := v.([]interface{})
            prcintervalArray := make([]interface{}, 0, len(prcintervalList))
            for i := range prcintervalList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.prc_interval.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.prc_interval.%d.max_interval", i)); ok {
                    itemMap["max-interval"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.prc_interval.%d.initial_offset", i)); ok {
                    itemMap["initial-offset"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.prc_interval.%d.second_offset", i)); ok {
                    itemMap["second-offset"] = v.(string)
                }
                if len(itemMap) > 0 {
                    prcintervalArray = append(prcintervalArray, itemMap)
                }
            }
            if len(prcintervalArray) > 0 {
                ipv6Map["prc-interval"] = prcintervalArray
            }
        }
        if v, ok := d.GetOk("ipv6.0.spf"); ok {
            spfList := v.([]interface{})
            spfArray := make([]interface{}, 0, len(spfList))
            for i := range spfList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.spf.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.spf.%d.max_interval", i)); ok {
                    itemMap["max-interval"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.spf.%d.initial_offset", i)); ok {
                    itemMap["initial-offset"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.spf.%d.second_offset", i)); ok {
                    itemMap["second-offset"] = v.(string)
                }
                if len(itemMap) > 0 {
                    spfArray = append(spfArray, itemMap)
                }
            }
            if len(spfArray) > 0 {
                ipv6Map["spf"] = spfArray
            }
        }
        if v, ok := d.GetOk("ipv6.0.default_metric"); ok {
            defaultmetricList := v.([]interface{})
            defaultmetricArray := make([]interface{}, 0, len(defaultmetricList))
            for i := range defaultmetricList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.default_metric.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("ipv6.0.default_metric.%d.metric", i)); ok {
                    itemMap["metric"] = v.(string)
                }
                if len(itemMap) > 0 {
                    defaultmetricArray = append(defaultmetricArray, itemMap)
                }
            }
            if len(defaultmetricArray) > 0 {
                ipv6Map["default-metric"] = defaultmetricArray
            }
        }
        if len(ipv6Map) > 0 {
            payload["ipv6"] = ipv6Map
        }
    }

    log.Println("Execute set-isis - Payload = ", payload)

    GaiaSetIsisRes, err := client.ApiCallSimple("set-isis", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetIsisRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetIsisRes.Success {
            errMsg = GaiaSetIsisRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetIsisRes.GetData()
        }

        debugLogOperation(
            "set-isis",        // resource type
            "command",                       // operation
            "set-isis",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-isis: %v", err)
    }
    if !GaiaSetIsisRes.Success {
        if GaiaSetIsisRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetIsisRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }



    d.SetId(fmt.Sprintf("set-isis-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetIsis(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSetIsis(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

