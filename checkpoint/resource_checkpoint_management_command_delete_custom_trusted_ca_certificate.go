package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementDeleteCustomTrustedCaCertificate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDeleteCustomTrustedCaCertificate,
		Read:   readManagementDeleteCustomTrustedCaCertificate,
		Delete: deleteManagementDeleteCustomTrustedCaCertificate,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object name.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object name.",
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

func createManagementDeleteCustomTrustedCaCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		if v, ok := d.GetOk("name"); ok {
			payload["name"] = v.(string)
		}
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	DeleteCustomTrustedCaCertificateRes, _ := client.ApiCall("delete-custom-trusted-ca-certificate", payload, client.GetSessionID(), true, false)
	if !DeleteCustomTrustedCaCertificateRes.Success {
		return fmt.Errorf(DeleteCustomTrustedCaCertificateRes.ErrorMsg)
	}
	d.SetId("delete-custom-trusted-ca-certificate-" + acctest.RandString(10))

	return readManagementDeleteCustomTrustedCaCertificate(d, m)
}

func readManagementDeleteCustomTrustedCaCertificate(d *schema.ResourceData, m interface{}) error {

	return nil
}

func deleteManagementDeleteCustomTrustedCaCertificate(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
