package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceManagementAciDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAciDataCenterServer,
		Read:   readManagementAciDataCenterServer,
		Update: updateManagementAciDataCenterServer,
		Delete: deleteManagementAciDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"urls": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Address of APIC cluster members.\nExample: http(s)://<host1 ip/url>.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "User ID of the Cisco APIC server.\nWhen using commonLoginLogic Domains use the following syntax:\napic:<domain>\\<username>.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the Cisco APIC server.",
			},
			"password_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the Cisco APIC server encoded in Base64.",
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

func createManagementAciDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	aciDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		aciDataCenterServer["name"] = v.(string)
	}

	aciDataCenterServer["type"] = "aci"

	if v, ok := d.GetOk("urls"); ok {
		aciDataCenterServer["urls"] = v
	}

	if v, ok := d.GetOk("username"); ok {
		aciDataCenterServer["username"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		aciDataCenterServer["password"] = v.(string)
	}

	if v, ok := d.GetOk("password_base64"); ok {
		aciDataCenterServer["password-base64"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		aciDataCenterServer["certificate-fingerprint"] = v.(string)
	}

	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		aciDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		aciDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		aciDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		aciDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		aciDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		aciDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create aciDataCenterServer - Map = ", aciDataCenterServer)

	addAciDataCenterServerRes, err := client.ApiCall("add-data-center-server", aciDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addAciDataCenterServerRes.Success {
		if addAciDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addAciDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addAciDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": aciDataCenterServer["name"],
	}
	showAciDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAciDataCenterServerRes.Success {
		return fmt.Errorf(showAciDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showAciDataCenterServerRes.GetData()["uid"].(string))
	return readManagementAciDataCenterServer(d, m)
}

func readManagementAciDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showAciDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAciDataCenterServerRes.Success {
		if objectNotFound(showAciDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showAciDataCenterServerRes.ErrorMsg)
	}
	aciDataCenterServer := showAciDataCenterServerRes.GetData()

	if v := aciDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if aciDataCenterServer["properties"] != nil {
		propsJson, ok := aciDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				if propName == "unsafe_auto_accept" {
					propValue, _ = strconv.ParseBool(propValue.(string))
				}
				if propName == "urls" {
					propValue = strings.Split(propValue.(string), ";")
				}
				_ = d.Set(propName, propValue)
			}
		}
	}

	if aciDataCenterServer["tags"] != nil {
		tagsJson, ok := aciDataCenterServer["tags"].([]interface{})
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

	if v := aciDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := aciDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := aciDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := aciDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementAciDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	aciDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		aciDataCenterServer["name"] = oldName
		aciDataCenterServer["new-name"] = newName
	} else {
		aciDataCenterServer["name"] = d.Get("name")
	}

	if d.HasChange("urls") {
		aciDataCenterServer["urls"] = d.Get("urls")
	}

	if d.HasChange("password") {
		aciDataCenterServer["password"] = d.Get("password")
	}

	if d.HasChange("password_base64") {
		aciDataCenterServer["password-base64"] = d.Get("password_base64")
	}

	if d.HasChange("username") {
		aciDataCenterServer["username"] = d.Get("username")
		if v := d.Get("password"); v != nil && v != "" {
			aciDataCenterServer["password"] = v
		}
		if v := d.Get("password_base64"); v != nil && v != "" {
			aciDataCenterServer["password-base64"] = v
		}
	}

	if d.HasChange("certificate_fingerprint") {
		aciDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}

	if d.HasChange("unsafe_auto_accept") {
		aciDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			aciDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			aciDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		aciDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		aciDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		aciDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		aciDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update aciDataCenterServer - Map = ", aciDataCenterServer)

	updateAciDataCenterServerRes, err := client.ApiCall("set-data-center-server", aciDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateAciDataCenterServerRes.Success {
		if updateAciDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateAciDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateAciDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementAciDataCenterServer(d, m)
}

func deleteManagementAciDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	aciDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete aciDataCenterServer")

	deleteAciDataCenterServerRes, err := client.ApiCall("delete-data-center-server", aciDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteAciDataCenterServerRes.Success {
		if deleteAciDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteAciDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
