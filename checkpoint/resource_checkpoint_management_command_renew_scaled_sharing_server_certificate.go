package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementRenewScaledSharingServerCertificate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRenewScaledSharingServerCertificate,
		Read:   readManagementRenewScaledSharingServerCertificate,
		Delete: deleteManagementRenewScaledSharingServerCertificate,
		Schema: map[string]*schema.Schema{
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Gateway or cluster unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: "Gateway or cluster name.",
			},
			"message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Operation status.",
			},
		},
	}
}

func createManagementRenewScaledSharingServerCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("name or uid must be specified")
	}

	RenewScaledSharingServerCertificateRes, err := client.ApiCall("renew-scaled-sharing-server-certificate", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !RenewScaledSharingServerCertificateRes.Success {
		return fmt.Errorf(RenewScaledSharingServerCertificateRes.ErrorMsg)
	}

	resetSicStatusProfile := RenewScaledSharingServerCertificateRes.GetData()

	if v := resetSicStatusProfile["message"]; v != nil {
		_ = d.Set("message", v)
	}

	log.Println("Read ResetSicStatus - Show JSON = ", resetSicStatusProfile)

	d.SetId("renew-scaled-sharing-server-certificate-" + acctest.RandString(10))

	return nil
}

func readManagementRenewScaledSharingServerCertificate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteManagementRenewScaledSharingServerCertificate(d *schema.ResourceData, m interface{}) error {

	d.SetId("")
	return nil
}
