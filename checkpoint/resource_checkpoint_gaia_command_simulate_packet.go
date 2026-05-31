package checkpoint

import (
        "context"
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaSimulatePacket() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSimulatePacket,
        Read:   readGaiaSimulatePacket,
        Delete: deleteGaiaSimulatePacket,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "source_ip": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Source IP, should match selected \"ip-version\" (which defaults to 4)`,
            },
            "destination_ip": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Destination IP, should match selected \"ip-version\" (which defaults to 4)`,
            },
            "ip_protocol": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `ip-protocol either in integer form: based on IANA Protocol Number in decimal format                      or a string of one of the following protocols: [UDP, TCP, ICMP]`,
            },
            "protocol_options": {
                Type:        schema.TypeList,
                Required:    true,
                ForceNew:    true,
                Description: `Protocol options required for the selected ip-protocol. please note, only the relevant protocol's options should be filled.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "tcp": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `TCP specific required options. required if ip-protocol is \"TCP\" or its IANA Protocol Number \"6\".`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "source_port": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Source port in the Decimal format. if not specified will default to 12345`,
                                    },
                                    "destination_port": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Destination port in the Decimal format. This parameter is mandatory for the TCP (6) and UDP (17) protocols.`,
                                    },
                                },
                            },
                        },
                        "udp": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `UDP specific required options. required if ip-protocol is \"UDP\" or its IANA Protocol Number \"17\".`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "source_port": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Source port in the Decimal format. if not specified will default to 12345`,
                                    },
                                    "destination_port": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Destination port in the Decimal format. This parameter is mandatory for the TCP (6) and UDP (17) protocols.`,
                                    },
                                },
                            },
                        },
                        "icmp": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `icmp specific required options. required if ip-protocol is \"icmp\" or its IANA Protocol Number in IPv4: \"1\" or in IPv6: \"58\".`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "type": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `a string of the desired icmp type in decimal format as seen in IANA icmp parameters`,
                                    },
                                    "code": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `a string of the desired icmp code in decimal format as seen in IANA icmp parameters,  will defualt to 0.`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "incoming_interface": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Incoming interface name for the packet, identified by the name. The simulated connection is inbound, in order to simulate a local outgoing connection, set incoming-interface to localhost`,
            },
            "application": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: `Name of the Application/Category as defined in SmartConsole.  You can specify multiple applications.`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "protocol": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Protocol to match for services that have \"Protocol Signature\" enabled`,
            },
            "check_access_rule_uid": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Rule uid to check why the packet didn't match this rule`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "policy_details": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "policy_name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "installed_on": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "access": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "action": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "check_access_rule": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "match": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "no_match_code": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "no_match_reason": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "no_match_column": {
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
                        "required_classifications": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "active_classifications": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "implied_rule_first_match": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "implied_rule_first_name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "layers": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "layer_uuid": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ordered_layer_number": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "match_status": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "action": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "rule": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "parent_rule": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "possible_rules": {
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
                        "classification_objects": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "application": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_network": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_network": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "service": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_access_role": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_access_role": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "user_authentication": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_security_zone": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_security_zone": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "vpn_source": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "vpn_destination": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "client_authentication": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "logical_server": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "xff_source_access_role": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_dynamic_object": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_dynamic_object": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_domain_object": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_domain_object": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_user_at_location": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_user_at_location": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_user_limitation": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_user_limitation": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "protocol": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "service_application": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "mab_protection_level": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "mab_application": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "scada_application": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "content_and_file": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "file": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "content": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "direction": {
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
                    },
                },
            },
            "nat": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "rules": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "rule_uuid": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "action": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "source": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_port": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_port": {
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
                        "action_messages": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "classification_objects": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "source_access_roles": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_access_roles": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_dnd": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_dnd": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_security_zones": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_security_zones": {
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
                    },
                },
            },
            "errors": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "nat": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "access": {
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
        },
    }
}

