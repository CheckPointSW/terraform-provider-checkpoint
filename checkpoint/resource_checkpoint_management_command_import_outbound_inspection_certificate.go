package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementImportOutboundInspectionCertificate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementImportOutboundInspectionCertificate,
		Read:   readManagementImportOutboundInspectionCertificate,
		Delete: deleteManagementImportOutboundInspectionCertificate,
		Schema: map[string]*schema.Schema{
			"base64_certificate": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Certificate file encoded in base64.<br/>Valid file format: p12.",
			},
			"base64_password": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Password (encoded in Base64 with padding) for the certificate file.",
			},
			"task_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Command asynchronous task unique identifier.",
			},
		},
	}
}

func createManagementImportOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("base64_certificate"); ok {
		payload["base64-certificate"] = v.(string)
	}

	if v, ok := d.GetOk("base64_password"); ok {
		payload["base64-password"] = v.(string)
	}

	ImportOutboundInspectionCertificateRes, _ := client.ApiCall("import-outbound-inspection-certificate", payload, client.GetSessionID(), true, false)
	if !ImportOutboundInspectionCertificateRes.Success {
		return fmt.Errorf(ImportOutboundInspectionCertificateRes.ErrorMsg)
	}

	d.SetId("import-outbound-inspection-certificate" + acctest.RandString(10))
	_ = d.Set("task_id", resolveTaskId(ImportOutboundInspectionCertificateRes.GetData()))

	return readManagementImportOutboundInspectionCertificate(d, m)
}

func readManagementImportOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementImportOutboundInspectionCertificate(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
