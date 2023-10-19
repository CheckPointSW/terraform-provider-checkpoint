package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataSourceManagementSimpleCluster() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSimpleClusterRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object unique identifier.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IPv6 address.",
			},
			"cluster_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster mode.",
			},
			"geo_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Cluster High Availability Geo mode. This setting applies only to a cluster deployed in a cloud. Available when the cluster mode equals \"cluster-xl-ha\".",
			},
			"advanced_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "N/A",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_persistence": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Handling established connections when installing a new policy.",
						},
						"sam": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "SAM.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"forward_to_other_sam_servers": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Forward SAM clients' requests to other SAM servers.",
									},
									"use_early_versions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Use early versions compatibility mode.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Use early versions compatibility mode.",
												},
												"compatibility_mode": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Early versions compatibility mode.",
												},
											},
										},
									},
									"purge_sam_file": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Purge SAM File.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Purge SAM File.",
												},
												"purge_when_size_reaches_to": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Purge SAM File When it Reaches to.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"enable_https_inspection": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API.",
			},
			"fetch_policy": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Security management server(s) to fetch the policy from.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hit_count": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Hit count tracks the number of connections each rule matches.",
			},
			"https_inspection": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "HTTPS inspection.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_on_failure": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to bypass all requests (Fail-open) in case of internal system error.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"site_categorization_allow_mode": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to 'background' in order to allowed requests until categorization is complete.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"deny_untrusted_server_cert": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with untrusted server certificate.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"deny_revoked_server_cert": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"deny_expired_server_cert": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with expired server certificate.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
					},
				},
			},
			"identity_awareness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Identity awareness blade enabled.",
			},
			"identity_awareness_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Gateway Identity Awareness settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"browser_based_authentication": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Browser Based Authentication source.",
						},
						"browser_based_authentication_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Browser Based Authentication settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authentication Settings for Browser Based Authentication.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Authentication method.",
												},
												"identity_provider": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Identity provider object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"identity provider\".",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"radius": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Radius server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"radius\".",
												},
												"users_directories": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Users directories.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"external_user_profile": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "External user profile.",
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Internal users.",
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Users from external directories.",
															},
															"specific": {
																Type:        schema.TypeSet,
																Computed:    true,
																Description: "LDAP AU objects identified by the name or UID. Must be set when \"users-from-external-directories\" was selected to be \"specific\".",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"browser_based_authentication_portal_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Browser Based Authentication portal settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"portal_web_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal web settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"aliases": {
																Type:        schema.TypeSet,
																Computed:    true,
																Description: "List of URL aliases that are redirected to the main portal URL.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"main_url": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The main URL for the web portal.",
															},
														},
													},
												},
												"certificate_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal certificate settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"base64_certificate": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate file encoded in Base64 with padding.  This file must be in the *.p12 format.",
															},
															"base64_password": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Password (encoded in Base64 with padding) for the certificate file.",
															},
														},
													},
												},
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"vpn": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"identity_agent": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Identity Agent source.",
						},
						"identity_agent_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity Agent settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"agents_interval_keepalive": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Agents send keepalive period (minutes).",
									},
									"user_reauthenticate_interval": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Agent reauthenticate time interval (minutes).",
									},
									"authentication_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authentication Settings for Identity Agent.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Authentication method.",
												},
												"radius": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Radius server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"radius\".",
												},
												"users_directories": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Users directories.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"external_user_profile": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "External user profile.",
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Internal users.",
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Users from external directories.",
															},
															"specific": {
																Type:        schema.TypeSet,
																Computed:    true,
																Description: "LDAP AU objects identified by the name or UID. Must be set when \"users-from-external-directories\" was selected to be \"specific\".",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"identity_agent_portal_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Identity Agent accessibility settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"vpn": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"identity_collector": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Identity Collector source.",
						},
						"identity_collector_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity Collector settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authorized_clients": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authorized Clients.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"client": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Host / Network Group Name or UID.",
												},
												"client_secret": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Client Secret.",
												},
											},
										},
									},
									"authentication_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authentication Settings for Identity Collector.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"users_directories": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Users directories.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"external_user_profile": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "External user profile.",
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Internal users.",
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Users from external directories.",
															},
															"specific": {
																Type:        schema.TypeSet,
																Computed:    true,
																Description: "LDAP AU objects identified by the name or UID. Must be set when \"users-from-external-directories\" was selected to be \"specific\".",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
											},
										},
									},
									"client_access_permissions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Identity Collector accessibility settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"vpn": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"identity_sharing_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity sharing settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"share_with_other_gateways": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable identity sharing with other gateways.",
									},
									"receive_from_other_gateways": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable receiving identity from other gateways.",
									},
									"receive_from": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Gateway(s) to receive identity from.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"proxy_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity-Awareness Proxy settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"detect_using_x_forward_for": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP.",
									},
								},
							},
						},
						"remote_access": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Remote Access Identity source.",
						},
					},
				},
			},
			"ips_update_policy": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Specifies whether the IPS will be downloaded from the Management or directly to the Gateway.",
			},
			"nat_hide_internal_interfaces": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Hide internal networks behind the Gateway's external IP.",
			},
			"nat_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to add automatic address translation rules.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address.",
						},
						"hide_behind": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Hide behind method. This parameter is forbidden in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"platform_portal_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Platform portal settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"portal_web_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal web settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aliases": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "List of URL aliases that are redirected to the main portal URL.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"main_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The main URL for the web portal.",
									},
								},
							},
						},
						"certificate_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal certificate settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"base64_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The certificate file encoded in Base64 with padding.  This file must be in the *.p12 format.",
									},
									"base64_password": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Password (encoded in Base64 with padding) for the certificate file.",
									},
								},
							},
						},
						"accessibility": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal access settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_access_from": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Allowed access to the web portal (based on interfaces, or security policy).",
									},
									"internal_access_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Configuration of the additional portal access settings for internal interfaces only.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"undefined": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
												},
												"dmz": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
												},
												"vpn": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"proxy_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Proxy Server for Gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_custom_proxy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use custom proxy settings for this network object.",
							Default:     false,
						},
						"proxy_server": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"port": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "N/A",
							Default:     80,
						},
					},
				},
			},
			"qos": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "QoS.",
			},
			"usercheck_portal_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "UserCheck portal settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "State of the web portal (enabled or disabled). The supported blades are: {'Application Control', 'URL Filtering', 'Data Loss Prevention', 'Anti Virus', 'Anti Bot', 'Threat Emulation', 'Threat Extraction', 'Data Awareness'}.",
						},
						"portal_web_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal web settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aliases": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "List of URL aliases that are redirected to the main portal URL.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"main_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The main URL for the web portal.",
									},
								},
							},
						},
						"certificate_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal certificate settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"base64_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The certificate file encoded in Base64 with padding.  This file must be in the *.p12 format.",
									},
									"base64_password": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Password (encoded in Base64 with padding) for the certificate file.",
									},
								},
							},
						},
						"accessibility": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal access settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_access_from": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Allowed access to the web portal (based on interfaces, or security policy).",
									},
									"internal_access_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Configuration of the additional portal access settings for internal interfaces only.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"undefined": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
												},
												"dmz": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
												},
												"vpn": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Controls portal access settings for interfaces that are part of a VPN Encryption Domain.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"zero_phishing": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Zero Phishing blade enabled.",
			},
			"zero_phishing_fqdn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Zero Phishing gateway FQDN.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"interface_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cluster interface type.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 address.",
						},
						"ipv4_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network address.",
						},
						"ipv6_network_mask": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network address.",
						},
						"ipv4_mask_length": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 network mask length.",
						},
						"ipv6_mask_length": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv6 network mask length.",
						},
						"anti_spoofing": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Anti spoofing.",
						},
						"anti_spoofing_settings": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Anti spoofing settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).",
									},
								},
							},
						},
						"multicast_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Multicast IP Address.",
						},
						"multicast_address_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Multicast Address Type.",
						},
						"security_zone": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Security zone.",
						},
						"security_zone_settings": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Security zone settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_calculated": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Security Zone is calculated according to where the interface leads to.",
									},
									"specific_zone": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Security Zone specified manually.",
									},
								},
							},
						},
						"topology": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topology.",
						},
						"topology_settings": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Topology settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_leads_to_dmz": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether this interface leads to demilitarized zone (perimeter network).",
									},
									"ip_address_behind_this_interface": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Ip address behind this interface.",
									},
									"specific_network": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Network behind this interface.",
									},
								},
							},
						},
						"topology_automatic_calculation": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Shows the automatic topology calculation.",
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
				},
			},
			"members": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cluster members.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IPv4 or IPv6 address.",
						},
						"sic_state": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secure Internal Communication name.",
						},
						"sic_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Secure Internal Communication state.",
						},
						"interfaces": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Network interfaces.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Object name. Should be unique in the domain.",
									},
									"ipv4_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv4 address.",
									},
									"ipv6_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv6 address.",
									},
									"ipv4_network_mask": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv4 network address.",
									},
									"ipv6_network_mask": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv6 network address.",
									},
									"ipv4_mask_length": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv4 network mask length.",
									},
									"ipv6_mask_length": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IPv6 network mask length.",
									},
								},
							},
						},
					},
				},
			},
			"anti_bot": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Anti-Bot blade enabled.",
			},
			"anti_virus": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Anti-Virus blade enabled.",
			},
			"application_control": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Application Control blade enabled.",
			},
			"content_awareness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Content Awareness blade enabled.",
			},
			"data_awareness": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Data Awareness blade enabled.",
			},
			"firewall": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Firewall blade enabled.",
			},
			"firewall_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Firewall settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_calculate_connections_hash_table_size_and_memory_pool": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Auto calculate connections hash table size and memory pool.",
						},
						"auto_maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Auto maximum limit for concurrent connections.",
						},
						"connections_hash_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Connections hash size.",
						},
						"maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum limit for concurrent connections.",
						},
						"maximum_memory_pool_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum memory pool size.",
						},
						"memory_pool_size": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Memory pool size.",
						},
					},
				},
			},
			"ips": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Intrusion Prevention System blade enabled.",
			},
			"threat_emulation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Threat Emulation blade enabled.",
			},
			"url_filtering": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "URL Filtering blade enabled.",
			},
			"dynamic_ip": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Dynamic IP address.",
			},
			"os_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "OS name.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster platform version.",
			},
			"hardware": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Cluster platform hardware.",
			},
			"sic_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication name.",
			},
			"sic_state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication state.",
			},
			"save_logs_locally": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Save logs locally.",
			},
			"send_alerts_to_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Server(s) to send alerts to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Backup server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_server": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"logs_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Logs settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable alert when free disk space is below threshold.",
						},
						"alert_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert when free disk space below metrics.",
						},
						"alert_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Alert when free disk space below threshold.",
						},
						"alert_when_free_disk_space_below_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert when free disk space below type.",
						},
						"before_delete_keep_logs_from_the_last_days": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable before delete keep logs from the last days.",
						},
						"before_delete_keep_logs_from_the_last_days_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Before delete keep logs from the last days threshold.",
						},
						"before_delete_run_script": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable Before delete run script.",
						},
						"before_delete_run_script_command": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Before delete run script command.",
						},
						"delete_index_files_older_than_days": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete index files older than days.",
						},
						"delete_index_files_older_than_days_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete index files older than days threshold.",
						},
						"delete_index_files_when_index_size_above": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete index files when index size above.",
						},
						"delete_index_files_when_index_size_above_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete index files when index size above metrics.",
						},
						"delete_index_files_when_index_size_above_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete index files when index size above threshold.",
						},
						"delete_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable delete when free disk space below.",
						},
						"delete_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete when free disk space below metric.",
						},
						"delete_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Delete when free disk space below threshold.",
						},
						"detect_new_citrix_ica_application_names": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable detect new citrix ica application names.",
						},
						"forward_logs_to_log_server": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable forward logs to log server.",
						},
						"forward_logs_to_log_server_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forward logs to log server name.",
						},
						"forward_logs_to_log_server_schedule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Forward logs to log server schedule name.",
						},
						"free_disk_space_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Free disk space metrics.",
						},
						"perform_log_rotate_before_log_forwarding": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable perform log rotate before log forwarding.",
						},
						"reject_connections_when_free_disk_space_below_threshold": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable reject connections when free disk space below threshold.",
						},
						"reserve_for_packet_capture_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Reserve for packet capture metrics.",
						},
						"reserve_for_packet_capture_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Reserve for packet capture threshold.",
						},
						"rotate_log_by_file_size": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable rotate log by file size.",
						},
						"rotate_log_file_size_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Log file size threshold.",
						},
						"rotate_log_on_schedule": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable rotate log on schedule.",
						},
						"rotate_log_schedule_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Rotate log schedule name.",
						},
						"stop_logging_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable stop logging when free disk space below.",
						},
						"stop_logging_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Stop logging when free disk space below threshold.",
						},
						"turn_on_qos_logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable turn on qos logging.",
						},
						"update_account_log_every": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Update account log in every amount of seconds.",
						},
					},
				},
			},
			"vpn": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "VPN blade enabled.",
			},
			"vpn_settings": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Gateway VPN settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Authentication.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_clients": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Collection of VPN Authentication clients identified by the name or UID.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"link_selection": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Link Selection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_selection": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IP selection",
									},
									"dns_resolving_hostname": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "DNS Resolving Hostname. Must be set when \"ip-selection\" was selected to be \"dns-resolving-from-hostname\".",
									},
									"ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IP Address. Must be set when \"ip-selection\" was selected to be \"use-selected-address-from-topology\" or \"use-statically-nated-ip\"",
									},
								},
							},
						},
						"maximum_concurrent_ike_negotiations": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum concurrent ike negotiations",
						},
						"maximum_concurrent_tunnels": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Maximum concurrent tunnels",
						},
						"office_mode": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Office Mode Permissions. When selected to be \"off\", all the other definitions are irrelevant.",
									},
									"group": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Group. Identified by name or UID. Must be set when \"office-mode-permissions\" was selected to be \"group\".",
									},
									"allocate_ip_address_from": {
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"radius_server": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Radius server used to authenticate the user.",
												},
												"use_allocate_method": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Use Allocate Method.",
												},
												"allocate_method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Using either Manual (IP Pool) or Automatic (DHCP). Must be set when \"use-allocate-method\" is true.",
												},
												"manual_network": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Manual Network. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"manual\".",
												},
												"dhcp_server": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "DHCP Server. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"virtual_ip_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Virtual IPV4 address for DHCP server replies. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"dhcp_mac_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Calculated MAC address for DHCP allocation. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"optional_parameters": {
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"use_primary_dns_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Primary DNS Server.",
															},
															"primary_dns_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Primary DNS Server. Identified by name or UID. Must be set when \"use-primary-dns-server\" is true and can not be set when \"use-primary-dns-server\" is false.",
															},
															"use_first_backup_dns_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use First Backup DNS Server.",
															},
															"first_backup_dns_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "First Backup DNS Server. Identified by name or UID. Must be set when \"use-first-backup-dns-server\" is true and can not be set when \"use-first-backup-dns-server\" is false.",
															},
															"use_second_backup_dns_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Second Backup DNS Server.",
															},
															"second_backup_dns_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Second Backup DNS Server. Identified by name or UID. Must be set when \"use-second-backup-dns-server\" is true and can not be set when \"use-second-backup-dns-server\" is false.",
															},
															"dns_suffixes": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "DNS Suffixes.",
															},
															"use_primary_wins_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Primary WINS Server.",
															},
															"primary_wins_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Primary WINS Server. Identified by name or UID. Must be set when \"use-primary-wins-server\" is true and can not be set when \"use-primary-wins-server\" is false.",
															},
															"use_first_backup_wins_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use First Backup WINS Server.",
															},
															"first_backup_wins_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "First Backup WINS Server. Identified by name or UID. Must be set when \"use-first-backup-wins-server\" is true and can not be set when \"use-first-backup-wins-server\" is false.",
															},
															"use_second_backup_wins_server": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Use Second Backup WINS Server.",
															},
															"second_backup_wins_server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Second Backup WINS Server. Identified by name or UID. Must be set when \"use-second-backup-wins-server\" is true and can not be set when \"use-second-backup-wins-server\" is false.",
															},
															"ip_lease_duration": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "IP Lease Duration in Minutes. The value must be in the range 2-32767.",
															},
														},
													},
												},
											},
										},
									},
									"support_multiple_interfaces": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support connectivity enhancement for gateways with multiple external interfaces.",
									},
									"perform_anti_spoofing": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Perform Anti-Spoofing on Office Mode addresses.",
									},
									"anti_spoofing_additional_addresses": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Additional IP Addresses for Anti-Spoofing. Identified by name or UID. Must be set when \"perform-anti-spoofings\" is true.",
									},
								},
							},
						},
						"remote_access": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Remote Access.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"support_l2tp": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support L2TP (relevant only when office mode is active).",
									},
									"l2tp_auth_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "L2TP Authentication Method. Must be set when \"support-l2tp\" is true.",
									},
									"l2tp_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "L2TP Certificate. Must be set when \"l2tp-auth-method\" was selected to be \"certificate\". Insert \"defaultCert\" when you want to use the default certificate.",
									},
									"allow_vpn_clients_to_route_traffic": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Allow VPN clients to route traffic.",
									},
									"support_nat_traversal_mechanism": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support NAT traversal mechanism (UDP encapsulation).",
									},
									"nat_traversal_service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Allocated NAT traversal UDP service. Identified by name or UID. Must be set when \"support-nat-traversal-mechanism\" is true.",
									},
									"support_visitor_mode": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Support Visitor Mode.",
									},
									"visitor_mode_service": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "TCP Service for Visitor Mode. Identified by name or UID. Must be set when \"support-visitor-mode\" is true.",
									},
									"visitor_mode_interface": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Interface for Visitor Mode. Must be set when \"support-visitor-mode\" is true. Insert IPV4 Address of existing interface or \"All IPs\" when you want all interfaces.",
									},
								},
							},
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway VPN domain identified by the name or UID.",
						},
						"vpn_domain_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Gateway VPN domain type.",
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
		},
	}
}

func dataSourceManagementSimpleClusterRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showClusterRes, err := client.ApiCall("show-simple-cluster", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showClusterRes.Success {
		return fmt.Errorf(showClusterRes.ErrorMsg)
	}

	cluster := showClusterRes.GetData()

	// If total interfaces above 50, Run show-simple-cluster with interface-limit
	if v := cluster["interfaces"]; v != nil {
		if total, ok := v.(map[string]interface{})["total"]; ok {
			totalInterfaces := int(total.(float64))
			if totalInterfaces > 50 {
				payload["limit-interfaces"] = totalInterfaces
				showClusterRes, err := client.ApiCall("show-simple-cluster", payload, client.GetSessionID(), true, client.IsProxyUsed())
				if err != nil {
					return fmt.Errorf(err.Error())
				}
				if !showClusterRes.Success {
					return fmt.Errorf(showClusterRes.ErrorMsg)
				}
				cluster = showClusterRes.GetData()
			}
		}
	}

	log.Println("Read Simple Cluster - Show JSON = ", cluster)

	if v := cluster["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := cluster["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := cluster["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := cluster["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if v := cluster["cluster-mode"]; v != nil {
		_ = d.Set("cluster_mode", v)
	}

	if v := cluster["geo-mode"]; v != nil {
		_ = d.Set("geo_mode", v)
	}

	if cluster["advanced-settings"] != nil {

		advancedSettingsMap, ok := cluster["advanced-settings"].(map[string]interface{})

		if ok {
			advancedSettingsMapToReturn := make(map[string]interface{})

			if v := advancedSettingsMap["connection-persistence"]; v != nil {
				advancedSettingsMapToReturn["connection_persistence"] = v
			}
			if v, ok := advancedSettingsMap["sam"]; ok {

				samMap, ok := v.(map[string]interface{})
				if ok {
					samMapToReturn := make(map[string]interface{})

					if v, _ := samMap["forward-to-other-sam-servers"]; v != nil {
						samMapToReturn["forward_to_other_sam_servers"] = v
					}
					if v, _ := samMap["use-early-versions"]; v != nil {
						samMapToReturn["use_early_versions"] = v
					}
					if v, _ := samMap["purge-sam-file"]; v != nil {
						samMapToReturn["purge_sam_file"] = v
					}
					advancedSettingsMapToReturn["sam"] = []interface{}{samMapToReturn}
				}
			}
			_ = d.Set("advanced_settings", []interface{}{advancedSettingsMapToReturn})

		}
	} else {
		_ = d.Set("advanced_settings", nil)
	}

	if v := cluster["enable-https-inspection"]; v != nil {
		_ = d.Set("enable_https_inspection", v)
	}

	if cluster["fetch-policy"] != nil {
		fetchPolicyJson, ok := cluster["fetch-policy"].([]interface{})
		if ok {
			_ = d.Set("fetch_policy", fetchPolicyJson)
		}
	} else {
		_ = d.Set("fetch_policy", nil)
	}

	if v := cluster["hit-count"]; v != nil {
		_ = d.Set("hit_count", v)
	}

	if cluster["https-inspection"] != nil {

		httpsInspectionMap, ok := cluster["https-inspection"].(map[string]interface{})

		if ok {
			httpsInspectionMapToReturn := make(map[string]interface{})

			if v, ok := httpsInspectionMap["bypass-on-failure"]; ok {

				bypassOnFailureMap, ok := v.(map[string]interface{})
				if ok {
					bypassOnFailureMapToReturn := make(map[string]interface{})

					if v, _ := bypassOnFailureMap["override-profile"]; v != nil {
						bypassOnFailureMapToReturn["override_profile"] = v
					}
					if v, _ := bypassOnFailureMap["value"]; v != nil {
						bypassOnFailureMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["bypass_on_failure"] = []interface{}{bypassOnFailureMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["site-categorization-allow-mode"]; ok {

				siteCategorizationAllowModeMap, ok := v.(map[string]interface{})
				if ok {
					siteCategorizationAllowModeMapToReturn := make(map[string]interface{})

					if v, _ := siteCategorizationAllowModeMap["override-profile"]; v != nil {
						siteCategorizationAllowModeMapToReturn["override_profile"] = v
					}
					if v, _ := siteCategorizationAllowModeMap["value"]; v != nil {
						siteCategorizationAllowModeMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["site_categorization_allow_mode"] = []interface{}{siteCategorizationAllowModeMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["deny-untrusted-server-cert"]; ok {

				denyUntrustedServerCertMap, ok := v.(map[string]interface{})
				if ok {
					denyUntrustedServerCertMapToReturn := make(map[string]interface{})

					if v, _ := denyUntrustedServerCertMap["override-profile"]; v != nil {
						denyUntrustedServerCertMapToReturn["override_profile"] = v
					}
					if v, _ := denyUntrustedServerCertMap["value"]; v != nil {
						denyUntrustedServerCertMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["deny_untrusted_server_cert"] = []interface{}{denyUntrustedServerCertMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["deny-revoked-server-cert"]; ok {

				denyRevokedServerCertMap, ok := v.(map[string]interface{})
				if ok {
					denyRevokedServerCertMapToReturn := make(map[string]interface{})

					if v, _ := denyRevokedServerCertMap["override-profile"]; v != nil {
						denyRevokedServerCertMapToReturn["override_profile"] = v
					}
					if v, _ := denyRevokedServerCertMap["value"]; v != nil {
						denyRevokedServerCertMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["deny_revoked_server_cert"] = []interface{}{denyRevokedServerCertMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["deny-expired-server-cert"]; ok {

				denyExpiredServerCertMap, ok := v.(map[string]interface{})
				if ok {
					denyExpiredServerCertMapToReturn := make(map[string]interface{})

					if v, _ := denyExpiredServerCertMap["override-profile"]; v != nil {
						denyExpiredServerCertMapToReturn["override_profile"] = v
					}
					if v, _ := denyExpiredServerCertMap["value"]; v != nil {
						denyExpiredServerCertMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["deny_expired_server_cert"] = []interface{}{denyExpiredServerCertMapToReturn}
				}
			}
			_ = d.Set("https_inspection", []interface{}{httpsInspectionMapToReturn})

		}
	} else {
		_ = d.Set("https_inspection", nil)
	}

	if v := cluster["identity-awareness"]; v != nil {
		_ = d.Set("identity_awareness", v)
	}

	if cluster["identity-awareness-settings"] != nil {

		identityAwarenessSettingsMap, ok := cluster["identity-awareness-settings"].(map[string]interface{})

		if ok {
			identityAwarenessSettingsMapToReturn := make(map[string]interface{})

			if v := identityAwarenessSettingsMap["browser-based-authentication"]; v != nil {
				identityAwarenessSettingsMapToReturn["browser_based_authentication"] = v
			}
			if v, ok := identityAwarenessSettingsMap["browser-based-authentication-settings"]; ok {

				browserBasedAuthenticationSettingsMap, ok := v.(map[string]interface{})
				if ok {
					browserBasedAuthenticationSettingsMapToReturn := make(map[string]interface{})

					if v, _ := browserBasedAuthenticationSettingsMap["authentication-settings"]; v != nil {
						browserBasedAuthenticationSettingsMapToReturn["authentication_settings"] = v
					}
					if v, _ := browserBasedAuthenticationSettingsMap["browser-based-authentication-portal-settings"]; v != nil {
						browserBasedAuthenticationSettingsMapToReturn["browser_based_authentication_portal_settings"] = v
					}
					identityAwarenessSettingsMapToReturn["browser_based_authentication_settings"] = []interface{}{browserBasedAuthenticationSettingsMapToReturn}
				}
			}
			if v := identityAwarenessSettingsMap["identity-agent"]; v != nil {
				identityAwarenessSettingsMapToReturn["identity_agent"] = v
			}
			if v, ok := identityAwarenessSettingsMap["identity-agent-settings"]; ok {

				identityAgentSettingsMap, ok := v.(map[string]interface{})
				if ok {
					identityAgentSettingsMapToReturn := make(map[string]interface{})

					if v, _ := identityAgentSettingsMap["agents-interval-keepalive"]; v != nil {
						identityAgentSettingsMapToReturn["agents_interval_keepalive"] = v
					}
					if v, _ := identityAgentSettingsMap["user-reauthenticate-interval"]; v != nil {
						identityAgentSettingsMapToReturn["user_reauthenticate_interval"] = v
					}
					if v, _ := identityAgentSettingsMap["authentication-settings"]; v != nil {
						identityAgentSettingsMapToReturn["authentication_settings"] = v
					}
					if v, _ := identityAgentSettingsMap["identity-agent-portal-settings"]; v != nil {
						identityAgentSettingsMapToReturn["identity_agent_portal_settings"] = v
					}
					identityAwarenessSettingsMapToReturn["identity_agent_settings"] = []interface{}{identityAgentSettingsMapToReturn}
				}
			}
			if v := identityAwarenessSettingsMap["identity-collector"]; v != nil {
				identityAwarenessSettingsMapToReturn["identity_collector"] = v
			}
			if v, ok := identityAwarenessSettingsMap["identity-collector-settings"]; ok {

				identityCollectorSettingsMap, ok := v.(map[string]interface{})
				if ok {
					identityCollectorSettingsMapToReturn := make(map[string]interface{})

					if v, _ := identityCollectorSettingsMap["authorized-clients"]; v != nil {
						identityCollectorSettingsMapToReturn["authorized_clients"] = v
					}
					if v, _ := identityCollectorSettingsMap["authentication-settings"]; v != nil {
						identityCollectorSettingsMapToReturn["authentication_settings"] = v
					}
					if v, _ := identityCollectorSettingsMap["client-access-permissions"]; v != nil {
						identityCollectorSettingsMapToReturn["client_access_permissions"] = v
					}
					identityAwarenessSettingsMapToReturn["identity_collector_settings"] = []interface{}{identityCollectorSettingsMapToReturn}
				}
			}
			if v, ok := identityAwarenessSettingsMap["identity-sharing-settings"]; ok {

				identitySharingSettingsMap, ok := v.(map[string]interface{})
				if ok {
					identitySharingSettingsMapToReturn := make(map[string]interface{})

					if v, _ := identitySharingSettingsMap["share-with-other-gateways"]; v != nil {
						identitySharingSettingsMapToReturn["share_with_other_gateways"] = v
					}
					if v, _ := identitySharingSettingsMap["receive-from-other-gateways"]; v != nil {
						identitySharingSettingsMapToReturn["receive_from_other_gateways"] = v
					}
					if v, _ := identitySharingSettingsMap["receive-from"]; v != nil {
						identitySharingSettingsMapToReturn["receive_from"] = v
					}
					identityAwarenessSettingsMapToReturn["identity_sharing_settings"] = []interface{}{identitySharingSettingsMapToReturn}
				}
			}
			if v, ok := identityAwarenessSettingsMap["proxy-settings"]; ok {

				proxySettingsMap, ok := v.(map[string]interface{})
				if ok {
					proxySettingsMapToReturn := make(map[string]interface{})

					if v, _ := proxySettingsMap["detect-using-x-forward-for"]; v != nil {
						proxySettingsMapToReturn["detect_using_x_forward_for"] = v
					}
					identityAwarenessSettingsMapToReturn["proxy_settings"] = []interface{}{proxySettingsMapToReturn}
				}
			}
			if v := identityAwarenessSettingsMap["remote-access"]; v != nil {
				identityAwarenessSettingsMapToReturn["remote_access"] = v
			}
			_ = d.Set("identity_awareness_settings", []interface{}{identityAwarenessSettingsMapToReturn})

		}
	} else {
		_ = d.Set("identity_awareness_settings", nil)
	}

	if v := cluster["ips-update-policy"]; v != nil {
		_ = d.Set("ips_update_policy", v)
	}

	if v := cluster["nat-hide-internal-interfaces"]; v != nil {
		_ = d.Set("nat_hide_internal_interfaces", v)
	}

	if cluster["nat-settings"] != nil {

		natSettingsMap := cluster["nat-settings"].(map[string]interface{})

		natSettingsMapToReturn := make(map[string]interface{})

		if v, _ := natSettingsMap["auto-rule"]; v != nil {
			natSettingsMapToReturn["auto_rule"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := natSettingsMap["ipv4-address"]; v != nil && v != "" {
			natSettingsMapToReturn["ipv4_address"] = v
		}
		if v, _ := natSettingsMap["ipv6-address"]; v != nil && v != "" {
			natSettingsMapToReturn["ipv6_address"] = v
		}
		if v, _ := natSettingsMap["hide-behind"]; v != nil {
			natSettingsMapToReturn["hide_behind"] = v
		}
		if v, _ := natSettingsMap["install-on"]; v != nil {
			natSettingsMapToReturn["install_on"] = v
		}
		if v, _ := natSettingsMap["method"]; v != nil {
			natSettingsMapToReturn["method"] = v
		}
		_ = d.Set("nat_settings", natSettingsMapToReturn)
	} else {
		_ = d.Set("nat_settings", nil)
	}

	if cluster["platform-portal-settings"] != nil {

		platformPortalSettingsMap, ok := cluster["platform-portal-settings"].(map[string]interface{})

		if ok {
			platformPortalSettingsMapToReturn := make(map[string]interface{})

			if v, ok := platformPortalSettingsMap["portal-web-settings"]; ok {

				portalWebSettingsMap, ok := v.(map[string]interface{})
				if ok {
					portalWebSettingsMapToReturn := make(map[string]interface{})

					if v, _ := portalWebSettingsMap["aliases"]; v != nil {
						portalWebSettingsMapToReturn["aliases"] = v
					}
					if v, _ := portalWebSettingsMap["main-url"]; v != nil {
						portalWebSettingsMapToReturn["main_url"] = v
					}
					platformPortalSettingsMapToReturn["portal_web_settings"] = []interface{}{portalWebSettingsMapToReturn}
				}
			}
			if v, ok := platformPortalSettingsMap["certificate-settings"]; ok {

				certificateSettingsMap, ok := v.(map[string]interface{})
				if ok {
					certificateSettingsMapToReturn := make(map[string]interface{})

					if v, _ := certificateSettingsMap["base64-certificate"]; v != nil {
						certificateSettingsMapToReturn["base64_certificate"] = v
					}
					if v, _ := certificateSettingsMap["base64-password"]; v != nil {
						certificateSettingsMapToReturn["base64_password"] = v
					}
					platformPortalSettingsMapToReturn["certificate_settings"] = []interface{}{certificateSettingsMapToReturn}
				}
			}
			if v, ok := platformPortalSettingsMap["accessibility"]; ok {

				accessibilityMap, ok := v.(map[string]interface{})
				if ok {
					accessibilityMapToReturn := make(map[string]interface{})

					if v, _ := accessibilityMap["allow-access-from"]; v != nil {
						accessibilityMapToReturn["allow_access_from"] = v
					}
					if v, _ := accessibilityMap["internal-access-settings"]; v != nil {
						accessibilityMapToReturn["internal_access_settings"] = v
					}
					platformPortalSettingsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
				}
			}
			_ = d.Set("platform_portal_settings", []interface{}{platformPortalSettingsMapToReturn})

		}
	} else {
		_ = d.Set("platform_portal_settings", nil)
	}

	if cluster["proxy-settings"] != nil {

		proxySettingsMap := cluster["proxy-settings"].(map[string]interface{})

		proxySettingsMapToReturn := make(map[string]interface{})

		if v, _ := proxySettingsMap["use-custom-proxy"]; v != nil {
			proxySettingsMapToReturn["use_custom_proxy"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := proxySettingsMap["proxy-server"]; v != nil {
			proxySettingsMapToReturn["proxy_server"] = v
		}
		if v, _ := proxySettingsMap["port"]; v != nil {
			proxySettingsMapToReturn["port"] = v
		}
		_ = d.Set("proxy_settings", proxySettingsMapToReturn)
	} else {
		_ = d.Set("proxy_settings", nil)
	}

	if v := cluster["qos"]; v != nil {
		_ = d.Set("qos", v)
	}

	if cluster["usercheck-portal-settings"] != nil {

		usercheckPortalSettingsMap, ok := cluster["usercheck-portal-settings"].(map[string]interface{})

		if ok {
			usercheckPortalSettingsMapToReturn := make(map[string]interface{})

			if v := usercheckPortalSettingsMap["enabled"]; v != nil {
				usercheckPortalSettingsMapToReturn["enabled"] = v
			}
			if v, ok := usercheckPortalSettingsMap["portal-web-settings"]; ok {

				portalWebSettingsMap, ok := v.(map[string]interface{})
				if ok {
					portalWebSettingsMapToReturn := make(map[string]interface{})

					if v, _ := portalWebSettingsMap["aliases"]; v != nil {
						portalWebSettingsMapToReturn["aliases"] = v
					}
					if v, _ := portalWebSettingsMap["main-url"]; v != nil {
						portalWebSettingsMapToReturn["main_url"] = v
					}
					usercheckPortalSettingsMapToReturn["portal_web_settings"] = []interface{}{portalWebSettingsMapToReturn}
				}
			}
			if v, ok := usercheckPortalSettingsMap["certificate-settings"]; ok {

				certificateSettingsMap, ok := v.(map[string]interface{})
				if ok {
					certificateSettingsMapToReturn := make(map[string]interface{})

					if v, _ := certificateSettingsMap["base64-certificate"]; v != nil {
						certificateSettingsMapToReturn["base64_certificate"] = v
					}
					if v, _ := certificateSettingsMap["base64-password"]; v != nil {
						certificateSettingsMapToReturn["base64_password"] = v
					}
					usercheckPortalSettingsMapToReturn["certificate_settings"] = []interface{}{certificateSettingsMapToReturn}
				}
			}
			if v, ok := usercheckPortalSettingsMap["accessibility"]; ok {

				accessibilityMap, ok := v.(map[string]interface{})
				if ok {
					accessibilityMapToReturn := make(map[string]interface{})

					if v, _ := accessibilityMap["allow-access-from"]; v != nil {
						accessibilityMapToReturn["allow_access_from"] = v
					}
					if v, _ := accessibilityMap["internal-access-settings"]; v != nil {
						accessibilityMapToReturn["internal_access_settings"] = v
					}
					usercheckPortalSettingsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
				}
			}
			_ = d.Set("usercheck_portal_settings", []interface{}{usercheckPortalSettingsMapToReturn})

		}
	} else {
		_ = d.Set("usercheck_portal_settings", nil)
	}

	if v := cluster["zero-phishing"]; v != nil {
		_ = d.Set("zero_phishing", v)
	}

	if v := cluster["zero-phishing-fqdn"]; v != nil {
		_ = d.Set("zero_phishing_fqdn", v)
	}

	if v := cluster["interfaces"]; v != nil {
		interfacesList := v.(map[string]interface{})["objects"].([]interface{})
		if len(interfacesList) > 0 {
			var interfacesListState []map[string]interface{}
			for i := range interfacesList {
				interfaceJson := interfacesList[i].(map[string]interface{})
				interfaceState := make(map[string]interface{})
				if v, _ := interfaceJson["name"]; v != nil {
					interfaceState["name"] = v
				}
				if v, _ := interfaceJson["ipv4-address"]; v != nil {
					interfaceState["ipv4_address"] = v
				}
				if v, _ := interfaceJson["ipv4-mask-length"]; v != nil {
					interfaceState["ipv4_mask_length"] = v
				}
				if v, _ := interfaceJson["ipv4-network-mask"]; v != nil {
					interfaceState["ipv4_network_mask"] = v
				}
				if v, _ := interfaceJson["ipv6-address"]; v != nil {
					interfaceState["ipv6_address"] = v
				}
				if v, _ := interfaceJson["ipv6-mask-length"]; v != nil {
					interfaceState["ipv6_mask_length"] = v
				}
				if v, _ := interfaceJson["ipv6-network-mask"]; v != nil {
					interfaceState["ipv6_network_mask"] = v
				}
				if v, _ := interfaceJson["interface-type"]; v != nil {
					interfaceState["interface_type"] = v
				}
				if v, _ := interfaceJson["anti-spoofing"]; v != nil {
					interfaceState["anti_spoofing"] = v
				}
				if v, _ := interfaceJson["anti-spoofing-settings"]; v != nil {
					antiSpoofingSettingsJson := v.(map[string]interface{})
					antiSpoofingSettingsState := make(map[string]interface{})
					if v, _ := antiSpoofingSettingsJson["action"]; v != nil {
						antiSpoofingSettingsState["action"] = v
					}
					interfaceState["anti_spoofing_settings"] = antiSpoofingSettingsState
				}
				if v, _ := interfaceJson["security-zone"]; v != nil {
					interfaceState["security_zone"] = v
				}
				if v, _ := interfaceJson["security-zone-settings"]; v != nil {
					securityZoneSettingsJson := v.(map[string]interface{})
					securityZoneSettingsState := make(map[string]interface{})
					if v, _ := securityZoneSettingsJson["auto-calculated"]; v != nil {
						securityZoneSettingsState["auto_calculated"] = v
					}
					if v, _ := securityZoneSettingsJson["specific-zone"]; v != nil {
						securityZoneSettingsState["specific_zone"] = v
					}
					interfaceState["security_zone_settings"] = securityZoneSettingsState
				}
				if v, _ := interfaceJson["topology"]; v != nil {
					interfaceState["topology"] = v
				}
				if v, _ := interfaceJson["topology-automatic-calculation"]; v != nil {
					interfaceState["topology_automatic_calculation"] = v
				}
				if v, _ := interfaceJson["topology-settings"]; v != nil {
					topologySettingsJson := v.(map[string]interface{})
					topologySettingsState := make(map[string]interface{})
					if v, _ := topologySettingsJson["interface-leads-to-dmz"]; v != nil {
						topologySettingsState["interface_leads_to_dmz"] = v
					}
					if v, _ := topologySettingsJson["ip-address-behind-this-interface"]; v != nil {
						topologySettingsState["ip_address_behind_this_interface"] = v
					}
					if v, _ := topologySettingsJson["specific-network"]; v != nil {
						topologySettingsState["specific_network"] = v
					}
					interfaceState["topology_settings"] = topologySettingsState
				}

				if v, _ := interfaceJson["color"]; v != nil {
					interfaceState["color"] = v
				}
				if v, _ := interfaceJson["comments"]; v != nil {
					interfaceState["comments"] = v
				}
				interfacesListState = append(interfacesListState, interfaceState)
			}
			_ = d.Set("interfaces", interfacesListState)
		} else {
			_ = d.Set("interfaces", interfacesList)
		}
	} else {
		_ = d.Set("interfaces", nil)
	}

	if v := cluster["cluster-members"]; v != nil {
		membersList := v.([]interface{})
		if len(membersList) > 0 {
			var membersListState []map[string]interface{}
			for i := range membersList {
				memberJson := membersList[i].(map[string]interface{})
				memberState := make(map[string]interface{})
				if v, _ := memberJson["name"]; v != nil {
					memberState["name"] = v
				}
				if v, _ := memberJson["ip-address"]; v != nil {
					memberState["ip_address"] = v
				}
				if v, _ := memberJson["interfaces"]; v != nil {
					memberInterfacesList := v.([]interface{})
					if len(memberInterfacesList) > 0 {
						var memberInterfacesState []map[string]interface{}
						for i := range memberInterfacesList {
							memberInterfaceJson := memberInterfacesList[i].(map[string]interface{})
							memberInterfaceState := make(map[string]interface{})
							if v, _ := memberInterfaceJson["name"]; v != nil {
								memberInterfaceState["name"] = v
							}
							if v, _ := memberInterfaceJson["ipv4-address"]; v != nil {
								memberInterfaceState["ipv4_address"] = v
							}
							if v, _ := memberInterfaceJson["ipv4-mask-length"]; v != nil {
								memberInterfaceState["ipv4_mask_length"] = v
							}
							if v, _ := memberInterfaceJson["ipv4-network-mask"]; v != nil {
								memberInterfaceState["ipv4_network_mask"] = v
							}
							if v, _ := memberInterfaceJson["ipv6-address"]; v != nil {
								memberInterfaceState["ipv6_address"] = v
							}
							if v, _ := memberInterfaceJson["ipv6-mask-length"]; v != nil {
								memberInterfaceState["ipv6_mask_length"] = v
							}
							if v, _ := memberInterfaceJson["ipv6-network-mask"]; v != nil {
								memberInterfaceState["ipv6_network_mask"] = v
							}
							memberInterfacesState = append(memberInterfacesState, memberInterfaceState)
						}
						memberState["interfaces"] = memberInterfacesState
					}
				}

				if v, _ := memberJson["sic-message"]; v != nil {
					memberState["sic_message"] = v
				}
				if v, _ := memberJson["sic-state"]; v != nil {
					memberState["sic_state"] = v
				}
				membersListState = append(membersListState, memberState)
			}
			_ = d.Set("members", membersListState)
		} else {
			_ = d.Set("members", membersList)
		}
	} else {
		_ = d.Set("members", nil)
	}

	if v := cluster["anti-bot"]; v != nil {
		_ = d.Set("anti_bot", v)
	}

	if v := cluster["anti-virus"]; v != nil {
		_ = d.Set("anti_virus", v)
	}

	if v := cluster["application-control"]; v != nil {
		_ = d.Set("application_control", v)
	}

	if v := cluster["content-awareness"]; v != nil {
		_ = d.Set("content_awareness", v)
	}

	if v := cluster["dynamic-ip"]; v != nil {
		_ = d.Set("dynamic_ip", v)
	}

	if v := cluster["firewall"]; v != nil {
		_ = d.Set("firewall", v)
	}

	if v := cluster["ips"]; v != nil {
		_ = d.Set("ips", v)
	}

	if v := cluster["threat-emulation"]; v != nil {
		_ = d.Set("threat_emulation", v)
	}

	if v := cluster["url-filtering"]; v != nil {
		_ = d.Set("url_filtering", v)
	}

	if v := cluster["vpn"]; v != nil {
		_ = d.Set("vpn", v)
	}

	if v := cluster["os-name"]; v != nil {
		_ = d.Set("os_name", v)
	}

	if v := cluster["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := cluster["hardware"]; v != nil {
		_ = d.Set("hardware", v)
	}

	if v := cluster["sic-name"]; v != nil {
		_ = d.Set("sic_name", v)
	}

	if v := cluster["sic-state"]; v != nil {
		_ = d.Set("sic_state", v)
	}

	if v := cluster["save-logs-locally"]; v != nil {
		_ = d.Set("save_logs_locally", v)
	}

	if v := cluster["send_alerts_to_server"]; v != nil {
		_ = d.Set("send_alerts_to_server", v)
	} else {
		_ = d.Set("send_alerts_to_server", nil)
	}

	if v := cluster["send-logs-to-backup-server"]; v != nil {
		_ = d.Set("send_logs_to_backup_server", v)
	} else {
		_ = d.Set("send_logs_to_backup_server", nil)
	}

	if v := cluster["send-logs-to-server"]; v != nil {
		_ = d.Set("send_logs_to_server", v)
	} else {
		_ = d.Set("send_logs_to_server", nil)
	}

	if v := cluster["logs-settings"]; v != nil {
		logSettingsJson := v.(map[string]interface{})
		logSettingsState := make(map[string]interface{})
		if v := logSettingsJson["alert-when-free-disk-space-below"]; v != nil {
			logSettingsState["alert_when_free_disk_space_below"] = v
		}
		if v := logSettingsJson["alert-when-free-disk-space-below-metrics"]; v != nil {
			logSettingsState["alert_when_free_disk_space_below_metrics"] = v
		}
		if v := logSettingsJson["alert_when_free_disk_space_below_threshold"]; v != nil {
			logSettingsState["alert_when_free_disk_space_below_threshold"] = v
		}
		if v := logSettingsJson["alert-when-free-disk-space-below-type"]; v != nil {
			logSettingsState["alert_when_free_disk_space_below_type"] = v
		}
		if v := logSettingsJson["before-delete-keep-logs-from-the-last-days"]; v != nil {
			logSettingsState["before_delete_keep_logs_from_the_last_days"] = v
		}
		if v := logSettingsJson["before-delete-keep-logs-from-the-last-days-threshold"]; v != nil {
			logSettingsState["before_delete_keep_logs_from_the_last_days_threshold"] = v
		}
		if v := logSettingsJson["before-delete-run-script"]; v != nil {
			logSettingsState["before_delete_run_script"] = v
		}
		if v := logSettingsJson["before-delete-run-script-command"]; v != nil {
			logSettingsState["before_delete_run_script_command"] = v
		}
		if v := logSettingsJson["delete-index-files-older-than-days"]; v != nil {
			logSettingsState["delete_index_files_older_than_days"] = v
		}
		if v := logSettingsJson["delete-index-files-older-than-days-threshold"]; v != nil {
			logSettingsState["delete_index_files_older_than_days_threshold"] = v
		}
		if v := logSettingsJson["delete-index-files-when-index-size-above"]; v != nil {
			logSettingsState["delete_index_files_when_index_size_above"] = v
		}
		if v := logSettingsJson["delete-index-files-when-index-size-above-metrics"]; v != nil {
			logSettingsState["delete_index_files_when_index_size_above_metrics"] = v
		}
		if v := logSettingsJson["delete-index-files-when-index-size-above-threshold"]; v != nil {
			logSettingsState["delete_index_files_when_index_size_above_threshold"] = v
		}
		if v := logSettingsJson["delete-when-free-disk-space-below"]; v != nil {
			logSettingsState["delete_when_free_disk_space_below"] = v
		}
		if v := logSettingsJson["delete-when-free-disk-space-below-metrics"]; v != nil {
			logSettingsState["delete_when_free_disk_space_below_metrics"] = v
		}
		if v := logSettingsJson["delete-when-free-disk-space-below-threshold"]; v != nil {
			logSettingsState["delete_when_free_disk_space_below_threshold"] = v
		}
		if v := logSettingsJson["detect-new-citrix-ica-application-names"]; v != nil {
			logSettingsState["detect_new_citrix_ica_application_names"] = v
		}
		if v := logSettingsJson["forward-logs-to-log-server"]; v != nil {
			logSettingsState["forward_logs_to_log_server"] = v
		}
		if v := logSettingsJson["forward-logs-to-log-server-name"]; v != nil {
			logSettingsState["forward_logs_to_log_server_name"] = v
		}
		if v := logSettingsJson["forward-logs-to-log-server-schedule-name"]; v != nil {
			logSettingsState["forward_logs_to_log_server_schedule_name"] = v
		}
		if v := logSettingsJson["perform-log-rotate-before-log-forwarding"]; v != nil {
			logSettingsState["perform_log_rotate_before_log_forwarding"] = v
		}
		if v := logSettingsJson["reject-connections-when-free-disk-space-below-threshold"]; v != nil {
			logSettingsState["reject_connections_when_free_disk_space_below_threshold"] = v
		}
		if v := logSettingsJson["reserve-for-packet-capture-metrics"]; v != nil {
			logSettingsState["reserve_for_packet_capture_metrics"] = v
		}
		if v := logSettingsJson["reserve-for-packet-capture-threshold"]; v != nil {
			logSettingsState["reserve_for_packet_capture_threshold"] = v
		}
		if v := logSettingsJson["rotate-log-by-file-size"]; v != nil {
			logSettingsState["rotate_log_by_file_size"] = v
		}
		if v := logSettingsJson["rotate-log-file-size-threshold"]; v != nil {
			logSettingsState["rotate_log_file_size_threshold"] = v
		}
		if v := logSettingsJson["rotate-log-on-schedule"]; v != nil {
			logSettingsState["rotate_log_on_schedule"] = v
		}
		if v := logSettingsJson["rotate-log-schedule-name"]; v != nil {
			logSettingsState["rotate_log_schedule_name"] = v
		}
		if v := logSettingsJson["stop-logging-when-free-disk-space-below"]; v != nil {
			logSettingsState["stop_logging_when_free_disk_space_below"] = v
		}
		if v := logSettingsJson["stop-logging-when-free-disk-space-below-metrics"]; v != nil {
			logSettingsState["stop_logging_when_free_disk_space_below_metrics"] = v
		}
		if v := logSettingsJson["stop-logging-when-free-disk-space-below-threshold"]; v != nil {
			logSettingsState["stop_logging_when_free_disk_space_below_threshold"] = v
		}
		if v := logSettingsJson["turn-on-qos-logging"]; v != nil {
			logSettingsState["turn_on_qos_logging"] = v
		}
		if v := logSettingsJson["update-account-log-every"]; v != nil {
			logSettingsState["update_account_log_every"] = v
		}
		_ = d.Set("logs_settings", logSettingsState)
	} else {
		_ = d.Set("logs_settings", nil)
	}

	if v := cluster["firewall-settings"]; v != nil {
		firewallSettingsJson := v.(map[string]interface{})
		firewallSettingsState := make(map[string]interface{})
		if v := firewallSettingsJson["auto-calculate-connections-hash-table-size-and-memory-pool"]; v != nil {
			firewallSettingsState["auto_calculate_connections_hash_table_size_and_memory_pool"] = v
		}
		if v := firewallSettingsJson["auto-maximum-limit-for-concurrent-connections"]; v != nil {
			firewallSettingsState["auto_maximum_limit_for_concurrent_connections"] = v
		}
		if v := firewallSettingsJson["connections-hash-size"]; v != nil {
			firewallSettingsState["connections_hash_size"] = v
		}
		if v := firewallSettingsJson["maximum-limit-for-concurrent-connections"]; v != nil {
			firewallSettingsState["maximum_limit_for_concurrent_connections"] = v
		}
		if v := firewallSettingsJson["maximum-memory-pool-size"]; v != nil {
			firewallSettingsState["maximum_memory_pool_size"] = v
		}
		if v := firewallSettingsJson["memory-pool-size"]; v != nil {
			firewallSettingsState["memory_pool_size"] = v
		}
		_ = d.Set("firewall_settings", firewallSettingsState)
	} else {
		_ = d.Set("firewall_settings", nil)
	}

	if v := cluster["vpn-settings"]; v != nil {
		vpnSettingsJson := v.(map[string]interface{})
		vpnSettingsState := make(map[string]interface{})
		if v := vpnSettingsJson["authentication"]; v != nil {
			authenticationJson := v.(map[string]interface{})
			authenticationState := make(map[string]interface{})
			if v := authenticationJson["authentication-clients"]; v != nil {
				clientsJson := v.([]interface{})
				var clientsIds = make([]string, 0)
				if len(clientsJson) > 0 {
					for _, client := range clientsJson {
						clientsIds = append(clientsIds, client.(map[string]interface{})["name"].(string))
					}
				}
				authenticationState["authentication_clients"] = clientsIds
			}
			vpnSettingsState["authentication"] = authenticationState
		}

		if v := vpnSettingsJson["link-selection"]; v != nil {
			linkSelectionJson := v.(map[string]interface{})
			linkSelectionState := make(map[string]interface{})
			if v := linkSelectionJson["ip-selection"]; v != nil {
				linkSelectionState["ip_selection"] = v
			}
			if v := linkSelectionJson["dns-resolving-hostname"]; v != nil {
				linkSelectionState["dns_resolving_hostname"] = v
			}
			if v := linkSelectionJson["ip-address"]; v != nil {
				linkSelectionState["ip_address"] = v
			}
			vpnSettingsState["link_selection"] = linkSelectionState
		}
		if v := vpnSettingsJson["maximum-concurrent-ike-negotiations"]; v != nil {
			vpnSettingsState["maximum_concurrent_ike_negotiations"] = v
		}
		if v := vpnSettingsJson["maximum-concurrent-tunnels"]; v != nil {
			vpnSettingsState["maximum_concurrent_tunnels"] = v
		}
		if v := vpnSettingsJson["vpn-domain-type"]; v != nil {
			vpnSettingsState["vpn_domain_type"] = v
		}
		if v := vpnSettingsJson["vpn-domain"]; v != nil {
			vpnSettingsState["vpn_domain"] = v.(map[string]interface{})["name"]
		}
		if v := vpnSettingsJson["remote-access"]; v != nil {
			remoteAccessJson := v.(map[string]interface{})
			remoteAccessState := make(map[string]interface{})
			if v := remoteAccessJson["support-l2tp"]; v != nil {
				remoteAccessState["support_l2tp"] = v
			}
			if v := remoteAccessJson["l2tp-auth-method"]; v != nil {
				remoteAccessState["l2tp_auth_method"] = v
			}
			if v := remoteAccessJson["l2tp-certificate"]; v != nil {
				remoteAccessState["l2tp_certificate"] = v
			}
			if v := remoteAccessJson["allow-vpn-clients-to-route-traffic"]; v != nil {
				remoteAccessState["allow_vpn_clients_to_route_traffic"] = v
			}
			if v := remoteAccessJson["support-nat-traversal-mechanism"]; v != nil {
				remoteAccessState["support_nat_traversal_mechanism"] = v
			}
			if v := remoteAccessJson["nat-traversal-service"]; v != nil {
				remoteAccessState["nat_traversal_service"] = v.(map[string]interface{})["name"]
			}
			if v := remoteAccessJson["support-visitor-mode"]; v != nil {
				remoteAccessState["support_visitor_mode"] = v
			}
			if v := remoteAccessJson["visitor-mode-service"]; v != nil {
				remoteAccessState["visitor_mode_service"] = v.(map[string]interface{})["name"]
			}
			if v := remoteAccessJson["visitor-mode-interface"]; v != nil {
				remoteAccessState["visitor_mode_interface"] = v
			}
			vpnSettingsState["remote_access"] = remoteAccessState
		}

		if v := vpnSettingsJson["office-mode"]; v != nil {
			officeModeJson := v.(map[string]interface{})
			officeModeState := make(map[string]interface{})
			if v := officeModeJson["mode"]; v != nil {
				officeModeState["mode"] = v
			}
			if v := officeModeJson["group"]; v != nil {
				officeModeState["group"] = v.(map[string]interface{})["name"]
			}
			if v := officeModeJson["support-multiple-interfaces"]; v != nil {
				officeModeState["support_multiple_interfaces"] = v
			}
			if v := officeModeJson["perform-anti-spoofing"]; v != nil {
				officeModeState["perform_anti_spoofing"] = v
			}
			if v := officeModeJson["anti-spoofing-additional-addresses"]; v != nil {
				officeModeState["anti_spoofing_additional_addresses"] = v.(map[string]interface{})["name"]
			}
			if v := officeModeJson["allocate-ip-address-from"]; v != nil {
				allocateIpAddressFromJson := v.(map[string]interface{})
				allocateIpAddressFromState := make(map[string]interface{})
				if v := allocateIpAddressFromJson["radius-server"]; v != nil {
					allocateIpAddressFromState["radius_server"] = v
				}
				if v := allocateIpAddressFromJson["use-allocate-method"]; v != nil {
					allocateIpAddressFromState["use_allocate_method"] = v
				}
				if v := allocateIpAddressFromJson["allocate-method"]; v != nil {
					allocateIpAddressFromState["allocate_method"] = v
				}
				if v := allocateIpAddressFromJson["manual-network"]; v != nil {
					allocateIpAddressFromState["manual_network"] = v.(map[string]interface{})["name"]
				}
				if v := allocateIpAddressFromJson["dhcp-server"]; v != nil {
					allocateIpAddressFromState["dhcp_server"] = v.(map[string]interface{})["name"]
				}
				if v := allocateIpAddressFromJson["virtual-ip-address"]; v != nil {
					allocateIpAddressFromState["virtual_ip_address"] = v
				}
				if v := allocateIpAddressFromJson["dhcp-mac-address"]; v != nil {
					allocateIpAddressFromState["dhcp_mac_address"] = v
				}
				if v := allocateIpAddressFromJson["optional-parameters"]; v != nil {
					optionalParametersJson := v.(map[string]interface{})
					optionalParametersState := make(map[string]interface{})
					if v := optionalParametersJson["use-primary-dns-server"]; v != nil {
						optionalParametersState["use_primary_dns_server"] = v
					}
					if v := optionalParametersJson["primary-dns-server"]; v != nil {
						optionalParametersState["primary-dns-server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["use-first-backup-dns-server"]; v != nil {
						optionalParametersState["use_first_backup_dns_server"] = v
					}
					if v := optionalParametersJson["first-backup-dns-server"]; v != nil {
						optionalParametersState["first_backup_dns_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["use-second-backup-dns-server"]; v != nil {
						optionalParametersState["use_second_backup_dns_server"] = v
					}
					if v := optionalParametersJson["second-backup-dns-server"]; v != nil {
						optionalParametersState["second_backup_dns_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["dns-suffixes"]; v != nil {
						optionalParametersState["dns_suffixes"] = v
					}
					if v := optionalParametersJson["use-primary-wins-server"]; v != nil {
						optionalParametersState["use_primary_wins_server"] = v
					}
					if v := optionalParametersJson["primary-wins-server"]; v != nil {
						optionalParametersState["primary_wins_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["use-first-backup-wins-server"]; v != nil {
						optionalParametersState["use_first_backup_wins_server"] = v
					}
					if v := optionalParametersJson["first-backup-wins-server"]; v != nil {
						optionalParametersState["first_backup_wins_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["use-second-backup-wins-server"]; v != nil {
						optionalParametersState["use_second_backup_wins_server"] = v
					}
					if v := optionalParametersJson["second-backup-wins-server"]; v != nil {
						optionalParametersState["second_backup_wins_server"] = v.(map[string]interface{})["name"]
					}
					if v := optionalParametersJson["ip-lease-duration"]; v != nil {
						optionalParametersState["ip_lease_duration"] = v
					}
					allocateIpAddressFromState["optional_parameters"] = optionalParametersState
				}
				officeModeState["allocate_ip_address_from"] = allocateIpAddressFromState
			}
			vpnSettingsState["office_mode"] = officeModeState
		}
		_ = d.Set("vpn-settings", vpnSettingsState)
	} else {
		_ = d.Set("vpn-settings", nil)
	}

	if v := cluster["tags"]; v != nil {
		tagsJson := v.([]interface{})
		var tagsIds = make([]string, 0)
		if len(tagsJson) > 0 {
			for _, tag := range tagsJson {
				tagsIds = append(tagsIds, tag.(map[string]interface{})["name"].(string))
			}
		}
		_ = d.Set("tags", tagsIds)
	} else {
		_ = d.Set("tags", nil)
	}

	if v := cluster["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := cluster["color"]; v != nil {
		_ = d.Set("color", v)
	}

	return nil
}
