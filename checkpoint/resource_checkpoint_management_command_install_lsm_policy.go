package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementInstallLsmPolicy() *schema.Resource {
	return &schema.Resource{
		Create: createManagementInstallLsmPolicy,
		Read:   readManagementInstallLsmPolicy,
		Delete: deleteManagementInstallLsmPolicy,
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

func createManagementInstallLsmPolicy(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	InstallLsmPolicyRes, _ := client.ApiCall("install-lsm-policy", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !InstallLsmPolicyRes.Success {
		return fmt.Errorf(InstallLsmPolicyRes.ErrorMsg)
	}

	d.SetId("install-lsm-policy" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(InstallLsmPolicyRes.GetData()))

	return readManagementInstallLsmPolicy(d, m)
}

func readManagementInstallLsmPolicy(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementInstallLsmPolicy(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
