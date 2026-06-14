package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSetBgp() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetBgp,
        Read:   readGaiaSetBgp,
        Delete: deleteGaiaSetBgp,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "as": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Autonomous System Number<br>The value can be one of the following:<br>'off'<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535<br><br>WARNING: Removing the AS number will result in all BGP configurations and any associated route-redistribution or inbound-route-filter configurations being removed.`,
            },
            "cluster_id": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Specifies the cluster-id used for route reflection.`,
            },
            "enable_communities": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `This option controls whether or not community information is included in BGP advertisements. This option must be enabled in order to configure the routing policy to filter incoming or outgoing advertisements based on community information.`,
            },
            "confederation": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure BGP Confederation parameters.<br><br>A BGP Confederation is a single large AS that is divided into sub-AS's called Routing Domains. The Routing Domains are only visible within the Confederation; to the outside world, the entire Confederation appears as one AS. Confederations improve BGP performance for large AS's by reducing IBGP mesh size. IBGP is used within each Routing Domain, but between them, a modified form of eBGP is used which preserves route metrics and other BGP attributes.<br><br>Both the Confederation identifier and Routing Domain identifier must be configured before BGP can be configured on this router when Confederations are in use. An AS cannot be configured at the same time as the Confederation equivalents.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "aspath_loops_permitted": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the number of times the Local AS can appear in an AS path for routes learned via BGP. Routes with numbers higher than the configured value are rejected. The default value is 1.`,
                        },
                        "identifier": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the identifier for the entire Confederation.<br><br>The Confederation identifier is used as the AS number in external BGP sessions, so it must be a globally unique, normally assigned AS number.<br><br>Both the Confederation identifer and Routing Domain identifier must be configured before BGP can be configured on this router when Confederations are in use. An AS cannot be configured at the same time as the Confederation equivalents.<br><br>The value can be one of the following:<br>'off'<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
                        },
                    },
                },
            },
            "dampening": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Weighted-route dampening minimizes the propagation of flapping routes across an internetwork.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "enabled": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Enable/disable weighted-route dampening.`,
                        },
                        "keep_history": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the period over which route flapping history for a given route is maintained. The default value is 1800.`,
                        },
                        "max_flap": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the upper limit of the instability accepted. The value must be higher than one plus the suppress-above value. The default value is 16.`,
                        },
                        "reachable_decay": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the length of time (seconds) it takes for the instability metric to reach one half of its current value when the route is reachable. The default value is 300.`,
                        },
                        "reuse_below": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the value of the instability metric at which a suppressed but reachable route becomes unsuppressed. The value must be less than the suppress-above value. The default value is 2.`,
                        },
                        "suppress_above": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the value of the instability metric at which a route is suppressed. While suppressed, the route is neither installed, nor advertised as reachable. The default value is 3.`,
                        },
                        "unreachable_decay": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the rate at which the instability metric is decayed when a route is unreachable. This value must be equal to or greater than the reachable decay. The default value is 900.`,
                        },
                    },
                },
            },
            "default_med": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Defines the Multi-Exit Discriminator metric (MED) used when advertising routes through BGP. If no value is specified, no metric is propagated. Any metrics configured in peer, Route Map, or route redistribution configurations will override the value configured here.`,
            },
            "default_route_gateway": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `A default route is generated if any BGP peer is up. This route has a higher rank than the default configured through static route configuration. If a specific BGP peer should not be considered for generating the default route, it should be explicitly suppressed via the peer-specific 'suppress-default-originate' configuration.`,
            },
            "enable_ecmp": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `Enables or disables ECMP (Equal-Cost Multi-Path) routing for IPv4 BGP routes.`,
            },
            "graceful_restart": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configures global settings for BGP Graceful Restart.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "restart_time": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the time (seconds) that BGP peers of this router should keep the routes advertised to them while this router restarts. The default value is 360.`,
                        },
                        "selection_deferral_time": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the time (seconds) that this router will wait for the End-of-RIB notification from each of its BGP peers after a restart. The default value is 360.`,
                        },
                    },
                },
            },
            "ping": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configures global settings for BGP ping.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "count": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the number of failed pings to an individual BGP peer with ping enabled before BGP will drop that peer. This value is common across all peers. The default value is 3.`,
                        },
                        "interval": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Set the interval between pings sent to all BGP peers with ping enabled. The default value is 2.`,
                        },
                    },
                },
            },
            "routing_domain": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Configure Routing Domain parameters.<br><br>In Confederation mode, the Routing Domain is the Confederation sub-AS. It acts as an independent AS within the Confederation, but is not visible outside the Confederation. The Routing Domain identifier (RDI) is the equivalent of its AS number. <br><br>Both the Confederation identifier and Routing Domain identifier must be configured before BGP can be configured on this router when Confederations are in use. An AS cannot be configured at the same time as the Confederation equivalents.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "aspath_loops_permitted": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Specifies the number of times the Local AS can appear in an AS path for routes learned via BGP. Routes with numbers higher than the configured value are rejected. The default value is 1.`,
                        },
                        "identifier": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `In Confederation mode, the Routing Domain identifier (RDI) identifies the Routing Domain, or Confederation sub-AS, to which this router belongs. The RDI is used as the AS number for peers within the Confederation, while the Confederation identifier is used outside the Confederation. The RDI does not have to be globally unique, since it is never used outside the Confederation.<br><br>Both the Confederation identifier and Routing Domain identifier must be configured before BGP can be configured on this router when Confederations are in use. An AS cannot be configured at the same time as the Confederation equivalents.<br><br>The value can be one of the following:<br>'off'<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
                        },
                    },
                },
            },
            "enable_synchronization": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `Enabling this option directs Internal BGP (IBGP) peers to check for a matching route from IGP protocols before installing a route.`,
            },
        },
    }
}

func createGaiaSetBgp(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("as"); ok {
        payload["as"] = v.(string)
    }

    if v, ok := d.GetOk("cluster_id"); ok {
        payload["cluster-id"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_communities"); ok {
        payload["enable-communities"] = v.(bool)
    }

    if v := d.Get("confederation"); len(v.([]interface{})) > 0 {
        _ = v
        confederationMap := make(map[string]interface{})
        if v, ok := d.GetOk("confederation.0.aspath_loops_permitted"); ok {
            confederationMap["aspath-loops-permitted"] = v.(string)
        }
        if v, ok := d.GetOk("confederation.0.identifier"); ok {
            confederationMap["identifier"] = v.(string)
        }
        if len(confederationMap) > 0 {
            payload["confederation"] = confederationMap
        }
    }

    if v := d.Get("dampening"); len(v.([]interface{})) > 0 {
        _ = v
        dampeningMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("dampening.0.enabled"); ok && v.(bool) {
            dampeningMap["enabled"] = v.(bool)
        }
        if v, ok := d.GetOk("dampening.0.keep_history"); ok {
            dampeningMap["keep-history"] = v.(string)
        }
        if v, ok := d.GetOk("dampening.0.max_flap"); ok {
            dampeningMap["max-flap"] = v.(string)
        }
        if v, ok := d.GetOk("dampening.0.reachable_decay"); ok {
            dampeningMap["reachable-decay"] = v.(string)
        }
        if v, ok := d.GetOk("dampening.0.reuse_below"); ok {
            dampeningMap["reuse-below"] = v.(string)
        }
        if v, ok := d.GetOk("dampening.0.suppress_above"); ok {
            dampeningMap["suppress-above"] = v.(string)
        }
        if v, ok := d.GetOk("dampening.0.unreachable_decay"); ok {
            dampeningMap["unreachable-decay"] = v.(string)
        }
        if len(dampeningMap) > 0 {
            payload["dampening"] = dampeningMap
        }
    }

    if v, ok := d.GetOk("default_med"); ok {
        payload["default-med"] = v.(string)
    }

    if v, ok := d.GetOk("default_route_gateway"); ok {
        payload["default-route-gateway"] = v.(string)
    }

    if v, ok := d.GetOkExists("enable_ecmp"); ok {
        payload["enable-ecmp"] = v.(bool)
    }

    if v := d.Get("graceful_restart"); len(v.([]interface{})) > 0 {
        _ = v
        gracefulrestartMap := make(map[string]interface{})
        if v, ok := d.GetOk("graceful_restart.0.restart_time"); ok {
            gracefulrestartMap["restart-time"] = v.(string)
        }
        if v, ok := d.GetOk("graceful_restart.0.selection_deferral_time"); ok {
            gracefulrestartMap["selection-deferral-time"] = v.(string)
        }
        if len(gracefulrestartMap) > 0 {
            payload["graceful-restart"] = gracefulrestartMap
        }
    }

    if v := d.Get("ping"); len(v.([]interface{})) > 0 {
        _ = v
        pingMap := make(map[string]interface{})
        if v, ok := d.GetOk("ping.0.count"); ok {
            pingMap["count"] = v.(string)
        }
        if v, ok := d.GetOk("ping.0.interval"); ok {
            pingMap["interval"] = v.(string)
        }
        if len(pingMap) > 0 {
            payload["ping"] = pingMap
        }
    }

    if v := d.Get("routing_domain"); len(v.([]interface{})) > 0 {
        _ = v
        routingdomainMap := make(map[string]interface{})
        if v, ok := d.GetOk("routing_domain.0.aspath_loops_permitted"); ok {
            routingdomainMap["aspath-loops-permitted"] = v.(string)
        }
        if v, ok := d.GetOk("routing_domain.0.identifier"); ok {
            routingdomainMap["identifier"] = v.(string)
        }
        if len(routingdomainMap) > 0 {
            payload["routing-domain"] = routingdomainMap
        }
    }

    if v, ok := d.GetOkExists("enable_synchronization"); ok {
        payload["enable-synchronization"] = v.(bool)
    }

    log.Println("Execute set-bgp - Payload = ", payload)

    // Pre-create cleanup: ensure no stale BGP state from a prior failed run.
    client.ApiCallSimple("set-bgp", map[string]interface{}{
        "as": "off",
    })
    client.ApiCallSimple("set-bgp", map[string]interface{}{
        "confederation": map[string]interface{}{"identifier": "off"},
    })
    GaiaSetBgpRes, err := client.ApiCallSimple("set-bgp", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetBgpRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetBgpRes.Success {
            errMsg = GaiaSetBgpRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetBgpRes.GetData()
        }

        debugLogOperation(
            "set-bgp",        // resource type
            "command",                       // operation
            "set-bgp",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-bgp: %v", err)
    }
    if !GaiaSetBgpRes.Success {
        if GaiaSetBgpRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetBgpRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }



    // Flush configuration so subsequent sessions see this change immediately.
    if _, saveErr := client.ApiCallSimple("save-to-network", map[string]interface{}{}); saveErr != nil {
        log.Printf("[WARN] save-to-network after successful create failed: %v", saveErr)
    }

    d.SetId(fmt.Sprintf("set-bgp-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetBgp(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSetBgp(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)
    // Clear confederation identifier before clearing AS (device enforces ordering).
    client.ApiCallSimple("set-bgp", map[string]interface{}{
        "confederation": map[string]interface{}{"identifier": "off"},
    })
    // Clear AS.
    client.ApiCallSimple("set-bgp", map[string]interface{}{
        "as": "off",
    })
    d.SetId("")
    return nil
}

