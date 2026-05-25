package checkpoint

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementIllumioDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementIllumioDataCenterServer,
		Read:   readManagementIllumioDataCenterServer,
		Update: updateManagementIllumioDataCenterServer,
		Delete: deleteManagementIllumioDataCenterServer,
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
				Description: "IP address or hostname of the Illumio PCE server.",
			},
			"org_id": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Organization ID in the Illumio PCE.",
			},
			"auth_username": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Authentication username.",
			},
			"secret": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Secret for authentication.",
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

func createManagementIllumioDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	illumioDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		illumioDataCenterServer["name"] = v.(string)
	}

	illumioDataCenterServer["type"] = "illumio"

	if v, ok := d.GetOk("hostname"); ok {
		illumioDataCenterServer["hostname"] = v.(string)
	}

	if v, ok := d.GetOk("org_id"); ok {
		illumioDataCenterServer["org-id"] = v.(int)
	}

	if v, ok := d.GetOk("auth_username"); ok {
		illumioDataCenterServer["auth-username"] = v.(string)
	}

	if v, ok := d.GetOk("secret"); ok {
		illumioDataCenterServer["secret"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		illumioDataCenterServer["certificate-fingerprint"] = v.(string)
	}

	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		illumioDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		illumioDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		illumioDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		illumioDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		illumioDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		illumioDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create illumioDataCenterServer - Map = ", illumioDataCenterServer)

	addIllumioDataCenterServerRes, err := client.ApiCall("add-data-center-server", illumioDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !addIllumioDataCenterServerRes.Success {
		if addIllumioDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf("%s", addIllumioDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addIllumioDataCenterServerRes.GetData())
		return fmt.Errorf("%s", msg)
	}
	payload := map[string]interface{}{
		"name": illumioDataCenterServer["name"],
	}
	showIllumioDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showIllumioDataCenterServerRes.Success {
		return fmt.Errorf("%s", showIllumioDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showIllumioDataCenterServerRes.GetData()["uid"].(string))
	return readManagementIllumioDataCenterServer(d, m)
}

func readManagementIllumioDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showIllumioDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showIllumioDataCenterServerRes.Success {
		if objectNotFound(showIllumioDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showIllumioDataCenterServerRes.ErrorMsg)
	}
	illumioDataCenterServer := showIllumioDataCenterServerRes.GetData()

	if v := illumioDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if illumioDataCenterServer["properties"] != nil {
		propsJson, ok := illumioDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				if propName == "unsafe_auto_accept" {
					propValue, _ = strconv.ParseBool(propValue.(string))
				} else if propName == "org_id" {
					propValue, _ = strconv.Atoi(propValue.(string))
				}
				_ = d.Set(propName, propValue)
			}
		}
	}

	if illumioDataCenterServer["tags"] != nil {
		tagsJson, ok := illumioDataCenterServer["tags"].([]interface{})
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

	if v := illumioDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := illumioDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := illumioDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := illumioDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementIllumioDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	illumioDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		illumioDataCenterServer["name"] = oldName
		illumioDataCenterServer["new-name"] = newName
	} else {
		illumioDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("hostname"); ok {
		illumioDataCenterServer["hostname"] = d.Get("hostname")
	}

	if ok := d.HasChange("org_id"); ok {
		illumioDataCenterServer["org-id"] = d.Get("org_id")
	}

	if ok := d.HasChange("auth_username"); ok {
		illumioDataCenterServer["auth-username"] = d.Get("auth_username")
	}

	if ok := d.HasChange("secret"); ok {
		illumioDataCenterServer["secret"] = d.Get("secret")
	}

	if ok := d.HasChange("certificate_fingerprint"); ok {
		illumioDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}

	if ok := d.HasChange("unsafe_auto_accept"); ok {
		illumioDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			illumioDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			illumioDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		illumioDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		illumioDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		illumioDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		illumioDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update illumioDataCenterServer - Map = ", illumioDataCenterServer)

	updateIllumioDataCenterServerRes, err := client.ApiCall("set-data-center-server", illumioDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !updateIllumioDataCenterServerRes.Success {
		if updateIllumioDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf("%s", updateIllumioDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateIllumioDataCenterServerRes.GetData())
		return fmt.Errorf("%s", msg)
	}

	return readManagementIllumioDataCenterServer(d, m)
}

func deleteManagementIllumioDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	illumioDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		illumioDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		illumioDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete illumioDataCenterServer")

	deleteIllumioDataCenterServerRes, err := client.ApiCall("delete-data-center-server", illumioDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteIllumioDataCenterServerRes.Success {
		if deleteIllumioDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteIllumioDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
