package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementIpsUpdateSchedule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementIpsUpdateScheduleRead,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "IPS Update Schedule status.",
			},
			"time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Time in format HH:mm.",
			},
			"recurrence": {
				Type:        schema.TypeList,
				Computed:    true,
				MaxItems:    1,
				Description: "Days recurrence.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"days": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Valid on specific days. Multiple options, support range of days in months. Example:[\"1\",\"3\",\"9-20\"].",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"minutes": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Valid on interval. The length of time in minutes between updates.",
						},
						"pattern": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Valid on \"Interval\", \"Daily\", \"Weekly\", \"Monthly\" base.",
						},
						"weekdays": {
							Type:        schema.TypeSet,
							Computed:    true,
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

func dataSourceManagementIpsUpdateScheduleRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showIpsUpdateScheduleRes, err := client.ApiCall("show-ips-update-schedule", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if !showIpsUpdateScheduleRes.Success {
		fmt.Errorf(showIpsUpdateScheduleRes.ErrorMsg)
	}

	ipsUpdateSchedule := showIpsUpdateScheduleRes.GetData()

	log.Println("Read Ips Update Schedule - Show JSON = ", ipsUpdateSchedule)

	d.SetId("show-ips-update-schedule-" + acctest.RandString(10))

	if v := ipsUpdateSchedule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := ipsUpdateSchedule["time"]; v != nil {
		_ = d.Set("time", v)
	}

	if ipsUpdateSchedule["recurrence"] != nil {
		recurrenceMap := ipsUpdateSchedule["recurrence"].(map[string]interface{})
		recurrenceList := make([]interface{}, 0)

		recurrenceMapToReturn := make(map[string]interface{})

		if recurrenceMap["days"] != nil {
			daysJson := recurrenceMap["days"].([]interface{})
			daysIds := make([]string, 0)
			if len(daysJson) > 0 {
				for _, day := range daysJson {
					day := day.(map[string]interface{})
					daysIds = append(daysIds, day["name"].(string))
				}
			}
			recurrenceMapToReturn["days"] = daysIds
		} else {
			recurrenceMapToReturn["days"] = nil
		}

		if v, _ := recurrenceMap["minutes"]; v != nil {
			recurrenceMapToReturn["minutes"] = v
		}
		if v, _ := recurrenceMap["pattern"]; v != nil {
			recurrenceMapToReturn["pattern"] = v
		}

		if recurrenceMap["weekdays"] != nil {
			weekdaysJson := recurrenceMap["weekdays"].([]interface{})
			weekdaysIds := make([]string, 0)
			if len(weekdaysJson) > 0 {
				for _, weekday := range weekdaysJson {
					weekday := weekday.(map[string]interface{})
					weekdaysIds = append(weekdaysIds, weekday["name"].(string))
				}
			}
			recurrenceMapToReturn["weekdays"] = weekdaysIds
		} else {
			recurrenceMapToReturn["weekdays"] = nil
		}

		recurrenceList = append(recurrenceList, recurrenceMapToReturn)
		_ = d.Set("recurrence", recurrenceList)
	} else {
		_ = d.Set("recurrence", nil)
	}

	return nil
}
