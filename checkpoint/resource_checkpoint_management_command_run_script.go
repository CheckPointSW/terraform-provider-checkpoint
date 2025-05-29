package checkpoint

import (
	"encoding/json"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementRunScript() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRunScript,
		Read:   readManagementRunScript,
		Delete: deleteManagementRunScript,
		Schema: map[string]*schema.Schema{
			"script_name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Script name.",
			},
			"script": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Script body.",
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
			"args": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Script arguments.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Comments string.",
			},
			"tasks": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Command asynchronous task unique identifiers",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"response": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Response message in JSON format",
			},
		},
	}
}

func createManagementRunScript(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("script_name"); ok {
		payload["script-name"] = v.(string)
	}

	if v, ok := d.GetOk("script"); ok {
		payload["script"] = v.(string)
	}

	if v, ok := d.GetOk("targets"); ok {
		payload["targets"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("args"); ok {
		payload["args"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		payload["comments"] = v.(string)
	}

	runScriptRes, err := client.ApiCall("run-script", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !runScriptRes.Success {
		return fmt.Errorf(runScriptRes.ErrorMsg)
	}

	taskIds := resolveTaskIds(runScriptRes.GetData())
	_ = d.Set("tasks", taskIds)

	var showTaskPayload = map[string]interface{}{}
	showTaskPayload["task-id"] = taskIds
	showTaskPayload["details-level"] = "full"
	showTaskRes, err := client.ApiCallSimple("show-task", showTaskPayload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTaskRes.Success {
		return fmt.Errorf(showTaskRes.ErrorMsg)
	}
	jsonResponse, err := json.Marshal(showTaskRes.GetData())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if jsonResponse != nil {
		_ = d.Set("response", string(jsonResponse))
	}

	d.SetId("run-script-" + acctest.RandString(10))

	return readManagementRunScript(d, m)
}

func readManagementRunScript(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementRunScript(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
