package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementCMEAccountsAzure() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEAccountsAzure,
		Update: updateManagementCMEAccountsAzure,
		Read:   readManagementCMEAccountsAzure,
		Delete: deleteManagementCMEAccountsAzure,
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
			"subscription": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Azure subscription ID.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"directory_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Azure Active Directory tenant ID.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"application_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The application ID with which the service principal is associated.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"client_secret": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The service principal's client secret.",
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

func createManagementCMEAccountsAzure(d *schema.ResourceData, m interface{}) error {
	return createUpdateAccountAzure(d, m)
}

func updateManagementCMEAccountsAzure(d *schema.ResourceData, m interface{}) error {
	return createUpdateAccountAzure(d, m)
}

func readManagementCMEAccountsAzure(d *schema.ResourceData, m interface{}) error {
	d.SetId(generateId())
	_ = d.Set("status_code", 200)
	return nil
}

func deleteManagementCMEAccountsAzure(d *schema.ResourceData, m interface{}) error {
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

func createUpdateAccountAzure(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name string
	var method string = "POST"

	url := "cme-api/v1/accounts/azure"

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

	if v, ok := d.GetOk("subscription"); ok {
		payload["subscription"] = v.(string)
	} else if !ok && !updateAccount {
		return fmt.Errorf("expected subscription id when creating new account")
	}
	if v, ok := d.GetOk("directory_id"); ok {
		payload["directory_id"] = v.(string)
	} else if !ok && !updateAccount {
		return fmt.Errorf("expected directory id when creating new account")
	}
	if v, ok := d.GetOk("application_id"); ok {
		payload["application_id"] = v.(string)
	} else if !ok && !updateAccount {
		return fmt.Errorf("expected application id when creating new account")
	}
	if v, ok := d.GetOk("client_secret"); ok {
		payload["client_secret"] = v.(string)
	} else if !ok && !updateAccount {
		return fmt.Errorf("expected client secret when creating new account")
	}
	if v, ok := d.GetOk("deletion_tolerance"); ok {
		payload["deletion_tolerance"] = v.(int)
	}

	log.Println("Set cme Azure account - Map = ", payload)

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

	return readManagementCMEAccountsAzure(d, m)
}
