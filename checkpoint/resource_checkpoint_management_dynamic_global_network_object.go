package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementDynamicGlobalNetworkObject() *schema.Resource {
	return &schema.Resource{
		Create: createManagementDynamicGlobalNetworkObject,
		Read:   readManagementDynamicGlobalNetworkObject,
		Update: updateManagementDynamicGlobalNetworkObject,
		Delete: deleteManagementDynamicGlobalNetworkObject,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
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

func createManagementDynamicGlobalNetworkObject(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	dynamicGlobalNetworkObject := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		dynamicGlobalNetworkObject["name"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		dynamicGlobalNetworkObject["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		dynamicGlobalNetworkObject["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		dynamicGlobalNetworkObject["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dynamicGlobalNetworkObject["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dynamicGlobalNetworkObject["ignore-errors"] = v.(bool)
	}

	log.Println("Create DynamicGlobalNetworkObject - Map = ", dynamicGlobalNetworkObject)

	addDynamicGlobalNetworkObjectRes, err := client.ApiCall("add-dynamic-global-network-object", dynamicGlobalNetworkObject, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addDynamicGlobalNetworkObjectRes.Success {
		if addDynamicGlobalNetworkObjectRes.ErrorMsg != "" {
			return fmt.Errorf(addDynamicGlobalNetworkObjectRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addDynamicGlobalNetworkObjectRes.GetData()["uid"].(string))

	return readManagementDynamicGlobalNetworkObject(d, m)
}

func readManagementDynamicGlobalNetworkObject(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showDynamicGlobalNetworkObjectRes, err := client.ApiCall("show-dynamic-global-network-object", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showDynamicGlobalNetworkObjectRes.Success {
		if objectNotFound(showDynamicGlobalNetworkObjectRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showDynamicGlobalNetworkObjectRes.ErrorMsg)
	}

	dynamicGlobalNetworkObject := showDynamicGlobalNetworkObjectRes.GetData()

	log.Println("Read DynamicGlobalNetworkObject - Show JSON = ", dynamicGlobalNetworkObject)

	if v := dynamicGlobalNetworkObject["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if dynamicGlobalNetworkObject["tags"] != nil {
		tagsJson, ok := dynamicGlobalNetworkObject["tags"].([]interface{})
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

	if v := dynamicGlobalNetworkObject["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := dynamicGlobalNetworkObject["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := dynamicGlobalNetworkObject["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := dynamicGlobalNetworkObject["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementDynamicGlobalNetworkObject(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	dynamicGlobalNetworkObject := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		dynamicGlobalNetworkObject["name"] = oldName
		dynamicGlobalNetworkObject["new-name"] = newName
	} else {
		dynamicGlobalNetworkObject["name"] = d.Get("name")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			dynamicGlobalNetworkObject["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			dynamicGlobalNetworkObject["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		dynamicGlobalNetworkObject["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		dynamicGlobalNetworkObject["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dynamicGlobalNetworkObject["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dynamicGlobalNetworkObject["ignore-errors"] = v.(bool)
	}

	log.Println("Update DynamicGlobalNetworkObject - Map = ", dynamicGlobalNetworkObject)

	updateDynamicGlobalNetworkObjectRes, err := client.ApiCall("set-dynamic-global-network-object", dynamicGlobalNetworkObject, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateDynamicGlobalNetworkObjectRes.Success {
		if updateDynamicGlobalNetworkObjectRes.ErrorMsg != "" {
			return fmt.Errorf(updateDynamicGlobalNetworkObjectRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementDynamicGlobalNetworkObject(d, m)
}

func deleteManagementDynamicGlobalNetworkObject(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	dynamicGlobalNetworkObjectPayload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		dynamicGlobalNetworkObjectPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		dynamicGlobalNetworkObjectPayload["ignore-errors"] = v.(bool)
	}
	log.Println("Delete DynamicGlobalNetworkObject")

	deleteDynamicGlobalNetworkObjectRes, err := client.ApiCall("delete-dynamic-global-network-object", dynamicGlobalNetworkObjectPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteDynamicGlobalNetworkObjectRes.Success {
		if deleteDynamicGlobalNetworkObjectRes.ErrorMsg != "" {
			return fmt.Errorf(deleteDynamicGlobalNetworkObjectRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
