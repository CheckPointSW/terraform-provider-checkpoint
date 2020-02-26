package commands

import (
	"encoding/json"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"io/ioutil"
	"os"
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

func InitClient() (checkpoint.ApiClient, error) {

	//get credentials from Environment variables
	server := os.Getenv("CHECKPOINT_SERVER")
	username := os.Getenv("CHECKPOINT_USERNAME")
	password := os.Getenv("CHECKPOINT_PASSWORD")

	if server == "" || username == "" || password == "" {
		return checkpoint.ApiClient{}, fmt.Errorf("missing parameters to initialize api client - (server, username, password)")
	}

	//install policy - only on management api
	if val, ok := os.LookupEnv("CHECKPOINT_CONTEXT"); ok {
		if val == "gaia_api" {
			return checkpoint.ApiClient{}, fmt.Errorf("install-policy is valid only to management api")
		}
	}

	args := checkpoint.ApiClientArgs{
		Port:                    checkpoint.DefaultPort,
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
		Timeout:                 checkpoint.TimeOut,
		Sleep:                   checkpoint.SleepTime,
	}

	s, err := GetSession()
	if err != nil {
		return checkpoint.ApiClient{}, err
	}
	if s.Sid != "" {
		args.Sid = s.Sid
	} else {
		return checkpoint.ApiClient{}, fmt.Errorf("session id can't be empty")
	}

	mgmt := checkpoint.APIClient(args)

	return *mgmt, nil

}
