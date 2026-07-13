package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementTestAiGuardApiKey() *schema.Resource {
	return &schema.Resource{
		Create: createManagementTestAiGuardApiKey,
		Read:   readManagementTestAiGuardApiKey,
		Delete: deleteManagementTestAiGuardApiKey,
		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Optional Lakera project ID to validate. If provided, also verifies the project belongs to the API key.",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Validation result message.",
			},
			"success": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the API key (and optional project) is valid.",
			},
		},
	}
}

func createManagementTestAiGuardApiKey(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("project_id"); ok {
		payload["project-id"] = v.(string)
	}

	TestAiGuardApiKeyRes, err := client.ApiCall("test-ai-guard-api-key", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !TestAiGuardApiKeyRes.Success {
		return fmt.Errorf(TestAiGuardApiKeyRes.ErrorMsg)
	}

	d.SetId("test-ai-guard-api-key-" + acctest.RandString(10))
	if v := TestAiGuardApiKeyRes.GetData()["message"]; v != nil {
		_ = d.Set("message", v)
	}
	if v := TestAiGuardApiKeyRes.GetData()["success"]; v != nil {
		_ = d.Set("success", v)
	}
	return nil
}

func readManagementTestAiGuardApiKey(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementTestAiGuardApiKey(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
