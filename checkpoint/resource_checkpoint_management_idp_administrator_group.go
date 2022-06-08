package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"strconv"
)

func resourceManagementIdpAdministratorGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementIdpAdministratorGroup,
		Read:   readManagementIdpAdministratorGroup,
		Update: updateManagementIdpAdministratorGroup,
		Delete: deleteManagementIdpAdministratorGroup,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"group_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Group ID or Name should be set base on the source attribute of 'groups' in the Saml Assertion.",
			},
			"multi_domain_profile": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Administrator multi-domain profile.",
			},
			"permissions_profile": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Administrator permissions profile. Permissions profile should not be provided when multi-domain-profile is set to \"Multi-Domain Super User\" or \"Domain Super User\".",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
						"profile": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Permission profile.",
						},
					},
				},
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
				Description: "Apply changes ignoring warnings.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
		},
	}
}

func createManagementIdpAdministratorGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	idpAdministratorGroup := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		idpAdministratorGroup["name"] = v.(string)
	}

	if v, ok := d.GetOk("group_id"); ok {
		idpAdministratorGroup["group-id"] = v.(string)
	}

	if v, ok := d.GetOk("multi_domain_profile"); ok {
		idpAdministratorGroup["multi-domain-profile"] = v.(string)
	}

	if v, ok := d.GetOk("permissions_profile"); ok {

		permissionsProfileList := v.([]interface{})

		if len(permissionsProfileList) > 0 {

			var permissionsProfilePayload []map[string]interface{}

			for i := range permissionsProfileList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("permissions_profile." + strconv.Itoa(i) + ".domain"); ok {
					Payload["domain"] = v.(string)
				}
				if v, ok := d.GetOk("permissions_profile." + strconv.Itoa(i) + ".profile"); ok {
					Payload["profile"] = v.(string)
				}
				permissionsProfilePayload = append(permissionsProfilePayload, Payload)
			}
			idpAdministratorGroup["permissions-profile"] = permissionsProfilePayload
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		idpAdministratorGroup["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		idpAdministratorGroup["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		idpAdministratorGroup["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		idpAdministratorGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		idpAdministratorGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Create IdpAdministratorGroup - Map = ", idpAdministratorGroup)

	addIdpAdministratorGroupRes, err := client.ApiCall("add-idp-administrator-group", idpAdministratorGroup, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addIdpAdministratorGroupRes.Success {
		if addIdpAdministratorGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addIdpAdministratorGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addIdpAdministratorGroupRes.GetData()["uid"].(string))

	return readManagementIdpAdministratorGroup(d, m)
}

func readManagementIdpAdministratorGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showIdpAdministratorGroupRes, err := client.ApiCall("show-idp-administrator-group", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

func updateManagementIdpAdministratorGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	idpAdministratorGroup := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		idpAdministratorGroup["name"] = oldName
		idpAdministratorGroup["new-name"] = newName
	} else {
		idpAdministratorGroup["name"] = d.Get("name")
	}

	if ok := d.HasChange("group_id"); ok {
		idpAdministratorGroup["group-id"] = d.Get("group_id")
	}

	if ok := d.HasChange("multi_domain_profile"); ok {
		idpAdministratorGroup["multi-domain-profile"] = d.Get("multi_domain_profile")
	}

	if d.HasChange("permissions_profile") {

		if v, ok := d.GetOk("permissions_profile"); ok {

			permissionsProfileList := v.([]interface{})

			var permissionsProfilePayload []map[string]interface{}

			for i := range permissionsProfileList {

				Payload := make(map[string]interface{})

				if d.HasChange("permissions_profile." + strconv.Itoa(i) + ".domain") {
					Payload["domain"] = d.Get("permissions_profile." + strconv.Itoa(i) + ".domain")
				}
				if d.HasChange("permissions_profile." + strconv.Itoa(i) + ".profile") {
					Payload["profile"] = d.Get("permissions_profile." + strconv.Itoa(i) + ".profile")
				}
				permissionsProfilePayload = append(permissionsProfilePayload, Payload)
			}
			idpAdministratorGroup["permissions-profile"] = permissionsProfilePayload
		} else {
			oldpermissionsProfile, _ := d.GetChange("permissions_profile")
			var permissionsProfileToDelete []interface{}
			for _, i := range oldpermissionsProfile.([]interface{}) {
				permissionsProfileToDelete = append(permissionsProfileToDelete, i.(map[string]interface{})["name"].(string))
			}
			idpAdministratorGroup["permissions-profile"] = map[string]interface{}{"remove": permissionsProfileToDelete}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			idpAdministratorGroup["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			idpAdministratorGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		idpAdministratorGroup["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		idpAdministratorGroup["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		idpAdministratorGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		idpAdministratorGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Update IdpAdministratorGroup - Map = ", idpAdministratorGroup)

	updateIdpAdministratorGroupRes, err := client.ApiCall("set-idp-administrator-group", idpAdministratorGroup, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateIdpAdministratorGroupRes.Success {
		if updateIdpAdministratorGroupRes.ErrorMsg != "" {
			return fmt.Errorf(updateIdpAdministratorGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementIdpAdministratorGroup(d, m)
}

func deleteManagementIdpAdministratorGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	idpAdministratorGroupPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete IdpAdministratorGroup")

	deleteIdpAdministratorGroupRes, err := client.ApiCall("delete-idp-administrator-group", idpAdministratorGroupPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteIdpAdministratorGroupRes.Success {
		if deleteIdpAdministratorGroupRes.ErrorMsg != "" {
			return fmt.Errorf(deleteIdpAdministratorGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
