package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementResetSic() *schema.Resource {
	return &schema.Resource{
		Create: createManagementResetSic,
		Read:   readManagementResetSic,
		Delete: deleteManagementResetSic,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Gateway, cluster member or Check Point host name.",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operation status.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementResetSic(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	ResetSicRes, _ := client.ApiCall("reset-sic", payload, client.GetSessionID(), true, false)
	if !ResetSicRes.Success {
		return fmt.Errorf(ResetSicRes.ErrorMsg)
	}

	d.SetId("reset-sic" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(ResetSicRes.GetData()))
	resetSicStatusProfile := ResetSicRes.GetData()

	log.Println("Read ResetSicStatus - Show JSON = ", resetSicStatusProfile)

	if v := resetSicStatusProfile["message"]; v != nil {
		_ = d.Set("message", v)
	}

	return readManagementResetSic(d, m)
}

func readManagementResetSic(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementResetSic(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
