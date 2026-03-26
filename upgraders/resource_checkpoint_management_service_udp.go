package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementServiceUdpV0 is the V0 schema where aggressive_aging was TypeMap.
func ResourceManagementServiceUdpV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"accept_replies": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "N/A",
				Default:     true,
			},
			"aggressive_aging": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Sets short (aggressive) timeouts for idle connections.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default_timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Default aggressive aging timeout in seconds.",
						},
						"enable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "N/A",
						},
						"timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Aggressive aging timeout in seconds.",
						},
						"use_default_timeout": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "N/A",
						},
					},
				},
			},
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
			},
			"match_by_protocol_signature": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "A value of true enables matching by the selected protocol's signature - the signature identifies the protocol as genuine. Select this option to limit the port to the specified protocol. If the selected protocol does not support matching by signature, this field cannot be set to true.",
				Default:     false,
			},
			"match_for_any": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.",
				Default:     true,
			},
			"override_default_settings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether this service is a Data Domain service which has been overridden.",
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The number of the port used to provide this service. To specify a port range, place a hyphen between the lowest and highest port numbers, for example 44-55.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Select the protocol type associated with the service, and by implication, the management server (if any) that enforces Content Security and Authentication for the service. Selecting a Protocol Type invokes the specific protocol handlers for each protocol type, thus enabling higher level of security by parsing the protocol, and higher level of connectivity by tracking dynamic actions (such as opening of ports).",
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Time (in seconds) before the session times out.",
				Default:     40,
			},
			"source_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Port number for the client side service. If specified, only those Source port Numbers will be Accepted, Dropped, or Rejected during packet inspection. Otherwise, the source port is not inspected.",
			},
			"sync_connections_on_cluster": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.",
				Default:     true,
			},
			"use_default_session_timeout": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use default virtual session timeout.",
				Default:     true,
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
		},
	}
}

// ResourceManagementServiceUdpStateUpgradeV0 converts aggressive_aging from TypeMap to TypeList.
func ResourceManagementServiceUdpStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "aggressive_aging"), nil
}
