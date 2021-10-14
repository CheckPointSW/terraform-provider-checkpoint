package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementIdentityTag() *schema.Resource {
	return &schema.Resource{
		Create: createManagementIdentityTag,
		Read:   readManagementIdentityTag,
		Update: updateManagementIdentityTag,
		Delete: deleteManagementIdentityTag,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"external_identifier": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "External identifier. For example: Cisco ISE security group tag.",
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

func createManagementIdentityTag(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	identityTag := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		identityTag["name"] = v.(string)
	}

	if v, ok := d.GetOk("external_identifier"); ok {
		identityTag["external-identifier"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		identityTag["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		identityTag["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		identityTag["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		identityTag["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		identityTag["ignore-errors"] = v.(bool)
	}

	log.Println("Create IdentityTag - Map = ", identityTag)

	addIdentityTagRes, err := client.ApiCall("add-identity-tag", identityTag, client.GetSessionID(), true, false)
	if err != nil || !addIdentityTagRes.Success {
		if addIdentityTagRes.ErrorMsg != "" {
			return fmt.Errorf(addIdentityTagRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addIdentityTagRes.GetData()["uid"].(string))

	return readManagementIdentityTag(d, m)
}

func readManagementIdentityTag(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showIdentityTagRes, err := client.ApiCall("show-identity-tag", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showIdentityTagRes.Success {
		if objectNotFound(showIdentityTagRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showIdentityTagRes.ErrorMsg)
	}

	identityTag := showIdentityTagRes.GetData()

	log.Println("Read IdentityTag - Show JSON = ", identityTag)

	if v := identityTag["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := identityTag["external-identifier"]; v != nil {
		_ = d.Set("external_identifier", v)
	}

	if identityTag["tags"] != nil {
		tagsJson, ok := identityTag["tags"].([]interface{})
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

	if v := identityTag["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := identityTag["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementIdentityTag(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	identityTag := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		identityTag["name"] = oldName
		identityTag["new-name"] = newName
	} else {
		identityTag["name"] = d.Get("name")
	}

	if ok := d.HasChange("external_identifier"); ok {
		identityTag["external-identifier"] = d.Get("external_identifier")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			identityTag["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			identityTag["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		identityTag["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		identityTag["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		identityTag["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		identityTag["ignore-errors"] = v.(bool)
	}

	log.Println("Update IdentityTag - Map = ", identityTag)

	updateIdentityTagRes, err := client.ApiCall("set-identity-tag", identityTag, client.GetSessionID(), true, false)
	if err != nil || !updateIdentityTagRes.Success {
		if updateIdentityTagRes.ErrorMsg != "" {
			return fmt.Errorf(updateIdentityTagRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementIdentityTag(d, m)
}

func deleteManagementIdentityTag(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	identityTagPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete IdentityTag")

	deleteIdentityTagRes, err := client.ApiCall("delete-identity-tag", identityTagPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteIdentityTagRes.Success {
		if deleteIdentityTagRes.ErrorMsg != "" {
			return fmt.Errorf(deleteIdentityTagRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
