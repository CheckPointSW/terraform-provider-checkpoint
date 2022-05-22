package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementMdPermissionsProfile() *schema.Resource {
	return &schema.Resource{
		Create: createManagementMdPermissionsProfile,
		Read:   readManagementMdPermissionsProfile,
		Update: updateManagementMdPermissionsProfile,
		Delete: deleteManagementMdPermissionsProfile,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"mds_provisioning": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Create and manage Multi-Domain Servers and Multi-Domain Log Servers.<br>Only a \"Super User\" permission-level profile can select this option.",
			},
			"manage_admins": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Create and manage Multi-Domain Security Management administrators with the same or lower permission level. For example, a Domain manager cannot create Superusers or global managers.<br>Only a 'Manager' permission-level profile can edit this permission.",
				Default:     true,
			},
			"manage_sessions": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Connect/disconnect Domain sessions, publish changes, and delete other administrator sessions.<br>Only a 'Manager' permission-level profile can edit this permission.",
			},
			"management_api_login": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Permission to log in to the Security Management Server and run API commands using these tools: mgmt_cli (Linux and Windows binaries), Gaia CLI (clish) and Web Services (REST). Useful if you want to prevent administrators from running automatic scripts on the Management.<br>Note: This permission is not required to run commands from within the API terminal in SmartConsole.",
				Default:     true,
			},
			"cme_operations": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Permission to read / edit the Cloud Management Extension (CME) configuration.",
				Default:     "disabled",
			},
			"global_vpn_management": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Lets the administrator select Enable global use for a Security Gateway shown in the MDS Gateways & Servers view.<br>Only a 'Manager' permission-level profile can edit this permission.",
			},
			"manage_global_assignments": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Controls the ability to create, edit and delete global assignment and not the ability to reassign, which is set according to the specific Domain's permission profile.",
			},
			"enable_default_profile_for_global_domains": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable the option to specify a default profile for all global domains.",
				Default:     true,
			},
			"default_profile_global_domains": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name or UID of the required default profile for all global domains.",
				Default:     "Read Only All",
			},
			"view_global_objects_in_domain": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Lets an administrator with no global objects permissions view the global objects in the domain. This option is required for valid domain management.",
				Default:     true,
			},
			"enable_default_profile_for_local_domains": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable the option to specify a default profile for all local domains.",
			},
			"default_profile_local_domains": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Name or UID of the required default profile for all local domains.",
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
			"domains_to_process": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
			"permission_level": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The level of the Multi Domain Permissions Profile.<br>The level cannot be changed after creation.",
				Default:     "manager",
			},
		},
	}
}

