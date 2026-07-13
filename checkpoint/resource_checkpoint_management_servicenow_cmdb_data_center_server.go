package checkpoint

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementServicenowCmdbDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementServicenowCmdbDataCenterServer,
		Read:   readManagementServicenowCmdbDataCenterServer,
		Update: updateManagementServicenowCmdbDataCenterServer,
		Delete: deleteManagementServicenowCmdbDataCenterServer,
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
				Description: "Instance hostname of the ServiceNow instance (e.g. instance.service-now.com).",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "User name for ServiceNow instance.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Password for ServiceNow instance.",
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

func createManagementServicenowCmdbDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	servicenowCmdbDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		servicenowCmdbDataCenterServer["name"] = v.(string)
	}

	servicenowCmdbDataCenterServer["type"] = "servicenow-cmdb"

	if v, ok := d.GetOk("hostname"); ok {
		servicenowCmdbDataCenterServer["hostname"] = v.(string)
	}
	if v, ok := d.GetOk("username"); ok {
		servicenowCmdbDataCenterServer["username"] = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		servicenowCmdbDataCenterServer["password"] = v.(string)
	}
	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		servicenowCmdbDataCenterServer["certificate-fingerprint"] = v.(string)
	}
	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		servicenowCmdbDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		servicenowCmdbDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		servicenowCmdbDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		servicenowCmdbDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		servicenowCmdbDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		servicenowCmdbDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create servicenowCmdbDataCenterServer - Map = ", servicenowCmdbDataCenterServer)

	addRes, err := client.ApiCall("add-data-center-server", servicenowCmdbDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
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
		"name": servicenowCmdbDataCenterServer["name"],
	}
	showRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showRes.Success {
		return fmt.Errorf("%s", showRes.ErrorMsg)
	}
	d.SetId(showRes.GetData()["uid"].(string))
	return readManagementServicenowCmdbDataCenterServer(d, m)
}

func readManagementServicenowCmdbDataCenterServer(d *schema.ResourceData, m interface{}) error {
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
	servicenowCmdbDataCenterServer := showRes.GetData()

	if v := servicenowCmdbDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if servicenowCmdbDataCenterServer["properties"] != nil {
		propsJson, ok := servicenowCmdbDataCenterServer["properties"].([]interface{})
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

	if servicenowCmdbDataCenterServer["tags"] != nil {
		tagsJson, ok := servicenowCmdbDataCenterServer["tags"].([]interface{})
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

	if v := servicenowCmdbDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := servicenowCmdbDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := servicenowCmdbDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := servicenowCmdbDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementServicenowCmdbDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	servicenowCmdbDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		servicenowCmdbDataCenterServer["name"] = oldName
		servicenowCmdbDataCenterServer["new-name"] = newName
	} else {
		servicenowCmdbDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("hostname"); ok {
		servicenowCmdbDataCenterServer["hostname"] = d.Get("hostname")
	}
	if ok := d.HasChange("username"); ok {
		servicenowCmdbDataCenterServer["username"] = d.Get("username")
	}
	if ok := d.HasChange("password"); ok {
		servicenowCmdbDataCenterServer["password"] = d.Get("password")
	}
	if ok := d.HasChange("certificate_fingerprint"); ok {
		servicenowCmdbDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}
	if ok := d.HasChange("unsafe_auto_accept"); ok {
		servicenowCmdbDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			servicenowCmdbDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			servicenowCmdbDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		servicenowCmdbDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		servicenowCmdbDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		servicenowCmdbDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		servicenowCmdbDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update servicenowCmdbDataCenterServer - Map = ", servicenowCmdbDataCenterServer)

	updateRes, err := client.ApiCall("set-data-center-server", servicenowCmdbDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
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

	return readManagementServicenowCmdbDataCenterServer(d, m)
}

func deleteManagementServicenowCmdbDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	servicenowCmdbDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		servicenowCmdbDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		servicenowCmdbDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete servicenowCmdbDataCenterServer")

	deleteRes, err := client.ApiCall("delete-data-center-server", servicenowCmdbDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteRes.Success {
		if deleteRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
