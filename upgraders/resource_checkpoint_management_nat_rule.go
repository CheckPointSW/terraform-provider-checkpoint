package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementNatRuleV0 is the V0 schema where position was TypeMap.
func ResourceManagementNatRuleV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"package": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the package.",
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
			"method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Nat method.",
				Default:     "static",
			},
			"original_destination": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Original destination.",
				Default:     "Any",
			},
			"original_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Original service.",
				Default:     "Any",
			},
			"original_source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Original source.",
				Default:     "Any",
			},
			"translated_destination": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Translated destination.",
				Default:     "Original",
			},
			"translated_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Translated service.",
				Default:     "Original",
			},
			"translated_source": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Translated source.",
				Default:     "Original",
			},
			"auto_generated": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Auto generated.",
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
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
		},
	}
}

// ResourceManagementNatRuleStateUpgradeV0 converts position from TypeMap to TypeList.
func ResourceManagementNatRuleStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "position"), nil
}
