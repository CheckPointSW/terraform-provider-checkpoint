package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"

	"strconv"
)

func resourceManagementDomainPermissionsProfile() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDomainPermissionsProfile,
		Read:   readManagementDomainPermissionsProfile,
		Update: updateManagementDomainPermissionsProfile,
		Delete: deleteManagementDomainPermissionsProfile,
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
	if _, ok := d.GetOk("endpoint"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("endpoint.manage_policies_and_software_deployment"); ok {
			res["manage-policies-and-software-deployment"] = v
		}
		if v, ok := d.GetOk("endpoint.edit_endpoint_policies"); ok {
			res["edit-endpoint-policies"] = v
		}
		if v, ok := d.GetOk("endpoint.policies_installation"); ok {
			res["policies-installation"] = v
		}
		if v, ok := d.GetOk("endpoint.edit_software_deployment"); ok {
			res["edit-software-deployment"] = v
		}
		if v, ok := d.GetOk("endpoint.software_deployment_installation"); ok {
			res["software-deployment-installation"] = v
		}
		if v, ok := d.GetOk("endpoint.allow_executing_push_operations"); ok {
			res["allow-executing-push-operations"] = v
		}
		if v, ok := d.GetOk("endpoint.authorize_preboot_users"); ok {
			res["authorize-preboot-users"] = v
		}
		if v, ok := d.GetOk("endpoint.recovery_media"); ok {
			res["recovery-media"] = v
		}
		if v, ok := d.GetOk("endpoint.remote_help"); ok {
			res["remote-help"] = v
		}
		if v, ok := d.GetOk("endpoint.reset_computer_data"); ok {
			res["reset-computer-data"] = v
		}
		domainPermissionsProfile["endpoint"] = res
	}

	if _, ok := d.GetOk("events_and_reports"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("events_and_reports.smart_event"); ok {
			res["smart-event"] = v.(string)
		}
		if v, ok := d.GetOk("events_and_reports.events"); ok {
			res["events"] = v.(string)
		}
		if v, ok := d.GetOk("events_and_reports.policy"); ok {
			res["policy"] = v.(string)
		}
		if v, ok := d.GetOk("events_and_reports.reports"); ok {
			res["reports"] = v
		}
		domainPermissionsProfile["events-and-reports"] = res
	}

	if _, ok := d.GetOk("gateways"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("gateways.smart_update"); ok {
			res["smart-update"] = v.(string)
		}
		if v, ok := d.GetOk("gateways.lsm_gw_db"); ok {
			res["lsm-gw-db"] = v.(string)
		}
		if v, ok := d.GetOk("gateways.manage_provisioning_profiles"); ok {
			res["manage-provisioning-profiles"] = v.(string)
		}
		if v, ok := d.GetOk("gateways.vsx_provisioning"); ok {
			res["vsx-provisioning"] = v
		}
		if v, ok := d.GetOk("gateways.system_backup"); ok {
			res["system-backup"] = v
		}
		if v, ok := d.GetOk("gateways.system_restore"); ok {
			res["system-restore"] = v
		}
		if v, ok := d.GetOk("gateways.open_shell"); ok {
			res["open-shell"] = v
		}
		if v, ok := d.GetOk("gateways.run_one_time_script"); ok {
			res["run-one-time-script"] = v
		}
		if v, ok := d.GetOk("gateways.run_repository_script"); ok {
			res["run-repository-script"] = v
		}
		if v, ok := d.GetOk("gateways.manage_repository_scripts"); ok {
			res["manage-repository-scripts"] = v.(string)
		}
		domainPermissionsProfile["gateways"] = res
	}

	if _, ok := d.GetOk("management"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("management.cme_operations"); ok {
			res["cme-operations"] = v.(string)
		}
		if v, ok := d.GetOk("management.manage_admins"); ok {
			res["manage-admins"] = v
		}
		if v, ok := d.GetOk("management.management_api_login"); ok {
			res["management-api-login"] = v
		}
		if v, ok := d.GetOk("management.manage_sessions"); ok {
			res["manage-sessions"] = v
		}
		if v, ok := d.GetOk("management.high_availability_operations"); ok {
			res["high-availability-operations"] = v
		}
		if v, ok := d.GetOk("management.approve_or_reject_sessions"); ok {
			res["approve-or-reject-sessions"] = v
		}
		if v, ok := d.GetOk("management.publish_sessions"); ok {
			res["publish-sessions"] = v
		}
		if v, ok := d.GetOk("management.manage_integration_with_cloud_services"); ok {
			res["manage-integration-with-cloud-services"] = v
		}
		domainPermissionsProfile["management"] = res
	}

	if _, ok := d.GetOk("monitoring_and_logging"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("monitoring_and_logging.monitoring"); ok {
			res["monitoring"] = v.(string)
		}
		if v, ok := d.GetOk("monitoring_and_logging.management_logs"); ok {
			res["management-logs"] = v.(string)
		}
		if v, ok := d.GetOk("monitoring_and_logging.track_logs"); ok {
			res["track-logs"] = v.(string)
		}
		if v, ok := d.GetOk("monitoring_and_logging.app_and_url_filtering_logs"); ok {
			res["app-and-url-filtering-logs"] = v
		}
		if v, ok := d.GetOk("monitoring_and_logging.https_inspection_logs"); ok {
			res["https-inspection-logs"] = v
		}
		if v, ok := d.GetOk("monitoring_and_logging.packet_capture_and_forensics"); ok {
			res["packet-capture-and-forensics"] = v
		}
		if v, ok := d.GetOk("monitoring_and_logging.show_packet_capture_by_default"); ok {
			res["show-packet-capture-by-default"] = v
		}
		if v, ok := d.GetOk("monitoring_and_logging.identities"); ok {
			res["identities"] = v
		}
		if v, ok := d.GetOk("monitoring_and_logging.show_identities_by_default"); ok {
			res["show-identities-by-default"] = v
		}
		if v, ok := d.GetOk("monitoring_and_logging.dlp_logs_including_confidential_fields"); ok {
			res["dlp-logs-including-confidential-fields"] = v
		}
		if v, ok := d.GetOk("monitoring_and_logging.manage_dlp_messages"); ok {
			res["manage-dlp-messages"] = v
		}
		domainPermissionsProfile["monitoring-and-logging"] = res
	}

	if _, ok := d.GetOk("threat_prevention"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("threat_prevention.policy_layers"); ok {
			res["policy-layers"] = v.(string)
		}
		if v, ok := d.GetOk("threat_prevention.edit_layers"); ok {
			res["edit-layers"] = v.(string)
		}
		if v, ok := d.GetOk("threat_prevention.edit_settings"); ok {
			res["edit-settings"] = v
		}
		if v, ok := d.GetOk("threat_prevention.policy_exceptions"); ok {
			res["policy-exceptions"] = v.(string)
		}
		if v, ok := d.GetOk("threat_prevention.profiles"); ok {
			res["profiles"] = v.(string)
		}
		if v, ok := d.GetOk("threat_prevention.protections"); ok {
			res["protections"] = v.(string)
		}
		if v, ok := d.GetOk("threat_prevention.install_policy"); ok {
			res["install-policy"] = v
		}
		if v, ok := d.GetOk("threat_prevention.ips_update"); ok {
			res["ips-update"] = v
		}
		domainPermissionsProfile["threat-prevention"] = res
	}

	if _, ok := d.GetOk("others"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("others.client_certificates"); ok {
			res["client-certificates"] = v
		}
		if v, ok := d.GetOk("others.edit_cp_users_db"); ok {
			res["edit-cp-users-db"] = v
		}
		if v, ok := d.GetOk("others.https_inspection"); ok {
			res["https-inspection"] = v.(string)
		}
		if v, ok := d.GetOk("others.ldap_users_db"); ok {
			res["ldap-users-db"] = v.(string)
		}
		if v, ok := d.GetOk("others.user_authority_access"); ok {
			res["user-authority-access"] = v.(string)
		}
		if v, ok := d.GetOk("others.user_device_mgmt_conf"); ok {
			res["user-device-mgmt-conf"] = v.(string)
		}
		domainPermissionsProfile["others"] = res
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
			return fmt.Errorf(addDomainPermissionsProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
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
		return fmt.Errorf(err.Error())
	}
	if !showDomainPermissionsProfileRes.Success {
		if objectNotFound(showDomainPermissionsProfileRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDomainPermissionsProfileRes.ErrorMsg)
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

		defaultEventsAndReports := map[string]interface{}{
			"show-policy":                          "custom",
			"dlp-policy":                           "write",
			"geo-control-policy":                   "write",
			"nat-policy":                           "true",
			"qos-policy":                           "true",
			"access-control-objects-and-settings":  "true",
			"app-control-and-url-filtering-update": "true",
			"install-policy":                       "true",
		}

		accessControlMap, ok := domainPermissionsProfile["access-control"].(map[string]interface{})

		if ok {
			accessControlMapToReturn := make(map[string]interface{})

			if v := accessControlMap["show-policy"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "access_control.0.show_policy", defaultEventsAndReports["show-policy"].(string)) {
				accessControlMapToReturn["show_policy"] = strconv.FormatBool(v.(bool))
			}
			if v, ok := accessControlMap["policy-layers"]; ok {

				policyLayersMap, ok := v.(map[string]interface{})
				if ok {
					policyLayersMapToReturn := make(map[string]interface{})

					if v, _ := policyLayersMap["edit-layers"]; v != nil {
						policyLayersMapToReturn["edit_layers"] = v
					}
					if v, _ := policyLayersMap["app-control-and-url-filtering"]; v != nil {
						policyLayersMapToReturn["app_control_and_url_filtering"] = v
					}
					if v, _ := policyLayersMap["content-awareness"]; v != nil {
						policyLayersMapToReturn["content_awareness"] = v
					}
					if v, _ := policyLayersMap["firewall"]; v != nil {
						policyLayersMapToReturn["firewall"] = v
					}
					if v, _ := policyLayersMap["mobile-access"]; v != nil {
						policyLayersMapToReturn["mobile_access"] = v
					}
					accessControlMapToReturn["policy_layers"] = []interface{}{policyLayersMapToReturn}
				}
			}
			if v := accessControlMap["dlp-policy"]; v != nil && isArgDefault(v.(string), d, "access_control.0.dlp_policy", defaultEventsAndReports["dlp-policy"].(string)) {
				accessControlMapToReturn["dlp_policy"] = v
			}
			if v := accessControlMap["geo-control-policy"]; v != nil && isArgDefault(v.(string), d, "access_control.0.geo_control_policy", defaultEventsAndReports["geo-control-policy"].(string)) {
				accessControlMapToReturn["geo_control_policy"] = v
			}
			if v := accessControlMap["nat-policy"]; v != nil && isArgDefault(v.(string), d, "access_control.0.nat_policy", defaultEventsAndReports["nat-policy"].(string)) {
				accessControlMapToReturn["nat_policy"] = v
			}
			if v := accessControlMap["qos-policy"]; v != nil && isArgDefault(v.(string), d, "access_control.0.qos_policy", defaultEventsAndReports["qos-policy"].(string)) {
				accessControlMapToReturn["qos_policy"] = v
			}
			if v := accessControlMap["access-control-objects-and-settings"]; v != nil && isArgDefault(v.(string), d, "access_control.0.access_control_objects_and_settings", defaultEventsAndReports["access-control-objects-and-settings"].(string)) {
				accessControlMapToReturn["access_control_objects_and_settings"] = v
			}
			if v := accessControlMap["app-control-and-url-filtering-update"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "access_control.0.app_control_and_url_filtering_update", defaultEventsAndReports["app-control-and-url-filtering-update"].(string)) {
				accessControlMapToReturn["app_control_and_url_filtering_update"] = strconv.FormatBool(v.(bool))
			}
			if v := accessControlMap["install-policy"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "access_control.0.install_policy", defaultEventsAndReports["install-policy"].(string)) {
				accessControlMapToReturn["install_policy"] = v
			}
			_ = d.Set("access_control", []interface{}{accessControlMapToReturn})

		}
	} else {
		_ = d.Set("access_control", nil)
	}

	if domainPermissionsProfile["endpoint"] != nil {

		endpointMap := domainPermissionsProfile["endpoint"].(map[string]interface{})

		endpointMapToReturn := make(map[string]interface{})

		if v, _ := endpointMap["manage-policies-and-software-deployment"]; v != nil {
			endpointMapToReturn["manage_policies_and_software_deployment"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := endpointMap["edit-endpoint-policies"]; v != nil {
			endpointMapToReturn["edit_endpoint_policies"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := endpointMap["policies-installation"]; v != nil {
			endpointMapToReturn["policies_installation"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := endpointMap["edit-software-deployment"]; v != nil {
			endpointMapToReturn["edit_software_deployment"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := endpointMap["software-deployment-installation"]; v != nil {
			endpointMapToReturn["software_deployment_installation"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := endpointMap["allow-executing-push-operations"]; v != nil {
			endpointMapToReturn["allow_executing_push_operations"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := endpointMap["authorize-preboot-users"]; v != nil {
			endpointMapToReturn["authorize_preboot_users"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := endpointMap["recovery-media"]; v != nil {
			endpointMapToReturn["recovery_media"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := endpointMap["remote-help"]; v != nil {
			endpointMapToReturn["remote_help"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := endpointMap["reset-computer-data"]; v != nil {
			endpointMapToReturn["reset_computer_data"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("endpoint", endpointMapToReturn)
	} else {
		_ = d.Set("endpoint", nil)
	}

	if domainPermissionsProfile["events-and-reports"] != nil {

		defaultEventsAndReports := map[string]interface{}{
			"smart-event": "custom",
			"events":      "write",
			"policy":      "write",
			"reports":     "true",
		}

		eventsAndReportsMap := domainPermissionsProfile["events-and-reports"].(map[string]interface{})

		eventsAndReportsMapToReturn := make(map[string]interface{})

		if v, _ := eventsAndReportsMap["smart-event"]; v != nil && isArgDefault(v.(string), d, "events_and_reports.smart_event", defaultEventsAndReports["smart-event"].(string)) {
			eventsAndReportsMapToReturn["smart_event"] = v
		}
		if v, _ := eventsAndReportsMap["events"]; v != nil && isArgDefault(v.(string), d, "events_and_reports.events", defaultEventsAndReports["events"].(string)) {
			eventsAndReportsMapToReturn["events"] = v
		}
		if v, _ := eventsAndReportsMap["policy"]; v != nil && isArgDefault(v.(string), d, "events_and_reports.policy", defaultEventsAndReports["policy"].(string)) {
			eventsAndReportsMapToReturn["policy"] = v
		}
		if v, _ := eventsAndReportsMap["reports"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "events_and_reports.reports", defaultEventsAndReports["reports"].(string)) {
			eventsAndReportsMapToReturn["reports"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("events_and_reports", eventsAndReportsMapToReturn)
	} else {
		_ = d.Set("events_and_reports", nil)
	}

	if domainPermissionsProfile["gateways"] != nil {

		defaultGateways := map[string]interface{}{
			"smart-update":                 "read",
			"lsm-gw-db":                    "disabled",
			"manage-provisioning-profiles": "disabled",
			"vsx-provisioning":             "false",
			"system-backup":                "false",
			"system-restore":               "false",
			"open-shell":                   "false",
			"run-one-time-script":          "false",
			"run-repository-script":        "false",
			"manage-repository-scripts":    "read",
		}

		gatewaysMap := domainPermissionsProfile["gateways"].(map[string]interface{})

		gatewaysMapToReturn := make(map[string]interface{})

		if v, _ := gatewaysMap["smart-update"]; v != nil && isArgDefault(v.(string), d, "gateways.smart_update", defaultGateways["smart-update"].(string)) {
			gatewaysMapToReturn["smart_update"] = v
		}
		if v, _ := gatewaysMap["lsm-gw-db"]; v != nil && isArgDefault(v.(string), d, "gateways.lsm_gw_db", defaultGateways["lsm-gw-db"].(string)) {
			gatewaysMapToReturn["lsm_gw_db"] = v
		}
		if v, _ := gatewaysMap["manage-provisioning-profiles"]; v != nil && isArgDefault(v.(string), d, "gateways.manage_provisioning_profiles", defaultGateways["manage-provisioning-profiles"].(string)) {
			gatewaysMapToReturn["manage_provisioning_profiles"] = v
		}
		if v, _ := gatewaysMap["vsx-provisioning"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "gateways.vsx_provisioning", defaultGateways["vsx-provisioning"].(string)) {
			gatewaysMapToReturn["vsx_provisioning"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["system-backup"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "gateways.system_backup", defaultGateways["system-backup"].(string)) {
			gatewaysMapToReturn["system_backup"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["system-restore"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "gateways.system_restore", defaultGateways["system-restore"].(string)) {
			gatewaysMapToReturn["system_restore"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["open-shell"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "gateways.open_shell", defaultGateways["open-shell"].(string)) {
			gatewaysMapToReturn["open_shell"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["run-one-time-script"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "gateways.run_one_time_script", defaultGateways["run-one-time-script"].(string)) {
			gatewaysMapToReturn["run_one_time_script"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["run-repository-script"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "gateways.run_repository_script", defaultGateways["run-repository-script"].(string)) {
			gatewaysMapToReturn["run_repository_script"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["manage-repository-scripts"]; v != nil && isArgDefault(v.(string), d, "gateways.manage_repository_scripts", defaultGateways["manage-repository-scripts"].(string)) {
			gatewaysMapToReturn["manage_repository_scripts"] = v
		}
		_ = d.Set("gateways", gatewaysMapToReturn)
	} else {
		_ = d.Set("gateways", nil)
	}

	if domainPermissionsProfile["management"] != nil {

		defaultManagement := map[string]interface{}{
			"manage-sessions":                        "false",
			"high-availability-operations":           "true",
			"approve-or-reject-sessions":             "false",
			"publish-sessions":                       "true",
			"manage-integration-with-cloud-services": "false",
		}

		managementMap := domainPermissionsProfile["management"].(map[string]interface{})

		managementMapToReturn := make(map[string]interface{})

		if v, _ := managementMap["cme-operations"]; v != nil {
			managementMapToReturn["cme_operations"] = v
		}
		if v, _ := managementMap["manage-admins"]; v != nil {
			managementMapToReturn["manage_admins"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["management-api-login"]; v != nil {
			managementMapToReturn["management_api_login"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["manage-sessions"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "management.manage_sessions", defaultManagement["manage-sessions"].(string)) {
			managementMapToReturn["manage_sessions"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["high-availability-operations"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "management.high_availability_operations", defaultManagement["high-availability-operations"].(string)) {
			managementMapToReturn["high_availability_operations"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["approve-or-reject-sessions"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "management.approve_or_reject_sessions", defaultManagement["approve-or-reject-sessions"].(string)) {
			managementMapToReturn["approve_or_reject_sessions"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["publish-sessions"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "management.publish_sessions", defaultManagement["publish-sessions"].(string)) {
			managementMapToReturn["publish_sessions"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["manage-integration-with-cloud-services"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "management.manage_integration_with_cloud_services", defaultManagement["manage-integration-with-cloud-services"].(string)) {
			managementMapToReturn["manage_integration_with_cloud_services"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("management", managementMapToReturn)
	} else {
		_ = d.Set("management", nil)
	}

	if domainPermissionsProfile["monitoring-and-logging"] != nil {

		defaultMonitoring := map[string]interface{}{
			"monitoring":                             "write",
			"management-logs":                        "write",
			"track-logs":                             "write",
			"app-and-url-filtering-logs":             "true",
			"https-inspection-logs":                  "true",
			"packet-capture-and-forensics":           "true",
			"show-packet-capture-by-default":         "true",
			"identities":                             "true",
			"show-identities-by-default":             "true",
			"dlp-logs-including-confidential-fields": "false",
			"manage-dlp-messages":                    "false",
		}

		monitoringAndLoggingMap := domainPermissionsProfile["monitoring-and-logging"].(map[string]interface{})

		monitoringAndLoggingMapToReturn := make(map[string]interface{})

		if v, _ := monitoringAndLoggingMap["monitoring"]; v != nil && isArgDefault(v.(string), d, "monitoring_and_logging.monitoring", defaultMonitoring["monitoring"].(string)) {
			monitoringAndLoggingMapToReturn["monitoring"] = v
		}
		if v, _ := monitoringAndLoggingMap["management-logs"]; v != nil && isArgDefault(v.(string), d, "monitoring_and_logging.management_logs", defaultMonitoring["management-logs"].(string)) {
			monitoringAndLoggingMapToReturn["management_logs"] = v
		}
		if v, _ := monitoringAndLoggingMap["track-logs"]; v != nil && isArgDefault(v.(string), d, "monitoring_and_logging.track_logs", defaultMonitoring["track-logs"].(string)) {
			monitoringAndLoggingMapToReturn["track_logs"] = v
		}
		if v, _ := monitoringAndLoggingMap["app-and-url-filtering-logs"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "monitoring_and_logging.app_and_url_filtering_logs", defaultMonitoring["app-and-url-filtering-logs"].(string)) {
			monitoringAndLoggingMapToReturn["app_and_url_filtering_logs"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["https-inspection-logs"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "monitoring_and_logging.https_inspection_logs", defaultMonitoring["https-inspection-logs"].(string)) {
			monitoringAndLoggingMapToReturn["https_inspection_logs"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["packet-capture-and-forensics"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "monitoring_and_logging.packet_capture_and_forensics", defaultMonitoring["packet-capture-and-forensics"].(string)) {
			monitoringAndLoggingMapToReturn["packet_capture_and_forensics"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["show-packet-capture-by-default"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "monitoring_and_logging.show_packet_capture_by_default", defaultMonitoring["show-packet-capture-by-default"].(string)) {
			monitoringAndLoggingMapToReturn["show_packet_capture_by_default"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["identities"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "monitoring_and_logging.identities", defaultMonitoring["identities"].(string)) {
			monitoringAndLoggingMapToReturn["identities"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["show-identities-by-default"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "monitoring_and_logging.show_identities_by_default", defaultMonitoring["show-identities-by-default"].(string)) {
			monitoringAndLoggingMapToReturn["show_identities_by_default"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["dlp-logs-including-confidential-fields"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "monitoring_and_logging.dlp_logs_including_confidential_fields", defaultMonitoring["dlp-logs-including-confidential-fields"].(string)) {
			monitoringAndLoggingMapToReturn["dlp_logs_including_confidential_fields"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["manage-dlp-messages"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "monitoring_and_logging.manage_dlp_messages", defaultMonitoring["manage-dlp-messages"].(string)) {
			monitoringAndLoggingMapToReturn["manage_dlp_messages"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("monitoring_and_logging", monitoringAndLoggingMapToReturn)
	} else {
		_ = d.Set("monitoring_and_logging", nil)
	}

	if domainPermissionsProfile["threat-prevention"] != nil {

		defaultThreatPrevention := map[string]interface{}{
			"policy-layers":     "write",
			"edit-layers":       "all",
			"edit-settings":     "true",
			"policy-exceptions": "write",
			"profiles":          "write",
			"protections":       "write",
			"install-policy":    "true",
			"ips-update":        "true",
		}

		threatPreventionMap := domainPermissionsProfile["threat-prevention"].(map[string]interface{})

		threatPreventionMapToReturn := make(map[string]interface{})

		if v, _ := threatPreventionMap["policy-layers"]; v != nil && isArgDefault(v.(string), d, "threat_prevention.policy_layers", defaultThreatPrevention["policy-layers"].(string)) {
			threatPreventionMapToReturn["policy_layers"] = v
		}
		if v, _ := threatPreventionMap["edit-layers"]; v != nil && isArgDefault(v.(string), d, "threat_prevention.edit_layers", defaultThreatPrevention["edit-layers"].(string)) {
			threatPreventionMapToReturn["edit_layers"] = v
		}
		if v, _ := threatPreventionMap["edit-settings"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "threat_prevention.edit_settings", defaultThreatPrevention["edit-settings"].(string)) {
			threatPreventionMapToReturn["edit_settings"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := threatPreventionMap["policy-exceptions"]; v != nil && isArgDefault(v.(string), d, "threat_prevention.policy_exceptions", defaultThreatPrevention["policy-exceptions"].(string)) {
			threatPreventionMapToReturn["policy_exceptions"] = v
		}
		if v, _ := threatPreventionMap["profiles"]; v != nil && isArgDefault(v.(string), d, "threat_prevention.profiles", defaultThreatPrevention["profiles"].(string)) {
			threatPreventionMapToReturn["profiles"] = v
		}
		if v, _ := threatPreventionMap["protections"]; v != nil && isArgDefault(v.(string), d, "threat_prevention.protections", defaultThreatPrevention["protections"].(string)) {
			threatPreventionMapToReturn["protections"] = v
		}
		if v, _ := threatPreventionMap["install-policy"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "threat_prevention.install_policy", defaultThreatPrevention["install-policy"].(string)) {
			threatPreventionMapToReturn["install_policy"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := threatPreventionMap["ips-update"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "threat_prevention.ips_update", defaultThreatPrevention["ips-update"].(string)) {
			threatPreventionMapToReturn["ips_update"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("threat_prevention", threatPreventionMapToReturn)
	} else {
		_ = d.Set("threat_prevention", nil)
	}

	if domainPermissionsProfile["others"] != nil {

		defaultOthers := map[string]interface{}{
			"client-certificates":   "true",
			"edit-cp-users-db":      "true",
			"https-inspection":      "write",
			"ldap-users-db":         "write",
			"user-authority-access": "write",
			"user-device-mgmt-conf": "read",
		}
		othersMap := domainPermissionsProfile["others"].(map[string]interface{})

		othersMapToReturn := make(map[string]interface{})

		if v, _ := othersMap["client-certificates"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "others.client_certificates", defaultOthers["client-certificates"].(string)) {
			othersMapToReturn["client_certificates"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := othersMap["edit-cp-users-db"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "others.edit_cp_users_db", defaultOthers["edit-cp-users-db"].(string)) {
			othersMapToReturn["edit_cp_users_db"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := othersMap["https-inspection"]; v != nil && isArgDefault(v.(string), d, "others.https_inspection", defaultOthers["https-inspection"].(string)) {
			othersMapToReturn["https_inspection"] = v
		}
		if v, _ := othersMap["ldap-users-db"]; v != nil && isArgDefault(v.(string), d, "others.ldap_users_db", defaultOthers["ldap-users-db"].(string)) {
			othersMapToReturn["ldap_users_db"] = v
		}
		if v, _ := othersMap["user-authority-access"]; v != nil && isArgDefault(v.(string), d, "others.user_authority_access", defaultOthers["user-authority-access"].(string)) {
			othersMapToReturn["user_authority_access"] = v
		}
		if v, _ := othersMap["user-device-mgmt-conf"]; v != nil && isArgDefault(v.(string), d, "others.user_device_mgmt_conf", defaultOthers["user-device-mgmt-conf"].(string)) {
			othersMapToReturn["user_device_mgmt_conf"] = v
		}

		_ = d.Set("others", othersMapToReturn)
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

		if _, ok := d.GetOk("endpoint"); ok {

			res := make(map[string]interface{})

			if d.HasChange("endpoint.manage_policies_and_software_deployment") {
				res["manage-policies-and-software-deployment"] = d.Get("endpoint.manage_policies_and_software_deployment")
			}
			if d.HasChange("endpoint.edit_endpoint_policies") {
				res["edit-endpoint-policies"] = d.Get("endpoint.edit_endpoint_policies")
			}
			if d.HasChange("endpoint.policies_installation") {
				res["policies-installation"] = d.Get("endpoint.policies_installation")
			}
			if d.HasChange("endpoint.edit_software_deployment") {
				res["edit-software-deployment"] = d.Get("endpoint.edit_software_deployment")
			}
			if d.HasChange("endpoint.software_deployment_installation") {
				res["software-deployment-installation"] = d.Get("endpoint.software_deployment_installation")
			}
			if d.HasChange("endpoint.allow_executing_push_operations") {
				res["allow-executing-push-operations"] = d.Get("endpoint.allow_executing_push_operations")
			}
			if d.HasChange("endpoint.authorize_preboot_users") {
				res["authorize-preboot-users"] = d.Get("endpoint.authorize_preboot_users")
			}
			if d.HasChange("endpoint.recovery_media") {
				res["recovery-media"] = d.Get("endpoint.recovery_media")
			}
			if d.HasChange("endpoint.remote_help") {
				res["remote-help"] = d.Get("endpoint.remote_help")
			}
			if d.HasChange("endpoint.reset_computer_data") {
				res["reset-computer-data"] = d.Get("endpoint.reset_computer_data")
			}
			domainPermissionsProfile["endpoint"] = res
		}
	}

	if d.HasChange("events_and_reports") {

		if _, ok := d.GetOk("events_and_reports"); ok {

			res := make(map[string]interface{})

			if d.HasChange("events_and_reports.smart_event") {
				res["smart-event"] = d.Get("events_and_reports.smart_event")
			}
			if d.HasChange("events_and_reports.events") {
				res["events"] = d.Get("events_and_reports.events")
			}
			if d.HasChange("events_and_reports.policy") {
				res["policy"] = d.Get("events_and_reports.policy")
			}
			if d.HasChange("events_and_reports.reports") {
				res["reports"] = d.Get("events_and_reports.reports")
			}
			domainPermissionsProfile["events-and-reports"] = res
		}
	}

	if d.HasChange("gateways") {

		if _, ok := d.GetOk("gateways"); ok {

			res := make(map[string]interface{})

			if d.HasChange("gateways.smart_update") {
				res["smart-update"] = d.Get("gateways.smart_update")
			}
			if d.HasChange("gateways.lsm_gw_db") {
				res["lsm-gw-db"] = d.Get("gateways.lsm_gw_db")
			}
			if d.HasChange("gateways.manage_provisioning_profiles") {
				res["manage-provisioning-profiles"] = d.Get("gateways.manage_provisioning_profiles")
			}
			if d.HasChange("gateways.vsx_provisioning") {
				res["vsx-provisioning"] = d.Get("gateways.vsx_provisioning")
			}
			if d.HasChange("gateways.system_backup") {
				res["system-backup"] = d.Get("gateways.system_backup")
			}
			if d.HasChange("gateways.system_restore") {
				res["system-restore"] = d.Get("gateways.system_restore")
			}
			if d.HasChange("gateways.open_shell") {
				res["open-shell"] = d.Get("gateways.open_shell")
			}
			if d.HasChange("gateways.run_one_time_script") {
				res["run-one-time-script"] = d.Get("gateways.run_one_time_script")
			}
			if d.HasChange("gateways.run_repository_script") {
				res["run-repository-script"] = d.Get("gateways.run_repository_script")
			}
			if d.HasChange("gateways.manage_repository_scripts") {
				res["manage-repository-scripts"] = d.Get("gateways.manage_repository_scripts")
			}
			domainPermissionsProfile["gateways"] = res
		}
	}

	if d.HasChange("management") {

		if _, ok := d.GetOk("management"); ok {

			res := make(map[string]interface{})

			if d.HasChange("management.cme_operations") {
				res["cme-operations"] = d.Get("management.cme_operations")
			}
			if d.HasChange("management.manage_admins") {
				res["manage-admins"] = d.Get("management.manage_admins")
			}
			if d.HasChange("management.management_api_login") {
				res["management-api-login"] = d.Get("management.management_api_login")
			}
			if d.HasChange("management.manage_sessions") {
				res["manage-sessions"] = d.Get("management.manage_sessions")
			}
			if d.HasChange("management.high_availability_operations") {
				res["high-availability-operations"] = d.Get("management.high_availability_operations")
			}
			if d.HasChange("management.approve_or_reject_sessions") {
				res["approve-or-reject-sessions"] = d.Get("management.approve_or_reject_sessions")
			}
			if d.HasChange("management.publish_sessions") {
				res["publish-sessions"] = d.Get("management.publish_sessions")
			}
			if d.HasChange("management.manage_integration_with_cloud_services") {
				res["manage-integration-with-cloud-services"] = d.Get("management.manage_integration_with_cloud_services")
			}
			domainPermissionsProfile["management"] = res
		}
	}

	if d.HasChange("monitoring_and_logging") {

		if _, ok := d.GetOk("monitoring_and_logging"); ok {

			res := make(map[string]interface{})

			if d.HasChange("monitoring_and_logging.monitoring") {
				res["monitoring"] = d.Get("monitoring_and_logging.monitoring")
			}
			if d.HasChange("monitoring_and_logging.management_logs") {
				res["management-logs"] = d.Get("monitoring_and_logging.management_logs")
			}
			if d.HasChange("monitoring_and_logging.track_logs") {
				res["track-logs"] = d.Get("monitoring_and_logging.track_logs")
			}
			if d.HasChange("monitoring_and_logging.app_and_url_filtering_logs") {
				res["app-and-url-filtering-logs"] = d.Get("monitoring_and_logging.app_and_url_filtering_logs")
			}
			if d.HasChange("monitoring_and_logging.https_inspection_logs") {
				res["https-inspection-logs"] = d.Get("monitoring_and_logging.https_inspection_logs")
			}
			if d.HasChange("monitoring_and_logging.packet_capture_and_forensics") {
				res["packet-capture-and-forensics"] = d.Get("monitoring_and_logging.packet_capture_and_forensics")
			}
			if d.HasChange("monitoring_and_logging.show_packet_capture_by_default") {
				res["show-packet-capture-by-default"] = d.Get("monitoring_and_logging.show_packet_capture_by_default")
			}
			if d.HasChange("monitoring_and_logging.identities") {
				res["identities"] = d.Get("monitoring_and_logging.identities")
			}
			if d.HasChange("monitoring_and_logging.show_identities_by_default") {
				res["show-identities-by-default"] = d.Get("monitoring_and_logging.show_identities_by_default")
			}
			if d.HasChange("monitoring_and_logging.dlp_logs_including_confidential_fields") {
				res["dlp-logs-including-confidential-fields"] = d.Get("monitoring_and_logging.dlp_logs_including_confidential_fields")
			}
			if d.HasChange("monitoring_and_logging.manage_dlp_messages") {
				res["manage-dlp-messages"] = d.Get("monitoring_and_logging.manage_dlp_messages")
			}
			domainPermissionsProfile["monitoring-and-logging"] = res
		}
	}

	if d.HasChange("threat_prevention") {

		if _, ok := d.GetOk("threat_prevention"); ok {

			res := make(map[string]interface{})

			if d.HasChange("threat_prevention.policy_layers") {
				res["policy-layers"] = d.Get("threat_prevention.policy_layers")
			}
			if d.HasChange("threat_prevention.edit_layers") {
				res["edit-layers"] = d.Get("threat_prevention.edit_layers")
			}
			if d.HasChange("threat_prevention.edit_settings") {
				res["edit-settings"] = d.Get("threat_prevention.edit_settings")
			}
			if d.HasChange("threat_prevention.policy_exceptions") {
				res["policy-exceptions"] = d.Get("threat_prevention.policy_exceptions")
			}
			if d.HasChange("threat_prevention.profiles") {
				res["profiles"] = d.Get("threat_prevention.profiles")
			}
			if d.HasChange("threat_prevention.protections") {
				res["protections"] = d.Get("threat_prevention.protections")
			}
			if d.HasChange("threat_prevention.install_policy") {
				res["install-policy"] = d.Get("threat_prevention.install_policy")
			}
			if d.HasChange("threat_prevention.ips_update") {
				res["ips-update"] = d.Get("threat_prevention.ips_update")
			}
			domainPermissionsProfile["threat-prevention"] = res
		}
	}

	if d.HasChange("others") {

		if _, ok := d.GetOk("others"); ok {

			res := make(map[string]interface{})

			if d.HasChange("others.client_certificates") {
				res["client-certificates"] = d.Get("others.client_certificates")
			}
			if d.HasChange("others.edit_cp_users_db") {
				res["edit-cp-users-db"] = d.Get("others.edit_cp_users_db")
			}
			if d.HasChange("others.https_inspection") {
				res["https-inspection"] = d.Get("others.https_inspection")
			}
			if d.HasChange("others.ldap_users_db") {
				res["ldap-users-db"] = d.Get("others.ldap_users_db")
			}
			if d.HasChange("others.user_authority_access") {
				res["user-authority-access"] = d.Get("others.user_authority_access")
			}
			if d.HasChange("others.user_device_mgmt_conf") {
				res["user-device-mgmt-conf"] = d.Get("others.user_device_mgmt_conf")
			}
			domainPermissionsProfile["others"] = res
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
			return fmt.Errorf(updateDomainPermissionsProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDomainPermissionsProfile(d, m)
}

func deleteManagementDomainPermissionsProfile(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	domainPermissionsProfilePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DomainPermissionsProfile")

	deleteDomainPermissionsProfileRes, err := client.ApiCall("delete-domain-permissions-profile", domainPermissionsProfilePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteDomainPermissionsProfileRes.Success {
		if deleteDomainPermissionsProfileRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDomainPermissionsProfileRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
