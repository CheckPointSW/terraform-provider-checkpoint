package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementAddThreatProtections() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAddThreatProtections,
		Read:   readManagementAddThreatProtections,
		Delete: deleteManagementAddThreatProtections,
		Schema: map[string]*schema.Schema{
			"package_format": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Protections package format.",
			},
			"package_path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Protections package path.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementAddThreatProtections(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("package_format"); ok {
		payload["package-format"] = v.(string)
	}

	if v, ok := d.GetOk("package_path"); ok {
		payload["package-path"] = v.(string)
	}

	AddThreatProtectionsRes, _ := client.ApiCall("add-threat-protections", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !AddThreatProtectionsRes.Success {
		return fmt.Errorf(AddThreatProtectionsRes.ErrorMsg)
	}

	d.SetId("add-threat-protections-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(AddThreatProtectionsRes.GetData()))
	return readManagementAddThreatProtections(d, m)
}

func readManagementAddThreatProtections(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementAddThreatProtections(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
