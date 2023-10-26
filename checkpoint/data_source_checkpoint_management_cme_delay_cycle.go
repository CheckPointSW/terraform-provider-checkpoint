package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementCMEDelayCycle() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEDelayCycleRead,
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
						"delay_cycle": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The delay cycle number.",
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

func dataSourceManagementCMEDelayCycleRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	cmeDelayCycleRes, err := client.ApiCall("cme-api/v1/generalConfiguration/delayCycle", nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cmeDelayCycleRes.Success {
		return fmt.Errorf(cmeDelayCycleRes.ErrorMsg)
	}

	cmeDelayCycleJson := cmeDelayCycleRes.GetData()
	log.Println("Read cme delay cycle - Show JSON = ", cmeDelayCycleJson)

	cmeDelayCycleToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if v := cmeDelayCycleJson["status-code"]; v != nil {
		_ = d.Set("status_code", int(math.Round(v.(float64))))
	}

	if cmeDelayCycleJson["result"] != nil {
		cmeDelayCycleResult, ok := cmeDelayCycleJson["result"]

		if ok {
			cmeDelayCycleResultJson := cmeDelayCycleResult.(map[string]interface{})
			tempObject := make(map[string]interface{})

			if v := cmeDelayCycleResultJson["delay_cycle"]; v != nil {
				var delay_cycle string = strconv.Itoa(int(math.Round(v.(float64))))
				tempObject["delay_cycle"] = delay_cycle
			}

			cmeDelayCycleToReturn["result"] = tempObject
		}
	} else if cmeDelayCycleJson["error"] != nil {
		errorResult, ok := cmeDelayCycleJson["error"]

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

			cmeDelayCycleToReturn["error"] = tempObject
		}
	} else {
		cmeDelayCycleToReturn["result"] = map[string]interface{}{}
		cmeDelayCycleToReturn["error"] = map[string]interface{}{}
	}

	d.SetId(generateId())
	_ = d.Set("result", cmeDelayCycleToReturn["result"])
	_ = d.Set("error", cmeDelayCycleToReturn["error"])

	if has_error {
		return fmt.Errorf(err_message)
	}

	return nil
}
