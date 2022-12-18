package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementVpnCommunityMeshed() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVpnCommunityMeshed,
		Read:   readManagementVpnCommunityMeshed,
		Update: updateManagementVpnCommunityMeshed,
		Delete: deleteManagementVpnCommunityMeshed,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"encryption_method": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The encryption method to be used.",
			},
			"encryption_suite": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The encryption suite to be used.",
			},
			"gateways": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Gateway objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"ike_phase_1": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom].",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_integrity": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The hash algorithm to be used.",
							Default:     "sha1",
						},
						"diffie_hellman_group": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Diffie-Hellman group to be used.",
							Default:     "group-2",
						},
						"encryption_algorithm": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The encryption algorithm to be used.",
							Default:     "aes-256",
						},
						"ike_p1_rekey_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates the time interval for IKE phase 1 renegotiation.",
							Default:     1440,
						},
					},
				},
			},
			"ike_phase_2": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom].",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"data_integrity": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The hash algorithm to be used.",
							Default:     "sha1",
						},
						"encryption_algorithm": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The encryption algorithm to be used.",
							Default:     "aes-128",
						},
						"ike_p2_use_pfs": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
							Default:     false,
						},
						"ike_p2_pfs_dh_grp": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The Diffie-Hellman group to be used.",
							Default:     "group-2",
						},
						"ike_p2_rekey_time": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Indicates the time interval for IKE phase 2 renegotiation.",
							Default:     1440,
						},
					},
				},
			},
			"override_vpn_domains": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The Overrides VPN Domains of the participants GWs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Participant gateway in override VPN domain identified by the name or UID.",
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "VPN domain network identified by the name or UID.",
						},
					},
				},
			},
			"shared_secrets": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Shared secrets for external gateways.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"external_gateway": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "External gateway identified by the name or UID.",
						},
						"shared_secret": {
							Type:        schema.TypeString,
							Optional:    true,
							Sensitive:   true,
							Description: "Shared secret.",
						},
					},
				},
			},
			"tunnel_granularity": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "VPN tunnel sharing option to be used.",
			},
			"granular_encryptions": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "VPN granular encryption settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"internal_gateway": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Internally managed Check Point gateway identified by name or UID, or 'Any' for all internal-gateways participants in this community.",
						},
						"external_gateway": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Externally managed or 3rd party gateway identified by name or UID.",
						},
						"encryption_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The encryption method to be used.",
						},
						"encryption_suite": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The encryption suite to be used.",
						},
						"ike_phase_1": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Ike Phase 1 settings. Only applicable when the encryption-suite is set to [custom].",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_integrity": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The hash algorithm to be used.",
										Default:     "sha1",
									},
									"diffie_hellman_group": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Diffie-Hellman group to be used.",
										Default:     "group-2",
									},
									"encryption_algorithm": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The encryption algorithm to be used.",
										Default:     "aes-256",
									},
									"ike_p1_rekey_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates the time interval for IKE phase 1 renegotiation.",
										Default:     1440,
									},
								},
							},
						},
						"ike_phase_2": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Ike Phase 2 settings. Only applicable when the encryption-suite is set to [custom].",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data_integrity": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The hash algorithm to be used.",
										Default:     "sha1",
									},
									"encryption_algorithm": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The encryption algorithm to be used.",
										Default:     "aes-128",
									},
									"ike_p2_use_pfs": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
										Default:     false,
									},
									"ike_p2_pfs_dh_grp": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Diffie-Hellman group to be used.",
										Default:     "group-2",
									},
									"ike_p2_rekey_time": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates the time interval for IKE phase 2 renegotiation.",
										Default:     1440,
									},
								},
							},
						},
					},
				},
			},
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"use_shared_secret": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether the shared secret should be used for all external gateways.",
				Default:     false,
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

