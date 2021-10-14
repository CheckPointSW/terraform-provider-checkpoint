package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementCheckpointHost() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCheckpointHost,
		Read:   readManagementCheckpointHost,
		Update: updateManagementCheckpointHost,
		Delete: deleteManagementCheckpointHost,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Checkpoint host interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Interface name.",
						},
						"subnet4": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 network address.",
						},
						"subnet6": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 network address.",
						},
						"mask_length4": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "IPv4 network mask length.",
						},
						"mask_length6": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "IPv6 network mask length.",
						},
						"subnet_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 network mask.",
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
				},
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv6 address.",
			},
			"nat_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to add automatic address translation rules.",
							Default:     false,
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 address.",
						},
						"hide_behind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Hide behind method. This parameter is not required in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"one_time_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Secure internal connection one time password.",
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
			"hardware": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Hardware name.",
				Default:     "Open server",
			},
			"os": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Operating system name.",
				Default:     "Gaia",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Checkpoint host platform version.",
				Default:     "R81",
			},
			"management_blades": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Management blades.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"network_policy_management": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Network Policy Management.",
							Default:     false,
						},
						"logging_and_status": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Logging & Status.",
							Default:     false,
						},
						"smart_event_server": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable SmartEvent server. </br>When activating SmartEvent server, blades 'logging-and-status' and 'smart-event-correlation' should be set to True. </br>To complete SmartEvent configuration, perform Install Database or Install Policy on your Security Management servers and Log servers. </br>Activating SmartEvent Server is not recommended in Management High Availability environment. For more information refer to sk25164.",
							Default:     false,
						},
						"smart_event_correlation": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable SmartEvent Correlation Unit.",
							Default:     false,
						},
						"endpoint_policy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Endpoint Policy. </br>To complete Endpoint Security Management configuration, perform Install Database on your Endpoint Management Server. </br>Field is not supported on Multi Domain Server environment.",
							Default:     false,
						},
						"compliance": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Compliance blade. Can be set when 'network-policy-management' was selected to be True.",
							Default:     false,
						},
						"user_directory": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable User Directory. Can be set when 'network-policy-management' was selected to be True.",
							Default:     false,
						},
						"secondary": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Secondary Management enabled.",
						},
						"identity-logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Identity Logging enabled.",
						},
					},
				},
			},
			"logs_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Logs settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"free_disk_space_metrics": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free disk space metrics.",
							Default:     "mbytes",
						},
						"accept_syslog_messages": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable accept syslog messages.",
							Default:     false,
						},
						"alert_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable alert when free disk space is below threshold.",
							Default:     true,
						},
						"alert_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Alert when free disk space below threshold.",
							Default:     20,
						},
						"alert_when_free_disk_space_below_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Alert when free disk space below type.",
							Default:     "popup alert",
						},
						"before_delete_keep_logs_from_the_last_days": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable before delete keep logs from the last days.",
							Default:     false,
						},
						"before_delete_keep_logs_from_the_last_days_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Before delete keep logs from the last days threshold.",
							Default:     3650,
						},
						"before_delete_run_script": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Before delete run script.",
							Default:     false,
						},
						"before_delete_run_script_command": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Before delete run script command.",
						},
						"delete_index_files_older_than_days": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable delete index files older than days.",
						},
						"delete_index_files_older_than_days_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete index files older than days threshold.",
							Default:     14,
						},
						"delete_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable delete when free disk space below.",
							Default:     true,
						},
						"delete_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete when free disk space below threshold.",
							Default:     5000,
						},
						"detect_new_citrix_ica_application_names": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable detect new citrix ica application names.",
							Default:     false,
						},
						"enable_log_indexing": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable log indexing.",
							Default:     true,
						},
						"forward_logs_to_log_server": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable forward logs to log server.",
							Default:     false,
						},
						"forward_logs_to_log_server_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Forward logs to log server name.",
						},
						"forward_logs_to_log_server_schedule_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Forward logs to log server schedule name.",
						},
						"rotate_log_by_file_size": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable rotate log by file size.",
						},
						"rotate_log_file_size_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Log file size threshold.",
						},
						"rotate_log_on_schedule": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable rotate log on schedule.",
							Default:     false,
						},
						"rotate_log_schedule_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Rotate log schedule name.",
						},
						"smart_event_intro_correletion_unit": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable SmartEvent intro correletion unit.",
						},
						"stop_logging_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable stop logging when free disk space below.",
							Default:     false,
						},
						"stop_logging_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Stop logging when free disk space below threshold.",
							Default:     100,
						},
						"turn_on_qos_logging": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable turn on qos logging.",
							Default:     true,
						},
						"update_account_log_every": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Update account log in every amount of seconds.",
							Default:     3600,
						},
					},
				},
			},
			"save_logs_locally": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable save logs locally.",
			},
			"send_alerts_to_server": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Server(s) to send alerts to identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Backup server(s) to send logs to identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_server": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Server(s) to send logs to identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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

