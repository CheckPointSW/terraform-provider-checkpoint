package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"reflect"
	"strconv"
)

func resourceManagementVpnCommunityRemoteAccess() *schema.Resource {
	return &schema.Resource{
		Create: createManagementVpnCommunityRemoteAccess,
		Read:   readManagementVpnCommunityRemoteAccess,
		Update: updateManagementVpnCommunityRemoteAccess,
		Delete: deleteManagementVpnCommunityRemoteAccess,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"gateways": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of Gateway objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_groups": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of User group objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"tags": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"ignore_warnings": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring warnings.",
			},
			"ignore_errors": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.",
			},
			"encryption": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Encryption settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"encryption_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The encryption method to be used.",
						},
						"ike_phase_1": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "IKE Phase 1 settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configured_data_integrity_algorithms": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "The list of configured hash algorithms.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"configured_diffie_hellman_groups": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "The list of configured Diffie-Hellman groups.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"configured_encryption_algorithms": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "The list of configured encryption algorithms.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"data_integrity": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The hash algorithm to be used.",
									},
									"diffie_hellman_group": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Diffie-Hellman group to be used.",
									},
									"encryption_algorithm": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The encryption algorithm to be used.",
									},
									"ike_p1_rekey_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Indicates the time interval for IKE phase 1 renegotiation.",
									},
									"ike_p1_rekey_time_unit": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates the time unit for the 'ike-p1-rekey-time-unit' parameter, rounded up to minutes scale.",
									},
									"multiple_key_exchanges": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of the Multiple Key Exchanges proposal object.",
									},
									"use_multiple_key_exchanges": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether to use a proposal with Multiple Key Exchanges.",
									},
									"use_standard_proposal": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether to use a proposal with a single Diffie-Hellman group.",
									},
								},
							},
						},
						"ike_phase_2": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "IKE Phase 2 settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configured_data_integrity_algorithms": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "The list of configured hash algorithms.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"configured_diffie_hellman_groups": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "The list of configured Diffie-Hellman groups.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"configured_encryption_algorithms": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "The list of configured encryption algorithms.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"data_integrity": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The hash algorithm to be used.",
									},
									"encryption_algorithm": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The encryption algorithm to be used.",
									},
									"enforce_encryption_alg_and_data_integrity_on_all_users": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enforce Encryption Algorithm and Data Integrity on all users.",
									},
									"ike_p2_pfs_dh_grp": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The Diffie-Hellman group to be used.",
									},
									"ike_p2_rekey_time": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Indicates the time interval for IKE phase 2 renegotiation.",
									},
									"ike_p2_rekey_time_unit": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Indicates the time unit for the 'ike-p2-rekey-time-unit' parameter, rounded up to minutes scale.",
									},
									"ike_p2_use_pfs": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
									},
									"multiple_key_exchanges": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Name of the Multiple Key Exchanges proposal object to use when PFS is enabled and multiple key exchanges are configured.",
									},
									"use_multiple_key_exchanges": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether to use a proposal with Multiple Key Exchanges when PFS is enabled.",
									},
									"use_standard_proposal": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Indicates whether to use a proposal with a single Diffie-Hellman group when PFS is enabled.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func createManagementVpnCommunityRemoteAccess(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if v, ok := d.GetOk("name"); ok {
		payload["name"] = v.(string)
	}

	if v, ok := d.GetOk("gateways"); ok {
		payload["gateways"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("user_groups"); ok {
		payload["user-groups"] = v.(*schema.Set).List()
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
			payload["override-vpn-domains"] = overrideVpnDomainsPayload
		}
	}

	if v, ok := d.GetOk("tags"); ok {
		payload["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		payload["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		payload["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	if _, ok := d.GetOk("encryption"); ok {
		encryptionPayload := make(map[string]interface{})
		if v, ok := d.GetOk("encryption.0.encryption_method"); ok {
			encryptionPayload["encryption-method"] = v.(string)
		}
		if _, ok := d.GetOk("encryption.0.ike_phase_1"); ok {
			ikePhase1Payload := make(map[string]interface{})
			if v, ok := d.GetOk("encryption.0.ike_phase_1.0.configured_data_integrity_algorithms"); ok {
				ikePhase1Payload["configured-data-integrity-algorithms"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_1.0.configured_diffie_hellman_groups"); ok {
				ikePhase1Payload["configured-diffie-hellman-groups"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_1.0.configured_encryption_algorithms"); ok {
				ikePhase1Payload["configured-encryption-algorithms"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_1.0.data_integrity"); ok {
				ikePhase1Payload["data-integrity"] = v.(string)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_1.0.diffie_hellman_group"); ok {
				ikePhase1Payload["diffie-hellman-group"] = v.(string)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_1.0.encryption_algorithm"); ok {
				ikePhase1Payload["encryption-algorithm"] = v.(string)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_1.0.ike_p1_rekey_time"); ok {
				ikePhase1Payload["ike-p1-rekey-time"] = v.(int)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_1.0.ike_p1_rekey_time_unit"); ok {
				ikePhase1Payload["ike-p1-rekey-time-unit"] = v.(string)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_1.0.multiple_key_exchanges"); ok {
				ikePhase1Payload["multiple-key-exchanges"] = v.(string)
			}
			if v, ok := d.GetOkExists("encryption.0.ike_phase_1.0.use_multiple_key_exchanges"); ok {
				ikePhase1Payload["use-multiple-key-exchanges"] = v.(bool)
			}
			if v, ok := d.GetOkExists("encryption.0.ike_phase_1.0.use_standard_proposal"); ok {
				ikePhase1Payload["use-standard-proposal"] = v.(bool)
			}
			encryptionPayload["ike-phase-1"] = ikePhase1Payload
		}
		if _, ok := d.GetOk("encryption.0.ike_phase_2"); ok {
			ikePhase2Payload := make(map[string]interface{})
			if v, ok := d.GetOk("encryption.0.ike_phase_2.0.configured_data_integrity_algorithms"); ok {
				ikePhase2Payload["configured-data-integrity-algorithms"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_2.0.configured_diffie_hellman_groups"); ok {
				ikePhase2Payload["configured-diffie-hellman-groups"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_2.0.configured_encryption_algorithms"); ok {
				ikePhase2Payload["configured-encryption-algorithms"] = v.(*schema.Set).List()
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_2.0.data_integrity"); ok {
				ikePhase2Payload["data-integrity"] = v.(string)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_2.0.encryption_algorithm"); ok {
				ikePhase2Payload["encryption-algorithm"] = v.(string)
			}
			if v, ok := d.GetOkExists("encryption.0.ike_phase_2.0.enforce_encryption_alg_and_data_integrity_on_all_users"); ok {
				ikePhase2Payload["enforce-encryption-alg-and-data-integrity-on-all-users"] = v.(bool)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_2.0.ike_p2_pfs_dh_grp"); ok {
				ikePhase2Payload["ike-p2-pfs-dh-grp"] = v.(string)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_2.0.ike_p2_rekey_time"); ok {
				ikePhase2Payload["ike-p2-rekey-time"] = v.(int)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_2.0.ike_p2_rekey_time_unit"); ok {
				ikePhase2Payload["ike-p2-rekey-time-unit"] = v.(string)
			}
			if v, ok := d.GetOkExists("encryption.0.ike_phase_2.0.ike_p2_use_pfs"); ok {
				ikePhase2Payload["ike-p2-use-pfs"] = v.(bool)
			}
			if v, ok := d.GetOk("encryption.0.ike_phase_2.0.multiple_key_exchanges"); ok {
				ikePhase2Payload["multiple-key-exchanges"] = v.(string)
			}
			if v, ok := d.GetOkExists("encryption.0.ike_phase_2.0.use_multiple_key_exchanges"); ok {
				ikePhase2Payload["use-multiple-key-exchanges"] = v.(bool)
			}
			if v, ok := d.GetOkExists("encryption.0.ike_phase_2.0.use_standard_proposal"); ok {
				ikePhase2Payload["use-standard-proposal"] = v.(bool)
			}
			encryptionPayload["ike-phase-2"] = ikePhase2Payload
		}
		payload["encryption"] = encryptionPayload
	}

	SetVpnCommunityRemoteAccessRes, _ := client.ApiCall("set-vpn-community-remote-access", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !SetVpnCommunityRemoteAccessRes.Success {
		return fmt.Errorf("%s", SetVpnCommunityRemoteAccessRes.ErrorMsg)
	}

	d.SetId(SetVpnCommunityRemoteAccessRes.GetData()["uid"].(string))

	return readManagementVpnCommunityRemoteAccess(d, m)
}

func updateManagementVpnCommunityRemoteAccess(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	var payload = map[string]interface{}{}

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		payload["name"] = oldName
		payload["new-name"] = newName
	} else {
		payload["name"] = d.Get("name")
	}

	if ok := d.HasChange("gateways"); ok {
		if v, ok := d.GetOk("gateways"); ok {
			payload["gateways"] = v.(*schema.Set).List()
		} else {
			oldGateways, _ := d.GetChange("gateways")
			payload["gateways"] = map[string]interface{}{"remove": oldGateways.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("user_groups"); ok {
		if v, ok := d.GetOk("user_groups"); ok {
			payload["user-groups"] = v.(*schema.Set).List()
		} else {
			oldUserGroups, _ := d.GetChange("gateways")
			payload["user-groups"] = map[string]interface{}{"remove": oldUserGroups.(*schema.Set).List()}
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
			payload["override-vpn-domains"] = overrideVpnDomainsPayload
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			payload["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			payload["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		payload["color"] = d.Get("color").(string)
	}

	if ok := d.HasChange("comments"); ok {
		payload["comments"] = d.Get("comments").(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v.(bool)
	}

	if d.HasChange("encryption") {
		if _, ok := d.GetOk("encryption"); ok {
			encryptionPayload := make(map[string]interface{})
			if v, ok := d.GetOk("encryption.0.encryption_method"); ok {
				encryptionPayload["encryption-method"] = v.(string)
			}
			if _, ok := d.GetOk("encryption.0.ike_phase_1"); ok {
				ikePhase1Payload := make(map[string]interface{})
				if v, ok := d.GetOk("encryption.0.ike_phase_1.0.configured_data_integrity_algorithms"); ok {
					ikePhase1Payload["configured-data-integrity-algorithms"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_1.0.configured_diffie_hellman_groups"); ok {
					ikePhase1Payload["configured-diffie-hellman-groups"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_1.0.configured_encryption_algorithms"); ok {
					ikePhase1Payload["configured-encryption-algorithms"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_1.0.data_integrity"); ok {
					ikePhase1Payload["data-integrity"] = v.(string)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_1.0.diffie_hellman_group"); ok {
					ikePhase1Payload["diffie-hellman-group"] = v.(string)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_1.0.encryption_algorithm"); ok {
					ikePhase1Payload["encryption-algorithm"] = v.(string)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_1.0.ike_p1_rekey_time"); ok {
					ikePhase1Payload["ike-p1-rekey-time"] = v.(int)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_1.0.ike_p1_rekey_time_unit"); ok {
					ikePhase1Payload["ike-p1-rekey-time-unit"] = v.(string)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_1.0.multiple_key_exchanges"); ok {
					ikePhase1Payload["multiple-key-exchanges"] = v.(string)
				}
				if v, ok := d.GetOkExists("encryption.0.ike_phase_1.0.use_multiple_key_exchanges"); ok {
					ikePhase1Payload["use-multiple-key-exchanges"] = v.(bool)
				}
				if v, ok := d.GetOkExists("encryption.0.ike_phase_1.0.use_standard_proposal"); ok {
					ikePhase1Payload["use-standard-proposal"] = v.(bool)
				}
				encryptionPayload["ike-phase-1"] = ikePhase1Payload
			}
			if _, ok := d.GetOk("encryption.0.ike_phase_2"); ok {
				ikePhase2Payload := make(map[string]interface{})
				if v, ok := d.GetOk("encryption.0.ike_phase_2.0.configured_data_integrity_algorithms"); ok {
					ikePhase2Payload["configured-data-integrity-algorithms"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_2.0.configured_diffie_hellman_groups"); ok {
					ikePhase2Payload["configured-diffie-hellman-groups"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_2.0.configured_encryption_algorithms"); ok {
					ikePhase2Payload["configured-encryption-algorithms"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_2.0.data_integrity"); ok {
					ikePhase2Payload["data-integrity"] = v.(string)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_2.0.encryption_algorithm"); ok {
					ikePhase2Payload["encryption-algorithm"] = v.(string)
				}
				if v, ok := d.GetOkExists("encryption.0.ike_phase_2.0.enforce_encryption_alg_and_data_integrity_on_all_users"); ok {
					ikePhase2Payload["enforce-encryption-alg-and-data-integrity-on-all-users"] = v.(bool)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_2.0.ike_p2_pfs_dh_grp"); ok {
					ikePhase2Payload["ike-p2-pfs-dh-grp"] = v.(string)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_2.0.ike_p2_rekey_time"); ok {
					ikePhase2Payload["ike-p2-rekey-time"] = v.(int)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_2.0.ike_p2_rekey_time_unit"); ok {
					ikePhase2Payload["ike-p2-rekey-time-unit"] = v.(string)
				}
				if v, ok := d.GetOkExists("encryption.0.ike_phase_2.0.ike_p2_use_pfs"); ok {
					ikePhase2Payload["ike-p2-use-pfs"] = v.(bool)
				}
				if v, ok := d.GetOk("encryption.0.ike_phase_2.0.multiple_key_exchanges"); ok {
					ikePhase2Payload["multiple-key-exchanges"] = v.(string)
				}
				if v, ok := d.GetOkExists("encryption.0.ike_phase_2.0.use_multiple_key_exchanges"); ok {
					ikePhase2Payload["use-multiple-key-exchanges"] = v.(bool)
				}
				if v, ok := d.GetOkExists("encryption.0.ike_phase_2.0.use_standard_proposal"); ok {
					ikePhase2Payload["use-standard-proposal"] = v.(bool)
				}
				encryptionPayload["ike-phase-2"] = ikePhase2Payload
			}
			payload["encryption"] = encryptionPayload
		}
	}

	SetVpnCommunityRemoteAccessRes, _ := client.ApiCall("set-vpn-community-remote-access", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if !SetVpnCommunityRemoteAccessRes.Success {
		return fmt.Errorf("%s", SetVpnCommunityRemoteAccessRes.ErrorMsg)
	}

	return readManagementVpnCommunityRemoteAccess(d, m)
}

func readManagementVpnCommunityRemoteAccess(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showVpnCommunityRemoteAccessRes, err := client.ApiCall("show-vpn-community-remote-access", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showVpnCommunityRemoteAccessRes.Success {
		if objectNotFound(showVpnCommunityRemoteAccessRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showVpnCommunityRemoteAccessRes.ErrorMsg)
	}

	vpnCommunityRemoteAccess := showVpnCommunityRemoteAccessRes.GetData()

	log.Println("Read VpnCommunityRemoteAccess - Show JSON = ", vpnCommunityRemoteAccess)

	if v := vpnCommunityRemoteAccess["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if vpnCommunityRemoteAccess["gateways"] != nil {
		gatewaysJson, ok := vpnCommunityRemoteAccess["gateways"].([]interface{})
		if ok {
			gwIds := make([]string, 0)
			if len(gatewaysJson) > 0 {
				for _, gw := range gatewaysJson {
					gwIds = append(gwIds, gw.(map[string]interface{})["name"].(string))
				}
			}
			_ = d.Set("gateways", gwIds)
		}
	} else {
		_ = d.Set("gateways", nil)
	}

	if vpnCommunityRemoteAccess["user-groups"] != nil {
		userGroupsJson, ok := vpnCommunityRemoteAccess["user-groups"].([]interface{})
		userGroupIds := make([]string, 0)
		if ok {
			if len(userGroupsJson) > 0 {
				for _, userGroup := range userGroupsJson {
					userGroupIds = append(userGroupIds, userGroup.(map[string]interface{})["name"].(string))
				}
			}
		}
		_, userGroupsInConf := d.GetOk("user_groups")
		defaultUserGroups := []string{"All Users"}
		if reflect.DeepEqual(defaultUserGroups, userGroupIds) && !userGroupsInConf {
			_ = d.Set("user_groups", []string{})
		} else {
			_ = d.Set("user_groups", userGroupIds)
		}
	} else {
		_ = d.Set("user_groups", nil)
	}

	if vpnCommunityRemoteAccess["override-vpn-domains"] != nil {
		overrideVpnDomainsList := vpnCommunityRemoteAccess["override-vpn-domains"].([]interface{})
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

	if vpnCommunityRemoteAccess["tags"] != nil {
		tagsJson, ok := vpnCommunityRemoteAccess["tags"].([]interface{})
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

	if v := vpnCommunityRemoteAccess["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := vpnCommunityRemoteAccess["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := vpnCommunityRemoteAccess["encryption"]; v != nil {
		encryptionShow := v.(map[string]interface{})
		encryptionState := make(map[string]interface{})
		if v := encryptionShow["encryption-method"]; v != nil {
			encryptionState["encryption_method"] = v
		}
		if v := encryptionShow["ike-phase-1"]; v != nil {
			ikePhase1Show := v.(map[string]interface{})
			ikePhase1State := make(map[string]interface{})
			if v := ikePhase1Show["configured-data-integrity-algorithms"]; v != nil {
				ikePhase1State["configured_data_integrity_algorithms"] = v
			}
			if v := ikePhase1Show["configured-diffie-hellman-groups"]; v != nil {
				ikePhase1State["configured_diffie_hellman_groups"] = v
			}
			if v := ikePhase1Show["configured-encryption-algorithms"]; v != nil {
				ikePhase1State["configured_encryption_algorithms"] = v
			}
			if v := ikePhase1Show["data-integrity"]; v != nil {
				ikePhase1State["data_integrity"] = v
			}
			if v := ikePhase1Show["diffie-hellman-group"]; v != nil {
				ikePhase1State["diffie_hellman_group"] = v
			}
			if v := ikePhase1Show["encryption-algorithm"]; v != nil {
				ikePhase1State["encryption_algorithm"] = v
			}
			if v := ikePhase1Show["ike-p1-rekey-time"]; v != nil {
				ikePhase1State["ike_p1_rekey_time"] = v
			}
			if v := ikePhase1Show["ike-p1-rekey-time-unit"]; v != nil {
				ikePhase1State["ike_p1_rekey_time_unit"] = v
			}
			if v := ikePhase1Show["multiple-key-exchanges"]; v != nil {
				ikePhase1State["multiple_key_exchanges"] = v.(map[string]interface{})["name"]
			}
			if v := ikePhase1Show["use-multiple-key-exchanges"]; v != nil {
				ikePhase1State["use_multiple_key_exchanges"] = v
			}
			if v := ikePhase1Show["use-standard-proposal"]; v != nil {
				ikePhase1State["use_standard_proposal"] = v
			}
			encryptionState["ike_phase_1"] = []interface{}{ikePhase1State}
		}
		if v := encryptionShow["ike-phase-2"]; v != nil {
			ikePhase2Show := v.(map[string]interface{})
			ikePhase2State := make(map[string]interface{})
			if v := ikePhase2Show["configured-data-integrity-algorithms"]; v != nil {
				ikePhase2State["configured_data_integrity_algorithms"] = v
			}
			if v := ikePhase2Show["configured-diffie-hellman-groups"]; v != nil {
				ikePhase2State["configured_diffie_hellman_groups"] = v
			}
			if v := ikePhase2Show["configured-encryption-algorithms"]; v != nil {
				ikePhase2State["configured_encryption_algorithms"] = v
			}
			if v := ikePhase2Show["data-integrity"]; v != nil {
				ikePhase2State["data_integrity"] = v
			}
			if v := ikePhase2Show["encryption-algorithm"]; v != nil {
				ikePhase2State["encryption_algorithm"] = v
			}
			if v := ikePhase2Show["enforce-encryption-alg-and-data-integrity-on-all-users"]; v != nil {
				ikePhase2State["enforce_encryption_alg_and_data_integrity_on_all_users"] = v
			}
			if v := ikePhase2Show["ike-p2-pfs-dh-grp"]; v != nil {
				ikePhase2State["ike_p2_pfs_dh_grp"] = v
			}
			if v := ikePhase2Show["ike-p2-rekey-time"]; v != nil {
				ikePhase2State["ike_p2_rekey_time"] = v
			}
			if v := ikePhase2Show["ike-p2-rekey-time-unit"]; v != nil {
				ikePhase2State["ike_p2_rekey_time_unit"] = v
			}
			if v := ikePhase2Show["ike-p2-use-pfs"]; v != nil {
				ikePhase2State["ike_p2_use_pfs"] = v
			}
			if v := ikePhase2Show["multiple-key-exchanges"]; v != nil {
				ikePhase2State["multiple_key_exchanges"] = v.(map[string]interface{})["name"]
			}
			if v := ikePhase2Show["use-multiple-key-exchanges"]; v != nil {
				ikePhase2State["use_multiple_key_exchanges"] = v
			}
			if v := ikePhase2Show["use-standard-proposal"]; v != nil {
				ikePhase2State["use_standard_proposal"] = v
			}
			encryptionState["ike_phase_2"] = []interface{}{ikePhase2State}
		}
		_ = d.Set("encryption", []interface{}{encryptionState})
	}

	return nil
}

func deleteManagementVpnCommunityRemoteAccess(d *schema.ResourceData, m interface{}) error {
	d.SetId("")
	return nil
}
