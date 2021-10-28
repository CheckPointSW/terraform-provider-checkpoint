package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementServiceCitrixTcp() *schema.Resource {
	return &schema.Resource{
		Create: createManagementServiceCitrixTcp,
		Read:   readManagementServiceCitrixTcp,
		Update: updateManagementServiceCitrixTcp,
		Delete: deleteManagementServiceCitrixTcp,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"application": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Citrix application name.",
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

func createManagementServiceCitrixTcp(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	serviceCitrixTcp := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		serviceCitrixTcp["name"] = v.(string)
	}

	if v, ok := d.GetOk("application"); ok {
		serviceCitrixTcp["application"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		serviceCitrixTcp["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		serviceCitrixTcp["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		serviceCitrixTcp["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceCitrixTcp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceCitrixTcp["ignore-errors"] = v.(bool)
	}

	log.Println("Create ServiceCitrixTcp - Map = ", serviceCitrixTcp)

	addServiceCitrixTcpRes, err := client.ApiCall("add-service-citrix-tcp", serviceCitrixTcp, client.GetSessionID(), true, false)
	if err != nil || !addServiceCitrixTcpRes.Success {
		if addServiceCitrixTcpRes.ErrorMsg != "" {
			return fmt.Errorf(addServiceCitrixTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addServiceCitrixTcpRes.GetData()["uid"].(string))

	return readManagementServiceCitrixTcp(d, m)
}

func readManagementServiceCitrixTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showServiceCitrixTcpRes, err := client.ApiCall("show-service-citrix-tcp", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceCitrixTcpRes.Success {
		if objectNotFound(showServiceCitrixTcpRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showServiceCitrixTcpRes.ErrorMsg)
	}

	serviceCitrixTcp := showServiceCitrixTcpRes.GetData()

	log.Println("Read ServiceCitrixTcp - Show JSON = ", serviceCitrixTcp)

	if v := serviceCitrixTcp["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceCitrixTcp["application"]; v != nil {
		_ = d.Set("application", v)
	}

	if serviceCitrixTcp["tags"] != nil {
		tagsJson, ok := serviceCitrixTcp["tags"].([]interface{})
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

	if v := serviceCitrixTcp["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := serviceCitrixTcp["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementServiceCitrixTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	serviceCitrixTcp := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		serviceCitrixTcp["name"] = oldName
		serviceCitrixTcp["new-name"] = newName
	} else {
		serviceCitrixTcp["name"] = d.Get("name")
	}

	if ok := d.HasChange("application"); ok {
		serviceCitrixTcp["application"] = d.Get("application")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			serviceCitrixTcp["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			serviceCitrixTcp["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		serviceCitrixTcp["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		serviceCitrixTcp["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceCitrixTcp["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceCitrixTcp["ignore-errors"] = v.(bool)
	}

	log.Println("Update ServiceCitrixTcp - Map = ", serviceCitrixTcp)

	updateServiceCitrixTcpRes, err := client.ApiCall("set-service-citrix-tcp", serviceCitrixTcp, client.GetSessionID(), true, false)
	if err != nil || !updateServiceCitrixTcpRes.Success {
		if updateServiceCitrixTcpRes.ErrorMsg != "" {
			return fmt.Errorf(updateServiceCitrixTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementServiceCitrixTcp(d, m)
}

func deleteManagementServiceCitrixTcp(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	serviceCitrixTcpPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ServiceCitrixTcp")

	deleteServiceCitrixTcpRes, err := client.ApiCall("delete-service-citrix-tcp", serviceCitrixTcpPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteServiceCitrixTcpRes.Success {
		if deleteServiceCitrixTcpRes.ErrorMsg != "" {
			return fmt.Errorf(deleteServiceCitrixTcpRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
