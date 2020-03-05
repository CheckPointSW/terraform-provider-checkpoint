package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementMigrateImportDomain() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementMigrateImportDomain,
            Read:   readManagementMigrateImportDomain,
            Delete: deleteManagementMigrateImportDomain,
            Schema: map[string]*schema.Schema{ 
            "file_path": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Path to the exported file to be imported. <br>Should be the full file path (example, \"/var/log/domain1_exported.tgz\").",
            },
            "domain_ip_address": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "IPv4 address.<br><font color=\"red\">Required only for</font> importing Security Management Server into Multi-Domain Server.",
            },
            "domain_name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Domain name. Should be unique in the MDS.<br><font color=\"red\">Required only for</font> importing Security Management Server into Multi-Domain Server.",
            },
            "domain_server_name": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Multi Domain server name.<br><font color=\"red\">Required only for</font> importing Security Management Server into Multi-Domain Server.",
            },
            "include_logs": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "Import logs from the input package.",
            },
        },
    }
}

func createManagementMigrateImportDomain(d *schema.ResourceData, m interface{}) error {
    return readManagementMigrateImportDomain(d, m)
}

func readManagementMigrateImportDomain(d *schema.ResourceData, m interface{}) error {

    client := m.(*checkpoint.ApiClient)

    var payload = map[string]interface{}{}
    if v, ok := d.GetOk("file_path"); ok {
        payload["file-path"] = v.(string)
    }

    if v, ok := d.GetOk("domain_ip_address"); ok {
        payload["domain-ip-address"] = v.(string)
    }

    if v, ok := d.GetOk("domain_name"); ok {
        payload["domain-name"] = v.(string)
    }

    if v, ok := d.GetOk("domain_server_name"); ok {
        payload["domain-server-name"] = v.(string)
    }

    if v, ok := d.GetOkExists("include_logs"); ok {
        payload["include-logs"] = v.(bool)
    }

    MigrateImportDomainRes, _ := client.ApiCall("migrate-import-domain", payload, client.GetSessionID(), true, false)
    if !MigrateImportDomainRes.Success {
        return fmt.Errorf(MigrateImportDomainRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementMigrateImportDomain(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

