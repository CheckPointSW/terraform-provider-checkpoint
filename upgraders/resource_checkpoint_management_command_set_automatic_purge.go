package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCommandSetAutomaticPurgeV0 is the V0 schema where scheduling was TypeMap.
func ResourceManagementCommandSetAutomaticPurgeV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Turn on/off the automatic-purge feature.",
			},
			"keep_sessions_by_count": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether or not to keep the latest N sessions. Note: when the automatic purge feature is enabled, this field and/or the \"keep-sessions-by-date\" field must be set to 'true'.",
			},
			"number_of_sessions_to_keep": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "When \"keep-sessions-by-count = true\" this sets the number of newest sessions to preserve, by the sessions's publish date.",
			},
			"keep_sessions_by_days": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether or not to keep the sessions for D days. Note: when the automatic purge feature is enabled, this field and/or the \"keep-sessions-by-count\" field must be set to 'true'.",
			},
			"number_of_days_to_keep": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "When \"keep-sessions-by-days = true\" this sets the number of days to keep the sessions.",
			},
			"scheduling": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "When to purge sessions that do not meet the \"keep\" criteria. Note: when the automatic purge feature is enabled, this field must be set.",
				ForceNew:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_date": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "The first time to check whether or not there are sessions to purge. ISO 8601. If timezone isn't specified in the input, the Management server's timezone is used. Instead - If you want to start immediately, type: \"now\". Note: when the automatic purge feature is enabled, this field must be set.",
						},
						"time_units": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: "Note: when the automatic purge feature is enabled, this field must be set.",
						},
						"check_interval": {
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Description: "Number of time-units between two purge checks.  Note: when the automatic purge feature is enabled, this field must be set.",
						},
					},
				},
			},
		},
	}
}

// ResourceManagementCommandSetAutomaticPurgeStateUpgradeV0 converts scheduling from TypeMap to TypeList.
func ResourceManagementCommandSetAutomaticPurgeStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "scheduling"), nil
}
