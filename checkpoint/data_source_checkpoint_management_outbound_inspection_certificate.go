package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementOutboundInspectionCertificate() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementOutboundInspectionCertificateRead,		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"issued_by": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The DN (Distinguished Name) of the certificate.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Certificate file encoded in base64.",
			},
			"valid_from": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The date, from which the certificate is valid. Format: YYYY-MM-DD.",
			},
			"valid_to": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "The certificate expiration date. Format: YYYY-MM-DD.",
			},
		},
	}
}

func dataSourceManagementOutboundInspectionCertificateRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})

	showOutboundInspectionCertificateRes, err := client.ApiCall("show-idp-default-assignment", payload, client.GetSessionID(), true, false)
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

	idpDefaultAssignment := showOutboundInspectionCertificateRes.GetData()

	log.Println("Read OutboundInspectionCertificate - Show JSON = ", idpDefaultAssignment)

	if v := idpDefaultAssignment["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := idpDefaultAssignment["issued-by"]; v != nil {
		_ = d.Set("issued_by", v)
	}

	if v := idpDefaultAssignment["base64-certificate"]; v != nil {
		_ = d.Set("base64-certificate", v)
	}

	if v := idpDefaultAssignment["valid-from"]; v != nil {
		_ = d.Set("valid_from", v)
	}

	if v := idpDefaultAssignment["valid-to"]; v != nil {
		_ = d.Set("valid_to", v)
	}

	return nil

}
