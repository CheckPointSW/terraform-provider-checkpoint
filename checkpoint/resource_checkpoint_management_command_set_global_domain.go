package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementSetGlobalDomain() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementSetGlobalDomain,
            Read:   readManagementSetGlobalDomain,
            Delete: deleteManagementSetGlobalDomain,
            Schema: map[string]*schema.Schema{ 
            "name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Object name.",
            },
            "servers": {
                Type:        schema.TypeMap,
                Optional:    true,
                Description: "Multi Domain Servers. When the field is provided, 'set-global-domain' command is executed asynchronously.",
                ForceNew:    true,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "add": {
                            Type:        schema.TypeSet,
                            Optional:    true,
                            Description: "Adds to collection of values",
                            Elem: &schema.Schema{
                                Type: schema.TypeString,
                            },
                        },
                    },
                },
            },
            "tags": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: "Collection of tag identifiers. Note: The list of tags can not be modified in a singlecommand together with the domain servers. To modify tags, please use the separate 'set-global-domain' command, without providing the list of domain servers.",
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

func createManagementSetGlobalDomain(d *schema.ResourceData, m interface{}) error {
    return readManagementSetGlobalDomain(d, m)
}

func readManagementSetGlobalDomain(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("name"); ok {
        payload["name"] = v.(string)
    }

    if _, ok := d.GetOk("servers"); ok {

        res := make(map[string]interface{})

        if v, ok := d.GetOk("servers.add"); ok {
            res["add"] = v
        }
        payload["servers"] = res
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

    SetGlobalDomainRes, _ := client.ApiCall("set-global-domain", payload, client.GetSessionID(), true, false)
    if !SetGlobalDomainRes.Success {
        return fmt.Errorf(SetGlobalDomainRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementSetGlobalDomain(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

