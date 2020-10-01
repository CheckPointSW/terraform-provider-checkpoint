package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementLogin() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLogin,
		Read:   readManagementLogin,
		Delete: deleteManagementLogin,
		Schema: map[string]*schema.Schema{
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Administrator user name.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Administrator password.",
			},
			"continue_last_session": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "When 'continue-last-session' is set to 'True', the new session would continue where the last session was stopped. This option is available when the administrator has only one session that can be continued. If there is more than one session, see 'switch-session' API.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Use domain to login to specific domain. Domain can be identified by name or UID.",
			},
			"enter_last_published_session": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Login to the last published session. Such login is done with the Read Only permissions.",
			},
			"read_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Login with Read Only permissions. This parameter is not considered in case continue-last-session is true.",
			},
			"session_comments": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Session comments.",
			},
			"session_description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Session description.",
			},
			"session_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Session unique name.",
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Session expiration timeout in seconds. Default 600 seconds.",
			},
		},
	}
}

func createManagementLogin(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = make(map[string]interface{})

	if v, ok := d.GetOk("user"); ok {
		payload["user"] = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		payload["password"] = v.(string)
	}
	if v, ok := d.GetOk("continue_last_session"); ok {
		payload["continue-last-session"] = v.(bool)
	}
	if v, ok := d.GetOk("domain"); ok {
		payload["domain"] = v.(string)
	}
	if v, ok := d.GetOk("enter_last_published_session"); ok {
		payload["enter-last-published-session"] = v.(bool)
	}
	if v, ok := d.GetOk("read_only"); ok {
		payload["read-only"] = v.(bool)
	}
	if v, ok := d.GetOk("session_comments"); ok {
		payload["session-comments"] = v.(string)
	}
	if v, ok := d.GetOk("session_description"); ok {
		payload["session-description"] = v.(string)
	}
	if v, ok := d.GetOk("session_name"); ok {
		payload["session-name"] = v.(string)
	}
	if v, ok := d.GetOk("session_timeout"); ok {
		payload["session-timeout"] = v.(int)
	}

	loginRes, _ := client.ApiCall("login", payload, "", true, false)
	if !loginRes.Success {
		return fmt.Errorf(loginRes.ErrorMsg)
	}

	d.SetId(loginRes.GetData()["sid"].(string))
	return readManagementLogin(d, m)
}

func readManagementLogin(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementLogin(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
