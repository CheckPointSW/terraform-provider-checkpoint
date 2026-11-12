package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementExportAccessRulebase() *schema.Resource {
	return &schema.Resource{
		Create: createManagementExportAccessRulebase,
		Read:   readManagementExportAccessRulebase,
		Delete: deleteManagementExportAccessRulebase,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object unique identifier.",
			},
			"package": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Name of the package.",
			},
			"show_expiration_settings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to calculate and show \"expiration date settings\" field in reply.",
			},
			"show_hits": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Show hitcount data.",
			},
			"use_object_dictionary": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "N/A",
			},
			"hits_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Hitcount settings, define the range if hits to show.",
				ForceNew:    true,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from_date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.",
						},
						"target": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Target gateway name or UID.",
						},
						"to_date": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Format: YYYY-MM-DD, YYYY-mm-ddThh:mm:ss.",
						},
					},
				},
			},
			"dereference_group_members": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to dereference \"members\" field by details level for every object in reply.",
			},
			"show_membership": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether to calculate and show \"groups\" field for every object in reply.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Asynchronous task unique identifier. Use show-task command to check the progress of the task.",
			},
		},
	}
}

func createManagementExportAccessRulebase(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	}

	if v, ok := d.GetOk("package"); ok {
		payload["package"] = v.(string)
	}

	if v, ok := d.GetOkExists("show_expiration_settings"); ok {
		payload["show-expiration-settings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_hits"); ok {
		payload["show-hits"] = v.(bool)
	}

	if v, ok := d.GetOkExists("use_object_dictionary"); ok {
		payload["use-object-dictionary"] = v.(bool)
	}

	if v, ok := d.GetOk("hits_settings"); ok {

		hitsSettingsList := v.([]interface{})

		if len(hitsSettingsList) > 0 {

			hitsSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("hits_settings.0.from_date"); ok {
				hitsSettingsPayload["from-date"] = v.(string)
			}
			if v, ok := d.GetOk("hits_settings.0.target"); ok {
				hitsSettingsPayload["target"] = v.(string)
			}
			if v, ok := d.GetOk("hits_settings.0.to_date"); ok {
				hitsSettingsPayload["to-date"] = v.(string)
			}
			payload["hits-settings"] = hitsSettingsPayload
		}
	}
	if v, ok := d.GetOkExists("dereference_group_members"); ok {
		payload["dereference-group-members"] = v.(bool)
	}

	if v, ok := d.GetOkExists("show_membership"); ok {
		payload["show-membership"] = v.(bool)
	}

	ExportAccessRulebaseRes, err := client.ApiCall("export-access-rulebase", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !ExportAccessRulebaseRes.Success {
		return fmt.Errorf(ExportAccessRulebaseRes.ErrorMsg)
	}

	d.SetId("export-access-rulebase-" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(ExportAccessRulebaseRes.GetData()))
	return nil
}

func readManagementExportAccessRulebase(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementExportAccessRulebase(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
