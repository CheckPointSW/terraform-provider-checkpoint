package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowDynamicLayer() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowDynamicLayer,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Object name. Must be unique in the domain.`,
            },
            "virtual_system_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "meta_info": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "last_modifier": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "last_modify_time": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "iso": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "posix": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "comments": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "tags": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "custom_fields": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "field_1": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "field_2": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "field_3": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "objects": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hosts": {
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
                                    "ip_address": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
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
                                },
                            },
                        },
                        "networks": {
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
                                    "subnet": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "subnet4": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "subnet6": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "subnet_mask": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "mask_length": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "mask_length4": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "mask_length6": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "broadcast": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "network_groups": {
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
                                    "members": {
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
                        "services_tcp": {
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
                                    "keep_connections_open_after_policy_installation": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "session_timeout": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "sync_connections_on_cluster": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "port": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "source_port": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "use_delayed_sync": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "delayed_sync_value": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "services_udp": {
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
                                    "keep_connections_open_after_policy_installation": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "session_timeout": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "sync_connections_on_cluster": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "port": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "source_port": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "accept_replies": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "services_other": {
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
                                    "keep_connections_open_after_policy_installation": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "session_timeout": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "sync_connections_on_cluster": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "accept_replies": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ip_protocol": {
                                        Type:        schema.TypeInt,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "service_groups": {
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
                                    "members": {
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
                        "application_sites": {
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
                                    "clone_of": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "services": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "negate": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "application_site_categories": {
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
                                    "clone_of": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "services": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "negate": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "application_site_groups": {
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
                                    "members": {
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
                        "dynamic_objects": {
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
                                },
                            },
                        },
                        "dns_domains": {
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
                                    "is_sub_domain": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "address_ranges": {
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
                                    "ip_address_first": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv4_address_first": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_address_first": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ip_address_last": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv4_address_last": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_address_last": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "groups_with_exclusion": {
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
                                    "include": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "except": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "wildcards": {
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
                                    "ipv4_address": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv4_mask_wildcard": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_address": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "ipv6_mask_wildcard": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "identity_tags": {
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
                                    "external_identifier": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "access_roles": {
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
                                    "ip_spoofing_protection": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "networks": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "users": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "machines": {
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
                        "access_layers": {
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
                                    "firewall": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "applications_and_url_filtering": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "content_awareness": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "detect_using_x_forward_for": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "mobile_access": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "implicit_cleanup_action": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "rulebase": {
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
                        "action": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "action_settings": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "enable_identity_captive_portal": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "inline_layer": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "track": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "accounting": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "alert": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "enable_firewall_session": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "per_connection": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "per_session": {
                                        Type:        schema.TypeBool,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "type": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
                            },
                        },
                        "source": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "source_negate": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "destination": {
                            Type:        schema.TypeSet,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "destination_negate": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "user_check": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "confirm": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "custom_frequency": {
                                        Type:        schema.TypeList,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "every": {
                                                    Type:        schema.TypeInt,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                                "unit": {
                                                    Type:        schema.TypeString,
                                                    Computed:    true,
                                                    Description: `N/A`,
                                                },
                                            },
                                        },
                                    },
                                    "frequency": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                    "interaction": {
                                        Type:        schema.TypeString,
                                        Computed:    true,
                                        Description: `N/A`,
                                    },
                                },
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
                        "service_negate": {
                            Type:        schema.TypeBool,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "uid": {
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

func readGaiaShowDynamicLayer(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-dynamic-layer - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-dynamic-layer", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && commandRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !commandRes.Success {
            errMsg = commandRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = commandRes.GetData()
        }

        debugLogOperation(
            "dynamic-layer",        // resource type
            "read",                       // operation
            "show-dynamic-layer",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-dynamic-layer: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["meta-info"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("meta_info", []interface{}{map[string]interface{}{
                "last_modifier": func() string { if _v, _ok := _m["last-modifier"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "last_modify_time": func() []interface{} {
                    if _nd, _ok := _m["last-modify-time	"].(map[string]interface{}); _ok {
                        return []interface{}{map[string]interface{}{
                            "iso": func() string { if _v, _ok := _nd["iso"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                            "posix": func() int { if f, ok := _nd["posix"].(float64); ok { return int(f) }; return 0 }(),
                        }}
                    }
                    return []interface{}{}
                }(),
                "comments": func() string { if _v, _ok := _m["comments"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "tags": func() string { if _v, _ok := _m["tags"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                "custom_fields": func() []interface{} {
                    if _nd, _ok := _m["custom-fields"].(map[string]interface{}); _ok {
                        return []interface{}{map[string]interface{}{
                            "field_1": func() string { if _v, _ok := _nd["field-1"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                            "field_2": func() string { if _v, _ok := _nd["field-2"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                            "field_3": func() string { if _v, _ok := _nd["field-3"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        }}
                    }
                    return []interface{}{}
                }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["objects"]; exists {
        if _m, _ok := v.(map[string]interface{}); _ok {
            d.Set("objects", []interface{}{map[string]interface{}{
                "hosts": func() []interface{} {
                    if _arr, _ok := _m["hosts"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "ip_address": fmt.Sprintf("%v", _im["ip-address"]),
                                    "ipv4_address": fmt.Sprintf("%v", _im["ipv4-address"]),
                                    "ipv6_address": fmt.Sprintf("%v", _im["ipv6-address"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "networks": func() []interface{} {
                    if _arr, _ok := _m["networks"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "subnet": fmt.Sprintf("%v", _im["subnet"]),
                                    "subnet4": fmt.Sprintf("%v", _im["subnet4"]),
                                    "subnet6": fmt.Sprintf("%v", _im["subnet6"]),
                                    "subnet_mask": fmt.Sprintf("%v", _im["subnet-mask"]),
                                    "mask_length": func() int { if f, ok := _im["mask-length"].(float64); ok { return int(f) }; return 0 }(),
                                    "mask_length4": func() int { if f, ok := _im["mask-length4"].(float64); ok { return int(f) }; return 0 }(),
                                    "mask_length6": func() int { if f, ok := _im["mask-length6"].(float64); ok { return int(f) }; return 0 }(),
                                    "broadcast": fmt.Sprintf("%v", _im["broadcast"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "network_groups": func() []interface{} {
                    if _arr, _ok := _m["network-groups"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "members": fmt.Sprintf("%v", _im["members"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "services_tcp": func() []interface{} {
                    if _arr, _ok := _m["services-tcp"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "keep_connections_open_after_policy_installation": func() bool { if b, ok := _im["keep-connections-open-after-policy-installation"].(bool); ok { return b }; if s, ok := _im["keep-connections-open-after-policy-installation"].(string); ok { return s == "true" }; return false }(),
                                    "session_timeout": func() int { if f, ok := _im["session-timeout"].(float64); ok { return int(f) }; return 0 }(),
                                    "sync_connections_on_cluster": func() bool { if b, ok := _im["sync-connections-on-cluster"].(bool); ok { return b }; if s, ok := _im["sync-connections-on-cluster"].(string); ok { return s == "true" }; return false }(),
                                    "port": fmt.Sprintf("%v", _im["port"]),
                                    "source_port": fmt.Sprintf("%v", _im["source-port"]),
                                    "use_delayed_sync": func() bool { if b, ok := _im["use-delayed-sync"].(bool); ok { return b }; if s, ok := _im["use-delayed-sync"].(string); ok { return s == "true" }; return false }(),
                                    "delayed_sync_value": func() int { if f, ok := _im["delayed-sync-value"].(float64); ok { return int(f) }; return 0 }(),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "services_udp": func() []interface{} {
                    if _arr, _ok := _m["services-udp"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "keep_connections_open_after_policy_installation": func() bool { if b, ok := _im["keep-connections-open-after-policy-installation"].(bool); ok { return b }; if s, ok := _im["keep-connections-open-after-policy-installation"].(string); ok { return s == "true" }; return false }(),
                                    "session_timeout": func() int { if f, ok := _im["session-timeout"].(float64); ok { return int(f) }; return 0 }(),
                                    "sync_connections_on_cluster": func() bool { if b, ok := _im["sync-connections-on-cluster"].(bool); ok { return b }; if s, ok := _im["sync-connections-on-cluster"].(string); ok { return s == "true" }; return false }(),
                                    "port": fmt.Sprintf("%v", _im["port"]),
                                    "source_port": fmt.Sprintf("%v", _im["source-port"]),
                                    "accept_replies": func() bool { if b, ok := _im["accept-replies"].(bool); ok { return b }; if s, ok := _im["accept-replies"].(string); ok { return s == "true" }; return false }(),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "services_other": func() []interface{} {
                    if _arr, _ok := _m["services-other"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "keep_connections_open_after_policy_installation": func() bool { if b, ok := _im["keep-connections-open-after-policy-installation"].(bool); ok { return b }; if s, ok := _im["keep-connections-open-after-policy-installation"].(string); ok { return s == "true" }; return false }(),
                                    "session_timeout": func() int { if f, ok := _im["session-timeout"].(float64); ok { return int(f) }; return 0 }(),
                                    "sync_connections_on_cluster": func() bool { if b, ok := _im["sync-connections-on-cluster"].(bool); ok { return b }; if s, ok := _im["sync-connections-on-cluster"].(string); ok { return s == "true" }; return false }(),
                                    "accept_replies": func() bool { if b, ok := _im["accept-replies"].(bool); ok { return b }; if s, ok := _im["accept-replies"].(string); ok { return s == "true" }; return false }(),
                                    "ip_protocol": func() int { if f, ok := _im["ip-protocol"].(float64); ok { return int(f) }; return 0 }(),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "service_groups": func() []interface{} {
                    if _arr, _ok := _m["service-groups"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "members": fmt.Sprintf("%v", _im["members"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "application_sites": func() []interface{} {
                    if _arr, _ok := _m["application-sites"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "clone_of": fmt.Sprintf("%v", _im["clone-of"]),
                                    "services": fmt.Sprintf("%v", _im["services"]),
                                    "negate": func() bool { if b, ok := _im["negate"].(bool); ok { return b }; if s, ok := _im["negate"].(string); ok { return s == "true" }; return false }(),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "application_site_categories": func() []interface{} {
                    if _arr, _ok := _m["application-site-categories"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "clone_of": fmt.Sprintf("%v", _im["clone-of"]),
                                    "services": fmt.Sprintf("%v", _im["services"]),
                                    "negate": func() bool { if b, ok := _im["negate"].(bool); ok { return b }; if s, ok := _im["negate"].(string); ok { return s == "true" }; return false }(),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "application_site_groups": func() []interface{} {
                    if _arr, _ok := _m["application-site-groups"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "members": fmt.Sprintf("%v", _im["members"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "dynamic_objects": func() []interface{} {
                    if _arr, _ok := _m["dynamic-objects"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "dns_domains": func() []interface{} {
                    if _arr, _ok := _m["dns-domains"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "is_sub_domain": func() bool { if b, ok := _im["is-sub-domain"].(bool); ok { return b }; if s, ok := _im["is-sub-domain"].(string); ok { return s == "true" }; return false }(),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "address_ranges": func() []interface{} {
                    if _arr, _ok := _m["address-ranges"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "ip_address_first": fmt.Sprintf("%v", _im["ip-address-first"]),
                                    "ipv4_address_first": fmt.Sprintf("%v", _im["ipv4-address-first"]),
                                    "ipv6_address_first": fmt.Sprintf("%v", _im["ipv6-address-first"]),
                                    "ip_address_last": fmt.Sprintf("%v", _im["ip-address-last"]),
                                    "ipv4_address_last": fmt.Sprintf("%v", _im["ipv4-address-last"]),
                                    "ipv6_address_last": fmt.Sprintf("%v", _im["ipv6-address-last"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "groups_with_exclusion": func() []interface{} {
                    if _arr, _ok := _m["groups-with-exclusion"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "include": fmt.Sprintf("%v", _im["include"]),
                                    "except": fmt.Sprintf("%v", _im["except"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "wildcards": func() []interface{} {
                    if _arr, _ok := _m["wildcards"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "ipv4_address": fmt.Sprintf("%v", _im["ipv4-address"]),
                                    "ipv4_mask_wildcard": fmt.Sprintf("%v", _im["ipv4-mask-wildcard"]),
                                    "ipv6_address": fmt.Sprintf("%v", _im["ipv6-address"]),
                                    "ipv6_mask_wildcard": fmt.Sprintf("%v", _im["ipv6-mask-wildcard"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "identity_tags": func() []interface{} {
                    if _arr, _ok := _m["identity-tags"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "external_identifier": fmt.Sprintf("%v", _im["external-identifier"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "access_roles": func() []interface{} {
                    if _arr, _ok := _m["access-roles"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "ip_spoofing_protection": func() bool { if b, ok := _im["ip-spoofing-protection"].(bool); ok { return b }; if s, ok := _im["ip-spoofing-protection"].(string); ok { return s == "true" }; return false }(),
                                    "networks": fmt.Sprintf("%v", _im["networks"]),
                                    "users": fmt.Sprintf("%v", _im["users"]),
                                    "machines": fmt.Sprintf("%v", _im["machines"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
                "access_layers": func() []interface{} {
                    if _arr, _ok := _m["access-layers"].([]interface{}); _ok {
                        _out := make([]interface{}, len(_arr))
                        for _i, _item := range _arr {
                            if _im, _ok := _item.(map[string]interface{}); _ok {
                                _out[_i] = map[string]interface{}{
                                    "name": fmt.Sprintf("%v", _im["name"]),
                                    "firewall": func() bool { if b, ok := _im["firewall"].(bool); ok { return b }; if s, ok := _im["firewall"].(string); ok { return s == "true" }; return false }(),
                                    "applications_and_url_filtering": func() bool { if b, ok := _im["applications-and-url-filtering"].(bool); ok { return b }; if s, ok := _im["applications-and-url-filtering"].(string); ok { return s == "true" }; return false }(),
                                    "content_awareness": func() bool { if b, ok := _im["content-awareness"].(bool); ok { return b }; if s, ok := _im["content-awareness"].(string); ok { return s == "true" }; return false }(),
                                    "detect_using_x_forward_for": func() bool { if b, ok := _im["detect-using-x-forward-for"].(bool); ok { return b }; if s, ok := _im["detect-using-x-forward-for"].(string); ok { return s == "true" }; return false }(),
                                    "mobile_access": func() bool { if b, ok := _im["mobile-access"].(bool); ok { return b }; if s, ok := _im["mobile-access"].(string); ok { return s == "true" }; return false }(),
                                    "implicit_cleanup_action": fmt.Sprintf("%v", _im["implicit-cleanup-action"]),
                                }
                            }
                        }
                        return _out
                    }
                    return nil
                }(),
            }})
        }
    }
    if v, exists := commandRes.GetData()["rulebase"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, len(raw))
            for i, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    mapped[i] = map[string]interface{}{
                        "name": func() string { if _v, _ok := m["name"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "action": func() string { if _v, _ok := m["action"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "action_settings": func() []interface{} {
                            if _obj, _ok := m["action-settings"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "enable_identity_captive_portal": func() bool { if b, ok := _obj["enable-identity-captive-portal"].(bool); ok { return b }; if s, ok := _obj["enable-identity-captive-portal"].(string); ok { return s == "true" }; return false }(),
                                }}
                            }
                            return nil
                        }(),
                        "inline_layer": func() string { if _v, _ok := m["inline-layer"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                        "track": func() []interface{} {
                            if _obj, _ok := m["track"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "accounting": func() bool { if b, ok := _obj["accounting"].(bool); ok { return b }; if s, ok := _obj["accounting"].(string); ok { return s == "true" }; return false }(),
                                    "alert": func() string { if _v, _ok := _obj["alert"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "enable_firewall_session": func() bool { if b, ok := _obj["enable-firewall-session"].(bool); ok { return b }; if s, ok := _obj["enable-firewall-session"].(string); ok { return s == "true" }; return false }(),
                                    "per_connection": func() bool { if b, ok := _obj["per-connection"].(bool); ok { return b }; if s, ok := _obj["per-connection"].(string); ok { return s == "true" }; return false }(),
                                    "per_session": func() bool { if b, ok := _obj["per-session"].(bool); ok { return b }; if s, ok := _obj["per-session"].(string); ok { return s == "true" }; return false }(),
                                    "type": func() string { if _v, _ok := _obj["type"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "source": func() []interface{} {
                            switch _ev := m["source"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "source_negate": func() bool { if b, ok := m["source-negate"].(bool); ok { return b }; if s, ok := m["source-negate"].(string); ok { return s == "true" }; return false }(),
                        "destination": func() []interface{} {
                            switch _ev := m["destination"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "destination_negate": func() bool { if b, ok := m["destination-negate"].(bool); ok { return b }; if s, ok := m["destination-negate"].(string); ok { return s == "true" }; return false }(),
                        "user_check": func() []interface{} {
                            if _obj, _ok := m["user-check"].(map[string]interface{}); _ok {
                                return []interface{}{map[string]interface{}{
                                    "confirm": func() string { if _v, _ok := _obj["confirm"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "custom_frequency": func() []interface{} {
                                        if _d2, _ok := _obj["custom-frequency"].(map[string]interface{}); _ok {
                                            return []interface{}{map[string]interface{}{
                                                "every": func() int { if f, ok := _d2["every"].(float64); ok { return int(f) }; return 0 }(),
                                                "unit": func() string { if _v, _ok := _d2["unit"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                            }}
                                        }
                                        return nil
                                    }(),
                                    "frequency": func() string { if _v, _ok := _obj["frequency"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                    "interaction": func() string { if _v, _ok := _obj["interaction"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                                }}
                            }
                            return nil
                        }(),
                        "service": func() []interface{} {
                            switch _ev := m["service"].(type) {
                            case string:
                                return []interface{}{_ev}
                            case []interface{}:
                                return _ev
                            default:
                                return []interface{}{}
                            }
                        }(),
                        "service_negate": func() bool { if b, ok := m["service-negate"].(bool); ok { return b }; if s, ok := m["service-negate"].(string); ok { return s == "true" }; return false }(),
                        "uid": func() string { if _v, _ok := m["uid"]; _ok && _v != nil { return fmt.Sprintf("%v", _v) }; return "" }(),
                    }
                }
            }
            d.Set("rulebase", mapped)
        }
    } else {
        d.Set("rulebase", []interface{}{})
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-dynamic-layer-" + acctest.RandString(10)))
    return nil
}

