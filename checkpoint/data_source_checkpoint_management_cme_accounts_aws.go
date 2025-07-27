package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementCMEAccountsAWS() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEAccountsAWSRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique account name for identification.",
			},
			"platform": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The platform of the account.",
			},
			"regions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Comma-separated list of AWS regions, in which the gateways are being deployed.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"gw_configurations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A list of GW configurations attached to the account",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"credentials_file": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The credentials file.",
			},
			"deletion_tolerance": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The number of CME cycles to wait when the cloud provider does not return a GW until its deletion.",
			},
			"access_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AWS access key.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AWS secret key.",
			},
			"sts_role": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AWS sts role.",
			},
			"sts_external_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AWS sts external id, must exist with sts role.",
			},
			"scan_gateways": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set true in order to scan gateways with AWS TGW.",
			},
			"scan_vpn": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set true in order to scan vpn with AWS TGW.",
			},
			"scan_load_balancers": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set true in order to scan load balancers access and NAT rules with AWS TGW.",
			},
			"scan_subnets": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set true in order to scan subnets with AWS GWLB.",
			},
			"scan_subnets_6": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Set true in order to scan IPv6 subnets with AWS GWLB.",
			},
			"communities": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "AWS communities.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sub_accounts": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "AWS sub accounts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique account name for identification.",
						},
						"credentials_file": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The credentials file.",
						},
						"access_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS access key.",
						},
						"secret_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS secret key.",
						},
						"sts_role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS sts role.",
						},
						"sts_external_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "AWS sts external id, must exist with sts role.",
						},
					},
				},
			},
			"domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The account's domain name in MDS environment.",
			},
		},
	}
}

func dataSourceManagementCMEAccountsAWSRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	log.Println("Read cme AWS account - name = ", name)

	url := CmeApiPath + "/accounts/" + name

	AWSAccountRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	account := AWSAccountRes.GetData()
	if checkIfRequestFailed(account) {
		errMessage := buildErrorMessage(account)
		return fmt.Errorf(errMessage)
	}

	d.SetId("cme-aws-account-" + name + "-" + acctest.RandString(10))

	AWSAccount := account["result"].(map[string]interface{})

	_ = d.Set("name", AWSAccount["name"])

	_ = d.Set("platform", AWSAccount["platform"])

	_ = d.Set("regions", AWSAccount["regions"])

	_ = d.Set("gw_configurations", AWSAccount["gw_configurations"])

	_ = d.Set("credentials_file", AWSAccount["credentials_file"])

	_ = d.Set("deletion_tolerance", AWSAccount["deletion_tolerance"])

	_ = d.Set("access_key", AWSAccount["access_key"])

	_ = d.Set("secret_key", AWSAccount["secret_key"])

	_ = d.Set("sts_role", AWSAccount["sts_role"])

	_ = d.Set("sts_external_id", AWSAccount["sts_external_id"])

	if AWSAccount["sync"] != nil {
		syncMap := AWSAccount["sync"].(map[string]interface{})
		_ = d.Set("scan_gateways", syncMap["scan_gateways"])
		_ = d.Set("scan_vpn", syncMap["scan_vpn"])
		_ = d.Set("scan_load_balancers", syncMap["scan_load_balancers"])
		_ = d.Set("scan_subnets", syncMap["scan_subnets"])
		_ = d.Set("scan_subnets_6", syncMap["scan_subnets_6"])
	} else {
		_ = d.Set("scan_gateways", nil)
		_ = d.Set("scan_vpn", nil)
		_ = d.Set("scan_load_balancers", nil)
		_ = d.Set("scan_subnets", nil)
		_ = d.Set("scan_subnets_6", nil)
	}
	_ = d.Set("communities", AWSAccount["communities"])

	if AWSAccount["sub_accounts"] != nil {
		subAccountsMap := AWSAccount["sub_accounts"].(map[string]interface{})
		if len(subAccountsMap) > 0 {
			var subAccountsListToReturn []map[string]interface{}
			for key, value := range subAccountsMap {
				subAccountMap := value.(map[string]interface{})
				subAccountMapToAdd := make(map[string]interface{})
				subAccountMapToAdd["name"] = key
				subAccountMapToAdd["credentials_file"] = subAccountMap["credentials_file"]
				subAccountMapToAdd["access_key"] = subAccountMap["access_key"]
				subAccountMapToAdd["secret_key"] = subAccountMap["secret_key"]
				subAccountMapToAdd["sts_role"] = subAccountMap["sts_role"]
				subAccountMapToAdd["sts_external_id"] = subAccountMap["sts_external_id"]
				subAccountsListToReturn = append(subAccountsListToReturn, subAccountMapToAdd)
			}
			_ = d.Set("sub_accounts", subAccountsListToReturn)
		} else {
			_ = d.Set("sub_accounts", []interface{}{})
		}
	} else {
		_ = d.Set("sub_accounts", nil)
	}
	_ = d.Set("domain", AWSAccount["domain"])
	return nil
}
