package main

import (
	chkp "api_go_sdk/APIFiles"
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	FILENAME = "sid.json"
)


func logout(client *chkp.ApiClient) error {
	logoutRes, err := client.ApiCall("logout",map[string]interface{}{}, client.GetSessionID(), true, false)
	if !logoutRes.Success{
		return err
	}
	return nil
}

type Session struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
}

func (s Session) Save() error {
	f, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(FILENAME, f, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetSid() (Session,error) {
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
		return Session{},err
	}
	return s,nil
}
func CheckSession(c *chkp.ApiClient, uid string) bool {
	if uid == "" || c.GetContext() != chkp.WebContext{
		return false
	}
	payload := map[string]interface{}{
		"uid": uid,
	}
	res, _ := c.ApiCall("show-session",payload,c.GetSessionID(),true,false)
	return res.Success
}
