package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementLsmRunScript() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLsmRunScript,
		Read:   readManagementLsmRunScript,
		Delete: deleteManagementLsmRunScript,
		Schema: map[string]*schema.Schema{
			"script_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The entire content of the script encoded in Base64.",
			},
			"script": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The entire content of the script.",
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
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementLsmRunScript(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("script_base64"); ok {
		payload["script-base64"] = v.(string)
	}

	if v, ok := d.GetOk("script"); ok {
		payload["script"] = v.(string)
	}

	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	LsmRunScriptRes, _ := client.ApiCall("lsm-run-script", payload, client.GetSessionID(), true, false)
	if !LsmRunScriptRes.Success {
		return fmt.Errorf(LsmRunScriptRes.ErrorMsg)
	}

	d.SetId("lsm-run-script" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(LsmRunScriptRes.GetData()))
	return readManagementLsmRunScript(d, m)
}

func readManagementLsmRunScript(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementLsmRunScript(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
