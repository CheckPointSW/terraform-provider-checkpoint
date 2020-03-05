package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementAddUpdatableObject() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementAddUpdatableObject,
            Read:   readManagementAddUpdatableObject,
            Delete: deleteManagementAddUpdatableObject,
            Schema: map[string]*schema.Schema{ 
            "uri": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "URI of the updatable object in the Updatable Objects Repository.",
            },
            "uid_in_updatable_objects_repository": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Unique identifier of the updatable object in the Updatable Objects Repository.",
            },
            "tags": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: "Collection of tag identifiers.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "color": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "Color of the object. Should be one of existing colors.",
            },
            "comments": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "Comments string.",
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

func createManagementAddUpdatableObject(d *schema.ResourceData, m interface{}) error {
    return readManagementAddUpdatableObject(d, m)
}

func readManagementAddUpdatableObject(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("uri"); ok {
        payload["uri"] = v.(string)
    }

    if v, ok := d.GetOk("uid_in_updatable_objects_repository"); ok {
        payload["uid-in-updatable-objects-repository"] = v.(string)
    }

    if v, ok := d.GetOk("tags"); ok {
        payload["tags"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("color"); ok {
        payload["color"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        payload["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        payload["ignore-errors"] = v.(bool)
    }

    AddUpdatableObjectRes, _ := client.ApiCall("add-updatable-object", payload, client.GetSessionID(), true, false)
    if !AddUpdatableObjectRes.Success {
        return fmt.Errorf(AddUpdatableObjectRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementAddUpdatableObject(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

