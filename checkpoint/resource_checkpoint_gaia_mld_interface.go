package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaMldInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaMldInterface,
        Read:   readGaiaMldInterface,
        Update: updateGaiaMldInterface,
        Delete: deleteGaiaMldInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                Description: `The name of the MLD interface`,
            },
            "last_listener_query_count": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `The number of queries to send when a listener leaves a group`,
            },
            "last_listener_query_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `The number of seconds between queries when a listener leaves a group`,
            },
            "loss_robustness": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The loss-robustness value`,
            },
            "query_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `The number of seconds between MLD general queries`,
            },
            "query_response_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `The maximum delay time in seconds for hosts to respond to an MLD membership query`,
            },
            "startup_query_count": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `The number of queries sent when MLD starts up`,
            },
            "startup_query_interval": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `The number of seconds between MLD startup queries`,
            },
            "mld_version": {
                Type:        schema.TypeString,
                Optional:    true,
                Computed:    true,
                Description: `The MLD version running`,
            },
            "reset": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `Reset all attributes of this interface to default values`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaMldInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("last_listener_query_count"); ok {
        payload["last-listener-query-count"] = v.(string)
    }

    if v, ok := d.GetOk("last_listener_query_interval"); ok {
        payload["last-listener-query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("loss_robustness"); ok {
        payload["loss-robustness"] = v.(string)
    }

    if v, ok := d.GetOk("query_interval"); ok {
        payload["query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("query_response_interval"); ok {
        payload["query-response-interval"] = v.(string)
    }

    if v, ok := d.GetOk("startup_query_count"); ok {
        payload["startup-query-count"] = v.(string)
    }

    if v, ok := d.GetOk("startup_query_interval"); ok {
        payload["startup-query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("mld_version"); ok {
        payload["mld-version"] = v.(string)
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    log.Println("Create MldInterface - Map = ", payload)

    addMldInterfaceRes, err := client.ApiCallSimple("set-mld-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addMldInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addMldInterfaceRes.Success {
            errMsg = addMldInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addMldInterfaceRes.GetData()
        }

        debugLogOperation(
            "mld-interface",        // resource type
            "create",                       // operation
            "set-mld-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add mld-interface: %v", err)
    }
    if !addMldInterfaceRes.Success {
        if addMldInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addMldInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("mld-interface-" + acctest.RandString(10)))
    return readGaiaMldInterface(d, m)
}

func readGaiaMldInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showMldInterfaceRes, err := client.ApiCallSimple("show-mld-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showMldInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showMldInterfaceRes.Success {
            errMsg = showMldInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showMldInterfaceRes.GetData()
        }

        debugLogOperation(
            "mld-interface",        // resource type
            "read",                       // operation
            "show-mld-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show mld-interface: %v", err)
    }
    if !showMldInterfaceRes.Success {
        if data := showMldInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        if strings.Contains(strings.ToLower(showMldInterfaceRes.ErrorMsg), "failed to get output") {
            // Daemon does not reflect config-db state yet; preserve existing state.
            return nil
        }
        if strings.Contains(strings.ToLower(showMldInterfaceRes.ErrorMsg), "general exception") {
            // Daemon does not reflect config-db state yet; preserve existing state.
            return nil
        }
        return fmt.Errorf(showMldInterfaceRes.ErrorMsg)
    }

    mldInterface := showMldInterfaceRes.GetData()

    log.Println("Read MldInterface - Show JSON = ", mldInterface)

    if v, exists := mldInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterface["mld-version"]; exists {
        d.Set("mld_version", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterface["loss-robustness"]; exists {
        d.Set("loss_robustness", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterface["query-interval"]; exists {
        d.Set("query_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterface["query-response-interval"]; exists {
        d.Set("query_response_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterface["last-listener-query-count"]; exists {
        d.Set("last_listener_query_count", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterface["last-listener-query-interval"]; exists {
        d.Set("last_listener_query_interval", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterface["startup-query-count"]; exists {
        d.Set("startup_query_count", fmt.Sprintf("%v", v))
    }
    if v, exists := mldInterface["startup-query-interval"]; exists {
        d.Set("startup_query_interval", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaMldInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOk("last_listener_query_count"); ok {
        payload["last-listener-query-count"] = v.(string)
    }

    if v, ok := d.GetOk("last_listener_query_interval"); ok {
        payload["last-listener-query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("loss_robustness"); ok {
        payload["loss-robustness"] = v.(string)
    }

    if v, ok := d.GetOk("query_interval"); ok {
        payload["query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("query_response_interval"); ok {
        payload["query-response-interval"] = v.(string)
    }

    if v, ok := d.GetOk("startup_query_count"); ok {
        payload["startup-query-count"] = v.(string)
    }

    if v, ok := d.GetOk("startup_query_interval"); ok {
        payload["startup-query-interval"] = v.(string)
    }

    if v, ok := d.GetOk("mld_version"); ok {
        payload["mld-version"] = v.(string)
    }

    if v, ok := d.GetOkExists("reset"); ok {
        payload["reset"] = v.(bool)
    }

    setMldInterfaceRes, err := client.ApiCallSimple("set-mld-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setMldInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setMldInterfaceRes.Success {
            errMsg = setMldInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setMldInterfaceRes.GetData()
        }

        debugLogOperation(
            "mld-interface",        // resource type
            "update",                       // operation
            "set-mld-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set mld-interface: %v", err)
    }
    if !setMldInterfaceRes.Success {
        return fmt.Errorf(setMldInterfaceRes.ErrorMsg)
    }

    return readGaiaMldInterface(d, m)
}

func deleteGaiaMldInterface(d *schema.ResourceData, m interface{}) error {


        // No API call - just remove the ID to indicate resource deletion
        d.SetId("")
        return nil
    }

    