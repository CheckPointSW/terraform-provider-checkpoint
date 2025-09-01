package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementRunAppControlUpdate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRunAppControlUpdate,
		Read:   readManagementRunAppControlUpdate,
		Delete: deleteManagementRunAppControlUpdate,
		Schema: map[string]*schema.Schema{
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementRunAppControlUpdate(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	RunAppControlUpdateRes, err := client.ApiCallSimple("run-app-control-update", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !RunAppControlUpdateRes.Success {
		return fmt.Errorf(RunAppControlUpdateRes.ErrorMsg)
	}

	d.SetId("app-control-update-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(RunAppControlUpdateRes.GetData()))
	return readManagementRunAppControlUpdate(d, m)
}

func readManagementRunAppControlUpdate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementRunAppControlUpdate(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
