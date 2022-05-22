package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementTime() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementTimeRead,
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
			"end": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "End time. Note: Each gateway may interpret this time differently according to its time zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date in format dd-MMM-yyyy.",
						},
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format. Time zone information is ignored.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
						"time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time in format HH:mm.",
						},
					},
				},
			},
			"end_never": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "End never.",
			},
			"hour_ranges": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Hours recurrence. Note: Each gateway may interpret this time differently according to its time zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Is hour range enabled.",
						},
						"from": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time in format HH:MM.",
						},
						"index": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Hour range index.",
						},
						"to": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time in format HH:MM.",
						},
					},
				},
			},
			"start": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Starting time. Note: Each gateway may interpret this time differently according to its time zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date in format dd-MMM-yyyy.",
						},
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format. Time zone information is ignored.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
						"time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time in format HH:mm.",
						},
					},
				},
			},
			"start_now": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Start immediately.",
			},
			"recurrence": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Computed:    true,
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
						"month": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Valid on month. Example: \"1\", \"2\",\"12\",\"Any\".",
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
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
		},
	}
}

func dataSourceManagementTimeRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showTimeRes, err := client.ApiCall("show-time", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTimeRes.Success {
		return fmt.Errorf(showTimeRes.ErrorMsg)
	}

	time := showTimeRes.GetData()

	log.Println("Read Time - Show JSON = ", time)

	if v := time["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := time["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if time["end"] != nil {

		endMap := time["end"].(map[string]interface{})

		endMapToReturn := make(map[string]interface{})

		if v, _ := endMap["date"]; v != nil {
			endMapToReturn["date"] = v
		}
		if v, _ := endMap["iso-8601"]; v != nil {
			endMapToReturn["iso_8601"] = v
		}
		if v, _ := endMap["posix"]; v != nil {
			endMapToReturn["posix"] = strconv.Itoa(int(math.Round(v.(float64))))
		}
		if v, _ := endMap["time"]; v != nil {
			endMapToReturn["time"] = v
		}

		_ = d.Set("end", endMapToReturn)
	}

	if v := time["end-never"]; v != nil {
		_ = d.Set("end_never", v)
	}

	if time["hours-ranges"] != nil {

		hoursRangesList, ok := time["hours-ranges"].([]interface{})

		if ok {

			if len(hoursRangesList) > 0 {

				var hourRangesListToReturn []map[string]interface{}

				for i := range hoursRangesList {

					hoursRangesMap := hoursRangesList[i].(map[string]interface{})

					hoursRangesMapToAdd := make(map[string]interface{})

					if v, _ := hoursRangesMap["enabled"]; v != nil {
						hoursRangesMapToAdd["enabled"] = v
					}
					if v, _ := hoursRangesMap["from"]; v != nil {
						hoursRangesMapToAdd["from"] = v
					}
					if v, _ := hoursRangesMap["index"]; v != nil {
						hoursRangesMapToAdd["index"] = strconv.Itoa(int(math.Round(v.(float64))))
					}
					if v, _ := hoursRangesMap["to"]; v != nil {
						hoursRangesMapToAdd["to"] = v
					}

					hourRangesListToReturn = append(hourRangesListToReturn, hoursRangesMapToAdd)
				}
				_ = d.Set("hours_ranges", hourRangesListToReturn)
			}
		}
	}

	if time["start"] != nil {

		endMap := time["start"].(map[string]interface{})

		endMapToReturn := make(map[string]interface{})

		if v, _ := endMap["date"]; v != nil {
			endMapToReturn["date"] = v
		}
		if v, _ := endMap["iso-8601"]; v != nil {
			endMapToReturn["iso_8601"] = v
		}
		if v, _ := endMap["posix"]; v != nil {
			endMapToReturn["posix"] = strconv.Itoa(int(math.Round(v.(float64))))
		}
		if v, _ := endMap["time"]; v != nil {
			endMapToReturn["time"] = v
		}

		_ = d.Set("start", endMapToReturn)
	}

	if v := time["start-now"]; v != nil {
		_ = d.Set("start_now", v)
	}

	if time["tags"] != nil {
		tagsJson, ok := time["tags"].([]interface{})
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

	if time["recurrence"] != nil {

		endMap := time["recurrence"].(map[string]interface{})

		endMapToReturn := make(map[string]interface{})

		if v, _ := endMap["days"]; v != nil {
			endMapToReturn["days"] = v
		}
		if v, _ := endMap["month"]; v != nil {
			endMapToReturn["month"] = v
		}
		if v, _ := endMap["pattern"]; v != nil {
			endMapToReturn["pattern"] = v
		}
		if v, _ := endMap["weekdays"]; v != nil {
			endMapToReturn["weekdays"] = v
		}

		_ = d.Set("recurrence", []interface{}{endMapToReturn})
	}

	if v := time["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := time["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := time["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := time["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil
}
