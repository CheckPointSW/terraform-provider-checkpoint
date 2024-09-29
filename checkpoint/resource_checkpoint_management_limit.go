package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementLimit() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLimit,
		Read:   readManagementLimit,
		Update: updateManagementLimit,
		Delete: deleteManagementLimit,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"enable_download": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable throughput limit for downloads from the internet to the organization.",
				Default:     false,
			},
			"download_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Rate for the maximum permitted bandwidth.",
				Default:     0,
			},
			"download_unit": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Unit for the maximum permitted bandwidth.",
				Default:     "mbps",
			},
			"enable_upload": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable throughput limit for uploads from the organization to the internet.",
				Default:     false,
			},
			"upload_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The Rate for the maximum permitted bandwidth.",
				Default:     0,
			},
			"upload_unit": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Unit for the maximum permitted bandwidth.",
				Default:     "mbps",
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

func createManagementLimit(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	limit := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		limit["name"] = v.(string)
	}

	if v, ok := d.GetOkExists("enable_download"); ok {
		limit["enable-download"] = v.(bool)
	}

	if v, ok := d.GetOk("download_rate"); ok {
		limit["download-rate"] = v.(int)
	}

	if v, ok := d.GetOk("download_unit"); ok {
		limit["download-unit"] = v.(string)
	}

	if v, ok := d.GetOkExists("enable_upload"); ok {
		limit["enable-upload"] = v.(bool)
	}

	if v, ok := d.GetOk("upload_rate"); ok {
		limit["upload-rate"] = v.(int)
	}

	if v, ok := d.GetOk("upload_unit"); ok {
		limit["upload-unit"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		limit["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		limit["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		limit["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		limit["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		limit["ignore-errors"] = v.(bool)
	}

	log.Println("Create Limit - Map = ", limit)

	addLimitRes, err := client.ApiCall("add-limit", limit, client.GetSessionID(), true, false)
	if err != nil || !addLimitRes.Success {
		if addLimitRes.ErrorMsg != "" {
			return fmt.Errorf(addLimitRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addLimitRes.GetData()["uid"].(string))

	return readManagementLimit(d, m)
}

func readManagementLimit(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showLimitRes, err := client.ApiCall("show-limit", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLimitRes.Success {
		if objectNotFound(showLimitRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLimitRes.ErrorMsg)
	}

	limit := showLimitRes.GetData()

	log.Println("Read Limit - Show JSON = ", limit)

	if v := limit["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := limit["enable-download"]; v != nil {
		_ = d.Set("enable_download", v)
	}

	if v := limit["download-rate"]; v != nil {
		_ = d.Set("download_rate", v)
	}

	if v := limit["download-unit"]; v != nil {
		_ = d.Set("download_unit", v)
	}

	if v := limit["enable-upload"]; v != nil {
		_ = d.Set("enable_upload", v)
	}

	if v := limit["upload-rate"]; v != nil {
		_ = d.Set("upload_rate", v)
	}

	if v := limit["upload-unit"]; v != nil {
		_ = d.Set("upload_unit", v)
	}

	if limit["tags"] != nil {
		tagsJson, ok := limit["tags"].([]interface{})
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

	if v := limit["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := limit["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := limit["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := limit["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementLimit(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	limit := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		limit["name"] = oldName
		limit["new-name"] = newName
	} else {
		limit["name"] = d.Get("name")
	}

	if v, ok := d.GetOkExists("enable_download"); ok {
		limit["enable-download"] = v.(bool)
	}

	if ok := d.HasChange("download_rate"); ok {
		limit["download-rate"] = d.Get("download_rate")
	}

	if ok := d.HasChange("download_unit"); ok {
		limit["download-unit"] = d.Get("download_unit")
	}

	if v, ok := d.GetOkExists("enable_upload"); ok {
		limit["enable-upload"] = v.(bool)
	}

	if ok := d.HasChange("upload_rate"); ok {
		limit["upload-rate"] = d.Get("upload_rate")
	}

	if ok := d.HasChange("upload_unit"); ok {
		limit["upload-unit"] = d.Get("upload_unit")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			limit["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			limit["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		limit["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		limit["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		limit["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		limit["ignore-errors"] = v.(bool)
	}

	log.Println("Update Limit - Map = ", limit)

	updateLimitRes, err := client.ApiCall("set-limit", limit, client.GetSessionID(), true, false)
	if err != nil || !updateLimitRes.Success {
		if updateLimitRes.ErrorMsg != "" {
			return fmt.Errorf(updateLimitRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementLimit(d, m)
}

func deleteManagementLimit(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	limitPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete Limit")

	deleteLimitRes, err := client.ApiCall("delete-limit", limitPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteLimitRes.Success {
		if deleteLimitRes.ErrorMsg != "" {
			return fmt.Errorf(deleteLimitRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
