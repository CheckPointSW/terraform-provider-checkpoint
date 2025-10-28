package checkpoint

import (
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementVoipDomainH323Gatekeeper() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVoipDomainH323Gatekeeper,
		Read:   readManagementVoipDomainH323Gatekeeper,
		Update: updateManagementVoipDomainH323Gatekeeper,
		Delete: deleteManagementVoipDomainH323Gatekeeper,
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
			"routing_mode": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The routing mode of the VoIP Domain H323 gatekeeper.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"direct": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the routing mode is direct.",
							Default:     false,
						},
						"call_setup": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the routing mode includes call setup (Q.931).",
							Default:     false,
						},
						"call_setup_and_call_control": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether the routing mode includes both call setup (Q.931) and call control (H.245).",
							Default:     false,
						},
					},
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

func createManagementVoipDomainH323Gatekeeper(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	voipDomainH323Gatekeeper := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		voipDomainH323Gatekeeper["name"] = v.(string)
	}

	if v, ok := d.GetOk("endpoints_domain"); ok {
		voipDomainH323Gatekeeper["endpoints-domain"] = v.(string)
	}

	if v, ok := d.GetOk("installed_at"); ok {
		voipDomainH323Gatekeeper["installed-at"] = v.(string)
	}

	if v, ok := d.GetOk("routing_mode"); ok {

		routingModeList := v.([]interface{})

		if len(routingModeList) > 0 {

			routingModePayload := make(map[string]interface{})

			if v, ok := d.GetOk("routing_mode.0.direct"); ok {
				routingModePayload["direct"] = v.(bool)
			}
			if v, ok := d.GetOk("routing_mode.0.call_setup"); ok {
				routingModePayload["call-setup"] = v.(bool)
			}
			if v, ok := d.GetOk("routing_mode.0.call_setup_and_call_control"); ok {
				routingModePayload["call-setup-and-call-control"] = v.(bool)
			}
			voipDomainH323Gatekeeper["routing-mode"] = routingModePayload
		}
	}
	if v, ok := d.GetOk("color"); ok {
		voipDomainH323Gatekeeper["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		voipDomainH323Gatekeeper["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		voipDomainH323Gatekeeper["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		voipDomainH323Gatekeeper["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		voipDomainH323Gatekeeper["ignore-errors"] = v.(bool)
	}

	log.Println("Create VoipDomainH323Gatekeeper - Map = ", voipDomainH323Gatekeeper)

	addVoipDomainH323GatekeeperRes, err := client.ApiCall("add-voip-domain-h323-gatekeeper", voipDomainH323Gatekeeper, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addVoipDomainH323GatekeeperRes.Success {
		if addVoipDomainH323GatekeeperRes.ErrorMsg != "" {
			return fmt.Errorf(addVoipDomainH323GatekeeperRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addVoipDomainH323GatekeeperRes.GetData()["uid"].(string))

	return readManagementVoipDomainH323Gatekeeper(d, m)
}

func readManagementVoipDomainH323Gatekeeper(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showVoipDomainH323GatekeeperRes, err := client.ApiCall("show-voip-domain-h323-gatekeeper", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVoipDomainH323GatekeeperRes.Success {
		if objectNotFound(showVoipDomainH323GatekeeperRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showVoipDomainH323GatekeeperRes.ErrorMsg)
	}

	voipDomainH323Gatekeeper := showVoipDomainH323GatekeeperRes.GetData()

	log.Println("Read VoipDomainH323Gatekeeper - Show JSON = ", voipDomainH323Gatekeeper)

	if v := voipDomainH323Gatekeeper["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := voipDomainH323Gatekeeper["endpoints-domain"]; v != nil {
		_ = d.Set("endpoints_domain", v.(map[string]interface{})["name"].(string))
	}

	if v := voipDomainH323Gatekeeper["installed-at"]; v != nil {
		_ = d.Set("installed_at", v.(map[string]interface{})["name"].(string))
	}

	if voipDomainH323Gatekeeper["routing-mode"] != nil {

		routingModeMap := voipDomainH323Gatekeeper["routing-mode"].(map[string]interface{})

		routingModeMapToReturn := make(map[string]interface{})

		if v := routingModeMap["direct"]; v != nil {
			routingModeMapToReturn["direct"] = v
		}
		if v := routingModeMap["call-setup"]; v != nil {
			routingModeMapToReturn["call_setup"] = v
		}
		if v := routingModeMap["call-setup-and-call-control"]; v != nil {
			routingModeMapToReturn["call_setup_and_call_control"] = v
		}
		_ = d.Set("routing_mode", []interface{}{routingModeMapToReturn})

	} else {
		_ = d.Set("routing_mode", nil)
	}

	if v := voipDomainH323Gatekeeper["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := voipDomainH323Gatekeeper["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if voipDomainH323Gatekeeper["tags"] != nil {
		tagsJson, ok := voipDomainH323Gatekeeper["tags"].([]interface{})
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

	if v := voipDomainH323Gatekeeper["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := voipDomainH323Gatekeeper["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementVoipDomainH323Gatekeeper(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	voipDomainH323Gatekeeper := make(map[string]interface{})

	voipDomainH323Gatekeeper["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			voipDomainH323Gatekeeper["new-name"] = v.(string)
		}
	}

	if ok := d.HasChange("endpoints_domain"); ok {
		if v, ok := d.GetOk("endpoints_domain"); ok {
			voipDomainH323Gatekeeper["endpoints-domain"] = v.(string)
		}
	}

	if ok := d.HasChange("installed_at"); ok {
		if v, ok := d.GetOk("installed_at"); ok {
			voipDomainH323Gatekeeper["installed-at"] = v.(string)
		}
	}

	if d.HasChange("routing_mode") {

		if v, ok := d.GetOk("routing_mode"); ok {

			routingModeList := v.([]interface{})

			if len(routingModeList) > 0 {

				routingModePayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("routing_mode.0.direct"); ok {
					routingModePayload["direct"] = v.(bool)
				}
				if v, ok := d.GetOkExists("routing_mode.0.call_setup"); ok {
					routingModePayload["call-setup"] = v.(bool)
				}
				if v, ok := d.GetOkExists("routing_mode.0.call_setup_and_call_control"); ok {
					routingModePayload["call-setup-and-call-control"] = v.(bool)
				}
				voipDomainH323Gatekeeper["routing-mode"] = routingModePayload
			}
		}
	}

	if ok := d.HasChange("color"); ok {
		if v, ok := d.GetOk("color"); ok {
			voipDomainH323Gatekeeper["color"] = v.(string)
		}
	}

	if ok := d.HasChange("comments"); ok {
		if v, ok := d.GetOk("comments"); ok {
			voipDomainH323Gatekeeper["comments"] = v.(string)
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			voipDomainH323Gatekeeper["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		voipDomainH323Gatekeeper["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		voipDomainH323Gatekeeper["ignore-errors"] = v.(bool)
	}

	log.Println("Update VoipDomainH323Gatekeeper - Map = ", voipDomainH323Gatekeeper)

	updateVoipDomainH323GatekeeperRes, err := client.ApiCall("set-voip-domain-h323-gatekeeper", voipDomainH323Gatekeeper, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateVoipDomainH323GatekeeperRes.Success {
		if updateVoipDomainH323GatekeeperRes.ErrorMsg != "" {
			return fmt.Errorf(updateVoipDomainH323GatekeeperRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementVoipDomainH323Gatekeeper(d, m)
}

func deleteManagementVoipDomainH323Gatekeeper(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	voipDomainH323GatekeeperPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete VoipDomainH323Gatekeeper")

	deleteVoipDomainH323GatekeeperRes, err := client.ApiCall("delete-voip-domain-h323-gatekeeper", voipDomainH323GatekeeperPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteVoipDomainH323GatekeeperRes.Success {
		if deleteVoipDomainH323GatekeeperRes.ErrorMsg != "" {
			return fmt.Errorf(deleteVoipDomainH323GatekeeperRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
