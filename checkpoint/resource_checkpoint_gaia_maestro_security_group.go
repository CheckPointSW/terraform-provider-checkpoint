package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaMaestroSecurityGroup() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaMaestroSecurityGroup,
        Read:   readGaiaMaestroSecurityGroup,
        Update: updateGaiaMaestroSecurityGroup,
        Delete: deleteGaiaMaestroSecurityGroup,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debug logging for this resource.",
            },
            "interfaces": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `Orchestrator port, or list of Orchestrator ports, that will be assigned to this Security Group. At least one of ‘id’ or ‘interface-name’ parameters must be provided`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Interface ID (e.g. \"1/13/1\")`,
                        },
                        "name": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Interface name (e.g. \"eth1-05\")`,
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Description of the interface`,
                        },
                    },
                },
            },
            "gateways": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `Single Gateway or list of Gateways to be assigned to new Security Group`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `ID of this Gateway`,
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Description of this GW`,
                        },
                    },
                },
            },
            "sites": {
                Type:        schema.TypeList,
                Optional:    true,
                Description: `List of Site descriptions. The security group will be assigned to sites automatically according to gateways associated with the Security Group`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "resource_id": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `ID of this site`,
                        },
                        "description": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Description of this site`,
                        },
                    },
                },
            },
            "ftw_configuration": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `First Time Wizard configuration`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hostname": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Hostname for Security Group`,
                        },
                        "is_vsx": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `Determines if this Security Group is a VSX`,
                        },
                        "one_time_password": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `One time password for Secure Internal Communication (SIC)`,
                        },
                        "admin_password": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Sensitive:   true,
                            Description: `Admin password for Security Group`,
                        },
                    },
                },
            },
            "mgmt_connectivity": {
                Type:        schema.TypeList,
                Required:    true,
                Description: `The IP addresses that will be used to manage this Security Group`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "ipv4_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `IPv4 address for Security Group`,
                        },
                        "ipv6_address": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `IPv6 address for Security Group. Supported starting from Gaia version R82.10`,
                        },
                        "ipv4_mask_length": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `IPv4 mask length for Security Group`,
                        },
                        "ipv6_mask_length": {
                            Type:        schema.TypeInt,
                            Optional:    true,
                            Description: `IPv6 mask length for Security Group. Supported starting from Gaia version R82.10`,
                        },
                        "default_gateway": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Default Gateway address for Security Group`,
                        },
                        "ipv6_default_gateway": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `Default Gateway IPv6 address for Security Group. Supported starting from Gaia version R82.10`,
                        },
                    },
                },
            },
            "description": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `New Security Group description`,
            },
            "mgmt_interface_settings": {
                Type:        schema.TypeList,
                Optional:    true,
                ForceNew:    true,
                Description: `Management interface settings of this Security Group. By default, values are create-mgmt-as-bond == True and bond-mode == 'active-backup'.`,
                MaxItems:    1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "create_mgmt_as_bond": {
                            Type:        schema.TypeBool,
                            Optional:    true,
                            Description: `If True, a magg interface will be created for MGMT traffic. Every assigned MGMT interface will be enslaved to this magg. If False, only one of the assigned MGMT interfaces will be used for MGMT traffic.`,
                        },
                        "bond_mode": {
                            Type:        schema.TypeString,
                            Optional:    true,
                            Description: `If create-mgmt-as-bond is true, this field determines the magg bond type. If create-mgmt-as-bond is false, this field will be ignored.Note that using \"xor\" or \"8023AD\" entails configuring a bond on the device this Maestro environment is connected to.`,
                        },
                    },
                },
            },
            "resource_id": {
                Type:        schema.TypeInt,
                Optional:    true,
                Description: `Security Group ID`,
            },
            "include_pending_changes": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: `If true, show pending Security Groups changes. If false, show deployed topology`,
            },
        },
    }
}

