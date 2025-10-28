package checkpoint

import (
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementLogicalServer() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLogicalServer,
		Read:   readManagementLogicalServer,
		Update: updateManagementLogicalServer,
		Delete: deleteManagementLogicalServer,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv6 address.",
			},
			"server_group": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Server group associated with the logical server.  Identified by name or UID.",
			},
			"server_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Type of server for the logical server.",
				Default:     "other",
			},
			"persistence_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates if persistence mode is enabled for the logical server.",
				Default:     true,
			},
			"persistency_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Persistency type for the logical server.",
				Default:     "by_service",
			},
			"balance_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Load balancing method for the logical server.",
				Default:     "random",
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

func createManagementLogicalServer(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	logicalServer := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		logicalServer["name"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		logicalServer["ipv4-address"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address"); ok {
		logicalServer["ipv6-address"] = v.(string)
	}

	if v, ok := d.GetOk("server_group"); ok {
		logicalServer["server-group"] = v.(string)
	}

	if v, ok := d.GetOk("server_type"); ok {
		logicalServer["server-type"] = v.(string)
	}

	if v, ok := d.GetOkExists("persistence_mode"); ok {
		logicalServer["persistence-mode"] = v.(bool)
	}

	if v, ok := d.GetOk("persistency_type"); ok {
		logicalServer["persistency-type"] = v.(string)
	}

	if v, ok := d.GetOk("balance_method"); ok {
		logicalServer["balance-method"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		logicalServer["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		logicalServer["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		logicalServer["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		logicalServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		logicalServer["ignore-errors"] = v.(bool)
	}

	log.Println("Create LogicalServer - Map = ", logicalServer)

	addLogicalServerRes, err := client.ApiCall("add-logical-server", logicalServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addLogicalServerRes.Success {
		if addLogicalServerRes.ErrorMsg != "" {
			return fmt.Errorf(addLogicalServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addLogicalServerRes.GetData()["uid"].(string))

	return readManagementLogicalServer(d, m)
}

func readManagementLogicalServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showLogicalServerRes, err := client.ApiCall("show-logical-server", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showLogicalServerRes.Success {
		if objectNotFound(showLogicalServerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showLogicalServerRes.ErrorMsg)
	}

	logicalServer := showLogicalServerRes.GetData()

	log.Println("Read LogicalServer - Show JSON = ", logicalServer)

	if v := logicalServer["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := logicalServer["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := logicalServer["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := logicalServer["server-group"]; v != nil {
		_ = d.Set("server_group", v)
	}

	if v := logicalServer["server-type"]; v != nil {
		_ = d.Set("server_type", v)
	}

	if v := logicalServer["persistence-mode"]; v != nil {
		_ = d.Set("persistence_mode", v)
	}

	if v := logicalServer["persistency-type"]; v != nil {
		_ = d.Set("persistency_type", v)
	}

	if v := logicalServer["balance-method"]; v != nil {
		_ = d.Set("balance_method", v)
	}

	if v := logicalServer["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := logicalServer["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if logicalServer["tags"] != nil {
		tagsJson, ok := logicalServer["tags"].([]interface{})
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

	if v := logicalServer["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := logicalServer["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementLogicalServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	logicalServer := make(map[string]interface{})

	logicalServer["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			logicalServer["new-name"] = v.(string)
		}
	}

	if ok := d.HasChange("ipv4_address"); ok {
		if v, ok := d.GetOk("ipv4_address"); ok {
			logicalServer["ipv4-address"] = v.(string)
		}
	}

	if ok := d.HasChange("ipv6_address"); ok {
		if v, ok := d.GetOk("ipv6_address"); ok {
			logicalServer["ipv6-address"] = v.(string)
		}
	}

	if ok := d.HasChange("server_group"); ok {
		if v, ok := d.GetOk("server_group"); ok {
			logicalServer["server-group"] = v.(string)
		}
	}

	if ok := d.HasChange("server_type"); ok {
		if v, ok := d.GetOk("server_type"); ok {
			logicalServer["server-type"] = v.(string)
		}
	}

	if v, ok := d.GetOkExists("persistence_mode"); ok {
		logicalServer["persistence-mode"] = v.(bool)
	}

	if ok := d.HasChange("persistency_type"); ok {
		if v, ok := d.GetOk("persistency_type"); ok {
			logicalServer["persistency-type"] = v.(string)
		}
	}

	if ok := d.HasChange("balance_method"); ok {
		if v, ok := d.GetOk("balance_method"); ok {
			logicalServer["balance-method"] = v.(string)
		}
	}

	if ok := d.HasChange("color"); ok {
		if v, ok := d.GetOk("color"); ok {
			logicalServer["color"] = v.(string)
		}
	}

	if ok := d.HasChange("comments"); ok {
		if v, ok := d.GetOk("comments"); ok {
			logicalServer["comments"] = v.(string)
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			logicalServer["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		logicalServer["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		logicalServer["ignore-errors"] = v.(bool)
	}

	log.Println("Update LogicalServer - Map = ", logicalServer)

	updateLogicalServerRes, err := client.ApiCall("set-logical-server", logicalServer, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateLogicalServerRes.Success {
		if updateLogicalServerRes.ErrorMsg != "" {
			return fmt.Errorf(updateLogicalServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementLogicalServer(d, m)
}

func deleteManagementLogicalServer(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	logicalServerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete LogicalServer")

	deleteLogicalServerRes, err := client.ApiCall("delete-logical-server", logicalServerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteLogicalServerRes.Success {
		if deleteLogicalServerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteLogicalServerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
