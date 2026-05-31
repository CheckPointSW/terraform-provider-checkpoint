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
func resourceGaiaLightshotPartition() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaLightshotPartition,
        Read:   readGaiaLightshotPartition,
        Update: updateGaiaLightshotPartition,
        Delete: deleteGaiaLightshotPartition,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "size": {
                Type:        schema.TypeInt,
                Required:    true,
                Description: `New size (GB) for setting light shotpartition`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "used": {
                Type:     schema.TypeInt,
                Computed: true,
            },
            "available": {
                Type:     schema.TypeInt,
                Computed: true,
            },
            "mount_on": {
                Type:     schema.TypeString,
                Computed: true,
            },
        },
    }
}

func createGaiaLightshotPartition(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("size"); ok {
        payload["size"] = v.(int)
    }

    log.Println("Create LightshotPartition - Map = ", payload)

    addLightshotPartitionRes, err := client.ApiCallSimple("set-lightshot-partition", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addLightshotPartitionRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addLightshotPartitionRes.Success {
            errMsg = addLightshotPartitionRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addLightshotPartitionRes.GetData()
        }

        debugLogOperation(
            "lightshot-partition",        // resource type
            "create",                       // operation
            "set-lightshot-partition",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add lightshot-partition: %v", err)
    }
    if !addLightshotPartitionRes.Success {
        if addLightshotPartitionRes.ErrorMsg != "" {
            return fmt.Errorf(addLightshotPartitionRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-lightshot-partition", addLightshotPartitionRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-lightshot-partition task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("set-lightshot-partition task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    d.SetId(fmt.Sprintf("lightshot-partition-" + acctest.RandString(10)))
    return readGaiaLightshotPartition(d, m)
}

func readGaiaLightshotPartition(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showLightshotPartitionRes, err := client.ApiCallSimple("show-lightshot-partition", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showLightshotPartitionRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showLightshotPartitionRes.Success {
            errMsg = showLightshotPartitionRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showLightshotPartitionRes.GetData()
        }

        debugLogOperation(
            "lightshot-partition",        // resource type
            "read",                       // operation
            "show-lightshot-partition",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show lightshot-partition: %v", err)
    }
    if !showLightshotPartitionRes.Success {
        if data := showLightshotPartitionRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showLightshotPartitionRes.ErrorMsg)
    }

    lightshotPartition := showLightshotPartitionRes.GetData()

    log.Println("Read LightshotPartition - Show JSON = ", lightshotPartition)

    if v, exists := lightshotPartition["size"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("size", int(f))
        }
    }
    if v, exists := lightshotPartition["used"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("used", int(f))
        }
    }
    if v, exists := lightshotPartition["available"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("available", int(f))
        }
    }
    if v, exists := lightshotPartition["mount-on"]; exists {
        d.Set("mount_on", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaLightshotPartition(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("size"); ok {
        payload["size"] = v.(int)
    }

    setLightshotPartitionRes, err := client.ApiCallSimple("set-lightshot-partition", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setLightshotPartitionRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setLightshotPartitionRes.Success {
            errMsg = setLightshotPartitionRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setLightshotPartitionRes.GetData()
        }

        debugLogOperation(
            "lightshot-partition",        // resource type
            "update",                       // operation
            "set-lightshot-partition",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set lightshot-partition: %v", err)
    }
    if !setLightshotPartitionRes.Success {
        return fmt.Errorf(setLightshotPartitionRes.ErrorMsg)
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "set-lightshot-partition", setLightshotPartitionRes, true, 0)
    if err != nil {
        return fmt.Errorf("set-lightshot-partition task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        errMsg := taskRes.Message
        if errMsg == "" {
            errMsg = fmt.Sprintf("set-lightshot-partition task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(errMsg)
    }

    return readGaiaLightshotPartition(d, m)
}

func deleteGaiaLightshotPartition(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    