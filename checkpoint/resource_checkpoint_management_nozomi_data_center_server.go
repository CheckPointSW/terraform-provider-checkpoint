package checkpoint

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementNozomiDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementNozomiDataCenterServer,
		Read:   readManagementNozomiDataCenterServer,
		Update: updateManagementNozomiDataCenterServer,
		Delete: deleteManagementNozomiDataCenterServer,
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
				Description: "IP address or hostname of the Nozomi Guardian or CMC server.",
			},
			"key_name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "API key name for Nozomi.",
			},
			"key_token": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "API key token for Nozomi.",
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

func createManagementNozomiDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	nozomiDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		nozomiDataCenterServer["name"] = v.(string)
	}

	nozomiDataCenterServer["type"] = "nozomi"

	if v, ok := d.GetOk("hostname"); ok {
		nozomiDataCenterServer["hostname"] = v.(string)
	}
	if v, ok := d.GetOk("key_name"); ok {
		nozomiDataCenterServer["key-name"] = v.(string)
	}
	if v, ok := d.GetOk("key_token"); ok {
		nozomiDataCenterServer["key-token"] = v.(string)
	}
	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		nozomiDataCenterServer["certificate-fingerprint"] = v.(string)
	}
	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		nozomiDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		nozomiDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		nozomiDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		nozomiDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		nozomiDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		nozomiDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create nozomiDataCenterServer - Map = ", nozomiDataCenterServer)

	addRes, err := client.ApiCall("add-data-center-server", nozomiDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
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
		"name": nozomiDataCenterServer["name"],
	}
	showRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showRes.Success {
		return fmt.Errorf("%s", showRes.ErrorMsg)
	}
	d.SetId(showRes.GetData()["uid"].(string))
	return readManagementNozomiDataCenterServer(d, m)
}

func readManagementNozomiDataCenterServer(d *schema.ResourceData, m interface{}) error {
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
	nozomiDataCenterServer := showRes.GetData()

	if v := nozomiDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if nozomiDataCenterServer["properties"] != nil {
		propsJson, ok := nozomiDataCenterServer["properties"].([]interface{})
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

	if nozomiDataCenterServer["tags"] != nil {
		tagsJson, ok := nozomiDataCenterServer["tags"].([]interface{})
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

	if v := nozomiDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := nozomiDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := nozomiDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := nozomiDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementNozomiDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	nozomiDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		nozomiDataCenterServer["name"] = oldName
		nozomiDataCenterServer["new-name"] = newName
	} else {
		nozomiDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("hostname"); ok {
		nozomiDataCenterServer["hostname"] = d.Get("hostname")
	}
	if ok := d.HasChange("key_name"); ok {
		nozomiDataCenterServer["key-name"] = d.Get("key_name")
	}
	if ok := d.HasChange("key_token"); ok {
		nozomiDataCenterServer["key-token"] = d.Get("key_token")
	}
	if ok := d.HasChange("certificate_fingerprint"); ok {
		nozomiDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}
	if ok := d.HasChange("unsafe_auto_accept"); ok {
		nozomiDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			nozomiDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			nozomiDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		nozomiDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		nozomiDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		nozomiDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		nozomiDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update nozomiDataCenterServer - Map = ", nozomiDataCenterServer)

	updateRes, err := client.ApiCall("set-data-center-server", nozomiDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
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

	return readManagementNozomiDataCenterServer(d, m)
}

func deleteManagementNozomiDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	nozomiDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		nozomiDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		nozomiDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete nozomiDataCenterServer")

	deleteRes, err := client.ApiCall("delete-data-center-server", nozomiDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteRes.Success {
		if deleteRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
