package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaDns() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaDns,
        Read:   readGaiaDns,
        Update: updateGaiaDns,
        Delete: deleteGaiaDns,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "primary": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Use empty-string in order to remove the setting`,
            },
            "secondary": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Use empty-string in order to remove the setting`,
            },
            "tertiary": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Use empty-string in order to remove the setting`,
            },
            "suffix": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Use empty-string in order to remove the setting`,
            },
            "forwarding_domains": {
                Type:        schema.TypeSet,
                Optional:    true,
                Description: `DNS proxy forwarding domains`,
                Set: func(v interface{}) int { return schema.HashString(v.(map[string]interface{})["suffix"].(string)) },
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "primary": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "secondary": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "tertiary": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                        "suffix": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
            "listening_interfaces": {
                Type:        schema.TypeSet,
                Optional:    true,
                Computed:    true,
                Description: `DNS proxy listening interfaces`,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaDns(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("primary"); ok {
        payload["primary"] = v.(string)
    }

    if v, ok := d.GetOk("secondary"); ok {
        payload["secondary"] = v.(string)
    }

    if v, ok := d.GetOk("tertiary"); ok {
        payload["tertiary"] = v.(string)
    }

    if v, ok := d.GetOk("suffix"); ok {
        payload["suffix"] = v.(string)
    }

    if v := d.Get("forwarding_domains"); len(v.(*schema.Set).List()) > 0 {
        payload["forwarding-domains"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("listening_interfaces"); ok {
        items := v.(*schema.Set).List()
        if len(items) == 1 && items[0].(string) == "all" {
            payload["listening-interfaces"] = "all"
        } else if len(items) > 0 {
            names := make([]string, len(items))
            for i, item := range items {
                names[i] = item.(string)
            }
            payload["listening-interfaces"] = map[string]interface{}{"add": names}
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create Dns - Map = ", payload)

    addDnsRes, err := client.ApiCallSimple("set-dns", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addDnsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addDnsRes.Success {
            errMsg = addDnsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addDnsRes.GetData()
        }

        debugLogOperation(
            "dns",        // resource type
            "create",                       // operation
            "set-dns",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add dns: %v", err)
    }
    if !addDnsRes.Success {
        if addDnsRes.ErrorMsg != "" {
            return fmt.Errorf(addDnsRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("dns-" + acctest.RandString(10)))
    return readGaiaDns(d, m)
}

func readGaiaDns(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showDnsRes, err := client.ApiCallSimple("show-dns", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showDnsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showDnsRes.Success {
            errMsg = showDnsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showDnsRes.GetData()
        }

        debugLogOperation(
            "dns",        // resource type
            "read",                       // operation
            "show-dns",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show dns: %v", err)
    }
    if !showDnsRes.Success {
        if data := showDnsRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showDnsRes.ErrorMsg)
    }

    dns := showDnsRes.GetData()

    log.Println("Read Dns - Show JSON = ", dns)

    if v, exists := dns["primary"]; exists {
        d.Set("primary", fmt.Sprintf("%v", v))
    }
    if v, exists := dns["secondary"]; exists {
        d.Set("secondary", fmt.Sprintf("%v", v))
    }
    if v, exists := dns["tertiary"]; exists {
        d.Set("tertiary", fmt.Sprintf("%v", v))
    }
    if v, exists := dns["suffix"]; exists {
        d.Set("suffix", fmt.Sprintf("%v", v))
    }
    if v, exists := dns["forwarding-domains"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    out = append(out, map[string]interface{}{
                        "primary":   fmt.Sprintf("%v", m["primary"]),
                        "secondary": fmt.Sprintf("%v", m["secondary"]),
                        "tertiary":  fmt.Sprintf("%v", m["tertiary"]),
                        "suffix":    fmt.Sprintf("%v", m["suffix"]),
                    })
                }
            }
            d.Set("forwarding_domains", out)
        }
    }
    if v, exists := dns["listening-interfaces"]; exists {
        switch val := v.(type) {
        case []interface{}:
            d.Set("listening_interfaces", val)
        case string:
            d.Set("listening_interfaces", []interface{}{val})
        default:
            d.Set("listening_interfaces", []interface{}{})
        }
    }
    if v, exists := dns["virtual-system-id"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("virtual_system_id", int(f))
        }
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaDns(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("primary"); ok {
        payload["primary"] = v.(string)
    }

    if v, ok := d.GetOk("secondary"); ok {
        payload["secondary"] = v.(string)
    }

    if v, ok := d.GetOk("tertiary"); ok {
        payload["tertiary"] = v.(string)
    }

    if v, ok := d.GetOk("suffix"); ok {
        payload["suffix"] = v.(string)
    }

    if v := d.Get("forwarding_domains"); len(v.(*schema.Set).List()) > 0 {
        payload["forwarding-domains"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("listening_interfaces"); ok {
        items := v.(*schema.Set).List()
        if len(items) == 1 && items[0].(string) == "all" {
            payload["listening-interfaces"] = "all"
        } else if len(items) > 0 {
            names := make([]string, len(items))
            for i, item := range items {
                names[i] = item.(string)
            }
            payload["listening-interfaces"] = map[string]interface{}{"add": names}
        }
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    setDnsRes, err := client.ApiCallSimple("set-dns", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setDnsRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setDnsRes.Success {
            errMsg = setDnsRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setDnsRes.GetData()
        }

        debugLogOperation(
            "dns",        // resource type
            "update",                       // operation
            "set-dns",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set dns: %v", err)
    }
    if !setDnsRes.Success {
        return fmt.Errorf(setDnsRes.ErrorMsg)
    }

    return readGaiaDns(d, m)
}

func deleteGaiaDns(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    