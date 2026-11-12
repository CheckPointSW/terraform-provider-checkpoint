package checkpoint

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementGuardicoreDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGuardicoreDataCenterServer,
		Read:   readManagementGuardicoreDataCenterServer,
		Update: updateManagementGuardicoreDataCenterServer,
		Delete: deleteManagementGuardicoreDataCenterServer,
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
				Description: "IP Address or hostname of the Guardicore Centra management server.",
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username for Guardicore Centra.",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Password for Guardicore Centra.",
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

func createManagementGuardicoreDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	guardicoreDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		guardicoreDataCenterServer["name"] = v.(string)
	}

	guardicoreDataCenterServer["type"] = "guardicore"

	if v, ok := d.GetOk("hostname"); ok {
		guardicoreDataCenterServer["hostname"] = v.(string)
	}
	if v, ok := d.GetOk("username"); ok {
		guardicoreDataCenterServer["username"] = v.(string)
	}
	if v, ok := d.GetOk("password"); ok {
		guardicoreDataCenterServer["password"] = v.(string)
	}
	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		guardicoreDataCenterServer["certificate-fingerprint"] = v.(string)
	}
	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		guardicoreDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		guardicoreDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		guardicoreDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		guardicoreDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		guardicoreDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		guardicoreDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create guardicoreDataCenterServer - Map = ", guardicoreDataCenterServer)

	addRes, err := client.ApiCall("add-data-center-server", guardicoreDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
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
		"name": guardicoreDataCenterServer["name"],
	}
	showRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showRes.Success {
		return fmt.Errorf("%s", showRes.ErrorMsg)
	}
	d.SetId(showRes.GetData()["uid"].(string))
	return readManagementGuardicoreDataCenterServer(d, m)
}

func readManagementGuardicoreDataCenterServer(d *schema.ResourceData, m interface{}) error {
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
	guardicoreDataCenterServer := showRes.GetData()

	if v := guardicoreDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if guardicoreDataCenterServer["properties"] != nil {
		propsJson, ok := guardicoreDataCenterServer["properties"].([]interface{})
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

	if guardicoreDataCenterServer["tags"] != nil {
		tagsJson, ok := guardicoreDataCenterServer["tags"].([]interface{})
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

	if v := guardicoreDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := guardicoreDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := guardicoreDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := guardicoreDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementGuardicoreDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	guardicoreDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		guardicoreDataCenterServer["name"] = oldName
		guardicoreDataCenterServer["new-name"] = newName
	} else {
		guardicoreDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("hostname"); ok {
		guardicoreDataCenterServer["hostname"] = d.Get("hostname")
	}
	if ok := d.HasChange("username"); ok {
		guardicoreDataCenterServer["username"] = d.Get("username")
	}
	if ok := d.HasChange("password"); ok {
		guardicoreDataCenterServer["password"] = d.Get("password")
	}
	if ok := d.HasChange("certificate_fingerprint"); ok {
		guardicoreDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}
	if ok := d.HasChange("unsafe_auto_accept"); ok {
		guardicoreDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			guardicoreDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			guardicoreDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		guardicoreDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		guardicoreDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		guardicoreDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		guardicoreDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update guardicoreDataCenterServer - Map = ", guardicoreDataCenterServer)

	updateRes, err := client.ApiCall("set-data-center-server", guardicoreDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
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

	return readManagementGuardicoreDataCenterServer(d, m)
}

func deleteManagementGuardicoreDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	guardicoreDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		guardicoreDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		guardicoreDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete guardicoreDataCenterServer")

	deleteRes, err := client.ApiCall("delete-data-center-server", guardicoreDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteRes.Success {
		if deleteRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
