package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementSetCpTrustedCaCertificate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSetCpTrustedCaCertificate,
		Read:   readManagementSetCpTrustedCaCertificate,
		Delete: deleteManagementSetCpTrustedCaCertificate,
		Schema: map[string]*schema.Schema{

			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Certificate Object uid.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Certificate Object name.",
			},
			"status": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates whether the trusted CP CA certificate is enabled/disabled.",
			},
		},
	}
}

func createManagementSetCpTrustedCaCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else {
		if v, ok := d.GetOk("uid"); ok {
			payload["uid"] = v.(string)
		}
	}

	if v, ok := d.GetOk("status"); ok {
		payload["status"] = v.(string)
	}

	SetCpTrustedCaCertificateRes, _ := client.ApiCall("set-cp-trusted-ca-certificate", payload, client.GetSessionID(), true, false)
	if !SetCpTrustedCaCertificateRes.Success {
		return fmt.Errorf(SetCpTrustedCaCertificateRes.ErrorMsg)
	}

	cpTrustedCaCertificateObj := SetCpTrustedCaCertificateRes.GetData()

	d.SetId(cpTrustedCaCertificateObj["uid"].(string))

	return readManagementSetCpTrustedCaCertificate(d, m)
}

func readManagementSetCpTrustedCaCertificate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementSetCpTrustedCaCertificate(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
