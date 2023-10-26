package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementCMEManagement() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEManagementRead,
		Schema: map[string]*schema.Schema{
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
					Schema: map[string]*schema.Schema{
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The host address.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the management.",
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

func dataSourceManagementCMEManagementRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	cmeManagementRes, err := client.ApiCall("cme-api/v1/management", nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cmeManagementRes.Success {
		return fmt.Errorf(cmeManagementRes.ErrorMsg)
	}

	cmeManagementJson := cmeManagementRes.GetData()
	log.Println("Read cme management - Show JSON = ", cmeManagementJson)

	cmeManagementToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if v := cmeManagementJson["status-code"]; v != nil {
		_ = d.Set("status_code", int(math.Round(v.(float64))))
	}

	if cmeManagementJson["result"] != nil {
		cmeManagementResult, ok := cmeManagementJson["result"]

		if ok {
			cmeVersionResultJson := cmeManagementResult.(map[string]interface{})
			tempObject := make(map[string]interface{})

			if v := cmeVersionResultJson["host"]; v != nil {
				tempObject["host"] = v.(string)
			}
			if v := cmeVersionResultJson["name"]; v != nil {
				tempObject["name"] = v.(string)
			}

			cmeManagementToReturn["result"] = tempObject
		}
	} else if cmeManagementJson["error"] != nil {
		errorResult, ok := cmeManagementJson["error"]

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

			cmeManagementToReturn["error"] = tempObject
		}
	} else {
		cmeManagementToReturn["result"] = map[string]interface{}{}
		cmeManagementToReturn["error"] = map[string]interface{}{}
	}

	d.SetId(generateId())
	_ = d.Set("result", cmeManagementToReturn["result"])
	_ = d.Set("error", cmeManagementToReturn["error"])

	if has_error {
		return fmt.Errorf(err_message)
	}

	return nil
}
