package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementMigrateExportDomain() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementMigrateExportDomain,
            Read:   readManagementMigrateExportDomain,
            Delete: deleteManagementMigrateExportDomain,
            Schema: map[string]*schema.Schema{ 
            "domain": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Domain can be identified by name or UID.<br><font color=\"red\">Required only for</font> exporting domain from Multi-Domain Server.",
            },
            "file_path": {
                Type:        schema.TypeString,
                Optional:    true,
                ForceNew:    true,
                Description: "Path in which the exported domain data will be saved. <br>Should be the directory path or the full file path with \".tgz\" <br>If no path was inserted the default will be: \"/var/log/&lt;domain name&gt;_&lt;date&gt;.tgz\".",
            },
            "include_logs": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Export logs.",
            },
        },
    }
}

func createManagementMigrateExportDomain(d *schema.ResourceData, m interface{}) error {
    return readManagementMigrateExportDomain(d, m)
}

func readManagementMigrateExportDomain(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("domain"); ok {
        payload["domain"] = v.(string)
    }

    if v, ok := d.GetOk("file_path"); ok {
        payload["file-path"] = v.(string)
    }

    if v, ok := d.GetOkExists("include_logs"); ok {
        payload["include-logs"] = v.(bool)
    }

    MigrateExportDomainRes, _ := client.ApiCall("migrate-export-domain", payload, client.GetSessionID(), true, false)
    if !MigrateExportDomainRes.Success {
        return fmt.Errorf(MigrateExportDomainRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementMigrateExportDomain(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

