package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementKeepalive() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementKeepalive,
            Read:   readManagementKeepalive,
            Delete: deleteManagementKeepalive,
            Schema: map[string]*schema.Schema{ 
        },
    }
}

func createManagementKeepalive(d *schema.ResourceData, m interface{}) error {
    return readManagementKeepalive(d, m)
}

func readManagementKeepalive(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    KeepaliveRes, _ := client.ApiCall("keepalive", payload, client.GetSessionID(), true, false)
    if !KeepaliveRes.Success {
        return fmt.Errorf(KeepaliveRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementKeepalive(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

