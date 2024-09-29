package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementAddCustomTrustedCaCertificate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAddCustomTrustedCaCertificate,
		Read:   readManagementAddCustomTrustedCaCertificate,
		Delete: deleteManagementAddCustomTrustedCaCertificate,
		Schema: map[string]*schema.Schema{
			"base64_certificate": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Certificate file encoded in base64.<br/>Valid file formats: x509.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object name.",
			},
			"added_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "By whom the certificate was added.",
			},
			"issued_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trusted CA certificate issued by.",
			},
			"issued_to": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trusted CA certificate issued to.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"valid_from": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Trusted CA certificate valid from date.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.\nTrusted CA certificate valid from date.",
						},
					},
				},
			},
			"valid_to": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Trusted CA certificate valid to date.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"iso_8601": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Date and time represented in international ISO 8601 format",
						},
						"posix": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.\nTrusted CA certificate valid from date.",
						},
					},
				},
			},
		},
	}
}

func createManagementAddCustomTrustedCaCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("base64_certificate"); ok {
		payload["base64-certificate"] = v.(string)
	}

	AddCustomTrustedCaCertificateRes, _ := client.ApiCall("add-custom-trusted-ca-certificate", payload, client.GetSessionID(), true, false)
	if !AddCustomTrustedCaCertificateRes.Success {
		return fmt.Errorf(AddCustomTrustedCaCertificateRes.ErrorMsg)
	}

	customTrustedCaCertificateObj := AddCustomTrustedCaCertificateRes.GetData()

	if v := customTrustedCaCertificateObj["uid"]; v != nil {
		d.SetId(v.(string))
		d.Set("uid", v)
	}

	if v := customTrustedCaCertificateObj["name"]; v != nil {
		d.Set("name", v)
	}

	if v := customTrustedCaCertificateObj["added-by"]; v != nil {
		_ = d.Set("added_by", v)
	}
	if v := customTrustedCaCertificateObj["base64-certificate"]; v != nil {
		_ = d.Set("base64_certificate", cleanseCertificate(v.(string)))
	}
	if v := customTrustedCaCertificateObj["issued-by"]; v != nil {
		_ = d.Set("issued_by", v)
	}
	if v := customTrustedCaCertificateObj["issued-to"]; v != nil {
		_ = d.Set("issued_to", v)
	}
	if customTrustedCaCertificateObj["tags"] != nil {
		tagsJson, ok := customTrustedCaCertificateObj["tags"].([]interface{})
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

	if v := customTrustedCaCertificateObj["valid-from"]; v != nil {

		localMap := v.(map[string]interface{})

		validFrom := make(map[string]interface{})

		if v := localMap["iso-8601"]; v != nil {
			validFrom["iso_8601"] = v
		}
		if v := localMap["posix"]; v != nil {
			validFrom["posix"] = v
		}
		_ = d.Set("valid_from", []interface{}{validFrom})
	}
	if v := customTrustedCaCertificateObj["valid-to"]; v != nil {

		localMap := v.(map[string]interface{})

		validTo := make(map[string]interface{})

		if v := localMap["iso-8601"]; v != nil {
			validTo["iso_8601"] = v
		}
		if v := localMap["posix"]; v != nil {
			validTo["posix"] = v
		}
		_ = d.Set("valid_to", []interface{}{validTo})
	}

	return readManagementAddCustomTrustedCaCertificate(d, m)
}

func readManagementAddCustomTrustedCaCertificate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementAddCustomTrustedCaCertificate(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
