package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementServerCertificate() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementServerCertificateRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"subject": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate's subject.",
			},
			"valid_from": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server certificate valid from date.",
			},
			"valid_to": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server certificate valid up to date.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server certificate comments.",
			},
		},
	}
}
func dataSourceManagementServerCertificateRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}
	showServerCertificateRes, err := client.ApiCall("show-server-certificate", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServerCertificateRes.Success {
		fmt.Errorf(showServerCertificateRes.ErrorMsg)
	}

	serverCertificate := showServerCertificateRes.GetData()

	log.Println("Read ServerCertificate - Show JSON ", serverCertificate)

	if v := serverCertificate["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := serverCertificate["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serverCertificate["subject"]; v != nil {
		_ = d.Set("subject", v)
	}
	if v := serverCertificate["valid-from"]; v != nil {
		_ = d.Set("valid_from", v)
	}
	if v := serverCertificate["valid-to"]; v != nil {
		_ = d.Set("valid_to", v)
	}
	if v := serverCertificate["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
