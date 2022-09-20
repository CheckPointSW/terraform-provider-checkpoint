package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementLoginToDomain() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLoginToDomain,
		Read:   readManagementLoginToDomain,
		Delete: deleteManagementLoginToDomain,
		Schema: map[string]*schema.Schema{
			"domain": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Domain identified by the name or UID.",
			},
			"continue_last_session": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "When 'continue-last-session' is set to 'True', the new session would continue where the last session was stopped. This option is available when the administrator has only one session that can be continued. If there is more than one session, see 'switch-session' API.",
			},
			"read_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Login with Read Only permissions. This parameter is not considered in case continue-last-session is true.",
			},
			"sid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Session unique identifier. Enter this session unique identifier in the 'X-chkp-sid' header of each request.",
			},
			"api_server_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "API Server version.",
			},
			"disk_space_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Information about the available disk space on the management server.",
			},
			"last_login_was_at": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Timestamp when administrator last accessed the management server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format.",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
						},
					},
				},
			},
			"login_message": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Login message.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message header.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Message content.",
						},
					},
				},
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Session expiration timeout in seconds.",
			},
			"standby": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if this management server is in the standby mode.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Session object unique identifier. This identifier may be used in the discard API to discard changes that were made in this session, when administrator is working from another session, or in the 'switch-session' API.",
			},
			"url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL that was used to reach the API server.",
			},
		},
	}
}

func createManagementLoginToDomain(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("domain"); ok {
		payload["domain"] = v.(string)
	}

	if v, ok := d.GetOk("continue_last_session"); ok {
		payload["continue-last-session"] = v.(bool)
	}

	if v, ok := d.GetOk("read_only"); ok {
		payload["read-only"] = v.(bool)
	}

	LoginToDomainRes, err := client.ApiCall("login-to-domain", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !LoginToDomainRes.Success {
		return fmt.Errorf(LoginToDomainRes.ErrorMsg)
	}

	loginToDomain := LoginToDomainRes.GetData()

	log.Println("Read Login To Domain - Show JSON = ", loginToDomain)

	if v := loginToDomain["sid"]; v != nil {
		_ = d.Set("sid", v)
	}

	if v := loginToDomain["api-server-version"]; v != nil {
		_ = d.Set("api_server_version", v)
	}

	if v := loginToDomain["disk-space-message"]; v != nil {
		_ = d.Set("disk_space_message", v)
	}

	if loginToDomain["last-login-was-at"] != nil {
		lastLoginWasAtMap := loginToDomain["last-login-was-at"].(map[string]interface{})

		lastLoginWasAtMapToReturn := make(map[string]interface{})

		if v, _ := lastLoginWasAtMap["iso-8601"]; v != nil {
			lastLoginWasAtMapToReturn["iso_8601"] = v
		}
		if v, _ := lastLoginWasAtMap["posix"]; v != nil {
			lastLoginWasAtMapToReturn["posix"] = v
		}

		_ = d.Set("last_login_was_at", lastLoginWasAtMapToReturn)
	} else {
		_ = d.Set("last_login_was_at", nil)
	}

	if loginToDomain["login-message"] != nil {
		loginMessageMap := loginToDomain["login-message"].(map[string]interface{})

		loginMessageMapToReturn := make(map[string]interface{})

		if v, _ := loginMessageMap["header"]; v != nil {
			loginMessageMapToReturn["header"] = v
		}
		if v, _ := loginMessageMap["message"]; v != nil {
			loginMessageMapToReturn["message"] = v
		}

		_ = d.Set("login_message", loginMessageMapToReturn)
	} else {
		_ = d.Set("login_message", nil)
	}

	if v := loginToDomain["read-only"]; v != nil {
		_ = d.Set("read_only", v)
	}

	if v := loginToDomain["session-timeout"]; v != nil {
		_ = d.Set("session_timeout", v)
	}

	if v := loginToDomain["standby"]; v != nil {
		_ = d.Set("standby", v)
	}

	if v := loginToDomain["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := loginToDomain["url"]; v != nil {
		_ = d.Set("url", v)
	}

	return readManagementLoginToDomain(d, m)
}

func readManagementLoginToDomain(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementLoginToDomain(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
