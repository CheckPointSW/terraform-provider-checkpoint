package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementAddApiKey() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAddApiKey,
		Read:   readManagementAddApiKey,
		Delete: deleteManagementAddApiKey,
		Schema: map[string]*schema.Schema{
			"admin_uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Administrator uid to generate API key for.",
			},
			"admin_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Administrator name to generate API key for.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Represents the API Key to be used for Login.",
			},
		},
	}
}

func createManagementAddApiKey(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("admin_uid"); ok {
		payload["admin-uid"] = v.(string)
	}

	if v, ok := d.GetOk("admin_name"); ok {
		payload["admin-name"] = v.(string)
	}

	AddApiKeyRes, _ := client.ApiCall("add-api-key", payload, client.GetSessionID(), true, false)
	if !AddApiKeyRes.Success {
		return fmt.Errorf(AddApiKeyRes.ErrorMsg)
	}

	d.SetId("add-api-key-" + acctest.RandString(10))
	if v, ok := AddApiKeyRes.GetData()["api-key"]; ok {
		_ = d.Set("api_key", v.(string))
	}
	return readManagementAddApiKey(d, m)
}

func readManagementAddApiKey(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementAddApiKey(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
