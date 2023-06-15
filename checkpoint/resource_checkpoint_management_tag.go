package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementTag() *schema.Resource {
	return &schema.Resource{
		Create: createManagementTag,
		Read:   readManagementTag,
		Update: updateManagementTag,
		Delete: deleteManagementTag,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain.",
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

func createManagementTag(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	tag := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		tag["name"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		tag["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		tag["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		tag["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		tag["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		tag["ignore-errors"] = v.(bool)
	}

	log.Println("Create Tag - Map = ", tag)

	addTagRes, err := client.ApiCall("add-tag", tag, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addTagRes.Success {
		if addTagRes.ErrorMsg != "" {
			return fmt.Errorf(addTagRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addTagRes.GetData()["uid"].(string))

	return readManagementTag(d, m)
}

func readManagementTag(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showTagRes, err := client.ApiCall("show-tag", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTagRes.Success {
		if objectNotFound(showTagRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showTagRes.ErrorMsg)
	}

	tag := showTagRes.GetData()

	log.Println("Read Tag - Show JSON = ", tag)

	if v := tag["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if tag["tags"] != nil {
		tagsJson, ok := tag["tags"].([]interface{})
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

	if v := tag["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := tag["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := tag["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := tag["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementTag(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	tag := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		tag["name"] = oldName
		tag["new-name"] = newName
	} else {
		tag["name"] = d.Get("name")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			tag["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			tag["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		tag["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		tag["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		tag["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		tag["ignore-errors"] = v.(bool)
	}

	log.Println("Update Tag - Map = ", tag)

	updateTagRes, err := client.ApiCall("set-tag", tag, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateTagRes.Success {
		if updateTagRes.ErrorMsg != "" {
			return fmt.Errorf(updateTagRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementTag(d, m)
}

func deleteManagementTag(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	tagPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		tagPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		tagPayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete Tag")

	deleteTagRes, err := client.ApiCall("delete-tag", tagPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteTagRes.Success {
		if deleteTagRes.ErrorMsg != "" {
			return fmt.Errorf(deleteTagRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
