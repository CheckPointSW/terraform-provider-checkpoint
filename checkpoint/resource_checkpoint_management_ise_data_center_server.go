package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceManagementIseDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementIseDataCenterServer,
		Read:   readManagementIseDataCenterServer,
		Update: updateManagementIseDataCenterServer,
		Delete: deleteManagementIseDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"hostnames": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Address of ISE administrator hostnames.\nExample: http(s)://<host1 ip/url>.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User ID of the ISE administrator server.\nWhen using commonLoginLogic Domains use the following syntax:\napic:<domain>\\<username>.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Password of the ISE administrator server.",
			},
			"password_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Password of the Cisco ISE administrator encoded in Base64.",
			},
			"certificate_fingerprint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.",
			},
			"unsafe_auto_accept": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname.\n\nWhen set to true, trust the current Data Center Server's certificate as-is.",
				Default:     false,
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"color": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Color of the object. Should be one of existing colors.",
				Default:     "black",
			},
			"comments": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings. By Setting this parameter to 'true' test connection failure will be ignored.",
				Default:     false,
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
				Default:     false,
			},
		},
	}
}

func createManagementIseDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	iseDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		iseDataCenterServer["name"] = v.(string)
	}

	iseDataCenterServer["type"] = "ise"

	if v, ok := d.GetOk("hostnames"); ok {
		iseDataCenterServer["hostnames"] = v
	}

	if v, ok := d.GetOk("username"); ok {
		iseDataCenterServer["username"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		iseDataCenterServer["password"] = v.(string)
	}

	if v, ok := d.GetOk("password_base64"); ok {
		iseDataCenterServer["password-base64"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		iseDataCenterServer["certificate-fingerprint"] = v.(string)
	}

	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		iseDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		iseDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		iseDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		iseDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		iseDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		iseDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create iseDataCenterServer - Map = ", iseDataCenterServer)

	addIseDataCenterServerRes, err := client.ApiCall("add-data-center-server", iseDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addIseDataCenterServerRes.Success {
		if addIseDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addIseDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addIseDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": iseDataCenterServer["name"],
	}
	showIseDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIseDataCenterServerRes.Success {
		return fmt.Errorf(showIseDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showIseDataCenterServerRes.GetData()["uid"].(string))
	return readManagementIseDataCenterServer(d, m)
}

func readManagementIseDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showIseDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIseDataCenterServerRes.Success {
		if objectNotFound(showIseDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showIseDataCenterServerRes.ErrorMsg)
	}
	iseDataCenterServer := showIseDataCenterServerRes.GetData()

	if v := iseDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if iseDataCenterServer["properties"] != nil {
		propsJson, ok := iseDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				if propName == "unsafe_auto_accept" {
					propValue, _ = strconv.ParseBool(propValue.(string))
				}
				if propName == "hostnames" {
					propValue = strings.Split(propValue.(string), ";")
				}
				_ = d.Set(propName, propValue)
			}
		}
	}

	if iseDataCenterServer["tags"] != nil {
		tagsJson, ok := iseDataCenterServer["tags"].([]interface{})
		if ok {
			tagsIds := make([]string, 0)
			if len(tagsJson) > 0 {
				for _, tags := range tagsJson {
					tags := tags.(map[string]interface{})
					tagsIds = append(tagsIds, tags["name"].(string))
				}
			}
			_ = d.Set("tags", tagsIds)
		}
	} else {
		_ = d.Set("tags", nil)
	}

	if v := iseDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := iseDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := iseDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := iseDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementIseDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	iseDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		iseDataCenterServer["name"] = oldName
		iseDataCenterServer["new-name"] = newName
	} else {
		iseDataCenterServer["name"] = d.Get("name")
	}

	if d.HasChange("hostnames") {
		iseDataCenterServer["hostnames"] = d.Get("hostnames")
	}

	if d.HasChange("password") {
		iseDataCenterServer["password"] = d.Get("password")
	}

	if d.HasChange("password_base64") {
		iseDataCenterServer["password-base64"] = d.Get("password_base64")
	}

	if d.HasChange("username") {
		iseDataCenterServer["username"] = d.Get("username")
		if v := d.Get("password"); v != nil && v != "" {
			iseDataCenterServer["password"] = v
		}
		if v := d.Get("password_base64"); v != nil && v != "" {
			iseDataCenterServer["password-base64"] = v
		}
	}

	if d.HasChange("certificate_fingerprint") {
		iseDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}

	if d.HasChange("unsafe_auto_accept") {
		iseDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			iseDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			iseDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		iseDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		iseDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		iseDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		iseDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update iseDataCenterServer - Map = ", iseDataCenterServer)

	updateIseDataCenterServerRes, err := client.ApiCall("set-data-center-server", iseDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateIseDataCenterServerRes.Success {
		if updateIseDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateIseDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateIseDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementIseDataCenterServer(d, m)
}

func deleteManagementIseDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	iseDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete iseDataCenterServer")

	deleteIseDataCenterServerRes, err := client.ApiCall("delete-data-center-server", iseDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteIseDataCenterServerRes.Success {
		if deleteIseDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteIseDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
