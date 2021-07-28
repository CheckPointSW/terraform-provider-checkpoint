package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
)

const (
	DefaultSessionFileName = "sid.json"
)

type Session struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
}

func GetSession(sessionFileName string) (Session, error) {
	if _, err := os.Stat(sessionFileName); os.IsNotExist(err) {
		_, err := os.Create(sessionFileName)
		if err != nil {
			return Session{}, err
		}
	}
	b, err := ioutil.ReadFile(sessionFileName)
	if err != nil {
		return Session{}, err
	}
	if len(b) == 0 {
		return Session{}, nil
	}
	var s Session
	if err = json.Unmarshal(b, &s); err != nil {
		return Session{}, err
	}
	return s, nil
}

func ResolveTaskId(data map[string]interface{}) interface{} {
	if data != nil {
		if v := data["tasks"]; v != nil {
			tasks := v.([]interface{})
			if len(tasks) > 0 {
				return tasks[0].(map[string]interface{})["task-id"]
			}
		}

		if v := data["task-id"]; v != nil {
			return v
		}
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

	sessionFileName := os.Getenv("CHECKPOINT_SESSIONFILENAME")
	if sessionFileName == "" {
		sessionFileName = DefaultSessionFileName
	}
	s, err := GetSession(sessionFileName)
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
