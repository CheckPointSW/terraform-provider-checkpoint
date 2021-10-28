package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetHaState() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetHaState,
		Read:   readManagementSetHaState,
		Delete: deleteManagementSetHaState,
		Schema: map[string]*schema.Schema{
			"new_state": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain server new state.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementSetHaState(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("new_state"); ok {
		payload["new-state"] = v.(string)
	}

	SetHaStateRes, _ := client.ApiCall("set-ha-state", payload, client.GetSessionID(), true, false)
	if !SetHaStateRes.Success {
		return fmt.Errorf(SetHaStateRes.ErrorMsg)
	}

	d.SetId("set-ha-state" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(SetHaStateRes.GetData()))

	return readManagementSetHaState(d, m)
}

func readManagementSetHaState(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetHaState(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
