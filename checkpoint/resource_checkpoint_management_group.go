package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGroup,
		Read:   readManagementGroup,
		Update: updateManagementGroup,
		Delete: deleteManagementGroup,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"members": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Network objects identified by the name or UID.",
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

func createManagementGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	group := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		group["name"] = v.(string)
	}
	if val, ok := d.GetOk("members"); ok {
		group["members"] = val.(*schema.Set).List()
	}
	if val, ok := d.GetOk("tags"); ok {
		group["tags"] = val.(*schema.Set).List()
	}

	if val, ok := d.GetOk("comments"); ok {
		group["comments"] = val.(string)
	}
	if val, ok := d.GetOk("color"); ok {
		group["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		group["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		group["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Group - Map = ", group)

	addGroupRes, err := client.ApiCall("add-group", group, client.GetSessionID(), true, false)
	if err != nil || !addGroupRes.Success {
		if addGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addGroupRes.GetData()["uid"].(string))

	return readManagementGroup(d, m)
}

func readManagementGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showGroupRes, err := client.ApiCall("show-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGroupRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showGroupRes.ErrorMsg)
	}

	group := showGroupRes.GetData()

	if v := group["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := group["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := group["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if group["members"] != nil {
		membersJson := group["members"].([]interface{})
		membersIds := make([]string, 0)
		if len(membersJson) > 0 {
			// Create slice of members names
			for _, member := range membersJson {
				member := member.(map[string]interface{})
				membersIds = append(membersIds, member["name"].(string))
			}
		}
		_ = d.Set("members", membersIds)
	} else {
		_ = d.Set("members", nil)
	}

	if group["tags"] != nil {
		tagsJson := group["tags"].([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			// Create slice of tag names
			for _, tag := range tagsJson {
				tag := tag.(map[string]interface{})
				tagsIds = append(tagsIds, tag["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	return nil
}

func updateManagementGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	group := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		group["name"] = oldName.(string)
		group["new-name"] = newName.(string)
	} else {
		group["name"] = d.Get("name")
	}

	if ok := d.HasChange("members"); ok {
		if v, ok := d.GetOk("members"); ok {
			group["members"] = v.(*schema.Set).List()
		} else {
			oldMembers, _ := d.GetChange("members")
			group["members"] = map[string]interface{}{"remove": oldMembers.(*schema.Set).List()}
		}
	}
	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			group["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			group["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("comments"); ok {
		group["comments"] = d.Get("comments")
	}
	if ok := d.HasChange("color"); ok {
		group["color"] = d.Get("color")
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		group["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		group["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Group - Map = ", group)
	setGroupRes, _ := client.ApiCall("set-group", group, client.GetSessionID(), true, false)
	if !setGroupRes.Success {
		return fmt.Errorf(setGroupRes.ErrorMsg)
	}

	return readManagementGroup(d, m)
}

func deleteManagementGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	deleteGroupRes, _ := client.ApiCall("delete-group", payload, client.GetSessionID(), true, false)
	if !deleteGroupRes.Success {
		return fmt.Errorf(deleteGroupRes.ErrorMsg)
	}
	d.SetId("")

	return nil
}
