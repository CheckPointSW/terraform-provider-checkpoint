package api_go_sdk

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

const OkResponseCode string = "200 OK"

// API Response struct represent http response (Created from httpResponse struct)
type APIResponse struct {
	StatusCode string
	data       map[string]interface{}
	Success    bool
	ErrorMsg   string
	resObj     map[string]interface{}
}

/*
Generate APIResponse from httpResponse object

httpResponse: input HTTP response object
errMsg: if there is an error message included, we include it in the APIResponse
return: The APIResponse object we generated

*/
func fromHTTPResponse(httpResponse *http.Response, errMsg string) (APIResponse, error) {
	var data map[string]interface{}
	errResp := json.NewDecoder(httpResponse.Body).Decode(&data)
	defer httpResponse.Body.Close()
	if errResp != nil {
		return APIResponse{}, errResp
	}
	return APIResponse{httpResponse.Status, data, httpResponse.Status == OkResponseCode, errMsg, map[string]interface{}{}}, nil
}

// Get response data (payload)
func (r *APIResponse) GetData() map[string]interface{} {
	return r.data
}

// Get response data (payload)
func (r *APIResponse) GetResTmp() map[string]interface{} {
	return r.resObj
}

// Convert API Response to a map
func (r *APIResponse) asGoMap() map[string]interface{} {
	dict := map[string]interface{}{
		"res_obj":     r.resObj,
		"success":     r.Success,
		"status_code": r.StatusCode,
		"data":        r.data,
	}

	if r.ErrorMsg != "" {
		dict["error_message"] = r.ErrorMsg
	}
	return dict
}

/*
Set the response success status
status: input status
*/
func (r *APIResponse) setSuccessStatus(status bool) {
	r.Success = status
}

func (r *APIResponse) buildGenericErrMsg() string {
	response := r.GetData()
	errMsg := "Failed to execute API call"

	if tasks := response["tasks"]; tasks != nil {
		tasksList := tasks.([]interface{})
		if len(tasksList) > 0 {
			for i := range tasksList {
				task := tasksList[i].(map[string]interface{})
				errMsg += "\nTask: " + task["task-name"].(string) + "\nMessage: "
				if taskDetails := task["task-details"]; taskDetails != nil {
					taskDetailsList := taskDetails.([]interface{})
					if len(taskDetailsList) > 0 {
						for j := range taskDetailsList {
							if v := taskDetailsList[j].(map[string]interface{})["fault-message"]; v != nil {
								errMsg += v.(string)
							}
						}
					}
				}
			}
		}
	} else {
		resCode := ""
		resMsg := ""
		if code := response["code"]; code != nil {
			resCode = code.(string)
		}
		if msg := response["message"]; msg != nil {
			resMsg = msg.(string)
		}

		errMsg +=
			"\nStatus: " + r.StatusCode +
				"\nCode: " + resCode +
				"\nMessage: " + resMsg

		if errorMsg := response["errors"]; errorMsg != nil {
			errMsg += "\nErrors: "
			errorMsgType := reflect.TypeOf(errorMsg).Kind()
			if errorMsgType == reflect.String {
				errMsg += errorMsg.(string) + "\n"
			} else {
				errorsList := response["errors"].([]interface{})
				for i := range errorsList {
					errMsg += "\n" + strconv.Itoa(i+1) + ". " + errorsList[i].(map[string]interface{})["message"].(string)
				}
			}
		}

		if warningMsg := response["warnings"]; warningMsg != nil {
			errMsg += "\nWarnings: "
			warningMsgType := reflect.TypeOf(warningMsg).Kind()
			if warningMsgType == reflect.String {
				errMsg += warningMsg.(string) + "\n"
			} else {
				warningsList := response["warnings"].([]interface{})
				for i := range warningsList {
					errMsg += "\n" + strconv.Itoa(i+1) + ". " + warningsList[i].(map[string]interface{})["message"].(string)
				}
			}
		}

		if blockingError := response["blocking-errors"]; blockingError != nil {
			errMsg += "\nBlocking errors: "
			warningMsgType := reflect.TypeOf(blockingError).Kind()
			if warningMsgType == reflect.String {
				errMsg += blockingError.(string) + "\n"
			} else {
				blockingErrorsList := response["blocking-errors"].([]interface{})
				for i := range blockingErrorsList {
					errMsg += "\n" + strconv.Itoa(i+1) + ". " + blockingErrorsList[i].(map[string]interface{})["message"].(string)
				}
			}
		}
	}

	return errMsg
}

func (r *APIResponse) setErrMsg(message string) {
	r.ErrorMsg = message
}
