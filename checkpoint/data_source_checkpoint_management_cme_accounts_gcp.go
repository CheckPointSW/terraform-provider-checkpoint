package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCMEAccountsGCP() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEAccountsGCPRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique account name for identification.",
			},
			"project_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project id.",
			},
			"credentials_file": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The credentials file.",
			},
			"credentials_data": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Base64 encoded string that represents the content of the credentials file.",
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

func dataSourceManagementCMEAccountsGCPRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	log.Println("Read cme GCP account - name = ", name)
	url := CmeApiPath + "/accounts/" + name

	GCPAccountRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	account := GCPAccountRes.GetData()
	if checkIfRequestFailed(account) {
		errMessage := buildErrorMessage(account)
		return fmt.Errorf(errMessage)
	}
	d.SetId("cme-gcp-account-" + name + "-" + acctest.RandString(10))

	GCPAccount := account["result"].(map[string]interface{})

	_ = d.Set("name", GCPAccount["name"])

	_ = d.Set("project_id", GCPAccount["project_id"])

	_ = d.Set("credentials_file", GCPAccount["credentials_file"])

	_ = d.Set("credentials_data", GCPAccount["credentials_data"])

	_ = d.Set("deletion_tolerance", GCPAccount["deletion_tolerance"])

	_ = d.Set("domain", GCPAccount["domain"])

	_ = d.Set("platform", GCPAccount["platform"])

	_ = d.Set("gw_configurations", GCPAccount["gw_configurations"])

	return nil
}
