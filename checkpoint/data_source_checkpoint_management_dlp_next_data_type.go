package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementDlpNextDataType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementDlpNextDataTypeRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DLP Next Data Type description in Infinity Portal.",
			},
			"external_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DLP Next Data Type unique identifier in Infinity Portal.",
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
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object icon.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementDlpNextDataTypeRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	showDlpNextDataTypeRes, err := client.ApiCall("show-dlp-next-data-type", payload, client.GetSessionID(), true, false)
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

	if v := dlpNextDataType["uid"]; v != nil {
		d.SetId(v.(string))
		_ = d.Set("uid", v)
	}

	if v := dlpNextDataType["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := dlpNextDataType["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := dlpNextDataType["external-id"]; v != nil {
		_ = d.Set("external_id", v)
	}

	if v := dlpNextDataType["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dlpNextDataType["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dlpNextDataType["icon"]; v != nil {
		_ = d.Set("icon", v)
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

	return nil

}
