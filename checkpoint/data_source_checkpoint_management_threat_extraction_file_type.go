package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementThreatExtractionFileType() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementThreatExtractionFileTypeRead,
		Schema: map[string]*schema.Schema{
			"file_type_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File type id.",
			},
			"file_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "File type extension.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "File type description.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable support for Threat Extraction.",
			},
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "File type icon.",
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

func dataSourceManagementThreatExtractionFileTypeRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("file_type_id"); ok {
		payload["file-type-id"] = v.(string)
	}

	if v, ok := d.GetOk("file_type"); ok {
		payload["file-type"] = v.(string)
	}

	showThreatExtractionFileTypeRes, err := client.ApiCall("show-threat-extraction-file-type", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showThreatExtractionFileTypeRes.Success {
		if objectNotFound(showThreatExtractionFileTypeRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showThreatExtractionFileTypeRes.ErrorMsg)
	}

	threatExtractionFileType := showThreatExtractionFileTypeRes.GetData()

	log.Println("Read ThreatExtractionFileType - Show JSON = ", threatExtractionFileType)

	if v := threatExtractionFileType["description"]; v != nil {
		_ = d.Set("description", v)
	}

	if v := threatExtractionFileType["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := threatExtractionFileType["file-type"]; v != nil {
		_ = d.Set("file_type", v)
	}

	if v := threatExtractionFileType["file-type-id"]; v != nil {
		d.SetId(v.(string))
		_ = d.Set("file_type_id", v)
	}

	if v := threatExtractionFileType["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	if threatExtractionFileType["tags"] != nil {
		tagsJson, ok := threatExtractionFileType["tags"].([]interface{})
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
