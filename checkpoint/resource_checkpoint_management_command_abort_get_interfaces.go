package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementAbortGetInterfaces() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAbortGetInterfaces,
		Read:   readManagementAbortGetInterfaces,
		Delete: deleteManagementAbortGetInterfaces,
		Schema: map[string]*schema.Schema{
			"task_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "get-interfaces task UID.",
			},
			"force_cleanup": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Forcefully abort the \"get-interfaces\" task.",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operation status",
			},
		},
	}
}

func createManagementAbortGetInterfaces(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("task_id"); ok {
		payload["task-id"] = v.(string)
	}

	if v, ok := d.GetOkExists("force_cleanup"); ok {
		payload["force-cleanup"] = v.(bool)
	}

	AbortGetInterfacesRes, err := client.ApiCall("abort-get-interfaces", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !AbortGetInterfacesRes.Success {
		return fmt.Errorf(AbortGetInterfacesRes.ErrorMsg)
	}

	abortGetInterfaces := AbortGetInterfacesRes.GetData()

	if v := abortGetInterfaces["message"]; v != nil {
		_ = d.Set("message", v)
	}

	d.SetId("abort-get-interfaces-" + acctest.RandString(10))
	return readManagementAbortGetInterfaces(d, m)
}

func readManagementAbortGetInterfaces(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementAbortGetInterfaces(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
