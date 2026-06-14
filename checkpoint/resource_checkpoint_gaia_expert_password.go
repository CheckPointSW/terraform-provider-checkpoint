package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaExpertPassword() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaExpertPassword,
        Read:   readGaiaExpertPassword,
        Update: updateGaiaExpertPassword,
        Delete: deleteGaiaExpertPassword,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "password": {
                Type:        schema.TypeString,
                Optional:    true,
                Sensitive:   true,
                Description: `expert new password`,
            },
            "password_hash": {
                Type:        schema.TypeString,
                Optional:    true,
                Sensitive:   true,
                Description: `An encrypted representation of the password`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaExpertPassword(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v, ok := d.GetOk("password_hash"); ok {
        payload["password-hash"] = v.(string)
    }

    log.Println("Create ExpertPassword - Map = ", payload)

    addExpertPasswordRes, err := client.ApiCallSimple("set-expert-password", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addExpertPasswordRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addExpertPasswordRes.Success {
            errMsg = addExpertPasswordRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addExpertPasswordRes.GetData()
        }

        debugLogOperation(
            "expert-password",        // resource type
            "create",                       // operation
            "set-expert-password",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add expert-password: %v", err)
    }
    if !addExpertPasswordRes.Success {
        if addExpertPasswordRes.ErrorMsg != "" {
            return fmt.Errorf(addExpertPasswordRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("expert-password-" + acctest.RandString(10)))
    return readGaiaExpertPassword(d, m)
}

func readGaiaExpertPassword(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showExpertPasswordRes, err := client.ApiCallSimple("show-expert-password", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showExpertPasswordRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showExpertPasswordRes.Success {
            errMsg = showExpertPasswordRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showExpertPasswordRes.GetData()
        }

        debugLogOperation(
            "expert-password",        // resource type
            "read",                       // operation
            "show-expert-password",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show expert-password: %v", err)
    }
    if !showExpertPasswordRes.Success {
        if data := showExpertPasswordRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showExpertPasswordRes.ErrorMsg)
    }

    expertPassword := showExpertPasswordRes.GetData()

    log.Println("Read ExpertPassword - Show JSON = ", expertPassword)

    if v, exists := expertPassword["password-hash"]; exists {
        d.Set("password_hash", fmt.Sprintf("%v", v))
    }
    if v, exists := expertPassword["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaExpertPassword(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v, ok := d.GetOk("password_hash"); ok {
        payload["password-hash"] = v.(string)
    }

    setExpertPasswordRes, err := client.ApiCallSimple("set-expert-password", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setExpertPasswordRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setExpertPasswordRes.Success {
            errMsg = setExpertPasswordRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setExpertPasswordRes.GetData()
        }

        debugLogOperation(
            "expert-password",        // resource type
            "update",                       // operation
            "set-expert-password",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set expert-password: %v", err)
    }
    if !setExpertPasswordRes.Success {
        return fmt.Errorf(setExpertPasswordRes.ErrorMsg)
    }

    return readGaiaExpertPassword(d, m)
}

func deleteGaiaExpertPassword(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    