package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementAppControlUpdateSchedule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAppControlUpdateScheduleRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"schedule_management_update": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Application Control & URL Filtering Update Schedule on Management Server.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable/Disable Application Control & URL Filtering Update Schedule on Management Server.",
						},
						"schedule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Schedule Configuration.",
							MaxItems:    1,
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
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"pattern": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Days recurrence pattern.",
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
					},
				},
			},
			"schedule_gateway_update": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Application Control & URL Filtering Update Schedule on Gateway.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable/Disable Application Control & URL Filtering Update Schedule on Gateway.",
						},
						"schedule": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Schedule Configuration.",
							MaxItems:    1,
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
										MaxItems:    1,
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
					},
				},
			},
		},
	}
}

func dataSourceManagementAppControlUpdateScheduleRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	appControlUpdateScheduleRes, err := client.ApiCallSimple("show-app-control-update-schedule", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !appControlUpdateScheduleRes.Success {
		return fmt.Errorf(appControlUpdateScheduleRes.ErrorMsg)
	}
	appControlUpdateScheduleData := appControlUpdateScheduleRes.GetData()

	if v := appControlUpdateScheduleData["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := appControlUpdateScheduleData["schedule-management-update"]; v != nil {

		innerMap := v.(map[string]interface{})

		mapToReturn := make(map[string]interface{})

		if v := innerMap["enabled"]; v != nil {
			mapToReturn["enabled"] = v
		}

		if v := innerMap["schedule"]; v != nil {
			innerScheduleMap := v.(map[string]interface{})

			scheduleMapToReturn := make(map[string]interface{})

			if v := innerScheduleMap["time"]; v != nil {
				scheduleMapToReturn["time"] = v
			}

			if v := innerScheduleMap["recurrence"]; v != nil {

				innerRecurrenceMap := v.(map[string]interface{})

				recurrenceMapToReturn := make(map[string]interface{})

				if v := innerRecurrenceMap["pattern"]; v != nil {
					recurrenceMapToReturn["pattern"] = v
				}

				if v := innerRecurrenceMap["weekdays"]; v != nil {
					recurrenceMapToReturn["weekdays"] = v.(*schema.Set).List()
				}

				if v := innerRecurrenceMap["days"]; v != nil {
					recurrenceMapToReturn["days"] = v.(*schema.Set).List()
				}

				scheduleMapToReturn["recurrence"] = []interface{}{recurrenceMapToReturn}
			}

			mapToReturn["schedule"] = []interface{}{scheduleMapToReturn}
		}
		_ = d.Set("schedule_management_update", []interface{}{mapToReturn})
	}

	if v := appControlUpdateScheduleData["schedule-gateway-update"]; v != nil {

		innerMap := v.(map[string]interface{})

		mapToReturn := make(map[string]interface{})

		if v := innerMap["enabled"]; v != nil {
			mapToReturn["enabled"] = v
		}

		if v := innerMap["schedule"]; v != nil {
			innerScheduleMap := v.(map[string]interface{})

			scheduleMapToReturn := make(map[string]interface{})

			if v := innerScheduleMap["time"]; v != nil {
				scheduleMapToReturn["time"] = v
			}

			if v := innerScheduleMap["recurrence"]; v != nil {

				innerRecurrenceMap := v.(map[string]interface{})

				recurrenceMapToReturn := make(map[string]interface{})

				if v := innerRecurrenceMap["pattern"]; v != nil {
					recurrenceMapToReturn["pattern"] = v
				}

				if v := innerRecurrenceMap["interval-hours"]; v != nil {
					recurrenceMapToReturn["interval_hours"] = v
				}

				if v := innerRecurrenceMap["interval-minutes"]; v != nil {
					recurrenceMapToReturn["interval_minutes"] = v
				}

				if v := innerRecurrenceMap["interval-seconds"]; v != nil {
					recurrenceMapToReturn["interval_seconds"] = v
				}

				if v := innerRecurrenceMap["weekdays"]; v != nil {
					recurrenceMapToReturn["weekdays"] = v.(*schema.Set).List()
				}

				if v := innerRecurrenceMap["days"]; v != nil {
					recurrenceMapToReturn["days"] = v.(*schema.Set).List()
				}

				scheduleMapToReturn["recurrence"] = []interface{}{recurrenceMapToReturn}
			}

			mapToReturn["schedule"] = []interface{}{scheduleMapToReturn}
		}
		_ = d.Set("schedule_gateway_update", []interface{}{mapToReturn})
	}

	return nil
}
