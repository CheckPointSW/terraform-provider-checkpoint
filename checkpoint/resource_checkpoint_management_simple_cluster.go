package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementSimpleCluster() *schema.Resource {
	return &schema.Resource{
		Create: createManagementSimpleCluster,
		Read:   readManagementSimpleCluster,
		Update: updateManagementSimpleCluster,
		Delete: deleteManagementSimpleCluster,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name. Should be unique in the domain.",
			},
			"ipv4_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv4 address.",
			},
			"ipv6_address": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "IPv6 address.",
			},
			"cluster_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster mode.",
				Default:     "cluster-xl-ha",
			},
			"geo_mode": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Cluster High Availability Geo mode. This setting applies only to a cluster deployed in a cloud. Available when the cluster mode equals \"cluster-xl-ha\".",
			},
			"advanced_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "N/A",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_persistence": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Handling established connections when installing a new policy.",
							Default:     "rematch-connections",
						},
						"sam": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "SAM.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"forward_to_other_sam_servers": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Forward SAM clients' requests to other SAM servers.",
										Default:     false,
									},
									"use_early_versions": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Use early versions compatibility mode.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Use early versions compatibility mode.",
													Default:     false,
												},
												"compatibility_mode": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Early versions compatibility mode.",
													Default:     "auth_opsec",
												},
											},
										},
									},
									"purge_sam_file": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Purge SAM File.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Purge SAM File.",
													Default:     false,
												},
												"purge_when_size_reaches_to": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Purge SAM File When it Reaches to.",
													Default:     100,
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
				Optional:    true,
				Description: "Enable HTTPS Inspection after defining an outbound inspection certificate. <br>To define the outbound certificate use outbound inspection certificate API.",
			},
			"fetch_policy": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Security management server(s) to fetch the policy from.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hit_count": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Hit count tracks the number of connections each rule matches.",
			},
			"https_inspection": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "HTTPS inspection.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_on_failure": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Set to be true in order to bypass all requests (Fail-open) in case of internal system error.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"site_categorization_allow_mode": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Set to 'background' in order to allowed requests until categorization is complete.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"deny_untrusted_server_cert": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Set to be true in order to drop traffic from servers with untrusted server certificate.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"deny_revoked_server_cert": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Override value.<br><font color=\"red\">Required only for</font> 'override-profile' is True.",
									},
								},
							},
						},
						"deny_expired_server_cert": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Set to be true in order to drop traffic from servers with expired server certificate.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Override profile of global configuration.",
									},
									"value": {
										Type:        schema.TypeBool,
										Optional:    true,
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
				Optional:    true,
				Description: "Identity awareness blade enabled.",
			},
			"identity_awareness_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Gateway Identity Awareness settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"browser_based_authentication": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Browser Based Authentication source.",
						},
						"browser_based_authentication_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Browser Based Authentication settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_settings": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Authentication Settings for Browser Based Authentication.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_method": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Authentication method.",
													Default:     "username and password",
												},
												"identity_provider": {
													Type:        schema.TypeSet,
													Optional:    true,
													Description: "Identity provider object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"identity provider\".",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"radius": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Radius server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"radius\".",
												},
												"users_directories": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Users directories.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"external_user_profile": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "External user profile.",
																Default:     true,
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Internal users.",
																Default:     true,
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Users from external directories.",
																Default:     "all gateways directories",
															},
															"specific": {
																Type:        schema.TypeSet,
																Optional:    true,
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
										Optional:    true,
										Description: "Browser Based Authentication portal settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"portal_web_settings": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Configuration of the portal web settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"aliases": {
																Type:        schema.TypeSet,
																Optional:    true,
																Description: "List of URL aliases that are redirected to the main portal URL.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"main_url": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The main URL for the web portal.",
															},
														},
													},
												},
												"certificate_settings": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Configuration of the portal certificate settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"base64_certificate": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "The certificate file encoded in Base64 with padding.  This file must be in the *.p12 format.",
															},
															"base64_password": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Password (encoded in Base64 with padding) for the certificate file.",
															},
														},
													},
												},
												"accessibility": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Configuration of the portal access settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"vpn": {
																			Type:        schema.TypeBool,
																			Optional:    true,
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
							Optional:    true,
							Description: "Enable Identity Agent source.",
						},
						"identity_agent_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Identity Agent settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"agents_interval_keepalive": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Agents send keepalive period (minutes).",
										Default:     5,
									},
									"user_reauthenticate_interval": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Agent reauthenticate time interval (minutes).",
										Default:     480,
									},
									"authentication_settings": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Authentication Settings for Identity Agent.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"authentication_method": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Authentication method.",
													Default:     "username and password",
												},
												"radius": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Radius server object identified by the name or UID. Must be set when \"authentication-method\" was selected to be \"radius\".",
												},
												"users_directories": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Users directories.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"external_user_profile": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "External user profile.",
																Default:     true,
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Internal users.",
																Default:     true,
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Users from external directories.",
																Default:     "all gateways directories",
															},
															"specific": {
																Type:        schema.TypeSet,
																Optional:    true,
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
										Optional:    true,
										Description: "Identity Agent accessibility settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Configuration of the portal access settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"vpn": {
																			Type:        schema.TypeBool,
																			Optional:    true,
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
							Optional:    true,
							Description: "Enable Identity Collector source.",
						},
						"identity_collector_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Identity Collector settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authorized_clients": {
										Type:        schema.TypeList,
										Required:    true,
										Description: "Authorized Clients.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"client": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Host / Network Group Name or UID.",
												},
												"client_secret": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Client Secret.",
												},
											},
										},
									},
									"authentication_settings": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Authentication Settings for Identity Collector.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"users_directories": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Users directories.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"external_user_profile": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "External user profile.",
																Default:     true,
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Internal users.",
																Default:     true,
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Users from external directories.",
																Default:     "all gateways directories",
															},
															"specific": {
																Type:        schema.TypeSet,
																Optional:    true,
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
										Optional:    true,
										Description: "Identity Collector accessibility settings.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													Optional:    true,
													Description: "Configuration of the portal access settings.",
													MaxItems:    1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																Optional:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																MaxItems:    1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"vpn": {
																			Type:        schema.TypeBool,
																			Optional:    true,
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
							Optional:    true,
							Description: "Identity sharing settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"share_with_other_gateways": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable identity sharing with other gateways.",
									},
									"receive_from_other_gateways": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable receiving identity from other gateways.",
									},
									"receive_from": {
										Type:        schema.TypeSet,
										Optional:    true,
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
							Optional:    true,
							Description: "Identity-Awareness Proxy settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"detect_using_x_forward_for": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to use X-Forward-For HTTP header, which is added by the proxy server to keep track of the original source IP.",
										Default:     false,
									},
								},
							},
						},
						"remote_access": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Remote Access Identity source.",
						},
					},
				},
			},
			"ips_update_policy": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Specifies whether the IPS will be downloaded from the Management or directly to the Gateway.",
			},
			"nat_hide_internal_interfaces": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Hide internal networks behind the Gateway's external IP.",
			},
			"nat_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "NAT settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_rule": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to add automatic address translation rules.",
							Default:     false,
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 address.",
						},
						"hide_behind": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Hide behind method. This parameter is forbidden in case \"method\" parameter is \"static\".",
						},
						"install_on": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Which gateway should apply the NAT translation.",
						},
						"method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "NAT translation method.",
						},
					},
				},
			},
			"platform_portal_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Platform portal settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"portal_web_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Configuration of the portal web settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aliases": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "List of URL aliases that are redirected to the main portal URL.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"main_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The main URL for the web portal.",
									},
								},
							},
						},
						"certificate_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Configuration of the portal certificate settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"base64_certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The certificate file encoded in Base64 with padding.  This file must be in the *.p12 format.",
									},
									"base64_password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Password (encoded in Base64 with padding) for the certificate file.",
									},
								},
							},
						},
						"accessibility": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Configuration of the portal access settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_access_from": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Allowed access to the web portal (based on interfaces, or security policy).",
									},
									"internal_access_settings": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Configuration of the additional portal access settings for internal interfaces only.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"undefined": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
												},
												"dmz": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
												},
												"vpn": {
													Type:        schema.TypeBool,
													Optional:    true,
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
				Optional:    true,
				Description: "Proxy Server for Gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_custom_proxy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Use custom proxy settings for this network object.",
							//Default:     false,
						},
						"proxy_server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "N/A",
						},
						"port": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "N/A",
							//Default:     80,
						},
					},
				},
			},
			"qos": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "QoS.",
			},
			"usercheck_portal_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "UserCheck portal settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "State of the web portal (enabled or disabled). The supported blades are: {'Application Control', 'URL Filtering', 'Data Loss Prevention', 'Anti Virus', 'Anti Bot', 'Threat Emulation', 'Threat Extraction', 'Data Awareness'}.",
						},
						"portal_web_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Configuration of the portal web settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"aliases": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "List of URL aliases that are redirected to the main portal URL.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"main_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The main URL for the web portal.",
									},
								},
							},
						},
						"certificate_settings": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Configuration of the portal certificate settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"base64_certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The certificate file encoded in Base64 with padding.  This file must be in the *.p12 format.",
									},
									"base64_password": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Password (encoded in Base64 with padding) for the certificate file.",
									},
								},
							},
						},
						"accessibility": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Configuration of the portal access settings.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allow_access_from": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Allowed access to the web portal (based on interfaces, or security policy).",
									},
									"internal_access_settings": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Configuration of the additional portal access settings for internal interfaces only.",
										MaxItems:    1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"undefined": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
												},
												"dmz": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
												},
												"vpn": {
													Type:        schema.TypeBool,
													Optional:    true,
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
				Optional:    true,
				Description: "Zero Phishing blade enabled.",
			},
			"zero_phishing_fqdn": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Zero Phishing gateway FQDN.",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"interface_type": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Cluster interface type.",
						},
						"ipv4_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 address.",
						},
						"ipv6_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 address.",
						},
						"ipv4_network_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 network address.",
						},
						"ipv6_network_mask": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 network address.",
						},
						"ipv4_mask_length": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 network mask length.",
						},
						"ipv6_mask_length": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv6 network mask length.",
						},
						"anti_spoofing": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Anti spoofing.",
							Default:     true,
						},
						"anti_spoofing_settings": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Anti spoofing settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).",
									},
								},
							},
						},
						"multicast_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Multicast IP Address.",
						},
						"multicast_address_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Multicast Address Type.",
						},
						"security_zone": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Security zone.",
							Default:     false,
						},
						"security_zone_settings": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Security zone settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_calculated": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Security Zone is calculated according to where the interface leads to.",
									},
									"specific_zone": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Security Zone specified manually.",
									},
								},
							},
						},
						"topology": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Topology.",
							Default:     "automatic",
						},
						"topology_settings": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Topology settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_leads_to_dmz": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether this interface leads to demilitarized zone (perimeter network).",
									},
									"ip_address_behind_this_interface": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Ip address behind this interface.",
									},
									"specific_network": {
										Type:        schema.TypeString,
										Optional:    true,
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
							Optional:    true,
							Default:     "black",
							Description: "Color of the object. Should be one of existing colors.",
						},
						"comments": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Comments string.",
						},
					},
				},
			},
			"members": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cluster members.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Object name. Should be unique in the domain.",
						},
						"ip_address": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "IPv4 or IPv6 address.",
						},
						"one_time_password": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "SIC one time password.",
						},
						"priority": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The member priority on the cluster.",
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
							Optional:    true,
							Description: "Network interfaces.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Object name. Should be unique in the domain.",
									},
									"ipv4_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv4 address.",
									},
									"ipv6_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv6 address.",
									},
									"ipv4_network_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv4 network address.",
									},
									"ipv6_network_mask": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv6 network address.",
									},
									"ipv4_mask_length": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IPv4 network mask length.",
									},
									"ipv6_mask_length": {
										Type:        schema.TypeString,
										Optional:    true,
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
				Optional:    true,
				Description: "Anti-Bot blade enabled.",
			},
			"anti_virus": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Anti-Virus blade enabled.",
			},
			"application_control": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Application Control blade enabled.",
			},
			"content_awareness": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Content Awareness blade enabled.",
			},
			"data_awareness": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Data Awareness blade enabled.",
			},
			"firewall": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Firewall blade enabled.",
			},
			"firewall_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Firewall settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_calculate_connections_hash_table_size_and_memory_pool": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Auto calculate connections hash table size and memory pool.",
						},
						"auto_maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Auto maximum limit for concurrent connections.",
						},
						"connections_hash_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Connections hash size.",
						},
						"maximum_limit_for_concurrent_connections": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum limit for concurrent connections.",
						},
						"maximum_memory_pool_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum memory pool size.",
						},
						"memory_pool_size": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Memory pool size.",
						},
					},
				},
			},
			"ips": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Intrusion Prevention System blade enabled.",
			},
			"ips_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Cluster IPS settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_all_under_load": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Disable/enable all IPS protections until CPU and memory levels are back to normal.",
						},
						"bypass_track_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Track options when all IPS protections are disabled until CPU/memory levels are back to normal.",
						},
						"top_cpu_consuming_protections": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Provides a way to reduce CPU levels on machines under load by disabling the top CPU consuming IPS protections.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"disable_period": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Duration (in hours) for disabling the protections.",
									},
									"disable_under_load": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Temporarily disable/enable top CPU consuming IPS protections.",
									},
								},
							},
						},
						"activation_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Defines whether the IPS blade operates in Detect Only mode or enforces the configured IPS Policy.",
						},
						"cpu_usage_low_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "CPU usage low threshold percentage (1-99).",
						},
						"cpu_usage_high_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "CPU usage high threshold percentage (1-99).",
						},
						"memory_usage_low_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Memory usage low threshold percentage (1-99).",
						},
						"memory_usage_high_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Memory usage high threshold percentage (1-99).",
						},
						"send_threat_cloud_info": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Help improve Check Point Threat Prevention product by sending anonymous information.",
						},
						"reject_on_cluster_fail_over": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Define the IPS connections during fail over reject packets or accept packets.",
						},
					},
				},
			},
			"threat_emulation": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Threat Emulation blade enabled.",
			},
			"url_filtering": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "URL Filtering blade enabled.",
			},
			"dynamic_ip": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Dynamic IP address.",
			},
			"os_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "OS name.",
				Default:     "Gaia",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Cluster platform version.",
			},
			"hardware": {
				Type:        schema.TypeString,
				Optional:    true,
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
				Optional:    true,
				Description: "Save logs locally.",
			},
			"send_alerts_to_server": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Server(s) to send alerts to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_backup_server": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Backup server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"send_logs_to_server": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Server(s) to send logs to.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpn": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "VPN blade enabled.",
			},
			"vpn_settings": {
				Type:        schema.TypeMap,
				Optional:    true,
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
										Optional:    true,
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
							Optional:    true,
							Description: "Link Selection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_selection": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IP selection",
										Default:     "use-main-address",
									},
									"dns_resolving_hostname": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "DNS Resolving Hostname. Must be set when \"ip-selection\" was selected to be \"dns-resolving-from-hostname\".",
									},
									"ip_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "IP Address. Must be set when \"ip-selection\" was selected to be \"use-selected-address-from-topology\" or \"use-statically-nated-ip\"",
									},
								},
							},
						},
						"maximum_concurrent_ike_negotiations": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum concurrent ike negotiations",
						},
						"maximum_concurrent_tunnels": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Maximum concurrent tunnels",
						},
						"office_mode": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Office Mode. Notation Wide Impact - Office Mode apply IPSec VPN Software Blade clients and to the Mobile Access Software Blade clients.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Office Mode Permissions. When selected to be \"off\", all the other definitions are irrelevant.",
										Default:     "off",
									},
									"group": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Group. Identified by name or UID. Must be set when \"office-mode-permissions\" was selected to be \"group\".",
									},
									"allocate_ip_address_from": {
										Type:        schema.TypeMap,
										Optional:    true,
										Description: "Allocate IP address Method. Allocate IP address by sequentially trying the given methods until success.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"radius_server": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Radius server used to authenticate the user.",
													Default:     false,
												},
												"use_allocate_method": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Use Allocate Method.",
													Default:     true,
												},
												"allocate_method": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Using either Manual (IP Pool) or Automatic (DHCP). Must be set when \"use-allocate-method\" is true.",
													Default:     "manual",
												},
												"manual_network": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Manual Network. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"manual\".",
												},
												"dhcp_server": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "DHCP Server. Identified by name or UID. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"virtual_ip_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Virtual IPV4 address for DHCP server replies. Must be set when \"allocate-method\" was selected to be \"automatic\".",
												},
												"dhcp_mac_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Calculated MAC address for DHCP allocation. Must be set when \"allocate-method\" was selected to be \"automatic\".",
													Default:     "per-machine",
												},
												"optional_parameters": {
													Type:        schema.TypeMap,
													Optional:    true,
													Description: "This configuration applies to all Office Mode methods except Automatic (using DHCP) and ipassignment.conf entries which contain this data.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"use_primary_dns_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use Primary DNS Server.",
																Default:     false,
															},
															"primary_dns_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Primary DNS Server. Identified by name or UID. Must be set when \"use-primary-dns-server\" is true and can not be set when \"use-primary-dns-server\" is false.",
															},
															"use_first_backup_dns_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use First Backup DNS Server.",
																Default:     false,
															},
															"first_backup_dns_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "First Backup DNS Server. Identified by name or UID. Must be set when \"use-first-backup-dns-server\" is true and can not be set when \"use-first-backup-dns-server\" is false.",
															},
															"use_second_backup_dns_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use Second Backup DNS Server.",
																Default:     false,
															},
															"second_backup_dns_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Second Backup DNS Server. Identified by name or UID. Must be set when \"use-second-backup-dns-server\" is true and can not be set when \"use-second-backup-dns-server\" is false.",
															},
															"dns_suffixes": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "DNS Suffixes.",
															},
															"use_primary_wins_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use Primary WINS Server.",
																Default:     false,
															},
															"primary_wins_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Primary WINS Server. Identified by name or UID. Must be set when \"use-primary-wins-server\" is true and can not be set when \"use-primary-wins-server\" is false.",
															},
															"use_first_backup_wins_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use First Backup WINS Server.",
																Default:     false,
															},
															"first_backup_wins_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "First Backup WINS Server. Identified by name or UID. Must be set when \"use-first-backup-wins-server\" is true and can not be set when \"use-first-backup-wins-server\" is false.",
															},
															"use_second_backup_wins_server": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Use Second Backup WINS Server.",
																Default:     false,
															},
															"second_backup_wins_server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Second Backup WINS Server. Identified by name or UID. Must be set when \"use-second-backup-wins-server\" is true and can not be set when \"use-second-backup-wins-server\" is false.",
															},
															"ip_lease_duration": {
																Type:        schema.TypeInt,
																Optional:    true,
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
										Optional:    true,
										Description: "Support connectivity enhancement for gateways with multiple external interfaces.",
										Default:     false,
									},
									"perform_anti_spoofing": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Perform Anti-Spoofing on Office Mode addresses.",
										Default:     false,
									},
									"anti_spoofing_additional_addresses": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Additional IP Addresses for Anti-Spoofing. Identified by name or UID. Must be set when \"perform-anti-spoofings\" is true.",
										Default:     "None",
									},
								},
							},
						},
						"remote_access": {
							Type:        schema.TypeMap,
							Optional:    true,
							Description: "Remote Access.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"support_l2tp": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Support L2TP (relevant only when office mode is active).",
										Default:     false,
									},
									"l2tp_auth_method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "L2TP Authentication Method. Must be set when \"support-l2tp\" is true.",
										Default:     "md5",
									},
									"l2tp_certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "L2TP Certificate. Must be set when \"l2tp-auth-method\" was selected to be \"certificate\". Insert \"defaultCert\" when you want to use the default certificate.",
									},
									"allow_vpn_clients_to_route_traffic": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Allow VPN clients to route traffic.",
										Default:     false,
									},
									"support_nat_traversal_mechanism": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Support NAT traversal mechanism (UDP encapsulation).",
										Default:     true,
									},
									"nat_traversal_service": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Allocated NAT traversal UDP service. Identified by name or UID. Must be set when \"support-nat-traversal-mechanism\" is true.",
										Default:     "VPN1_IPSEC_encapsulation",
									},
									"support_visitor_mode": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Support Visitor Mode.",
										Default:     false,
									},
									"visitor_mode_service": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "TCP Service for Visitor Mode. Identified by name or UID. Must be set when \"support-visitor-mode\" is true.",
										Default:     "https",
									},
									"visitor_mode_interface": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Interface for Visitor Mode. Must be set when \"support-visitor-mode\" is true. Insert IPV4 Address of existing interface or \"All IPs\" when you want all interfaces.",
										Default:     "All IPs",
									},
								},
							},
						},
						"vpn_domain": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Gateway VPN domain identified by the name or UID.",
						},
						"vpn_domain_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Gateway VPN domain type.",
						},
						"vpn_domain_exclude_external_ip_addresses": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Exclude the external IP addresses from the VPN domain of this Security Gateway.",
						},
					},
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

