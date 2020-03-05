package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementPutFile() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementPutFile,
            Read:   readManagementPutFile,
            Delete: deleteManagementPutFile,
            Schema: map[string]*schema.Schema{ 
            "targets": {
                Type:        schema.TypeSet,
                Required:    true,
                ForceNew:    true,
                Description: "On what targets to execute this command. Targets may be identified by their name, or object unique identifier.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "file_content": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
            },
            "file_name": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
            },
            "file_path": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
            },
            "comments": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "Comments string.",
            },
        },
    }
}

func createManagementPutFile(d *schema.ResourceData, m interface{}) error {
    return readManagementPutFile(d, m)
}

func readManagementPutFile(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("targets"); ok {
        payload["targets"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("file_content"); ok {
        payload["file-content"] = v.(string)
    }

    if v, ok := d.GetOk("file_name"); ok {
        payload["file-name"] = v.(string)
    }

    if v, ok := d.GetOk("file_path"); ok {
        payload["file-path"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    PutFileRes, _ := client.ApiCall("put-file", payload, client.GetSessionID(), true, false)
    if !PutFileRes.Success {
        return fmt.Errorf(PutFileRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementPutFile(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

