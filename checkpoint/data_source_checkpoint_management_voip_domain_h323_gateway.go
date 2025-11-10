package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementVoipDomainH323Gateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementVoipDomainH323GatewayRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"endpoints_domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The related endpoints domain to which the VoIP domain will connect.  Identified by name or UID.",
			},
			"installed_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The machine the VoIP is installed at.  Identified by name or UID.",
			},
			"routing_mode": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The routing mode of the VoIP Domain H323 gateway.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"call_setup": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the routing mode includes call setup (Q.931).",
						},
						"call_setup_and_call_control": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether the routing mode includes both call setup (Q.931) and call control (H.245).",
						},
					},
				},
			},
			"color": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Color of the object. Should be one of existing colors.",
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"icon": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object icon.",
			},
		},
	}
}

func dataSourceManagementVoipDomainH323GatewayRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	showVoipDomainH323GatewayRes, err := client.ApiCall("show-voip-domain-h323-gateway", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVoipDomainH323GatewayRes.Success {
		if objectNotFound(showVoipDomainH323GatewayRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showVoipDomainH323GatewayRes.ErrorMsg)
	}

	voipDomainH323Gateway := showVoipDomainH323GatewayRes.GetData()

	log.Println("Read VoipDomainH323Gateway - Show JSON = ", voipDomainH323Gateway)

	if v := voipDomainH323Gateway["uid"]; v != nil {
		d.SetId(v.(string))
	}

	if v := voipDomainH323Gateway["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := voipDomainH323Gateway["endpoints-domain"]; v != nil {
		_ = d.Set("endpoints_domain", v.(map[string]interface{})["name"].(string))
	}

	if v := voipDomainH323Gateway["installed-at"]; v != nil {
		_ = d.Set("installed_at", v.(map[string]interface{})["name"].(string))
	}

	if voipDomainH323Gateway["routing-mode"] != nil {

		routingModeMap := voipDomainH323Gateway["routing-mode"].(map[string]interface{})

		routingModeMapToReturn := make(map[string]interface{})

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

	if v := voipDomainH323Gateway["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := voipDomainH323Gateway["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if voipDomainH323Gateway["tags"] != nil {
		tagsJson, ok := voipDomainH323Gateway["tags"].([]interface{})
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

	if v := voipDomainH323Gateway["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	return nil

}
