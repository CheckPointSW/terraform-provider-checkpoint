package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementAssignGlobalAssignment() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementAssignGlobalAssignment,
            Read:   readManagementAssignGlobalAssignment,
            Delete: deleteManagementAssignGlobalAssignment,
            Schema: map[string]*schema.Schema{ 
            "dependent_domains": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "global_domains": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
        },
    }
}

func createManagementAssignGlobalAssignment(d *schema.ResourceData, m interface{}) error {
    return readManagementAssignGlobalAssignment(d, m)
}

func readManagementAssignGlobalAssignment(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("dependent_domains"); ok {
        payload["dependent-domains"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("global_domains"); ok {
        payload["global-domains"] = v.(*schema.Set).List()
    }

    AssignGlobalAssignmentRes, _ := client.ApiCall("assign-global-assignment", payload, client.GetSessionID(), true, false)
    if !AssignGlobalAssignmentRes.Success {
        return fmt.Errorf(AssignGlobalAssignmentRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementAssignGlobalAssignment(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

