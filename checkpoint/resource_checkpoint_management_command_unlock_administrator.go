package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementUnlockAdministrator() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementUnlockAdministrator,
            Read:   readManagementUnlockAdministrator,
            Delete: deleteManagementUnlockAdministrator,
            Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Object name.",
            },
        },
    }
}

func createManagementUnlockAdministrator(d *schema.ResourceData, m interface{}) error {
    return readManagementUnlockAdministrator(d, m)
}

func readManagementUnlockAdministrator(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    UnlockAdministratorRes, _ := client.ApiCall("unlock-administrator", payload, client.GetSessionID(), true, false)
    if !UnlockAdministratorRes.Success {
        return fmt.Errorf(UnlockAdministratorRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementUnlockAdministrator(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

