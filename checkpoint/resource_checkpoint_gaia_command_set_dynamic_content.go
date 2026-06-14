package checkpoint

import (
        "context"
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"
)
func dataGaiaSetDynamicContent() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaSetDynamicContent,
        Read:   readGaiaSetDynamicContent,
        Delete: deleteGaiaSetDynamicContent,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "comments": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Comments for this operation.`,
            },
            "tags": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: `List of tags for this operation.`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "custom_fields": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `List of custom fields for this operation.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "field_1": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `First Custom Field`,
                        },
                        "field_2": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Second Custom Field`,
                        },
                        "field_3": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Third Custom Field`,
                        },
                    },
                },
            },
            "dry_run": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `Perform validation without applying changes.`,
            },
            "referenced_objects": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `List of object names defined externally (\"internet\" , \"any\" , \"_GW_\" are already referenced).`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "application_sites": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `List of Application/Site objects as configured in SmartConsole and identified by the name.`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "application_site_categories": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `List of Application/Site Category objects as configured in SmartConsole and identified by the name.`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "services_tcp": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `List of TCP service objects as configured in SmartConsole and identified by the name.`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "services_udp": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `List of UDP service objects as configured in SmartConsole and identified by the name.`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "services_icmp": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `List of ICMP service objects as configured in SmartConsole and identified by the name.`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "updatable_objects": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `List of Updatable objects as configured in SmartConsole and identified by the name.`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "access_layers": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `List of Policy Layers in the Access Control Policy as configured in SmartConsole and identified by the name.`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },
            "objects": {
                Type:        schema.TypeList,
                Required:    true,
                ForceNew:    true,
                Description: `List of objects to create.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hosts": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "ip_address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv4 or IPv6 address. If both addresses are required then use the 'ipv4-address' and 'ipv6-address' fields explicitly.`,
                                    },
                                    "ipv4_address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv4 address.`,
                                    },
                                    "ipv6_address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv6 address.`,
                                    },
                                },
                            },
                        },
                        "networks": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "subnet": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv4 or IPv6 network address. If both addresses are required then use the 'subnet4' and 'subnet6' fields explicitly.`,
                                    },
                                    "subnet4": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv4 address.`,
                                    },
                                    "subnet6": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv6 address.`,
                                    },
                                    "subnet_mask": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv4 network mask.`,
                                    },
                                    "mask_length": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv4 or IPv6 network mask length.                     If both masks are required then use the 'mask-length4' and 'mask-length6' fields explicitly.                    Instead of the IPv4 mask length it is possible to specify the IPv4 mask itself in 'subnet-mask' field.`,
                                    },
                                    "mask_length4": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv4 network mask length.`,
                                    },
                                    "mask_length6": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv6 network mask length.`,
                                    },
                                    "broadcast": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Allow broadcast address inclusion.`,
                                    },
                                },
                            },
                        },
                        "network_groups": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "members": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Collection of Network objects identified by the name.`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "services_tcp": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "keep_connections_open_after_policy_installation": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Keep connections open after policy has been                   installed even if they are not allowed under the new policy.                       This overrides the settings on the Connection Persistence page of the Security Gateway object.                           If you change this property, the change will not affect open connections,                               but only future connections.`,
                                    },
                                    "session_timeout": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Time (in seconds) before the session times out.`,
                                    },
                                    "sync_connections_on_cluster": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Enables the state synchronization in a ClusterXL or OPSEC-certified cluster.`,
                                    },
                                    "port": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `The number of the port used to provide this service.                    To specify a port range, place a hyphen between the lowest and highest port numbers,                        for example 44-55.`,
                                    },
                                    "source_port": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Port number for the client-side service.                     If specified, only packets with these source port numbers will be accepted,                     dropped, or rejected during packet inspection.                     Otherwise, the packets are not matched to this service.`,
                                    },
                                    "use_delayed_sync": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Enable this option to delay notifying the Security Gateway about a connection,                    so that the connection will only be synchronized if it still exists x seconds after the connection is initiated.                     This feature uses SecureXL that is enabled by default.`,
                                    },
                                    "delayed_sync_value": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specify the delay (in seconds) when the synchronization will start after connection initiation. Relevant only if \"use-delayed-sync\" was set to \"true\".`,
                                    },
                                },
                            },
                        },
                        "services_udp": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "keep_connections_open_after_policy_installation": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Keep connections open after policy has been                   installed even if they are not allowed under the new policy.                       This overrides the settings on the Connection Persistence page of the Security Gateway object.                           If you change this property, the change will not affect open connections,                               but only future connections.`,
                                    },
                                    "session_timeout": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Time (in seconds) before the session times out.`,
                                    },
                                    "sync_connections_on_cluster": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Enables onthe state synchronization in a ClusterXL or OPSEC-certified cluster.`,
                                    },
                                    "port": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `The number of the port used to provide this service.         To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.`,
                                    },
                                    "source_port": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Port number for the client-side service.                     If specified, only packets with these source port numbers will be accepted,                     dropped, or rejected during packet inspection.                     Otherwise, the packets are not matched to this service.`,
                                    },
                                    "accept_replies": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specifies whether to accept UDP replies for this service.`,
                                    },
                                },
                            },
                        },
                        "services_other": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "keep_connections_open_after_policy_installation": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Keep connections open after policy has been installed even if they are not allowed under the new policy.                     This overrides the settings on the Connection Persistence page in the Security Gateway object.                     If you change this property, the change will not affect open connections, but only future connections.`,
                                    },
                                    "session_timeout": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Time (in seconds) before the session times out.`,
                                    },
                                    "sync_connections_on_cluster": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Enables the state synchronization in a ClusterXL or OPSEC-certified cluster.`,
                                    },
                                    "accept_replies": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specifies whether to accept replies for this service.`,
                                    },
                                    "ip_protocol": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IP protocol number.`,
                                    },
                                },
                            },
                        },
                        "service_groups": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "members": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Collection of Service objects identified by the name.`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "application_sites": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "clone_of": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Name of existing Application/Category to be cloned.`,
                                    },
                                    "services": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Collection of Service objects identified by the name. You can specify any service with the value 'any' or 'Any'.`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "negate": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specifies whether this object is negated.`,
                                    },
                                },
                            },
                        },
                        "application_site_categories": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "clone_of": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Name of existing Application/Category to be cloned.`,
                                    },
                                    "services": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Collection of Service objects identified by the name. You can specify any service with the value 'any' or 'Any'.`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "negate": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specifies whether this object is negated.`,
                                    },
                                },
                            },
                        },
                        "application_site_groups": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "members": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Collection of Application/Site objects identified by the name.`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                },
                            },
                        },
                        "dynamic_objects": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                },
                            },
                        },
                        "dns_domains": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "is_sub_domain": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specifies whether to match sub-domains in addition to the domain itself.`,
                                    },
                                },
                            },
                        },
                        "address_ranges": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "ip_address_first": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `First IP address in the range.                     If both IPv4 and IPv6 address ranges are required,                         then use the 'ipv4-address-first' and the 'ipv6-address-first' fields instead.`,
                                    },
                                    "ipv4_address_first": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `First IPv4 address in the range.`,
                                    },
                                    "ipv6_address_first": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `First IPv6 address in the range.`,
                                    },
                                    "ip_address_last": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Last IP address in the range.                     If both IPv4 and IPv6 address ranges are required,                         then use the 'ipv4-address-last' and the 'ipv6-address-last' fields instead.`,
                                    },
                                    "ipv4_address_last": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Last IPv4 address in the range.`,
                                    },
                                    "ipv6_address_last": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Last IPv6 address in the range.`,
                                    },
                                },
                            },
                        },
                        "groups_with_exclusion": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "include": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Name of an object which the group includes.`,
                                    },
                                    "except": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Name of an object which the group excludes.`,
                                    },
                                },
                            },
                        },
                        "wildcards": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "ipv4_address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv4 address.`,
                                    },
                                    "ipv4_mask_wildcard": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv4 mask wildcard.`,
                                    },
                                    "ipv6_address": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv6 address.`,
                                    },
                                    "ipv6_mask_wildcard": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `IPv6 mask wildcard.`,
                                    },
                                },
                            },
                        },
                        "identity_tags": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "external_identifier": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `External identifier. For example: Cisco ISE security group tag.`,
                                    },
                                },
                            },
                        },
                        "access_roles": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "ip_spoofing_protection": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Enforce IP spoofing protection.`,
                                    },
                                    "networks": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Collection of Network objects identified by the name that can access the system.                         Level of details in the output corresponds to the number of details for search.                             You can specify any Access Role with the value 'any' or 'Any'.`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "users": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Users that can access the system.                     Level of details in the output corresponds to the number of details for search.                     Valid options: 'any', 'all identified'.`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "source": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Active Directory name or 'Identity Tag' or 'Guests'.`,
                                                },
                                                "selection": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Distinguished Name (DN) or Identity Tag name or 'Unauthenticated Guests'.`,
                                                },
                                                "ad_entity_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Active directory entity type.`,
                                                },
                                            },
                                        },
                                    },
                                    "machines": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Machines that can access the system.                         Level of details in the output corresponds to the number of details for search.                         Valid options: 'any', 'all identified'.`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "source": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Active Directory name or 'Identity Tag'.`,
                                                },
                                                "selection": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Distinguished Name (DN) or Identity Tag name.`,
                                                },
                                                "ad_entity_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Active directory entity type.`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "access_layers": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Object name. Must be unique in the domain.`,
                                    },
                                    "firewall": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Whether to enable the Firewall blade on the layer.`,
                                    },
                                    "applications_and_url_filtering": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Whether to enable the Application Control & URL Filtering blades on the layer.`,
                                    },
                                    "content_awareness": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Whether to enable the Content Awareness blade on the layer.`,
                                    },
                                    "detect_using_x_forward_for": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Whether to use the 'X-Forward-For' HTTP header,                     which is added by the proxy server to keep track of the original source IP address.`,
                                    },
                                    "mobile_access": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Whether to enable the Mobile Access blade on the layer.`,
                                    },
                                    "implicit_cleanup_action": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specifies the default 'catch-all' action for traffic that does not match any explicit or implied rules in the layer.`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "access_layers_content": {
                Type:        schema.TypeList,
                Required:    true,
                ForceNew:    true,
                Description: `List of layers to apply. Supported layers : Layers created using this API, externally referenced layers (layers marked as 'dynamic layers' in the SmartConsole)`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Layer name.`,
                        },
                        "operation": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Layer operation.`,
                        },
                        "rulebase": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Rules of the layer.`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Rule name, Must be unique in the layer.`,
                                    },
                                    "action": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Action. Valid options: \"Accept\", \"Drop\", \"Ask\", \"Drop with Block message\", \"Inform\", \"Reject\", \"Apply Layer\".`,
                                    },
                                    "action_settings": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Action settings.`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "enable_identity_captive_portal": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Redirect HTTP traffic to an authentication (Captive Portal). After the user is authenticated, new connections from this source are inspected without requiring authentication.`,
                                                },
                                            },
                                        },
                                    },
                                    "inline_layer": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Inline Layer identified by the name.  Relevant only if \"the action\" was set to \"Apply Layer\".`,
                                    },
                                    "track": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Track Settings.`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "accounting": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Turns the Accounting on and off.`,
                                                },
                                                "alert": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Type of alert for the track. Valid options: \"None\", \"Alert\", \"Snmp\", \"Mail\", \"User Alert 1\", \"User Alert 2\", \"User Alert 3\".`,
                                                },
                                                "enable_firewall_session": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Specifies whether to generate a session log for connections that are inspected only by the Firewall blade.`,
                                                },
                                                "per_connection": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Specifies whether to generate a log for each connection. If set to 'true', may decrease the Security Gateway performance because of the number of generated logs.`,
                                                },
                                                "per_session": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Specifies whether to generate a log for each session. If set to 'true', may decrease the Security Gateway performance because of the number of generated logs.`,
                                                },
                                                "type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Track type. Valid options: \"Log\", \"Extended Log\", \"Detailed Log\", \"None\".`,
                                                },
                                            },
                                        },
                                    },
                                    "source": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Collection of network objects identified by their name or 'any'.`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "source_negate": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specifies whether to negate the source.`,
                                    },
                                    "destination": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Collection of network objects identified by their name or 'any'.`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "destination_negate": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specifies whether to negate the destination.`,
                                    },
                                    "user_check": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `UserCheck settings.`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "confirm": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Valid options: \"per rule\", \"per category\", \"per application/site\", \"per data type\" .`,
                                                },
                                                "custom_frequency": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Configure how often the user sees the configured message when the action is \"ask\", \"inform\", or \"block\". Relevant only if \"frequency\" was set to \"custom frequency\".`,
                                                    MaxItems:    1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "every": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                ForceNew:    true,
                                                                Description: `Valid values: 1 - 999 .`,
                                                            },
                                                            "unit": {
                                                                Type:        schema.TypeString,
                                                                Optional:    true,
                                                                ForceNew:    true,
                                                                Description: `Valid options: hours, days, weeks, months .`,
                                                            },
                                                        },
                                                    },
                                                },
                                                "frequency": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Configure how often the user sees the configured message when the action is \"ask\", \"inform\", or \"block\". Valid options: \"once a day\", \"once a week\", \"once a month\", \"custom frequency\" .`,
                                                },
                                                "interaction": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    ForceNew:    true,
                                                    Description: `Add the relevant interaction text. Need to be relevant to the rule action.`,
                                                },
                                            },
                                        },
                                    },
                                    "service": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Collection of Service and Application objects identified by the name. You can specify any object with the value 'any' or 'Any'.`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "service_negate": {
                                        Type:        schema.TypeBool,
                                        Optional:    true,
                                        ForceNew:    true,
                                        Description: `Specifies whether to negate this service.`,
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "validation_warnings": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "layer": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rule": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "object": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "message": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "validation_errors": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "layer": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "rule": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "object": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "message": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "error_code": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "change_summary": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
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
                                    "rules": {
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
                        "objects": {
                            Type:        schema.TypeList,
                            Computed:    true,
                            Description: `N/A`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "create": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "modify": {
                                        Type:        schema.TypeSet,
                                        Computed:    true,
                                        Description: `N/A`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "delete": {
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
        },
    }
}

func createGaiaSetDynamicContent(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    if v := d.Get("tags"); len(v.(*schema.Set).List()) > 0 {
        payload["tags"] = v.(*schema.Set).List()
    }

    if v := d.Get("custom_fields"); len(v.([]interface{})) > 0 {
        _ = v
        customfieldsMap := make(map[string]interface{})
        if v, ok := d.GetOk("custom_fields.0.field_1"); ok {
            customfieldsMap["field-1"] = v.(string)
        }
        if v, ok := d.GetOk("custom_fields.0.field_2"); ok {
            customfieldsMap["field-2"] = v.(string)
        }
        if v, ok := d.GetOk("custom_fields.0.field_3"); ok {
            customfieldsMap["field-3"] = v.(string)
        }
        if len(customfieldsMap) > 0 {
            payload["custom-fields"] = customfieldsMap
        }
    }

    if v, ok := d.GetOkExists("dry_run"); ok {
        payload["dry-run"] = v.(bool)
    }

    if v := d.Get("referenced_objects"); len(v.([]interface{})) > 0 {
        _ = v
        referencedobjectsMap := make(map[string]interface{})
        if v := d.Get("referenced_objects.0.application_sites"); len(v.(*schema.Set).List()) > 0 {
            referencedobjectsMap["application-sites"] = v.(*schema.Set).List()
        }
        if v := d.Get("referenced_objects.0.application_site_categories"); len(v.(*schema.Set).List()) > 0 {
            referencedobjectsMap["application-site-categories"] = v.(*schema.Set).List()
        }
        if v := d.Get("referenced_objects.0.services_tcp"); len(v.(*schema.Set).List()) > 0 {
            referencedobjectsMap["services-tcp"] = v.(*schema.Set).List()
        }
        if v := d.Get("referenced_objects.0.services_udp"); len(v.(*schema.Set).List()) > 0 {
            referencedobjectsMap["services-udp"] = v.(*schema.Set).List()
        }
        if v := d.Get("referenced_objects.0.services_icmp"); len(v.(*schema.Set).List()) > 0 {
            referencedobjectsMap["services-icmp"] = v.(*schema.Set).List()
        }
        if v := d.Get("referenced_objects.0.updatable_objects"); len(v.(*schema.Set).List()) > 0 {
            referencedobjectsMap["updatable-objects"] = v.(*schema.Set).List()
        }
        if v := d.Get("referenced_objects.0.access_layers"); len(v.(*schema.Set).List()) > 0 {
            referencedobjectsMap["access-layers"] = v.(*schema.Set).List()
        }
        if len(referencedobjectsMap) > 0 {
            payload["referenced-objects"] = referencedobjectsMap
        }
    }

    if v := d.Get("objects"); len(v.([]interface{})) > 0 {
        _ = v
        objectsMap := make(map[string]interface{})
        if v, ok := d.GetOk("objects.0.hosts"); ok {
            hostsList := v.([]interface{})
            hostsArray := make([]interface{}, 0, len(hostsList))
            for i := range hostsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.hosts.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.hosts.%d.ip_address", i)); ok {
                    itemMap["ip-address"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.hosts.%d.ipv4_address", i)); ok {
                    itemMap["ipv4-address"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.hosts.%d.ipv6_address", i)); ok {
                    itemMap["ipv6-address"] = v.(string)
                }
                if len(itemMap) > 0 {
                    hostsArray = append(hostsArray, itemMap)
                }
            }
            if len(hostsArray) > 0 {
                objectsMap["hosts"] = hostsArray
            }
        }
        if v, ok := d.GetOk("objects.0.networks"); ok {
            networksList := v.([]interface{})
            networksArray := make([]interface{}, 0, len(networksList))
            for i := range networksList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.networks.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.networks.%d.subnet", i)); ok {
                    itemMap["subnet"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.networks.%d.subnet4", i)); ok {
                    itemMap["subnet4"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.networks.%d.subnet6", i)); ok {
                    itemMap["subnet6"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.networks.%d.subnet_mask", i)); ok {
                    itemMap["subnet-mask"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.networks.%d.mask_length", i)); ok {
                    itemMap["mask-length"] = v.(int)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.networks.%d.mask_length4", i)); ok {
                    itemMap["mask-length4"] = v.(int)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.networks.%d.mask_length6", i)); ok {
                    itemMap["mask-length6"] = v.(int)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.networks.%d.broadcast", i)); ok {
                    itemMap["broadcast"] = v.(string)
                }
                if len(itemMap) > 0 {
                    networksArray = append(networksArray, itemMap)
                }
            }
            if len(networksArray) > 0 {
                objectsMap["networks"] = networksArray
            }
        }
        if v, ok := d.GetOk("objects.0.network_groups"); ok {
            networkgroupsList := v.([]interface{})
            networkgroupsArray := make([]interface{}, 0, len(networkgroupsList))
            for i := range networkgroupsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.network_groups.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if sv := d.Get(fmt.Sprintf("objects.0.network_groups.%d.members", i)); sv != nil {
                    if set, ok := sv.(*schema.Set); ok && set.Len() > 0 {
                        itemMap["members"] = set.List()
                    }
                }
                if len(itemMap) > 0 {
                    networkgroupsArray = append(networkgroupsArray, itemMap)
                }
            }
            if len(networkgroupsArray) > 0 {
                objectsMap["network-groups"] = networkgroupsArray
            }
        }
        if v, ok := d.GetOk("objects.0.services_tcp"); ok {
            servicestcpList := v.([]interface{})
            servicestcpArray := make([]interface{}, 0, len(servicestcpList))
            for i := range servicestcpList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_tcp.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v := d.Get(fmt.Sprintf("objects.0.services_tcp.%d.keep_connections_open_after_policy_installation", i)).(bool); v {
                    itemMap["keep-connections-open-after-policy-installation"] = v
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_tcp.%d.session_timeout", i)); ok {
                    itemMap["session-timeout"] = v.(int)
                }
                if v := d.Get(fmt.Sprintf("objects.0.services_tcp.%d.sync_connections_on_cluster", i)).(bool); v {
                    itemMap["sync-connections-on-cluster"] = v
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_tcp.%d.port", i)); ok {
                    itemMap["port"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_tcp.%d.source_port", i)); ok {
                    itemMap["source-port"] = v.(string)
                }
                if v := d.Get(fmt.Sprintf("objects.0.services_tcp.%d.use_delayed_sync", i)).(bool); v {
                    itemMap["use-delayed-sync"] = v
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_tcp.%d.delayed_sync_value", i)); ok {
                    itemMap["delayed-sync-value"] = v.(int)
                }
                if len(itemMap) > 0 {
                    servicestcpArray = append(servicestcpArray, itemMap)
                }
            }
            if len(servicestcpArray) > 0 {
                objectsMap["services-tcp"] = servicestcpArray
            }
        }
        if v, ok := d.GetOk("objects.0.services_udp"); ok {
            servicesudpList := v.([]interface{})
            servicesudpArray := make([]interface{}, 0, len(servicesudpList))
            for i := range servicesudpList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_udp.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v := d.Get(fmt.Sprintf("objects.0.services_udp.%d.keep_connections_open_after_policy_installation", i)).(bool); v {
                    itemMap["keep-connections-open-after-policy-installation"] = v
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_udp.%d.session_timeout", i)); ok {
                    itemMap["session-timeout"] = v.(int)
                }
                if v := d.Get(fmt.Sprintf("objects.0.services_udp.%d.sync_connections_on_cluster", i)).(bool); v {
                    itemMap["sync-connections-on-cluster"] = v
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_udp.%d.port", i)); ok {
                    itemMap["port"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_udp.%d.source_port", i)); ok {
                    itemMap["source-port"] = v.(string)
                }
                if v := d.Get(fmt.Sprintf("objects.0.services_udp.%d.accept_replies", i)).(bool); v {
                    itemMap["accept-replies"] = v
                }
                if len(itemMap) > 0 {
                    servicesudpArray = append(servicesudpArray, itemMap)
                }
            }
            if len(servicesudpArray) > 0 {
                objectsMap["services-udp"] = servicesudpArray
            }
        }
        if v, ok := d.GetOk("objects.0.services_other"); ok {
            servicesotherList := v.([]interface{})
            servicesotherArray := make([]interface{}, 0, len(servicesotherList))
            for i := range servicesotherList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_other.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v := d.Get(fmt.Sprintf("objects.0.services_other.%d.keep_connections_open_after_policy_installation", i)).(bool); v {
                    itemMap["keep-connections-open-after-policy-installation"] = v
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_other.%d.session_timeout", i)); ok {
                    itemMap["session-timeout"] = v.(int)
                }
                if v := d.Get(fmt.Sprintf("objects.0.services_other.%d.sync_connections_on_cluster", i)).(bool); v {
                    itemMap["sync-connections-on-cluster"] = v
                }
                if v := d.Get(fmt.Sprintf("objects.0.services_other.%d.accept_replies", i)).(bool); v {
                    itemMap["accept-replies"] = v
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.services_other.%d.ip_protocol", i)); ok {
                    itemMap["ip-protocol"] = v.(int)
                }
                if len(itemMap) > 0 {
                    servicesotherArray = append(servicesotherArray, itemMap)
                }
            }
            if len(servicesotherArray) > 0 {
                objectsMap["services-other"] = servicesotherArray
            }
        }
        if v, ok := d.GetOk("objects.0.service_groups"); ok {
            servicegroupsList := v.([]interface{})
            servicegroupsArray := make([]interface{}, 0, len(servicegroupsList))
            for i := range servicegroupsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.service_groups.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if sv := d.Get(fmt.Sprintf("objects.0.service_groups.%d.members", i)); sv != nil {
                    if set, ok := sv.(*schema.Set); ok && set.Len() > 0 {
                        itemMap["members"] = set.List()
                    }
                }
                if len(itemMap) > 0 {
                    servicegroupsArray = append(servicegroupsArray, itemMap)
                }
            }
            if len(servicegroupsArray) > 0 {
                objectsMap["service-groups"] = servicegroupsArray
            }
        }
        if v, ok := d.GetOk("objects.0.application_sites"); ok {
            applicationsitesList := v.([]interface{})
            applicationsitesArray := make([]interface{}, 0, len(applicationsitesList))
            for i := range applicationsitesList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.application_sites.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.application_sites.%d.clone_of", i)); ok {
                    itemMap["clone-of"] = v.(string)
                }
                if sv := d.Get(fmt.Sprintf("objects.0.application_sites.%d.services", i)); sv != nil {
                    if set, ok := sv.(*schema.Set); ok && set.Len() > 0 {
                        itemMap["services"] = set.List()
                    }
                }
                if v := d.Get(fmt.Sprintf("objects.0.application_sites.%d.negate", i)).(bool); v {
                    itemMap["negate"] = v
                }
                if len(itemMap) > 0 {
                    applicationsitesArray = append(applicationsitesArray, itemMap)
                }
            }
            if len(applicationsitesArray) > 0 {
                objectsMap["application-sites"] = applicationsitesArray
            }
        }
        if v, ok := d.GetOk("objects.0.application_site_categories"); ok {
            applicationsitecategoriesList := v.([]interface{})
            applicationsitecategoriesArray := make([]interface{}, 0, len(applicationsitecategoriesList))
            for i := range applicationsitecategoriesList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.application_site_categories.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.application_site_categories.%d.clone_of", i)); ok {
                    itemMap["clone-of"] = v.(string)
                }
                if sv := d.Get(fmt.Sprintf("objects.0.application_site_categories.%d.services", i)); sv != nil {
                    if set, ok := sv.(*schema.Set); ok && set.Len() > 0 {
                        itemMap["services"] = set.List()
                    }
                }
                if v := d.Get(fmt.Sprintf("objects.0.application_site_categories.%d.negate", i)).(bool); v {
                    itemMap["negate"] = v
                }
                if len(itemMap) > 0 {
                    applicationsitecategoriesArray = append(applicationsitecategoriesArray, itemMap)
                }
            }
            if len(applicationsitecategoriesArray) > 0 {
                objectsMap["application-site-categories"] = applicationsitecategoriesArray
            }
        }
        if v, ok := d.GetOk("objects.0.application_site_groups"); ok {
            applicationsitegroupsList := v.([]interface{})
            applicationsitegroupsArray := make([]interface{}, 0, len(applicationsitegroupsList))
            for i := range applicationsitegroupsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.application_site_groups.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if sv := d.Get(fmt.Sprintf("objects.0.application_site_groups.%d.members", i)); sv != nil {
                    if set, ok := sv.(*schema.Set); ok && set.Len() > 0 {
                        itemMap["members"] = set.List()
                    }
                }
                if len(itemMap) > 0 {
                    applicationsitegroupsArray = append(applicationsitegroupsArray, itemMap)
                }
            }
            if len(applicationsitegroupsArray) > 0 {
                objectsMap["application-site-groups"] = applicationsitegroupsArray
            }
        }
        if v, ok := d.GetOk("objects.0.dynamic_objects"); ok {
            dynamicobjectsList := v.([]interface{})
            dynamicobjectsArray := make([]interface{}, 0, len(dynamicobjectsList))
            for i := range dynamicobjectsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.dynamic_objects.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if len(itemMap) > 0 {
                    dynamicobjectsArray = append(dynamicobjectsArray, itemMap)
                }
            }
            if len(dynamicobjectsArray) > 0 {
                objectsMap["dynamic-objects"] = dynamicobjectsArray
            }
        }
        if v, ok := d.GetOk("objects.0.dns_domains"); ok {
            dnsdomainsList := v.([]interface{})
            dnsdomainsArray := make([]interface{}, 0, len(dnsdomainsList))
            for i := range dnsdomainsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.dns_domains.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v := d.Get(fmt.Sprintf("objects.0.dns_domains.%d.is_sub_domain", i)).(bool); v {
                    itemMap["is-sub-domain"] = v
                }
                if len(itemMap) > 0 {
                    dnsdomainsArray = append(dnsdomainsArray, itemMap)
                }
            }
            if len(dnsdomainsArray) > 0 {
                objectsMap["dns-domains"] = dnsdomainsArray
            }
        }
        if v, ok := d.GetOk("objects.0.address_ranges"); ok {
            addressrangesList := v.([]interface{})
            addressrangesArray := make([]interface{}, 0, len(addressrangesList))
            for i := range addressrangesList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.address_ranges.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.address_ranges.%d.ip_address_first", i)); ok {
                    itemMap["ip-address-first"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.address_ranges.%d.ipv4_address_first", i)); ok {
                    itemMap["ipv4-address-first"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.address_ranges.%d.ipv6_address_first", i)); ok {
                    itemMap["ipv6-address-first"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.address_ranges.%d.ip_address_last", i)); ok {
                    itemMap["ip-address-last"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.address_ranges.%d.ipv4_address_last", i)); ok {
                    itemMap["ipv4-address-last"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.address_ranges.%d.ipv6_address_last", i)); ok {
                    itemMap["ipv6-address-last"] = v.(string)
                }
                if len(itemMap) > 0 {
                    addressrangesArray = append(addressrangesArray, itemMap)
                }
            }
            if len(addressrangesArray) > 0 {
                objectsMap["address-ranges"] = addressrangesArray
            }
        }
        if v, ok := d.GetOk("objects.0.groups_with_exclusion"); ok {
            groupswithexclusionList := v.([]interface{})
            groupswithexclusionArray := make([]interface{}, 0, len(groupswithexclusionList))
            for i := range groupswithexclusionList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.groups_with_exclusion.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.groups_with_exclusion.%d.include", i)); ok {
                    itemMap["include"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.groups_with_exclusion.%d.except", i)); ok {
                    itemMap["except"] = v.(string)
                }
                if len(itemMap) > 0 {
                    groupswithexclusionArray = append(groupswithexclusionArray, itemMap)
                }
            }
            if len(groupswithexclusionArray) > 0 {
                objectsMap["groups-with-exclusion"] = groupswithexclusionArray
            }
        }
        if v, ok := d.GetOk("objects.0.wildcards"); ok {
            wildcardsList := v.([]interface{})
            wildcardsArray := make([]interface{}, 0, len(wildcardsList))
            for i := range wildcardsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.wildcards.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.wildcards.%d.ipv4_address", i)); ok {
                    itemMap["ipv4-address"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.wildcards.%d.ipv4_mask_wildcard", i)); ok {
                    itemMap["ipv4-mask-wildcard"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.wildcards.%d.ipv6_address", i)); ok {
                    itemMap["ipv6-address"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.wildcards.%d.ipv6_mask_wildcard", i)); ok {
                    itemMap["ipv6-mask-wildcard"] = v.(string)
                }
                if len(itemMap) > 0 {
                    wildcardsArray = append(wildcardsArray, itemMap)
                }
            }
            if len(wildcardsArray) > 0 {
                objectsMap["wildcards"] = wildcardsArray
            }
        }
        if v, ok := d.GetOk("objects.0.identity_tags"); ok {
            identitytagsList := v.([]interface{})
            identitytagsArray := make([]interface{}, 0, len(identitytagsList))
            for i := range identitytagsList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.identity_tags.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.identity_tags.%d.external_identifier", i)); ok {
                    itemMap["external-identifier"] = v.(string)
                }
                if len(itemMap) > 0 {
                    identitytagsArray = append(identitytagsArray, itemMap)
                }
            }
            if len(identitytagsArray) > 0 {
                objectsMap["identity-tags"] = identitytagsArray
            }
        }
        if v, ok := d.GetOk("objects.0.access_roles"); ok {
            accessrolesList := v.([]interface{})
            accessrolesArray := make([]interface{}, 0, len(accessrolesList))
            for i := range accessrolesList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.access_roles.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v := d.Get(fmt.Sprintf("objects.0.access_roles.%d.ip_spoofing_protection", i)).(bool); v {
                    itemMap["ip-spoofing-protection"] = v
                }
                if sv := d.Get(fmt.Sprintf("objects.0.access_roles.%d.networks", i)); sv != nil {
                    if set, ok := sv.(*schema.Set); ok && set.Len() > 0 {
                        itemMap["networks"] = set.List()
                    }
                }
                if sv := d.Get(fmt.Sprintf("objects.0.access_roles.%d.users", i)); len(sv.([]interface{})) > 0 {
                    usersList := sv.([]interface{})
                    usersArr := make([]interface{}, 0, len(usersList))
                    for j := range usersList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("objects.0.access_roles.%d.users.%d.source", i, j)); ok {
                            innerMap["source"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("objects.0.access_roles.%d.users.%d.selection", i, j)); ok {
                            innerMap["selection"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("objects.0.access_roles.%d.users.%d.ad_entity_type", i, j)); ok {
                            innerMap["ad-entity-type"] = iv.(string)
                        }
                        if len(innerMap) > 0 {
                            usersArr = append(usersArr, innerMap)
                        }
                    }
                    if len(usersArr) > 0 {
                        itemMap["users"] = usersArr
                    }
                }
                if sv := d.Get(fmt.Sprintf("objects.0.access_roles.%d.machines", i)); len(sv.([]interface{})) > 0 {
                    machinesList := sv.([]interface{})
                    machinesArr := make([]interface{}, 0, len(machinesList))
                    for j := range machinesList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("objects.0.access_roles.%d.machines.%d.source", i, j)); ok {
                            innerMap["source"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("objects.0.access_roles.%d.machines.%d.selection", i, j)); ok {
                            innerMap["selection"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("objects.0.access_roles.%d.machines.%d.ad_entity_type", i, j)); ok {
                            innerMap["ad-entity-type"] = iv.(string)
                        }
                        if len(innerMap) > 0 {
                            machinesArr = append(machinesArr, innerMap)
                        }
                    }
                    if len(machinesArr) > 0 {
                        itemMap["machines"] = machinesArr
                    }
                }
                if len(itemMap) > 0 {
                    accessrolesArray = append(accessrolesArray, itemMap)
                }
            }
            if len(accessrolesArray) > 0 {
                objectsMap["access-roles"] = accessrolesArray
            }
        }
        if v, ok := d.GetOk("objects.0.access_layers"); ok {
            accesslayersList := v.([]interface{})
            accesslayersArray := make([]interface{}, 0, len(accesslayersList))
            for i := range accesslayersList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.access_layers.%d.name", i)); ok {
                    itemMap["name"] = v.(string)
                }
                if v := d.Get(fmt.Sprintf("objects.0.access_layers.%d.firewall", i)).(bool); v {
                    itemMap["firewall"] = v
                }
                if v := d.Get(fmt.Sprintf("objects.0.access_layers.%d.applications_and_url_filtering", i)).(bool); v {
                    itemMap["applications-and-url-filtering"] = v
                }
                if v := d.Get(fmt.Sprintf("objects.0.access_layers.%d.content_awareness", i)).(bool); v {
                    itemMap["content-awareness"] = v
                }
                if v := d.Get(fmt.Sprintf("objects.0.access_layers.%d.detect_using_x_forward_for", i)).(bool); v {
                    itemMap["detect-using-x-forward-for"] = v
                }
                if v := d.Get(fmt.Sprintf("objects.0.access_layers.%d.mobile_access", i)).(bool); v {
                    itemMap["mobile-access"] = v
                }
                if v, ok := d.GetOk(fmt.Sprintf("objects.0.access_layers.%d.implicit_cleanup_action", i)); ok {
                    itemMap["implicit-cleanup-action"] = v.(string)
                }
                if len(itemMap) > 0 {
                    accesslayersArray = append(accesslayersArray, itemMap)
                }
            }
            if len(accesslayersArray) > 0 {
                objectsMap["access-layers"] = accesslayersArray
            }
        }
        if len(objectsMap) > 0 {
            payload["objects"] = objectsMap
        }
    }

    if v := d.Get("access_layers_content"); len(v.([]interface{})) > 0 {
        accesslayerscontentList := v.([]interface{})
        accesslayerscontentArray := make([]interface{}, 0, len(accesslayerscontentList))
        for i := range accesslayerscontentList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.operation", i)); ok {
                itemMap["operation"] = v.(string)
            }
            if sv := d.Get(fmt.Sprintf("access_layers_content.%d.rulebase", i)); len(sv.([]interface{})) > 0 {
                rulebaseList := sv.([]interface{})
                rulebaseArr := make([]interface{}, 0, len(rulebaseList))
                for j := range rulebaseList {
                    innerMap := make(map[string]interface{})
                    if iv, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.rulebase.%d.name", i, j)); ok {
                        innerMap["name"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.rulebase.%d.action", i, j)); ok {
                        innerMap["action"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.rulebase.%d.action_settings", i, j)); ok {
                        if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                            rawDict := ivList[0].(map[string]interface{})
                            action_settingsMap := make(map[string]interface{})
                            if sv, ok := rawDict["enable_identity_captive_portal"]; ok && sv.(bool) {
                                action_settingsMap["enable-identity-captive-portal"] = sv.(bool)
                            }
                            if len(action_settingsMap) > 0 {
                                innerMap["action-settings"] = action_settingsMap
                            }
                        }
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.rulebase.%d.inline_layer", i, j)); ok {
                        innerMap["inline-layer"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.rulebase.%d.track", i, j)); ok {
                        if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                            rawDict := ivList[0].(map[string]interface{})
                            trackMap := make(map[string]interface{})
                            if sv, ok := rawDict["accounting"]; ok && sv.(bool) {
                                trackMap["accounting"] = sv.(bool)
                            }
                            if sv, ok := rawDict["alert"]; ok && sv.(string) != "" {
                                trackMap["alert"] = sv.(string)
                            }
                            if sv, ok := rawDict["enable_firewall_session"]; ok && sv.(bool) {
                                trackMap["enable-firewall-session"] = sv.(bool)
                            }
                            if sv, ok := rawDict["per_connection"]; ok && sv.(bool) {
                                trackMap["per-connection"] = sv.(bool)
                            }
                            if sv, ok := rawDict["per_session"]; ok && sv.(bool) {
                                trackMap["per-session"] = sv.(bool)
                            }
                            if sv, ok := rawDict["type"]; ok && sv.(string) != "" {
                                trackMap["type"] = sv.(string)
                            }
                            if len(trackMap) > 0 {
                                innerMap["track"] = trackMap
                            }
                        }
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.rulebase.%d.source", i, j)); ok {
                        ivList := iv.(*schema.Set).List()
                        if len(ivList) == 1 && strings.EqualFold(fmt.Sprintf("%v", ivList[0]), "any") {
                            innerMap["source"] = ivList[0]
                        } else {
                            innerMap["source"] = ivList
                        }
                    }
                    if v := d.Get(fmt.Sprintf("access_layers_content.%d.rulebase.%d.source_negate", i, j)).(bool); v {
                        innerMap["source-negate"] = v
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.rulebase.%d.destination", i, j)); ok {
                        ivList := iv.(*schema.Set).List()
                        if len(ivList) == 1 && strings.EqualFold(fmt.Sprintf("%v", ivList[0]), "any") {
                            innerMap["destination"] = ivList[0]
                        } else {
                            innerMap["destination"] = ivList
                        }
                    }
                    if v := d.Get(fmt.Sprintf("access_layers_content.%d.rulebase.%d.destination_negate", i, j)).(bool); v {
                        innerMap["destination-negate"] = v
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.rulebase.%d.user_check", i, j)); ok {
                        if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                            rawDict := ivList[0].(map[string]interface{})
                            user_checkMap := make(map[string]interface{})
                            if sv, ok := rawDict["confirm"]; ok && sv.(string) != "" {
                                user_checkMap["confirm"] = sv.(string)
                            }
                            if sv, ok := rawDict["custom_frequency"]; ok {
                                user_checkMap["custom-frequency"] = sv
                            }
                            if sv, ok := rawDict["frequency"]; ok && sv.(string) != "" {
                                user_checkMap["frequency"] = sv.(string)
                            }
                            if sv, ok := rawDict["interaction"]; ok && sv.(string) != "" {
                                user_checkMap["interaction"] = sv.(string)
                            }
                            if len(user_checkMap) > 0 {
                                innerMap["user-check"] = user_checkMap
                            }
                        }
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("access_layers_content.%d.rulebase.%d.service", i, j)); ok {
                        ivList := iv.(*schema.Set).List()
                        if len(ivList) == 1 && strings.EqualFold(fmt.Sprintf("%v", ivList[0]), "any") {
                            innerMap["service"] = ivList[0]
                        } else {
                            innerMap["service"] = ivList
                        }
                    }
                    if v := d.Get(fmt.Sprintf("access_layers_content.%d.rulebase.%d.service_negate", i, j)).(bool); v {
                        innerMap["service-negate"] = v
                    }
                    if len(innerMap) > 0 {
                        rulebaseArr = append(rulebaseArr, innerMap)
                    }
                }
                if len(rulebaseArr) > 0 {
                    itemMap["rulebase"] = rulebaseArr
                }
            }
            if len(itemMap) > 0 {
                accesslayerscontentArray = append(accesslayerscontentArray, itemMap)
            }
        }
        if len(accesslayerscontentArray) > 0 {
            payload["access-layers-content"] = accesslayerscontentArray
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Execute set-dynamic-content - Payload = ", payload)

    GaiaSetDynamicContentRes, err := client.ApiCallSimple("set-dynamic-content", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaSetDynamicContentRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaSetDynamicContentRes.Success {
            errMsg = GaiaSetDynamicContentRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaSetDynamicContentRes.GetData()
        }

        debugLogOperation(
            "set-dynamic-content",        // resource type
            "command",                       // operation
            "set-dynamic-content",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute set-dynamic-content: %v", err)
    }
    if !GaiaSetDynamicContentRes.Success {
        if GaiaSetDynamicContentRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaSetDynamicContentRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-dynamic-content", GaiaSetDynamicContentRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-dynamic-content task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        msg := taskRes.Message
        if msg == "" {
            msg = fmt.Sprintf("set-dynamic-content task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
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
        _respData = GaiaSetDynamicContentRes.GetData()
    }
    d.Set("validation_warnings", []interface{}{})
    if v, exists := _respData["validation-warnings"]; exists {
        if arr, ok := v.([]interface{}); ok {
            var _vwResult []interface{}
            for _, item := range arr {
                if m, ok := item.(map[string]interface{}); ok {
                    _vwResult = append(_vwResult, map[string]interface{}{
                        "layer":   fmt.Sprintf("%v", m["layer"]),
                        "rule":    fmt.Sprintf("%v", m["rule"]),
                        "object":  fmt.Sprintf("%v", m["object"]),
                        "message": fmt.Sprintf("%v", m["message"]),
                    })
                }
            }
            d.Set("validation_warnings", _vwResult)
        }
    }
    d.Set("validation_errors", []interface{}{})
    if v, exists := _respData["validation-errors"]; exists {
        if arr, ok := v.([]interface{}); ok {
            var _veResult []interface{}
            for _, item := range arr {
                if m, ok := item.(map[string]interface{}); ok {
                    _veResult = append(_veResult, map[string]interface{}{
                        "layer":      fmt.Sprintf("%v", m["layer"]),
                        "rule":       fmt.Sprintf("%v", m["rule"]),
                        "object":     fmt.Sprintf("%v", m["object"]),
                        "message":    fmt.Sprintf("%v", m["message"]),
                        "error_code": fmt.Sprintf("%v", m["error-code"]),
                    })
                }
            }
            d.Set("validation_errors", _veResult)
        }
    }
    if v, exists := _respData["change-summary"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            var layerNames []interface{}
            if arr, ok := m["layers"].([]interface{}); ok {
                for _, item := range arr {
                    if obj, ok := item.(map[string]interface{}); ok {
                        if name, ok := obj["name"].(string); ok {
                            layerNames = append(layerNames, name)
                        }
                    } else {
                        layerNames = append(layerNames, fmt.Sprintf("%v", item))
                    }
                }
            }
            var objectNames []interface{}
            if arr, ok := m["objects"].([]interface{}); ok {
                for _, item := range arr {
                    if obj, ok := item.(map[string]interface{}); ok {
                        if name, ok := obj["name"].(string); ok {
                            objectNames = append(objectNames, name)
                        }
                    } else {
                        objectNames = append(objectNames, fmt.Sprintf("%v", item))
                    }
                }
            }
            d.Set("change_summary", []interface{}{map[string]interface{}{
                "layers":  layerNames,
                "objects": objectNames,
            }})
        }
    }


    d.SetId(fmt.Sprintf("set-dynamic-content-" + acctest.RandString(10)))
    return nil
}

func readGaiaSetDynamicContent(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaSetDynamicContent(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

