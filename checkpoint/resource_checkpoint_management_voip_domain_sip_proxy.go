package checkpoint

import (
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementVoipDomainSipProxy() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVoipDomainSipProxy,
		Read:   readManagementVoipDomainSipProxy,
		Update: updateManagementVoipDomainSipProxy,
		Delete: deleteManagementVoipDomainSipProxy,
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

func createManagementVoipDomainSipProxy(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	voipDomainSipProxy := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		voipDomainSipProxy["name"] = v.(string)
	}

	if v, ok := d.GetOk("endpoints_domain"); ok {
		voipDomainSipProxy["endpoints-domain"] = v.(string)
	}

	if v, ok := d.GetOk("installed_at"); ok {
		voipDomainSipProxy["installed-at"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		voipDomainSipProxy["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		voipDomainSipProxy["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		voipDomainSipProxy["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		voipDomainSipProxy["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		voipDomainSipProxy["ignore-errors"] = v.(bool)
	}

	log.Println("Create VoipDomainSipProxy - Map = ", voipDomainSipProxy)

	addVoipDomainSipProxyRes, err := client.ApiCall("add-voip-domain-sip-proxy", voipDomainSipProxy, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addVoipDomainSipProxyRes.Success {
		if addVoipDomainSipProxyRes.ErrorMsg != "" {
			return fmt.Errorf(addVoipDomainSipProxyRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addVoipDomainSipProxyRes.GetData()["uid"].(string))

	return readManagementVoipDomainSipProxy(d, m)
}

func readManagementVoipDomainSipProxy(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showVoipDomainSipProxyRes, err := client.ApiCall("show-voip-domain-sip-proxy", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVoipDomainSipProxyRes.Success {
		if objectNotFound(showVoipDomainSipProxyRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showVoipDomainSipProxyRes.ErrorMsg)
	}

	voipDomainSipProxy := showVoipDomainSipProxyRes.GetData()

	log.Println("Read VoipDomainSipProxy - Show JSON = ", voipDomainSipProxy)

	if v := voipDomainSipProxy["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := voipDomainSipProxy["endpoints-domain"]; v != nil {
		_ = d.Set("endpoints_domain", v.(map[string]interface{})["name"].(string))
	}

	if v := voipDomainSipProxy["installed-at"]; v != nil {
		_ = d.Set("installed_at", v.(map[string]interface{})["name"].(string))
	}

	if v := voipDomainSipProxy["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := voipDomainSipProxy["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if voipDomainSipProxy["tags"] != nil {
		tagsJson, ok := voipDomainSipProxy["tags"].([]interface{})
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

	if v := voipDomainSipProxy["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := voipDomainSipProxy["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementVoipDomainSipProxy(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	voipDomainSipProxy := make(map[string]interface{})

	voipDomainSipProxy["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			voipDomainSipProxy["new-name"] = v.(string)
		}
	}

	if ok := d.HasChange("endpoints_domain"); ok {
		if v, ok := d.GetOk("endpoints_domain"); ok {
			voipDomainSipProxy["endpoints-domain"] = v.(string)
		}
	}

	if ok := d.HasChange("installed_at"); ok {
		if v, ok := d.GetOk("installed_at"); ok {
			voipDomainSipProxy["installed-at"] = v.(string)
		}
	}

	if ok := d.HasChange("color"); ok {
		if v, ok := d.GetOk("color"); ok {
			voipDomainSipProxy["color"] = v.(string)
		}
	}

	if ok := d.HasChange("comments"); ok {
		if v, ok := d.GetOk("comments"); ok {
			voipDomainSipProxy["comments"] = v.(string)
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			voipDomainSipProxy["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		voipDomainSipProxy["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		voipDomainSipProxy["ignore-errors"] = v.(bool)
	}

	log.Println("Update VoipDomainSipProxy - Map = ", voipDomainSipProxy)

	updateVoipDomainSipProxyRes, err := client.ApiCall("set-voip-domain-sip-proxy", voipDomainSipProxy, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateVoipDomainSipProxyRes.Success {
		if updateVoipDomainSipProxyRes.ErrorMsg != "" {
			return fmt.Errorf(updateVoipDomainSipProxyRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementVoipDomainSipProxy(d, m)
}

func deleteManagementVoipDomainSipProxy(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	voipDomainSipProxyPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete VoipDomainSipProxy")

	deleteVoipDomainSipProxyRes, err := client.ApiCall("delete-voip-domain-sip-proxy", voipDomainSipProxyPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteVoipDomainSipProxyRes.Success {
		if deleteVoipDomainSipProxyRes.ErrorMsg != "" {
			return fmt.Errorf(deleteVoipDomainSipProxyRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
