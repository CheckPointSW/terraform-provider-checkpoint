package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementHaFullSync() *schema.Resource {
	return &schema.Resource{
		Create: createManagementHaFullSync,
		Read:   readManagementHaFullSync,
		Delete: deleteManagementHaFullSync,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Peer name (Multi Domain Server, Domain Server or Security Management Server).",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Peer unique identifier (Multi Domain Server, Domain Server or Security Management Server).",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementHaFullSync(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	}

	HaFullSyncRes, _ := client.ApiCall("ha-full-sync", payload, client.GetSessionID(), true, false)
	if !HaFullSyncRes.Success {
		return fmt.Errorf(HaFullSyncRes.ErrorMsg)
	}

	d.SetId("ha-full-sync" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(HaFullSyncRes.GetData()))
	return readManagementHaFullSync(d, m)
}

func readManagementHaFullSync(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementHaFullSync(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
