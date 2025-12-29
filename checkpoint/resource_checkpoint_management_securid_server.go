package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementSecuridServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSecuridServer,
		Read:   readManagementSecuridServer,
		Update: updateManagementSecuridServer,
		Delete: deleteManagementSecuridServer,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"config_file_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Configuration file name. Required only when 'base64-config-file-content' is not empty.",
			},
			"base64_config_file_content": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Base64 encoded configuration file for authentication.",
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

func createManagementSecuridServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	securidServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		securidServer["name"] = v.(string)
	}

	if v, ok := d.GetOk("config_file_name"); ok {
		securidServer["config-file-name"] = v.(string)
	}

	if v, ok := d.GetOk("base64_config_file_content"); ok {
		securidServer["base64-config-file-content"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		securidServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		securidServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		securidServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		securidServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		securidServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create SecuridServer - Map = ", securidServer)

	addSecuridServerRes, err := client.ApiCallSimple("add-securid-server", securidServer)
	if err != nil || !addSecuridServerRes.Success {
		if addSecuridServerRes.ErrorMsg != "" {
			return fmt.Errorf(addSecuridServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addSecuridServerRes.GetData()["uid"].(string))

	return readManagementSecuridServer(d, m)
}

func readManagementSecuridServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showSecuridServerRes, err := client.ApiCallSimple("show-securid-server", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSecuridServerRes.Success {
		if objectNotFound(showSecuridServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showSecuridServerRes.ErrorMsg)
	}

	securidServer := showSecuridServerRes.GetData()

	log.Println("Read SecuridServer - Show JSON = ", securidServer)

	if v := securidServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := securidServer["config-file-name"]; v != nil {
		_ = d.Set("config_file_name", v)
	}

	if v := securidServer["base64-config-file-content"]; v != nil {
		_ = d.Set("base64_config_file_content", v)
	}

	if securidServer["tags"] != nil {
		tagsJson, ok := securidServer["tags"].([]interface{})
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

	if v := securidServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := securidServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := securidServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := securidServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementSecuridServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	securidServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		securidServer["name"] = oldName
		securidServer["new-name"] = newName
	} else {
		securidServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("config_file_name"); ok {
		securidServer["config-file-name"] = d.Get("config_file_name")
	}

	if ok := d.HasChange("base64_config_file_content"); ok {
		securidServer["base64-config-file-content"] = d.Get("base64_config_file_content")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			securidServer["tags"] = v.(*schema.Set).List()
		}
	}

	if ok := d.HasChange("color"); ok {
		securidServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		securidServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		securidServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		securidServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update SecuridServer - Map = ", securidServer)

	updateSecuridServerRes, err := client.ApiCallSimple("set-securid-server", securidServer)
	if err != nil || !updateSecuridServerRes.Success {
		if updateSecuridServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateSecuridServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementSecuridServer(d, m)
}

func deleteManagementSecuridServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	securidServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		securidServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		securidServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete SecuridServer")

	deleteSecuridServerRes, err := client.ApiCallSimple("delete-securid-server", securidServerPayload)
	if err != nil || !deleteSecuridServerRes.Success {
		if deleteSecuridServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteSecuridServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
