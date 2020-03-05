package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementAddApiKey() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementAddApiKey,
            Read:   readManagementAddApiKey,
            Delete: deleteManagementAddApiKey,
            Schema: map[string]*schema.Schema{ 
            "admin_uid": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Administrator uid to generate API key for.",
            },
            "admin_name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Administrator name to generate API key for.",
            },
        },
    }
}

func createManagementAddApiKey(d *schema.ResourceData, m interface{}) error {
    return readManagementAddApiKey(d, m)
}

func readManagementAddApiKey(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("admin_uid"); ok {
        payload["admin-uid"] = v.(string)
    }

    if v, ok := d.GetOk("admin_name"); ok {
        payload["admin-name"] = v.(string)
    }

    AddApiKeyRes, _ := client.ApiCall("add-api-key", payload, client.GetSessionID(), true, false)
    if !AddApiKeyRes.Success {
        return fmt.Errorf(AddApiKeyRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementAddApiKey(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

