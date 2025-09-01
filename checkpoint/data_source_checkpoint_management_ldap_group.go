package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementLdapGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementLdapGroupRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"account_unit": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "LDAP account unit of the group. Identified by name or UID.",
			},
			"scope": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Group's scope. There are three possible ways of defining a group, based on the users defined on the Account Unit.",
			},
			"account_unit_branch": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Branch of the selected LDAP Account Unit.",
			},
			"sub_tree_prefix": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Sub tree prefix of the selected branch. <font color=\"red\">Relevant only when</font> 'scope' is set to 'only_sub_prefix'. Must be in DN syntax.",
			},
			"group_prefix": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Group name in the selected branch. <font color=\"red\">Required only when</font> 'scope' is set to 'only_group_in_branch'. Must be in DN syntax.",
			},
			"apply_filter_for_dynamic_group": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicate whether to apply LDAP filter for dynamic group. <font color=\"red\">Relevant only when</font> 'scope' is not set to 'only_group_in_branch'.",
			},
			"ldap_filter": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "LDAP filter for the dynamic group. <font color=\"red\">Relevant only when</font> 'apply-filter-for-dynamic-group' is set to 'true'.",
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
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementLdapGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showLdapGroupRes, err := client.ApiCallSimple("show-ldap-group", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLdapGroupRes.Success {
		return fmt.Errorf(showLdapGroupRes.ErrorMsg)
	}

	ldapGroup := showLdapGroupRes.GetData()

	log.Println("Read LdapGroup - Show JSON = ", ldapGroup)

	if v := ldapGroup["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := ldapGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := ldapGroup["account-unit"]; v != nil {
		_ = d.Set("account_unit", v.(map[string]interface{})["name"].(string))
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

	return nil
}
