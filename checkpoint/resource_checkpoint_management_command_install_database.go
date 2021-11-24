package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementInstallDatabase() *schema.Resource {
	return &schema.Resource{
		Create: createManagementInstallDatabase,
		Read:   readManagementInstallDatabase,
		Delete: deleteManagementInstallDatabase,
		Schema: map[string]*schema.Schema{
			"targets": {
				Type:        schema.TypeSet,
				Required:    true,
				ForceNew:    true,
				Description: "Check Point host(s) with one or more Management Software Blades enabled. The targets can be identified by their name or unique identifier.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"tasks": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Command asynchronous task unique identifiers",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createManagementInstallDatabase(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	InstallDatabaseRes, _ := client.ApiCall("install-database", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !InstallDatabaseRes.Success {
		return fmt.Errorf(InstallDatabaseRes.ErrorMsg)
	}

	d.SetId("install-database-" + acctest.RandString(10))

	_ = d.Set("tasks", resolveTaskIds(InstallDatabaseRes.GetData()))

	return readManagementInstallDatabase(d, m)
}

func readManagementInstallDatabase(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementInstallDatabase(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
