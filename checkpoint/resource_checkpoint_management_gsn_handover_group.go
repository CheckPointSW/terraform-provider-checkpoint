package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementGsnHandoverGroup() *schema.Resource {
	return &schema.Resource{
		Create: createManagementGsnHandoverGroup,
		Read:   readManagementGsnHandoverGroup,
		Update: updateManagementGsnHandoverGroup,
		Delete: deleteManagementGsnHandoverGroup,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"enforce_gtp": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable enforce GTP signal packet rate limit from this group.",
			},
			"gtp_rate": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Limit of the GTP rate in PDU/sec.",
			},
			"members": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of GSN handover group members identified by the name or UID.",
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

func createManagementGsnHandoverGroup(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	gsnHandoverGroup := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		gsnHandoverGroup["name"] = v.(string)
	}

	if v, ok := d.GetOkExists("enforce_gtp"); ok {
		gsnHandoverGroup["enforce-gtp"] = v.(bool)
	}

	if v, ok := d.GetOk("gtp_rate"); ok {
		gsnHandoverGroup["gtp-rate"] = v.(int)
	}

	if v, ok := d.GetOk("members"); ok {
		gsnHandoverGroup["members"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("tags"); ok {
		gsnHandoverGroup["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		gsnHandoverGroup["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		gsnHandoverGroup["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		gsnHandoverGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		gsnHandoverGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Create GsnHandoverGroup - Map = ", gsnHandoverGroup)

	addGsnHandoverGroupRes, err := client.ApiCall("add-gsn-handover-group", gsnHandoverGroup, client.GetSessionID(), true, false)
	if err != nil || !addGsnHandoverGroupRes.Success {
		if addGsnHandoverGroupRes.ErrorMsg != "" {
			return fmt.Errorf(addGsnHandoverGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addGsnHandoverGroupRes.GetData()["uid"].(string))

	return readManagementGsnHandoverGroup(d, m)
}

func readManagementGsnHandoverGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showGsnHandoverGroupRes, err := client.ApiCall("show-gsn-handover-group", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showGsnHandoverGroupRes.Success {
		if objectNotFound(showGsnHandoverGroupRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showGsnHandoverGroupRes.ErrorMsg)
	}

	gsnHandoverGroup := showGsnHandoverGroupRes.GetData()

	log.Println("Read GsnHandoverGroup - Show JSON = ", gsnHandoverGroup)

	if v := gsnHandoverGroup["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := gsnHandoverGroup["enforce-gtp"]; v != nil {
		_ = d.Set("enforce_gtp", v)
	}

	if v := gsnHandoverGroup["gtp-rate"]; v != nil {
		_ = d.Set("gtp_rate", v)
	}

	if gsnHandoverGroup["members"] != nil {
		membersJson, ok := gsnHandoverGroup["members"].([]interface{})
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

	if gsnHandoverGroup["tags"] != nil {
		tagsJson, ok := gsnHandoverGroup["tags"].([]interface{})
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

	if v := gsnHandoverGroup["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := gsnHandoverGroup["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementGsnHandoverGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	gsnHandoverGroup := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		gsnHandoverGroup["name"] = oldName
		gsnHandoverGroup["new-name"] = newName
	} else {
		gsnHandoverGroup["name"] = d.Get("name")
	}

	if v, ok := d.GetOkExists("enforce_gtp"); ok {
		gsnHandoverGroup["enforce-gtp"] = v.(bool)
	}

	if ok := d.HasChange("gtp_rate"); ok {
		gsnHandoverGroup["gtp-rate"] = d.Get("gtp_rate")
	}

	if d.HasChange("members") {
		if v, ok := d.GetOk("members"); ok {
			gsnHandoverGroup["members"] = v.(*schema.Set).List()
		} else {
			oldMembers, _ := d.GetChange("members")
			gsnHandoverGroup["members"] = map[string]interface{}{"remove": oldMembers.(*schema.Set).List()}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			gsnHandoverGroup["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			gsnHandoverGroup["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		gsnHandoverGroup["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		gsnHandoverGroup["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		gsnHandoverGroup["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		gsnHandoverGroup["ignore-errors"] = v.(bool)
	}

	log.Println("Update GsnHandoverGroup - Map = ", gsnHandoverGroup)

	updateGsnHandoverGroupRes, err := client.ApiCall("set-gsn-handover-group", gsnHandoverGroup, client.GetSessionID(), true, false)
	if err != nil || !updateGsnHandoverGroupRes.Success {
		if updateGsnHandoverGroupRes.ErrorMsg != "" {
			return fmt.Errorf(updateGsnHandoverGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementGsnHandoverGroup(d, m)
}

func deleteManagementGsnHandoverGroup(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	gsnHandoverGroupPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete GsnHandoverGroup")

	deleteGsnHandoverGroupRes, err := client.ApiCall("delete-gsn-handover-group", gsnHandoverGroupPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteGsnHandoverGroupRes.Success {
		if deleteGsnHandoverGroupRes.ErrorMsg != "" {
			return fmt.Errorf(deleteGsnHandoverGroupRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
