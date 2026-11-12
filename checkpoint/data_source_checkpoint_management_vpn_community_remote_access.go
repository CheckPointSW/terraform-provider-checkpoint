package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"reflect"
)

func dataSourceManagementVpnCommunityRemoteAccess() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementVpnCommunityRemoteAccessRead,
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
			"gateways": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Gateway objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_groups": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of User group objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"tags": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of tag identifiers.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
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
			"encryption": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Encryption settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"encryption_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The encryption method to be used.",
						},
						"ike_phase_1": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "IKE Phase 1 settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configured_data_integrity_algorithms": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "The list of configured hash algorithms.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"configured_diffie_hellman_groups": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "The list of configured Diffie-Hellman groups.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"configured_encryption_algorithms": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "The list of configured encryption algorithms.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"multiple_key_exchanges": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Multiple Key Exchanges proposal object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"additional_key_exchange_1_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 1 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_2_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 2 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_3_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 3 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_4_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 4 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_5_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 5 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_6_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 6 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_7_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 7 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"key_exchange_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Key-Exchange methods to use. Can contain only Diffie-Hellman groups.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"use_multiple_key_exchanges": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether to use a proposal with Multiple Key Exchanges.",
									},
									"use_standard_proposal": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether to use a proposal with a single Diffie-Hellman group.",
									},
								},
							},
						},
						"ike_phase_2": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "IKE Phase 2 settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"configured_data_integrity_algorithms": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "The list of configured hash algorithms.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"configured_diffie_hellman_groups": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "The list of configured Diffie-Hellman groups.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"configured_encryption_algorithms": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "The list of configured encryption algorithms.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
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
									"enforce_encryption_alg_and_data_integrity_on_all_users": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enforce Encryption Algorithm and Data Integrity on all users.",
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
									"ike_p2_use_pfs": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether Perfect Forward Secrecy (PFS) is being used for IKE phase 2.",
									},
									"multiple_key_exchanges": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Multiple Key Exchanges proposal object to use when PFS is enabled and multiple key exchanges are configured.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"additional_key_exchange_1_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 1 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_2_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 2 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_3_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 3 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_4_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 4 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_5_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 5 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_6_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 6 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"additional_key_exchange_7_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Additional Key-Exchange 7 methods to use.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"key_exchange_methods": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Key-Exchange methods to use. Can contain only Diffie-Hellman groups.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"use_multiple_key_exchanges": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Indicates whether to use a proposal with Multiple Key Exchanges when PFS is enabled.",
									},
									"use_standard_proposal": {
										Type:        schema.TypeBool,
										Computed:    true,
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

func dataSourceManagementVpnCommunityRemoteAccessRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showVpnCommunityRemoteAccessRes, err := client.ApiCall("show-vpn-community-remote-access", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showVpnCommunityRemoteAccessRes.Success {
		return fmt.Errorf("%s", showVpnCommunityRemoteAccessRes.ErrorMsg)
	}

	vpnCommunityRemoteAccess := showVpnCommunityRemoteAccessRes.GetData()

	log.Println("Read VpnCommunityRemoteAccess - Show JSON = ", vpnCommunityRemoteAccess)

	if v := vpnCommunityRemoteAccess["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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
			if v := ikePhase1Show["multiple-key-exchanges"]; v != nil {
				multipleKeyExchangesShow := v.(map[string]interface{})
				multipleKeyExchangesState := make(map[string]interface{})
				if v := multipleKeyExchangesShow["additional-key-exchange-1-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_1_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-2-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_2_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-3-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_3_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-4-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_4_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-5-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_5_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-6-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_6_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-7-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_7_methods"] = v
				}
				if v := multipleKeyExchangesShow["key-exchange-methods"]; v != nil {
					multipleKeyExchangesState["key_exchange_methods"] = v
				}
				ikePhase1State["multiple_key_exchanges"] = []interface{}{multipleKeyExchangesState}
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
			if v := ikePhase2Show["ike-p2-use-pfs"]; v != nil {
				ikePhase2State["ike_p2_use_pfs"] = v
			}
			if v := ikePhase2Show["multiple-key-exchanges"]; v != nil {
				multipleKeyExchangesShow := v.(map[string]interface{})
				multipleKeyExchangesState := make(map[string]interface{})
				if v := multipleKeyExchangesShow["additional-key-exchange-1-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_1_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-2-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_2_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-3-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_3_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-4-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_4_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-5-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_5_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-6-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_6_methods"] = v
				}
				if v := multipleKeyExchangesShow["additional-key-exchange-7-methods"]; v != nil {
					multipleKeyExchangesState["additional_key_exchange_7_methods"] = v
				}
				if v := multipleKeyExchangesShow["key-exchange-methods"]; v != nil {
					multipleKeyExchangesState["key_exchange_methods"] = v
				}
				ikePhase2State["multiple_key_exchanges"] = []interface{}{multipleKeyExchangesState}
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
