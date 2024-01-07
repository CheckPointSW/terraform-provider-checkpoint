package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCMEAPIVersions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEAPIVersionsRead,
		Schema: map[string]*schema.Schema{
			"current_version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current CME API version.",
			},
			"supported_versions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Supported CME API versions.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementCMEAPIVersionsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	log.Println("Read cme api versions")
	url := CmeApiPath + "/api-versions"

	cmeVersionRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	cmeAPIVersionsJson := cmeVersionRes.GetData()
	if checkIfRequestFailed(cmeAPIVersionsJson) {
		errMessage := buildErrorMessage(cmeAPIVersionsJson)
		return fmt.Errorf(errMessage)
	}

	d.SetId("cme-api-versions-" + acctest.RandString(10))

	cmeAPIVersions := cmeAPIVersionsJson["result"].(map[string]interface{})

	_ = d.Set("current_version", cmeAPIVersions["current_version"])

	_ = d.Set("supported_versions", cmeAPIVersions["supported_versions"])

	return nil
}
