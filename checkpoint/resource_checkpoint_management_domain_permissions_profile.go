package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/upgraders"
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"strconv"
)

func resourceManagementDomainPermissionsProfile() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDomainPermissionsProfile,
		Read:   readManagementDomainPermissionsProfile,
		Update: updateManagementDomainPermissionsProfile,
		Delete: deleteManagementDomainPermissionsProfile,
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementDomainPermissionsProfileV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementDomainPermissionsProfileStateUpgradeV0,
				Version: 0,
			},
		},
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
				Type:        schema.TypeList,
				MaxItems:    1,
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

func createManagementDomainPermissionsProfile(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	domainPermissionsProfile := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		domainPermissionsProfile["name"] = v.(string)
	}

	if v, ok := d.GetOk("permission_type"); ok {
		domainPermissionsProfile["permission-type"] = v.(string)
	}

	if v, ok := d.GetOkExists("edit_common_objects"); ok {
		domainPermissionsProfile["edit-common-objects"] = v.(bool)
	}

	if v, ok := d.GetOk("access_control"); ok {

		accessControlList := v.([]interface{})

		if len(accessControlList) > 0 {

			accessControlPayload := make(map[string]interface{})

			if v, ok := d.GetOk("access_control.0.show_policy"); ok {
				accessControlPayload["show-policy"] = v.(bool)
			}
			if _, ok := d.GetOk("access_control.0.policy_layers"); ok {

				policyLayersPayload := make(map[string]interface{})

				if v, ok := d.GetOk("access_control.0.policy_layers.0.edit_layers"); ok {
					policyLayersPayload["edit-layers"] = v.(string)
				}
				if v, ok := d.GetOk("access_control.0.policy_layers.0.app_control_and_url_filtering"); ok {
					policyLayersPayload["app-control-and-url-filtering"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("access_control.0.policy_layers.0.content_awareness"); ok {
					policyLayersPayload["content-awareness"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("access_control.0.policy_layers.0.firewall"); ok {
					policyLayersPayload["firewall"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("access_control.0.policy_layers.0.mobile_access"); ok {
					policyLayersPayload["mobile-access"] = strconv.FormatBool(v.(bool))
				}
				accessControlPayload["policy-layers"] = policyLayersPayload
			}
			if v, ok := d.GetOk("access_control.0.dlp_policy"); ok {
				accessControlPayload["dlp-policy"] = v.(string)
			}
			if v, ok := d.GetOk("access_control.0.geo_control_policy"); ok {
				accessControlPayload["geo-control-policy"] = v.(string)
			}
			if v, ok := d.GetOk("access_control.0.nat_policy"); ok {
				accessControlPayload["nat-policy"] = v.(string)
			}
			if v, ok := d.GetOk("access_control.0.qos_policy"); ok {
				accessControlPayload["qos-policy"] = v.(string)
			}
			if v, ok := d.GetOk("access_control.0.access_control_objects_and_settings"); ok {
				accessControlPayload["access-control-objects-and-settings"] = v.(string)
			}
			if v, ok := d.GetOk("access_control.0.app_control_and_url_filtering_update"); ok {
				accessControlPayload["app-control-and-url-filtering-update"] = v.(bool)
			}
			if v, ok := d.GetOk("access_control.0.install_policy"); ok {
				accessControlPayload["install-policy"] = v.(bool)
			}
			domainPermissionsProfile["access-control"] = accessControlPayload
		}
	}
	if v, ok := d.GetOk("endpoint"); ok {

		endpointList := v.([]interface{})

		if len(endpointList) > 0 {

			endpointPayload := make(map[string]interface{})

			if v, ok := d.GetOkExists("endpoint.0.manage_policies_and_software_deployment"); ok {
				endpointPayload["manage-policies-and-software-deployment"] = v.(bool)
			}
			if v, ok := d.GetOkExists("endpoint.0.edit_endpoint_policies"); ok {
				endpointPayload["edit-endpoint-policies"] = v.(bool)
			}
			if v, ok := d.GetOkExists("endpoint.0.policies_installation"); ok {
				endpointPayload["policies-installation"] = v.(bool)
			}
			if v, ok := d.GetOkExists("endpoint.0.edit_software_deployment"); ok {
				endpointPayload["edit-software-deployment"] = v.(bool)
			}
			if v, ok := d.GetOkExists("endpoint.0.software_deployment_installation"); ok {
				endpointPayload["software-deployment-installation"] = v.(bool)
			}
			if v, ok := d.GetOkExists("endpoint.0.allow_executing_push_operations"); ok {
				endpointPayload["allow-executing-push-operations"] = v.(bool)
			}
			if v, ok := d.GetOkExists("endpoint.0.authorize_preboot_users"); ok {
				endpointPayload["authorize-preboot-users"] = v.(bool)
			}
			if v, ok := d.GetOkExists("endpoint.0.recovery_media"); ok {
				endpointPayload["recovery-media"] = v.(bool)
			}
			if v, ok := d.GetOkExists("endpoint.0.remote_help"); ok {
				endpointPayload["remote-help"] = v.(bool)
			}
			if v, ok := d.GetOkExists("endpoint.0.reset_computer_data"); ok {
				endpointPayload["reset-computer-data"] = v.(bool)
			}
			domainPermissionsProfile["endpoint"] = endpointPayload
		}
	}

	if v, ok := d.GetOk("events_and_reports"); ok {

		eventsAndReportsList := v.([]interface{})

		if len(eventsAndReportsList) > 0 {

			eventsAndReportsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("events_and_reports.0.smart_event"); ok {
				eventsAndReportsPayload["smart-event"] = v.(string)
			}
			if v, ok := d.GetOk("events_and_reports.0.events"); ok {
				eventsAndReportsPayload["events"] = v.(string)
			}
			if v, ok := d.GetOk("events_and_reports.0.policy"); ok {
				eventsAndReportsPayload["policy"] = v.(string)
			}
			if v, ok := d.GetOkExists("events_and_reports.0.reports"); ok {
				eventsAndReportsPayload["reports"] = v.(bool)
			}
			domainPermissionsProfile["events-and-reports"] = eventsAndReportsPayload
		}
	}

	if v, ok := d.GetOk("gateways"); ok {

		gatewaysList := v.([]interface{})

		if len(gatewaysList) > 0 {

			gatewaysPayload := make(map[string]interface{})

			if v, ok := d.GetOk("gateways.0.smart_update"); ok {
				gatewaysPayload["smart-update"] = v.(string)
			}
			if v, ok := d.GetOk("gateways.0.lsm_gw_db"); ok {
				gatewaysPayload["lsm-gw-db"] = v.(string)
			}
			if v, ok := d.GetOk("gateways.0.manage_provisioning_profiles"); ok {
				gatewaysPayload["manage-provisioning-profiles"] = v.(string)
			}
			if v, ok := d.GetOkExists("gateways.0.vsx_provisioning"); ok {
				gatewaysPayload["vsx-provisioning"] = v.(bool)
			}
			if v, ok := d.GetOkExists("gateways.0.system_backup"); ok {
				gatewaysPayload["system-backup"] = v.(bool)
			}
			if v, ok := d.GetOkExists("gateways.0.system_restore"); ok {
				gatewaysPayload["system-restore"] = v.(bool)
			}
			if v, ok := d.GetOkExists("gateways.0.open_shell"); ok {
				gatewaysPayload["open-shell"] = v.(bool)
			}
			if v, ok := d.GetOkExists("gateways.0.run_one_time_script"); ok {
				gatewaysPayload["run-one-time-script"] = v.(bool)
			}
			if v, ok := d.GetOkExists("gateways.0.run_repository_script"); ok {
				gatewaysPayload["run-repository-script"] = v.(bool)
			}
			if v, ok := d.GetOk("gateways.0.manage_repository_scripts"); ok {
				gatewaysPayload["manage-repository-scripts"] = v.(string)
			}
			domainPermissionsProfile["gateways"] = gatewaysPayload
		}
	}

	if v, ok := d.GetOk("management"); ok {

		managementList := v.([]interface{})

		if len(managementList) > 0 {

			managementPayload := make(map[string]interface{})

			if v, ok := d.GetOk("management.0.cme_operations"); ok {
				managementPayload["cme-operations"] = v.(string)
			}
			if v, ok := d.GetOkExists("management.0.manage_admins"); ok {
				managementPayload["manage-admins"] = v.(bool)
			}
			if v, ok := d.GetOkExists("management.0.management_api_login"); ok {
				managementPayload["management-api-login"] = v.(bool)
			}
			if v, ok := d.GetOkExists("management.0.manage_sessions"); ok {
				managementPayload["manage-sessions"] = v.(bool)
			}
			if v, ok := d.GetOkExists("management.0.high_availability_operations"); ok {
				managementPayload["high-availability-operations"] = v.(bool)
			}
			if v, ok := d.GetOkExists("management.0.approve_or_reject_sessions"); ok {
				managementPayload["approve-or-reject-sessions"] = v.(bool)
			}
			if v, ok := d.GetOkExists("management.0.publish_sessions"); ok {
				managementPayload["publish-sessions"] = v.(bool)
			}
			if v, ok := d.GetOkExists("management.0.manage_integration_with_cloud_services"); ok {
				managementPayload["manage-integration-with-cloud-services"] = v.(bool)
			}
			domainPermissionsProfile["management"] = managementPayload
		}
	}

	if v, ok := d.GetOk("monitoring_and_logging"); ok {

		monitoringAndLoggingList := v.([]interface{})

		if len(monitoringAndLoggingList) > 0 {

			monitoringAndLoggingPayload := make(map[string]interface{})

			if v, ok := d.GetOk("monitoring_and_logging.0.monitoring"); ok {
				monitoringAndLoggingPayload["monitoring"] = v.(string)
			}
			if v, ok := d.GetOk("monitoring_and_logging.0.management_logs"); ok {
				monitoringAndLoggingPayload["management-logs"] = v.(string)
			}
			if v, ok := d.GetOk("monitoring_and_logging.0.track_logs"); ok {
				monitoringAndLoggingPayload["track-logs"] = v.(string)
			}
			if v, ok := d.GetOkExists("monitoring_and_logging.0.app_and_url_filtering_logs"); ok {
				monitoringAndLoggingPayload["app-and-url-filtering-logs"] = v.(bool)
			}
			if v, ok := d.GetOkExists("monitoring_and_logging.0.https_inspection_logs"); ok {
				monitoringAndLoggingPayload["https-inspection-logs"] = v.(bool)
			}
			if v, ok := d.GetOkExists("monitoring_and_logging.0.packet_capture_and_forensics"); ok {
				monitoringAndLoggingPayload["packet-capture-and-forensics"] = v.(bool)
			}
			if v, ok := d.GetOkExists("monitoring_and_logging.0.show_packet_capture_by_default"); ok {
				monitoringAndLoggingPayload["show-packet-capture-by-default"] = v.(bool)
			}
			if v, ok := d.GetOkExists("monitoring_and_logging.0.identities"); ok {
				monitoringAndLoggingPayload["identities"] = v.(bool)
			}
			if v, ok := d.GetOkExists("monitoring_and_logging.0.show_identities_by_default"); ok {
				monitoringAndLoggingPayload["show-identities-by-default"] = v.(bool)
			}
			if v, ok := d.GetOkExists("monitoring_and_logging.0.dlp_logs_including_confidential_fields"); ok {
				monitoringAndLoggingPayload["dlp-logs-including-confidential-fields"] = v.(bool)
			}
			if v, ok := d.GetOkExists("monitoring_and_logging.0.manage_dlp_messages"); ok {
				monitoringAndLoggingPayload["manage-dlp-messages"] = v.(bool)
			}
			domainPermissionsProfile["monitoring-and-logging"] = monitoringAndLoggingPayload
		}
	}

	if v, ok := d.GetOk("threat_prevention"); ok {

		threatPreventionList := v.([]interface{})

		if len(threatPreventionList) > 0 {

			threatPreventionPayload := make(map[string]interface{})

			if v, ok := d.GetOk("threat_prevention.0.policy_layers"); ok {
				threatPreventionPayload["policy-layers"] = v.(string)
			}
			if v, ok := d.GetOk("threat_prevention.0.edit_layers"); ok {
				threatPreventionPayload["edit-layers"] = v.(string)
			}
			if v, ok := d.GetOkExists("threat_prevention.0.edit_settings"); ok {
				threatPreventionPayload["edit-settings"] = v.(bool)
			}
			if v, ok := d.GetOk("threat_prevention.0.policy_exceptions"); ok {
				threatPreventionPayload["policy-exceptions"] = v.(string)
			}
			if v, ok := d.GetOk("threat_prevention.0.profiles"); ok {
				threatPreventionPayload["profiles"] = v.(string)
			}
			if v, ok := d.GetOk("threat_prevention.0.protections"); ok {
				threatPreventionPayload["protections"] = v.(string)
			}
			if v, ok := d.GetOkExists("threat_prevention.0.install_policy"); ok {
				threatPreventionPayload["install-policy"] = v.(bool)
			}
			if v, ok := d.GetOkExists("threat_prevention.0.ips_update"); ok {
				threatPreventionPayload["ips-update"] = v.(bool)
			}
			domainPermissionsProfile["threat-prevention"] = threatPreventionPayload
		}
	}

	if v, ok := d.GetOk("others"); ok {

		othersList := v.([]interface{})

		if len(othersList) > 0 {

			othersPayload := make(map[string]interface{})

			if v, ok := d.GetOkExists("others.0.client_certificates"); ok {
				othersPayload["client-certificates"] = v.(bool)
			}
			if v, ok := d.GetOkExists("others.0.edit_cp_users_db"); ok {
				othersPayload["edit-cp-users-db"] = v.(bool)
			}
			if v, ok := d.GetOk("others.0.https_inspection"); ok {
				othersPayload["https-inspection"] = v.(string)
			}
			if v, ok := d.GetOk("others.0.ldap_users_db"); ok {
				othersPayload["ldap-users-db"] = v.(string)
			}
			if v, ok := d.GetOk("others.0.user_authority_access"); ok {
				othersPayload["user-authority-access"] = v.(string)
			}
			if v, ok := d.GetOk("others.0.user_device_mgmt_conf"); ok {
				othersPayload["user-device-mgmt-conf"] = v.(string)
			}
			domainPermissionsProfile["others"] = othersPayload
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		domainPermissionsProfile["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		domainPermissionsProfile["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		domainPermissionsProfile["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		domainPermissionsProfile["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		domainPermissionsProfile["ignore-errors"] = v.(bool)
	}

	log.Println("Create DomainPermissionsProfile - Map = ", domainPermissionsProfile)

	addDomainPermissionsProfileRes, err := client.ApiCall("add-domain-permissions-profile", domainPermissionsProfile, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addDomainPermissionsProfileRes.Success {
		if addDomainPermissionsProfileRes.ErrorMsg != "" {
			return fmt.Errorf("%s", addDomainPermissionsProfileRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}

	d.SetId(addDomainPermissionsProfileRes.GetData()["uid"].(string))

	return readManagementDomainPermissionsProfile(d, m)
}

func readManagementDomainPermissionsProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showDomainPermissionsProfileRes, err := client.ApiCall("show-domain-permissions-profile", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showDomainPermissionsProfileRes.Success {
		if objectNotFound(showDomainPermissionsProfileRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showDomainPermissionsProfileRes.ErrorMsg)
	}

	domainPermissionsProfile := showDomainPermissionsProfileRes.GetData()

	log.Println("Read DomainPermissionsProfile - Show JSON = ", domainPermissionsProfile)

	if v := domainPermissionsProfile["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := domainPermissionsProfile["permission-type"]; v != nil {
		_ = d.Set("permission_type", v)
	}

	if v := domainPermissionsProfile["edit-common-objects"]; v != nil {
		_ = d.Set("edit_common_objects", v)
	}

	if domainPermissionsProfile["access-control"] != nil {

		accessControlMap, ok := domainPermissionsProfile["access-control"].(map[string]interface{})

		if ok {
			accessControlMapToReturn := make(map[string]interface{})

			if v := accessControlMap["show-policy"]; v != nil {
				accessControlMapToReturn["show_policy"] = v.(bool)
			}
			if v, ok := accessControlMap["policy-layers"]; ok {

				policyLayersMap, ok := v.(map[string]interface{})
				if ok {
					policyLayersMapToReturn := make(map[string]interface{})

					if v, _ := policyLayersMap["edit-layers"]; v != nil {
						policyLayersMapToReturn["edit_layers"] = v.(string)
					}
					if v, _ := policyLayersMap["app-control-and-url-filtering"]; v != nil {
						policyLayersMapToReturn["app_control_and_url_filtering"] = v.(bool)
					}
					if v, _ := policyLayersMap["content-awareness"]; v != nil {
						policyLayersMapToReturn["content_awareness"] = v.(bool)
					}
					if v, _ := policyLayersMap["firewall"]; v != nil {
						policyLayersMapToReturn["firewall"] = v.(bool)
					}
					if v, _ := policyLayersMap["mobile-access"]; v != nil {
						policyLayersMapToReturn["mobile_access"] = v.(bool)
					}
					accessControlMapToReturn["policy_layers"] = []interface{}{policyLayersMapToReturn}
				}
			}
			if v := accessControlMap["dlp-policy"]; v != nil {
				accessControlMapToReturn["dlp_policy"] = v.(string)
			}
			if v := accessControlMap["geo-control-policy"]; v != nil {
				accessControlMapToReturn["geo_control_policy"] = v.(string)
			}
			if v := accessControlMap["nat-policy"]; v != nil {
				accessControlMapToReturn["nat_policy"] = v.(string)
			}
			if v := accessControlMap["qos-policy"]; v != nil {
				accessControlMapToReturn["qos_policy"] = v.(string)
			}
			if v := accessControlMap["access-control-objects-and-settings"]; v != nil {
				accessControlMapToReturn["access_control_objects_and_settings"] = v.(string)
			}
			if v := accessControlMap["app-control-and-url-filtering-update"]; v != nil {
				accessControlMapToReturn["app_control_and_url_filtering_update"] = v.(bool)
			}
			if v := accessControlMap["install-policy"]; v != nil {
				accessControlMapToReturn["install_policy"] = v.(bool)
			}
			_ = d.Set("access_control", []interface{}{accessControlMapToReturn})

		}
	} else {
		_ = d.Set("access_control", nil)
	}

	if domainPermissionsProfile["endpoint"] != nil {

		endpointMap := domainPermissionsProfile["endpoint"].(map[string]interface{})

		endpointMapToReturn := make(map[string]interface{})

		if v := endpointMap["manage-policies-and-software-deployment"]; v != nil {
			endpointMapToReturn["manage_policies_and_software_deployment"] = v.(bool)
		}
		if v := endpointMap["edit-endpoint-policies"]; v != nil {
			endpointMapToReturn["edit_endpoint_policies"] = v.(bool)
		}
		if v := endpointMap["policies-installation"]; v != nil {
			endpointMapToReturn["policies_installation"] = v.(bool)
		}
		if v := endpointMap["edit-software-deployment"]; v != nil {
			endpointMapToReturn["edit_software_deployment"] = v.(bool)
		}
		if v := endpointMap["software-deployment-installation"]; v != nil {
			endpointMapToReturn["software_deployment_installation"] = v.(bool)
		}
		if v := endpointMap["allow-executing-push-operations"]; v != nil {
			endpointMapToReturn["allow_executing_push_operations"] = v.(bool)
		}
		if v := endpointMap["authorize-preboot-users"]; v != nil {
			endpointMapToReturn["authorize_preboot_users"] = v.(bool)
		}
		if v := endpointMap["recovery-media"]; v != nil {
			endpointMapToReturn["recovery_media"] = v.(bool)
		}
		if v := endpointMap["remote-help"]; v != nil {
			endpointMapToReturn["remote_help"] = v.(bool)
		}
		if v := endpointMap["reset-computer-data"]; v != nil {
			endpointMapToReturn["reset_computer_data"] = v.(bool)
		}
		_ = d.Set("endpoint", []interface{}{endpointMapToReturn})

	} else {
		_ = d.Set("endpoint", nil)
	}

	if domainPermissionsProfile["events-and-reports"] != nil {

		eventsAndReportsMap := domainPermissionsProfile["events-and-reports"].(map[string]interface{})

		eventsAndReportsMapToReturn := make(map[string]interface{})

		if v, _ := eventsAndReportsMap["smart-event"]; v != nil {
			eventsAndReportsMapToReturn["smart_event"] = v.(string)
		}
		if v, _ := eventsAndReportsMap["events"]; v != nil {
			eventsAndReportsMapToReturn["events"] = v.(string)
		}
		if v, _ := eventsAndReportsMap["policy"]; v != nil {
			eventsAndReportsMapToReturn["policy"] = v.(string)
		}
		if v, _ := eventsAndReportsMap["reports"]; v != nil {
			eventsAndReportsMapToReturn["reports"] = v.(bool)
		}
		_ = d.Set("events_and_reports", []interface{}{eventsAndReportsMapToReturn})

	} else {
		_ = d.Set("events_and_reports", nil)
	}

	if domainPermissionsProfile["gateways"] != nil {

		gatewaysMap := domainPermissionsProfile["gateways"].(map[string]interface{})

		gatewaysMapToReturn := make(map[string]interface{})

		if v, _ := gatewaysMap["smart-update"]; v != nil {
			gatewaysMapToReturn["smart_update"] = v.(string)
		}
		if v, _ := gatewaysMap["lsm-gw-db"]; v != nil {
			gatewaysMapToReturn["lsm_gw_db"] = v.(string)
		}
		if v, _ := gatewaysMap["manage-provisioning-profiles"]; v != nil {
			gatewaysMapToReturn["manage_provisioning_profiles"] = v.(string)
		}
		if v, _ := gatewaysMap["vsx-provisioning"]; v != nil {
			gatewaysMapToReturn["vsx_provisioning"] = v.(bool)
		}
		if v, _ := gatewaysMap["system-backup"]; v != nil {
			gatewaysMapToReturn["system_backup"] = v.(bool)
		}
		if v, _ := gatewaysMap["system-restore"]; v != nil {
			gatewaysMapToReturn["system_restore"] = v.(bool)
		}
		if v, _ := gatewaysMap["open-shell"]; v != nil {
			gatewaysMapToReturn["open_shell"] = v.(bool)
		}
		if v, _ := gatewaysMap["run-one-time-script"]; v != nil {
			gatewaysMapToReturn["run_one_time_script"] = v.(bool)
		}
		if v, _ := gatewaysMap["run-repository-script"]; v != nil {
			gatewaysMapToReturn["run_repository_script"] = v.(bool)
		}
		if v, _ := gatewaysMap["manage-repository-scripts"]; v != nil {
			gatewaysMapToReturn["manage_repository_scripts"] = v.(string)
		}
		_ = d.Set("gateways", []interface{}{gatewaysMapToReturn})

	} else {
		_ = d.Set("gateways", nil)
	}

	if domainPermissionsProfile["management"] != nil {

		managementMap := domainPermissionsProfile["management"].(map[string]interface{})

		managementMapToReturn := make(map[string]interface{})

		if v, _ := managementMap["cme-operations"]; v != nil {
			managementMapToReturn["cme_operations"] = v.(string)
		}
		if v, _ := managementMap["manage-admins"]; v != nil {
			managementMapToReturn["manage_admins"] = v.(bool)
		}
		if v, _ := managementMap["management-api-login"]; v != nil {
			managementMapToReturn["management_api_login"] = v.(bool)
		}
		if v, _ := managementMap["manage-sessions"]; v != nil {
			managementMapToReturn["manage_sessions"] = v.(bool)
		}
		if v, _ := managementMap["high-availability-operations"]; v != nil {
			managementMapToReturn["high_availability_operations"] = v.(bool)
		}
		if v, _ := managementMap["approve-or-reject-sessions"]; v != nil {
			managementMapToReturn["approve_or_reject_sessions"] = v.(bool)
		}
		if v, _ := managementMap["publish-sessions"]; v != nil {
			managementMapToReturn["publish_sessions"] = v.(bool)
		}
		if v, _ := managementMap["manage-integration-with-cloud-services"]; v != nil {
			managementMapToReturn["manage_integration_with_cloud_services"] = v.(bool)
		}
		_ = d.Set("management", []interface{}{managementMapToReturn})

	} else {
		_ = d.Set("management", nil)
	}

	if domainPermissionsProfile["monitoring-and-logging"] != nil {

		monitoringAndLoggingMap := domainPermissionsProfile["monitoring-and-logging"].(map[string]interface{})

		monitoringAndLoggingMapToReturn := make(map[string]interface{})

		if v, _ := monitoringAndLoggingMap["monitoring"]; v != nil {
			monitoringAndLoggingMapToReturn["monitoring"] = v.(string)
		}
		if v, _ := monitoringAndLoggingMap["management-logs"]; v != nil {
			monitoringAndLoggingMapToReturn["management_logs"] = v.(string)
		}
		if v, _ := monitoringAndLoggingMap["track-logs"]; v != nil {
			monitoringAndLoggingMapToReturn["track_logs"] = v.(string)
		}
		if v, _ := monitoringAndLoggingMap["app-and-url-filtering-logs"]; v != nil {
			monitoringAndLoggingMapToReturn["app_and_url_filtering_logs"] = v.(bool)
		}
		if v, _ := monitoringAndLoggingMap["https-inspection-logs"]; v != nil {
			monitoringAndLoggingMapToReturn["https_inspection_logs"] = v.(bool)
		}
		if v, _ := monitoringAndLoggingMap["packet-capture-and-forensics"]; v != nil {
			monitoringAndLoggingMapToReturn["packet_capture_and_forensics"] = v.(bool)
		}
		if v, _ := monitoringAndLoggingMap["show-packet-capture-by-default"]; v != nil {
			monitoringAndLoggingMapToReturn["show_packet_capture_by_default"] = v.(bool)
		}
		if v, _ := monitoringAndLoggingMap["identities"]; v != nil {
			monitoringAndLoggingMapToReturn["identities"] = v.(bool)
		}
		if v, _ := monitoringAndLoggingMap["show-identities-by-default"]; v != nil {
			monitoringAndLoggingMapToReturn["show_identities_by_default"] = v.(bool)
		}
		if v, _ := monitoringAndLoggingMap["dlp-logs-including-confidential-fields"]; v != nil {
			monitoringAndLoggingMapToReturn["dlp_logs_including_confidential_fields"] = v.(bool)
		}
		if v, _ := monitoringAndLoggingMap["manage-dlp-messages"]; v != nil {
			monitoringAndLoggingMapToReturn["manage_dlp_messages"] = v.(bool)
		}
		_ = d.Set("monitoring_and_logging", []interface{}{monitoringAndLoggingMapToReturn})

	} else {
		_ = d.Set("monitoring_and_logging", nil)
	}

	if domainPermissionsProfile["threat-prevention"] != nil {

		threatPreventionMap := domainPermissionsProfile["threat-prevention"].(map[string]interface{})

		threatPreventionMapToReturn := make(map[string]interface{})

		if v, _ := threatPreventionMap["policy-layers"]; v != nil {
			threatPreventionMapToReturn["policy_layers"] = v.(string)
		}
		if v, _ := threatPreventionMap["edit-layers"]; v != nil {
			threatPreventionMapToReturn["edit_layers"] = v.(string)
		}
		if v, _ := threatPreventionMap["edit-settings"]; v != nil {
			threatPreventionMapToReturn["edit_settings"] = v.(bool)
		}
		if v, _ := threatPreventionMap["policy-exceptions"]; v != nil {
			threatPreventionMapToReturn["policy_exceptions"] = v.(string)
		}
		if v, _ := threatPreventionMap["profiles"]; v != nil {
			threatPreventionMapToReturn["profiles"] = v.(string)
		}
		if v, _ := threatPreventionMap["protections"]; v != nil {
			threatPreventionMapToReturn["protections"] = v.(string)
		}
		if v, _ := threatPreventionMap["install-policy"]; v != nil {
			threatPreventionMapToReturn["install_policy"] = v.(bool)
		}
		if v, _ := threatPreventionMap["ips-update"]; v != nil {
			threatPreventionMapToReturn["ips_update"] = v.(bool)
		}
		_ = d.Set("threat_prevention", []interface{}{threatPreventionMapToReturn})

	} else {
		_ = d.Set("threat_prevention", nil)
	}

	if domainPermissionsProfile["others"] != nil {

		othersMap := domainPermissionsProfile["others"].(map[string]interface{})

		othersMapToReturn := make(map[string]interface{})

		if v, _ := othersMap["client-certificates"]; v != nil {
			othersMapToReturn["client_certificates"] = v.(bool)
		}
		if v, _ := othersMap["edit-cp-users-db"]; v != nil {
			othersMapToReturn["edit_cp_users_db"] = v.(bool)
		}
		if v, _ := othersMap["https-inspection"]; v != nil {
			othersMapToReturn["https_inspection"] = v.(string)
		}
		if v, _ := othersMap["ldap-users-db"]; v != nil {
			othersMapToReturn["ldap_users_db"] = v.(string)
		}
		if v, _ := othersMap["user-authority-access"]; v != nil {
			othersMapToReturn["user_authority_access"] = v.(string)
		}
		if v, _ := othersMap["user-device-mgmt-conf"]; v != nil {
			othersMapToReturn["user_device_mgmt_conf"] = v.(string)
		}
		_ = d.Set("others", []interface{}{othersMapToReturn})

	} else {
		_ = d.Set("others", nil)
	}

	if domainPermissionsProfile["tags"] != nil {
		tagsJson, ok := domainPermissionsProfile["tags"].([]interface{})
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

	if v := domainPermissionsProfile["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := domainPermissionsProfile["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := domainPermissionsProfile["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := domainPermissionsProfile["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDomainPermissionsProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	domainPermissionsProfile := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		domainPermissionsProfile["name"] = oldName
		domainPermissionsProfile["new-name"] = newName
	} else {
		domainPermissionsProfile["name"] = d.Get("name")
	}

	if ok := d.HasChange("permission_type"); ok {
		domainPermissionsProfile["permission-type"] = d.Get("permission_type")
	}

	if v, ok := d.GetOkExists("edit_common_objects"); ok {
		domainPermissionsProfile["edit-common-objects"] = v.(bool)
	}

	if d.HasChange("access_control") {

		if v, ok := d.GetOk("access_control"); ok {

			accessControlList := v.([]interface{})

			if len(accessControlList) > 0 {

				accessControlPayload := make(map[string]interface{})

				if d.HasChange("access_control.0.show_policy") {
					accessControlPayload["show-policy"] = d.Get("access_control.0.show_policy").(bool)
				}
				if d.HasChange("access_control.0.policy_layers") {

					policyLayersPayload := make(map[string]interface{})

					if d.HasChange("access_control.0.policy_layers.0.edit_layers") {
						policyLayersPayload["edit-layers"] = d.Get("access_control.0.policy_layers.0.edit_layers").(string)
					}
					if d.HasChange("access_control.0.policy_layers.0.app_control_and_url_filtering") {
						policyLayersPayload["app-control-and-url-filtering"] = d.Get("access_control.0.policy_layers.0.app_control_and_url_filtering")
					}
					if d.HasChange("access_control.0.policy_layers.0.content_awareness") {
						policyLayersPayload["content-awareness"] = d.Get("access_control.0.policy_layers.0.content_awareness")
					}
					if d.HasChange("access_control.0.policy_layers.0.firewall") {
						policyLayersPayload["firewall"] = d.Get("access_control.0.policy_layers.0.firewall")
					}
					if d.HasChange("access_control.0.policy_layers.0.mobile_access") {
						policyLayersPayload["mobile-access"] = d.Get("access_control.0.policy_layers.0.mobile_access")
					}
					accessControlPayload["policy-layers"] = policyLayersPayload
				}
				if d.HasChange("access_control.0.dlp_policy") {
					accessControlPayload["dlp-policy"] = d.Get("access_control.0.dlp_policy").(string)
				}
				if d.HasChange("access_control.0.geo_control_policy") {
					accessControlPayload["geo-control-policy"] = d.Get("access_control.0.geo_control_policy").(string)
				}
				if d.HasChange("access_control.0.nat_policy") {
					accessControlPayload["nat-policy"] = d.Get("access_control.0.nat_policy").(string)
				}
				if d.HasChange("access_control.0.qos_policy") {
					accessControlPayload["qos-policy"] = d.Get("access_control.0.qos_policy").(string)
				}
				if d.HasChange("access_control.0.access_control_objects_and_settings") {
					accessControlPayload["access-control-objects-and-settings"] = d.Get("access_control.0.access_control_objects_and_settings").(string)
				}
				if d.HasChange("access_control.0.app_control_and_url_filtering_update") {
					accessControlPayload["app-control-and-url-filtering-update"] = d.Get("access_control.0.app_control_and_url_filtering_update").(bool)
				}
				if d.HasChange("access_control.0.install_policy") {
					accessControlPayload["install-policy"] = d.Get("access_control.0.install_policy").(bool)
				}
				domainPermissionsProfile["access-control"] = accessControlPayload
			}
		}
	}

	if d.HasChange("endpoint") {

		if v, ok := d.GetOk("endpoint"); ok {

			endpointList := v.([]interface{})

			if len(endpointList) > 0 {

				endpointPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("endpoint.0.manage_policies_and_software_deployment"); ok {
					endpointPayload["manage-policies-and-software-deployment"] = v.(bool)
				}
				if v, ok := d.GetOkExists("endpoint.0.edit_endpoint_policies"); ok {
					endpointPayload["edit-endpoint-policies"] = v.(bool)
				}
				if v, ok := d.GetOkExists("endpoint.0.policies_installation"); ok {
					endpointPayload["policies-installation"] = v.(bool)
				}
				if v, ok := d.GetOkExists("endpoint.0.edit_software_deployment"); ok {
					endpointPayload["edit-software-deployment"] = v.(bool)
				}
				if v, ok := d.GetOkExists("endpoint.0.software_deployment_installation"); ok {
					endpointPayload["software-deployment-installation"] = v.(bool)
				}
				if v, ok := d.GetOkExists("endpoint.0.allow_executing_push_operations"); ok {
					endpointPayload["allow-executing-push-operations"] = v.(bool)
				}
				if v, ok := d.GetOkExists("endpoint.0.authorize_preboot_users"); ok {
					endpointPayload["authorize-preboot-users"] = v.(bool)
				}
				if v, ok := d.GetOkExists("endpoint.0.recovery_media"); ok {
					endpointPayload["recovery-media"] = v.(bool)
				}
				if v, ok := d.GetOkExists("endpoint.0.remote_help"); ok {
					endpointPayload["remote-help"] = v.(bool)
				}
				if v, ok := d.GetOkExists("endpoint.0.reset_computer_data"); ok {
					endpointPayload["reset-computer-data"] = v.(bool)
				}
				domainPermissionsProfile["endpoint"] = endpointPayload
			}
		}
	}

	if d.HasChange("events_and_reports") {

		if v, ok := d.GetOk("events_and_reports"); ok {

			eventsAndReportsList := v.([]interface{})

			if len(eventsAndReportsList) > 0 {

				eventsAndReportsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("events_and_reports.0.smart_event"); ok {
					eventsAndReportsPayload["smart-event"] = v.(string)
				}
				if v, ok := d.GetOk("events_and_reports.0.events"); ok {
					eventsAndReportsPayload["events"] = v.(string)
				}
				if v, ok := d.GetOk("events_and_reports.0.policy"); ok {
					eventsAndReportsPayload["policy"] = v.(string)
				}
				if v, ok := d.GetOkExists("events_and_reports.0.reports"); ok {
					eventsAndReportsPayload["reports"] = v.(bool)
				}
				domainPermissionsProfile["events-and-reports"] = eventsAndReportsPayload
			}
		}
	}

	if d.HasChange("gateways") {

		if v, ok := d.GetOk("gateways"); ok {

			gatewaysList := v.([]interface{})

			if len(gatewaysList) > 0 {

				gatewaysPayload := make(map[string]interface{})

				if v, ok := d.GetOk("gateways.0.smart_update"); ok {
					gatewaysPayload["smart-update"] = v.(string)
				}
				if v, ok := d.GetOk("gateways.0.lsm_gw_db"); ok {
					gatewaysPayload["lsm-gw-db"] = v.(string)
				}
				if v, ok := d.GetOk("gateways.0.manage_provisioning_profiles"); ok {
					gatewaysPayload["manage-provisioning-profiles"] = v.(string)
				}
				if v, ok := d.GetOkExists("gateways.0.vsx_provisioning"); ok {
					gatewaysPayload["vsx-provisioning"] = v.(bool)
				}
				if v, ok := d.GetOkExists("gateways.0.system_backup"); ok {
					gatewaysPayload["system-backup"] = v.(bool)
				}
				if v, ok := d.GetOkExists("gateways.0.system_restore"); ok {
					gatewaysPayload["system-restore"] = v.(bool)
				}
				if v, ok := d.GetOkExists("gateways.0.open_shell"); ok {
					gatewaysPayload["open-shell"] = v.(bool)
				}
				if v, ok := d.GetOkExists("gateways.0.run_one_time_script"); ok {
					gatewaysPayload["run-one-time-script"] = v.(bool)
				}
				if v, ok := d.GetOkExists("gateways.0.run_repository_script"); ok {
					gatewaysPayload["run-repository-script"] = v.(bool)
				}
				if v, ok := d.GetOk("gateways.0.manage_repository_scripts"); ok {
					gatewaysPayload["manage-repository-scripts"] = v.(string)
				}
				domainPermissionsProfile["gateways"] = gatewaysPayload
			}
		}
	}

	if d.HasChange("management") {

		if v, ok := d.GetOk("management"); ok {

			managementList := v.([]interface{})

			if len(managementList) > 0 {

				managementPayload := make(map[string]interface{})

				if v, ok := d.GetOk("management.0.cme_operations"); ok {
					managementPayload["cme-operations"] = v.(string)
				}
				if v, ok := d.GetOkExists("management.0.manage_admins"); ok {
					managementPayload["manage-admins"] = v.(bool)
				}
				if v, ok := d.GetOkExists("management.0.management_api_login"); ok {
					managementPayload["management-api-login"] = v.(bool)
				}
				if v, ok := d.GetOkExists("management.0.manage_sessions"); ok {
					managementPayload["manage-sessions"] = v.(bool)
				}
				if v, ok := d.GetOkExists("management.0.high_availability_operations"); ok {
					managementPayload["high-availability-operations"] = v.(bool)
				}
				if v, ok := d.GetOkExists("management.0.approve_or_reject_sessions"); ok {
					managementPayload["approve-or-reject-sessions"] = v.(bool)
				}
				if v, ok := d.GetOkExists("management.0.publish_sessions"); ok {
					managementPayload["publish-sessions"] = v.(bool)
				}
				if v, ok := d.GetOkExists("management.0.manage_integration_with_cloud_services"); ok {
					managementPayload["manage-integration-with-cloud-services"] = v.(bool)
				}
				domainPermissionsProfile["management"] = managementPayload
			}
		}
	}

	if d.HasChange("monitoring_and_logging") {

		if v, ok := d.GetOk("monitoring_and_logging"); ok {

			monitoringAndLoggingList := v.([]interface{})

			if len(monitoringAndLoggingList) > 0 {

				monitoringAndLoggingPayload := make(map[string]interface{})

				if v, ok := d.GetOk("monitoring_and_logging.0.monitoring"); ok {
					monitoringAndLoggingPayload["monitoring"] = v.(string)
				}
				if v, ok := d.GetOk("monitoring_and_logging.0.management_logs"); ok {
					monitoringAndLoggingPayload["management-logs"] = v.(string)
				}
				if v, ok := d.GetOk("monitoring_and_logging.0.track_logs"); ok {
					monitoringAndLoggingPayload["track-logs"] = v.(string)
				}
				if v, ok := d.GetOkExists("monitoring_and_logging.0.app_and_url_filtering_logs"); ok {
					monitoringAndLoggingPayload["app-and-url-filtering-logs"] = v.(bool)
				}
				if v, ok := d.GetOkExists("monitoring_and_logging.0.https_inspection_logs"); ok {
					monitoringAndLoggingPayload["https-inspection-logs"] = v.(bool)
				}
				if v, ok := d.GetOkExists("monitoring_and_logging.0.packet_capture_and_forensics"); ok {
					monitoringAndLoggingPayload["packet-capture-and-forensics"] = v.(bool)
				}
				if v, ok := d.GetOkExists("monitoring_and_logging.0.show_packet_capture_by_default"); ok {
					monitoringAndLoggingPayload["show-packet-capture-by-default"] = v.(bool)
				}
				if v, ok := d.GetOkExists("monitoring_and_logging.0.identities"); ok {
					monitoringAndLoggingPayload["identities"] = v.(bool)
				}
				if v, ok := d.GetOkExists("monitoring_and_logging.0.show_identities_by_default"); ok {
					monitoringAndLoggingPayload["show-identities-by-default"] = v.(bool)
				}
				if v, ok := d.GetOkExists("monitoring_and_logging.0.dlp_logs_including_confidential_fields"); ok {
					monitoringAndLoggingPayload["dlp-logs-including-confidential-fields"] = v.(bool)
				}
				if v, ok := d.GetOkExists("monitoring_and_logging.0.manage_dlp_messages"); ok {
					monitoringAndLoggingPayload["manage-dlp-messages"] = v.(bool)
				}
				domainPermissionsProfile["monitoring-and-logging"] = monitoringAndLoggingPayload
			}
		}
	}

	if d.HasChange("threat_prevention") {

		if v, ok := d.GetOk("threat_prevention"); ok {

			threatPreventionList := v.([]interface{})

			if len(threatPreventionList) > 0 {

				threatPreventionPayload := make(map[string]interface{})

				if v, ok := d.GetOk("threat_prevention.0.policy_layers"); ok {
					threatPreventionPayload["policy-layers"] = v.(string)
				}
				if v, ok := d.GetOk("threat_prevention.0.edit_layers"); ok {
					threatPreventionPayload["edit-layers"] = v.(string)
				}
				if v, ok := d.GetOkExists("threat_prevention.0.edit_settings"); ok {
					threatPreventionPayload["edit-settings"] = v.(bool)
				}
				if v, ok := d.GetOk("threat_prevention.0.policy_exceptions"); ok {
					threatPreventionPayload["policy-exceptions"] = v.(string)
				}
				if v, ok := d.GetOk("threat_prevention.0.profiles"); ok {
					threatPreventionPayload["profiles"] = v.(string)
				}
				if v, ok := d.GetOk("threat_prevention.0.protections"); ok {
					threatPreventionPayload["protections"] = v.(string)
				}
				if v, ok := d.GetOkExists("threat_prevention.0.install_policy"); ok {
					threatPreventionPayload["install-policy"] = v.(bool)
				}
				if v, ok := d.GetOkExists("threat_prevention.0.ips_update"); ok {
					threatPreventionPayload["ips-update"] = v.(bool)
				}
				domainPermissionsProfile["threat-prevention"] = threatPreventionPayload
			}
		}
	}

	if d.HasChange("others") {

		if v, ok := d.GetOk("others"); ok {

			othersList := v.([]interface{})

			if len(othersList) > 0 {

				othersPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("others.0.client_certificates"); ok {
					othersPayload["client-certificates"] = v.(bool)
				}
				if v, ok := d.GetOkExists("others.0.edit_cp_users_db"); ok {
					othersPayload["edit-cp-users-db"] = v.(bool)
				}
				if v, ok := d.GetOk("others.0.https_inspection"); ok {
					othersPayload["https-inspection"] = v.(string)
				}
				if v, ok := d.GetOk("others.0.ldap_users_db"); ok {
					othersPayload["ldap-users-db"] = v.(string)
				}
				if v, ok := d.GetOk("others.0.user_authority_access"); ok {
					othersPayload["user-authority-access"] = v.(string)
				}
				if v, ok := d.GetOk("others.0.user_device_mgmt_conf"); ok {
					othersPayload["user-device-mgmt-conf"] = v.(string)
				}
				domainPermissionsProfile["others"] = othersPayload
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			domainPermissionsProfile["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			domainPermissionsProfile["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		domainPermissionsProfile["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		domainPermissionsProfile["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		domainPermissionsProfile["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		domainPermissionsProfile["ignore-errors"] = v.(bool)
	}

	log.Println("Update DomainPermissionsProfile - Map = ", domainPermissionsProfile)

	updateDomainPermissionsProfileRes, err := client.ApiCall("set-domain-permissions-profile", domainPermissionsProfile, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateDomainPermissionsProfileRes.Success {
		if updateDomainPermissionsProfileRes.ErrorMsg != "" {
			return fmt.Errorf("%s", updateDomainPermissionsProfileRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}

	return readManagementDomainPermissionsProfile(d, m)
}

func deleteManagementDomainPermissionsProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	domainPermissionsProfilePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		domainPermissionsProfilePayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		domainPermissionsProfilePayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete DomainPermissionsProfile")

	deleteDomainPermissionsProfileRes, err := client.ApiCall("delete-domain-permissions-profile", domainPermissionsProfilePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteDomainPermissionsProfileRes.Success {
		if deleteDomainPermissionsProfileRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteDomainPermissionsProfileRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
