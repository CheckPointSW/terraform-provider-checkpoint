package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementGlobalAssignmentV0 is the V0 schema where assignment_up_to_date was TypeMap.
func ResourceManagementGlobalAssignmentV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"dependent_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "N/A",
			},
			"global_access_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Global domain access policy that is assigned to a dependent domain.",
			},
			"global_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "N/A",
			},
			"global_threat_prevention_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Global domain threat prevention policy that is assigned to a dependent domain.",
			},
			"manage_protection_actions": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "N/A",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
			"assignment_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"assignment_up_to_date": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "The time when the assignment was assigned.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
					},
				},
			},
		},
	}
}

// ResourceManagementGlobalAssignmentStateUpgradeV0 converts assignment_up_to_date from TypeMap to TypeList.
func ResourceManagementGlobalAssignmentStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "assignment_up_to_date"), nil
}
