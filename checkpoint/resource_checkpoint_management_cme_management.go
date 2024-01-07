package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementCMEManagement() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEManagement,
		Update: updateManagementCMEManagement,
		Read:   readManagementCMEManagement,
		Delete: deleteManagementCMEManagement,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The name of the management server.",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The management's domain name in MDS environment.",
			},
			"host": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The host of the management server.",
			},
		},
	}
}

func createManagementCMEManagement(d *schema.ResourceData, m interface{}) error {
	err := createUpdateManagementCMEManagement(d, m)
	if err != nil {
		return err
	}
	d.SetId("cme-management-" + acctest.RandString(10))
	return readManagementCMEManagement(d, m)
}

func updateManagementCMEManagement(d *schema.ResourceData, m interface{}) error {
	err := createUpdateManagementCMEManagement(d, m)
	if err != nil {
		return err
	}
	return readManagementCMEManagement(d, m)
}

func createUpdateManagementCMEManagement(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}
	if v, ok := d.GetOk("domain"); ok {
		payload["domain"] = v.(string)
	}

	log.Println("Update cme management - payload = ", payload)

	url := CmeApiPath + "/management"

	cmeManagementRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed(), "PUT")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeManagementRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}
	return nil
}

func readManagementCMEManagement(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	log.Println("Read cme management")
	url := CmeApiPath + "/management"

	cmeManagementRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeManagementRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}
	cmeManagementData := data["result"].(map[string]interface{})

	_ = d.Set("name", cmeManagementData["name"])

	_ = d.Set("domain", cmeManagementData["domain"])

	_ = d.Set("host", cmeManagementData["host"])

	return nil
}

func deleteManagementCMEManagement(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
