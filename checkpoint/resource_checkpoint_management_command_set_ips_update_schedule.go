package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementSetIpsUpdateSchedule() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementSetIpsUpdateSchedule,
            Read:   readManagementSetIpsUpdateSchedule,
            Delete: deleteManagementSetIpsUpdateSchedule,
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
                Type:        schema.TypeMap,
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
    return readManagementSetIpsUpdateSchedule(d, m)
}

func readManagementSetIpsUpdateSchedule(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    if v, ok := d.GetOk("time"); ok {
        payload["time"] = v.(string)
    }

    if _, ok := d.GetOk("recurrence"); ok {

        res := make(map[string]interface{})

        if v, ok := d.GetOk("recurrence.days"); ok {
            res["days"] = v
        }
        if v, ok := d.GetOk("recurrence.minutes"); ok {
            res["minutes"] = v
        }
        if v, ok := d.GetOk("recurrence.pattern"); ok {
            res["pattern"] = v.(string)
        }
        if v, ok := d.GetOk("recurrence.weekdays"); ok {
            res["weekdays"] = v
        }
        payload["recurrence"] = res
    }

    SetIpsUpdateScheduleRes, _ := client.ApiCall("set-ips-update-schedule", payload, client.GetSessionID(), true, false)
    if !SetIpsUpdateScheduleRes.Success {
        return fmt.Errorf(SetIpsUpdateScheduleRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementSetIpsUpdateSchedule(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

