package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementInstallLsmSettings() *schema.Resource {
	return &schema.Resource{
		Create: createManagementInstallLsmSettings,
		Read:   readManagementInstallLsmSettings,
		Delete: deleteManagementInstallLsmSettings,
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
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementInstallLsmSettings(d *schema.ResourceData, m interface{}) error {
	return readManagementInstallLsmSettings(d, m)
}

func readManagementInstallLsmSettings(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	InstallLsmSettingsRes, _ := client.ApiCall("install-lsm-settings", payload, client.GetSessionID(), true, false)
	if !InstallLsmSettingsRes.Success {
		return fmt.Errorf(InstallLsmSettingsRes.ErrorMsg)
	}

	d.SetId("install-lsm-settings" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(InstallLsmSettingsRes.GetData()))
	return nil
}

func deleteManagementInstallLsmSettings(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
