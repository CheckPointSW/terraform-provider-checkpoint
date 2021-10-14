package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetAutomaticPurge() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetAutomaticPurge,
		Read:   readManagementSetAutomaticPurge,
		Delete: deleteManagementSetAutomaticPurge,
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

func createManagementSetAutomaticPurge(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("enabled"); ok {
		payload["enabled"] = v.(bool)
	}

	if v, ok := d.GetOkExists("keep_sessions_by_count"); ok {
		payload["keep-sessions-by-count"] = v.(bool)
	}

	if v, ok := d.GetOk("number_of_sessions_to_keep"); ok {
		payload["number-of-sessions-to-keep"] = v.(int)
	}

	if v, ok := d.GetOkExists("keep_sessions_by_days"); ok {
		payload["keep-sessions-by-days"] = v.(bool)
	}

	if v, ok := d.GetOk("number_of_days_to_keep"); ok {
		payload["number-of-days-to-keep"] = v.(int)
	}

	if _, ok := d.GetOk("scheduling"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("scheduling.start_date"); ok {
			res["start-date"] = v.(string)
		}
		if v, ok := d.GetOk("scheduling.time_units"); ok {
			res["time-units"] = v.(string)
		}
		if v, ok := d.GetOk("scheduling.check_interval"); ok {
			res["check-interval"] = v.(int)
		}
		payload["scheduling"] = res
	}

	SetAutomaticPurgeRes, _ := client.ApiCall("set-automatic-purge", payload, client.GetSessionID(), true, false)
	if !SetAutomaticPurgeRes.Success {
		return fmt.Errorf(SetAutomaticPurgeRes.ErrorMsg)
	}

	d.SetId("set-automatic-purge" + acctest.RandString(10))
	return readManagementSetAutomaticPurge(d, m)
}

func readManagementSetAutomaticPurge(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetAutomaticPurge(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
