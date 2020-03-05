package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementDeleteApiKey() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementDeleteApiKey,
            Read:   readManagementDeleteApiKey,
            Delete: deleteManagementDeleteApiKey,
            Schema: map[string]*schema.Schema{ 
            "api_key": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "API key to be deleted.",
            },
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

func createManagementDeleteApiKey(d *schema.ResourceData, m interface{}) error {
    return readManagementDeleteApiKey(d, m)
}

func readManagementDeleteApiKey(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("api_key"); ok {
        payload["api-key"] = v.(string)
    }

    if v, ok := d.GetOk("admin_uid"); ok {
        payload["admin-uid"] = v.(string)
    }

    if v, ok := d.GetOk("admin_name"); ok {
        payload["admin-name"] = v.(string)
    }

    DeleteApiKeyRes, _ := client.ApiCall("delete-api-key", payload, client.GetSessionID(), true, false)
    if !DeleteApiKeyRes.Success {
        return fmt.Errorf(DeleteApiKeyRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementDeleteApiKey(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

