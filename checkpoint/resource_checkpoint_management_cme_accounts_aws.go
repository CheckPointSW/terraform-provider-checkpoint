package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementCMEAccountsAWS() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEAccountsAWS,
		Update: updateManagementCMEAccountsAWS,
		Read:   readManagementCMEAccountsAWS,
		Delete: deleteManagementCMEAccountsAWS,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
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
				Required:    true,
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
				Optional:    true,
				Description: "The credentials file.",
			},
			"deletion_tolerance": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of CME cycles to wait when the cloud provider does not return a GW until its deletion.",
			},
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS access key.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS secret key.",
				Sensitive:   true,
			},
			"sts_role": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS sts role.",
			},
			"sts_external_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS sts external id, must exist with sts role.",
			},
			"scan_gateways": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set true in order to scan gateways with AWS TGW.",
			},
			"scan_vpn": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set true in order to scan vpn with AWS TGW.",
			},
			"scan_load_balancers": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set true in order to scan load balancers access and NAT rules with AWS TGW.",
			},
			"scan_subnets": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set true in order to scan subnets with AWS GWLB.",
			},
			"communities": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "AWS communities.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sub_accounts": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "AWS sub accounts.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Unique account name for identification.",
						},
						"credentials_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The credentials file.",
						},
						"access_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "AWS access key.",
						},
						"secret_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "AWS secret key.",
							Sensitive:   true,
						},
						"sts_role": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "AWS sts role.",
						},
						"sts_external_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "AWS sts external id, must exist with sts role.",
						},
					},
				},
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The account's domain name in MDS environment.",
			},
		},
	}
}

func readManagementCMEAccountsAWS(d *schema.ResourceData, m interface{}) error {
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
		if cmeObjectNotFound(account) {
			d.SetId("")
			return nil
		}
		errMessage := buildErrorMessage(account)
		return fmt.Errorf(errMessage)
	}

	AWSAccount := account["result"].(map[string]interface{})

	_ = d.Set("name", AWSAccount["name"])

	_ = d.Set("platform", AWSAccount["platform"])

	_ = d.Set("regions", AWSAccount["regions"])

	_ = d.Set("gw_configurations", AWSAccount["gw_configurations"])

	_ = d.Set("credentials_file", AWSAccount["credentials_file"])

	_ = d.Set("deletion_tolerance", AWSAccount["deletion_tolerance"])

	_ = d.Set("access_key", AWSAccount["access_key"])

	_ = d.Set("sts_role", AWSAccount["sts_role"])

	_ = d.Set("sts_external_id", AWSAccount["sts_external_id"])

	if AWSAccount["sync"] != nil {
		syncMap := AWSAccount["sync"].(map[string]interface{})
		_ = d.Set("scan_gateways", syncMap["gateway"])
		_ = d.Set("scan_vpn", syncMap["vpn"])
		_ = d.Set("scan_load_balancers", syncMap["lb"])
		_ = d.Set("scan_subnets", syncMap["scan-subnets"])
	} else {
		_ = d.Set("scan_gateways", nil)
		_ = d.Set("scan_vpn", nil)
		_ = d.Set("scan_load_balancers", nil)
		_ = d.Set("scan_subnets", nil)
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
				if v, _ := subAccountMap["secret_key"]; v != nil {
					if v, ok := d.GetOk("sub_accounts"); ok {
						subAccountsList := v.([]interface{})
						if len(subAccountsList) > 0 {
							for i := range subAccountsList {
								if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".name"); ok {
									if key == v.(string) {
										if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".secret_key"); ok {
											subAccountMapToAdd["secret_key"] = v
											break
										}
									}
								}
							}
						}
					}
				} else {
					subAccountMapToAdd["secret_key"] = nil
				}
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

