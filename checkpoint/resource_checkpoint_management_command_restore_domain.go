package checkpoint

import (
    "fmt"
    checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
    "github.com/hashicorp/terraform/helper/schema"
)

    func resourceManagementRestoreDomain() *schema.Resource {   
        return &schema.Resource{
            Create: createManagementRestoreDomain,
            Read:   readManagementRestoreDomain,
            Delete: deleteManagementRestoreDomain,
            Schema: map[string]*schema.Schema{ 
            "file_path": {
                Type:        schema.TypeString,
                Required:    true,
                ForceNew:    true,
                Description: "Path to the backup file to be restored. <br>Should be the full file path (example, \"/var/log/domain1_backup.tgz\").",
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
            "verify_only": {
                Type:        schema.TypeBool,
                Optional:    true,
                ForceNew:    true,
                Description: "If true, verify that the import operation is valid for this input file and this environment <br>Note: Restore operation will not be executed.",
            },
        },
    }
}

func createManagementRestoreDomain(d *schema.ResourceData, m interface{}) error {
    return readManagementRestoreDomain(d, m)
}

func readManagementRestoreDomain(d *schema.ResourceData, m interface{}) error {

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

    if v, ok := d.GetOkExists("verify_only"); ok {
        payload["verify-only"] = v.(bool)
    }

    RestoreDomainRes, _ := client.ApiCall("restore-domain", payload, client.GetSessionID(), true, false)
    if !RestoreDomainRes.Success {
        return fmt.Errorf(RestoreDomainRes.ErrorMsg)
    }

    d.SetId("ff")
    return nil
}

func deleteManagementRestoreDomain(d *schema.ResourceData, m interface{}) error {

    d.SetId("")
    return nil
}

