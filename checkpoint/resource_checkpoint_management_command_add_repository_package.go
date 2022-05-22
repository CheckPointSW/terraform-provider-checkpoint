package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementAddRepositoryPackage() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAddRepositoryPackage,
		Read:   readManagementAddRepositoryPackage,
		Delete: deleteManagementAddRepositoryPackage,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the repository package.",
			},
			"path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The path of the repository package.<br><font color=\"red\">Required only for</font> adding package from local.",
			},
			"source": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The source of the repository package.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementAddRepositoryPackage(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("path"); ok {
		payload["path"] = v.(string)
	}

	if v, ok := d.GetOk("source"); ok {
		payload["source"] = v.(string)
	}

	AddRepositoryPackageRes, _ := client.ApiCall("add-repository-package", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !AddRepositoryPackageRes.Success {
		return fmt.Errorf(AddRepositoryPackageRes.ErrorMsg)
	}

	d.SetId("add-repository-package" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(AddRepositoryPackageRes.GetData()))
	return readManagementAddRepositoryPackage(d, m)
}

func readManagementAddRepositoryPackage(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementAddRepositoryPackage(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
