package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementCMEDelayCycle() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEDelayCycle,
		Update: updateManagementCMEDelayCycle,
		Read:   readManagementCMEDelayCycle,
		Delete: deleteManagementCMEDelayCycle,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"delay_cycle": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Time to wait in seconds after each poll cycle.",
			},
		},
	}
}

func createManagementCMEDelayCycle(d *schema.ResourceData, m interface{}) error {
	err := createUpdateManagementCMEDelayCycle(d, m)
	if err != nil {
		return err
	}
	d.SetId("cme-delay-cycle-" + acctest.RandString(10))
	return readManagementCMEDelayCycle(d, m)
}

func updateManagementCMEDelayCycle(d *schema.ResourceData, m interface{}) error {
	err := createUpdateManagementCMEDelayCycle(d, m)
	if err != nil {
		return err
	}
	return readManagementCMEDelayCycle(d, m)
}

func createUpdateManagementCMEDelayCycle(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})
	if v, ok := d.GetOk("delay_cycle"); ok {
		payload["delay_cycle"] = v.(int)
	}

	log.Println("Update cme delay cycle - payload = ", payload)

	url := CmeApiPath + "/generalConfiguration/delayCycle"

	cmeDelayCycleRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed(), "PUT")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeDelayCycleRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}
	return nil
}

func readManagementCMEDelayCycle(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	log.Println("Read cme delay cycle")
	url := CmeApiPath + "/generalConfiguration/delayCycle"

	cmeDelayCycleRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeDelayCycleRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	cmeDelayCycleData := data["result"].(map[string]interface{})

	_ = d.Set("delay_cycle", cmeDelayCycleData["delay_cycle"])

	return nil
}

func deleteManagementCMEDelayCycle(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
