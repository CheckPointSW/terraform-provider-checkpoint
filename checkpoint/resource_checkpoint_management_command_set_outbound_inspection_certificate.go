package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetOutboundInspectionCertificate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetOutboundInspectionCertificate,
		Read:   readManagementSetOutboundInspectionCertificate,
		Delete: deleteManagementSetOutboundInspectionCertificate,
		Schema: map[string]*schema.Schema{
			"issued_by": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The DN (Distinguished Name) of the certificate.",
			},
			"base64_password": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Password (encoded in Base64 with padding) for the certificate file.",
			},
			"valid_from": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The date, from which the certificate is valid. Format: YYYY-MM-DD.",
			},
			"valid_to": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The certificate expiration date. Format: YYYY-MM-DD.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementSetOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("issued_by"); ok {
		payload["issued-by"] = v.(string)
	}

	if v, ok := d.GetOk("base64_password"); ok {
		payload["base64-password"] = v.(string)
	}

	if v, ok := d.GetOk("valid_from"); ok {
		payload["valid-from"] = v.(string)
	}

	if v, ok := d.GetOk("valid_to"); ok {
		payload["valid-to"] = v.(string)
	}

	SetOutboundInspectionCertificateRes, _ := client.ApiCall("set-outbound-inspection-certificate", payload, client.GetSessionID(), true, false)
	if !SetOutboundInspectionCertificateRes.Success {
		return fmt.Errorf(SetOutboundInspectionCertificateRes.ErrorMsg)
	}

	d.SetId("set-outbound-inspection-certificate" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(SetOutboundInspectionCertificateRes.GetData()))
	return readManagementSetOutboundInspectionCertificate(d, m)
}

func readManagementSetOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
