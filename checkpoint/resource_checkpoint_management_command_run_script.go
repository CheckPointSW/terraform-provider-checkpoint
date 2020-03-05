package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementRunScript() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementRunScript,
            Read:   readManagementRunScript,
            Delete: deleteManagementRunScript,
            Schema: map[string]*schema.Schema{ 
            "script_name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Script name.",
            },
            "script": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Script body.",
            },
            "targets": {
                Type:        schema.TypeSet,
                Required:    true,
                ForceNew:    true,
                Description: "On what targets to execute this command. Targets may be identified by their name, or object unique identifier.",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "args": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "Script arguments.",
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

func createManagementRunScript(d *schema.ResourceData, m interface{}) error {
    return readManagementRunScript(d, m)
}

func readManagementRunScript(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("script_name"); ok {
        payload["script-name"] = v.(string)
    }

    if v, ok := d.GetOk("script"); ok {
        payload["script"] = v.(string)
    }

    if v, ok := d.GetOk("targets"); ok {
        payload["targets"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("args"); ok {
        payload["args"] = v.(string)
    }

    if v, ok := d.GetOk("comments"); ok {
        payload["comments"] = v.(string)
    }

    RunScriptRes, _ := client.ApiCall("run-script", payload, client.GetSessionID(), true, false)
    if !RunScriptRes.Success {
        return fmt.Errorf(RunScriptRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementRunScript(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

