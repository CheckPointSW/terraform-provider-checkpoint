package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceGaiaUsers() *schema.Resource {
	return &schema.Resource{
		Create: addGaiaUsers,
		Read:   showGaiaUsers,
		Update: setGaiaUsers,
		Delete: deleteGaiaUsers,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Username",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "User ID",
			},
			"homedir": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Home directory path",
			},
			"primary_system_group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Primary system group ID",
			},
			"secondary_system_groups": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Secondary system groups",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User password",
				Sensitive:   true,
			},
			"password-hash": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User password hash",
				Sensitive:   true,
			},
			"real-name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User's real name"
			},
			"shell": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "User's shell",
			},
			"allow_access_using": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "Allowed access methods",
			},
			"must_change_password": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Force user to change password on next login",
			},
			"roles": {
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User roles",
			},
			"unlock": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Unlock user account",
			},
			"requires_two_factor_authentication": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Require two-factor authentication for user",
			},
		},
	}
}

func addGaiaUsers(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok { payload["name"] = v.(string) } else { return fmt.Errorf("name is required") }
	if v, ok := d.GetOk("uid"); ok { payload["uid"] = v.(string) } else { return fmt.Errorf("uid is required") }
	if v, ok := d.GetOk("homedir"); ok { payload["homedir"] = v.(string) }
	if v, ok := d.GetOk("primary_system_group_id"); ok { payload["primary_system_group_id"] = v.(string) }
	if v, ok := d.GetOk("secondary_system_groups"); ok { payload["secondary_system_groups"] = v.([]string) }
	if v, ok := d.GetOk("password"); ok { payload["password"] = v.(string) }
	if v, ok := d.GetOk("password-hash"); ok { payload["password-hash"] = v.(string) }
	if v, ok := d.GetOk("real-name"); ok { payload["real-name"] = v.(string) }
	if v, ok := d.GetOk("shell"); ok { payload["shell"] = v.(string) }
	if v, ok := d.GetOk("allow_access_using"); ok { payload["allow_access_using"] = v.([]string) }
	if v, ok := d.GetOk("must_change_password"); ok { payload["must_change_password"] = v.(bool) }
	if v, ok := d.GetOk("roles"); ok { payload["roles"] = v.([]string) }
	if v, ok := d.GetOk("requires_two_factor_authentication"); ok { payload["requires_two_factor_authentication"] = v.(bool) }

	addGaiaUserRes, _ := client.ApiCall("add-user", payload, client.GetSessionID(), true, client.IsProxyUsed(), "POST")
	if !addGaiaUserRes.Success {
		return fmt.Errorf(addGaiaUserRes.ErrorMsg)
	}

	// Set Schema UID = Object key
	d.SetId(addGaiaUserRes.GetData()["name"].(string))
	return readGaiaUsers(d, m)
}

func showGaiaUsers(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok { payload["name"] = v.(string) } else { return fmt.Errorf("name is required") }

	showGaiaUserRes, err := client.ApiCall("show-user", payload, client.GetSessionID(), true, client.IsProxyUsed(), "POST")
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGaiaUserRes.Success {
		// Handle deletion of an object from other clients - Object not found
		if objectNotFound(showGaiaUserRes.GetData()["code"].(string)) {
			d.SetId("") // Destroy resource
			return nil
		}
		return fmt.Errorf(showGaiaUserRes.ErrorMsg)
	}
	jsonResponse, err := json.Marshal(showGaiaUserRes.GetData())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if jsonResponse != nil {
		_ = d.Set("response", string(jsonResponse))
	}

	_ = d.Set("name", jsonResponse["name"].(string))

	return nil
}

func setGaiaUsers(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok { payload["name"] = v.(string) } else { return fmt.Errorf("name is required") }
	if v, ok := d.GetOk("uid"); ok { payload["uid"] = v.(string) }
	if v, ok := d.GetOk("homedir"); ok { payload["homedir"] = v.(string) }
	if v, ok := d.GetOk("primary_system_group_id"); ok { payload["primary_system_group_id"] = v.(string) }
	if v, ok := d.GetOk("secondary_system_groups"); ok { payload["secondary_system_groups"] = v.([]string) }
	if v, ok := d.GetOk("password"); ok { payload["password"] = v.(string) }
	if v, ok := d.GetOk("password-hash"); ok { payload["password-hash"] = v.(string) }
	if v, ok := d.GetOk("real-name"); ok { payload["real-name"] = v.(string) }
	if v, ok := d.GetOk("shell"); ok { payload["shell"] = v.(string) }
	if v, ok := d.GetOk("allow_access_using"); ok { payload["allow_access_using"] = v.([]string) }
	if v, ok := d.GetOk("must_change_password"); ok { payload["must_change_password"] = v.(bool) }
	if v, ok := d.GetOk("roles"); ok { payload["roles"] = v.([]string) }
	if v, ok := d.GetOk("unlock"); ok { payload["unlock"] = v.(bool) }
	if v, ok := d.GetOk("requires_two_factor_authentication"); ok { payload["requires_two_factor_authentication"] = v.(bool) }

	setGaiaUserRes, _ := client.ApiCall("set-user", payload, client.GetSessionID(), true, client.IsProxyUsed(), "POST")
	if !setGaiaUserRes.Success {
		return fmt.Errorf(setGaiaUserRes.ErrorMsg)
	}
	return readGaiaUsers(d, m)
}

func deleteGaiaUsers(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok { payload["name"] = v.(string) } else { return fmt.Errorf("name is required") }

	deleteGaiaUserRes, _ := client.ApiCall("delete-user", payload, client.GetSessionID(), true, client.IsProxyUsed(), "POST")
	if !deleteGaiaUserRes.Success {
		return fmt.Errorf(deleteGaiaUserRes.ErrorMsg)
	}
	d.SetId("") // Destroy resource
	return nil
}
