package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementVsxProvisioningTool() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVsxProvisioningTool,
		Read:   readManagementVsxProvisioningTool,
		Delete: deleteManagementVsxProvisioningTool,
		Schema: map[string]*schema.Schema{
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
			"operation": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the provisioning operation to run. Each operation has its own specific parameters.<br> The available operations are:<ul><li><i>add-vsx-gateway</i> - Adds a new VSX gateway</li><li><i>add-vsx-cluster</i> - Adds a new VSX cluster*</li><li><i>add-vsx-cluster-member</i> - Adds a new VSX cluster member*</li><li><i>add-vd</i> - Adds a new Virtual Device (VS/VSB/VSW/VR) to a VSX gateway or VSX cluster</li><li><i>add-vd-interface</i> - Adds a new virtual interface to a Virtual Device</li><li><i>add-physical-interface</i> - Adds a physical interface to a VSX gateway or VSX cluster</li><li><i>add-route</i> - Adds a route to a Virtual Device</li><li><i>attach-bridge</i> - Attaches a bridge interface to a Virtual System</li><li><i>remove-vsx</i> - Removes a VSX gateway or VSX cluster</li><li><i>remove-vd</i> - Removes a Virtual Device</li><li><i>remove-vd-interface</i> - Removes an interface from a Virtual Device</li><li><i>remove-physical-interface</i> - Removes a physical interface from a VSX gateway or VSX cluster</li><li><i>remove-route</i> - Removes a route from a Virtual Device</li><li><i>set-vd</i> - Modifies a Virtual Device</li><li><i>set-vd-interface</i> - Modifies an interface on a Virtual Device</li><li><i>set-physical-interface</i> - Modifies a physical interface on a VSX cluster or VSX gateway</li></ul><br> * When adding a VSX Cluster, you must also add at least 2 cluster members<br> * Adding cluster members is only allowed when adding a new VSX cluster<br> * To add members to an existing cluster, use vsx-run-operation.",
			},
			"add_physical_interface_params": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Parameters for the operation to add a physical interface to a VSX gateway or VSX Cluster.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the interface.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the VSX Gateway or Cluster object.",
						},
						"vlan_trunk": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "True if this interface is a VLAN trunk.",
							Default:     false,
						},
					},
				},
			},
			"add_route_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to add a route to a Virtual System or Virtual Router.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Route destination. To specify the default route, use 'default' for IPv4 and 'default6' for IPv6.",
						},
						"next_hop": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Next hop IP address.",
						},
						"leads_to": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Virtual Router for this route<br/>This VD must have an existing connection to the VR.",
						},
						"vd": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the Virtual System, Virtual Switch, or Virtual Router.",
						},
						"netmask": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Subnet mask for this route.",
						},
						"prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "CIDR prefix for this route.",
						},
						"propagate": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Propagate this route to adjacent virtual devices.",
							Default:     false,
						},
					},
				},
			},
			"add_vd_interface_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to add a new interface to a Virtual Device.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"leads_to": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Virtual Switch or Virtual Router for this interface.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the interface.",
						},
						"vd": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the Virtual System, Virtual Switch, or Virtual Router.",
						},
						"anti_spoofing": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The anti-spoofing enforcement setting of this interface.",
						},
						"anti_spoofing_tracking": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The anti-spoofing tracking setting of this interface.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "IPv4 Address of this interface with optional CIDR prefix.<br/>Required if this interface belongs to a Virtual System or Virtual Router.",
						},
						"ipv4_netmask": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "IPv4 Subnet mask of this interface.",
						},
						"ipv4_prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "IPv4 CIDR prefix of this interface.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "IPv6 Address of this interface<br/>Required if this interface belongs to a Virtual System or Virtual Router.",
						},
						"ipv6_netmask": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "IPv6 Subnet mask of this interface.",
						},
						"ipv6_prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "IPv6 CIDR prefix of this interface.",
						},
						"mtu": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "MTU of this interface.",
						},
						"propagate": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Propagate IPv4 route to adjacent virtual devices.",
							Default:     false,
						},
						"propagate6": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Propagate IPv6 route to adjacent virtual devices.",
							Default:     false,
						},
						"specific_group": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Specific group for interface topology.<br/>Only for use with topology option 'internal_specific'.",
						},
						"topology": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Topology of this interface.<br/>Automatic topology calculation based on routes must be disabled for this VS.",
						},
						"vti_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							MaxItems:    1,
							Description: "VTI settings for this interface. This Virtual System must have VPN blade enabled.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"local_ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The IPv4 address of the VPN tunnel on this Virtual System.",
									},
									"peer_name": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The name of the remote peer object as defined in the VPN community.",
									},
									"remote_ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "The IPv4 address of the VPN tunnel on the remote VPN peer.",
									},
									"tunnel_id": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Optional unique Tunnel ID.<br/>Automatically assigned by the system if empty.",
									},
								},
							},
						},
					},
				},
			},
			"add_vd_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to add a new Virtual Device (VS/VSB/VSW/VR).",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interfaces": {
							Type:        schema.TypeList,
							Required:    true,
							ForceNew:    true,
							Description: "The list of interfaces for this new Virtual Device.<br/>Optional if this new VD is a Virtual Switch.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"leads_to": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Virtual Switch or Virtual Router for this interface.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of the interface.",
									},
									"anti_spoofing": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The anti-spoofing enforcement setting of this interface.",
									},
									"anti_spoofing_tracking": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The anti-spoofing tracking setting of this interface.",
									},
									"ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv4 Address of this interface with optional CIDR prefix.<br/>Required if this interface belongs to a Virtual System or Virtual Router.",
									},
									"ipv4_netmask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv4 Subnet mask of this interface.",
									},
									"ipv4_prefix": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv4 CIDR prefix of this interface.",
									},
									"ipv6_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv6 Address of this interface<br/>Required if this interface belongs to a Virtual System or Virtual Router.",
									},
									"ipv6_netmask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv6 Subnet mask of this interface.",
									},
									"ipv6_prefix": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv6 CIDR prefix of this interface.",
									},
									"mtu": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "MTU of this interface.",
									},
									"propagate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Propagate IPv4 route to adjacent virtual devices.",
									},
									"propagate6": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Propagate IPv6 route to adjacent virtual devices.",
									},
									"specific_group": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specific group for interface topology.<br/>Only for use with topology option 'internal_specific'.",
									},
									"topology": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Topology of this interface.<br/>Automatic topology calculation based on routes must be disabled for this VS.",
									},
								},
							},
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type of the Virtual Device <br><br>vs - Virtual Firewall<br>vr - Virtual Router<br>vsw - Virtual Switch<br>vsbm - Virtual Firewall in bridge mode.",
						},
						"vd": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the Virtual System, Virtual Switch, or Virtual Router.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the VSX Gateway or Cluster object.",
						},
						"calc_topology_auto": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Calculate interface topology automatically based on routes.<br/>Relevant only for Virtual Systems.<br/>Do not use for virtual devices.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Main IPv4 Address.<br/>Required if this device is a Virtual System.<br/>Do not use for other virtual devices.",
						},
						"ipv4_instances": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "Number of IPv4 instances for the Virtual System.<br/>Must be greater or equal to 1.<br/>Only relevant for Virtual Systems and Virtual Systems in bridge mode.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Main IPv6 Address.<br/>Required if this device is a Virtual System.<br/>Do not use for other virtual devices.",
						},
						"ipv6_instances": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "Number of IPv6 instances for the Virtual System.<br/>Only relevant for Virtual Systems and Virtual Systems in bridge mode.",
						},
						"routes": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: "The list of routes for this new Virtual Device (VS or VR only).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Route destination. To specify the default route, use 'default' for IPv4 and 'default6' for IPv6.",
									},
									"next_hop": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Next hop IP address.",
									},
									"leads_to": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Virtual Router for this route<br/>This VD must have an existing connection to the VR.",
									},
									"netmask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Subnet mask for this route.",
									},
									"prefix": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "CIDR prefix for this route.",
									},
									"propagate": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Propagate this route to adjacent virtual devices.",
										Default:     false,
									},
								},
							},
						},
						"vs_mtu": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "MTU of the Virtual System.<br/>Only relevant for Virtual Systems in bridge mode.<br/>Do not use for other virtual devices.",
						},
					},
				},
			},
			"add_vsx_cluster_params": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Parameters for the operation to add a new VSX Cluster.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cluster_type": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Cluster type for the VSX Cluster Object.<br/>Starting in R81.10, only VSLS can be configured during cluster creation.<br/>To use High Availability ('ha'), first create the cluster as VSLS and then run vsx_util on the Management.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Main IPv4 Address of the VSX Gateway or Cluster object.<br/>Optional if main IPv6 Address is defined.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Main IPv6 Address of the VSX Gateway or Cluster object.<br/>Optional if main IPv4 Address is defined.",
						},
						"members": {
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Description: "The list of cluster members for this new VSX Cluster. Minimum: 2.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Main IPv4 Address of the VSX Cluster member.<br/>Mandatory if the VSX Cluster has an IPv4 Address.",
									},
									"ipv6_address": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Main IPv6 Address of the VSX Cluster member.<br/>Mandatory if the VSX Cluster has an IPv6 Address.",
									},
									"name": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Name of the new VSX Cluster member.",
									},
									"sic_otp": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "SIC one-time-password of the VSX Gateway or Cluster member.<br/>Password must be between 4-127 characters in length.",
									},
									"sync_ip": {
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Description: "Sync IP address for the VSX Cluster member.",
									},
								},
							},
						},
						"sync_if_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Sync interface name for the VSX Cluster.",
						},
						"sync_netmask": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Sync interface netmask for the VSX Cluster.",
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Version of the VSX Gateway or Cluster object.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the VSX Gateway or Cluster object.",
						},
						"rule_drop": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a default drop rule to the VSX Gateway or Cluster initial policy.",
							Default:     "enable",
						},
						"rule_https": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow HTTPS traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
						"rule_ping": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow ping traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
						"rule_ping6": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow ping6 traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
						"rule_snmp": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow SNMP traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
						"rule_ssh": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow SSH traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
					},
				},
			},
			"add_vsx_gateway_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to add a new VSX Gateway.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Main IPv4 Address of the VSX Gateway or Cluster object.<br/>Optional if main IPv6 Address is defined.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Main IPv6 Address of the VSX Gateway or Cluster object.<br/>Optional if main IPv4 Address is defined.",
						},
						"sic_otp": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "SIC one-time-password of the VSX Gateway or Cluster member.<br/>Password must be between 4-127 characters in length.",
						},
						"version": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Version of the VSX Gateway or Cluster object.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the VSX Gateway or Cluster object.",
						},
						"rule_drop": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a default drop rule to the VSX Gateway or Cluster initial policy.",
							Default:     "enable",
						},
						"rule_https": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow HTTPS traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
						"rule_ping": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow ping traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
						"rule_ping6": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow ping6 traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
						"rule_snmp": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow SNMP traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
						"rule_ssh": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Add a rule to allow SSH traffic to the VSX Gateway or Cluster initial policy.",
							Default:     "disable",
						},
					},
				},
			},
			"attach_bridge_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to attach a new bridge interface to a Virtual System.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ifs1": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the first interface for the bridge.",
						},
						"ifs2": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the second interface for the bridge.",
						},
						"vd": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the Virtual System, Virtual Switch, or Virtual Router.",
						},
					},
				},
			},
			"remove_physical_interface_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to remove a physical interface from a VSX (Gateway or Cluster).",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the interface.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the VSX Gateway or Cluster object.",
						},
					},
				},
			},
			"remove_route_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to remove a route from a Virtual System or Virtual Router.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"destination": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Route destination. To specify the default route, use 'default' for IPv4 and 'default6' for IPv6.",
						},
						"vd": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the Virtual System, Virtual Switch, or Virtual Router.",
						},
						"netmask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Subnet mask for this route.",
						},
						"prefix": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "CIDR prefix for this route.",
						},
					},
				},
			},
			"remove_vd_interface_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to remove a logical interface from a Virtual Device.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"leads_to": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Virtual Switch or Virtual Router for this interface.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the interface.",
						},
						"vd": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the Virtual System, Virtual Switch, or Virtual Router.",
						},
					},
				},
			},
			"remove_vd_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to remove a Virtual Device.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vd": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the Virtual System, Virtual Switch, or Virtual Router.",
						},
					},
				},
			},
			"remove_vsx_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to remove a VSX Gateway or VSX Cluster.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the VSX Gateway or Cluster object.",
						},
					},
				},
			},
			"set_physical_interface_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to change the configuration of a physical interface.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the interface.",
						},
						"vlan_trunk": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "True if this interface is a VLAN trunk.",
						},
						"vsx_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the VSX Gateway or Cluster object.",
						},
					},
				},
			},
			"set_vd_interface_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to change the configuration of a logical interface.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"leads_to": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Virtual Switch or Virtual Router for this interface.",
						},
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the interface.",
						},
						"vd": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Name of the Virtual System, Virtual Switch, or Virtual Router.",
						},
						"anti_spoofing": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The anti-spoofing enforcement setting of this interface.",
						},
						"anti_spoofing_tracking": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The anti-spoofing tracking setting of this interface.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "IPv4 Address of this interface with optional CIDR prefix.<br/>Required if this interface belongs to a Virtual System or Virtual Router.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "IPv6 Address of this interface<br/>Required if this interface belongs to a Virtual System or Virtual Router.",
						},
						"mtu": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "MTU of this interface.",
						},
						"new_leads_to": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "New Virtual Switch or Virtual Router for this interface.",
						},
						"propagate": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Propagate IPv4 route to adjacent virtual devices.",
							Default:     false,
						},
						"propagate6": {
							Type:        schema.TypeBool,
							Optional:    true,
							ForceNew:    true,
							Description: "Propagate IPv6 route to adjacent virtual devices.",
							Default:     false,
						},
						"specific_group": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Specific group for interface topology.<br/>Only for use with topology option 'internal_specific'.",
						},
						"topology": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Topology of this interface.<br/>Automatic topology calculation based on routes must be disabled for this VS.",
						},
					},
				},
			},
			"set_vd_params": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Parameters for the operation to change the configuration of a Virtual Device.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vd": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name of the Virtual System, Virtual Switch, or Virtual Router.",
						},
						"calc_topology_auto": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Calculate interface topology automatically based on routes.<br/>Relevant only for Virtual Systems.<br/>Do not use for virtual devices.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Main IPv4 Address.<br/>Relevant only if this device is a Virtual System.<br/>Do not use for other virtual devices.",
						},
						"ipv4_instances": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of IPv4 instances for the Virtual System.<br/>Must be greater or equal to 1.<br/>Only relevant for Virtual Systems and Virtual Systems in bridge mode.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Main IPv6 Address.<br/>Relevant only if this device is a Virtual System.<br/>Do not use for other virtual devices.",
						},
						"ipv6_instances": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Number of IPv6 instances for the Virtual System.<br/>Only relevant for Virtual Systems and Virtual Systems in bridge mode.",
						},
						"vs_mtu": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "MTU of the Virtual System.<br/>Only relevant for Virtual Systems in bridge mode.<br/>Do not use for other virtual devices.",
						},
					},
				},
			},
		},
	}
}

