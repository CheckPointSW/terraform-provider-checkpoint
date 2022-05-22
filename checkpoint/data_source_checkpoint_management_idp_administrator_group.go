package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementIdpAdministratorGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementIdpAdministratorGroupRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Group ID or Name should be set base on the source attribute of 'groups' in the Saml Assertion.",
			},
			"multi_domain_profile": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Administrator multi-domain profile.",
			},
			"permissions_profile": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Administrator permissions profile. Permissions profile should not be provided when multi-domain-profile is set to \"Multi-Domain Super User\" or \"Domain Super User\".",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"profile": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Permission profile.",
						},
					},
				},
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

func dataSourceManagementIdpAdministratorGroupRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showIdpAdministratorGroupRes, err := client.ApiCall("show-idp-administrator-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIdpAdministratorGroupRes.Success {
		if objectNotFound(showIdpAdministratorGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showIdpAdministratorGroupRes.ErrorMsg)
	}

	idpAdministratorGroup := showIdpAdministratorGroupRes.GetData()

	log.Println("Read IdpAdministratorGroup - Show JSON = ", idpAdministratorGroup)

	if v := idpAdministratorGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := idpAdministratorGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := idpAdministratorGroup["group-id"]; v != nil {
		_ = d.Set("group_id", v)
	}

	if v := idpAdministratorGroup["multi-domain-profile"]; v != nil {
		_ = d.Set("multi_domain_profile", v)
	}

	if idpAdministratorGroup["permissions-profile"] != nil {

		permissionsProfileList, ok := idpAdministratorGroup["permissions-profile"].([]interface{})

		if ok {

			if len(permissionsProfileList) > 0 {

				var permissionsProfileListToReturn []map[string]interface{}

				for i := range permissionsProfileList {

					permissionsProfileMap := permissionsProfileList[i].(map[string]interface{})

					permissionsProfileMapToAdd := make(map[string]interface{})

					if v, _ := permissionsProfileMap["domain"]; v != nil {
						permissionsProfileMapToAdd["domain"] = v
					}
					if v, _ := permissionsProfileMap["profile"]; v != nil {
						permissionsProfileMapToAdd["profile"] = v
					}
					permissionsProfileListToReturn = append(permissionsProfileListToReturn, permissionsProfileMapToAdd)
				}
			}
		}
	}

	if idpAdministratorGroup["tags"] != nil {
		tagsJson, ok := idpAdministratorGroup["tags"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("tags", tagsIds)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if v := idpAdministratorGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := idpAdministratorGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := idpAdministratorGroup["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := idpAdministratorGroup["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil
}
