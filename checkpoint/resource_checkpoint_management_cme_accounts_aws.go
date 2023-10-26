package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"regions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Comma-separated list of AWS regions, in which tahe gateways are being deployed.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"credentials_file": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The credentials file.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"deletion_tolerance": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The number of CME cycles to wait when the cloud provider does not return a GW until its deletion.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v < 0 {
						errs = append(errs, fmt.Errorf("%v must not be a number lower then 0", key))
					}
					return
				},
			},
			"access_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS access key.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" || len(v) > 30 {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS secret key.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" || len(v) > 50 {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"sts_role": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS sts role.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" || len(v) > 50 {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"sts_external_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "AWS sts external id, must exist with sts role.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" || len(v) > 50 {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
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
			"communities": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "AWS communities.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"scan_subnets": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Set true in order to scan subnets with AWS GWLB.",
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
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(string)
								if v == "" {
									errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
								}
								return
							},
						},
						"credentials_file": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The credentials file.",
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(string)
								if v == "" {
									errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
								}
								return
							},
						},
						"access_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "AWS access key.",
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(string)
								if v == "" || len(v) > 30 {
									errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
								}
								return
							},
						},
						"secret_key": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "AWS secret key.",
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(string)
								if v == "" || len(v) > 50 {
									errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
								}
								return
							},
						},
						"sts_role": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "AWS sts role.",
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(string)
								if v == "" || len(v) > 50 {
									errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
								}
								return
							},
						},
						"sts_external_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "AWS sts external id, must exist with sts role.",
							ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
								v := val.(string)
								if v == "" || len(v) > 50 {
									errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
								}
								return
							},
						},
					},
				},
			},
			"status_code": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Result status code.",
			},
			"result": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"error": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"details": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error detials.",
						},
						"error_code": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Error code.",
						},
						"message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error message.",
						},
					},
				},
			},
		},
	}
}

func createManagementCMEAccountsAWS(d *schema.ResourceData, m interface{}) error {
	return createUpdateAccountAWS(d, m)
}

func updateManagementCMEAccountsAWS(d *schema.ResourceData, m interface{}) error {
	return createUpdateAccountAWS(d, m)
}

func readManagementCMEAccountsAWS(d *schema.ResourceData, m interface{}) error {
	d.SetId(generateId())
	_ = d.Set("status_code", 200)
	return nil
}

func deleteManagementCMEAccountsAWS(d *schema.ResourceData, m interface{}) error {
	var name string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	res, err := deleteAccount(name, m)

	if res != nil {
		if v, ok := res["result"]; ok {
			_ = d.Set("result", v)
		}
		if v, ok := res["error"]; ok {
			_ = d.Set("error", v)
		}
	}

	if err != nil {
		return err
	}

	_ = d.Set("status_code", 200)
	d.SetId("")
	return nil
}

func createUpdateAccountAWS(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string
	var method string = "POST"

	url := "cme-api/v1/accounts/aws"

	payload := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	isExist, err := checkAccountExisting(name, m)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	var updateAccount bool = false
	if isExist {
		method = "PUT"
		url += "/" + name
		updateAccount = true
	} else {
		payload["name"] = name
	}

	if v, ok := d.GetOk("regions"); ok {
		payload["regions"] = v.([]interface{})
	} else if !ok && !updateAccount {
		return fmt.Errorf("expected regions when creating new account")
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
	if v, ok := d.GetOk("scan_gateways"); ok {
		payload["scan_gateways"] = v.(bool)
	}
	if v, ok := d.GetOk("scan_vpn"); ok {
		payload["scan_vpn"] = v.(bool)
	}
	if v, ok := d.GetOk("scan_load_balancers"); ok {
		payload["scan_load_balancers"] = v.(bool)
	}
	if v, ok := d.GetOk("communities"); ok {
		payload["communities"] = v.([]interface{})
	}
	if v, ok := d.GetOk("scan_subnets"); ok {
		payload["scan_subnets"] = v.(bool)
	}
	if v, ok := d.GetOk("sub_accounts"); ok {
		var tempList []map[string]interface{}
		for _, subAccount := range v.([]interface{}) {
			tempObject := make(map[string]interface{})
			for key, value := range subAccount.(map[string]interface{}) {
				strValue := value.(string)
				if strValue != "" {
					tempObject[key] = strValue
				}
			}
			tempList = append(tempList, tempObject)
		}
		payload["sub_accounts"] = tempList
	}

	log.Println("Set cme AWS account - Map = ", payload)

	cmeAccoutsRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed(), method)

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cmeAccoutsRes.Success {
		return fmt.Errorf(cmeAccoutsRes.ErrorMsg)
	}

	cmeAccountsJson := cmeAccoutsRes.GetData()
	cmeAccountsToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if cmeAccountsJson["error"] != nil {
		errorResult, ok := cmeAccountsJson["error"]

		if ok {
			errorResultJson := errorResult.(map[string]interface{})
			tempObject := make(map[string]interface{})

			if v := errorResultJson["details"]; v != nil {
				tempObject["details"] = v.(string)
				err_message = v.(string)
				has_error = true
			}
			if v := errorResultJson["error_code"]; v != nil {
				var error_code string = strconv.Itoa(int(math.Round(v.(float64))))
				tempObject["error_code"] = error_code
				has_error = true
			}
			if v := errorResultJson["message"]; v != nil {
				tempObject["message"] = v
				has_error = true
			}

			cmeAccountsToReturn["error"] = tempObject
		}
	} else {
		cmeAccountsToReturn["result"] = map[string]interface{}{}
		cmeAccountsToReturn["error"] = map[string]interface{}{}
	}

	_ = d.Set("result", cmeAccountsToReturn["result"])
	_ = d.Set("error", cmeAccountsToReturn["error"])

	if has_error {
		return fmt.Errorf(err_message)
	}

	return readManagementCMEAccountsAWS(d, m)
}