func createManagementMdPermissionsProfile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	mdPermissionsProfile := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		mdPermissionsProfile["name"] = v.(string)
	}

	if v, ok := d.GetOkExists("mds_provisioning"); ok {
		mdPermissionsProfile["mds-provisioning"] = v.(bool)
	}

	if v, ok := d.GetOkExists("manage_admins"); ok {
		mdPermissionsProfile["manage-admins"] = v.(bool)
	}

	if v, ok := d.GetOkExists("manage_sessions"); ok {
		mdPermissionsProfile["manage-sessions"] = v.(bool)
	}

	if v, ok := d.GetOkExists("management_api_login"); ok {
		mdPermissionsProfile["management-api-login"] = v.(bool)
	}

	if v, ok := d.GetOk("cme_operations"); ok {
		mdPermissionsProfile["cme-operations"] = v.(string)
	}

	if v, ok := d.GetOkExists("global_vpn_management"); ok {
		mdPermissionsProfile["global-vpn-management"] = v.(bool)
	}

	if v, ok := d.GetOkExists("manage_global_assignments"); ok {
		mdPermissionsProfile["manage-global-assignments"] = v.(bool)
	}

	if v, ok := d.GetOkExists("enable_default_profile_for_global_domains"); ok {
		mdPermissionsProfile["enable-default-profile-for-global-domains"] = v.(bool)
	}

	if v, ok := d.GetOk("default_profile_global_domains"); ok {
		mdPermissionsProfile["default-profile-global-domains"] = v.(string)
	}

	if v, ok := d.GetOkExists("view_global_objects_in_domain"); ok {
		mdPermissionsProfile["view-global-objects-in-domain"] = v.(bool)
	}

	if v, ok := d.GetOkExists("enable_default_profile_for_local_domains"); ok {
		mdPermissionsProfile["enable-default-profile-for-local-domains"] = v.(bool)
	}

	if v, ok := d.GetOk("default_profile_local_domains"); ok {
		mdPermissionsProfile["default-profile-local-domains"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		mdPermissionsProfile["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		mdPermissionsProfile["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		mdPermissionsProfile["comments"] = v.(string)
	}

	if v, ok := d.GetOk("domains_to_process"); ok {
		mdPermissionsProfile["domains-to-process"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mdPermissionsProfile["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mdPermissionsProfile["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOk("permission_level"); ok {
		mdPermissionsProfile["permission-level"] = v.(string)
	}

	log.Println("Create MdPermissionsProfile - Map = ", mdPermissionsProfile)

	addMdPermissionsProfileRes, err := client.ApiCall("add-md-permissions-profile", mdPermissionsProfile, client.GetSessionID(), true, false)
	if err != nil || !addMdPermissionsProfileRes.Success {
		if addMdPermissionsProfileRes.ErrorMsg != "" {
			return fmt.Errorf(addMdPermissionsProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addMdPermissionsProfileRes.GetData()["uid"].(string))

	return readManagementMdPermissionsProfile(d, m)
}

func readManagementMdPermissionsProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showMdPermissionsProfileRes, err := client.ApiCall("show-md-permissions-profile", payload, client.GetSessionID(), true, false)
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

	if v := mdPermissionsProfile["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := mdPermissionsProfile["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	if v := mdPermissionsProfile["permission-level"]; v != nil {
		_ = d.Set("permission_level", v)
	}

	return nil

}

func updateManagementMdPermissionsProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	mdPermissionsProfile := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		mdPermissionsProfile["name"] = oldName
		mdPermissionsProfile["new-name"] = newName
	} else {
		mdPermissionsProfile["name"] = d.Get("name")
	}

	if v, ok := d.GetOkExists("mds_provisioning"); ok {
		mdPermissionsProfile["mds-provisioning"] = v.(bool)
	}

	if v, ok := d.GetOkExists("manage_admins"); ok {
		mdPermissionsProfile["manage-admins"] = v.(bool)
	}

	if v, ok := d.GetOkExists("manage_sessions"); ok {
		mdPermissionsProfile["manage-sessions"] = v.(bool)
	}

	if v, ok := d.GetOkExists("management_api_login"); ok {
		mdPermissionsProfile["management-api-login"] = v.(bool)
	}

	if ok := d.HasChange("cme_operations"); ok {
		mdPermissionsProfile["cme-operations"] = d.Get("cme_operations")
	}

	if v, ok := d.GetOkExists("global_vpn_management"); ok {
		mdPermissionsProfile["global-vpn-management"] = v.(bool)
	}

	if v, ok := d.GetOkExists("manage_global_assignments"); ok {
		mdPermissionsProfile["manage-global-assignments"] = v.(bool)
	}

	if v, ok := d.GetOkExists("enable_default_profile_for_global_domains"); ok {
		mdPermissionsProfile["enable-default-profile-for-global-domains"] = v.(bool)
	}

	if ok := d.HasChange("default_profile_global_domains"); ok {
		mdPermissionsProfile["default-profile-global-domains"] = d.Get("default_profile_global_domains")
	}

	if v, ok := d.GetOkExists("view_global_objects_in_domain"); ok {
		mdPermissionsProfile["view-global-objects-in-domain"] = v.(bool)
	}

	if v, ok := d.GetOkExists("enable_default_profile_for_local_domains"); ok {
		mdPermissionsProfile["enable-default-profile-for-local-domains"] = v.(bool)
	}

	if ok := d.HasChange("default_profile_local_domains"); ok {
		mdPermissionsProfile["default-profile-local-domains"] = d.Get("default_profile_local_domains")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			mdPermissionsProfile["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			mdPermissionsProfile["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		mdPermissionsProfile["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		mdPermissionsProfile["comments"] = d.Get("comments")
	}

	if d.HasChange("domains_to_process") {
		if v, ok := d.GetOk("domains_to_process"); ok {
			mdPermissionsProfile["domains_to_process"] = v.(*schema.Set).List()
		} else {
			oldDomains_To_Process, _ := d.GetChange("domains_to_process")
			mdPermissionsProfile["domains_to_process"] = map[string]interface{}{"remove": oldDomains_To_Process.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		mdPermissionsProfile["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		mdPermissionsProfile["ignore-errors"] = v.(bool)
	}

	if ok := d.HasChange("permission_level"); ok {
		mdPermissionsProfile["permission-level"] = d.Get("permission_level")
	}

	log.Println("Update MdPermissionsProfile - Map = ", mdPermissionsProfile)

	updateMdPermissionsProfileRes, err := client.ApiCall("set-md-permissions-profile", mdPermissionsProfile, client.GetSessionID(), true, false)
	if err != nil || !updateMdPermissionsProfileRes.Success {
		if updateMdPermissionsProfileRes.ErrorMsg != "" {
			return fmt.Errorf(updateMdPermissionsProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementMdPermissionsProfile(d, m)
}

func deleteManagementMdPermissionsProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	mdPermissionsProfilePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete MdPermissionsProfile")

	deleteMdPermissionsProfileRes, err := client.ApiCall("delete-md-permissions-profile", mdPermissionsProfilePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteMdPermissionsProfileRes.Success {
		if deleteMdPermissionsProfileRes.ErrorMsg != "" {
			return fmt.Errorf(deleteMdPermissionsProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
