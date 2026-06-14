package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaRouteRedistributionToIsis() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaRouteRedistributionToIsis,
        Read:   readGaiaRouteRedistributionToIsis,
        Update: updateGaiaRouteRedistributionToIsis,
        Delete: deleteGaiaRouteRedistributionToIsis,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "from": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure policy for exporting routes to IS-IS`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "aggregate": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Specifies the aggregate route to redistribute into IS-IS`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "all_ipv4_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches all IPv4 aggregate routes`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches all IPv6 aggregate routes<br><br>Note: IPv6 state must be enabled`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches specific IPv4 or IPv6 aggregate routes. The aggregate routes have to be already configured.<br><br>Note: IPv6 state must be enabled for IPv6 aggregate routes.`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv4 or IPv6 network`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the metric to be added to routes redistributed via this rule`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "kernel": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Redistribution of kernel routes into IS-IS.<br><br>Note: It may be inadvisable in certain cases to redistribute kernel routes into another protocol. Kernel routes usually exist upon startup of routed, before the routing table has settled, when error conditions or bad routes may be present. Use caution when configuring route redistribution from the kernel.`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "all_ipv4_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies<br><br>Note: IPv6 state must be enabled`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.<br><br>Note: When network objects are specified, previous objects will be overwritten`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv4 or IPv6 network`,
                                                },
                                                "restrict": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted`,
                                                },
                                                "match_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Defines how routes are matched to the network. The match types are as follows:<br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table>`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule`,
                                                },
                                                "range": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Description: `Specifies the mask length range<br><br>Note: The match-type needs to be of type \"range\" and network must be IPv4`,
                                                    MaxItems:    1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "from": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the lower limit of the range of mask lengths`,
                                                            },
                                                            "to": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the upper limit of the range of mask lengths`,
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "nat_pool": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Redistribution of NAT pools into IS-IS`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "all_ipv4_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches all IPv4 NAT pools`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches all IPv6 NAT pools`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches specific IPv4 or IPv6 NAT pools. The NAT pool has to be already configured.<br><br>Note: IPv6 state must be enabled for IPv6 NAT pools.`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv4 or IPv6 network`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the metric to be added to routes redistributed via this rule`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "rip": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Redistribution of RIP routes into IS-IS`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "all_ipv4_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this configuration to all routes from the given protocol described by an IPv4 network, unless a more specific route redistribution rule applies.<br><br>Note: When network objects are specified, previous objects will be overwritten`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv4 network`,
                                                },
                                                "restrict": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted`,
                                                },
                                                "match_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Defines how routes are matched to the network. The match types are as follows:<br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table>`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule`,
                                                },
                                                "range": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Description: `Specifies the mask length range<br><br>Note: The match-type needs to be of type \"range\"`,
                                                    MaxItems:    1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "from": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the lower limit of the range of mask lengths`,
                                                            },
                                                            "to": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the upper limit of the range of mask lengths`,
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "ripng": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Redistribution of RIPng routes into IS-IS.<br><br>Note: IPv6 state needs to be enabled.`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this configuration to all routes from the given protocol described by an IPv6 network, unless a more specific route redistribution rule applies.<br><br>Note: When network objects are specified, previous objects will be overwritten`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 network`,
                                                },
                                                "restrict": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted`,
                                                },
                                                "match_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Defines how routes are matched to the network. The match types are as follows:<br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table>`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "static_route": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Redistribution of static routes into IS-IS`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "all_ipv4_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches all IPv4 static route`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "default": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches the default IPv4 static route`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches all IPv6 static route`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "default6": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches the default IPv6 static route`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches specific IPv4 or IPv6 static routes. The static route has to be already configured.<br><br>Note: IPv6 state must be enabled for IPv6 static routes.`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv4 or IPv6 network`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the metric to be added to routes redistributed via this rule`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "bgp_as_number": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Configures Autonomous System numbers of the BGP group from which to export routes into IS-IS`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "as_number": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Configured Autonomous System Number. Valid Values are 1 - 4294967295 or 0.1 - 65535.65535.<br><br>The ASN format can be changed to dotted or plain format using the following command 'set format asn dotted/plain'.`,
                                    },
                                    "all_ipv4_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies<br><br>Note: IPv6 state must be enabled`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.<br><br>Note: When network objects are specified, previous objects will be overwritten`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv4 or IPv6 network`,
                                                },
                                                "restrict": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted`,
                                                },
                                                "match_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Defines how routes are matched to the network. The match types are as follows:<br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table>`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule`,
                                                },
                                                "range": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Description: `Specifies the mask length range<br><br>Note: The match-type needs to be of type \"range\" and network must be IPv4`,
                                                    MaxItems:    1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "from": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the lower limit of the range of mask lengths`,
                                                            },
                                                            "to": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the upper limit of the range of mask lengths`,
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "bgp_as_path": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Configures the redistribution of BGP routes, whose AS path matches a given regular expression into IS-IS`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "aspath_regex": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Configures the redistribution of BGP routes, whose AS path matches the given regular expression.<br><br>Valid Values are regular expressions surrounded by double quotes (\"). The regular expression can only have digits, a colon (:) and the following special characters:<br><br><table class=\"table\"><tr> <th>Regular Expression</th> <th>Description</th> </tr><tr> <td>.</td> <td>Match any single character</td> </tr><tr> <td>\</td> <td>Match the character right after the backslash. Also for recalling</td> </tr><tr> <td>^</td> <td>Match the characters or null string at the beginning of the value</td> </tr><tr> <td>$</td> <td>Match the characters or null string at the end of the value</td> </tr><tr> <td>?</td> <td>Match zero or one occurrences of the pattern before the '?' character</td> </tr><tr> <td>*</td> <td>Match zero or more occurrences of the pattern before the '*' character</td> </tr><tr> <td>+</td> <td>Match one or more occurrences of the pattern before the '+' character</td> </tr><tr> <td>|</td> <td>Match one of the patterns on either side of the '|' character</td> </tr><tr> <td>_</td> <td>Match comma (,), left brace ({), right brace (}), beginning of value (^), end of value ($) or a whitespace</td> </tr><tr> <td>[]</td> <td>Match set of characters or a range of characters separated by a hyphen (-) within []</td> </tr><tr> <td>()</td> <td>Group one or more patterns into a single pattern</td> </tr><tr> <td>{m,n}</td> <td>At least m and at most n repetitions of the pattern before {m,n}</td> </tr><tr> <td>{m}</td> <td>Exactly m repetitions of the pattern before {m}</td> </tr><tr> <td>{m,}</td> <td>m or more repetitions of the pattern before {m}</td> </tr></table>`,
                                    },
                                    "origin": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Specifies the completeness of the AS path information. Only a single origin should be used with a regular expression.<br><br>Any - Matches any routes, regardless of origin.<br>IGP - Route was learned from an interior routing protocol and the AS path is probably complete.<br>EGP - Route was learned from an exterior routing protocol that does not support AS paths and the path is probably incomplete.<br>incomplete - Use when the AS path information is incomplete.`,
                                    },
                                    "all_ipv4_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies<br><br>Note: IPv6 state must be enabled`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this configuration to all routes from the given protocol described by a network, unless a more specific route redistribution rule applies.<br><br>Note: When network objects are specified, previous objects will be overwritten`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv4 or IPv6 network`,
                                                },
                                                "restrict": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted`,
                                                },
                                                "match_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Defines how routes are matched to the network. The match types are as follows:<br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table>`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule`,
                                                },
                                                "range": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Description: `Specifies the mask length range<br><br>Note: The match-type needs to be of type \"range\" and network must be IPv4`,
                                                    MaxItems:    1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "from": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the lower limit of the range of mask lengths`,
                                                            },
                                                            "to": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the upper limit of the range of mask lengths`,
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "interface": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Configures the redistribution of all directly connected routes from an interface into IS-IS`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "interface": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Specifies the name of the interface`,
                                    },
                                    "metric": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule<br><br>Wide metric type: 1 - 4294967295<br>Narrow metric type: 1 - 63`,
                                    },
                                },
                            },
                        },
                        "ospf2": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Configures the redistribution of IPv4 OSPF routes into IS-IS`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "instance": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Redistribute routes from a specific OSPF instance`,
                                    },
                                    "all_ipv4_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this configuration to all routes from the given protocol described by an IPv4 network, unless a more specific route redistribution rule applies.<br><br>Note: When network objects are specified, previous objects will be overwritten`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv4 network`,
                                                },
                                                "restrict": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted`,
                                                },
                                                "match_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Defines how routes are matched to the network. The match types are as follows:<br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table>`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule`,
                                                },
                                                "range": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Description: `Specifies the mask length range<br><br>Note: The match-type needs to be of type \"range\"`,
                                                    MaxItems:    1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "from": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the lower limit of the range of mask lengths`,
                                                            },
                                                            "to": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the upper limit of the range of mask lengths`,
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "ospf2ase": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Configures the redistribution of OSPF Autonomous System External routes into IS-IS`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "instance": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Redistribute routes from a specific OSPF instance`,
                                    },
                                    "all_ipv4_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv4 routes from this protocol, unless a more specific route redistribution rule applies`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this configuration to all routes from the given protocol described by an IPv4 network, unless a more specific route redistribution rule applies.<br><br>Note: When network objects are specified, previous objects will be overwritten`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv4 network`,
                                                },
                                                "restrict": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted`,
                                                },
                                                "match_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Defines how routes are matched to the network. The match types are as follows:<br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table>`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule`,
                                                },
                                                "range": {
                                                    Type:        schema.TypeList,
                                                    Optional:    true,
                                                    Description: `Specifies the mask length range<br><br>Note: The match-type needs to be of type \"range\"`,
                                                    MaxItems:    1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "from": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the lower limit of the range of mask lengths`,
                                                            },
                                                            "to": {
                                                                Type:        schema.TypeInt,
                                                                Optional:    true,
                                                                Description: `Specifies the upper limit of the range of mask lengths`,
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "ospf3": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Configures the redistribution of IPv6 OSPF routes into IS-IS.<br><br>Note: IPv6 state needs to be enabled.`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "instance": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Redistribute routes from a specific OSPF instance`,
                                    },
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this configuration to all routes from the given protocol described by an IPv6 network, unless a more specific route redistribution rule applies.<br><br>Note: When network objects are specified, previous objects will be overwritten`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 network`,
                                                },
                                                "restrict": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted`,
                                                },
                                                "match_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Defines how routes are matched to the network. The match types are as follows:<br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table>`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                        "ospf3ase": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Configures the redistribution of IPv6 OSPF Autonomous System External routes into IS-IS.<br><br>Note: IPv6 state needs to be enabled.`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "instance": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Redistribute routes from a specific OSPF instance`,
                                    },
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this route redistrution rule to all IPv6 routes from this protocol, unless a more specific route redistribution rule applies`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IS-IS metric value to routes matching this rule`,
                                                },
                                                "enable": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Enables or disables the metric value`,
                                                },
                                            },
                                        },
                                    },
                                    "network": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Applies this configuration to all routes from the given protocol described by an IPv6 network, unless a more specific route redistribution rule applies.<br><br>Note: When network objects are specified, previous objects will be overwritten`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 network`,
                                                },
                                                "restrict": {
                                                    Type:        schema.TypeBool,
                                                    Optional:    true,
                                                    Description: `Specifies whether to accept or restrict routes that match the given rule. By default routes are accepted`,
                                                },
                                                "match_type": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Defines how routes are matched to the network. The match types are as follows:<br><br><table class=\"table\"><tr> <th>Match Type</th> <th>Description</th> </tr><tr> <td>Normal</td> <td>Matches any route contained within the specified network</td> </tr><tr> <td>Exact</td> <td>Matches only routes with the prefix and mask length exactly equal to the specified network</td> </tr><tr> <td>Refines</td> <td>Matches only routes that are more specific than the specified network</td> </tr><tr> <td>Range</td> <td>Matches any route whose IP prefix equals the specified network and whose mask length falls within the specified mask length range (Network needs to be IPv4 in order to specify this value)</td> </tr></table>`,
                                                },
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies the IS-IS metric to be added to routes redistributed via this rule`,
                                                },
                                            },
                                        },
                                    },
                                },
                            },
                        },
                    },
                },
            },
            "level": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configures IS-IS for the specified level`,
            },
            "reset": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Removes IS-IS Route Redistribution configuration`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaRouteRedistributionToIsis(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("from"); len(v.([]interface{})) > 0 {
        _ = v
        fromMap := make(map[string]interface{})
        if v, ok := d.GetOk("from.0.aggregate"); ok {
            _ = v
            aggregateMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.aggregate.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.aggregate.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.aggregate.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    aggregateMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.aggregate.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.aggregate.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.aggregate.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    aggregateMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.aggregate.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.aggregate.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.aggregate.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    aggregateMap["network"] = networkArray
                }
            }
            if len(aggregateMap) > 0 {
                fromMap["aggregate"] = aggregateMap
            }
        }
        if v, ok := d.GetOk("from.0.kernel"); ok {
            _ = v
            kernelMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.kernel.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.kernel.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.kernel.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    kernelMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.kernel.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.kernel.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.kernel.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    kernelMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.kernel.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v := d.Get(fmt.Sprintf("from.0.kernel.0.network.%d.restrict", i)).(bool); v {
                        itemMap["restrict"] = v
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.match_type", i)); ok {
                        itemMap["match-type"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if sv, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.range", i)); ok {
                        if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                            rawDict := ivList[0].(map[string]interface{})
                            rangeMap := make(map[string]interface{})
                            if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                rangeMap["from"] = sv.(int)
                            }
                            if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                rangeMap["to"] = sv.(int)
                            }
                            if len(rangeMap) > 0 {
                                itemMap["range"] = rangeMap
                            }
                        }
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    kernelMap["network"] = networkArray
                }
            }
            if len(kernelMap) > 0 {
                fromMap["kernel"] = kernelMap
            }
        }
        if v, ok := d.GetOk("from.0.nat_pool"); ok {
            _ = v
            natpoolMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.nat_pool.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    natpoolMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.nat_pool.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    natpoolMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.nat_pool.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.nat_pool.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.nat_pool.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    natpoolMap["network"] = networkArray
                }
            }
            if len(natpoolMap) > 0 {
                fromMap["nat-pool"] = natpoolMap
            }
        }
        if v, ok := d.GetOk("from.0.rip"); ok {
            _ = v
            ripMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.rip.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.rip.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.rip.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    ripMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.rip.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.rip.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v := d.Get(fmt.Sprintf("from.0.rip.0.network.%d.restrict", i)).(bool); v {
                        itemMap["restrict"] = v
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.rip.0.network.%d.match_type", i)); ok {
                        itemMap["match-type"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.rip.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if sv, ok := d.GetOk(fmt.Sprintf("from.0.rip.0.network.%d.range", i)); ok {
                        if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                            rawDict := ivList[0].(map[string]interface{})
                            rangeMap := make(map[string]interface{})
                            if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                rangeMap["from"] = sv.(int)
                            }
                            if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                rangeMap["to"] = sv.(int)
                            }
                            if len(rangeMap) > 0 {
                                itemMap["range"] = rangeMap
                            }
                        }
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    ripMap["network"] = networkArray
                }
            }
            if len(ripMap) > 0 {
                fromMap["rip"] = ripMap
            }
        }
        if v, ok := d.GetOk("from.0.ripng"); ok {
            _ = v
            ripngMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.ripng.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.ripng.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.ripng.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    ripngMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.ripng.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.ripng.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v := d.Get(fmt.Sprintf("from.0.ripng.0.network.%d.restrict", i)).(bool); v {
                        itemMap["restrict"] = v
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.ripng.0.network.%d.match_type", i)); ok {
                        itemMap["match-type"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.ripng.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    ripngMap["network"] = networkArray
                }
            }
            if len(ripngMap) > 0 {
                fromMap["ripng"] = ripngMap
            }
        }
        if v, ok := d.GetOk("from.0.static_route"); ok {
            _ = v
            staticrouteMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.static_route.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    staticrouteMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.static_route.0.default"); ok {
                _ = v
                defaultMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.default.0.metric"); ok {
                    defaultMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.default.0.enable"); ok && v.(bool) {
                    defaultMap["enable"] = v.(bool)
                }
                if len(defaultMap) > 0 {
                    staticrouteMap["default"] = defaultMap
                }
            }
            if v, ok := d.GetOk("from.0.static_route.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    staticrouteMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.static_route.0.default6"); ok {
                _ = v
                default6Map := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.default6.0.metric"); ok {
                    default6Map["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.default6.0.enable"); ok && v.(bool) {
                    default6Map["enable"] = v.(bool)
                }
                if len(default6Map) > 0 {
                    staticrouteMap["default6"] = default6Map
                }
            }
            if v, ok := d.GetOk("from.0.static_route.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.static_route.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.static_route.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    staticrouteMap["network"] = networkArray
                }
            }
            if len(staticrouteMap) > 0 {
                fromMap["static-route"] = staticrouteMap
            }
        }
        if v, ok := d.GetOk("from.0.bgp_as_number"); ok {
            bgpasnumberList := v.([]interface{})
            bgpasnumberArray := make([]interface{}, 0, len(bgpasnumberList))
            for i := range bgpasnumberList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.as_number", i)); ok {
                    itemMap["as-number"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.all_ipv4_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv4_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv4_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv4_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv4_routesMap) > 0 {
                            itemMap["all-ipv4-routes"] = all_ipv4_routesMap
                        }
                    }
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.bgp_as_number.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.range", i, j)); ok {
                            if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                                rawDict := ivList[0].(map[string]interface{})
                                rangeMap := make(map[string]interface{})
                                if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                    rangeMap["from"] = sv.(int)
                                }
                                if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                    rangeMap["to"] = sv.(int)
                                }
                                if len(rangeMap) > 0 {
                                    innerMap["range"] = rangeMap
                                }
                            }
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    bgpasnumberArray = append(bgpasnumberArray, itemMap)
                }
            }
            if len(bgpasnumberArray) > 0 {
                fromMap["bgp-as-number"] = bgpasnumberArray
            }
        }
        if v, ok := d.GetOk("from.0.bgp_as_path"); ok {
            bgpaspathList := v.([]interface{})
            bgpaspathArray := make([]interface{}, 0, len(bgpaspathList))
            for i := range bgpaspathList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.aspath_regex", i)); ok {
                    itemMap["aspath-regex"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.origin", i)); ok {
                    itemMap["origin"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.all_ipv4_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv4_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv4_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv4_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv4_routesMap) > 0 {
                            itemMap["all-ipv4-routes"] = all_ipv4_routesMap
                        }
                    }
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.bgp_as_path.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.range", i, j)); ok {
                            if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                                rawDict := ivList[0].(map[string]interface{})
                                rangeMap := make(map[string]interface{})
                                if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                    rangeMap["from"] = sv.(int)
                                }
                                if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                    rangeMap["to"] = sv.(int)
                                }
                                if len(rangeMap) > 0 {
                                    innerMap["range"] = rangeMap
                                }
                            }
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    bgpaspathArray = append(bgpaspathArray, itemMap)
                }
            }
            if len(bgpaspathArray) > 0 {
                fromMap["bgp-as-path"] = bgpaspathArray
            }
        }
        if v, ok := d.GetOk("from.0.interface"); ok {
            interfaceList := v.([]interface{})
            interfaceArray := make([]interface{}, 0, len(interfaceList))
            for i := range interfaceList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.interface.%d.interface", i)); ok {
                    itemMap["interface"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("from.0.interface.%d.metric", i)); ok {
                    itemMap["metric"] = v.(string)
                }
                if len(itemMap) > 0 {
                    interfaceArray = append(interfaceArray, itemMap)
                }
            }
            if len(interfaceArray) > 0 {
                fromMap["interface"] = interfaceArray
            }
        }
        if v, ok := d.GetOk("from.0.ospf2"); ok {
            ospf2List := v.([]interface{})
            ospf2Array := make([]interface{}, 0, len(ospf2List))
            for i := range ospf2List {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.instance", i)); ok {
                    itemMap["instance"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.all_ipv4_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv4_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv4_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv4_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv4_routesMap) > 0 {
                            itemMap["all-ipv4-routes"] = all_ipv4_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.ospf2.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.ospf2.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.network.%d.range", i, j)); ok {
                            if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                                rawDict := ivList[0].(map[string]interface{})
                                rangeMap := make(map[string]interface{})
                                if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                    rangeMap["from"] = sv.(int)
                                }
                                if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                    rangeMap["to"] = sv.(int)
                                }
                                if len(rangeMap) > 0 {
                                    innerMap["range"] = rangeMap
                                }
                            }
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    ospf2Array = append(ospf2Array, itemMap)
                }
            }
            if len(ospf2Array) > 0 {
                fromMap["ospf2"] = ospf2Array
            }
        }
        if v, ok := d.GetOk("from.0.ospf2ase"); ok {
            ospf2aseList := v.([]interface{})
            ospf2aseArray := make([]interface{}, 0, len(ospf2aseList))
            for i := range ospf2aseList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.instance", i)); ok {
                    itemMap["instance"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.all_ipv4_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv4_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv4_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv4_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv4_routesMap) > 0 {
                            itemMap["all-ipv4-routes"] = all_ipv4_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.ospf2ase.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.range", i, j)); ok {
                            if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                                rawDict := ivList[0].(map[string]interface{})
                                rangeMap := make(map[string]interface{})
                                if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                    rangeMap["from"] = sv.(int)
                                }
                                if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                    rangeMap["to"] = sv.(int)
                                }
                                if len(rangeMap) > 0 {
                                    innerMap["range"] = rangeMap
                                }
                            }
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    ospf2aseArray = append(ospf2aseArray, itemMap)
                }
            }
            if len(ospf2aseArray) > 0 {
                fromMap["ospf2ase"] = ospf2aseArray
            }
        }
        if v, ok := d.GetOk("from.0.ospf3"); ok {
            ospf3List := v.([]interface{})
            ospf3Array := make([]interface{}, 0, len(ospf3List))
            for i := range ospf3List {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.instance", i)); ok {
                    itemMap["instance"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.ospf3.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.ospf3.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    ospf3Array = append(ospf3Array, itemMap)
                }
            }
            if len(ospf3Array) > 0 {
                fromMap["ospf3"] = ospf3Array
            }
        }
        if v, ok := d.GetOk("from.0.ospf3ase"); ok {
            ospf3aseList := v.([]interface{})
            ospf3aseArray := make([]interface{}, 0, len(ospf3aseList))
            for i := range ospf3aseList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.instance", i)); ok {
                    itemMap["instance"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.ospf3ase.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    ospf3aseArray = append(ospf3aseArray, itemMap)
                }
            }
            if len(ospf3aseArray) > 0 {
                fromMap["ospf3ase"] = ospf3aseArray
            }
        }
        if len(fromMap) > 0 {
            payload["from"] = fromMap
        }
    }

    if v, ok := d.GetOk("level"); ok {
        payload["level"] = v.(string)
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    log.Println("Create RouteRedistributionToIsis - Map = ", payload)

    addRouteRedistributionToIsisRes, err := client.ApiCallSimple("set-route-redistribution-to-isis", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addRouteRedistributionToIsisRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addRouteRedistributionToIsisRes.Success {
            errMsg = addRouteRedistributionToIsisRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addRouteRedistributionToIsisRes.GetData()
        }

        debugLogOperation(
            "route-redistribution-to-isis",        // resource type
            "create",                       // operation
            "set-route-redistribution-to-isis",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add route-redistribution-to-isis: %v", err)
    }
    if !addRouteRedistributionToIsisRes.Success {
        if addRouteRedistributionToIsisRes.ErrorMsg != "" {
            return fmt.Errorf(addRouteRedistributionToIsisRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("route-redistribution-to-isis-" + acctest.RandString(10)))
    return readGaiaRouteRedistributionToIsis(d, m)
}

func readGaiaRouteRedistributionToIsis(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("level"); ok {
        payload["level"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showRouteRedistributionToIsisRes, err := client.ApiCallSimple("show-route-redistribution-to-isis", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showRouteRedistributionToIsisRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showRouteRedistributionToIsisRes.Success {
            errMsg = showRouteRedistributionToIsisRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showRouteRedistributionToIsisRes.GetData()
        }

        debugLogOperation(
            "route-redistribution-to-isis",        // resource type
            "read",                       // operation
            "show-route-redistribution-to-isis",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show route-redistribution-to-isis: %v", err)
    }
    if !showRouteRedistributionToIsisRes.Success {
        if data := showRouteRedistributionToIsisRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showRouteRedistributionToIsisRes.ErrorMsg)
    }

    routeRedistributionToIsis := showRouteRedistributionToIsisRes.GetData()

    log.Println("Read RouteRedistributionToIsis - Show JSON = ", routeRedistributionToIsis)

    if v, exists := routeRedistributionToIsis["isis"]; exists {
        if isisItems, ok := v.([]interface{}); ok && len(isisItems) > 0 {
            apiItem, _ := isisItems[0].(map[string]interface{})
            if val, ok := apiItem["level"]; ok {
                d.Set("level", fmt.Sprintf("%v", val))
            }
            apiFrom, _ := apiItem["from"].(map[string]interface{})
            if apiFrom == nil {
                apiFrom = map[string]interface{}{}
            }
            fromObj := map[string]interface{}{}
            buildAIR := func(m map[string]interface{}) map[string]interface{} {
                r := map[string]interface{}{}
                if val, ok := m["metric"]; ok { r["metric"] = fmt.Sprintf("%v", val) }
                if val, ok := m["enable"]; ok { if b, ok := val.(bool); ok { r["enable"] = b } }
                return r
            }
            buildNet := func(lst []interface{}, hasRestrict bool) []interface{} {
                out := make([]interface{}, 0, len(lst))
                for _, rw := range lst {
                    if nm, ok := rw.(map[string]interface{}); ok {
                        ni := map[string]interface{}{}
                        if val, ok := nm["address"]; ok { ni["address"] = fmt.Sprintf("%v", val) }
                        if hasRestrict {
                            if val, ok := nm["restrict"]; ok { if b, ok := val.(bool); ok { ni["restrict"] = b } }
                            if val, ok := nm["match-type"]; ok { ni["match_type"] = fmt.Sprintf("%v", val) }
                        }
                        if val, ok := nm["metric"]; ok { ni["metric"] = fmt.Sprintf("%v", val) }
                        if rv, ok := nm["range"]; ok {
                            if rm, ok := rv.(map[string]interface{}); ok {
                                re := map[string]interface{}{}
                                if fv, ok := rm["from"]; ok { var n int; if _, e := fmt.Sscanf(fmt.Sprintf("%v", fv), "%d", &n); e == nil { re["from"] = n } }
                                if tv, ok := rm["to"]; ok { var n int; if _, e := fmt.Sscanf(fmt.Sprintf("%v", tv), "%d", &n); e == nil { re["to"] = n } }
                                if len(re) > 0 { ni["range"] = []interface{}{re} }
                            }
                        }
                        if len(ni) > 0 { out = append(out, ni) }
                    }
                }
                return out
            }
            // aggregate: both ipv4 and ipv6, no range
            if raw, ok := apiFrom["aggregate"]; ok {
                if pm, ok := raw.(map[string]interface{}); ok {
                    protoObj := map[string]interface{}{}
                    if ar, ok := pm["all-ipv4-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv4_routes"] = []interface{}{buildAIR(am)} } }
                    if ar, ok := pm["all-ipv6-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv6_routes"] = []interface{}{buildAIR(am)} } }
                    if nr, ok := pm["network"]; ok { if nl, ok := nr.([]interface{}); ok { protoObj["network"] = buildNet(nl, false) } }
                    fromObj["aggregate"] = []interface{}{protoObj}
                }
            }
            // kernel: both ipv4 and ipv6, with range
            if raw, ok := apiFrom["kernel"]; ok {
                if pm, ok := raw.(map[string]interface{}); ok {
                    protoObj := map[string]interface{}{}
                    if ar, ok := pm["all-ipv4-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv4_routes"] = []interface{}{buildAIR(am)} } }
                    if ar, ok := pm["all-ipv6-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv6_routes"] = []interface{}{buildAIR(am)} } }
                    if nr, ok := pm["network"]; ok { if nl, ok := nr.([]interface{}); ok { protoObj["network"] = buildNet(nl, true) } }
                    fromObj["kernel"] = []interface{}{protoObj}
                }
            }
            // nat_pool: both ipv4 and ipv6, no range
            if raw, ok := apiFrom["nat-pool"]; ok {
                if pm, ok := raw.(map[string]interface{}); ok {
                    protoObj := map[string]interface{}{}
                    if ar, ok := pm["all-ipv4-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv4_routes"] = []interface{}{buildAIR(am)} } }
                    if ar, ok := pm["all-ipv6-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv6_routes"] = []interface{}{buildAIR(am)} } }
                    if nr, ok := pm["network"]; ok { if nl, ok := nr.([]interface{}); ok { protoObj["network"] = buildNet(nl, false) } }
                    fromObj["nat_pool"] = []interface{}{protoObj}
                }
            }
            // rip: ipv4 only, with range
            if raw, ok := apiFrom["rip"]; ok {
                if pm, ok := raw.(map[string]interface{}); ok {
                    protoObj := map[string]interface{}{}
                    if ar, ok := pm["all-ipv4-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv4_routes"] = []interface{}{buildAIR(am)} } }
                    if nr, ok := pm["network"]; ok { if nl, ok := nr.([]interface{}); ok { protoObj["network"] = buildNet(nl, true) } }
                    fromObj["rip"] = []interface{}{protoObj}
                }
            }
            // ripng: ipv6 only, with range
            if raw, ok := apiFrom["ripng"]; ok {
                if pm, ok := raw.(map[string]interface{}); ok {
                    protoObj := map[string]interface{}{}
                    if ar, ok := pm["all-ipv6-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv6_routes"] = []interface{}{buildAIR(am)} } }
                    if nr, ok := pm["network"]; ok { if nl, ok := nr.([]interface{}); ok { protoObj["network"] = buildNet(nl, true) } }
                    fromObj["ripng"] = []interface{}{protoObj}
                }
            }
            // static-route: both ipv4 and ipv6, default and default6
            if raw, ok := apiFrom["static-route"]; ok {
                if pm, ok := raw.(map[string]interface{}); ok {
                    protoObj := map[string]interface{}{}
                    if ar, ok := pm["all-ipv4-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv4_routes"] = []interface{}{buildAIR(am)} } }
                    if ar, ok := pm["all-ipv6-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { protoObj["all_ipv6_routes"] = []interface{}{buildAIR(am)} } }
                    if dr, ok := pm["default"]; ok { if dm, ok := dr.(map[string]interface{}); ok { protoObj["default"] = []interface{}{buildAIR(dm)} } }
                    if dr, ok := pm["default6"]; ok { if dm, ok := dr.(map[string]interface{}); ok { protoObj["default6"] = []interface{}{buildAIR(dm)} } }
                    if nr, ok := pm["network"]; ok { if nl, ok := nr.([]interface{}); ok { protoObj["network"] = buildNet(nl, false) } }
                    fromObj["static_route"] = []interface{}{protoObj}
                }
            }
            // bgp-as-number (list): both ipv4 and ipv6, with range
            if raw, ok := apiFrom["bgp-as-number"]; ok {
                if items, ok := raw.([]interface{}); ok {
                    out := make([]interface{}, 0, len(items))
                    for _, ri := range items {
                        if im, ok := ri.(map[string]interface{}); ok {
                            item := map[string]interface{}{}
                            if val, ok := im["as-number"]; ok { item["as_number"] = fmt.Sprintf("%v", val) }
                            if ar, ok := im["all-ipv4-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { item["all_ipv4_routes"] = []interface{}{buildAIR(am)} } }
                            if ar, ok := im["all-ipv6-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { item["all_ipv6_routes"] = []interface{}{buildAIR(am)} } }
                            if nr, ok := im["network"]; ok { if nl, ok := nr.([]interface{}); ok { item["network"] = buildNet(nl, true) } }
                            if len(item) > 0 { out = append(out, item) }
                        }
                    }
                    fromObj["bgp_as_number"] = out
                }
            }
            // bgp-as-path (list): both ipv4 and ipv6, with range
            if raw, ok := apiFrom["bgp-as-path"]; ok {
                if items, ok := raw.([]interface{}); ok {
                    out := make([]interface{}, 0, len(items))
                    for _, ri := range items {
                        if im, ok := ri.(map[string]interface{}); ok {
                            item := map[string]interface{}{}
                            if val, ok := im["aspath-regex"]; ok { item["aspath_regex"] = fmt.Sprintf("%v", val) }
                            if val, ok := im["origin"]; ok { item["origin"] = fmt.Sprintf("%v", val) }
                            if ar, ok := im["all-ipv4-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { item["all_ipv4_routes"] = []interface{}{buildAIR(am)} } }
                            if ar, ok := im["all-ipv6-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { item["all_ipv6_routes"] = []interface{}{buildAIR(am)} } }
                            if nr, ok := im["network"]; ok { if nl, ok := nr.([]interface{}); ok { item["network"] = buildNet(nl, true) } }
                            if len(item) > 0 { out = append(out, item) }
                        }
                    }
                    fromObj["bgp_as_path"] = out
                }
            }
            // interface (list)
            if raw, ok := apiFrom["interface"]; ok {
                if items, ok := raw.([]interface{}); ok {
                    out := make([]interface{}, 0, len(items))
                    for _, ri := range items {
                        if im, ok := ri.(map[string]interface{}); ok {
                            item := map[string]interface{}{}
                            if val, ok := im["interface"]; ok { item["interface"] = fmt.Sprintf("%v", val) }
                            if val, ok := im["metric"]; ok { item["metric"] = fmt.Sprintf("%v", val) }
                            if len(item) > 0 { out = append(out, item) }
                        }
                    }
                    fromObj["interface"] = out
                }
            }
            // ospf2, ospf2ase (list, ipv4, with range)
            for _, e2 := range []struct{ a, t string }{
                {"ospf2",    "ospf2"},
                {"ospf2ase", "ospf2ase"},
            } {
                if raw, ok := apiFrom[e2.a]; ok {
                    if items, ok := raw.([]interface{}); ok {
                        out := make([]interface{}, 0, len(items))
                        for _, ri := range items {
                            if im, ok := ri.(map[string]interface{}); ok {
                                item := map[string]interface{}{}
                                if val, ok := im["instance"]; ok { item["instance"] = fmt.Sprintf("%v", val) }
                                if ar, ok := im["all-ipv4-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { item["all_ipv4_routes"] = []interface{}{buildAIR(am)} } }
                                if nr, ok := im["network"]; ok { if nl, ok := nr.([]interface{}); ok { item["network"] = buildNet(nl, true) } }
                                if len(item) > 0 { out = append(out, item) }
                            }
                        }
                        fromObj[e2.t] = out
                    }
                }
            }
            // ospf3, ospf3ase (list, ipv6)
            for _, e2 := range []struct{ a, t string }{
                {"ospf3",    "ospf3"},
                {"ospf3ase", "ospf3ase"},
            } {
                if raw, ok := apiFrom[e2.a]; ok {
                    if items, ok := raw.([]interface{}); ok {
                        out := make([]interface{}, 0, len(items))
                        for _, ri := range items {
                            if im, ok := ri.(map[string]interface{}); ok {
                                item := map[string]interface{}{}
                                if val, ok := im["instance"]; ok { item["instance"] = fmt.Sprintf("%v", val) }
                                if ar, ok := im["all-ipv6-routes"]; ok { if am, ok := ar.(map[string]interface{}); ok { item["all_ipv6_routes"] = []interface{}{buildAIR(am)} } }
                                if nr, ok := im["network"]; ok { if nl, ok := nr.([]interface{}); ok { item["network"] = buildNet(nl, true) } }
                                if len(item) > 0 { out = append(out, item) }
                            }
                        }
                        fromObj[e2.t] = out
                    }
                }
            }
            d.Set("from", []interface{}{fromObj})
        }
    }
    if v, exists := routeRedistributionToIsis["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaRouteRedistributionToIsis(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("from"); len(v.([]interface{})) > 0 {
        _ = v
        fromMap := make(map[string]interface{})
        if v, ok := d.GetOk("from.0.aggregate"); ok {
            _ = v
            aggregateMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.aggregate.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.aggregate.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.aggregate.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    aggregateMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.aggregate.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.aggregate.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.aggregate.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    aggregateMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.aggregate.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.aggregate.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.aggregate.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    aggregateMap["network"] = networkArray
                }
            }
            if len(aggregateMap) > 0 {
                fromMap["aggregate"] = aggregateMap
            }
        }
        if v, ok := d.GetOk("from.0.kernel"); ok {
            _ = v
            kernelMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.kernel.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.kernel.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.kernel.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    kernelMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.kernel.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.kernel.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.kernel.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    kernelMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.kernel.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v := d.Get(fmt.Sprintf("from.0.kernel.0.network.%d.restrict", i)).(bool); v {
                        itemMap["restrict"] = v
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.match_type", i)); ok {
                        itemMap["match-type"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if sv, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.range", i)); ok {
                        if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                            rawDict := ivList[0].(map[string]interface{})
                            rangeMap := make(map[string]interface{})
                            if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                rangeMap["from"] = sv.(int)
                            }
                            if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                rangeMap["to"] = sv.(int)
                            }
                            if len(rangeMap) > 0 {
                                itemMap["range"] = rangeMap
                            }
                        }
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    kernelMap["network"] = networkArray
                }
            }
            if len(kernelMap) > 0 {
                fromMap["kernel"] = kernelMap
            }
        }
        if v, ok := d.GetOk("from.0.nat_pool"); ok {
            _ = v
            natpoolMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.nat_pool.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    natpoolMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.nat_pool.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    natpoolMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.nat_pool.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.nat_pool.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.nat_pool.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    natpoolMap["network"] = networkArray
                }
            }
            if len(natpoolMap) > 0 {
                fromMap["nat-pool"] = natpoolMap
            }
        }
        if v, ok := d.GetOk("from.0.rip"); ok {
            _ = v
            ripMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.rip.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.rip.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.rip.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    ripMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.rip.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.rip.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v := d.Get(fmt.Sprintf("from.0.rip.0.network.%d.restrict", i)).(bool); v {
                        itemMap["restrict"] = v
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.rip.0.network.%d.match_type", i)); ok {
                        itemMap["match-type"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.rip.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if sv, ok := d.GetOk(fmt.Sprintf("from.0.rip.0.network.%d.range", i)); ok {
                        if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                            rawDict := ivList[0].(map[string]interface{})
                            rangeMap := make(map[string]interface{})
                            if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                rangeMap["from"] = sv.(int)
                            }
                            if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                rangeMap["to"] = sv.(int)
                            }
                            if len(rangeMap) > 0 {
                                itemMap["range"] = rangeMap
                            }
                        }
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    ripMap["network"] = networkArray
                }
            }
            if len(ripMap) > 0 {
                fromMap["rip"] = ripMap
            }
        }
        if v, ok := d.GetOk("from.0.ripng"); ok {
            _ = v
            ripngMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.ripng.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.ripng.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.ripng.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    ripngMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.ripng.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.ripng.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v := d.Get(fmt.Sprintf("from.0.ripng.0.network.%d.restrict", i)).(bool); v {
                        itemMap["restrict"] = v
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.ripng.0.network.%d.match_type", i)); ok {
                        itemMap["match-type"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.ripng.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    ripngMap["network"] = networkArray
                }
            }
            if len(ripngMap) > 0 {
                fromMap["ripng"] = ripngMap
            }
        }
        if v, ok := d.GetOk("from.0.static_route"); ok {
            _ = v
            staticrouteMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.static_route.0.all_ipv4_routes"); ok {
                _ = v
                allipv4routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.all_ipv4_routes.0.metric"); ok {
                    allipv4routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.all_ipv4_routes.0.enable"); ok && v.(bool) {
                    allipv4routesMap["enable"] = v.(bool)
                }
                if len(allipv4routesMap) > 0 {
                    staticrouteMap["all-ipv4-routes"] = allipv4routesMap
                }
            }
            if v, ok := d.GetOk("from.0.static_route.0.default"); ok {
                _ = v
                defaultMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.default.0.metric"); ok {
                    defaultMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.default.0.enable"); ok && v.(bool) {
                    defaultMap["enable"] = v.(bool)
                }
                if len(defaultMap) > 0 {
                    staticrouteMap["default"] = defaultMap
                }
            }
            if v, ok := d.GetOk("from.0.static_route.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.all_ipv6_routes.0.enable"); ok && v.(bool) {
                    allipv6routesMap["enable"] = v.(bool)
                }
                if len(allipv6routesMap) > 0 {
                    staticrouteMap["all-ipv6-routes"] = allipv6routesMap
                }
            }
            if v, ok := d.GetOk("from.0.static_route.0.default6"); ok {
                _ = v
                default6Map := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.default6.0.metric"); ok {
                    default6Map["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.default6.0.enable"); ok && v.(bool) {
                    default6Map["enable"] = v.(bool)
                }
                if len(default6Map) > 0 {
                    staticrouteMap["default6"] = default6Map
                }
            }
            if v, ok := d.GetOk("from.0.static_route.0.network"); ok {
                networkList := v.([]interface{})
                networkArray := make([]interface{}, 0, len(networkList))
                for i := range networkList {
                    itemMap := make(map[string]interface{})
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.static_route.0.network.%d.address", i)); ok {
                        itemMap["address"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.static_route.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
                    }
                    if len(itemMap) > 0 {
                        networkArray = append(networkArray, itemMap)
                    }
                }
                if len(networkArray) > 0 {
                    staticrouteMap["network"] = networkArray
                }
            }
            if len(staticrouteMap) > 0 {
                fromMap["static-route"] = staticrouteMap
            }
        }
        if v, ok := d.GetOk("from.0.bgp_as_number"); ok {
            bgpasnumberList := v.([]interface{})
            bgpasnumberArray := make([]interface{}, 0, len(bgpasnumberList))
            for i := range bgpasnumberList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.as_number", i)); ok {
                    itemMap["as-number"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.all_ipv4_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv4_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv4_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv4_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv4_routesMap) > 0 {
                            itemMap["all-ipv4-routes"] = all_ipv4_routesMap
                        }
                    }
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.bgp_as_number.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.range", i, j)); ok {
                            if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                                rawDict := ivList[0].(map[string]interface{})
                                rangeMap := make(map[string]interface{})
                                if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                    rangeMap["from"] = sv.(int)
                                }
                                if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                    rangeMap["to"] = sv.(int)
                                }
                                if len(rangeMap) > 0 {
                                    innerMap["range"] = rangeMap
                                }
                            }
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    bgpasnumberArray = append(bgpasnumberArray, itemMap)
                }
            }
            if len(bgpasnumberArray) > 0 {
                fromMap["bgp-as-number"] = bgpasnumberArray
            }
        }
        if v, ok := d.GetOk("from.0.bgp_as_path"); ok {
            bgpaspathList := v.([]interface{})
            bgpaspathArray := make([]interface{}, 0, len(bgpaspathList))
            for i := range bgpaspathList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.aspath_regex", i)); ok {
                    itemMap["aspath-regex"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.origin", i)); ok {
                    itemMap["origin"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.all_ipv4_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv4_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv4_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv4_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv4_routesMap) > 0 {
                            itemMap["all-ipv4-routes"] = all_ipv4_routesMap
                        }
                    }
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.bgp_as_path.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.range", i, j)); ok {
                            if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                                rawDict := ivList[0].(map[string]interface{})
                                rangeMap := make(map[string]interface{})
                                if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                    rangeMap["from"] = sv.(int)
                                }
                                if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                    rangeMap["to"] = sv.(int)
                                }
                                if len(rangeMap) > 0 {
                                    innerMap["range"] = rangeMap
                                }
                            }
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    bgpaspathArray = append(bgpaspathArray, itemMap)
                }
            }
            if len(bgpaspathArray) > 0 {
                fromMap["bgp-as-path"] = bgpaspathArray
            }
        }
        if v, ok := d.GetOk("from.0.interface"); ok {
            interfaceList := v.([]interface{})
            interfaceArray := make([]interface{}, 0, len(interfaceList))
            for i := range interfaceList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.interface.%d.interface", i)); ok {
                    itemMap["interface"] = v.(string)
                }
                if v, ok := d.GetOk(fmt.Sprintf("from.0.interface.%d.metric", i)); ok {
                    itemMap["metric"] = v.(string)
                }
                if len(itemMap) > 0 {
                    interfaceArray = append(interfaceArray, itemMap)
                }
            }
            if len(interfaceArray) > 0 {
                fromMap["interface"] = interfaceArray
            }
        }
        if v, ok := d.GetOk("from.0.ospf2"); ok {
            ospf2List := v.([]interface{})
            ospf2Array := make([]interface{}, 0, len(ospf2List))
            for i := range ospf2List {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.instance", i)); ok {
                    itemMap["instance"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.all_ipv4_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv4_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv4_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv4_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv4_routesMap) > 0 {
                            itemMap["all-ipv4-routes"] = all_ipv4_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.ospf2.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.ospf2.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2.%d.network.%d.range", i, j)); ok {
                            if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                                rawDict := ivList[0].(map[string]interface{})
                                rangeMap := make(map[string]interface{})
                                if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                    rangeMap["from"] = sv.(int)
                                }
                                if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                    rangeMap["to"] = sv.(int)
                                }
                                if len(rangeMap) > 0 {
                                    innerMap["range"] = rangeMap
                                }
                            }
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    ospf2Array = append(ospf2Array, itemMap)
                }
            }
            if len(ospf2Array) > 0 {
                fromMap["ospf2"] = ospf2Array
            }
        }
        if v, ok := d.GetOk("from.0.ospf2ase"); ok {
            ospf2aseList := v.([]interface{})
            ospf2aseArray := make([]interface{}, 0, len(ospf2aseList))
            for i := range ospf2aseList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.instance", i)); ok {
                    itemMap["instance"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.all_ipv4_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv4_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv4_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv4_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv4_routesMap) > 0 {
                            itemMap["all-ipv4-routes"] = all_ipv4_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.ospf2ase.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf2ase.%d.network.%d.range", i, j)); ok {
                            if ivList, ok := iv.([]interface{}); ok && len(ivList) > 0 {
                                rawDict := ivList[0].(map[string]interface{})
                                rangeMap := make(map[string]interface{})
                                if sv, ok := rawDict["from"]; ok && sv.(int) != 0 {
                                    rangeMap["from"] = sv.(int)
                                }
                                if sv, ok := rawDict["to"]; ok && sv.(int) != 0 {
                                    rangeMap["to"] = sv.(int)
                                }
                                if len(rangeMap) > 0 {
                                    innerMap["range"] = rangeMap
                                }
                            }
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    ospf2aseArray = append(ospf2aseArray, itemMap)
                }
            }
            if len(ospf2aseArray) > 0 {
                fromMap["ospf2ase"] = ospf2aseArray
            }
        }
        if v, ok := d.GetOk("from.0.ospf3"); ok {
            ospf3List := v.([]interface{})
            ospf3Array := make([]interface{}, 0, len(ospf3List))
            for i := range ospf3List {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.instance", i)); ok {
                    itemMap["instance"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.ospf3.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.ospf3.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    ospf3Array = append(ospf3Array, itemMap)
                }
            }
            if len(ospf3Array) > 0 {
                fromMap["ospf3"] = ospf3Array
            }
        }
        if v, ok := d.GetOk("from.0.ospf3ase"); ok {
            ospf3aseList := v.([]interface{})
            ospf3aseArray := make([]interface{}, 0, len(ospf3aseList))
            for i := range ospf3aseList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.instance", i)); ok {
                    itemMap["instance"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok && sv.(bool) {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.ospf3ase.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v := d.Get(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.restrict", i, j)).(bool); v {
                            innerMap["restrict"] = v
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.metric", i, j)); ok {
                            innerMap["metric"] = iv.(string)
                        }
                        if len(innerMap) > 0 {
                            networkArr = append(networkArr, innerMap)
                        }
                    }
                    if len(networkArr) > 0 {
                        itemMap["network"] = networkArr
                    }
                }
                if len(itemMap) > 0 {
                    ospf3aseArray = append(ospf3aseArray, itemMap)
                }
            }
            if len(ospf3aseArray) > 0 {
                fromMap["ospf3ase"] = ospf3aseArray
            }
        }
        if len(fromMap) > 0 {
            payload["from"] = fromMap
        }
    }

    if v, ok := d.GetOk("level"); ok {
        payload["level"] = v.(string)
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    setRouteRedistributionToIsisRes, err := client.ApiCallSimple("set-route-redistribution-to-isis", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setRouteRedistributionToIsisRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setRouteRedistributionToIsisRes.Success {
            errMsg = setRouteRedistributionToIsisRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setRouteRedistributionToIsisRes.GetData()
        }

        debugLogOperation(
            "route-redistribution-to-isis",        // resource type
            "update",                       // operation
            "set-route-redistribution-to-isis",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set route-redistribution-to-isis: %v", err)
    }
    if !setRouteRedistributionToIsisRes.Success {
        return fmt.Errorf(setRouteRedistributionToIsisRes.ErrorMsg)
    }

    return readGaiaRouteRedistributionToIsis(d, m)
}

func deleteGaiaRouteRedistributionToIsis(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    