package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementGetPlatform() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGetPlatform,
		Read:   readManagementGetPlatform,
		Delete: deleteManagementGetPlatform,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Gateway, cluster or Check Point host name.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementGetPlatform(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	GetPlatformRes, _ := client.ApiCall("get-platform", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !GetPlatformRes.Success {
		return fmt.Errorf(GetPlatformRes.ErrorMsg)
	}

	d.SetId("get-platform" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(GetPlatformRes.GetData()))
	return readManagementGetPlatform(d, m)
}

func readManagementGetPlatform(d *schema.ResourceData, m interface{}) error {


	return nil
}

func deleteManagementGetPlatform(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
