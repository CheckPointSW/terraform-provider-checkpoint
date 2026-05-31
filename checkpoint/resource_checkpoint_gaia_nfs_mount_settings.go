package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaNfsMountSettings() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaNfsMountSettings,
        Read:   readGaiaNfsMountSettings,
        Update: updateGaiaNfsMountSettings,
        Delete: deleteGaiaNfsMountSettings,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "timeout": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Nfs timeout in seconds.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaNfsMountSettings(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("timeout"); ok {
        payload["timeout"] = v.(int)
    }

    log.Println("Create NfsMountSettings - Map = ", payload)

    addNfsMountSettingsRes, err := client.ApiCallSimple("set-nfs-mount-settings", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addNfsMountSettingsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addNfsMountSettingsRes.Success {
            errMsg = addNfsMountSettingsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addNfsMountSettingsRes.GetData()
        }

        debugLogOperation(
            "nfs-mount-settings",        // resource type
            "create",                       // operation
            "set-nfs-mount-settings",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add nfs-mount-settings: %v", err)
    }
    if !addNfsMountSettingsRes.Success {
        if addNfsMountSettingsRes.ErrorMsg != "" {
            return fmt.Errorf(addNfsMountSettingsRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("nfs-mount-settings-" + acctest.RandString(10)))
    return readGaiaNfsMountSettings(d, m)
}

func readGaiaNfsMountSettings(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showNfsMountSettingsRes, err := client.ApiCallSimple("show-nfs-mount-settings", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showNfsMountSettingsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showNfsMountSettingsRes.Success {
            errMsg = showNfsMountSettingsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showNfsMountSettingsRes.GetData()
        }

        debugLogOperation(
            "nfs-mount-settings",        // resource type
            "read",                       // operation
            "show-nfs-mount-settings",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show nfs-mount-settings: %v", err)
    }
    if !showNfsMountSettingsRes.Success {
        if data := showNfsMountSettingsRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showNfsMountSettingsRes.ErrorMsg)
    }

    nfsMountSettings := showNfsMountSettingsRes.GetData()

    log.Println("Read NfsMountSettings - Show JSON = ", nfsMountSettings)

    if v, exists := nfsMountSettings["NFS-Version-4-Settings"]; exists {
        if settingsMap, ok := v.(map[string]interface{}); ok {
            if timeout, timeoutExists := settingsMap["timeout"]; timeoutExists {
                d.Set("timeout", timeout)
            }
        }
    }
    if v, exists := nfsMountSettings["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaNfsMountSettings(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("timeout"); ok {
        payload["timeout"] = v.(int)
    }

    setNfsMountSettingsRes, err := client.ApiCallSimple("set-nfs-mount-settings", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setNfsMountSettingsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setNfsMountSettingsRes.Success {
            errMsg = setNfsMountSettingsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setNfsMountSettingsRes.GetData()
        }

        debugLogOperation(
            "nfs-mount-settings",        // resource type
            "update",                       // operation
            "set-nfs-mount-settings",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set nfs-mount-settings: %v", err)
    }
    if !setNfsMountSettingsRes.Success {
        return fmt.Errorf(setNfsMountSettingsRes.ErrorMsg)
    }

    return readGaiaNfsMountSettings(d, m)
}

func deleteGaiaNfsMountSettings(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    