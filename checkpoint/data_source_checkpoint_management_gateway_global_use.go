package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceManagementSetGatewayGlobalUse() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementSetGatewayGlobalUseRead,

		Schema: map[string]*schema.Schema{
			"target": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "On what target to execute this command. Target may be identified by its object name, or object unique identifier.",
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Object unique identifier.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target name.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether global use is enabled on the target.",
			},
			"domain": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Information about the domain that holds the Object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Must be unique in the domain.",
						},
						"uid": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object unique identifier.",
						},
						"domain_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Domain type.",
						},
					},
				},
			},
			"enable_identity_sharing": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether Identity Awareness Sharing global use is enabled on the target.",
			},
			"enable_vpn": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether VPN global use is enabled on the target.",
			},
			"identity_sharing_domains": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Domains that Identity Awareness Sharing global use applied to them.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceManagementSetGatewayGlobalUseRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("target"); ok {
		payload["target"] = v.(string)
	}

	ShowGatewayGlobalUseRes, _ := client.ApiCall("show-gateway-global-use", payload, client.GetSessionID(), true, false)
	if !ShowGatewayGlobalUseRes.Success {
		return fmt.Errorf("%s", ShowGatewayGlobalUseRes.ErrorMsg)
	}

	showGatewatGlobalUseData := ShowGatewayGlobalUseRes.GetData()

	_ = d.Set("uid", showGatewatGlobalUseData["uid"])
	d.SetId(showGatewatGlobalUseData["uid"].(string))

	if v := showGatewatGlobalUseData["name"]; v != nil {
		d.Set("name", v)
	}
	if v := showGatewatGlobalUseData["enabled"]; v != nil {
		d.Set("enabled", v)
	}
	if v := showGatewatGlobalUseData["domain"]; v != nil {

		innerMap := v.(map[string]interface{})

		mapToReturn := make(map[string]interface{})

		if v := innerMap["name"]; v != nil {
			mapToReturn["name"] = v
		}
		if v := innerMap["uid"]; v != nil {
			mapToReturn["uid"] = v
		}
		if v := innerMap["domain-type"]; v != nil {
			mapToReturn["domain_type"] = v
		}

		_ = d.Set("domain", []interface{}{mapToReturn})
	}
	if v := showGatewatGlobalUseData["enable-identity-sharing"]; v != nil {
		_ = d.Set("enable_identity_sharing", v)
	}

	if v := showGatewatGlobalUseData["enable-vpn"]; v != nil {
		_ = d.Set("enable_vpn", v)
	}

	if v := showGatewatGlobalUseData["identity-sharing-domains"]; v != nil {
		identitysharingdomainsIdsList := v.([]interface{})
		var identitysharingdomainsIds = make([]string, 0)
		if len(identitysharingdomainsIdsList) > 0 {
			for _, item := range identitysharingdomainsIdsList {
				identitysharingdomainsIds = append(identitysharingdomainsIds, item.(map[string]interface{})["name"].(string))
			}
		}
		_ = d.Set("identity_sharing_domains", identitysharingdomainsIds)
	}

	return nil
}
