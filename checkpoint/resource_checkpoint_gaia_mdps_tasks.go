package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaMdpsTasks() *schema.Resource {
    return &schema.Resource{
        Create: createGaiaMdpsTasks,
        Read:   readGaiaMdpsTasks,
        Update: updateGaiaMdpsTasks,
        Delete: deleteGaiaMdpsTasks,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "external_address": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `External address to communicate with via the Management plane`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "os_service": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `OS Service to run on Management Plane, see 'chkconfig --list' (R82 and below) or 'systemctl list-units --type=service' (R82.10 and above)`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "os_process": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `OS Process to run on Management Plane (see <a href='https://support.checkpoint.com/results/sk/sk97638'>sk97638</a> for more information).`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "cp_port_protocol": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `Check Point Port and Protocol to use on the Management plane (see <a href='https://support.checkpoint.com/results/sk/sk52421'>sk52421</a> for more information).`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "port": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Port number`,
                        },
                        "protocol": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Protocol type`,
                        },
                    },
                },
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

// buildCpPortProtocolArray converts Terraform ResourceData list to API objects.
func buildCpPortProtocolArray(d *schema.ResourceData) []interface{} {
    v := d.Get("cp_port_protocol")
    if len(v.([]interface{})) == 0 {
        return nil
    }
    cpportprotocolList := v.([]interface{})
    cpportprotocolArray := make([]interface{}, 0, len(cpportprotocolList))
    for i := range cpportprotocolList {
        itemMap := make(map[string]interface{})
        if v, ok := d.GetOk(fmt.Sprintf("cp_port_protocol.%d.port", i)); ok {
            itemMap["port"] = v.(int)
        }
        if v, ok := d.GetOk(fmt.Sprintf("cp_port_protocol.%d.protocol", i)); ok {
            itemMap["protocol"] = v.(string)
        }
        if len(itemMap) > 0 {
            cpportprotocolArray = append(cpportprotocolArray, itemMap)
        }
    }
    return cpportprotocolArray
}

func createGaiaMdpsTasks(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    // set-mdps-tasks uses add/remove semantics for list fields, not flat arrays.
    if v := d.Get("external_address"); len(v.([]interface{})) > 0 {
        payload["external-address"] = map[string]interface{}{"add": v.([]interface{})}
    }

    if v := d.Get("os_service"); len(v.([]interface{})) > 0 {
        payload["os-service"] = map[string]interface{}{"add": v.([]interface{})}
    }

    if v := d.Get("os_process"); len(v.([]interface{})) > 0 {
        payload["os-process"] = map[string]interface{}{"add": v.([]interface{})}
    }

    if arr := buildCpPortProtocolArray(d); len(arr) > 0 {
        payload["cp-port-protocol"] = map[string]interface{}{"add": arr}
    }

    log.Println("Create MdpsTasks - Map = ", payload)

    addMdpsTasksRes, err := client.ApiCallSimple("set-mdps-tasks", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addMdpsTasksRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addMdpsTasksRes.Success {
            errMsg = addMdpsTasksRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addMdpsTasksRes.GetData()
        }

        debugLogOperation(
            "mdps-tasks",        // resource type
            "create",                       // operation
            "set-mdps-tasks",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add mdps-tasks: %v", err)
    }
    if !addMdpsTasksRes.Success {
        if addMdpsTasksRes.ErrorMsg != "" {
            return fmt.Errorf(addMdpsTasksRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("mdps-tasks-" + acctest.RandString(10)))
    return readGaiaMdpsTasks(d, m)
}

func readGaiaMdpsTasks(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showMdpsTasksRes, err := client.ApiCallSimple("show-mdps-tasks", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showMdpsTasksRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showMdpsTasksRes.Success {
            errMsg = showMdpsTasksRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showMdpsTasksRes.GetData()
        }

        debugLogOperation(
            "mdps-tasks",        // resource type
            "read",                       // operation
            "show-mdps-tasks",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show mdps-tasks: %v", err)
    }
    if !showMdpsTasksRes.Success {
        if data := showMdpsTasksRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showMdpsTasksRes.ErrorMsg)
    }

    mdpsTasks := showMdpsTasksRes.GetData()

    log.Println("Read MdpsTasks - Show JSON = ", mdpsTasks)

    if v, exists := mdpsTasks["external-address"]; exists {
        d.Set("external_address", v.([]interface{}))
    }
    if v, exists := mdpsTasks["os-service"]; exists {
        d.Set("os_service", v.([]interface{}))
    }
    if v, exists := mdpsTasks["os-process"]; exists {
        d.Set("os_process", v.([]interface{}))
    }
    if v, exists := mdpsTasks["cp-port-protocol"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    port := 0
                    if f, ok := m["port"].(float64); ok {
                        port = int(f)
                    }
                    out = append(out, map[string]interface{}{
                        "port":     port,
                        "protocol": fmt.Sprintf("%v", m["protocol"]),
                    })
                }
            }
            d.Set("cp_port_protocol", out)
        }
    }
    if v, exists := mdpsTasks["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaMdpsTasks(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    // For string-list fields, compute add/remove diff from old and new state.
    for _, field := range []string{"external_address", "os_service", "os_process"} {
        if !d.HasChange(field) {
            continue
        }
        old, new := d.GetChange(field)
        oldList := old.([]interface{})
        newList := new.([]interface{})

        apiKey := strings.ReplaceAll(field, "_", "-")
        op := map[string]interface{}{}
        if len(oldList) > 0 {
            op["remove"] = oldList
        }
        if len(newList) > 0 {
            op["add"] = newList
        }
        if len(op) > 0 {
            payload[apiKey] = op
        }
    }

    // For cp_port_protocol, remove old items and add new items.
    if d.HasChange("cp_port_protocol") {
        oldRaw, _ := d.GetChange("cp_port_protocol")
        oldList := oldRaw.([]interface{})
        oldArr := make([]interface{}, 0, len(oldList))
        for _, item := range oldList {
            if m, ok := item.(map[string]interface{}); ok {
                itemMap := make(map[string]interface{})
                if v, ok := m["port"]; ok {
                    itemMap["port"] = v
                }
                if v, ok := m["protocol"]; ok {
                    itemMap["protocol"] = v
                }
                if len(itemMap) > 0 {
                    oldArr = append(oldArr, itemMap)
                }
            }
        }
        newArr := buildCpPortProtocolArray(d)

        op := map[string]interface{}{}
        if len(oldArr) > 0 {
            op["remove"] = oldArr
        }
        if len(newArr) > 0 {
            op["add"] = newArr
        }
        if len(op) > 0 {
            payload["cp-port-protocol"] = op
        }
    }

    if len(payload) == 0 {
        return readGaiaMdpsTasks(d, m)
    }

    setMdpsTasksRes, err := client.ApiCallSimple("set-mdps-tasks", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setMdpsTasksRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setMdpsTasksRes.Success {
            errMsg = setMdpsTasksRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setMdpsTasksRes.GetData()
        }

        debugLogOperation(
            "mdps-tasks",        // resource type
            "update",                       // operation
            "set-mdps-tasks",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set mdps-tasks: %v", err)
    }
    if !setMdpsTasksRes.Success {
        return fmt.Errorf(setMdpsTasksRes.ErrorMsg)
    }

    return readGaiaMdpsTasks(d, m)
}

func deleteGaiaMdpsTasks(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

