package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetLoginMessage() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetLoginMessage,
		Read:   readManagementSetLoginMessage,
		Delete: deleteManagementSetLoginMessage,
		Schema: map[string]*schema.Schema{
			"header": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "commonLoginLogic message header.",
			},
			"message": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "commonLoginLogic message body.",
			},
			"show_message": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Whether to show login message.",
			},
			"warning": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Add warning sign.",
			},
		},
	}
}

func createManagementSetLoginMessage(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("header"); ok {
		payload["header"] = v.(string)
	}

	if v, ok := d.GetOk("message"); ok {
		payload["message"] = v.(string)
	}

	if v, ok := d.GetOkExists("show_message"); ok {
		payload["show-message"] = v.(bool)
	}

	if v, ok := d.GetOkExists("warning"); ok {
		payload["warning"] = v.(bool)
	}

	SetLoginMessageRes, _ := client.ApiCall("set-login-message", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !SetLoginMessageRes.Success {
		return fmt.Errorf(SetLoginMessageRes.ErrorMsg)
	}

	d.SetId("set-login-message-" + acctest.RandString(10))
	return readManagementSetLoginMessage(d, m)
}

func readManagementSetLoginMessage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetLoginMessage(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
