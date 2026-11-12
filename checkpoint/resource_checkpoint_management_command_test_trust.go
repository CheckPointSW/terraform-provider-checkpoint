package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementTestTrust() *schema.Resource {
	return &schema.Resource{
		Create: createManagementTestTrust,
		Read:   readManagementTestTrust,
		Delete: deleteManagementTestTrust,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Object unique identifier.",
			},
			"trust_method": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Establish the trust communication method.",
			},
			"domains_to_process": {
				Type:        schema.TypeSet,
				Optional:    true,
				ForceNew:    true,
				Description: "Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ig...",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func createManagementTestTrust(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	}

	if v, ok := d.GetOk("trust_method"); ok {
		payload["trust-method"] = v.(string)
	}

	if v, ok := d.GetOk("domains_to_process"); ok {
		payload["domains-to-process"] = v.(*schema.Set).List()
	}

	TestTrustRes, err := client.ApiCall("test-trust", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !TestTrustRes.Success {
		return fmt.Errorf("%s", TestTrustRes.ErrorMsg)
	}

	d.SetId("test-trust-" + acctest.RandString(10))
	return nil
}

func readManagementTestTrust(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementTestTrust(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
