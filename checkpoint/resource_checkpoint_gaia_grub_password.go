package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaGrubPassword() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaGrubPassword,
        Read:   readGaiaGrubPassword,
        Update: updateGaiaGrubPassword,
        Delete: deleteGaiaGrubPassword,
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
                Description: `GRUB new password`,
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

func createGaiaGrubPassword(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v, ok := d.GetOk("password_hash"); ok {
        payload["password-hash"] = v.(string)
    }

    log.Println("Create GrubPassword - Map = ", payload)

    addGrubPasswordRes, err := client.ApiCallSimple("set-grub-password", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addGrubPasswordRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addGrubPasswordRes.Success {
            errMsg = addGrubPasswordRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addGrubPasswordRes.GetData()
        }

        debugLogOperation(
            "grub-password",        // resource type
            "create",                       // operation
            "set-grub-password",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add grub-password: %v", err)
    }
    if !addGrubPasswordRes.Success {
        if addGrubPasswordRes.ErrorMsg != "" {
            return fmt.Errorf(addGrubPasswordRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("grub-password-" + acctest.RandString(10)))
    return readGaiaGrubPassword(d, m)
}

func readGaiaGrubPassword(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showGrubPasswordRes, err := client.ApiCallSimple("show-grub-password", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showGrubPasswordRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showGrubPasswordRes.Success {
            errMsg = showGrubPasswordRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showGrubPasswordRes.GetData()
        }

        debugLogOperation(
            "grub-password",        // resource type
            "read",                       // operation
            "show-grub-password",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show grub-password: %v", err)
    }
    if !showGrubPasswordRes.Success {
        if data := showGrubPasswordRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showGrubPasswordRes.ErrorMsg)
    }

    grubPassword := showGrubPasswordRes.GetData()

    log.Println("Read GrubPassword - Show JSON = ", grubPassword)

    if v, exists := grubPassword["password-hash"]; exists {
        d.Set("password_hash", fmt.Sprintf("%v", v))
    }
    if v, exists := grubPassword["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaGrubPassword(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("password"); ok {
        payload["password"] = v.(string)
    }

    if v, ok := d.GetOk("password_hash"); ok {
        payload["password-hash"] = v.(string)
    }

    setGrubPasswordRes, err := client.ApiCallSimple("set-grub-password", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setGrubPasswordRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setGrubPasswordRes.Success {
            errMsg = setGrubPasswordRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setGrubPasswordRes.GetData()
        }

        debugLogOperation(
            "grub-password",        // resource type
            "update",                       // operation
            "set-grub-password",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set grub-password: %v", err)
    }
    if !setGrubPasswordRes.Success {
        return fmt.Errorf(setGrubPasswordRes.ErrorMsg)
    }

    return readGaiaGrubPassword(d, m)
}

func deleteGaiaGrubPassword(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    