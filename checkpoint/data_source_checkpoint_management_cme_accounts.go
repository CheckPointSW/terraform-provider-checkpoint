package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCMEAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEAccountsRead,
		Schema: map[string]*schema.Schema{
			"result": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Response data - contains all accounts",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique account name for identification.",
						},
						"platform": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The platform of the account.",
						},
						"gw_configurations": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A list of GW configurations attached to the account",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"deletion_tolerance": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The number of CME cycles to wait when the cloud provider does not return a GW until its deletion.",
						},
						"domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The account's domain name in MDS environment.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementCMEAccountsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	log.Println("Read cme accounts")

	url := CmeApiPath + "/accounts"
	AccountsRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	data := AccountsRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}
	d.SetId("cme-accounts-" + acctest.RandString(10))

	accountsList := data["result"].([]interface{})
	var accountsListToReturn []map[string]interface{}
	if len(accountsList) > 0 {
		for i := range accountsList {
			singleAccount := accountsList[i].(map[string]interface{})
			tempObject := make(map[string]interface{})
			tempObject["name"] = singleAccount["name"]
			tempObject["platform"] = singleAccount["platform"]
			tempObject["gw_configurations"] = singleAccount["gw_configurations"]
			tempObject["deletion_tolerance"] = singleAccount["deletion_tolerance"]
			tempObject["domain"] = singleAccount["domain"]
			accountsListToReturn = append(accountsListToReturn, tempObject)
		}
		_ = d.Set("result", accountsListToReturn)
	} else {
		_ = d.Set("result", []interface{}{})
	}
	return nil
}
