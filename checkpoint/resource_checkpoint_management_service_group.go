package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementServiceGroup,
		Read:   readManagementServiceGroup,
		Update: updateManagementServiceGroup,
		Delete: deleteManagementServiceGroup,
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

func createManagementServiceGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	serviceGroup := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		serviceGroup["name"] = v.(string)
	}
	if val, ok := d.GetOk("members"); ok {
		serviceGroup["members"] = val.(*schema.Set).List()
	}
	if val, ok := d.GetOk("tags"); ok {
		serviceGroup["tags"] = val.(*schema.Set).List()
	}

	if val, ok := d.GetOk("comments"); ok {
		serviceGroup["comments"] = val.(string)
	}
	if val, ok := d.GetOk("color"); ok {
		serviceGroup["color"] = val.(string)
	}
	if val, ok := d.GetOkExists("ignore_errors"); ok {
		serviceGroup["ignore-errors"] = val.(bool)
	}
	if val, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceGroup["ignore-warnings"] = val.(bool)
	}

	log.Println("Create Service Group - Map = ", serviceGroup)

	addServiceGroupRes, err := client.ApiCall("add-service-group", serviceGroup, client.GetSessionID(), true, false)
	if err != nil || !addServiceGroupRes.Success {
		if addServiceGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addServiceGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addServiceGroupRes.GetData()["uid"].(string))

	return readManagementServiceGroup(d, m)
}

func readManagementServiceGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showServiceGroupRes, err := client.ApiCall("show-service-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showServiceGroupRes.Success {
		// Handle delete resource from other clients
		if objectNotFound(showServiceGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showServiceGroupRes.ErrorMsg)
	}

	serviceGroup := showServiceGroupRes.GetData()

	if v := serviceGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := serviceGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := serviceGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if serviceGroup["members"] != nil {
		membersJson := serviceGroup["members"].([]interface{})
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

	if serviceGroup["tags"] != nil {
		tagsJson := serviceGroup["tags"].([]interface{})
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

func updateManagementServiceGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	serviceGroup := make(map[string]interface{})

	if d.HasChange("name") {
		oldName, newName := d.GetChange("name")
		serviceGroup["name"] = oldName.(string)
		serviceGroup["new-name"] = newName.(string)
	} else {
		serviceGroup["name"] = d.Get("name")
	}

	if ok := d.HasChange("members"); ok {
		if v, ok := d.GetOk("members"); ok {
			serviceGroup["members"] = v.(*schema.Set).List()
		} else {
			oldMembers, _ := d.GetChange("members")
			serviceGroup["members"] = map[string]interface{}{"remove": oldMembers.(*schema.Set).List()}
		}
	}
	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			serviceGroup["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			serviceGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("comments"); ok {
		serviceGroup["comments"] = d.Get("comments")
	}

	if ok := d.HasChange("color"); ok {
		serviceGroup["color"] = d.Get("color")
	}
	if v, ok := d.GetOkExists("ignore_errors"); ok {
		serviceGroup["ignore-errors"] = v.(bool)
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		serviceGroup["ignore-warnings"] = v.(bool)
	}

	log.Println("Update Service Group - Map = ", serviceGroup)
	setserviceGroupRes, _ := client.ApiCall("set-service-group", serviceGroup, client.GetSessionID(), true, false)
	if !setserviceGroupRes.Success {
		return fmt.Errorf(setserviceGroupRes.ErrorMsg)
	}

	return readManagementServiceGroup(d, m)
}

func deleteManagementServiceGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	deleteServiceGroupRes, _ := client.ApiCall("delete-service-group", payload, client.GetSessionID(), true, false)
	if !deleteServiceGroupRes.Success {
		return fmt.Errorf(deleteServiceGroupRes.ErrorMsg)
	}
	d.SetId("")

	return nil
}
