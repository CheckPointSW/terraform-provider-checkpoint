package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementMdPermissionsProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementMdPermissionsProfileRead,
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
			"mds_provisioning": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Create and manage Multi-Domain Servers and Multi-Domain Log Servers.<br>Only a \"Super User\" permission-level profile can select this option.",
			},
			"manage_admins": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Create and manage Multi-Domain Security Management administrators with the same or lower permission level. For example, a Domain manager cannot create Superusers or global managers.<br>Only a 'Manager' permission-level profile can edit this permission.",
			},
			"manage_sessions": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Connect/disconnect Domain sessions, publish changes, and delete other administrator sessions.<br>Only a 'Manager' permission-level profile can edit this permission.",
			},
			"management_api_login": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Permission to log in to the Security Management Server and run API commands using these tools: mgmt_cli (Linux and Windows binaries), Gaia CLI (clish) and Web Services (REST). Useful if you want to prevent administrators from running automatic scripts on the Management.<br>Note: This permission is not required to run commands from within the API terminal in SmartConsole.",
			},
			"cme_operations": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Permission to read / edit the Cloud Management Extension (CME) configuration.",
			},
			"global_vpn_management": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Lets the administrator select Enable global use for a Security Gateway shown in the MDS Gateways & Servers view.<br>Only a 'Manager' permission-level profile can edit this permission.",
			},
			"manage_global_assignments": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Controls the ability to create, edit and delete global assignment and not the ability to reassign, which is set according to the specific Domain's permission profile.",
			},
			"enable_default_profile_for_global_domains": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable the option to specify a default profile for all global domains.",
			},
			"default_profile_global_domains": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name or UID of the required default profile for all global domains.",
			},
			"view_global_objects_in_domain": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Lets an administrator with no global objects permissions view the global objects in the domain. This option is required for valid domain management.",
			},
			"enable_default_profile_for_local_domains": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable the option to specify a default profile for all local domains.",
			},
			"default_profile_local_domains": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name or UID of the required default profile for all local domains.",
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
			"domains_to_process": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"permission_level": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The level of the Multi Domain Permissions Profile.<br>The level cannot be changed after creation.",
			},
		},
	}
}

func dataSourceManagementMdPermissionsProfileRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showMdPermissionsProfileRes, err := client.ApiCall("show-md-permissions-profile", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showMdPermissionsProfileRes.Success {
		if objectNotFound(showMdPermissionsProfileRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showMdPermissionsProfileRes.ErrorMsg)
	}

	mdPermissionsProfile := showMdPermissionsProfileRes.GetData()

	log.Println("Read MdPermissionsProfile - Show JSON = ", mdPermissionsProfile)

	if v := mdPermissionsProfile["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := mdPermissionsProfile["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := mdPermissionsProfile["mds-provisioning"]; v != nil {
		_ = d.Set("mds_provisioning", v)
	}

	if v := mdPermissionsProfile["manage-admins"]; v != nil {
		_ = d.Set("manage_admins", v)
	}

	if v := mdPermissionsProfile["manage-sessions"]; v != nil {
		_ = d.Set("manage_sessions", v)
	}

	if v := mdPermissionsProfile["management-api-login"]; v != nil {
		_ = d.Set("management_api_login", v)
	}

	if v := mdPermissionsProfile["cme-operations"]; v != nil {
		_ = d.Set("cme_operations", v)
	}

	if v := mdPermissionsProfile["global-vpn-management"]; v != nil {
		_ = d.Set("global_vpn_management", v)
	}

	if v := mdPermissionsProfile["manage-global-assignments"]; v != nil {
		_ = d.Set("manage_global_assignments", v)
	}

	if v := mdPermissionsProfile["enable-default-profile-for-global-domains"]; v != nil {
		_ = d.Set("enable_default_profile_for_global_domains", v)
	}

	if v := mdPermissionsProfile["default-profile-global-domains"]; v != nil {
		_ = d.Set("default_profile_global_domains", v.(map[string]interface{})["name"])
	}

	if v := mdPermissionsProfile["view-global-objects-in-domain"]; v != nil {
		_ = d.Set("view_global_objects_in_domain", v)
	}

	if v := mdPermissionsProfile["enable-default-profile-for-local-domains"]; v != nil {
		_ = d.Set("enable_default_profile_for_local_domains", v)
	}

	if v := mdPermissionsProfile["default-profile-local-domains"]; v != nil {
		_ = d.Set("default_profile_local_domains", v.(map[string]interface{})["name"])
	}

	if mdPermissionsProfile["tags"] != nil {
		tagsJson, ok := mdPermissionsProfile["tags"].([]interface{})
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

	if v := mdPermissionsProfile["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := mdPermissionsProfile["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if mdPermissionsProfile["domains_to_process"] != nil {
		domainsToProcessJson, ok := mdPermissionsProfile["domains_to_process"].([]interface{})
		if ok {
			domainsToProcessIds := make([]string, 0)
			if len(domainsToProcessJson) > 0 {
				for _, domains_to_process := range domainsToProcessJson {
					domains_to_process := domains_to_process.(map[string]interface{})
					domainsToProcessIds = append(domainsToProcessIds, domains_to_process["name"].(string))
				}
			}
			_ = d.Set("domains_to_process", domainsToProcessIds)
		}
	} else {
		_ = d.Set("domains_to_process", nil)
	}

	if v := mdPermissionsProfile["permission-level"]; v != nil {
		_ = d.Set("permission_level", v)
	}

	return nil
}