func createManagementVsxProvisioningTool(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("operation"); ok {
		payload["operation"] = v.(string)
	}

	if _, ok := d.GetOk("add_physical_interface_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("add_physical_interface_params.0.name"); ok {
			res["name"] = v.(string)
		}
		if v, ok := d.GetOk("add_physical_interface_params.0.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		if v, ok := d.GetOk("add_physical_interface_params.0.vlan_trunk"); ok {
			res["vlan-trunk"] = v
		}
		payload["add-physical-interface-params"] = res
	}

	if _, ok := d.GetOk("add_route_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("add_route_params.0.destination"); ok {
			res["destination"] = v.(string)
		}
		if v, ok := d.GetOk("add_route_params.0.next_hop"); ok {
			res["next-hop"] = v.(string)
		}
		if v, ok := d.GetOk("add_route_params.0.leads_to"); ok {
			res["leads-to"] = v.(string)
		}
		if v, ok := d.GetOk("add_route_params.0.vd"); ok {
			res["vd"] = v.(string)
		}
		if v, ok := d.GetOk("add_route_params.0.netmask"); ok {
			res["netmask"] = v.(string)
		}
		if v, ok := d.GetOk("add_route_params.0.prefix"); ok {
			res["prefix"] = v.(string)
		}
		if v, ok := d.GetOk("add_route_params.0.propagate"); ok {
			res["propagate"] = v
		}
		payload["add-route-params"] = res
	}

	if _, ok := d.GetOk("add_vd_interface_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("add_vd_interface_params.0.leads_to"); ok {
			res["leads-to"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.name"); ok {
			res["name"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.vd"); ok {
			res["vd"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.anti_spoofing"); ok {
			res["anti-spoofing"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.anti_spoofing_tracking"); ok {
			res["anti-spoofing-tracking"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.ipv4_address"); ok {
			res["ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.ipv4_netmask"); ok {
			res["ipv4-netmask"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.ipv4_prefix"); ok {
			res["ipv4-prefix"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.ipv6_address"); ok {
			res["ipv6-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.ipv6_netmask"); ok {
			res["ipv6-netmask"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.ipv6_prefix"); ok {
			res["ipv6-prefix"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.mtu"); ok {
			res["mtu"] = v
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.propagate"); ok {
			res["propagate"] = v
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.propagate6"); ok {
			res["propagate6"] = v
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.specific_group"); ok {
			res["specific-group"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.topology"); ok {
			res["topology"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_interface_params.0.vti_settings"); ok {

			innerMap := v.([]interface{})[0].(map[string]interface{})

			mapToReturn := make(map[string]interface{})

			if v := innerMap["local_ipv4_address"]; v != nil {
				mapToReturn["local-ipv4-address"] = v
			}
			if v := innerMap["peer_name"]; v != nil {
				mapToReturn["peer-name"] = v
			}
			if v := innerMap["remote_ipv4_address"]; v != nil {
				mapToReturn["remote-ipv4-address"] = v
			}
			if v := innerMap["tunnel_id"]; v != nil {
				mapToReturn["tunnel-id"] = v
			}
			res["vti-settings"] = mapToReturn
		}
		payload["add-vd-interface-params"] = res
	}

	if _, ok := d.GetOk("add_vd_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("add_vd_params.0.interfaces"); ok {

			interfacesList := v.([]interface{})

			var interfacesPayload []map[string]interface{}

			for i := range interfacesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".name"); ok {
					Payload["name"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".leads_to"); ok {
					Payload["leads-to"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".anti_spoofing"); ok {
					Payload["anti-spoofing"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".anti_spoofing_tracking"); ok {
					Payload["anti-spoofing-tracking"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".ipv4_address"); ok {
					Payload["ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".ipv4_netmask"); ok {
					Payload["ipv4-netmask"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".ipv4_prefix"); ok {
					Payload["ipv4-prefix"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".ipv6_address"); ok {
					Payload["ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".ipv6_netmask"); ok {
					Payload["ipv6-netmask"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".ipv6_prefix"); ok {
					Payload["ipv6-prefix"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".mtu"); ok {
					Payload["mtu"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".propagate"); ok {
					Payload["propagate"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".propagate6"); ok {
					Payload["propagate6"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".specific_group"); ok {
					Payload["specific-group"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.interfaces." + strconv.Itoa(i) + ".topology"); ok {
					Payload["topology"] = v.(string)
				}
				interfacesPayload = append(interfacesPayload, Payload)
			}

			res["interfaces"] = interfacesPayload
		}
		if v, ok := d.GetOk("add_vd_params.0.type"); ok {
			res["type"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_params.0.vd"); ok {
			res["vd"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_params.0.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_params.0.calc_topology_auto"); ok {
			res["calc-topology-auto"] = v
		}
		if v, ok := d.GetOk("add_vd_params.0.ipv4_address"); ok {
			res["ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_params.0.ipv4_instances"); ok {
			res["ipv4-instances"] = v
		}
		if v, ok := d.GetOk("add_vd_params.0.ipv6_address"); ok {
			res["ipv6-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_vd_params.0.ipv6_instances"); ok {
			res["ipv6-instances"] = v
		}
		if v, ok := d.GetOk("add_vd_params.0.routes"); ok {

			routesList := v.([]interface{})

			var routesPayload []map[string]interface{}

			for i := range routesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("add_vd_params.0.routes." + strconv.Itoa(i) + ".destination"); ok {
					Payload["destination"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.routes." + strconv.Itoa(i) + ".next_hop"); ok {
					Payload["next-hop"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.routes." + strconv.Itoa(i) + ".leads_to"); ok {
					Payload["leads-to"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.routes." + strconv.Itoa(i) + ".netmask"); ok {
					Payload["netmask"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.routes." + strconv.Itoa(i) + ".prefix"); ok {
					Payload["prefix"] = v.(string)
				}
				if v, ok := d.GetOk("add_vd_params.0.routes." + strconv.Itoa(i) + ".propagate"); ok {
					Payload["propagate"] = v.(string)
				}
				routesPayload = append(routesPayload, Payload)
			}
			res["routes"] = routesPayload
		}
		if v, ok := d.GetOk("add_vd_params.0.vs_mtu"); ok {
			res["vs-mtu"] = v
		}
		payload["add-vd-params"] = res
	}

	if _, ok := d.GetOk("add_vsx_cluster_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("add_vsx_cluster_params.0.cluster_type"); ok {
			res["cluster-type"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.ipv4_address"); ok {
			res["ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.ipv6_address"); ok {
			res["ipv6-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.members"); ok {

			membersList := v.([]interface{})

			if len(membersList) > 0 {

				var membersObjectsPayload []map[string]interface{}

				for i := range membersList {

					memberObject := membersList[i].(map[string]interface{})

					objectPayload := make(map[string]interface{})

					if v := memberObject["ipv4_address"]; v != nil {
						objectPayload["ipv4-address"] = v
					}
					if v := memberObject["ipv6_address"]; v != nil {
						objectPayload["ipv6-address"] = v
					}
					if v := memberObject["name"]; v != nil {
						objectPayload["name"] = v
					}
					if v := memberObject["sic_otp"]; v != nil {
						objectPayload["sic-otp"] = v
					}
					if v := memberObject["sync_ip"]; v != nil {
						objectPayload["sync-ip"] = v
					}
					membersObjectsPayload = append(membersObjectsPayload, objectPayload)
				}
				res["members"] = membersObjectsPayload
			}

		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.sync_if_name"); ok {
			res["sync-if-name"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.sync_netmask"); ok {
			res["sync-netmask"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.version"); ok {
			res["version"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.rule_drop"); ok {
			res["rule-drop"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.rule_https"); ok {
			res["rule-https"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.rule_ping"); ok {
			res["rule-ping"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.rule_ping6"); ok {
			res["rule-ping6"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.rule_snmp"); ok {
			res["rule-snmp"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_cluster_params.0.rule_ssh"); ok {
			res["rule-ssh"] = v.(string)
		}
		payload["add-vsx-cluster-params"] = res
	}

	if _, ok := d.GetOk("add_vsx_gateway_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("add_vsx_gateway_params.0.ipv4_address"); ok {
			res["ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.ipv6_address"); ok {
			res["ipv6-address"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.sic_otp"); ok {
			res["sic-otp"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.version"); ok {
			res["version"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.rule_drop"); ok {
			res["rule-drop"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.rule_https"); ok {
			res["rule-https"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.rule_ping"); ok {
			res["rule-ping"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.rule_ping6"); ok {
			res["rule-ping6"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.rule_snmp"); ok {
			res["rule-snmp"] = v.(string)
		}
		if v, ok := d.GetOk("add_vsx_gateway_params.0.rule_ssh"); ok {
			res["rule-ssh"] = v.(string)
		}
		payload["add-vsx-gateway-params"] = res
	}

	if _, ok := d.GetOk("attach_bridge_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("attach_bridge_params.0.ifs1"); ok {
			res["ifs1"] = v.(string)
		}
		if v, ok := d.GetOk("attach_bridge_params.0.ifs2"); ok {
			res["ifs2"] = v.(string)
		}
		if v, ok := d.GetOk("attach_bridge_params.0.vd"); ok {
			res["vd"] = v.(string)
		}
		payload["attach-bridge-params"] = res
	}

	if _, ok := d.GetOk("remove_physical_interface_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("remove_physical_interface_params.0.name"); ok {
			res["name"] = v.(string)
		}
		if v, ok := d.GetOk("remove_physical_interface_params.0.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		payload["remove-physical-interface-params"] = res
	}

	if _, ok := d.GetOk("remove_route_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("remove_route_params.0.destination"); ok {
			res["destination"] = v.(string)
		}
		if v, ok := d.GetOk("remove_route_params.0.vd"); ok {
			res["vd"] = v.(string)
		}
		if v, ok := d.GetOk("remove_route_params.0.netmask"); ok {
			res["netmask"] = v.(string)
		}
		if v, ok := d.GetOk("remove_route_params.0.prefix"); ok {
			res["prefix"] = v.(string)
		}
		payload["remove-route-params"] = res
	}

	if _, ok := d.GetOk("remove_vd_interface_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("remove_vd_interface_params.0.leads_to"); ok {
			res["leads-to"] = v.(string)
		}
		if v, ok := d.GetOk("remove_vd_interface_params.0.name"); ok {
			res["name"] = v.(string)
		}
		if v, ok := d.GetOk("remove_vd_interface_params.0.vd"); ok {
			res["vd"] = v.(string)
		}
		payload["remove-vd-interface-params"] = res
	}

	if _, ok := d.GetOk("remove_vd_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("remove_vd_params.0.vd"); ok {
			res["vd"] = v.(string)
		}
		payload["remove-vd-params"] = res
	}

	if _, ok := d.GetOk("remove_vsx_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("remove_vsx_params.0.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		payload["remove-vsx-params"] = res
	}

	if _, ok := d.GetOk("set_physical_interface_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("set_physical_interface_params.0.name"); ok {
			res["name"] = v.(string)
		}
		if v, ok := d.GetOk("set_physical_interface_params.0.vlan_trunk"); ok {
			res["vlan-trunk"] = v
		}
		if v, ok := d.GetOk("set_physical_interface_params.0.vsx_name"); ok {
			res["vsx-name"] = v.(string)
		}
		payload["set-physical-interface-params"] = res
	}

	if _, ok := d.GetOk("set_vd_interface_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("set_vd_interface_params.0.leads_to"); ok {
			res["leads-to"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.name"); ok {
			res["name"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.vd"); ok {
			res["vd"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.anti_spoofing"); ok {
			res["anti-spoofing"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.anti_spoofing_tracking"); ok {
			res["anti-spoofing-tracking"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.ipv4_address"); ok {
			res["ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.ipv6_address"); ok {
			res["ipv6-address"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.mtu"); ok {
			res["mtu"] = v
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.new_leads_to"); ok {
			res["new-leads-to"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.propagate"); ok {
			res["propagate"] = v
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.propagate6"); ok {
			res["propagate6"] = v
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.specific_group"); ok {
			res["specific-group"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_interface_params.0.topology"); ok {
			res["topology"] = v.(string)
		}
		payload["set-vd-interface-params"] = res
	}

	if _, ok := d.GetOk("set_vd_params"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("set_vd_params.0.vd"); ok {
			res["vd"] = v
		}
		if v, ok := d.GetOk("set_vd_params.0.calc_topology_auto"); ok {
			res["calc-topology-auto"] = v
		}
		if v, ok := d.GetOk("set_vd_params.0.ipv4_address"); ok {
			res["ipv4-address"] = v
		}
		if v, ok := d.GetOk("set_vd_params.0.ipv4_instances"); ok {
			res["ipv4-instances"] = v
		}
		if v, ok := d.GetOk("set_vd_params.0.ipv6_address"); ok {
			res["ipv6-address"] = v.(string)
		}
		if v, ok := d.GetOk("set_vd_params.0.ipv6_instances"); ok {
			res["ipv6-instances"] = v
		}
		if v, ok := d.GetOk("set_vd_params.0.vs_mtu"); ok {
			res["vs-mtu"] = v
		}
		payload["set-vd-params"] = res
	}

	vsxProvisioningToolRes, err := client.ApiCall("vsx-provisioning-tool", payload, client.GetSessionID(), true, client.IsProxyUsed())

	log.Println("vsx-provisioning-tool result is ", vsxProvisioningToolRes)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	if !vsxProvisioningToolRes.Success {
		return fmt.Errorf(vsxProvisioningToolRes.ErrorMsg)
	}

	d.SetId("vsx-provisioning-tool-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(vsxProvisioningToolRes.GetData()))
	return readManagementVsxProvisioningTool(d, m)
}

func readManagementVsxProvisioningTool(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementVsxProvisioningTool(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
