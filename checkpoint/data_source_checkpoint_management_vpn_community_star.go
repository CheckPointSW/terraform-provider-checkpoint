package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
)

func dataSourceManagementVpnCommunityStar() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementVpnCommunityStarRead,
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
			"center_gateways": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Gateway objects representing center gateways identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"encryption_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encryption method to be used.",
			},
			"encryption_suite": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The encryption suite to be used.",
			},
			"ike_phase_1": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom].",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_integrity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hash algorithm to be used.",
						},
						"diffie_hellman_group": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Diffie-Hellman group to be used.",
						},
						"encryption_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The encryption algorithm to be used.",
						},
					},
				},
			},
			"ike_phase_2": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom].",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_integrity": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The hash algorithm to be used.",
						},
						"encryption_algorithm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The encryption algorithm to be used.",
						},
					},
				},
			},
			"mesh_center_gateways": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the meshed community is in center.",
			},
			"override_vpn_domains": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The Overrides VPN Domains of the participants GWs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Participant gateway in override VPN domain identified by the name or UID.",
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "VPN domain network identified by the name or UID.",
						},
					},
				},
			},
			"satellite_gateways": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Gateway objects representing satellite gateways identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"shared_secrets": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Shared secrets for external gateways.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "External gateway identified by the name or UID.",
						},
						"shared_secret": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Shared secret.",
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_shared_secret": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether the shared secret should be used for all external gateways.",
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
		},
	}
}

func dataSourceManagementVpnCommunityStarRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showVpnCommunityStarRes, err := client.ApiCall("show-vpn-community-star", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVpnCommunityStarRes.Success {
		return fmt.Errorf(showVpnCommunityStarRes.ErrorMsg)
	}

	vpnCommunityStar := showVpnCommunityStarRes.GetData()

	log.Println("Read VpnCommunityStar - Show JSON = ", vpnCommunityStar)

	if v := vpnCommunityStar["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := vpnCommunityStar["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if vpnCommunityStar["center_gateways"] != nil {
		centerGatewaysJson, ok := vpnCommunityStar["center_gateways"].([]interface{})
		if ok {
			centerGatewaysIds := make([]string, 0)
			if len(centerGatewaysJson) > 0 {
				for _, center_gateways := range centerGatewaysJson {
					center_gateways := center_gateways.(map[string]interface{})
					centerGatewaysIds = append(centerGatewaysIds, center_gateways["name"].(string))
				}
			}
			_ = d.Set("center_gateways", centerGatewaysIds)
		}
	} else {
		_ = d.Set("center_gateways", nil)
	}

	if v := vpnCommunityStar["encryption-method"]; v != nil {
		_ = d.Set("encryption_method", v)
	}

	if v := vpnCommunityStar["encryption-suite"]; v != nil {
		_ = d.Set("encryption_suite", v)
	}

	if vpnCommunityStar["ike-phase-1"] != nil {

		ikePhase1Map := vpnCommunityStar["ike-phase-1"].(map[string]interface{})

		ikePhase1MapToReturn := make(map[string]interface{})

		if v, _ := ikePhase1Map["data-integrity"]; v != nil {
			ikePhase1MapToReturn["data_integrity"] = v
		}
		if v, _ := ikePhase1Map["diffie-hellman-group"]; v != nil {
			ikePhase1MapToReturn["diffie_hellman_group"] = v
		}
		if v, _ := ikePhase1Map["encryption-algorithm"]; v != nil {
			ikePhase1MapToReturn["encryption_algorithm"] = v
		}

		_, ikePhase1InConf := d.GetOk("ike_phase_1")
		defaultIkePhase1 := map[string]interface{}{"encryption_algorithm": "aes-256", "diffie_hellman_group": "group-2", "data_integrity": "sha1"}
		if reflect.DeepEqual(defaultIkePhase1, ikePhase1MapToReturn) && !ikePhase1InConf {
			_ = d.Set("ike_phase_1", map[string]interface{}{})
		} else {
			_ = d.Set("ike_phase_1", ikePhase1MapToReturn)
		}

	} else {
		_ = d.Set("ike_phase_1", nil)
	}

	if vpnCommunityStar["ike-phase-2"] != nil {

		ikePhase2Map := vpnCommunityStar["ike-phase-2"].(map[string]interface{})

		ikePhase2MapToReturn := make(map[string]interface{})

		if v, _ := ikePhase2Map["data-integrity"]; v != nil {
			ikePhase2MapToReturn["data_integrity"] = v
		}
		if v, _ := ikePhase2Map["encryption-algorithm"]; v != nil {
			ikePhase2MapToReturn["encryption_algorithm"] = v
		}

		_, ikePhase2InConf := d.GetOk("ike_phase_2")
		defaultIkePhase2 := map[string]interface{}{"encryption_algorithm": "aes-128", "data_integrity": "sha1"}
		if reflect.DeepEqual(defaultIkePhase2, ikePhase2MapToReturn) && !ikePhase2InConf {
			_ = d.Set("ike_phase_2", map[string]interface{}{})
		} else {
			_ = d.Set("ike_phase_2", ikePhase2MapToReturn)
		}

	} else {
		_ = d.Set("ike_phase_2", nil)
	}

	if v := vpnCommunityStar["mesh-center-gateways"]; v != nil {
		_ = d.Set("mesh_center_gateways", v)
	}

	if vpnCommunityStar["override-vpn-domains"] != nil {

		overrideVpnDomainsList, ok := vpnCommunityStar["override-vpn-domains"].([]interface{})

		if ok {

			if len(overrideVpnDomainsList) > 0 {

				var overrideVpnDomainsListToReturn []map[string]interface{}

				for i := range overrideVpnDomainsList {

					overrideVpnDomainsMap := overrideVpnDomainsList[i].(map[string]interface{})

					overrideVpnDomainsMapToAdd := make(map[string]interface{})

					if v, _ := overrideVpnDomainsMap["gateway"]; v != nil {
						overrideVpnDomainsMapToAdd["gateway"] = v
					}
					if v, _ := overrideVpnDomainsMap["vpn-domain"]; v != nil {
						overrideVpnDomainsMapToAdd["vpn_domain"] = v
					}
					overrideVpnDomainsListToReturn = append(overrideVpnDomainsListToReturn, overrideVpnDomainsMapToAdd)
				}
			}
		}
	}

	if vpnCommunityStar["satellite_gateways"] != nil {
		satelliteGatewaysJson, ok := vpnCommunityStar["satellite_gateways"].([]interface{})
		if ok {
			satelliteGatewaysIds := make([]string, 0)
			if len(satelliteGatewaysJson) > 0 {
				for _, satellite_gateways := range satelliteGatewaysJson {
					satellite_gateways := satellite_gateways.(map[string]interface{})
					satelliteGatewaysIds = append(satelliteGatewaysIds, satellite_gateways["name"].(string))
				}
			}
			_ = d.Set("satellite_gateways", satelliteGatewaysIds)
		}
	} else {
		_ = d.Set("satellite_gateways", nil)
	}

	if vpnCommunityStar["shared-secrets"] != nil {

		sharedSecretsList, ok := vpnCommunityStar["shared-secrets"].([]interface{})

		if ok {

			if len(sharedSecretsList) > 0 {

				var sharedSecretsListToReturn []map[string]interface{}

				for i := range sharedSecretsList {

					sharedSecretsMap := sharedSecretsList[i].(map[string]interface{})

					sharedSecretsMapToAdd := make(map[string]interface{})

					if v, _ := sharedSecretsMap["external-gateway"]; v != nil {
						sharedSecretsMapToAdd["external_gateway"] = v
					}
					if v, _ := sharedSecretsMap["shared-secret"]; v != nil {
						sharedSecretsMapToAdd["shared_secret"] = v
					}
					sharedSecretsListToReturn = append(sharedSecretsListToReturn, sharedSecretsMapToAdd)
				}
			}
		}
	}

	if vpnCommunityStar["tags"] != nil {
		tagsJson, ok := vpnCommunityStar["tags"].([]interface{})
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

	if v := vpnCommunityStar["use-shared-secret"]; v != nil {
		_ = d.Set("use_shared_secret", v)
	}

	if v := vpnCommunityStar["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := vpnCommunityStar["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
