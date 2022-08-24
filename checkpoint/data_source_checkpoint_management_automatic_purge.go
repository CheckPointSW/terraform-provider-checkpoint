package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementAutomaticPurge() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAutomaticPurgeRead,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Turn on/off the automatic-purge feature.",
			},
			"keep_sessions_by_count": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether or not to keep the latest N sessions.",
			},
			"number_of_sessions_to_keep": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of newest sessions to preserve, by the sessions's publish date.",
			},
			"keep_sessions_by_days": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether or not to keep the sessions for D days.",
			},
			"number_of_days_to_keep": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "When \"keep-sessions-by-days = true\" this sets the number of days to keep the sessions.",
			},
			"scheduling": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "When to purge sessions that do not meet the \"keep\" criteria. Note: when the automatic purge feature is enabled, this field must be set.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The first time to check whether or not there are sessions to purge. ISO 8601. If timezone isn't specified in the input, the Management server's timezone is used. Instead - If you want to start immediately, type: \"now\". Note: when the automatic purge feature is enabled, this field must be set.",
						},
						"time_units": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Note: when the automatic purge feature is enabled, this field must be set.",
						},
						"check_interval": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of time-units between two purge checks.  Note: when the automatic purge feature is enabled, this field must be set.",
						},
						"last_check": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last time purge check was executed.",
						},
						"next_check": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Next time purge check will be executed.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementAutomaticPurgeRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showAutomaticPurgeRes, err := client.ApiCall("show-automatic-purge", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if !showAutomaticPurgeRes.Success {
		fmt.Errorf(showAutomaticPurgeRes.ErrorMsg)
	}

	automaticPurge := showAutomaticPurgeRes.GetData()

	log.Println("Read Automatic Purge - Show JSON = ", automaticPurge)

	d.SetId("show-automatic-purge-" + acctest.RandString(10))

	if v := automaticPurge["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := automaticPurge["keep-sessions-by-count"]; v != nil {
		_ = d.Set("keep_sessions_by_count", v)
	}

	if v := automaticPurge["number-of-sessions-to-keep"]; v != nil {
		_ = d.Set("number_of_sessions_to_keep", v)
	}

	if v := automaticPurge["keep-sessions-by-days"]; v != nil {
		_ = d.Set("keep_sessions_by_days", v)
	}

	if v := automaticPurge["number-of-days-to-keep"]; v != nil {
		_ = d.Set("number_of_days_to_keep", v)
	}

	if automaticPurge["scheduling"] != nil {
		schedulingMap := automaticPurge["scheduling"].(map[string]interface{})

		schedulingMapToReturn := make(map[string]interface{})

		if v, _ := schedulingMap["start-date"]; v != nil {
			schedulingMapToReturn["start_date"] = v
		}
		if v, _ := schedulingMap["time-units"]; v != nil {
			schedulingMapToReturn["time_units"] = v
		}
		if v, _ := schedulingMap["check-interval"]; v != nil {
			schedulingMapToReturn["check_interval"] = v
		}
		if v, _ := schedulingMap["last-check"]; v != nil {
			schedulingMapToReturn["last_check"] = v
		}
		if v, _ := schedulingMap["next-check"]; v != nil {
			schedulingMapToReturn["next_check"] = v
		}

		_ = d.Set("scheduling", []interface{}{schedulingMapToReturn})
	} else {
		_ = d.Set("scheduling", nil)
	}

	return nil
}
