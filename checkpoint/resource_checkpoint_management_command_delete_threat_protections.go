package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementDeleteThreatProtections() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDeleteThreatProtections,
		Read:   readManagementDeleteThreatProtections,
		Delete: deleteManagementDeleteThreatProtections,
		Schema: map[string]*schema.Schema{
			"package_format": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Protections package format.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementDeleteThreatProtections(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("package_format"); ok {
		payload["package-format"] = v.(string)
	}

	DeleteThreatProtectionsRes, _ := client.ApiCall("delete-threat-protections", payload, client.GetSessionID(), true, false)
	if !DeleteThreatProtectionsRes.Success {
		return fmt.Errorf(DeleteThreatProtectionsRes.ErrorMsg)
	}

	d.SetId("delete-threat-protections-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(DeleteThreatProtectionsRes.GetData()))
	return readManagementDeleteThreatProtections(d, m)
}

func readManagementDeleteThreatProtections(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementDeleteThreatProtections(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
