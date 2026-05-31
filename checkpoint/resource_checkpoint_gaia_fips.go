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
func resourceGaiaFips() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaFips,
        Read:   readGaiaFips,
        Update: updateGaiaFips,
        Delete: deleteGaiaFips,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "enabled": {
                Type:        schema.TypeBool,
                Required:    true,
                Description: `FIPS mode enabled status`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "task_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `Asynchronous task unique identifier`,
            },
        },
    }
}

func createGaiaFips(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    log.Println("Create Fips - Map = ", payload)

    addFipsRes, err := client.ApiCallSimple("set-fips", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addFipsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addFipsRes.Success {
            errMsg = addFipsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addFipsRes.GetData()
        }

        debugLogOperation(
            "fips",        // resource type
            "create",                       // operation
            "set-fips",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add fips: %v", err)
    }
    if !addFipsRes.Success {
        if addFipsRes.ErrorMsg != "" {
            return fmt.Errorf(addFipsRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-fips", addFipsRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-fips task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("set-fips task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    d.SetId(fmt.Sprintf("fips-" + acctest.RandString(10)))
    return readGaiaFips(d, m)
}

func readGaiaFips(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showFipsRes, err := client.ApiCallSimple("show-fips", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showFipsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showFipsRes.Success {
            errMsg = showFipsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showFipsRes.GetData()
        }

        debugLogOperation(
            "fips",        // resource type
            "read",                       // operation
            "show-fips",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show fips: %v", err)
    }
    if !showFipsRes.Success {
        if data := showFipsRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showFipsRes.ErrorMsg)
    }

    fips := showFipsRes.GetData()

    log.Println("Read Fips - Show JSON = ", fips)

    if v, exists := fips["enabled"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("enabled", b)
        } else if s, ok := v.(string); ok {
            d.Set("enabled", s == "true")
        }
    }
    if v, exists := fips["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaFips(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOkExists("enabled"); ok {
        payload["enabled"] = v.(bool)
    }

    setFipsRes, err := client.ApiCallSimple("set-fips", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setFipsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setFipsRes.Success {
            errMsg = setFipsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setFipsRes.GetData()
        }

        debugLogOperation(
            "fips",        // resource type
            "update",                       // operation
            "set-fips",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set fips: %v", err)
    }
    if !setFipsRes.Success {
        return fmt.Errorf(setFipsRes.ErrorMsg)
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-fips", setFipsRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-fips task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("set-fips task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }
    d.Set("task_id", taskRes.TaskID)

    return readGaiaFips(d, m)
}

func deleteGaiaFips(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    