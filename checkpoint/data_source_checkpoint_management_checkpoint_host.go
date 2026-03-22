package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementCheckpointHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCheckpointHostRead,
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
			"interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Checkpoint host interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Interface name.",
						},
						"subnet4": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network address.",
						},
						"subnet6": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network address.",
						},
						"mask_length4": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPv4 network mask length.",
						},
						"mask_length6": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "IPv6 network mask length.",
						},
						"subnet_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network mask.",
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
				},
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 address.",
			},
			"nat_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to add automatic address translation rules.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address.",
						},
						"hide_behind": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"hardware": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Hardware name.",
			},
			"os": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operating system name.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Checkpoint host platform version.",
			},
			"sic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of the Secure Internal Connection Trust.",
			},
			"sic_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State the Secure Internal Connection Trust.",
			},
			"management_blades": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Management blades.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_policy_management": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Network Policy Management.",
						},
						"logging_and_status": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Logging & Status.",
						},
						"smart_event_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable SmartEvent server. </br>When activating SmartEvent server, blades 'logging-and-status' and 'smart-event-correlation' should be set to True. </br>To complete SmartEvent configuration, perform Install Database or Install Policy on your Security Management servers and Log servers. </br>Activating SmartEvent Server is not recommended in Management High Availability environment. For more information refer to sk25164.",
						},
						"smart_event_correlation": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable SmartEvent Correlation Unit.",
						},
						"endpoint_policy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Endpoint Policy. </br>To complete Endpoint Security Management configuration, perform Install Database on your Endpoint Management Server. </br>Field is not supported on Multi Domain Server environment.",
						},
						"compliance": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Compliance blade. Can be set when 'network-policy-management' was selected to be True.",
						},
						"user_directory": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable User Directory. Can be set when 'network-policy-management' was selected to be True.",
						},
						"secondary": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Secondary Management enabled.",
						},
						"identity_logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Identity Logging enabled.",
						},
					},
				},
			},
			"logs_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Logs settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"free_disk_space_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Free disk space metrics.",
						},
						"accept_syslog_messages": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable accept syslog messages.",
						},
						"alert_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable alert when free disk space is below threshold.",
						},
						"alert_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Alert when free disk space below threshold.",
						},
						"alert_when_free_disk_space_below_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert when free disk space below type.",
						},
						"before_delete_keep_logs_from_the_last_days": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable before delete keep logs from the last days.",
						},
						"before_delete_keep_logs_from_the_last_days_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Before delete keep logs from the last days threshold.",
						},
						"before_delete_run_script": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Before delete run script.",
						},
						"before_delete_run_script_command": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Before delete run script command.",
						},
						"delete_index_files_older_than_days": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete index files older than days.",
						},
						"delete_index_files_older_than_days_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete index files older than days threshold.",
						},
						"delete_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete when free disk space below.",
						},
						"delete_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete when free disk space below threshold.",
						},
						"detect_new_citrix_ica_application_names": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable detect new citrix ica application names.",
						},
						"enable_log_indexing": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable log indexing.",
						},
						"forward_logs_to_log_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable forward logs to log server.",
						},
						"forward_logs_to_log_server_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forward logs to log server name.",
						},
						"forward_logs_to_log_server_schedule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forward logs to log server schedule name.",
						},
						"rotate_log_by_file_size": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable rotate log by file size.",
						},
						"rotate_log_file_size_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Log file size threshold.",
						},
						"rotate_log_on_schedule": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable rotate log on schedule.",
						},
						"rotate_log_schedule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rotate log schedule name.",
						},
						"smart_event_intro_correletion_unit": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable SmartEvent intro correletion unit.",
						},
						"stop_logging_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable stop logging when free disk space below.",
						},
						"stop_logging_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Stop logging when free disk space below threshold.",
						},
						"turn_on_qos_logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable turn on qos logging.",
						},
						"update_account_log_every": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Update account log in every amount of seconds.",
						},
					},
				},
			},
			"save_logs_locally": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable save logs locally.",
			},
			"send_alerts_to_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Server(s) to send alerts to identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Backup server(s) to send logs to identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Server(s) to send logs to identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func dataSourceManagementCheckpointHostRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showCheckpointHostRes, err := client.ApiCall("show-checkpoint-host", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showCheckpointHostRes.Success {
		return fmt.Errorf(showCheckpointHostRes.ErrorMsg)
	}

	checkpointHost := showCheckpointHostRes.GetData()

	log.Println("Read CheckpointHost - Show JSON = ", checkpointHost)

	if v := checkpointHost["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := checkpointHost["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if checkpointHost["interfaces"] != nil {

		interfacesList, ok := checkpointHost["interfaces"].([]interface{})

		if ok {

			if len(interfacesList) > 0 {

				var interfacesListToReturn []map[string]interface{}

				for i := range interfacesList {

					interfacesMap := interfacesList[i].(map[string]interface{})

					interfacesMapToAdd := make(map[string]interface{})

					if v, _ := interfacesMap["name"]; v != nil {
						interfacesMapToAdd["name"] = v
					}
					if v, _ := interfacesMap["subnet4"]; v != nil {
						interfacesMapToAdd["subnet4"] = v
					}
					if v, _ := interfacesMap["subnet6"]; v != nil {
						interfacesMapToAdd["subnet6"] = v
					}
					if v, _ := interfacesMap["mask-length4"]; v != nil {
						interfacesMapToAdd["mask_length4"] = v
					}
					if v, _ := interfacesMap["mask-length6"]; v != nil {
						interfacesMapToAdd["mask_length6"] = v
					}
					if v, _ := interfacesMap["subnet-mask"]; v != nil {
						interfacesMapToAdd["subnet_mask"] = v
					}
					if v, _ := interfacesMap["color"]; v != nil {
						interfacesMapToAdd["color"] = v
					}
					if v, _ := interfacesMap["comments"]; v != nil {
						interfacesMapToAdd["comments"] = v
					}
					if v, _ := interfacesMap["ignore-warnings"]; v != nil {
						interfacesMapToAdd["ignore_warnings"] = v
					}
					if v, _ := interfacesMap["ignore-errors"]; v != nil {
						interfacesMapToAdd["ignore_errors"] = v
					}
					interfacesListToReturn = append(interfacesListToReturn, interfacesMapToAdd)
				}
				_ = d.Set("interfaces", interfacesListToReturn)
			}
		}
	}

	if v := checkpointHost["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := checkpointHost["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if checkpointHost["nat-settings"] != nil {

		natSettingsMap := checkpointHost["nat-settings"].(map[string]interface{})

		natSettingsMapToReturn := make(map[string]interface{})

		if v := natSettingsMap["auto-rule"]; v != nil {
			natSettingsMapToReturn["auto_rule"] = v
		}
		if v := natSettingsMap["ipv4-address"]; v != nil {
			natSettingsMapToReturn["ipv4_address"] = v
		}
		if v := natSettingsMap["ipv6-address"]; v != nil {
			natSettingsMapToReturn["ipv6_address"] = v
		}
		if v := natSettingsMap["hide-behind"]; v != nil {
			natSettingsMapToReturn["hide_behind"] = v
		}
		if v := natSettingsMap["install-on"]; v != nil {
			natSettingsMapToReturn["install_on"] = v
		}
		if v := natSettingsMap["method"]; v != nil {
			natSettingsMapToReturn["method"] = v
		}

		_ = d.Set("nat_settings", []interface{}{natSettingsMapToReturn})

	} else {
		_ = d.Set("nat_settings", nil)
	}

	if v := checkpointHost["one-time-password"]; v != nil {
		_ = d.Set("one_time_password", v)
	}

	if v := checkpointHost["hardware"]; v != nil {
		_ = d.Set("hardware", v)
	}

	if v := checkpointHost["os"]; v != nil {
		_ = d.Set("os", v)
	}

	if v := checkpointHost["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := checkpointHost["sic_name"]; v != nil {
		_ = d.Set("sic_name", v)
	}

	if v := checkpointHost["sic_state"]; v != nil {
		_ = d.Set("sic_state", v)
	}

	if checkpointHost["management-blades"] != nil {

		managementBladesMap := checkpointHost["management-blades"].(map[string]interface{})

		managementBladesMapToReturn := make(map[string]interface{})

		if v := managementBladesMap["network-policy-management"]; v != nil {
			managementBladesMapToReturn["network_policy_management"] = v
		}
		if v := managementBladesMap["secondary"]; v != nil {
			managementBladesMapToReturn["secondary"] = v
		}
		if v := managementBladesMap["logging-and-status"]; v != nil {
			managementBladesMapToReturn["logging_and_status"] = v
		}
		if v := managementBladesMap["endpoint-policy"]; v != nil {
			managementBladesMapToReturn["endpoint_policy"] = v
		}
		if v := managementBladesMap["identity-logging"]; v != nil {
			managementBladesMapToReturn["identity_logging"] = v
		}
		if v := managementBladesMap["smart-event-correlation"]; v != nil {
			managementBladesMapToReturn["smart_event_correlation"] = v
		}
		if v := managementBladesMap["smart-event-server"]; v != nil {
			managementBladesMapToReturn["smart_event_server"] = v
		}
		if v := managementBladesMap["compliance"]; v != nil {
			managementBladesMapToReturn["compliance"] = v
		}
		if v := managementBladesMap["user-directory"]; v != nil {
			managementBladesMapToReturn["user_directory"] = v
		}

		_ = d.Set("management_blades", []interface{}{managementBladesMapToReturn})

	} else {
		_ = d.Set("management_blades", nil)
	}

	if checkpointHost["logs-settings"] != nil {

		logsSettingsMap := checkpointHost["logs-settings"].(map[string]interface{})

		logsSettingsMapToReturn := make(map[string]interface{})

		if v := logsSettingsMap["accept-syslog-messages"]; v != nil {
			logsSettingsMapToReturn["accept_syslog_messages"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below-type"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_type"] = v
		}
		if v := logsSettingsMap["before-delete-keep-logs-from-the-last-days"]; v != nil {
			logsSettingsMapToReturn["before_delete_keep_logs_from_the_last_days"] = v
		}
		if v := logsSettingsMap["before-delete-keep-logs-from-the-last-days-threshold"]; v != nil {
			logsSettingsMapToReturn["before_delete_keep_logs_from_the_last_days_threshold"] = v
		}
		if v := logsSettingsMap["before-delete-run-script"]; v != nil {
			logsSettingsMapToReturn["before_delete_run_script"] = v
		}
		if v := logsSettingsMap["before-delete-run-script-command"]; v != nil {
			logsSettingsMapToReturn["before_delete_run_script_command"] = v
		}
		if v := logsSettingsMap["delete-index-files-older-than-days"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_older_than_days"] = v
		}
		if v := logsSettingsMap["delete-index-files-older-than-days-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_older_than_days_threshold"] = v
		}
		if v := logsSettingsMap["delete-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below"] = v
		}
		if v := logsSettingsMap["delete-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsMap["delete-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["detect-new-citrix-ica-application-names"]; v != nil {
			logsSettingsMapToReturn["detect_new_citrix_ica_application_names"] = v
		}
		if v := logsSettingsMap["distribute-logs-between-all-active-servers"]; v != nil {
			logsSettingsMapToReturn["distribute_logs_between_all_active_servers"] = v
		}
		if v := logsSettingsMap["enable-log-indexing"]; v != nil {
			logsSettingsMapToReturn["enable_log_indexing"] = v
		}
		if v := logsSettingsMap["forward-logs-to-log-server"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server"] = v
		}
		if v := logsSettingsMap["forward-logs-to-log-server-name"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server_name"] = v
		}
		if v := logsSettingsMap["forward-logs-to-log-server-schedule-name"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server_schedule_name"] = v
		}
		if v := logsSettingsMap["rotate-log-by-file-size"]; v != nil {
			logsSettingsMapToReturn["rotate_log_by_file_size"] = v
		}
		if v := logsSettingsMap["rotate-log-file-size-threshold"]; v != nil {
			logsSettingsMapToReturn["rotate_log_file_size_threshold"] = v
		}
		if v := logsSettingsMap["rotate-log-on-schedule"]; v != nil {
			logsSettingsMapToReturn["rotate_log_on_schedule"] = v
		}
		if v := logsSettingsMap["rotate-log-schedule-name"]; v != nil {
			logsSettingsMapToReturn["rotate_log_schedule_name"] = v
		}
		if v := logsSettingsMap["smart-event-intro-correlation-unit"]; v != nil {
			logsSettingsMapToReturn["smart_event_intro_correlation_unit"] = v
		}
		if v := logsSettingsMap["stop-logging-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below"] = v
		}
		if v := logsSettingsMap["stop-logging-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsMap["stop-logging-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["turn-on-qos-logging"]; v != nil {
			logsSettingsMapToReturn["turn_on_qos_logging"] = v
		}
		if v := logsSettingsMap["update-account-log-every"]; v != nil {
			logsSettingsMapToReturn["update_account_log_every"] = v
		}

		_ = d.Set("logs_settings", []interface{}{logsSettingsMapToReturn})

	} else {
		_ = d.Set("logs_settings", nil)
	}

	if v := checkpointHost["save-logs-locally"]; v != nil {
		_ = d.Set("save_logs_locally", v)
	}

	if checkpointHost["send_alerts_to_server"] != nil {
		sendAlertsToServerJson, ok := checkpointHost["send_alerts_to_server"].([]interface{})
		if ok {
			sendAlertsToServerIds := make([]string, 0)
			if len(sendAlertsToServerJson) > 0 {
				for _, sendAlertsToServer := range sendAlertsToServerJson {
					sendAlertsToServer := sendAlertsToServer.(map[string]interface{})
					sendAlertsToServerIds = append(sendAlertsToServerIds, sendAlertsToServer["name"].(string))
				}
			}
			_ = d.Set("send_alerts_to_server", sendAlertsToServerIds)
		}
	} else {
		_ = d.Set("send_alerts_to_server", nil)
	}

	if checkpointHost["send_logs_to_backup_server"] != nil {
		sendLogsToBackupServerJson, ok := checkpointHost["send_logs_to_backup_server"].([]interface{})
		if ok {
			sendLogsToBackupServerIds := make([]string, 0)
			if len(sendLogsToBackupServerJson) > 0 {
				for _, sendLogsToBackupServer := range sendLogsToBackupServerJson {
					sendLogsToBackupServer := sendLogsToBackupServer.(map[string]interface{})
					sendLogsToBackupServerIds = append(sendLogsToBackupServerIds, sendLogsToBackupServer["name"].(string))
				}
			}
			_ = d.Set("send_logs_to_backup_server", sendLogsToBackupServerIds)
		}
	} else {
		_ = d.Set("send_logs_to_backup_server", nil)
	}

	if checkpointHost["send_logs_to_server"] != nil {
		sendLogsToServerJson, ok := checkpointHost["send_logs_to_server"].([]interface{})
		if ok {
			sendLogsToServerIds := make([]string, 0)
			if len(sendLogsToServerJson) > 0 {
				for _, sendLogsToServer := range sendLogsToServerJson {
					sendLogsToServer := sendLogsToServer.(map[string]interface{})
					sendLogsToServerIds = append(sendLogsToServerIds, sendLogsToServer["name"].(string))
				}
			}
			_ = d.Set("send_logs_to_server", sendLogsToServerIds)
		}
	} else {
		_ = d.Set("send_logs_to_server", nil)
	}

	if checkpointHost["tags"] != nil {
		tagsJson, ok := checkpointHost["tags"].([]interface{})
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

	if v := checkpointHost["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := checkpointHost["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
