package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func resourceManagementDlpNextDataType() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDlpNextDataType,
		Read:   readManagementDlpNextDataType,
		Update: updateManagementDlpNextDataType,
		Delete: deleteManagementDlpNextDataType,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"external_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "DLP Next Data Type unique identifier in Infinity Portal.",
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
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the data type.",
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

func createManagementDlpNextDataType(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dlpNextDataType := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		dlpNextDataType["name"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		dlpNextDataType["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dlpNextDataType["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dlpNextDataType["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dlpNextDataType["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dlpNextDataType["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOk("external_id"); ok {
		dlpNextDataType["external-id"] = v.(string)
	}

	log.Println("Create DlpNextDataType - Map = ", dlpNextDataType)

	addDlpNextDataTypeRes, err := client.ApiCall("add-dlp-next-data-type", dlpNextDataType, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addDlpNextDataTypeRes.Success {
		if addDlpNextDataTypeRes.ErrorMsg != "" {
			return fmt.Errorf(addDlpNextDataTypeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDlpNextDataTypeRes.GetData()["uid"].(string))

	return readManagementDlpNextDataType(d, m)
}

func readManagementDlpNextDataType(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showDlpNextDataTypeRes, err := client.ApiCall("show-dlp-next-data-type", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDlpNextDataTypeRes.Success {
		if objectNotFound(showDlpNextDataTypeRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDlpNextDataTypeRes.ErrorMsg)
	}

	dlpNextDataType := showDlpNextDataTypeRes.GetData()

	log.Println("Read DlpNextDataType - Show JSON = ", dlpNextDataType)

	if v := dlpNextDataType["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if dlpNextDataType["tags"] != nil {
		tagsJson, ok := dlpNextDataType["tags"].([]interface{})
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

	if v := dlpNextDataType["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dlpNextDataType["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dlpNextDataType["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := dlpNextDataType["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dlpNextDataType["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	if v := dlpNextDataType["external-id"]; v != nil {
		_ = d.Set("external_id", v)
	}

	return nil

}

func updateManagementDlpNextDataType(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dlpNextDataType := make(map[string]interface{})

	dlpNextDataType["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			dlpNextDataType["new-name"] = v.(string)
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dlpNextDataType["tags"] = v.(*schema.Set).List()
		}
	}

	if ok := d.HasChange("color"); ok {
		dlpNextDataType["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dlpNextDataType["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dlpNextDataType["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dlpNextDataType["ignore-errors"] = v.(bool)
	}

	if ok := d.HasChange("external_id"); ok {
		dlpNextDataType["external-id"] = d.Get("external_id")
	}

	log.Println("Update DlpNextDataType - Map = ", dlpNextDataType)

	updateDlpNextDataTypeRes, err := client.ApiCall("set-dlp-next-data-type", dlpNextDataType, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateDlpNextDataTypeRes.Success {
		if updateDlpNextDataTypeRes.ErrorMsg != "" {
			return fmt.Errorf(updateDlpNextDataTypeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDlpNextDataType(d, m)
}

func deleteManagementDlpNextDataType(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dlpNextDataTypePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete DlpNextDataType")

	deleteDlpNextDataTypeRes, err := client.ApiCall("delete-dlp-next-data-type", dlpNextDataTypePayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteDlpNextDataTypeRes.Success {
		if deleteDlpNextDataTypeRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDlpNextDataTypeRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
