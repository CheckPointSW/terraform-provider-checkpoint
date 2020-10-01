package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
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
	client := m.(*checkpoint.ApiClient)

	logoutRes, _ := client.ApiCall("logout", make(map[string]interface{}), "", true, false)
	if !logoutRes.Success {
		return fmt.Errorf(logoutRes.ErrorMsg)
	}

	d.SetId("logout-" + acctest.RandString(10))
	return readManagementLogout(d, m)
}

func readManagementLogout(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementLogout(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
