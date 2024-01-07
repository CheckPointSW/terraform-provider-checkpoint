package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCMEManagement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEManagementRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the management server.",
			},
			"domain": {
				Type:        schema.TypeString,
				Computed:    true,
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

func dataSourceManagementCMEManagementRead(d *schema.ResourceData, m interface{}) error {
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

	d.SetId("cme-management-" + acctest.RandString(10))

	cmeManagementData := data["result"].(map[string]interface{})

	_ = d.Set("name", cmeManagementData["name"])

	_ = d.Set("domain", cmeManagementData["domain"])

	_ = d.Set("host", cmeManagementData["host"])

	return nil
}
