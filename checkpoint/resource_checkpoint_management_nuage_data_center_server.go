package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceManagementNuageDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementNuageDataCenterServer,
		Read:   readManagementNuageDataCenterServer,
		Update: updateManagementNuageDataCenterServer,
		Delete: deleteManagementNuageDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		DeprecationMessage: "This resource will be deprecated In R82.10",
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"hostname": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "IP address or hostname of the Nuage server.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Username of the Nuage administrator.",
			},
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Organization name or enterprise.",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the Nuage administrator.",
			},
			"password_base64": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Password of the Nuage administrator encoded in Base64.",
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

func createManagementNuageDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	nuageDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		nuageDataCenterServer["name"] = v.(string)
	}

	nuageDataCenterServer["type"] = "nuage"

	if v, ok := d.GetOk("hostname"); ok {
		nuageDataCenterServer["hostname"] = v.(string)
	}

	if v, ok := d.GetOk("username"); ok {
		nuageDataCenterServer["username"] = v.(string)
	}

	if v, ok := d.GetOk("organization"); ok {
		nuageDataCenterServer["organization"] = v.(string)
	}

	if v, ok := d.GetOk("password"); ok {
		nuageDataCenterServer["password"] = v.(string)
	}

	if v, ok := d.GetOk("password_base64"); ok {
		nuageDataCenterServer["password-base64"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		nuageDataCenterServer["certificate-fingerprint"] = v.(string)
	}

	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		nuageDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		nuageDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		nuageDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		nuageDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		nuageDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		nuageDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create nuageDataCenterServer - Map = ", nuageDataCenterServer)

	addNuageDataCenterServerRes, err := client.ApiCall("add-data-center-server", nuageDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addNuageDataCenterServerRes.Success {
		if addNuageDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addNuageDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addNuageDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": nuageDataCenterServer["name"],
	}
	showNuageDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNuageDataCenterServerRes.Success {
		return fmt.Errorf(showNuageDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showNuageDataCenterServerRes.GetData()["uid"].(string))
	return readManagementNuageDataCenterServer(d, m)
}

func readManagementNuageDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showNuageDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showNuageDataCenterServerRes.Success {
		if objectNotFound(showNuageDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showNuageDataCenterServerRes.ErrorMsg)
	}
	nuageDataCenterServer := showNuageDataCenterServerRes.GetData()

	if v := nuageDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if nuageDataCenterServer["properties"] != nil {
		propsJson, ok := nuageDataCenterServer["properties"].([]interface{})
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

	if nuageDataCenterServer["tags"] != nil {
		tagsJson, ok := nuageDataCenterServer["tags"].([]interface{})
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

	if v := nuageDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := nuageDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := nuageDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := nuageDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementNuageDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	nuageDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		nuageDataCenterServer["name"] = oldName
		nuageDataCenterServer["new-name"] = newName
	} else {
		nuageDataCenterServer["name"] = d.Get("name")
	}

	if d.HasChange("organization") {
		nuageDataCenterServer["organization"] = d.Get("organization")
	}

	if d.HasChange("hostname") {
		nuageDataCenterServer["hostname"] = d.Get("hostname")
	}

	if d.HasChange("password") {
		nuageDataCenterServer["password"] = d.Get("password")
	}

	if d.HasChange("password_base64") {
		nuageDataCenterServer["password-base64"] = d.Get("password_base64")
	}

	if d.HasChange("username") {
		nuageDataCenterServer["username"] = d.Get("username")
		if v := d.Get("password"); v != nil && v != "" {
			nuageDataCenterServer["password"] = v
		}
		if v := d.Get("password_base64"); v != nil && v != "" {
			nuageDataCenterServer["password-base64"] = v
		}
	}

	if d.HasChange("certificate_fingerprint") {
		nuageDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}

	if d.HasChange("unsafe_auto_accept") {
		nuageDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			nuageDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			nuageDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		nuageDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		nuageDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		nuageDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		nuageDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update nuageDataCenterServer - Map = ", nuageDataCenterServer)

	updateNuageDataCenterServerRes, err := client.ApiCall("set-data-center-server", nuageDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateNuageDataCenterServerRes.Success {
		if updateNuageDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateNuageDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateNuageDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementNuageDataCenterServer(d, m)
}

func deleteManagementNuageDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	nuageDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		nuageDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		nuageDataCenterServerPayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete nuageDataCenterServer")

	deleteNuageDataCenterServerRes, err := client.ApiCall("delete-data-center-server", nuageDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())

	if err != nil || !deleteNuageDataCenterServerRes.Success {
		if deleteNuageDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteNuageDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
