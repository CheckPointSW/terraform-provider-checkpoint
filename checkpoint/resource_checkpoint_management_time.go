package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"math"
	"strconv"
)

func resourceManagementTime() *schema.Resource {
	return &schema.Resource{
		Create: createManagementTime,
		Read:   readManagementTime,
		Update: updateManagementTime,
		Delete: deleteManagementTime,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"end": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "End time. Note: Each gateway may interpret this time differently according to its time zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date in format dd-MMM-yyyy.",
						},
						"time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time in format HH:mm.",
						},
					},
				},
			},
			"end_never": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "End never.",
				Default:     true,
			},
			"hours_ranges": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Hours recurrence. Note: Each gateway may interpret this time differently according to its time zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Is hour range enabled.",
						},
						"from": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time in format HH:MM.",
						},
						"index": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Hour range index.",
						},
						"to": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time in format HH:MM.",
						},
					},
				},
			},
			"start": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Starting time. Note: Each gateway may interpret this time differently according to its time zone.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Date in format dd-MMM-yyyy.",
						},
						"time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Time in format HH:mm.",
						},
					},
				},
			},
			"start_now": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Start immediately.",
				Default:     true,
			},
			"recurrence": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Days recurrence.",
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
						"month": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Valid on month. Example: \"1\", \"2\",\"12\",\"Any\".",
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func createManagementTime(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	time := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		time["name"] = v.(string)
	}

	if _, ok := d.GetOk("end"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("end.date"); ok {
			res["date"] = v
		}
		if v, ok := d.GetOk("end.time"); ok {
			res["time"] = v
		}
		time["end"] = res
	}

	if v, ok := d.GetOkExists("end_never"); ok {
		time["end-never"] = v.(bool)
	}

	if v, ok := d.GetOk("hours_ranges"); ok {

		hoursRangesList := v.([]interface{})

		if len(hoursRangesList) > 0 {
			var hourRangesPayload []map[string]interface{}

			for i := range hoursRangesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("hours_ranges." + strconv.Itoa(i) + ".enabled"); ok {
					Payload["enabled"] = v.(bool)
				}
				if v, ok := d.GetOk("hours_ranges." + strconv.Itoa(i) + ".from"); ok {
					Payload["from"] = v.(string)
				}
				if v, ok := d.GetOk("hours_ranges." + strconv.Itoa(i) + ".index"); ok {
					Payload["index"] = v.(int)
				}
				if v, ok := d.GetOk("hours_ranges." + strconv.Itoa(i) + ".to"); ok {
					Payload["to"] = v.(string)
				}
				hourRangesPayload = append(hourRangesPayload, Payload)
			}
			time["hours-ranges"] = hourRangesPayload
		}
	}

	if _, ok := d.GetOk("start"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("start.date"); ok {
			res["date"] = v
		}
		if v, ok := d.GetOk("start.time"); ok {
			res["time"] = v
		}
		time["start"] = res
	}

	if v, ok := d.GetOkExists("start_now"); ok {
		time["start-now"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		time["tags"] = v.(*schema.Set).List()
	}

	if _, ok := d.GetOk("recurrence"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("recurrence.days"); ok {
			res["days"] = v
		}
		if v, ok := d.GetOk("recurrence.month"); ok {
			res["month"] = v
		}
		if v, ok := d.GetOk("recurrence.pattern"); ok {
			res["pattern"] = v.(string)
		}
		if v, ok := d.GetOk("recurrence.weekdays"); ok {
			res["weekdays"] = v
		}
		time["recurrence"] = res
	}

	if v, ok := d.GetOk("color"); ok {
		time["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		time["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		time["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		time["ignore-errors"] = v.(bool)
	}

	log.Println("Create Time - Map = ", time)

	addTimeRes, err := client.ApiCall("add-time", time, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addTimeRes.Success {
		if addTimeRes.ErrorMsg != "" {
			return fmt.Errorf(addTimeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addTimeRes.GetData()["uid"].(string))

	return readManagementTime(d, m)
}

func readManagementTime(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showTimeRes, err := client.ApiCall("show-time", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTimeRes.Success {
		if objectNotFound(showTimeRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showTimeRes.ErrorMsg)
	}

	time := showTimeRes.GetData()

	log.Println("Read Time - Show JSON = ", time)

	if v := time["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if time["end"] != nil {
		defaultEndMap := map[string]interface{}{
			"date":     "01-Jan-1970",
			"time":     "00:00",
		}
		endMap := time["end"].(map[string]interface{})

		endMapToReturn := make(map[string]interface{})

		if v, _ := endMap["date"]; v != nil && isArgDefault(v.(string), d, "end.date", defaultEndMap["date"].(string)) {
			endMapToReturn["date"] = v
		}
		if v, _ := endMap["time"]; v != nil && isArgDefault(v.(string), d, "end.time", defaultEndMap["time"].(string)) {
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
						hoursRangesMapToAdd["index"] = int(math.Round(v.(float64)))
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

		defaultStartMap := map[string]interface{}{
			"date":     "01-Jan-1970",
			"time":     "00:00",
		}
		startMap := time["start"].(map[string]interface{})

		startMapToReturn := make(map[string]interface{})

		if v, _ := startMap["date"]; v != nil && isArgDefault(v.(string), d, "end.date", defaultStartMap["date"].(string)) {
			startMapToReturn["date"] = v
		}
		if v, _ := startMap["time"]; v != nil && isArgDefault(v.(string), d, "end.time", defaultStartMap["time"].(string)) {
			startMapToReturn["time"] = v
		}

		_ = d.Set("start", startMapToReturn)
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
		defaultRecurrenceMap := map[string]interface{}{
			"pattern": "Daily",
			"month":   "Any",
		}
		endMap := time["recurrence"].(map[string]interface{})

		endMapToReturn := make(map[string]interface{})

		if v, _ := endMap["days"]; v != nil {
			tagsJson, ok := endMap["days"].([]interface{})
			if ok {
				tagsIds := make([]string, 0)
				if len(tagsJson) > 0 {
					for _, tags := range tagsJson {
						tagsIds = append(tagsIds, tags.(string))
					}
					endMapToReturn["days"] = tagsIds
				}
			}
		}
		if v, _ := endMap["month"]; v != nil && isArgDefault(v.(string), d, "recurrence.month", defaultRecurrenceMap["month"].(string)) {
			endMapToReturn["month"] = v
		}
		if v, _ := endMap["pattern"]; v != nil && isArgDefault(v.(string), d, "recurrence.pattern", defaultRecurrenceMap["pattern"].(string)) {
			endMapToReturn["pattern"] = v
		}
		if v, _ := endMap["weekdays"]; v != nil {
			tagsJson, ok := endMap["weekdays"].([]interface{})
			if ok {
				tagsIds := make([]string, 0)
				if len(tagsJson) > 0 {
					for _, tags := range tagsJson {
						tagsIds = append(tagsIds, tags.(string))
					}
					endMapToReturn["weekdays"] = tagsIds
				}
			}
		}
		if len(endMapToReturn) > 0 {
			_ = d.Set("recurrence", []interface{}{endMapToReturn})
		}
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

func updateManagementTime(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	time := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		time["name"] = oldName
		time["new-name"] = newName
	} else {
		time["name"] = d.Get("name")
	}

	if d.HasChange("end") {
		defaultEndMap := map[string]interface{}{
			"date":     "01-Jan-1970",
			"time":     "00:00",
		}

		res := make(map[string]interface{})

		if v := d.Get("end.date"); v != nil && v != "" {
			res["date"] = d.Get("end.date")
		}
		if v := d.Get("end.time"); v != nil && v != "" {
			res["time"] = d.Get("end.time")
		}
		if len(res) > 0 {
			time["end"] = res
		} else {
			time["end"] = defaultEndMap
		}
	}

	if v, ok := d.GetOk("end_never"); ok {
		time["end-never"] = v.(bool)
	}

	if d.HasChange("hours_ranges") {
		var serversPayload []map[string]interface{}
		defaultAmountOfRanges := 3
		for i := 0; i < defaultAmountOfRanges; i++ {
			serverPayload := make(map[string]interface{})
			if v, ok := d.GetOk("hours_ranges." + strconv.Itoa(i) + ".enabled"); ok {
				serverPayload["enabled"] = v.(bool)
			} else {
				serverPayload["enabled"] = false
			}
			if v, ok := d.GetOk("hours_ranges." + strconv.Itoa(i) + ".from"); ok {
				serverPayload["from"] = v
			} else {
				serverPayload["from"] = "00:00"
			}
			if v, ok := d.GetOk("hours_ranges." + strconv.Itoa(i) + ".index"); ok {
				serverPayload["index"] = v.(int)
			} else {
				serverPayload["index"] = i + 1
			}
			if v, ok := d.GetOk("hours_ranges." + strconv.Itoa(i) + ".to"); ok {
				serverPayload["to"] = v
			} else {
				serverPayload["to"] = "00:00"
			}
			serversPayload = append(serversPayload, serverPayload)
		}
		time["hours-ranges"] = serversPayload
	}

	if d.HasChange("start") {
		defaultStartMap := map[string]interface{}{
			"date":     "01-Jan-1970",
			"time":     "00:00",
		}
		res := make(map[string]interface{})

		if v := d.Get("start.date"); v != nil && v != "" {
			res["date"] = d.Get("start.date")
		}
		if v := d.Get("start.time"); v != nil && v != "" {
			res["time"] = d.Get("start.time")
		}
		if len(res) > 0 {
			time["start"] = res
		} else {
			time["start"] = defaultStartMap
		}

	}

	if d.HasChange("start_now"){
		time["start-now"] = d.Get("start_now")
	}

	if d.HasChange("recurrence") {
		defaultRecurrenceMap := map[string]interface{}{
			"pattern": "Daily",
			"month":   "Any",
		}
		res := make(map[string]interface{})

		if d.HasChange("recurrence.days") {
			res["days"] = d.Get("recurrence.days")
		} else {
			res["days"] = nil
		}
		if d.HasChange("recurrence.month") {
			res["month"] = d.Get("recurrence.month")
		} else {
			res["month"] = defaultRecurrenceMap["month"]
		}
		if d.HasChange("recurrence.pattern") {
			res["pattern"] = d.Get("recurrence.pattern")
		} else {
			res["pattern"] = defaultRecurrenceMap["pattern"]
		}
		if d.HasChange("recurrence.weekdays") {
			res["weekdays"] = d.Get("recurrence.weekdays")
		} else {
			res["weekdays"] = nil
		}
		time["recurrence"] = res
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			time["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			time["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		time["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		time["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOk("ignore_warnings"); ok {
		time["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOk("ignore_errors"); ok {
		time["ignore-errors"] = v.(bool)
	}

	log.Println("Update Time - Map = ", time)

	updateTimeRes, err := client.ApiCall("set-time", time, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateTimeRes.Success {
		if updateTimeRes.ErrorMsg != "" {
			return fmt.Errorf(updateTimeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementTime(d, m)
}

func deleteManagementTime(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	timePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete Time")

	deleteTimeRes, err := client.ApiCall("delete-time", timePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteTimeRes.Success {
		if deleteTimeRes.ErrorMsg != "" {
			return fmt.Errorf(deleteTimeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
