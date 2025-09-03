package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementLogExporter() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLogExporter,
		Read:   readManagementLogExporter,
		Update: updateManagementLogExporter,
		Delete: deleteManagementLogExporter,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"target_server": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Target server port to which logs are exported.",
			},
			"target_port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Port number of the target server.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Protocol used to send logs to the target server.",
				Default:     "udp",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to enable export.",
				Default:     true,
			},
			"attachments": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Log exporter attachments.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_link_to_log_attachment": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to add link to log attachment in SmartView.",
							Default:     false,
						},
						"add_link_to_log_details": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to add link to log details in SmartView.",
							Default:     false,
						},
						"add_log_attachment_id": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to add log attachment ID.",
							Default:     false,
						},
					},
				},
			},
			"data_manipulation": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Log exporter data manipulation.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregate_log_updates": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to aggregate log updates.",
						},
						"format": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Logs format.",
						},
					},
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createManagementLogExporter(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	logExporter := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		logExporter["name"] = v.(string)
	}

	if v, ok := d.GetOk("target_server"); ok {
		logExporter["target-server"] = v.(string)
	}

	if v, ok := d.GetOk("target_port"); ok {
		logExporter["target-port"] = v.(int)
	}

	if v, ok := d.GetOk("protocol"); ok {
		logExporter["protocol"] = v.(string)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		logExporter["enabled"] = v.(bool)
	}

	if _, ok := d.GetOk("attachments"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOkExists("attachments.0.add_link_to_log_attachment"); ok {
			res["add-link-to-log-attachment"] = v.(bool)
		}
		if v, ok := d.GetOkExists("attachments.0.add_link_to_log_details"); ok {
			res["add-link-to-log-details"] = v.(bool)
		}
		if v, ok := d.GetOkExists("attachments.0.add_log_attachment_id"); ok {
			res["add-log-attachment-id"] = v.(bool)
		}

		logExporter["attachments"] = res
	}

	if _, ok := d.GetOk("data_manipulation"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOkExists("data_manipulation.0.aggregate_log_updates"); ok {
			res["aggregate-log-updates"] = v.(bool)
		}
		if v, ok := d.GetOk("data_manipulation.0.format"); ok {
			res["format"] = v.(string)
		}

		logExporter["data-manipulation"] = res
	}

	if v, ok := d.GetOk("color"); ok {
		logExporter["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		logExporter["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		logExporter["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		logExporter["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		logExporter["ignore-errors"] = v.(bool)
	}

	log.Println("Create LogExporter - Map = ", logExporter)

	addLogExporterRes, err := client.ApiCallSimple("add-log-exporter", logExporter)
	if err != nil || !addLogExporterRes.Success {
		if addLogExporterRes.ErrorMsg != "" {
			return fmt.Errorf(addLogExporterRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addLogExporterRes.GetData()["uid"].(string))

	return readManagementLogExporter(d, m)
}

func readManagementLogExporter(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showLogExporterRes, err := client.ApiCallSimple("show-log-exporter", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLogExporterRes.Success {
		if objectNotFound(showLogExporterRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLogExporterRes.ErrorMsg)
	}

	logExporter := showLogExporterRes.GetData()

	log.Println("Read LogExporter - Show JSON = ", logExporter)

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
			attachmentsMapToReturn["add_link_to_log_attachment"] = v
		}
		if v, _ := attachmentsMap["add-link-to-log-details"]; v != nil {
			attachmentsMapToReturn["add_link_to_log_details"] = v
		}
		if v, _ := attachmentsMap["add-log-attachment-id"]; v != nil {
			attachmentsMapToReturn["add_log_attachment_id"] = v
		}

		_ = d.Set("attachments", []interface{}{attachmentsMapToReturn})
	} else {
		_ = d.Set("attachments", nil)
	}

	if logExporter["data-manipulation"] != nil {

		dataManipulationMap := logExporter["data-manipulation"].(map[string]interface{})

		dataManipulationMapToReturn := make(map[string]interface{})

		if v, _ := dataManipulationMap["aggregate-log-updates"]; v != nil {
			dataManipulationMapToReturn["aggregate_log_updates"] = v
		}
		if v, _ := dataManipulationMap["format"]; v != nil {
			dataManipulationMapToReturn["format"] = v
		}

		_ = d.Set("attachments", []interface{}{dataManipulationMapToReturn})
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

	if v := logExporter["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := logExporter["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementLogExporter(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	logExporter := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		logExporter["name"] = oldName
		logExporter["new-name"] = newName
	} else {
		logExporter["name"] = d.Get("name")
	}

	if ok := d.HasChange("target_server"); ok {
		logExporter["target-server"] = d.Get("target_server")
	}

	if ok := d.HasChange("target_port"); ok {
		logExporter["target-port"] = d.Get("target_port")
	}

	if ok := d.HasChange("protocol"); ok {
		logExporter["protocol"] = d.Get("protocol")
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		logExporter["enabled"] = v.(bool)
	}

	if d.HasChange("attachments") {

		if _, ok := d.GetOk("attachments"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOkExists("attachments.0.add_link_to_log_attachment"); ok {
				res["add-link-to-log-attachment"] = v.(bool)
			}

			if v, ok := d.GetOkExists("attachments.0.add_link_to_log_details"); ok {
				res["add-link-to-log-details"] = v.(bool)
			}

			if v, ok := d.GetOkExists("attachments.0.add_log_attachment_id"); ok {
				res["add-log-attachment-id"] = v.(bool)
			}

			logExporter["attachments"] = res
		}
	}

	if d.HasChange("data_manipulation") {

		if _, ok := d.GetOk("data_manipulation"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOkExists("data_manipulation.0.aggregate_log_updates"); ok {
				res["aggregate-log-updates"] = v.(bool)
			}

			if v, ok := d.GetOk("data_manipulation.0.format"); ok {
				res["format"] = v.(string)
			}

			logExporter["data-manipulation"] = res
		}
	}

	if ok := d.HasChange("color"); ok {
		logExporter["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		logExporter["comments"] = d.Get("comments")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			logExporter["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		logExporter["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		logExporter["ignore-errors"] = v.(bool)
	}

	log.Println("Update LogExporter - Map = ", logExporter)

	updateLogExporterRes, err := client.ApiCallSimple("set-log-exporter", logExporter)
	if err != nil || !updateLogExporterRes.Success {
		if updateLogExporterRes.ErrorMsg != "" {
			return fmt.Errorf(updateLogExporterRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementLogExporter(d, m)
}

func deleteManagementLogExporter(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	logExporterPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete LogExporter")

	deleteLogExporterRes, err := client.ApiCallSimple("delete-log-exporter", logExporterPayload)
	if err != nil || !deleteLogExporterRes.Success {
		if deleteLogExporterRes.ErrorMsg != "" {
			return fmt.Errorf(deleteLogExporterRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
