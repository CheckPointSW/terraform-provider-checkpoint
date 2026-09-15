package checkpoint

import (
	"github.com/CheckPointSW/terraform-provider-checkpoint/v3/upgraders"
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
		SchemaVersion: 1,
		StateUpgraders: []schema.StateUpgrader{
			{
				Type:    upgraders.ResourceManagementSimpleClusterV0().CoreConfigSchema().ImpliedType(),
				Upgrade: upgraders.ResourceManagementSimpleClusterStateUpgradeV0,
				Version: 0,
			},
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
						"bypass_on_client_failure": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Bypass HTTPS inspection on client failure.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to override the value inherited from the profile.",
									},
									"value": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to bypass on client failure.",
									},
								},
							},
						},
						"bypass_under_load": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Bypass HTTPS inspection under load.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to bypass under load.",
									},
								},
							},
						},
						"outbound_certificate": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Outbound HTTPS inspection certificate.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to override the value inherited from the profile.",
									},
									"value": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Outbound certificate identified by the name or UID.",
									},
								},
							},
						},
						"deployment_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "HTTPS inspection deployment mode.",
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
						"identity_web_api": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Identity Web API source.",
						},
						"identity_web_api_settings": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Identity Web API settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_settings": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Authentication Settings for Identity Web Api.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"users_directories": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Users directories.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"specific": {
																Type:        schema.TypeSet,
																Optional:    true,
																Description: "LDAP AU objects identified by the name or UID. Must be set when 'users-from-external-directories' was selected to be 'specific'.",
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"external_user_profile": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "External user profile.",
															},
															"internal_users": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Internal users.",
															},
															"users_from_external_directories": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Users from external directories.",
															},
														},
													},
												},
											},
										},
									},
									"authorized_clients": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
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
													Sensitive:   true,
													Description: "Client Secret.",
												},
											},
										},
									},
									"client_access_permissions": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Identity Web Api accessibility settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Configuration of the portal access settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"allow_access_from": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Allowed access to the web portal (based on interfaces, or security policy).",
															},
															"internal_access_settings": {
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "Configuration of the additional portal access settings for internal interfaces only.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"dmz": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'DMZ'.",
																		},
																		"undefined": {
																			Type:        schema.TypeBool,
																			Optional:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to 'Undefined'.",
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
															"ip_address": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Optional IP address to be used for the portal URL.",
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
																Sensitive:   true,
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
													Sensitive:   true,
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
									"cache_mode": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Identity cache mode.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"override_profile": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether to override the value inherited from the profile.",
												},
												"value": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether the identity cache is enabled.",
												},
											},
										},
									},
									"cache_mode_duration": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Identity cache mode duration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"override_profile": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether to override the value inherited from the profile.",
												},
												"value": {
													Type:        schema.TypeInt,
													Optional:    true,
													Description: "Identity cache duration in minutes.",
												},
											},
										},
									},
									"receive_from_infinity_identity": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether to receive identities from Infinity Identity.",
									},
									"scaled_sharing": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Whether scaled identity sharing is enabled.",
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
						"identity_based_enforcement": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "ON: Configures this object as a PEP-only object - identity-based enforcement (PEP) is enabled.<br>OFF: Configures this object as a PDP-only object - identity-ba",
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
						"apply_control_connections": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "This option performs NAT on VPN control connections to and from this object.",
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
									"ip_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Optional IP address to be used for the portal URL.",
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
										Sensitive:   true,
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
									"ip_address": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Optional IP address to be used for the portal URL.",
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
										Sensitive:   true,
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
				Deprecated:  "use zero_phishing_settings.manual_fqdn instead - the API replaced zero-phishing-fqdn with zero-phishing-settings",
			},
			"interfaces": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Network interfaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"dynamic_ip": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The Topology of interface with Dynamic IP is set to Automatic - External.",
						},
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
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Anti spoofing settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"exclude_packets": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Don't check packets from excluded network.",
									},
									"excluded_network_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Excluded network name.",
									},
									"excluded_network_uid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Excluded network UID.",
									},
									"spoof_tracking": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Spoof tracking.",
									},
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
							Type:        schema.TypeList,
							MaxItems:    1,
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
							Type:        schema.TypeList,
							MaxItems:    1,
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
						"auto_generate_ip": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Use an automatically generated IP address for the Gateway object (applies only to Smart-1 Cloud).",
						},
						"trust_method": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Trust method to use for establishing communication.",
						},
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
							Sensitive:   true,
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
									"anti_spoofing_settings": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "N/A",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"action": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).",
												},
												"exclude_packets": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Don't check packets from excluded network.",
												},
												"excluded_network_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Excluded network name.",
												},
												"excluded_network_uid": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Excluded network UID.",
												},
												"spoof_tracking": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Spoof tracking.",
												},
											},
										},
									},
									"security_zone_settings": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "N/A",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"auto_calculated": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Security Zone is calculated according to where the interface leads to.",
												},
											},
										},
									},
									"topology_settings": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "N/A",
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
													Description: "Network settings behind this interface.",
												},
												"specific_network": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Network behind this interface.",
												},
											},
										},
									},
									"anti_spoofing": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "N/A",
									},
									"dynamic_ip": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "The Topology of interface with Dynamic IP is set to Automatic - External.",
									},
									"security_zone": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "N/A",
									},
									"topology": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "N/A",
									},
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
				Type:        schema.TypeList,
				MaxItems:    1,
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
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Gateway VPN settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"interfaces": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Enhanced Link Selection Interfaces.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the interface.",
									},
									"ip_version": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The IP version of the interface's IP address (IPv4/IPv6).",
									},
									"next_hop_ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The IP address of the next hop.",
									},
									"priority": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Priority of a 'Backup' interface.",
									},
									"redundancy_mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Interface redundancy mode (Active/Backup).",
									},
									"static_nat_ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The NATed IPv4 address that hides the source IPv4 address of outgoing connections (applies only to IPv4).",
									},
								},
							},
						},
						"authentication": {
							Type:        schema.TypeList,
							MaxItems:    1,
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
									"single_authentication_client": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Settings for clients that support only single authentication method.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Allow clients that support only single authentication method.",
													Default:     true,
												},
												"allow_multiple_authentication_clients": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Allow clients that support multiple authentication methods to connect.",
													Default:     true,
												},
												"display_name": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Display name for the authentication method.",
													Default:     "standard",
												},
												"method": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Authentication method type.",
													Default:     "defined-on-user-record",
												},
												"secur_id": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "SecurID authentication settings, relevant only when method is \"secur-id\".",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Server object identified by the name or UID.",
															},
															"token_card_type": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Token card type.",
																Default:     "any",
															},
														},
													},
												},
												"radius": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "RADIUS authentication settings, relevant only when method is \"radius\".",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Server object identified by the name or UID.",
															},
															"ask_user_password": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Ask user for password during authentication.",
																Default:     false,
															},
														},
													},
												},
												"personal_certificate": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Personal certificate authentication settings, relevant only when method is \"personal-certificate\".",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fetch_username_from": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Fetch username from.",
																Default:     "subject-dn",
															},
															"storage_type": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Certificate storage type.",
																Default:     "any",
															},
															"source": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Certificate source field.",
																Default:     "subject",
															},
															"dn_part": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "DN part to extract.",
															},
															"dn_concurrence": {
																Type:        schema.TypeInt,
																Optional:    true,
																Description: "DN part occurrence number.",
																Default:     1,
															},
														},
													},
												},
												"client_display_settings": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Client display configuration settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"headline": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Display headline for authentication dialog.",
															},
															"username_label": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Label for username field.",
															},
															"password_label": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Label for password field.",
															},
														},
													},
												},
											},
										},
									},
									"override_global_dynamic_id_settings": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Override global dynamic ID settings.",
										Default:     false,
									},
									"dynamic_id_settings": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Dynamic ID settings, relevant only when \"override-global-dynamic-id-settings\" is true.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sms_provider_and_email_settings": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "SMS provider and email configuration.",
												},
												"sms_provider_credentials": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "SMS provider credentials configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"username": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "SMS provider username.",
															},
															"password": {
																Type:        schema.TypeString,
																Optional:    true,
																Sensitive:   true,
																Description: "SMS provider password.",
															},
															"api_id": {
																Type:        schema.TypeString,
																Optional:    true,
																Sensitive:   true,
																Description: "SMS provider API ID.",
															},
														},
													},
												},
												"advanced_settings": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Advanced Dynamic ID configuration settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dynamic_id_message": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Dynamic ID message displayed to users.",
															},
															"otp_settings": {
																Type:        schema.TypeList,
																MaxItems:    1,
																Optional:    true,
																Description: "One Time Password configuration settings.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"length": {
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Length of one time password.",
																			Default:     6,
																		},
																		"expiration": {
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "One time password expiration (in minutes).",
																			Default:     5,
																		},
																		"max_attempts": {
																			Type:        schema.TypeInt,
																			Optional:    true,
																			Description: "Number of times users can attempt to enter the one time password before the entire authentication process restarts.",
																			Default:     3,
																		},
																	},
																},
															},
															"enable_display_user_details": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Enable display of user details.",
																Default:     false,
															},
															"country_code": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Country code for SMS services.",
															},
															"user_details_retrieval": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "User details retrieval method.",
																Default:     "internal-or-ldap-or-local",
															},
														},
													},
												},
											},
										},
									},
									"send_machine_certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Configure when to send machine certificate.",
										Default:     "when-available",
									},
								},
							},
						},
						"link_selection": {
							Type:        schema.TypeList,
							MaxItems:    1,
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
									"route_selection_method": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Outgoing route selection method when initiating a tunnel.",
										Default:     "os-routing-table",
									},
									"responding_traffic": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Responding traffic route selection method.",
									},
									"source_ip_selection": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Source IP address selection method for outgoing traffic.",
									},
									"selected_ip": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Selected IP address. Must be set when \"source-ip-selection\" was selected to be \"manual\".",
									},
									"outgoing_link_tracking": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Outgoing link tracking method.",
									},
									"probing_settings": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Probing settings configuration. Only available when \"ip-selection\" is \"use-probing-with-high-availability\" or \"use-probing-with-load-sharing\".",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"probed_interfaces": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Specifies whether to probe all addresses defined in the topology tab or specific addresses.",
													Default:     "all",
												},
												"probed_interface_list": {
													Type:        schema.TypeSet,
													Optional:    true,
													Description: "List of specific IP addresses to probe. Only relevant when \"probed-interfaces\" is set to \"specific\".",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"use_primary_address": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "Whether to use a primary address for high availability probing.",
													Default:     false,
												},
												"primary_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Primary IP address to use. Must be one of the addresses from \"probed-interface-list\". Required when \"use-primary-address\" is true.",
												},
												"probing_method": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Probing method.",
													Default:     "ongoing",
												},
											},
										},
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
							Type:        schema.TypeList,
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
										Type:        schema.TypeList,
										MaxItems:    1,
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
													Type:        schema.TypeList,
													MaxItems:    1,
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
							Type:        schema.TypeList,
							MaxItems:    1,
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
						"advanced": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Advanced VPN settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tunnel_sharing_mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Tunnel sharing mode.",
									},
									"shutdown_on_gateway_restart": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Shutdown VPN tunnels on gateway restart.",
									},
									"enable_wire_mode": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable wire mode.",
									},
									"wire_mode_interfaces": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Wire mode interfaces.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"enable_wire_mode_log_traffic": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Log traffic in wire mode.",
									},
									"enable_nat_traversal": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable NAT traversal.",
									},
								},
							},
						},
						"exported_routes": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Exported routes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"internal_interfaces": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Export internal interfaces.",
									},
									"static_routes": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Export static routes.",
									},
									"custom_routes": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Export custom routes.",
									},
									"custom_routes_object": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Custom routes object identified by the name or UID.",
									},
								},
							},
						},
						"vpn_clients": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "VPN clients settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_endpoint_security_vpn": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable Endpoint Security VPN client.",
									},
									"enable_cp_mobile_for_windows": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable Check Point Mobile for Windows client.",
									},
									"enable_secu_remote": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable SecuRemote client.",
									},
									"enable_capsule_vpn_connect": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable Capsule VPN Connect client.",
									},
									"enable_ssl_network_extender": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Enable SSL Network Extender client.",
									},
									"gateway_authentication_certificate": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Gateway authentication certificate.",
									},
								},
							},
						},
						"enable_clientless_vpn": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable clientless VPN.",
						},
						"clientless_vpn_settings": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Clientless VPN settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate_gateway_authentication": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Certificate gateway authentication.",
									},
									"client_authentication": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Client authentication.",
									},
									"concurrent_servers_or_processes": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Number of concurrent servers or processes.",
									},
									"accept_only_3des": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Accept only 3DES.",
									},
								},
							},
						},
						"saml_portal_settings": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "SAML portal settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"portal_web_settings": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Configuration of the SAML portal web settings.",
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
												"ip_address": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Optional IP address to be used for the portal URL.",
												},
												"main_url": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The main URL for the portal.",
												},
											},
										},
									},
									"certificate_settings": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Configuration of the SAML portal certificate.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"base64_certificate": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The certificate file encoded in Base64 with padding.",
												},
												"base64_password": {
													Type:        schema.TypeString,
													Optional:    true,
													Sensitive:   true,
													Description: "Certificate file password.",
												},
											},
										},
									},
									"accessibility": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Configuration of the portal access settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allow_access_from": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Allowed access to the SAML portal.",
												},
												"internal_access_settings": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Configuration of the additional portal access settings for internal interfaces only.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"undefined": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Controls portal access settings for internal interfaces, whose topology is set to \"Undefined\".",
															},
															"dmz": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Controls portal access settings for internal interfaces, whose topology is set to \"DMZ\".",
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
			"anti_spam_and_email_security": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables Anti-Spam & Email-Security blade.",
			},
			"auto_topology_custom_recalculation_time": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Auto topology custom recalculation time (seconds).",
			},
			"auto_topology_use_custom_recalculation_time": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Auto topology to use custom recalculation time instead of default.",
			},
			"data_loss_prevention": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Data Loss Prevention blade.",
			},
			"mobile_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Mobile Access blade.",
			},
			"monitoring": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables Real Time Monitoring blade.",
			},
			"policy_server": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Policy Server blade.",
			},
			"rtm_counters_report": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables monitoring blades system counters report (e.g CPU Usage,Memory Usage).",
			},
			"rtm_traffic_report": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables monitoring blades traffic report.",
			},
			"rtm_traffic_report_per_connection": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Enables Monitoring blade traffic report per connection.",
			},
			"threat_extraction": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Threat Extraction blade enabled.",
			},
			"threat_prevention_mode": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed.",
			},
			"workforce_ai": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Workforce AI Security blade enabled. Requires content awareness blade and version R82.20 or higher to be enabled.",
			},
			"application_control_and_url_filtering_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Gateway Application Control and URL filtering settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"global_settings_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Whether to override global settings or not.",
						},
						"override_global_settings": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "override global settings object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fail_mode": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Fail mode - allow or block all requests.",
									},
									"website_categorization": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Website categorization object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"custom_mode": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Custom mode object.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"social_networking_widgets": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "Social networking widgets mode.",
															},
															"url_filtering": {
																Type:        schema.TypeString,
																Optional:    true,
																Description: "URL filtering mode.",
															},
														},
													},
												},
												"mode": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Website categorization mode.",
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
			"cluster_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "ClusterXL and VRRP Settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"member_recovery_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "In a High Availability cluster, each member is given a priority. The member with the highest priority serves as the gateway. If this gateway fails, control is p",
						},
						"state_synchronization": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Cluster State Synchronization settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"delayed": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Start synchronizing with delay of seconds, as defined by delayed-seconds, after connection initiation. Disabled when state-synchronization disabled.",
									},
									"delayed_seconds": {
										Type:        schema.TypeInt,
										Optional:    true,
										Description: "Start synchronizing X seconds after connection initiation . The values must be in a range between 2 and 3600.",
									},
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Use State Synchronization.",
									},
								},
							},
						},
						"track_changes_of_cluster_members": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Track changes in the status of Cluster Members.",
						},
						"use_virtual_mac": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Use Virtual MAC. By enabling Virtual MAC in ClusterXL High Availability New mode, or Load Sharing Unicast mode, all cluster members associate the same Virtual M",
						},
					},
				},
			},
			"communication_with_servers_behind_nat": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Gateway behind NAT communications settings with the server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"override_profile": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether to override the Server (Check Point Host) object configuration.",
						},
						"value": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "according-to-topology: Use the original or translated IP address of the server based on the Topology of Security Gateway interfaces.<br>original-ip-only: Use on",
						},
					},
				},
			},
			"zero_phishing_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Fqdn settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway_fqdn_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Manual Fqdn.",
						},
						"manual_fqdn": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Zero Phishing gateway FQDN.",
						},
					},
				},
			},
			"dns_server": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "DNS Server.",
			},
			"logs_settings": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Logs settings that apply to Quantum Security Gateways that run Gaia OS.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable alert when free disk space is below threshold.",
						},
						"alert_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Alert when free disk space below threshold.",
						},
						"alert_when_free_disk_space_below_type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Alert when free disk space below type.",
						},
						"before_delete_keep_logs_from_the_last_days": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable before delete keep logs from the last days.",
						},
						"before_delete_keep_logs_from_the_last_days_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Before delete keep logs from the last days threshold.",
						},
						"before_delete_run_script": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable Before delete run script.",
						},
						"before_delete_run_script_command": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Before delete run script command.",
						},
						"delete_index_files_older_than_days": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable delete index files older than days.",
						},
						"delete_index_files_older_than_days_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete index files older than days threshold.",
						},
						"delete_index_files_when_index_size_above": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable delete index files when index size above.",
						},
						"delete_index_files_when_index_size_above_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete index files when index size above threshold.",
						},
						"delete_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable delete when free disk space below.",
						},
						"delete_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Delete when free disk space below threshold.",
						},
						"detect_new_citrix_ica_application_names": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable detect new Citrix ICA application names.",
						},
						"distribute_logs_between_all_active_servers": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Distribute logs between all active servers.",
						},
						"forward_logs_to_log_server": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable forward logs to log server.",
						},
						"forward_logs_to_log_server_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Forward logs to log server name.",
						},
						"forward_logs_to_log_server_schedule_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Forward logs to log server schedule name.",
						},
						"free_disk_space_metrics": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Free disk space metrics.",
						},
						"include_tcp_state_information": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Include TCP state information. Relevant only when Firewall blade is enabled.",
						},
						"perform_log_rotate_before_log_forwarding": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable perform log rotate before log forwarding.",
						},
						"reject_connections_when_free_disk_space_below_threshold": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable reject connections when free disk space below threshold.",
						},
						"reserve_for_packet_capture_metrics": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Reserve for packet capture metrics.",
						},
						"reserve_for_packet_capture_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Reserve for packet capture threshold.",
						},
						"rotate_log_by_file_size": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable rotate log by file size.",
						},
						"rotate_log_file_size_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Log file size threshold.",
						},
						"rotate_log_on_schedule": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable rotate log on schedule.",
						},
						"rotate_log_schedule_name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Rotate log schedule name.",
						},
						"stop_logging_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable stop logging when free disk space below.",
						},
						"stop_logging_when_free_disk_space_below_threshold": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Stop logging when free disk space below threshold.",
						},
						"turn_on_qos_logging": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Enable turn on QoS Logging.",
						},
						"update_account_log_every": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Update account log in every amount of seconds.",
						},
					},
				},
			},
			"show_portals_certificate": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to show the portals certificate value in the reply.",
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
			if _, ok := d.GetOk("https_inspection.0.bypass_on_client_failure"); ok {

				bypassOnClientFailurePayload := make(map[string]interface{})

				if v, ok := d.GetOk("https_inspection.0.bypass_on_client_failure.0.override_profile"); ok {
					bypassOnClientFailurePayload["override-profile"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("https_inspection.0.bypass_on_client_failure.0.value"); ok {
					bypassOnClientFailurePayload["value"] = strconv.FormatBool(v.(bool))
				}
				httpsInspectionPayload["bypass-on-client-failure"] = bypassOnClientFailurePayload
			}
			if _, ok := d.GetOk("https_inspection.0.bypass_under_load"); ok {

				bypassUnderLoadPayload := make(map[string]interface{})

				if v, ok := d.GetOk("https_inspection.0.bypass_under_load.0.value"); ok {
					bypassUnderLoadPayload["value"] = strconv.FormatBool(v.(bool))
				}
				httpsInspectionPayload["bypass-under-load"] = bypassUnderLoadPayload
			}
			if _, ok := d.GetOk("https_inspection.0.outbound_certificate"); ok {

				outboundCertificatePayload := make(map[string]interface{})

				if v, ok := d.GetOk("https_inspection.0.outbound_certificate.0.override_profile"); ok {
					outboundCertificatePayload["override-profile"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("https_inspection.0.outbound_certificate.0.value"); ok {
					outboundCertificatePayload["value"] = v.(string)
				}
				httpsInspectionPayload["outbound-certificate"] = outboundCertificatePayload
			}
			if v, ok := d.GetOk("https_inspection.0.deployment_mode"); ok {
				httpsInspectionPayload["deployment-mode"] = v.(string)
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

				if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings"); ok {
					authenticationSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.authentication_method"); ok {
						authenticationSettingsPayload["authentication-method"] = v.(string)
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.identity_provider"); ok {
						authenticationSettingsPayload["identity-provider"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.radius"); ok {
						authenticationSettingsPayload["radius"] = v.(string)
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories"); ok {

						usersDirectoriesPayload := make(map[string]interface{})

						if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories.0.external_user_profile"); ok {
							usersDirectoriesPayload["external-user-profile"] = v.(bool)
						}
						if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories.0.internal_users"); ok {
							usersDirectoriesPayload["internal-users"] = v.(bool)
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories.0.specific"); ok {
							usersDirectoriesPayload["specific"] = v.(*schema.Set).List()
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories.0.users_from_external_directories"); ok {
							usersDirectoriesPayload["users-from-external-directories"] = v.(string)
						}
						authenticationSettingsPayload["users-directories"] = usersDirectoriesPayload
					}
					browserBasedAuthenticationSettingsPayload["authentication-settings"] = authenticationSettingsPayload
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings"); ok {
					browserBasedAuthenticationPortalSettingsPayload := make(map[string]interface{})

					if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility"); ok {

						accessibilityPayload := make(map[string]interface{})

						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.allow_access_from"); ok {
							accessibilityPayload["allow-access-from"] = v.(string)
						}
						if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.internal_access_settings"); ok {

							internalAccessSettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.internal_access_settings.0.dmz"); ok {
								internalAccessSettingsPayload["dmz"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.internal_access_settings.0.undefined"); ok {
								internalAccessSettingsPayload["undefined"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.internal_access_settings.0.vpn"); ok {
								internalAccessSettingsPayload["vpn"] = v.(bool)
							}
							accessibilityPayload["internal-access-settings"] = internalAccessSettingsPayload
						}
						browserBasedAuthenticationPortalSettingsPayload["accessibility"] = accessibilityPayload
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.certificate_settings"); ok {

						certificateSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.certificate_settings.0.base64_certificate"); ok {
							certificateSettingsPayload["base64-certificate"] = v.(string)
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.certificate_settings.0.base64_password"); ok {
							certificateSettingsPayload["base64-password"] = v.(string)
						}
						browserBasedAuthenticationPortalSettingsPayload["certificate-settings"] = certificateSettingsPayload
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.portal_web_settings"); ok {

						portalWebSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.portal_web_settings.0.aliases"); ok {
							portalWebSettingsPayload["aliases"] = v.(*schema.Set).List()
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.portal_web_settings.0.ip_address"); ok {
							portalWebSettingsPayload["ip-address"] = v.(string)
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.portal_web_settings.0.main_url"); ok {
							portalWebSettingsPayload["main-url"] = v.(string)
						}
						browserBasedAuthenticationPortalSettingsPayload["portal-web-settings"] = portalWebSettingsPayload
					}
					browserBasedAuthenticationSettingsPayload["browser-based-authentication-portal-settings"] = browserBasedAuthenticationPortalSettingsPayload
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
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings"); ok {
					authenticationSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.authentication_method"); ok {
						authenticationSettingsPayload["authentication-method"] = v.(string)
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.radius"); ok {
						authenticationSettingsPayload["radius"] = v.(string)
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories"); ok {

						usersDirectoriesPayload := make(map[string]interface{})

						if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories.0.external_user_profile"); ok {
							usersDirectoriesPayload["external-user-profile"] = v.(bool)
						}
						if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories.0.internal_users"); ok {
							usersDirectoriesPayload["internal-users"] = v.(bool)
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories.0.specific"); ok {
							usersDirectoriesPayload["specific"] = v.(*schema.Set).List()
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories.0.users_from_external_directories"); ok {
							usersDirectoriesPayload["users-from-external-directories"] = v.(string)
						}
						authenticationSettingsPayload["users-directories"] = usersDirectoriesPayload
					}
					identityAgentSettingsPayload["authentication-settings"] = authenticationSettingsPayload
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings"); ok {
					identityAgentPortalSettingsPayload := make(map[string]interface{})

					if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility"); ok {

						accessibilityPayload := make(map[string]interface{})

						if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.allow_access_from"); ok {
							accessibilityPayload["allow-access-from"] = v.(string)
						}
						if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.internal_access_settings"); ok {

							internalAccessSettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.internal_access_settings.0.dmz"); ok {
								internalAccessSettingsPayload["dmz"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.internal_access_settings.0.undefined"); ok {
								internalAccessSettingsPayload["undefined"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.internal_access_settings.0.vpn"); ok {
								internalAccessSettingsPayload["vpn"] = v.(bool)
							}
							accessibilityPayload["internal-access-settings"] = internalAccessSettingsPayload
						}
						identityAgentPortalSettingsPayload["accessibility"] = accessibilityPayload
					}
					identityAgentSettingsPayload["identity-agent-portal-settings"] = identityAgentPortalSettingsPayload
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
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings"); ok {
					authenticationSettingsPayload := make(map[string]interface{})

					if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories"); ok {

						usersDirectoriesPayload := make(map[string]interface{})

						if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories.0.external_user_profile"); ok {
							usersDirectoriesPayload["external-user-profile"] = v.(bool)
						}
						if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories.0.internal_users"); ok {
							usersDirectoriesPayload["internal-users"] = v.(bool)
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories.0.specific"); ok {
							usersDirectoriesPayload["specific"] = v.(*schema.Set).List()
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories.0.users_from_external_directories"); ok {
							usersDirectoriesPayload["users-from-external-directories"] = v.(string)
						}
						authenticationSettingsPayload["users-directories"] = usersDirectoriesPayload
					}
					identityCollectorSettingsPayload["authentication-settings"] = authenticationSettingsPayload
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions"); ok {
					clientAccessPermissionsPayload := make(map[string]interface{})

					if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility"); ok {

						accessibilityPayload := make(map[string]interface{})

						if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.allow_access_from"); ok {
							accessibilityPayload["allow-access-from"] = v.(string)
						}
						if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings"); ok {

							internalAccessSettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings.0.dmz"); ok {
								internalAccessSettingsPayload["dmz"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings.0.undefined"); ok {
								internalAccessSettingsPayload["undefined"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings.0.vpn"); ok {
								internalAccessSettingsPayload["vpn"] = v.(bool)
							}
							accessibilityPayload["internal-access-settings"] = internalAccessSettingsPayload
						}
						clientAccessPermissionsPayload["accessibility"] = accessibilityPayload
					}
					identityCollectorSettingsPayload["client-access-permissions"] = clientAccessPermissionsPayload
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
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.cache_mode"); ok {

					cacheModePayload := make(map[string]interface{})

					if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.cache_mode.0.override_profile"); ok {
						cacheModePayload["override-profile"] = v.(bool)
					}
					if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.cache_mode.0.value"); ok {
						cacheModePayload["value"] = v.(bool)
					}
					identitySharingSettingsPayload["cache-mode"] = cacheModePayload
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.cache_mode_duration"); ok {

					cacheModeDurationPayload := make(map[string]interface{})

					if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.cache_mode_duration.0.override_profile"); ok {
						cacheModeDurationPayload["override-profile"] = v.(bool)
					}
					if v, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.cache_mode_duration.0.value"); ok {
						cacheModeDurationPayload["value"] = v.(int)
					}
					identitySharingSettingsPayload["cache-mode-duration"] = cacheModeDurationPayload
				}
				if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.receive_from_infinity_identity"); ok {
					identitySharingSettingsPayload["receive-from-infinity-identity"] = v.(bool)
				}
				if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.scaled_sharing"); ok {
					identitySharingSettingsPayload["scaled-sharing"] = v.(bool)
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
			if v, ok := d.GetOk("identity_awareness_settings.0.identity_based_enforcement"); ok {
				identityAwarenessSettingsPayload["identity-based-enforcement"] = v.(string)
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

	if v, ok := d.GetOk("nat_settings"); ok {

		natSettingsList := v.([]interface{})

		if len(natSettingsList) > 0 {

			natSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOkExists("nat_settings.0.auto_rule"); ok {
				natSettingsPayload["auto-rule"] = v.(bool)
			}
			if v, ok := d.GetOk("nat_settings.0.ipv4_address"); ok {
				natSettingsPayload["ipv4-address"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.0.ipv6_address"); ok {
				natSettingsPayload["ipv6-address"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.0.hide_behind"); ok {
				natSettingsPayload["hide-behind"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.0.install_on"); ok {
				natSettingsPayload["install-on"] = v.(string)
			}
			if v, ok := d.GetOk("nat_settings.0.method"); ok {
				natSettingsPayload["method"] = v.(string)
			}
			if v, ok := d.GetOkExists("nat_settings.0.apply_control_connections"); ok {
				natSettingsPayload["apply-control-connections"] = v.(bool)
			}
			cluster["nat-settings"] = natSettingsPayload
		}
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
				if v, ok := d.GetOk("platform_portal_settings.0.portal_web_settings.0.ip_address"); ok {
					portalWebSettingsPayload["ip-address"] = v.(string)
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

	if v, ok := d.GetOk("proxy_settings"); ok {

		proxySettingsList := v.([]interface{})

		if len(proxySettingsList) > 0 {

			proxySettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("proxy_settings.0.use_custom_proxy"); ok {
				proxySettingsPayload["use-custom-proxy"] = v.(bool)
			}
			if v, ok := d.GetOk("proxy_settings.0.proxy_server"); ok {
				proxySettingsPayload["proxy-server"] = v.(string)
			}
			if v, ok := d.GetOk("proxy_settings.0.port"); ok {
				proxySettingsPayload["port"] = v.(int)
			}
			cluster["proxy-settings"] = proxySettingsPayload
		}
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
				if v, ok := d.GetOk("usercheck_portal_settings.0.portal_web_settings.0.ip_address"); ok {
					portalWebSettingsPayload["ip-address"] = v.(string)
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
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.action"); ok {
						antiSpoofingSettings["action"] = v.(string)
					}
					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.exclude_packets"); ok {
						antiSpoofingSettings["exclude-packets"] = v.(bool)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.excluded_network_name"); ok {
						antiSpoofingSettings["excluded-network-name"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.excluded_network_uid"); ok {
						antiSpoofingSettings["excluded-network-uid"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.spoof_tracking"); ok {
						antiSpoofingSettings["spoof-tracking"] = v.(string)
					}
					interfacePayload["anti-spoofing-settings"] = antiSpoofingSettings
				}
				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".dynamic_ip"); ok {
					interfacePayload["dynamic-ip"] = v.(bool)
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
					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone_settings.0.auto_calculated"); ok {
						securityZoneSettings["auto-calculated"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings.0.specific_zone"); ok {
						securityZoneSettings["specific-zone"] = v.(string)
					}
					interfacePayload["security-zone-settings"] = securityZoneSettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology"); ok {
					interfacePayload["topology"] = v.(string)
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings"); ok {
					topologySettings := make(map[string]interface{})

					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".topology_settings.0.interface_leads_to_dmz"); ok {
						topologySettings["interface-leads-to-dmz"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.0.ip_address_behind_this_interface"); ok {
						topologySettings["ip-address-behind-this-interface"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.0.specific_network"); ok {
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
				if v, ok := d.GetOkExists("members." + strconv.Itoa(i) + ".auto_generate_ip"); ok {
					memberPayload["auto-generate-ip"] = v.(bool)
				}
				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".trust_method"); ok {
					memberPayload["trust-method"] = v.(string)
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
							if v, ok := d.GetOkExists(memberInterfacePrefix + ".anti_spoofing"); ok {
								interfacePayload["anti-spoofing"] = v.(bool)
							}
							if v, ok := d.GetOkExists(memberInterfacePrefix + ".dynamic_ip"); ok {
								interfacePayload["dynamic-ip"] = v.(bool)
							}
							if v, ok := d.GetOkExists(memberInterfacePrefix + ".security_zone"); ok {
								interfacePayload["security-zone"] = v.(bool)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".topology"); ok {
								interfacePayload["topology"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv6_mask_length"); ok {
								interfacePayload["ipv6-mask-length"] = v.(string)
							}
							if _, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings"); ok {
								antiSpoofingSettingsPayload := make(map[string]interface{})
								if v, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings.0.action"); ok {
									antiSpoofingSettingsPayload["action"] = v.(string)
								}
								if v, ok := d.GetOkExists(memberInterfacePrefix + ".anti_spoofing_settings.0.exclude_packets"); ok {
									antiSpoofingSettingsPayload["exclude-packets"] = v.(bool)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings.0.excluded_network_name"); ok {
									antiSpoofingSettingsPayload["excluded-network-name"] = v.(string)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings.0.excluded_network_uid"); ok {
									antiSpoofingSettingsPayload["excluded-network-uid"] = v.(string)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings.0.spoof_tracking"); ok {
									antiSpoofingSettingsPayload["spoof-tracking"] = v.(string)
								}
								interfacePayload["anti-spoofing-settings"] = antiSpoofingSettingsPayload
							}
							if _, ok := d.GetOk(memberInterfacePrefix + ".security_zone_settings"); ok {
								securityZoneSettingsPayload := make(map[string]interface{})
								if v, ok := d.GetOkExists(memberInterfacePrefix + ".security_zone_settings.0.auto_calculated"); ok {
									securityZoneSettingsPayload["auto-calculated"] = v.(bool)
								}
								interfacePayload["security-zone-settings"] = securityZoneSettingsPayload
							}
							if _, ok := d.GetOk(memberInterfacePrefix + ".topology_settings"); ok {
								topologySettingsPayload := make(map[string]interface{})
								if v, ok := d.GetOkExists(memberInterfacePrefix + ".topology_settings.0.interface_leads_to_dmz"); ok {
									topologySettingsPayload["interface-leads-to-dmz"] = v.(bool)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".topology_settings.0.ip_address_behind_this_interface"); ok {
									topologySettingsPayload["ip-address-behind-this-interface"] = v.(string)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".topology_settings.0.specific_network"); ok {
									topologySettingsPayload["specific-network"] = v.(string)
								}
								interfacePayload["topology-settings"] = topologySettingsPayload
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

	if v, ok := d.GetOk("firewall_settings"); ok {

		firewallSettingsList := v.([]interface{})

		if len(firewallSettingsList) > 0 {

			firewallSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOkExists("firewall_settings.0.auto_calculate_connections_hash_table_size_and_memory_pool"); ok {
				firewallSettingsPayload["auto-calculate-connections-hash-table-size-and-memory-pool"] = v.(bool)
			}
			if v, ok := d.GetOkExists("firewall_settings.0.auto_maximum_limit_for_concurrent_connections"); ok {
				firewallSettingsPayload["auto-maximum-limit-for-concurrent-connections"] = v.(bool)
			}
			if v, ok := d.GetOk("firewall_settings.0.connections_hash_size"); ok {
				firewallSettingsPayload["connections-hash-size"] = v.(int)
			}
			if v, ok := d.GetOk("firewall_settings.0.maximum_limit_for_concurrent_connections"); ok {
				firewallSettingsPayload["maximum-limit-for-concurrent-connections"] = v.(int)
			}
			if v, ok := d.GetOk("firewall_settings.0.maximum_memory_pool_size"); ok {
				firewallSettingsPayload["maximum-memory-pool-size"] = v.(int)
			}
			if v, ok := d.GetOk("firewall_settings.0.memory_pool_size"); ok {
				firewallSettingsPayload["memory-pool-size"] = v.(int)
			}
			cluster["firewall-settings"] = firewallSettingsPayload
		}
	}

	// VPN settings
	if v, ok := d.GetOk("vpn_settings"); ok {

		vpnSettingsList := v.([]interface{})

		if len(vpnSettingsList) > 0 {

			vpnSettingsPayload := make(map[string]interface{})

			if _, ok := d.GetOk("vpn_settings.0.advanced"); ok {

				advancedPayload := make(map[string]interface{})

				if v, ok := d.GetOk("vpn_settings.0.advanced.0.tunnel_sharing_mode"); ok {
					advancedPayload["tunnel-sharing-mode"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.advanced.0.shutdown_on_gateway_restart"); ok {
					advancedPayload["shutdown-on-gateway-restart"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.advanced.0.enable_wire_mode"); ok {
					advancedPayload["enable-wire-mode"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.advanced.0.wire_mode_interfaces"); ok {
					advancedPayload["wire-mode-interfaces"] = v.(*schema.Set).List()
				}
				if v, ok := d.GetOk("vpn_settings.0.advanced.0.enable_wire_mode_log_traffic"); ok {
					advancedPayload["enable-wire-mode-log-traffic"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.advanced.0.enable_nat_traversal"); ok {
					advancedPayload["enable-nat-traversal"] = strconv.FormatBool(v.(bool))
				}
				vpnSettingsPayload["advanced"] = advancedPayload
			}
			if _, ok := d.GetOk("vpn_settings.0.authentication"); ok {

				authenticationPayload := make(map[string]interface{})

				if v, ok := d.GetOk("vpn_settings.0.authentication.0.authentication_clients"); ok {
					authenticationPayload["authentication-clients"] = v.(*schema.Set).List()
				}
				if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client"); ok {

					singleAuthenticationClientPayload := make(map[string]interface{})

					if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.single_authentication_client.0.enabled"); ok {
						singleAuthenticationClientPayload["enabled"] = v.(bool)
					}
					if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.single_authentication_client.0.allow_multiple_authentication_clients"); ok {
						singleAuthenticationClientPayload["allow-multiple-authentication-clients"] = v.(bool)
					}
					if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.display_name"); ok {
						singleAuthenticationClientPayload["display-name"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.method"); ok {
						singleAuthenticationClientPayload["method"] = v.(string)
					}
					if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.secur_id"); ok {

						securIdPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.secur_id.0.server"); ok {
							securIdPayload["server"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.secur_id.0.token_card_type"); ok {
							securIdPayload["token-card-type"] = v.(string)
						}
						singleAuthenticationClientPayload["secur-id"] = securIdPayload
					}
					if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.radius"); ok {

						radiusPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.radius.0.server"); ok {
							radiusPayload["server"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.single_authentication_client.0.radius.0.ask_user_password"); ok {
							radiusPayload["ask-user-password"] = v.(bool)
						}
						singleAuthenticationClientPayload["radius"] = radiusPayload
					}
					if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate"); ok {

						personalCertificatePayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.fetch_username_from"); ok {
							personalCertificatePayload["fetch-username-from"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.storage_type"); ok {
							personalCertificatePayload["storage-type"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.source"); ok {
							personalCertificatePayload["source"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.dn_part"); ok {
							personalCertificatePayload["dn-part"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.dn_concurrence"); ok {
							personalCertificatePayload["dn-concurrence"] = v.(int)
						}
						singleAuthenticationClientPayload["personal-certificate"] = personalCertificatePayload
					}
					if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.client_display_settings"); ok {

						clientDisplaySettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.client_display_settings.0.headline"); ok {
							clientDisplaySettingsPayload["headline"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.client_display_settings.0.username_label"); ok {
							clientDisplaySettingsPayload["username-label"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.client_display_settings.0.password_label"); ok {
							clientDisplaySettingsPayload["password-label"] = v.(string)
						}
						singleAuthenticationClientPayload["client-display-settings"] = clientDisplaySettingsPayload
					}
					authenticationPayload["single-authentication-client"] = singleAuthenticationClientPayload
				}
				if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.override_global_dynamic_id_settings"); ok {
					authenticationPayload["override-global-dynamic-id-settings"] = v.(bool)
				}
				if _, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings"); ok {

					dynamicIdSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_and_email_settings"); ok {
						dynamicIdSettingsPayload["sms-provider-and-email-settings"] = v.(string)
					}
					if _, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_credentials"); ok {

						smsProviderCredentialsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_credentials.0.username"); ok {
							smsProviderCredentialsPayload["username"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_credentials.0.password"); ok {
							smsProviderCredentialsPayload["password"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_credentials.0.api_id"); ok {
							smsProviderCredentialsPayload["api-id"] = v.(string)
						}
						dynamicIdSettingsPayload["sms-provider-credentials"] = smsProviderCredentialsPayload
					}
					if _, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings"); ok {

						advancedSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.dynamic_id_message"); ok {
							advancedSettingsPayload["dynamic-id-message"] = v.(string)
						}
						if _, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.otp_settings"); ok {

							otpSettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.otp_settings.0.length"); ok {
								otpSettingsPayload["length"] = v.(int)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.otp_settings.0.expiration"); ok {
								otpSettingsPayload["expiration"] = v.(int)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.otp_settings.0.max_attempts"); ok {
								otpSettingsPayload["max-attempts"] = v.(int)
							}
							advancedSettingsPayload["otp-settings"] = otpSettingsPayload
						}
						if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.enable_display_user_details"); ok {
							advancedSettingsPayload["enable-display-user-details"] = v.(bool)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.country_code"); ok {
							advancedSettingsPayload["country-code"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.user_details_retrieval"); ok {
							advancedSettingsPayload["user-details-retrieval"] = v.(string)
						}
						dynamicIdSettingsPayload["advanced-settings"] = advancedSettingsPayload
					}
					authenticationPayload["dynamic-id-settings"] = dynamicIdSettingsPayload
				}
				if v, ok := d.GetOk("vpn_settings.0.authentication.0.send_machine_certificate"); ok {
					authenticationPayload["send-machine-certificate"] = v.(string)
				}
				vpnSettingsPayload["authentication"] = authenticationPayload
			}
			if _, ok := d.GetOk("vpn_settings.0.certificates"); ok {

				certificatesPayload := make(map[string]interface{})

				if v, ok := d.GetOk("vpn_settings.0.certificates.0.name"); ok {
					certificatesPayload["name"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.certificates.0.certificate_authority"); ok {
					certificatesPayload["certificate-authority"] = v.(string)
				}
				if _, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment"); ok {

					enrollmentPayload := make(map[string]interface{})

					if _, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings"); ok {

						enrollmentSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings.0.distinguished_name"); ok {
							enrollmentSettingsPayload["distinguished-name"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings.0.alternate_names"); ok {

							alternateNamesList := v.([]interface{})

							if len(alternateNamesList) > 0 {

								var alternateNamesPayload []map[string]interface{}

								for j := range alternateNamesList {

									alternateNamesMapToAdd := make(map[string]interface{})

									if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings.0.alternate_names." + strconv.Itoa(j) + ".name_type"); ok {
										alternateNamesMapToAdd["name-type"] = v.(string)
									}
									if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings.0.alternate_names." + strconv.Itoa(j) + ".value"); ok {
										alternateNamesMapToAdd["value"] = v.(string)
									}
									alternateNamesPayload = append(alternateNamesPayload, alternateNamesMapToAdd)
								}
								enrollmentSettingsPayload["alternate-names"] = alternateNamesPayload
							}
						}
						enrollmentPayload["enrollment-settings"] = enrollmentSettingsPayload
					}
					if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_type"); ok {
						enrollmentPayload["enrollment-type"] = v.(string)
					}
					certificatesPayload["enrollment"] = enrollmentPayload
				}
				if v, ok := d.GetOk("vpn_settings.0.certificates.0.stored_at"); ok {
					certificatesPayload["stored-at"] = v.(string)
				}
				vpnSettingsPayload["certificates"] = certificatesPayload
			}
			if _, ok := d.GetOk("vpn_settings.0.exported_routes"); ok {

				exportedRoutesPayload := make(map[string]interface{})

				if v, ok := d.GetOk("vpn_settings.0.exported_routes.0.internal_interfaces"); ok {
					exportedRoutesPayload["internal-interfaces"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.exported_routes.0.static_routes"); ok {
					exportedRoutesPayload["static-routes"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.exported_routes.0.custom_routes"); ok {
					exportedRoutesPayload["custom-routes"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.exported_routes.0.custom_routes_object"); ok {
					exportedRoutesPayload["custom-routes-object"] = v.(string)
				}
				vpnSettingsPayload["exported-routes"] = exportedRoutesPayload
			}
			if _, ok := d.GetOk("vpn_settings.0.link_selection"); ok {

				linkSelectionPayload := make(map[string]interface{})

				if v, ok := d.GetOk("vpn_settings.0.link_selection.0.ip_selection"); ok {
					linkSelectionPayload["ip-selection"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.link_selection.0.ip_address"); ok {
					linkSelectionPayload["ip-address"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.link_selection.0.dns_resolving_hostname"); ok {
					linkSelectionPayload["dns-resolving-hostname"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.link_selection.0.route_selection_method"); ok {
					linkSelectionPayload["route-selection-method"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.link_selection.0.responding_traffic"); ok {
					linkSelectionPayload["responding-traffic"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.link_selection.0.source_ip_selection"); ok {
					linkSelectionPayload["source-ip-selection"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.link_selection.0.selected_ip"); ok {
					linkSelectionPayload["selected-ip"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.link_selection.0.outgoing_link_tracking"); ok {
					linkSelectionPayload["outgoing-link-tracking"] = v.(string)
				}
				if _, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings"); ok {

					probingSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings.0.probed_interfaces"); ok {
						probingSettingsPayload["probed-interfaces"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings.0.probed_interface_list"); ok {
						probingSettingsPayload["probed-interface-list"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOkExists("vpn_settings.0.link_selection.0.probing_settings.0.use_primary_address"); ok {
						probingSettingsPayload["use-primary-address"] = v.(bool)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings.0.primary_address"); ok {
						probingSettingsPayload["primary-address"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings.0.probing_method"); ok {
						probingSettingsPayload["probing-method"] = v.(string)
					}
					linkSelectionPayload["probing-settings"] = probingSettingsPayload
				}
				vpnSettingsPayload["link-selection"] = linkSelectionPayload
			}
			if v, ok := d.GetOk("vpn_settings.0.maximum_concurrent_ike_negotiations"); ok {
				vpnSettingsPayload["maximum-concurrent-ike-negotiations"] = v.(int)
			}
			if v, ok := d.GetOk("vpn_settings.0.maximum_concurrent_tunnels"); ok {
				vpnSettingsPayload["maximum-concurrent-tunnels"] = v.(int)
			}
			if _, ok := d.GetOk("vpn_settings.0.office_mode"); ok {

				officeModePayload := make(map[string]interface{})

				if v, ok := d.GetOk("vpn_settings.0.office_mode.0.mode"); ok {
					officeModePayload["mode"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.office_mode.0.group"); ok {
					officeModePayload["group"] = v.(string)
				}
				if _, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from"); ok {

					allocateIpAddressFromPayload := make(map[string]interface{})

					if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.radius_server"); ok {
						allocateIpAddressFromPayload["radius-server"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.use_allocate_method"); ok {
						allocateIpAddressFromPayload["use-allocate-method"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.allocate_method"); ok {
						allocateIpAddressFromPayload["allocate-method"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.manual_network"); ok {
						allocateIpAddressFromPayload["manual-network"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.dhcp_server"); ok {
						allocateIpAddressFromPayload["dhcp-server"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.virtual_ip_address"); ok {
						allocateIpAddressFromPayload["virtual-ip-address"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.dhcp_mac_address"); ok {
						allocateIpAddressFromPayload["dhcp-mac-address"] = v.(string)
					}
					if _, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters"); ok {

						optionalParametersPayload := make(map[string]interface{})

						if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_primary_dns_server"); ok {
							optionalParametersPayload["use-primary-dns-server"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.primary_dns_server"); ok {
							optionalParametersPayload["primary-dns-server"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_first_backup_dns_server"); ok {
							optionalParametersPayload["use-first-backup-dns-server"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.first_backup_dns_server"); ok {
							optionalParametersPayload["first-backup-dns-server"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_second_backup_dns_server"); ok {
							optionalParametersPayload["use-second-backup-dns-server"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.second_backup_dns_server"); ok {
							optionalParametersPayload["second-backup-dns-server"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.dns_suffixes"); ok {
							optionalParametersPayload["dns-suffixes"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_primary_wins_server"); ok {
							optionalParametersPayload["use-primary-wins-server"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.primary_wins_server"); ok {
							optionalParametersPayload["primary-wins-server"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_first_backup_wins_server"); ok {
							optionalParametersPayload["use-first-backup-wins-server"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.first_backup_wins_server"); ok {
							optionalParametersPayload["first-backup-wins-server"] = v.(string)
						}
						if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_second_backup_wins_server"); ok {
							optionalParametersPayload["use-second-backup-wins-server"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.second_backup_wins_server"); ok {
							optionalParametersPayload["second-backup-wins-server"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.ip_lease_duration"); ok {
							optionalParametersPayload["ip-lease-duration"] = v
						}
						allocateIpAddressFromPayload["optional-parameters"] = optionalParametersPayload
					}
					officeModePayload["allocate-ip-address-from"] = allocateIpAddressFromPayload
				}
				if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.support_multiple_interfaces"); ok {
					officeModePayload["support-multiple-interfaces"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.perform_anti_spoofing"); ok {
					officeModePayload["perform-anti-spoofing"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.office_mode.0.anti_spoofing_additional_addresses"); ok {
					officeModePayload["anti-spoofing-additional-addresses"] = v.(string)
				}
				vpnSettingsPayload["office-mode"] = officeModePayload
			}
			if _, ok := d.GetOk("vpn_settings.0.remote_access"); ok {

				remoteAccessPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("vpn_settings.0.remote_access.0.support_l2tp"); ok {
					remoteAccessPayload["support-l2tp"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.remote_access.0.l2tp_auth_method"); ok {
					remoteAccessPayload["l2tp-auth-method"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.remote_access.0.l2tp_certificate"); ok {
					remoteAccessPayload["l2tp-certificate"] = v.(string)
				}
				if v, ok := d.GetOkExists("vpn_settings.0.remote_access.0.allow_vpn_clients_to_route_traffic"); ok {
					remoteAccessPayload["allow-vpn-clients-to-route-traffic"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOkExists("vpn_settings.0.remote_access.0.support_nat_traversal_mechanism"); ok {
					remoteAccessPayload["support-nat-traversal-mechanism"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.remote_access.0.nat_traversal_service"); ok {
					remoteAccessPayload["nat-traversal-service"] = v.(string)
				}
				if v, ok := d.GetOkExists("vpn_settings.0.remote_access.0.support_visitor_mode"); ok {
					remoteAccessPayload["support-visitor-mode"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.remote_access.0.visitor_mode_service"); ok {
					remoteAccessPayload["visitor-mode-service"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.remote_access.0.visitor_mode_interface"); ok {
					remoteAccessPayload["visitor-mode-interface"] = v.(string)
				}
				vpnSettingsPayload["remote-access"] = remoteAccessPayload
			}
			if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings"); ok {

				samlPortalSettingsPayload := make(map[string]interface{})

				if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.portal_web_settings"); ok {

					portalWebSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.portal_web_settings.0.aliases"); ok {
						portalWebSettingsPayload["aliases"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.portal_web_settings.0.ip_address"); ok {
						portalWebSettingsPayload["ip-address"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.portal_web_settings.0.main_url"); ok {
						portalWebSettingsPayload["main-url"] = v.(string)
					}
					samlPortalSettingsPayload["portal-web-settings"] = portalWebSettingsPayload
				}
				if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.certificate_settings"); ok {

					certificateSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.certificate_settings.0.base64_certificate"); ok {
						certificateSettingsPayload["base64-certificate"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.certificate_settings.0.base64_password"); ok {
						certificateSettingsPayload["base64-password"] = v.(string)
					}
					samlPortalSettingsPayload["certificate-settings"] = certificateSettingsPayload
				}
				if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.accessibility"); ok {

					accessibilityPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.accessibility.0.allow_access_from"); ok {
						accessibilityPayload["allow-access-from"] = v.(string)
					}
					if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.accessibility.0.internal_access_settings"); ok {

						internalAccessSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.accessibility.0.internal_access_settings.0.undefined"); ok {
							internalAccessSettingsPayload["undefined"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.accessibility.0.internal_access_settings.0.dmz"); ok {
							internalAccessSettingsPayload["dmz"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.accessibility.0.internal_access_settings.0.vpn"); ok {
							internalAccessSettingsPayload["vpn"] = strconv.FormatBool(v.(bool))
						}
						accessibilityPayload["internal-access-settings"] = internalAccessSettingsPayload
					}
					samlPortalSettingsPayload["accessibility"] = accessibilityPayload
				}
				vpnSettingsPayload["saml-portal-settings"] = samlPortalSettingsPayload
			}
			if _, ok := d.GetOk("vpn_settings.0.vpn_clients"); ok {

				vpnClientsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("vpn_settings.0.vpn_clients.0.enable_endpoint_security_vpn"); ok {
					vpnClientsPayload["enable-endpoint-security-vpn"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.vpn_clients.0.enable_cp_mobile_for_windows"); ok {
					vpnClientsPayload["enable-cp-mobile-for-windows"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.vpn_clients.0.enable_secu_remote"); ok {
					vpnClientsPayload["enable-secu-remote"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.vpn_clients.0.enable_capsule_vpn_connect"); ok {
					vpnClientsPayload["enable-capsule-vpn-connect"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.vpn_clients.0.enable_ssl_network_extender"); ok {
					vpnClientsPayload["enable-ssl-network-extender"] = strconv.FormatBool(v.(bool))
				}
				if v, ok := d.GetOk("vpn_settings.0.vpn_clients.0.gateway_authentication_certificate"); ok {
					vpnClientsPayload["gateway-authentication-certificate"] = v.(string)
				}
				vpnSettingsPayload["vpn-clients"] = vpnClientsPayload
			}
			if v, ok := d.GetOk("vpn_settings.0.vpn_domain"); ok {
				vpnSettingsPayload["vpn-domain"] = v.(string)
			}
			if v, ok := d.GetOkExists("vpn_settings.0.vpn_domain_exclude_external_ip_addresses"); ok {
				vpnSettingsPayload["vpn-domain-exclude-external-ip-addresses"] = v.(bool)
			}
			if v, ok := d.GetOk("vpn_settings.0.vpn_domain_type"); ok {
				vpnSettingsPayload["vpn-domain-type"] = v.(string)
			}
			if v, ok := d.GetOk("vpn_settings.0.enable_clientless_vpn"); ok {
				vpnSettingsPayload["enable-clientless-vpn"] = v.(bool)
			}
			if _, ok := d.GetOk("vpn_settings.0.clientless_vpn_settings"); ok {

				clientlessVpnSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("vpn_settings.0.clientless_vpn_settings.0.certificate_gateway_authentication"); ok {
					clientlessVpnSettingsPayload["certificate-gateway-authentication"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.clientless_vpn_settings.0.client_authentication"); ok {
					clientlessVpnSettingsPayload["client-authentication"] = v.(string)
				}
				if v, ok := d.GetOk("vpn_settings.0.clientless_vpn_settings.0.concurrent_servers_or_processes"); ok {
					clientlessVpnSettingsPayload["concurrent-servers-or-processes"] = v
				}
				if v, ok := d.GetOk("vpn_settings.0.clientless_vpn_settings.0.accept_only_3des"); ok {
					clientlessVpnSettingsPayload["accept-only-3des"] = strconv.FormatBool(v.(bool))
				}
				vpnSettingsPayload["clientless-vpn-settings"] = clientlessVpnSettingsPayload
			}
			cluster["vpn-settings"] = vpnSettingsPayload
		}
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

	if v, ok := d.GetOkExists("anti_spam_and_email_security"); ok {
		cluster["anti-spam-and-email-security"] = v.(bool)
	}
	if v, ok := d.GetOk("auto_topology_custom_recalculation_time"); ok {
		cluster["auto-topology-custom-recalculation-time"] = v.(int)
	}
	if v, ok := d.GetOkExists("auto_topology_use_custom_recalculation_time"); ok {
		cluster["auto-topology-use-custom-recalculation-time"] = v.(bool)
	}
	if v, ok := d.GetOkExists("data_loss_prevention"); ok {
		cluster["data-loss-prevention"] = v.(bool)
	}
	if v, ok := d.GetOkExists("mobile_access"); ok {
		cluster["mobile-access"] = v.(bool)
	}
	if v, ok := d.GetOkExists("monitoring"); ok {
		cluster["monitoring"] = v.(bool)
	}
	if v, ok := d.GetOkExists("policy_server"); ok {
		cluster["policy-server"] = v.(bool)
	}
	if v, ok := d.GetOkExists("rtm_counters_report"); ok {
		cluster["rtm-counters-report"] = v.(bool)
	}
	if v, ok := d.GetOkExists("rtm_traffic_report"); ok {
		cluster["rtm-traffic-report"] = v.(bool)
	}
	if v, ok := d.GetOkExists("rtm_traffic_report_per_connection"); ok {
		cluster["rtm-traffic-report-per-connection"] = v.(bool)
	}
	if v, ok := d.GetOkExists("threat_extraction"); ok {
		cluster["threat-extraction"] = v.(bool)
	}
	if v, ok := d.GetOk("threat_prevention_mode"); ok {
		cluster["threat-prevention-mode"] = v.(string)
	}
	if v, ok := d.GetOkExists("workforce_ai"); ok {
		cluster["workforce-ai"] = v.(bool)
	}
	if _, ok := d.GetOk("application_control_and_url_filtering_settings"); ok {

		applicationControlAndUrlFilteringSettingsPayload := make(map[string]interface{})

		if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.global_settings_mode"); ok {
			applicationControlAndUrlFilteringSettingsPayload["global-settings-mode"] = v.(string)
		}
		if _, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings"); ok {

			overrideGlobalSettingsPayload := make(map[string]interface{})

			if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.fail_mode"); ok {
				overrideGlobalSettingsPayload["fail-mode"] = v.(string)
			}
			if _, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization"); ok {

				websiteCategorizationPayload := make(map[string]interface{})

				if _, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization.0.custom_mode"); ok {

					customModePayload := make(map[string]interface{})

					if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization.0.custom_mode.0.social_networking_widgets"); ok {
						customModePayload["social-networking-widgets"] = v.(string)
					}
					if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization.0.custom_mode.0.url_filtering"); ok {
						customModePayload["url-filtering"] = v.(string)
					}
					websiteCategorizationPayload["custom-mode"] = customModePayload
				}
				if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization.0.mode"); ok {
					websiteCategorizationPayload["mode"] = v.(string)
				}
				overrideGlobalSettingsPayload["website-categorization"] = websiteCategorizationPayload
			}
			applicationControlAndUrlFilteringSettingsPayload["override-global-settings"] = overrideGlobalSettingsPayload
		}
		cluster["application-control-and-url-filtering-settings"] = applicationControlAndUrlFilteringSettingsPayload
	}
	if _, ok := d.GetOk("cluster_settings"); ok {

		clusterSettingsPayload := make(map[string]interface{})

		if v, ok := d.GetOk("cluster_settings.0.member_recovery_mode"); ok {
			clusterSettingsPayload["member-recovery-mode"] = v.(string)
		}
		if _, ok := d.GetOk("cluster_settings.0.state_synchronization"); ok {

			stateSynchronizationPayload := make(map[string]interface{})

			if v, ok := d.GetOkExists("cluster_settings.0.state_synchronization.0.delayed"); ok {
				stateSynchronizationPayload["delayed"] = v.(bool)
			}
			if v, ok := d.GetOk("cluster_settings.0.state_synchronization.0.delayed_seconds"); ok {
				stateSynchronizationPayload["delayed-seconds"] = v.(int)
			}
			if v, ok := d.GetOkExists("cluster_settings.0.state_synchronization.0.enabled"); ok {
				stateSynchronizationPayload["enabled"] = v.(bool)
			}
			clusterSettingsPayload["state-synchronization"] = stateSynchronizationPayload
		}
		if v, ok := d.GetOk("cluster_settings.0.track_changes_of_cluster_members"); ok {
			clusterSettingsPayload["track-changes-of-cluster-members"] = v.(string)
		}
		if v, ok := d.GetOkExists("cluster_settings.0.use_virtual_mac"); ok {
			clusterSettingsPayload["use-virtual-mac"] = v.(bool)
		}
		cluster["cluster-settings"] = clusterSettingsPayload
	}
	if _, ok := d.GetOk("communication_with_servers_behind_nat"); ok {

		communicationWithServersBehindNatPayload := make(map[string]interface{})

		if v, ok := d.GetOkExists("communication_with_servers_behind_nat.0.override_profile"); ok {
			communicationWithServersBehindNatPayload["override-profile"] = v.(bool)
		}
		if v, ok := d.GetOk("communication_with_servers_behind_nat.0.value"); ok {
			communicationWithServersBehindNatPayload["value"] = v.(string)
		}
		cluster["communication-with-servers-behind-nat"] = communicationWithServersBehindNatPayload
	}
	if _, ok := d.GetOk("zero_phishing_settings"); ok {

		zeroPhishingSettingsPayload := make(map[string]interface{})

		if v, ok := d.GetOk("zero_phishing_settings.0.gateway_fqdn_mode"); ok {
			zeroPhishingSettingsPayload["gateway-fqdn-mode"] = v.(string)
		}
		if v, ok := d.GetOk("zero_phishing_settings.0.manual_fqdn"); ok {
			zeroPhishingSettingsPayload["manual-fqdn"] = v.(string)
		}
		cluster["zero-phishing-settings"] = zeroPhishingSettingsPayload
	}
	if _, ok := d.GetOk("logs_settings"); ok {
		logsSettingsPayload := make(map[string]interface{})
		if v, ok := d.GetOkExists("logs_settings.0.alert_when_free_disk_space_below"); ok {
			logsSettingsPayload["alert-when-free-disk-space-below"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.alert_when_free_disk_space_below_threshold"); ok {
			logsSettingsPayload["alert-when-free-disk-space-below-threshold"] = v.(int)
		}
		if v, ok := d.GetOk("logs_settings.0.alert_when_free_disk_space_below_type"); ok {
			logsSettingsPayload["alert-when-free-disk-space-below-type"] = v.(string)
		}
		if v, ok := d.GetOkExists("logs_settings.0.before_delete_keep_logs_from_the_last_days"); ok {
			logsSettingsPayload["before-delete-keep-logs-from-the-last-days"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.before_delete_keep_logs_from_the_last_days_threshold"); ok {
			logsSettingsPayload["before-delete-keep-logs-from-the-last-days-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.0.before_delete_run_script"); ok {
			logsSettingsPayload["before-delete-run-script"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.before_delete_run_script_command"); ok {
			logsSettingsPayload["before-delete-run-script-command"] = v.(string)
		}
		if v, ok := d.GetOkExists("logs_settings.0.delete_index_files_older_than_days"); ok {
			logsSettingsPayload["delete-index-files-older-than-days"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.delete_index_files_older_than_days_threshold"); ok {
			logsSettingsPayload["delete-index-files-older-than-days-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.0.delete_index_files_when_index_size_above"); ok {
			logsSettingsPayload["delete-index-files-when-index-size-above"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.delete_index_files_when_index_size_above_threshold"); ok {
			logsSettingsPayload["delete-index-files-when-index-size-above-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.0.delete_when_free_disk_space_below"); ok {
			logsSettingsPayload["delete-when-free-disk-space-below"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.delete_when_free_disk_space_below_threshold"); ok {
			logsSettingsPayload["delete-when-free-disk-space-below-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.0.detect_new_citrix_ica_application_names"); ok {
			logsSettingsPayload["detect-new-citrix-ica-application-names"] = v.(bool)
		}
		if v, ok := d.GetOkExists("logs_settings.0.distribute_logs_between_all_active_servers"); ok {
			logsSettingsPayload["distribute-logs-between-all-active-servers"] = v.(bool)
		}
		if v, ok := d.GetOkExists("logs_settings.0.forward_logs_to_log_server"); ok {
			logsSettingsPayload["forward-logs-to-log-server"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.forward_logs_to_log_server_name"); ok {
			logsSettingsPayload["forward-logs-to-log-server-name"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.0.forward_logs_to_log_server_schedule_name"); ok {
			logsSettingsPayload["forward-logs-to-log-server-schedule-name"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.0.free_disk_space_metrics"); ok {
			logsSettingsPayload["free-disk-space-metrics"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.0.include_tcp_state_information"); ok {
			logsSettingsPayload["include-tcp-state-information"] = v.(string)
		}
		if v, ok := d.GetOkExists("logs_settings.0.perform_log_rotate_before_log_forwarding"); ok {
			logsSettingsPayload["perform-log-rotate-before-log-forwarding"] = v.(bool)
		}
		if v, ok := d.GetOkExists("logs_settings.0.reject_connections_when_free_disk_space_below_threshold"); ok {
			logsSettingsPayload["reject-connections-when-free-disk-space-below-threshold"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.reserve_for_packet_capture_metrics"); ok {
			logsSettingsPayload["reserve-for-packet-capture-metrics"] = v.(string)
		}
		if v, ok := d.GetOk("logs_settings.0.reserve_for_packet_capture_threshold"); ok {
			logsSettingsPayload["reserve-for-packet-capture-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.0.rotate_log_by_file_size"); ok {
			logsSettingsPayload["rotate-log-by-file-size"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.rotate_log_file_size_threshold"); ok {
			logsSettingsPayload["rotate-log-file-size-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.0.rotate_log_on_schedule"); ok {
			logsSettingsPayload["rotate-log-on-schedule"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.rotate_log_schedule_name"); ok {
			logsSettingsPayload["rotate-log-schedule-name"] = v.(string)
		}
		if v, ok := d.GetOkExists("logs_settings.0.stop_logging_when_free_disk_space_below"); ok {
			logsSettingsPayload["stop-logging-when-free-disk-space-below"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.stop_logging_when_free_disk_space_below_threshold"); ok {
			logsSettingsPayload["stop-logging-when-free-disk-space-below-threshold"] = v.(int)
		}
		if v, ok := d.GetOkExists("logs_settings.0.turn_on_qos_logging"); ok {
			logsSettingsPayload["turn-on-qos-logging"] = v.(bool)
		}
		if v, ok := d.GetOk("logs_settings.0.update_account_log_every"); ok {
			logsSettingsPayload["update-account-log-every"] = v.(int)
		}
		cluster["logs-settings"] = logsSettingsPayload
	}

	if v, ok := d.GetOkExists("show_portals_certificate"); ok {
		cluster["show-portals-certificate"] = v.(bool)
	}

	addClusterRes, err := client.ApiCall("add-simple-cluster", cluster, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !addClusterRes.Success {
		if addClusterRes.ErrorMsg != "" {
			return fmt.Errorf("%s", addClusterRes.ErrorMsg)
		}
		msg := createTaskFailMessage("add-simple-cluster", addClusterRes.GetData())
		return fmt.Errorf("%s", msg)
	}

	// add-simple-cluster returns task-id. Call show-simple-cluster for object uid.
	showClusterRes, err := client.ApiCall("show-simple-cluster", map[string]interface{}{"name": d.Get("name")}, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showClusterRes.Success {
		return fmt.Errorf("%s", showClusterRes.ErrorMsg)
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
		return fmt.Errorf("%s", err.Error())
	}
	if !showClusterRes.Success {
		if objectNotFound(showClusterRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showClusterRes.ErrorMsg)
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
					return fmt.Errorf("%s", err.Error())
				}
				if !showClusterRes.Success {
					return fmt.Errorf("%s", showClusterRes.ErrorMsg)
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
			if v, ok := httpsInspectionMap["bypass-on-client-failure"]; ok {

				bypassOnClientFailureMap, ok := v.(map[string]interface{})
				if ok {
					bypassOnClientFailureMapToReturn := make(map[string]interface{})

					if v, _ := bypassOnClientFailureMap["override-profile"]; v != nil {
						bypassOnClientFailureMapToReturn["override_profile"] = v
					}
					if v, _ := bypassOnClientFailureMap["value"]; v != nil {
						bypassOnClientFailureMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["bypass_on_client_failure"] = []interface{}{bypassOnClientFailureMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["bypass-under-load"]; ok {

				bypassUnderLoadMap, ok := v.(map[string]interface{})
				if ok {
					bypassUnderLoadMapToReturn := make(map[string]interface{})

					if v, _ := bypassUnderLoadMap["value"]; v != nil {
						bypassUnderLoadMapToReturn["value"] = v
					}
					httpsInspectionMapToReturn["bypass_under_load"] = []interface{}{bypassUnderLoadMapToReturn}
				}
			}
			if v, ok := httpsInspectionMap["outbound-certificate"]; ok {

				outboundCertificateMap, ok := v.(map[string]interface{})
				if ok {
					outboundCertificateMapToReturn := make(map[string]interface{})

					if v, _ := outboundCertificateMap["override-profile"]; v != nil {
						outboundCertificateMapToReturn["override_profile"] = v
					}
					if v, _ := outboundCertificateMap["value"]; v != nil {
						outboundCertificateMapToReturn["value"] = v.(map[string]interface{})["name"]
					}
					httpsInspectionMapToReturn["outbound_certificate"] = []interface{}{outboundCertificateMapToReturn}
				}
			}
			if v, _ := httpsInspectionMap["deployment-mode"]; v != nil {
				httpsInspectionMapToReturn["deployment_mode"] = v
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
						authenticationSettingsMap := v.(map[string]interface{})
						authenticationSettingsMapToReturn := make(map[string]interface{})
						if v, _ := authenticationSettingsMap["authentication-method"]; v != nil {
							authenticationSettingsMapToReturn["authentication_method"] = v
						}
						if v, _ := authenticationSettingsMap["identity-provider"]; v != nil {
							authenticationSettingsMapToReturn["identity_provider"] = v
						}
						if v, _ := authenticationSettingsMap["radius"]; v != nil {
							authenticationSettingsMapToReturn["radius"] = v
						}
						if v, _ := authenticationSettingsMap["users-directories"]; v != nil {
							usersDirectoriesMap := v.(map[string]interface{})
							usersDirectoriesMapToReturn := make(map[string]interface{})
							if v, _ := usersDirectoriesMap["external-user-profile"]; v != nil {
								usersDirectoriesMapToReturn["external_user_profile"] = v
							}
							if v, _ := usersDirectoriesMap["internal-users"]; v != nil {
								usersDirectoriesMapToReturn["internal_users"] = v
							}
							if v, _ := usersDirectoriesMap["specific"]; v != nil {
								usersDirectoriesMapToReturn["specific"] = v
							}
							if v, _ := usersDirectoriesMap["users-from-external-directories"]; v != nil {
								usersDirectoriesMapToReturn["users_from_external_directories"] = v
							}
							authenticationSettingsMapToReturn["users_directories"] = []interface{}{usersDirectoriesMapToReturn}
						}
						browserBasedAuthenticationSettingsMapToReturn["authentication_settings"] = []interface{}{authenticationSettingsMapToReturn}
					}
					if v, _ := browserBasedAuthenticationSettingsMap["browser-based-authentication-portal-settings"]; v != nil {
						browserBasedAuthenticationPortalSettingsMap := v.(map[string]interface{})
						browserBasedAuthenticationPortalSettingsMapToReturn := make(map[string]interface{})
						if v, _ := browserBasedAuthenticationPortalSettingsMap["accessibility"]; v != nil {
							accessibilityMap := v.(map[string]interface{})
							accessibilityMapToReturn := make(map[string]interface{})
							if v, _ := accessibilityMap["allow-access-from"]; v != nil {
								accessibilityMapToReturn["allow_access_from"] = v
							}
							if v, _ := accessibilityMap["internal-access-settings"]; v != nil {
								internalAccessSettingsMap := v.(map[string]interface{})
								internalAccessSettingsMapToReturn := make(map[string]interface{})
								if v, _ := internalAccessSettingsMap["dmz"]; v != nil {
									internalAccessSettingsMapToReturn["dmz"] = v
								}
								if v, _ := internalAccessSettingsMap["undefined"]; v != nil {
									internalAccessSettingsMapToReturn["undefined"] = v
								}
								if v, _ := internalAccessSettingsMap["vpn"]; v != nil {
									internalAccessSettingsMapToReturn["vpn"] = v
								}
								accessibilityMapToReturn["internal_access_settings"] = []interface{}{internalAccessSettingsMapToReturn}
							}
							browserBasedAuthenticationPortalSettingsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
						}
						if v, _ := browserBasedAuthenticationPortalSettingsMap["certificate-settings"]; v != nil {
							certificateSettingsMap := v.(map[string]interface{})
							certificateSettingsMapToReturn := make(map[string]interface{})
							if v, _ := certificateSettingsMap["base64-certificate"]; v != nil {
								certificateSettingsMapToReturn["base64_certificate"] = v
							}
							if v, _ := certificateSettingsMap["base64-password"]; v != nil {
								certificateSettingsMapToReturn["base64_password"] = v
							}
							browserBasedAuthenticationPortalSettingsMapToReturn["certificate_settings"] = []interface{}{certificateSettingsMapToReturn}
						}
						if v, _ := browserBasedAuthenticationPortalSettingsMap["portal-web-settings"]; v != nil {
							portalWebSettingsMap := v.(map[string]interface{})
							portalWebSettingsMapToReturn := make(map[string]interface{})
							if v, _ := portalWebSettingsMap["aliases"]; v != nil {
								portalWebSettingsMapToReturn["aliases"] = v
							}
							if v, _ := portalWebSettingsMap["main-url"]; v != nil {
								portalWebSettingsMapToReturn["main_url"] = v
							}
							if v, _ := portalWebSettingsMap["ip-address"]; v != nil {
								portalWebSettingsMapToReturn["ip_address"] = v
							}
							browserBasedAuthenticationPortalSettingsMapToReturn["portal_web_settings"] = []interface{}{portalWebSettingsMapToReturn}
						}
						browserBasedAuthenticationSettingsMapToReturn["browser_based_authentication_portal_settings"] = []interface{}{browserBasedAuthenticationPortalSettingsMapToReturn}
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
						authenticationSettingsMap := v.(map[string]interface{})
						authenticationSettingsMapToReturn := make(map[string]interface{})
						if v, _ := authenticationSettingsMap["authentication-method"]; v != nil {
							authenticationSettingsMapToReturn["authentication_method"] = v
						}
						if v, _ := authenticationSettingsMap["radius"]; v != nil {
							authenticationSettingsMapToReturn["radius"] = v
						}
						if v, _ := authenticationSettingsMap["users-directories"]; v != nil {
							usersDirectoriesMap := v.(map[string]interface{})
							usersDirectoriesMapToReturn := make(map[string]interface{})
							if v, _ := usersDirectoriesMap["external-user-profile"]; v != nil {
								usersDirectoriesMapToReturn["external_user_profile"] = v
							}
							if v, _ := usersDirectoriesMap["internal-users"]; v != nil {
								usersDirectoriesMapToReturn["internal_users"] = v
							}
							if v, _ := usersDirectoriesMap["specific"]; v != nil {
								usersDirectoriesMapToReturn["specific"] = v
							}
							if v, _ := usersDirectoriesMap["users-from-external-directories"]; v != nil {
								usersDirectoriesMapToReturn["users_from_external_directories"] = v
							}
							authenticationSettingsMapToReturn["users_directories"] = []interface{}{usersDirectoriesMapToReturn}
						}
						identityAgentSettingsMapToReturn["authentication_settings"] = []interface{}{authenticationSettingsMapToReturn}
					}
					if v, _ := identityAgentSettingsMap["identity-agent-portal-settings"]; v != nil {
						identityAgentPortalSettingsMap := v.(map[string]interface{})
						identityAgentPortalSettingsMapToReturn := make(map[string]interface{})
						if v, _ := identityAgentPortalSettingsMap["accessibility"]; v != nil {
							accessibilityMap := v.(map[string]interface{})
							accessibilityMapToReturn := make(map[string]interface{})
							if v, _ := accessibilityMap["allow-access-from"]; v != nil {
								accessibilityMapToReturn["allow_access_from"] = v
							}
							if v, _ := accessibilityMap["internal-access-settings"]; v != nil {
								internalAccessSettingsMap := v.(map[string]interface{})
								internalAccessSettingsMapToReturn := make(map[string]interface{})
								if v, _ := internalAccessSettingsMap["dmz"]; v != nil {
									internalAccessSettingsMapToReturn["dmz"] = v
								}
								if v, _ := internalAccessSettingsMap["undefined"]; v != nil {
									internalAccessSettingsMapToReturn["undefined"] = v
								}
								if v, _ := internalAccessSettingsMap["vpn"]; v != nil {
									internalAccessSettingsMapToReturn["vpn"] = v
								}
								accessibilityMapToReturn["internal_access_settings"] = []interface{}{internalAccessSettingsMapToReturn}
							}
							identityAgentPortalSettingsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
						}
						identityAgentSettingsMapToReturn["identity_agent_portal_settings"] = []interface{}{identityAgentPortalSettingsMapToReturn}
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
						authenticationSettingsMap := v.(map[string]interface{})
						authenticationSettingsMapToReturn := make(map[string]interface{})
						if v, _ := authenticationSettingsMap["users-directories"]; v != nil {
							usersDirectoriesMap := v.(map[string]interface{})
							usersDirectoriesMapToReturn := make(map[string]interface{})
							if v, _ := usersDirectoriesMap["external-user-profile"]; v != nil {
								usersDirectoriesMapToReturn["external_user_profile"] = v
							}
							if v, _ := usersDirectoriesMap["internal-users"]; v != nil {
								usersDirectoriesMapToReturn["internal_users"] = v
							}
							if v, _ := usersDirectoriesMap["specific"]; v != nil {
								usersDirectoriesMapToReturn["specific"] = v
							}
							if v, _ := usersDirectoriesMap["users-from-external-directories"]; v != nil {
								usersDirectoriesMapToReturn["users_from_external_directories"] = v
							}
							authenticationSettingsMapToReturn["users_directories"] = []interface{}{usersDirectoriesMapToReturn}
						}
						identityCollectorSettingsMapToReturn["authentication_settings"] = []interface{}{authenticationSettingsMapToReturn}
					}
					if v, _ := identityCollectorSettingsMap["client-access-permissions"]; v != nil {
						clientAccessPermissionsMap := v.(map[string]interface{})
						clientAccessPermissionsMapToReturn := make(map[string]interface{})
						if v, _ := clientAccessPermissionsMap["accessibility"]; v != nil {
							accessibilityMap := v.(map[string]interface{})
							accessibilityMapToReturn := make(map[string]interface{})
							if v, _ := accessibilityMap["allow-access-from"]; v != nil {
								accessibilityMapToReturn["allow_access_from"] = v
							}
							if v, _ := accessibilityMap["internal-access-settings"]; v != nil {
								internalAccessSettingsMap := v.(map[string]interface{})
								internalAccessSettingsMapToReturn := make(map[string]interface{})
								if v, _ := internalAccessSettingsMap["dmz"]; v != nil {
									internalAccessSettingsMapToReturn["dmz"] = v
								}
								if v, _ := internalAccessSettingsMap["undefined"]; v != nil {
									internalAccessSettingsMapToReturn["undefined"] = v
								}
								if v, _ := internalAccessSettingsMap["vpn"]; v != nil {
									internalAccessSettingsMapToReturn["vpn"] = v
								}
								accessibilityMapToReturn["internal_access_settings"] = []interface{}{internalAccessSettingsMapToReturn}
							}
							clientAccessPermissionsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
						}
						identityCollectorSettingsMapToReturn["client_access_permissions"] = []interface{}{clientAccessPermissionsMapToReturn}
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
					if v, _ := identitySharingSettingsMap["cache-mode"]; v != nil {
						cacheModeMap := v.(map[string]interface{})
						cacheModeMapToReturn := make(map[string]interface{})
						if v, _ := cacheModeMap["override-profile"]; v != nil {
							cacheModeMapToReturn["override_profile"] = v
						}
						if v, _ := cacheModeMap["value"]; v != nil {
							cacheModeMapToReturn["value"] = v
						}
						identitySharingSettingsMapToReturn["cache_mode"] = []interface{}{cacheModeMapToReturn}
					}
					if v, _ := identitySharingSettingsMap["cache-mode-duration"]; v != nil {
						cacheModeDurationMap := v.(map[string]interface{})
						cacheModeDurationMapToReturn := make(map[string]interface{})
						if v, _ := cacheModeDurationMap["override-profile"]; v != nil {
							cacheModeDurationMapToReturn["override_profile"] = v
						}
						if v, _ := cacheModeDurationMap["value"]; v != nil {
							cacheModeDurationMapToReturn["value"] = v
						}
						identitySharingSettingsMapToReturn["cache_mode_duration"] = []interface{}{cacheModeDurationMapToReturn}
					}
					if v, _ := identitySharingSettingsMap["receive-from-infinity-identity"]; v != nil {
						identitySharingSettingsMapToReturn["receive_from_infinity_identity"] = v
					}
					if v, _ := identitySharingSettingsMap["scaled-sharing"]; v != nil {
						identitySharingSettingsMapToReturn["scaled_sharing"] = v
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
			if v := identityAwarenessSettingsMap["identity-based-enforcement"]; v != nil {
				identityAwarenessSettingsMapToReturn["identity_based_enforcement"] = v
			}
			if v := identityAwarenessSettingsMap["identity-web-api"]; v != nil {
				identityAwarenessSettingsMapToReturn["identity_web_api"] = v
			}
			if v := identityAwarenessSettingsMap["identity-web-api-settings"]; v != nil {
				identityWebApiSettingsShow := v.(map[string]interface{})
				identityWebApiSettingsState := make(map[string]interface{})
				if v := identityWebApiSettingsShow["authentication-settings"]; v != nil {
					authenticationSettingsShow := v.(map[string]interface{})
					authenticationSettingsState := make(map[string]interface{})
					if v := authenticationSettingsShow["users-directories"]; v != nil {
						usersDirectoriesShow := v.(map[string]interface{})
						usersDirectoriesState := make(map[string]interface{})
						if v := usersDirectoriesShow["external-user-profile"]; v != nil {
							usersDirectoriesState["external_user_profile"] = v
						}
						if v := usersDirectoriesShow["internal-users"]; v != nil {
							usersDirectoriesState["internal_users"] = v
						}
						if v := usersDirectoriesShow["specific"]; v != nil {
							usersDirectoriesState["specific"] = v
						}
						if v := usersDirectoriesShow["users-from-external-directories"]; v != nil {
							usersDirectoriesState["users_from_external_directories"] = v
						}
						if v := usersDirectoriesShow["specific"]; v != nil {
							usersDirectoriesState["specific"] = v
						}
						authenticationSettingsState["users_directories"] = []interface{}{usersDirectoriesState}
					}
					identityWebApiSettingsState["authentication_settings"] = []interface{}{authenticationSettingsState}
				}
				if v := identityWebApiSettingsShow["authorized-clients"]; v != nil {
					authorizedClientsShow := v.(map[string]interface{})
					authorizedClientsState := make(map[string]interface{})
					if v := authorizedClientsShow["client"]; v != nil {
						authorizedClientsState["client"] = v
					}
					identityWebApiSettingsState["authorized_clients"] = []interface{}{authorizedClientsState}
				}
				if v := identityWebApiSettingsShow["client-access-permissions"]; v != nil {
					clientAccessPermissionsShow := v.(map[string]interface{})
					clientAccessPermissionsState := make(map[string]interface{})
					if v := clientAccessPermissionsShow["accessibility"]; v != nil {
						accessibilityShow := v.(map[string]interface{})
						accessibilityState := make(map[string]interface{})
						if v := accessibilityShow["allow-access-from"]; v != nil {
							accessibilityState["allow_access_from"] = v
						}
						if v := accessibilityShow["internal-access-settings"]; v != nil {
							internalAccessSettingsShow := v.(map[string]interface{})
							internalAccessSettingsState := make(map[string]interface{})
							if v := internalAccessSettingsShow["dmz"]; v != nil {
								internalAccessSettingsState["dmz"] = v
							}
							if v := internalAccessSettingsShow["undefined"]; v != nil {
								internalAccessSettingsState["undefined"] = v
							}
							if v := internalAccessSettingsShow["vpn"]; v != nil {
								internalAccessSettingsState["vpn"] = v
							}
							accessibilityState["internal_access_settings"] = []interface{}{internalAccessSettingsState}
						}
						clientAccessPermissionsState["accessibility"] = []interface{}{accessibilityState}
					}
					if v := clientAccessPermissionsShow["certificate-settings"]; v != nil {
						certificateSettingsShow := v.(map[string]interface{})
						certificateSettingsState := make(map[string]interface{})
						if v := certificateSettingsShow["certificate"]; v != nil {
							certificateSettingsState["certificate"] = v
						}
						if v := certificateSettingsShow["certificate-dn"]; v != nil {
							certificateSettingsState["certificate_dn"] = v
						}
						if v := certificateSettingsShow["certificate-valid-from"]; v != nil {
							certificateSettingsState["certificate_valid_from"] = v
						}
						if v := certificateSettingsShow["certificate-valid-to"]; v != nil {
							certificateSettingsState["certificate_valid_to"] = v
						}
						clientAccessPermissionsState["certificate_settings"] = []interface{}{certificateSettingsState}
					}
					if v := clientAccessPermissionsShow["portal-web-settings"]; v != nil {
						portalWebSettingsShow := v.(map[string]interface{})
						portalWebSettingsState := make(map[string]interface{})
						if v := portalWebSettingsShow["aliases"]; v != nil {
							portalWebSettingsState["aliases"] = v
						}
						if v := portalWebSettingsShow["ip-address"]; v != nil {
							portalWebSettingsState["ip_address"] = v
						}
						if v := portalWebSettingsShow["main-url"]; v != nil {
							portalWebSettingsState["main_url"] = v
						}
						clientAccessPermissionsState["portal_web_settings"] = []interface{}{portalWebSettingsState}
					}
					identityWebApiSettingsState["client_access_permissions"] = []interface{}{clientAccessPermissionsState}
				}
				identityAwarenessSettingsMapToReturn["identity_web_api_settings"] = []interface{}{identityWebApiSettingsState}
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

		if v := natSettingsMap["auto-rule"]; v != nil {
			natSettingsMapToReturn["auto_rule"] = v
		}
		if v := natSettingsMap["ipv4-address"]; v != nil {
			natSettingsMapToReturn["ipv4_address"] = v
		}
		if v := natSettingsMap["ipv6-address"]; v != nil {
			natSettingsMapToReturn["ipv6_address"] = v
		}
		if v := natSettingsMap["hide-behind"]; v != nil {
			natSettingsMapToReturn["hide_behind"] = v
		}
		if v := natSettingsMap["install-on"]; v != nil {
			natSettingsMapToReturn["install_on"] = v
		}
		if v := natSettingsMap["method"]; v != nil {
			natSettingsMapToReturn["method"] = v
		}
		if v := natSettingsMap["apply-control-connections"]; v != nil {
			natSettingsMapToReturn["apply_control_connections"] = v
		}
		_ = d.Set("nat_settings", []interface{}{natSettingsMapToReturn})

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
					if v, _ := portalWebSettingsMap["ip-address"]; v != nil {
						portalWebSettingsMapToReturn["ip_address"] = v
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

		if v := proxySettingsMap["use-custom-proxy"]; v != nil {
			proxySettingsMapToReturn["use_custom_proxy"] = v
		}
		if v := proxySettingsMap["proxy-server"]; v != nil {
			proxySettingsMapToReturn["proxy_server"] = v
		}
		if v := proxySettingsMap["port"]; v != nil {
			proxySettingsMapToReturn["port"] = v
		}
		_ = d.Set("proxy_settings", []interface{}{proxySettingsMapToReturn})

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
					if v, _ := portalWebSettingsMap["ip-address"]; v != nil {
						portalWebSettingsMapToReturn["ip_address"] = v
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
					interfaceState["anti_spoofing_settings"] = []interface{}{antiSpoofingSettingsState}
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
					interfaceState["security_zone_settings"] = []interface{}{securityZoneSettingsState}
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
					interfaceState["topology_settings"] = []interface{}{topologySettingsState}
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
		_ = d.Set("firewall_settings", []interface{}{firewallSettingsState})
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
			if v := authenticationJson["single-authentication-client"]; v != nil {
				singleAuthenticationClientJson := v.(map[string]interface{})
				singleAuthenticationClientState := make(map[string]interface{})
				if v := singleAuthenticationClientJson["enabled"]; v != nil {
					singleAuthenticationClientState["enabled"] = v
				}
				if v := singleAuthenticationClientJson["allow-multiple-authentication-clients"]; v != nil {
					singleAuthenticationClientState["allow_multiple_authentication_clients"] = v
				}
				if v := singleAuthenticationClientJson["display-name"]; v != nil {
					singleAuthenticationClientState["display_name"] = v
				}
				if v := singleAuthenticationClientJson["method"]; v != nil {
					singleAuthenticationClientState["method"] = v
				}
				if v := singleAuthenticationClientJson["secur-id"]; v != nil {
					securIdJson := v.(map[string]interface{})
					securIdState := make(map[string]interface{})
					if v := securIdJson["server"]; v != nil {
						securIdState["server"] = v.(map[string]interface{})["name"]
					}
					if v := securIdJson["token-card-type"]; v != nil {
						securIdState["token_card_type"] = v
					}
					singleAuthenticationClientState["secur_id"] = []interface{}{securIdState}
				}
				if v := singleAuthenticationClientJson["radius"]; v != nil {
					radiusJson := v.(map[string]interface{})
					radiusState := make(map[string]interface{})
					if v := radiusJson["server"]; v != nil {
						radiusState["server"] = v.(map[string]interface{})["name"]
					}
					if v := radiusJson["ask-user-password"]; v != nil {
						radiusState["ask_user_password"] = v
					}
					singleAuthenticationClientState["radius"] = []interface{}{radiusState}
				}
				if v := singleAuthenticationClientJson["personal-certificate"]; v != nil {
					personalCertificateJson := v.(map[string]interface{})
					personalCertificateState := make(map[string]interface{})
					if v := personalCertificateJson["fetch-username-from"]; v != nil {
						personalCertificateState["fetch_username_from"] = v
					}
					if v := personalCertificateJson["storage-type"]; v != nil {
						personalCertificateState["storage_type"] = v
					}
					if v := personalCertificateJson["source"]; v != nil {
						personalCertificateState["source"] = v
					}
					if v := personalCertificateJson["dn-part"]; v != nil {
						personalCertificateState["dn_part"] = v
					}
					if v := personalCertificateJson["dn-concurrence"]; v != nil {
						personalCertificateState["dn_concurrence"] = v
					}
					singleAuthenticationClientState["personal_certificate"] = []interface{}{personalCertificateState}
				}
				if v := singleAuthenticationClientJson["client-display-settings"]; v != nil {
					clientDisplaySettingsJson := v.(map[string]interface{})
					clientDisplaySettingsState := make(map[string]interface{})
					if v := clientDisplaySettingsJson["headline"]; v != nil {
						clientDisplaySettingsState["headline"] = v
					}
					if v := clientDisplaySettingsJson["username-label"]; v != nil {
						clientDisplaySettingsState["username_label"] = v
					}
					if v := clientDisplaySettingsJson["password-label"]; v != nil {
						clientDisplaySettingsState["password_label"] = v
					}
					singleAuthenticationClientState["client_display_settings"] = []interface{}{clientDisplaySettingsState}
				}
				authenticationState["single_authentication_client"] = []interface{}{singleAuthenticationClientState}
			}
			if v := authenticationJson["override-global-dynamic-id-settings"]; v != nil {
				authenticationState["override_global_dynamic_id_settings"] = v
			}
			if v := authenticationJson["dynamic-id-settings"]; v != nil {
				dynamicIdSettingsJson := v.(map[string]interface{})
				dynamicIdSettingsState := make(map[string]interface{})
				if v := dynamicIdSettingsJson["sms-provider-and-email-settings"]; v != nil {
					dynamicIdSettingsState["sms_provider_and_email_settings"] = v
				}
				if v := dynamicIdSettingsJson["sms-provider-credentials"]; v != nil {
					smsProviderCredentialsJson := v.(map[string]interface{})
					smsProviderCredentialsState := make(map[string]interface{})
					if v := smsProviderCredentialsJson["username"]; v != nil {
						smsProviderCredentialsState["username"] = v
					}
					if v := smsProviderCredentialsJson["password"]; v != nil {
						smsProviderCredentialsState["password"] = v
					}
					if v := smsProviderCredentialsJson["api-id"]; v != nil {
						smsProviderCredentialsState["api_id"] = v
					}
					dynamicIdSettingsState["sms_provider_credentials"] = []interface{}{smsProviderCredentialsState}
				}
				if v := dynamicIdSettingsJson["advanced-settings"]; v != nil {
					advancedSettingsJson := v.(map[string]interface{})
					advancedSettingsState := make(map[string]interface{})
					if v := advancedSettingsJson["dynamic-id-message"]; v != nil {
						advancedSettingsState["dynamic_id_message"] = v
					}
					if v := advancedSettingsJson["otp-settings"]; v != nil {
						otpSettingsJson := v.(map[string]interface{})
						otpSettingsState := make(map[string]interface{})
						if v := otpSettingsJson["length"]; v != nil {
							otpSettingsState["length"] = v
						}
						if v := otpSettingsJson["expiration"]; v != nil {
							otpSettingsState["expiration"] = v
						}
						if v := otpSettingsJson["max-attempts"]; v != nil {
							otpSettingsState["max_attempts"] = v
						}
						advancedSettingsState["otp_settings"] = []interface{}{otpSettingsState}
					}
					if v := advancedSettingsJson["enable-display-user-details"]; v != nil {
						advancedSettingsState["enable_display_user_details"] = v
					}
					if v := advancedSettingsJson["country-code"]; v != nil {
						advancedSettingsState["country_code"] = v
					}
					if v := advancedSettingsJson["user-details-retrieval"]; v != nil {
						advancedSettingsState["user_details_retrieval"] = v
					}
					dynamicIdSettingsState["advanced_settings"] = []interface{}{advancedSettingsState}
				}
				authenticationState["dynamic_id_settings"] = []interface{}{dynamicIdSettingsState}
			}
			if v := authenticationJson["send-machine-certificate"]; v != nil {
				authenticationState["send_machine_certificate"] = v
			}
			vpnSettingsState["authentication"] = []interface{}{authenticationState}
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
			if v := linkSelectionJson["route-selection-method"]; v != nil {
				linkSelectionState["route_selection_method"] = v
			}
			if v := linkSelectionJson["responding-traffic"]; v != nil {
				linkSelectionState["responding_traffic"] = v
			}
			if v := linkSelectionJson["source-ip-selection"]; v != nil {
				linkSelectionState["source_ip_selection"] = v
			}
			if v := linkSelectionJson["selected-ip"]; v != nil {
				linkSelectionState["selected_ip"] = v
			}
			if v := linkSelectionJson["outgoing-link-tracking"]; v != nil {
				linkSelectionState["outgoing_link_tracking"] = v
			}
			if v := linkSelectionJson["probing-settings"]; v != nil {
				probingSettingsJson := v.(map[string]interface{})
				probingSettingsState := make(map[string]interface{})
				if v := probingSettingsJson["probed-interfaces"]; v != nil {
					probingSettingsState["probed_interfaces"] = v
				}
				if v := probingSettingsJson["probed-interface-list"]; v != nil {
					probingSettingsState["probed_interface_list"] = v
				}
				if v := probingSettingsJson["use-primary-address"]; v != nil {
					probingSettingsState["use_primary_address"] = v
				}
				if v := probingSettingsJson["primary-address"]; v != nil {
					probingSettingsState["primary_address"] = v
				}
				if v := probingSettingsJson["probing-method"]; v != nil {
					probingSettingsState["probing_method"] = v
				}
				linkSelectionState["probing_settings"] = []interface{}{probingSettingsState}
			}
			vpnSettingsState["link_selection"] = []interface{}{linkSelectionState}
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
			vpnSettingsState["remote_access"] = []interface{}{remoteAccessState}
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
					allocateIpAddressFromState["optional_parameters"] = []interface{}{optionalParametersState}
				}
				officeModeState["allocate_ip_address_from"] = []interface{}{allocateIpAddressFromState}
			}
			vpnSettingsState["office_mode"] = []interface{}{officeModeState}
		}

		if v := vpnSettingsJson["advanced"]; v != nil {
			advancedJson := v.(map[string]interface{})
			advancedState := make(map[string]interface{})
			if v := advancedJson["tunnel-sharing-mode"]; v != nil {
				advancedState["tunnel_sharing_mode"] = v
			}
			if v := advancedJson["shutdown-on-gateway-restart"]; v != nil {
				advancedState["shutdown_on_gateway_restart"] = v
			}
			if v := advancedJson["enable-wire-mode"]; v != nil {
				advancedState["enable_wire_mode"] = v
			}
			if v := advancedJson["wire-mode-interfaces"]; v != nil {
				interfacesJson := v.([]interface{})
				var interfacesIds = make([]string, 0)
				if len(interfacesJson) > 0 {
					for _, iface := range interfacesJson {
						interfacesIds = append(interfacesIds, iface.(map[string]interface{})["name"].(string))
					}
				}
				advancedState["wire_mode_interfaces"] = interfacesIds
			}
			if v := advancedJson["enable-wire-mode-log-traffic"]; v != nil {
				advancedState["enable_wire_mode_log_traffic"] = v
			}
			if v := advancedJson["enable-nat-traversal"]; v != nil {
				advancedState["enable_nat_traversal"] = v
			}
			vpnSettingsState["advanced"] = []interface{}{advancedState}
		}

		if v := vpnSettingsJson["exported-routes"]; v != nil {
			exportedRoutesJson := v.(map[string]interface{})
			exportedRoutesState := make(map[string]interface{})
			if v := exportedRoutesJson["internal-interfaces"]; v != nil {
				exportedRoutesState["internal_interfaces"] = v
			}
			if v := exportedRoutesJson["static-routes"]; v != nil {
				exportedRoutesState["static_routes"] = v
			}
			if v := exportedRoutesJson["custom-routes"]; v != nil {
				exportedRoutesState["custom_routes"] = v
			}
			if v := exportedRoutesJson["custom-routes-object"]; v != nil {
				exportedRoutesState["custom_routes_object"] = v.(map[string]interface{})["name"]
			}
			vpnSettingsState["exported_routes"] = []interface{}{exportedRoutesState}
		}

		if v := vpnSettingsJson["vpn-clients"]; v != nil {
			vpnClientsJson := v.(map[string]interface{})
			vpnClientsState := make(map[string]interface{})
			if v := vpnClientsJson["enable-endpoint-security-vpn"]; v != nil {
				vpnClientsState["enable_endpoint_security_vpn"] = v
			}
			if v := vpnClientsJson["enable-cp-mobile-for-windows"]; v != nil {
				vpnClientsState["enable_cp_mobile_for_windows"] = v
			}
			if v := vpnClientsJson["enable-secu-remote"]; v != nil {
				vpnClientsState["enable_secu_remote"] = v
			}
			if v := vpnClientsJson["enable-capsule-vpn-connect"]; v != nil {
				vpnClientsState["enable_capsule_vpn_connect"] = v
			}
			if v := vpnClientsJson["enable-ssl-network-extender"]; v != nil {
				vpnClientsState["enable_ssl_network_extender"] = v
			}
			if v := vpnClientsJson["gateway-authentication-certificate"]; v != nil {
				vpnClientsState["gateway_authentication_certificate"] = v
			}
			vpnSettingsState["vpn_clients"] = []interface{}{vpnClientsState}
		}

		if v := vpnSettingsJson["enable-clientless-vpn"]; v != nil {
			vpnSettingsState["enable_clientless_vpn"] = v
		}

		if v := vpnSettingsJson["clientless-vpn-settings"]; v != nil {
			clientlessVpnSettingsJson := v.(map[string]interface{})
			clientlessVpnSettingsState := make(map[string]interface{})
			if v := clientlessVpnSettingsJson["certificate-gateway-authentication"]; v != nil {
				clientlessVpnSettingsState["certificate_gateway_authentication"] = v
			}
			if v := clientlessVpnSettingsJson["client-authentication"]; v != nil {
				clientlessVpnSettingsState["client_authentication"] = v
			}
			if v := clientlessVpnSettingsJson["concurrent-servers-or-processes"]; v != nil {
				clientlessVpnSettingsState["concurrent_servers_or_processes"] = v
			}
			if v := clientlessVpnSettingsJson["accept-only-3des"]; v != nil {
				clientlessVpnSettingsState["accept_only_3des"] = v
			}
			vpnSettingsState["clientless_vpn_settings"] = []interface{}{clientlessVpnSettingsState}
		}

		if v := vpnSettingsJson["saml-portal-settings"]; v != nil {
			samlPortalSettingsJson := v.(map[string]interface{})
			samlPortalSettingsState := make(map[string]interface{})
			if v := samlPortalSettingsJson["portal-web-settings"]; v != nil {
				portalWebSettingsJson := v.(map[string]interface{})
				portalWebSettingsState := make(map[string]interface{})
				if v := portalWebSettingsJson["aliases"]; v != nil {
					portalWebSettingsState["aliases"] = v
				}
				if v := portalWebSettingsJson["ip-address"]; v != nil {
					portalWebSettingsState["ip_address"] = v
				}
				if v := portalWebSettingsJson["main-url"]; v != nil {
					portalWebSettingsState["main_url"] = v
				}
				samlPortalSettingsState["portal_web_settings"] = []interface{}{portalWebSettingsState}
			}
			if v := samlPortalSettingsJson["accessibility"]; v != nil {
				accessibilityJson := v.(map[string]interface{})
				accessibilityState := make(map[string]interface{})
				if v := accessibilityJson["allow-access-from"]; v != nil {
					accessibilityState["allow_access_from"] = v
				}
				if v := accessibilityJson["internal-access-settings"]; v != nil {
					internalAccessSettingsJson := v.(map[string]interface{})
					internalAccessSettingsState := make(map[string]interface{})
					if v := internalAccessSettingsJson["undefined"]; v != nil {
						internalAccessSettingsState["undefined"] = v
					}
					if v := internalAccessSettingsJson["dmz"]; v != nil {
						internalAccessSettingsState["dmz"] = v
					}
					if v := internalAccessSettingsJson["vpn"]; v != nil {
						internalAccessSettingsState["vpn"] = v
					}
					accessibilityState["internal_access_settings"] = []interface{}{internalAccessSettingsState}
				}
				samlPortalSettingsState["accessibility"] = []interface{}{accessibilityState}
			}
			vpnSettingsState["saml_portal_settings"] = []interface{}{samlPortalSettingsState}
		}

		if v := vpnSettingsJson["interfaces"]; v != nil {
			interfacesList := v.([]interface{})
			var interfacesListState []map[string]interface{}
			for i := range interfacesList {
				interfacesShow := interfacesList[i].(map[string]interface{})
				interfacesState := make(map[string]interface{})
				if v := interfacesShow["interface-name"]; v != nil {
					interfacesState["interface_name"] = v
				}
				if v := interfacesShow["ip-version"]; v != nil {
					interfacesState["ip_version"] = v
				}
				if v := interfacesShow["next-hop-ip"]; v != nil {
					interfacesState["next_hop_ip"] = v
				}
				if v := interfacesShow["priority"]; v != nil {
					interfacesState["priority"] = v
				}
				if v := interfacesShow["redundancy-mode"]; v != nil {
					interfacesState["redundancy_mode"] = v
				}
				if v := interfacesShow["static-nat-ip"]; v != nil {
					interfacesState["static_nat_ip"] = v
				}
				interfacesListState = append(interfacesListState, interfacesState)
			}
			vpnSettingsState["interfaces"] = interfacesListState
		}
		_ = d.Set("vpn_settings", []interface{}{vpnSettingsState})
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

	if v := cluster["anti-spam-and-email-security"]; v != nil {
		_ = d.Set("anti_spam_and_email_security", v)
	}

	if v := cluster["auto-topology-custom-recalculation-time"]; v != nil {
		_ = d.Set("auto_topology_custom_recalculation_time", v)
	}

	if v := cluster["auto-topology-use-custom-recalculation-time"]; v != nil {
		_ = d.Set("auto_topology_use_custom_recalculation_time", v)
	}

	if v := cluster["data-loss-prevention"]; v != nil {
		_ = d.Set("data_loss_prevention", v)
	}

	if v := cluster["mobile-access"]; v != nil {
		_ = d.Set("mobile_access", v)
	}

	if v := cluster["monitoring"]; v != nil {
		_ = d.Set("monitoring", v)
	}

	if v := cluster["policy-server"]; v != nil {
		_ = d.Set("policy_server", v)
	}

	if v := cluster["rtm-counters-report"]; v != nil {
		_ = d.Set("rtm_counters_report", v)
	}

	if v := cluster["rtm-traffic-report"]; v != nil {
		_ = d.Set("rtm_traffic_report", v)
	}

	if v := cluster["rtm-traffic-report-per-connection"]; v != nil {
		_ = d.Set("rtm_traffic_report_per_connection", v)
	}

	if v := cluster["threat-extraction"]; v != nil {
		_ = d.Set("threat_extraction", v)
	}

	if v := cluster["threat-prevention-mode"]; v != nil {
		_ = d.Set("threat_prevention_mode", v)
	}

	if v := cluster["workforce-ai"]; v != nil {
		_ = d.Set("workforce_ai", v)
	}

	if v := cluster["application-control-and-url-filtering-settings"]; v != nil {
		applicationControlAndUrlFilteringSettingsJson := v.(map[string]interface{})
		applicationControlAndUrlFilteringSettingsState := make(map[string]interface{})
		if v := applicationControlAndUrlFilteringSettingsJson["global-settings-mode"]; v != nil {
			applicationControlAndUrlFilteringSettingsState["global_settings_mode"] = v
		}
		if v := applicationControlAndUrlFilteringSettingsJson["override-global-settings"]; v != nil {
			overrideGlobalSettingsJson := v.(map[string]interface{})
			overrideGlobalSettingsState := make(map[string]interface{})
			if v := overrideGlobalSettingsJson["fail-mode"]; v != nil {
				overrideGlobalSettingsState["fail_mode"] = v
			}
			if v := overrideGlobalSettingsJson["website-categorization"]; v != nil {
				websiteCategorizationJson := v.(map[string]interface{})
				websiteCategorizationState := make(map[string]interface{})
				if v := websiteCategorizationJson["custom-mode"]; v != nil {
					customModeJson := v.(map[string]interface{})
					customModeState := make(map[string]interface{})
					if v := customModeJson["social-networking-widgets"]; v != nil {
						customModeState["social_networking_widgets"] = v
					}
					if v := customModeJson["url-filtering"]; v != nil {
						customModeState["url_filtering"] = v
					}
					websiteCategorizationState["custom_mode"] = []interface{}{customModeState}
				}
				if v := websiteCategorizationJson["mode"]; v != nil {
					websiteCategorizationState["mode"] = v
				}
				overrideGlobalSettingsState["website_categorization"] = []interface{}{websiteCategorizationState}
			}
			applicationControlAndUrlFilteringSettingsState["override_global_settings"] = []interface{}{overrideGlobalSettingsState}
		}
		_ = d.Set("application_control_and_url_filtering_settings", []interface{}{applicationControlAndUrlFilteringSettingsState})
	}

	if v := cluster["cluster-settings"]; v != nil {
		clusterSettingsJson := v.(map[string]interface{})
		clusterSettingsState := make(map[string]interface{})
		if v := clusterSettingsJson["member-recovery-mode"]; v != nil {
			clusterSettingsState["member_recovery_mode"] = v
		}
		if v := clusterSettingsJson["state-synchronization"]; v != nil {
			stateSynchronizationJson := v.(map[string]interface{})
			stateSynchronizationState := make(map[string]interface{})
			if v := stateSynchronizationJson["delayed"]; v != nil {
				stateSynchronizationState["delayed"] = v
			}
			if v := stateSynchronizationJson["delayed-seconds"]; v != nil {
				stateSynchronizationState["delayed_seconds"] = v
			}
			if v := stateSynchronizationJson["enabled"]; v != nil {
				stateSynchronizationState["enabled"] = v
			}
			clusterSettingsState["state_synchronization"] = []interface{}{stateSynchronizationState}
		}
		if v := clusterSettingsJson["track-changes-of-cluster-members"]; v != nil {
			clusterSettingsState["track_changes_of_cluster_members"] = v
		}
		if v := clusterSettingsJson["use-virtual-mac"]; v != nil {
			clusterSettingsState["use_virtual_mac"] = v
		}
		_ = d.Set("cluster_settings", []interface{}{clusterSettingsState})
	}

	if v := cluster["communication-with-servers-behind-nat"]; v != nil {
		communicationWithServersBehindNatJson := v.(map[string]interface{})
		communicationWithServersBehindNatState := make(map[string]interface{})
		if v := communicationWithServersBehindNatJson["override-profile"]; v != nil {
			communicationWithServersBehindNatState["override_profile"] = v
		}
		if v := communicationWithServersBehindNatJson["value"]; v != nil {
			communicationWithServersBehindNatState["value"] = v
		}
		_ = d.Set("communication_with_servers_behind_nat", []interface{}{communicationWithServersBehindNatState})
	}

	if v := cluster["zero-phishing-settings"]; v != nil {
		zeroPhishingSettingsJson := v.(map[string]interface{})
		zeroPhishingSettingsState := make(map[string]interface{})
		if v := zeroPhishingSettingsJson["gateway-fqdn-mode"]; v != nil {
			zeroPhishingSettingsState["gateway_fqdn_mode"] = v
		}
		if v := zeroPhishingSettingsJson["manual-fqdn"]; v != nil {
			zeroPhishingSettingsState["manual_fqdn"] = v
		}
		_ = d.Set("zero_phishing_settings", []interface{}{zeroPhishingSettingsState})
	}

	if v := cluster["dns-server"]; v != nil {
		_ = d.Set("dns_server", v)
	}

	if v := cluster["logs-settings"]; v != nil {
		logsSettingsShow := v.(map[string]interface{})
		logsSettingsState := make(map[string]interface{})
		if v := logsSettingsShow["alert-when-free-disk-space-below"]; v != nil {
			logsSettingsState["alert_when_free_disk_space_below"] = v
		}
		if v := logsSettingsShow["alert-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsState["alert_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsShow["alert-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsState["alert_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsShow["alert-when-free-disk-space-below-type"]; v != nil {
			logsSettingsState["alert_when_free_disk_space_below_type"] = v
		}
		if v := logsSettingsShow["before-delete-keep-logs-from-the-last-days"]; v != nil {
			logsSettingsState["before_delete_keep_logs_from_the_last_days"] = v
		}
		if v := logsSettingsShow["before-delete-keep-logs-from-the-last-days-threshold"]; v != nil {
			logsSettingsState["before_delete_keep_logs_from_the_last_days_threshold"] = v
		}
		if v := logsSettingsShow["before-delete-run-script"]; v != nil {
			logsSettingsState["before_delete_run_script"] = v
		}
		if v := logsSettingsShow["before-delete-run-script-command"]; v != nil {
			logsSettingsState["before_delete_run_script_command"] = v
		}
		if v := logsSettingsShow["delete-index-files-older-than-days"]; v != nil {
			logsSettingsState["delete_index_files_older_than_days"] = v
		}
		if v := logsSettingsShow["delete-index-files-older-than-days-threshold"]; v != nil {
			logsSettingsState["delete_index_files_older_than_days_threshold"] = v
		}
		if v := logsSettingsShow["delete-index-files-when-index-size-above"]; v != nil {
			logsSettingsState["delete_index_files_when_index_size_above"] = v
		}
		if v := logsSettingsShow["delete-index-files-when-index-size-above-metrics"]; v != nil {
			logsSettingsState["delete_index_files_when_index_size_above_metrics"] = v
		}
		if v := logsSettingsShow["delete-index-files-when-index-size-above-threshold"]; v != nil {
			logsSettingsState["delete_index_files_when_index_size_above_threshold"] = v
		}
		if v := logsSettingsShow["delete-when-free-disk-space-below"]; v != nil {
			logsSettingsState["delete_when_free_disk_space_below"] = v
		}
		if v := logsSettingsShow["delete-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsState["delete_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsShow["delete-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsState["delete_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsShow["detect-new-citrix-ica-application-names"]; v != nil {
			logsSettingsState["detect_new_citrix_ica_application_names"] = v
		}
		if v := logsSettingsShow["distribute-logs-between-all-active-servers"]; v != nil {
			logsSettingsState["distribute_logs_between_all_active_servers"] = v
		}
		if v := logsSettingsShow["forward-logs-to-log-server"]; v != nil {
			logsSettingsState["forward_logs_to_log_server"] = v
		}
		if v := logsSettingsShow["forward-logs-to-log-server-name"]; v != nil {
			logsSettingsState["forward_logs_to_log_server_name"] = v
		}
		if v := logsSettingsShow["forward-logs-to-log-server-schedule-name"]; v != nil {
			logsSettingsState["forward_logs_to_log_server_schedule_name"] = v
		}
		if v := logsSettingsShow["include-tcp-state-information"]; v != nil {
			logsSettingsState["include_tcp_state_information"] = v
		}
		if v := logsSettingsShow["perform-log-rotate-before-log-forwarding"]; v != nil {
			logsSettingsState["perform_log_rotate_before_log_forwarding"] = v
		}
		if v := logsSettingsShow["reject-connections-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsState["reject_connections_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsShow["reserve-for-packet-capture-metrics"]; v != nil {
			logsSettingsState["reserve_for_packet_capture_metrics"] = v
		}
		if v := logsSettingsShow["reserve-for-packet-capture-threshold"]; v != nil {
			logsSettingsState["reserve_for_packet_capture_threshold"] = v
		}
		if v := logsSettingsShow["rotate-log-by-file-size"]; v != nil {
			logsSettingsState["rotate_log_by_file_size"] = v
		}
		if v := logsSettingsShow["rotate-log-file-size-threshold"]; v != nil {
			logsSettingsState["rotate_log_file_size_threshold"] = v
		}
		if v := logsSettingsShow["rotate-log-on-schedule"]; v != nil {
			logsSettingsState["rotate_log_on_schedule"] = v
		}
		if v := logsSettingsShow["rotate-log-schedule-name"]; v != nil {
			logsSettingsState["rotate_log_schedule_name"] = v
		}
		if v := logsSettingsShow["stop-logging-when-free-disk-space-below"]; v != nil {
			logsSettingsState["stop_logging_when_free_disk_space_below"] = v
		}
		if v := logsSettingsShow["stop-logging-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsState["stop_logging_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsShow["stop-logging-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsState["stop_logging_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsShow["turn-on-qos-logging"]; v != nil {
			logsSettingsState["turn_on_qos_logging"] = v
		}
		if v := logsSettingsShow["update-account-log-every"]; v != nil {
			logsSettingsState["update_account_log_every"] = v
		}
		_ = d.Set("logs_settings", []interface{}{logsSettingsState})
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
				if _, ok := d.GetOk("https_inspection.0.bypass_on_client_failure"); ok {

					bypassOnClientFailurePayload := make(map[string]interface{})

					if v, ok := d.GetOk("https_inspection.0.bypass_on_client_failure.0.override_profile"); ok {
						bypassOnClientFailurePayload["override-profile"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("https_inspection.0.bypass_on_client_failure.0.value"); ok {
						bypassOnClientFailurePayload["value"] = strconv.FormatBool(v.(bool))
					}
					httpsInspectionPayload["bypass-on-client-failure"] = bypassOnClientFailurePayload
				}
				if _, ok := d.GetOk("https_inspection.0.bypass_under_load"); ok {

					bypassUnderLoadPayload := make(map[string]interface{})

					if v, ok := d.GetOk("https_inspection.0.bypass_under_load.0.value"); ok {
						bypassUnderLoadPayload["value"] = strconv.FormatBool(v.(bool))
					}
					httpsInspectionPayload["bypass-under-load"] = bypassUnderLoadPayload
				}
				if _, ok := d.GetOk("https_inspection.0.outbound_certificate"); ok {

					outboundCertificatePayload := make(map[string]interface{})

					if v, ok := d.GetOk("https_inspection.0.outbound_certificate.0.override_profile"); ok {
						outboundCertificatePayload["override-profile"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("https_inspection.0.outbound_certificate.0.value"); ok {
						outboundCertificatePayload["value"] = v.(string)
					}
					httpsInspectionPayload["outbound-certificate"] = outboundCertificatePayload
				}
				if v, ok := d.GetOk("https_inspection.0.deployment_mode"); ok {
					httpsInspectionPayload["deployment-mode"] = v.(string)
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

					if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings"); ok {
						authenticationSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.authentication_method"); ok {
							authenticationSettingsPayload["authentication-method"] = v.(string)
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.identity_provider"); ok {
							authenticationSettingsPayload["identity-provider"] = v.(*schema.Set).List()
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.radius"); ok {
							authenticationSettingsPayload["radius"] = v.(string)
						}
						if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories"); ok {

							usersDirectoriesPayload := make(map[string]interface{})

							if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories.0.external_user_profile"); ok {
								usersDirectoriesPayload["external-user-profile"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories.0.internal_users"); ok {
								usersDirectoriesPayload["internal-users"] = v.(bool)
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories.0.specific"); ok {
								usersDirectoriesPayload["specific"] = v.(*schema.Set).List()
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.authentication_settings.0.users_directories.0.users_from_external_directories"); ok {
								usersDirectoriesPayload["users-from-external-directories"] = v.(string)
							}
							authenticationSettingsPayload["users-directories"] = usersDirectoriesPayload
						}
						browserBasedAuthenticationSettingsPayload["authentication-settings"] = authenticationSettingsPayload
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings"); ok {
						browserBasedAuthenticationPortalSettingsPayload := make(map[string]interface{})

						if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility"); ok {

							accessibilityPayload := make(map[string]interface{})

							if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.allow_access_from"); ok {
								accessibilityPayload["allow-access-from"] = v.(string)
							}
							if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.internal_access_settings"); ok {

								internalAccessSettingsPayload := make(map[string]interface{})

								if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.internal_access_settings.0.dmz"); ok {
									internalAccessSettingsPayload["dmz"] = v.(bool)
								}
								if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.internal_access_settings.0.undefined"); ok {
									internalAccessSettingsPayload["undefined"] = v.(bool)
								}
								if v, ok := d.GetOkExists("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.accessibility.0.internal_access_settings.0.vpn"); ok {
									internalAccessSettingsPayload["vpn"] = v.(bool)
								}
								accessibilityPayload["internal-access-settings"] = internalAccessSettingsPayload
							}
							browserBasedAuthenticationPortalSettingsPayload["accessibility"] = accessibilityPayload
						}
						if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.certificate_settings"); ok {

							certificateSettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.certificate_settings.0.base64_certificate"); ok {
								certificateSettingsPayload["base64-certificate"] = v.(string)
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.certificate_settings.0.base64_password"); ok {
								certificateSettingsPayload["base64-password"] = v.(string)
							}
							browserBasedAuthenticationPortalSettingsPayload["certificate-settings"] = certificateSettingsPayload
						}
						if _, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.portal_web_settings"); ok {

							portalWebSettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.portal_web_settings.0.aliases"); ok {
								portalWebSettingsPayload["aliases"] = v.(*schema.Set).List()
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.portal_web_settings.0.ip_address"); ok {
								portalWebSettingsPayload["ip-address"] = v.(string)
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.browser_based_authentication_settings.0.browser_based_authentication_portal_settings.0.portal_web_settings.0.main_url"); ok {
								portalWebSettingsPayload["main-url"] = v.(string)
							}
							browserBasedAuthenticationPortalSettingsPayload["portal-web-settings"] = portalWebSettingsPayload
						}
						browserBasedAuthenticationSettingsPayload["browser-based-authentication-portal-settings"] = browserBasedAuthenticationPortalSettingsPayload
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
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings"); ok {
						authenticationSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.authentication_method"); ok {
							authenticationSettingsPayload["authentication-method"] = v.(string)
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.radius"); ok {
							authenticationSettingsPayload["radius"] = v.(string)
						}
						if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories"); ok {

							usersDirectoriesPayload := make(map[string]interface{})

							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories.0.external_user_profile"); ok {
								usersDirectoriesPayload["external-user-profile"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories.0.internal_users"); ok {
								usersDirectoriesPayload["internal-users"] = v.(bool)
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories.0.specific"); ok {
								usersDirectoriesPayload["specific"] = v.(*schema.Set).List()
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.authentication_settings.0.users_directories.0.users_from_external_directories"); ok {
								usersDirectoriesPayload["users-from-external-directories"] = v.(string)
							}
							authenticationSettingsPayload["users-directories"] = usersDirectoriesPayload
						}
						identityAgentSettingsPayload["authentication-settings"] = authenticationSettingsPayload
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings"); ok {
						identityAgentPortalSettingsPayload := make(map[string]interface{})

						if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility"); ok {

							accessibilityPayload := make(map[string]interface{})

							if v, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.allow_access_from"); ok {
								accessibilityPayload["allow-access-from"] = v.(string)
							}
							if _, ok := d.GetOk("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.internal_access_settings"); ok {

								internalAccessSettingsPayload := make(map[string]interface{})

								if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.internal_access_settings.0.dmz"); ok {
									internalAccessSettingsPayload["dmz"] = v.(bool)
								}
								if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.internal_access_settings.0.undefined"); ok {
									internalAccessSettingsPayload["undefined"] = v.(bool)
								}
								if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_agent_settings.0.identity_agent_portal_settings.0.accessibility.0.internal_access_settings.0.vpn"); ok {
									internalAccessSettingsPayload["vpn"] = v.(bool)
								}
								accessibilityPayload["internal-access-settings"] = internalAccessSettingsPayload
							}
							identityAgentPortalSettingsPayload["accessibility"] = accessibilityPayload
						}
						identityAgentSettingsPayload["identity-agent-portal-settings"] = identityAgentPortalSettingsPayload
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
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings"); ok {
						authenticationSettingsPayload := make(map[string]interface{})

						if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories"); ok {

							usersDirectoriesPayload := make(map[string]interface{})

							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories.0.external_user_profile"); ok {
								usersDirectoriesPayload["external-user-profile"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories.0.internal_users"); ok {
								usersDirectoriesPayload["internal-users"] = v.(bool)
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories.0.specific"); ok {
								usersDirectoriesPayload["specific"] = v.(*schema.Set).List()
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.authentication_settings.0.users_directories.0.users_from_external_directories"); ok {
								usersDirectoriesPayload["users-from-external-directories"] = v.(string)
							}
							authenticationSettingsPayload["users-directories"] = usersDirectoriesPayload
						}
						identityCollectorSettingsPayload["authentication-settings"] = authenticationSettingsPayload
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions"); ok {
						clientAccessPermissionsPayload := make(map[string]interface{})

						if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility"); ok {

							accessibilityPayload := make(map[string]interface{})

							if v, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.allow_access_from"); ok {
								accessibilityPayload["allow-access-from"] = v.(string)
							}
							if _, ok := d.GetOk("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings"); ok {

								internalAccessSettingsPayload := make(map[string]interface{})

								if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings.0.dmz"); ok {
									internalAccessSettingsPayload["dmz"] = v.(bool)
								}
								if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings.0.undefined"); ok {
									internalAccessSettingsPayload["undefined"] = v.(bool)
								}
								if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_collector_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings.0.vpn"); ok {
									internalAccessSettingsPayload["vpn"] = v.(bool)
								}
								accessibilityPayload["internal-access-settings"] = internalAccessSettingsPayload
							}
							clientAccessPermissionsPayload["accessibility"] = accessibilityPayload
						}
						identityCollectorSettingsPayload["client-access-permissions"] = clientAccessPermissionsPayload
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
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.cache_mode"); ok {

						cacheModePayload := make(map[string]interface{})

						if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.cache_mode.0.override_profile"); ok {
							cacheModePayload["override-profile"] = v.(bool)
						}
						if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.cache_mode.0.value"); ok {
							cacheModePayload["value"] = v.(bool)
						}
						identitySharingSettingsPayload["cache-mode"] = cacheModePayload
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.cache_mode_duration"); ok {

						cacheModeDurationPayload := make(map[string]interface{})

						if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.cache_mode_duration.0.override_profile"); ok {
							cacheModeDurationPayload["override-profile"] = v.(bool)
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.identity_sharing_settings.cache_mode_duration.0.value"); ok {
							cacheModeDurationPayload["value"] = v.(int)
						}
						identitySharingSettingsPayload["cache-mode-duration"] = cacheModeDurationPayload
					}
					if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.receive_from_infinity_identity"); ok {
						identitySharingSettingsPayload["receive-from-infinity-identity"] = v.(bool)
					}
					if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_sharing_settings.scaled_sharing"); ok {
						identitySharingSettingsPayload["scaled-sharing"] = v.(bool)
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
				if v, ok := d.GetOk("identity_awareness_settings.0.identity_based_enforcement"); ok {
					identityAwarenessSettingsPayload["identity-based-enforcement"] = v.(string)
				}
				if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_web_api"); ok {
					identityAwarenessSettingsPayload["identity-web-api"] = v.(bool)
				}
				if _, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings"); ok {
					identityWebApiSettingsPayload := make(map[string]interface{})
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.authentication_settings"); ok {
						authenticationSettingsPayload := make(map[string]interface{})
						if _, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.authentication_settings.0.users_directories"); ok {
							usersDirectoriesPayload := make(map[string]interface{})
							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_web_api_settings.0.authentication_settings.0.users_directories.0.external_user_profile"); ok {
								usersDirectoriesPayload["external-user-profile"] = v.(bool)
							}
							if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_web_api_settings.0.authentication_settings.0.users_directories.0.internal_users"); ok {
								usersDirectoriesPayload["internal-users"] = v.(bool)
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.authentication_settings.0.users_directories.0.users_from_external_directories"); ok {
								usersDirectoriesPayload["users-from-external-directories"] = v.(string)
							}
							if v, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.authentication_settings.0.users_directories.0.specific"); ok {
								usersDirectoriesPayload["specific"] = v.(*schema.Set).List()
							}
							authenticationSettingsPayload["users-directories"] = usersDirectoriesPayload
						}
						identityWebApiSettingsPayload["authentication-settings"] = authenticationSettingsPayload
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.authorized_clients"); ok {
						authorizedClientsPayload := make(map[string]interface{})
						if v, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.authorized_clients.0.client"); ok {
							authorizedClientsPayload["client"] = v.(string)
						}
						if v, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.authorized_clients.0.client_secret"); ok {
							authorizedClientsPayload["client-secret"] = v.(string)
						}
						identityWebApiSettingsPayload["authorized-clients"] = authorizedClientsPayload
					}
					if _, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.client_access_permissions"); ok {
						clientAccessPermissionsPayload := make(map[string]interface{})
						if _, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.client_access_permissions.0.accessibility"); ok {
							accessibilityPayload := make(map[string]interface{})
							if v, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.client_access_permissions.0.accessibility.0.allow_access_from"); ok {
								accessibilityPayload["allow-access-from"] = v.(string)
							}
							if _, ok := d.GetOk("identity_awareness_settings.0.identity_web_api_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings"); ok {
								internalAccessSettingsPayload := make(map[string]interface{})
								if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_web_api_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings.0.dmz"); ok {
									internalAccessSettingsPayload["dmz"] = v.(bool)
								}
								if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_web_api_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings.0.undefined"); ok {
									internalAccessSettingsPayload["undefined"] = v.(bool)
								}
								if v, ok := d.GetOkExists("identity_awareness_settings.0.identity_web_api_settings.0.client_access_permissions.0.accessibility.0.internal_access_settings.0.vpn"); ok {
									internalAccessSettingsPayload["vpn"] = v.(bool)
								}
								accessibilityPayload["internal-access-settings"] = internalAccessSettingsPayload
							}
							clientAccessPermissionsPayload["accessibility"] = accessibilityPayload
						}
						identityWebApiSettingsPayload["client-access-permissions"] = clientAccessPermissionsPayload
					}
					identityAwarenessSettingsPayload["identity-web-api-settings"] = identityWebApiSettingsPayload
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

		if v, ok := d.GetOk("nat_settings"); ok {

			natSettingsList := v.([]interface{})

			if len(natSettingsList) > 0 {

				natSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("nat_settings.0.auto_rule"); ok {
					natSettingsPayload["auto-rule"] = v.(bool)
				}
				if v, ok := d.GetOk("nat_settings.0.ipv4_address"); ok {
					natSettingsPayload["ipv4-address"] = v.(string)
				}
				if v, ok := d.GetOk("nat_settings.0.ipv6_address"); ok {
					natSettingsPayload["ipv6-address"] = v.(string)
				}
				if v, ok := d.GetOk("nat_settings.0.hide_behind"); ok {
					natSettingsPayload["hide-behind"] = v.(string)
				}
				if v, ok := d.GetOk("nat_settings.0.install_on"); ok {
					natSettingsPayload["install-on"] = v.(string)
				}
				if v, ok := d.GetOk("nat_settings.0.method"); ok {
					natSettingsPayload["method"] = v.(string)
				}
				if v, ok := d.GetOkExists("nat_settings.0.apply_control_connections"); ok {
					natSettingsPayload["apply-control-connections"] = v.(bool)
				}
				cluster["nat-settings"] = natSettingsPayload
			}
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
					if v, ok := d.GetOk("platform_portal_settings.0.portal_web_settings.0.ip_address"); ok {
						portalWebSettingsPayload["ip-address"] = v.(string)
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

		if v, ok := d.GetOk("proxy_settings"); ok {

			proxySettingsList := v.([]interface{})

			if len(proxySettingsList) > 0 {

				proxySettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("proxy_settings.0.use_custom_proxy"); ok {
					proxySettingsPayload["use-custom-proxy"] = v.(bool)
				}
				if v, ok := d.GetOk("proxy_settings.0.proxy_server"); ok {
					proxySettingsPayload["proxy-server"] = v.(string)
				}
				if v, ok := d.GetOk("proxy_settings.0.port"); ok {
					proxySettingsPayload["port"] = v.(int)
				}
				cluster["proxy-settings"] = proxySettingsPayload
			}
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
					if v, ok := d.GetOk("usercheck_portal_settings.0.portal_web_settings.0.ip_address"); ok {
						portalWebSettingsPayload["ip-address"] = v.(string)
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
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.action"); ok {
						antiSpoofingSettings["action"] = v.(string)
					}
					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.exclude_packets"); ok {
						antiSpoofingSettings["exclude-packets"] = v.(bool)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.excluded_network_name"); ok {
						antiSpoofingSettings["excluded-network-name"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.excluded_network_uid"); ok {
						antiSpoofingSettings["excluded-network-uid"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".anti_spoofing_settings.0.spoof_tracking"); ok {
						antiSpoofingSettings["spoof-tracking"] = v.(string)
					}
					interfacePayload["anti-spoofing-settings"] = antiSpoofingSettings
				}
				if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".dynamic_ip"); ok {
					interfacePayload["dynamic-ip"] = v.(bool)
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
					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".security_zone_settings.0.auto_calculated"); ok {
						securityZoneSettings["auto-calculated"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".security_zone_settings.0.specific_zone"); ok {
						securityZoneSettings["specific-zone"] = v.(string)
					}
					interfacePayload["security-zone-settings"] = securityZoneSettings
				}
				if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology"); ok {
					interfacePayload["topology"] = v.(string)
				}
				if _, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings"); ok {
					topologySettings := make(map[string]interface{})

					if v, ok := d.GetOkExists("interfaces." + strconv.Itoa(i) + ".topology_settings.0.interface_leads_to_dmz"); ok {
						topologySettings["interface-leads-to-dmz"] = v
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.0.ip_address_behind_this_interface"); ok {
						topologySettings["ip-address-behind-this-interface"] = v.(string)
					}
					if v, ok := d.GetOk("interfaces." + strconv.Itoa(i) + ".topology_settings.0.specific_network"); ok {
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
				if v, ok := d.GetOkExists("members." + strconv.Itoa(i) + ".auto_generate_ip"); ok {
					memberPayload["auto-generate-ip"] = v.(bool)
				}
				if v, ok := d.GetOk("members." + strconv.Itoa(i) + ".trust_method"); ok {
					memberPayload["trust-method"] = v.(string)
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
							if v, ok := d.GetOkExists(memberInterfacePrefix + ".anti_spoofing"); ok {
								interfacePayload["anti-spoofing"] = v.(bool)
							}
							if v, ok := d.GetOkExists(memberInterfacePrefix + ".dynamic_ip"); ok {
								interfacePayload["dynamic-ip"] = v.(bool)
							}
							if v, ok := d.GetOkExists(memberInterfacePrefix + ".security_zone"); ok {
								interfacePayload["security-zone"] = v.(bool)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".topology"); ok {
								interfacePayload["topology"] = v.(string)
							}
							if v, ok := d.GetOk(memberInterfacePrefix + ".ipv6_mask_length"); ok {
								interfacePayload["ipv6-mask-length"] = v.(string)
							}
							if _, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings"); ok {
								antiSpoofingSettingsPayload := make(map[string]interface{})
								if v, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings.0.action"); ok {
									antiSpoofingSettingsPayload["action"] = v.(string)
								}
								if v, ok := d.GetOkExists(memberInterfacePrefix + ".anti_spoofing_settings.0.exclude_packets"); ok {
									antiSpoofingSettingsPayload["exclude-packets"] = v.(bool)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings.0.excluded_network_name"); ok {
									antiSpoofingSettingsPayload["excluded-network-name"] = v.(string)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings.0.excluded_network_uid"); ok {
									antiSpoofingSettingsPayload["excluded-network-uid"] = v.(string)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".anti_spoofing_settings.0.spoof_tracking"); ok {
									antiSpoofingSettingsPayload["spoof-tracking"] = v.(string)
								}
								interfacePayload["anti-spoofing-settings"] = antiSpoofingSettingsPayload
							}
							if _, ok := d.GetOk(memberInterfacePrefix + ".security_zone_settings"); ok {
								securityZoneSettingsPayload := make(map[string]interface{})
								if v, ok := d.GetOkExists(memberInterfacePrefix + ".security_zone_settings.0.auto_calculated"); ok {
									securityZoneSettingsPayload["auto-calculated"] = v.(bool)
								}
								interfacePayload["security-zone-settings"] = securityZoneSettingsPayload
							}
							if _, ok := d.GetOk(memberInterfacePrefix + ".topology_settings"); ok {
								topologySettingsPayload := make(map[string]interface{})
								if v, ok := d.GetOkExists(memberInterfacePrefix + ".topology_settings.0.interface_leads_to_dmz"); ok {
									topologySettingsPayload["interface-leads-to-dmz"] = v.(bool)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".topology_settings.0.ip_address_behind_this_interface"); ok {
									topologySettingsPayload["ip-address-behind-this-interface"] = v.(string)
								}
								if v, ok := d.GetOk(memberInterfacePrefix + ".topology_settings.0.specific_network"); ok {
									topologySettingsPayload["specific-network"] = v.(string)
								}
								interfacePayload["topology-settings"] = topologySettingsPayload
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
		if v, ok := d.GetOk("firewall_settings"); ok {

			firewallSettingsList := v.([]interface{})

			if len(firewallSettingsList) > 0 {

				firewallSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("firewall_settings.0.auto_calculate_connections_hash_table_size_and_memory_pool"); ok {
					firewallSettingsPayload["auto-calculate-connections-hash-table-size-and-memory-pool"] = v.(bool)
				}
				if v, ok := d.GetOkExists("firewall_settings.0.auto_maximum_limit_for_concurrent_connections"); ok {
					firewallSettingsPayload["auto-maximum-limit-for-concurrent-connections"] = v.(bool)
				}
				if v, ok := d.GetOk("firewall_settings.0.connections_hash_size"); ok {
					firewallSettingsPayload["connections-hash-size"] = v.(int)
				}
				if v, ok := d.GetOk("firewall_settings.0.maximum_limit_for_concurrent_connections"); ok {
					firewallSettingsPayload["maximum-limit-for-concurrent-connections"] = v.(int)
				}
				if v, ok := d.GetOk("firewall_settings.0.maximum_memory_pool_size"); ok {
					firewallSettingsPayload["maximum-memory-pool-size"] = v.(int)
				}
				if v, ok := d.GetOk("firewall_settings.0.memory_pool_size"); ok {
					firewallSettingsPayload["memory-pool-size"] = v.(int)
				}
				cluster["firewall-settings"] = firewallSettingsPayload
			}
		}
	}

	// VPN settings
	if ok := d.HasChange("vpn_settings"); ok {
		if v, ok := d.GetOk("vpn_settings"); ok {

			vpnSettingsList := v.([]interface{})

			if len(vpnSettingsList) > 0 {

				vpnSettingsPayload := make(map[string]interface{})

				if _, ok := d.GetOk("vpn_settings.0.advanced"); ok {

					advancedPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.advanced.0.tunnel_sharing_mode"); ok {
						advancedPayload["tunnel-sharing-mode"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.0.advanced.0.shutdown_on_gateway_restart"); ok {
						advancedPayload["shutdown-on-gateway-restart"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.advanced.0.enable_wire_mode"); ok {
						advancedPayload["enable-wire-mode"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("vpn_settings.0.advanced.0.wire_mode_interfaces"); ok {
						advancedPayload["wire-mode-interfaces"] = v.(*schema.Set).List()
					}
					if v, ok := d.GetOkExists("vpn_settings.0.advanced.0.enable_wire_mode_log_traffic"); ok {
						advancedPayload["enable-wire-mode-log-traffic"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.advanced.0.enable_nat_traversal"); ok {
						advancedPayload["enable-nat-traversal"] = strconv.FormatBool(v.(bool))
					}
					vpnSettingsPayload["advanced"] = advancedPayload
				}
				if _, ok := d.GetOk("vpn_settings.0.authentication"); ok {

					authenticationPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.authentication.0.authentication_clients"); ok {
						authenticationPayload["authentication-clients"] = v.(*schema.Set).List()
					}
					if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client"); ok {

						singleAuthenticationClientPayload := make(map[string]interface{})

						if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.single_authentication_client.0.enabled"); ok {
							singleAuthenticationClientPayload["enabled"] = v.(bool)
						}
						if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.single_authentication_client.0.allow_multiple_authentication_clients"); ok {
							singleAuthenticationClientPayload["allow-multiple-authentication-clients"] = v.(bool)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.display_name"); ok {
							singleAuthenticationClientPayload["display-name"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.method"); ok {
							singleAuthenticationClientPayload["method"] = v.(string)
						}
						if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.secur_id"); ok {

							securIdPayload := make(map[string]interface{})

							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.secur_id.0.server"); ok {
								securIdPayload["server"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.secur_id.0.token_card_type"); ok {
								securIdPayload["token-card-type"] = v.(string)
							}
							singleAuthenticationClientPayload["secur-id"] = securIdPayload
						}
						if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.radius"); ok {

							radiusPayload := make(map[string]interface{})

							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.radius.0.server"); ok {
								radiusPayload["server"] = v.(string)
							}
							if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.single_authentication_client.0.radius.0.ask_user_password"); ok {
								radiusPayload["ask-user-password"] = v.(bool)
							}
							singleAuthenticationClientPayload["radius"] = radiusPayload
						}
						if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate"); ok {

							personalCertificatePayload := make(map[string]interface{})

							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.fetch_username_from"); ok {
								personalCertificatePayload["fetch-username-from"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.storage_type"); ok {
								personalCertificatePayload["storage-type"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.source"); ok {
								personalCertificatePayload["source"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.dn_part"); ok {
								personalCertificatePayload["dn-part"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.personal_certificate.0.dn_concurrence"); ok {
								personalCertificatePayload["dn-concurrence"] = v.(int)
							}
							singleAuthenticationClientPayload["personal-certificate"] = personalCertificatePayload
						}
						if _, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.client_display_settings"); ok {

							clientDisplaySettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.client_display_settings.0.headline"); ok {
								clientDisplaySettingsPayload["headline"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.client_display_settings.0.username_label"); ok {
								clientDisplaySettingsPayload["username-label"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.single_authentication_client.0.client_display_settings.0.password_label"); ok {
								clientDisplaySettingsPayload["password-label"] = v.(string)
							}
							singleAuthenticationClientPayload["client-display-settings"] = clientDisplaySettingsPayload
						}
						authenticationPayload["single-authentication-client"] = singleAuthenticationClientPayload
					}
					if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.override_global_dynamic_id_settings"); ok {
						authenticationPayload["override-global-dynamic-id-settings"] = v.(bool)
					}
					if _, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings"); ok {

						dynamicIdSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_and_email_settings"); ok {
							dynamicIdSettingsPayload["sms-provider-and-email-settings"] = v.(string)
						}
						if _, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_credentials"); ok {

							smsProviderCredentialsPayload := make(map[string]interface{})

							if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_credentials.0.username"); ok {
								smsProviderCredentialsPayload["username"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_credentials.0.password"); ok {
								smsProviderCredentialsPayload["password"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.sms_provider_credentials.0.api_id"); ok {
								smsProviderCredentialsPayload["api-id"] = v.(string)
							}
							dynamicIdSettingsPayload["sms-provider-credentials"] = smsProviderCredentialsPayload
						}
						if _, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings"); ok {

							advancedSettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.dynamic_id_message"); ok {
								advancedSettingsPayload["dynamic-id-message"] = v.(string)
							}
							if _, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.otp_settings"); ok {

								otpSettingsPayload := make(map[string]interface{})

								if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.otp_settings.0.length"); ok {
									otpSettingsPayload["length"] = v.(int)
								}
								if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.otp_settings.0.expiration"); ok {
									otpSettingsPayload["expiration"] = v.(int)
								}
								if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.otp_settings.0.max_attempts"); ok {
									otpSettingsPayload["max-attempts"] = v.(int)
								}
								advancedSettingsPayload["otp-settings"] = otpSettingsPayload
							}
							if v, ok := d.GetOkExists("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.enable_display_user_details"); ok {
								advancedSettingsPayload["enable-display-user-details"] = v.(bool)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.country_code"); ok {
								advancedSettingsPayload["country-code"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.authentication.0.dynamic_id_settings.0.advanced_settings.0.user_details_retrieval"); ok {
								advancedSettingsPayload["user-details-retrieval"] = v.(string)
							}
							dynamicIdSettingsPayload["advanced-settings"] = advancedSettingsPayload
						}
						authenticationPayload["dynamic-id-settings"] = dynamicIdSettingsPayload
					}
					if v, ok := d.GetOk("vpn_settings.0.authentication.0.send_machine_certificate"); ok {
						authenticationPayload["send-machine-certificate"] = v.(string)
					}
					vpnSettingsPayload["authentication"] = authenticationPayload
				}
				if _, ok := d.GetOk("vpn_settings.0.certificates"); ok {

					certificatesPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.certificates.0.name"); ok {
						certificatesPayload["name"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.certificates.0.certificate_authority"); ok {
						certificatesPayload["certificate-authority"] = v.(string)
					}
					if _, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment"); ok {

						enrollmentPayload := make(map[string]interface{})

						if _, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings"); ok {

							enrollmentSettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings.0.distinguished_name"); ok {
								enrollmentSettingsPayload["distinguished-name"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings.0.alternate_names"); ok {

								alternateNamesList := v.([]interface{})

								if len(alternateNamesList) > 0 {

									var alternateNamesPayload []map[string]interface{}

									for j := range alternateNamesList {

										alternateNamesMapToAdd := make(map[string]interface{})

										if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings.0.alternate_names." + strconv.Itoa(j) + ".name_type"); ok {
											alternateNamesMapToAdd["name-type"] = v.(string)
										}
										if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_settings.0.alternate_names." + strconv.Itoa(j) + ".value"); ok {
											alternateNamesMapToAdd["value"] = v.(string)
										}
										alternateNamesPayload = append(alternateNamesPayload, alternateNamesMapToAdd)
									}
									enrollmentSettingsPayload["alternate-names"] = alternateNamesPayload
								}
							}
							enrollmentPayload["enrollment-settings"] = enrollmentSettingsPayload
						}
						if v, ok := d.GetOk("vpn_settings.0.certificates.0.enrollment.0.enrollment_type"); ok {
							enrollmentPayload["enrollment-type"] = v.(string)
						}
						certificatesPayload["enrollment"] = enrollmentPayload
					}
					if v, ok := d.GetOk("vpn_settings.0.certificates.0.stored_at"); ok {
						certificatesPayload["stored-at"] = v.(string)
					}
					vpnSettingsPayload["certificates"] = certificatesPayload
				}
				if _, ok := d.GetOk("vpn_settings.0.exported_routes"); ok {

					exportedRoutesPayload := make(map[string]interface{})

					if v, ok := d.GetOkExists("vpn_settings.0.exported_routes.0.internal_interfaces"); ok {
						exportedRoutesPayload["internal-interfaces"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.exported_routes.0.static_routes"); ok {
						exportedRoutesPayload["static-routes"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.exported_routes.0.custom_routes"); ok {
						exportedRoutesPayload["custom-routes"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("vpn_settings.0.exported_routes.0.custom_routes_object"); ok {
						exportedRoutesPayload["custom-routes-object"] = v.(string)
					}
					vpnSettingsPayload["exported-routes"] = exportedRoutesPayload
				}
				if _, ok := d.GetOk("vpn_settings.0.link_selection"); ok {

					linkSelectionPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.ip_selection"); ok {
						linkSelectionPayload["ip-selection"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.ip_address"); ok {
						linkSelectionPayload["ip-address"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.dns_resolving_hostname"); ok {
						linkSelectionPayload["dns-resolving-hostname"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.route_selection_method"); ok {
						linkSelectionPayload["route-selection-method"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.responding_traffic"); ok {
						linkSelectionPayload["responding-traffic"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.source_ip_selection"); ok {
						linkSelectionPayload["source-ip-selection"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.selected_ip"); ok {
						linkSelectionPayload["selected-ip"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.link_selection.0.outgoing_link_tracking"); ok {
						linkSelectionPayload["outgoing-link-tracking"] = v.(string)
					}
					if _, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings"); ok {

						probingSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings.0.probed_interfaces"); ok {
							probingSettingsPayload["probed-interfaces"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings.0.probed_interface_list"); ok {
							probingSettingsPayload["probed-interface-list"] = v.(*schema.Set).List()
						}
						if v, ok := d.GetOkExists("vpn_settings.0.link_selection.0.probing_settings.0.use_primary_address"); ok {
							probingSettingsPayload["use-primary-address"] = v.(bool)
						}
						if v, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings.0.primary_address"); ok {
							probingSettingsPayload["primary-address"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.link_selection.0.probing_settings.0.probing_method"); ok {
							probingSettingsPayload["probing-method"] = v.(string)
						}
						linkSelectionPayload["probing-settings"] = probingSettingsPayload
					}
					vpnSettingsPayload["link-selection"] = linkSelectionPayload
				}
				if v, ok := d.GetOk("vpn_settings.0.maximum_concurrent_ike_negotiations"); ok {
					vpnSettingsPayload["maximum-concurrent-ike-negotiations"] = v.(int)
				}
				if v, ok := d.GetOk("vpn_settings.0.maximum_concurrent_tunnels"); ok {
					vpnSettingsPayload["maximum-concurrent-tunnels"] = v.(int)
				}
				if _, ok := d.GetOk("vpn_settings.0.office_mode"); ok {

					officeModePayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.office_mode.0.mode"); ok {
						officeModePayload["mode"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.office_mode.0.group"); ok {
						officeModePayload["group"] = v.(string)
					}
					if _, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from"); ok {

						allocateIpAddressFromPayload := make(map[string]interface{})

						if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.radius_server"); ok {
							allocateIpAddressFromPayload["radius-server"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.use_allocate_method"); ok {
							allocateIpAddressFromPayload["use-allocate-method"] = strconv.FormatBool(v.(bool))
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.allocate_method"); ok {
							allocateIpAddressFromPayload["allocate-method"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.manual_network"); ok {
							allocateIpAddressFromPayload["manual-network"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.dhcp_server"); ok {
							allocateIpAddressFromPayload["dhcp-server"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.virtual_ip_address"); ok {
							allocateIpAddressFromPayload["virtual-ip-address"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.dhcp_mac_address"); ok {
							allocateIpAddressFromPayload["dhcp-mac-address"] = v.(string)
						}
						if _, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters"); ok {

							optionalParametersPayload := make(map[string]interface{})

							if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_primary_dns_server"); ok {
								optionalParametersPayload["use-primary-dns-server"] = strconv.FormatBool(v.(bool))
							}
							if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.primary_dns_server"); ok {
								optionalParametersPayload["primary-dns-server"] = v.(string)
							}
							if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_first_backup_dns_server"); ok {
								optionalParametersPayload["use-first-backup-dns-server"] = strconv.FormatBool(v.(bool))
							}
							if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.first_backup_dns_server"); ok {
								optionalParametersPayload["first-backup-dns-server"] = v.(string)
							}
							if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_second_backup_dns_server"); ok {
								optionalParametersPayload["use-second-backup-dns-server"] = strconv.FormatBool(v.(bool))
							}
							if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.second_backup_dns_server"); ok {
								optionalParametersPayload["second-backup-dns-server"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.dns_suffixes"); ok {
								optionalParametersPayload["dns-suffixes"] = v.(string)
							}
							if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_primary_wins_server"); ok {
								optionalParametersPayload["use-primary-wins-server"] = strconv.FormatBool(v.(bool))
							}
							if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.primary_wins_server"); ok {
								optionalParametersPayload["primary-wins-server"] = v.(string)
							}
							if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_first_backup_wins_server"); ok {
								optionalParametersPayload["use-first-backup-wins-server"] = strconv.FormatBool(v.(bool))
							}
							if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.first_backup_wins_server"); ok {
								optionalParametersPayload["first-backup-wins-server"] = v.(string)
							}
							if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.use_second_backup_wins_server"); ok {
								optionalParametersPayload["use-second-backup-wins-server"] = strconv.FormatBool(v.(bool))
							}
							if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.second_backup_wins_server"); ok {
								optionalParametersPayload["second-backup-wins-server"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.office_mode.0.allocate_ip_address_from.0.optional_parameters.0.ip_lease_duration"); ok {
								optionalParametersPayload["ip-lease-duration"] = v
							}
							allocateIpAddressFromPayload["optional-parameters"] = optionalParametersPayload
						}
						officeModePayload["allocate-ip-address-from"] = allocateIpAddressFromPayload
					}
					if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.support_multiple_interfaces"); ok {
						officeModePayload["support-multiple-interfaces"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.office_mode.0.perform_anti_spoofing"); ok {
						officeModePayload["perform-anti-spoofing"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("vpn_settings.0.office_mode.0.anti_spoofing_additional_addresses"); ok {
						officeModePayload["anti-spoofing-additional-addresses"] = v.(string)
					}
					vpnSettingsPayload["office-mode"] = officeModePayload
				}
				if _, ok := d.GetOk("vpn_settings.0.remote_access"); ok {

					remoteAccessPayload := make(map[string]interface{})

					if v, ok := d.GetOkExists("vpn_settings.0.remote_access.0.support_l2tp"); ok {
						remoteAccessPayload["support-l2tp"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("vpn_settings.0.remote_access.0.l2tp_auth_method"); ok {
						remoteAccessPayload["l2tp-auth-method"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.remote_access.0.l2tp_certificate"); ok {
						remoteAccessPayload["l2tp-certificate"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.0.remote_access.0.allow_vpn_clients_to_route_traffic"); ok {
						remoteAccessPayload["allow-vpn-clients-to-route-traffic"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.remote_access.0.support_nat_traversal_mechanism"); ok {
						remoteAccessPayload["support-nat-traversal-mechanism"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("vpn_settings.0.remote_access.0.nat_traversal_service"); ok {
						remoteAccessPayload["nat-traversal-service"] = v.(string)
					}
					if v, ok := d.GetOkExists("vpn_settings.0.remote_access.0.support_visitor_mode"); ok {
						remoteAccessPayload["support-visitor-mode"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("vpn_settings.0.remote_access.0.visitor_mode_service"); ok {
						remoteAccessPayload["visitor-mode-service"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.remote_access.0.visitor_mode_interface"); ok {
						remoteAccessPayload["visitor-mode-interface"] = v.(string)
					}
					vpnSettingsPayload["remote-access"] = remoteAccessPayload
				}
				if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings"); ok {

					samlPortalSettingsPayload := make(map[string]interface{})

					if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.portal_web_settings"); ok {

						portalWebSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.portal_web_settings.0.aliases"); ok {
							portalWebSettingsPayload["aliases"] = v.(*schema.Set).List()
						}
						if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.portal_web_settings.0.ip_address"); ok {
							portalWebSettingsPayload["ip-address"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.portal_web_settings.0.main_url"); ok {
							portalWebSettingsPayload["main-url"] = v.(string)
						}
						samlPortalSettingsPayload["portal-web-settings"] = portalWebSettingsPayload
					}
					if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.certificate_settings"); ok {

						certificateSettingsPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.certificate_settings.0.base64_certificate"); ok {
							certificateSettingsPayload["base64-certificate"] = v.(string)
						}
						if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.certificate_settings.0.base64_password"); ok {
							certificateSettingsPayload["base64-password"] = v.(string)
						}
						samlPortalSettingsPayload["certificate-settings"] = certificateSettingsPayload
					}
					if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.accessibility"); ok {

						accessibilityPayload := make(map[string]interface{})

						if v, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.accessibility.0.allow_access_from"); ok {
							accessibilityPayload["allow-access-from"] = v.(string)
						}
						if _, ok := d.GetOk("vpn_settings.0.saml_portal_settings.0.accessibility.0.internal_access_settings"); ok {

							internalAccessSettingsPayload := make(map[string]interface{})

							if v, ok := d.GetOkExists("vpn_settings.0.saml_portal_settings.0.accessibility.0.internal_access_settings.0.undefined"); ok {
								internalAccessSettingsPayload["undefined"] = strconv.FormatBool(v.(bool))
							}
							if v, ok := d.GetOkExists("vpn_settings.0.saml_portal_settings.0.accessibility.0.internal_access_settings.0.dmz"); ok {
								internalAccessSettingsPayload["dmz"] = strconv.FormatBool(v.(bool))
							}
							if v, ok := d.GetOkExists("vpn_settings.0.saml_portal_settings.0.accessibility.0.internal_access_settings.0.vpn"); ok {
								internalAccessSettingsPayload["vpn"] = strconv.FormatBool(v.(bool))
							}
							accessibilityPayload["internal-access-settings"] = internalAccessSettingsPayload
						}
						samlPortalSettingsPayload["accessibility"] = accessibilityPayload
					}
					vpnSettingsPayload["saml-portal-settings"] = samlPortalSettingsPayload
				}
				if _, ok := d.GetOk("vpn_settings.0.vpn_clients"); ok {

					vpnClientsPayload := make(map[string]interface{})

					if v, ok := d.GetOkExists("vpn_settings.0.vpn_clients.0.enable_endpoint_security_vpn"); ok {
						vpnClientsPayload["enable-endpoint-security-vpn"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.vpn_clients.0.enable_cp_mobile_for_windows"); ok {
						vpnClientsPayload["enable-cp-mobile-for-windows"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.vpn_clients.0.enable_secu_remote"); ok {
						vpnClientsPayload["enable-secu-remote"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.vpn_clients.0.enable_capsule_vpn_connect"); ok {
						vpnClientsPayload["enable-capsule-vpn-connect"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOkExists("vpn_settings.0.vpn_clients.0.enable_ssl_network_extender"); ok {
						vpnClientsPayload["enable-ssl-network-extender"] = strconv.FormatBool(v.(bool))
					}
					if v, ok := d.GetOk("vpn_settings.0.vpn_clients.0.gateway_authentication_certificate"); ok {
						vpnClientsPayload["gateway-authentication-certificate"] = v.(string)
					}
					vpnSettingsPayload["vpn-clients"] = vpnClientsPayload
				}
				if v, ok := d.GetOk("vpn_settings.0.vpn_domain"); ok {
					vpnSettingsPayload["vpn-domain"] = v.(string)
				}
				if v, ok := d.GetOkExists("vpn_settings.0.vpn_domain_exclude_external_ip_addresses"); ok {
					vpnSettingsPayload["vpn-domain-exclude-external-ip-addresses"] = v.(bool)
				}
				if v, ok := d.GetOk("vpn_settings.0.vpn_domain_type"); ok {
					vpnSettingsPayload["vpn-domain-type"] = v.(string)
				}
				if v, ok := d.GetOkExists("vpn_settings.0.enable_clientless_vpn"); ok {
					vpnSettingsPayload["enable-clientless-vpn"] = v.(bool)
				}
				if _, ok := d.GetOk("vpn_settings.0.clientless_vpn_settings"); ok {

					clientlessVpnSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("vpn_settings.0.clientless_vpn_settings.0.certificate_gateway_authentication"); ok {
						clientlessVpnSettingsPayload["certificate-gateway-authentication"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.clientless_vpn_settings.0.client_authentication"); ok {
						clientlessVpnSettingsPayload["client-authentication"] = v.(string)
					}
					if v, ok := d.GetOk("vpn_settings.0.clientless_vpn_settings.0.concurrent_servers_or_processes"); ok {
						clientlessVpnSettingsPayload["concurrent-servers-or-processes"] = v
					}
					if v, ok := d.GetOkExists("vpn_settings.0.clientless_vpn_settings.0.accept_only_3des"); ok {
						clientlessVpnSettingsPayload["accept-only-3des"] = strconv.FormatBool(v.(bool))
					}
					vpnSettingsPayload["clientless-vpn-settings"] = clientlessVpnSettingsPayload
				}
				if v, ok := d.GetOk("vpn_settings.0.interfaces"); ok {
					interfacesList := v.([]interface{})
					if len(interfacesList) > 0 {
						var interfacesPayload []map[string]interface{}
						for i := range interfacesList {
							interfacesItem := make(map[string]interface{})
							if v, ok := d.GetOk("vpn_settings.0.interfaces." + strconv.Itoa(i) + ".interface_name"); ok {
								interfacesItem["interface-name"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.interfaces." + strconv.Itoa(i) + ".ip_version"); ok {
								interfacesItem["ip-version"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.interfaces." + strconv.Itoa(i) + ".next_hop_ip"); ok {
								interfacesItem["next-hop-ip"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.interfaces." + strconv.Itoa(i) + ".priority"); ok {
								interfacesItem["priority"] = v.(int)
							}
							if v, ok := d.GetOk("vpn_settings.0.interfaces." + strconv.Itoa(i) + ".redundancy_mode"); ok {
								interfacesItem["redundancy-mode"] = v.(string)
							}
							if v, ok := d.GetOk("vpn_settings.0.interfaces." + strconv.Itoa(i) + ".static_nat_ip"); ok {
								interfacesItem["static-nat-ip"] = v.(string)
							}
							interfacesPayload = append(interfacesPayload, interfacesItem)
						}
						vpnSettingsPayload["interfaces"] = interfacesPayload
					}
				}
				cluster["vpn-settings"] = vpnSettingsPayload
			}
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
		if ok := d.HasChange("anti_spam_and_email_security"); ok {
			cluster["anti-spam-and-email-security"] = d.Get("anti_spam_and_email_security")
		}
		if ok := d.HasChange("auto_topology_custom_recalculation_time"); ok {
			cluster["auto-topology-custom-recalculation-time"] = d.Get("auto_topology_custom_recalculation_time")
		}
		if ok := d.HasChange("auto_topology_use_custom_recalculation_time"); ok {
			cluster["auto-topology-use-custom-recalculation-time"] = d.Get("auto_topology_use_custom_recalculation_time")
		}
		if ok := d.HasChange("data_loss_prevention"); ok {
			cluster["data-loss-prevention"] = d.Get("data_loss_prevention")
		}
		if ok := d.HasChange("mobile_access"); ok {
			cluster["mobile-access"] = d.Get("mobile_access")
		}
		if ok := d.HasChange("monitoring"); ok {
			cluster["monitoring"] = d.Get("monitoring")
		}
		if ok := d.HasChange("policy_server"); ok {
			cluster["policy-server"] = d.Get("policy_server")
		}
		if ok := d.HasChange("rtm_counters_report"); ok {
			cluster["rtm-counters-report"] = d.Get("rtm_counters_report")
		}
		if ok := d.HasChange("rtm_traffic_report"); ok {
			cluster["rtm-traffic-report"] = d.Get("rtm_traffic_report")
		}
		if ok := d.HasChange("rtm_traffic_report_per_connection"); ok {
			cluster["rtm-traffic-report-per-connection"] = d.Get("rtm_traffic_report_per_connection")
		}
		if ok := d.HasChange("threat_extraction"); ok {
			cluster["threat-extraction"] = d.Get("threat_extraction")
		}
		if ok := d.HasChange("threat_prevention_mode"); ok {
			cluster["threat-prevention-mode"] = d.Get("threat_prevention_mode")
		}
		if ok := d.HasChange("workforce_ai"); ok {
			cluster["workforce-ai"] = d.Get("workforce_ai")
		}
		if ok := d.HasChange("application_control_and_url_filtering_settings"); ok {
			if _, ok := d.GetOk("application_control_and_url_filtering_settings"); ok {

				applicationControlAndUrlFilteringSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.global_settings_mode"); ok {
					applicationControlAndUrlFilteringSettingsPayload["global-settings-mode"] = v.(string)
				}
				if _, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings"); ok {

					overrideGlobalSettingsPayload := make(map[string]interface{})

					if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.fail_mode"); ok {
						overrideGlobalSettingsPayload["fail-mode"] = v.(string)
					}
					if _, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization"); ok {

						websiteCategorizationPayload := make(map[string]interface{})

						if _, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization.0.custom_mode"); ok {

							customModePayload := make(map[string]interface{})

							if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization.0.custom_mode.0.social_networking_widgets"); ok {
								customModePayload["social-networking-widgets"] = v.(string)
							}
							if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization.0.custom_mode.0.url_filtering"); ok {
								customModePayload["url-filtering"] = v.(string)
							}
							websiteCategorizationPayload["custom-mode"] = customModePayload
						}
						if v, ok := d.GetOk("application_control_and_url_filtering_settings.0.override_global_settings.0.website_categorization.0.mode"); ok {
							websiteCategorizationPayload["mode"] = v.(string)
						}
						overrideGlobalSettingsPayload["website-categorization"] = websiteCategorizationPayload
					}
					applicationControlAndUrlFilteringSettingsPayload["override-global-settings"] = overrideGlobalSettingsPayload
				}
				cluster["application-control-and-url-filtering-settings"] = applicationControlAndUrlFilteringSettingsPayload
			}
		}
		if ok := d.HasChange("cluster_settings"); ok {
			if _, ok := d.GetOk("cluster_settings"); ok {

				clusterSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("cluster_settings.0.member_recovery_mode"); ok {
					clusterSettingsPayload["member-recovery-mode"] = v.(string)
				}
				if _, ok := d.GetOk("cluster_settings.0.state_synchronization"); ok {

					stateSynchronizationPayload := make(map[string]interface{})

					if v, ok := d.GetOkExists("cluster_settings.0.state_synchronization.0.delayed"); ok {
						stateSynchronizationPayload["delayed"] = v.(bool)
					}
					if v, ok := d.GetOk("cluster_settings.0.state_synchronization.0.delayed_seconds"); ok {
						stateSynchronizationPayload["delayed-seconds"] = v.(int)
					}
					if v, ok := d.GetOkExists("cluster_settings.0.state_synchronization.0.enabled"); ok {
						stateSynchronizationPayload["enabled"] = v.(bool)
					}
					clusterSettingsPayload["state-synchronization"] = stateSynchronizationPayload
				}
				if v, ok := d.GetOk("cluster_settings.0.track_changes_of_cluster_members"); ok {
					clusterSettingsPayload["track-changes-of-cluster-members"] = v.(string)
				}
				if v, ok := d.GetOkExists("cluster_settings.0.use_virtual_mac"); ok {
					clusterSettingsPayload["use-virtual-mac"] = v.(bool)
				}
				cluster["cluster-settings"] = clusterSettingsPayload
			}
		}
		if ok := d.HasChange("communication_with_servers_behind_nat"); ok {
			if _, ok := d.GetOk("communication_with_servers_behind_nat"); ok {

				communicationWithServersBehindNatPayload := make(map[string]interface{})

				if v, ok := d.GetOkExists("communication_with_servers_behind_nat.0.override_profile"); ok {
					communicationWithServersBehindNatPayload["override-profile"] = v.(bool)
				}
				if v, ok := d.GetOk("communication_with_servers_behind_nat.0.value"); ok {
					communicationWithServersBehindNatPayload["value"] = v.(string)
				}
				cluster["communication-with-servers-behind-nat"] = communicationWithServersBehindNatPayload
			}
		}
		if ok := d.HasChange("zero_phishing_settings"); ok {
			if _, ok := d.GetOk("zero_phishing_settings"); ok {

				zeroPhishingSettingsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("zero_phishing_settings.0.gateway_fqdn_mode"); ok {
					zeroPhishingSettingsPayload["gateway-fqdn-mode"] = v.(string)
				}
				if v, ok := d.GetOk("zero_phishing_settings.0.manual_fqdn"); ok {
					zeroPhishingSettingsPayload["manual-fqdn"] = v.(string)
				}
				cluster["zero-phishing-settings"] = zeroPhishingSettingsPayload
			}
		}
		if ok := d.HasChange("dns_server"); ok {
			if v, ok := d.GetOkExists("dns_server"); ok {
				cluster["dns-server"] = v.(bool)
			}
		}

		if ok := d.HasChange("logs_settings"); ok {
			if _, ok := d.GetOk("logs_settings"); ok {
				logsSettingsPayload := make(map[string]interface{})
				if v, ok := d.GetOkExists("logs_settings.0.alert_when_free_disk_space_below"); ok {
					logsSettingsPayload["alert-when-free-disk-space-below"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.alert_when_free_disk_space_below_threshold"); ok {
					logsSettingsPayload["alert-when-free-disk-space-below-threshold"] = v.(int)
				}
				if v, ok := d.GetOk("logs_settings.0.alert_when_free_disk_space_below_type"); ok {
					logsSettingsPayload["alert-when-free-disk-space-below-type"] = v.(string)
				}
				if v, ok := d.GetOkExists("logs_settings.0.before_delete_keep_logs_from_the_last_days"); ok {
					logsSettingsPayload["before-delete-keep-logs-from-the-last-days"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.before_delete_keep_logs_from_the_last_days_threshold"); ok {
					logsSettingsPayload["before-delete-keep-logs-from-the-last-days-threshold"] = v.(int)
				}
				if v, ok := d.GetOkExists("logs_settings.0.before_delete_run_script"); ok {
					logsSettingsPayload["before-delete-run-script"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.before_delete_run_script_command"); ok {
					logsSettingsPayload["before-delete-run-script-command"] = v.(string)
				}
				if v, ok := d.GetOkExists("logs_settings.0.delete_index_files_older_than_days"); ok {
					logsSettingsPayload["delete-index-files-older-than-days"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.delete_index_files_older_than_days_threshold"); ok {
					logsSettingsPayload["delete-index-files-older-than-days-threshold"] = v.(int)
				}
				if v, ok := d.GetOkExists("logs_settings.0.delete_index_files_when_index_size_above"); ok {
					logsSettingsPayload["delete-index-files-when-index-size-above"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.delete_index_files_when_index_size_above_threshold"); ok {
					logsSettingsPayload["delete-index-files-when-index-size-above-threshold"] = v.(int)
				}
				if v, ok := d.GetOkExists("logs_settings.0.delete_when_free_disk_space_below"); ok {
					logsSettingsPayload["delete-when-free-disk-space-below"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.delete_when_free_disk_space_below_threshold"); ok {
					logsSettingsPayload["delete-when-free-disk-space-below-threshold"] = v.(int)
				}
				if v, ok := d.GetOkExists("logs_settings.0.detect_new_citrix_ica_application_names"); ok {
					logsSettingsPayload["detect-new-citrix-ica-application-names"] = v.(bool)
				}
				if v, ok := d.GetOkExists("logs_settings.0.distribute_logs_between_all_active_servers"); ok {
					logsSettingsPayload["distribute-logs-between-all-active-servers"] = v.(bool)
				}
				if v, ok := d.GetOkExists("logs_settings.0.forward_logs_to_log_server"); ok {
					logsSettingsPayload["forward-logs-to-log-server"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.forward_logs_to_log_server_name"); ok {
					logsSettingsPayload["forward-logs-to-log-server-name"] = v.(string)
				}
				if v, ok := d.GetOk("logs_settings.0.forward_logs_to_log_server_schedule_name"); ok {
					logsSettingsPayload["forward-logs-to-log-server-schedule-name"] = v.(string)
				}
				if v, ok := d.GetOk("logs_settings.0.free_disk_space_metrics"); ok {
					logsSettingsPayload["free-disk-space-metrics"] = v.(string)
				}
				if v, ok := d.GetOk("logs_settings.0.include_tcp_state_information"); ok {
					logsSettingsPayload["include-tcp-state-information"] = v.(string)
				}
				if v, ok := d.GetOkExists("logs_settings.0.perform_log_rotate_before_log_forwarding"); ok {
					logsSettingsPayload["perform-log-rotate-before-log-forwarding"] = v.(bool)
				}
				if v, ok := d.GetOkExists("logs_settings.0.reject_connections_when_free_disk_space_below_threshold"); ok {
					logsSettingsPayload["reject-connections-when-free-disk-space-below-threshold"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.reserve_for_packet_capture_metrics"); ok {
					logsSettingsPayload["reserve-for-packet-capture-metrics"] = v.(string)
				}
				if v, ok := d.GetOk("logs_settings.0.reserve_for_packet_capture_threshold"); ok {
					logsSettingsPayload["reserve-for-packet-capture-threshold"] = v.(int)
				}
				if v, ok := d.GetOkExists("logs_settings.0.rotate_log_by_file_size"); ok {
					logsSettingsPayload["rotate-log-by-file-size"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.rotate_log_file_size_threshold"); ok {
					logsSettingsPayload["rotate-log-file-size-threshold"] = v.(int)
				}
				if v, ok := d.GetOkExists("logs_settings.0.rotate_log_on_schedule"); ok {
					logsSettingsPayload["rotate-log-on-schedule"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.rotate_log_schedule_name"); ok {
					logsSettingsPayload["rotate-log-schedule-name"] = v.(string)
				}
				if v, ok := d.GetOkExists("logs_settings.0.stop_logging_when_free_disk_space_below"); ok {
					logsSettingsPayload["stop-logging-when-free-disk-space-below"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.stop_logging_when_free_disk_space_below_threshold"); ok {
					logsSettingsPayload["stop-logging-when-free-disk-space-below-threshold"] = v.(int)
				}
				if v, ok := d.GetOkExists("logs_settings.0.turn_on_qos_logging"); ok {
					logsSettingsPayload["turn-on-qos-logging"] = v.(bool)
				}
				if v, ok := d.GetOk("logs_settings.0.update_account_log_every"); ok {
					logsSettingsPayload["update-account-log-every"] = v.(int)
				}
				cluster["logs-settings"] = logsSettingsPayload
			}
		}

		if ok := d.HasChange("show_portals_certificate"); ok {
			if v, ok := d.GetOkExists("show_portals_certificate"); ok {
				cluster["show-portals-certificate"] = v.(bool)
			}
		}

		updateSimpleClusterRes, err := client.ApiCall("set-simple-cluster", cluster, client.GetSessionID(), true, client.IsProxyUsed())
		if err != nil {
			return fmt.Errorf("%s", err.Error())
		}
		if !updateSimpleClusterRes.Success {
			if updateSimpleClusterRes.ErrorMsg != "" {
				return fmt.Errorf("%s", updateSimpleClusterRes.ErrorMsg)
			}
			msg := createTaskFailMessage("set-simple-cluster", updateSimpleClusterRes.GetData())
			return fmt.Errorf("%s", msg)
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
			return fmt.Errorf("%s", deleteClusterRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
