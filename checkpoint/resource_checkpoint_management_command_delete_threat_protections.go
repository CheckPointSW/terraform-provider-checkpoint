package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementDeleteThreatProtections() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementDeleteThreatProtections,
            Read:   readManagementDeleteThreatProtections,
            Delete: deleteManagementDeleteThreatProtections,
            Schema: map[string]*schema.Schema{ 
            "package_format": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "Protections package format.",
            },
        },
    }
}

func createManagementDeleteThreatProtections(d *schema.ResourceData, m interface{}) error {
    return readManagementDeleteThreatProtections(d, m)
}

func readManagementDeleteThreatProtections(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("package_format"); ok {
        payload["package-format"] = v.(string)
    }

    DeleteThreatProtectionsRes, _ := client.ApiCall("delete-threat-protections", payload, client.GetSessionID(), true, false)
    if !DeleteThreatProtectionsRes.Success {
        return fmt.Errorf(DeleteThreatProtectionsRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementDeleteThreatProtections(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

