package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementUpdateUpdatableObjectsRepositoryContent() *schema.Resource {
	return &schema.Resource{
		Create: createManagementUpdateUpdatableObjectsRepositoryContent,
		Read:   readManagementUpdateUpdatableObjectsRepositoryContent,
		Delete: deleteManagementUpdateUpdatableObjectsRepositoryContent,
		Schema: map[string]*schema.Schema{
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementUpdateUpdatableObjectsRepositoryContent(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	UpdateUpdatableObjectsRepositoryContentRes, _ := client.ApiCall("update-updatable-objects-repository-content", payload, client.GetSessionID(), true, false)
	if !UpdateUpdatableObjectsRepositoryContentRes.Success {
		return fmt.Errorf(UpdateUpdatableObjectsRepositoryContentRes.ErrorMsg)
	}

	d.SetId("update-updatable-objects-repository-content-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(UpdateUpdatableObjectsRepositoryContentRes.GetData()))

	return readManagementUpdateUpdatableObjectsRepositoryContent(d, m)
}

func readManagementUpdateUpdatableObjectsRepositoryContent(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementUpdateUpdatableObjectsRepositoryContent(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
