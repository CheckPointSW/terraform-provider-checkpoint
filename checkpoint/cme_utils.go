package checkpoint

import (
	"fmt"
	"log"
	"math"
	"strconv"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
)

func checkAccountExisting(name string, m interface{}) (bool, error) {
	client := m.(*checkpoint.ApiClient)

	res, err := client.ApiCall("cme-api/v1/accounts/"+name, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	if !res.Success {
		return false, fmt.Errorf(res.ErrorMsg)
	}

	resJson := res.GetData()

	if resJson["result"] != nil {
		results, ok := resJson["result"]
		if ok {
			resultsJson := results.(map[string]interface{})
			if len(resultsJson) != 0 {
				return true, nil
			}
		}
	}

	return false, nil
}

func deleteAccount(name string, m interface{}) (map[string]interface{}, error) {
	client := m.(*checkpoint.ApiClient)

	isExist, err := checkAccountExisting(name, m)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if !isExist {
		return nil, fmt.Errorf("account %v is not exist", name)
	}

	url := "cme-api/v1/accounts/" + name

	log.Println("Delete cme account - Name = ", name)

	res, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "DELETE")

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	if !res.Success {
		return nil, fmt.Errorf(res.ErrorMsg)
	}

	resJson := res.GetData()
	resToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if errorResult, ok := resJson["error"]; ok {
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

		resToReturn["error"] = tempObject
	} else {
		resToReturn["result"] = map[string]interface{}{}
		resToReturn["error"] = map[string]interface{}{}
	}

	if has_error {
		return resToReturn, fmt.Errorf(err_message)
	}

	return resToReturn, nil
}

func checkGWConfigurationExisting(name string, m interface{}) (bool, error) {
	client := m.(*checkpoint.ApiClient)

	res, err := client.ApiCall("cme-api/v1/gwConfigurations/"+name, nil, client.GetSessionID(), true, client.IsProxyUsed(), "GET")

	if err != nil {
		return false, fmt.Errorf(err.Error())
	}
	if !res.Success {
		return false, fmt.Errorf(res.ErrorMsg)
	}

	resJson := res.GetData()

	if resJson["result"] != nil {
		results, ok := resJson["result"]
		if ok {
			resultsJson := results.(map[string]interface{})
			if len(resultsJson) != 0 {
				return true, nil
			}
		}
	}

	return false, nil
}

func deleteGWConfiguration(name string, m interface{}) (map[string]interface{}, error) {
	client := m.(*checkpoint.ApiClient)

	isExist, err := checkGWConfigurationExisting(name, m)

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if !isExist {
		return nil, fmt.Errorf("gw configuration %v is not exist", name)
	}

	url := "cme-api/v1/gwConfigurations/" + name

	log.Println("Delete cme GW configuration - Name = ", name)

	res, err := client.ApiCall(url, nil, client.GetSessionID(), true, client.IsProxyUsed(), "DELETE")

	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}
	if !res.Success {
		return nil, fmt.Errorf(res.ErrorMsg)
	}

	resJson := res.GetData()
	resToReturn := make(map[string]interface{})

	var has_error bool = false
	var err_message string

	if errorResult, ok := resJson["error"]; ok {
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

		resToReturn["error"] = tempObject
	} else {
		resToReturn["result"] = map[string]interface{}{}
		resToReturn["error"] = map[string]interface{}{}
	}

	if has_error {
		return resToReturn, fmt.Errorf(err_message)
	}

	return resToReturn, nil
}
