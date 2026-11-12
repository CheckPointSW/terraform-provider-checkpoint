/*
APIClient.go
version 1.0

A library for communicating with Check Point's management server using Golang
written by: Check Point software technologies inc.
June 2019
tested with Check Point R81.20

-----------------------------------------------------------------------------

This is the main module, it contains all of the important command such as ApiCall, ApiQuery, etc.

*/

package api_go_sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	InProgress           string        = "in progress"
	DefaultPort          int           = 443
	Limit                int           = 50
	Filename             string        = "fingerprints.json"
	TimeOut              time.Duration = time.Second * 10
	SleepTime            time.Duration = time.Second * 2
	GaiaContext          string        = "gaia_api"
	WebContext           string        = "web_api"
	DefaultProxyPort                   = -1
	DefaultProxyHost                   = ""
	AutoPublishBatchSize int           = 100
)

// Check Point API Client (Management/GAIA)
type ApiClient struct {
	port                    int
	isPortDefault_          bool
	fingerprint             string
	sid                     string
	server                  string
	domain                  string
	proxyHost               string
	proxyPort               int
	isProxyUsed             bool
	apiVersion              string
	ignoreServerCertificate bool
	acceptServerCertificate bool
	debugFile               string
	httpDebugLevel          string
	context                 string
	timeout                 time.Duration
	sleep                   time.Duration
	userAgent               string
	cloudMgmtId             string
	autoPublishBatchSize    int
	activeCallsLock         sync.Mutex
	autoPublishLock         sync.Mutex
	totalCallsLock          sync.Mutex
	duringPublish           bool
	activeCallsCtr          int
	totalCallsCtr           int
}

// ApiClient constructor
// Input ApiClientArgs
// Returns new client instance
func APIClient(apiCA ApiClientArgs) *ApiClient {
	isPortDefault := false
	proxyUsed := false

	if apiCA.Port == -1 || apiCA.Port == DefaultPort {
		apiCA.Port = DefaultPort
		isPortDefault = true
	}
	if apiCA.ProxyPort != DefaultProxyPort && apiCA.ProxyHost != DefaultProxyHost {
		proxyUsed = true
	}

	// The context of using the client - defaults to web api
	if apiCA.Context == "" {
		apiCA.Context = WebContext
	}

	if apiCA.Timeout == -1 || apiCA.Timeout == TimeOut {
		apiCA.Timeout = TimeOut
	} else {
		apiCA.Timeout = apiCA.Timeout * time.Second
	}

	if apiCA.Sleep == -1 {
		apiCA.Sleep = SleepTime
	}

	if apiCA.UserAgent == "" {
		apiCA.UserAgent = "golang-api-wrapper"
	}

	return &ApiClient{
		port:                    apiCA.Port,
		isPortDefault_:          isPortDefault,
		fingerprint:             apiCA.Fingerprint,
		sid:                     apiCA.Sid,
		server:                  apiCA.Server,
		domain:                  "",
		proxyHost:               apiCA.ProxyHost,
		proxyPort:               apiCA.ProxyPort,
		isProxyUsed:             proxyUsed,
		apiVersion:              apiCA.ApiVersion,
		ignoreServerCertificate: apiCA.IgnoreServerCertificate,
		acceptServerCertificate: apiCA.AcceptServerCertificate,
		debugFile:               apiCA.DebugFile,
		httpDebugLevel:          apiCA.HttpDebugLevel,
		context:                 apiCA.Context,
		autoPublishBatchSize:    apiCA.AutoPublishBatchSize,
		timeout:                 apiCA.Timeout,
		sleep:                   apiCA.Sleep,
		userAgent:               apiCA.UserAgent,
		cloudMgmtId:             apiCA.CloudMgmtId,
	}
}

// Returns the port of API client
func (c *ApiClient) GetPort() int {
	return c.port
}

// Returns the context of API client
func (c *ApiClient) GetContext() string {
	return c.context
}

// Returns the fingerprint of API client
func (c *ApiClient) getFingerprint() string {
	return c.fingerprint
}

