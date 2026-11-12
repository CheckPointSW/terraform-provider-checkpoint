package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementThreatEmulationImage() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatEmulationImageRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Image name.",
			},
			"image_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Image id.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Image description.",
			},
			"display_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Image display name.",
			},
			"image_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Image type.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object icon.",
			},
		},
	}
}

func dataSourceManagementThreatEmulationImageRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("image_id"); ok {
		payload["image-id"] = v.(string)
	}

	showThreatEmulationImageRes, err := client.ApiCall("show-threat-emulation-image", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatEmulationImageRes.Success {
		if objectNotFound(showThreatEmulationImageRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatEmulationImageRes.ErrorMsg)
	}

	threatEmulationImage := showThreatEmulationImageRes.GetData()

	log.Println("Read ThreatEmulationImage - Show JSON = ", threatEmulationImage)

	if v := threatEmulationImage["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := threatEmulationImage["display-name"]; v != nil {
		_ = d.Set("display_name", v)
	}

	if v := threatEmulationImage["image-id"]; v != nil {
		d.SetId(v.(string))
		_ = d.Set("image_id", v)
	}

	if v := threatEmulationImage["image-type"]; v != nil {
		_ = d.Set("image_type", v)
	}

	if v := threatEmulationImage["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if threatEmulationImage["tags"] != nil {
		tagsJson, ok := threatEmulationImage["tags"].([]interface{})
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

	if v := threatEmulationImage["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	return nil

}
