package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
    "log"
    "strings"

)
func resourceGaiaLicense() *schema.Resource {   
    return &schema.Resource{
        Create: createGaiaLicense,
        Read:   readGaiaLicense,
        Delete: deleteGaiaLicense,
        Schema: map[string]*schema.Schema{
            "debug": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Enable debug logging for this resource.",
            },
            "license": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: `The license string received from the User Center - without 'cplic put'`,
            },
            "target": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `The remote target to deploy the license on - used for central licenses only`,
            },
            "signature": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `The license signature to show details for`,
            },
            "member_id": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: `Relevant for commands on Scalable and ElasticXL platforms only.<br>When member-id is provided in the login request,<br>show commands during the session will be executed on the specified member,<br>unless a different member-id is provided in a successive requests<br>Set operations will be performed on all members`,
            },
            "ip_addr": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "expiration": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "sku": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "ck": {
                Type:        schema.TypeString,
                Computed:    true,
                Description: `N/A`,
            },
            "central": {
                Type:        schema.TypeBool,
                Computed:    true,
                Description: `N/A`,
            },
        },
    }
}

func createGaiaLicense(d *schema.ResourceData, m interface{}) error {
    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := make(map[string]interface{})

    if v, ok := d.GetOk("license"); ok {
        payload["license"] = v.(string)
    }

    if v, ok := d.GetOk("target"); ok {
        payload["target"] = v.(string)
    }

    log.Println("Create License - Map = ", payload)

    addLicenseRes, err := client.ApiCallSimple("add-license", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && addLicenseRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !addLicenseRes.Success {
            errMsg = addLicenseRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = addLicenseRes.GetData()
        }

        debugLogOperation(
            "license",        // resource type
            "create",                       // operation
            "add-license",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to add license: %v", err)
    }
    if !addLicenseRes.Success {
        if addLicenseRes.ErrorMsg != "" {
            return fmt.Errorf(addLicenseRes.ErrorMsg)
        }
        return fmt.Errorf("Unknown error occurred")
    }

    // Extract API-assigned fields from Create response before calling Read.
    if data := addLicenseRes.GetData(); data != nil {
        if v, exists := data["signature"]; exists {
            d.Set("signature", v)
        }
    }

    d.SetId(fmt.Sprintf("license-" + acctest.RandString(10)))
    return readGaiaLicense(d, m)
}

func readGaiaLicense(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("signature"); ok {
        payload["signature"] = v.(string)
    }

    if v, ok := d.GetOk("member_id"); ok {
        payload["member-id"] = v.(string)
    }

    showLicenseRes, err := client.ApiCallSimple("show-license", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && showLicenseRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !showLicenseRes.Success {
            errMsg = showLicenseRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = showLicenseRes.GetData()
        }

        debugLogOperation(
            "license",        // resource type
            "read",                       // operation
            "show-license",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to show license: %v", err)
    }
    if !showLicenseRes.Success {
        if data := showLicenseRes.GetData(); data != nil {
            if code, exists := data["code"]; exists {
                if strings.Contains(strings.ToLower(code.(string)), "not_found") || strings.Contains(strings.ToLower(code.(string)), "object_not_found") {
                    d.SetId("")
                    return nil
                }
            }
        }
        return fmt.Errorf(showLicenseRes.ErrorMsg)
    }

    license := showLicenseRes.GetData()

    log.Println("Read License - Show JSON = ", license)

    if v, exists := license["ip_addr"]; exists {
        d.Set("ip_addr", fmt.Sprintf("%v", v))
    }
    if v, exists := license["expiration"]; exists {
        d.Set("expiration", fmt.Sprintf("%v", v))
    }
    if v, exists := license["signature"]; exists {
        d.Set("signature", fmt.Sprintf("%v", v))
    }
    if v, exists := license["SKU"]; exists {
        d.Set("sku", fmt.Sprintf("%v", v))
    }
    if v, exists := license["CK"]; exists {
        d.Set("ck", fmt.Sprintf("%v", v))
    }
    if v, exists := license["central"]; exists {
        if b, ok := v.(bool); ok {
            d.Set("central", b)
        } else if s, ok := v.(string); ok {
            d.Set("central", s == "true")
        }
    }
    if v, exists := license["target"]; exists {
        d.Set("target", fmt.Sprintf("%v", v))
    }
    if v, exists := license["member-id"]; exists {
        d.Set("member_id", fmt.Sprintf("%v", v))
    }
    d.SetId(d.Id())
    return nil
}

func deleteGaiaLicense(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)
    ensureDebugServerFromClient(client)

    payload := map[string]interface{}{}

    if v, ok := d.GetOk("target"); ok {
        payload["target"] = v.(string)
    }

    if v, ok := d.GetOk("signature"); ok {
        payload["signature"] = v.(string)
    }

    deleteLicenseRes, err := client.ApiCallSimple("delete-license", payload)
    // DEBUG: generic logger
    if resourceDebugEnabled(d) {
        success := err == nil && deleteLicenseRes.Success
        errMsg := ""
        if err != nil {
            errMsg = err.Error()
        } else if !deleteLicenseRes.Success {
            errMsg = deleteLicenseRes.ErrorMsg
        }

        var respData map[string]interface{}
        if err == nil {
            respData = deleteLicenseRes.GetData()
        }

        debugLogOperation(
            "license",        // resource type
            "delete",                       // operation
            "delete-license",         // API call name
            payload,                        // request payload
            respData,                       // response data (if any)
            success,
            errMsg,
        )
    }
    if err != nil {
        return fmt.Errorf("Failed to delete license: %v", err)
    }
    if !deleteLicenseRes.Success {
        return fmt.Errorf(deleteLicenseRes.ErrorMsg)
    }

    d.SetId("")
    return nil
}

