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
			"triggers": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Triggers a logout if there are any changes to objects in this list.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createManagementLogout(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	logoutRes, _ := client.ApiCall("logout", make(map[string]interface{}), client.GetSessionID(), true, client.IsProxyUsed())
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
