package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementCMEManagement() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEManagement,
		Update: updateManagementCMEManagement,
		Read:   readManagementCMEManagement,
		Delete: deleteManagementCMEManagement,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "New name to the management.",
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

func createManagementCMEManagement(d *schema.ResourceData, m interface{}) error {
	return updateManagementCMEManagement(d, m)
}

func updateManagementCMEManagement(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	log.Println("Update cme management - Map = ", payload)

	cmeManagementRes, err := client.ApiCall("cme-api/v1/management", payload, client.GetSessionID(), true, client.IsProxyUsed(), "PUT")

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cmeManagementRes.Success {
		return fmt.Errorf(cmeManagementRes.ErrorMsg)
	}

	cmeManagementJson := cmeManagementRes.GetData()
	cmeManagementToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if cmeManagementJson["error"] != nil {
		errorResult, ok := cmeManagementJson["error"]

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

			cmeManagementToReturn["error"] = tempObject
		}
	} else {
		cmeManagementToReturn["result"] = map[string]interface{}{}
		cmeManagementToReturn["error"] = map[string]interface{}{}
	}

	_ = d.Set("result", cmeManagementToReturn["result"])
	_ = d.Set("error", cmeManagementToReturn["error"])

	if has_error {
		return fmt.Errorf(err_message)
	}

	return readManagementCMEManagement(d, m)
}

func readManagementCMEManagement(d *schema.ResourceData, m interface{}) error {
	d.SetId(generateId())
	_ = d.Set("status_code", 200)
	return nil
}

func deleteManagementCMEManagement(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
