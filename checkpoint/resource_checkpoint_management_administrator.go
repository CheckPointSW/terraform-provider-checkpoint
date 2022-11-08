package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementAdministrator() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAdministrator,
		Read:   readManagementAdministrator,
		Update: updateManagementAdministrator,
		Delete: deleteManagementAdministrator,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Authentication method.",
				Default:     "check point password",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Administrator email.",
			},
			"expiration_date": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.",
			},
			"multi_domain_profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Administrator multi-domain profile.",
			},
			"must_change_password": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "True if administrator must change password on the next login.",
				Default:     true,
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Administrator password.",
			},
			"password_hash": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Administrator password hash.",
			},
			"permissions_profile": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Administrator permissions profile. Permissions profile should not be provided when multi-domain-profile is set to \"Multi-Domain Super User\" or \"Domain Super User\".",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:     schema.TypeString,
							Required: true,
						},
						"profile": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"phone_number": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Administrator phone number.",
			},
			"radius_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "RADIUS server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"RADIUS\".",
			},
			"tacacs_server": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "TACACS server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"TACACS\".",
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring warnings.\nApply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
			"sic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Secure Internal Connection Trust.",
			},
		},
	}
}

func createManagementAdministrator(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	administrator := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		administrator["name"] = v.(string)
	}

	if v, ok := d.GetOk("authentication_method"); ok {
		administrator["authentication-method"] = v.(string)
	}

	if v, ok := d.GetOk("email"); ok {
		administrator["email"] = v.(string)
	}

	if v, ok := d.GetOk("expiration_date"); ok {
		administrator["expiration-date"] = v.(string)
	}

	if v, ok := d.GetOk("multi_domain_profile"); ok {
		administrator["multi-domain-profile"] = v.(string)
	}

	if v, ok := d.GetOkExists("must_change_password"); ok {
		administrator["must-change-password"] = v.(bool)
	}

	if v, ok := d.GetOk("password"); ok {
		administrator["password"] = v.(string)
	}

	if v, ok := d.GetOk("password_hash"); ok {
		administrator["password-hash"] = v.(string)
	}

	if v, ok := d.GetOk("permissions_profile"); ok {
		permissionsProfileList := v.([]interface{})

		if v, _ := d.GetOk("permissions_profile.0.domain"); v.(string) == "SMC User" {
			if len(permissionsProfileList) == 1 {

				if v, ok := d.GetOk("permissions_profile.0.profile"); ok {
					administrator["permissions-profile"] = v.(string)
				}

			}

		} else {
			administrator["permissions-profile"] = permissionsProfileList
		}
	}

	if v, ok := d.GetOk("phone_number"); ok {
		administrator["phone-number"] = v.(string)
	}

	if v, ok := d.GetOk("radius_server"); ok {
		administrator["radius-sever"] = v.(string)
	}

	if v, ok := d.GetOk("tacacs_server"); ok {
		administrator["tacacs-sever"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		administrator["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		administrator["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		administrator["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		administrator["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		administrator["ignore-errors"] = v.(bool)
	}

	log.Println("Create Administrator - Map = ", administrator)

	addAdministratorRes, err := client.ApiCall("add-administrator", administrator, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addAdministratorRes.Success {
		if addAdministratorRes.ErrorMsg != "" {
			return fmt.Errorf(addAdministratorRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addAdministratorRes.GetData()["uid"].(string))
	return readManagementAdministrator(d, m)
}

func readManagementAdministrator(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showAdministratorRes, err := client.ApiCall("show-administrator", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAdministratorRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showAdministratorRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showAdministratorRes.ErrorMsg)
	}

	administrator := showAdministratorRes.GetData()
	log.Println("Read Administrator - Show JSON = ", administrator)

	if v := administrator["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := administrator["authentication-method"]; v != nil {
		_ = d.Set("authentication_method", v)
	}

	if v := administrator["email"]; v != nil {
		_ = d.Set("email", v)
	}

	if v := administrator["expiration-date"]; v != nil {
		_ = d.Set("expiration_date", v)
	}

	if administrator["multi-domain-profile"] != nil {
		if multiDomainProfileMap, ok := administrator["multi-domain-profile"].(map[string]interface{}); ok {
			if v, _ := multiDomainProfileMap["name"]; v != nil {
				_ = d.Set("multi_domain_profile", v)
			}
		}
	}

	if v := administrator["must-change-password"]; v != nil {
		_ = d.Set("must_change_password", v)
	}

	if v := administrator["password"]; v != nil {
		_ = d.Set("password", v)
	}

	if v := administrator["password-hash"]; v != nil {
		_ = d.Set("password_hash", v)
	}

	if v := administrator["must-change-password"]; v != nil {
		_ = d.Set("must_change_password", v)
	}

	if administrator["permissions-profile"] != nil {
		var permissionsProfileListToReturn []map[string]interface{}

		if permissionsProfileList, ok := administrator["permissions-profile"].([]interface{}); ok {

			for i := range permissionsProfileList {
				permissionsProfileMap := permissionsProfileList[i].(map[string]interface{})

				permissionsProfileMapToAdd := make(map[string]interface{})

				if profile, _ := permissionsProfileMap["profile"]; profile != nil {
					if v, _ := profile.(map[string]interface{})["name"]; v != nil {
						permissionsProfileMapToAdd["profile"] = v.(string)
					}
				}
				if domain, _ := permissionsProfileMap["domain"]; domain != nil {
					if v, _ := domain.(map[string]interface{})["name"]; v != nil {
						permissionsProfileMapToAdd["domain"] = v.(string)
					}
				}
				permissionsProfileListToReturn = append(permissionsProfileListToReturn, permissionsProfileMapToAdd)
			}

		} else if v, ok := administrator["permissions-profile"].(map[string]interface{}); ok {
			permissionsProfileListToReturn = []map[string]interface{}{
				{
					"domain":  "SMC User",
					"profile": v["name"].(string),
				},
			}
		}
		_ = d.Set("permissions_profile", permissionsProfileListToReturn)

	}

	if v := administrator["phone-number"]; v != nil {
		_ = d.Set("phone_number", v)
	}

	if v := administrator["radius-server"]; v != nil {
		_ = d.Set("radius_server", v)
	}

	if v := administrator["tacacs-server"]; v != nil {
		_ = d.Set("tacacs_server", v)
	}

	if administrator["tags"] != nil {
		tagsJson := administrator["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	if v := administrator["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := administrator["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := administrator["sic-name"]; v != nil {
		_ = d.Set("sic_name", v)
	}

	return nil
}

func updateManagementAdministrator(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	administrator := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		administrator["name"] = oldName
		administrator["new-name"] = newName
	} else {
		administrator["name"] = d.Get("name")
	}

	if d.HasChange("authentication_method") {
		administrator["authentication-method"] = d.Get("authentication_method")
	}

	if d.HasChange("email") {
		administrator["email"] = d.Get("email")
	}

	if d.HasChange("expiration_date") {
		administrator["expiration-date"] = d.Get("expiration_date")
	}

	if d.HasChange("multi_domain_profile") {
		administrator["multi-domain-profile"] = d.Get("multi_domain_profile")
	}

	if d.HasChange("email") {
		administrator["email"] = d.Get("email")
	}

	if v, ok := d.GetOkExists("must_change_password"); ok {
		administrator["must-change-password"] = v.(bool)
	}

	if d.HasChange("password") {
		administrator["password"] = d.Get("password")
	}

	if d.HasChange("password_hash") {
		administrator["password-hash"] = d.Get("password_hash")
	}

	if d.HasChange("permissions_profile") {

		if v, ok := d.GetOk("permissions_profile"); ok {
			permissionsProfileList := v.([]interface{})

			if len(permissionsProfileList) == 1 {
				if v, _ := d.GetOk("permissions_profile.0.domain"); v.(string) == "SMC User" {

					if v, ok := d.GetOk("permissions_profile.0.profile"); ok {
						administrator["permissions-profile"] = v.(string)
					}
				}

			} else {
				administrator["permissions-profile"] = permissionsProfileList
			}
		}
	}

	if d.HasChange("phone_number") {
		administrator["phone-number"] = d.Get("phone_number")
	}

	if d.HasChange("radius_server") {
		administrator["radius-server"] = d.Get("radius_server")
	}

	if d.HasChange("tacacs_server") {
		administrator["tacacs-server"] = d.Get("tacacs_server")
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			administrator["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			administrator["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if d.HasChange("color") {
		administrator["color"] = d.Get("color")
	}

	if d.HasChange("comments") {
		administrator["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		administrator["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		administrator["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Administrator - Map = ", administrator)
	updateAdministratorRes, err := client.ApiCall("set-administrator", administrator, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateAdministratorRes.Success {
		if updateAdministratorRes.ErrorMsg != "" {
			return fmt.Errorf(updateAdministratorRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementAdministrator(d, m)
}

func deleteManagementAdministrator(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	administratorPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	deleteAdministratorRes, err := client.ApiCall("delete-administrator", administratorPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteAdministratorRes.Success {
		if deleteAdministratorRes.ErrorMsg != "" {
			return fmt.Errorf(deleteAdministratorRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
