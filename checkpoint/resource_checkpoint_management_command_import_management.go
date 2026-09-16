package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementImportManagement() *schema.Resource {
	return &schema.Resource{
		Create: createManagementImportManagement,
		Read:   readManagementImportManagement,
		Delete: deleteManagementImportManagement,
		Schema: map[string]*schema.Schema{
			"file_path": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Path to the exported database file to be imported.",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Domain name to be imported. Must be unique in the Multi-Domain Server. Required only for importing the Security Management Server into the Multi-Domain Server.",
			},
			"domain_ip_address": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "IPv4 address for the imported Domain. Required only for importing the Security Management Server into the Multi-Domain Server.",
			},
			"domain_server_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Multi-Domain Server name for the imported Domain. Required only for importing the Security Management Server into the Multi-Domain Server.",
			},
			"include_logs": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "Import logs without log indexes.",
			},
			"include_logs_indexes": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "Import logs with log indexes.",
			},
			"include_endpoint_configuration": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "Include import of the Endpoint Security Management configuration files.",
			},
			"include_endpoint_database": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "Include import of the Endpoint Security Management database.",
			},
			"verify_domain_restore": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "If true, verify that the restore operation is valid for this input file and this environment. <br>Note: Restore operation will not be executed.",
			},
			"pre_import_verification_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     true,
				Description: "If true, only runs the pre-import verifications instead of the full import.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Ignoring the verification warnings. By Setting this parameter to 'true' import will not be blocked by warnings.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Asynchronous task unique identifier.",
			},
			"login_required": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "If set to \"True\", session is expired and login is required.",
			},
			"prepare_background_import": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "If 'true', the import will run in the background and 'Prepare' phase will be achieved. You can continue making changes on the Management Server during...",
			},
			"keep_cloud_sharing": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Preserve the connection of the Management Server to Check Point's Infinity Portal.<br>Use this flag after ensuring that the original Management Server...",
			},
			"domain_ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "IPv6 address for the imported Domain.",
			},
			"days_of_logs": {
				Type:        schema.TypeInt,
				Optional:    true,
				ForceNew:    true,
				Description: "Export <N> last days of logs.",
			},
			"complete_background_import": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "If 'true', import the changes you made during the 'Prepare' phase, and the 'Complete' phase will be achieved. You can't make any changes during the 'C...",
			},
			"change_ips": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "New IP addresses (IPv4, IPv6, or both) of the servers.<br><font color='red'>Required only if</font> one or more of the servers in the Security Managem...",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"new_ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The new IPv4 address of the server that migrates to a new IP address.",
						},
						"server_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The object name of the server that migrates to a new IP address.",
						},
					},
				},
			},
			"cancel_prepare_import": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "If 'true', cancels the import in background process. If you do not run this command within the number of days defined in 'show-background-upgrade-sett...",
			},
		},
	}
}

func createManagementImportManagement(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("file_path"); ok {
		payload["file-path"] = v.(string)
	}

	if v, ok := d.GetOk("domain_name"); ok {
		payload["domain-name"] = v.(string)
	}

	if v, ok := d.GetOk("domain_ip_address"); ok {
		payload["domain-ip-address"] = v.(string)
	}

	if v, ok := d.GetOk("domain_server_name"); ok {
		payload["domain-server-name"] = v.(string)
	}

	if v, ok := d.GetOkExists("include_logs"); ok {
		payload["include-logs"] = v.(bool)
	}

	if v, ok := d.GetOkExists("include_logs_indexes"); ok {
		payload["include-logs-indexes"] = v.(bool)
	}

	if v, ok := d.GetOkExists("include_endpoint_configuration"); ok {
		payload["include-endpoint-configuration"] = v.(bool)
	}

	if v, ok := d.GetOkExists("include_endpoint_database"); ok {
		payload["include-endpoint-database"] = v.(bool)
	}

	if v, ok := d.GetOkExists("verify_domain_restore"); ok {
		payload["verify-domain-restore"] = v.(bool)
	}

	if v, ok := d.GetOkExists("pre_import_verification_only"); ok {
		payload["pre-import-verification-only"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("prepare_background_import"); ok {
		payload["prepare-background-import"] = v.(bool)
	}

	if v, ok := d.GetOkExists("keep_cloud_sharing"); ok {
		payload["keep-cloud-sharing"] = v.(bool)
	}

	if v, ok := d.GetOk("domain_ipv6_address"); ok {
		payload["domain-ipv6-address"] = v.(string)
	}

	if v, ok := d.GetOk("days_of_logs"); ok {
		payload["days-of-logs"] = v.(int)
	}

	if v, ok := d.GetOkExists("complete_background_import"); ok {
		payload["complete-background-import"] = v.(bool)
	}

	if _, ok := d.GetOk("change_ips"); ok {
		changeIpsPayload := make(map[string]interface{})
		if v, ok := d.GetOk("change_ips.0.new_ipv4_address"); ok {
			changeIpsPayload["new-ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("change_ips.0.server_name"); ok {
			changeIpsPayload["server-name"] = v.(string)
		}
		payload["change-ips"] = changeIpsPayload
	}

	if v, ok := d.GetOkExists("cancel_prepare_import"); ok {
		payload["cancel-prepare-import"] = v.(bool)
	}

	ImportManagementRes, err := client.ApiCall("import-management", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !ImportManagementRes.Success {
		return fmt.Errorf("%s", ImportManagementRes.ErrorMsg)
	}

	importManagement := ImportManagementRes.GetData()

	if v := importManagement["login-required"]; v != nil {
		_ = d.Set("login_required", v)
	}

	d.SetId("import-management-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(ImportManagementRes.GetData()))
	return readManagementImportManagement(d, m)
}

func readManagementImportManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementImportManagement(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
