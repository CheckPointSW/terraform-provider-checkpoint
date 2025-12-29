package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementLdapGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLdapGroup,
		Read:   readManagementLdapGroup,
		Update: updateManagementLdapGroup,
		Delete: deleteManagementLdapGroup,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"account_unit": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "LDAP account unit of the group.  Identified by name or UID.",
			},
			"scope": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Group's scope. There are three possible ways of defining a group, based on the users defined on the Account Unit.",
				Default:     "all_account_unit_users",
			},
			"account_unit_branch": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Branch of the selected LDAP Account Unit.",
			},
			"sub_tree_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Sub tree prefix of the selected branch. <font color=\"red\">Relevant only when</font> 'scope' is set to 'only_sub_prefix'. Must be in DN syntax.",
			},
			"group_prefix": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Group name in the selected branch. <font color=\"red\">Required only when</font> 'scope' is set to 'only_group_in_branch'. Must be in DN syntax.",
			},
			"apply_filter_for_dynamic_group": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicate whether to apply LDAP filter for dynamic group. <font color=\"red\">Relevant only when</font> 'scope' is not set to 'only_group_in_branch'.",
				Default:     false,
			},
			"ldap_filter": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "LDAP filter for the dynamic group. <font color=\"red\">Relevant only when</font> 'apply-filter-for-dynamic-group' is set to 'true'.",
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementLdapGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	ldapGroup := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		ldapGroup["name"] = v.(string)
	}

	if v, ok := d.GetOk("account_unit"); ok {
		ldapGroup["account-unit"] = v.(string)
	}

	if v, ok := d.GetOk("scope"); ok {
		ldapGroup["scope"] = v.(string)
	}

	if v, ok := d.GetOk("account_unit_branch"); ok {
		ldapGroup["account-unit-branch"] = v.(string)
	}

	if v, ok := d.GetOk("sub_tree_prefix"); ok {
		ldapGroup["sub-tree-prefix"] = v.(string)
	}

	if v, ok := d.GetOk("group_prefix"); ok {
		ldapGroup["group-prefix"] = v.(string)
	}

	if v, ok := d.GetOkExists("apply_filter_for_dynamic_group"); ok {
		ldapGroup["apply-filter-for-dynamic-group"] = v.(bool)
	}

	if v, ok := d.GetOk("ldap_filter"); ok {
		ldapGroup["ldap-filter"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		ldapGroup["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		ldapGroup["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		ldapGroup["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		ldapGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		ldapGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Create LdapGroup - Map = ", ldapGroup)

	addLdapGroupRes, err := client.ApiCallSimple("add-ldap-group", ldapGroup)
	if err != nil || !addLdapGroupRes.Success {
		if addLdapGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addLdapGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addLdapGroupRes.GetData()["uid"].(string))

	return readManagementLdapGroup(d, m)
}

func readManagementLdapGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showLdapGroupRes, err := client.ApiCallSimple("show-ldap-group", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLdapGroupRes.Success {
		if objectNotFound(showLdapGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLdapGroupRes.ErrorMsg)
	}

	ldapGroup := showLdapGroupRes.GetData()

	log.Println("Read LdapGroup - Show JSON = ", ldapGroup)

	if v := ldapGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := ldapGroup["account-unit"]; v != nil {
		_ = d.Set("account_unit", v)
	}

	if v := ldapGroup["scope"]; v != nil {
		_ = d.Set("scope", v)
	}

	if v := ldapGroup["account-unit-branch"]; v != nil {
		_ = d.Set("account_unit_branch", v)
	}

	if v := ldapGroup["sub-tree-prefix"]; v != nil {
		_ = d.Set("sub_tree_prefix", v)
	}

	if v := ldapGroup["group-prefix"]; v != nil {
		_ = d.Set("group_prefix", v)
	}

	if v := ldapGroup["apply-filter-for-dynamic-group"]; v != nil {
		_ = d.Set("apply_filter_for_dynamic_group", v)
	}

	if v := ldapGroup["ldap-filter"]; v != nil {
		_ = d.Set("ldap_filter", v)
	}

	if v := ldapGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := ldapGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if ldapGroup["tags"] != nil {
		tagsJson, ok := ldapGroup["tags"].([]interface{})
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

	if v := ldapGroup["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := ldapGroup["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementLdapGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	ldapGroup := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		ldapGroup["name"] = oldName
		ldapGroup["new-name"] = newName
	} else {
		ldapGroup["name"] = d.Get("name")
	}

	if ok := d.HasChange("account_unit"); ok {
		ldapGroup["account-unit"] = d.Get("account_unit")
	}

	if ok := d.HasChange("scope"); ok {
		ldapGroup["scope"] = d.Get("scope")
	}

	if ok := d.HasChange("account_unit_branch"); ok {
		ldapGroup["account-unit-branch"] = d.Get("account_unit_branch")
	}

	if ok := d.HasChange("sub_tree_prefix"); ok {
		ldapGroup["sub-tree-prefix"] = d.Get("sub_tree_prefix")
	}

	if ok := d.HasChange("group_prefix"); ok {
		ldapGroup["group-prefix"] = d.Get("group_prefix")
	}

	if v, ok := d.GetOkExists("apply_filter_for_dynamic_group"); ok {
		ldapGroup["apply-filter-for-dynamic-group"] = v.(bool)
	}

	if ok := d.HasChange("ldap_filter"); ok {
		ldapGroup["ldap-filter"] = d.Get("ldap_filter")
	}

	if ok := d.HasChange("color"); ok {
		ldapGroup["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		ldapGroup["comments"] = d.Get("comments")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			ldapGroup["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		ldapGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		ldapGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Update LdapGroup - Map = ", ldapGroup)

	updateLdapGroupRes, err := client.ApiCallSimple("set-ldap-group", ldapGroup)
	if err != nil || !updateLdapGroupRes.Success {
		if updateLdapGroupRes.ErrorMsg != "" {
			return fmt.Errorf(updateLdapGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementLdapGroup(d, m)
}

func deleteManagementLdapGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	ldapGroupPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete LdapGroup")

	deleteLdapGroupRes, err := client.ApiCallSimple("delete-ldap-group", ldapGroupPayload)
	if err != nil || !deleteLdapGroupRes.Success {
		if deleteLdapGroupRes.ErrorMsg != "" {
			return fmt.Errorf(deleteLdapGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
