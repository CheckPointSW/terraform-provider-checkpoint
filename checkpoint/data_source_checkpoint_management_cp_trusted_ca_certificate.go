package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCpTrustedCaCertificate() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementCpTrustedCaCertificateRead,

		Schema: map[string]*schema.Schema{

			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate Object uid.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate Object name.",
			},
			"added_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "By whom the certificate was added.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate file encoded in base64.",
			},
			"base64_public_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Public Certificate file encoded in base64 (pem format).",
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
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Indicates whether the trusted CP CA certificate is enabled/disabled.",
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

func dataSourceManagementCpTrustedCaCertificateRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	cpTrustedCaCertificateObjRes, err := client.ApiCall("show-cp-trusted-ca-certificate", payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cpTrustedCaCertificateObjRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(cpTrustedCaCertificateObjRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(cpTrustedCaCertificateObjRes.ErrorMsg)
	}

	cpTrustedCaCertificateObj := cpTrustedCaCertificateObjRes.GetData()

	log.Println("Read CP Trusted CA Certificate Object - Show JSON = ", cpTrustedCaCertificateObj)

	if v := cpTrustedCaCertificateObj["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := cpTrustedCaCertificateObj["name"]; v != nil {
		_ = d.Set("name", v)
	}
	if v := cpTrustedCaCertificateObj["added-by"]; v != nil {
		_ = d.Set("added_by", v)
	}
	if v := cpTrustedCaCertificateObj["base64-certificate"]; v != nil {
		_ = d.Set("base64_certificate", v)
	}
	if v := cpTrustedCaCertificateObj["base64-public-certificate"]; v != nil {
		_ = d.Set("base64_public_certificate", v)
	}
	if v := cpTrustedCaCertificateObj["issued-by"]; v != nil {
		_ = d.Set("issued_by", v)
	}
	if v := cpTrustedCaCertificateObj["issued-to"]; v != nil {
		_ = d.Set("issued_to", v)
	}
	if v := cpTrustedCaCertificateObj["status"]; v != nil {
		_ = d.Set("status", v)
	}
	if cpTrustedCaCertificateObj["tags"] != nil {
		tagsJson, ok := cpTrustedCaCertificateObj["tags"].([]interface{})
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

	if v := cpTrustedCaCertificateObj["status"]; v != nil {
		_ = d.Set("status", v)
	}
	if v := cpTrustedCaCertificateObj["valid-from"]; v != nil {

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
	if v := cpTrustedCaCertificateObj["valid-to"]; v != nil {

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
	return nil
}
