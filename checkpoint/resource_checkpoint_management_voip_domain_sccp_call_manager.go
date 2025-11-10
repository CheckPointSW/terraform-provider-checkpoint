package checkpoint

import (
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementVoipDomainSccpCallManager() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVoipDomainSccpCallManager,
		Read:   readManagementVoipDomainSccpCallManager,
		Update: updateManagementVoipDomainSccpCallManager,
		Delete: deleteManagementVoipDomainSccpCallManager,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"endpoints_domain": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The related endpoints domain to which the VoIP domain will connect.  Identified by name or UID.",
			},
			"installed_at": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The machine the VoIP is installed at.  Identified by name or UID.",
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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

func createManagementVoipDomainSccpCallManager(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	voipDomainSccpCallManager := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		voipDomainSccpCallManager["name"] = v.(string)
	}

	if v, ok := d.GetOk("endpoints_domain"); ok {
		voipDomainSccpCallManager["endpoints-domain"] = v.(string)
	}

	if v, ok := d.GetOk("installed_at"); ok {
		voipDomainSccpCallManager["installed-at"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		voipDomainSccpCallManager["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		voipDomainSccpCallManager["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		voipDomainSccpCallManager["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		voipDomainSccpCallManager["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		voipDomainSccpCallManager["ignore-errors"] = v.(bool)
	}

	log.Println("Create VoipDomainSccpCallManager - Map = ", voipDomainSccpCallManager)

	addVoipDomainSccpCallManagerRes, err := client.ApiCall("add-voip-domain-sccp-call-manager", voipDomainSccpCallManager, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addVoipDomainSccpCallManagerRes.Success {
		if addVoipDomainSccpCallManagerRes.ErrorMsg != "" {
			return fmt.Errorf(addVoipDomainSccpCallManagerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addVoipDomainSccpCallManagerRes.GetData()["uid"].(string))

	return readManagementVoipDomainSccpCallManager(d, m)
}

func readManagementVoipDomainSccpCallManager(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showVoipDomainSccpCallManagerRes, err := client.ApiCall("show-voip-domain-sccp-call-manager", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVoipDomainSccpCallManagerRes.Success {
		if objectNotFound(showVoipDomainSccpCallManagerRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showVoipDomainSccpCallManagerRes.ErrorMsg)
	}

	voipDomainSccpCallManager := showVoipDomainSccpCallManagerRes.GetData()

	log.Println("Read VoipDomainSccpCallManager - Show JSON = ", voipDomainSccpCallManager)

	if v := voipDomainSccpCallManager["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := voipDomainSccpCallManager["endpoints-domain"]; v != nil {
		_ = d.Set("endpoints_domain", v.(map[string]interface{})["name"].(string))
	}

	if v := voipDomainSccpCallManager["installed-at"]; v != nil {
		_ = d.Set("installed_at", v.(map[string]interface{})["name"].(string))
	}

	if v := voipDomainSccpCallManager["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := voipDomainSccpCallManager["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if voipDomainSccpCallManager["tags"] != nil {
		tagsJson, ok := voipDomainSccpCallManager["tags"].([]interface{})
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

	if v := voipDomainSccpCallManager["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := voipDomainSccpCallManager["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementVoipDomainSccpCallManager(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	voipDomainSccpCallManager := make(map[string]interface{})

	voipDomainSccpCallManager["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			voipDomainSccpCallManager["new-name"] = v.(string)
		}
	}

	if ok := d.HasChange("endpoints_domain"); ok {
		if v, ok := d.GetOk("endpoints_domain"); ok {
			voipDomainSccpCallManager["endpoints-domain"] = v.(string)
		}
	}

	if ok := d.HasChange("installed_at"); ok {
		if v, ok := d.GetOk("installed_at"); ok {
			voipDomainSccpCallManager["installed-at"] = v.(string)
		}
	}

	if ok := d.HasChange("color"); ok {
		if v, ok := d.GetOk("color"); ok {
			voipDomainSccpCallManager["color"] = v.(string)
		}
	}

	if ok := d.HasChange("comments"); ok {
		if v, ok := d.GetOk("comments"); ok {
			voipDomainSccpCallManager["comments"] = v.(string)
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			voipDomainSccpCallManager["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		voipDomainSccpCallManager["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		voipDomainSccpCallManager["ignore-errors"] = v.(bool)
	}

	log.Println("Update VoipDomainSccpCallManager - Map = ", voipDomainSccpCallManager)

	updateVoipDomainSccpCallManagerRes, err := client.ApiCall("set-voip-domain-sccp-call-manager", voipDomainSccpCallManager, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateVoipDomainSccpCallManagerRes.Success {
		if updateVoipDomainSccpCallManagerRes.ErrorMsg != "" {
			return fmt.Errorf(updateVoipDomainSccpCallManagerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementVoipDomainSccpCallManager(d, m)
}

func deleteManagementVoipDomainSccpCallManager(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	voipDomainSccpCallManagerPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete VoipDomainSccpCallManager")

	deleteVoipDomainSccpCallManagerRes, err := client.ApiCall("delete-voip-domain-sccp-call-manager", voipDomainSccpCallManagerPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteVoipDomainSccpCallManagerRes.Success {
		if deleteVoipDomainSccpCallManagerRes.ErrorMsg != "" {
			return fmt.Errorf(deleteVoipDomainSccpCallManagerRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
