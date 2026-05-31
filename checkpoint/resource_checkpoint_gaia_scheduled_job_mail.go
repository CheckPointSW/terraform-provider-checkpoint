package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaScheduledJobMail() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaScheduledJobMail,
        Read:   readGaiaScheduledJobMail,
        Update: updateGaiaScheduledJobMail,
        Delete: deleteGaiaScheduledJobMail,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "email_address": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `New e-mail address to send reports to`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaScheduledJobMail(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("email_address"); ok {
        payload["email-address"] = v.(string)
    }

    log.Println("Create ScheduledJobMail - Map = ", payload)

    addScheduledJobMailRes, err := client.ApiCallSimple("set-scheduled-job-mail", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addScheduledJobMailRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addScheduledJobMailRes.Success {
            errMsg = addScheduledJobMailRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addScheduledJobMailRes.GetData()
        }

        debugLogOperation(
            "scheduled-job-mail",        // resource type
            "create",                       // operation
            "set-scheduled-job-mail",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add scheduled-job-mail: %v", err)
    }
    if !addScheduledJobMailRes.Success {
        if addScheduledJobMailRes.ErrorMsg != "" {
            return fmt.Errorf(addScheduledJobMailRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("scheduled-job-mail-" + acctest.RandString(10)))
    return readGaiaScheduledJobMail(d, m)
}

func readGaiaScheduledJobMail(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showScheduledJobMailRes, err := client.ApiCallSimple("show-scheduled-job-mail", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showScheduledJobMailRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showScheduledJobMailRes.Success {
            errMsg = showScheduledJobMailRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showScheduledJobMailRes.GetData()
        }

        debugLogOperation(
            "scheduled-job-mail",        // resource type
            "read",                       // operation
            "show-scheduled-job-mail",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show scheduled-job-mail: %v", err)
    }
    if !showScheduledJobMailRes.Success {
        if data := showScheduledJobMailRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showScheduledJobMailRes.ErrorMsg)
    }

    scheduledJobMail := showScheduledJobMailRes.GetData()

    log.Println("Read ScheduledJobMail - Show JSON = ", scheduledJobMail)

    if v, exists := scheduledJobMail["email-address"]; exists {
        d.Set("email_address", fmt.Sprintf("%v", v))
    }
    if v, exists := scheduledJobMail["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaScheduledJobMail(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("email_address"); ok {
        payload["email-address"] = v.(string)
    }

    setScheduledJobMailRes, err := client.ApiCallSimple("set-scheduled-job-mail", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setScheduledJobMailRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setScheduledJobMailRes.Success {
            errMsg = setScheduledJobMailRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setScheduledJobMailRes.GetData()
        }

        debugLogOperation(
            "scheduled-job-mail",        // resource type
            "update",                       // operation
            "set-scheduled-job-mail",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set scheduled-job-mail: %v", err)
    }
    if !setScheduledJobMailRes.Success {
        return fmt.Errorf(setScheduledJobMailRes.ErrorMsg)
    }

    return readGaiaScheduledJobMail(d, m)
}

func deleteGaiaScheduledJobMail(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    