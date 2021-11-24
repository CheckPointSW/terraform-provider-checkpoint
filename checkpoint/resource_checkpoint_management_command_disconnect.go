package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementDisconnect() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDisconnect,
		Read:   readManagementDisconnect,
		Delete: deleteManagementDisconnect,
		Schema: map[string]*schema.Schema{
			"discard": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Discard all changes committed during the session.",
			},
		},
	}
}

func createManagementDisconnect(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("discard"); ok {
		payload["discard"] = v.(bool)
	}

	DisconnectRes, _ := client.ApiCall("disconnect", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !DisconnectRes.Success {
		return fmt.Errorf(DisconnectRes.ErrorMsg)
	}

	d.SetId("discard-" + acctest.RandString(10))
	return readManagementDisconnect(d, m)
}

func readManagementDisconnect(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementDisconnect(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
