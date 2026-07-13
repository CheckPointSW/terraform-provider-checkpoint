package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementSetThreatExtractionFileType() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetThreatExtractionFileType,
		Read:   readManagementSetThreatExtractionFileType,
		Delete: deleteManagementSetThreatExtractionFileType,
		Schema: map[string]*schema.Schema{
			"file_type_id": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "File type id.",
			},
			"file_type": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "File type extension.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Enable support for Threat Extraction.",
			},
			"description": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "File type description.",
			},
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "File type icon.",
			},
			"tags": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Computed:    true,
				Description: "Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.",
			},
		},
	}
}

func createManagementSetThreatExtractionFileType(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("file_type_id"); ok {
		payload["file-type-id"] = v.(string)
	}

	if v, ok := d.GetOk("file_type"); ok {
		payload["file-type"] = v.(string)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		payload["enabled"] = v.(bool)
	}

	SetThreatExtractionFileTypeRes, err := client.ApiCall("set-threat-extraction-file-type", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetThreatExtractionFileTypeRes.Success {
		return fmt.Errorf(SetThreatExtractionFileTypeRes.ErrorMsg)
	}

	d.SetId("set-threat-extraction-file-type-" + acctest.RandString(10))
	if v := SetThreatExtractionFileTypeRes.GetData()["description"]; v != nil {
		_ = d.Set("description", v)
	}
	if v := SetThreatExtractionFileTypeRes.GetData()["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}
	if v := SetThreatExtractionFileTypeRes.GetData()["file-type"]; v != nil {
		_ = d.Set("file_type", v)
	}
	if v := SetThreatExtractionFileTypeRes.GetData()["file-type-id"]; v != nil {
		_ = d.Set("file_type_id", v)
	}
	if v := SetThreatExtractionFileTypeRes.GetData()["icon"]; v != nil {
		_ = d.Set("icon", v)
	}
	if v := SetThreatExtractionFileTypeRes.GetData()["domain"]; v != nil {
		_ = d.Set("domain", v)
	}
	if v := SetThreatExtractionFileTypeRes.GetData()["tags"]; v != nil {
		_ = d.Set("tags", v)
	}
	return nil
}

func readManagementSetThreatExtractionFileType(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetThreatExtractionFileType(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
