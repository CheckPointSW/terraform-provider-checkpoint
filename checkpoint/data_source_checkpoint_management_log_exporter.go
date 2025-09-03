package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementLogExporter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementLogExporterRead,
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
			"target_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target server port to which logs are exported.",
			},
			"target_port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port number of the target server.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Protocol used to send logs to the target server.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to enable export.",
			},
			"attachments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log exporter attachments.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_link_to_log_attachment": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to add link to log attachment in SmartView.",
						},
						"add_link_to_log_details": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to add link to log details in SmartView.",
						},
						"add_log_attachment_id": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to add log attachment ID.",
						},
					},
				},
			},
			"data_manipulation": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log exporter data manipulation.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregate_log_updates": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to aggregate log updates.",
						},
						"format": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Logs format.",
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

func dataSourceManagementLogExporterRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showLogExporterRes, err := client.ApiCallSimple("show-log-exporter", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLogExporterRes.Success {
		return fmt.Errorf(showLogExporterRes.ErrorMsg)
	}

	logExporter := showLogExporterRes.GetData()

	log.Println("Read LogExporter - Show JSON = ", logExporter)

	if v := logExporter["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := logExporter["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := logExporter["target-server"]; v != nil {
		_ = d.Set("target_server", v)
	}

	if v := logExporter["target-port"]; v != nil {
		_ = d.Set("target_port", v)
	}

	if v := logExporter["protocol"]; v != nil {
		_ = d.Set("protocol", v)
	}

	if v := logExporter["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if logExporter["attachments"] != nil {

		attachmentsMap := logExporter["attachments"].(map[string]interface{})

		attachmentsMapToReturn := make(map[string]interface{})

		if v, _ := attachmentsMap["add-link-to-log-attachment"]; v != nil {
			attachmentsMapToReturn["add_link_to_log_attachment"] = v.(bool)
		}
		if v, _ := attachmentsMap["add-link-to-log-details"]; v != nil {
			attachmentsMapToReturn["add_link_to_log_details"] = v.(bool)
		}
		if v, _ := attachmentsMap["add-log-attachment-id"]; v != nil {
			attachmentsMapToReturn["add_log_attachment_id"] = v.(bool)
		}

		_ = d.Set("attachments", []interface{}{attachmentsMapToReturn})
	} else {
		_ = d.Set("attachments", nil)
	}

	if logExporter["data-manipulation"] != nil {

		dataManipulationMap := logExporter["data-manipulation"].(map[string]interface{})

		dataManipulationMapToReturn := make(map[string]interface{})

		if v, _ := dataManipulationMap["aggregate-log-updates"]; v != nil {
			dataManipulationMapToReturn["aggregate_log_updates"] = v.(bool)
		}
		if v, _ := dataManipulationMap["format"]; v != nil {
			dataManipulationMapToReturn["format"] = v
		}

		_ = d.Set("data_manipulation", []interface{}{dataManipulationMapToReturn})
	} else {
		_ = d.Set("data_manipulation", nil)
	}

	if v := logExporter["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := logExporter["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if logExporter["tags"] != nil {
		tagsJson, ok := logExporter["tags"].([]interface{})
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

	return nil
}
