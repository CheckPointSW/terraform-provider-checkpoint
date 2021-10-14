package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementGetAttachment() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGetAttachment,
		Read:   readManagementGetAttachment,
		Delete: deleteManagementGetAttachment,
		Schema: map[string]*schema.Schema{
			"attachment_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Attachment identifier from a log record.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Log uid from a log record.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementGetAttachment(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("attachment_id"); ok {
		payload["attachment-id"] = v.(string)
	}

	if v, ok := d.GetOk("uid"); ok {
		payload["id"] = v.(string)
	}

	GetAttachmentRes, _ := client.ApiCall("get-attachment", payload, client.GetSessionID(), true, false)
	if !GetAttachmentRes.Success {
		return fmt.Errorf(GetAttachmentRes.ErrorMsg)
	}

	d.SetId("get-attachment" + acctest.RandString(10))

	_ = d.Set("task_id", resolveTaskId(GetAttachmentRes.GetData()))

	return readManagementGetAttachment(d, m)
}

func readManagementGetAttachment(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementGetAttachment(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
