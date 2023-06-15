package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceManagementOpenStackDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementOpenStackDataCenterServer,
		Read:   readManagementOpenStackDataCenterServer,
		Update: updateManagementOpenStackDataCenterServer,
		Delete: deleteManagementOpenStackDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "URL of the OpenStack server.\nhttp(s)://<host>:<port>/<version>\nExample: https://1.2.3.4:5000/v2.0",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Username of the OpenStack server.\nTo login to specific domain insert domain name before username.\nExample: <domain>/<username>",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the OpenStack server.",
			},
			"password_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the OpenStack server encoded in Base64.",
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

func createManagementOpenStackDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	openstackDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		openstackDataCenterServer["name"] = v.(string)
	}

	openstackDataCenterServer["type"] = "openstack"

	if v, ok := d.GetOk("hostname"); ok {
		openstackDataCenterServer["hostname"] = v.(string)
	}

	if v, ok := d.GetOk("username"); ok {
		openstackDataCenterServer["username"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		openstackDataCenterServer["password"] = v.(string)
	}

	if v, ok := d.GetOk("password_base64"); ok {
		openstackDataCenterServer["password-base64"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		openstackDataCenterServer["certificate-fingerprint"] = v.(string)
	}

	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		openstackDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		openstackDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		openstackDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		openstackDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		openstackDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		openstackDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create openstackDataCenterServer - Map = ", openstackDataCenterServer)

	addOpenStackDataCenterServerRes, err := client.ApiCall("add-data-center-server", openstackDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addOpenStackDataCenterServerRes.Success {
		if addOpenStackDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addOpenStackDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addOpenStackDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": openstackDataCenterServer["name"],
	}
	showOpenStackDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showOpenStackDataCenterServerRes.Success {
		return fmt.Errorf(showOpenStackDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showOpenStackDataCenterServerRes.GetData()["uid"].(string))
	return readManagementOpenStackDataCenterServer(d, m)
}

func readManagementOpenStackDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showOpenStackDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showOpenStackDataCenterServerRes.Success {
		if objectNotFound(showOpenStackDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showOpenStackDataCenterServerRes.ErrorMsg)
	}
	openstackDataCenterServer := showOpenStackDataCenterServerRes.GetData()

	if v := openstackDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if openstackDataCenterServer["properties"] != nil {
		propsJson, ok := openstackDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				if propName == "unsafe_auto_accept" {
					propValue, _ = strconv.ParseBool(propValue.(string))
				}
				_ = d.Set(propName, propValue)
			}
		}
	}

	if openstackDataCenterServer["tags"] != nil {
		tagsJson, ok := openstackDataCenterServer["tags"].([]interface{})
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

	if v := openstackDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := openstackDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := openstackDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := openstackDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementOpenStackDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	openstackDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		openstackDataCenterServer["name"] = oldName
		openstackDataCenterServer["new-name"] = newName
	} else {
		openstackDataCenterServer["name"] = d.Get("name")
	}

	if d.HasChange("hostname") {
		openstackDataCenterServer["hostname"] = d.Get("hostname")
	}

	if d.HasChange("password") {
		openstackDataCenterServer["password"] = d.Get("password")
	}

	if d.HasChange("password_base64") {
		openstackDataCenterServer["password-base64"] = d.Get("password_base64")
	}

	if d.HasChange("username") {
		openstackDataCenterServer["username"] = d.Get("username")
		if v := d.Get("password"); v != nil && v != "" {
			openstackDataCenterServer["password"] = v
		}
		if v := d.Get("password_base64"); v != nil && v != "" {
			openstackDataCenterServer["password-base64"] = v
		}
	}

	if d.HasChange("certificate_fingerprint") {
		openstackDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}

	if d.HasChange("unsafe_auto_accept") {
		openstackDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			openstackDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			openstackDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		openstackDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		openstackDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		openstackDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		openstackDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update openstackDataCenterServer - Map = ", openstackDataCenterServer)

	updateOpenStackDataCenterServerRes, err := client.ApiCall("set-data-center-server", openstackDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateOpenStackDataCenterServerRes.Success {
		if updateOpenStackDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateOpenStackDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateOpenStackDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementOpenStackDataCenterServer(d, m)
}

func deleteManagementOpenStackDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	openstackDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		openstackDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		openstackDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete openstackDataCenterServer")

	deleteOpenStackDataCenterServerRes, err := client.ApiCall("delete-data-center-server", openstackDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteOpenStackDataCenterServerRes.Success {
		if deleteOpenStackDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteOpenStackDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
