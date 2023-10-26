package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementCMEGWConfigurations() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCMEGWConfiguration,
		Update: updateManagementCMEGWConfiguration,
		Read:   readManagementCMEGWConfiguration,
		Delete: deleteManagementCMEGWConfiguration,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GW configuration.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The GW configuration version.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"base64_sic_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Base64 SIC key.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The policy type.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"related_account": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The CME account name.",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v == "" {
						errs = append(errs, fmt.Errorf("%v must not be an empty string", key))
					}
					return
				},
			},
			"blades": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "The blades to set active for this GW configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{},
				},
			},
			"platform": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GW configuration platform (azure, gcp, aws).",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != "azure" && v != "gcp" && v != "aws" {
						errs = append(errs, fmt.Errorf("%v must be either one of the following: azure, gcp, aws", key))
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

func createManagementCMEGWConfiguration(d *schema.ResourceData, m interface{}) error {
	return createUpdateGWConfiguration(d, m)
}

func updateManagementCMEGWConfiguration(d *schema.ResourceData, m interface{}) error {
	return createUpdateGWConfiguration(d, m)
}

func readManagementCMEGWConfiguration(d *schema.ResourceData, m interface{}) error {
	d.SetId(generateId())
	_ = d.Set("status_code", 200)
	return nil
}

func deleteManagementCMEGWConfiguration(d *schema.ResourceData, m interface{}) error {
	var name string
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}

	res, err := deleteGWConfiguration(name, m)

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

func createUpdateGWConfiguration(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var name, platform string
	var method string = "POST"

	url := "cme-api/v1/gwConfigurations"

	payload := make(map[string]interface{})
	if v, ok := d.GetOk("name"); ok {
		name = v.(string)
	}
	if v, ok := d.GetOk("platform"); ok {
		platform = v.(string)
		url += "/" + platform
	} else {
		return fmt.Errorf("platform must be declared and have either one of the following: azure, gcp, aws")
	}

	isExist, err := checkGWConfigurationExisting(name, m)

	if err != nil {
		return fmt.Errorf(err.Error())
	}

	var updateGWConfiguration bool = false
	if isExist {
		method = "PUT"
		url += "/" + name
		updateGWConfiguration = true
	} else {
		payload["name"] = name
	}

	if v, ok := d.GetOk("version"); ok {
		payload["version"] = v.(string)
	} else if !ok && !updateGWConfiguration {
		return fmt.Errorf("expected version when creating new gw configuration")
	}
	if v, ok := d.GetOk("base64_sic_key"); ok {
		payload["base64_sic_key"] = v.(string)
	} else if !ok && !updateGWConfiguration {
		return fmt.Errorf("expected base64 sic key when creating new gw configuration")
	}
	if v, ok := d.GetOk("policy"); ok {
		payload["policy"] = v.(string)
	} else if !ok && !updateGWConfiguration {
		return fmt.Errorf("expected policy when creating new gw configuration")
	}
	if v, ok := d.GetOk("related_account"); ok {
		payload["related_account"] = v.(string)
	} else if !ok && !updateGWConfiguration {
		return fmt.Errorf("expected related account when creating new gw configuration")
	}
	if v, ok := d.GetOk("blades"); ok {
		tempObject := make(map[string]interface{})
		for k, v := range v.(map[string]interface{}) {
			tempObject[k], err = strconv.ParseBool(v.(string))
			if err != nil {
				return fmt.Errorf("expected boolean value instead got %v", v)
			}
		}

		payload["blades"] = tempObject
	}

	log.Println("Set cme GW configuration - Map = ", payload)

	cmeGWConfigurationRes, err := client.ApiCall(url, payload, client.GetSessionID(), true, client.IsProxyUsed(), method)

	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !cmeGWConfigurationRes.Success {
		return fmt.Errorf(cmeGWConfigurationRes.ErrorMsg)
	}

	cmeGWConfigurationJson := cmeGWConfigurationRes.GetData()
	cmeGWConfigurationToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if cmeGWConfigurationJson["error"] != nil {
		errorResult, ok := cmeGWConfigurationJson["error"]

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

			cmeGWConfigurationToReturn["error"] = tempObject
		}
	} else {
		cmeGWConfigurationToReturn["result"] = map[string]interface{}{}
		cmeGWConfigurationToReturn["error"] = map[string]interface{}{}
	}

	_ = d.Set("result", cmeGWConfigurationToReturn["result"])
	_ = d.Set("error", cmeGWConfigurationToReturn["error"])

	if has_error {
		return fmt.Errorf(err_message)
	}

	return readManagementCMEGWConfiguration(d, m)
}
