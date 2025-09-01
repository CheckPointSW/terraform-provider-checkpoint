package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementSecuridServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSecuridServerRead,
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
			"config_file_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configuration file name. Required only when 'base64-config-file-content' is not empty.",
			},
			"base64_config_file_content": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Base64 encoded configuration file for authentication.",
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

func dataSourceManagementSecuridServerRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showSecuridServerRes, err := client.ApiCallSimple("show-securid-server", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSecuridServerRes.Success {
		return fmt.Errorf(showSecuridServerRes.ErrorMsg)
	}

	securidServer := showSecuridServerRes.GetData()

	log.Println("Read SecuridServer - Show JSON = ", securidServer)

	if v := securidServer["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

	return nil

}
