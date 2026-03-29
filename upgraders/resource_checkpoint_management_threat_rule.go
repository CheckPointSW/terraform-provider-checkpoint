package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementThreatRuleV0 is the V0 schema where position and track_settings were TypeMap.
func ResourceManagementThreatRuleV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"layer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Layer that the rule belongs to identified by the name or UID.",
			},
			"position": {
				Type:        schema.TypeMap,
				Required:    true,
				Description: "Position in the rulebase.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"top": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule on top of specific section identified by uid or name. Select value 'top' for entire rule base.",
						},
						"above": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule above specific section/rule identified by uid or name.",
						},
						"below": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule below specific section/rule identified by uid or name.",
						},
						"bottom": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add rule in the bottom of specific section identified by uid or name. Select value 'bottom' for entire rule base.",
						},
					},
				},
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Rule name.",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Action-the enforced profile.",
				Default:     "Optimized",
			},
			"destination": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"destination_negate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if negate is set for destination.",
				Default:     false,
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable/Disable the rule.",
				Default:     true,
			},
			"install_on": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protected_scope": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of objects defining Protected Scope identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protected_scope_negate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if negate is set for Protected Scope.",
				Default:     false,
			},
			"service": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_negate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if negate is set for service.",
				Default:     false,
			},
			"source": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source_negate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if negate is set for source.",
				Default:     false,
			},
			"track": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Packet tracking.",
				Default:     "Log",
			},
			"track_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Threat rule track settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"packet_capture": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Packet capture.",
						},
					},
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"exceptions": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of rule's exceptions identified by UID",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

// ResourceManagementThreatRuleStateUpgradeV0 converts position and track_settings from TypeMap to TypeList.
func ResourceManagementThreatRuleStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "position", "track_settings"), nil
}
