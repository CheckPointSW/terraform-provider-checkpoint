package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func resourceManagementGcpDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGcpDataCenterServer,
		Read:   readManagementGcpDataCenterServer,
		Update: updateManagementGcpDataCenterServer,
		Delete: deleteManagementGcpDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"authentication_method": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "key-authentication\nUses the Service Account private key file to authenticate.\nvm-instance-authentication\nUses the Service Account VM Instance to authenticate.\nThis option requires the Security Management Server deployed in a GCP, and runs as a Service Account with the required permissions.",
			},
			"private_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A Service Account Key JSON file, encoded in base64.\nRequired for authentication-method: key-authentication.",
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

func createManagementGcpDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	gcpDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		gcpDataCenterServer["name"] = v.(string)
	}

	gcpDataCenterServer["type"] = "gcp"

	if v, ok := d.GetOk("authentication_method"); ok {
		gcpDataCenterServer["authentication-method"] = v.(string)
	}

	if v, ok := d.GetOk("private_key"); ok {
		gcpDataCenterServer["private-key"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		gcpDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		gcpDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		gcpDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		gcpDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		gcpDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create gcpDataCenterServer - Map = ", gcpDataCenterServer)

	addGcpDataCenterServerRes, err := client.ApiCall("add-data-center-server", gcpDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addGcpDataCenterServerRes.Success {
		if addGcpDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(addGcpDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addGcpDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}
	payload := map[string]interface{}{
		"name": gcpDataCenterServer["name"],
	}
	showGcpDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGcpDataCenterServerRes.Success {
		return fmt.Errorf(showGcpDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showGcpDataCenterServerRes.GetData()["uid"].(string))
	return readManagementGcpDataCenterServer(d, m)
}

func readManagementGcpDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showGcpDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGcpDataCenterServerRes.Success {
		if objectNotFound(showGcpDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showGcpDataCenterServerRes.ErrorMsg)
	}
	gcpDataCenterServer := showGcpDataCenterServerRes.GetData()

	if v := gcpDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if gcpDataCenterServer["properties"] != nil {
		propsJson, ok := gcpDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				_ = d.Set(propName, propValue)
			}
		}
	}

	if gcpDataCenterServer["tags"] != nil {
		tagsJson, ok := gcpDataCenterServer["tags"].([]interface{})
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

	if v := gcpDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := gcpDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := gcpDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := gcpDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementGcpDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	gcpDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		gcpDataCenterServer["name"] = oldName
		gcpDataCenterServer["new-name"] = newName
	} else {
		gcpDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("private_key"); ok {
		gcpDataCenterServer["private-key"] = d.Get("private_key")
	}

	if ok := d.HasChange("authentication_method"); ok {
		gcpDataCenterServer["authentication-method"] = d.Get("authentication_method")
		if gcpDataCenterServer["authentication-method"] == "key-authentication" {
			gcpDataCenterServer["private-key"] = d.Get("private_key")
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			gcpDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			gcpDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		gcpDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		gcpDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		gcpDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		gcpDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update gcpDataCenterServer - Map = ", gcpDataCenterServer)

	updateGcpDataCenterServerRes, err := client.ApiCall("set-data-center-server", gcpDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !updateGcpDataCenterServerRes.Success {
		if updateGcpDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateGcpDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateGcpDataCenterServerRes.GetData())
		return fmt.Errorf(msg)
	}

	return readManagementGcpDataCenterServer(d, m)
}

func deleteManagementGcpDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	gcpDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete gcpDataCenterServer")

	deleteGcpDataCenterServerRes, err := client.ApiCall("delete-data-center-server", gcpDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteGcpDataCenterServerRes.Success {
		if deleteGcpDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteGcpDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
