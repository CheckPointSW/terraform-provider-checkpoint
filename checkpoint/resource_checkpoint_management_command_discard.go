package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementDiscard() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDiscard,
		Read:   readManagementDiscard,
		Delete: deleteManagementDiscard,
		Schema: map[string]*schema.Schema{},
	}
}

func createManagementDiscard(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	DiscardRes, _ := client.ApiCall("discard", payload, client.GetSessionID(), true, false)
	if !DiscardRes.Success {
		return fmt.Errorf(DiscardRes.ErrorMsg)
	}

	d.SetId("discard-" + acctest.RandString(10))
	return readManagementDiscard(d, m)
}

func readManagementDiscard(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementDiscard(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