func createGaiaSimulatePacket(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("ip_version"); ok {
        payload["ip-version"] = v.(string)
    }

    if v, ok := d.GetOk("source_ip"); ok {
        payload["source-ip"] = v.(string)
    }

    if v, ok := d.GetOk("destination_ip"); ok {
        payload["destination-ip"] = v.(string)
    }

    if v, ok := d.GetOk("ip_protocol"); ok {
        payload["ip-protocol"] = v.(int)
    }

    if v := d.Get("protocol_options"); len(v.([]interface{})) > 0 {
        _ = v
        protocoloptionsMap := make(map[string]interface{})
        if v, ok := d.GetOk("protocol_options.0.tcp"); ok {
            _ = v
            TCPMap := make(map[string]interface{})
            if v, ok := d.GetOk("protocol_options.0.tcp.0.source_port"); ok {
                TCPMap["source-port"] = v.(string)
            }
            if v, ok := d.GetOk("protocol_options.0.tcp.0.destination_port"); ok {
                TCPMap["destination-port"] = v.(string)
            }
            if len(TCPMap) > 0 {
                protocoloptionsMap["TCP"] = TCPMap
            }
        }
        if v, ok := d.GetOk("protocol_options.0.udp"); ok {
            _ = v
            UDPMap := make(map[string]interface{})
            if v, ok := d.GetOk("protocol_options.0.udp.0.source_port"); ok {
                UDPMap["source-port"] = v.(string)
            }
            if v, ok := d.GetOk("protocol_options.0.udp.0.destination_port"); ok {
                UDPMap["destination-port"] = v.(string)
            }
            if len(UDPMap) > 0 {
                protocoloptionsMap["UDP"] = UDPMap
            }
        }
        if v, ok := d.GetOk("protocol_options.0.icmp"); ok {
            _ = v
            icmpMap := make(map[string]interface{})
            if v, ok := d.GetOk("protocol_options.0.icmp.0.type"); ok {
                icmpMap["type"] = v.(string)
            }
            if v, ok := d.GetOk("protocol_options.0.icmp.0.code"); ok {
                icmpMap["code"] = v.(string)
            }
            if len(icmpMap) > 0 {
                protocoloptionsMap["icmp"] = icmpMap
            }
        }
        if len(protocoloptionsMap) > 0 {
            payload["protocol-options"] = protocoloptionsMap
        }
    }

    if v, ok := d.GetOk("incoming_interface"); ok {
        payload["incoming-interface"] = v.(string)
    }

    if v := d.Get("application"); len(v.(*schema.Set).List()) > 0 {
        payload["application"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("protocol"); ok {
        payload["protocol"] = v.(string)
    }

    if v, ok := d.GetOk("check_access_rule_uid"); ok {
        payload["check-access-rule-uid"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute simulate-packet - Payload = ", payload)

    GaiaSimulatePacketRes, err := client.ApiCallSimple("simulate-packet", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSimulatePacketRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSimulatePacketRes.Success {
            errMsg = GaiaSimulatePacketRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSimulatePacketRes.GetData()
        }

        debugLogOperation(
            "simulate-packet",        // resource type
            "command",                       // operation
            "simulate-packet",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute simulate-packet: %v", err)
    }
    if !GaiaSimulatePacketRes.Success {
        if GaiaSimulatePacketRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSimulatePacketRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "simulate-packet", GaiaSimulatePacketRes, true, 0)
    if err != nil {
        return fmt.Errorf("simulate-packet task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        msg := taskRes.Message
        if msg == "" {
            msg = fmt.Sprintf("simulate-packet task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(msg)
    }

    _taskDetailsRes, _tdErr := client.ApiCallSimple("show-task", map[string]interface{}{"task-id": taskRes.TaskID})
    var _respData map[string]interface{}
    if _tdErr == nil && _taskDetailsRes.Success {
        _td := _taskDetailsRes.GetData()
        if _tasks, _ok := _td["tasks"].([]interface{}); _ok && len(_tasks) > 0 {
            if _task, _ok := _tasks[0].(map[string]interface{}); _ok {
                if _details, _ok := _task["task-details"].([]interface{}); _ok && len(_details) > 0 {
                    if _d0, _ok := _details[0].(map[string]interface{}); _ok {
                        _respData = _d0
                    }
                }
            }
        }
    }
    if _respData == nil {
        _respData = GaiaSimulatePacketRes.GetData()
    }
    if _tasks, _ok := _respData["tasks"].([]interface{}); _ok && len(_tasks) > 0 {
        if _task, _ok := _tasks[0].(map[string]interface{}); _ok {
            if _details, _ok := _task["task-details"].([]interface{}); _ok && len(_details) > 0 {
                if _d0, _ok := _details[0].(map[string]interface{}); _ok {
                    _respData = _d0
                }
            }
        }
    }
    if v, exists := _respData["policy-details"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            d.Set("policy_details", []interface{}{map[string]interface{}{
                "policy_name": toString(m["policy-name"]),
                "installed_on": toString(m["installed-on"]),
            }})
        }
    }
    if v, exists := _respData["access"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            accessMap := map[string]interface{}{
                "action":                   toString(m["action"]),
                "implied_rule_first_match":  toString(m["implied-rule-first-match"]) == "true",
                "implied_rule_first_name":   toString(m["implied-rule-first-name"]),
            }
            if arr, ok := m["required-classifications"].([]interface{}); ok {
                strs := make([]interface{}, len(arr))
                for i, e := range arr { strs[i] = fmt.Sprintf("%v", e) }
                accessMap["required_classifications"] = strs
            }
            if arr, ok := m["active-classifications"].([]interface{}); ok {
                strs := make([]interface{}, len(arr))
                for i, e := range arr { strs[i] = fmt.Sprintf("%v", e) }
                accessMap["active_classifications"] = strs
            }
            if arr, ok := m["classification-objects"].([]interface{}); ok {
                strs := make([]interface{}, len(arr))
                for i, e := range arr { strs[i] = fmt.Sprintf("%v", e) }
                accessMap["classification_objects"] = strs
            }
            d.Set("access", []interface{}{accessMap})
        }
    }
    if v, exists := _respData["NAT"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            natMap := map[string]interface{}{}
            if arr, ok := m["rules"].([]interface{}); ok {
                strs := make([]interface{}, len(arr))
                for i, e := range arr { strs[i] = fmt.Sprintf("%v", e) }
                natMap["rules"] = strs
            }
            if arr, ok := m["action-messages"].([]interface{}); ok {
                strs := make([]interface{}, len(arr))
                for i, e := range arr { strs[i] = fmt.Sprintf("%v", e) }
                natMap["action_messages"] = strs
            }
            if arr, ok := m["classification-objects"].([]interface{}); ok {
                strs := make([]interface{}, len(arr))
                for i, e := range arr { strs[i] = fmt.Sprintf("%v", e) }
                natMap["classification_objects"] = strs
            }
            d.Set("nat", []interface{}{natMap})
        }
    }
    if v, exists := _respData["errors"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            errorsMap := map[string]interface{}{}
            if arr, ok := m["NAT"].([]interface{}); ok {
                strs := make([]interface{}, len(arr))
                for i, e := range arr { strs[i] = fmt.Sprintf("%v", e) }
                errorsMap["nat"] = strs
            }
            if arr, ok := m["access"].([]interface{}); ok {
                strs := make([]interface{}, len(arr))
                for i, e := range arr { strs[i] = fmt.Sprintf("%v", e) }
                errorsMap["access"] = strs
            }
            d.Set("errors", []interface{}{errorsMap})
        }
    }


    d.SetId(fmt.Sprintf("simulate-packet-" + acctest.RandString(10)))
    return nil
}

func readGaiaSimulatePacket(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSimulatePacket(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

