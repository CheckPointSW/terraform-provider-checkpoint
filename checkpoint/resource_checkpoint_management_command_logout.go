package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementLogout() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLogout,
		Read:   readManagementLogout,
		Delete: deleteManagementLogout,
		Schema: map[string]*schema.Schema{
			"message": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Operation status.",
			},
		},
	}
}

func createManagementLogout(d *schema.ResourceData, m interface{}) error {

	return readManagementLogout(d, m)
}

func readManagementLogout(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	logoutRes, _ := client.ApiCall("logout", make(map[string]interface{}), "", true, false)
	if !logoutRes.Success {
		return fmt.Errorf(logoutRes.ErrorMsg)
	}
	// Set Schema UID = Session UID
	d.SetId(logoutRes.GetData()["message"].(string))
	return nil
}

func deleteManagementLogout(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
