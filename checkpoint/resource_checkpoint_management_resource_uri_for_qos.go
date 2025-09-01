package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementResourceUriForQos() *schema.Resource {
	return &schema.Resource{
		Create: createManagementResourceUriForQos,
		Read:   readManagementResourceUriForQos,
		Update: updateManagementResourceUriForQos,
		Delete: deleteManagementResourceUriForQos,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"search_for_url": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "URL string that will be matched to an HTTP connection.",
				Default:     "*",
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

func createManagementResourceUriForQos(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	resourceUriForQos := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		resourceUriForQos["name"] = v.(string)
	}

	if v, ok := d.GetOk("search_for_url"); ok {
		resourceUriForQos["search-for-url"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		resourceUriForQos["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		resourceUriForQos["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		resourceUriForQos["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceUriForQos["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceUriForQos["ignore-errors"] = v.(bool)
	}

	log.Println("Create ResourceUriForQos - Map = ", resourceUriForQos)

	addResourceUriForQosRes, err := client.ApiCallSimple("add-resource-uri-for-qos", resourceUriForQos)
	if err != nil || !addResourceUriForQosRes.Success {
		if addResourceUriForQosRes.ErrorMsg != "" {
			return fmt.Errorf(addResourceUriForQosRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addResourceUriForQosRes.GetData()["uid"].(string))

	return readManagementResourceUriForQos(d, m)
}

func readManagementResourceUriForQos(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showResourceUriForQosRes, err := client.ApiCallSimple("show-resource-uri-for-qos", payload)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceUriForQosRes.Success {
		if objectNotFound(showResourceUriForQosRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showResourceUriForQosRes.ErrorMsg)
	}

	resourceUriForQos := showResourceUriForQosRes.GetData()

	log.Println("Read ResourceUriForQos - Show JSON = ", resourceUriForQos)

	if v := resourceUriForQos["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceUriForQos["search-for-url"]; v != nil {
		_ = d.Set("search_for_url", v)
	}

	if resourceUriForQos["tags"] != nil {
		tagsJson, ok := resourceUriForQos["tags"].([]interface{})
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

	if v := resourceUriForQos["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceUriForQos["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := resourceUriForQos["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := resourceUriForQos["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementResourceUriForQos(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	resourceUriForQos := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		resourceUriForQos["name"] = oldName
		resourceUriForQos["new-name"] = newName
	} else {
		resourceUriForQos["name"] = d.Get("name")
	}

	if ok := d.HasChange("search_for_url"); ok {
		resourceUriForQos["search-for-url"] = d.Get("search_for_url")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			resourceUriForQos["tags"] = v.(*schema.Set).List()
		}
	}

	if ok := d.HasChange("color"); ok {
		resourceUriForQos["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		resourceUriForQos["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceUriForQos["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceUriForQos["ignore-errors"] = v.(bool)
	}

	log.Println("Update ResourceUriForQos - Map = ", resourceUriForQos)

	updateResourceUriForQosRes, err := client.ApiCallSimple("set-resource-uri-for-qos", resourceUriForQos)
	if err != nil || !updateResourceUriForQosRes.Success {
		if updateResourceUriForQosRes.ErrorMsg != "" {
			return fmt.Errorf(updateResourceUriForQosRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementResourceUriForQos(d, m)
}

func deleteManagementResourceUriForQos(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	resourceUriForQosPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ResourceUriForQos")

	deleteResourceUriForQosRes, err := client.ApiCallSimple("delete-resource-uri-for-qos", resourceUriForQosPayload)
	if err != nil || !deleteResourceUriForQosRes.Success {
		if deleteResourceUriForQosRes.ErrorMsg != "" {
			return fmt.Errorf(deleteResourceUriForQosRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
