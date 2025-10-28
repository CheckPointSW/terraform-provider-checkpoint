package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementSmartConsoleIdleTimeout() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSmartConsoleIdleTimeoutRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to perform logout after being idle.",
			},
			"timeout_duration": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of minutes that the SmartConsole will automatically logout after being idle.<br>Updating the interval will take effect only on the next login.",
			},
		},
	}
}

func dataSourceManagementSmartConsoleIdleTimeoutRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	smartConsoleIdleTimeoutRes, err := client.ApiCallSimple("show-smart-console-idle-timeout", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !smartConsoleIdleTimeoutRes.Success {
		return fmt.Errorf(smartConsoleIdleTimeoutRes.ErrorMsg)
	}
	smartConsoleIdleTimeoutData := smartConsoleIdleTimeoutRes.GetData()

	if v := smartConsoleIdleTimeoutData["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := smartConsoleIdleTimeoutData["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := smartConsoleIdleTimeoutData["timeout-duration"]; v != nil {
		_ = d.Set("timeout_duration", v)
	}

	return nil

}
