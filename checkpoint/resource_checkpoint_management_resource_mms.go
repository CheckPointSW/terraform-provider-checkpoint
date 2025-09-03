package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementResourceMms() *schema.Resource {
	return &schema.Resource{
		Create: createManagementResourceMms,
		Read:   readManagementResourceMms,
		Update: updateManagementResourceMms,
		Delete: deleteManagementResourceMms,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"track": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Logs the activity when a packet matches on a Firewall Rule with the Resource.",
				Default:     "none",
			},
			"action": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Accepts or Drops traffic that matches a Firewall Rule using the Resource.",
				Default:     "accept",
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

func createManagementResourceMms(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	resourceMms := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		resourceMms["name"] = v.(string)
	}

	if v, ok := d.GetOk("track"); ok {
		resourceMms["track"] = v.(string)
	}

	if v, ok := d.GetOk("action"); ok {
		resourceMms["action"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		resourceMms["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		resourceMms["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		resourceMms["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceMms["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceMms["ignore-errors"] = v.(bool)
	}

	log.Println("Create ResourceMms - Map = ", resourceMms)

	addResourceMmsRes, err := client.ApiCallSimple("add-resource-mms", resourceMms)
	if err != nil || !addResourceMmsRes.Success {
		if addResourceMmsRes.ErrorMsg != "" {
			return fmt.Errorf(addResourceMmsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addResourceMmsRes.GetData()["uid"].(string))

	return readManagementResourceMms(d, m)
}

func readManagementResourceMms(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showResourceMmsRes, err := client.ApiCallSimple("show-resource-mms", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceMmsRes.Success {
		if objectNotFound(showResourceMmsRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showResourceMmsRes.ErrorMsg)
	}

	resourceMms := showResourceMmsRes.GetData()

	log.Println("Read ResourceMms - Show JSON = ", resourceMms)

	if v := resourceMms["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceMms["track"]; v != nil {
		_ = d.Set("track", v)
	}

	if v := resourceMms["action"]; v != nil {
		_ = d.Set("action", v)
	}

	if resourceMms["tags"] != nil {
		tagsJson, ok := resourceMms["tags"].([]interface{})
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

	if v := resourceMms["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceMms["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := resourceMms["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := resourceMms["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementResourceMms(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	resourceMms := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		resourceMms["name"] = oldName
		resourceMms["new-name"] = newName
	} else {
		resourceMms["name"] = d.Get("name")
	}

	if ok := d.HasChange("track"); ok {
		resourceMms["track"] = d.Get("track")
	}

	if ok := d.HasChange("action"); ok {
		resourceMms["action"] = d.Get("action")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			resourceMms["tags"] = v.(*schema.Set).List()
		}
	}

	if ok := d.HasChange("color"); ok {
		resourceMms["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		resourceMms["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceMms["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceMms["ignore-errors"] = v.(bool)
	}

	log.Println("Update ResourceMms - Map = ", resourceMms)

	updateResourceMmsRes, err := client.ApiCallSimple("set-resource-mms", resourceMms)
	if err != nil || !updateResourceMmsRes.Success {
		if updateResourceMmsRes.ErrorMsg != "" {
			return fmt.Errorf(updateResourceMmsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementResourceMms(d, m)
}

func deleteManagementResourceMms(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	resourceMmsPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ResourceMms")

	deleteResourceMmsRes, err := client.ApiCallSimple("delete-resource-mms", resourceMmsPayload)
	if err != nil || !deleteResourceMmsRes.Success {
		if deleteResourceMmsRes.ErrorMsg != "" {
			return fmt.Errorf(deleteResourceMmsRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
