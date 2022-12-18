package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
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
						"ike_p1_rekey_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates the time interval for IKE phase 1 renegotiation.",
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
						"ike_p2_use_pfs": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
						},
						"ike_p2_pfs_dh_grp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Diffie-Hellman group to be used.",
						},
						"ike_p2_rekey_time": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Indicates the time interval for IKE phase 2 renegotiation.",
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
							Sensitive:   true,
							Description: "Shared secret.",
						},
					},
				},
			},
			"tunnel_granularity": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "VPN tunnel sharing option to be used.",
			},
			"granular_encryptions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "VPN granular encryption settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"internal_gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Internally managed Check Point gateway identified by name or UID, or 'Any' for all internal-gateways participants in this community.",
						},
						"external_gateway": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Externally managed or 3rd party gateway identified by name or UID.",
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
									"ike_p1_rekey_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates the time interval for IKE phase 1 renegotiation.",
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
									"ike_p2_use_pfs": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
									},
									"ike_p2_pfs_dh_grp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Diffie-Hellman group to be used.",
									},
									"ike_p2_rekey_time": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Indicates the time interval for IKE phase 2 renegotiation.",
									},
								},
							},
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

	showVpnCommunityStarRes, err := client.ApiCall("show-vpn-community-star", payload, client.GetSessionID(), true, client.IsProxyUsed())
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

	if vpnCommunityStar["center-gateways"] != nil {
		centerGatewaysJson, ok := vpnCommunityStar["center-gateways"].([]interface{})
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
		if v := ikePhase1Map["ike-p1-rekey-time"]; v != nil {
			ikePhase1MapToReturn["ike_p1_rekey_time"] = strconv.Itoa(int(v.(float64)))
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
		if v := ikePhase2Map["ike-p2-use-pfs"]; v != nil {
			ikePhase2MapToReturn["ike_p2_use_pfs"] = strconv.FormatBool(v.(bool))
		}
		if v := ikePhase2Map["ike-p2-pfs-dh-grp"]; v != nil {
			ikePhase2MapToReturn["ike_p2_pfs_dh_grp"] = v
		}
		if v := ikePhase2Map["ike-p2-rekey-time"]; v != nil {
			ikePhase2MapToReturn["ike_p2_rekey_time"] = strconv.Itoa(int(v.(float64)))
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
		overrideVpnDomainsList := vpnCommunityStar["override-vpn-domains"].([]interface{})
		var overrideVpnDomainsListToReturn []map[string]interface{}
		if len(overrideVpnDomainsList) > 0 {
			for i := range overrideVpnDomainsList {

				overrideVpnDomainsMap := overrideVpnDomainsList[i].(map[string]interface{})

				overrideVpnDomainsMapToAdd := make(map[string]interface{})

				if v, _ := overrideVpnDomainsMap["gateway"]; v != nil {
					overrideVpnDomainsMapToAdd["gateway"] = v.(map[string]interface{})["name"].(string)
				}
				if v, _ := overrideVpnDomainsMap["vpn-domain"]; v != nil {
					overrideVpnDomainsMapToAdd["vpn_domain"] = v.(map[string]interface{})["name"].(string)
				}
				overrideVpnDomainsListToReturn = append(overrideVpnDomainsListToReturn, overrideVpnDomainsMapToAdd)
			}
		}
		_ = d.Set("override_vpn_domains", overrideVpnDomainsListToReturn)
	} else {
		_ = d.Set("override_vpn_domains", nil)
	}

	if vpnCommunityStar["satellite-gateways"] != nil {
		satelliteGatewaysJson, ok := vpnCommunityStar["satellite-gateways"].([]interface{})
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
		sharedSecretsList := vpnCommunityStar["shared-secrets"].([]interface{})
		var sharedSecretsListToReturn []map[string]interface{}
		if len(sharedSecretsList) > 0 {
			for i := range sharedSecretsList {
				sharedSecretsMap := sharedSecretsList[i].(map[string]interface{})
				externalGateway := ""
				sharedSecret := "N/A"
				if v, _ := sharedSecretsMap["external-gateway"]; v != nil {
					externalGateway = v.(map[string]interface{})["name"].(string)
					if val, ok := d.GetOk("shared_secrets"); ok {
						sharedSecretsList := val.([]interface{})
						if len(sharedSecretsList) > 0 {
							for i := range sharedSecretsList {
								if v, ok := d.GetOk("shared_secrets." + strconv.Itoa(i) + ".external_gateway"); ok {
									if externalGateway == v.(string) {
										sharedSecret = d.Get("shared_secrets." + strconv.Itoa(i) + ".shared_secret").(string)
										break
									}
								}
							}
						}
					}
				}
				if externalGateway != "" {
					sharedSecretsMapToAdd := make(map[string]interface{})
					sharedSecretsMapToAdd["external_gateway"] = externalGateway
					sharedSecretsMapToAdd["shared_secret"] = sharedSecret
					sharedSecretsListToReturn = append(sharedSecretsListToReturn, sharedSecretsMapToAdd)
				}
			}
		}
		_ = d.Set("shared_secrets", sharedSecretsListToReturn)
	} else {
		_ = d.Set("shared_secrets", nil)
	}

	if v := vpnCommunityStar["tunnel-granularity"]; v != nil {
		_ = d.Set("tunnel_granularity", v)
	}

	if vpnCommunityStar["granular-encryptions"] != nil {
		granularEncryptions, ok := vpnCommunityStar["granular-encryptions"].([]interface{})
		if ok {
			if len(granularEncryptions) > 0 {
				var granularEncryptionsState []map[string]interface{}
				for i := range granularEncryptions {
					granularEncryptionShow := granularEncryptions[i].(map[string]interface{})
					granularEncryptionState := make(map[string]interface{})
					if granularEncryptionShow["internal-gateway"] != nil {
						var internalGatewayName string
						v := granularEncryptionShow["internal-gateway"]
						if obj, ok := v.(map[string]interface{}); ok {
							if obj["name"] != nil {
								internalGatewayName = obj["name"].(string)
							}
						} else if val, ok := v.(string); ok {
							internalGatewayName = val
						}
						granularEncryptionState["internal_gateway"] = internalGatewayName
					}

					if granularEncryptionShow["external-gateway"] != nil {
						var externalGatewayName string
						v := granularEncryptionShow["external-gateway"]
						if obj, ok := v.(map[string]interface{}); ok {
							if obj["name"] != nil {
								externalGatewayName = obj["name"].(string)
							}
						} else if val, ok := v.(string); ok {
							externalGatewayName = val
						}
						granularEncryptionState["external_gateway"] = externalGatewayName
					}

					if v := granularEncryptionShow["encryption-method"]; v != nil {
						granularEncryptionState["encryption_method"] = v
					}

					if v := granularEncryptionShow["encryption-suite"]; v != nil {
						granularEncryptionState["encryption_suite"] = v
					}

					if v := granularEncryptionShow["ike-phase-1"]; v != nil {
						ikePhase1Show := v.(map[string]interface{})
						ikePhase1State := make(map[string]interface{})
						if v := ikePhase1Show["encryption-algorithm"]; v != nil {
							ikePhase1State["encryption_algorithm"] = v
						}
						if v := ikePhase1Show["data-integrity"]; v != nil {
							ikePhase1State["data_integrity"] = v
						}
						if v := ikePhase1Show["diffie-hellman-group"]; v != nil {
							ikePhase1State["diffie_hellman_group"] = v
						}
						if v := ikePhase1Show["ike-p1-rekey-time"]; v != nil {
							ikePhase1State["ike_p1_rekey_time"] = strconv.Itoa(int(v.(float64)))
						}
						granularEncryptionState["ike_phase_1"] = ikePhase1State
					}

					if v := granularEncryptionShow["ike-phase-2"]; v != nil {
						ikePhase2Show := v.(map[string]interface{})
						ikePhase2State := make(map[string]interface{})
						if v := ikePhase2Show["encryption-algorithm"]; v != nil {
							ikePhase2State["encryption_algorithm"] = v
						}
						if v := ikePhase2Show["data-integrity"]; v != nil {
							ikePhase2State["data_integrity"] = v
						}
						if v := ikePhase2Show["ike-p2-use-pfs"]; v != nil {
							ikePhase2State["ike_p2_use_pfs"] = strconv.FormatBool(v.(bool))
						}
						if v := ikePhase2Show["ike-p2-pfs-dh-grp"]; v != nil {
							ikePhase2State["ike_p2_pfs_dh_grp"] = v
						}
						if v := ikePhase2Show["ike-p2-rekey-time"]; v != nil {
							ikePhase2State["ike_p2_rekey_time"] = strconv.Itoa(int(v.(float64)))
						}
						granularEncryptionState["ike_phase_2"] = ikePhase2State
					}
					granularEncryptionsState = append(granularEncryptionsState, granularEncryptionState)
				}
				_ = d.Set("granular_encryptions", granularEncryptionsState)
			} else {
				_ = d.Set("granular_encryptions", nil)
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