func createManagementCMEAccountsAWS(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}
	if v, ok := d.GetOk("scan_gateways"); ok {
		payload["scan_gateways"] = v.(bool)
	}
	if v, ok := d.GetOk("scan_vpn"); ok {
		payload["scan_vpn"] = v.(bool)
	}
	if v, ok := d.GetOk("scan_load_balancers"); ok {
		payload["scan_load_balancers"] = v.(bool)
	}
	if v, ok := d.GetOk("scan_subnets"); ok {
		payload["scan_subnets"] = v.(bool)
	}
	if v, ok := d.GetOk("regions"); ok {
		payload["regions"] = v.([]interface{})
	}
	if v, ok := d.GetOk("credentials_file"); ok {
		payload["credentials_file"] = v.(string)
	}
	if v, ok := d.GetOk("deletion_tolerance"); ok {
		payload["deletion_tolerance"] = v.(int)
	}
	if v, ok := d.GetOk("access_key"); ok {
		payload["access_key"] = v.(string)
	}
	if v, ok := d.GetOk("secret_key"); ok {
		payload["secret_key"] = v.(string)
	}
	if v, ok := d.GetOk("sts_role"); ok {
		payload["sts_role"] = v.(string)
	}
	if v, ok := d.GetOk("sts_external_id"); ok {
		payload["sts_external_id"] = v.(string)
	}
	if v, ok := d.GetOk("communities"); ok {
		payload["communities"] = v.([]interface{})
	}
	if v, ok := d.GetOk("sub_accounts"); ok {
		subAccountsList := v.([]interface{})
		if len(subAccountsList) > 0 {
			var subAccountsPayload []map[string]interface{}
			for i := range subAccountsList {
				tempObject := make(map[string]interface{})
				if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".name"); ok {
					tempObject["name"] = v.(string)
				}
				if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".credentials_file"); ok {
					tempObject["credentials_file"] = v.(string)
				}
				if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".access_key"); ok {
					tempObject["access_key"] = v.(string)
				}
				if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".secret_key"); ok {
					tempObject["secret_key"] = v.(string)
				}
				if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".sts_role"); ok {
					tempObject["sts_role"] = v.(string)
				}
				if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".sts_external_id"); ok {
					tempObject["sts_external_id"] = v.(string)
				}
				subAccountsPayload = append(subAccountsPayload, tempObject)
			}
			payload["sub_accounts"] = subAccountsPayload
		} else {
			payload["sub_accounts"] = subAccountsList
		}
	}
	if v, ok := d.GetOk("domain"); ok {
		payload["domain"] = v.(string)
	}
	log.Println("Create cme AWS account - name = ", payload["name"])

	url := CmeApiPath + "/accounts/aws"

	cmeAccountsRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeAccountsRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	d.SetId("cme-aws-account-" + d.Get("name").(string) + "-" + acctest.RandString(10))

	return readManagementCMEAccountsAWS(d, m)
}

func updateManagementCMEAccountsAWS(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := make(map[string]interface{})

	if d.HasChange("scan_gateways") {
		payload["scan_gateways"] = d.Get("scan_gateways")
	}
	if d.HasChange("scan_vpn") {
		payload["scan_vpn"] = d.Get("scan_vpn")
	}
	if d.HasChange("scan_load_balancers") {
		payload["scan_load_balancers"] = d.Get("scan_load_balancers")
	}
	if d.HasChange("scan_subnets") {
		payload["scan_subnets"] = d.Get("scan_subnets")
	}
	if d.HasChange("regions") {
		payload["regions"] = d.Get("regions")
	}
	if d.HasChange("credentials_file") {
		payload["credentials_file"] = d.Get("credentials_file")
	}
	if d.HasChange("deletion_tolerance") {
		payload["deletion_tolerance"] = d.Get("deletion_tolerance")
	}
	if d.HasChange("access_key") {
		payload["access_key"] = d.Get("access_key")
	}
	if d.HasChange("secret_key") {
		payload["secret_key"] = d.Get("secret_key")
	}
	if d.HasChange("sts_role") {
		payload["sts_role"] = d.Get("sts_role")
	}
	if d.HasChange("sts_external_id") {
		payload["sts_external_id"] = d.Get("sts_external_id")
	}
	if d.HasChange("communities") {
		payload["communities"] = d.Get("communities")
	}
	if d.HasChange("sub_accounts") {
		if v, ok := d.GetOk("sub_accounts"); ok {
			subAccountsList := v.([]interface{})
			if len(subAccountsList) > 0 {
				var subAccountsPayload []map[string]interface{}
				for i := range subAccountsList {
					tempObject := make(map[string]interface{})
					if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".name"); ok {
						tempObject["name"] = v.(string)
					}
					if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".credentials_file"); ok {
						tempObject["credentials_file"] = v.(string)
					}
					if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".access_key"); ok {
						tempObject["access_key"] = v.(string)
					}
					if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".secret_key"); ok {
						tempObject["secret_key"] = v.(string)
					}
					if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".sts_role"); ok {
						tempObject["sts_role"] = v.(string)
					}
					if v, ok := d.GetOk("sub_accounts." + strconv.Itoa(i) + ".sts_external_id"); ok {
						tempObject["sts_external_id"] = v.(string)
					}
					subAccountsPayload = append(subAccountsPayload, tempObject)
				}
				payload["sub_accounts"] = subAccountsPayload
			} else {
				payload["sub_accounts"] = subAccountsList
			}
		} else {
			payload["sub_accounts"] = v.([]interface{})
		}
	}
	if d.HasChange("domain") {
		payload["domain"] = d.Get("domain")
	}
	var name string

	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	log.Println("Set cme AWS account - name = ", name)

	url := CmeApiPath + "/accounts/aws/" + name
	cmeAccountsRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed(), "PUT")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := cmeAccountsRes.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	return readManagementCMEAccountsAWS(d, m)
}

func deleteManagementCMEAccountsAWS(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	log.Println("Delete cme AWS account - name = ", name)

	url := CmeApiPath + "/accounts/" + name

	res, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "DELETE")

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	data := res.GetData()
	if checkIfRequestFailed(data) {
		errMessage := buildErrorMessage(data)
		return fmt.Errorf(errMessage)
	}

	d.SetId("")
	return nil
}
