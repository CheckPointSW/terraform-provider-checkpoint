package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementCMEAccounts() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEAccountsRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A name of an account.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
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
				Type:        schema.TypeList,
				Computed:    true,
				Description: "N/A",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the account.",
						},
						"platform": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The platform of the account.",
						},
						"project_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the project.",
						},
						"credentials_file": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The credential file.",
						},
						"deletion_tolerance": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Number of the deletion_tolerance.",
						},
						"subscription": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subscription ID.",
						},
						"directory_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Directory ID.",
						},
						"application_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Application ID.",
						},
						"client_secret": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Client secret.",
						},
						"access_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Access key.",
						},
						"secret_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sts external id.",
						},
						"sts_role": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sts role.",
						},
						"sts_external_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Sts external id.",
						},
						"gw_configurations": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A list of gateway configurations.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"regions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A list of regions of the account.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"communities": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A list of AWS communities.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"sync": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "AWS sync.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
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
					},
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

func dataSourceManagementCMEAccountsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var url string = "cme-api/v1/accounts"
	var filter bool = false

	if v, ok := d.GetOk("name"); ok {
		url += "/" + v.(string)
		filter = true
	}

	cmeAccountsRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cmeAccountsRes.Success {
		return fmt.Errorf(cmeAccountsRes.ErrorMsg)
	}
	cmeAccountsJson := cmeAccountsRes.GetData()
	log.Println("Read cme accounts - Show JSON = ", cmeAccountsJson)

	cmeAccountsToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if v := cmeAccountsJson["status-code"]; v != nil {
		_ = d.Set("status_code", int(math.Round(v.(float64))))
	}

	if cmeAccountsJson["result"] != nil {
		if !filter {
			cmeAccountsResultList, ok := cmeAccountsJson["result"].([]interface{})
			var objectDictToReturn []map[string]interface{}

			if ok {
				for i := range cmeAccountsResultList {
					cmeAccountsResultJson := cmeAccountsResultList[i].(map[string]interface{})
					tempObject := readSingleAccount(cmeAccountsResultJson)
					objectDictToReturn = append(objectDictToReturn, tempObject)
				}

				cmeAccountsToReturn["result"] = objectDictToReturn
			}
		} else {
			cmeAccountsResultList, ok := cmeAccountsJson["result"]
			var objectDictToReturn []map[string]interface{}

			if ok {
				cmeAccountsResultJson := cmeAccountsResultList.(map[string]interface{})
				tempObject := readSingleAccount(cmeAccountsResultJson)
				objectDictToReturn = append(objectDictToReturn, tempObject)

				cmeAccountsToReturn["result"] = objectDictToReturn
			}
		}
	} else if cmeAccountsJson["error"] != nil {
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
				tempObject["message"] = v.(string)
				has_error = true
			}

			cmeAccountsToReturn["error"] = tempObject
		}
	} else {
		cmeAccountsToReturn["result"] = map[string]interface{}{}
		cmeAccountsToReturn["error"] = map[string]interface{}{}
	}

	d.SetId(generateId())
	_ = d.Set("result", cmeAccountsToReturn["result"])
	_ = d.Set("error", cmeAccountsToReturn["error"])

	if has_error {
		return fmt.Errorf(err_message)
	}

	return nil
}

func readSingleAccount(cmeAccountsResultJson map[string]interface{}) map[string]interface{} {
	tempObject := make(map[string]interface{})

	for key, value := range cmeAccountsResultJson {
		switch key {
		case "name", "platform", "subscription", "directory_id", "application_id", "client_secret",
			"credentials_file", "access_key", "sts_role", "sts_external_id", "project_id":
			tempObject[key] = value.(string)
		case "deletion_tolerance":
			tempObject[key] = int(math.Round(value.(float64)))
		case "gw_configurations", "regions", "communities":
			vList := value.([]interface{})
			var tempList = make([]string, len(vList))

			for j := range vList {
				tempList[j] = vList[j].(string)
			}

			tempObject[key] = tempList
		case "sync":
			vMap := value.(map[string]interface{})
			tempMap := make(map[string]interface{})

			for k, v := range vMap {
				tempMap[k] = strconv.FormatBool(v.(bool))
			}

			tempObject[key] = tempMap
		case "sub_accounts":
			vMap := value.(map[string]interface{})
			var tempList []map[string]interface{}

			for k, v := range vMap {
				tempSubAccountMap := make(map[string]interface{})
				tempSubAccountMap["name"] = k
				for key, value := range v.(map[string]interface{}) {
					tempSubAccountMap[key] = value.(string)
				}
				tempList = append(tempList, tempSubAccountMap)
			}

			tempObject[key] = tempList
		}
	}

	return tempObject
}
