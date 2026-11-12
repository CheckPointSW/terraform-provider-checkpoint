package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementSimpleGateway() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementSimpleGatewayRead,
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
			"advanced_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "N/A",
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"bypass_on_failure": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to bypass all requests (Fail-open) in case of internal system error.",
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
									"profile_value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The value inherited from the profile.",
									},
								},
							},
						},
						"site_categorization_allow_mode": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to 'background' in order to allowed requests until categorization is complete.",
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
									"profile_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value inherited from the profile.",
									},
								},
							},
						},
						"deny_untrusted_server_cert": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with untrusted server certificate.",
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
									"profile_value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The value inherited from the profile.",
									},
								},
							},
						},
						"deny_revoked_server_cert": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL).",
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
									"profile_value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The value inherited from the profile.",
									},
								},
							},
						},
						"deny_expired_server_cert": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Set to be true in order to drop traffic from servers with expired server certificate.",
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
									"profile_value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The value inherited from the profile.",
									},
								},
							},
						},
						"bypass_on_client_failure": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Bypass HTTPS inspection on client failure.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to override the value inherited from the profile.",
									},
									"profile_value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "The value inherited from the profile.",
									},
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to bypass on client failure.",
									},
								},
							},
						},
						"bypass_under_load": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Bypass HTTPS inspection under load.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"value": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to bypass under load.",
									},
								},
							},
						},
						"outbound_certificate": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Outbound HTTPS inspection certificate.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"override_profile": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to override the value inherited from the profile.",
									},
									"profile_value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value inherited from the profile.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Outbound certificate identified by the name or UID.",
									},
								},
							},
						},
						"deployment_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "HTTPS inspection deployment mode.",
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authentication Settings for Browser Based Authentication.",
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"portal_web_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal web settings.",
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
															"ip_address": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Optional IP address to be used for the portal URL.",
															},
														},
													},
												},
												"certificate_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal certificate settings.",
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
															"certificate": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate.",
															},
															"certificate_dn": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate distinguished name.",
															},
															"certificate_valid_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The date from which the certificate is valid.",
															},
															"certificate_valid_to": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The date until which the certificate is valid.",
															},
														},
													},
												},
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
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
												"certificate_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal certificate.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"certificate": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate.",
															},
															"certificate_dn": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate distinguished name.",
															},
															"certificate_valid_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The date from which the certificate is valid.",
															},
															"certificate_valid_to": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The date until which the certificate is valid.",
															},
														},
													},
												},
												"portal_web_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal web settings.",
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
															"ip_address": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Optional IP address to be used for the portal URL.",
															},
															"main_url": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The main URL for the portal.",
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"users_directories": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Users directories.",
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
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
												"certificate_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal certificate.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"certificate": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate.",
															},
															"certificate_dn": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate distinguished name.",
															},
															"certificate_valid_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The date from which the certificate is valid.",
															},
															"certificate_valid_to": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The date until which the certificate is valid.",
															},
														},
													},
												},
												"portal_web_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal web settings.",
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
															"ip_address": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Optional IP address to be used for the portal URL.",
															},
															"main_url": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The main URL for the portal.",
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
						"identity_web_api": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Identity Web API source enabled.",
						},
						"identity_web_api_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity Web API settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"authentication_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authentication settings for Identity Web API.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"users_directories": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Users directories.",
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
																Description: "LDAP AU objects identified by the name or UID.",
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
									"authorized_clients": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Authorized clients.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"client": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Host / Network Group Name or UID.",
												},
											},
										},
									},
									"client_access_permissions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Identity Web API accessibility settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal access settings.",
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
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"undefined": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to \"Undefined\".",
																		},
																		"dmz": {
																			Type:        schema.TypeBool,
																			Computed:    true,
																			Description: "Controls portal access settings for internal interfaces, whose topology is set to \"DMZ\".",
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
												"certificate_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal certificate.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"certificate": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate.",
															},
															"certificate_dn": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The certificate distinguished name.",
															},
															"certificate_valid_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The date from which the certificate is valid.",
															},
															"certificate_valid_to": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The date until which the certificate is valid.",
															},
														},
													},
												},
												"portal_web_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the portal web settings.",
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
															"ip_address": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Optional IP address to be used for the portal URL.",
															},
															"main_url": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The main URL for the portal.",
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
						"ad_query": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "AD Query source enabled.",
						},
						"collecting_identities": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This gateway collects identities.",
						},
						"identity_based_enforcement": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Configures this object as a PEP-only object - identity-based enforcement (PEP).",
						},
						"radius_accounting": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Radius Accounting source enabled.",
						},
						"terminal_servers": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Terminal Servers source enabled.",
						},
						"identity_sharing_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity sharing settings.",
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
									"cache_mode": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Identity cache mode.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"override_profile": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to override the value inherited from the profile.",
												},
												"value": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether the identity cache is enabled.",
												},
												"profile_value": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "The value inherited from the profile.",
												},
											},
										},
									},
									"cache_mode_duration": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Identity cache mode duration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"override_profile": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to override the value inherited from the profile.",
												},
												"value": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Identity cache duration in minutes.",
												},
												"profile_value": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The duration inherited from the profile, in minutes.",
												},
											},
										},
									},
									"receive_from_infinity_identity": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether to receive identities from Infinity Identity.",
									},
									"scaled_sharing": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether scaled identity sharing is enabled.",
									},
								},
							},
						},
						"proxy_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Identity-Awareness Proxy settings.",
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
				Type:        schema.TypeList,
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
						"apply_control_connections": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This option performs NAT on VPN control connections to and from this object.",
						},
					},
				},
			},
			"platform_portal_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Platform portal settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"portal_web_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Configuration of the portal web settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Optional: IP address for the web portal to use, if your DNS server fails to resolve the main portal URL. Note: If your DNS server resolves the main po...",
									},
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The certificate.",
									},
									"certificate_dn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The DN (Distinguished Name) of the certificate.",
									},
									"certificate_valid_from": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The date, from which the certificate is valid.",
									},
									"certificate_valid_to": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The certificate expiration date.",
									},
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
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "State of the web portal (enabled or disabled).",
						},
					},
				},
			},
			"proxy_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Proxy Server for Gateway.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_custom_proxy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Use custom proxy settings for this network object.",
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip_address": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Optional: IP address for the web portal to use, if your DNS server fails to resolve the main portal URL. Note: If your DNS server resolves the main po...",
									},
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
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The certificate.",
									},
									"certificate_dn": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The DN (Distinguished Name) of the certificate.",
									},
									"certificate_valid_from": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The date, from which the certificate is valid.",
									},
									"certificate_valid_to": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The certificate expiration date.",
									},
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
				Deprecated:  "use zero_phishing_settings.manual_fqdn instead - the API replaced zero-phishing-fqdn with zero-phishing-settings",
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
						"dynamic_ip": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable dynamic interface.",
						},
						"network_interface_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of network interface.",
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
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Anti spoofing settings",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"action": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "If packets will be rejected (the Prevent option) or whether the packets will be monitored (the Detect option).",
									},
									"exclude_packets": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Don't check packets from excluded network.",
									},
									"excluded_network_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Excluded network name.",
									},
									"excluded_network_uid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Excluded network UID.",
									},
									"spoof_tracking": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Spoof tracking.",
									},
								},
							},
						},
						"security_zone": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Security zone.",
						},
						"security_zone_settings": {
							Type:        schema.TypeList,
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
							Type:        schema.TypeList,
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
			"firewall": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Firewall blade enabled.",
			},
			"firewall_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Firewall settings",
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
			"icap_server": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "ICAP Server enabled.",
			},
			"ips": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Intrusion Prevention System blade enabled.",
			},
			"ips_settings": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Gateway IPS settings.",
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
					},
				},
			},
			"threat_emulation": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Threat Emulation blade enabled.",
			},
			"threat_extraction": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Threat Extraction blade enabled.",
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
				Description: "Gateway platform version.",
			},
			"hardware": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway platform hardware type.",
			},
			"one_time_password": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "SIC one time password.",
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
			"export_logs_to_servers": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Export logs to syslog/SIEM servers. NOTE:After you configure a Log Exporter, you must run Install Database.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"application_control_and_url_filtering_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Gateway Application Control and URL filtering settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"global_settings_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether to override global settings or not.",
						},
						"override_global_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "override global settings object.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"fail_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Fail mode - allow or block all requests.",
									},
									"website_categorization": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Website categorization object.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mode": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Website categorization mode.",
												},
												"custom_mode": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Custom mode object.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"social_networking_widgets": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Social networking widgets mode.",
															},
															"url_filtering": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "URL filtering mode.",
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
			"logs_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Logs settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_when_free_disk_space_below": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable alert when free disk space is below threshold.",
						},
						"free_disk_space_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Free disk space metrics.",
						},
						"delete_index_files_when_index_size_above_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete index files when index size above metrics",
						},
						"delete_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Delete when free disk space below metric.",
						},
						"stop_logging_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Stop logging when free disk space below metrics",
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
						"alert_when_free_disk_space_below_metrics": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert when free disk space below metrics.",
						},
						"distribute_logs_between_all_active_servers": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Distribute logs between all active servers.",
						},
						"include_tcp_state_information": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Include TCP state information. Relevant only when Firewall blade is enabled.",
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
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Gateway VPN settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication": {
							Type:        schema.TypeList,
							Computed:    true,
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
									"single_authentication_client": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Settings for clients that support only single authentication method.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Allow clients that support only single authentication method.",
												},
												"allow_multiple_authentication_clients": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Allow clients that support multiple authentication methods to connect.",
												},
												"display_name": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Display name for the authentication method.",
												},
												"method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Authentication method type.",
												},
												"secur_id": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "SecurID authentication settings, relevant only when method is \"secur-id\".",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Server object identified by the name or UID.",
															},
															"token_card_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Token card type.",
															},
														},
													},
												},
												"radius": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "RADIUS authentication settings, relevant only when method is \"radius\".",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"server": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Server object identified by the name or UID.",
															},
															"ask_user_password": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Ask user for password during authentication.",
															},
														},
													},
												},
												"personal_certificate": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Personal certificate authentication settings, relevant only when method is \"personal-certificate\".",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"fetch_username_from": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Fetch username from.",
															},
															"storage_type": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Certificate storage type.",
															},
															"source": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Certificate source field.",
															},
															"dn_part": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "DN part to extract.",
															},
															"dn_concurrence": {
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "DN part occurrence number.",
															},
														},
													},
												},
												"client_display_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Client display configuration settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"headline": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Display headline for authentication dialog.",
															},
															"username_label": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Label for username field.",
															},
															"password_label": {
																Type:        schema.TypeString,
																Computed:    true,
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
										Computed:    true,
										Description: "Override global dynamic ID settings.",
									},
									"dynamic_id_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Dynamic ID settings, relevant only when \"override-global-dynamic-id-settings\" is true.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"sms_provider_and_email_settings": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "SMS provider and email configuration.",
												},
												"sms_provider_credentials": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "SMS provider credentials configuration.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"username": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "SMS provider username.",
															},
															"api_id": {
																Type:        schema.TypeString,
																Computed:    true,
																Sensitive:   true,
																Description: "SMS provider API ID.",
															},
														},
													},
												},
												"advanced_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Advanced Dynamic ID configuration settings.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"dynamic_id_message": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Dynamic ID message displayed to users.",
															},
															"otp_settings": {
																Type:        schema.TypeList,
																Computed:    true,
																Description: "One Time Password configuration settings.",
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"length": {
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Length of one time password.",
																		},
																		"expiration": {
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "One time password expiration (in minutes).",
																		},
																		"max_attempts": {
																			Type:        schema.TypeInt,
																			Computed:    true,
																			Description: "Number of times users can attempt to enter the one time password before the entire authentication process restarts.",
																		},
																	},
																},
															},
															"enable_display_user_details": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Enable display of user details.",
															},
															"country_code": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Country code for SMS services.",
															},
															"user_details_retrieval": {
																Type:        schema.TypeString,
																Computed:    true,
																Description: "User details retrieval method.",
															},
														},
													},
												},
											},
										},
									},
									"send_machine_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Configure when to send machine certificate.",
									},
								},
							},
						},
						"link_selection": {
							Type:        schema.TypeList,
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
									"route_selection_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Outgoing route selection method when initiating a tunnel.",
									},
									"responding_traffic": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Responding traffic route selection method.",
									},
									"source_ip_selection": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Source IP address selection method for outgoing traffic.",
									},
									"selected_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Selected IP address. Must be set when \"source-ip-selection\" was selected to be \"manual\".",
									},
									"outgoing_link_tracking": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Outgoing link tracking method.",
									},
									"probing_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Probing settings configuration. Only available when \"ip-selection\" is \"use-probing-with-high-availability\" or \"use-probing-with-load-sharing\".",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"probed_interfaces": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Specifies whether to probe all addresses defined in the topology tab or specific addresses.",
												},
												"probed_interface_list": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "List of specific IP addresses to probe. Only relevant when \"probed-interfaces\" is set to \"specific\".",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"use_primary_address": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Whether to use a primary address for high availability probing.",
												},
												"primary_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Primary IP address to use. Must be one of the addresses from \"probed-interface-list\". Required when \"use-primary-address\" is true.",
												},
												"probing_method": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Probing method.",
												},
											},
										},
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
							Type:        schema.TypeList,
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
										Type:        schema.TypeList,
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
													Type:        schema.TypeList,
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
							Type:        schema.TypeList,
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
						"vpn_domain_exclude_external_ip_addresses": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Exclude the external IP addresses from the VPN domain of this Security Gateway.",
						},
						"advanced": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Advanced VPN settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"tunnel_sharing_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Tunnel sharing mode.",
									},
									"shutdown_on_gateway_restart": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Shutdown VPN tunnels on gateway restart.",
									},
									"enable_wire_mode": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable wire mode.",
									},
									"wire_mode_interfaces": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Wire mode interfaces.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"enable_wire_mode_log_traffic": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Log traffic in wire mode.",
									},
									"enable_nat_traversal": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable NAT traversal.",
									},
								},
							},
						},
						"exported_routes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Exported routes.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"internal_interfaces": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Export internal interfaces.",
									},
									"static_routes": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Export static routes.",
									},
									"custom_routes": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Export custom routes.",
									},
									"custom_routes_object": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Custom routes object identified by the name or UID.",
									},
								},
							},
						},
						"vpn_clients": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "VPN clients settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enable_endpoint_security_vpn": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable Endpoint Security VPN client.",
									},
									"enable_cp_mobile_for_windows": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable Check Point Mobile for Windows client.",
									},
									"enable_secu_remote": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable SecuRemote client.",
									},
									"enable_capsule_vpn_connect": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable Capsule VPN Connect client.",
									},
									"enable_ssl_network_extender": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Enable SSL Network Extender client.",
									},
									"gateway_authentication_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Gateway authentication certificate.",
									},
								},
							},
						},
						"enable_clientless_vpn": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enable clientless VPN.",
						},
						"clientless_vpn_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Clientless VPN settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"certificate_gateway_authentication": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Certificate gateway authentication.",
									},
									"client_authentication": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client authentication.",
									},
									"concurrent_servers_or_processes": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of concurrent servers or processes.",
									},
									"accept_only_3des": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Accept only 3DES.",
									},
								},
							},
						},
						"saml_portal_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "SAML portal settings.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Whether the SAML portal is enabled.",
									},
									"portal_web_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Configuration of the SAML portal web settings.",
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
												"ip_address": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Optional IP address to be used for the portal URL.",
												},
												"main_url": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The main URL for the portal.",
												},
											},
										},
									},
									"certificate_settings": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Configuration of the SAML portal certificate.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"certificate": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The certificate.",
												},
												"certificate_dn": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The certificate distinguished name.",
												},
												"certificate_valid_from": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The date from which the certificate is valid.",
												},
												"certificate_valid_to": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The date until which the certificate is valid.",
												},
											},
										},
									},
									"accessibility": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Configuration of the portal access settings.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"allow_access_from": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Allowed access to the SAML portal.",
												},
												"internal_access_settings": {
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Configuration of the additional portal access settings for internal interfaces only.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"undefined": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Controls portal access settings for internal interfaces, whose topology is set to \"Undefined\".",
															},
															"dmz": {
																Type:        schema.TypeBool,
																Computed:    true,
																Description: "Controls portal access settings for internal interfaces, whose topology is set to \"DMZ\".",
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
						"certificates": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "VPN certificates.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Certificate name.",
									},
									"stored_at": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Where the certificate is stored.",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Certificate status.",
									},
									"distinguished_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Certificate distinguished name.",
									},
									"base64_certificate": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The certificate encoded in Base64.",
									},
									"certificate_authority": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Certificate authority identified by the name or UID.",
									},
									"expiration_date": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Certificate expiration date.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"iso_8601": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Date and time represented in international ISO 8601 format.",
												},
												"posix": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.",
												},
											},
										},
									},
								},
							},
						},
						"interfaces": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "VPN link selection interfaces.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"interface_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the interface.",
									},
									"next_hop_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP address of the next hop.",
									},
									"static_nat_ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The NATed IPv4 address that hides the source IPv4 address of outgoing connections.",
									},
									"priority": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Priority of a \"Backup\" interface.",
									},
									"redundancy_mode": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Interface redundancy mode (Active/Backup).",
									},
									"ip_version": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The IP version of the interface's IP address.",
									},
								},
							},
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
			"accept_syslog_messages": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enables the gateway accept syslog messages, relevant only when Logging and Status blade is enabled.",
			},
			"anti_spam_and_email_security": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Anti-Spam & Email-Security blade enabled.",
			},
			"auto_generate_ip": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Use an automatically generated IP address for the Gateway object (applies only to Smart-1 Cloud).",
			},
			"auto_topology_custom_recalculation_time": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Auto topology custom recalculation time (seconds).",
			},
			"auto_topology_use_custom_recalculation_time": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Auto topology to use custom recalculation time instead of default.",
			},
			"autonomous_system_number": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The Autonomous System Number (ASN) for this Check Point Security Gateway.",
			},
			"data_loss_prevention": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Data Loss Prevention.",
			},
			"dns_server": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "DNS Server.",
			},
			"enable_log_indexing": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable log indexing, The Log Indexing uses more storage to provide fast log queries. Relevant only when Logging and Status blade is enabled.",
			},
			"externally_managed": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Externally Managed Check Point Gateway.",
			},
			"hardware_subtype": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Gateway type (relevant only for Spark gateways).",
			},
			"install_policy_without_push": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Specifies whether the policy is pushed to the gateway during policy installation, or whether the gateway should fetch the policy.",
			},
			"interfaces_topology_settings": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Topology setting for all interfaces on a Security Gateway. Default for Security Gateways that run Gaia OS: 'per interface'. Default for Quantum Spark ...",
			},
			"legacy_url_filtering": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Legacy URL Filtering enabled.",
			},
			"log_server": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Logging & Status.",
			},
			"mobile_access": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Mobile-Access blade enabled.",
			},
			"monitoring": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Monitoring blade enabled.",
			},
			"network_policy_management": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Management blade enabled.",
			},
			"policy_server": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Policy-Server blade enabled.",
			},
			"rtm_counters_report": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enables monitoring blades system counters report (e.g CPU Usage,Memory Usage).",
			},
			"rtm_traffic_report": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enables monitoring blades traffic report.",
			},
			"rtm_traffic_report_per_connection": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enables Monitoring blade traffic report per connection.",
			},
			"sic_message": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Secure Internal Communication message.",
			},
			"smart_event_intro_correlation_unit": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enables the gateway use SmartEvent intro correlation unit with one Security Gateway Software Blade, relevant only when Logging and Status blade is ena...",
			},
			"threat_prevention_mode": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The mode of Threat Prevention to use. When using Autonomous Threat Prevention, disabling the Threat Prevention blades is not allowed.",
			},
			"trust_method": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Trust method that was used for establishing communication.",
			},
			"workforce_ai": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Workforce AI Security blade enabled. Requires content awareness blade and version R82.20 or higher to be enabled.",
			},
			"communication_with_servers_behind_nat": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Gateway behind NAT communications settings with the server.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"override_profile": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether to override the Server (Check Point Host) object configuration.",
						},
						"value": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "according-to-topology: Use the original or translated IP address of the server based on the Topology of Security Gateway interfaces.<br>original-ip-on...",
						},
					},
				},
			},
			"smb_logs_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Logs settings that apply to Quantum Spark Appliances that run Gaia Embedded OS.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"alert_when_queue_is_full": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Alert when queue is full enabled.",
						},
						"alert_when_queue_is_full_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Alert when queue is full type.",
						},
						"detect_new_citrix_ica_application_names": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Detect new citrix ica application names enabled.",
						},
						"stop_logging_when_queue_reaches_maximal_capacity": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Stop logging when queue reaches maximal capacity enabled.",
						},
						"stop_logging_when_queue_reaches_maximal_capacity_threshold": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Stop logging when queue reaches maximal capacity threshold.",
						},
						"turn_on_qos_logging": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Turn on qos logging enabled.",
						},
						"update_account_log_every": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Update account log in every amount of seconds.",
						},
					},
				},
			},
			"zero_phishing_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Fqdn settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"gateway_fqdn_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Manual Fqdn.",
						},
						"manual_fqdn": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Zero Phishing gateway FQDN.",
						},
					},
				},
			},
			"fetch_policy_scheduler": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Fetch policy functionality settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if the Security Gateway will fetch policy according to a schedule (true) or manually (false).",
						},
						"schedule": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Scheduled Event for fetching policy by.<br><font color='red'>Will be applied only</font> when the `enabled` field is set to true.<br>When not defined ...",
						},
					},
				},
			},
			"trust_details": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Details for trusted communication.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"authentication_token": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Authentication token to use on the Gateway side to establish the communication between the Gateway and the Management Server (applies only to Smart-1 ...",
						},
						"cloud_communication_details": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Details about the communication status with cloud (applies only to Smart-1 Cloud).",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "IP address used for communication between the Gateway and the Management Server (used when 'auto-generate-ip=true').",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of communication between the Gateway and the Management Server.",
									},
								},
							},
						},
						"gateway_mac_address": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Use the Security Gateway MAC address, relevant for the gateway_mac_address identification-method.",
						},
						"identification_method": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "How to identify the gateway (relevant for DAIP gateways only).",
						},
						"status": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the trusted communication with the Security Gateway.",
						},
						"token_expiration_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Details about the communication status with cloud (applies only to Smart-1 Cloud).",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementSimpleGatewayRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showGatewayRes, err := client.ApiCall("show-simple-gateway", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showGatewayRes.Success {
		return fmt.Errorf("%s", showGatewayRes.ErrorMsg)
	}

	gateway := showGatewayRes.GetData()

	log.Println("Read Simple Gateway - Show JSON = ", gateway)

	if v := gateway["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

	if v := gateway["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := gateway["ipv4-address"]; v != nil {
		_ = d.Set("ipv4_address", v)
	}

	if v := gateway["ipv6-address"]; v != nil {
		_ = d.Set("ipv6_address", v)
	}

	if gateway["advanced-settings"] != nil {

		advancedSettingsMap, ok := gateway["advanced-settings"].(map[string]interface{})

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

	if v := gateway["enable-https-inspection"]; v != nil {
		_ = d.Set("enable_https_inspection", v)
	}

	if gateway["fetch-policy"] != nil {
		fetchPolicyJson, ok := gateway["fetch-policy"].([]interface{})
		if ok {
			fetchPolicyIds := make([]string, 0)
			if len(fetchPolicyJson) > 0 {
				for _, fetchPolicy := range fetchPolicyJson {
					fetchPolicyIdentifier := ""
					if fetchPolicyObject, ok := fetchPolicy.(map[string]interface{}); ok {
						fetchPolicyIdentifier = fetchPolicyObject["name"].(string)
					} else {
						fetchPolicyIdentifier = fetchPolicy.(string)
					}

					if fetchPolicyIdentifier != "" {
						fetchPolicyIds = append(fetchPolicyIds, fetchPolicyIdentifier)
					}
				}
			}
			_ = d.Set("fetch_policy", fetchPolicyIds)
		}
	} else {
		_ = d.Set("fetch_policy", nil)
	}

	if v := gateway["hit-count"]; v != nil {
		_ = d.Set("hit_count", v)
	}

	if gateway["https-inspection"] != nil {

		httpsInspectionMap, ok := gateway["https-inspection"].(map[string]interface{})

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
					if v, _ := bypassOnFailureMap["profile-value"]; v != nil {
						bypassOnFailureMapToReturn["profile_value"] = v
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
					if v, _ := siteCategorizationAllowModeMap["profile-value"]; v != nil {
						siteCategorizationAllowModeMapToReturn["profile_value"] = v
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
					if v, _ := denyUntrustedServerCertMap["profile-value"]; v != nil {
						denyUntrustedServerCertMapToReturn["profile_value"] = v
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
					if v, _ := denyRevokedServerCertMap["profile-value"]; v != nil {
						denyRevokedServerCertMapToReturn["profile_value"] = v
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
					if v, _ := denyExpiredServerCertMap["profile-value"]; v != nil {
						denyExpiredServerCertMapToReturn["profile_value"] = v
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
					if v, _ := bypassOnClientFailureMap["profile-value"]; v != nil {
						bypassOnClientFailureMapToReturn["profile_value"] = v
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
					if v, _ := outboundCertificateMap["profile-value"]; v != nil {
						outboundCertificateMapToReturn["profile_value"] = v.(map[string]interface{})["name"]
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

	if v := gateway["identity-awareness"]; v != nil {
		_ = d.Set("identity_awareness", v)
	}

	if gateway["identity-awareness-settings"] != nil {

		identityAwarenessSettingsMap, ok := gateway["identity-awareness-settings"].(map[string]interface{})

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
								if v, _ := internalAccessSettingsMap["undefined"]; v != nil {
									internalAccessSettingsMapToReturn["undefined"] = v
								}
								if v, _ := internalAccessSettingsMap["dmz"]; v != nil {
									internalAccessSettingsMapToReturn["dmz"] = v
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
							if v, _ := certificateSettingsMap["certificate"]; v != nil {
								certificateSettingsMapToReturn["certificate"] = v
							}
							if v, _ := certificateSettingsMap["certificate-dn"]; v != nil {
								certificateSettingsMapToReturn["certificate_dn"] = v
							}
							if v, _ := certificateSettingsMap["certificate-valid-from"]; v != nil {
								certificateSettingsMapToReturn["certificate_valid_from"] = v
							}
							if v, _ := certificateSettingsMap["certificate-valid-to"]; v != nil {
								certificateSettingsMapToReturn["certificate_valid_to"] = v
							}
							browserBasedAuthenticationPortalSettingsMapToReturn["certificate_settings"] = []interface{}{certificateSettingsMapToReturn}
						}
						if v, _ := browserBasedAuthenticationPortalSettingsMap["portal-web-settings"]; v != nil {
							portalWebSettingsMap := v.(map[string]interface{})
							portalWebSettingsMapToReturn := make(map[string]interface{})
							if v, _ := portalWebSettingsMap["aliases"]; v != nil {
								portalWebSettingsMapToReturn["aliases"] = v
							}
							if v, _ := portalWebSettingsMap["ip-address"]; v != nil {
								portalWebSettingsMapToReturn["ip_address"] = v
							}
							if v, _ := portalWebSettingsMap["main-url"]; v != nil {
								portalWebSettingsMapToReturn["main_url"] = v
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
						identityAgentSettingsMapToReturn["authentication_settings"] = v
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
								if v, _ := internalAccessSettingsMap["undefined"]; v != nil {
									internalAccessSettingsMapToReturn["undefined"] = v
								}
								if v, _ := internalAccessSettingsMap["dmz"]; v != nil {
									internalAccessSettingsMapToReturn["dmz"] = v
								}
								if v, _ := internalAccessSettingsMap["vpn"]; v != nil {
									internalAccessSettingsMapToReturn["vpn"] = v
								}
								accessibilityMapToReturn["internal_access_settings"] = []interface{}{internalAccessSettingsMapToReturn}
							}
							identityAgentPortalSettingsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
						}
						if v, _ := identityAgentPortalSettingsMap["certificate-settings"]; v != nil {
							certificateSettingsMap := v.(map[string]interface{})
							certificateSettingsMapToReturn := make(map[string]interface{})
							if v, _ := certificateSettingsMap["certificate"]; v != nil {
								certificateSettingsMapToReturn["certificate"] = v
							}
							if v, _ := certificateSettingsMap["certificate-dn"]; v != nil {
								certificateSettingsMapToReturn["certificate_dn"] = v
							}
							if v, _ := certificateSettingsMap["certificate-valid-from"]; v != nil {
								certificateSettingsMapToReturn["certificate_valid_from"] = v
							}
							if v, _ := certificateSettingsMap["certificate-valid-to"]; v != nil {
								certificateSettingsMapToReturn["certificate_valid_to"] = v
							}
							identityAgentPortalSettingsMapToReturn["certificate_settings"] = []interface{}{certificateSettingsMapToReturn}
						}
						if v, _ := identityAgentPortalSettingsMap["portal-web-settings"]; v != nil {
							portalWebSettingsMap := v.(map[string]interface{})
							portalWebSettingsMapToReturn := make(map[string]interface{})
							if v, _ := portalWebSettingsMap["aliases"]; v != nil {
								portalWebSettingsMapToReturn["aliases"] = v
							}
							if v, _ := portalWebSettingsMap["ip-address"]; v != nil {
								portalWebSettingsMapToReturn["ip_address"] = v
							}
							if v, _ := portalWebSettingsMap["main-url"]; v != nil {
								portalWebSettingsMapToReturn["main_url"] = v
							}
							identityAgentPortalSettingsMapToReturn["portal_web_settings"] = []interface{}{portalWebSettingsMapToReturn}
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
						identityCollectorSettingsMapToReturn["authentication_settings"] = v
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
								if v, _ := internalAccessSettingsMap["undefined"]; v != nil {
									internalAccessSettingsMapToReturn["undefined"] = v
								}
								if v, _ := internalAccessSettingsMap["dmz"]; v != nil {
									internalAccessSettingsMapToReturn["dmz"] = v
								}
								if v, _ := internalAccessSettingsMap["vpn"]; v != nil {
									internalAccessSettingsMapToReturn["vpn"] = v
								}
								accessibilityMapToReturn["internal_access_settings"] = []interface{}{internalAccessSettingsMapToReturn}
							}
							clientAccessPermissionsMapToReturn["accessibility"] = []interface{}{accessibilityMapToReturn}
						}
						if v, _ := clientAccessPermissionsMap["certificate-settings"]; v != nil {
							certificateSettingsMap := v.(map[string]interface{})
							certificateSettingsMapToReturn := make(map[string]interface{})
							if v, _ := certificateSettingsMap["certificate"]; v != nil {
								certificateSettingsMapToReturn["certificate"] = v
							}
							if v, _ := certificateSettingsMap["certificate-dn"]; v != nil {
								certificateSettingsMapToReturn["certificate_dn"] = v
							}
							if v, _ := certificateSettingsMap["certificate-valid-from"]; v != nil {
								certificateSettingsMapToReturn["certificate_valid_from"] = v
							}
							if v, _ := certificateSettingsMap["certificate-valid-to"]; v != nil {
								certificateSettingsMapToReturn["certificate_valid_to"] = v
							}
							clientAccessPermissionsMapToReturn["certificate_settings"] = []interface{}{certificateSettingsMapToReturn}
						}
						if v, _ := clientAccessPermissionsMap["portal-web-settings"]; v != nil {
							portalWebSettingsMap := v.(map[string]interface{})
							portalWebSettingsMapToReturn := make(map[string]interface{})
							if v, _ := portalWebSettingsMap["aliases"]; v != nil {
								portalWebSettingsMapToReturn["aliases"] = v
							}
							if v, _ := portalWebSettingsMap["ip-address"]; v != nil {
								portalWebSettingsMapToReturn["ip_address"] = v
							}
							if v, _ := portalWebSettingsMap["main-url"]; v != nil {
								portalWebSettingsMapToReturn["main_url"] = v
							}
							clientAccessPermissionsMapToReturn["portal_web_settings"] = []interface{}{portalWebSettingsMapToReturn}
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
						if v, _ := cacheModeMap["profile-value"]; v != nil {
							cacheModeMapToReturn["profile_value"] = v
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
						if v, _ := cacheModeDurationMap["profile-value"]; v != nil {
							cacheModeDurationMapToReturn["profile_value"] = v
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
			if v := identityAwarenessSettingsMap["identity-web-api"]; v != nil {
				identityAwarenessSettingsMapToReturn["identity_web_api"] = v
			}
			if v := identityAwarenessSettingsMap["identity-web-api-settings"]; v != nil {
				identityWebApiSettingsJson := v.(map[string]interface{})
				identityWebApiSettingsState := make(map[string]interface{})
				if v := identityWebApiSettingsJson["authentication-settings"]; v != nil {
					authenticationSettingsJson := v.(map[string]interface{})
					authenticationSettingsState := make(map[string]interface{})
					if v := authenticationSettingsJson["users-directories"]; v != nil {
						usersDirectoriesJson := v.(map[string]interface{})
						usersDirectoriesState := make(map[string]interface{})
						if v := usersDirectoriesJson["external-user-profile"]; v != nil {
							usersDirectoriesState["external_user_profile"] = v
						}
						if v := usersDirectoriesJson["internal-users"]; v != nil {
							usersDirectoriesState["internal_users"] = v
						}
						if v := usersDirectoriesJson["users-from-external-directories"]; v != nil {
							usersDirectoriesState["users_from_external_directories"] = v
						}
						if v := usersDirectoriesJson["specific"]; v != nil {
							usersDirectoriesState["specific"] = v
						}
						authenticationSettingsState["users_directories"] = []interface{}{usersDirectoriesState}
					}
					identityWebApiSettingsState["authentication_settings"] = []interface{}{authenticationSettingsState}
				}
				if v := identityWebApiSettingsJson["authorized-clients"]; v != nil {
					authorizedClientsList := v.([]interface{})
					if len(authorizedClientsList) > 0 {
						var authorizedClientsListState []map[string]interface{}
						for i := range authorizedClientsList {
							authorizedClientsItemJson := authorizedClientsList[i].(map[string]interface{})
							authorizedClientsItemState := make(map[string]interface{})
							if v := authorizedClientsItemJson["client"]; v != nil {
								authorizedClientsItemState["client"] = v
							}
							authorizedClientsListState = append(authorizedClientsListState, authorizedClientsItemState)
						}
						identityWebApiSettingsState["authorized_clients"] = authorizedClientsListState
					}
				}
				if v := identityWebApiSettingsJson["client-access-permissions"]; v != nil {
					clientAccessPermissionsJson := v.(map[string]interface{})
					clientAccessPermissionsState := make(map[string]interface{})
					if v := clientAccessPermissionsJson["accessibility"]; v != nil {
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
						clientAccessPermissionsState["accessibility"] = []interface{}{accessibilityState}
					}
					if v := clientAccessPermissionsJson["certificate-settings"]; v != nil {
						certificateSettingsJson := v.(map[string]interface{})
						certificateSettingsState := make(map[string]interface{})
						if v := certificateSettingsJson["certificate"]; v != nil {
							certificateSettingsState["certificate"] = v
						}
						if v := certificateSettingsJson["certificate-dn"]; v != nil {
							certificateSettingsState["certificate_dn"] = v
						}
						if v := certificateSettingsJson["certificate-valid-from"]; v != nil {
							certificateSettingsState["certificate_valid_from"] = v
						}
						if v := certificateSettingsJson["certificate-valid-to"]; v != nil {
							certificateSettingsState["certificate_valid_to"] = v
						}
						clientAccessPermissionsState["certificate_settings"] = []interface{}{certificateSettingsState}
					}
					if v := clientAccessPermissionsJson["portal-web-settings"]; v != nil {
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
						clientAccessPermissionsState["portal_web_settings"] = []interface{}{portalWebSettingsState}
					}
					identityWebApiSettingsState["client_access_permissions"] = []interface{}{clientAccessPermissionsState}
				}
				identityAwarenessSettingsMapToReturn["identity_web_api_settings"] = []interface{}{identityWebApiSettingsState}
			}
			if v := identityAwarenessSettingsMap["ad-query"]; v != nil {
				identityAwarenessSettingsMapToReturn["ad_query"] = v
			}
			if v := identityAwarenessSettingsMap["collecting-identities"]; v != nil {
				identityAwarenessSettingsMapToReturn["collecting_identities"] = v
			}
			if v := identityAwarenessSettingsMap["identity-based-enforcement"]; v != nil {
				identityAwarenessSettingsMapToReturn["identity_based_enforcement"] = v
			}
			if v := identityAwarenessSettingsMap["radius-accounting"]; v != nil {
				identityAwarenessSettingsMapToReturn["radius_accounting"] = v
			}
			if v := identityAwarenessSettingsMap["terminal-servers"]; v != nil {
				identityAwarenessSettingsMapToReturn["terminal_servers"] = v
			}
			_ = d.Set("identity_awareness_settings", []interface{}{identityAwarenessSettingsMapToReturn})

		}
	} else {
		_ = d.Set("identity_awareness_settings", nil)
	}

	if v := gateway["ips-update-policy"]; v != nil {
		_ = d.Set("ips_update_policy", v)
	}

	if v := gateway["nat-hide-internal-interfaces"]; v != nil {
		_ = d.Set("nat_hide_internal_interfaces", v)
	}

	if gateway["nat-settings"] != nil {

		natSettingsMap := gateway["nat-settings"].(map[string]interface{})

		natSettingsMapToReturn := make(map[string]interface{})

		if v := natSettingsMap["auto-rule"]; v != nil {
			natSettingsMapToReturn["auto_rule"] = v
		}
		if v := natSettingsMap["hide-behind"]; v != nil {
			natSettingsMapToReturn["hide_behind"] = v
		}
		if v := natSettingsMap["install-on"]; v != nil {
			natSettingsMapToReturn["install_on"] = v
		}
		if v := natSettingsMap["ipv4-address"]; v != nil {
			natSettingsMapToReturn["ipv4_address"] = v
		}
		if v := natSettingsMap["ipv6-address"]; v != nil {
			natSettingsMapToReturn["ipv6_address"] = v
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

	if gateway["platform-portal-settings"] != nil {

		platformPortalSettingsMap, ok := gateway["platform-portal-settings"].(map[string]interface{})

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
					if v, _ := certificateSettingsMap["certificate"]; v != nil {
						certificateSettingsMapToReturn["certificate"] = v
					}
					if v, _ := certificateSettingsMap["certificate-dn"]; v != nil {
						certificateSettingsMapToReturn["certificate_dn"] = v
					}
					if v, _ := certificateSettingsMap["certificate-valid-from"]; v != nil {
						certificateSettingsMapToReturn["certificate_valid_from"] = v
					}
					if v, _ := certificateSettingsMap["certificate-valid-to"]; v != nil {
						certificateSettingsMapToReturn["certificate_valid_to"] = v
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
			if v := platformPortalSettingsMap["enabled"]; v != nil {
				platformPortalSettingsMapToReturn["enabled"] = v
			}
			_ = d.Set("platform_portal_settings", []interface{}{platformPortalSettingsMapToReturn})

		}
	} else {
		_ = d.Set("platform_portal_settings", nil)
	}

	if gateway["proxy-settings"] != nil {

		proxySettingsMap := gateway["proxy-settings"].(map[string]interface{})

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

	if v := gateway["qos"]; v != nil {
		_ = d.Set("qos", v)
	}

	if gateway["usercheck-portal-settings"] != nil {

		usercheckPortalSettingsMap, ok := gateway["usercheck-portal-settings"].(map[string]interface{})

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
					if v, _ := certificateSettingsMap["certificate"]; v != nil {
						certificateSettingsMapToReturn["certificate"] = v
					}
					if v, _ := certificateSettingsMap["certificate-dn"]; v != nil {
						certificateSettingsMapToReturn["certificate_dn"] = v
					}
					if v, _ := certificateSettingsMap["certificate-valid-from"]; v != nil {
						certificateSettingsMapToReturn["certificate_valid_from"] = v
					}
					if v, _ := certificateSettingsMap["certificate-valid-to"]; v != nil {
						certificateSettingsMapToReturn["certificate_valid_to"] = v
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

	if v := gateway["zero-phishing"]; v != nil {
		_ = d.Set("zero_phishing", v)
	}

	if v := gateway["zero-phishing-fqdn"]; v != nil {
		_ = d.Set("zero_phishing_fqdn", v)
	}

	if v := gateway["interfaces"]; v != nil {
		interfacesList := v.([]interface{})
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
				if v, _ := interfaceJson["anti-spoofing"]; v != nil {
					interfaceState["anti_spoofing"] = v
				}
				if v, _ := interfaceJson["anti-spoofing-settings"]; v != nil {
					antiSpoofingSettingsJson := v.(map[string]interface{})
					antiSpoofingSettingsState := make(map[string]interface{})
					if v, _ := antiSpoofingSettingsJson["action"]; v != nil {
						antiSpoofingSettingsState["action"] = v
					}
					if v := antiSpoofingSettingsJson["exclude-packets"]; v != nil {
						antiSpoofingSettingsState["exclude_packets"] = v
					}
					if v := antiSpoofingSettingsJson["excluded-network-name"]; v != nil {
						antiSpoofingSettingsState["excluded_network_name"] = v
					}
					if v := antiSpoofingSettingsJson["excluded-network-uid"]; v != nil {
						antiSpoofingSettingsState["excluded_network_uid"] = v
					}
					if v := antiSpoofingSettingsJson["spoof-tracking"]; v != nil {
						antiSpoofingSettingsState["spoof_tracking"] = v
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

	if v := gateway["anti-bot"]; v != nil {
		_ = d.Set("anti_bot", v)
	}

	if v := gateway["anti-virus"]; v != nil {
		_ = d.Set("anti_virus", v)
	}

	if v := gateway["application-control"]; v != nil {
		_ = d.Set("application_control", v)
	}

	if gateway["application-control-and-url-filtering-settings"] != nil {

		applicationControlSettingsMap := gateway["application-control-and-url-filtering-settings"].(map[string]interface{})

		applicationControlSettingsMapToReturn := make(map[string]interface{})

		if v, _ := applicationControlSettingsMap["global-settings-mode"]; v != nil {
			applicationControlSettingsMapToReturn["global_settings_mode"] = v
		}

		if overrideGlobal, _ := applicationControlSettingsMap["override-global-settings"]; overrideGlobal != nil {
			overrideGlobalSettingsMap := overrideGlobal.(map[string]interface{})
			overrideGlobalMapToReturn := make(map[string]interface{})
			if v, _ := overrideGlobalSettingsMap["fail-mode"]; v != nil {
				overrideGlobalMapToReturn["fail_mode"] = v
			}

			if websiteCategorization, _ := overrideGlobalSettingsMap["website-categorization"]; websiteCategorization != nil {
				websiteCategorizationMap := websiteCategorization.(map[string]interface{})
				websiteCategorizationMapToReturn := make(map[string]interface{})

				if v, _ := websiteCategorizationMap["mode"]; v != nil {
					websiteCategorizationMapToReturn["mode"] = v
				}

				if customMode, _ := websiteCategorizationMap["custom-mode"]; customMode != nil {
					customModeMap := customMode.(map[string]interface{})
					customModeMapToReturn := make(map[string]interface{})
					if v, _ := customModeMap["social-networking-widgets"]; v != nil {
						customModeMapToReturn["social_networking_widgets"] = v
					}

					if v, _ := customModeMap["url-filtering"]; v != nil {
						customModeMapToReturn["url_filtering"] = v
					}
					websiteCategorizationMapToReturn["custom_mode"] = []interface{}{customModeMapToReturn}
				}
				overrideGlobalMapToReturn["website_categorization"] = []interface{}{websiteCategorizationMapToReturn}
			}
			applicationControlSettingsMapToReturn["override_global_settings"] = []interface{}{overrideGlobalMapToReturn}
		}

		_ = d.Set("application_control_and_url_filtering_settings", []interface{}{applicationControlSettingsMapToReturn})
	} else {
		_ = d.Set("application_control_and_url_filtering_settings", []interface{}{})
	}

	if v := gateway["content-awareness"]; v != nil {
		_ = d.Set("content_awareness", v)
	}

	if v := gateway["dynamic-ip"]; v != nil {
		_ = d.Set("dynamic_ip", v)
	}

	if v := gateway["firewall"]; v != nil {
		_ = d.Set("firewall", v)
	}

	if v := gateway["icap-server"]; v != nil {
		_ = d.Set("icap_server", v)
	}

	if v := gateway["ips"]; v != nil {
		_ = d.Set("ips", v)
	}

	if gateway["ips-settings"] != nil {

		ipsSettingsMap, ok := gateway["ips-settings"].(map[string]interface{})

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
			_ = d.Set("ips_settings", []interface{}{ipsSettingsMapToReturn})

		}
	} else {
		_ = d.Set("ips_settings", nil)
	}

	if v := gateway["threat-emulation"]; v != nil {
		_ = d.Set("threat_emulation", v)
	}

	if v := gateway["threat-extraction"]; v != nil {
		_ = d.Set("threat_extraction", v)
	}

	if v := gateway["url-filtering"]; v != nil {
		_ = d.Set("url_filtering", v)
	}

	if v := gateway["vpn"]; v != nil {
		_ = d.Set("vpn", v)
	}

	if v := gateway["os-name"]; v != nil {
		_ = d.Set("os_name", v)
	}

	if v := gateway["version"]; v != nil {
		_ = d.Set("version", v)
	}

	if v := gateway["hardware"]; v != nil {
		_ = d.Set("hardware", v)
	}

	if v := gateway["sic-name"]; v != nil {
		_ = d.Set("sic_name", v)
	}

	if v := gateway["sic-state"]; v != nil {
		_ = d.Set("sic_state", v)
	}

	if v := gateway["save-logs-locally"]; v != nil {
		_ = d.Set("save_logs_locally", v)
	}

	if v := gateway["send_alerts_to_server"]; v != nil {
		_ = d.Set("send_alerts_to_server", v)
	} else {
		_ = d.Set("send_alerts_to_server", nil)
	}

	if v := gateway["send-logs-to-backup-server"]; v != nil {
		_ = d.Set("send_logs_to_backup_server", v)
	} else {
		_ = d.Set("send_logs_to_backup_server", nil)
	}

	if v := gateway["send-logs-to-server"]; v != nil {
		_ = d.Set("send_logs_to_server", v)
	} else {
		_ = d.Set("send_logs_to_server", nil)
	}

	if v := gateway["export-logs-to-servers"]; v != nil {
		_ = d.Set("export_logs_to_servers", v)
	} else {
		_ = d.Set("export_logs_to_servers", nil)
	}

	if gateway["logs-settings"] != nil {

		logsSettingsMap := gateway["logs-settings"].(map[string]interface{})

		logsSettingsMapToReturn := make(map[string]interface{})

		if v := logsSettingsMap["alert-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["free_disk_space_metrics"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["alert-when-free-disk-space-below-type"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_type"] = v
		}
		if v := logsSettingsMap["before-delete-keep-logs-from-the-last-days"]; v != nil {
			logsSettingsMapToReturn["before_delete_keep_logs_from_the_last_days"] = v
		}
		if v := logsSettingsMap["before-delete-keep-logs-from-the-last-days-threshold"]; v != nil {
			logsSettingsMapToReturn["before_delete_keep_logs_from_the_last_days_threshold"] = v
		}
		if v := logsSettingsMap["before-delete-run-script"]; v != nil {
			logsSettingsMapToReturn["before_delete_run_script"] = v
		}
		if v := logsSettingsMap["before-delete-run-script-command"]; v != nil {
			logsSettingsMapToReturn["before_delete_run_script_command"] = v
		}
		if v := logsSettingsMap["delete-index-files-older-than-days"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_older_than_days"] = v
		}
		if v := logsSettingsMap["delete-index-files-older-than-days-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_older_than_days_threshold"] = v
		}
		if v := logsSettingsMap["delete-index-files-when-index-size-above"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_when_index_size_above"] = v
		}
		if v := logsSettingsMap["delete-index-files-when-index-size-above-metrics"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_when_index_size_above_metrics"] = v
		}
		if v := logsSettingsMap["delete-index-files-when-index-size-above-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_index_files_when_index_size_above_threshold"] = v
		}
		if v := logsSettingsMap["delete-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below"] = v
		}
		if v := logsSettingsMap["delete-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsMap["delete-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["delete_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["detect-new-citrix-ica-application-names"]; v != nil {
			logsSettingsMapToReturn["detect_new_citrix_ica_application_names"] = v
		}
		if v := logsSettingsMap["forward-logs-to-log-server"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server"] = v
		}
		if v := logsSettingsMap["forward-logs-to-log-server-name"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server_name"] = v
		}
		if v := logsSettingsMap["forward-logs-to-log-server-schedule-name"]; v != nil {
			logsSettingsMapToReturn["forward_logs_to_log_server_schedule_name"] = v
		}
		if v := logsSettingsMap["perform-log-rotate-before-log-forwarding"]; v != nil {
			logsSettingsMapToReturn["perform_log_rotate_before_log_forwarding"] = v
		}
		if v := logsSettingsMap["reject-connections-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["reject_connections_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["reserve-for-packet-capture-metrics"]; v != nil {
			logsSettingsMapToReturn["reserve_for_packet_capture_metrics"] = v
		}
		if v := logsSettingsMap["reserve-for-packet-capture-threshold"]; v != nil {
			logsSettingsMapToReturn["reserve_for_packet_capture_threshold"] = v
		}
		if v := logsSettingsMap["rotate-log-by-file-size"]; v != nil {
			logsSettingsMapToReturn["rotate_log_by_file_size"] = v
		}
		if v := logsSettingsMap["rotate-log-file-size-threshold"]; v != nil {
			logsSettingsMapToReturn["rotate_log_file_size_threshold"] = v
		}
		if v := logsSettingsMap["rotate-log-on-schedule"]; v != nil {
			logsSettingsMapToReturn["rotate_log_on_schedule"] = v
		}
		if v := logsSettingsMap["rotate-log-schedule-name"]; v != nil {
			logsSettingsMapToReturn["rotate_log_schedule_name"] = v
		}
		if v := logsSettingsMap["stop-logging-when-free-disk-space-below"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below"] = v
		}
		if v := logsSettingsMap["stop-logging-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsMap["stop-logging-when-free-disk-space-below-threshold"]; v != nil {
			logsSettingsMapToReturn["stop_logging_when_free_disk_space_below_threshold"] = v
		}
		if v := logsSettingsMap["turn-on-qos-logging"]; v != nil {
			logsSettingsMapToReturn["turn_on_qos_logging"] = v
		}
		if v := logsSettingsMap["update-account-log-every"]; v != nil {
			logsSettingsMapToReturn["update_account_log_every"] = v
		}

		if v := logsSettingsMap["alert-when-free-disk-space-below-metrics"]; v != nil {
			logsSettingsMapToReturn["alert_when_free_disk_space_below_metrics"] = v
		}
		if v := logsSettingsMap["distribute-logs-between-all-active-servers"]; v != nil {
			logsSettingsMapToReturn["distribute_logs_between_all_active_servers"] = v
		}
		if v := logsSettingsMap["include-tcp-state-information"]; v != nil {
			logsSettingsMapToReturn["include_tcp_state_information"] = v
		}
		_ = d.Set("logs_settings", []interface{}{logsSettingsMapToReturn})

	} else {
		_ = d.Set("logs_settings", nil)
	}

	if gateway["firewall-settings"] != nil {

		firewallSettingsMap := gateway["firewall-settings"].(map[string]interface{})

		firewallSettingsMapToReturn := make(map[string]interface{})

		if v := firewallSettingsMap["auto-calculate-connections-hash-table-size-and-memory-pool"]; v != nil {
			firewallSettingsMapToReturn["auto_calculate_connections_hash_table_size_and_memory_pool"] = v
		}
		if v := firewallSettingsMap["auto-maximum-limit-for-concurrent-connections"]; v != nil {
			firewallSettingsMapToReturn["auto_maximum_limit_for_concurrent_connections"] = v
		}
		if v := firewallSettingsMap["connections-hash-size"]; v != nil {
			firewallSettingsMapToReturn["connections_hash_size"] = v
		}
		if v := firewallSettingsMap["maximum-limit-for-concurrent-connections"]; v != nil {
			firewallSettingsMapToReturn["maximum_limit_for_concurrent_connections"] = v
		}
		if v := firewallSettingsMap["maximum-memory-pool-size"]; v != nil {
			firewallSettingsMapToReturn["maximum_memory_pool_size"] = v
		}
		if v := firewallSettingsMap["memory-pool-size"]; v != nil {
			firewallSettingsMapToReturn["memory_pool_size"] = v
		}

		_ = d.Set("firewall_settings", []interface{}{firewallSettingsMapToReturn})

	} else {
		_ = d.Set("firewall_settings", nil)
	}

	if v := gateway["vpn-settings"]; v != nil {
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
			if v := samlPortalSettingsJson["enabled"]; v != nil {
				samlPortalSettingsState["enabled"] = v
			}
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
			if v := samlPortalSettingsJson["certificate-settings"]; v != nil {
				certificateSettingsJson := v.(map[string]interface{})
				certificateSettingsState := make(map[string]interface{})
				if v := certificateSettingsJson["certificate"]; v != nil {
					certificateSettingsState["certificate"] = v
				}
				if v := certificateSettingsJson["certificate-dn"]; v != nil {
					certificateSettingsState["certificate_dn"] = v
				}
				if v := certificateSettingsJson["certificate-valid-from"]; v != nil {
					certificateSettingsState["certificate_valid_from"] = v
				}
				if v := certificateSettingsJson["certificate-valid-to"]; v != nil {
					certificateSettingsState["certificate_valid_to"] = v
				}
				samlPortalSettingsState["certificate_settings"] = []interface{}{certificateSettingsState}
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

		if v := vpnSettingsJson["certificates"]; v != nil {
			certificatesList := v.([]interface{})
			if len(certificatesList) > 0 {
				var certificatesListState []map[string]interface{}
				for i := range certificatesList {
					certificateJson := certificatesList[i].(map[string]interface{})
					certificateState := make(map[string]interface{})
					if v := certificateJson["name"]; v != nil {
						certificateState["name"] = v
					}
					if v := certificateJson["stored-at"]; v != nil {
						certificateState["stored_at"] = v
					}
					if v := certificateJson["status"]; v != nil {
						certificateState["status"] = v
					}
					if v := certificateJson["distinguished-name"]; v != nil {
						certificateState["distinguished_name"] = v
					}
					if v := certificateJson["base64-certificate"]; v != nil {
						certificateState["base64_certificate"] = v
					}
					if v := certificateJson["certificate-authority"]; v != nil {
						certificateState["certificate_authority"] = v.(map[string]interface{})["name"]
					}
					if v := certificateJson["expiration-date"]; v != nil {
						expirationDateJson := v.(map[string]interface{})
						expirationDateState := make(map[string]interface{})
						if v := expirationDateJson["iso-8601"]; v != nil {
							expirationDateState["iso_8601"] = v
						}
						if v := expirationDateJson["posix"]; v != nil {
							expirationDateState["posix"] = v
						}
						certificateState["expiration_date"] = []interface{}{expirationDateState}
					}
					certificatesListState = append(certificatesListState, certificateState)
				}
				vpnSettingsState["certificates"] = certificatesListState
			}
		}

		if v := vpnSettingsJson["interfaces"]; v != nil {
			interfacesList := v.([]interface{})
			if len(interfacesList) > 0 {
				var interfacesListState []map[string]interface{}
				for i := range interfacesList {
					interfacesItemJson := interfacesList[i].(map[string]interface{})
					interfacesItemState := make(map[string]interface{})
					if v := interfacesItemJson["interface-name"]; v != nil {
						interfacesItemState["interface_name"] = v
					}
					if v := interfacesItemJson["next-hop-ip"]; v != nil {
						interfacesItemState["next_hop_ip"] = v
					}
					if v := interfacesItemJson["static-nat-ip"]; v != nil {
						interfacesItemState["static_nat_ip"] = v
					}
					if v := interfacesItemJson["priority"]; v != nil {
						interfacesItemState["priority"] = v
					}
					if v := interfacesItemJson["redundancy-mode"]; v != nil {
						interfacesItemState["redundancy_mode"] = v
					}
					if v := interfacesItemJson["ip-version"]; v != nil {
						interfacesItemState["ip_version"] = v
					}
					interfacesListState = append(interfacesListState, interfacesItemState)
				}
				vpnSettingsState["interfaces"] = interfacesListState
			}
		}

		_ = d.Set("vpn_settings", []interface{}{vpnSettingsState})
	} else {
		_ = d.Set("vpn_settings", nil)
	}

	if v := gateway["tags"]; v != nil {
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

	if v := gateway["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := gateway["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := gateway["accept-syslog-messages"]; v != nil {
		_ = d.Set("accept_syslog_messages", v)
	}

	if v := gateway["anti-spam-and-email-security"]; v != nil {
		_ = d.Set("anti_spam_and_email_security", v)
	}

	if v := gateway["auto-generate-ip"]; v != nil {
		_ = d.Set("auto_generate_ip", v)
	}

	if v := gateway["auto-topology-custom-recalculation-time"]; v != nil {
		_ = d.Set("auto_topology_custom_recalculation_time", v)
	}

	if v := gateway["auto-topology-use-custom-recalculation-time"]; v != nil {
		_ = d.Set("auto_topology_use_custom_recalculation_time", v)
	}

	if v := gateway["autonomous-system-number"]; v != nil {
		_ = d.Set("autonomous_system_number", v)
	}

	if v := gateway["data-loss-prevention"]; v != nil {
		_ = d.Set("data_loss_prevention", v)
	}

	if v := gateway["dns-server"]; v != nil {
		_ = d.Set("dns_server", v)
	}

	if v := gateway["enable-log-indexing"]; v != nil {
		_ = d.Set("enable_log_indexing", v)
	}

	if v := gateway["externally-managed"]; v != nil {
		_ = d.Set("externally_managed", v)
	}

	if v := gateway["hardware-subtype"]; v != nil {
		_ = d.Set("hardware_subtype", v)
	}

	if v := gateway["install-policy-without-push"]; v != nil {
		_ = d.Set("install_policy_without_push", v)
	}

	if v := gateway["interfaces-topology-settings"]; v != nil {
		_ = d.Set("interfaces_topology_settings", v)
	}

	if v := gateway["legacy-url-filtering"]; v != nil {
		_ = d.Set("legacy_url_filtering", v)
	}

	if v := gateway["log-server"]; v != nil {
		_ = d.Set("log_server", v)
	}

	if v := gateway["mobile-access"]; v != nil {
		_ = d.Set("mobile_access", v)
	}

	if v := gateway["monitoring"]; v != nil {
		_ = d.Set("monitoring", v)
	}

	if v := gateway["network-policy-management"]; v != nil {
		_ = d.Set("network_policy_management", v)
	}

	if v := gateway["policy-server"]; v != nil {
		_ = d.Set("policy_server", v)
	}

	if v := gateway["rtm-counters-report"]; v != nil {
		_ = d.Set("rtm_counters_report", v)
	}

	if v := gateway["rtm-traffic-report"]; v != nil {
		_ = d.Set("rtm_traffic_report", v)
	}

	if v := gateway["rtm-traffic-report-per-connection"]; v != nil {
		_ = d.Set("rtm_traffic_report_per_connection", v)
	}

	if v := gateway["sic-message"]; v != nil {
		_ = d.Set("sic_message", v)
	}

	if v := gateway["smart-event-intro-correlation-unit"]; v != nil {
		_ = d.Set("smart_event_intro_correlation_unit", v)
	}

	if v := gateway["threat-prevention-mode"]; v != nil {
		_ = d.Set("threat_prevention_mode", v)
	}

	if v := gateway["trust-method"]; v != nil {
		_ = d.Set("trust_method", v)
	}

	if v := gateway["workforce-ai"]; v != nil {
		_ = d.Set("workforce_ai", v)
	}

	if v := gateway["communication-with-servers-behind-nat"]; v != nil {
		communicationWithServersBehindNatShow := v.(map[string]interface{})
		communicationWithServersBehindNatState := make(map[string]interface{})
		if v := communicationWithServersBehindNatShow["override-profile"]; v != nil {
			communicationWithServersBehindNatState["override_profile"] = v
		}
		if v := communicationWithServersBehindNatShow["value"]; v != nil {
			communicationWithServersBehindNatState["value"] = v
		}
		_ = d.Set("communication_with_servers_behind_nat", []interface{}{communicationWithServersBehindNatState})
	}

	if v := gateway["smb-logs-settings"]; v != nil {
		smbLogsSettingsShow := v.(map[string]interface{})
		smbLogsSettingsState := make(map[string]interface{})
		if v := smbLogsSettingsShow["alert-when-queue-is-full"]; v != nil {
			smbLogsSettingsState["alert_when_queue_is_full"] = v
		}
		if v := smbLogsSettingsShow["alert-when-queue-is-full-type"]; v != nil {
			smbLogsSettingsState["alert_when_queue_is_full_type"] = v
		}
		if v := smbLogsSettingsShow["detect-new-citrix-ica-application-names"]; v != nil {
			smbLogsSettingsState["detect_new_citrix_ica_application_names"] = v
		}
		if v := smbLogsSettingsShow["stop-logging-when-queue-reaches-maximal-capacity"]; v != nil {
			smbLogsSettingsState["stop_logging_when_queue_reaches_maximal_capacity"] = v
		}
		if v := smbLogsSettingsShow["stop-logging-when-queue-reaches-maximal-capacity-threshold"]; v != nil {
			smbLogsSettingsState["stop_logging_when_queue_reaches_maximal_capacity_threshold"] = v
		}
		if v := smbLogsSettingsShow["turn-on-qos-logging"]; v != nil {
			smbLogsSettingsState["turn_on_qos_logging"] = v
		}
		if v := smbLogsSettingsShow["update-account-log-every"]; v != nil {
			smbLogsSettingsState["update_account_log_every"] = v
		}
		_ = d.Set("smb_logs_settings", []interface{}{smbLogsSettingsState})
	}

	if v := gateway["zero-phishing-settings"]; v != nil {
		zeroPhishingSettingsShow := v.(map[string]interface{})
		zeroPhishingSettingsState := make(map[string]interface{})
		if v := zeroPhishingSettingsShow["gateway-fqdn-mode"]; v != nil {
			zeroPhishingSettingsState["gateway_fqdn_mode"] = v
		}
		if v := zeroPhishingSettingsShow["manual-fqdn"]; v != nil {
			zeroPhishingSettingsState["manual_fqdn"] = v
		}
		_ = d.Set("zero_phishing_settings", []interface{}{zeroPhishingSettingsState})
	}

	if v := gateway["fetch-policy-scheduler"]; v != nil {
		fetchPolicySchedulerShow := v.(map[string]interface{})
		fetchPolicySchedulerState := make(map[string]interface{})
		if v := fetchPolicySchedulerShow["enabled"]; v != nil {
			fetchPolicySchedulerState["enabled"] = v
		}
		if v := fetchPolicySchedulerShow["schedule"]; v != nil {
			fetchPolicySchedulerState["schedule"] = v.(map[string]interface{})["name"]
		}
		_ = d.Set("fetch_policy_scheduler", []interface{}{fetchPolicySchedulerState})
	}

	if v := gateway["trust-details"]; v != nil {
		trustDetailsShow := v.(map[string]interface{})
		trustDetailsState := make(map[string]interface{})
		if v := trustDetailsShow["authentication-token"]; v != nil {
			trustDetailsState["authentication_token"] = v
		}
		if v := trustDetailsShow["cloud-communication-details"]; v != nil {
			cloudCommunicationDetailsShow := v.(map[string]interface{})
			cloudCommunicationDetailsState := make(map[string]interface{})
			if v := cloudCommunicationDetailsShow["ip"]; v != nil {
				cloudCommunicationDetailsState["ip"] = v
			}
			if v := cloudCommunicationDetailsShow["status"]; v != nil {
				cloudCommunicationDetailsState["status"] = v
			}
			trustDetailsState["cloud_communication_details"] = []interface{}{cloudCommunicationDetailsState}
		}
		if v := trustDetailsShow["gateway-mac-address"]; v != nil {
			trustDetailsState["gateway_mac_address"] = v
		}
		if v := trustDetailsShow["identification-method"]; v != nil {
			trustDetailsState["identification_method"] = v
		}
		if v := trustDetailsShow["status"]; v != nil {
			trustDetailsState["status"] = v
		}
		if v := trustDetailsShow["token-expiration-date"]; v != nil {
			trustDetailsState["token_expiration_date"] = v
		}
		_ = d.Set("trust_details", []interface{}{trustDetailsState})
	}

	return nil
}
