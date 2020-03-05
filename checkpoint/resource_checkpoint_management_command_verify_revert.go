package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementVerifyRevert() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementVerifyRevert,
            Read:   readManagementVerifyRevert,
            Delete: deleteManagementVerifyRevert,
            Schema: map[string]*schema.Schema{ 
            "to_session": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Session unique identifier. Specify the session you would like to verify a revert operation to.",
            },
        },
    }
}

func createManagementVerifyRevert(d *schema.ResourceData, m interface{}) error {
    return readManagementVerifyRevert(d, m)
}

func readManagementVerifyRevert(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("to_session"); ok {
        payload["to-session"] = v.(string)
    }

    VerifyRevertRes, _ := client.ApiCall("verify-revert", payload, client.GetSessionID(), true, false)
    if !VerifyRevertRes.Success {
        return fmt.Errorf(VerifyRevertRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementVerifyRevert(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

