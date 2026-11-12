package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strconv"
)

func resourceManagementSetThreatExtractionFileTypes() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetThreatExtractionFileTypes,
		Read:   readManagementSetThreatExtractionFileTypes,
		Delete: deleteManagementSetThreatExtractionFileTypes,
		Schema: map[string]*schema.Schema{
			"file_types": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of Threat Extraction file type updates. Each entry sets 'enabled' on the file type identified by 'file-type-id' or 'file-type'.",
				ForceNew:    true,
				Elem: &schema.Resource{
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
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable support for Threat Extraction.",
						},
					},
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementSetThreatExtractionFileTypes(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("file_types"); ok {

		fileTypesList := v.([]interface{})

		if len(fileTypesList) > 0 {

			var fileTypesPayload []map[string]interface{}

			for i := range fileTypesList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("file_types." + strconv.Itoa(i) + ".file_type_id"); ok {
					Payload["file-type-id"] = v.(string)
				}
				if v, ok := d.GetOk("file_types." + strconv.Itoa(i) + ".file_type"); ok {
					Payload["file-type"] = v.(string)
				}
				if v, ok := d.GetOkExists("file_types." + strconv.Itoa(i) + ".enabled"); ok {
					Payload["enabled"] = v.(bool)
				}
				fileTypesPayload = append(fileTypesPayload, Payload)
			}
			payload["file-types"] = fileTypesPayload
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	SetThreatExtractionFileTypesRes, err := client.ApiCall("set-threat-extraction-file-types", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !SetThreatExtractionFileTypesRes.Success {
		return fmt.Errorf(SetThreatExtractionFileTypesRes.ErrorMsg)
	}

	d.SetId("set-threat-extraction-file-types-" + acctest.RandString(10))
	if v := SetThreatExtractionFileTypesRes.GetData()["file-types"]; v != nil {
		_ = d.Set("file_types", v)
	}
	return nil
}

func readManagementSetThreatExtractionFileTypes(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementSetThreatExtractionFileTypes(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
