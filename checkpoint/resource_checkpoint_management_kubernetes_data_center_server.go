package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
	"strings"
)

func resourceManagementKubernetesDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementKubernetesDataCenterServer,
		Read:   readManagementKubernetesDataCenterServer,
		Update: updateManagementKubernetesDataCenterServer,
		Delete: deleteManagementKubernetesDataCenterServer,
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
				Sensitive:   true,
				Description: "IP address or hostname of the Kubernetes server.",
			},
			"token_file": {
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				Description: "Kubernetes access token encoded in base64.",
			},
			"ca_certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Kubernetes public certificate key encoded in base64.",
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

func createManagementKubernetesDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	kubernetesDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		kubernetesDataCenterServer["name"] = v.(string)
	}

	kubernetesDataCenterServer["type"] = "kubernetes"

	if v, ok := d.GetOk("hostname"); ok {
		kubernetesDataCenterServer["hostname"] = v.(string)
	}

	if v, ok := d.GetOk("token_file"); ok {
		kubernetesDataCenterServer["token-file"] = v.(string)
	}

	if v, ok := d.GetOk("ca_certificate"); ok {
		kubernetesDataCenterServer["ca-certificate"] = v.(string)
	}

	if v, ok := d.GetOk("unsafe_auto_accept"); ok {
		kubernetesDataCenterServer["unsafe-auto-accept"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		kubernetesDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		kubernetesDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		kubernetesDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		kubernetesDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		kubernetesDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create kubernetesDataCenterServer - Map = ", kubernetesDataCenterServer)

	addKubernetesDataCenterServerRes, err := client.ApiCall("add-data-center-server", kubernetesDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addKubernetesDataCenterServerRes.Success {
		if addKubernetesDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addKubernetesDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addKubernetesDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": kubernetesDataCenterServer["name"],
	}
	showKubernetesDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showKubernetesDataCenterServerRes.Success {
		return fmt.Errorf(showKubernetesDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showKubernetesDataCenterServerRes.GetData()["uid"].(string))
	return readManagementKubernetesDataCenterServer(d, m)
}

func readManagementKubernetesDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showKubernetesDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showKubernetesDataCenterServerRes.Success {
		if objectNotFound(showKubernetesDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showKubernetesDataCenterServerRes.ErrorMsg)
	}
	kubernetesDataCenterServer := showKubernetesDataCenterServerRes.GetData()

	if v := kubernetesDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if kubernetesDataCenterServer["properties"] != nil {
		propsJson, ok := kubernetesDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				if propMap["name"] != nil {
					propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
					propValue := propMap["value"]
					if propName == "unsafe_auto_accept" {
						propValue, _ = strconv.ParseBool(propValue.(string))
					}
					_ = d.Set(propName, propValue)
				}
			}
		}
	}

	if kubernetesDataCenterServer["tags"] != nil {
		tagsJson, ok := kubernetesDataCenterServer["tags"].([]interface{})
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

	if v := kubernetesDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := kubernetesDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := kubernetesDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := kubernetesDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementKubernetesDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	kubernetesDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		kubernetesDataCenterServer["name"] = oldName
		kubernetesDataCenterServer["new-name"] = newName
	} else {
		kubernetesDataCenterServer["name"] = d.Get("name")
	}

	if d.HasChange("hostname") {
		kubernetesDataCenterServer["hostname"] = d.Get("hostname")
	}

	if d.HasChange("token_file") {
		kubernetesDataCenterServer["token-file"] = d.Get("token_file")
	}

	if d.HasChange("ca_certificate") {
		kubernetesDataCenterServer["ca-certificate"] = d.Get("ca_certificate")
	}

	if d.HasChange("unsafe_auto_accept") {
		kubernetesDataCenterServer["unsafe-auto-accept"] = d.Get("unsafe_auto_accept")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			kubernetesDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			kubernetesDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		kubernetesDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		kubernetesDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		kubernetesDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		kubernetesDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update kubernetesDataCenterServer - Map = ", kubernetesDataCenterServer)

	updateKubernetesDataCenterServerRes, err := client.ApiCall("set-data-center-server", kubernetesDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateKubernetesDataCenterServerRes.Success {
		if updateKubernetesDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateKubernetesDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateKubernetesDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementKubernetesDataCenterServer(d, m)
}

func deleteManagementKubernetesDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	kubernetesDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		kubernetesDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		kubernetesDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete kubernetesDataCenterServer")

	deleteKubernetesDataCenterServerRes, err := client.ApiCall("delete-data-center-server", kubernetesDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteKubernetesDataCenterServerRes.Success {
		if deleteKubernetesDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteKubernetesDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
