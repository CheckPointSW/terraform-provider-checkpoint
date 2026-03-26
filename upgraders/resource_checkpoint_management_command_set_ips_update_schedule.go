package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandSetIpsUpdateScheduleV0 is the V0 schema where recurrence was TypeMap.
func ResourceManagementCommandSetIpsUpdateScheduleV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Enable/Disable IPS Update Schedule.",
			},
			"time": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Time in format HH:mm.",
			},
			"recurrence": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Days recurrence.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"days": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Valid on specific days. Multiple options, support range of days in months. Example:[\"1\",\"3\",\"9-20\"].",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"minutes": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Valid on interval. The length of time in minutes between updates.",
						},
						"pattern": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Valid on \"Interval\", \"Daily\", \"Weekly\", \"Monthly\" base.",
						},
						"weekdays": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Valid on weekdays. Example: \"Sun\", \"Mon\"...\"Sat\".",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

// ResourceManagementCommandSetIpsUpdateScheduleStateUpgradeV0 converts recurrence from TypeMap to TypeList.
func ResourceManagementCommandSetIpsUpdateScheduleStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "recurrence"), nil
}
