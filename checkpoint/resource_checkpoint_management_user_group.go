package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementUserGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementUserGroup,
		Read:   readManagementUserGroup,
		Update: updateManagementUserGroup,
		Delete: deleteManagementUserGroup,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"email": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email Address.",
			},
			"members": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of User Group objects identified by the name or UID.",
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

func createManagementUserGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	userGroup := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		userGroup["name"] = v.(string)
	}

	if v, ok := d.GetOk("email"); ok {
		userGroup["email"] = v.(string)
	}

	if v, ok := d.GetOk("members"); ok {
		userGroup["members"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		userGroup["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		userGroup["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		userGroup["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		userGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		userGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Create UserGroup - Map = ", userGroup)

	addUserGroupRes, err := client.ApiCall("add-user-group", userGroup, client.GetSessionID(), true, false)
	if err != nil || !addUserGroupRes.Success {
		if addUserGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addUserGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addUserGroupRes.GetData()["uid"].(string))

	return readManagementUserGroup(d, m)
}

func readManagementUserGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showUserGroupRes, err := client.ApiCall("show-user-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showUserGroupRes.Success {
		if objectNotFound(showUserGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showUserGroupRes.ErrorMsg)
	}

	userGroup := showUserGroupRes.GetData()

	log.Println("Read UserGroup - Show JSON = ", userGroup)

	if v := userGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := userGroup["email"]; v != nil {
		_ = d.Set("email", v)
	}

	if userGroup["members"] != nil {
		membersJson, ok := userGroup["members"].([]interface{})
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

	if userGroup["tags"] != nil {
		tagsJson, ok := userGroup["tags"].([]interface{})
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

	if v := userGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := userGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementUserGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	userGroup := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		userGroup["name"] = oldName
		userGroup["new-name"] = newName
	} else {
		userGroup["name"] = d.Get("name")
	}

	if ok := d.HasChange("email"); ok {
		userGroup["email"] = d.Get("email")
	}

	if d.HasChange("members") {
		if v, ok := d.GetOk("members"); ok {
			userGroup["members"] = v.(*schema.Set).List()
		} else {
			oldMembers, _ := d.GetChange("members")
			userGroup["members"] = map[string]interface{}{"remove": oldMembers.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			userGroup["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			userGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		userGroup["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		userGroup["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		userGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		userGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Update UserGroup - Map = ", userGroup)

	updateUserGroupRes, err := client.ApiCall("set-user-group", userGroup, client.GetSessionID(), true, false)
	if err != nil || !updateUserGroupRes.Success {
		if updateUserGroupRes.ErrorMsg != "" {
			return fmt.Errorf(updateUserGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementUserGroup(d, m)
}

func deleteManagementUserGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	userGroupPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete UserGroup")

	deleteUserGroupRes, err := client.ApiCall("delete-user-group", userGroupPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteUserGroupRes.Success {
		if deleteUserGroupRes.ErrorMsg != "" {
			return fmt.Errorf(deleteUserGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
