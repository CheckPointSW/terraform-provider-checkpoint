package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementSetThreatEmulationFileType() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetThreatEmulationFileType,
		Read:   readManagementSetThreatEmulationFileType,
		Delete: deleteManagementSetThreatEmulationFileType,
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
				Description: "Enable support for Threat Emulation.",
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
			"supported_platforms": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Supported platforms for Threat Emulation.",
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

func createManagementSetThreatEmulationFileType(d *schema.ResourceData, m interface{}) error {

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

	SetThreatEmulationFileTypeRes, err := client.ApiCall("set-threat-emulation-file-type", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetThreatEmulationFileTypeRes.Success {
		return fmt.Errorf(SetThreatEmulationFileTypeRes.ErrorMsg)
	}

	d.SetId("set-threat-emulation-file-type-" + acctest.RandString(10))
	if v := SetThreatEmulationFileTypeRes.GetData()["description"]; v != nil {
		_ = d.Set("description", v)
	}
	if v := SetThreatEmulationFileTypeRes.GetData()["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}
	if v := SetThreatEmulationFileTypeRes.GetData()["file-type"]; v != nil {
		_ = d.Set("file_type", v)
	}
	if v := SetThreatEmulationFileTypeRes.GetData()["file-type-id"]; v != nil {
		_ = d.Set("file_type_id", v)
	}
	if v := SetThreatEmulationFileTypeRes.GetData()["icon"]; v != nil {
		_ = d.Set("icon", v)
	}
	if v := SetThreatEmulationFileTypeRes.GetData()["supported-platforms"]; v != nil {
		_ = d.Set("supported_platforms", v)
	}
	if v := SetThreatEmulationFileTypeRes.GetData()["tags"]; v != nil {
		_ = d.Set("tags", v)
	}
	return nil
}

func readManagementSetThreatEmulationFileType(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetThreatEmulationFileType(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
