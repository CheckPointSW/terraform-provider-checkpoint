package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
)

func dataSourceManagementVpnCommunityMeshed() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementVpnCommunityMeshedRead,
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
			"gateways": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Gateway objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
							Type:        schema.TypeInt,
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
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
						},
						"ike_p2_pfs_dh_grp": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The Diffie-Hellman group to be used.",
						},
						"ike_p2_rekey_time": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Indicates the time interval for IKE phase 2 renegotiation.",
						},
					},
				},
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
							Computed: true,
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
										Type:        schema.TypeInt,
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
										Default:     "sha1",
									},
									"encryption_algorithm": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The encryption algorithm to be used.",
									},
									"ike_p2_use_pfs": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
									},
									"ike_p2_pfs_dh_grp": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The Diffie-Hellman group to be used.",
									},
									"ike_p2_rekey_time": {
										Type:        schema.TypeInt,
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

func dataSourceManagementVpnCommunityMeshedRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showVpnCommunityMeshedRes, err := client.ApiCall("show-vpn-community-meshed", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVpnCommunityMeshedRes.Success {
		return fmt.Errorf(showVpnCommunityMeshedRes.ErrorMsg)
	}

	vpnCommunityMeshed := showVpnCommunityMeshedRes.GetData()

	log.Println("Read VpnCommunityMeshed - Show JSON = ", vpnCommunityMeshed)

	if v := vpnCommunityMeshed["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := vpnCommunityMeshed["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := vpnCommunityMeshed["encryption-method"]; v != nil {
		_ = d.Set("encryption_method", v)
	}

	if v := vpnCommunityMeshed["encryption-suite"]; v != nil {
		_ = d.Set("encryption_suite", v)
	}

	if vpnCommunityMeshed["gateways"] != nil {
		gatewaysJson, ok := vpnCommunityMeshed["gateways"].([]interface{})
		if ok {
			gatewaysIds := make([]string, 0)
			if len(gatewaysJson) > 0 {
				for _, gateways := range gatewaysJson {
					gateways := gateways.(map[string]interface{})
					gatewaysIds = append(gatewaysIds, gateways["name"].(string))
				}
			}
			_ = d.Set("gateways", gatewaysIds)
		}
	} else {
		_ = d.Set("gateways", nil)
	}

	if vpnCommunityMeshed["ike-phase-1"] != nil {

		ikePhase1Map := vpnCommunityMeshed["ike-phase-1"].(map[string]interface{})

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
			ikePhase1MapToReturn["ike_p1_rekey_time"] = v
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

	if vpnCommunityMeshed["ike-phase-2"] != nil {

		ikePhase2Map := vpnCommunityMeshed["ike-phase-2"].(map[string]interface{})

		ikePhase2MapToReturn := make(map[string]interface{})

		if v, _ := ikePhase2Map["data-integrity"]; v != nil {
			ikePhase2MapToReturn["data_integrity"] = v
		}
		if v, _ := ikePhase2Map["encryption-algorithm"]; v != nil {
			ikePhase2MapToReturn["encryption_algorithm"] = v
		}
		if v := ikePhase2Map["ike-p2-use-pfs"]; v != nil {
			ikePhase2MapToReturn["ike_p2_use_pfs"] = v
		}
		if v := ikePhase2Map["ike-p2-pfs-dh-grp"]; v != nil {
			ikePhase2MapToReturn["ike_p2_pfs_dh_grp"] = v
		}
		if v := ikePhase2Map["ike-p2-rekey-time"]; v != nil {
			ikePhase2MapToReturn["ike_p2_rekey_time"] = v
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

	if vpnCommunityMeshed["override-vpn-domains"] != nil {

		overrideVpnDomainsList, ok := vpnCommunityMeshed["override-vpn-domains"].([]interface{})

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

	if vpnCommunityMeshed["shared-secrets"] != nil {

		sharedSecretsList, ok := vpnCommunityMeshed["shared-secrets"].([]interface{})

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

	if v := vpnCommunityMeshed["tunnel-granularity"]; v != nil {
		_ = d.Set("tunnel_granularity", v)
	}

	if vpnCommunityMeshed["granular-encryptions"] != nil {
		granularEncryptions, ok := vpnCommunityMeshed["granular-encryptions"].([]interface{})
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
						}else if val, ok := v.(string); ok {
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
						}else if val, ok := v.(string); ok {
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
							ikePhase1State["ike_p1_rekey_time"] = v
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
							ikePhase2State["ike_p2_use_pfs"] = v
						}
						if v := ikePhase2Show["ike-p2-pfs-dh-grp"]; v != nil {
							ikePhase2State["ike_p2_pfs_dh_grp"] = v
						}
						if v := ikePhase2Show["ike-p2-rekey-time"]; v != nil {
							ikePhase2State["ike_p2_rekey_time"] = v
						}
						granularEncryptionState["ike_phase_2"] = ikePhase2State
					}
					granularEncryptionsState = append(granularEncryptionsState, granularEncryptionState)
				}
				_ = d.Set("granular_encryptions", granularEncryptionsState)
			}else{
				_ = d.Set("granular_encryptions", nil)
			}
		}
	}

	if vpnCommunityMeshed["tags"] != nil {
		tagsJson, ok := vpnCommunityMeshed["tags"].([]interface{})
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

	if v := vpnCommunityMeshed["use-shared-secret"]; v != nil {
		_ = d.Set("use_shared_secret", v)
	}

	if v := vpnCommunityMeshed["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := vpnCommunityMeshed["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	return nil
}
