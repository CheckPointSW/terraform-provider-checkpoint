package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetAppControlUpdateSchedule() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetAppControlUpdateSchedule,
		Read:   readManagementSetAppControlUpdateSchedule,
		Delete: deleteManagementSetAppControlUpdateSchedule,
		Schema: map[string]*schema.Schema{
			"schedule_management_update": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Application Control & URL Filtering Update Schedule on Management Server.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable/Disable Application Control & URL Filtering Update Schedule on Management Server.",
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
					},
				},
			},
			"schedule_gateway_update": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Application Control & URL Filtering Update Schedule on Gateway.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable/Disable Application Control & URL Filtering Update Schedule on Gateway.",
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
												},
												"interval_hours": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The amount of hours between updates. <font color=\"red\">Required only when</font> pattern is set to 'Interval'.",
												},
												"interval_minutes": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The amount of minutes between updates. <font color=\"red\">Required only when</font> pattern is set to 'Interval'.",
												},
												"interval_seconds": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "The amount of seconds between updates. <font color=\"red\">Required only when</font> pattern is set to 'Interval'.",
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
					},
				},
			},
		},
	}
}

func createManagementSetAppControlUpdateSchedule(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("schedule_management_update"); ok {

		scheduleManagementUpdateList := v.([]interface{})
		if len(scheduleManagementUpdateList) > 0 {

			scheduleManagementPayload := make(map[string]interface{})

			if v, ok := d.GetOkExists("schedule_management_update.0.enabled"); ok {
				scheduleManagementPayload["enabled"] = v.(bool)
			}

			if v, ok := d.GetOk("schedule_management_update.0.schedule"); ok {

				scheduleList := v.([]interface{})

				if len(scheduleList) > 0 {

					schedulePayload := make(map[string]interface{})

					if v, ok := d.GetOk("schedule_management_update.0.schedule.0.time"); ok {
						schedulePayload["time"] = v.(string)
					}

					if _, ok := d.GetOk("schedule_management_update.0.schedule.0.recurrence"); ok {

						recurrencePayLoad := make(map[string]interface{})

						if v, ok := d.GetOk("schedule_management_update.0.schedule.0.recurrence.0.pattern"); ok {
							recurrencePayLoad["pattern"] = v.(string)
						}
						if v, ok := d.GetOk("schedule_management_update.0.schedule.0.recurrence.0.weekdays"); ok {
							recurrencePayLoad["weekdays"] = v.(*schema.Set).List()
						}
						if v, ok := d.GetOk("schedule_management_update.0.schedule.0.recurrence.0.days"); ok {
							recurrencePayLoad["days"] = v.(*schema.Set).List()
						}

						schedulePayload["recurrence"] = recurrencePayLoad
					}
					scheduleManagementPayload["schedule"] = schedulePayload
				}
			}
			payload["schedule-management-update"] = scheduleManagementPayload
		}
	}

	if v, ok := d.GetOk("schedule_gateway_update"); ok {

		scheduleGatewayUpdateList := v.([]interface{})
		if len(scheduleGatewayUpdateList) > 0 {

			scheduleGatewayPayload := make(map[string]interface{})

			if v, ok := d.GetOkExists("schedule_gateway_update.0.enabled"); ok {
				scheduleGatewayPayload["enabled"] = v.(bool)
			}

			if v, ok := d.GetOk("schedule_gateway_update.0.schedule"); ok {

				scheduleList := v.([]interface{})

				if len(scheduleList) > 0 {

					schedulePayload := make(map[string]interface{})

					if v, ok := d.GetOk("schedule_gateway_update.0.schedule.0.time"); ok {
						schedulePayload["time"] = v.(string)
					}

					if _, ok := d.GetOk("schedule_gateway_update.0.schedule.0.recurrence"); ok {

						recurrencePayLoad := make(map[string]interface{})

						if v, ok := d.GetOk("schedule_gateway_update.0.schedule.0.recurrence.0.pattern"); ok {
							recurrencePayLoad["pattern"] = v.(string)
						}
						if v, ok := d.GetOk("schedule_gateway_update.0.schedule.0.recurrence.0.interval_hours"); ok {
							recurrencePayLoad["interval-hours"] = v.(int)
						}
						if v, ok := d.GetOk("schedule_gateway_update.0.schedule.0.recurrence.0.interval_minutes"); ok {
							recurrencePayLoad["interval-minutes"] = v.(int)
						}
						if v, ok := d.GetOk("schedule_gateway_update.0.schedule.0.recurrence.0.interval_seconds"); ok {
							recurrencePayLoad["interval-seconds"] = v.(int)
						}
						if v, ok := d.GetOk("schedule_gateway_update.0.schedule.0.recurrence.0.weekdays"); ok {
							recurrencePayLoad["weekdays"] = v.(*schema.Set).List()
						}
						if v, ok := d.GetOk("schedule_gateway_update.0.schedule.0.recurrence.0.days"); ok {
							recurrencePayLoad["days"] = v.(*schema.Set).List()
						}

						schedulePayload["recurrence"] = recurrencePayLoad
					}
					scheduleGatewayPayload["schedule"] = schedulePayload
				}
			}
			payload["schedule-gateway-update"] = scheduleGatewayPayload
		}
	}

	SetAppControlUpdateScheduleRes, err := client.ApiCallSimple("set-app-control-update-schedule", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetAppControlUpdateScheduleRes.Success {
		return fmt.Errorf(SetAppControlUpdateScheduleRes.ErrorMsg)
	}

	d.SetId("set-app-control-update-schedule-" + acctest.RandString(10))
	return readManagementSetAppControlUpdateSchedule(d, m)
}

func readManagementSetAppControlUpdateSchedule(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetAppControlUpdateSchedule(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
