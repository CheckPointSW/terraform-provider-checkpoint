package api_go_sdk

import (
	"encoding/json"
	"net/http"
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
