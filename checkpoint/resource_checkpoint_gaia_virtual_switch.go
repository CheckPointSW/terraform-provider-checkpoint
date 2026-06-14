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
func resourceGaiaVirtualSwitch() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaVirtualSwitch,
        Read:   readGaiaVirtualSwitch,
        Update: updateGaiaVirtualSwitch,
        Delete: deleteGaiaVirtualSwitch,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "resource_id": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Virtual switch identifier`,
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Virtual switch name`,
            },
            "interface": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Network interface to be added`,
            },
            "set_if_exist": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `If another virtual switch with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original virtual switch's fields will be overwritten by the fields provided in the request payload!`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "action": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "status": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "message": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "vsxd_task_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "vs_id": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaVirtualSwitch(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    if v, ok := d.GetOkExists("set_if_exist"); ok {
        payload["set-if-exist"] = v.(bool)
    }

    log.Println("Create VirtualSwitch - Map = ", payload)

    addVirtualSwitchRes, err := client.ApiCallSimple("add-virtual-switch", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addVirtualSwitchRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addVirtualSwitchRes.Success {
            errMsg = addVirtualSwitchRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addVirtualSwitchRes.GetData()
        }

        debugLogOperation(
            "virtual-switch",        // resource type
            "create",                       // operation
            "add-virtual-switch",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add virtual-switch: %v", err)
    }
    if !addVirtualSwitchRes.Success {
        if addVirtualSwitchRes.ErrorMsg != "" {
            return fmt.Errorf(addVirtualSwitchRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "add-virtual-switch", addVirtualSwitchRes, true, 0)
    if err != nil {
        return fmt.Errorf("add-virtual-switch task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("add-virtual-switch task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    d.SetId(fmt.Sprintf("virtual-switch-" + acctest.RandString(10)))
    return readGaiaVirtualSwitch(d, m)
}

func readGaiaVirtualSwitch(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showVirtualSwitchRes, err := client.ApiCallSimple("show-virtual-switch", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showVirtualSwitchRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showVirtualSwitchRes.Success {
            errMsg = showVirtualSwitchRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showVirtualSwitchRes.GetData()
        }

        debugLogOperation(
            "virtual-switch",        // resource type
            "read",                       // operation
            "show-virtual-switch",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show virtual-switch: %v", err)
    }
    if !showVirtualSwitchRes.Success {
        if data := showVirtualSwitchRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showVirtualSwitchRes.ErrorMsg)
    }

    virtualSwitch := showVirtualSwitchRes.GetData()

    log.Println("Read VirtualSwitch - Show JSON = ", virtualSwitch)

    if v, exists := virtualSwitch["id"]; exists {
        d.Set("resource_id", fmt.Sprintf("%v", v))
    }
    if v, exists := virtualSwitch["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := virtualSwitch["interfaces"]; exists {
        if iList, ok := v.([]interface{}); ok && len(iList) > 0 {
            d.Set("interface", fmt.Sprintf("%v", iList[0]))
        }
    }
    if v, exists := virtualSwitch["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaVirtualSwitch(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("interface"); ok {
        payload["interface"] = v.(string)
    }

    setVirtualSwitchRes, err := client.ApiCallSimple("set-virtual-switch", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setVirtualSwitchRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setVirtualSwitchRes.Success {
            errMsg = setVirtualSwitchRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setVirtualSwitchRes.GetData()
        }

        debugLogOperation(
            "virtual-switch",        // resource type
            "update",                       // operation
            "set-virtual-switch",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set virtual-switch: %v", err)
    }
    if !setVirtualSwitchRes.Success {
        return fmt.Errorf(setVirtualSwitchRes.ErrorMsg)
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-virtual-switch", setVirtualSwitchRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-virtual-switch task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("set-virtual-switch task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    return readGaiaVirtualSwitch(d, m)
}

func deleteGaiaVirtualSwitch(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(string)
    }

    deleteVirtualSwitchRes, err := client.ApiCallSimple("delete-virtual-switch", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteVirtualSwitchRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteVirtualSwitchRes.Success {
            errMsg = deleteVirtualSwitchRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteVirtualSwitchRes.GetData()
        }

        debugLogOperation(
            "virtual-switch",        // resource type
            "delete",                       // operation
            "delete-virtual-switch",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete virtual-switch: %v", err)
    }
    if !deleteVirtualSwitchRes.Success {
        return fmt.Errorf(deleteVirtualSwitchRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

