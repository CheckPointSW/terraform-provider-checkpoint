package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementDomainPermissionsProfileV0 is the V0 schema where endpoint, events_and_reports,
// gateways, management, monitoring_and_logging, threat_prevention, and others were TypeMap.
func ResourceManagementDomainPermissionsProfileV0() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"permission_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type of the Permissions Profile.",
				Default:     "customized",
			},
			"edit_common_objects": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Define and manage objects in the Check Point database: Network Objects, Services, Custom Application Site, VPN Community, Users, Servers, Resources, Time, UserCheck, and Limit.<br>Only a 'Customized' permission-type profile can edit this permission.",
			},
			"access_control": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Access Control permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"show_policy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Select to let administrators work with Access Control rules and NAT rules. If not selected, administrators cannot see these rules.",
						},
						"policy_layers": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Layer editing permissions.<br>Available only if show-policy is set to true.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"edit_layers": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "\"By Software Blades\" - Edit Access Control layers that contain the blades enabled in the Permissions Profile.<br>\"By Selected Profile In A Layer Editor\" - Administrators can only edit the layer if the Access Control layer editor gives editing permission to their profiles.",
									},
									"app_control_and_url_filtering": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Use Application and URL Filtering in Access Control rules.<br>Available only if edit-layers is set to \"By Software Blades\".",
									},
									"content_awareness": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Use specified data types in Access Control rules.<br>Available only if edit-layers is set to \"By Software Blades\".",
									},
									"firewall": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Work with Access Control and other Software Blades that do not have their own Policies.<br>Available only if edit-layers is set to \"By Software Blades\".",
									},
									"mobile_access": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Work with Mobile Access rules.<br>Available only if edit-layers is set to \"By Software Blades\".",
									},
								},
							},
						},
						"dlp_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Configure DLP rules and Policies.",
						},
						"geo_control_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Work with Access Control rules that control traffic to and from specified countries.",
						},
						"nat_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Work with NAT in Access Control rules.",
						},
						"qos_policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Work with QoS Policies and rules.",
						},
						"access_control_objects_and_settings": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Allow editing of the following objet types: VPN Community, Access Role, Custom application group,Custom application, Custom category, Limit, Application - Match Settings, Application Category - Match Settings,Override Categorization, Application and URL filtering blade - Advanced Settings, Content Awareness blade - Advanced Settings.",
						},
						"app_control_and_url_filtering_update": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Install Application and URL Filtering updates.",
						},
						"install_policy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Install Access Control Policies.",
						},
					},
				},
			},
			"endpoint": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Endpoint permissions. Not supported for Multi-Domain Servers.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_policies_and_software_deployment": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The administrator can work with policies, rules and actions.",
						},
						"edit_endpoint_policies": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Available only if manage-policies-and-software-deployment is set to true.",
						},
						"policies_installation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The administrator can install policies on endpoint computers.",
						},
						"edit_software_deployment": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The administrator can define deployment rules, create packages for export, and configure advanced package settings.<br>Available only if manage-policies-and-software-deployment is set to true.",
						},
						"software_deployment_installation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The administrator can deploy packages and install endpoint clients.",
						},
						"allow_executing_push_operations": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The administrator can start operations that the Security Management Server pushes directly to client computers with no policy installation required.",
						},
						"authorize_preboot_users": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The administrator can add and remove the users who are permitted to log on to Endpoint Security client computers with Full Disk Encryption.",
						},
						"recovery_media": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The administrator can create recovery media on endpoint computers and devices.",
						},
						"remote_help": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The administrator can use the Remote Help feature to reset user passwords and give access to locked out users.",
						},
						"reset_computer_data": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The administrator can reset a computer, which deletes all information about the computer from the Security Management Server.",
						},
					},
				},
			},
			"events_and_reports": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Events and Reports permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"smart_event": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "'Custom' - Configure SmartEvent permissions.",
						},
						"events": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Work with event queries on the Events tab. Create custom event queries.<br>Available only if smart-event is set to 'Custom'.",
						},
						"policy": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Configure SmartEvent Policy rules and install SmartEvent Policies.<br>Available only if smart-event is set to 'Custom'.",
						},
						"reports": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Create and run SmartEvent reports.<br>Available only if smart-event is set to 'Custom'.",
						},
					},
				},
			},
			"gateways": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Gateways permissions. <br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"smart_update": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Install, update and delete Check Point licenses. This includes permissions to use SmartUpdate to manage licenses.",
						},
						"lsm_gw_db": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Access to objects defined in LSM gateway tables. These objects are managed in the SmartProvisioning GUI or LSMcli command-line.<br>Note: 'Write' permission on lsm-gw-db allows administrator to run a script on SmartLSM gateway in Expert mode.",
						},
						"manage_provisioning_profiles": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Administrator can add, edit, delete, and assign provisioning profiles to gateways (both LSM and non-LSM).<br>Available for edit only if lsm-gw-db is set with 'Write' permission.<br>Note: 'Read' permission on lsm-gw-db enables 'Read' permission for manage-provisioning-profiles.",
						},
						"vsx_provisioning": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Create and configure Virtual Systems and other VSX virtual objects.",
						},
						"system_backup": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Backup Security Gateways.",
						},
						"system_restore": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Restore Security Gateways from saved backups.",
						},
						"open_shell": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Use the SmartConsole CLI to run commands.",
						},
						"run_one_time_script": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Run user scripts from the command line.",
						},
						"run_repository_script": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Run scripts from the repository.",
						},
						"manage_repository_scripts": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Add, change and remove scripts in the repository.",
						},
					},
				},
			},
			"management": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Management permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cme_operations": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Permission to read / edit the Cloud Management Extension (CME) configuration.<br>Not supported for Multi-Domain Servers.",
						},
						"manage_admins": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Controls the ability to manage Administrators, Permission Profiles, Trusted clients,API settings and Policy settings.<br>Only a \"Read Write All\" permission-type profile can edit this permission.<br>Not supported for Multi-Domain Servers.",
						},
						"management_api_login": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Permission to log in to the Security Management Server and run API commands using thesetools: mgmt_cli (Linux and Windows binaries), Gaia CLI (clish) and Web Services (REST). Useful if you want to prevent administrators from running automatic scripts on the Management.<br>Note: This permission is not required to run commands from within the API terminal in SmartConsole.<br>Not supported for Multi-Domain Servers.",
						},
						"manage_sessions": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Lets you disconnect, discard, publish, or take over other administrator sessions.<br>Only a \"Read Write All\" permission-type profile can edit this permission.",
						},
						"high_availability_operations": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Configure and work with Domain High Availability.<br>Only a 'Customized' permission-type profile can edit this permission.",
						},
						"approve_or_reject_sessions": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Approve / reject other sessions.",
						},
						"publish_sessions": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Allow session publishing without an approval.",
						},
						"manage_integration_with_cloud_services": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Manage integration with Cloud Services.",
						},
					},
				},
			},
			"monitoring_and_logging": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Monitoring and Logging permissions.<br>'Customized' permission-type profile can edit all these permissions. \"Read Write All\" permission-type can edit only dlp-logs-including-confidential-fields and manage-dlp-messages permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitoring": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "See monitoring views and reports.",
						},
						"management_logs": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "See Multi-Domain Server audit logs.",
						},
						"track_logs": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Use the log tracking features in SmartConsole.",
						},
						"app_and_url_filtering_logs": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Work with Application and URL Filtering logs.",
						},
						"https_inspection_logs": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "See logs generated by HTTPS Inspection.",
						},
						"packet_capture_and_forensics": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "See logs generated by the IPS and Forensics features.",
						},
						"show_packet_capture_by_default": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable packet capture by default.",
						},
						"identities": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Show user and computer identity information in logs.",
						},
						"show_identities_by_default": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Show user and computer identity information in logs by default.",
						},
						"dlp_logs_including_confidential_fields": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Show DLP logs including confidential fields.",
						},
						"manage_dlp_messages": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "View/Release/Discard DLP messages.<br>Available only if dlp-logs-including-confidential-fields is set to true.",
						},
					},
				},
			},
			"threat_prevention": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Threat Prevention permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_layers": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Configure Threat Prevention Policy rules.<br>Note: To have policy-layers permissions you must set policy-exceptionsand profiles permissions. To have 'Write' permissions for policy-layers, policy-exceptions must be set with 'Write' permission as well.",
						},
						"edit_layers": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "'ALL' -  Gives permission to edit all layers.<br>\"By Selected Profile In A Layer Editor\" -  Administrators can only edit the layer if the Threat Prevention layer editor gives editing permission to their profiles.<br>Available only if policy-layers is set to 'Write'.",
						},
						"edit_settings": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Work with general Threat Prevention settings.",
						},
						"policy_exceptions": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Configure exceptions to Threat Prevention rules.<br>Note: To have policy-exceptions you must set the protections permission.",
						},
						"profiles": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Configure Threat Prevention profiles.",
						},
						"protections": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Work with malware protections.",
						},
						"install_policy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Install Policies.",
						},
						"ips_update": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Update IPS protections.<br>Note: You do not have to log into the User Center to receive IPS updates.",
						},
					},
				},
			},
			"others": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Additional permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_certificates": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Create and manage client certificates for Mobile Access.",
						},
						"edit_cp_users_db": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Work with user accounts and groups.",
						},
						"https_inspection": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Enable and configure HTTPS Inspection rules.",
						},
						"ldap_users_db": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Work with the LDAP database and user accounts, groups and OUs.",
						},
						"user_authority_access": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Work with Check Point User Authority authentication.",
						},
						"user_device_mgmt_conf": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Gives access to the UDM (User & Device Management) web-based application that handles security challenges in a \"bring your own device\" (BYOD) workspace.",
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

// ResourceManagementDomainPermissionsProfileStateUpgradeV0 converts the TypeMap fields to TypeList.
func ResourceManagementDomainPermissionsProfileStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState,
		"endpoint", "events_and_reports", "gateways", "management",
		"monitoring_and_logging", "threat_prevention", "others",
	), nil
}
