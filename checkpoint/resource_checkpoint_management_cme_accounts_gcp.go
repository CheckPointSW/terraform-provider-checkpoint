package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementCMEAccountsGCP() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEAccountsGCP,
		Update: updateManagementCMEAccountsGCP,
		Read:   readManagementCMEAccountsGCP,
		Delete: deleteManagementCMEAccountsGCP,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The account name.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"project_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The project id.",
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

func createManagementCMEAccountsGCP(d *schema.ResourceData, m interface{}) error {
	return createUpdateAccountGCP(d, m)
}

func updateManagementCMEAccountsGCP(d *schema.ResourceData, m interface{}) error {
	return createUpdateAccountGCP(d, m)
}

func readManagementCMEAccountsGCP(d *schema.ResourceData, m interface{}) error {
	d.SetId(generateId())
	_ = d.Set("status_code", 200)
	return nil
}

func deleteManagementCMEAccountsGCP(d *schema.ResourceData, m interface{}) error {
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

func createUpdateAccountGCP(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string
	var method string = "POST"

	url := "cme-api/v1/accounts/gcp"

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

	if v, ok := d.GetOk("project_id"); ok {
		payload["project_id"] = v.(string)
	} else if !ok && !updateAccount {
		return fmt.Errorf("expected project id when creating new account")
	}
	if v, ok := d.GetOk("credentials_file"); ok {
		payload["credentials_file"] = v.(string)
	} else if !ok && !updateAccount {
		return fmt.Errorf("expected credentials file when creating new account")
	}
	if v, ok := d.GetOk("deletion_tolerance"); ok {
		payload["deletion_tolerance"] = v.(int)
	}

	log.Println("Set cme GCP account - Map = ", payload)

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

	return readManagementCMEAccountsGCP(d, m)
}
