package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementTacacsGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementTacacsGroup,
		Read:   readManagementTacacsGroup,
		Update: updateManagementTacacsGroup,
		Delete: deleteManagementTacacsGroup,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Must be unique in the domain",
			},
			"members": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tacacs servers identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createManagementTacacsGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	tacacsGroupPayload := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		tacacsGroupPayload["name"] = v.(string)
	}

	if v, ok := d.GetOk("members"); ok {
		tacacsGroupPayload["members"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		tacacsGroupPayload["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		tacacsGroupPayload["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		tacacsGroupPayload["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		tacacsGroupPayload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		tacacsGroupPayload["ignore-errors"] = v.(bool)
	}

	log.Println("Create Tacacs Group - Map = ", tacacsGroupPayload)

	addTacacsGroupRes, err := client.ApiCall("add-tacacs-group", tacacsGroupPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addTacacsGroupRes.Success {
		if addTacacsGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addTacacsGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addTacacsGroupRes.GetData()["uid"].(string))

	return readManagementTacacsGroup(d, m)
}

func readManagementTacacsGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showTacacsGroupRes, err := client.ApiCall("show-tacacs-group", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showTacacsGroupRes.Success {
		if objectNotFound(showTacacsGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showTacacsGroupRes.ErrorMsg)
	}

	tacacsGroup := showTacacsGroupRes.GetData()

	log.Println("Read Tacacs Group - Show JSON = ", tacacsGroup)

	if v := tacacsGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if tacacsGroup["members"] != nil {
		membersJson, ok := tacacsGroup["members"].([]interface{})
		if ok {
			membersIds := make([]string, 0)
			if len(membersJson) > 0 {
				for _, members := range membersJson {
					members := members.(map[string]interface{})
					membersIds = append(membersIds, members["name"].(string))
				}
			}
			_ = d.Set("members", membersIds)
		}
	} else {
		_ = d.Set("members", nil)
	}

	if tacacsGroup["tags"] != nil {
		tagsJson, ok := tacacsGroup["tags"].([]interface{})
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

	if v := tacacsGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := tacacsGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil

}

func updateManagementTacacsGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	tacacsGroup := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		tacacsGroup["name"] = oldName
		tacacsGroup["new-name"] = newName
	} else {
		tacacsGroup["name"] = d.Get("name")
	}

	if d.HasChange("members") {
		if v, ok := d.GetOk("members"); ok {
			tacacsGroup["members"] = v.(*schema.Set).List()
		} else {
			oldMembers, _ := d.GetChange("members")
			tacacsGroup["members"] = map[string]interface{}{"remove": oldMembers.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			tacacsGroup["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			tacacsGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		tacacsGroup["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		tacacsGroup["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		tacacsGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		tacacsGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Update Tacacs Group - Map = ", tacacsGroup)

	updateTacacsGroupRes, err := client.ApiCall("set-tacacs-group", tacacsGroup, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateTacacsGroupRes.Success {
		if updateTacacsGroupRes.ErrorMsg != "" {
			return fmt.Errorf(updateTacacsGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementTacacsGroup(d, m)
}

func deleteManagementTacacsGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	tacacsGroupPayload := map[string]interface{}{
		"uid":             d.Id(),
		"ignore-warnings": "true",
	}

	log.Println("Delete Tacacs Group")

	deleteTacacsGroupRes, err := client.ApiCall("delete-tacacs-group", tacacsGroupPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteTacacsGroupRes.Success {
		if deleteTacacsGroupRes.ErrorMsg != "" {
			return fmt.Errorf(deleteTacacsGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