// Returns true if API port is set to default
func (c *ApiClient) IsPortDefault() bool {
	return c.isPortDefault_
}

// Returns true if client use proxy
func (c *ApiClient) IsProxyUsed() bool {
	return c.isProxyUsed
}

// Set API port
func (c *ApiClient) SetPort(portToSet int) {
	if portToSet == DefaultPort {
		c.isPortDefault_ = true
	} else {
		c.isPortDefault_ = false
	}
	c.port = portToSet
}

// Set API sleep time
func (c *ApiClient) SetSleepTime(sleepTime time.Duration) {
	c.sleep = sleepTime
}

// Set API client timeout
func (c *ApiClient) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// Returns session id
func (c *ApiClient) GetSessionID() string {
	return c.sid
}

// Returns number of batch size
func (c *ApiClient) GetAutoPublishBatchSize() int {
	return c.autoPublishBatchSize
}

func (c *ApiClient) SetAutoPublishBatchSize(autoPublishBatchSize int) {
	c.autoPublishBatchSize = autoPublishBatchSize
}

func (c *ApiClient) increaseActiveCalls() {
	c.activeCallsLock.Lock()
	c.activeCallsCtr++
	c.activeCallsLock.Unlock()
}

func (c *ApiClient) decreaseActiveCalls() {
	c.activeCallsLock.Lock()
	c.activeCallsCtr--
	c.activeCallsLock.Unlock()
}

func (c *ApiClient) ResetTotalCallsCounter() {
	c.totalCallsCtr = 0
}

func (c *ApiClient) DisableAutoPublish() {
	c.autoPublishBatchSize = -1
	c.totalCallsCtr = 0
}

// Deprecated: Do not use. Use ApiLogin instead
func (c *ApiClient) Login(username string, password string, continueLastSession bool, domain string, readOnly bool, payload string) (APIResponse, error) {
	credentials := map[string]interface{}{
		"user":     username,
		"password": password,
	}
	return c.commonLoginLogic(credentials, continueLastSession, domain, readOnly, make(map[string]interface{}))
}

// Deprecated: Do not use. Use ApiLoginWithApiKey instead
func (c *ApiClient) LoginWithApiKey(apiKey string, continueLastSession bool, domain string, readOnly bool, payload string) (APIResponse, error) {
	credentials := map[string]interface{}{
		"api-key": apiKey,
	}
	return c.commonLoginLogic(credentials, continueLastSession, domain, readOnly, make(map[string]interface{}))
}

/*
Performs login API call to the management server using username and password

username: Check Point admin name
password: Check Point admin password
continue_last_session: [optional] It is possible to continue the last Check Point session or to create a new one
domain: [optional] The name, UID or IP-Address of the domain to login.
read_only: [optional] Login with Read Only permissions. This parameter is not considered in case continue-last-session is true.
payload: [optional] More settings for the login command
returns: APIResponse, error
side-effects: updates the class's uid and server variables
*/
func (c *ApiClient) ApiLogin(username string, password string, continueLastSession bool, domain string, readOnly bool, payload map[string]interface{}) (APIResponse, error) {
	credentials := map[string]interface{}{
		"user":     username,
		"password": password,
	}
	return c.commonLoginLogic(credentials, continueLastSession, domain, readOnly, payload)
}

/*
Performs login API call to the management server using api key

api_key: Check Point api-key
continue_last_session: [optional] It is possible to continue the last Check Point session
or to create a new one
domain: [optional] The name, UID or IP-Address of the domain to login.
read_only: [optional] Login with Read Only permissions. This parameter is not considered in case
continue-last-session is true.
payload: [optional] More settings for the login command
returns: APIResponse object
side-effects: updates the class's uid and server variables
*/
func (c *ApiClient) ApiLoginWithApiKey(apiKey string, continueLastSession bool, domain string, readOnly bool, payload map[string]interface{}) (APIResponse, error) {
	credentials := map[string]interface{}{
		"api-key": apiKey,
	}
	return c.commonLoginLogic(credentials, continueLastSession, domain, readOnly, payload)
}

