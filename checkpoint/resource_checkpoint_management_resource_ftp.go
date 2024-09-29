package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementResourceFtp() *schema.Resource {
	return &schema.Resource{
		Create: createManagementResourceFtp,
		Read:   readManagementResourceFtp,
		Update: updateManagementResourceFtp,
		Delete: deleteManagementResourceFtp,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"resource_matching_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "GET allows Downloads from the server to the client. PUT allows Uploads from the client to the server.",
			},
			"exception_track": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The UID or Name of the exception track to be used to log actions taken as a result of a match on the resource.",
				Default:     "None",
			},
			"resources_path": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Refers to a location on the FTP server.",
				Default:     "*",
			},

			"cvp": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Configure CVP inspection on mail messages.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cvp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Select to enable the Content Vectoring Protocol.",
							Default:     false,
						},
						"server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application.",
						},
						"allowed_to_modify_content": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Configures the CVP server to inspect but not modify content.",
							Default:     true,
						},
						"reply_order": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Designates when the CVP server returns data to the Security Gateway security server.",
							Default:     "return_data_after_content_is_approved",
						},
					},
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

func createManagementResourceFtp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	resourceFtp := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		resourceFtp["name"] = v.(string)
	}

	if v, ok := d.GetOk("resource_matching_method"); ok {
		resourceFtp["resource-matching-method"] = v.(string)
	}

	if v, ok := d.GetOk("exception_track"); ok {
		resourceFtp["exception-track"] = v.(string)
	}

	if v, ok := d.GetOk("resources_path"); ok {
		resourceFtp["resources-path"] = v.(string)
	}

	if v, ok := d.GetOk("cvp"); ok {

		res := make(map[string]interface{})

		v := v.([]interface{})

		cvpMap := v[0].(map[string]interface{})

		if v := cvpMap["enable_cvp"]; v != nil {
			res["enable-cvp"] = v
		}
		if v := cvpMap["server"]; v != nil {
			if len(v.(string)) > 0 {
				res["server"] = v
			}
		}
		if v := cvpMap["allowed_to_modify_content"]; v != nil {
			res["allowed-to-modify-content"] = v
		}
		if v := cvpMap["reply_order"]; v != nil {
			res["reply-order"] = v
		}

		resourceFtp["cvp"] = res
	}

	if v, ok := d.GetOk("tags"); ok {
		resourceFtp["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		resourceFtp["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		resourceFtp["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceFtp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceFtp["ignore-errors"] = v.(bool)
	}

	log.Println("Create ResourceFtp - Map = ", resourceFtp)

	addResourceFtpRes, err := client.ApiCall("add-resource-ftp", resourceFtp, client.GetSessionID(), true, false)
	if err != nil || !addResourceFtpRes.Success {
		if addResourceFtpRes.ErrorMsg != "" {
			return fmt.Errorf(addResourceFtpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addResourceFtpRes.GetData()["uid"].(string))

	return readManagementResourceFtp(d, m)
}

func readManagementResourceFtp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showResourceFtpRes, err := client.ApiCall("show-resource-ftp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceFtpRes.Success {
		if objectNotFound(showResourceFtpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showResourceFtpRes.ErrorMsg)
	}

	resourceFtp := showResourceFtpRes.GetData()

	log.Println("Read ResourceFtp - Show JSON = ", resourceFtp)

	if v := resourceFtp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceFtp["exception-track"]; v != nil {

		objMap := v.(map[string]interface{})
		if v := objMap["name"]; v != nil {
			_ = d.Set("exception_track", v)
		}
	}

	if v := resourceFtp["resources-path"]; v != nil {
		_ = d.Set("resources_path", v)
	}

	if v := resourceFtp["resource-matching-method"]; v != nil {
		_ = d.Set("resource_matching_method", v)
	}

	if resourceFtp["cvp"] != nil {

		cvpMap := resourceFtp["cvp"].(map[string]interface{})

		cvpMapToReturn := make(map[string]interface{})

		if v, _ := cvpMap["enable-cvp"]; v != nil {
			cvpMapToReturn["enable_cvp"] = v
		}
		if v, _ := cvpMap["server"]; v != nil {

			objMap := v.(map[string]interface{})

			if v := objMap["name"]; v != nil {
				cvpMapToReturn["server"] = v
			}
		}
		if v, _ := cvpMap["cvp-server-is-allowed-to-modify-content"]; v != nil {
			cvpMapToReturn["allowed_to_modify_content"] = v
		}
		if v, _ := cvpMap["reply-order"]; v != nil {
			cvpMapToReturn["reply_order"] = v
		}
		_ = d.Set("cvp", []interface{}{cvpMapToReturn})
	} else {
		_ = d.Set("cvp", nil)
	}

	if resourceFtp["tags"] != nil {
		tagsJson, ok := resourceFtp["tags"].([]interface{})
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

	if v := resourceFtp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceFtp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := resourceFtp["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := resourceFtp["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementResourceFtp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	resourceFtp := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		resourceFtp["name"] = oldName
		resourceFtp["new-name"] = newName
	} else {
		resourceFtp["name"] = d.Get("name")
	}

	if ok := d.HasChange("exception_track"); ok {
		resourceFtp["exception-track"] = d.Get("exception_track")
	}

	if ok := d.HasChange("resources_path"); ok {
		resourceFtp["resources-path"] = d.Get("resources_path")
	}

	if ok := d.HasChange("resource_matching_method"); ok {
		resourceFtp["resource-matching-method"] = d.Get("resource_matching_method")
	}

	if d.HasChange("cvp") {

		if v, ok := d.GetOk("cvp"); ok {

			res := make(map[string]interface{})

			v := v.([]interface{})

			cvpMap := v[0].(map[string]interface{})

			if v := cvpMap["enable_cvp"]; v != nil {
				res["enable-cvp"] = v
			}
			if v := cvpMap["server"]; v != nil {
				if len(v.(string)) > 0 {
					res["server"] = v
				}

			}
			if v := cvpMap["allowed_to_modify_content"]; v != nil {
				res["allowed-to-modify-content"] = v
			}
			if v := cvpMap["reply_order"]; v != nil {
				res["reply-order"] = v
			}

			resourceFtp["cvp"] = res
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			resourceFtp["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			resourceFtp["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		resourceFtp["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		resourceFtp["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceFtp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceFtp["ignore-errors"] = v.(bool)
	}

	log.Println("Update ResourceFtp - Map = ", resourceFtp)

	updateResourceFtpRes, err := client.ApiCall("set-resource-ftp", resourceFtp, client.GetSessionID(), true, false)
	if err != nil || !updateResourceFtpRes.Success {
		if updateResourceFtpRes.ErrorMsg != "" {
			return fmt.Errorf(updateResourceFtpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementResourceFtp(d, m)
}

func deleteManagementResourceFtp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	resourceFtpPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ResourceFtp")

	deleteResourceFtpRes, err := client.ApiCall("delete-resource-ftp", resourceFtpPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteResourceFtpRes.Success {
		if deleteResourceFtpRes.ErrorMsg != "" {
			return fmt.Errorf(deleteResourceFtpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
