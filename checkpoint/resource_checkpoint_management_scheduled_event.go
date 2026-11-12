package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceManagementScheduledEvent() *schema.Resource {
	return &schema.Resource{
		Create: createManagementScheduledEvent,
		Read:   readManagementScheduledEvent,
		Update: updateManagementScheduledEvent,
		Delete: deleteManagementScheduledEvent,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"schedule": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Schedule Configuration.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time in format HH:mm.",
						},
						"recurrence": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Days recurrence.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pattern": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Days recurrence pattern.",
										Default:     "Interval",
									},
									"interval_hours": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The amount of hours between updates. <font color=\"red\">Required only when</font> pattern is set to 'Interval'.",
										Default:     0,
									},
									"interval_minutes": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The amount of minutes between updates. <font color=\"red\">Required only when</font> pattern is set to 'Interval'.",
										Default:     0,
									},
									"interval_seconds": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "The amount of seconds between updates. <font color=\"red\">Required only when</font> pattern is set to 'Interval'.",
										Default:     0,
									},
									"weekdays": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Days of the week to run the update.<br> Valid values: group of values from {'Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'}. <font color=\"red\">Required only when</font> pattern is set to 'Weekly'.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"days": {
										Type:        schema.TypeSet,
										Optional:    true,
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
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
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

func createManagementScheduledEvent(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	scheduledEvent := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		scheduledEvent["name"] = v.(string)
	}

	if v, ok := d.GetOk("schedule"); ok {

		scheduleList := v.([]interface{})

		if len(scheduleList) > 0 {

			schedulePayload := make(map[string]interface{})

			if v, ok := d.GetOk("schedule.0.time"); ok {
				schedulePayload["time"] = v.(string)
			}
			if _, ok := d.GetOk("schedule.0.recurrence"); ok {

				recurrencePayload := make(map[string]interface{})

				if v, ok := d.GetOk("schedule.0.recurrence.0.pattern"); ok {
					recurrencePayload["pattern"] = v.(string)
				}
				if v, ok := d.GetOk("schedule.0.recurrence.0.interval_hours"); ok {
					recurrencePayload["interval-hours"] = v
				}
				if v, ok := d.GetOk("schedule.0.recurrence.0.interval_minutes"); ok {
					recurrencePayload["interval-minutes"] = v
				}
				if v, ok := d.GetOk("schedule.0.recurrence.0.interval_seconds"); ok {
					recurrencePayload["interval-seconds"] = v
				}
				if v, ok := d.GetOk("schedule.0.recurrence.0.weekdays"); ok {
					recurrencePayload["weekdays"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("schedule.0.recurrence.0.days"); ok {
					recurrencePayload["days"] = v.(*schema.Set).List()
				}
				schedulePayload["recurrence"] = recurrencePayload
			}
			scheduledEvent["schedule"] = schedulePayload
		}
	}
	if v, ok := d.GetOk("color"); ok {
		scheduledEvent["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		scheduledEvent["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		scheduledEvent["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		scheduledEvent["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		scheduledEvent["ignore-errors"] = v.(bool)
	}

	log.Println("Create ScheduledEvent - Map = ", scheduledEvent)

	addScheduledEventRes, err := client.ApiCall("add-scheduled-event", scheduledEvent, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addScheduledEventRes.Success {
		if addScheduledEventRes.ErrorMsg != "" {
			return fmt.Errorf(addScheduledEventRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addScheduledEventRes.GetData()["uid"].(string))

	return readManagementScheduledEvent(d, m)
}

func readManagementScheduledEvent(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showScheduledEventRes, err := client.ApiCall("show-scheduled-event", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

			if v, _ := recurrenceMap["pattern"]; v != nil {
				recurrenceMapToReturn["pattern"] = v
			}
			if v, _ := recurrenceMap["interval-hours"]; v != nil {
				recurrenceMapToReturn["interval_hours"] = v
			}
			if v, _ := recurrenceMap["interval-minutes"]; v != nil {
				recurrenceMapToReturn["interval_minutes"] = v
			}
			if v, _ := recurrenceMap["interval-seconds"]; v != nil {
				recurrenceMapToReturn["interval_seconds"] = v
			}
			if v, _ := recurrenceMap["weekdays"]; v != nil {
				recurrenceMapToReturn["weekdays"] = v
			}
			if v, _ := recurrenceMap["days"]; v != nil {
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

	if v := scheduledEvent["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := scheduledEvent["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementScheduledEvent(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	scheduledEvent := make(map[string]interface{})

	scheduledEvent["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			scheduledEvent["new-name"] = v.(string)
		}
	}

	if d.HasChange("schedule") {

		if v, ok := d.GetOk("schedule"); ok {

			scheduleList := v.([]interface{})

			if len(scheduleList) > 0 {

				schedulePayload := make(map[string]interface{})

				if v, ok := d.GetOk("schedule.0.time"); ok {
					schedulePayload["time"] = v.(string)
				}
				if _, ok := d.GetOk("schedule.0.recurrence"); ok {

					recurrencePayload := make(map[string]interface{})

					if v, ok := d.GetOk("schedule.0.recurrence.0.pattern"); ok {
						recurrencePayload["pattern"] = v.(string)
					}
					if v, ok := d.GetOk("schedule.0.recurrence.0.interval_hours"); ok {
						recurrencePayload["interval-hours"] = v
					}
					if v, ok := d.GetOk("schedule.0.recurrence.0.interval_minutes"); ok {
						recurrencePayload["interval-minutes"] = v
					}
					if v, ok := d.GetOk("schedule.0.recurrence.0.interval_seconds"); ok {
						recurrencePayload["interval-seconds"] = v
					}
					if v, ok := d.GetOk("schedule.0.recurrence.0.weekdays"); ok {
						recurrencePayload["weekdays"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("schedule.0.recurrence.0.days"); ok {
						recurrencePayload["days"] = v.(*schema.Set).List()
					}
					schedulePayload["recurrence"] = recurrencePayload
				}
				scheduledEvent["schedule"] = schedulePayload
			}
		}
	}

	if ok := d.HasChange("color"); ok {
		scheduledEvent["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		scheduledEvent["comments"] = d.Get("comments")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			scheduledEvent["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		scheduledEvent["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		scheduledEvent["ignore-errors"] = v.(bool)
	}

	log.Println("Update ScheduledEvent - Map = ", scheduledEvent)

	updateScheduledEventRes, err := client.ApiCall("set-scheduled-event", scheduledEvent, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateScheduledEventRes.Success {
		if updateScheduledEventRes.ErrorMsg != "" {
			return fmt.Errorf(updateScheduledEventRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementScheduledEvent(d, m)
}

func deleteManagementScheduledEvent(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	scheduledEventPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ScheduledEvent")

	deleteScheduledEventRes, err := client.ApiCall("delete-scheduled-event", scheduledEventPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteScheduledEventRes.Success {
		if deleteScheduledEventRes.ErrorMsg != "" {
			return fmt.Errorf(deleteScheduledEventRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
