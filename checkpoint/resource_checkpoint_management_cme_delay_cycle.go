package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementCMEDelayCycle() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEDelayCycle,
		Update: updateManagementCMEDelayCycle,
		Read:   readManagementCMEDelayCycle,
		Delete: deleteManagementCMEDelayCycle,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"delay_cycle": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "New delay cycle to set.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v <= 0 {
						errs = append(errs, fmt.Errorf("%v must be between a positive number. Got: %v", key, v))
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
							Description: "Error details.",
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

func createManagementCMEDelayCycle(d *schema.ResourceData, m interface{}) error {
	return updateManagementCMEDelayCycle(d, m)
}

func updateManagementCMEDelayCycle(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})
	if v, ok := d.GetOk("delay_cycle"); ok {
		payload["delay_cycle"] = v.(int)
	}

	log.Println("Update cme delay cycle - Map = ", payload)

	cmeDelayCycleRes, err := client.ApiCall("cme-api/v1/generalConfiguration/delayCycle", payload, client.GetSessionID(), true, client.IsProxyUsed(), "PUT")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cmeDelayCycleRes.Success {
		return fmt.Errorf(cmeDelayCycleRes.ErrorMsg)
	}

	cmeDelayCycleJson := cmeDelayCycleRes.GetData()
	cmeDelayCycleToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if cmeDelayCycleJson["error"] != nil {
		errorResult, ok := cmeDelayCycleJson["error"]

		if ok {
			errorResultJson := errorResult.(map[string]interface{})
			tempObject := make(map[string]interface{})

			if v := errorResultJson["details"]; v != nil {
				tempObject["details"] = v.(string)
				has_error = true
			}
			if v := errorResultJson["error_code"]; v != nil {
				var error_code string = strconv.Itoa(int(math.Round(v.(float64))))
				tempObject["error_code"] = error_code
				has_error = true
			}
			if v := errorResultJson["message"]; v != nil {
				tempObject["message"] = v
				err_message = v.(string)
				has_error = true
			}

			cmeDelayCycleToReturn["error"] = tempObject
		}
	} else {
		cmeDelayCycleToReturn["result"] = map[string]interface{}{}
		cmeDelayCycleToReturn["error"] = map[string]interface{}{}
	}

	_ = d.Set("result", cmeDelayCycleToReturn["result"])
	_ = d.Set("error", cmeDelayCycleToReturn["error"])

	if has_error {
		return fmt.Errorf(err_message)
	}

	return readManagementCMEDelayCycle(d, m)
}

func readManagementCMEDelayCycle(d *schema.ResourceData, m interface{}) error {
	d.SetId(generateId())
	_ = d.Set("status_code", 200)
	return nil
}

func deleteManagementCMEDelayCycle(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
