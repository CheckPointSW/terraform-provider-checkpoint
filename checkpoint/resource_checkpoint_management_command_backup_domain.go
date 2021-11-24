package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementBackupDomain() *schema.Resource {
	return &schema.Resource{
		Create: createManagementBackupDomain,
		Read:   readManagementBackupDomain,
		Delete: deleteManagementBackupDomain,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain can be identified by name or UID.",
			},
			"file_path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Path in which the backup domain data will be saved. <br>Should be the directory path or the full file path with \".tgz\" <br>If no path was inserted the default will be: \"/var/log/&lt;domain name&gt;_&lt;date&gt;.tgz\".",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementBackupDomain(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("domain"); ok {
		payload["domain"] = v.(string)
	}

	if v, ok := d.GetOk("file_path"); ok {
		payload["file-path"] = v.(string)
	}

	BackupDomainRes, _ := client.ApiCall("backup-domain", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !BackupDomainRes.Success {
		return fmt.Errorf(BackupDomainRes.ErrorMsg)
	}

	d.SetId("backup-domain-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(BackupDomainRes.GetData()))
	return readManagementBackupDomain(d, m)
}

func readManagementBackupDomain(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementBackupDomain(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
