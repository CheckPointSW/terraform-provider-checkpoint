package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetSyncWithUserCenter() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetSyncWithUserCenter,
		Read:   readManagementSetSyncWithUserCenter,
		Delete: deleteManagementSetSyncWithUserCenter,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Synchronize information once a day.",
			},
		},
	}
}

func createManagementSetSyncWithUserCenter(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("enabled"); ok {
		payload["enabled"] = v.(bool)
	}

	SetSyncWithUserCenterRes, err := client.ApiCallSimple("set-sync-with-user-center", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetSyncWithUserCenterRes.Success {
		return fmt.Errorf(SetSyncWithUserCenterRes.ErrorMsg)
	}

	d.SetId("set-sync-with-user-center-" + acctest.RandString(10))
	return readManagementSetSyncWithUserCenter(d, m)
}

func readManagementSetSyncWithUserCenter(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetSyncWithUserCenter(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