func createGaiaMaestroSecurityGroup(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v := d.Get("interfaces"); len(v.([]interface{})) > 0 {
        interfacesList := v.([]interface{})
        interfacesArray := make([]interface{}, 0, len(interfacesList))
        for i := range interfacesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("interfaces.%d.resource_id", i)); ok {
                itemMap["id"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("interfaces.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("interfaces.%d.description", i)); ok {
                itemMap["description"] = v.(string)
            }
            if len(itemMap) > 0 {
                interfacesArray = append(interfacesArray, itemMap)
            }
        }
        if len(interfacesArray) > 0 {
            payload["interfaces"] = interfacesArray
        }
    }

    if v := d.Get("gateways"); len(v.([]interface{})) > 0 {
        gatewaysList := v.([]interface{})
        gatewaysArray := make([]interface{}, 0, len(gatewaysList))
        for i := range gatewaysList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("gateways.%d.resource_id", i)); ok {
                itemMap["id"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("gateways.%d.description", i)); ok {
                itemMap["description"] = v.(string)
            }
            if len(itemMap) > 0 {
                gatewaysArray = append(gatewaysArray, itemMap)
            }
        }
        if len(gatewaysArray) > 0 {
            payload["gateways"] = gatewaysArray
        }
    }

    if v := d.Get("sites"); len(v.([]interface{})) > 0 {
        sitesList := v.([]interface{})
        sitesArray := make([]interface{}, 0, len(sitesList))
        for i := range sitesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("sites.%d.resource_id", i)); ok {
                itemMap["id"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("sites.%d.description", i)); ok {
                itemMap["description"] = v.(string)
            }
            if len(itemMap) > 0 {
                sitesArray = append(sitesArray, itemMap)
            }
        }
        if len(sitesArray) > 0 {
            payload["sites"] = sitesArray
        }
    }

    if v := d.Get("ftw_configuration"); len(v.([]interface{})) > 0 {
        _ = v
        ftwconfigurationMap := make(map[string]interface{})
        if v, ok := d.GetOk("ftw_configuration.0.hostname"); ok {
            ftwconfigurationMap["hostname"] = v.(string)
        }
        if v, ok := d.GetOkExists("ftw_configuration.0.is_vsx"); ok && v.(bool) {
            ftwconfigurationMap["is-vsx"] = v.(bool)
        }
        if v, ok := d.GetOk("ftw_configuration.0.one_time_password"); ok {
            ftwconfigurationMap["one-time-password"] = v.(string)
        }
        if v, ok := d.GetOk("ftw_configuration.0.admin_password"); ok {
            ftwconfigurationMap["admin-password"] = v.(string)
        }
        if len(ftwconfigurationMap) > 0 {
            payload["ftw-configuration"] = ftwconfigurationMap
        }
    }

    if v := d.Get("mgmt_connectivity"); len(v.([]interface{})) > 0 {
        _ = v
        mgmtconnectivityMap := make(map[string]interface{})
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv4_address"); ok {
            mgmtconnectivityMap["ipv4-address"] = v.(string)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv6_address"); ok {
            mgmtconnectivityMap["ipv6-address"] = v.(string)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv4_mask_length"); ok {
            mgmtconnectivityMap["ipv4-mask-length"] = v.(int)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv6_mask_length"); ok {
            mgmtconnectivityMap["ipv6-mask-length"] = v.(int)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.default_gateway"); ok {
            mgmtconnectivityMap["default-gateway"] = v.(string)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv6_default_gateway"); ok {
            mgmtconnectivityMap["ipv6-default-gateway"] = v.(string)
        }
        if len(mgmtconnectivityMap) > 0 {
            payload["mgmt-connectivity"] = mgmtconnectivityMap
        }
    }

    if v, ok := d.GetOk("description"); ok {
        payload["description"] = v.(string)
    }

    if v := d.Get("mgmt_interface_settings"); len(v.([]interface{})) > 0 {
        _ = v
        mgmtinterfacesettingsMap := make(map[string]interface{})
        if v, ok := d.GetOkExists("mgmt_interface_settings.0.create_mgmt_as_bond"); ok && v.(bool) {
            mgmtinterfacesettingsMap["create-mgmt-as-bond"] = v.(bool)
        }
        if v, ok := d.GetOk("mgmt_interface_settings.0.bond_mode"); ok {
            mgmtinterfacesettingsMap["bond-mode"] = v.(string)
        }
        if len(mgmtinterfacesettingsMap) > 0 {
            payload["mgmt-interface-settings"] = mgmtinterfacesettingsMap
        }
    }

    log.Println("Create MaestroSecurityGroup - Map = ", payload)

    addMaestroSecurityGroupRes, err := client.ApiCallSimple("add-maestro-security-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addMaestroSecurityGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addMaestroSecurityGroupRes.Success {
            errMsg = addMaestroSecurityGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addMaestroSecurityGroupRes.GetData()
        }

        debugLogOperation(
            "maestro-security-group",        // resource type
            "create",                       // operation
            "add-maestro-security-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add maestro-security-group: %v", err)
    }
    if !addMaestroSecurityGroupRes.Success {
        if addMaestroSecurityGroupRes.ErrorMsg != "" {
            return fmt.Errorf(addMaestroSecurityGroupRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    // Extract API-assigned id from Create response so Read can look up the resource.
    if data := addMaestroSecurityGroupRes.GetData(); data != nil {
        if v, exists := data["id"]; exists {
            if f, ok := v.(float64); ok {
                d.Set("resource_id", int(f))
            }
        }
    }

    d.SetId(fmt.Sprintf("maestro-security-group-" + acctest.RandString(10)))
    return readGaiaMaestroSecurityGroup(d, m)
}

func readGaiaMaestroSecurityGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    if v, ok := d.GetOkExists("include_pending_changes"); ok {
        payload["include-pending-changes"] = v.(bool)
    }

    showMaestroSecurityGroupRes, err := client.ApiCallSimple("show-maestro-security-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showMaestroSecurityGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showMaestroSecurityGroupRes.Success {
            errMsg = showMaestroSecurityGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showMaestroSecurityGroupRes.GetData()
        }

        debugLogOperation(
            "maestro-security-group",        // resource type
            "read",                       // operation
            "show-maestro-security-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show maestro-security-group: %v", err)
    }
    if !showMaestroSecurityGroupRes.Success {
        if data := showMaestroSecurityGroupRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showMaestroSecurityGroupRes.ErrorMsg)
    }

    maestroSecurityGroup := showMaestroSecurityGroupRes.GetData()

    log.Println("Read MaestroSecurityGroup - Show JSON = ", maestroSecurityGroup)

    if v, exists := maestroSecurityGroup["id"]; exists {
        if f, ok := v.(float64); ok {
            d.Set("resource_id", int(f))
        } else if s, ok := v.(string); ok {
            var _n int
            if _, _err := fmt.Sscanf(s, "%d", &_n); _err == nil {
                d.Set("resource_id", _n)
            }
        }
    }
    if v, exists := maestroSecurityGroup["interfaces"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    out = append(out, map[string]interface{}{
                        "resource_id": fmt.Sprintf("%v", m["id"]),
                        "name":        fmt.Sprintf("%v", m["name"]),
                        "description": fmt.Sprintf("%v", m["description"]),
                    })
                }
            }
            d.Set("interfaces", out)
        }
    }
    if v, exists := maestroSecurityGroup["gateways"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    out = append(out, map[string]interface{}{
                        "resource_id": fmt.Sprintf("%v", m["id"]),
                        "description": fmt.Sprintf("%v", m["description"]),
                    })
                }
            }
            d.Set("gateways", out)
        }
    }
    if v, exists := maestroSecurityGroup["sites"]; exists {
        if items, ok := v.([]interface{}); ok {
            out := make([]interface{}, 0, len(items))
            for _, item := range items {
                if m, ok := item.(map[string]interface{}); ok {
                    out = append(out, map[string]interface{}{
                        "resource_id": func() int { if f, ok := m["id"].(float64); ok { return int(f) }; return 0 }(),
                        "description": fmt.Sprintf("%v", m["description"]),
                    })
                }
            }
            d.Set("sites", out)
        }
    }
    if v, exists := maestroSecurityGroup["ftw-configuration"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            d.Set("ftw_configuration", []interface{}{map[string]interface{}{
                "hostname":          fmt.Sprintf("%v", m["hostname"]),
                "is_vsx":            func() bool { b, _ := m["is-vsx"].(bool); return b }(),
                "one_time_password": fmt.Sprintf("%v", m["one-time-password"]),
                "admin_password":    fmt.Sprintf("%v", m["admin-password"]),
            }})
        }
    }
    if v, exists := maestroSecurityGroup["mgmt-connectivity"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            d.Set("mgmt_connectivity", []interface{}{map[string]interface{}{
                "ipv4_address":         fmt.Sprintf("%v", m["ipv4-address"]),
                "ipv6_address":         fmt.Sprintf("%v", m["ipv6-address"]),
                "ipv4_mask_length":     func() int { if f, ok := m["ipv4-mask-length"].(float64); ok { return int(f) }; return 0 }(),
                "ipv6_mask_length":     func() int { if f, ok := m["ipv6-mask-length"].(float64); ok { return int(f) }; return 0 }(),
                "default_gateway":      fmt.Sprintf("%v", m["default-gateway"]),
                "ipv6_default_gateway": fmt.Sprintf("%v", m["ipv6-default-gateway"]),
            }})
        }
    }
    if v, exists := maestroSecurityGroup["mgmt-interface-settings"]; exists {
        if m, ok := v.(map[string]interface{}); ok {
            d.Set("mgmt_interface_settings", []interface{}{map[string]interface{}{
                "create_mgmt_as_bond": func() bool { b, _ := m["create-mgmt-as-bond"].(bool); return b }(),
                "bond_mode":          fmt.Sprintf("%v", m["bond-mode"]),
            }})
        }
    }
    if v, exists := maestroSecurityGroup["description"]; exists {
        d.Set("description", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func updateGaiaMaestroSecurityGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v := d.Get("interfaces"); len(v.([]interface{})) > 0 {
        interfacesList := v.([]interface{})
        interfacesArray := make([]interface{}, 0, len(interfacesList))
        for i := range interfacesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("interfaces.%d.resource_id", i)); ok {
                itemMap["id"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("interfaces.%d.name", i)); ok {
                itemMap["name"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("interfaces.%d.description", i)); ok {
                itemMap["description"] = v.(string)
            }
            if len(itemMap) > 0 {
                interfacesArray = append(interfacesArray, itemMap)
            }
        }
        if len(interfacesArray) > 0 {
            payload["interfaces"] = interfacesArray
        }
    }

    if v := d.Get("gateways"); len(v.([]interface{})) > 0 {
        gatewaysList := v.([]interface{})
        gatewaysArray := make([]interface{}, 0, len(gatewaysList))
        for i := range gatewaysList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("gateways.%d.resource_id", i)); ok {
                itemMap["id"] = v.(string)
            }
            if v, ok := d.GetOk(fmt.Sprintf("gateways.%d.description", i)); ok {
                itemMap["description"] = v.(string)
            }
            if len(itemMap) > 0 {
                gatewaysArray = append(gatewaysArray, itemMap)
            }
        }
        if len(gatewaysArray) > 0 {
            payload["gateways"] = gatewaysArray
        }
    }

    if v := d.Get("sites"); len(v.([]interface{})) > 0 {
        sitesList := v.([]interface{})
        sitesArray := make([]interface{}, 0, len(sitesList))
        for i := range sitesList {
            itemMap := make(map[string]interface{})
            if v, ok := d.GetOk(fmt.Sprintf("sites.%d.resource_id", i)); ok {
                itemMap["id"] = v.(int)
            }
            if v, ok := d.GetOk(fmt.Sprintf("sites.%d.description", i)); ok {
                itemMap["description"] = v.(string)
            }
            if len(itemMap) > 0 {
                sitesArray = append(sitesArray, itemMap)
            }
        }
        if len(sitesArray) > 0 {
            payload["sites"] = sitesArray
        }
    }

    if v := d.Get("ftw_configuration"); len(v.([]interface{})) > 0 {
        _ = v
        ftwconfigurationMap := make(map[string]interface{})
        if v, ok := d.GetOk("ftw_configuration.0.hostname"); ok {
            ftwconfigurationMap["hostname"] = v.(string)
        }
        if v, ok := d.GetOkExists("ftw_configuration.0.is_vsx"); ok && v.(bool) {
            ftwconfigurationMap["is-vsx"] = v.(bool)
        }
        if v, ok := d.GetOk("ftw_configuration.0.one_time_password"); ok {
            ftwconfigurationMap["one-time-password"] = v.(string)
        }
        if v, ok := d.GetOk("ftw_configuration.0.admin_password"); ok {
            ftwconfigurationMap["admin-password"] = v.(string)
        }
        if len(ftwconfigurationMap) > 0 {
            payload["ftw-configuration"] = ftwconfigurationMap
        }
    }

    if v := d.Get("mgmt_connectivity"); len(v.([]interface{})) > 0 {
        _ = v
        mgmtconnectivityMap := make(map[string]interface{})
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv4_address"); ok {
            mgmtconnectivityMap["ipv4-address"] = v.(string)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv6_address"); ok {
            mgmtconnectivityMap["ipv6-address"] = v.(string)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv4_mask_length"); ok {
            mgmtconnectivityMap["ipv4-mask-length"] = v.(int)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv6_mask_length"); ok {
            mgmtconnectivityMap["ipv6-mask-length"] = v.(int)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.default_gateway"); ok {
            mgmtconnectivityMap["default-gateway"] = v.(string)
        }
        if v, ok := d.GetOk("mgmt_connectivity.0.ipv6_default_gateway"); ok {
            mgmtconnectivityMap["ipv6-default-gateway"] = v.(string)
        }
        if len(mgmtconnectivityMap) > 0 {
            payload["mgmt-connectivity"] = mgmtconnectivityMap
        }
    }

    if v, ok := d.GetOk("description"); ok {
        payload["description"] = v.(string)
    }

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    setMaestroSecurityGroupRes, err := client.ApiCallSimple("set-maestro-security-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && setMaestroSecurityGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !setMaestroSecurityGroupRes.Success {
            errMsg = setMaestroSecurityGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = setMaestroSecurityGroupRes.GetData()
        }

        debugLogOperation(
            "maestro-security-group",        // resource type
            "update",                       // operation
            "set-maestro-security-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to set maestro-security-group: %v", err)
    }
    if !setMaestroSecurityGroupRes.Success {
        return fmt.Errorf(setMaestroSecurityGroupRes.ErrorMsg)
    }

    return readGaiaMaestroSecurityGroup(d, m)
}

func deleteGaiaMaestroSecurityGroup(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("resource_id"); ok {
        payload["id"] = v.(int)
    }

    deleteMaestroSecurityGroupRes, err := client.ApiCallSimple("delete-maestro-security-group", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteMaestroSecurityGroupRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteMaestroSecurityGroupRes.Success {
            errMsg = deleteMaestroSecurityGroupRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteMaestroSecurityGroupRes.GetData()
        }

        debugLogOperation(
            "maestro-security-group",        // resource type
            "delete",                       // operation
            "delete-maestro-security-group",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete maestro-security-group: %v", err)
    }
    if !deleteMaestroSecurityGroupRes.Success {
        return fmt.Errorf(deleteMaestroSecurityGroupRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

