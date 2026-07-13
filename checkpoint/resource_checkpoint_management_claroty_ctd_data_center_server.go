package checkpoint

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementClarotyCtdDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementClarotyCtdDataCenterServer,
		Read:   readManagementClarotyCtdDataCenterServer,
		Update: updateManagementClarotyCtdDataCenterServer,
		Delete: deleteManagementClarotyCtdDataCenterServer,
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
				Description: "IP address or hostname of the Claroty CTD server.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User name for Claroty CTD.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Password for Claroty CTD.",
			},
			"certificate_fingerprint": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specify the SHA-1 or SHA-256 fingerprint of the Data Center Server's certificate.",
			},
			"unsafe_auto_accept": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "When set to false, the current Data Center Server's certificate should be trusted, either by providing the certificate-fingerprint argument or by relying on a previously trusted certificate of this hostname. When set to true, trust the current Data Center Server's certificate as-is.",
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

func createManagementClarotyCtdDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	clarotyCtdDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		clarotyCtdDataCenterServer["name"] = v.(string)
	}

	clarotyCtdDataCenterServer["type"] = "claroty-ctd"

	if v, ok := d.GetOk("hostname"); ok {
		clarotyCtdDataCenterServer["hostname"] = v.(string)
	}
	if v, ok := d.GetOk("username"); ok {
		clarotyCtdDataCenterServer["username"] = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		clarotyCtdDataCenterServer["password"] = v.(string)
	}
	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		clarotyCtdDataCenterServer["certificate-fingerprint"] = v.(string)
	}
	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		clarotyCtdDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		clarotyCtdDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		clarotyCtdDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		clarotyCtdDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		clarotyCtdDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		clarotyCtdDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create clarotyCtdDataCenterServer - Map = ", clarotyCtdDataCenterServer)

	addRes, err := client.ApiCall("add-data-center-server", clarotyCtdDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !addRes.Success {
		if addRes.ErrorMsg != "" {
			return fmt.Errorf("%s", addRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addRes.GetData())
		return fmt.Errorf("%s", msg)
	}
	payload := map[string]interface{}{
		"name": clarotyCtdDataCenterServer["name"],
	}
	showRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showRes.Success {
		return fmt.Errorf("%s", showRes.ErrorMsg)
	}
	d.SetId(showRes.GetData()["uid"].(string))
	return readManagementClarotyCtdDataCenterServer(d, m)
}

func readManagementClarotyCtdDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showRes.Success {
		if objectNotFound(showRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showRes.ErrorMsg)
	}
	clarotyCtdDataCenterServer := showRes.GetData()

	if v := clarotyCtdDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if clarotyCtdDataCenterServer["properties"] != nil {
		propsJson, ok := clarotyCtdDataCenterServer["properties"].([]interface{})
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

	if clarotyCtdDataCenterServer["tags"] != nil {
		tagsJson, ok := clarotyCtdDataCenterServer["tags"].([]interface{})
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

	if v := clarotyCtdDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := clarotyCtdDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := clarotyCtdDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := clarotyCtdDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementClarotyCtdDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	clarotyCtdDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		clarotyCtdDataCenterServer["name"] = oldName
		clarotyCtdDataCenterServer["new-name"] = newName
	} else {
		clarotyCtdDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("hostname"); ok {
		clarotyCtdDataCenterServer["hostname"] = d.Get("hostname")
	}
	if ok := d.HasChange("username"); ok {
		clarotyCtdDataCenterServer["username"] = d.Get("username")
	}
	if ok := d.HasChange("password"); ok {
		clarotyCtdDataCenterServer["password"] = d.Get("password")
	}
	if ok := d.HasChange("certificate_fingerprint"); ok {
		clarotyCtdDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}
	if ok := d.HasChange("unsafe_auto_accept"); ok {
		clarotyCtdDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			clarotyCtdDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			clarotyCtdDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		clarotyCtdDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		clarotyCtdDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		clarotyCtdDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		clarotyCtdDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update clarotyCtdDataCenterServer - Map = ", clarotyCtdDataCenterServer)

	updateRes, err := client.ApiCall("set-data-center-server", clarotyCtdDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !updateRes.Success {
		if updateRes.ErrorMsg != "" {
			return fmt.Errorf("%s", updateRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateRes.GetData())
		return fmt.Errorf("%s", msg)
	}

	return readManagementClarotyCtdDataCenterServer(d, m)
}

func deleteManagementClarotyCtdDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	clarotyCtdDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		clarotyCtdDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		clarotyCtdDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete clarotyCtdDataCenterServer")

	deleteRes, err := client.ApiCall("delete-data-center-server", clarotyCtdDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteRes.Success {
		if deleteRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
