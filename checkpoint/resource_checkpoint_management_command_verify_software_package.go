package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementVerifySoftwarePackage() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVerifySoftwarePackage,
		Read:   readManagementVerifySoftwarePackage,
		Delete: deleteManagementVerifySoftwarePackage,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The name of the software package.",
			},
			"targets": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "On what targets to execute this command. Targets may be identified by their name, or object unique identifier.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"concurrency_limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "The number of targets, on which the same package is installed at the same time.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementVerifySoftwarePackage(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("concurrency_limit"); ok {
		payload["concurrency-limit"] = v.(int)
	}

	VerifySoftwarePackageRes, _ := client.ApiCall("verify-software-package", payload, client.GetSessionID(), true, false)
	if !VerifySoftwarePackageRes.Success {
		return fmt.Errorf(VerifySoftwarePackageRes.ErrorMsg)
	}

	d.SetId("verify-software-package-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(VerifySoftwarePackageRes.GetData()))

	return readManagementVerifySoftwarePackage(d, m)
}

func readManagementVerifySoftwarePackage(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementVerifySoftwarePackage(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
