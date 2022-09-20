package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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

	ImportManagementRes, err := client.ApiCall("import-management", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !ImportManagementRes.Success {
		return fmt.Errorf(ImportManagementRes.ErrorMsg)
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
