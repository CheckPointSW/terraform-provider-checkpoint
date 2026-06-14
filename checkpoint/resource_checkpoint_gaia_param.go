package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaParam() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaParam,
        Read:   readGaiaParam,
        Update: updateGaiaParam,
        Delete: deleteGaiaParam,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "param_list": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `List of parameters to be set`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "param_path": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Parameter's full path`,
                        },
                        "value": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Parameter's value. it can be sanitized-ascii, int or boolean or object. can not send it with use-default in the same request`,
                        },
                        "comments": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Comments to be added to the parameter, length is limited to 256 characters`,
                        },
                        "use_default": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Set the parameter back to its default value. can not send it with value in the same request`,
                        },
                        "volatile": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Set parameter with the value specified untill the next reboot`,
                        },
                        "virtual_system_id": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `VSX vs-id which present the context id, can be 'all' or a string represent spesific VS Ids, for example: '2,4,7-10'`,
                        },
                    },
                },
            },
            "virtual_system_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `VSX vs-id which present the context switch`,
            },
            "dry_run": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `run the request without saving to data base, return only with the changed values.`,
            },
            "use_regex": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `indicates that parameter param_path includes *`,
            },
            "param_path": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Paramter full path or prefix`,
            },
            "filter_by": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `show parameters by filtering type in the following format: <fiter_type>=true when fiter_type can be one of (modified,with-comments,not-default), for example modified=true`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaParam(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("param_list"); len(v.([]interface{})) > 0 {
        paramlistList := v.([]interface{})
        paramlistArray := make([]interface{}, 0, len(paramlistList))
        for i := range paramlistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("param_list.%d.param_path", i)); ok {
                itemMap["param-path"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("param_list.%d.value", i)); ok {
                itemMap["value"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("param_list.%d.comments", i)); ok {
                itemMap["comments"] = v.(string)
            }
            if v := d.Get(fmt.Sprintf("param_list.%d.use_default", i)).(bool); v {
                itemMap["use-default"] = v
            }
            if v := d.Get(fmt.Sprintf("param_list.%d.volatile", i)).(bool); v {
                itemMap["volatile"] = v
            }
            if v, ok := d.GetOk(fmt.Sprintf("param_list.%d.virtual_system_id", i)); ok {
                itemMap["virtual-system-id"] = v.(string)
            }
            if len(itemMap) > 0 {
                paramlistArray = append(paramlistArray, itemMap)
            }
        }
        if len(paramlistArray) > 0 {
            payload["param-list"] = paramlistArray
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(string)
    }

    if v, ok := d.GetOkExists("dry_run"); ok {
        payload["dry-run"] = v.(bool)
    }

    if v, ok := d.GetOkExists("use_regex"); ok {
        payload["use-regex"] = v.(bool)
    }

    log.Println("Create Param - Map = ", payload)

    addParamRes, err := client.ApiCallSimple("set-param", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addParamRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addParamRes.Success {
            errMsg = addParamRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addParamRes.GetData()
        }

        debugLogOperation(
            "param",        // resource type
            "create",                       // operation
            "set-param",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add param: %v", err)
    }
    if !addParamRes.Success {
        if addParamRes.ErrorMsg != "" {
            return fmt.Errorf(addParamRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("param-" + acctest.RandString(10)))
    return readGaiaParam(d, m)
}

func readGaiaParam(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(string)
    }

    if v, ok := d.GetOk("param_path"); ok {
        payload["param-path"] = v.(string)
    }

    if v, ok := d.GetOk("filter_by"); ok {
        payload["filter-by"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    // Build show-param payload from param_list when top-level param_path is not set.
    // show-param requires a top-level param-path string; param-list is not accepted.
    // Use the first param_path from param_list state as the lookup key.
    if _, ok := payload["param-path"]; !ok {
        if v := d.Get("param_list"); len(v.([]interface{})) > 0 {
            for _, item := range v.([]interface{}) {
                if m, ok := item.(map[string]interface{}); ok {
                    if p, ok := m["param_path"].(string); ok && p != "" {
                        payload["param-path"] = p
                        break
                    }
                }
            }
        }
    }

    showParamRes, err := client.ApiCallSimple("show-param", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showParamRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showParamRes.Success {
            errMsg = showParamRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showParamRes.GetData()
        }

        debugLogOperation(
            "param",        // resource type
            "read",                       // operation
            "show-param",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show param: %v", err)
    }
    if !showParamRes.Success {
        if data := showParamRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showParamRes.ErrorMsg)
    }

    param := showParamRes.GetData()

    log.Println("Read Param - Show JSON = ", param)

    if v, exists := param["param-list"]; exists {
        d.Set("param_list", v.([]interface{}))
    }
    if v, exists := param["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaParam(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("param_list"); len(v.([]interface{})) > 0 {
        paramlistList := v.([]interface{})
        paramlistArray := make([]interface{}, 0, len(paramlistList))
        for i := range paramlistList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("param_list.%d.param_path", i)); ok {
                itemMap["param-path"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("param_list.%d.value", i)); ok {
                itemMap["value"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("param_list.%d.comments", i)); ok {
                itemMap["comments"] = v.(string)
            }
            if v := d.Get(fmt.Sprintf("param_list.%d.use_default", i)).(bool); v {
                itemMap["use-default"] = v
            }
            if v := d.Get(fmt.Sprintf("param_list.%d.volatile", i)).(bool); v {
                itemMap["volatile"] = v
            }
            if v, ok := d.GetOk(fmt.Sprintf("param_list.%d.virtual_system_id", i)); ok {
                itemMap["virtual-system-id"] = v.(string)
            }
            if len(itemMap) > 0 {
                paramlistArray = append(paramlistArray, itemMap)
            }
        }
        if len(paramlistArray) > 0 {
            payload["param-list"] = paramlistArray
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(string)
    }

    if v, ok := d.GetOkExists("dry_run"); ok {
        payload["dry-run"] = v.(bool)
    }

    if v, ok := d.GetOkExists("use_regex"); ok {
        payload["use-regex"] = v.(bool)
    }

    setParamRes, err := client.ApiCallSimple("set-param", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setParamRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setParamRes.Success {
            errMsg = setParamRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setParamRes.GetData()
        }

        debugLogOperation(
            "param",        // resource type
            "update",                       // operation
            "set-param",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set param: %v", err)
    }
    if !setParamRes.Success {
        return fmt.Errorf(setParamRes.ErrorMsg)
    }

    return readGaiaParam(d, m)
}

func deleteGaiaParam(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    