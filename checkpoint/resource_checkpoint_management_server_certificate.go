package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementServerCertificate() *schema.Resource {
	return &schema.Resource{
		Create: createManagementServerCertificate,
		Read:   readManagementServerCertificate,
		Update: updateManagementServerCertificate,
		Delete: deleteManagementServerCertificate,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"base64_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Certificate file encoded in base64.<br/>Valid file formats: p12.",
			},
			"base64_password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Base64 encoded password of the certificate file.",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Server certificate comments.",
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
		},
	}
}

func createManagementServerCertificate(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	serverCertificate := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		serverCertificate["name"] = v.(string)
	}

	if v, ok := d.GetOk("base64_certificate"); ok {
		serverCertificate["base64-certificate"] = v.(string)
	}

	if v, ok := d.GetOk("base64_password"); ok {
		serverCertificate["base64-password"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		serverCertificate["comments"] = v.(string)
	}

	log.Println("Create ServerCertificate - Map = ", serverCertificate)

	addServerCertificateRes, err := client.ApiCall("add-server-certificate", serverCertificate, client.GetSessionID(), true, false)
	if err != nil || !addServerCertificateRes.Success {
		if addServerCertificateRes.ErrorMsg != "" {
			return fmt.Errorf(addServerCertificateRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addServerCertificateRes.GetData()["uid"].(string))

	return readManagementServerCertificate(d, m)
}

func readManagementServerCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showServerCertificateRes, err := client.ApiCall("show-server-certificate", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServerCertificateRes.Success {
		if objectNotFound(showServerCertificateRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showServerCertificateRes.ErrorMsg)
	}

	serverCertificate := showServerCertificateRes.GetData()

	log.Println("Read ServerCertificate - Show JSON = ", serverCertificate)

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

	return nil

}

func updateManagementServerCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	serverCertificate := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		serverCertificate["name"] = oldName
		serverCertificate["new-name"] = newName
	} else {
		serverCertificate["name"] = d.Get("name")
	}

	if ok := d.HasChange("base64_certificate"); ok {
		serverCertificate["base64-certificate"] = d.Get("base64_certificate")
	}

	if ok := d.HasChange("base64_password"); ok {
		serverCertificate["base64-password"] = d.Get("base64_password")
	}

	if ok := d.HasChange("comments"); ok {
		serverCertificate["comments"] = d.Get("comments")
	}

	log.Println("Update ServerCertificate - Map = ", serverCertificate)

	updateServerCertificateRes, err := client.ApiCall("set-server-certificate", serverCertificate, client.GetSessionID(), true, false)
	if err != nil || !updateServerCertificateRes.Success {
		if updateServerCertificateRes.ErrorMsg != "" {
			return fmt.Errorf(updateServerCertificateRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementServerCertificate(d, m)
}

func deleteManagementServerCertificate(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	serverCertificatePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ServerCertificate")

	deleteServerCertificateRes, err := client.ApiCall("delete-server-certificate", serverCertificatePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteServerCertificateRes.Success {
		if deleteServerCertificateRes.ErrorMsg != "" {
			return fmt.Errorf(deleteServerCertificateRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
