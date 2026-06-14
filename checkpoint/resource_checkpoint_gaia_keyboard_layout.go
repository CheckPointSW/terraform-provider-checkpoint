package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaKeyboardLayout() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaKeyboardLayout,
        Read:   readGaiaKeyboardLayout,
        Update: updateGaiaKeyboardLayout,
        Delete: deleteGaiaKeyboardLayout,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "keyboard_layout": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Available languages: be-latin1 - Belgian, bg - Bulgarian, br-abnt2 - Brazilian, cf - Central African Republic, cz-lat2 - Czechoslovakian, de - German, dvorak - Dvorák, dk - Danish, et - Estonian, fi - Finnish, fr - French, fr_CH - Swiss French, sg - Swiss German, hu - Hungarian, is-latin1 - Icelandic, it - Italian, jp106 - Japanese, no - Norwegian, pl - Polish, pt-latin1 - Portuguese, ru - Russian, es - Spanish, se-latin1 - Swedish, trq - Turkish, uk - Great Britain, us - US `,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaKeyboardLayout(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("keyboard_layout"); ok {
        payload["keyboard-layout"] = v.(string)
    }

    log.Println("Create KeyboardLayout - Map = ", payload)

    addKeyboardLayoutRes, err := client.ApiCallSimple("set-keyboard-layout", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addKeyboardLayoutRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addKeyboardLayoutRes.Success {
            errMsg = addKeyboardLayoutRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addKeyboardLayoutRes.GetData()
        }

        debugLogOperation(
            "keyboard-layout",        // resource type
            "create",                       // operation
            "set-keyboard-layout",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add keyboard-layout: %v", err)
    }
    if !addKeyboardLayoutRes.Success {
        if addKeyboardLayoutRes.ErrorMsg != "" {
            return fmt.Errorf(addKeyboardLayoutRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("keyboard-layout-" + acctest.RandString(10)))
    return readGaiaKeyboardLayout(d, m)
}

func readGaiaKeyboardLayout(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showKeyboardLayoutRes, err := client.ApiCallSimple("show-keyboard-layout", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showKeyboardLayoutRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showKeyboardLayoutRes.Success {
            errMsg = showKeyboardLayoutRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showKeyboardLayoutRes.GetData()
        }

        debugLogOperation(
            "keyboard-layout",        // resource type
            "read",                       // operation
            "show-keyboard-layout",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show keyboard-layout: %v", err)
    }
    if !showKeyboardLayoutRes.Success {
        if data := showKeyboardLayoutRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showKeyboardLayoutRes.ErrorMsg)
    }

    keyboardLayout := showKeyboardLayoutRes.GetData()

    log.Println("Read KeyboardLayout - Show JSON = ", keyboardLayout)

    if v, exists := keyboardLayout["keyboard-layout"]; exists {
        d.Set("keyboard_layout", fmt.Sprintf("%v", v))
    }
    if v, exists := keyboardLayout["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaKeyboardLayout(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("keyboard_layout"); ok {
        payload["keyboard-layout"] = v.(string)
    }

    setKeyboardLayoutRes, err := client.ApiCallSimple("set-keyboard-layout", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setKeyboardLayoutRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setKeyboardLayoutRes.Success {
            errMsg = setKeyboardLayoutRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setKeyboardLayoutRes.GetData()
        }

        debugLogOperation(
            "keyboard-layout",        // resource type
            "update",                       // operation
            "set-keyboard-layout",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set keyboard-layout: %v", err)
    }
    if !setKeyboardLayoutRes.Success {
        return fmt.Errorf(setKeyboardLayoutRes.ErrorMsg)
    }

    return readGaiaKeyboardLayout(d, m)
}

func deleteGaiaKeyboardLayout(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    