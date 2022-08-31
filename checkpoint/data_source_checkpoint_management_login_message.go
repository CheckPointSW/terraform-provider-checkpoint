package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementLoginMessage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementLoginMessageRead,
		Schema: map[string]*schema.Schema{
			"header": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Login message header.",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Login message body.",
			},
			"show_message": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether to show login message.",
			},
			"warning": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Add warning sign.",
			},
		},
	}
}

func dataSourceManagementLoginMessageRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showLoginMessageRes, err := client.ApiCall("show-login-message", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		fmt.Errorf(err.Error())
	}
	if !showLoginMessageRes.Success {
		fmt.Errorf(showLoginMessageRes.ErrorMsg)
	}

	loginMessage := showLoginMessageRes.GetData()

	log.Println("Read Login Message - Show JSON = ", loginMessage)

	d.SetId("login-message-" + acctest.RandString(10))

	if v := loginMessage["header"]; v != nil {
		_ = d.Set("header", v)
	}

	if v := loginMessage["message"]; v != nil {
		_ = d.Set("message", v)
	}

	if v := loginMessage["show-message"]; v != nil {
		_ = d.Set("show_message", v)
	}

	if v := loginMessage["warning"]; v != nil {
		_ = d.Set("warning", v)
	}

	return nil
}
