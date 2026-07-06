package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementScheduledEvent() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementScheduledEventRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"schedule": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Schedule Configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time in format HH:mm.",
						},
						"recurrence": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Days recurrence.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pattern": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Days recurrence pattern.",
									},
									"interval_hours": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The amount of hours between updates. <font color=\"red\">Required only when</font> pattern is set to 'Interval'.",
									},
									"interval_minutes": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The amount of minutes between updates. <font color=\"red\">Required only when</font> pattern is set to 'Interval'.",
									},
									"interval_seconds": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The amount of seconds between updates. <font color=\"red\">Required only when</font> pattern is set to 'Interval'.",
									},
									"weekdays": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Days of the week to run the update.<br> Valid values: group of values from {'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'}. <font color=\"red\">Required only when</font> pattern is set to 'Weekly'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"days": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Days of the month to run the update.<br> Valid values: interval in the range of 1 to 31. <font color=\"red\">Required only when</font> pattern is set to 'Monthly'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object icon.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementScheduledEventRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	showScheduledEventRes, err := client.ApiCall("show-scheduled-event", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showScheduledEventRes.Success {
		if objectNotFound(showScheduledEventRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showScheduledEventRes.ErrorMsg)
	}

	scheduledEvent := showScheduledEventRes.GetData()

	log.Println("Read ScheduledEvent - Show JSON = ", scheduledEvent)

	if v := scheduledEvent["uid"]; v != nil {
		d.SetId(v.(string))
	}

	if v := scheduledEvent["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if scheduledEvent["schedule"] != nil {

		scheduleMap := scheduledEvent["schedule"].(map[string]interface{})

		scheduleMapToReturn := make(map[string]interface{})

		if v := scheduleMap["time"]; v != nil {
			scheduleMapToReturn["time"] = v
		}
		if v := scheduleMap["recurrence"]; v != nil {

			recurrenceMap := v.(map[string]interface{})

			recurrenceMapToReturn := make(map[string]interface{})

			if v := recurrenceMap["pattern"]; v != nil {
				recurrenceMapToReturn["pattern"] = v
			}
			if v := recurrenceMap["interval-hours"]; v != nil {
				recurrenceMapToReturn["interval_hours"] = v
			}
			if v := recurrenceMap["interval-minutes"]; v != nil {
				recurrenceMapToReturn["interval_minutes"] = v
			}
			if v := recurrenceMap["interval-seconds"]; v != nil {
				recurrenceMapToReturn["interval_seconds"] = v
			}
			if v := recurrenceMap["weekdays"]; v != nil {
				recurrenceMapToReturn["weekdays"] = v
			}
			if v := recurrenceMap["days"]; v != nil {
				recurrenceMapToReturn["days"] = v
			}

			scheduleMapToReturn["recurrence"] = []interface{}{recurrenceMapToReturn}
		}

		_ = d.Set("schedule", []interface{}{scheduleMapToReturn})

	} else {
		_ = d.Set("schedule", nil)
	}

	if v := scheduledEvent["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := scheduledEvent["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := scheduledEvent["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	if scheduledEvent["tags"] != nil {
		tagsJson, ok := scheduledEvent["tags"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("tags", tagsIds)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if v := scheduledEvent["uid"]; v != nil {
		_ = d.Set("uid", v)
	}

	return nil

}
