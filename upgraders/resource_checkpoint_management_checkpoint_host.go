package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementCheckpointHostV0 is the V0 schema where nat_settings, management_blades,
// and logs_settings were TypeMap.
func ResourceManagementCheckpointHostV0() *schema.Resource {
	return &schema.Resource{
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

// ResourceManagementCheckpointHostStateUpgradeV0 converts nat_settings, management_blades, and logs_settings from TypeMap to TypeList.
func ResourceManagementCheckpointHostStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "nat_settings", "management_blades", "logs_settings"), nil
}
