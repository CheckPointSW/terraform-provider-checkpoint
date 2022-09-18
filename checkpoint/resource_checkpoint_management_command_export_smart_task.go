package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementExportSmartTask() *schema.Resource {
	return &schema.Resource{
		Create: createManagementExportSmartTask,
		Read:   readManagementExportSmartTask,
		Delete: deleteManagementExportSmartTask,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Name of task to be exported.",
			},
			"file_path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Path to the SmartTask file to be exported. <br>Should be the full file path (example, \"/home/admin/exported-smart-task.txt)\".<br>If no path was inserted the default will be: \"/var/log/<task_name>.txt\".",
			},
		},
	}
}

func createManagementExportSmartTask(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("file_path"); ok {
		payload["file-path"] = v.(string)
	}

	ExportSmartTaskRes, err := client.ApiCall("export-smart-task", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !ExportSmartTaskRes.Success {
		return fmt.Errorf(ExportSmartTaskRes.ErrorMsg)
	}

	exportSmartTask := ExportSmartTaskRes.GetData()

	if v := exportSmartTask["file-path"]; v != nil {
		_ = d.Set("file_path", v)
	}

	d.SetId("export-smart-task-" + acctest.RandString(10))
	return readManagementExportSmartTask(d, m)
}

func readManagementExportSmartTask(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementExportSmartTask(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
