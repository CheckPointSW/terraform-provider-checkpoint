package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementDeleteUpdatableObject() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementDeleteUpdatableObject,
            Read:   readManagementDeleteUpdatableObject,
            Delete: deleteManagementDeleteUpdatableObject,
            Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Object name.",
            },
            "ignore_warnings": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Apply changes ignoring warnings.",
            },
            "ignore_errors": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
            },
        },
    }
}

func createManagementDeleteUpdatableObject(d *schema.ResourceData, m interface{}) error {
    return readManagementDeleteUpdatableObject(d, m)
}

func readManagementDeleteUpdatableObject(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        payload["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        payload["ignore-errors"] = v.(bool)
    }

    DeleteUpdatableObjectRes, _ := client.ApiCall("delete-updatable-object", payload, client.GetSessionID(), true, false)
    if !DeleteUpdatableObjectRes.Success {
        return fmt.Errorf(DeleteUpdatableObjectRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementDeleteUpdatableObject(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

