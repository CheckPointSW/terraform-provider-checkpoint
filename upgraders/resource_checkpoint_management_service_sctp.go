package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementServiceSctpV0 is the V0 schema where aggressive_aging was TypeMap.
func ResourceManagementServiceSctpV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
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
							Default:     0,
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
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
			},
			"match_for_any": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether this service is used when 'Any' is set as the rule's service and there are several service objects with the same source port and protocol.",
			},
			"port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45.",
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Time (in seconds) before the session times out.",
			},
			"source_port": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Source port number. To specify a port range add a hyphen between the lowest and the highest port numbers, for example 44-45.",
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

// ResourceManagementServiceSctpStateUpgradeV0 converts aggressive_aging from TypeMap to TypeList.
func ResourceManagementServiceSctpStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "aggressive_aging"), nil
}
