package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementOutboundInspectionCertificate() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementOutboundInspectionCertificateRead,
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
			"issued_by": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DN (Distinguished Name) of the certificate.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Password (encoded in Base64 with padding) for the certificate file.",
			},
			"base64_public_certificate": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Password (encoded in Base64 with padding) for the certificate file.",
			},
			"valid_from": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date, from which the certificate is valid. Format: YYYY-MM-DD.",
			},
			"valid_to": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The certificate expiration date. Format: YYYY-MM-DD.",
			},
			"is_default": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Is the certificate the default certificate.",
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
func dataSourceManagementOutboundInspectionCertificateRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showOutboundInspectionCertificateRes, err := client.ApiCall("show-outbound-inspection-certificate", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showOutboundInspectionCertificateRes.Success {
		if objectNotFound(showOutboundInspectionCertificateRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showOutboundInspectionCertificateRes.ErrorMsg)
	}

	outboundInspectionCertificate := showOutboundInspectionCertificateRes.GetData()

	log.Println("Read OutboundInspectionCertificate - Show JSON = ", outboundInspectionCertificate)

	if v := outboundInspectionCertificate["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := outboundInspectionCertificate["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := outboundInspectionCertificate["issued-by"]; v != nil {
		_ = d.Set("issued_by", v)
	}

	if v := outboundInspectionCertificate["base64-certificate"]; v != nil {
		_ = d.Set("base64_certificate", cleanseCertificate(v.(string)))
	}

	if v := outboundInspectionCertificate["base64-public-certificate"]; v != nil {
		_ = d.Set("base64_public_certificate", cleanseCertificate(v.(string)))
	}

	if v := outboundInspectionCertificate["valid-from"]; v != nil {
		_ = d.Set("valid_from", v)
	}

	if v := outboundInspectionCertificate["valid-to"]; v != nil {
		_ = d.Set("valid_to", v)
	}

	if v := outboundInspectionCertificate["is-default"]; v != nil {
		_ = d.Set("is_default", v)
	}

	if outboundInspectionCertificate["tags"] != nil {
		tagsJson, ok := outboundInspectionCertificate["tags"].([]interface{})
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

	if v := outboundInspectionCertificate["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := outboundInspectionCertificate["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}
