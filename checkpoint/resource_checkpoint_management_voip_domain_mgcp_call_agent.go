package checkpoint

import (
	"fmt"
	"log"

	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceManagementVoipDomainMgcpCallAgent() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVoipDomainMgcpCallAgent,
		Read:   readManagementVoipDomainMgcpCallAgent,
		Update: updateManagementVoipDomainMgcpCallAgent,
		Delete: deleteManagementVoipDomainMgcpCallAgent,
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

func createManagementVoipDomainMgcpCallAgent(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	voipDomainMgcpCallAgent := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		voipDomainMgcpCallAgent["name"] = v.(string)
	}

	if v, ok := d.GetOk("endpoints_domain"); ok {
		voipDomainMgcpCallAgent["endpoints-domain"] = v.(string)
	}

	if v, ok := d.GetOk("installed_at"); ok {
		voipDomainMgcpCallAgent["installed-at"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		voipDomainMgcpCallAgent["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		voipDomainMgcpCallAgent["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		voipDomainMgcpCallAgent["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		voipDomainMgcpCallAgent["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		voipDomainMgcpCallAgent["ignore-errors"] = v.(bool)
	}

	log.Println("Create VoipDomainMgcpCallAgent - Map = ", voipDomainMgcpCallAgent)

	addVoipDomainMgcpCallAgentRes, err := client.ApiCall("add-voip-domain-mgcp-call-agent", voipDomainMgcpCallAgent, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addVoipDomainMgcpCallAgentRes.Success {
		if addVoipDomainMgcpCallAgentRes.ErrorMsg != "" {
			return fmt.Errorf(addVoipDomainMgcpCallAgentRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addVoipDomainMgcpCallAgentRes.GetData()["uid"].(string))

	return readManagementVoipDomainMgcpCallAgent(d, m)
}

func readManagementVoipDomainMgcpCallAgent(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showVoipDomainMgcpCallAgentRes, err := client.ApiCall("show-voip-domain-mgcp-call-agent", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

	if v := voipDomainMgcpCallAgent["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := voipDomainMgcpCallAgent["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementVoipDomainMgcpCallAgent(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	voipDomainMgcpCallAgent := make(map[string]interface{})

	voipDomainMgcpCallAgent["uid"] = d.Id()

	if ok := d.HasChange("name"); ok {
		if v, ok := d.GetOk("name"); ok {
			voipDomainMgcpCallAgent["new-name"] = v.(string)
		}
	}

	if ok := d.HasChange("endpoints_domain"); ok {
		if v, ok := d.GetOk("endpoints_domain"); ok {
			voipDomainMgcpCallAgent["endpoints-domain"] = v.(string)
		}
	}

	if ok := d.HasChange("installed_at"); ok {
		if v, ok := d.GetOk("installed_at"); ok {
			voipDomainMgcpCallAgent["installed-at"] = v.(string)
		}
	}

	if ok := d.HasChange("color"); ok {
		if v, ok := d.GetOk("color"); ok {
			voipDomainMgcpCallAgent["color"] = v.(string)
		}
	}

	if ok := d.HasChange("comments"); ok {
		if v, ok := d.GetOk("comments"); ok {
			voipDomainMgcpCallAgent["comments"] = v.(string)
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			voipDomainMgcpCallAgent["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		voipDomainMgcpCallAgent["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		voipDomainMgcpCallAgent["ignore-errors"] = v.(bool)
	}

	log.Println("Update VoipDomainMgcpCallAgent - Map = ", voipDomainMgcpCallAgent)

	updateVoipDomainMgcpCallAgentRes, err := client.ApiCall("set-voip-domain-mgcp-call-agent", voipDomainMgcpCallAgent, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateVoipDomainMgcpCallAgentRes.Success {
		if updateVoipDomainMgcpCallAgentRes.ErrorMsg != "" {
			return fmt.Errorf(updateVoipDomainMgcpCallAgentRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementVoipDomainMgcpCallAgent(d, m)
}

func deleteManagementVoipDomainMgcpCallAgent(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	voipDomainMgcpCallAgentPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete VoipDomainMgcpCallAgent")

	deleteVoipDomainMgcpCallAgentRes, err := client.ApiCall("delete-voip-domain-mgcp-call-agent", voipDomainMgcpCallAgentPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteVoipDomainMgcpCallAgentRes.Success {
		if deleteVoipDomainMgcpCallAgentRes.ErrorMsg != "" {
			return fmt.Errorf(deleteVoipDomainMgcpCallAgentRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
