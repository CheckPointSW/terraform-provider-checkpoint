package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementRepositoryPackage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementRepositoryPackageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the repository package.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique identifier of the 'show-repository-package' task. Use the 'show-task' command to check the progress of the task.",
			},
		},
	}
}

func dataSourceManagementRepositoryPackageRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)

	payload := make(map[string]interface{})

	payload["name"] = name

	showRepositoryPackageRes, err := client.ApiCall("show-repository-package", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRepositoryPackageRes.Success {
		return fmt.Errorf(showRepositoryPackageRes.ErrorMsg)
	}

	RepositoryPackage := showRepositoryPackageRes.GetData()

	log.Println("Read RepositoryPackage - Show JSON = ", RepositoryPackage)

	if v := RepositoryPackage["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	} else {
		d.SetId("ff")
	}

	if v, ok := d.GetOk("name"); ok {
		_ = d.Set("name", v.(string))
	}

	if v, ok := d.GetOk("task_id"); ok {
		_ = d.Set("task_id", v.(string))
	}

	return nil
}
