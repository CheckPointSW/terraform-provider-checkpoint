package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

    "context"
    
)
func resourceGaiaLightshot() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaLightshot,
        Read:   readGaiaLightshot,
        Delete: deleteGaiaLightshot,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Name of lightshot to add to lightshots list`,
            },
            "lightshot": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "size": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "date": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "result": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: "Either the task-id, or a JSON string of the object inside task-details",
            },
            "task_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: "The task ID of the async operation",
            },
        },
    }
}

func createGaiaLightshot(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    log.Println("Create Lightshot - Map = ", payload)

    addLightshotRes, err := client.ApiCall("add-lightshot", payload, client.GetSessionID(), false, client.IsProxyUsed())
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addLightshotRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addLightshotRes.Success {
            errMsg = addLightshotRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addLightshotRes.GetData()
        }

        debugLogOperation(
            "lightshot",        // resource type
            "create",                       // operation
            "add-lightshot",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add lightshot: %v", err)
    }
    if !addLightshotRes.Success {
        if addLightshotRes.ErrorMsg != "" {
            return fmt.Errorf(addLightshotRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    res, err := HandleTaskCreate(context.Background(), client, "add-lightshot", addLightshotRes, true, 0)
    if err != nil {
        // transport/SDK error
        return fmt.Errorf("add-lightshot task polling failed: %v", err)
    }
    if !res.IsSuccess() {
        // business failure
        return fmt.Errorf("add-lightshot task %s %s: %s", res.TaskID, res.Status, res.Message)
    }

    // Persist to TF state
    _ = d.Set("task_id", res.TaskID)
    // If Message is pretty JSON, store as-is; otherwise store task-id
    out := res.Message
    if strings.TrimSpace(out) == "" { out = res.TaskID }
    _ = d.Set("result", out)

    // Use task-id as stable ID
    d.SetId(fmt.Sprintf("lightshot-" + acctest.RandString(10)))
    return readGaiaLightshot(d, m)
}

func readGaiaLightshot(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	ensureDebugServerFromClient(client)

	payload := map[string]interface{}{}

	showLightshotRes, err := client.ApiCallSimple("show-lightshots", payload)
	// DEBUG: generic logger
	if resourceDebugEnabled(d) {
		success := err == nil && showLightshotRes.Success
		errMsg := ""
		if err != nil {
			errMsg = err.Error()
		} else if !showLightshotRes.Success {
			errMsg = showLightshotRes.ErrorMsg
		}

		var respData map[string]interface{}
		if err == nil {
			respData = showLightshotRes.GetData()
		}

		debugLogOperation(
			"lightshot",        // resource type
			"show",                       // operation
			"show-lightshots",         // API call name
			payload,                        // request payload
			respData,                       // response data (if any)
			success,
			errMsg,
		)
	}
	if err != nil {
		return fmt.Errorf("Failed to show lightshot: %v", err)
	}
	if !showLightshotRes.Success {
		if data := showLightshotRes.GetData(); data != nil {
			if code, exists := data["code"]; exists {
				if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
					d.SetId("")
					return nil
				}
			}
		}
		return fmt.Errorf(showLightshotRes.ErrorMsg)
	}

	
	lightshot := showLightshotRes.GetData()

	log.Println("Read Lightshot - Show JSON = ", lightshot)

        if items, ok := lightshot["lightshots"].([]interface{}); ok {
                nameVal := ""
                if n, ok := d.GetOk("name"); ok {
                        nameVal = n.(string)
                }
                for _, item := range items {
                        if inner, ok := item.(map[string]interface{}); ok {
                                if fmt.Sprintf("%v", inner["name"]) == nameVal {
                                        d.Set("lightshot", []interface{}{map[string]interface{}{
                                                "name":        fmt.Sprintf("%v", inner["name"]),
                                                "description": fmt.Sprintf("%v", inner["description"]),
                                                "size":        fmt.Sprintf("%v", inner["size"]),
                                                "date":        fmt.Sprintf("%v", inner["date"]),
                                        }})
                                        break
                                }
                        }
                }
        }

        d.SetId(d.Id())
        return nil
}


func deleteGaiaLightshot(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteLightshotRes, err := client.ApiCallSimple("delete-lightshot", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteLightshotRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteLightshotRes.Success {
            errMsg = deleteLightshotRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteLightshotRes.GetData()
        }

        debugLogOperation(
            "lightshot",        // resource type
            "delete",                       // operation
            "delete-lightshot",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete lightshot: %v", err)
    }
    if !deleteLightshotRes.Success {
        return fmt.Errorf(deleteLightshotRes.ErrorMsg)
    }

    res, err := HandleTaskCreate(context.Background(), client, "delete-lightshot", deleteLightshotRes, true, 0)
    if err != nil {
        // transport/SDK error
        return fmt.Errorf("delete-lightshot task polling failed: %v", err)
    }
    if !res.IsSuccess() {
        // business failure
        return fmt.Errorf("delete-lightshot task %s %s: %s", res.TaskID, res.Status, res.Message)
    }

    // Persist to TF state
    _ = d.Set("task_id", res.TaskID)
    // If Message is pretty JSON, store as-is; otherwise store task-id
    out := res.Message
    if strings.TrimSpace(out) == "" { out = res.TaskID }
    _ = d.Set("result", out)

    // Use task-id as stable ID
    d.SetId("")
    return nil
}

