package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementServiceOtherV0 is the V0 schema where aggressive_aging was TypeMap.
func ResourceManagementServiceOtherV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"accept_replies": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Specifies whether Other Service replies are to be accepted.",
				Default:     false,
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Contains an INSPECT expression that defines the action to take if a rule containing this service is matched. Example: set r_mhandler &open_ssl_handler sets a handler on the connection.",
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
							Default:     600,
						},
						"enable": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "N/A",
							Default:     true,
						},
						"timeout": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Aggressive aging timeout in seconds.",
							Default:     600,
						},
						"use_default_timeout": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "N/A",
							Default:     true,
						},
					},
				},
			},
			"ip_protocol": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "IP protocol number.",
			},
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
			},
			"match": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Contains an INSPECT expression that defines the matching criteria. The connection is examined against the expression during the first packet. Example: tcp, dport = 21, direction = 0 matches incoming FTP control connections.",
			},
			"match_for_any": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.",
			},
			"override_default_settings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether this service is a Data Domain service which has been overridden.",
				Default:     false,
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Time (in seconds) before the session times out.",
			},
			"sync_connections_on_cluster": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables state-synchronized High Availability or Load Sharing on a ClusterXL or OPSEC-certified cluster.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_default_session_timeout": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Use default virtual session timeout.",
				Default:     true,
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

// ResourceManagementServiceOtherStateUpgradeV0 converts aggressive_aging from TypeMap to TypeList.
func ResourceManagementServiceOtherStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "aggressive_aging"), nil
}
