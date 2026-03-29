package checkpoint

import (
	"encoding/json"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

//var lock sync.Mutex

const (
	DefaultSessionFilename = "sid.json"
)

type Session struct {
	Sid string `json:"sid"`
	Uid string `json:"uid"`
}

func (s *Session) Save(sessionFileName string) error {
	f, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(sessionFileName, f, 0644)
	if err != nil {
		return err
	}
	return nil
}

func GetSession(sessionFileName string) (Session, error) {
	if _, err := os.Stat(sessionFileName); os.IsNotExist(err) {
		_, err := os.Create(sessionFileName)
		if err != nil {
			return Session{}, err
		}
	}
	b, err := ioutil.ReadFile(sessionFileName)
	if err != nil || len(b) == 0 {
		return Session{}, err
	}
	var s Session
	if err = json.Unmarshal(b, &s); err != nil {
		return Session{}, err
	}
	return s, nil
}

func CheckSession(c *checkpoint.ApiClient, uid string) bool {
	if uid == "" || c.GetContext() != checkpoint.WebContext {
		return false
	}
	payload := map[string]interface{}{
		"uid": uid,
	}
	res, _ := c.ApiCall("show-session", payload, c.GetSessionID(), true, c.IsProxyUsed())
	return res.Success
}

func Compare(a, b []string) []string {
	for i := len(a) - 1; i >= 0; i-- {
		for _, vD := range b {
			if a[i] == vD {
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
	return a
}

func resolveTaskId(data map[string]interface{}) interface{} {
	if v := data["tasks"]; v != nil {
		if v, ok := v.([]interface{}); ok {
			if len(v) == 1 {
				return v[0].(map[string]interface{})["task-id"]
			}
		}
	}
	if v := data["task-id"]; v != nil {
		return v
	}
	return nil
}

func resolveTaskIds(data map[string]interface{}) []interface{} {
	if data["tasks"] != nil {
		if tasksJson, ok := data["tasks"].([]interface{}); ok {
			tasksIds := make([]interface{}, 0)
			if len(tasksJson) > 0 {
				for _, task := range tasksJson {
					taskJson := task.(map[string]interface{})
					tasksIds = append(tasksIds, taskJson["task-id"])
				}
			}
			return tasksIds
		}
	}
	return nil
}

func createTaskFailMessage(command string, data map[string]interface{}) string {
	msg := fmt.Sprintf("fail to %s.", command)
	if data != nil {
		if v, ok := data["tasks"].([]interface{}); ok {
			if len(v) > 0 {
				task := v[0].(map[string]interface{})
				msg += fmt.Sprintf(" task-id [%s]", task["task-id"])
				if task["status"] != "succeeded" {
					if len(task["task-details"].([]interface{})) > 0 {
						myTask := task["task-details"].([]interface{})[0].(map[string]interface{})
						if v, ok := myTask["fault-message"]; ok {
							msg += "\nMessage: " + v.(string)
						}
						if v, ok := myTask["statusDescription"]; ok {
							msg += "\nDescription: " + v.(string)
						}
					}
				}
			}
		}
	}
	return msg
}

// converts object type to source for machines and users.
func getTypeToSource() map[string]string {
	TypeToSource := map[string]string{
		"identity-tag":      "Identity Tag",
		"user-group":        "Internal User Groups",
		"CpmiAnyObject":     "Guests",
		"CpmiExternalGroup": "LDAP groups",
	}
	return TypeToSource
}

func getKeysToFixedKeys() map[string]string {
	KeysToFixedKeys := map[string]string{
		"PREDEFINED":          "predefined",
		"Type in Data Center": "type-in-data-center",
		"Name in Data Center": "name-in-data-center",
		"IP Address":          "ip-address",
		"TAG":                 "tag",
	}
	return KeysToFixedKeys
}
func isArgDefault(v string, d *schema.ResourceData, arg string, defaultVal string) bool {
	_, ok := d.GetOk(arg)
	isDefault := v == defaultVal && ok
	return v != defaultVal || isDefault
}

func resolveListOfIdentifiers(fieldName string, jsonResponse interface{}, d *schema.ResourceData) []string {
	res := make([]string, 0)
	key := "name" // by default we use name as object identifier

	if v, ok := d.GetOk("fields_with_uid_identifier"); ok {
		fieldsSupportUidList := v.(*schema.Set).List()
		if len(fieldsSupportUidList) > 0 {
			for _, field := range fieldsSupportUidList {
				if field == fieldName {
					key = "uid"
					break
				}
			}
		}
	}

	if arr, ok := jsonResponse.([]interface{}); ok {
		if len(arr) > 0 {
			for _, obj := range arr {
				res = append(res, obj.(map[string]interface{})[key].(string))
			}
		}
	} else {
		if obj, ok := jsonResponse.(map[string]interface{}); ok {
			res = append(res, obj[key].(string))
		}
	}

	return res
}

func resolveObjectIdentifier(fieldName string, jsonResponse interface{}, d *schema.ResourceData) string {
	return resolveListOfIdentifiers(fieldName, jsonResponse, d)[0]
}

// removing prefix. suffix and '\n' that return with the cert from the server.
func cleanseCertificate(cert string) string {

	cert = strings.TrimPrefix(cert, "-----BEGIN CERTIFICATE-----\n")
	cert = strings.TrimSuffix(cert, "\n-----END CERTIFICATE-----\n")
	cert = strings.ReplaceAll(cert, "\n", "")

	return cert
}

func removeCnPrefix(issueBy string) string {

	issueBy = strings.TrimPrefix(issueBy, "CN=")
	return issueBy

}

func convertDateFormat(dateStr string) (string, error) {
	inputLayout := "02-Jan-06"
	outputLayout := "2006-01-02"

	// Parse the input date string using the input layout
	t, err := time.Parse(inputLayout, dateStr)
	if err != nil {
		fmt.Println("Error:", err)
		return "", err
	}

	// Format the parsed time using the output layout
	return t.Format(outputLayout), nil
}
