package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetGatewayGlobalUse() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetGatewayGlobalUse,
		Read:   readManagementSetGatewayGlobalUse,
		Delete: deleteManagementSetGatewayGlobalUse,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				ForceNew:    true,
				Description: "Indicates whether global use is enabled on the target.",
			},
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "On what target to execute this command. Target may be identified by its object name, or object unique identifier.",
			},
		},
	}
}

func createManagementSetGatewayGlobalUse(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("enabled"); ok {
		payload["enabled"] = v.(bool)
	}

	if v, ok := d.GetOk("target"); ok {
		payload["target"] = v.(string)
	}

	SetGatewayGlobalUseRes, _ := client.ApiCall("set-gateway-global-use", payload, client.GetSessionID(), true, false)
	if !SetGatewayGlobalUseRes.Success {
		return fmt.Errorf(SetGatewayGlobalUseRes.ErrorMsg)
	}

	res := SetGatewayGlobalUseRes.GetData()

	_ = d.Set("uid", res["uid"])
	d.SetId(res["uid"].(string))
	return readManagementSetGatewayGlobalUse(d, m)
}

func readManagementSetGatewayGlobalUse(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetGatewayGlobalUse(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
