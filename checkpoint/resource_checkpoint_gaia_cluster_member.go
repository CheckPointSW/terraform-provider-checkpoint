package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaClusterMember() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaClusterMember,
        Read:   readGaiaClusterMember,
        Delete: deleteGaiaClusterMember,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Enable debug logging for this resource.",
            },
            "method": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Method used for adding the member:<br>        1. serial-number - Retrieve from new member using \"show-serial-number\"<br>        2. hostname - Retrieve from new member using \"show-hostname\"<br>        3. request-id - Retrieve from new member using \"show-cluster-request-id\"`,
            },
            "identifier": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `Identifier of member`,
            },
            "site_id": {
                Type:        schema.TypeInt,
                Required:    true,
                ForceNew:    true,
                Description: `Site id to add member to`,
            },
            "member": {
                Type:        schema.TypeList,
                Computed:    true,
                Description: `N/A`,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "hostname": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "serial_number": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "request_id": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "site_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "member_id": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "model": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "version": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "member_status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "site_status": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "state": {
                            Type:        schema.TypeString,
                            Computed:    true,
                            Description: `N/A`,
                        },
                        "installed_jumbo_take": {
                            Type:        schema.TypeInt,
                            Computed:    true,
                            Description: `N/A`,
                        },
                    },
                },
            },
        },
    }
}

func createGaiaClusterMember(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("method"); ok {
        payload["method"] = v.(string)
    }

    if v, ok := d.GetOk("identifier"); ok {
        payload["identifier"] = v.(string)
    }

    if v, ok := d.GetOk("site_id"); ok {
        payload["site-id"] = v.(int)
    }

    log.Println("Create ClusterMember - Map = ", payload)

    addClusterMemberRes, err := client.ApiCallSimple("add-cluster-member", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addClusterMemberRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addClusterMemberRes.Success {
            errMsg = addClusterMemberRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addClusterMemberRes.GetData()
        }

        debugLogOperation(
            "cluster-member",        // resource type
            "create",                       // operation
            "add-cluster-member",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add cluster-member: %v", err)
    }
    if !addClusterMemberRes.Success {
        if addClusterMemberRes.ErrorMsg != "" {
            return fmt.Errorf(addClusterMemberRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    d.SetId(fmt.Sprintf("cluster-member-" + acctest.RandString(10)))
    return readGaiaClusterMember(d, m)
}

func readGaiaClusterMember(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    showClusterMemberRes, err := client.ApiCallSimple("show-cluster-members", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showClusterMemberRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showClusterMemberRes.Success {
            errMsg = showClusterMemberRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showClusterMemberRes.GetData()
        }

        debugLogOperation(
            "cluster-member",        // resource type
            "read",                       // operation
            "show-cluster-members",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show cluster-member: %v", err)
    }
    if !showClusterMemberRes.Success {
        if data := showClusterMemberRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showClusterMemberRes.ErrorMsg)
    }

    clusterMember := showClusterMemberRes.GetData()

    log.Println("Read ClusterMember - Show JSON = ", clusterMember)

    if v, exists := clusterMember["pending-gateways"]; exists {
        d.Set("pending_gateways", v.([]interface{}))
    }
    // show-cluster-members returns a members list; find entry matching method+identifier.
    methodVal := ""
    if mv, ok := d.GetOk("method"); ok {
        methodVal = mv.(string)
    }
    identifierVal := ""
    if iv, ok := d.GetOk("identifier"); ok {
        identifierVal = iv.(string)
    }
    if members, ok := clusterMember["members"].([]interface{}); ok {
        found := false
        for _, obj := range members {
            if item, ok := obj.(map[string]interface{}); ok {
                if fmt.Sprintf("%v", item[methodVal]) == identifierVal {
                    memberMap := map[string]interface{}{
                        "hostname":             fmt.Sprintf("%v", item["hostname"]),
                        "serial_number":        fmt.Sprintf("%v", item["serial-number"]),
                        "request_id":           fmt.Sprintf("%v", item["request-id"]),
                        "site_id":              func() int { if f, ok := item["site-id"].(float64); ok { return int(f) }; return 0 }(),
                        "member_id":            func() int { if f, ok := item["member-id"].(float64); ok { return int(f) }; return 0 }(),
                        "model":                fmt.Sprintf("%v", item["model"]),
                        "version":              fmt.Sprintf("%v", item["version"]),
                        "member_status":        fmt.Sprintf("%v", item["member-status"]),
                        "site_status":          fmt.Sprintf("%v", item["site-status"]),
                        "state":                fmt.Sprintf("%v", item["state"]),
                        "installed_jumbo_take": func() int { if f, ok := item["installed-jumbo-take"].(float64); ok { return int(f) }; return 0 }(),
                    }
                    d.Set("member", []interface{}{memberMap})
                    found = true
                    break
                }
            }
        }
        if !found {
            d.SetId("")
            return nil
        }
    }
    d.SetId(d.Id())
    return nil
}

func deleteGaiaClusterMember(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("method"); ok {
        payload["method"] = v.(string)
    }

    if v, ok := d.GetOk("identifier"); ok {
        payload["identifier"] = v.(string)
    }

    deleteClusterMemberRes, err := client.ApiCallSimple("delete-cluster-member", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteClusterMemberRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteClusterMemberRes.Success {
            errMsg = deleteClusterMemberRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteClusterMemberRes.GetData()
        }

        debugLogOperation(
            "cluster-member",        // resource type
            "delete",                       // operation
            "delete-cluster-member",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete cluster-member: %v", err)
    }
    if !deleteClusterMemberRes.Success {
        return fmt.Errorf(deleteClusterMemberRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

