package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementExport() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementExport,
            Read:   readManagementExport,
            Delete: deleteManagementExport,
            Schema: map[string]*schema.Schema{ 
            "exclude_classes": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "exclude_topics": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "export_files_by_class": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
            },
            "include_classes": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "include_topics": {
                Type:        schema.TypeSet,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },
            "query_limit": {
                Type:        schema.TypeInt,
                Optional:    true,
                ForceNew:    true,
                Description: "N/A",
            },
        },
    }
}

func createManagementExport(d *schema.ResourceData, m interface{}) error {
    return readManagementExport(d, m)
}

func readManagementExport(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("exclude_classes"); ok {
        payload["exclude-classes"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("exclude_topics"); ok {
        payload["exclude-topics"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOkExists("export_files_by_class"); ok {
        payload["export-files-by-class"] = v.(bool)
    }

    if v, ok := d.GetOk("include_classes"); ok {
        payload["include-classes"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("include_topics"); ok {
        payload["include-topics"] = v.(*schema.Set).List()
    }

    if v, ok := d.GetOk("query_limit"); ok {
        payload["query-limit"] = v.(int)
    }

    ExportRes, _ := client.ApiCall("export", payload, client.GetSessionID(), true, false)
    if !ExportRes.Success {
        return fmt.Errorf(ExportRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementExport(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

