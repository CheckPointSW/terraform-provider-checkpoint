package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaRouteRedistributionToOspf3() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaRouteRedistributionToOspf3,
        Read:   readGaiaRouteRedistributionToOspf3,
        Update: updateGaiaRouteRedistributionToOspf3,
        Delete: deleteGaiaRouteRedistributionToOspf3,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "from": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Configure policy for exporting routes to IPv6 OSPF`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "aggregate": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Specifies the aggregate route to redistribute into IPv6 OSPF`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches all IPv6 aggregate routes`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                        Description: `Matches specific IPv6 aggregate routes. The aggregate routes have to be already configured.`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 network`,
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
                            Description: `Redistribution of kernel routes into IPv6 OSPF.<br><br>Note: It may be inadvisable in certain cases to redistribute kernel routes into another protocol. Kernel routes usually exist upon startup of routed, before the routing table has settled, when error conditions or bad routes may be present. Use caution when configuring route redistribution from the kernel.`,
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
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                                    Description: `Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule`,
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
                            Description: `Redistribution of NAT pools into IPv6 OSPF`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches all IPv4 NAT pools`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                        Description: `Matches specific IPv6 NAT pools. The NAT pool has to be already configured.`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 network`,
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
                        "ripng": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Redistribution of RIPng routes into IPv6 OSPF`,
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
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                                    Description: `Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule`,
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
                            Description: `Redistribution of static routes into IPv6 OSPF`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "all_ipv6_routes": {
                                        Type:        schema.TypeList,
                                        Optional:    true,
                                        Description: `Matches all IPv4 static route`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                        Description: `Matches the default IPv4 static route`,
                                        MaxItems:    1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "metric": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                        Description: `Matches specific IPv6 static routes. The static route has to be already configured.`,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "address": {
                                                    Type:        schema.TypeString,
                                                    Optional:    true,
                                                    Description: `Specifies IPv6 network`,
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
                            Description: `Configures Autonomous System numbers of the BGP group from which to export routes into IPv6 OSPF`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "as_number": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Configured Autonomous System Number. Valid Values are 1 - 4294967295 or 0.1 - 65535.65535.<br><br>The ASN format can be changed to dotted or plain format using the following command 'set format asn dotted/plain'.`,
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
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                                    Description: `Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule`,
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
                            Description: `Configures the redistribution of BGP routes, whose AS path matches a given regular expression into IPv6 OSPF`,
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
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                                    Description: `Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule`,
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
                            Description: `Configures the redistribution of all directly connected routes from an interface into IPv6 OSPF`,
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
                                        Description: `Specifies the OSPF metric to be added to routes redistributed via this rule<br><br>The metric used by OSPF is a cost, representing the overhead required (i.e. due to bandwidth) to reach a destination. Routes with higher OSPF cost are more expensive`,
                                    },
                                },
                            },
                        },
                        "isis": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Configures the redistribution of IS-IS routes into IPv6 OSPF`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "level": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Specifies which IS-IS level the route redistribution is applied to`,
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
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                                    Description: `Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule`,
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
                            Description: `Configures the redistribution of IPv6 OSPF routes into IPv6 OSPF`,
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
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                                    Description: `Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule`,
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
                            Description: `Configures the redistribution of IPv6 OSPF Autonomous System External routes into IPv6 OSPF`,
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
                                                    Description: `Specifies IPv6 OSPF metric value to routes matching this rule`,
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
                                                    Description: `Specifies the IPv6 OSPF metric to be added to routes redistributed via this rule`,
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
            "instance": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Configures OSPF3 for specified instance`,
            },
            "reset": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Removes OSPF3 Route Redistribution configuration`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaRouteRedistributionToOspf3(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("from"); len(v.([]interface{})) > 0 {
        _ = v
        fromMap := make(map[string]interface{})
        if v, ok := d.GetOk("from.0.aggregate"); ok {
            _ = v
            aggregateMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.aggregate.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.aggregate.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.aggregate.0.all_ipv6_routes.0.enable"); ok {
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
            if v, ok := d.GetOk("from.0.kernel.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.kernel.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.kernel.0.all_ipv6_routes.0.enable"); ok {
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
                    if v, ok := d.GetOkExists(fmt.Sprintf("from.0.kernel.0.network.%d.restrict", i)); ok {
                        itemMap["restrict"] = v.(bool)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.match_type", i)); ok {
                        itemMap["match-type"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
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
            if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.nat_pool.0.all_ipv6_routes.0.enable"); ok {
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
        if v, ok := d.GetOk("from.0.ripng"); ok {
            _ = v
            ripngMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.ripng.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.ripng.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.ripng.0.all_ipv6_routes.0.enable"); ok {
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
                    if v, ok := d.GetOkExists(fmt.Sprintf("from.0.ripng.0.network.%d.restrict", i)); ok {
                        itemMap["restrict"] = v.(bool)
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
            if v, ok := d.GetOk("from.0.static_route.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.all_ipv6_routes.0.enable"); ok {
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
                if v, ok := d.GetOkExists("from.0.static_route.0.default6.0.enable"); ok {
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
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok {
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
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.metric", i, j)); ok {
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
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok {
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
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.metric", i, j)); ok {
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
        if v, ok := d.GetOk("from.0.isis"); ok {
            isisList := v.([]interface{})
            isisArray := make([]interface{}, 0, len(isisList))
            for i := range isisList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.isis.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.isis.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.network.%d.metric", i, j)); ok {
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
                    isisArray = append(isisArray, itemMap)
                }
            }
            if len(isisArray) > 0 {
                fromMap["isis"] = isisArray
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
                        if sv, ok := rawDict["enable"]; ok {
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
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.ospf3.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
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
                        if sv, ok := rawDict["enable"]; ok {
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
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
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

    if v, ok := d.GetOk("instance"); ok {
        payload["instance"] = v.(string)
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    log.Println("Create RouteRedistributionToOspf3 - Map = ", payload)

    addRouteRedistributionToOspf3Res, err := client.ApiCallSimple("set-route-redistribution-to-ospf3", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addRouteRedistributionToOspf3Res.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addRouteRedistributionToOspf3Res.Success {
            errMsg = addRouteRedistributionToOspf3Res.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addRouteRedistributionToOspf3Res.GetData()
        }

        debugLogOperation(
            "route-redistribution-to-ospf3",        // resource type
            "create",                       // operation
            "set-route-redistribution-to-ospf3",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add route-redistribution-to-ospf3: %v", err)
    }
    if !addRouteRedistributionToOspf3Res.Success {
        if addRouteRedistributionToOspf3Res.ErrorMsg != "" {
            return fmt.Errorf(addRouteRedistributionToOspf3Res.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("route-redistribution-to-ospf3-" + acctest.RandString(10)))
    return readGaiaRouteRedistributionToOspf3(d, m)
}

func readGaiaRouteRedistributionToOspf3(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("instance"); ok {
        payload["instance"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showRouteRedistributionToOspf3Res, err := client.ApiCallSimple("show-route-redistribution-to-ospf3", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showRouteRedistributionToOspf3Res.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showRouteRedistributionToOspf3Res.Success {
            errMsg = showRouteRedistributionToOspf3Res.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showRouteRedistributionToOspf3Res.GetData()
        }

        debugLogOperation(
            "route-redistribution-to-ospf3",        // resource type
            "read",                       // operation
            "show-route-redistribution-to-ospf3",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show route-redistribution-to-ospf3: %v", err)
    }
    if !showRouteRedistributionToOspf3Res.Success {
        if data := showRouteRedistributionToOspf3Res.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showRouteRedistributionToOspf3Res.ErrorMsg)
    }

    routeRedistributionToOspf3 := showRouteRedistributionToOspf3Res.GetData()

    log.Println("Read RouteRedistributionToOspf3 - Show JSON = ", routeRedistributionToOspf3)

    if v, exists := routeRedistributionToOspf3["ospf3"]; exists {
        // Bug 1: API returns redistribution data under "ospf3"; remap to Terraform "from" field.
        if ospf3Items, ok := v.([]interface{}); ok && len(ospf3Items) > 0 {
            apiItem, _ := ospf3Items[0].(map[string]interface{})
            if val, ok := apiItem["instance"]; ok {
                d.Set("instance", fmt.Sprintf("%v", val))
            }
            apiFrom, _ := apiItem["from"].(map[string]interface{})
            if apiFrom == nil {
                apiFrom = map[string]interface{}{}
            }
            fromObj := map[string]interface{}{}
            buildAIR := func(m map[string]interface{}) map[string]interface{} {
                r := map[string]interface{}{}
                if val, ok := m["metric"]; ok {
                    r["metric"] = fmt.Sprintf("%v", val)
                }
                if val, ok := m["enable"]; ok {
                    r["enable"] = val.(bool)
                }
                return r
            }
            buildNet := func(lst []interface{}, hasRestrict bool) []interface{} {
                out := make([]interface{}, 0, len(lst))
                for _, rw := range lst {
                    if nm, ok := rw.(map[string]interface{}); ok {
                        ni := map[string]interface{}{}
                        if val, ok := nm["address"]; ok {
                            ni["address"] = fmt.Sprintf("%v", val)
                        }
                        if hasRestrict {
                            if val, ok := nm["restrict"]; ok {
                                ni["restrict"] = val.(bool)
                            }
                            if val, ok := nm["match-type"]; ok {
                                ni["match_type"] = fmt.Sprintf("%v", val)
                            }
                        }
                        if val, ok := nm["metric"]; ok {
                            ni["metric"] = fmt.Sprintf("%v", val)
                        }
                        if len(ni) > 0 {
                            out = append(out, ni)
                        }
                    }
                }
                return out
            }
            for _, e := range []struct {
                apiKey string
                tfKey  string
                hasR   bool
                hasD6  bool
            }{
                {"aggregate",    "aggregate",    false, false},
                {"kernel",       "kernel",       true,  false},
                {"nat-pool",     "nat_pool",     false, false},
                {"ripng",        "ripng",        true,  false},
                {"static-route", "static_route", false, true},
            } {
                if raw, ok := apiFrom[e.apiKey]; ok {
                    if pm, ok := raw.(map[string]interface{}); ok {
                        protoObj := map[string]interface{}{}
                        if ar, ok := pm["all-ipv6-routes"]; ok {
                            if am, ok := ar.(map[string]interface{}); ok {
                                protoObj["all_ipv6_routes"] = []interface{}{buildAIR(am)}
                            }
                        }
                        if e.hasD6 {
                            if dr, ok := pm["default6"]; ok {
                                if dm, ok := dr.(map[string]interface{}); ok {
                                    protoObj["default6"] = []interface{}{buildAIR(dm)}
                                }
                            }
                        }
                        if nr, ok := pm["network"]; ok {
                            if nl, ok := nr.([]interface{}); ok {
                                protoObj["network"] = buildNet(nl, e.hasR)
                            }
                        }
                        fromObj[e.tfKey] = []interface{}{protoObj}
                    }
                }
            }
            // bgp-as-number
            if raw, ok := apiFrom["bgp-as-number"]; ok {
                if items, ok := raw.([]interface{}); ok {
                    out := make([]interface{}, 0, len(items))
                    for _, ri := range items {
                        if im, ok := ri.(map[string]interface{}); ok {
                            item := map[string]interface{}{}
                            if val, ok := im["as-number"]; ok {
                                item["as_number"] = fmt.Sprintf("%v", val)
                            }
                            if ar, ok := im["all-ipv6-routes"]; ok {
                                if am, ok := ar.(map[string]interface{}); ok {
                                    item["all_ipv6_routes"] = []interface{}{buildAIR(am)}
                                }
                            }
                            if nr, ok := im["network"]; ok {
                                if nl, ok := nr.([]interface{}); ok {
                                    item["network"] = buildNet(nl, true)
                                }
                            }
                            if len(item) > 0 {
                                out = append(out, item)
                            }
                        }
                    }
                    fromObj["bgp_as_number"] = out
                }
            }
            // bgp-as-path
            if raw, ok := apiFrom["bgp-as-path"]; ok {
                if items, ok := raw.([]interface{}); ok {
                    out := make([]interface{}, 0, len(items))
                    for _, ri := range items {
                        if im, ok := ri.(map[string]interface{}); ok {
                            item := map[string]interface{}{}
                            if val, ok := im["aspath-regex"]; ok {
                                item["aspath_regex"] = fmt.Sprintf("%v", val)
                            }
                            if val, ok := im["origin"]; ok {
                                item["origin"] = fmt.Sprintf("%v", val)
                            }
                            if ar, ok := im["all-ipv6-routes"]; ok {
                                if am, ok := ar.(map[string]interface{}); ok {
                                    item["all_ipv6_routes"] = []interface{}{buildAIR(am)}
                                }
                            }
                            if nr, ok := im["network"]; ok {
                                if nl, ok := nr.([]interface{}); ok {
                                    item["network"] = buildNet(nl, true)
                                }
                            }
                            if len(item) > 0 {
                                out = append(out, item)
                            }
                        }
                    }
                    fromObj["bgp_as_path"] = out
                }
            }
            // interface
            if raw, ok := apiFrom["interface"]; ok {
                if items, ok := raw.([]interface{}); ok {
                    out := make([]interface{}, 0, len(items))
                    for _, ri := range items {
                        if im, ok := ri.(map[string]interface{}); ok {
                            item := map[string]interface{}{}
                            if val, ok := im["interface"]; ok {
                                item["interface"] = fmt.Sprintf("%v", val)
                            }
                            if val, ok := im["metric"]; ok {
                                item["metric"] = fmt.Sprintf("%v", val)
                            }
                            if len(item) > 0 {
                                out = append(out, item)
                            }
                        }
                    }
                    fromObj["interface"] = out
                }
            }
            // isis, ospf3, ospf3ase
            for _, e2 := range []struct{ a, t string }{
                {"isis",     "isis"},
                {"ospf3",    "ospf3"},
                {"ospf3ase", "ospf3ase"},
            } {
                if raw, ok := apiFrom[e2.a]; ok {
                    if items, ok := raw.([]interface{}); ok {
                        out := make([]interface{}, 0, len(items))
                        for _, ri := range items {
                            if im, ok := ri.(map[string]interface{}); ok {
                                item := map[string]interface{}{}
                                if e2.a == "isis" {
                                    if val, ok := im["level"]; ok {
                                        item["level"] = fmt.Sprintf("%v", val)
                                    }
                                } else {
                                    if val, ok := im["instance"]; ok {
                                        item["instance"] = fmt.Sprintf("%v", val)
                                    }
                                }
                                if ar, ok := im["all-ipv6-routes"]; ok {
                                    if am, ok := ar.(map[string]interface{}); ok {
                                        item["all_ipv6_routes"] = []interface{}{buildAIR(am)}
                                    }
                                }
                                if nr, ok := im["network"]; ok {
                                    if nl, ok := nr.([]interface{}); ok {
                                        item["network"] = buildNet(nl, true)
                                    }
                                }
                                if len(item) > 0 {
                                    out = append(out, item)
                                }
                            }
                        }
                        fromObj[e2.t] = out
                    }
                }
            }
            d.Set("from", []interface{}{fromObj})
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaRouteRedistributionToOspf3(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("from"); len(v.([]interface{})) > 0 {
        _ = v
        fromMap := make(map[string]interface{})
        if v, ok := d.GetOk("from.0.aggregate"); ok {
            _ = v
            aggregateMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.aggregate.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.aggregate.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.aggregate.0.all_ipv6_routes.0.enable"); ok {
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
            if v, ok := d.GetOk("from.0.kernel.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.kernel.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.kernel.0.all_ipv6_routes.0.enable"); ok {
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
                    if v, ok := d.GetOkExists(fmt.Sprintf("from.0.kernel.0.network.%d.restrict", i)); ok {
                        itemMap["restrict"] = v.(bool)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.match_type", i)); ok {
                        itemMap["match-type"] = v.(string)
                    }
                    if v, ok := d.GetOk(fmt.Sprintf("from.0.kernel.0.network.%d.metric", i)); ok {
                        itemMap["metric"] = v.(string)
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
            if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.nat_pool.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.nat_pool.0.all_ipv6_routes.0.enable"); ok {
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
        if v, ok := d.GetOk("from.0.ripng"); ok {
            _ = v
            ripngMap := make(map[string]interface{})
            if v, ok := d.GetOk("from.0.ripng.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.ripng.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.ripng.0.all_ipv6_routes.0.enable"); ok {
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
                    if v, ok := d.GetOkExists(fmt.Sprintf("from.0.ripng.0.network.%d.restrict", i)); ok {
                        itemMap["restrict"] = v.(bool)
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
            if v, ok := d.GetOk("from.0.static_route.0.all_ipv6_routes"); ok {
                _ = v
                allipv6routesMap := make(map[string]interface{})
                if v, ok := d.GetOk("from.0.static_route.0.all_ipv6_routes.0.metric"); ok {
                    allipv6routesMap["metric"] = v.(string)
                }
                if v, ok := d.GetOkExists("from.0.static_route.0.all_ipv6_routes.0.enable"); ok {
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
                if v, ok := d.GetOkExists("from.0.static_route.0.default6.0.enable"); ok {
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
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok {
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
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_number.%d.network.%d.metric", i, j)); ok {
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
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok {
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
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.bgp_as_path.%d.network.%d.metric", i, j)); ok {
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
        if v, ok := d.GetOk("from.0.isis"); ok {
            isisList := v.([]interface{})
            isisArray := make([]interface{}, 0, len(isisList))
            for i := range isisList {
                itemMap := make(map[string]interface{})
                if v, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.level", i)); ok {
                    itemMap["level"] = v.(string)
                }
                if sv, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.all_ipv6_routes", i)); ok {
                    if ivList, ok := sv.([]interface{}); ok && len(ivList) > 0 {
                        rawDict := ivList[0].(map[string]interface{})
                        all_ipv6_routesMap := make(map[string]interface{})
                        if sv, ok := rawDict["metric"]; ok && sv.(string) != "" {
                            all_ipv6_routesMap["metric"] = sv.(string)
                        }
                        if sv, ok := rawDict["enable"]; ok {
                            all_ipv6_routesMap["enable"] = sv.(bool)
                        }
                        if len(all_ipv6_routesMap) > 0 {
                            itemMap["all-ipv6-routes"] = all_ipv6_routesMap
                        }
                    }
                }
                if sv := d.Get(fmt.Sprintf("from.0.isis.%d.network", i)); len(sv.([]interface{})) > 0 {
                    networkList := sv.([]interface{})
                    networkArr := make([]interface{}, 0, len(networkList))
                    for j := range networkList {
                        innerMap := make(map[string]interface{})
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.network.%d.address", i, j)); ok {
                            innerMap["address"] = iv.(string)
                        }
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.isis.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.network.%d.match_type", i, j)); ok {
                            innerMap["match-type"] = iv.(string)
                        }
                        if iv, ok := d.GetOk(fmt.Sprintf("from.0.isis.%d.network.%d.metric", i, j)); ok {
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
                    isisArray = append(isisArray, itemMap)
                }
            }
            if len(isisArray) > 0 {
                fromMap["isis"] = isisArray
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
                        if sv, ok := rawDict["enable"]; ok {
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
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.ospf3.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
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
                        if sv, ok := rawDict["enable"]; ok {
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
                        if v, ok := d.GetOkExists(fmt.Sprintf("from.0.ospf3ase.%d.network.%d.restrict", i, j)); ok {
                            innerMap["restrict"] = v.(bool)
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

    if v, ok := d.GetOk("instance"); ok {
        payload["instance"] = v.(string)
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    setRouteRedistributionToOspf3Res, err := client.ApiCallSimple("set-route-redistribution-to-ospf3", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setRouteRedistributionToOspf3Res.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setRouteRedistributionToOspf3Res.Success {
            errMsg = setRouteRedistributionToOspf3Res.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setRouteRedistributionToOspf3Res.GetData()
        }

        debugLogOperation(
            "route-redistribution-to-ospf3",        // resource type
            "update",                       // operation
            "set-route-redistribution-to-ospf3",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set route-redistribution-to-ospf3: %v", err)
    }
    if !setRouteRedistributionToOspf3Res.Success {
        return fmt.Errorf(setRouteRedistributionToOspf3Res.ErrorMsg)
    }

    return readGaiaRouteRedistributionToOspf3(d, m)
}

func deleteGaiaRouteRedistributionToOspf3(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    