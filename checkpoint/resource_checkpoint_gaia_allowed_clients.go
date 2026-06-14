package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaAllowedClients() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaAllowedClients,
        Read:   readGaiaAllowedClients,
        Update: updateGaiaAllowedClients,
        Delete: deleteGaiaAllowedClients,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "allowed_networks": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "subnet": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `The network subnet`,
                        },
                        "mask_length": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `The network mask length`,
                        },
                    },
                },
            },
            "allowed_hosts": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `N/A`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "allowed_any_host": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `N/A`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaAllowedClients(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("allowed_networks"); len(v.([]interface{})) > 0 {
        allowednetworksList := v.([]interface{})
        allowednetworksArray := make([]interface{}, 0, len(allowednetworksList))
        for i := range allowednetworksList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("allowed_networks.%d.subnet", i)); ok {
                itemMap["subnet"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("allowed_networks.%d.mask_length", i)); ok {
                itemMap["mask-length"] = v.(int)
            }
            if len(itemMap) > 0 {
                allowednetworksArray = append(allowednetworksArray, itemMap)
            }
        }
        if len(allowednetworksArray) > 0 {
            payload["allowed-networks"] = allowednetworksArray
        }
    }

    if v := d.Get("allowed_hosts"); len(v.([]interface{})) > 0 {
        payload["allowed-hosts"] = v.([]interface{})
    }

    if v, ok := d.GetOkExists("allowed_any_host"); ok {
        payload["allowed-any-host"] = v.(bool)
    }

    log.Println("Create AllowedClients - Map = ", payload)

    addAllowedClientsRes, err := client.ApiCallSimple("set-allowed-clients", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addAllowedClientsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addAllowedClientsRes.Success {
            errMsg = addAllowedClientsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addAllowedClientsRes.GetData()
        }

        debugLogOperation(
            "allowed-clients",        // resource type
            "create",                       // operation
            "set-allowed-clients",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add allowed-clients: %v", err)
    }
    if !addAllowedClientsRes.Success {
        if addAllowedClientsRes.ErrorMsg != "" {
            return fmt.Errorf(addAllowedClientsRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("allowed-clients-" + acctest.RandString(10)))
    return readGaiaAllowedClients(d, m)
}

func readGaiaAllowedClients(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showAllowedClientsRes, err := client.ApiCallSimple("show-allowed-clients", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showAllowedClientsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showAllowedClientsRes.Success {
            errMsg = showAllowedClientsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showAllowedClientsRes.GetData()
        }

        debugLogOperation(
            "allowed-clients",        // resource type
            "read",                       // operation
            "show-allowed-clients",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show allowed-clients: %v", err)
    }
    if !showAllowedClientsRes.Success {
        if data := showAllowedClientsRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showAllowedClientsRes.ErrorMsg)
    }

    allowedClients := showAllowedClientsRes.GetData()

    log.Println("Read AllowedClients - Show JSON = ", allowedClients)

    if v, exists := allowedClients["allowed-networks"]; exists {
        if raw, ok := v.([]interface{}); ok {
            mapped := make([]interface{}, 0, len(raw))
            for _, item := range raw {
                if m, ok := item.(map[string]interface{}); ok {
                    entry := map[string]interface{}{
                        "subnet": fmt.Sprintf("%v", m["subnet"]),
                    }
                    if f, ok := m["mask-length"].(float64); ok {
                        entry["mask_length"] = int(f)
                    }
                    mapped = append(mapped, entry)
                }
            }
            d.Set("allowed_networks", mapped)
        }
    }
    if v, exists := allowedClients["allowed-hosts"]; exists {
        d.Set("allowed_hosts", v.([]interface{}))
    }
    if v, exists := allowedClients["allowed-any-host"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("allowed_any_host", b)
        } else if s, ok := v.(string); ok {
            d.Set("allowed_any_host", s == "true")
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaAllowedClients(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("allowed_networks"); len(v.([]interface{})) > 0 {
        allowednetworksList := v.([]interface{})
        allowednetworksArray := make([]interface{}, 0, len(allowednetworksList))
        for i := range allowednetworksList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("allowed_networks.%d.subnet", i)); ok {
                itemMap["subnet"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("allowed_networks.%d.mask_length", i)); ok {
                itemMap["mask-length"] = v.(int)
            }
            if len(itemMap) > 0 {
                allowednetworksArray = append(allowednetworksArray, itemMap)
            }
        }
        if len(allowednetworksArray) > 0 {
            payload["allowed-networks"] = allowednetworksArray
        }
    }

    if v := d.Get("allowed_hosts"); len(v.([]interface{})) > 0 {
        payload["allowed-hosts"] = v.([]interface{})
    }

    if v, ok := d.GetOkExists("allowed_any_host"); ok {
        payload["allowed-any-host"] = v.(bool)
    }

    setAllowedClientsRes, err := client.ApiCallSimple("set-allowed-clients", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setAllowedClientsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setAllowedClientsRes.Success {
            errMsg = setAllowedClientsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setAllowedClientsRes.GetData()
        }

        debugLogOperation(
            "allowed-clients",        // resource type
            "update",                       // operation
            "set-allowed-clients",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set allowed-clients: %v", err)
    }
    if !setAllowedClientsRes.Success {
        return fmt.Errorf(setAllowedClientsRes.ErrorMsg)
    }

    return readGaiaAllowedClients(d, m)
}

func deleteGaiaAllowedClients(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    