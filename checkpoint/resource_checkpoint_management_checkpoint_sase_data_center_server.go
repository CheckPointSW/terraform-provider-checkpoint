package checkpoint

import (
	"fmt"
	"log"
	"strings"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceManagementCheckpointSaseDataCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementCheckpointSaseDataCenterServer,
		Read:   readManagementCheckpointSaseDataCenterServer,
		Update: updateManagementCheckpointSaseDataCenterServer,
		Delete: deleteManagementCheckpointSaseDataCenterServer,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
			},
			"connect_to": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "connected-portal: Connect to the connected Check Point Portal Account. other-portal: Connect to a different Check Point Portal Account.",
			},
			"hostname": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL from Check Point Portal. Required for connect-to: other-portal.",
			},
			"client_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Client ID for Check Point SASE account. Required for connect-to: other-portal.",
			},
			"secret_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "Secret key for Check Point SASE account. Required for connect-to: other-portal.",
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

func createManagementCheckpointSaseDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	checkpointSaseDataCenterServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		checkpointSaseDataCenterServer["name"] = v.(string)
	}

	checkpointSaseDataCenterServer["type"] = "checkpoint-sase"

	if v, ok := d.GetOk("connect_to"); ok {
		checkpointSaseDataCenterServer["connect-to"] = v.(string)
	}

	if v, ok := d.GetOk("hostname"); ok {
		checkpointSaseDataCenterServer["hostname"] = v.(string)
	}

	if v, ok := d.GetOk("client_id"); ok {
		checkpointSaseDataCenterServer["client-id"] = v.(string)
	}

	if v, ok := d.GetOk("secret_key"); ok {
		checkpointSaseDataCenterServer["secret-key"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		checkpointSaseDataCenterServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		checkpointSaseDataCenterServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		checkpointSaseDataCenterServer["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		checkpointSaseDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		checkpointSaseDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create checkpointSaseDataCenterServer - Map = ", checkpointSaseDataCenterServer)

	addCheckpointSaseDataCenterServerRes, err := client.ApiCall("add-data-center-server", checkpointSaseDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !addCheckpointSaseDataCenterServerRes.Success {
		if addCheckpointSaseDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf("%s", addCheckpointSaseDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-data-center-server", addCheckpointSaseDataCenterServerRes.GetData())
		return fmt.Errorf("%s", msg)
	}
	payload := map[string]interface{}{
		"name": checkpointSaseDataCenterServer["name"],
	}
	showCheckpointSaseDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showCheckpointSaseDataCenterServerRes.Success {
		return fmt.Errorf("%s", showCheckpointSaseDataCenterServerRes.ErrorMsg)
	}
	d.SetId(showCheckpointSaseDataCenterServerRes.GetData()["uid"].(string))
	return readManagementCheckpointSaseDataCenterServer(d, m)
}

func readManagementCheckpointSaseDataCenterServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showCheckpointSaseDataCenterServerRes, err := client.ApiCall("show-data-center-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showCheckpointSaseDataCenterServerRes.Success {
		if objectNotFound(showCheckpointSaseDataCenterServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showCheckpointSaseDataCenterServerRes.ErrorMsg)
	}
	checkpointSaseDataCenterServer := showCheckpointSaseDataCenterServerRes.GetData()

	if v := checkpointSaseDataCenterServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if checkpointSaseDataCenterServer["properties"] != nil {
		propsJson, ok := checkpointSaseDataCenterServer["properties"].([]interface{})
		if ok {
			for _, prop := range propsJson {
				propMap := prop.(map[string]interface{})
				propName := strings.ReplaceAll(propMap["name"].(string), "-", "_")
				propValue := propMap["value"]
				_ = d.Set(propName, propValue)
			}
		}
	}

	if checkpointSaseDataCenterServer["tags"] != nil {
		tagsJson, ok := checkpointSaseDataCenterServer["tags"].([]interface{})
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

	if v := checkpointSaseDataCenterServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := checkpointSaseDataCenterServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := checkpointSaseDataCenterServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := checkpointSaseDataCenterServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementCheckpointSaseDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	checkpointSaseDataCenterServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		checkpointSaseDataCenterServer["name"] = oldName
		checkpointSaseDataCenterServer["new-name"] = newName
	} else {
		checkpointSaseDataCenterServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("secret_key"); ok {
		checkpointSaseDataCenterServer["secret-key"] = d.Get("secret_key")
	}

	if ok := d.HasChange("connect_to"); ok {
		checkpointSaseDataCenterServer["connect-to"] = d.Get("connect_to")
		if checkpointSaseDataCenterServer["connect-to"] == "other-portal" {
			checkpointSaseDataCenterServer["secret-key"] = d.Get("secret_key")
		}
	}

	if ok := d.HasChange("hostname"); ok {
		checkpointSaseDataCenterServer["hostname"] = d.Get("hostname")
	}

	if ok := d.HasChange("client_id"); ok {
		checkpointSaseDataCenterServer["client-id"] = d.Get("client_id")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			checkpointSaseDataCenterServer["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			checkpointSaseDataCenterServer["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		checkpointSaseDataCenterServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		checkpointSaseDataCenterServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		checkpointSaseDataCenterServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		checkpointSaseDataCenterServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update checkpointSaseDataCenterServer - Map = ", checkpointSaseDataCenterServer)

	updateCheckpointSaseDataCenterServerRes, err := client.ApiCall("set-data-center-server", checkpointSaseDataCenterServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !updateCheckpointSaseDataCenterServerRes.Success {
		if updateCheckpointSaseDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf("%s", updateCheckpointSaseDataCenterServerRes.ErrorMsg)
		}
		msg := createTaskFailMessage("set-data-center-server", updateCheckpointSaseDataCenterServerRes.GetData())
		return fmt.Errorf("%s", msg)
	}

	return readManagementCheckpointSaseDataCenterServer(d, m)
}

func deleteManagementCheckpointSaseDataCenterServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	checkpointSaseDataCenterServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		checkpointSaseDataCenterServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		checkpointSaseDataCenterServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete checkpointSaseDataCenterServer")

	deleteCheckpointSaseDataCenterServerRes, err := client.ApiCall("delete-data-center-server", checkpointSaseDataCenterServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteCheckpointSaseDataCenterServerRes.Success {
		if deleteCheckpointSaseDataCenterServerRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteCheckpointSaseDataCenterServerRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
