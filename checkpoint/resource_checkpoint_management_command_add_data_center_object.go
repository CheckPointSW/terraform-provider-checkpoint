package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementAddDataCenterObject() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementAddDataCenterObject,
            Read:   readManagementAddDataCenterObject,
            Delete: deleteManagementAddDataCenterObject,
            Schema: map[string]*schema.Schema{ 
            "data_center_name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Name of the Data Center Server the object is in.",
            },
            "data_center_uid": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Unique identifier of the Data Center Server the object is in.",
            },
            "uri": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "URI of the object in the Data Center Server.",
            },
            "uid_in_data_center": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Unique identifier of the object in the Data Center Server.",
            },
            "name": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "Override default name on data-center.",
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
            "groups": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: "Collection of group identifiers.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
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

func createManagementAddDataCenterObject(d *schema.ResourceData, m interface{}) error {
    return readManagementAddDataCenterObject(d, m)
}

func readManagementAddDataCenterObject(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("data_center_name"); ok {
        payload["data-center-name"] = v.(string)
    }

    if v, ok := d.GetOk("data_center_uid"); ok {
        payload["data-center-uid"] = v.(string)
    }

    if v, ok := d.GetOk("uri"); ok {
        payload["uri"] = v.(string)
    }

    if v, ok := d.GetOk("uid_in_data_center"); ok {
        payload["uid-in-data-center"] = v.(string)
    }

    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
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

    if v, ok := d.GetOk("groups"); ok {
        payload["groups"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("ignore_warnings"); ok {
        payload["ignore-warnings"] = v.(bool)
    }

    if v, ok := d.GetOkExists("ignore_errors"); ok {
        payload["ignore-errors"] = v.(bool)
    }

    AddDataCenterObjectRes, _ := client.ApiCall("add-data-center-object", payload, client.GetSessionID(), true, false)
    if !AddDataCenterObjectRes.Success {
        return fmt.Errorf(AddDataCenterObjectRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementAddDataCenterObject(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

