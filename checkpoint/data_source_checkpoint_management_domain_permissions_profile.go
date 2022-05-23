package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementDomainPermissionsProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDomainPermissionsProfileRead,
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
			"permission_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the Permissions Profile.",
			},
			"edit_common_objects": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Define and manage objects in the Check Point database: Network Objects, Services, Custom Application Site, VPN Community, Users, Servers, Resources, Time, UserCheck, and Limit.<br>Only a 'Customized' permission-type profile can edit this permission.",
			},
			"access_control": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Access Control permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"show_policy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select to let administrators work with Access Control rules and NAT rules. If not selected, administrators cannot see these rules.",
						},
						"policy_layers": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Layer editing permissions.<br>Available only if show-policy is set to true.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"edit_layers": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "\"By Software Blades\" - Edit Access Control layers that contain the blades enabled in the Permissions Profile.<br>\"By Selected Profile In A Layer Editor\" - Administrators can only edit the layer if the Access Control layer editor gives editing permission to their profiles.",
									},
									"app_control_and_url_filtering": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Use Application and URL Filtering in Access Control rules.<br>Available only if edit-layers is set to \"By Software Blades\".",
									},
									"content_awareness": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Use specified data types in Access Control rules.<br>Available only if edit-layers is set to \"By Software Blades\".",
									},
									"firewall": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Work with Access Control and other Software Blades that do not have their own Policies.<br>Available only if edit-layers is set to \"By Software Blades\".",
									},
									"mobile_access": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Work with Mobile Access rules.<br>Available only if edit-layers is set to \"By Software Blades\".",
									},
								},
							},
						},
						"dlp_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configure DLP rules and Policies.",
						},
						"geo_control_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Work with Access Control rules that control traffic to and from specified countries.",
						},
						"nat_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Work with NAT in Access Control rules.",
						},
						"qos_policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Work with QoS Policies and rules.",
						},
						"access_control_objects_and_settings": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Allow editing of the following objet types: VPN Community, Access Role, Custom application group,Custom application, Custom category, Limit, Application - Match Settings, Application Category - Match Settings,Override Categorization, Application and URL filtering blade - Advanced Settings, Content Awareness blade - Advanced Settings.",
						},
						"app_control_and_url_filtering_update": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Install Application and URL Filtering updates.",
						},
						"install_policy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Install Access Control Policies.",
						},
					},
				},
			},
			"endpoint": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Endpoint permissions. Not supported for Multi-Domain Servers.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"manage_policies_and_software_deployment": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The administrator can work with policies, rules and actions.",
						},
						"edit_endpoint_policies": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Available only if manage-policies-and-software-deployment is set to true.",
						},
						"policies_installation": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The administrator can install policies on endpoint computers.",
						},
						"edit_software_deployment": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The administrator can define deployment rules, create packages for export, and configure advanced package settings.<br>Available only if manage-policies-and-software-deployment is set to true.",
						},
						"software_deployment_installation": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The administrator can deploy packages and install endpoint clients.",
						},
						"allow_executing_push_operations": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The administrator can start operations that the Security Management Server pushes directly to client computers with no policy installation required.",
						},
						"authorize_preboot_users": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The administrator can add and remove the users who are permitted to log on to Endpoint Security client computers with Full Disk Encryption.",
						},
						"recovery_media": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The administrator can create recovery media on endpoint computers and devices.",
						},
						"remote_help": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The administrator can use the Remote Help feature to reset user passwords and give access to locked out users.",
						},
						"reset_computer_data": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The administrator can reset a computer, which deletes all information about the computer from the Security Management Server.",
						},
					},
				},
			},
			"events_and_reports": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Events and Reports permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"smart_event": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "'Custom' - Configure SmartEvent permissions.",
						},
						"events": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Work with event queries on the Events tab. Create custom event queries.<br>Available only if smart-event is set to 'Custom'.",
						},
						"policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configure SmartEvent Policy rules and install SmartEvent Policies.<br>Available only if smart-event is set to 'Custom'.",
						},
						"reports": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Create and run SmartEvent reports.<br>Available only if smart-event is set to 'Custom'.",
						},
					},
				},
			},
			"gateways": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Gateways permissions. <br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"smart_update": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Install, update and delete Check Point licenses. This includes permissions to use SmartUpdate to manage licenses.",
						},
						"lsm_gw_db": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access to objects defined in LSM gateway tables. These objects are managed in the SmartProvisioning GUI or LSMcli command-line.<br>Note: 'Write' permission on lsm-gw-db allows administrator to run a script on SmartLSM gateway in Expert mode.",
						},
						"manage_provisioning_profiles": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Administrator can add, edit, delete, and assign provisioning profiles to gateways (both LSM and non-LSM).<br>Available for edit only if lsm-gw-db is set with 'Write' permission.<br>Note: 'Read' permission on lsm-gw-db enables 'Read' permission for manage-provisioning-profiles.",
						},
						"vsx_provisioning": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Create and configure Virtual Systems and other VSX virtual objects.",
						},
						"system_backup": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Backup Security Gateways.",
						},
						"system_restore": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Restore Security Gateways from saved backups.",
						},
						"open_shell": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use the SmartConsole CLI to run commands.",
						},
						"run_one_time_script": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Run user scripts from the command line.",
						},
						"run_repository_script": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Run scripts from the repository.",
						},
						"manage_repository_scripts": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Add, change and remove scripts in the repository.",
						},
					},
				},
			},
			"management": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Management permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cme_operations": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Permission to read / edit the Cloud Management Extension (CME) configuration.<br>Not supported for Multi-Domain Servers.",
						},
						"manage_admins": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Controls the ability to manage Administrators, Permission Profiles, Trusted clients,API settings and Policy settings.<br>Only a \"Read Write All\" permission-type profile can edit this permission.<br>Not supported for Multi-Domain Servers.",
						},
						"management_api_login": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Permission to log in to the Security Management Server and run API commands using thesetools: mgmt_cli (Linux and Windows binaries), Gaia CLI (clish) and Web Services (REST). Useful if you want to prevent administrators from running automatic scripts on the Management.<br>Note: This permission is not required to run commands from within the API terminal in SmartConsole.<br>Not supported for Multi-Domain Servers.",
						},
						"manage_sessions": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Lets you disconnect, discard, publish, or take over other administrator sessions.<br>Only a \"Read Write All\" permission-type profile can edit this permission.",
						},
						"high_availability_operations": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Configure and work with Domain High Availability.<br>Only a 'Customized' permission-type profile can edit this permission.",
						},
						"approve_or_reject_sessions": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Approve / reject other sessions.",
						},
						"publish_sessions": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Allow session publishing without an approval.",
						},
						"manage_integration_with_cloud_services": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Manage integration with Cloud Services.",
						},
					},
				},
			},
			"monitoring_and_logging": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Monitoring and Logging permissions.<br>'Customized' permission-type profile can edit all these permissions. \"Read Write All\" permission-type can edit only dlp-logs-including-confidential-fields and manage-dlp-messages permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"monitoring": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "See monitoring views and reports.",
						},
						"management_logs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "See Multi-Domain Server audit logs.",
						},
						"track_logs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Use the log tracking features in SmartConsole.",
						},
						"app_and_url_filtering_logs": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Work with Application and URL Filtering logs.",
						},
						"https_inspection_logs": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "See logs generated by HTTPS Inspection.",
						},
						"packet_capture_and_forensics": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "See logs generated by the IPS and Forensics features.",
						},
						"show_packet_capture_by_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable packet capture by default.",
						},
						"identities": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Show user and computer identity information in logs.",
						},
						"show_identities_by_default": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Show user and computer identity information in logs by default.",
						},
						"dlp_logs_including_confidential_fields": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Show DLP logs including confidential fields.",
						},
						"manage_dlp_messages": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "View/Release/Discard DLP messages.<br>Available only if dlp-logs-including-confidential-fields is set to true.",
						},
					},
				},
			},
			"threat_prevention": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Threat Prevention permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_layers": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configure Threat Prevention Policy rules.<br>Note: To have policy-layers permissions you must set policy-exceptionsand profiles permissions. To have 'Write' permissions for policy-layers, policy-exceptions must be set with 'Write' permission as well.",
						},
						"edit_layers": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "'ALL' -  Gives permission to edit all layers.<br>\"By Selected Profile In A Layer Editor\" -  Administrators can only edit the layer if the Threat Prevention layer editor gives editing permission to their profiles.<br>Available only if policy-layers is set to 'Write'.",
						},
						"edit_settings": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Work with general Threat Prevention settings.",
						},
						"policy_exceptions": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configure exceptions to Threat Prevention rules.<br>Note: To have policy-exceptions you must set the protections permission.",
						},
						"profiles": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configure Threat Prevention profiles.",
						},
						"protections": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Work with malware protections.",
						},
						"install_policy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Install Policies.",
						},
						"ips_update": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Update IPS protections.<br>Note: You do not have to log into the User Center to receive IPS updates.",
						},
					},
				},
			},
			"others": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Additional permissions.<br>Only a 'Customized' permission-type profile can edit these permissions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"client_certificates": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Create and manage client certificates for Mobile Access.",
						},
						"edit_cp_users_db": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Work with user accounts and groups.",
						},
						"https_inspection": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Enable and configure HTTPS Inspection rules.",
						},
						"ldap_users_db": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Work with the LDAP database and user accounts, groups and OUs.",
						},
						"user_authority_access": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Work with Check Point User Authority authentication.",
						},
						"user_device_mgmt_conf": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gives access to the UDM (User & Device Management) web-based application that handles security challenges in a \"bring your own device\" (BYOD) workspace.",
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

func dataSourceManagementDomainPermissionsProfileRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := domainPermissionsProfile["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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
				accessControlMapToReturn["show_policy"] = v
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
			if v := accessControlMap["dlp-policy"]; v != nil {
				accessControlMapToReturn["dlp_policy"] = v
			}
			if v := accessControlMap["geo-control-policy"]; v != nil {
				accessControlMapToReturn["geo_control_policy"] = v
			}
			if v := accessControlMap["nat-policy"]; v != nil {
				accessControlMapToReturn["nat_policy"] = v
			}
			if v := accessControlMap["qos-policy"]; v != nil {
				accessControlMapToReturn["qos_policy"] = v
			}
			if v := accessControlMap["access-control-objects-and-settings"]; v != nil {
				accessControlMapToReturn["access_control_objects_and_settings"] = v
			}
			if v := accessControlMap["app-control-and-url-filtering-update"]; v != nil {
				accessControlMapToReturn["app_control_and_url_filtering_update"] = v
			}
			if v := accessControlMap["install-policy"]; v != nil {
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

		eventsAndReportsMap := domainPermissionsProfile["events-and-reports"].(map[string]interface{})

		eventsAndReportsMapToReturn := make(map[string]interface{})

		if v, _ := eventsAndReportsMap["smart-event"]; v != nil {
			eventsAndReportsMapToReturn["smart_event"] = v
		}
		if v, _ := eventsAndReportsMap["events"]; v != nil {
			eventsAndReportsMapToReturn["events"] = v
		}
		if v, _ := eventsAndReportsMap["policy"]; v != nil {
			eventsAndReportsMapToReturn["policy"] = v
		}
		if v, _ := eventsAndReportsMap["reports"]; v != nil {
			eventsAndReportsMapToReturn["reports"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("events_and_reports", eventsAndReportsMapToReturn)
	} else {
		_ = d.Set("events_and_reports", nil)
	}

	if domainPermissionsProfile["gateways"] != nil {

		gatewaysMap := domainPermissionsProfile["gateways"].(map[string]interface{})

		gatewaysMapToReturn := make(map[string]interface{})

		if v, _ := gatewaysMap["smart-update"]; v != nil {
			gatewaysMapToReturn["smart_update"] = v
		}
		if v, _ := gatewaysMap["lsm-gw-db"]; v != nil {
			gatewaysMapToReturn["lsm_gw_db"] = v
		}
		if v, _ := gatewaysMap["manage-provisioning-profiles"]; v != nil {
			gatewaysMapToReturn["manage_provisioning_profiles"] = v
		}
		if v, _ := gatewaysMap["vsx-provisioning"]; v != nil {
			gatewaysMapToReturn["vsx_provisioning"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["system-backup"]; v != nil {
			gatewaysMapToReturn["system_backup"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["system-restore"]; v != nil {
			gatewaysMapToReturn["system_restore"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["open-shell"]; v != nil {
			gatewaysMapToReturn["open_shell"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["run-one-time-script"]; v != nil {
			gatewaysMapToReturn["run_one_time_script"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["run-repository-script"]; v != nil {
			gatewaysMapToReturn["run_repository_script"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := gatewaysMap["manage-repository-scripts"]; v != nil {
			gatewaysMapToReturn["manage_repository_scripts"] = v
		}
		_ = d.Set("gateways", gatewaysMapToReturn)
	} else {
		_ = d.Set("gateways", nil)
	}

	if domainPermissionsProfile["management"] != nil {

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
		if v, _ := managementMap["manage-sessions"]; v != nil {
			managementMapToReturn["manage_sessions"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["high-availability-operations"]; v != nil {
			managementMapToReturn["high_availability_operations"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["approve-or-reject-sessions"]; v != nil {
			managementMapToReturn["approve_or_reject_sessions"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["publish-sessions"]; v != nil {
			managementMapToReturn["publish_sessions"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := managementMap["manage-integration-with-cloud-services"]; v != nil {
			managementMapToReturn["manage_integration_with_cloud_services"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("management", managementMapToReturn)
	} else {
		_ = d.Set("management", nil)
	}

	if domainPermissionsProfile["monitoring-and-logging"] != nil {

		monitoringAndLoggingMap := domainPermissionsProfile["monitoring-and-logging"].(map[string]interface{})

		monitoringAndLoggingMapToReturn := make(map[string]interface{})

		if v, _ := monitoringAndLoggingMap["monitoring"]; v != nil {
			monitoringAndLoggingMapToReturn["monitoring"] = v
		}
		if v, _ := monitoringAndLoggingMap["management-logs"]; v != nil {
			monitoringAndLoggingMapToReturn["management_logs"] = v
		}
		if v, _ := monitoringAndLoggingMap["track-logs"]; v != nil {
			monitoringAndLoggingMapToReturn["track_logs"] = v
		}
		if v, _ := monitoringAndLoggingMap["app-and-url-filtering-logs"]; v != nil {
			monitoringAndLoggingMapToReturn["app_and_url_filtering_logs"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["https-inspection-logs"]; v != nil {
			monitoringAndLoggingMapToReturn["https_inspection_logs"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["packet-capture-and-forensics"]; v != nil {
			monitoringAndLoggingMapToReturn["packet_capture_and_forensics"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["show-packet-capture-by-default"]; v != nil {
			monitoringAndLoggingMapToReturn["show_packet_capture_by_default"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["identities"]; v != nil {
			monitoringAndLoggingMapToReturn["identities"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["show-identities-by-default"]; v != nil {
			monitoringAndLoggingMapToReturn["show_identities_by_default"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["dlp-logs-including-confidential-fields"]; v != nil {
			monitoringAndLoggingMapToReturn["dlp_logs_including_confidential_fields"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := monitoringAndLoggingMap["manage-dlp-messages"]; v != nil {
			monitoringAndLoggingMapToReturn["manage_dlp_messages"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("monitoring_and_logging", monitoringAndLoggingMapToReturn)
	} else {
		_ = d.Set("monitoring_and_logging", nil)
	}

	if domainPermissionsProfile["threat-prevention"] != nil {

		threatPreventionMap := domainPermissionsProfile["threat-prevention"].(map[string]interface{})

		threatPreventionMapToReturn := make(map[string]interface{})

		if v, _ := threatPreventionMap["policy-layers"]; v != nil {
			threatPreventionMapToReturn["policy_layers"] = v
		}
		if v, _ := threatPreventionMap["edit-layers"]; v != nil {
			threatPreventionMapToReturn["edit_layers"] = v
		}
		if v, _ := threatPreventionMap["edit-settings"]; v != nil {
			threatPreventionMapToReturn["edit_settings"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := threatPreventionMap["policy-exceptions"]; v != nil {
			threatPreventionMapToReturn["policy_exceptions"] = v
		}
		if v, _ := threatPreventionMap["profiles"]; v != nil {
			threatPreventionMapToReturn["profiles"] = v
		}
		if v, _ := threatPreventionMap["protections"]; v != nil {
			threatPreventionMapToReturn["protections"] = v
		}
		if v, _ := threatPreventionMap["install-policy"]; v != nil {
			threatPreventionMapToReturn["install_policy"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := threatPreventionMap["ips-update"]; v != nil {
			threatPreventionMapToReturn["ips_update"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("threat_prevention", threatPreventionMapToReturn)
	} else {
		_ = d.Set("threat_prevention", nil)
	}

	if domainPermissionsProfile["others"] != nil {

		othersMap := domainPermissionsProfile["others"].(map[string]interface{})

		othersMapToReturn := make(map[string]interface{})

		if v, _ := othersMap["client-certificates"]; v != nil {
			othersMapToReturn["client_certificates"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := othersMap["edit-cp-users-db"]; v != nil {
			othersMapToReturn["edit_cp_users_db"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := othersMap["https-inspection"]; v != nil {
			othersMapToReturn["https_inspection"] = v
		}
		if v, _ := othersMap["ldap-users-db"]; v != nil {
			othersMapToReturn["ldap_users_db"] = v
		}
		if v, _ := othersMap["user-authority-access"]; v != nil {
			othersMapToReturn["user_authority_access"] = v
		}
		if v, _ := othersMap["user-device-mgmt-conf"]; v != nil {
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

	return nil

}
