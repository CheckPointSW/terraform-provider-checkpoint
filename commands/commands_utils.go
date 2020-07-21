package commands

import (
	"encoding/json"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const (
	FILENAME = "sid.json"
)

type Session struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
}

func GetSession() (Session, error) {
	if _, err := os.Stat(FILENAME); os.IsNotExist(err) {
		_, err := os.Create(FILENAME)
		if err != nil {
			return Session{}, err
		}
	}
	b, err := ioutil.ReadFile(FILENAME)
	if err != nil || len(b) == 0 {
		return Session{}, err
	}
	var s Session
	if err = json.Unmarshal(b, &s); err != nil {
		return Session{}, err
	}
	return s, nil
}

func LogToFile(filename string, msg string) error {
	fullMsg := "[" + time.Now().String() + "] " + msg
	err := ioutil.WriteFile(filename, []byte(fullMsg), 0644)
	if err != nil {
		return err
	}
	return nil
}

func InitClient() (checkpoint.ApiClient, error) {
	// Default values
	port := checkpoint.DefaultPort
	timeout := checkpoint.TimeOut

	// Get credentials from Environment variables
	server := os.Getenv("CHECKPOINT_SERVER")
	username := os.Getenv("CHECKPOINT_USERNAME")
	password := os.Getenv("CHECKPOINT_PASSWORD")
	portVal := os.Getenv("CHECKPOINT_PORT")
	timeoutVal := os.Getenv("CHECKPOINT_TIMEOUT")

	var err error
	if portVal != "" {
		port, err = strconv.Atoi(portVal)
		if err != nil {
			return checkpoint.ApiClient{}, fmt.Errorf("failed to parse CHECKPOINT_PORT to integer")
		}
	}

	if timeoutVal != "" {
		timeoutInteger, err := strconv.Atoi(timeoutVal)
		if err != nil {
			return checkpoint.ApiClient{}, fmt.Errorf("failed to parse CHECKPOINT_TIMEOUT to integer")
		}
		timeout = time.Duration(timeoutInteger)
	}

	if server == "" || username == "" || password == "" {
		return checkpoint.ApiClient{}, fmt.Errorf("missing at least one required parameter to initialize API client (CHECKPOINT_SERVER, CHECKPOINT_USERNAME, CHECKPOINT_PASSWORD)")
	}

	// install policy/publish - only on management api
	if val, ok := os.LookupEnv("CHECKPOINT_CONTEXT"); ok {
		if val == "gaia_api" {
			return checkpoint.ApiClient{}, fmt.Errorf("install-policy/publish is valid only on management api (CHECKPOINT_CONTEXT = gaia_api)")
		}
	}

	args := checkpoint.ApiClientArgs{
		Port:                    port,
		Fingerprint:             "",
		Sid:                     "",
		Server:                  server,
		ProxyHost:               "",
		ProxyPort:               -1,
		ApiVersion:              "",
		IgnoreServerCertificate: false,
		AcceptServerCertificate: false,
		DebugFile:               "deb.txt",
		Context:                 "web_api",
		Timeout:                 timeout,
		Sleep:                   checkpoint.SleepTime,
	}

	s, err := GetSession()
	if err != nil {
		return checkpoint.ApiClient{}, err
	}
	if s.Sid != "" {
		args.Sid = s.Sid
	} else {
		return checkpoint.ApiClient{}, fmt.Errorf("session id not found. Verify 'sid.json' file exists in working directory")
	}

	mgmt := checkpoint.APIClient(args)

	return *mgmt, nil
}
