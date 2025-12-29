package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceManagementProxmoxDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementProxmoxDataCenterServer,
		Read:   readManagementProxmoxDataCenterServer,
		Update: updateManagementProxmoxDataCenterServer,
		Delete: deleteManagementProxmoxDataCenterServer,
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
				Description: "IP Address or hostname of the Proxmox server.",
			},
			"token_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "API Token Id, in format Username@Realm!TokenName",
			},
			"secret": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Secret token API.",
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
			"automatic_refresh": {
				Computed:    true,
				Type:        schema.TypeBool,
				Description: "Indicates whether the data center server's content is automatically updated.",
			},
			"data_center_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Data Center type.",
			},
		},
	}
}

func createManagementProxmoxDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	proxmoxDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		proxmoxDataCenterServer["name"] = v.(string)
	}

	proxmoxDataCenterServer["type"] = "proxmox"

	if v, ok := d.GetOk("hostname"); ok {
		proxmoxDataCenterServer["hostname"] = v.(string)
	}

	if v, ok := d.GetOk("token_id"); ok {
		proxmoxDataCenterServer["token-id"] = v.(string)
	}

	if v, ok := d.GetOk("secret"); ok {
		proxmoxDataCenterServer["secret"] = v.(string)
	}

	if v, ok := d.GetOk("certificate_fingerprint"); ok {
		proxmoxDataCenterServer["certificate-fingerprint"] = v.(string)
	}

	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		proxmoxDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		proxmoxDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		proxmoxDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		proxmoxDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		proxmoxDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		proxmoxDataCenterServer["ignore-errors"] = v.(bool)
	}

	if v, ok := d.GetOkExists("automatic_refresh"); ok {
		proxmoxDataCenterServer["automatic-refresh"] = v.(bool)
	}

	log.Println("Create proxmoxDataCenterServer - Map = ", proxmoxDataCenterServer)

	addProxmoxDataCenterServerRes, err := client.ApiCall("add-data-center-server", proxmoxDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addProxmoxDataCenterServerRes.Success {
		if addProxmoxDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addProxmoxDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addProxmoxDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": proxmoxDataCenterServer["name"],
	}
	showProxmoxDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showProxmoxDataCenterServerRes.Success {
		return fmt.Errorf(showProxmoxDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showProxmoxDataCenterServerRes.GetData()["uid"].(string))
	return readManagementProxmoxDataCenterServer(d, m)
}

func readManagementProxmoxDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showProxmoxDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showProxmoxDataCenterServerRes.Success {
		if objectNotFound(showProxmoxDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showProxmoxDataCenterServerRes.ErrorMsg)
	}
	proxmoxDataCenterServer := showProxmoxDataCenterServerRes.GetData()

	log.Println("Read Proxmox Data Center - Show JSON = ", proxmoxDataCenterServer)

	if v := proxmoxDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := proxmoxDataCenterServer["data-center-type"]; v != nil {
		_ = d.Set("data_center_type", v)
	}

	// Hostname, username, realm, token-id, and secret are part of the "properties".
	if proxmoxDataCenterServer["properties"] != nil {
		propsJson, ok := proxmoxDataCenterServer["properties"].([]interface{})
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

	if proxmoxDataCenterServer["tags"] != nil {
		tagsJson, ok := proxmoxDataCenterServer["tags"].([]interface{})
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
	if v := proxmoxDataCenterServer["automatic-refresh"]; v != nil {
		_ = d.Set("automatic_refresh", v)
	}
	if v := proxmoxDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := proxmoxDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := proxmoxDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := proxmoxDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementProxmoxDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	proxmoxDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		proxmoxDataCenterServer["name"] = oldName
		proxmoxDataCenterServer["new-name"] = newName
	} else {
		proxmoxDataCenterServer["name"] = d.Get("name")
	}

	if d.HasChange("hostname") {
		proxmoxDataCenterServer["hostname"] = d.Get("hostname")
	}

	if d.HasChange("secret") {
		proxmoxDataCenterServer["secret"] = d.Get("secret")
	}

	if d.HasChange("token_id") {
		proxmoxDataCenterServer["token-id"] = d.Get("token_id")
		if v := d.Get("secret"); v != nil && v != "" {
			proxmoxDataCenterServer["secret"] = v
		}
	}

	if d.HasChange("certificate_fingerprint") {
		proxmoxDataCenterServer["certificate-fingerprint"] = d.Get("certificate_fingerprint")
	}

	if d.HasChange("unsafe_auto_accept") {
		proxmoxDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			proxmoxDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			proxmoxDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		proxmoxDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		proxmoxDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		proxmoxDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		proxmoxDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update proxmoxDataCenterServer - Map = ", proxmoxDataCenterServer)

	updateProxmoxDataCenterServerRes, err := client.ApiCall("set-data-center-server", proxmoxDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateProxmoxDataCenterServerRes.Success {
		if updateProxmoxDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateProxmoxDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateProxmoxDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementProxmoxDataCenterServer(d, m)
}

func deleteManagementProxmoxDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	proxmoxDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		proxmoxDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		proxmoxDataCenterServerPayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete proxmoxDataCenterServer")

	deleteProxmoxDataCenterServerRes, err := client.ApiCall("delete-data-center-server", proxmoxDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteProxmoxDataCenterServerRes.Success {
		if deleteProxmoxDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteProxmoxDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
