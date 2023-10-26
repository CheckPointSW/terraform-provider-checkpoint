package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceManagementCMEVersion() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementCMEVersionRead,
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
						"take": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Take number of the installed CME.",
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

func dataSourceManagementCMEVersionRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	cmeVersionRes, err := client.ApiCall("cme-api/v1/generalConfiguration/cmeVersion", nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cmeVersionRes.Success {
		return fmt.Errorf(cmeVersionRes.ErrorMsg)
	}
	cmeVersionJson := cmeVersionRes.GetData()
	log.Println("Read cme version - Show JSON = ", cmeVersionJson)

	cmeVersionToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if v := cmeVersionJson["status-code"]; v != nil {
		_ = d.Set("status_code", int(math.Round(v.(float64))))
	}

	if cmeVersionJson["result"] != nil {
		cmeVersionResult, ok := cmeVersionJson["result"]

		if ok {
			cmeVersionResultJson := cmeVersionResult.(map[string]interface{})
			tempObject := make(map[string]interface{})

			if v := cmeVersionResultJson["take"]; v != nil {
				var take string = strconv.Itoa(int(math.Round(v.(float64))))
				tempObject["take"] = take
			}

			cmeVersionToReturn["result"] = tempObject
		}
	} else if cmeVersionJson["error"] != nil {
		errorResult, ok := cmeVersionJson["error"]

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

			cmeVersionToReturn["error"] = tempObject
		}
	} else {
		cmeVersionToReturn["result"] = map[string]interface{}{}
		cmeVersionToReturn["error"] = map[string]interface{}{}
	}

	d.SetId(generateId())
	_ = d.Set("result", cmeVersionToReturn["result"])
	_ = d.Set("error", cmeVersionToReturn["error"])

	if has_error {
		return fmt.Errorf(err_message)
	}

	return nil
}
