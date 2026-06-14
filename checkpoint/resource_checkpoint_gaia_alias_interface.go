package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaAliasInterface() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaAliasInterface,
        Read:   readGaiaAliasInterface,
        Delete: deleteGaiaAliasInterface,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Enable debug logging for this resource.",
            },
            "parent": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `N/A`,
            },
            "ipv4_address": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `N/A`,
            },
            "ipv4_mask_length": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `N/A`,
            },
            "virtual_system_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: `Virtual System ID. Relevant for VSNext setups`,
            },
            "name": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
        },
    }
}

func createGaiaAliasInterface(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("parent"); ok {
        payload["parent"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_address"); ok {
        payload["ipv4-address"] = v.(string)
    }

    if v, ok := d.GetOk("ipv4_mask_length"); ok {
        payload["ipv4-mask-length"] = v.(int)
    }

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

    log.Println("Create AliasInterface - Map = ", payload)

    addAliasInterfaceRes, err := client.ApiCallSimple("add-alias-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addAliasInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addAliasInterfaceRes.Success {
            errMsg = addAliasInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addAliasInterfaceRes.GetData()
        }

        debugLogOperation(
            "alias-interface",        // resource type
            "create",                       // operation
            "add-alias-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add alias-interface: %v", err)
    }
    if !addAliasInterfaceRes.Success {
        if addAliasInterfaceRes.ErrorMsg != "" {
            return fmt.Errorf(addAliasInterfaceRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    // Extract API-assigned fields from Create response before calling Read.
    if data := addAliasInterfaceRes.GetData(); data != nil {
        if v, exists := data["name"]; exists {
            d.Set("name", v)
        }
    }

    d.SetId(fmt.Sprintf("alias-interface-" + acctest.RandString(10)))
    return readGaiaAliasInterface(d, m)
}

func readGaiaAliasInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    // name is Computed (auto-assigned by device); only send it if already in state.
    delete(payload, "name")
    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    showAliasInterfaceRes, err := client.ApiCallSimple("show-alias-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showAliasInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showAliasInterfaceRes.Success {
            errMsg = showAliasInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showAliasInterfaceRes.GetData()
        }

        debugLogOperation(
            "alias-interface",        // resource type
            "read",                       // operation
            "show-alias-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show alias-interface: %v", err)
    }
    if !showAliasInterfaceRes.Success {
        if data := showAliasInterfaceRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showAliasInterfaceRes.ErrorMsg)
    }

    aliasInterface := showAliasInterfaceRes.GetData()

    log.Println("Read AliasInterface - Show JSON = ", aliasInterface)

    if v, exists := aliasInterface["parent"]; exists {
        d.Set("parent", fmt.Sprintf("%v", v))
    }
    if v, exists := aliasInterface["ipv4-address"]; exists {
        d.Set("ipv4_address", fmt.Sprintf("%v", v))
    }
    if v, exists := aliasInterface["ipv4-mask-length"]; exists {
        d.Set("ipv4_mask_length", fmt.Sprintf("%v", v))
    }
    if v, exists := aliasInterface["name"]; exists {
        d.Set("name", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func deleteGaiaAliasInterface(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("virtual_system_id"); ok {
        payload["virtual-system-id"] = v.(int)
    }

   payload["name"] = d.Get("name")
    // name is Computed (auto-assigned by device); only send it if already in state.
    delete(payload, "name")
    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    deleteAliasInterfaceRes, err := client.ApiCallSimple("delete-alias-interface", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteAliasInterfaceRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteAliasInterfaceRes.Success {
            errMsg = deleteAliasInterfaceRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteAliasInterfaceRes.GetData()
        }

        debugLogOperation(
            "alias-interface",        // resource type
            "delete",                       // operation
            "delete-alias-interface",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete alias-interface: %v", err)
    }
    if !deleteAliasInterfaceRes.Success {
        return fmt.Errorf(deleteAliasInterfaceRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