func (c *ApiClient) commonLoginLogic(credentials map[string]interface{}, continueLastSession bool, domain string, readOnly bool, payload map[string]interface{}) (APIResponse, error) {

	if c.context == WebContext {
		credentials["continue-last-session"] = continueLastSession
		credentials["read-only"] = readOnly
	}

	if domain != "" {
		credentials["domain"] = domain
	}

	if payload != nil {
		for k, v := range payload {
			credentials[k] = v
		}
	}

	loginRes, errCall := c.apiCall("login", credentials, "", false, c.IsProxyUsed(), true)
	if errCall != nil {
		return loginRes, errCall
	}
	if loginRes.Success {
		c.sid = loginRes.data["sid"].(string)
		c.domain = domain
		if c.apiVersion == "" {
			c.apiVersion = loginRes.data["api-server-version"].(string)
		}
	}

	return loginRes, nil
}

/*
Performs a web-service API request to the management server

command: the command is placed in the URL field
payload: a JSON object (or a string representing a JSON object) with the command arguments
sid: The Check Point session-id. when omitted use self.sid.
waitForTask: determines the behavior when the API server responds with a "task-id".

	by default, the function will periodically check the status of the task
	and will not return until the task is completed.
	when wait_for_task=False, it is up to the user to call the "show-task" API and check
	the status of the command.

useProxy: Determines if the user wants to use the proxy server and port provider.
method: HTTP request method - POST by default
return: APIResponse object
side-effects: updates the class's uid and server variables
*/
func (c *ApiClient) ApiCall(command string, payload map[string]interface{}, sid string, waitForTask bool, useProxy bool, method ...string) (APIResponse, error) {
	return c.apiCall(command, payload, sid, waitForTask, useProxy, false, method...)
}

func (c *ApiClient) ApiCallSimple(command string, payload map[string]interface{}) (APIResponse, error) {
	return c.apiCall(command, payload, c.sid, true, c.IsProxyUsed(), false)
}

