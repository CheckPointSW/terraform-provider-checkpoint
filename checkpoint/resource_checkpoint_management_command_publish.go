package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementPublish() *schema.Resource {
	return &schema.Resource{
		Create: createManagementPublish,
		Read:   readManagementPublish,
		Delete: deleteManagementPublish,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Session unique identifier. Specify it to publish a different session than the one you currently use.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
			"triggers": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Triggers a publish if there are any changes to objects in this list.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createManagementPublish(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = make(map[string]interface{})

	if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	}

	publishRes, _ := client.ApiCall("publish", payload, client.GetSessionID(), true, false)
	if !publishRes.Success {
		return fmt.Errorf(publishRes.ErrorMsg)
	}

	d.SetId("publish-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(publishRes.GetData()))

	return readManagementPublish(d, m)
}

func readManagementPublish(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementPublish(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
