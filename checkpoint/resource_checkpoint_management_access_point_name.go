package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func resourceManagementAccessPointName() *schema.Resource {
	return &schema.Resource{
		Create: createManagementAccessPointName,
		Read:   readManagementAccessPointName,
		Update: updateManagementAccessPointName,
		Delete: deleteManagementAccessPointName,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"apn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "APN name.",
			},
			"enforce_end_user_domain": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enable enforce end user domain.",
			},
			"block_traffic_other_end_user_domains": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Block MS to MS traffic between this and other APN end user domains.",
				Default:     true,
			},
			"block_traffic_this_end_user_domain": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Block MS to MS traffic within this end user domain.",
				Default:     true,
			},
			"end_user_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "End user domain name or UID.",
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

func createManagementAccessPointName(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	accessPointName := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		accessPointName["name"] = v.(string)
	}

	if v, ok := d.GetOk("apn"); ok {
		accessPointName["apn"] = v.(string)
	}

	if v, ok := d.GetOkExists("enforce_end_user_domain"); ok {
		accessPointName["enforce-end-user-domain"] = v.(bool)
	}

	if v, ok := d.GetOkExists("block_traffic_other_end_user_domains"); ok {
		accessPointName["block-traffic-other-end-user-domains"] = v.(bool)
	}

	if v, ok := d.GetOkExists("block_traffic_this_end_user_domain"); ok {
		accessPointName["block-traffic-this-end-user-domain"] = v.(bool)
	}

	if v, ok := d.GetOk("end_user_domain"); ok {
		accessPointName["end-user-domain"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		accessPointName["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		accessPointName["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		accessPointName["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		accessPointName["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		accessPointName["ignore-errors"] = v.(bool)
	}

	log.Println("Create AccessPointName - Map = ", accessPointName)

	addAccessPointNameRes, err := client.ApiCall("add-access-point-name", accessPointName, client.GetSessionID(), true, false)
	if err != nil || !addAccessPointNameRes.Success {
		if addAccessPointNameRes.ErrorMsg != "" {
			return fmt.Errorf(addAccessPointNameRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addAccessPointNameRes.GetData()["uid"].(string))

	return readManagementAccessPointName(d, m)
}

func readManagementAccessPointName(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showAccessPointNameRes, err := client.ApiCall("show-access-point-name", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showAccessPointNameRes.Success {
		if objectNotFound(showAccessPointNameRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showAccessPointNameRes.ErrorMsg)
	}

	accessPointName := showAccessPointNameRes.GetData()

	log.Println("Read AccessPointName - Show JSON = ", accessPointName)

	if v := accessPointName["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := accessPointName["apn"]; v != nil {
		_ = d.Set("apn", v)
	}

	if v := accessPointName["enforce-end-user-domain"]; v != nil {
		_ = d.Set("enforce_end_user_domain", v)
	}

	if v := accessPointName["block-traffic-other-end-user-domains"]; v != nil {
		_ = d.Set("block_traffic_other_end_user_domains", v)
	}

	if v := accessPointName["block-traffic-this-end-user-domain"]; v != nil {
		_ = d.Set("block_traffic_this_end_user_domain", v)
	}

	if v := accessPointName["end-user-domain"]; v != nil {
		_ = d.Set("end_user_domain", v)
	}

	if accessPointName["tags"] != nil {
		tagsJson, ok := accessPointName["tags"].([]interface{})
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

	if v := accessPointName["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := accessPointName["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}

func updateManagementAccessPointName(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	accessPointName := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		accessPointName["name"] = oldName
		accessPointName["new-name"] = newName
	} else {
		accessPointName["name"] = d.Get("name")
	}

	if ok := d.HasChange("apn"); ok {
		accessPointName["apn"] = d.Get("apn")
	}

	if v, ok := d.GetOkExists("enforce_end_user_domain"); ok {
		accessPointName["enforce-end-user-domain"] = v.(bool)
	}

	if v, ok := d.GetOkExists("block_traffic_other_end_user_domains"); ok {
		accessPointName["block-traffic-other-end-user-domains"] = v.(bool)
	}

	if v, ok := d.GetOkExists("block_traffic_this_end_user_domain"); ok {
		accessPointName["block-traffic-this-end-user-domain"] = v.(bool)
	}

	if ok := d.HasChange("end_user_domain"); ok {
		accessPointName["end-user-domain"] = d.Get("end_user_domain")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			accessPointName["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			accessPointName["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		accessPointName["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		accessPointName["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		accessPointName["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		accessPointName["ignore-errors"] = v.(bool)
	}

	log.Println("Update AccessPointName - Map = ", accessPointName)

	updateAccessPointNameRes, err := client.ApiCall("set-access-point-name", accessPointName, client.GetSessionID(), true, false)
	if err != nil || !updateAccessPointNameRes.Success {
		if updateAccessPointNameRes.ErrorMsg != "" {
			return fmt.Errorf(updateAccessPointNameRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementAccessPointName(d, m)
}

func deleteManagementAccessPointName(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	accessPointNamePayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete AccessPointName")

	deleteAccessPointNameRes, err := client.ApiCall("delete-access-point-name", accessPointNamePayload, client.GetSessionID(), true, false)
	if err != nil || !deleteAccessPointNameRes.Success {
		if deleteAccessPointNameRes.ErrorMsg != "" {
			return fmt.Errorf(deleteAccessPointNameRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