func (c *ApiClient) apiCall(command string, payload map[string]interface{}, sid string, waitForTask bool, useProxy bool, internal bool, method ...string) (APIResponse, error) {
	reqMethod := "POST"
	if len(method) > 0 {
		providedMethod := method[0]
		if !isValidHTTPMethod(providedMethod) {
			return APIResponse{}, fmt.Errorf("invalid HTTP method: %s", providedMethod)
		}
		reqMethod = providedMethod
	}
	fp, errFP := getFingerprint(c.server, c.port)
	if errFP != nil {
		return APIResponse{}, errFP
	}

	c.fingerprint = fp
	fpAuthentication, err := c.CheckFingerprint()
	if !fpAuthentication {
		return APIResponse{}, errors.New("fingerprint doesn't match, someone might be trying to steal your information\n")
	}
	if err != nil {
		return APIResponse{}, err
	}

	if payload == nil {
		payload = map[string]interface{}{}
	}

	_data, err := json.Marshal(payload)
	if err != nil {
		return APIResponse{}, err
	}

	if sid == "" {
		sid = c.sid
	}

	var client *Client
	if useProxy {
		client, err = CreateProxyClient(c.server, c.proxyHost, sid, c.proxyPort, c.timeout)
		if err != nil {
			return APIResponse{}, err
		}
	} else {
		client, err = CreateClient(c.server, sid, c.timeout)
		if err != nil {
			return APIResponse{}, err
		}
	}

	url := "https://" + c.server + ":" + strconv.Itoa(c.port)

	if c.cloudMgmtId != "" {
		url += "/" + c.cloudMgmtId
	}

	url += "/" + c.context

	if c.apiVersion != "" {
		url += "/v" + c.apiVersion
	}

	url += "/" + command

	client.fingerprint = c.fingerprint

	client.SetDebugLevel(c.httpDebugLevel)

	spotReader := bytes.NewReader(_data)

	req, err := http.NewRequest(reqMethod, url, spotReader)
	if err != nil {
		return APIResponse{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", c.userAgent)
	req.Header.Set("Accept", "*/*")

	if command != "login" {
		req.Header.Set("X-chkp-sid", sid)
	}

	if !internal && c.autoPublishBatchSize > 0 {
		waitToRun := true
		for waitToRun {
			if c.totalCallsCtr+1 <= c.autoPublishBatchSize && !c.duringPublish {
				c.totalCallsLock.Lock()
				if c.totalCallsCtr+1 <= c.autoPublishBatchSize && !c.duringPublish {
					c.totalCallsCtr++
					waitToRun = false
				}
				c.totalCallsLock.Unlock()
			}
			if waitToRun {
				time.Sleep(time.Second)
			}
		}
		c.increaseActiveCalls()
	}

	response, err := client.client.Do(req)

	if err != nil {
		if !internal && c.autoPublishBatchSize > 0 {
			c.decreaseActiveCalls()
		}
		return APIResponse{}, err
	}

	res, err := fromHTTPResponse(response, "")
	if err != nil {
		if !internal && c.autoPublishBatchSize > 0 {
			c.decreaseActiveCalls()
		}
		return APIResponse{}, err
	}

	if !res.Success {
		res.ErrorMsg = res.buildGenericErrMsg()
	}

	if waitForTask == true && res.Success && command != "show-task" {
		if _, ok := res.data["task-id"]; ok {
			res, err = c.waitForTask(res.data["task-id"].(string))
			if err != nil {
				if !internal && c.autoPublishBatchSize > 0 {
					c.decreaseActiveCalls()
				}
				return APIResponse{}, err
			}
		} else if _, ok := res.data["tasks"]; ok {
			tasks := res.data["tasks"].([]interface{})
			if len(tasks) > 0 {
				res = c.waitForTasks(tasks)
			}
		}
	}

	if !internal && c.autoPublishBatchSize > 0 {
		c.decreaseActiveCalls()
		if c.totalCallsCtr > 0 && c.totalCallsCtr%c.autoPublishBatchSize == 0 && !c.duringPublish {
			c.autoPublishLock.Lock()
			if c.totalCallsCtr > 0 && c.totalCallsCtr%c.autoPublishBatchSize == 0 && !c.duringPublish {
				c.duringPublish = true
				c.autoPublishLock.Unlock()
				for c.activeCallsCtr > 0 {
					//	 Waiting for other calls to finish
					fmt.Println("Waiting to start auto publish (Active calls " + strconv.Itoa(c.activeCallsCtr) + ")")
					time.Sleep(time.Second)
				}
				// Going to publish
				fmt.Println("Start auto publish...")
				publishRes, _ := c.apiCall("publish", map[string]interface{}{}, c.GetSessionID(), true, c.IsProxyUsed(), true)

				if !publishRes.Success {
					fmt.Println("Auto publish failed. Message: " + publishRes.ErrorMsg)
				} else {
					fmt.Println("Auto publish finished successfully")
				}
				c.totalCallsCtr = 0
				c.duringPublish = false
			} else {
				c.autoPublishLock.Unlock()
			}
		}
	}

	return res, nil
}

/*
*
The APIs that return a list of objects are limited by the number of objects that they return.
To get the full list of objects, there's a need to make repeated API calls each time using a different offset
until all the objects are returned.
This API makes such repeated API calls and return the full list objects.
note: this function calls gen_api_query and iterates over the generator until it gets all the objects,
then returns.

command: name of API command. This command should be an API that returns an array of

	objects (for example: show-hosts, show networks, ...)

details_level: query APIs always take a details-level argument.

	possible values are "standard", "full", "uid"

container_key: name of the key that holds the objects in the JSON response (usually "objects").
include_container_key: If set to False the 'data' field of the APIResponse object

	will be a list of the wanted objects. Otherwise, the date field of the APIResponse will be a dictionary in the following

format: { container_key: [ List of the wanted objects], "total": size of the list}
payload: a JSON object (or a string representing a JSON object) with the command arguments
return: if include-container-key is False:

	an APIResponse object whose .data member contains a list of the objects requested: [ , , , ...]
	if include-container-key is True:
	an APIResponse object whose .data member contains a dict: { container_key: [...], "total": n }
*/
func (c *ApiClient) ApiQuery(command string, detailsLevel string, containerKey string, includeContainerKey bool, payload map[string]interface{}) (APIResponse, error) {

	var apiRes = APIResponse{}

	if containerKey == "" {
		containerKey = "objects"
	}

	containerKeys := []string{containerKey}
	var err error
	serverResponse := c.genApiQuery(command, detailsLevel, containerKeys, payload, &err)

	if err != nil {
		return APIResponse{}, err
	}

	if len(serverResponse) == 0 {

	} else {

		apiRes = serverResponse[len(serverResponse)-1]

		_, ok := apiRes.data[containerKey]
		if apiRes.Success && includeContainerKey == false && ok {

			m := map[string]interface{}{}
			for y, x := range apiRes.data[containerKey].([]interface{}) {

				m[fmt.Sprintf("%d", y)] = x
			}
			apiRes.data = m

		}
	}

	return apiRes, nil

}

/*
This is a generator function that yields the list of wanted objects received so far from the management server.
This is in contrast to normal API calls that return only a limited number of objects.
This function can be used to show progress when requesting many objects (i.e. "Received x/y objects.")

command: name of API command. This command should be an API that returns an array of objects

	(for example: show-hosts, show networks, ...)

details_level: query APIs always take a details-level argument. Possible values are "standard", "full", "uid"
container_keys: the field in the .data dict that contains the objects
payload: a JSON object (or a string representing a JSON object) with the command arguments
returns: an APIResponse object as detailed above
*/
func (c *ApiClient) genApiQuery(command string, detailsLevel string, containerKeys []string, payload map[string]interface{}, err_output *error) []APIResponse {

	const objLimit int = Limit
	var finished bool = false

	allObjects := map[string][]interface{}{}

	if len(containerKeys) == 0 {
		containerKeys = []string{"objects"}
	}

	for _, key := range containerKeys {
		allObjects[key] = []interface{}{}
	}

	iterations := 0
	if payload == nil {
		payload = map[string]interface{}{}
	}
	payload["limit"] = objLimit
	payload["offset"] = iterations * objLimit
	payload["details-level"] = detailsLevel
	apiRes, err := c.apiCall(command, payload, c.sid, false, c.IsProxyUsed(), true)

	if err != nil {
		print(err.Error())
	}

	var serverResponse []APIResponse

	for _, containerKey := range containerKeys {

		if apiRes.data == nil {
			print(containerKey)
		}
		_, ok := apiRes.data[containerKey]
		if !ok {
			finished = true

			serverResponse = append(serverResponse, apiRes)
			break
		}
	}

	for !finished {
		if !apiRes.Success {
			print("FAILED!\n")
			os.Exit(1)
		}

		totalObjects := apiRes.data["total"]
		receivedObjects := apiRes.data["to"]

		if receivedObjects == nil {
			receivedObjects = float64(0)
		}

		i := 0
		for _, containerKey := range containerKeys {

			for _, data := range (apiRes.data[containerKey]).([]interface{}) {
				allObjects[containerKey] = append(allObjects[containerKey], data)
			}
			apiRes.data[containerKey] = allObjects[containerKey]
			i++
		}

		serverResponse = append(serverResponse, apiRes)

		if totalObjects == receivedObjects {
			break
		}
		iterations += 1
		payload["limit"] = objLimit
		payload["offset"] = iterations * objLimit
		payload["details-level"] = detailsLevel
		apiRes, err = c.apiCall(command, payload, c.sid, false, c.IsProxyUsed(), true)

		if err != nil {
			print("Error communicating with server, please check your connection.")
			*err_output = err
			return nil
		}
	}
	*err_output = nil

	return serverResponse
}

/*
*
When the server needs to perform an API call that may take a long time (e.g. run-script, install-policy,
publish), the server responds with a 'task-id'.
Using the show-task API it is possible to check on the status of this task until its completion.
Every two seconds, this function will check for the status of the task.
The function will return when the task (and its sub-tasks) are no longer in-progress.

task_id: The task identifier.
return: APIResponse object (response of show-task command).
*/
func (c *ApiClient) waitForTask(taskId string) (APIResponse, error) {

	taskComplete := false
	var taskResult APIResponse
	var err error

	payload := map[string]interface{}{"task-id": taskId, "details-level": "full"}

	for !taskComplete {
		taskResult, err = c.apiCall("show-task", payload, c.sid, false, c.IsProxyUsed(), true)

		if err != nil {
			return APIResponse{}, err
		}

		attemptsCounter := 0

		for taskResult.Success == false {
			if attemptsCounter < 5 {
				attemptsCounter++
				time.Sleep(c.sleep)
				taskResult, err = c.apiCall("show-task", payload, c.sid, false, c.IsProxyUsed(), true)

				if err != nil {
					return APIResponse{}, err
				}

			} else {
				fmt.Println("ERROR: Failed to handle asynchronous tasks as synchronous, tasks result is undefined ", taskResult)
			}

		}

		completedTasks := 0
		totalTasks := 0
		for _, task := range taskResult.GetData()["tasks"].([]interface{}) {
			totalTasks++
			taskMap := task.(map[string]interface{})
			if taskMap["status"] != nil && taskMap["status"].(string) != InProgress {
				completedTasks++
			}
		}

		if completedTasks == totalTasks {
			taskComplete = true
		} else {
			time.Sleep(c.sleep)
		}

	}

	checkTasksStatus(&taskResult)
	return taskResult, nil

}

/*
*
The version of waitForTask function for the collection of tasks

task_objects: A list of task objects
return: APIResponse object (response of show-task command).
*/
func (c *ApiClient) waitForTasks(taskObjects []interface{}) APIResponse {

	var tasks []string
	for _, taskObj := range taskObjects {
		taskId := taskObj.(map[string]interface{})["task-id"]
		tasks = append(tasks, taskId.(string))
		c.waitForTask(taskId.(string))
	}

	payload := map[string]interface{}{
		"task-id":       tasks,
		"details-level": "full",
	}
	taskRes, err := c.apiCall("show-task", payload, c.GetSessionID(), false, c.IsProxyUsed(), true)

	if err != nil {
		fmt.Println("Problem showing tasks, try again")

	}
	checkTasksStatus(&taskRes)
	return taskRes

}

/*
*
This method checks if one of the tasks failed and if so, changes the response status to be False

task_result: api_response returned from "show-task" command
return:
*/
func checkTasksStatus(taskResult *APIResponse) {
	if v := taskResult.data["tasks"]; v != nil {
		for _, task := range taskResult.data["tasks"].([]interface{}) {
			if task.(map[string]interface{})["status"] == "failed" || task.(map[string]interface{})["status"] == "partially succeeded" {
				taskResult.setSuccessStatus(false)
				taskResult.StatusCode = ""
				taskResult.setErrMsg(taskResult.buildGenericErrMsg())
				break
			}
		}
	}
}

/*
   @===================@
   |  FINGERPRINT AREA |
   @===================@
*/

/*
*
This function checks if the server's certificate is stored in the local fingerprints file.
If the server's fingerprint is not found, an HTTPS connection is made to the server
and the user is asked if he or she accepts the server's fingerprint.
If the fingerprint is trusted, it is stored in the fingerprint file.

return: False if the user does not accept the server certificate, True in all other cases.
*/
func (c *ApiClient) CheckFingerprint() (bool, error) {

	if c.ignoreServerCertificate {
		return true, nil
	}

	//read the fingerprint form a local file
	var localFp, err = c.loadFingerprintFromFile()

	if err != nil {
		return false, err
	}

	var serverFp, errFP = getFingerprint(c.server, c.port)
	if errFP != nil {
		return false, errFP
	}

	if c.fingerprint == serverFp {
		return true, nil
	}

	if localFp == "" || strings.Replace(localFp, ":", "", -1) != strings.Replace(serverFp, ":", "", -1) {
		if serverFp == "" {
			return false, nil
		}

		if c.acceptServerCertificate {
			c.saveFingerprintToFile(c.server, c.fingerprint)
			return true, nil
		}

		if localFp == "" {
			fmt.Fprintf(os.Stderr, "You currently do not have a record of this server's fingerprint.\n")
		} else {
			fmt.Fprintf(os.Stderr, "The server's fingerprint is different from your local record of this server's fingerprint.\n You maybe a victim to a Man-in-the-Middle attack, please beware.\n")
		}
		fmt.Fprintf(os.Stderr, "Server's fingerprint: %s\n", (serverFp))

		if c.askYesOrNoQuestion("Do you accept this fingerprint?\n") {
			if c.saveFingerprintToFile(c.server, serverFp) == nil {
				fmt.Fprintf(os.Stderr, "Fingerprint saved.\n")
			} else {
				fmt.Fprintf(os.Stderr, "Could not save fingerprint to file. Continuing anyway.\n")
			}
		} else {
			return false, nil
		}
	}
	c.fingerprint = serverFp
	return true, nil
}

func (c *ApiClient) loadFingerprintFromFile() (string, error) {
	objmap, err := c.fpFileToMap()

	if err != nil {
		return "", err
	}

	//Objmap contains json data now
	if val, ok := objmap[c.server]; ok {
		return val, nil

	} else {
		err = c.saveFingerprintToFile(c.server, c.fingerprint)
		if err != nil {
			return "", err
		}
		return c.fingerprint, nil
	}

}

/*
*
This function takes the content of the file $FILENAME (which is a json file)
and parses it's content to a map (from string to string)

return: returns the map described above, error if happened
*/
func (c *ApiClient) fpFileToMap() (map[string]string, error) {

	//creates file if file doesn't exist
	c.createEmptyJsonFile(Filename)

	var data []byte
	var err error
	data, err = ioutil.ReadFile(Filename)
	if err != nil {
		return nil, err
	}
	//File opened
	var objmap map[string]string
	err = json.Unmarshal(data, &objmap)

	//Error occurs here
	if err != nil {
		return nil, err
	}
	//Process ends here

	//Objmap contains json data now
	return objmap, nil

}

/*
*
store a server's fingerprint into a local file.

server: the IP address/name of the Check Point management server.
fingerprint: A SHA1 fingerprint of the server's certificate.
filename: The file in which to store the certificates. The file will hold a JSON structure in which

	the key is the server and the value is its fingerprint.

return: 'True' if everything went well. 'False' if there was some kind of error storing the fingerprint.
*/
func (c *ApiClient) saveFingerprintToFile(server string, fingerprint string) error {

	objmap, err := c.fpFileToMap()
	if err != nil {
		return err
	}
	//Objmap contains json data now

	if val, ok := objmap[c.server]; ok {
		if val == fingerprint {
			return nil
		}
	}
	//File opened but does not contain server fp
	objmap[c.server] = fingerprint

	jsonmap, errJSON := json.Marshal(objmap)
	if errJSON != nil {
		return err
	}

	errWriting := ioutil.WriteFile(Filename, jsonmap, 0644)
	if errWriting != nil {
		return errWriting
	}
	return nil

}

/**
Simply creates a new empty json file with the name $name

return: error if happened
*/

func (c *ApiClient) createEmptyJsonFile(name string) error {

	// check if file exists
	var _, err = os.Stat(name)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(name)
		if err != nil {
			return err
		}
		defer file.Close()
		file.WriteString("{}")
	}
	return nil

}

/* @=========@
   |  Utils  |
   @=========@ */

func (c *ApiClient) askYesOrNoQuestion(question string) bool {
	fmt.Println(question)
	var answer string
	_, _ = fmt.Scanln(&answer)
	return strings.ToLower(answer) == "y" || strings.ToLower(answer) == "yes"
}

func isValidHTTPMethod(method string) bool {
	validMethods := []string{
		"GET", "POST", "PUT", "DELETE",
	}
	for _, validMethod := range validMethods {
		if method == validMethod {
			return true
		}
	}
	return false
}
