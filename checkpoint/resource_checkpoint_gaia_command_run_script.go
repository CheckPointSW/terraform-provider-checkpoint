package checkpoint

import (
        "context"
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaRunScript() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaRunScript,
        Read:   readGaiaRunScript,
        Delete: deleteGaiaRunScript,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Script description`,
            },
            "script": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Script body`,
            },
            "args": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Script arguments, separated by space character. Note: don't send sensitive data on this parameter.`,
            },
            "environment_variables": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Define environment variables to be used in the script, it's better to send sensitive data on environment variables since it's not stored.`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Variable's name`,
                        },
                        "value": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            ForceNew:    true,
                            Description: `Variable's value`,
                        },
                    },
                },
            },
            "return_value": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
            "output": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "error": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaRunScript(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("description"); ok {
        payload["description"] = v.(string)
    }

    if v, ok := d.GetOk("script"); ok {
        payload["script"] = v.(string)
    }

    if v, ok := d.GetOk("args"); ok {
        payload["args"] = v.(string)
    }

    if v := d.Get("environment_variables"); len(v.([]interface{})) > 0 {
        environmentvariablesList := v.([]interface{})
        environmentvariablesArray := make([]interface{}, 0, len(environmentvariablesList))
        for i := range environmentvariablesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("environment_variables.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("environment_variables.%d.value", i)); ok {
                itemMap["value"] = v.(string)
            }
            if len(itemMap) > 0 {
                environmentvariablesArray = append(environmentvariablesArray, itemMap)
            }
        }
        if len(environmentvariablesArray) > 0 {
            payload["environment-variables"] = environmentvariablesArray
        }
    }

    log.Println("Execute run-script - Payload = ", payload)

    GaiaRunScriptRes, err := client.ApiCallSimple("run-script", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaRunScriptRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaRunScriptRes.Success {
            errMsg = GaiaRunScriptRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaRunScriptRes.GetData()
        }

        debugLogOperation(
            "run-script",        // resource type
            "command",                       // operation
            "run-script",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute run-script: %v", err)
    }
    if !GaiaRunScriptRes.Success {
        if GaiaRunScriptRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaRunScriptRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    taskRes, err := HandleTaskCreate(context.Background(), client, "run-script", GaiaRunScriptRes, true, 0)
    if err != nil {
        return fmt.Errorf("run-script task polling failed: %v", err)
    }
    if !taskRes.IsSuccess() {
        msg := taskRes.Message
        if msg == "" {
            msg = fmt.Sprintf("run-script task %s ended with status: %s", taskRes.TaskID, taskRes.Status)
        }
        return fmt.Errorf(msg)
    }

    _taskDetailsRes, _tdErr := client.ApiCallSimple("show-task", map[string]interface{}{"task-id": taskRes.TaskID})
    var _respData map[string]interface{}
    if _tdErr == nil && _taskDetailsRes.Success {
        _td := _taskDetailsRes.GetData()
        if _tasks, _ok := _td["tasks"].([]interface{}); _ok && len(_tasks) > 0 {
            if _task, _ok := _tasks[0].(map[string]interface{}); _ok {
                if _details, _ok := _task["task-details"].([]interface{}); _ok && len(_details) > 0 {
                    if _d0, _ok := _details[0].(map[string]interface{}); _ok {
                        _respData = _d0
                    }
                }
            }
        }
    }
    if _respData == nil {
        _respData = GaiaRunScriptRes.GetData()
    }
    if v, exists := _respData["return-value"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("return_value", int(f))
        }
    }
    if v, exists := _respData["output"]; exists {
        d.Set("output", toString(v))
    }
    if v, exists := _respData["error"]; exists {
        d.Set("error", toString(v))
    }


    d.SetId(fmt.Sprintf("run-script-" + acctest.RandString(10)))
    return nil
}

func readGaiaRunScript(d *schema.ResourceData, m interface{}) error {
    return nil
}

func deleteGaiaRunScript(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

