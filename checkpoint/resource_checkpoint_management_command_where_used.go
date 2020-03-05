package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementWhereUsed() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementWhereUsed,
            Read:   readManagementWhereUsed,
            Delete: deleteManagementWhereUsed,
            Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Object name.",
            },
            "dereference_group_members": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Indicates whether to dereference \"members\" field by details level for every object in reply.",
            },
            "show_membership": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Indicates whether to calculate and show \"groups\" field for every object in reply.",
            },
            "indirect": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Search for indirect usage.",
            },
            "indirect_max_depth": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: "Maximum nesting level during indirect usage search.",
            },
        },
    }
}

func createManagementWhereUsed(d *schema.ResourceData, m interface{}) error {
    return readManagementWhereUsed(d, m)
}

func readManagementWhereUsed(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("dereference_group_members"); ok {
        payload["dereference-group-members"] = v.(bool)
    }

    if v, ok := d.GetOkExists("show_membership"); ok {
        payload["show-membership"] = v.(bool)
    }

    if v, ok := d.GetOkExists("indirect"); ok {
        payload["indirect"] = v.(bool)
    }

    if v, ok := d.GetOk("indirect_max_depth"); ok {
        payload["indirect-max-depth"] = v.(int)
    }

    WhereUsedRes, _ := client.ApiCall("where-used", payload, client.GetSessionID(), true, false)
    if !WhereUsedRes.Success {
        return fmt.Errorf(WhereUsedRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementWhereUsed(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