func createManagementSimpleCluster(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	cluster := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		cluster["name"] = v.(string)
	}

	if v, ok := d.GetOk("ipv4_address"); ok {
		cluster["ipv4-address"] = v.(string)
	}

	if v, ok := d.GetOk("ipv6_address"); ok {
		cluster["ipv6-address"] = v.(string)
	}

	if v, ok := d.GetOk("cluster_mode"); ok {
		cluster["cluster-mode"] = v.(string)
	}

	if v, ok := d.GetOkExists("geo_mode"); ok {
		cluster["geo-mode"] = v.(bool)
	}

	if v, ok := d.GetOk("advanced_settings"); ok {

		advancedSettingsList := v.([]interface{})

		if len(advancedSettingsList) > 0 {

			advancedSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("advanced_settings.0.connection_persistence"); ok {
				advancedSettingsPayload["connection-persistence"] = v.(string)
			}
			if _, ok := d.GetOk("advanced_settings.0.sam"); ok {

				samPayload := make(map[string]interface{})

				if v, ok := d.GetOk("advanced_settings.0.sam.0.forward_to_other_sam_servers"); ok {
					samPayload["forward-to-other-sam-servers"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("advanced_settings.0.sam.0.use_early_versions"); ok {
					samPayload["use-early-versions"] = v
				}
				if v, ok := d.GetOk("advanced_settings.0.sam.0.purge_sam_file"); ok {
					samPayload["purge-sam-file"] = v
				}
				advancedSettingsPayload["sam"] = samPayload
			}
			cluster["advanced-settings"] = advancedSettingsPayload
		}
	}

	if v, ok := d.GetOkExists("enable_https_inspection"); ok {
		cluster["enable-https-inspection"] = v.(bool)
	}

	if v, ok := d.GetOk("fetch_policy"); ok {
		cluster["fetch-policy"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("hit_count"); ok {
		cluster["hit-count"] = v.(bool)
	}

	if v, ok := d.GetOk("https_inspection"); ok {

		httpsInspectionList := v.([]interface{})

		if len(httpsInspectionList) > 0 {

			httpsInspectionPayload := make(map[string]interface{})

			if _, ok := d.GetOk("https_inspection.0.bypass_on_failure"); ok {

				bypassOnFailurePayload := make(map[string]interface{})

				if v, ok := d.GetOk("https_inspection.0.bypass_on_failure.0.override_profile"); ok {
					bypassOnFailurePayload["override-profile"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("https_inspection.0.bypass_on_failure.0.value"); ok {
					bypassOnFailurePayload["value"] = strconv.FormatBool(v.(bool))
				}
				httpsInspectionPayload["bypass-on-failure"] = bypassOnFailurePayload
			}
			if _, ok := d.GetOk("https_inspection.0.site_categorization_allow_mode"); ok {

				siteCategorizationAllowModePayload := make(map[string]interface{})

				if v, ok := d.GetOk("https_inspection.0.site_categorization_allow_mode.0.override_profile"); ok {
					siteCategorizationAllowModePayload["override-profile"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("https_inspection.0.site_categorization_allow_mode.0.value"); ok {
					siteCategorizationAllowModePayload["value"] = v.(string)
				}
				httpsInspectionPayload["site-categorization-allow-mode"] = siteCategorizationAllowModePayload
			}
			if _, ok := d.GetOk("https_inspection.0.deny_untrusted_server_cert"); ok {

				denyUntrustedServerCertPayload := make(map[string]interface{})

				if v, ok := d.GetOk("https_inspection.0.deny_untrusted_server_cert.0.override_profile"); ok {
					denyUntrustedServerCertPayload["override-profile"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("https_inspection.0.deny_untrusted_server_cert.0.value"); ok {
					denyUntrustedServerCertPayload["value"] = strconv.FormatBool(v.(bool))
				}
				httpsInspectionPayload["deny-untrusted-server-cert"] = denyUntrustedServerCertPayload
			}
			if _, ok := d.GetOk("https_inspection.0.deny_revoked_server_cert"); ok {

				denyRevokedServerCertPayload := make(map[string]interface{})

				if v, ok := d.GetOk("https_inspection.0.deny_revoked_server_cert.0.override_profile"); ok {
					denyRevokedServerCertPayload["override-profile"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("https_inspection.0.deny_revoked_server_cert.0.value"); ok {
					denyRevokedServerCertPayload["value"] = strconv.FormatBool(v.(bool))
				}
				httpsInspectionPayload["deny-revoked-server-cert"] = denyRevokedServerCertPayload
			}
			if _, ok := d.GetOk("https_inspection.0.deny_expired_server_cert"); ok {

				denyExpiredServerCertPayload := make(map[string]interface{})

				if v, ok := d.GetOk("https_inspection.0.deny_expired_server_cert.0.override_profile"); ok {
					denyExpiredServerCertPayload["override-profile"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("https_inspection.0.deny_expired_server_cert.0.value"); ok {
					denyExpiredServerCertPayload["value"] = strconv.FormatBool(v.(bool))
				}
				httpsInspectionPayload["deny-expired-server-cert"] = denyExpiredServerCertPayload
			}
			cluster["https-inspection"] = httpsInspectionPayload
		}
	}

	if v, ok := d.GetOkExists("identity_awareness"); ok {
		cluster["identity-awareness"] = v.(bool)
	}

	if v, ok := d.GetOk("identity_awareness_settings"); ok {

		identityAwarenessSettingsList := v.([]interface{})

		if len(identityAwarenessSettingsList) > 0 {

			identityAwarenessSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication"); ok {
				identityAwarenessSettingsPayload["browser-based-authentication"] = v.(bool)
			}
			if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings"); ok {

				browserBasedAuthenticationSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings"); ok {
					browserBasedAuthenticationSettingsPayload["authentication-settings"] = v
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings"); ok {
					browserBasedAuthenticationSettingsPayload["browser-based-authentication-portal-settings"] = v
				}
				identityAwarenessSettingsPayload["browser-based-authentication-settings"] = browserBasedAuthenticationSettingsPayload
			}
			if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent"); ok {
				identityAwarenessSettingsPayload["identity-agent"] = v.(bool)
			}
			if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings"); ok {

				identityAgentSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.agents_interval_keepalive"); ok {
					identityAgentSettingsPayload["agents-interval-keepalive"] = v
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.user_reauthenticate_interval"); ok {
					identityAgentSettingsPayload["user-reauthenticate-interval"] = v
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings"); ok {
					identityAgentSettingsPayload["authentication-settings"] = v
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings"); ok {
					identityAgentSettingsPayload["identity-agent-portal-settings"] = v
				}
				identityAwarenessSettingsPayload["identity-agent-settings"] = identityAgentSettingsPayload
			}
			if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector"); ok {
				identityAwarenessSettingsPayload["identity-collector"] = v.(bool)
			}
			if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings"); ok {

				identityCollectorSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authorized_clients"); ok {
					identityCollectorSettingsPayload["authorized-clients"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings"); ok {
					identityCollectorSettingsPayload["authentication-settings"] = v
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions"); ok {
					identityCollectorSettingsPayload["client-access-permissions"] = v
				}
				identityAwarenessSettingsPayload["identity-collector-settings"] = identityCollectorSettingsPayload
			}
			if _, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings"); ok {

				identitySharingSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.0.share_with_other_gateways"); ok {
					identitySharingSettingsPayload["share-with-other-gateways"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.0.receive_from_other_gateways"); ok {
					identitySharingSettingsPayload["receive-from-other-gateways"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.0.receive_from"); ok {
					identitySharingSettingsPayload["receive-from"] = v.(*schema.Set).List()
				}
				identityAwarenessSettingsPayload["identity-sharing-settings"] = identitySharingSettingsPayload
			}
			if _, ok := d.GetOk("identity_awareness_settings.0.proxy_settings"); ok {

				proxySettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("identity_awareness_settings.0.proxy_settings.0.detect_using_x_forward_for"); ok {
					proxySettingsPayload["detect-using-x-forward-for"] = strconv.FormatBool(v.(bool))
				}
				identityAwarenessSettingsPayload["proxy-settings"] = proxySettingsPayload
			}
			if v, ok := d.GetOk("identity_awareness_settings.0.remote_access"); ok {
				identityAwarenessSettingsPayload["remote-access"] = v.(bool)
			}
			cluster["identity-awareness-settings"] = identityAwarenessSettingsPayload
		}
	}

	if v, ok := d.GetOk("ips_update_policy"); ok {
		cluster["ips-update-policy"] = v.(string)
	}

	if v, ok := d.GetOkExists("nat_hide_internal_interfaces"); ok {
		cluster["nat-hide-internal-interfaces"] = v.(bool)
	}

	if _, ok := d.GetOk("nat_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("nat_settings.auto_rule"); ok {
			res["auto-rule"] = v
		}
		if v, ok := d.GetOk("nat_settings.ipv4_address"); ok {
			res["ipv4-address"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.ipv6_address"); ok {
			res["ipv6-address"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.hide_behind"); ok {
			res["hide-behind"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.install_on"); ok {
			res["install-on"] = v.(string)
		}
		if v, ok := d.GetOk("nat_settings.method"); ok {
			res["method"] = v.(string)
		}
		cluster["nat-settings"] = res
	}

	if v, ok := d.GetOk("platform_portal_settings"); ok {

		platformPortalSettingsList := v.([]interface{})

		if len(platformPortalSettingsList) > 0 {

			platformPortalSettingsPayload := make(map[string]interface{})

			if _, ok := d.GetOk("platform_portal_settings.0.portal_web_settings"); ok {

				portalWebSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("platform_portal_settings.0.portal_web_settings.0.aliases"); ok {
					portalWebSettingsPayload["aliases"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("platform_portal_settings.0.portal_web_settings.0.main_url"); ok {
					portalWebSettingsPayload["main-url"] = v.(string)
				}
				platformPortalSettingsPayload["portal-web-settings"] = portalWebSettingsPayload
			}
			if _, ok := d.GetOk("platform_portal_settings.0.certificate_settings"); ok {

				certificateSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("platform_portal_settings.0.certificate_settings.0.base64_certificate"); ok {
					certificateSettingsPayload["base64-certificate"] = v.(string)
				}
				if v, ok := d.GetOk("platform_portal_settings.0.certificate_settings.0.base64_password"); ok {
					certificateSettingsPayload["base64-password"] = v.(string)
				}
				platformPortalSettingsPayload["certificate-settings"] = certificateSettingsPayload
			}
			if _, ok := d.GetOk("platform_portal_settings.0.accessibility"); ok {

				accessibilityPayload := make(map[string]interface{})

				if v, ok := d.GetOk("platform_portal_settings.0.accessibility.0.allow_access_from"); ok {
					accessibilityPayload["allow-access-from"] = v.(string)
				}
				if v, ok := d.GetOk("platform_portal_settings.0.accessibility.0.internal_access_settings"); ok {
					accessibilityPayload["internal-access-settings"] = v
				}
				platformPortalSettingsPayload["accessibility"] = accessibilityPayload
			}
			cluster["platform-portal-settings"] = platformPortalSettingsPayload
		}
	}

	if _, ok := d.GetOk("proxy_settings"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("proxy_settings.use_custom_proxy"); ok {
			res["use-custom-proxy"] = v
		}
		if v, ok := d.GetOk("proxy_settings.proxy_server"); ok {
			res["proxy-server"] = v.(string)
		}
		if v, ok := d.GetOk("proxy_settings.port"); ok {
			res["port"] = v
		}
		cluster["proxy-settings"] = res
	}

	if v, ok := d.GetOkExists("qos"); ok {
		cluster["qos"] = v.(bool)
	}

	if v, ok := d.GetOk("usercheck_portal_settings"); ok {

		usercheckPortalSettingsList := v.([]interface{})

		if len(usercheckPortalSettingsList) > 0 {

			usercheckPortalSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("usercheck_portal_settings.0.enabled"); ok {
				usercheckPortalSettingsPayload["enabled"] = v.(bool)
			}
			if _, ok := d.GetOk("usercheck_portal_settings.0.portal_web_settings"); ok {

				portalWebSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("usercheck_portal_settings.0.portal_web_settings.0.aliases"); ok {
					portalWebSettingsPayload["aliases"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("usercheck_portal_settings.0.portal_web_settings.0.main_url"); ok {
					portalWebSettingsPayload["main-url"] = v.(string)
				}
				usercheckPortalSettingsPayload["portal-web-settings"] = portalWebSettingsPayload
			}
			if _, ok := d.GetOk("usercheck_portal_settings.0.certificate_settings"); ok {

				certificateSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("usercheck_portal_settings.0.certificate_settings.0.base64_certificate"); ok {
					certificateSettingsPayload["base64-certificate"] = v.(string)
				}
				if v, ok := d.GetOk("usercheck_portal_settings.0.certificate_settings.0.base64_password"); ok {
					certificateSettingsPayload["base64-password"] = v.(string)
				}
				usercheckPortalSettingsPayload["certificate-settings"] = certificateSettingsPayload
			}
			if _, ok := d.GetOk("usercheck_portal_settings.0.accessibility"); ok {

				accessibilityPayload := make(map[string]interface{})

				if v, ok := d.GetOk("usercheck_portal_settings.0.accessibility.0.allow_access_from"); ok {
					accessibilityPayload["allow-access-from"] = v.(string)
				}
				if v, ok := d.GetOk("usercheck_portal_settings.0.accessibility.0.internal_access_settings"); ok {
					accessibilityPayload["internal-access-settings"] = v
				}
				usercheckPortalSettingsPayload["accessibility"] = accessibilityPayload
			}
			cluster["usercheck-portal-settings"] = usercheckPortalSettingsPayload
		}
	}

	if v, ok := d.GetOkExists("zero_phishing"); ok {
		cluster["zero-phishing"] = v.(bool)
	}

	if v, ok := d.GetOk("zero_phishing_fqdn"); ok {
		cluster["zero-phishing-fqdn"] = v.(string)
	}

	if v, ok := d.GetOk("interfaces"); ok {
		interfacesList := v.([]interface{})
		if len(interfacesList) > 0 {
			var interfacesPayload []map[string]interface{}
			for i := range interfacesList {

				interfacePayload := make(map[string]interface{})

				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".name"); ok {
					interfacePayload["name"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".interface_type"); ok {
					interfacePayload["interface-type"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_address"); ok {
					interfacePayload["ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_address"); ok {
					interfacePayload["ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_network_mask"); ok {
					interfacePayload["ipv4-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_network_mask"); ok {
					interfacePayload["ipv6-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_mask_length"); ok {
					interfacePayload["ipv4-mask-length"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_mask_length"); ok {
					interfacePayload["ipv6-mask-length"] = v.(string)
				}
				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".anti_spoofing"); ok {
					interfacePayload["anti-spoofing"] = v
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings"); ok {
					antiSpoofingSettings := make(map[string]interface{})
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.action"); ok {
						antiSpoofingSettings["action"] = v.(string)
					}
					interfacePayload["anti-spoofing-settings"] = antiSpoofingSettings
				}

				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".multicast_address"); ok {
					interfacePayload["multicast-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".multicast_address_type"); ok {
					interfacePayload["multicast-address-type"] = v.(string)
				}

				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone"); ok {
					interfacePayload["security-zone"] = v
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings"); ok {
					securityZoneSettings := make(map[string]interface{})
					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone_settings.auto_calculated"); ok {
						securityZoneSettings["auto-calculated"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings.specific_zone"); ok {
						securityZoneSettings["specific-zone"] = v.(string)
					}
					interfacePayload["security-zone-settings"] = securityZoneSettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology"); ok {
					interfacePayload["topology"] = v.(string)
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings"); ok {
					topologySettings := make(map[string]interface{})

					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".topology_settings.interface_leads_to_dmz"); ok {
						topologySettings["interface-leads-to-dmz"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.ip_address_behind_this_interface"); ok {
						topologySettings["ip-address-behind-this-interface"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.specific_network"); ok {
						topologySettings["specific-network"] = v.(string)
					}
					interfacePayload["topology-settings"] = topologySettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".color"); ok {
					interfacePayload["color"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".comments"); ok {
					interfacePayload["comments"] = v.(string)
				}
				interfacesPayload = append(interfacesPayload, interfacePayload)
			}
			cluster["interfaces"] = interfacesPayload
		}
	}

	if v, ok := d.GetOk("members"); ok {
		membersList := v.([]interface{})
		if len(membersList) > 0 {
			var membersPayload []map[string]interface{}
			for i := range membersList {
				memberPayload := make(map[string]interface{})

				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".name"); ok {
					memberPayload["name"] = v.(string)
				}

				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".ip_address"); ok {
					memberPayload["ip-address"] = v.(string)
				}

				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".one_time_password"); ok {
					memberPayload["one-time-password"] = v.(string)
				}

				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".interfaces"); ok {
					interfacesList := v.([]interface{})
					if len(interfacesList) > 0 {
						var interfacesPayload []map[string]interface{}
						for j := range interfacesList {
							interfacePayload := make(map[string]interface{})
							memberInterfacePrefix := "members." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j)
							if v, ok := d.GetOk(memberInterfacePrefix + ".name"); ok {
								interfacePayload["name"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv4_address"); ok {
								interfacePayload["ipv4-address"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv6_address"); ok {
								interfacePayload["ipv6-address"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv4_network_mask"); ok {
								interfacePayload["ipv4-network-mask"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv6_network_mask"); ok {
								interfacePayload["ipv6-network-mask"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv4_mask_length"); ok {
								interfacePayload["ipv4-mask-length"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv6_mask_length"); ok {
								interfacePayload["ipv6-mask-length"] = v.(string)
							}
							interfacesPayload = append(interfacesPayload, interfacePayload)
						}
						memberPayload["interfaces"] = interfacesPayload
					}
				}
				membersPayload = append(membersPayload, memberPayload)
			}
			cluster["members"] = membersPayload
		}
	}

	// Platform
	if v, ok := d.GetOk("os_name"); ok {
		cluster["os-name"] = v.(string)
	}

	if v, ok := d.GetOk("version"); ok {
		cluster["version"] = v.(string)
	}

	if v, ok := d.GetOk("hardware"); ok {
		cluster["hardware"] = v.(string)
	}

	// Blades
	if v, ok := d.GetOkExists("anti_bot"); ok {
		cluster["anti-bot"] = v
	}

	if v, ok := d.GetOkExists("anti_virus"); ok {
		cluster["anti-virus"] = v
	}

	if v, ok := d.GetOkExists("application_control"); ok {
		cluster["application-control"] = v
	}

	if v, ok := d.GetOkExists("content_awareness"); ok {
		cluster["content-awareness"] = v
	}

	if v, ok := d.GetOkExists("data_awareness"); ok {
		cluster["data-awareness"] = v
	}

	if v, ok := d.GetOkExists("ips"); ok {
		cluster["ips"] = v
	}

	if v, ok := d.GetOk("ips_settings"); ok {

		ipsSettingsList := v.([]interface{})

		if len(ipsSettingsList) > 0 {

			ipsSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("ips_settings.0.bypass_all_under_load"); ok {
				ipsSettingsPayload["bypass-all-under-load"] = v.(bool)
			}
			if v, ok := d.GetOk("ips_settings.0.bypass_track_method"); ok {
				ipsSettingsPayload["bypass-track-method"] = v.(string)
			}
			if _, ok := d.GetOk("ips_settings.0.top_cpu_consuming_protections"); ok {

				topCpuConsumingProtectionsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("ips_settings.0.top_cpu_consuming_protections.0.disable_period"); ok {
					topCpuConsumingProtectionsPayload["disable-period"] = v
				}
				if v, ok := d.GetOk("ips_settings.0.top_cpu_consuming_protections.0.disable_under_load"); ok {
					topCpuConsumingProtectionsPayload["disable-under-load"] = strconv.FormatBool(v.(bool))
				}
				ipsSettingsPayload["top-cpu-consuming-protections"] = topCpuConsumingProtectionsPayload
			}
			if v, ok := d.GetOk("ips_settings.0.activation_mode"); ok {
				ipsSettingsPayload["activation-mode"] = v.(string)
			}
			if v, ok := d.GetOk("ips_settings.0.cpu_usage_low_threshold"); ok {
				ipsSettingsPayload["cpu-usage-low-threshold"] = v.(int)
			}
			if v, ok := d.GetOk("ips_settings.0.cpu_usage_high_threshold"); ok {
				ipsSettingsPayload["cpu-usage-high-threshold"] = v.(int)
			}
			if v, ok := d.GetOk("ips_settings.0.memory_usage_low_threshold"); ok {
				ipsSettingsPayload["memory-usage-low-threshold"] = v.(int)
			}
			if v, ok := d.GetOk("ips_settings.0.memory_usage_high_threshold"); ok {
				ipsSettingsPayload["memory-usage-high-threshold"] = v.(int)
			}
			if v, ok := d.GetOk("ips_settings.0.send_threat_cloud_info"); ok {
				ipsSettingsPayload["send-threat-cloud-info"] = v.(bool)
			}
			if v, ok := d.GetOk("ips_settings.0.reject_on_cluster_fail_over"); ok {
				ipsSettingsPayload["reject-on-cluster-fail-over"] = v.(bool)
			}
			cluster["ips-settings"] = ipsSettingsPayload
		}
	}

	if v, ok := d.GetOkExists("threat_emulation"); ok {
		cluster["threat-emulation"] = v
	}

	if v, ok := d.GetOkExists("url_filtering"); ok {
		cluster["url-filtering"] = v
	}

	if v, ok := d.GetOkExists("vpn"); ok {
		cluster["vpn"] = v
	}

	if v, ok := d.GetOkExists("firewall"); ok {
		cluster["firewall"] = v
	}

	if _, ok := d.GetOk("firewall_settings"); ok {
		firewallSettings := make(map[string]interface{})
		if v, ok := d.GetOkExists("firewall_settings.auto_calculate_connections_hash_table_size_and_memory_pool"); ok {
			firewallSettings["auto-calculate-connections-hash-table-size-and-memory-pool"] = v
		}
		if v, ok := d.GetOkExists("firewall_settings.auto_maximum_limit_for_concurrent_connections"); ok {
			firewallSettings["auto-maximum-limit-for-concurrent-connections"] = v
		}
		if v, ok := d.GetOk("firewall_settings.connections_hash_size"); ok {
			firewallSettings["connections-hash-size"] = v.(int)
		}
		if v, ok := d.GetOk("firewall_settings.maximum_limit_for_concurrent_connections"); ok {
			firewallSettings["maximum-limit-for-concurrent-connections"] = v.(int)
		}
		if v, ok := d.GetOk("firewall_settings.maximum_memory_pool_size"); ok {
			firewallSettings["maximum-memory-pool-size"] = v.(int)
		}
		if v, ok := d.GetOk("firewall_settings.memory_pool_size"); ok {
			firewallSettings["memory-pool-size"] = v.(int)
		}
		cluster["firewall-settings"] = firewallSettings
	}

	// VPN settings
	if _, ok := d.GetOk("vpn_settings"); ok {
		vpnSettings := make(map[string]interface{})

		if _, ok := d.GetOk("vpn_settings.authentication"); ok {
			authentication := make(map[string]interface{})
			if v, ok := d.GetOk("vpn_settings.authentication.authentication_clients"); ok {
				authentication["authentication-clients"] = v.(*schema.Set).List()
			}
			vpnSettings["authentication"] = authentication
		}

		if _, ok := d.GetOk("vpn_settings.link_selection"); ok {
			linkSelection := make(map[string]interface{})
			if v, ok := d.GetOk("vpn_settings.link_selection.ip_selection"); ok {
				linkSelection["ip-selection"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.link_selection.dns_resolving_hostname"); ok {
				linkSelection["dns-resolving-hostname"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.link_selection.ip_address"); ok {
				linkSelection["ip-address"] = v.(string)
			}
			vpnSettings["link-selection"] = linkSelection
		}

		if v, ok := d.GetOk("vpn_settings.maximum_concurrent_ike_negotiations"); ok {
			vpnSettings["maximum-concurrent-ike-negotiations"] = v.(int)
		}
		if v, ok := d.GetOk("vpn_settings.maximum_concurrent_tunnels"); ok {
			vpnSettings["maximum-concurrent-tunnels"] = v.(int)
		}

		if _, ok := d.GetOk("vpn_settings.office_mode"); ok {
			officeMode := make(map[string]interface{})
			if v, ok := d.GetOk("vpn_settings.office_mode.mode"); ok {
				officeMode["mode"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.office_mode.group"); ok {
				officeMode["group"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.office_mode.support_multiple_interfaces"); ok {
				officeMode["support-multiple-interfaces"] = v
			}
			if v, ok := d.GetOkExists("vpn_settings.office_mode.perform_anti_spoofing"); ok {
				officeMode["perform-anti-spoofing"] = v
			}
			if v, ok := d.GetOk("vpn_settings.office_mode.anti_spoofing_additional_addresses"); ok {
				officeMode["anti-spoofing-additional-addresses"] = v.(string)
			}
			if _, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from"); ok {
				allocateIpAddressFrom := make(map[string]interface{})
				if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.radius_server"); ok {
					allocateIpAddressFrom["radius-server"] = v
				}
				if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.use_allocate_method"); ok {
					allocateIpAddressFrom["use-allocate-method"] = v
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.allocate_method"); ok {
					allocateIpAddressFrom["allocate-method"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.manual_network"); ok {
					allocateIpAddressFrom["manual-network"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.dhcp_server"); ok {
					allocateIpAddressFrom["dhcp-server"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.virtual_ip_address"); ok {
					allocateIpAddressFrom["virtual-ip-address"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.dhcp_mac_address"); ok {
					allocateIpAddressFrom["dhcp-mac-address"] = v.(string)
				}
				if _, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters"); ok {
					optionalParameters := make(map[string]interface{})
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_dns_server"); ok {
						optionalParameters["use-primary-dns-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_dns_server"); ok {
						optionalParameters["primary-dns-server"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_dns_server"); ok {
						optionalParameters["use-first-backup-dns-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_dns_server"); ok {
						optionalParameters["first-backup-dns-server"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_dns_server"); ok {
						optionalParameters["use-second-backup-dns-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_dns_server"); ok {
						optionalParameters["second-backup-dns-server"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.dns_suffixes"); ok {
						optionalParameters["dns-suffixes"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_wins_server"); ok {
						optionalParameters["use-primary-wins-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_wins_server"); ok {
						optionalParameters["primary-wins-server"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_wins_server"); ok {
						optionalParameters["use-first-backup-wins-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_wins_server"); ok {
						optionalParameters["first-backup-wins-server"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_wins_server"); ok {
						optionalParameters["use-second-backup-wins-server"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_wins_server"); ok {
						optionalParameters["second-backup-wins-server"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.ip_lease_duration"); ok {
						optionalParameters["ip-lease-duration"] = v.(int)
					}
					allocateIpAddressFrom["optional-parameters"] = optionalParameters
				}
				officeMode["allocate-ip-address-from"] = allocateIpAddressFrom
			}
			vpnSettings["office-mode"] = officeMode
		}

		if _, ok := d.GetOk("vpn_settings.remote_access"); ok {
			remoteAccess := make(map[string]interface{})
			if v, ok := d.GetOkExists("vpn_settings.remote_access.support_l2tp"); ok {
				remoteAccess["support-l2tp"] = v
			}
			if v, ok := d.GetOk("vpn_settings.remote_access.l2tp_auth_method"); ok {
				remoteAccess["l2tp-auth-method"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.remote_access.l2tp_certificate"); ok {
				remoteAccess["l2tp-certificate"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.allow_vpn_clients_to_route_traffic"); ok {
				remoteAccess["allow-vpn-clients-to-route-traffic"] = v
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.support_nat_traversal_mechanism"); ok {
				remoteAccess["support-nat-traversal-mechanism"] = v
			}
			if v, ok := d.GetOk("vpn_settings.remote_access.nat_traversal_service"); ok {
				remoteAccess["nat-traversal-service"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.support_visitor_mode"); ok {
				remoteAccess["support-visitor-mode"] = v
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.visitor_mode_service"); ok {
				remoteAccess["visitor-mode-service"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.remote_access.visitor_mode_interface"); ok {
				remoteAccess["visitor-mode-interface"] = v.(string)
			}
			vpnSettings["remote-access"] = remoteAccess
		}

		if v, ok := d.GetOk("vpn_settings.vpn_domain"); ok {
			vpnSettings["vpn-domain"] = v.(string)
		}
		if v, ok := d.GetOk("vpn_settings.vpn_domain_type"); ok {
			vpnSettings["vpn-domain-type"] = v.(string)
		}
		if v, ok := d.GetOkExists("vpn_settings.vpn_domain_exclude_external_ip_addresses"); ok {
			vpnSettings["vpn-domain-exclude-external-ip-addresses"] = v
		}
		cluster["vpn-settings"] = vpnSettings
	}

	// Logs
	if v, ok := d.GetOkExists("save_logs_locally"); ok {
		cluster["save-logs-locally"] = v
	}

	if v, ok := d.GetOk("send_alerts_to_server"); ok {
		cluster["send-alerts-to-server"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("send_logs_to_backup_server"); ok {
		cluster["send-logs-to-backup-server"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("send_logs_to_server"); ok {
		cluster["send-logs-to-server"] = v.(*schema.Set).List()
	}

	// General
	if v, ok := d.GetOk("tags"); ok {
		cluster["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("comments"); ok {
		cluster["comments"] = v.(string)
	}

	if v, ok := d.GetOk("color"); ok {
		cluster["color"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		cluster["ignore-warnings"] = v
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		cluster["ignore-errors"] = v
	}

	log.Println("Create Simple Cluster - Map = ", cluster)

	addClusterRes, err := client.ApiCall("add-simple-cluster", cluster, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !addClusterRes.Success {
		if addClusterRes.ErrorMsg != "" {
			return fmt.Errorf(addClusterRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-simple-cluster", addClusterRes.GetData())
		return fmt.Errorf(msg)
	}

	// add-simple-cluster returns task-id. Call show-simple-cluster for object uid.
	showClusterRes, err := client.ApiCall("show-simple-cluster", map[string]interface{}{"name": d.Get("name")}, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showClusterRes.Success {
		return fmt.Errorf(showClusterRes.ErrorMsg)
	}

	d.SetId(showClusterRes.GetData()["uid"].(string))

	return readManagementSimpleCluster(d, m)
}

func readManagementSimpleCluster(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showClusterRes, err := client.ApiCall("show-simple-cluster", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showClusterRes.Success {
		if objectNotFound(showClusterRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
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
				if v, _ := memberJson["priority"]; v != nil {
					memberState["priority"] = v
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

	if cluster["ips-settings"] != nil {

		ipsSettingsMap, ok := cluster["ips-settings"].(map[string]interface{})

		if ok {
			ipsSettingsMapToReturn := make(map[string]interface{})

			if v := ipsSettingsMap["bypass-all-under-load"]; v != nil {
				ipsSettingsMapToReturn["bypass_all_under_load"] = v
			}
			if v := ipsSettingsMap["bypass-track-method"]; v != nil {
				ipsSettingsMapToReturn["bypass_track_method"] = v
			}
			if v, ok := ipsSettingsMap["top-cpu-consuming-protections"]; ok {

				topCpuConsumingProtectionsMap, ok := v.(map[string]interface{})
				if ok {
					topCpuConsumingProtectionsMapToReturn := make(map[string]interface{})

					if v, _ := topCpuConsumingProtectionsMap["disable-period"]; v != nil {
						topCpuConsumingProtectionsMapToReturn["disable_period"] = v
					}
					if v, _ := topCpuConsumingProtectionsMap["disable-under-load"]; v != nil {
						topCpuConsumingProtectionsMapToReturn["disable_under_load"] = v
					}
					ipsSettingsMapToReturn["top_cpu_consuming_protections"] = []interface{}{topCpuConsumingProtectionsMapToReturn}
				}
			}
			if v := ipsSettingsMap["activation-mode"]; v != nil {
				ipsSettingsMapToReturn["activation_mode"] = v
			}
			if v := ipsSettingsMap["cpu-usage-low-threshold"]; v != nil {
				ipsSettingsMapToReturn["cpu_usage_low_threshold"] = v
			}
			if v := ipsSettingsMap["cpu-usage-high-threshold"]; v != nil {
				ipsSettingsMapToReturn["cpu_usage_high_threshold"] = v
			}
			if v := ipsSettingsMap["memory-usage-low-threshold"]; v != nil {
				ipsSettingsMapToReturn["memory_usage_low_threshold"] = v
			}
			if v := ipsSettingsMap["memory-usage-high-threshold"]; v != nil {
				ipsSettingsMapToReturn["memory_usage_high_threshold"] = v
			}
			if v := ipsSettingsMap["send-threat-cloud-info"]; v != nil {
				ipsSettingsMapToReturn["send_threat_cloud_info"] = v
			}
			if v := ipsSettingsMap["reject-on-cluster-fail-over"]; v != nil {
				ipsSettingsMapToReturn["reject_on_cluster_fail_over"] = v
			}
			_ = d.Set("ips_settings", []interface{}{ipsSettingsMapToReturn})

		}
	} else {
		_ = d.Set("ips_settings", nil)
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
		if v := vpnSettingsJson["vpn-domain-exclude-external-ip-addresses"]; v != nil {
			vpnSettingsState["vpn_domain_exclude_external_ip_addresses"] = v
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
		_ = d.Set("vpn_settings", vpnSettingsState)
	} else {
		_ = d.Set("vpn_settings", nil)
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

func updateManagementSimpleCluster(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)
	cluster := make(map[string]interface{})

	cluster["uid"] = d.Id()

	if d.HasChange("name") {
		if v, ok := d.GetOk("name"); ok {
			cluster["new-name"] = v
		}
	}

	if ok := d.HasChange("ipv4_address"); ok {
		if v, ok := d.GetOk("ipv4_address"); ok {
			cluster["ipv4-address"] = v
		}
	}

	if ok := d.HasChange("ipv6_address"); ok {
		if v, ok := d.GetOk("ipv6_address"); ok {
			cluster["ipv6-address"] = v
		}
	}

	if ok := d.HasChange("cluster_mode"); ok {
		if v, ok := d.GetOk("cluster_mode"); ok {
			cluster["cluster-mode"] = v.(string)
		}
	}

	if ok := d.HasChanges("geo_mode"); ok {
		if v, ok := d.GetOk("geo_mode"); ok {
			cluster["geo-mode"] = v.(bool)
		}
	}

	if d.HasChange("advanced_settings") {

		if v, ok := d.GetOk("advanced_settings"); ok {

			advancedSettingsList := v.([]interface{})

			if len(advancedSettingsList) > 0 {

				advancedSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("advanced_settings.0.connection_persistence"); ok {
					advancedSettingsPayload["connection-persistence"] = v.(string)
				}
				if _, ok := d.GetOk("advanced_settings.0.sam"); ok {

					samPayload := make(map[string]interface{})

					if v, ok := d.GetOk("advanced_settings.0.sam.0.forward_to_other_sam_servers"); ok {
						samPayload["forward-to-other-sam-servers"] = v
					}
					if v, ok := d.GetOk("advanced_settings.0.sam.0.use_early_versions"); ok {
						samPayload["use-early-versions"] = v
					}
					if v, ok := d.GetOk("advanced_settings.0.sam.0.purge_sam_file"); ok {
						samPayload["purge-sam-file"] = v
					}
					advancedSettingsPayload["sam"] = samPayload
				}
				cluster["advanced-settings"] = advancedSettingsPayload
			}
		}
	}

	if d.HasChange("enable_https_inspection") {
		if v, ok := d.GetOkExists("enable_https_inspection"); ok {
			cluster["enable-https-inspection"] = v.(bool)
		}
	}

	if d.HasChange("fetch_policy") {
		if v, ok := d.GetOk("fetch_policy"); ok {
			cluster["fetch_policy"] = v.(*schema.Set).List()
		}
		//else {
		//	oldFetchPolicy, _ := d.GetChange("fetch_policy")
		//	if oldFetchPolicy != nil {
		//		cluster["fetch-policy"] = map[string]interface{}{"remove": oldFetchPolicy.(*schema.Set).List()}
		//	}
		//}
	}

	if d.HasChange("hit_count") {
		if v, ok := d.GetOkExists("hit_count"); ok {
			cluster["hit-count"] = v
		}
	}

	if d.HasChange("https_inspection") {

		if v, ok := d.GetOk("https_inspection"); ok {

			httpsInspectionList := v.([]interface{})

			if len(httpsInspectionList) > 0 {

				httpsInspectionPayload := make(map[string]interface{})

				if _, ok := d.GetOk("https_inspection.0.bypass_on_failure"); ok {

					bypassOnFailurePayload := make(map[string]interface{})

					if v, ok := d.GetOk("https_inspection.0.bypass_on_failure.0.override_profile"); ok {
						bypassOnFailurePayload["override-profile"] = v
					}
					if v, ok := d.GetOk("https_inspection.0.bypass_on_failure.0.value"); ok {
						bypassOnFailurePayload["value"] = v
					}
					httpsInspectionPayload["bypass-on-failure"] = bypassOnFailurePayload
				}
				if _, ok := d.GetOk("https_inspection.0.site_categorization_allow_mode"); ok {

					siteCategorizationAllowModePayload := make(map[string]interface{})

					if v, ok := d.GetOk("https_inspection.0.site_categorization_allow_mode.0.override_profile"); ok {
						siteCategorizationAllowModePayload["override-profile"] = v
					}
					if v, ok := d.GetOk("https_inspection.0.site_categorization_allow_mode.0.value"); ok {
						siteCategorizationAllowModePayload["value"] = v
					}
					httpsInspectionPayload["site-categorization-allow-mode"] = siteCategorizationAllowModePayload
				}
				if _, ok := d.GetOk("https_inspection.0.deny_untrusted_server_cert"); ok {

					denyUntrustedServerCertPayload := make(map[string]interface{})

					if v, ok := d.GetOk("https_inspection.0.deny_untrusted_server_cert.0.override_profile"); ok {
						denyUntrustedServerCertPayload["override-profile"] = v
					}
					if v, ok := d.GetOk("https_inspection.0.deny_untrusted_server_cert.0.value"); ok {
						denyUntrustedServerCertPayload["value"] = v
					}
					httpsInspectionPayload["deny-untrusted-server-cert"] = denyUntrustedServerCertPayload
				}
				if _, ok := d.GetOk("https_inspection.0.deny_revoked_server_cert"); ok {

					denyRevokedServerCertPayload := make(map[string]interface{})

					if v, ok := d.GetOk("https_inspection.0.deny_revoked_server_cert.0.override_profile"); ok {
						denyRevokedServerCertPayload["override-profile"] = v
					}
					if v, ok := d.GetOk("https_inspection.0.deny_revoked_server_cert.0.value"); ok {
						denyRevokedServerCertPayload["value"] = v
					}
					httpsInspectionPayload["deny-revoked-server-cert"] = denyRevokedServerCertPayload
				}
				if _, ok := d.GetOk("https_inspection.0.deny_expired_server_cert"); ok {

					denyExpiredServerCertPayload := make(map[string]interface{})

					if v, ok := d.GetOk("https_inspection.0.deny_expired_server_cert.0.override_profile"); ok {
						denyExpiredServerCertPayload["override-profile"] = v
					}
					if v, ok := d.GetOk("https_inspection.0.deny_expired_server_cert.0.value"); ok {
						denyExpiredServerCertPayload["value"] = v
					}
					httpsInspectionPayload["deny-expired-server-cert"] = denyExpiredServerCertPayload
				}
				cluster["https-inspection"] = httpsInspectionPayload
			}
		}
	}

	if d.HasChange("identity_awareness") {
		if v, ok := d.GetOkExists("identity_awareness"); ok {
			cluster["identity-awareness"] = v.(bool)
		}
	}

	if d.HasChange("identity_awareness_settings") {

		if v, ok := d.GetOk("identity_awareness_settings"); ok {

			identityAwarenessSettingsList := v.([]interface{})

			if len(identityAwarenessSettingsList) > 0 {

				identityAwarenessSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication"); ok {
					identityAwarenessSettingsPayload["browser-based-authentication"] = v.(bool)
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings"); ok {

					browserBasedAuthenticationSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings"); ok {
						browserBasedAuthenticationSettingsPayload["authentication-settings"] = v
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings"); ok {
						browserBasedAuthenticationSettingsPayload["browser-based-authentication-portal-settings"] = v
					}
					identityAwarenessSettingsPayload["browser-based-authentication-settings"] = browserBasedAuthenticationSettingsPayload
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent"); ok {
					identityAwarenessSettingsPayload["identity-agent"] = v.(bool)
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings"); ok {

					identityAgentSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.agents_interval_keepalive"); ok {
						identityAgentSettingsPayload["agents-interval-keepalive"] = v
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.user_reauthenticate_interval"); ok {
						identityAgentSettingsPayload["user-reauthenticate-interval"] = v
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings"); ok {
						identityAgentSettingsPayload["authentication-settings"] = v
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings"); ok {
						identityAgentSettingsPayload["identity-agent-portal-settings"] = v
					}
					identityAwarenessSettingsPayload["identity-agent-settings"] = identityAgentSettingsPayload
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector"); ok {
					identityAwarenessSettingsPayload["identity-collector"] = v.(bool)
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings"); ok {

					identityCollectorSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authorized_clients"); ok {
						identityCollectorSettingsPayload["authorized-clients"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings"); ok {
						identityCollectorSettingsPayload["authentication-settings"] = v
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions"); ok {
						identityCollectorSettingsPayload["client-access-permissions"] = v
					}
					identityAwarenessSettingsPayload["identity-collector-settings"] = identityCollectorSettingsPayload
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings"); ok {

					identitySharingSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.0.share_with_other_gateways"); ok {
						identitySharingSettingsPayload["share-with-other-gateways"] = v
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.0.receive_from_other_gateways"); ok {
						identitySharingSettingsPayload["receive-from-other-gateways"] = v
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.0.receive_from"); ok {
						identitySharingSettingsPayload["receive-from"] = v.(*schema.Set).List()
					}
					identityAwarenessSettingsPayload["identity-sharing-settings"] = identitySharingSettingsPayload
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.proxy_settings"); ok {

					proxySettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("identity_awareness_settings.0.proxy_settings.0.detect_using_x_forward_for"); ok {
						proxySettingsPayload["detect-using-x-forward-for"] = v
					}
					identityAwarenessSettingsPayload["proxy-settings"] = proxySettingsPayload
				}
				if v, ok := d.GetOk("identity_awareness_settings.0.remote_access"); ok {
					identityAwarenessSettingsPayload["remote-access"] = v.(bool)
				}
				cluster["identity-awareness-settings"] = identityAwarenessSettingsPayload
			}
		}
	}

	if ok := d.HasChange("ips_update_policy"); ok {
		if v, ok := d.GetOk("ips_update_policy"); ok {
			cluster["ips-update-policy"] = v
		}
	}
	if ok := d.HasChange("nat_hide_internal_interfaces"); ok {
		if v, ok := d.GetOkExists("nat_hide_internal_interfaces"); ok {
			cluster["nat-hide-internal-interfaces"] = v.(bool)
		}
	}

	if d.HasChange("nat_settings") {

		if _, ok := d.GetOk("nat_settings"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOk("nat_settings.auto_rule"); ok {
				res["auto-rule"] = v
			}
			if v, ok := d.GetOk("nat_settings.ipv4_address"); ok {
				res["ipv4-address"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.ipv6_address"); ok {
				res["ipv6-address"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.hide_behind"); ok {
				res["hide-behind"] = v
			}
			if v, ok := d.GetOk("nat_settings.install_on"); ok {
				res["install-on"] = v
			}
			if v, ok := d.GetOk("nat_settings.method"); ok {
				res["method"] = v
			}
			cluster["nat-settings"] = res
		}
	}

	if d.HasChange("platform_portal_settings") {

		if v, ok := d.GetOk("platform_portal_settings"); ok {

			platformPortalSettingsList := v.([]interface{})

			if len(platformPortalSettingsList) > 0 {

				platformPortalSettingsPayload := make(map[string]interface{})

				if _, ok := d.GetOk("platform_portal_settings.0.portal_web_settings"); ok {

					portalWebSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("platform_portal_settings.0.portal_web_settings.0.aliases"); ok {
						portalWebSettingsPayload["aliases"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("platform_portal_settings.0.portal_web_settings.0.main_url"); ok {
						portalWebSettingsPayload["main-url"] = v.(string)
					}
					platformPortalSettingsPayload["portal-web-settings"] = portalWebSettingsPayload
				}
				if _, ok := d.GetOk("platform_portal_settings.0.certificate_settings"); ok {

					certificateSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("platform_portal_settings.0.certificate_settings.0.base64_certificate"); ok {
						certificateSettingsPayload["base64-certificate"] = v.(string)
					}
					if v, ok := d.GetOk("platform_portal_settings.0.certificate_settings.0.base64_password"); ok {
						certificateSettingsPayload["base64-password"] = v.(string)
					}
					platformPortalSettingsPayload["certificate-settings"] = certificateSettingsPayload
				}
				if _, ok := d.GetOk("platform_portal_settings.0.accessibility"); ok {

					accessibilityPayload := make(map[string]interface{})

					if v, ok := d.GetOk("platform_portal_settings.0.accessibility.0.allow_access_from"); ok {
						accessibilityPayload["allow-access-from"] = v.(string)
					}
					if v, ok := d.GetOk("platform_portal_settings.0.accessibility.0.internal_access_settings"); ok {
						accessibilityPayload["internal-access-settings"] = v
					}
					platformPortalSettingsPayload["accessibility"] = accessibilityPayload
				}
				cluster["platform-portal-settings"] = platformPortalSettingsPayload
			}
		}
	}

	if d.HasChange("proxy_settings") {

		if _, ok := d.GetOk("proxy_settings"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOk("proxy_settings.use_custom_proxy"); ok {
				res["use-custom-proxy"] = v
			}
			if v, ok := d.GetOk("proxy_settings.proxy_server"); ok {
				res["proxy-server"] = v
			}
			if v, ok := d.GetOk("proxy_settings.port"); ok {
				res["port"] = v
			}
			cluster["proxy-settings"] = res
		}
	}

	if d.HasChange("qos") {
		if v, ok := d.GetOkExists("qos"); ok {
			cluster["qos"] = v.(bool)
		}
	}

	if d.HasChange("usercheck_portal_settings") {

		if v, ok := d.GetOk("usercheck_portal_settings"); ok {

			usercheckPortalSettingsList := v.([]interface{})

			if len(usercheckPortalSettingsList) > 0 {

				usercheckPortalSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("usercheck_portal_settings.0.enabled"); ok {
					usercheckPortalSettingsPayload["enabled"] = v.(bool)
				}
				if _, ok := d.GetOk("usercheck_portal_settings.0.portal_web_settings"); ok {

					portalWebSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("usercheck_portal_settings.0.portal_web_settings.0.aliases"); ok {
						portalWebSettingsPayload["aliases"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("usercheck_portal_settings.0.portal_web_settings.0.main_url"); ok {
						portalWebSettingsPayload["main-url"] = v.(string)
					}
					usercheckPortalSettingsPayload["portal-web-settings"] = portalWebSettingsPayload
				}
				if _, ok := d.GetOk("usercheck_portal_settings.0.certificate_settings"); ok {

					certificateSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("usercheck_portal_settings.0.certificate_settings.0.base64_certificate"); ok {
						certificateSettingsPayload["base64-certificate"] = v.(string)
					}
					if v, ok := d.GetOk("usercheck_portal_settings.0.certificate_settings.0.base64_password"); ok {
						certificateSettingsPayload["base64-password"] = v.(string)
					}
					usercheckPortalSettingsPayload["certificate-settings"] = certificateSettingsPayload
				}
				if _, ok := d.GetOk("usercheck_portal_settings.0.accessibility"); ok {

					accessibilityPayload := make(map[string]interface{})

					if v, ok := d.GetOk("usercheck_portal_settings.0.accessibility.0.allow_access_from"); ok {
						accessibilityPayload["allow-access-from"] = v.(string)
					}
					if v, ok := d.GetOk("usercheck_portal_settings.0.accessibility.0.internal_access_settings"); ok {
						accessibilityPayload["internal-access-settings"] = v
					}
					usercheckPortalSettingsPayload["accessibility"] = accessibilityPayload
				}
				cluster["usercheck-portal-settings"] = usercheckPortalSettingsPayload
			}
		}
	}

	if ok := d.HasChange("zero_phishing"); ok {
		if v, ok := d.GetOkExists("zero_phishing"); ok {
			cluster["zero-phishing"] = v.(bool)
		}
	}

	if ok := d.HasChange("zero_phishing_fqdn"); ok {
		if v, ok := d.GetOk("zero_phishing_fqdn"); ok {
			cluster["zero-phishing-fqdn"] = v
		}
	}

	if d.HasChange("interfaces") {
		if v, ok := d.GetOk("interfaces"); ok {
			interfacesList := v.([]interface{})
			var interfacesPayload []map[string]interface{}
			for i := range interfacesList {
				interfacePayload := make(map[string]interface{})

				interfacePayload["name"] = d.Get("interfaces." + strconv.Itoa(i) + ".name").(string)

				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_address"); ok {
					interfacePayload["ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_address"); ok {
					interfacePayload["ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_network_mask"); ok {
					interfacePayload["ipv4-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_network_mask"); ok {
					interfacePayload["ipv6-network-mask"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv4_mask_length"); ok {
					interfacePayload["ipv4-mask-length"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".ipv6_mask_length"); ok {
					interfacePayload["ipv6-mask-length"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".interface_type"); ok {
					interfacePayload["interface-type"] = v.(string)
				}
				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".anti_spoofing"); ok {
					interfacePayload["anti-spoofing"] = v
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings"); ok {
					antiSpoofingSettings := make(map[string]interface{})
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.action"); ok {
						antiSpoofingSettings["action"] = v.(string)
					}
					interfacePayload["anti-spoofing-settings"] = antiSpoofingSettings
				}
				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone"); ok {
					interfacePayload["security-zone"] = v
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".multicast_address"); ok {
					interfacePayload["multicast-address"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".multicast_address_type"); ok {
					interfacePayload["multicast-address-type"] = v.(string)
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings"); ok {
					securityZoneSettings := make(map[string]interface{})
					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone_settings.auto_calculated"); ok {
						securityZoneSettings["auto-calculated"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings.specific_zone"); ok {
						securityZoneSettings["specific-zone"] = v.(string)
					}
					interfacePayload["security-zone-settings"] = securityZoneSettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology"); ok {
					interfacePayload["topology"] = v.(string)
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings"); ok {
					topologySettings := make(map[string]interface{})

					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".topology_settings.interface_leads_to_dmz"); ok {
						topologySettings["interface-leads-to-dmz"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.ip_address_behind_this_interface"); ok {
						topologySettings["ip-address-behind-this-interface"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.specific_network"); ok {
						topologySettings["specific-network"] = v.(string)
					}
					interfacePayload["topology-settings"] = topologySettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".color"); ok {
					interfacePayload["color"] = v.(string)
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".comments"); ok {
					interfacePayload["comments"] = v.(string)
				}
				interfacesPayload = append(interfacesPayload, interfacePayload)
			}
			cluster["interfaces"] = interfacesPayload
		}
		//else {
		//	// Remove interface
		//	oldInterfaces, _ := d.GetChange("interfaces")
		//	if oldInterfaces != nil {
		//		var interfacesToDelete []interface{}
		//		for _, inter := range oldInterfaces.([]interface{}) {
		//			interfacesToDelete = append(interfacesToDelete, inter.(map[string]interface{})["name"].(string))
		//		}
		//		cluster["interfaces"] = map[string]interface{}{"remove": interfacesToDelete}
		//	}
		//}
	}

	if ok := d.HasChange("members"); ok {
		if v, ok := d.GetOk("members"); ok {
			membersList := v.([]interface{})
			var membersPayload []map[string]interface{}
			for i := range membersList {
				memberPayload := make(map[string]interface{})

				memberPayload["name"] = d.Get("members." + strconv.Itoa(i) + ".name").(string)

				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".ip_address"); ok {
					memberPayload["ip-address"] = v
				}

				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".priority"); ok {
					memberPayload["priority"] = v
				}

				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".one_time_password"); ok {
					memberPayload["one-time-password"] = v
				}

				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".interfaces"); ok {
					interfacesList := v.([]interface{})
					if len(interfacesList) > 0 {
						var interfacesPayload []map[string]interface{}
						for j := range interfacesList {
							interfacePayload := make(map[string]interface{})
							memberInterfacePrefix := "members." + strconv.Itoa(i) + ".interfaces." + strconv.Itoa(j)
							if v, ok := d.GetOk(memberInterfacePrefix + ".name"); ok {
								interfacePayload["name"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv4_address"); ok {
								interfacePayload["ipv4-address"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv6_address"); ok {
								interfacePayload["ipv6-address"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv4_network_mask"); ok {
								interfacePayload["ipv4-network-mask"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv6_network_mask"); ok {
								interfacePayload["ipv6-network-mask"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv4_mask_length"); ok {
								interfacePayload["ipv4-mask-length"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv6_mask_length"); ok {
								interfacePayload["ipv6-mask-length"] = v.(string)
							}
							interfacesPayload = append(interfacesPayload, interfacePayload)
						}
						memberPayload["interfaces"] = interfacesPayload
					}
				}
				membersPayload = append(membersPayload, memberPayload)
			}
			cluster["members"] = membersPayload
		}
		//else {
		//	oldMembers, _ := d.GetChange("members")
		//	if oldMembers != nil {
		//		var membersToDelete []interface{}
		//		for _, member := range oldMembers.([]interface{}) {
		//			membersToDelete = append(membersToDelete, member.(map[string]interface{})["name"].(string))
		//		}
		//		cluster["members"] = map[string]interface{}{"remove": membersToDelete}
		//	}
		//}
	}

	if ok := d.HasChange("one_time_password"); ok {
		if v, ok := d.GetOk("one_time_password"); ok {
			cluster["one-time-password"] = v.(string)
		}
	}

	if ok := d.HasChange("os_name"); ok {
		if v, ok := d.GetOk("os_name"); ok {
			cluster["os-name"] = v.(string)
		}
	}

	if ok := d.HasChange("version"); ok {
		if v, ok := d.GetOk("version"); ok {
			cluster["version"] = v.(string)
		}
	}

	if ok := d.HasChange("hardware"); ok {
		if v, ok := d.GetOk("hardware"); ok {
			cluster["hardware"] = v
		}
	}

	// Blades
	if ok := d.HasChange("anti_bot"); ok {
		if v, ok := d.GetOkExists("anti_bot"); ok {
			cluster["anti-bot"] = v
		}
	}

	if ok := d.HasChange("anti_virus"); ok {
		if v, ok := d.GetOkExists("anti_virus"); ok {
			cluster["anti-virus"] = v
		}
	}

	if ok := d.HasChange("application_control"); ok {
		if v, ok := d.GetOkExists("application_control"); ok {
			cluster["application-control"] = v
		}
	}

	if ok := d.HasChange("content_awareness"); ok {
		if v, ok := d.GetOkExists("content_awareness"); ok {
			cluster["content-awareness"] = v
		}
	}

	if ok := d.HasChange("data_awareness"); ok {
		if v, ok := d.GetOkExists("data_awareness"); ok {
			cluster["data-awareness"] = v
		}
	}

	if ok := d.HasChange("ips"); ok {
		if v, ok := d.GetOkExists("ips"); ok {
			cluster["ips"] = v
		}
	}

	if d.HasChange("ips_settings") {

		if v, ok := d.GetOk("ips_settings"); ok {

			ipsSettingsList := v.([]interface{})

			if len(ipsSettingsList) > 0 {

				ipsSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("ips_settings.0.bypass_all_under_load"); ok {
					ipsSettingsPayload["bypass-all-under-load"] = v.(bool)
				}
				if v, ok := d.GetOk("ips_settings.0.bypass_track_method"); ok {
					ipsSettingsPayload["bypass-track-method"] = v.(string)
				}
				if _, ok := d.GetOk("ips_settings.0.top_cpu_consuming_protections"); ok {
					topCpuConsumingProtectionsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("ips_settings.0.top_cpu_consuming_protections.0.disable_period"); ok {
						topCpuConsumingProtectionsPayload["disable-period"] = v
					}
					if v, ok := d.GetOk("ips_settings.0.top_cpu_consuming_protections.0.disable_under_load"); ok {
						topCpuConsumingProtectionsPayload["disable-under-load"] = strconv.FormatBool(v.(bool))
					}
					ipsSettingsPayload["top-cpu-consuming-protections"] = topCpuConsumingProtectionsPayload
				}
				if v, ok := d.GetOk("ips_settings.0.activation_mode"); ok {
					ipsSettingsPayload["activation-mode"] = v.(string)
				}
				if v, ok := d.GetOk("ips_settings.0.cpu_usage_low_threshold"); ok {
					ipsSettingsPayload["cpu-usage-low-threshold"] = v.(int)
				}
				if v, ok := d.GetOk("ips_settings.0.cpu_usage_high_threshold"); ok {
					ipsSettingsPayload["cpu-usage-high-threshold"] = v.(int)
				}
				if v, ok := d.GetOk("ips_settings.0.memory_usage_low_threshold"); ok {
					ipsSettingsPayload["memory-usage-low-threshold"] = v.(int)
				}
				if v, ok := d.GetOk("ips_settings.0.memory_usage_high_threshold"); ok {
					ipsSettingsPayload["memory-usage-high-threshold"] = v.(int)
				}
				if v, ok := d.GetOk("ips_settings.0.send_threat_cloud_info"); ok {
					ipsSettingsPayload["send-threat-cloud-info"] = v.(bool)
				}
				if v, ok := d.GetOk("ips_settings.0.reject_on_cluster_fail_over"); ok {
					ipsSettingsPayload["reject-on-cluster-fail-over"] = v.(bool)
				}
				cluster["ips-settings"] = ipsSettingsPayload
			}
		}
	}

	if ok := d.HasChange("threat_emulation"); ok {
		if v, ok := d.GetOkExists("threat_emulation"); ok {
			cluster["threat-emulation"] = v
		}
	}

	if ok := d.HasChange("url_filtering"); ok {
		if v, ok := d.GetOkExists("url_filtering"); ok {
			cluster["url-filtering"] = v
		}
	}

	if ok := d.HasChange("vpn"); ok {
		if v, ok := d.GetOkExists("vpn"); ok {
			cluster["vpn"] = v
		}
	}

	if ok := d.HasChange("firewall"); ok {
		if v, ok := d.GetOkExists("firewall"); ok {
			cluster["firewall"] = v
		}
	}

	if ok := d.HasChange("firewall_settings"); ok {
		if _, ok := d.GetOk("firewall_settings"); ok {
			firewallSettings := make(map[string]interface{})
			if v, ok := d.GetOkExists("firewall_settings.auto_calculate_connections_hash_table_size_and_memory_pool"); ok {
				firewallSettings["auto-calculate-connections-hash-table-size-and-memory-pool"] = v
			}
			if v, ok := d.GetOkExists("firewall_settings.auto_maximum_limit_for_concurrent_connections"); ok {
				firewallSettings["auto-maximum-limit-for-concurrent-connections"] = v
			}
			if v, ok := d.GetOk("firewall_settings.connections_hash_size"); ok {
				firewallSettings["connections-hash-size"] = v
			}
			if v, ok := d.GetOk("firewall_settings.maximum_limit_for_concurrent_connections"); ok {
				firewallSettings["maximum-limit-for-concurrent-connections"] = v
			}
			if v, ok := d.GetOk("firewall_settings.maximum_memory_pool_size"); ok {
				firewallSettings["maximum-memory-pool-size"] = v
			}
			if v, ok := d.GetOk("firewall_settings.memory_pool_size"); ok {
				firewallSettings["memory-pool-size"] = v
			}
			cluster["firewall-settings"] = firewallSettings
		}
	}

	// VPN settings
	if ok := d.HasChange("vpn_settings"); ok {
		if _, ok := d.GetOk("vpn_settings"); ok {
			vpnSettings := make(map[string]interface{})

			if _, ok := d.GetOk("vpn_settings.authentication"); ok {
				if ok := d.HasChange("vpn_settings.authentication.authentication_clients"); ok {
					authentication := make(map[string]interface{})
					if v, ok := d.GetOk("vpn_settings.authentication.authentication_clients"); ok {
						authentication["authentication-clients"] = v.(*schema.Set).List()
					}
					//else {
					//	oldValues, _ := d.GetChange("vpn_settings.authentication.authentication_clients")
					//	if oldValues != nil {
					//		authentication["authentication-clients"] = map[string]interface{}{"remove": oldValues.(*schema.Set).List()}
					//	}
					//}
					vpnSettings["authentication"] = authentication
				}
			}

			if _, ok := d.GetOk("vpn_settings.link_selection"); ok {
				linkSelection := make(map[string]interface{})
				if v, ok := d.GetOk("vpn_settings.link_selection.ip_selection"); ok {
					linkSelection["ip-selection"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.link_selection.dns_resolving_hostname"); ok {
					linkSelection["dns-resolving-hostname"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.link_selection.ip_address"); ok {
					linkSelection["ip-address"] = v.(string)
				}
				vpnSettings["link-selection"] = linkSelection
			}

			if ok := d.HasChange("vpn_settings.maximum_concurrent_ike_negotiations"); ok {
				if v, ok := d.GetOk("vpn_settings.maximum_concurrent_ike_negotiations"); ok {
					vpnSettings["maximum-concurrent-ike-negotiations"] = v.(int)
				}
			}

			if ok := d.HasChange("vpn_settings.maximum_concurrent_tunnels"); ok {
				if v, ok := d.GetOk("vpn_settings.maximum_concurrent_tunnels"); ok {
					vpnSettings["maximum-concurrent-tunnels"] = v.(int)
				}
			}

			if _, ok := d.GetOk("vpn_settings.office_mode"); ok {
				officeMode := make(map[string]interface{})

				if v, ok := d.GetOk("vpn_settings.office_mode.mode"); ok {
					officeMode["mode"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.group"); ok {
					officeMode["group"] = v.(string)
				}
				if v, ok := d.GetOkExists("vpn_settings.office_mode.support_multiple_interfaces"); ok {
					officeMode["support-multiple-interfaces"] = v
				}
				if v, ok := d.GetOkExists("vpn_settings.office_mode.perform_anti_spoofing"); ok {
					officeMode["perform-anti-spoofing"] = v
				}
				if v, ok := d.GetOk("vpn_settings.office_mode.anti_spoofing_additional_addresses"); ok {
					officeMode["anti-spoofing-additional-addresses"] = v.(string)
				}
				if _, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from"); ok {
					allocateIpAddressFrom := make(map[string]interface{})

					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.radius_server"); ok {
						allocateIpAddressFrom["radius-server"] = v
					}
					if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.use_allocate_method"); ok {
						allocateIpAddressFrom["use-allocate-method"] = v
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.allocate_method"); ok {
						allocateIpAddressFrom["allocate-method"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.manual_network"); ok {
						allocateIpAddressFrom["manual-network"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.dhcp_server"); ok {
						allocateIpAddressFrom["dhcp-server"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.virtual_ip_address"); ok {
						allocateIpAddressFrom["virtual-ip-address"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.dhcp_mac_address"); ok {
						allocateIpAddressFrom["dhcp-mac-address"] = v.(string)
					}

					if _, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters"); ok {
						optionalParameters := make(map[string]interface{})

						if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_dns_server"); ok {
							optionalParameters["use-primary-dns-server"] = v
						}
						if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_dns_server"); ok {
							optionalParameters["primary-dns-server"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_dns_server"); ok {
							optionalParameters["use-first-backup-dns-server"] = v
						}
						if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_dns_server"); ok {
							optionalParameters["first-backup-dns-server"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_dns_server"); ok {
							optionalParameters["use-second-backup-dns-server"] = v
						}
						if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_dns_server"); ok {
							optionalParameters["second-backup-dns-server"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.dns_suffixes"); ok {
							optionalParameters["dns-suffixes"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_primary_wins_server"); ok {
							optionalParameters["use-primary-wins-server"] = v
						}
						if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.primary_wins_server"); ok {
							optionalParameters["primary-wins-server"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_first_backup_wins_server"); ok {
							optionalParameters["use-first-backup-wins-server"] = v
						}
						if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.first_backup_wins_server"); ok {
							optionalParameters["first-backup-wins-server"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.use_second_backup_wins_server"); ok {
							optionalParameters["use-second-backup-wins-server"] = v
						}
						if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.second_backup_wins_server"); ok {
							optionalParameters["second-backup-wins-server"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.office_mode.allocate_ip_address_from.optional_parameters.ip_lease_duration"); ok {
							optionalParameters["ip-lease-duration"] = v.(int)
						}
						allocateIpAddressFrom["optional-parameters"] = optionalParameters
					}
					officeMode["allocate-ip-address-from"] = allocateIpAddressFrom
				}
				vpnSettings["office-mode"] = officeMode
			}

			if _, ok := d.GetOk("vpn_settings.remote_access"); ok {
				remoteAccess := make(map[string]interface{})
				if v, ok := d.GetOkExists("vpn_settings.remote_access.support_l2tp"); ok {
					remoteAccess["support-l2tp"] = v
				}
				if v, ok := d.GetOk("vpn_settings.remote_access.l2tp_auth_method"); ok {
					remoteAccess["l2tp-auth-method"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.remote_access.l2tp_certificate"); ok {
					remoteAccess["l2tp-certificate"] = v.(string)
				}
				if v, ok := d.GetOkExists("vpn_settings.remote_access.allow_vpn_clients_to_route_traffic"); ok {
					remoteAccess["allow-vpn-clients-to-route-traffic"] = v
				}
				if v, ok := d.GetOkExists("vpn_settings.remote_access.support_nat_traversal_mechanism"); ok {
					remoteAccess["support-nat-traversal-mechanism"] = v
				}
				if v, ok := d.GetOk("vpn_settings.remote_access.nat_traversal_service"); ok {
					remoteAccess["nat-traversal-service"] = v.(string)
				}
				if v, ok := d.GetOkExists("vpn_settings.remote_access.support_visitor_mode"); ok {
					remoteAccess["support-visitor-mode"] = v
				}
				if v, ok := d.GetOk("vpn_settings.remote_access.visitor_mode_service"); ok {
					remoteAccess["visitor-mode-service"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.remote_access.visitor_mode_interface"); ok {
					remoteAccess["visitor-mode-interface"] = v.(string)
				}
				vpnSettings["remote-access"] = remoteAccess
			}

			if v, ok := d.GetOk("vpn_settings.vpn_domain"); ok {
				vpnSettings["vpn-domain"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.vpn_domain_type"); ok {
				vpnSettings["vpn-domain-type"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.vpn_domain_exclude_external_ip_addresses"); ok {
				vpnSettings["vpn-domain-exclude-external-ip-addresses"] = v
			}
			cluster["vpn-settings"] = vpnSettings
		}
	}

	// Logs
	if ok := d.HasChange("save_logs_locally"); ok {
		if v, ok := d.GetOkExists("save_logs_locally"); ok {
			cluster["save-logs-locally"] = v
		}
	}

	if ok := d.HasChange("send_alerts_to_server"); ok {
		if v, ok := d.GetOk("send_alerts_to_server"); ok {
			cluster["send-alerts-to-server"] = v.(*schema.Set).List()
		}
		//else {
		//	oldValues, _ := d.GetChange("send_alerts_to_server")
		//	if oldValues != nil {
		//		cluster["send-alerts-to-server"] = map[string]interface{}{"remove": oldValues.(*schema.Set).List()}
		//	}
		//}
	}

	if ok := d.HasChange("send_logs_to_backup_server"); ok {
		if v, ok := d.GetOk("send_logs_to_backup_server"); ok {
			cluster["send-logs-to-backup-server"] = v.(*schema.Set).List()
		}
		//else {
		//	oldValues, _ := d.GetChange("send_logs_to_backup_server")
		//	if oldValues != nil {
		//		cluster["send-logs-to-backup-server"] = map[string]interface{}{"remove": oldValues.(*schema.Set).List()}
		//	}
		//}
	}
	if ok := d.HasChange("send_logs_to_server"); ok {
		if v, ok := d.GetOk("send_logs_to_server"); ok {
			cluster["send-logs-to-server"] = v.(*schema.Set).List()
		}
		//else {
		//	oldValues, _ := d.GetChange("send_logs_to_server")
		//	if oldValues != nil {
		//		cluster["send-logs-to-server"] = map[string]interface{}{"remove": oldValues.(*schema.Set).List()}
		//	}
		//}
	}

	if ok := d.HasChange("tags"); ok {
		if v, ok := d.GetOk("tags"); ok {
			cluster["tags"] = v.(*schema.Set).List()
		}
		//else {
		//	oldTags, _ := d.GetChange("tags")
		//	if oldTags != nil {
		//		cluster["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		//	}
		//}
	}

	if ok := d.HasChange("comments"); ok {
		if v, ok := d.GetOk("comments"); ok {
			cluster["comments"] = v.(string)
		}
	}

	if ok := d.HasChange("color"); ok {
		if v, ok := d.GetOk("color"); ok {
			cluster["color"] = v.(string)
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		cluster["ignore-warnings"] = v
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		cluster["ignore-errors"] = v
	}

	log.Println("Update Simple Cluster - Map = ", cluster)

	if len(cluster) != 3 {
		updateSimpleClusterRes, err := client.ApiCall("set-simple-cluster", cluster, client.GetSessionID(), true, client.IsProxyUsed())
		if err != nil {
			return fmt.Errorf(err.Error())
		}
		if !updateSimpleClusterRes.Success {
			if updateSimpleClusterRes.ErrorMsg != "" {
				return fmt.Errorf(updateSimpleClusterRes.ErrorMsg)
			}
			msg := createTaskFailMessage("set-simple-cluster", updateSimpleClusterRes.GetData())
			return fmt.Errorf(msg)
		}
	} else {
		// Payload contain only required fields: uid, ignore-warnings and ignore-errors
		// We got empty update, skip update API call...
		log.Println("Got empty update. Skip update API call...")
	}

	return readManagementSimpleCluster(d, m)
}

func deleteManagementSimpleCluster(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}
	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		payload["ignore-warnings"] = v
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		payload["ignore-errors"] = v
	}
	deleteClusterRes, err := client.ApiCall("delete-simple-cluster", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil || !deleteClusterRes.Success {
		if deleteClusterRes.ErrorMsg != "" {
			return fmt.Errorf(deleteClusterRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
