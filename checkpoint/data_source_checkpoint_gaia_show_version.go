package checkpoint

import (
        "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
)
func dataGaiaShowVersion() *schema.Resource {   
    return &schema.Resource{
        Read:   readGaiaShowVersion,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                Description: "Enable debugging for this resource only.",
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "product_version": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "os_build": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "os_kernel_version": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "os_edition": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func readGaiaShowVersion(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    log.Println("Execute show-version - Payload = ", payload)
    commandRes, err := client.ApiCallSimple("show-version", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && commandRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !commandRes.Success {
            errMsg = commandRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = commandRes.GetData()
        }

        debugLogOperation(
            "version",        // resource type
            "read",                       // operation
            "show-version",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to execute show-version: %v", err)
    }
    if !commandRes.Success {
        return fmt.Errorf(commandRes.ErrorMsg)
    }

    if v, exists := commandRes.GetData()["product-version"]; exists {
        d.Set("product_version", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["os-build"]; exists {
        d.Set("os_build", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["os-kernel-version"]; exists {
        d.Set("os_kernel_version", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["os-edition"]; exists {
        d.Set("os_edition", fmt.Sprintf("%v", v))
    }
    if v, exists := commandRes.GetData()["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(fmt.Sprintf("show-version-" + acctest.RandString(10)))
    return nil
}

