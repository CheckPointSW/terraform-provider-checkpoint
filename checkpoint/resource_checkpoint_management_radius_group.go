package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementRadiusGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementRadiusGroup,
		Read:   readManagementRadiusGroup,
		Update: updateManagementRadiusGroup,
		Delete: deleteManagementRadiusGroup,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"members": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of radius servers identified by the name or UID.",
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
			"color": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "black",
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Comments string.",
			},
			"ignore_warnings": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
		},
	}
}

func createManagementRadiusGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	radiusGroup := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		radiusGroup["name"] = v.(string)
	}

	if v, ok := d.GetOk("members"); ok {
		radiusGroup["members"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		radiusGroup["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		radiusGroup["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		radiusGroup["comments"] = v.(string)
	}

	if v, ok := d.GetOk("ignore_warnings"); ok {
		radiusGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOk("ignore_errors"); ok {
		radiusGroup["ignore-errors"] = v.(bool)
	}

	addRadiusGroupRes, err := client.ApiCall("add-radius-group", radiusGroup, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addRadiusGroupRes.Success {
		if addRadiusGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addRadiusGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addRadiusGroupRes.GetData()["uid"].(string))
	return readManagementRadiusGroup(d, m)
}

func readManagementRadiusGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showRadiusGroupRes, err := client.ApiCall("show-radius-group", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showRadiusGroupRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showRadiusGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showRadiusGroupRes.ErrorMsg)
	}

	radiusGroup := showRadiusGroupRes.GetData()

	if v := radiusGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if radiusGroup["members"] != nil {
		membersJson := radiusGroup["members"].([]interface{})
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

	if radiusGroup["tags"] != nil {
		tagsJson := radiusGroup["tags"].([]interface{})
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

	if v := radiusGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := radiusGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementRadiusGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	radiusGroup := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		radiusGroup["name"] = oldName.(string)
		radiusGroup["new-name"] = newName.(string)
	} else {
		radiusGroup["name"] = d.Get("name")
	}

	if ok := d.HasChange("comments"); ok {
		radiusGroup["comments"] = d.Get("comments")
	}
	if ok := d.HasChange("color"); ok {
		radiusGroup["color"] = d.Get("color")
	}

	if ok := d.HasChange("members"); ok {
		if v, ok := d.GetOk("members"); ok {
			radiusGroup["members"] = v.(*schema.Set).List()
		} else {
			oldMembers, _ := d.GetChange("members")
			radiusGroup["members"] = map[string]interface{}{"remove": oldMembers.(*schema.Set).List()}
		}
	}
	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			radiusGroup["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			radiusGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		radiusGroup["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		radiusGroup["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Radius Group - Map = ", radiusGroup)
	setRadiusGroupRes, _ := client.ApiCall("set-radius-group", radiusGroup, client.GetSessionID(), true, client.IsProxyUsed())
	if !setRadiusGroupRes.Success {
		return fmt.Errorf(setRadiusGroupRes.ErrorMsg)
	}

	return readManagementRadiusGroup(d, m)
}

func deleteManagementRadiusGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid":             d.Id(),
		"ignore-warnings": "true",
	}
	deleteRadiusGroupRes, _ := client.ApiCall("delete-radius-group", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !deleteRadiusGroupRes.Success {
		return fmt.Errorf(deleteRadiusGroupRes.ErrorMsg)
	}
	d.SetId("")

	return nil
}
