package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaScheduledJob() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaScheduledJob,
        Read:   readGaiaScheduledJob,
        Update: updateGaiaScheduledJob,
        Delete: deleteGaiaScheduledJob,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Scheduled job name`,
            },
            "command": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Scheduled command (expert CLI style)`,
            },
            "recurrence": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `Recurrence schedule`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Job recurrence type`,
                        },
                        "interval": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Time interval in minutes. Relevant for \"interval\" recurrence type`,
                        },
                        "time_of_day": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Time of day in 24 hour format. Relevant for \"daily\", \"weekly\" and \"monthly\" recurrence types`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "hour": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Time hour`,
                                    },
                                    "minute": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Time minute`,
                                    },
                                },
                            },
                        },
                        "hourly": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Hours of day in 24 hour format. Can choose multiple hours. Relevant for \"hourly\" recurrence type`,
                            MaxItems:    1,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "hours_of_day": {
                                        Type:        schema.TypeSet,
                                        Optional:    true,
                                        Description: `Hours of day in 24 hour format`,
                                        Elem: &schema.Schema{
                                            Type: schema.TypeString,
                                        },
                                    },
                                    "minute": {
                                        Type:        schema.TypeInt,
                                        Optional:    true,
                                        Description: `Time minute`,
                                    },
                                },
                            },
                        },
                        "weekdays": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `Days of the week. Relevant for \"weekly\" recurrence type`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "days": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `Days of the month. Relevant for \"monthly\" recurrence type`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                        "months": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: `Month numbers. Relevant for \"monthly\" recurrence type`,
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaScheduledJob(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("command"); ok {
        payload["command"] = v.(string)
    }

    if v := d.Get("recurrence"); len(v.([]interface{})) > 0 {
        _ = v
        recurrenceMap := make(map[string]interface{})
        if v, ok := d.GetOk("recurrence.0.type"); ok {
            recurrenceMap["type"] = v.(string)
        }
        if v, ok := d.GetOk("recurrence.0.interval"); ok {
            recurrenceMap["interval"] = v.(int)
        }
        if v, ok := d.GetOk("recurrence.0.time_of_day"); ok {
            _ = v
            timeofdayMap := make(map[string]interface{})
            if v, ok := d.GetOk("recurrence.0.time_of_day.0.hour"); ok {
                timeofdayMap["hour"] = v.(int)
            }
            if v, ok := d.GetOk("recurrence.0.time_of_day.0.minute"); ok {
                timeofdayMap["minute"] = v.(int)
            }
            if len(timeofdayMap) > 0 {
                recurrenceMap["time-of-day"] = timeofdayMap
            }
        }
        if v, ok := d.GetOk("recurrence.0.hourly"); ok {
            _ = v
            hourlyMap := make(map[string]interface{})
            if v := d.Get("recurrence.0.hourly.0.hours_of_day"); len(v.(*schema.Set).List()) > 0 {
                hourlyMap["hours-of-day"] = v.(*schema.Set).List()
            }
            if v, ok := d.GetOk("recurrence.0.hourly.0.minute"); ok {
                hourlyMap["minute"] = v.(int)
            }
            if len(hourlyMap) > 0 {
                recurrenceMap["hourly"] = hourlyMap
            }
        }
        if v := d.Get("recurrence.0.weekdays"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["weekdays"] = v.(*schema.Set).List()
        }
        if v := d.Get("recurrence.0.days"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["days"] = v.(*schema.Set).List()
        }
        if v := d.Get("recurrence.0.months"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["months"] = v.(*schema.Set).List()
        }
        if len(recurrenceMap) > 0 {
            payload["recurrence"] = recurrenceMap
        }
    }

    log.Println("Create ScheduledJob - Map = ", payload)

    addScheduledJobRes, err := client.ApiCallSimple("add-scheduled-job", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addScheduledJobRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addScheduledJobRes.Success {
            errMsg = addScheduledJobRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addScheduledJobRes.GetData()
        }

        debugLogOperation(
            "scheduled-job",        // resource type
            "create",                       // operation
            "add-scheduled-job",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add scheduled-job: %v", err)
    }
    if !addScheduledJobRes.Success {
        if addScheduledJobRes.ErrorMsg != "" {
            return fmt.Errorf(addScheduledJobRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("scheduled-job-" + acctest.RandString(10)))
    return readGaiaScheduledJob(d, m)
}

func readGaiaScheduledJob(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showScheduledJobRes, err := client.ApiCallSimple("show-scheduled-job", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showScheduledJobRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showScheduledJobRes.Success {
            errMsg = showScheduledJobRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showScheduledJobRes.GetData()
        }

        debugLogOperation(
            "scheduled-job",        // resource type
            "read",                       // operation
            "show-scheduled-job",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show scheduled-job: %v", err)
    }
    if !showScheduledJobRes.Success {
        if data := showScheduledJobRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showScheduledJobRes.ErrorMsg)
    }

    scheduledJob := showScheduledJobRes.GetData()

    log.Println("Read ScheduledJob - Show JSON = ", scheduledJob)

    if v, exists := scheduledJob["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := scheduledJob["command"]; exists {
        cmd := fmt.Sprintf("%v", v)
        cmd = strings.ReplaceAll(cmd, "\\ ", " ")
        d.Set("command", cmd)
    }
    if v, exists := scheduledJob["recurrence"]; exists {
        d.Set("recurrence", v)
    }
    if v, exists := scheduledJob["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaScheduledJob(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("command"); ok {
        payload["command"] = v.(string)
    }

    if v := d.Get("recurrence"); len(v.([]interface{})) > 0 {
        _ = v
        recurrenceMap := make(map[string]interface{})
        if v, ok := d.GetOk("recurrence.0.type"); ok {
            recurrenceMap["type"] = v.(string)
        }
        if v, ok := d.GetOk("recurrence.0.interval"); ok {
            recurrenceMap["interval"] = v.(int)
        }
        if v, ok := d.GetOk("recurrence.0.time_of_day"); ok {
            _ = v
            timeofdayMap := make(map[string]interface{})
            if v, ok := d.GetOk("recurrence.0.time_of_day.0.hour"); ok {
                timeofdayMap["hour"] = v.(int)
            }
            if v, ok := d.GetOk("recurrence.0.time_of_day.0.minute"); ok {
                timeofdayMap["minute"] = v.(int)
            }
            if len(timeofdayMap) > 0 {
                recurrenceMap["time-of-day"] = timeofdayMap
            }
        }
        if v, ok := d.GetOk("recurrence.0.hourly"); ok {
            _ = v
            hourlyMap := make(map[string]interface{})
            if v := d.Get("recurrence.0.hourly.0.hours_of_day"); len(v.(*schema.Set).List()) > 0 {
                hourlyMap["hours-of-day"] = v.(*schema.Set).List()
            }
            if v, ok := d.GetOk("recurrence.0.hourly.0.minute"); ok {
                hourlyMap["minute"] = v.(int)
            }
            if len(hourlyMap) > 0 {
                recurrenceMap["hourly"] = hourlyMap
            }
        }
        if v := d.Get("recurrence.0.weekdays"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["weekdays"] = v.(*schema.Set).List()
        }
        if v := d.Get("recurrence.0.days"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["days"] = v.(*schema.Set).List()
        }
        if v := d.Get("recurrence.0.months"); len(v.(*schema.Set).List()) > 0 {
            recurrenceMap["months"] = v.(*schema.Set).List()
        }
        if len(recurrenceMap) > 0 {
            payload["recurrence"] = recurrenceMap
        }
    }

    setScheduledJobRes, err := client.ApiCallSimple("set-scheduled-job", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setScheduledJobRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setScheduledJobRes.Success {
            errMsg = setScheduledJobRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setScheduledJobRes.GetData()
        }

        debugLogOperation(
            "scheduled-job",        // resource type
            "update",                       // operation
            "set-scheduled-job",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set scheduled-job: %v", err)
    }
    if !setScheduledJobRes.Success {
        return fmt.Errorf(setScheduledJobRes.ErrorMsg)
    }

    return readGaiaScheduledJob(d, m)
}

func deleteGaiaScheduledJob(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteScheduledJobRes, err := client.ApiCallSimple("delete-scheduled-job", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteScheduledJobRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteScheduledJobRes.Success {
            errMsg = deleteScheduledJobRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteScheduledJobRes.GetData()
        }

        debugLogOperation(
            "scheduled-job",        // resource type
            "delete",                       // operation
            "delete-scheduled-job",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete scheduled-job: %v", err)
    }
    if !deleteScheduledJobRes.Success {
        return fmt.Errorf(deleteScheduledJobRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

