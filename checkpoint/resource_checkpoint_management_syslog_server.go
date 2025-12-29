package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementSyslogServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSyslogServer,
		Read:   readManagementSyslogServer,
		Update: updateManagementSyslogServer,
		Delete: deleteManagementSyslogServer,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"host": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Host server object identified by the name or UID.",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Port number.",
				Default:     514,
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "RFC version.",
				Default:     "bsd",
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementSyslogServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	syslogServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		syslogServer["name"] = v.(string)
	}

	if v, ok := d.GetOk("host"); ok {
		syslogServer["host"] = v.(string)
	}

	if v, ok := d.GetOk("port"); ok {
		syslogServer["port"] = v.(int)
	}

	if v, ok := d.GetOk("version"); ok {
		syslogServer["version"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		syslogServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		syslogServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		syslogServer["comments"] = v.(string)
	}

	if v, ok := d.GetOk("ignore_warnings"); ok {
		syslogServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOk("ignore_errors"); ok {
		syslogServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create SyslogServer - Map = ", syslogServer)

	addSyslogServerRes, err := client.ApiCallSimple("add-syslog-server", syslogServer)
	if err != nil || !addSyslogServerRes.Success {
		if addSyslogServerRes.ErrorMsg != "" {
			return fmt.Errorf(addSyslogServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addSyslogServerRes.GetData()["uid"].(string))

	return readManagementSyslogServer(d, m)
}

func readManagementSyslogServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showSyslogServerRes, err := client.ApiCallSimple("show-syslog-server", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showSyslogServerRes.Success {
		if objectNotFound(showSyslogServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showSyslogServerRes.ErrorMsg)
	}

	syslogServer := showSyslogServerRes.GetData()

	log.Println("Read SyslogServer - Show JSON = ", syslogServer)

	if v := syslogServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := syslogServer["host"]; v != nil {
		_ = d.Set("host", v.(map[string]interface{})["name"].(string))
	}

	if v := syslogServer["port"]; v != nil {
		_ = d.Set("port", v)
	}

	if v := syslogServer["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if syslogServer["tags"] != nil {
		tagsJson, ok := syslogServer["tags"].([]interface{})
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

	if v := syslogServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := syslogServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := syslogServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := syslogServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementSyslogServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	syslogServer := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		syslogServer["name"] = oldName
		syslogServer["new-name"] = newName
	} else {
		syslogServer["name"] = d.Get("name")
	}

	if ok := d.HasChange("host"); ok {
		syslogServer["host"] = d.Get("host")
	}

	if ok := d.HasChange("port"); ok {
		syslogServer["port"] = d.Get("port")
	}

	if ok := d.HasChange("version"); ok {
		syslogServer["version"] = d.Get("version")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			syslogServer["tags"] = v.(*schema.Set).List()
		}
	}

	if ok := d.HasChange("color"); ok {
		syslogServer["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		syslogServer["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		syslogServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		syslogServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update SyslogServer - Map = ", syslogServer)

	updateSyslogServerRes, err := client.ApiCallSimple("set-syslog-server", syslogServer)
	if err != nil || !updateSyslogServerRes.Success {
		if updateSyslogServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateSyslogServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementSyslogServer(d, m)
}

func deleteManagementSyslogServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	syslogServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		syslogServerPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		syslogServerPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Delete SyslogServer")

	deleteSyslogServerRes, err := client.ApiCallSimple("delete-syslog-server", syslogServerPayload)
	if err != nil || !deleteSyslogServerRes.Success {
		if deleteSyslogServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteSyslogServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
