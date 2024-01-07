package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCMEDelayCycle() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEDelayCycleRead,
		Schema: map[string]*schema.Schema{
			"delay_cycle": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Time to wait in seconds after each poll cycle.",
			},
		},
	}
}

func dataSourceManagementCMEDelayCycleRead(d *schema.ResourceData, m interface{}) error {
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

	d.SetId("cme-delay-cycle-" + acctest.RandString(10))

	cmeDelayCycleData := data["result"].(map[string]interface{})

	_ = d.Set("delay_cycle", cmeDelayCycleData["delay_cycle"])

	return nil
}
