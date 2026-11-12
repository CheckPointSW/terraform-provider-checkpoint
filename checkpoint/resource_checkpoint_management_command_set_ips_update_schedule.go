package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/v3/upgraders"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementSetIpsUpdateSchedule() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetIpsUpdateSchedule,
		Read:   readManagementSetIpsUpdateSchedule,
		Delete: deleteManagementSetIpsUpdateSchedule,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementCommandSetIpsUpdateScheduleV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementCommandSetIpsUpdateScheduleStateUpgradeV0,
				Version: 0,
			},
		},
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
				Type:        schema.TypeList,
				MaxItems:    1,
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

func createManagementSetIpsUpdateSchedule(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("enabled"); ok {
		payload["enabled"] = v.(bool)
	}

	if v, ok := d.GetOk("time"); ok {
		payload["time"] = v.(string)
	}

	if v, ok := d.GetOk("recurrence"); ok {

		recurrenceList := v.([]interface{})

		if len(recurrenceList) > 0 {

			recurrencePayload := make(map[string]interface{})

			if v, ok := d.GetOk("recurrence.0.days"); ok {
				recurrencePayload["days"] = v
			}
			if v, ok := d.GetOk("recurrence.0.minutes"); ok {
				recurrencePayload["minutes"] = v.(int)
			}
			if v, ok := d.GetOk("recurrence.0.pattern"); ok {
				recurrencePayload["pattern"] = v.(string)
			}
			if v, ok := d.GetOk("recurrence.0.weekdays"); ok {
				recurrencePayload["weekdays"] = v
			}
			payload["recurrence"] = recurrencePayload
		}
	}

	SetIpsUpdateScheduleRes, _ := client.ApiCall("set-ips-update-schedule", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !SetIpsUpdateScheduleRes.Success {
		return fmt.Errorf("%s", SetIpsUpdateScheduleRes.ErrorMsg)
	}

	d.SetId("set-ips-update-schedule-" + acctest.RandString(10))
	return readManagementSetIpsUpdateSchedule(d, m)
}

func readManagementSetIpsUpdateSchedule(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetIpsUpdateSchedule(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
