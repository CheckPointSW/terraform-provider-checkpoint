package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaGetFile() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaGetFile,
        Read:   readGaiaGetFile,
        Delete: deleteGaiaGetFile,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:     true,
                Description: "Enable debugging for this resource only.",
            },
            "file_name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Filename include the desired path. The file will be created in the user home directory if the full path wasn't provided`,
            },
            "text_content": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "group_ownership": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "user_ownership": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "permissions": {
                Type:        schema.TypeInt,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaGetFile(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("file_name"); ok {
        payload["file-name"] = v.(string)
    }

    log.Println("Execute get-file - Payload = ", payload)

    GaiaGetFileRes, err := client.ApiCallSimple("get-file", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaGetFileRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaGetFileRes.Success {
            errMsg = GaiaGetFileRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaGetFileRes.GetData()
        }

        debugLogOperation(
            "get-file",        // resource type
            "command",                       // operation
            "get-file",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute get-file: %v", err)
    }
    if !GaiaGetFileRes.Success {
        if GaiaGetFileRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaGetFileRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaGetFileRes.GetData()
    if v, exists := _respData["text-content"]; exists {
        d.Set("text_content", toString(v))
    }
    if v, exists := _respData["group-ownership"]; exists {
        d.Set("group_ownership", toString(v))
    }
    if v, exists := _respData["user-ownership"]; exists {
        d.Set("user_ownership", toString(v))
    }
    if v, exists := _respData["permissions"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("permissions", int(f))
        }
    }


    d.SetId(fmt.Sprintf("get-file-" + acctest.RandString(10)))
    return nil
}

func readGaiaGetFile(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("file_name"); ok {
        payload["file-name"] = v.(string)
    }

    log.Println("Execute get-file - Payload = ", payload)

    GaiaGetFileRes, err := client.ApiCallSimple("get-file", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaGetFileRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaGetFileRes.Success {
            errMsg = GaiaGetFileRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaGetFileRes.GetData()
        }

        debugLogOperation(
            "get-file",        // resource type
            "read",                       // operation
            "get-file",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute get-file: %v", err)
    }
    if !GaiaGetFileRes.Success {
        if GaiaGetFileRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaGetFileRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaGetFileRes.GetData()
    if v, exists := _respData["text-content"]; exists {
        d.Set("text_content", toString(v))
    }
    if v, exists := _respData["group-ownership"]; exists {
        d.Set("group_ownership", toString(v))
    }
    if v, exists := _respData["user-ownership"]; exists {
        d.Set("user_ownership", toString(v))
    }
    if v, exists := _respData["permissions"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("permissions", int(f))
        }
    }
    return nil
}

func deleteGaiaGetFile(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

