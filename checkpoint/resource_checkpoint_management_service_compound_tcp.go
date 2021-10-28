package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementServiceCompoundTcp() *schema.Resource {
	return &schema.Resource{
		Create: createManagementServiceCompoundTcp,
		Read:   readManagementServiceCompoundTcp,
		Update: updateManagementServiceCompoundTcp,
		Delete: deleteManagementServiceCompoundTcp,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"compound_service": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Compound service type.",
				Default:     "pointcast",
			},
			"keep_connections_open_after_policy_installation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Keep connections open after policy has been installed even if they are not allowed under the new policy. This overrides the settings in the Connection Persistence page. If you change this property, the change will not affect open connections, but only future connections.",
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

func createManagementServiceCompoundTcp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	serviceCompoundTcp := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		serviceCompoundTcp["name"] = v.(string)
	}

	if v, ok := d.GetOk("compound_service"); ok {
		serviceCompoundTcp["compound-service"] = v.(string)
	}

	if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
		serviceCompoundTcp["keep-connections-open-after-policy-installation"] = v.(bool)
	}

	if v, ok := d.GetOk("tags"); ok {
		serviceCompoundTcp["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		serviceCompoundTcp["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		serviceCompoundTcp["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceCompoundTcp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceCompoundTcp["ignore-errors"] = v.(bool)
	}

	log.Println("Create ServiceCompoundTcp - Map = ", serviceCompoundTcp)

	addServiceCompoundTcpRes, err := client.ApiCall("add-service-compound-tcp", serviceCompoundTcp, client.GetSessionID(), true, false)
	if err != nil || !addServiceCompoundTcpRes.Success {
		if addServiceCompoundTcpRes.ErrorMsg != "" {
			return fmt.Errorf(addServiceCompoundTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addServiceCompoundTcpRes.GetData()["uid"].(string))

	return readManagementServiceCompoundTcp(d, m)
}

func readManagementServiceCompoundTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showServiceCompoundTcpRes, err := client.ApiCall("show-service-compound-tcp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceCompoundTcpRes.Success {
		if objectNotFound(showServiceCompoundTcpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showServiceCompoundTcpRes.ErrorMsg)
	}

	serviceCompoundTcp := showServiceCompoundTcpRes.GetData()

	log.Println("Read ServiceCompoundTcp - Show JSON = ", serviceCompoundTcp)

	if v := serviceCompoundTcp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceCompoundTcp["compound-service"]; v != nil {
		_ = d.Set("compound_service", v)
	}

	if v := serviceCompoundTcp["keep-connections-open-after-policy-installation"]; v != nil {
		_ = d.Set("keep_connections_open_after_policy_installation", v)
	}

	if serviceCompoundTcp["tags"] != nil {
		tagsJson, ok := serviceCompoundTcp["tags"].([]interface{})
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

	if v := serviceCompoundTcp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceCompoundTcp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementServiceCompoundTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	serviceCompoundTcp := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		serviceCompoundTcp["name"] = oldName
		serviceCompoundTcp["new-name"] = newName
	} else {
		serviceCompoundTcp["name"] = d.Get("name")
	}

	if ok := d.HasChange("compound_service"); ok {
		serviceCompoundTcp["compound-service"] = d.Get("compound_service")
	}

	if v, ok := d.GetOkExists("keep_connections_open_after_policy_installation"); ok {
		serviceCompoundTcp["keep-connections-open-after-policy-installation"] = v.(bool)
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			serviceCompoundTcp["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			serviceCompoundTcp["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		serviceCompoundTcp["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		serviceCompoundTcp["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceCompoundTcp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceCompoundTcp["ignore-errors"] = v.(bool)
	}

	log.Println("Update ServiceCompoundTcp - Map = ", serviceCompoundTcp)

	updateServiceCompoundTcpRes, err := client.ApiCall("set-service-compound-tcp", serviceCompoundTcp, client.GetSessionID(), true, false)
	if err != nil || !updateServiceCompoundTcpRes.Success {
		if updateServiceCompoundTcpRes.ErrorMsg != "" {
			return fmt.Errorf(updateServiceCompoundTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementServiceCompoundTcp(d, m)
}

func deleteManagementServiceCompoundTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	serviceCompoundTcpPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ServiceCompoundTcp")

	deleteServiceCompoundTcpRes, err := client.ApiCall("delete-service-compound-tcp", serviceCompoundTcpPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteServiceCompoundTcpRes.Success {
		if deleteServiceCompoundTcpRes.ErrorMsg != "" {
			return fmt.Errorf(deleteServiceCompoundTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
