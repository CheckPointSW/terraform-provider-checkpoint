package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaPutFile() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaPutFile,
        Read:   readGaiaPutFile,
        Delete: deleteGaiaPutFile,
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
                Optional:    true,
                ForceNew:    true,
                Description: `File content as string, for new line use \n`,
            },
            "override": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: `overwrite file content`,
            },
            "group_ownership": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Group file owner`,
            },
            "user_ownership": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `User file owner`,
            },
            "permissions": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `File permissions, provided in octal mode`,
            },
            "virtual_system_id": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaPutFile(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("file_name"); ok {
        payload["file-name"] = v.(string)
    }

    if v, ok := d.GetOk("text_content"); ok {
        payload["text-content"] = v.(string)
    }

    if v, ok := d.GetOkExists("override"); ok {
        payload["override"] = v.(bool)
    }

    if v, ok := d.GetOk("group_ownership"); ok {
        payload["group-ownership"] = v.(string)
    }

    if v, ok := d.GetOk("user_ownership"); ok {
        payload["user-ownership"] = v.(string)
    }

    if v, ok := d.GetOk("permissions"); ok {
        payload["permissions"] = v.(int)
    }

    log.Println("Execute put-file - Payload = ", payload)

    GaiaPutFileRes, err := client.ApiCallSimple("put-file", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaPutFileRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaPutFileRes.Success {
            errMsg = GaiaPutFileRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaPutFileRes.GetData()
        }

        debugLogOperation(
            "put-file",        // resource type
            "command",                       // operation
            "put-file",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute put-file: %v", err)
    }
    if !GaiaPutFileRes.Success {
        if GaiaPutFileRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaPutFileRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaPutFileRes.GetData()
    if v, exists := _respData["virtual-system-id"]; exists {
        d.Set("virtual_system_id", toString(v))
    }


    d.SetId(fmt.Sprintf("put-file-" + acctest.RandString(10)))
    return nil
}

func readGaiaPutFile(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("file_name"); ok {
        payload["file-name"] = v.(string)
    }

    if v, ok := d.GetOk("text_content"); ok {
        payload["text-content"] = v.(string)
    }

    if v, ok := d.GetOkExists("override"); ok {
        payload["override"] = v.(bool)
    }

    if v, ok := d.GetOk("group_ownership"); ok {
        payload["group-ownership"] = v.(string)
    }

    if v, ok := d.GetOk("user_ownership"); ok {
        payload["user-ownership"] = v.(string)
    }

    if v, ok := d.GetOk("permissions"); ok {
        payload["permissions"] = v.(int)
    }

    log.Println("Execute put-file - Payload = ", payload)

    GaiaPutFileRes, err := client.ApiCallSimple("put-file", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && GaiaPutFileRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !GaiaPutFileRes.Success {
            errMsg = GaiaPutFileRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = GaiaPutFileRes.GetData()
        }

        debugLogOperation(
            "put-file",        // resource type
            "read",                       // operation
            "put-file",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute put-file: %v", err)
    }
    if !GaiaPutFileRes.Success {
        if GaiaPutFileRes.ErrorMsg != "" {
            return fmt.Errorf(GaiaPutFileRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    _respData := GaiaPutFileRes.GetData()
    if v, exists := _respData["virtual-system-id"]; exists {
        d.Set("virtual_system_id", toString(v))
    }
    return nil
}

func deleteGaiaPutFile(d *schema.ResourceData, m interface{}) error {
    d.SetId("")
    return nil
}

