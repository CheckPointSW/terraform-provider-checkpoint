package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
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
		},
	}
}

func createManagementPublish(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	var uid string
	var payload = make(map[string]interface{})

	if v, ok := d.GetOk("uid"); ok {
		uid = v.(string)
		payload["uid"] = uid
		log.Println("publish other session uid - ", uid)
	} else {
		// Publish current session
		s, err := GetSession()
		if err != nil {
			return err
		}
		uid = s.Uid
		log.Println("publish current session uid - ", uid)
	}
	publishRes, _ := client.ApiCall("publish", payload, client.GetSessionID(), true, false)
	if !publishRes.Success {
		return fmt.Errorf(publishRes.ErrorMsg)
	}

	d.SetId(uid)
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
