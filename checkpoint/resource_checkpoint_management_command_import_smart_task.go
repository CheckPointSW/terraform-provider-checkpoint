package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementImportSmartTask() *schema.Resource {
	return &schema.Resource{
		Create: createManagementImportSmartTask,
		Read:   readManagementImportSmartTask,
		Delete: deleteManagementImportSmartTask,
		Schema: map[string]*schema.Schema{
			"file_path": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Path to the SmartTask file to be imported. <br>Should be the full file path (example, \"/home/admin/exported-smart-task.txt\").",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operation status.",
			},
		},
	}
}

func createManagementImportSmartTask(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("file_path"); ok {
		payload["file-path"] = v.(string)
	}

	ImportSmartTaskRes, err := client.ApiCall("import-smart-task", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !ImportSmartTaskRes.Success {
		return fmt.Errorf(ImportSmartTaskRes.ErrorMsg)
	}

	importSmartTask := ImportSmartTaskRes.GetData()

	if v := importSmartTask["message"]; v != nil {
		_ = d.Set("message", v)
	}

	d.SetId("import-smart-task-" + acctest.RandString(10))
	return readManagementImportSmartTask(d, m)
}

func readManagementImportSmartTask(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementImportSmartTask(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
