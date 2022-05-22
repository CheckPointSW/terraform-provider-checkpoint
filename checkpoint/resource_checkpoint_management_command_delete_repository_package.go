package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementDeleteRepositoryPackage() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDeleteRepositoryPackage,
		Read:   readManagementDeleteRepositoryPackage,
		Delete: deleteManagementDeleteRepositoryPackage,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the software package.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementDeleteRepositoryPackage(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	DeleteRepositoryPackageRes, _ := client.ApiCall("delete-repository-package", payload, client.GetSessionID(), true, false)
	if !DeleteRepositoryPackageRes.Success {
		return fmt.Errorf(DeleteRepositoryPackageRes.ErrorMsg)
	}

	d.SetId("delete-repository-package" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(DeleteRepositoryPackageRes.GetData()))
	return readManagementDeleteRepositoryPackage(d, m)
}

func readManagementDeleteRepositoryPackage(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementDeleteRepositoryPackage(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
