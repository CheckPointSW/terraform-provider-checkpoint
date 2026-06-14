package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSetBgpConfederation() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetBgpConfederation,
        Read:   readGaiaSetBgpConfederation,
        Delete: deleteGaiaSetBgpConfederation,
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
            "member_as": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Specify the Routing Domain identifier of the Confederation peer group to configure.<br><br>If the peer group specified is the local Routing Domain, it will run IBGP in a full mesh (just as an internal peer group normally would in non-Confederation mode). Otherwise, if an external Routing Domain within the Confederation is specified, the peer group will run a modified version of eBGP, which preserves route metrics and other BGP attributes.<br><br>The value can be one of the following:<br>An integer from 1-4294967295<br>A float from 0.1-65535.65535`,
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
        },
    }
}

func createGaiaSetBgpConfederation(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("description"); ok {
        payload["description"] = v.(string)
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

    if v, ok := d.GetOk("member_as"); ok {
        payload["member-as"] = v.(string)
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

    log.Println("Execute set-bgp-confederation - Payload = ", payload)

    GaiaSetBgpConfederationRes, err := client.ApiCallSimple("set-bgp-confederation", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetBgpConfederationRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetBgpConfederationRes.Success {
            errMsg = GaiaSetBgpConfederationRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetBgpConfederationRes.GetData()
        }

        debugLogOperation(
            "set-bgp-confederation",        // resource type
            "command",                       // operation
            "set-bgp-confederation",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-bgp-confederation: %v", err)
    }
    if !GaiaSetBgpConfederationRes.Success {
        if GaiaSetBgpConfederationRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetBgpConfederationRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }



    // Flush configuration so subsequent sessions see this change immediately.
    if _, saveErr := client.ApiCallSimple("save-to-network", map[string]interface{}{}); saveErr != nil {
        log.Printf("[WARN] save-to-network after successful create failed: %v", saveErr)
    }

    d.SetId(fmt.Sprintf("set-bgp-confederation-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetBgpConfederation(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSetBgpConfederation(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

