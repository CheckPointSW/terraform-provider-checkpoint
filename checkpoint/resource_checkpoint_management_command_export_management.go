package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementExportManagement() *schema.Resource {
	return &schema.Resource{
		Create: createManagementExportManagement,
		Read:   readManagementExportManagement,
		Delete: deleteManagementExportManagement,
		Schema: map[string]*schema.Schema{
			"file_path": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Path in which the exported database file is saved.<br><font color=\"red\">Required only</font> when not using pre-export-verification-only flag.",
			},
			"domain_name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Domain name to be exported.<br><font color=\"red\">Required only for</font> exporting a Domain from the Multi-Domain Server or backing up Domain.",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Target version.",
			},
			"include_logs": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Export logs without log indexes.",
			},
			"include_logs_indexes": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Export logs with log indexes.",
			},
			"include_endpoint_configuration": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Include export of the Endpoint Security Management configuration files.",
			},
			"include_endpoint_database": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Include export of the Endpoint Security Management database.",
			},
			"is_domain_backup": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "If true, the exported Domain will be suitable for import on the same Multi-Domain Server only.",
			},
			"is_smc_to_mds": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "If true, the exported Security Management Server will be suitable for import on the Multi-Domain Server only.",
			},
			"pre_export_verification_only": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "If true, only runs the pre-export verifications instead of the full export.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Default:     false,
				Description: "Ignoring the verification warnings. By Setting this parameter to 'true' export will not be blocked by warnings.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementExportManagement(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("file_path"); ok {
		payload["file-path"] = v.(string)
	}

	if v, ok := d.GetOk("domain_name"); ok {
		payload["domain-name"] = v.(string)
	}

	if v, ok := d.GetOk("version"); ok {
		payload["version"] = v.(string)
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

	if v, ok := d.GetOkExists("is_domain_backup"); ok {
		payload["is-domain-backup"] = v.(bool)
	}

	if v, ok := d.GetOkExists("is_smc_to_mds"); ok {
		payload["is-smc-to-mds"] = v.(bool)
	}

	if v, ok := d.GetOkExists("pre_export_verification_only"); ok {
		payload["pre-export-verification-only"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	ExportManagementRes, err := client.ApiCall("export-management", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !ExportManagementRes.Success {
		return fmt.Errorf(ExportManagementRes.ErrorMsg)
	}

	d.SetId("export-management-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(ExportManagementRes.GetData()))

	return readManagementExportManagement(d, m)
}

func readManagementExportManagement(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementExportManagement(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