func createManagementCheckpointHost(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	checkpointHost := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		checkpointHost["name"] = v.(string)
	}

	if v, ok := d.GetOk("interfaces"); ok {

		interfacesList := v.([]interface{})

		if len(interfacesList) > 0 {

			var interfacesPayload []map[string]interface{}

			for i := range interfacesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".name"); ok {
					Payload["name"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".subnet4"); ok {
					Payload["subnet4"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".subnet6"); ok {
					Payload["subnet6"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".mask_length4"); ok {
					Payload["mask-length4"] = v.(int)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".mask_length6"); ok {
					Payload["mask-length6"] = v.(int)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".subnet_mask"); ok {
					Payload["subnet-mask"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".color"); ok {
					Payload["color"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".comments"); ok {
					Payload["comments"] = v.(string)
				}
				interfacesPayload = append(interfacesPayload, Payload)
			}
			checkpointHost["interfaces"] = interfacesPayload
		}
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		checkpointHost["ipv4-address"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address"); ok {
		checkpointHost["ipv6-address"] = v.(string)
	}

	if _, ok := d.GetOk("nat_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("nat_settings.auto_rule"); ok {
			res["auto-rule"] = v.(bool)
		}
		if v, ok := d.GetOk("nat_settings.ipv4_address"); ok {
			res["ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.ipv6_address"); ok {
			res["ipv6-address"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.hide_behind"); ok {
			res["hide-behind"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.install_on"); ok {
			res["install-on"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.method"); ok {
			res["method"] = v.(string)
		}
		checkpointHost["nat-settings"] = res
	}

	if v, ok := d.GetOk("one_time_password"); ok {
		checkpointHost["one-time-password"] = v.(string)
	}

	if v, ok := d.GetOk("hardware"); ok {
		checkpointHost["hardware"] = v.(string)
	}

	if v, ok := d.GetOk("os"); ok {
		checkpointHost["os"] = v.(string)
	}

	if v, ok := d.GetOk("version"); ok {
		checkpointHost["version"] = v.(string)
	}

	if _, ok := d.GetOk("management_blades"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOkExists("management_blades.network_policy_management"); ok {
			res["network-policy-management"] = v
		}
		if v, ok := d.GetOkExists("management_blades.logging_and_status"); ok {
			res["logging-and-status"] = v
		}
		if v, ok := d.GetOkExists("management_blades.smart_event_server"); ok {
			res["smart-event-server"] = v
		}
		if v, ok := d.GetOkExists("management_blades.smart_event_correlation"); ok {
			res["smart-event-correlation"] = v
		}
		if v, ok := d.GetOkExists("management_blades.endpoint_policy"); ok {
			res["endpoint-policy"] = v
		}
		if v, ok := d.GetOkExists("management_blades.compliance"); ok {
			res["compliance"] = v
		}
		if v, ok := d.GetOkExists("management_blades.user_directory"); ok {
			res["user-directory"] = v
		}
		checkpointHost["management-blades"] = res
	}

	if _, ok := d.GetOk("logs_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("logs_settings.free_disk_space_metrics"); ok {
			res["free-disk-space-metrics"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.accept_syslog_messages"); ok {
			res["accept-syslog-messages"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.alert_when_free_disk_space_below"); ok {
			res["alert-when-free-disk-space-below"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.alert_when_free_disk_space_below_threshold"); ok {
			res["alert-when-free-disk-space-below-threshold"] = v.(int)
		}
		if v, ok := d.GetOk("logs_settings.alert_when_free_disk_space_below_type"); ok {
			res["alert-when-free-disk-space-below-type"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.before_delete_keep_logs_from_the_last_days"); ok {
			res["before-delete-keep-logs-from-the-last-days"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.before_delete_keep_logs_from_the_last_days_threshold"); ok {
			res["before-delete-keep-logs-from-the-last-days-threshold"] = v.(int)
		}
		if v, ok := d.GetOk("logs_settings.before_delete_run_script"); ok {
			res["before-delete-run-script"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.before_delete_run_script_command"); ok {
			res["before-delete-run-script-command"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.delete_index_files_older_than_days"); ok {
			res["delete-index-files-older-than-days"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.delete_index_files_older_than_days_threshold"); ok {
			res["delete-index-files-older-than-days-threshold"] = v.(int)
		}
		if v, ok := d.GetOk("logs_settings.delete_when_free_disk_space_below"); ok {
			res["delete-when-free-disk-space-below"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.delete_when_free_disk_space_below_threshold"); ok {
			res["delete-when-free-disk-space-below-threshold"] = v
		}
		if v, ok := d.GetOk("logs_settings.detect_new_citrix_ica_application_names"); ok {
			res["detect-new-citrix-ica-application-names"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.enable_log_indexing"); ok {
			res["enable-log-indexing"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.forward_logs_to_log_server"); ok {
			res["forward-logs-to-log-server"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.forward_logs_to_log_server_name"); ok {
			res["forward-logs-to-log-server-name"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.forward_logs_to_log_server_schedule_name"); ok {
			res["forward-logs-to-log-server-schedule-name"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.rotate_log_by_file_size"); ok {
			res["rotate-log-by-file-size"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.rotate_log_file_size_threshold"); ok {
			res["rotate-log-file-size-threshold"] = v.(int)
		}
		if v, ok := d.GetOk("logs_settings.rotate_log_on_schedule"); ok {
			res["rotate-log-on-schedule"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.rotate_log_schedule_name"); ok {
			res["rotate-log-schedule-name"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.smart_event_intro_correletion_unit"); ok {
			res["smart-event-intro-correletion-unit"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.stop_logging_when_free_disk_space_below"); ok {
			res["stop-logging-when-free-disk-space-below"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.stop_logging_when_free_disk_space_below_threshold"); ok {
			res["stop-logging-when-free-disk-space-below-threshold"] = v.(int)
		}
		if v, ok := d.GetOk("logs_settings.turn_on_qos_logging"); ok {
			res["turn-on-qos-logging"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.update_account_log_every"); ok {
			res["update-account-log-every"] = v.(bool)
		}
		checkpointHost["logs-settings"] = res
	}

	if v, ok := d.GetOkExists("save_logs_locally"); ok {
		checkpointHost["save-logs-locally"] = v.(bool)
	}

	if v, ok := d.GetOk("send_alerts_to_server"); ok {
		checkpointHost["send-alerts-to-server"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("send_logs_to_backup_server"); ok {
		checkpointHost["send-logs-to-backup-server"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("send_logs_to_server"); ok {
		checkpointHost["send-logs-to-server"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		checkpointHost["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		checkpointHost["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		checkpointHost["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		checkpointHost["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		checkpointHost["ignore-errors"] = v.(bool)
	}

	log.Println("Create CheckpointHost - Map = ", checkpointHost)

	addCheckpointHostRes, err := client.ApiCall("add-checkpoint-host", checkpointHost, client.GetSessionID(), true, false)
	if err != nil || !addCheckpointHostRes.Success {
		if addCheckpointHostRes.ErrorMsg != "" {
			return fmt.Errorf(addCheckpointHostRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addCheckpointHostRes.GetData()["uid"].(string))

	return readManagementCheckpointHost(d, m)
}

func readManagementCheckpointHost(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showCheckpointHostRes, err := client.ApiCall("show-checkpoint-host", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showCheckpointHostRes.Success {
		if objectNotFound(showCheckpointHostRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showCheckpointHostRes.ErrorMsg)
	}

	checkpointHost := showCheckpointHostRes.GetData()

	log.Println("Read CheckpointHost - Show JSON = ", checkpointHost)

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

		if v, _ := natSettingsMap["auto-rule"]; v != nil {
			natSettingsMapToReturn["auto_rule"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := natSettingsMap["ipv4-address"]; v != nil && v != "" {
			natSettingsMapToReturn["ipv4_address"] = v
		}
		if v, _ := natSettingsMap["ipv6-address"]; v != nil && v != "" {
			natSettingsMapToReturn["ipv6_address"] = v
		}
		if v, _ := natSettingsMap["hide-behind"]; v != nil {
			natSettingsMapToReturn["hide_behind"] = v
		}
		if v, _ := natSettingsMap["install-on"]; v != nil {
			natSettingsMapToReturn["install_on"] = v
		}
		if v, _ := natSettingsMap["method"]; v != nil {
			natSettingsMapToReturn["method"] = v
		}

		_, natSettingsInConf := d.GetOk("nat_settings")
		defaultNatSettings := map[string]interface{}{"auto_rule": "false"}
		if reflect.DeepEqual(defaultNatSettings, natSettingsMapToReturn) && !natSettingsInConf {
			_ = d.Set("nat_settings", map[string]interface{}{})
		} else {
			_ = d.Set("nat_settings", natSettingsMapToReturn)
		}

	} else {
		_ = d.Set("nat_settings", nil)
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

		if v, _ := managementBladesMap["network-policy-management"]; v != nil {
			managementBladesMapToReturn["network_policy_management"] = v.(bool)
		}
		if v, _ := managementBladesMap["logging-and-status"]; v != nil {
			managementBladesMapToReturn["logging_and_status"] = v.(bool)
		}
		if v, _ := managementBladesMap["smart-event-server"]; v != nil {
			managementBladesMapToReturn["smart_event_server"] = v.(bool)
		}
		if v, _ := managementBladesMap["smart-event-correlation"]; v != nil {
			managementBladesMapToReturn["smart_event_correlation"] = v.(bool)
		}
		if v, _ := managementBladesMap["endpoint-policy"]; v != nil {
			managementBladesMapToReturn["endpoint_policy"] = v.(bool)
		}
		if v, _ := managementBladesMap["compliance"]; v != nil {
			managementBladesMapToReturn["compliance"] = v.(bool)
		}
		if v, _ := managementBladesMap["user-directory"]; v != nil {
			managementBladesMapToReturn["user_directory"] = v.(bool)
		}
		if v, _ := managementBladesMap["secondary"]; v != nil {
			managementBladesMapToReturn["secondary"] = v.(bool)
		}
		if v, _ := managementBladesMap["identity-logging"]; v != nil {
			managementBladesMapToReturn["identity_logging"] = v.(bool)
		}

		_, managementBladesInConf := d.GetOk("management_blades")
		defaultManagementBlades := map[string]interface{}{"network_policy_management": false, "logging_and_status": false, "smart_event_server": false, "smart_event_correlation": false, "endpoint_policy": false, "compliance": false, "user_directory": false, "secondary": true, "identity_logging": false}
		if reflect.DeepEqual(defaultManagementBlades, managementBladesMapToReturn) && !managementBladesInConf {
			_ = d.Set("management_blades", map[string]interface{}{})
		} else {
			_ = d.Set("management_blades", managementBladesMapToReturn)
		}

	} else {
		_ = d.Set("management_blades", nil)
	}

	if checkpointHost["logs-settings"] != nil {

		logsSettingsMap := checkpointHost["logs-settings"].(map[string]interface{})

		logsSettingsMapToReturn := make(map[string]interface{})

		if v, _ := logsSettingsMap["free-disk-space-metrics"]; v != nil {
			logsSettingsMapToReturn["free_disk_space_metrics"] = v
		}
		if v, _ := logsSettingsMap["accept-syslog-messages"]; v != nil {
			logsSettingsMapToReturn["accept_syslog_messages"] = v.(bool)
		}
		if v, _ := logsSettingsMap["alert-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below"] = v.(bool)
		}
		if v, _ := logsSettingsMap["alert-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_threshold"] = v
		}
		if v, _ := logsSettingsMap["alert-when-free-disk-space-below-type"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_type"] = v
		}
		if v, _ := logsSettingsMap["before-delete-keep-logs-from-the-last-days"]; v != nil {
			logsSettingsMapToReturn["before_delete_keep_logs_from_the_last_days"] = v.(bool)
		}
		if v, _ := logsSettingsMap["before-delete-keep-logs-from-the-last-days-threshold"]; v != nil {
			logsSettingsMapToReturn["before_delete_keep_logs_from_the_last_days_threshold"] = v
		}
		if v, _ := logsSettingsMap["before-delete-run-script"]; v != nil {
			logsSettingsMapToReturn["before_delete_run_script"] = v.(bool)
		}
		if v, _ := logsSettingsMap["before-delete-run-script-command"]; v != nil {
			logsSettingsMapToReturn["before_delete_run_script_command"] = v
		}
		if v, _ := logsSettingsMap["delete-index-files-older-than-days"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_older_than_days"] = v.(bool)
		}
		if v, _ := logsSettingsMap["delete-index-files-older-than-days-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_older_than_days_threshold"] = v
		}
		if v, _ := logsSettingsMap["delete-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below"] = v.(bool)
		}
		if v, _ := logsSettingsMap["delete-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below_threshold"] = v
		}
		if v, _ := logsSettingsMap["detect-new-citrix-ica-application-names"]; v != nil {
			logsSettingsMapToReturn["detect_new_citrix_ica_application_names"] = v.(bool)
		}
		if v, _ := logsSettingsMap["enable-log-indexing"]; v != nil {
			logsSettingsMapToReturn["enable_log_indexing"] = v.(bool)
		}
		if v, _ := logsSettingsMap["forward-logs-to-log-server"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server"] = v.(bool)
		}
		if v, _ := logsSettingsMap["forward-logs-to-log-server-name"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server_name"] = v
		}
		if v, _ := logsSettingsMap["forward-logs-to-log-server-schedule-name"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server_schedule_name"] = v
		}
		if v, _ := logsSettingsMap["rotate-log-by-file-size"]; v != nil {
			logsSettingsMapToReturn["rotate_log_by_file_size"] = v.(bool)
		}
		if v, _ := logsSettingsMap["rotate-log-file-size-threshold"]; v != nil {
			logsSettingsMapToReturn["rotate_log_file_size_threshold"] = v
		}
		if v, _ := logsSettingsMap["rotate-log-on-schedule"]; v != nil {
			logsSettingsMapToReturn["rotate_log_on_schedule"] = v.(bool)
		}
		if v, _ := logsSettingsMap["rotate-log-schedule-name"]; v != nil {
			logsSettingsMapToReturn["rotate_log_schedule_name"] = v
		}
		if v, _ := logsSettingsMap["smart-event-intro-correletion-unit"]; v != nil {
			logsSettingsMapToReturn["smart_event_intro_correletion_unit"] = v.(bool)
		}
		if v, _ := logsSettingsMap["stop-logging-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below"] = v.(bool)
		}
		if v, _ := logsSettingsMap["stop-logging-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below_threshold"] = v
		}
		if v, _ := logsSettingsMap["turn-on-qos-logging"]; v != nil {
			logsSettingsMapToReturn["turn_on_qos_logging"] = v.(bool)
		}
		if v, _ := logsSettingsMap["update-account-log-every"]; v != nil {
			logsSettingsMapToReturn["update_account_log_every"] = v
		}
		_ = d.Set("logs_settings", logsSettingsMapToReturn)
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

func updateManagementCheckpointHost(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	checkpointHost := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		checkpointHost["name"] = oldName
		checkpointHost["new-name"] = newName
	} else {
		checkpointHost["name"] = d.Get("name")
	}

	if d.HasChange("interfaces") {

		if v, ok := d.GetOk("interfaces"); ok {

			interfacesList := v.([]interface{})

			var interfacesPayload []map[string]interface{}

			for i := range interfacesList {

				Payload := make(map[string]interface{})

				if d.HasChange("interfaces." + strconv.Itoa(i) + ".name") {
					Payload["name"] = d.Get("interfaces." + strconv.Itoa(i) + ".name")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".subnet4") {
					Payload["subnet4"] = d.Get("interfaces." + strconv.Itoa(i) + ".subnet4")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".subnet6") {
					Payload["subnet6"] = d.Get("interfaces." + strconv.Itoa(i) + ".subnet6")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".mask_length4") {
					Payload["mask-length4"] = d.Get("interfaces." + strconv.Itoa(i) + ".mask_length4")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".mask_length6") {
					Payload["mask-length6"] = d.Get("interfaces." + strconv.Itoa(i) + ".mask_length6")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".subnet_mask") {
					Payload["subnet-mask"] = d.Get("interfaces." + strconv.Itoa(i) + ".subnet_mask")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".color") {
					Payload["color"] = d.Get("interfaces." + strconv.Itoa(i) + ".color")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".comments") {
					Payload["comments"] = d.Get("interfaces." + strconv.Itoa(i) + ".comments")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".ignore_warnings") {
					Payload["ignore-warnings"] = d.Get("interfaces." + strconv.Itoa(i) + ".ignore_warnings")
				}
				if d.HasChange("interfaces." + strconv.Itoa(i) + ".ignore_errors") {
					Payload["ignore-errors"] = d.Get("interfaces." + strconv.Itoa(i) + ".ignore_errors")
				}
				interfacesPayload = append(interfacesPayload, Payload)
			}
			checkpointHost["interfaces"] = interfacesPayload
		} else {
			oldinterfaces, _ := d.GetChange("interfaces")
			var interfacesToDelete []interface{}
			for _, oldInterface := range oldinterfaces.([]interface{}) {
				interfacesToDelete = append(interfacesToDelete, oldInterface.(map[string]interface{})["name"].(string))
			}
			checkpointHost["interfaces"] = map[string]interface{}{"remove": interfacesToDelete}
		}
	}

	if ok := d.HasChange("ipv4_address"); ok {
		checkpointHost["ipv4-address"] = d.Get("ipv4_address")
	}

	if ok := d.HasChange("ipv6_address"); ok {
		checkpointHost["ipv6-address"] = d.Get("ipv6_address")
	}

	if d.HasChange("nat_settings") {

		if _, ok := d.GetOk("nat_settings"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOk("nat_settings.auto_rule"); ok {
				res["auto-rule"] = v
			}
			if v, ok := d.GetOk("nat_settings.ipv4_address"); ok {
				res["ipv4-address"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.ipv6_address"); ok {
				res["ipv6-address"] = v.(string)
			}
			if d.HasChange("nat_settings.hide_behind") {
				res["hide-behind"] = d.Get("nat_settings.hide_behind")
			}
			if d.HasChange("nat_settings.install_on") {
				res["install-on"] = d.Get("nat_settings.install_on")
			}
			if d.HasChange("nat_settings.method") {
				res["method"] = d.Get("nat_settings.method")
			}
			checkpointHost["nat-settings"] = res
		} else {
			checkpointHost["nat-settings"] = map[string]interface{}{"auto-rule": false}
		}
	}

	if ok := d.HasChange("one_time_password"); ok {
		checkpointHost["one-time-password"] = d.Get("one_time_password")
	}

	if ok := d.HasChange("hardware"); ok {
		checkpointHost["hardware"] = d.Get("hardware")
	}

	if ok := d.HasChange("os"); ok {
		checkpointHost["os"] = d.Get("os")
	}

	if ok := d.HasChange("version"); ok {
		checkpointHost["version"] = d.Get("version")
	}

	if d.HasChange("management_blades") {

		if _, ok := d.GetOk("management_blades"); ok {

			res := make(map[string]interface{})

			if d.HasChange("management_blades.network_policy_management") {
				res["network-policy-management"] = d.Get("management_blades.network_policy_management")
			}
			if d.HasChange("management_blades.logging_and_status") {
				res["logging-and-status"] = d.Get("management_blades.logging_and_status")
			}
			if d.HasChange("management_blades.smart_event_server") {
				res["smart-event-server"] = d.Get("management_blades.smart_event_server")
			}
			if d.HasChange("management_blades.smart_event_correlation") {
				res["smart-event-correlation"] = d.Get("management_blades.smart_event_correlation")
			}
			if d.HasChange("management_blades.endpoint_policy") {
				res["endpoint-policy"] = d.Get("management_blades.endpoint_policy")
			}
			if d.HasChange("management_blades.compliance") {
				res["compliance"] = d.Get("management_blades.compliance")
			}
			if d.HasChange("management_blades.user_directory") {
				res["user-directory"] = d.Get("management_blades.user_directory")
			}
			checkpointHost["management-blades"] = res
		} else {
			checkpointHost["management-blades"] = map[string]interface{}{"logging-and-status": false, "smart-event-server": false, "smart-event-correlation": false, "network-policy-management": false, "user-directory": false, "compliance": false, "endpoint-policy": false, "secondary": true, "identity-logging": false}
		}
	}

	if d.HasChange("logs_settings") {

		if _, ok := d.GetOk("logs_settings"); ok {

			res := make(map[string]interface{})

			if d.HasChange("logs_settings.free_disk_space_metrics") {
				res["free-disk-space-metrics"] = d.Get("logs_settings.free_disk_space_metrics")
			}
			if d.HasChange("logs_settings.accept_syslog_messages") {
				res["accept-syslog-messages"] = d.Get("logs_settings.accept_syslog_messages")
			}
			if d.HasChange("logs_settings.alert_when_free_disk_space_below") {
				res["alert-when-free-disk-space-below"] = d.Get("logs_settings.alert_when_free_disk_space_below")
			}
			if d.HasChange("logs_settings.alert_when_free_disk_space_below_threshold") {
				res["alert-when-free-disk-space-below-threshold"] = d.Get("logs_settings.alert_when_free_disk_space_below_threshold")
			}
			if d.HasChange("logs_settings.alert_when_free_disk_space_below_type") {
				res["alert-when-free-disk-space-below-type"] = d.Get("logs_settings.alert_when_free_disk_space_below_type")
			}
			if d.HasChange("logs_settings.before_delete_keep_logs_from_the_last_days") {
				res["before-delete-keep-logs-from-the-last-days"] = d.Get("logs_settings.before_delete_keep_logs_from_the_last_days")
			}
			if d.HasChange("logs_settings.before_delete_keep_logs_from_the_last_days_threshold") {
				res["before-delete-keep-logs-from-the-last-days-threshold"] = d.Get("logs_settings.before_delete_keep_logs_from_the_last_days_threshold")
			}
			if d.HasChange("logs_settings.before_delete_run_script") {
				res["before-delete-run-script"] = d.Get("logs_settings.before_delete_run_script")
			}
			if d.HasChange("logs_settings.before_delete_run_script_command") {
				res["before-delete-run-script-command"] = d.Get("logs_settings.before_delete_run_script_command")
			}
			if d.HasChange("logs_settings.delete_index_files_older_than_days") {
				res["delete-index-files-older-than-days"] = d.Get("logs_settings.delete_index_files_older_than_days")
			}
			if d.HasChange("logs_settings.delete_index_files_older_than_days_threshold") {
				res["delete-index-files-older-than-days-threshold"] = d.Get("logs_settings.delete_index_files_older_than_days_threshold")
			}
			if d.HasChange("logs_settings.delete_when_free_disk_space_below") {
				res["delete-when-free-disk-space-below"] = d.Get("logs_settings.delete_when_free_disk_space_below")
			}
			if d.HasChange("logs_settings.delete_when_free_disk_space_below_threshold") {
				res["delete-when-free-disk-space-below-threshold"] = d.Get("logs_settings.delete_when_free_disk_space_below_threshold")
			}
			if d.HasChange("logs_settings.detect_new_citrix_ica_application_names") {
				res["detect-new-citrix-ica-application-names"] = d.Get("logs_settings.detect_new_citrix_ica_application_names")
			}
			if d.HasChange("logs_settings.enable_log_indexing") {
				res["enable-log-indexing"] = d.Get("logs_settings.enable_log_indexing")
			}
			if d.HasChange("logs_settings.forward_logs_to_log_server") {
				res["forward-logs-to-log-server"] = d.Get("logs_settings.forward_logs_to_log_server")
			}
			if d.HasChange("logs_settings.forward_logs_to_log_server_name") {
				res["forward-logs-to-log-server-name"] = d.Get("logs_settings.forward_logs_to_log_server_name")
			}
			if d.HasChange("logs_settings.forward_logs_to_log_server_schedule_name") {
				res["forward-logs-to-log-server-schedule-name"] = d.Get("logs_settings.forward_logs_to_log_server_schedule_name")
			}
			if d.HasChange("logs_settings.rotate_log_by_file_size") {
				res["rotate-log-by-file-size"] = d.Get("logs_settings.rotate_log_by_file_size")
			}
			if d.HasChange("logs_settings.rotate_log_file_size_threshold") {
				res["rotate-log-file-size-threshold"] = d.Get("logs_settings.rotate_log_file_size_threshold")
			}
			if d.HasChange("logs_settings.rotate_log_on_schedule") {
				res["rotate-log-on-schedule"] = d.Get("logs_settings.rotate_log_on_schedule")
			}
			if d.HasChange("logs_settings.rotate_log_schedule_name") {
				res["rotate-log-schedule-name"] = d.Get("logs_settings.rotate_log_schedule_name")
			}
			if d.HasChange("logs_settings.smart_event_intro_correletion_unit") {
				res["smart-event-intro-correletion-unit"] = d.Get("logs_settings.smart_event_intro_correletion_unit")
			}
			if d.HasChange("logs_settings.stop_logging_when_free_disk_space_below") {
				res["stop-logging-when-free-disk-space-below"] = d.Get("logs_settings.stop_logging_when_free_disk_space_below")
			}
			if d.HasChange("logs_settings.stop_logging_when_free_disk_space_below_threshold") {
				res["stop-logging-when-free-disk-space-below-threshold"] = d.Get("logs_settings.stop_logging_when_free_disk_space_below_threshold")
			}
			if d.HasChange("logs_settings.turn_on_qos_logging") {
				res["turn-on-qos-logging"] = d.Get("logs_settings.turn_on_qos_logging")
			}
			if d.HasChange("logs_settings.update_account_log_every") {
				res["update-account-log-every"] = d.Get("logs_settings.update_account_log_every")
			}
			checkpointHost["logs-settings"] = res
		}
	}

	if v, ok := d.GetOkExists("save_logs_locally"); ok {
		checkpointHost["save-logs-locally"] = v.(bool)
	}

	if d.HasChange("send_alerts_to_server") {
		if v, ok := d.GetOk("send_alerts_to_server"); ok {
			checkpointHost["send_alerts_to_server"] = v.(*schema.Set).List()
		} else {
			oldsendAlertsToServer, _ := d.GetChange("send_alerts_to_server")
			checkpointHost["send_alerts_to_server"] = map[string]interface{}{"remove": oldsendAlertsToServer.(*schema.Set).List()}
		}
	}

	if d.HasChange("send_logs_to_backup_server") {
		if v, ok := d.GetOk("send_logs_to_backup_server"); ok {
			checkpointHost["send_logs_to_backup_server"] = v.(*schema.Set).List()
		} else {
			oldsendLogsToBackupServer, _ := d.GetChange("send_logs_to_backup_server")
			checkpointHost["send_logs_to_backup_server"] = map[string]interface{}{"remove": oldsendLogsToBackupServer.(*schema.Set).List()}
		}
	}

	if d.HasChange("send_logs_to_server") {
		if v, ok := d.GetOk("send_logs_to_server"); ok {
			checkpointHost["send_logs_to_server"] = v.(*schema.Set).List()
		} else {
			oldsendLogsToServer, _ := d.GetChange("send_logs_to_server")
			checkpointHost["send_logs_to_server"] = map[string]interface{}{"remove": oldsendLogsToServer.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			checkpointHost["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			checkpointHost["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		checkpointHost["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		checkpointHost["comments"] = d.Get("comments")
	}

	if d.HasChange("groups") {
		if v, ok := d.GetOk("groups"); ok {
			checkpointHost["groups"] = v.(*schema.Set).List()
		} else {
			oldGroups, _ := d.GetChange("groups")
			checkpointHost["groups"] = map[string]interface{}{"remove": oldGroups.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		checkpointHost["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		checkpointHost["ignore-errors"] = v.(bool)
	}

	log.Println("Update CheckpointHost - Map = ", checkpointHost)

	updateCheckpointHostRes, err := client.ApiCall("set-checkpoint-host", checkpointHost, client.GetSessionID(), true, false)
	if err != nil || !updateCheckpointHostRes.Success {
		if updateCheckpointHostRes.ErrorMsg != "" {
			return fmt.Errorf(updateCheckpointHostRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementCheckpointHost(d, m)
}

func deleteManagementCheckpointHost(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	checkpointHostPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete CheckpointHost")

	deleteCheckpointHostRes, err := client.ApiCall("delete-checkpoint-host", checkpointHostPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteCheckpointHostRes.Success {
		if deleteCheckpointHostRes.ErrorMsg != "" {
			return fmt.Errorf(deleteCheckpointHostRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")
	return nil
}
