package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "context"
    "strings"

)
func resourceGaiaTimeAndDate() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaTimeAndDate,
        Read:   readGaiaTimeAndDate,
        Update: updateGaiaTimeAndDate,
        Delete: deleteGaiaTimeAndDate,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "time": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Time to set, in HH:MM[:SS] format`,
            },
            "timezone": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Timezone in Area / Region format. See timezones list via 'show-timezones'`,
                DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
                norm := func(s string) string {
                    return strings.ReplaceAll(strings.ReplaceAll(s, " / ", "/"), " /", "/")  
                }
                return norm(old) == norm(new)
            },
            },
            "date": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Date to set, in YYYY-MM-DD format`,
                DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
                normalize := func(s string) string {
                    if len(s) == 10 && s[2] == '-' && s[5] == '-' {
                        return s[6:10] + "-" + s[3:5] + "-" + s[0:2]
                    }
                    return s
                }
                return normalize(old) == normalize(new)
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

func createGaiaTimeAndDate(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("time"); ok {
        payload["time"] = v.(string)
    }

    if v, ok := d.GetOk("timezone"); ok {
        payload["timezone"] = strings.ReplaceAll(strings.ReplaceAll(v.(string), " / ", "/"), "/", " / ")
    }

    if v, ok := d.GetOk("date"); ok {
        payload["date"] = v.(string)
    }

    log.Println("Create TimeAndDate - Map = ", payload)

    addTimeAndDateRes, err := client.ApiCallSimple("set-time-and-date", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addTimeAndDateRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addTimeAndDateRes.Success {
            errMsg = addTimeAndDateRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addTimeAndDateRes.GetData()
        }

        debugLogOperation(
            "time-and-date",        // resource type
            "create",                       // operation
            "set-time-and-date",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add time-and-date: %v", err)
    }
    if !addTimeAndDateRes.Success {
        if addTimeAndDateRes.ErrorMsg != "" {
            return fmt.Errorf(addTimeAndDateRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-time-and-date", addTimeAndDateRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-time-and-date task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("set-time-and-date task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    d.SetId(fmt.Sprintf("time-and-date-" + acctest.RandString(10)))
    return readGaiaTimeAndDate(d, m)
}

func readGaiaTimeAndDate(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showTimeAndDateRes, err := client.ApiCallSimple("show-time-and-date", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showTimeAndDateRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showTimeAndDateRes.Success {
            errMsg = showTimeAndDateRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showTimeAndDateRes.GetData()
        }

        debugLogOperation(
            "time-and-date",        // resource type
            "read",                       // operation
            "show-time-and-date",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show time-and-date: %v", err)
    }
    if !showTimeAndDateRes.Success {
        if data := showTimeAndDateRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showTimeAndDateRes.ErrorMsg)
    }

    timeAndDate := showTimeAndDateRes.GetData()

    log.Println("Read TimeAndDate - Show JSON = ", timeAndDate)

    if v, exists := timeAndDate["date"]; exists {
        if s, ok := v.(string); ok {
            if len(s) == 10 && s[2] == '-' && s[5] == '-' {
                s = s[6:10] + "-" + s[3:5] + "-" + s[0:2]
            }
            d.Set("date", s)
        }
    }
    if v, exists := timeAndDate["timezone"]; exists {
        if s, ok := v.(string); ok {
            d.Set("timezone", strings.ReplaceAll(s, " / ", "/"))
        }
    }
    if v, exists := timeAndDate["iso8601"]; exists {
        d.Set("iso8601", fmt.Sprintf("%v", v))
    }
    if v, exists := timeAndDate["posix"]; exists {
        d.Set("posix", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaTimeAndDate(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("time"); ok {
        payload["time"] = v.(string)
    }

    if v, ok := d.GetOk("timezone"); ok {
        payload["timezone"] = strings.ReplaceAll(strings.ReplaceAll(v.(string), " / ", "/"), "/", " / ")
    }

    if v, ok := d.GetOk("date"); ok {
        payload["date"] = v.(string)
    }

    setTimeAndDateRes, err := client.ApiCallSimple("set-time-and-date", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setTimeAndDateRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setTimeAndDateRes.Success {
            errMsg = setTimeAndDateRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setTimeAndDateRes.GetData()
        }

        debugLogOperation(
            "time-and-date",        // resource type
            "update",                       // operation
            "set-time-and-date",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set time-and-date: %v", err)
    }
    if !setTimeAndDateRes.Success {
        return fmt.Errorf(setTimeAndDateRes.ErrorMsg)
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-time-and-date", setTimeAndDateRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-time-and-date task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("set-time-and-date task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    return readGaiaTimeAndDate(d, m)
}

func deleteGaiaTimeAndDate(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    