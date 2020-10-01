package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementRevertToRevision() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRevertToRevision,
		Read:   readManagementRevertToRevision,
		Delete: deleteManagementRevertToRevision,
		Schema: map[string]*schema.Schema{
			"to_session": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Session unique identifier. Specify the session  id you would like to revert your database to.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementRevertToRevision(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("to_session"); ok {
		payload["to-session"] = v.(string)
	}

	RevertToRevisionRes, _ := client.ApiCall("revert-to-revision", payload, client.GetSessionID(), true, false)
	if !RevertToRevisionRes.Success {
		return fmt.Errorf(RevertToRevisionRes.ErrorMsg)
	}

	d.SetId("revert-to-revision-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(RevertToRevisionRes.GetData()))

	return readManagementRevertToRevision(d, m)
}

func readManagementRevertToRevision(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementRevertToRevision(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
