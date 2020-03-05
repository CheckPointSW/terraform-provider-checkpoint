package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementSetApiSettings() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementSetApiSettings,
            Read:   readManagementSetApiSettings,
            Delete: deleteManagementSetApiSettings,
            Schema: map[string]*schema.Schema{ 
            "accepted_api_calls_from": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "Clients allowed to connect to the API Server.",
            },
            "automatic_start": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "MGMT API will start after server will start.",
            },
        },
    }
}

func createManagementSetApiSettings(d *schema.ResourceData, m interface{}) error {
    return readManagementSetApiSettings(d, m)
}

func readManagementSetApiSettings(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("accepted_api_calls_from"); ok {
        payload["accepted-api-calls-from"] = v.(string)
    }

    if v, ok := d.GetOkExists("automatic_start"); ok {
        payload["automatic-start"] = v.(bool)
    }

    SetApiSettingsRes, _ := client.ApiCall("set-api-settings", payload, client.GetSessionID(), true, false)
    if !SetApiSettingsRes.Success {
        return fmt.Errorf(SetApiSettingsRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementSetApiSettings(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

