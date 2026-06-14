package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"
    "strconv"

)
func resourceGaiaPbrTable() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaPbrTable,
        Read:   readGaiaPbrTable,
        Update: updateGaiaPbrTable,
        Delete: deleteGaiaPbrTable,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "table": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `Name of PBR Table`,
            },
            "static_routes": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `List of static routes configured on PBR Table`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `IPv4 address of route`,
                        },
                        "mask_length": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `Mask length of route`,
                        },
                        "type": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Type of next-hop. Possible values: blackhole, gateway, reject`,
                        },
                        "next_hop": {
                            Type:        schema.TypeList,
                            Optional:    true,
                            Description: `Static next-hop. Contains a list of next-hop gateways. Each gateway is formatted in the following manner:{\"gateway\": IP address or logical name, \"priority\": default or integer 1-8}`,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "gateway": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `IP address or logical name for the static next-hop gateway`,
                                    },
                                    "priority": {
                                        Type:        schema.TypeString,
                                        Optional:    true,
                                        Description: `Priority defines which gateway to select as the next-hop. The lower the priority, the higher the preference. Possible values: default or integer 1-8`,
                                    },
                                },
                            },
                        },
                        "ping": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Configures ping monitoring of the given IPv4 static route. Possible values: true, false`,
                        },
                    },
                },
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "static_route_limit": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `The maximum number of configured static-routes to show in response`,
            },
            "static_route_offset": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `The number of configured static-routes to initially skip`,
            },
            "static_route_order": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Sorts the static-routes by address in either ascending or descending order.`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaPbrTable(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("table"); ok {
        payload["table"] = v.(string)
    }

    if v := d.Get("static_routes"); len(v.([]interface{})) > 0 {
        staticroutesList := v.([]interface{})
        staticroutesArray := make([]interface{}, 0, len(staticroutesList))
        for i := range staticroutesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("static_routes.%d.address", i)); ok {
                itemMap["address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("static_routes.%d.mask_length", i)); ok {
                itemMap["mask-length"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("static_routes.%d.type", i)); ok {
                itemMap["type"] = v.(string)
            }
            if sv := d.Get(fmt.Sprintf("static_routes.%d.next_hop", i)); len(sv.([]interface{})) > 0 {
                next_hopList := sv.([]interface{})
                next_hopArr := make([]interface{}, 0, len(next_hopList))
                for j := range next_hopList {
                    innerMap := make(map[string]interface{})
                    if iv, ok := d.GetOk(fmt.Sprintf("static_routes.%d.next_hop.%d.gateway", i, j)); ok {
                        innerMap["gateway"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("static_routes.%d.next_hop.%d.priority", i, j)); ok {
                        pStr := iv.(string)
                        if pStr != "default" {
                            if pInt, err := strconv.Atoi(pStr); err == nil {
                                innerMap["priority"] = pInt
                            }
                        }
                    }
                    if len(innerMap) > 0 {
                        next_hopArr = append(next_hopArr, innerMap)
                    }
                }
                if len(next_hopArr) > 0 {
                    itemMap["next-hop"] = next_hopArr
                }
            }
            if v := d.Get(fmt.Sprintf("static_routes.%d.ping", i)).(bool); v {
                itemMap["ping"] = v
            }
            if len(itemMap) > 0 {
                staticroutesArray = append(staticroutesArray, itemMap)
            }
        }
        if len(staticroutesArray) > 0 {
            payload["static-routes"] = map[string]interface{}{"add": staticroutesArray}
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create PbrTable - Map = ", payload)

    addPbrTableRes, err := client.ApiCallSimple("set-pbr-table", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addPbrTableRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addPbrTableRes.Success {
            errMsg = addPbrTableRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addPbrTableRes.GetData()
        }

        debugLogOperation(
            "pbr-table",        // resource type
            "create",                       // operation
            "set-pbr-table",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add pbr-table: %v", err)
    }
    if !addPbrTableRes.Success {
        if addPbrTableRes.ErrorMsg != "" {
            return fmt.Errorf(addPbrTableRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("pbr-table-" + acctest.RandString(10)))
    return readGaiaPbrTable(d, m)
}

func readGaiaPbrTable(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("table"); ok {
        payload["table"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("static_route_limit"); ok {
        payload["static-route-limit"] = v.(int)
    }

    if v, ok := d.GetOk("static_route_offset"); ok {
        payload["static-route-offset"] = v.(int)
    }

    if v, ok := d.GetOk("static_route_order"); ok {
        payload["static-route-order"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showPbrTableRes, err := client.ApiCallSimple("show-pbr-table", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showPbrTableRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showPbrTableRes.Success {
            errMsg = showPbrTableRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showPbrTableRes.GetData()
        }

        debugLogOperation(
            "pbr-table",        // resource type
            "read",                       // operation
            "show-pbr-table",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show pbr-table: %v", err)
    }
    if !showPbrTableRes.Success {
        if data := showPbrTableRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showPbrTableRes.ErrorMsg)
    }

    pbrTable := showPbrTableRes.GetData()

    log.Println("Read PbrTable - Show JSON = ", pbrTable)

    if v, exists := pbrTable["table"]; exists {
        d.Set("table", fmt.Sprintf("%v", v))
    }
    if v, exists := pbrTable["from-static-route"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("from_static_route", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("from_static_route", _n)
            }
        }
    }
    if v, exists := pbrTable["to-static-route"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("to_static_route", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("to_static_route", _n)
            }
        }
    }
    if v, exists := pbrTable["total-static-routes"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("total_static_routes", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("total_static_routes", _n)
            }
        }
    }
    if v, exists := pbrTable["static-routes"]; exists {
        d.Set("static_routes", v.([]interface{}))
    }
    if v, exists := pbrTable["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    if v, exists := pbrTable["virtual-system-id"]; exists {
        d.Set("virtual_system_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaPbrTable(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("table"); ok {
        payload["table"] = v.(string)
    }

    if v := d.Get("static_routes"); len(v.([]interface{})) > 0 {
        staticroutesList := v.([]interface{})
        staticroutesArray := make([]interface{}, 0, len(staticroutesList))
        for i := range staticroutesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("static_routes.%d.address", i)); ok {
                itemMap["address"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("static_routes.%d.mask_length", i)); ok {
                itemMap["mask-length"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("static_routes.%d.type", i)); ok {
                itemMap["type"] = v.(string)
            }
            if sv := d.Get(fmt.Sprintf("static_routes.%d.next_hop", i)); len(sv.([]interface{})) > 0 {
                next_hopList := sv.([]interface{})
                next_hopArr := make([]interface{}, 0, len(next_hopList))
                for j := range next_hopList {
                    innerMap := make(map[string]interface{})
                    if iv, ok := d.GetOk(fmt.Sprintf("static_routes.%d.next_hop.%d.gateway", i, j)); ok {
                        innerMap["gateway"] = iv.(string)
                    }
                    if iv, ok := d.GetOk(fmt.Sprintf("static_routes.%d.next_hop.%d.priority", i, j)); ok {
                        pStr := iv.(string)
                        if pStr != "default" {
                            if pInt, err := strconv.Atoi(pStr); err == nil {
                                innerMap["priority"] = pInt
                            }
                        }
                    }
                    if len(innerMap) > 0 {
                        next_hopArr = append(next_hopArr, innerMap)
                    }
                }
                if len(next_hopArr) > 0 {
                    itemMap["next-hop"] = next_hopArr
                }
            }
            if v := d.Get(fmt.Sprintf("static_routes.%d.ping", i)).(bool); v {
                itemMap["ping"] = v
            }
            if len(itemMap) > 0 {
                staticroutesArray = append(staticroutesArray, itemMap)
            }
        }
        if len(staticroutesArray) > 0 {
            payload["static-routes"] = map[string]interface{}{"set": staticroutesArray}
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    setPbrTableRes, err := client.ApiCallSimple("set-pbr-table", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setPbrTableRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setPbrTableRes.Success {
            errMsg = setPbrTableRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setPbrTableRes.GetData()
        }

        debugLogOperation(
            "pbr-table",        // resource type
            "update",                       // operation
            "set-pbr-table",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set pbr-table: %v", err)
    }
    if !setPbrTableRes.Success {
        return fmt.Errorf(setPbrTableRes.ErrorMsg)
    }

    return readGaiaPbrTable(d, m)
}

func deleteGaiaPbrTable(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("table"); ok {
        payload["table"] = v.(string)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    deletePbrTableRes, err := client.ApiCallSimple("delete-pbr-table", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deletePbrTableRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deletePbrTableRes.Success {
            errMsg = deletePbrTableRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deletePbrTableRes.GetData()
        }

        debugLogOperation(
            "pbr-table",        // resource type
            "delete",                       // operation
            "delete-pbr-table",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete pbr-table: %v", err)
    }
    if !deletePbrTableRes.Success {
        return fmt.Errorf(deletePbrTableRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

