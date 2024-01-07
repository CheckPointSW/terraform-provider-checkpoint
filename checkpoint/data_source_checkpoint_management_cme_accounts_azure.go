package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCMEAccountsAzure() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEAccountsAzureRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique account name for identification.",
			},
			"subscription": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Azure subscription ID.",
			},
			"directory_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Azure Active Directory tenant ID.",
			},
			"application_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The application ID with which the service principal is associated.",
			},
			"client_secret": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The service principal's client secret.",
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
		},
	}
}

func dataSourceManagementCMEAccountsAzureRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	log.Println("Read cme Azure account - name = ", name)
	url := CmeApiPath + "/accounts/" + name

	AzureAccountRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	account := AzureAccountRes.GetData()
	if checkIfRequestFailed(account) {
		errMessage := buildErrorMessage(account)
		return fmt.Errorf(errMessage)
	}
	d.SetId("cme-azure-account-" + name + "-" + acctest.RandString(10))

	AzureAccount := account["result"].(map[string]interface{})

	_ = d.Set("name", AzureAccount["name"])

	_ = d.Set("subscription", AzureAccount["subscription"])

	_ = d.Set("directory_id", AzureAccount["directory_id"])

	_ = d.Set("application_id", AzureAccount["application_id"])

	_ = d.Set("client_secret", AzureAccount["client_secret"])

	_ = d.Set("deletion_tolerance", AzureAccount["deletion_tolerance"])

	_ = d.Set("domain", AzureAccount["domain"])

	_ = d.Set("platform", AzureAccount["platform"])

	_ = d.Set("gw_configurations", AzureAccount["gw_configurations"])

	return nil
}
