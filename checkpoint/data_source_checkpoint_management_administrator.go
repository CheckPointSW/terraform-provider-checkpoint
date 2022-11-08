package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementAdministrator() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAdministratorRead,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Authentication method.",
			},
			"email": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Administrator email.",
			},
			"expiration_date": {
				Type:     schema.TypeMap,
				Computed: true,
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
			"multi_domain_profile": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Administrator multi-domain profile. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
			},
			"must_change_password": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if administrator must change password on the next login.",
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
				Computed:    true,
				Description: "Administrator phone number.",
			},
			"radius_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "RADIUS server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"RADIUS\". Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
			},
			"sic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Secure Internal Connection Trust.",
			},
			"tacacs_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "TACACS server object identified by the name or UID . Must be set when \"authentication-method\" was selected to be \"TACACS\". Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
		},
	}
}

func dataSourceManagementAdministratorRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showAdministratorRes, err := client.ApiCall("show-administrator", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAdministratorRes.Success {
		return fmt.Errorf(showAdministratorRes.ErrorMsg)
	}

	administrator := showAdministratorRes.GetData()
	log.Println("Read Administrator - Show JSON = ", administrator)

	if v := administrator["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := administrator["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
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
