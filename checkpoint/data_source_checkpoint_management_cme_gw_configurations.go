package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementCMEGWConfigurations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEGWConfigurationsRead,
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
							Description: "The name of the configuration.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The version of the configuration.",
						},
						"sic_key": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The configuration sic key.",
						},
						"policy": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configuration policy.",
						},
						"related_account": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Number of the deletion_tolerance.",
						},
						"blades": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Active blades",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
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

func dataSourceManagementCMEGWConfigurationsRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var url string = "cme-api/v1/gwConfigurations"
	var filter bool = false

	if v, ok := d.GetOk("name"); ok {
		url += "/" + v.(string)
		filter = true
	}

	cmeGWConfigurationsRes, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cmeGWConfigurationsRes.Success {
		return fmt.Errorf(cmeGWConfigurationsRes.ErrorMsg)
	}
	cmeGWConfigurationsJson := cmeGWConfigurationsRes.GetData()
	log.Println("Read cme GW configuration - Show JSON = ", cmeGWConfigurationsJson)

	cmeGWConfigurationsToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if v := cmeGWConfigurationsJson["status-code"]; v != nil {
		_ = d.Set("status_code", int(math.Round(v.(float64))))
	}

	if cmeGWConfigurationsJson["result"] != nil {
		if !filter {
			cmeGWConfigurationsResultList, ok := cmeGWConfigurationsJson["result"].([]interface{})
			var objectDictToReturn []map[string]interface{}

			if ok {
				for i := range cmeGWConfigurationsResultList {
					cmeGWConfigurationsResultJson := cmeGWConfigurationsResultList[i].(map[string]interface{})
					tempObject := readSingleConfiguration(cmeGWConfigurationsResultJson)
					objectDictToReturn = append(objectDictToReturn, tempObject)
				}
				log.Println("gw configurations = ", objectDictToReturn)

				cmeGWConfigurationsToReturn["result"] = objectDictToReturn
			}
		} else {
			cmeGWConfigurationsResultList, ok := cmeGWConfigurationsJson["result"]
			var objectDictToReturn []map[string]interface{}

			if ok {
				cmeGWConfigurationsResultJson := cmeGWConfigurationsResultList.(map[string]interface{})
				tempObject := readSingleConfiguration(cmeGWConfigurationsResultJson)
				objectDictToReturn = append(objectDictToReturn, tempObject)

				cmeGWConfigurationsToReturn["result"] = objectDictToReturn
			}
		}
	} else if cmeGWConfigurationsJson["error"] != nil {
		errorResult, ok := cmeGWConfigurationsJson["error"]

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

			cmeGWConfigurationsToReturn["error"] = tempObject
		}
	} else {
		cmeGWConfigurationsToReturn["result"] = map[string]interface{}{}
		cmeGWConfigurationsToReturn["error"] = map[string]interface{}{}
	}

	d.SetId(generateId())
	_ = d.Set("result", cmeGWConfigurationsToReturn["result"])
	_ = d.Set("error", cmeGWConfigurationsToReturn["error"])

	if has_error {
		return fmt.Errorf(err_message)
	}

	return nil
}

func readSingleConfiguration(cmeGWConfigurationsResultJson map[string]interface{}) map[string]interface{} {
	tempObject := make(map[string]interface{})
	blades := make(map[string]interface{})

	for key, value := range cmeGWConfigurationsResultJson {
		switch key {
		case "name", "version", "sic_key", "policy", "related_account":
			tempObject[key] = value.(string)
		default:
			if key != "one-time-password" {
				blades[key] = strconv.FormatBool(value.(bool))
				tempObject["blades"] = blades
			}
		}
	}

	return tempObject
}
