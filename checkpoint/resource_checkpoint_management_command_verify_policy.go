package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementVerifyPolicy() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementVerifyPolicy,
            Read:   readManagementVerifyPolicy,
            Delete: deleteManagementVerifyPolicy,
            Schema: map[string]*schema.Schema{ 
            "policy_package": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Policy package identified by the name or UID.",
            },
        },
    }
}

func createManagementVerifyPolicy(d *schema.ResourceData, m interface{}) error {
    return readManagementVerifyPolicy(d, m)
}

func readManagementVerifyPolicy(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("policy_package"); ok {
        payload["policy-package"] = v.(string)
    }

    VerifyPolicyRes, _ := client.ApiCall("verify-policy", payload, client.GetSessionID(), true, false)
    if !VerifyPolicyRes.Success {
        return fmt.Errorf(VerifyPolicyRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementVerifyPolicy(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

