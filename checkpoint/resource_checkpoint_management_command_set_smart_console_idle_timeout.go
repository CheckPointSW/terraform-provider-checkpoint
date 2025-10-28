package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetSmartConsoleIdleTimeout() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetSmartConsoleIdleTimeout,
		Read:   readManagementSetSmartConsoleIdleTimeout,
		Delete: deleteManagementSetSmartConsoleIdleTimeout,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to perform logout after being idle.",
			},
			"timeout_duration": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Number of minutes that the SmartConsole will automatically logout after being idle.<br>Updating the interval will take effect only on the next login.",
			},
		},
	}
}

func createManagementSetSmartConsoleIdleTimeout(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOkExists("enabled"); ok {
		payload["enabled"] = v.(bool)
	}

	if v, ok := d.GetOk("timeout_duration"); ok {
		payload["timeout-duration"] = v.(int)
	}

	SetSmartConsoleIdleTimeoutRes, err := client.ApiCall("set-smart-console-idle-timeout", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetSmartConsoleIdleTimeoutRes.Success {
		return fmt.Errorf(SetSmartConsoleIdleTimeoutRes.ErrorMsg)
	}

	d.SetId("set-smart-console-idle-timeout-" + acctest.RandString(10))
	return nil
}

func readManagementSetSmartConsoleIdleTimeout(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetSmartConsoleIdleTimeout(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
