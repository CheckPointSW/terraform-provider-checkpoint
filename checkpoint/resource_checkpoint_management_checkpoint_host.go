package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"math"
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
				Sensitive:   true,
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
			res["auto-rule"] = v
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
			res["forward-logs-to-log-server"] = v
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

	addCheckpointHostRes, err := client.ApiCall("add-checkpoint-host", checkpointHost, client.GetSessionID(), true, client.IsProxyUsed())
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

	showCheckpointHostRes, err := client.ApiCall("show-checkpoint-host", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

		_, managementBladesInConf := d.GetOk("managemen" +
			"t_blades")
		defaultManagementBlades := map[string]interface{}{"network_policy_management": false, "logging_and_status": false, "smart_event_server": false, "smart_event_correlation": false, "endpoint_policy": false, "compliance": false, "user_directory": false, "secondary": true, "identity_logging": false}
		if reflect.DeepEqual(defaultManagementBlades, managementBladesMapToReturn) && !managementBladesInConf {
			_ = d.Set("management_blades", map[string]interface{}{})
		} else {
			_ = d.Set("management_blades", managementBladesMapToReturn)
		}

	} else {
		_ = d.Set("management_blades", nil)
	}

	if v := checkpointHost["logs-settings"]; v != nil {
		logSettingsJson := v.(map[string]interface{})
		logSettingsState := make(map[string]interface{})
		defaultLogsSettings := map[string]interface{}{
			"alert_when_free_disk_space_below":                     "true",
			"free_disk_space_metrics":                              "mbytes",
			"alert_when_free_disk_space_below_type":                "popup alert",
			"alert_when_free_disk_space_below_threshold":           "20",
			"before_delete_keep_logs_from_the_last_days":           "false",
			"before_delete_keep_logs_from_the_last_days_threshold": "3664",
			"before_delete_run_script":                             "false",
			"before_delete_run_script_command":                     "",
			"delete_index_files_older_than_days":                   "false",
			"delete_index_files_older_than_days_threshold":         "14",
			"delete_index_files_when_index_size_above":             "false",
			"delete_index_files_when_index_size_above_threshold":   "100000",
			"delete_when_free_disk_space_below":                    "true",
			"delete_when_free_disk_space_below_threshold":          "5000",
			"detect_new_citrix_ica_application_names":              "false",
			"forward_logs_to_log_server":                           "false",
			"perform_log_rotate_before_log_forwarding":             "false",
			"rotate_log_by_file_size":                              "false",
			"rotate_log_file_size_threshold":                       "1000",
			"rotate_log_on_schedule":                               "false",
			"rotate-log-schedule-name":                             "mgmt_schd",
			"smart_event_intro_correletion_unit":                   "false",
			"stop_logging_when_free_disk_space_below":              "false",
			"stop_logging_when_free_disk_space_below_threshold":    "100",
			"turn_on_qos_logging":                                  "true",
			"update_account_log_every":                             "3600",
		}
		if v := logSettingsJson["alert-when-free-disk-space-below"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.alert_when_free_disk_space_below", defaultLogsSettings["alert_when_free_disk_space_below"].(string)) {
			logSettingsState["alert_when_free_disk_space_below"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["free-disk-space-metrics"]; v != nil && isArgDefault(v.(string), d, "logs_settings.free_disk_space_metrics", defaultLogsSettings["free_disk_space_metrics"].(string)) {
			logSettingsState["free_disk_space_metrics"] = v.(string)
		}
		if v := logSettingsJson["alert-when-free-disk-space-below-threshold"]; v != nil && isArgDefault(strconv.Itoa(int(math.Round(v.(float64)))), d, "logs_settings.alert_when_free_disk_space_below_threshold", defaultLogsSettings["alert_when_free_disk_space_below_threshold"].(string)) {
			logSettingsState["alert_when_free_disk_space_below_threshold"] = strconv.Itoa(int(math.Round(v.(float64))))
		}
		if v := logSettingsJson["alert-when-free-disk-space-below-type"]; v != nil && isArgDefault(v.(string), d, "logs_settings.alert_when_free_disk_space_below_type", defaultLogsSettings["alert_when_free_disk_space_below_type"].(string)) {
			logSettingsState["alert_when_free_disk_space_below_type"] = v.(string)
		}
		if v := logSettingsJson["before-delete-keep-logs-from-the-last-days"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.before_delete_keep_logs_from_the_last_days", defaultLogsSettings["before_delete_keep_logs_from_the_last_days"].(string)) {
			logSettingsState["before_delete_keep_logs_from_the_last_days"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["before-delete-keep-logs-from-the-last-days-threshold"]; v != nil && isArgDefault(strconv.Itoa(int(math.Round(v.(float64)))), d, "logs_settings.before_delete_keep_logs_from_the_last_days_threshold", defaultLogsSettings["before_delete_keep_logs_from_the_last_days_threshold"].(string)) {
			logSettingsState["before_delete_keep_logs_from_the_last_days_threshold"] = strconv.Itoa(int(math.Round(v.(float64))))
		}
		if v := logSettingsJson["before-delete-run-script"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.before_delete_run_script", defaultLogsSettings["before_delete_run_script"].(string)) {
			logSettingsState["before_delete_run_script"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["before-delete-run-script-command"]; v != nil && isArgDefault(v.(string), d, "logs_settings.before_delete_run_script_command", defaultLogsSettings["before_delete_run_script_command"].(string)) {
			logSettingsState["before_delete_run_script_command"] = v.(string)
		}
		if v := logSettingsJson["delete-index-files-older-than-days"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.delete_index_files_older_than_days", defaultLogsSettings["delete_index_files_older_than_days"].(string)) {
			logSettingsState["delete_index_files_older_than_days"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["delete-index-files-older-than-days-threshold"]; v != nil && isArgDefault(strconv.Itoa(int(math.Round(v.(float64)))), d, "logs_settings.delete_index_files_older_than_days_threshold", defaultLogsSettings["delete_index_files_older_than_days_threshold"].(string)) {
			logSettingsState["delete_index_files_older_than_days_threshold"] = strconv.Itoa(int(math.Round(v.(float64))))
		}
		if v := logSettingsJson["delete-when-free-disk-space-below"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.delete_when_free_disk_space_below", defaultLogsSettings["delete_when_free_disk_space_below"].(string)) {
			logSettingsState["delete_when_free_disk_space_below"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["delete-when-free-disk-space-below-threshold"]; v != nil && isArgDefault(strconv.Itoa(int(math.Round(v.(float64)))), d, "logs_settings.delete_when_free_disk_space_below_threshold", defaultLogsSettings["delete_when_free_disk_space_below_threshold"].(string)) {
			logSettingsState["delete_when_free_disk_space_below_threshold"] = strconv.Itoa(int(math.Round(v.(float64))))
		}
		if v := logSettingsJson["detect-new-citrix-ica-application-names"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.detect_new_citrix_ica_application_names", defaultLogsSettings["detect_new_citrix_ica_application_names"].(string)) {
			logSettingsState["detect_new_citrix_ica_application_names"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["forward-logs-to-log-server"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.forward_logs_to_log_server", defaultLogsSettings["forward_logs_to_log_server"].(string)) {
			logSettingsState["forward_logs_to_log_server"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["forward-logs-to-log-server-name"]; v != nil {
			logSettingsState["forward_logs_to_log_server_name"] = v.(string)
		}
		if v := logSettingsJson["forward-logs-to-log-server-schedule-name"]; v != nil {
			logSettingsState["forward_logs_to_log_server_schedule_name"] = v.(string)
		}
		if v := logSettingsJson["rotate-log-by-file-size"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.rotate_log_by_file_size", defaultLogsSettings["rotate_log_by_file_size"].(string)) {
			logSettingsState["rotate_log_by_file_size"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["rotate-log-file-size-threshold"]; v != nil && isArgDefault(strconv.Itoa(int(math.Round(v.(float64)))), d, "logs_settings.rotate_log_file_size_threshold", defaultLogsSettings["rotate_log_file_size_threshold"].(string)) {
			logSettingsState["rotate_log_file_size_threshold"] = strconv.Itoa(int(math.Round(v.(float64))))
		}
		if v := logSettingsJson["rotate-log-on-schedule"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.rotate_log_on_schedule", defaultLogsSettings["rotate_log_on_schedule"].(string)) {
			logSettingsState["rotate_log_on_schedule"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["rotate-log-schedule-name"]; v != nil {
			logSettingsState["rotate_log_schedule_name"] = v.(string)
		}
		if v := logSettingsJson["stop-logging-when-free-disk-space-below"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.stop_logging_when_free_disk_space_below", defaultLogsSettings["stop_logging_when_free_disk_space_below"].(string)) {
			logSettingsState["stop_logging_when_free_disk_space_below"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["smart-event-intro-correletion-unit"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.smart_event_intro_correletion_unit", defaultLogsSettings["smart_event_intro_correletion_unit"].(string)) {
			logSettingsState["smart_event_intro_correletion_unit"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["stop-logging-when-free-disk-space-below-threshold"]; v != nil && isArgDefault(strconv.Itoa(int(math.Round(v.(float64)))), d, "logs_settings.stop_logging_when_free_disk_space_below_threshold", defaultLogsSettings["stop_logging_when_free_disk_space_below_threshold"].(string)) {
			logSettingsState["stop_logging_when_free_disk_space_below_threshold"] = strconv.Itoa(int(math.Round(v.(float64))))
		}
		if v := logSettingsJson["turn-on-qos-logging"]; v != nil && isArgDefault(strconv.FormatBool(v.(bool)), d, "logs_settings.turn_on_qos_logging", defaultLogsSettings["turn_on_qos_logging"].(string)) {
			logSettingsState["turn_on_qos_logging"] = strconv.FormatBool(v.(bool))
		}
		if v := logSettingsJson["update-account-log-every"]; v != nil && isArgDefault(strconv.Itoa(int(math.Round(v.(float64)))), d, "logs_settings.update_account_log_every", defaultLogsSettings["update_account_log_every"].(string)) {
			logSettingsState["update_account_log_every"] = strconv.Itoa(int(math.Round(v.(float64))))
		}
		_ = d.Set("logs_settings", logSettingsState)
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

	if v, ok := d.GetOk("management_blades"); ok {
		defaultLogsSettings := map[string]interface{}{
			"network-policy-management": "false",
			"logging-and-status":        "false",
			"smart-event-server":        "false",
			"smart-event-correlation":   "false",
			"endpoint-policy":           "false",
		}
		logsSettingsJson := v.(map[string]interface{})
		res := make(map[string]interface{})
		if val, ok := logsSettingsJson["network_policy_management"]; ok {
			res["network-policy-management"] = val
		} else {
			res["network-policy-management"] = defaultLogsSettings["network-policy-management"]
		}
		if val, ok := logsSettingsJson["logging_and_status"]; ok {
			res["logging-and-status"] = val
		} else {
			res["logging-and-status"] = defaultLogsSettings["logging-and-status"]
		}
		if val, ok := logsSettingsJson["smart_event_server"]; ok {
			res["smart-event-server"] = val
		} else {
			res["smart-event-server"] = defaultLogsSettings["smart-event-server"]
		}
		if val, ok := logsSettingsJson["smart_event_correlation"]; ok {
			res["smart-event-correlation"] = val
		} else {
			res["smart-event-correlation"] = defaultLogsSettings["smart-event-correlation"]
		}
		if val, ok := logsSettingsJson["endpoint_policy"]; ok {
			res["endpoint-policy"] = val
		} else {
			res["endpoint-policy"] = defaultLogsSettings["endpoint-policy"]
		}
		if val, ok := logsSettingsJson["compliance"]; ok {
			res["compliance"] = val
		}
		if val, ok := logsSettingsJson["user_directory"]; ok {
			res["user-directory"] = val
		}
		checkpointHost["management-blades"] = res
	}

	if ok := d.HasChange("logs_settings"); ok {
		defaultLogsSettings := map[string]interface{}{
			"alert-when-free-disk-space-below":                     "true",
			"free-disk-space-metrics":                              "mbytes",
			"alert-when-free-disk-space-below-type":                "popup alert",
			"alert-when-free-disk-space-below-threshold":           20,
			"before-delete-keep-logs-from-the-last-days":           "false",
			"before-delete-keep-logs-from-the-last-days-threshold": 3664,
			"before-delete-run-script":                             "false",
			"before-delete-run-script-command":                     "",
			"delete-index-files-older-than-days":                   "false",
			"delete-index-files-older-than-days-threshold":         14,
			"delete-when-free-disk-space-below":                    "true",
			"delete-when-free-disk-space-below-threshold":          5000,
			"detect-new-citrix-ica-application-names":              "false",
			"forward-logs-to-log-server":                           "false",
			"rotate-log-by-file-size":                              "false",
			"rotate-log-file-size-threshold":                       1000,
			"rotate-log-on-schedule":                               "false",
			"rotate-log-schedule-name":                             "mgmt_schd",
			"smart-event-intro-correletion-unit":                   "false",
			"stop-logging-when-free-disk-space-below":              "false",
			"stop-logging-when-free-disk-space-below-threshold":    100,
			"turn-on-qos-logging":                                  "true",
			"update-account-log-every":                             3600,
		}
		if v, ok := d.GetOk("logs_settings"); ok {
			logsSettingsJson := v.(map[string]interface{})
			logsSettings := make(map[string]interface{})
			if val, ok := logsSettingsJson["alert_when_free_disk_space_below"]; ok {
				logsSettings["alert-when-free-disk-space-below"] = val
			} else {
				logsSettings["alert-when-free-disk-space-below"] = defaultLogsSettings["alert-when-free-disk-space-below"]
			}
			if val, ok := logsSettingsJson["alert_when_free_disk_space_below_metrics"]; ok {
				logsSettings["free-disk-space-metrics"] = val
			} else {
				logsSettings["free-disk-space-metrics"] = defaultLogsSettings["free-disk-space-metrics"]
			}
			if val, ok := logsSettingsJson["alert_when_free_disk_space_below_threshold"]; ok {
				logsSettings["alert-when-free-disk-space-below-threshold"] = val
			} else {
				logsSettings["alert-when-free-disk-space-below-threshold"] = defaultLogsSettings["alert-when-free-disk-space-below-threshold"]
			}
			if val, ok := logsSettingsJson["alert_when_free_disk_space_below_type"]; ok {
				logsSettings["alert-when-free-disk-space-below-type"] = val
			} else {
				logsSettings["alert-when-free-disk-space-below-type"] = defaultLogsSettings["alert-when-free-disk-space-below-type"]
			}
			if val, ok := logsSettingsJson["before_delete_keep_logs_from_the_last_days"]; ok {
				logsSettings["before-delete-keep-logs-from-the-last-days"] = val
			} else {
				logsSettings["before-delete-keep-logs-from-the-last-days"] = defaultLogsSettings["before-delete-keep-logs-from-the-last-days"]
			}
			if val, ok := logsSettingsJson["before_delete_keep_logs_from_the_last_days_threshold"]; ok {
				logsSettings["before-delete-keep-logs-from-the-last-days-threshold"] = val
			} else {
				logsSettings["before-delete-keep-logs-from-the-last-days-threshold"] = defaultLogsSettings["before-delete-keep-logs-from-the-last-days-threshold"]
			}
			if val, ok := logsSettingsJson["before_delete_run_script"]; ok {
				logsSettings["before-delete-run-script"] = val
			} else {
				logsSettings["before-delete-run-script"] = defaultLogsSettings["before-delete-run-script"]
			}
			if val, ok := logsSettingsJson["before_delete_run_script_command"]; ok {
				logsSettings["before-delete-run-script-command"] = val
			} else {
				logsSettings["before-delete-run-script-command"] = defaultLogsSettings["before-delete-run-script-command"]
			}
			if val, ok := logsSettingsJson["delete_index_files_older_than_days"]; ok {
				logsSettings["delete-index-files-older-than-days"] = val
			} else {
				logsSettings["delete-index-files-older-than-days"] = defaultLogsSettings["delete-index-files-older-than-days"]
			}
			if val, ok := logsSettingsJson["delete_index_files_older_than_days_threshold"]; ok {
				logsSettings["delete-index-files-older-than-days-threshold"] = val
			} else {
				logsSettings["delete-index-files-older-than-days-threshold"] = defaultLogsSettings["delete-index-files-older-than-days-threshold"]
			}
			if val, ok := logsSettingsJson["delete_when_free_disk_space_below"]; ok {
				logsSettings["delete-when-free-disk-space-below"] = val
			} else {
				logsSettings["delete-when-free-disk-space-below"] = defaultLogsSettings["delete-when-free-disk-space-below"]
			}
			if val, ok := logsSettingsJson["delete_when_free_disk_space_below_threshold"]; ok {
				logsSettings["delete-when-free-disk-space-below-threshold"] = val
			} else {
				logsSettings["delete-when-free-disk-space-below-threshold"] = defaultLogsSettings["delete-when-free-disk-space-below-threshold"]
			}
			if val, ok := logsSettingsJson["detect_new_citrix_ica_application_names"]; ok {
				logsSettings["detect-new-citrix-ica-application-names"] = val
			} else {
				logsSettings["detect-new-citrix-ica-application-names"] = defaultLogsSettings["detect-new-citrix-ica-application-names"]
			}
			if val, ok := logsSettingsJson["forward_logs_to_log_server"]; ok {
				logsSettings["forward-logs-to-log-server"] = val
			} else {
				logsSettings["forward-logs-to-log-server"] = defaultLogsSettings["forward-logs-to-log-server"]
			}
			if val, ok := logsSettingsJson["forward_logs_to_log_server_name"]; ok {
				logsSettings["forward-logs-to-log-server-name"] = val
			}
			if val, ok := logsSettingsJson["forward_logs_to_log_server_schedule_name"]; ok {
				logsSettings["forward-logs-to-log-server-schedule-name"] = val
			}
			if val, ok := logsSettingsJson["rotate_log_by_file_size"]; ok {
				logsSettings["rotate-log-by-file-size"] = val
			} else {
				logsSettings["rotate-log-by-file-size"] = defaultLogsSettings["rotate-log-by-file-size"]
			}
			if val, ok := logsSettingsJson["rotate_log_file_size_threshold"]; ok {
				logsSettings["rotate-log-file-size-threshold"] = val
			} else {
				logsSettings["rotate-log-file-size-threshold"] = defaultLogsSettings["rotate-log-file-size-threshold"]
			}
			if val, ok := logsSettingsJson["rotate_log_on_schedule"]; ok {
				logsSettings["rotate-log-on-schedule"] = val
			} else {
				logsSettings["rotate-log-on-schedule"] = defaultLogsSettings["rotate-log-on-schedule"]
			}
			if val, ok := logsSettingsJson["rotate_log_schedule_name"]; ok {
				logsSettings["rotate-log-schedule-name"] = val
			}
			if val, ok := logsSettingsJson["smart_event_intro_correletion_unit"]; ok {
				logsSettings["smart-event-intro-correletion-unit"] = val
			} else {
				logsSettings["smart-event-intro-correletion-unit"] = defaultLogsSettings["smart-event-intro-correletion-unit"]
			}
			if val, ok := logsSettingsJson["stop_logging_when_free_disk_space_below"]; ok {
				logsSettings["stop-logging-when-free-disk-space-below"] = val
			} else {
				logsSettings["stop-logging-when-free-disk-space-below"] = defaultLogsSettings["stop-logging-when-free-disk-space-below"]
			}
			if val, ok := logsSettingsJson["stop_logging_when_free_disk_space_below_threshold"]; ok {
				logsSettings["stop-logging-when-free-disk-space-below-threshold"] = val
			} else {
				logsSettings["stop-logging-when-free-disk-space-below-threshold"] = defaultLogsSettings["stop-logging-when-free-disk-space-below-threshold"]
			}
			if val, ok := logsSettingsJson["turn_on_qos_logging"]; ok {
				logsSettings["turn-on-qos-logging"] = val
			} else {
				logsSettings["turn-on-qos-logging"] = defaultLogsSettings["turn-on-qos-logging"]
			}
			if val, ok := logsSettingsJson["update_account_log_every"]; ok {
				logsSettings["update-account-log-every"] = val
			} else {
				logsSettings["update-account-log-every"] = defaultLogsSettings["update-account-log-every"]
			}

			checkpointHost["logs-settings"] = logsSettings
		} else {
			checkpointHost["logs-settings"] = defaultLogsSettings
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

	updateCheckpointHostRes, err := client.ApiCall("set-checkpoint-host", checkpointHost, client.GetSessionID(), true, client.IsProxyUsed())
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

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		checkpointHostPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		checkpointHostPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete CheckpointHost")

	deleteCheckpointHostRes, err := client.ApiCall("delete-checkpoint-host", checkpointHostPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteCheckpointHostRes.Success {
		if deleteCheckpointHostRes.ErrorMsg != "" {
			return fmt.Errorf(deleteCheckpointHostRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")
	return nil
}
