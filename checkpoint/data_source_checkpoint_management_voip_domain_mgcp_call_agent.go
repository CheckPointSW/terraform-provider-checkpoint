package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
)

func dataSourceManagementVoipDomainMgcpCallAgent() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementVoipDomainMgcpCallAgentRead,
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

func dataSourceManagementVoipDomainMgcpCallAgentRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	} else if v, ok := d.GetOk("uid"); ok {
		payload["uid"] = v.(string)
	} else {
		return fmt.Errorf("Either name or uid must be specified")
	}

	showVoipDomainMgcpCallAgentRes, err := client.ApiCall("show-voip-domain-mgcp-call-agent", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVoipDomainMgcpCallAgentRes.Success {
		if objectNotFound(showVoipDomainMgcpCallAgentRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showVoipDomainMgcpCallAgentRes.ErrorMsg)
	}

	voipDomainMgcpCallAgent := showVoipDomainMgcpCallAgentRes.GetData()

	log.Println("Read VoipDomainMgcpCallAgent - Show JSON = ", voipDomainMgcpCallAgent)

	if v := voipDomainMgcpCallAgent["uid"]; v != nil {
		d.SetId(v.(string))
	}

	if v := voipDomainMgcpCallAgent["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := voipDomainMgcpCallAgent["endpoints-domain"]; v != nil {
		_ = d.Set("endpoints_domain", v.(map[string]interface{})["name"].(string))
	}

	if v := voipDomainMgcpCallAgent["installed-at"]; v != nil {
		_ = d.Set("installed_at", v.(map[string]interface{})["name"].(string))
	}

	if v := voipDomainMgcpCallAgent["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := voipDomainMgcpCallAgent["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if voipDomainMgcpCallAgent["tags"] != nil {
		tagsJson, ok := voipDomainMgcpCallAgent["tags"].([]interface{})
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

	if v := voipDomainMgcpCallAgent["icon"]; v != nil {
		_ = d.Set("icon", v)
	}

	return nil

}