func createManagementVpnCommunityMeshed(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	vpnCommunityMeshed := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		vpnCommunityMeshed["name"] = v.(string)
	}

	if v, ok := d.GetOk("encryption_method"); ok {
		vpnCommunityMeshed["encryption-method"] = v.(string)
	}

	if v, ok := d.GetOk("encryption_suite"); ok {
		vpnCommunityMeshed["encryption-suite"] = v.(string)
	}

	if v, ok := d.GetOk("gateways"); ok {
		vpnCommunityMeshed["gateways"] = v.(*schema.Set).List()
	}

	if _, ok := d.GetOk("ike_phase_1"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("ike_phase_1.data_integrity"); ok {
			res["data-integrity"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_1.diffie_hellman_group"); ok {
			res["diffie-hellman-group"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_1.encryption_algorithm"); ok {
			res["encryption-algorithm"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_1.ike_p1_rekey_time"); ok {
			res["ike-p1-rekey-time"] = v.(string)
		}
		vpnCommunityMeshed["ike-phase-1"] = res
	}

	if _, ok := d.GetOk("ike_phase_2"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("ike_phase_2.data_integrity"); ok {
			res["data-integrity"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_2.encryption_algorithm"); ok {
			res["encryption-algorithm"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_2.ike_p2_use_pfs"); ok {
			res["ike-p2-use-pfs"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_2.ike_p2_pfs_dh_grp"); ok {
			res["ike-p2-pfs-dh-grp"] = v.(string)
		}
		if v, ok := d.GetOk("ike_phase_2.ike_p2_rekey_time"); ok {
			res["ike-p2-rekey-time"] = v.(string)
		}
		vpnCommunityMeshed["ike-phase-2"] = res
	}

	if v, ok := d.GetOk("override_vpn_domains"); ok {

		overrideVpnDomainsList := v.([]interface{})

		if len(overrideVpnDomainsList) > 0 {

			var overrideVpnDomainsPayload []map[string]interface{}

			for i := range overrideVpnDomainsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("override_vpn_domains." + strconv.Itoa(i) + ".gateway"); ok {
					Payload["gateway"] = v.(string)
				}
				if v, ok := d.GetOk("override_vpn_domains." + strconv.Itoa(i) + ".vpn_domain"); ok {
					Payload["vpn-domain"] = v.(string)
				}
				overrideVpnDomainsPayload = append(overrideVpnDomainsPayload, Payload)
			}
			vpnCommunityMeshed["overrideVpnDomains"] = overrideVpnDomainsPayload
		}
	}

	if v, ok := d.GetOk("shared_secrets"); ok {

		sharedSecretsList := v.([]interface{})

		if len(sharedSecretsList) > 0 {

			var sharedSecretsPayload []map[string]interface{}

			for i := range sharedSecretsList {

				Payload := make(map[string]interface{})

				if v, ok := d.GetOk("shared_secrets." + strconv.Itoa(i) + ".external_gateway"); ok {
					Payload["external-gateway"] = v.(string)
				}
				if v, ok := d.GetOk("shared_secrets." + strconv.Itoa(i) + ".shared_secret"); ok {
					Payload["shared-secret"] = v.(string)
				}
				sharedSecretsPayload = append(sharedSecretsPayload, Payload)
			}
			vpnCommunityMeshed["shared-secrets"] = sharedSecretsPayload
		}
	}

	if v, ok := d.GetOk("tunnel_granularity"); ok {
		vpnCommunityMeshed["tunnel-granularity"] = v.(string)
	}

	if v, ok := d.GetOk("granular_encryptions"); ok {
		granularEncryptions := v.([]interface{})
		if len(granularEncryptions) > 0 {
			var granularEncryptionsPayload []map[string]interface{}

			for i := range granularEncryptions {
				payload := make(map[string]interface{})
				if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".internal_gateway"); ok {
					payload["internal-gateway"] = v.(string)
				}
				if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".external_gateway"); ok {
					payload["external-gateway"] = v.(string)
				}
				if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".encryption_method"); ok {
					payload["encryption-method"] = v.(string)
				}
				if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".encryption_suite"); ok {
					payload["encryption-suite"] = v.(string)
				}
				if _, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1"); ok {
					ikePhase1Payload := make(map[string]interface{})
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.encryption_algorithm"); ok {
						ikePhase1Payload["encryption-algorithm"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.data_integrity"); ok {
						ikePhase1Payload["data-integrity"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.diffie_hellman_group"); ok {
						ikePhase1Payload["diffie-hellman-group"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.ike_p1_rekey_time"); ok {
						ikePhase1Payload["ike-p1-rekey-time"] = v.(string)
					}
					payload["ike-phase-1"] = ikePhase1Payload
				}
				if _, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2"); ok {
					ikePhase2Payload := make(map[string]interface{})
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.encryption_algorithm"); ok {
						ikePhase2Payload["encryption-algorithm"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.data_integrity"); ok {
						ikePhase2Payload["data-integrity"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_use_pfs"); ok {
						ikePhase2Payload["ike-p2-use-pfs"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_pfs_dh_grp"); ok {
						ikePhase2Payload["ike-p2-pfs-dh-grp"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_rekey_time"); ok {
						ikePhase2Payload["ike-p2-rekey-time"] = v.(string)
					}
					payload["ike-phase-2"] = ikePhase2Payload
				}
				granularEncryptionsPayload = append(granularEncryptionsPayload, payload)
			}
			vpnCommunityMeshed["granular-encryptions"] = granularEncryptionsPayload
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		vpnCommunityMeshed["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("use_shared_secret"); ok {
		vpnCommunityMeshed["use-shared-secret"] = v.(bool)
	}

	if v, ok := d.GetOk("color"); ok {
		vpnCommunityMeshed["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		vpnCommunityMeshed["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		vpnCommunityMeshed["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		vpnCommunityMeshed["ignore-errors"] = v.(bool)
	}

	log.Println("Create VpnCommunityMeshed - Map = ", vpnCommunityMeshed)

	addVpnCommunityMeshedRes, err := client.ApiCall("add-vpn-community-meshed", vpnCommunityMeshed, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !addVpnCommunityMeshedRes.Success {
		if addVpnCommunityMeshedRes.ErrorMsg != "" {
			return fmt.Errorf(addVpnCommunityMeshedRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addVpnCommunityMeshedRes.GetData()["uid"].(string))

	return readManagementVpnCommunityMeshed(d, m)
}

func readManagementVpnCommunityMeshed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showVpnCommunityMeshedRes, err := client.ApiCall("show-vpn-community-meshed", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showVpnCommunityMeshedRes.Success {
		if objectNotFound(showVpnCommunityMeshedRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showVpnCommunityMeshedRes.ErrorMsg)
	}

	vpnCommunityMeshed := showVpnCommunityMeshedRes.GetData()

	log.Println("Read VpnCommunityMeshed - Show JSON = ", vpnCommunityMeshed)

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

	if vpnCommunityMeshed["override-vpn-domains"] != nil {
		overrideVpnDomainsList := vpnCommunityMeshed["override-vpn-domains"].([]interface{})
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

	if vpnCommunityMeshed["shared-secrets"] != nil {
		sharedSecretsList := vpnCommunityMeshed["shared-secrets"].([]interface{})
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

	if v := vpnCommunityMeshed["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := vpnCommunityMeshed["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementVpnCommunityMeshed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	vpnCommunityMeshed := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		vpnCommunityMeshed["name"] = oldName
		vpnCommunityMeshed["new-name"] = newName
	} else {
		vpnCommunityMeshed["name"] = d.Get("name")
	}

	if ok := d.HasChange("encryption_method"); ok {
		vpnCommunityMeshed["encryption-method"] = d.Get("encryption_method")
	}

	if ok := d.HasChange("encryption_suite"); ok {
		vpnCommunityMeshed["encryption-suite"] = d.Get("encryption_suite")
	}

	if d.HasChange("gateways") {
		if v, ok := d.GetOk("gateways"); ok {
			vpnCommunityMeshed["gateways"] = v.(*schema.Set).List()
		} else {
			oldGateways, _ := d.GetChange("gateways")
			vpnCommunityMeshed["gateways"] = map[string]interface{}{"remove": oldGateways.(*schema.Set).List()}
		}
	}

	if d.HasChange("ike_phase_1") {

		if _, ok := d.GetOk("ike_phase_1"); ok {

			res := make(map[string]interface{})

			if d.HasChange("ike_phase_1.data_integrity") {
				res["data-integrity"] = d.Get("ike_phase_1.data_integrity")
			}
			if d.HasChange("ike_phase_1.diffie_hellman_group") {
				res["diffie-hellman-group"] = d.Get("ike_phase_1.diffie_hellman_group")
			}
			if d.HasChange("ike_phase_1.encryption_algorithm") {
				res["encryption-algorithm"] = d.Get("ike_phase_1.encryption_algorithm")
			}
			if d.HasChange("ike_phase_1.ike_p1_rekey_time") {
				res["ike-p1-rekey-time"] = d.Get("ike_phase_1.ike_p1_rekey_time")
			}
			vpnCommunityMeshed["ike-phase-1"] = res
		} else {
			vpnCommunityMeshed["ike-phase-1"] = map[string]interface{}{"encryption-algorithm": "aes-256", "diffie-hellman-group": "group-2", "data-integrity": "sha1"}
		}
	}

	if d.HasChange("ike_phase_2") {

		if _, ok := d.GetOk("ike_phase_2"); ok {

			res := make(map[string]interface{})

			if d.HasChange("ike_phase_2.data_integrity") {
				res["data-integrity"] = d.Get("ike_phase_2.data_integrity")
			}
			if d.HasChange("ike_phase_2.encryption_algorithm") {
				res["encryption-algorithm"] = d.Get("ike_phase_2.encryption_algorithm")
			}
			if d.HasChange("ike_phase_2.ike_p2_use_pfs") {
				res["ike-p2-use-pfs"] = d.Get("ike_phase_2.ike_p2_use_pfs")
			}
			if d.HasChange("ike_phase_2.ike_p2_pfs_dh_grp") {
				res["ike-p2-pfs-dh-grp"] = d.Get("ike_phase_2.ike_p2_pfs_dh_grp")
			}
			if d.HasChange("ike_phase_2.ike_p2_rekey_time") {
				res["ike-p2-rekey-time"] = d.Get("ike_phase_2.ike_p2_rekey_time")
			}
			vpnCommunityMeshed["ike-phase-2"] = res
		} else {
			vpnCommunityMeshed["ike-phase-2"] = map[string]interface{}{"encryption-algorithm": "aes-128", "data-integrity": "sha1"}
		}
	}

	if d.HasChange("override_vpn_domains") {

		if v, ok := d.GetOk("override_vpn_domains"); ok {

			overrideVpnDomainsList := v.([]interface{})

			var overrideVpnDomainsPayload []map[string]interface{}

			for i := range overrideVpnDomainsList {

				Payload := make(map[string]interface{})

				if d.HasChange("override_vpn_domains." + strconv.Itoa(i) + ".gateway") {
					Payload["gateway"] = d.Get("override_vpn_domains." + strconv.Itoa(i) + ".gateway")
				}
				if d.HasChange("override_vpn_domains." + strconv.Itoa(i) + ".vpn_domain") {
					Payload["vpn-domain"] = d.Get("override_vpn_domains." + strconv.Itoa(i) + ".vpn_domain")
				}
				overrideVpnDomainsPayload = append(overrideVpnDomainsPayload, Payload)
			}
			vpnCommunityMeshed["override-vpn-domains"] = overrideVpnDomainsPayload
		} else {
			oldoverrideVpnDomains, _ := d.GetChange("override_vpn_domains")
			var overrideVpnDomainsToDelete []interface{}
			for _, i := range oldoverrideVpnDomains.([]interface{}) {
				overrideVpnDomainsToDelete = append(overrideVpnDomainsToDelete, i.(map[string]interface{})["name"].(string))
			}
			vpnCommunityMeshed["override-vpn-domains"] = map[string]interface{}{"remove": overrideVpnDomainsToDelete}
		}
	}

	if d.HasChange("shared_secrets") {

		if v, ok := d.GetOk("shared_secrets"); ok {

			sharedSecretsList := v.([]interface{})

			var sharedSecretsPayload []map[string]interface{}

			for i := range sharedSecretsList {

				Payload := make(map[string]interface{})

				if d.HasChange("shared_secrets." + strconv.Itoa(i) + ".external_gateway") {
					Payload["external-gateway"] = d.Get("shared_secrets." + strconv.Itoa(i) + ".external_gateway")
				}
				if d.HasChange("shared_secrets." + strconv.Itoa(i) + ".shared_secret") {
					Payload["shared-secret"] = d.Get("shared_secrets." + strconv.Itoa(i) + ".shared_secret")
				}
				sharedSecretsPayload = append(sharedSecretsPayload, Payload)
			}
			vpnCommunityMeshed["shared-secrets"] = sharedSecretsPayload
		} else {
			oldsharedSecrets, _ := d.GetChange("shared_secrets")
			var sharedSecretsToDelete []interface{}
			for _, i := range oldsharedSecrets.([]interface{}) {
				sharedSecretsToDelete = append(sharedSecretsToDelete, i.(map[string]interface{})["name"].(string))
			}
			vpnCommunityMeshed["shared-secrets"] = map[string]interface{}{"remove": sharedSecretsToDelete}
		}
	}

	if d.HasChange("tunnel_granularity") {
		vpnCommunityMeshed["tunnel-granularity"] = d.Get("tunnel_granularity")
	}

	if d.HasChange("granular_encryptions") {
		if v, ok := d.GetOk("granular_encryptions"); ok {
			granularEncryptions := v.([]interface{})
			if len(granularEncryptions) > 0 {
				var granularEncryptionsPayload []map[string]interface{}

				for i := range granularEncryptions {
					payload := make(map[string]interface{})
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".internal_gateway"); ok {
						payload["internal-gateway"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".external_gateway"); ok {
						payload["external-gateway"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".encryption_method"); ok {
						payload["encryption-method"] = v.(string)
					}
					if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".encryption_suite"); ok {
						payload["encryption-suite"] = v.(string)
					}
					if _, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1"); ok {
						ikePhase1Payload := make(map[string]interface{})
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.encryption_algorithm"); ok {
							ikePhase1Payload["encryption-algorithm"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.data_integrity"); ok {
							ikePhase1Payload["data-integrity"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.diffie_hellman_group"); ok {
							ikePhase1Payload["diffie-hellman-group"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_1.ike_p1_rekey_time"); ok {
							ikePhase1Payload["ike-p1-rekey-time"] = v.(string)
						}
						payload["ike-phase-1"] = ikePhase1Payload
					}
					if _, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2"); ok {
						ikePhase2Payload := make(map[string]interface{})
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.encryption_algorithm"); ok {
							ikePhase2Payload["encryption-algorithm"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.data_integrity"); ok {
							ikePhase2Payload["data-integrity"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_use_pfs"); ok {
							ikePhase2Payload["ike-p2-use-pfs"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_pfs_dh_grp"); ok {
							ikePhase2Payload["ike-p2-pfs-dh-grp"] = v.(string)
						}
						if v, ok := d.GetOk("granular_encryptions." + strconv.Itoa(i) + ".ike_phase_2.ike_p2_rekey_time"); ok {
							ikePhase2Payload["ike-p2-rekey-time"] = v.(int)
						}
						payload["ike-phase-2"] = ikePhase2Payload
					}
					granularEncryptionsPayload = append(granularEncryptionsPayload, payload)
				}
				vpnCommunityMeshed["granular-encryptions"] = granularEncryptionsPayload
			}
		} else {
			granularEncryptions, _ := d.GetChange("granular_encryptions")
			oldValues := granularEncryptions.([]interface{})
			if len(oldValues) > 0 {
				var toRemove []interface{}
				for _, v := range oldValues {
					obj := make(map[string]interface{})
					obj["internal-gateway"] = v.(map[string]interface{})["internal_gateway"].(string)
					obj["external-gateway"] = v.(map[string]interface{})["external_gateway"].(string)
					toRemove = append(toRemove, obj)
				}
				vpnCommunityMeshed["granular-encryptions"] = map[string]interface{}{"remove": toRemove}
			}
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			vpnCommunityMeshed["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			vpnCommunityMeshed["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if v, ok := d.GetOkExists("use_shared_secret"); ok {
		vpnCommunityMeshed["use-shared-secret"] = v.(bool)
	}

	if ok := d.HasChange("color"); ok {
		vpnCommunityMeshed["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		vpnCommunityMeshed["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		vpnCommunityMeshed["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		vpnCommunityMeshed["ignore-errors"] = v.(bool)
	}

	log.Println("Update VpnCommunityMeshed - Map = ", vpnCommunityMeshed)

	updateVpnCommunityMeshedRes, err := client.ApiCall("set-vpn-community-meshed", vpnCommunityMeshed, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !updateVpnCommunityMeshedRes.Success {
		if updateVpnCommunityMeshedRes.ErrorMsg != "" {
			return fmt.Errorf(updateVpnCommunityMeshedRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementVpnCommunityMeshed(d, m)
}

func deleteManagementVpnCommunityMeshed(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	vpnCommunityMeshedPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete VpnCommunityMeshed")

	deleteVpnCommunityMeshedRes, err := client.ApiCall("delete-vpn-community-meshed", vpnCommunityMeshedPayload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteVpnCommunityMeshedRes.Success {
		if deleteVpnCommunityMeshedRes.ErrorMsg != "" {
			return fmt.Errorf(deleteVpnCommunityMeshedRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
