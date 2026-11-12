package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementCMEVersion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEVersionRead,
		Schema: map[string]*schema.Schema{
			"take": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "CME take number.",
			},
		},
	}
}

func dataSourceManagementCMEVersionRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	log.Println("Read cme version")
	url := CmeApiPath + "/generalConfiguration/cmeVersion"

	cmeVersionRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}

	cmeVersionJson := cmeVersionRes.GetData()
	if checkIfRequestFailed(cmeVersionJson) {
		errMessage := buildErrorMessage(cmeVersionJson)
		return fmt.Errorf("%s", errMessage)
	}

	d.SetId("cme-version-" + acctest.RandString(10))

	cmeVersion := cmeVersionJson["result"].(map[string]interface{})

	_ = d.Set("take", cmeVersion["take"])

	return nil
}
