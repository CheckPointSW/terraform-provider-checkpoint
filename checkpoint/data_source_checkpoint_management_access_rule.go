package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strings"
)

func dataSourceManagementAccessRule() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementAccessRuleRead,

		Schema: map[string]*schema.Schema{
			"layer": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Layer that the rule belongs to identified by the name or UID.",
			},
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
			"action": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "\"Accept\", \"Drop\", \"Ask\", \"Inform\", \"Reject\", \"User Auth\", \"Client Auth\", \"Apply Layer\".",
			},
			"action_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Action settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"user_auth_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"allowed_http_servers": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Restricts which HTTP servers authenticated users may access.",
									},
									"destination": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "How destination hosts are matched against the user database.",
									},
									"source": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "How source users are matched against the user database.",
									},
								},
							},
						},
						"client_auth_settings": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"destination": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "How destination hosts are matched against the user database.",
									},
									"require_desktop_config_verification": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "When true, the rule only applies if the client's desktop security policy configuration has been verified by the gateway.",
									},
									"sessions_limit": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Maximum number of concurrent sessions allowed per user. Only used when unlimited-sessions is false. Must be 1 or greater.",
									},
									"sign_on_method": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Mechanism used to authenticate the client.",
									},
									"sign_on_type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Determines whether a standard or specific sign-on policy is applied. Allowed values: standard, specific.",
									},
									"source": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "How source users are matched against the user database.",
									},
									"timeout": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Controls how long an authenticated session remains valid.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enable": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "When true, authenticated sessions expire after the duration defined by minutes. When false, authenticated sessions never expire.",
												},
												"minutes": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Number of minutes before an authenticated session expires. Only used when enable is true. Must be 0 or greater.",
												},
												"refreshable": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "When true, the session timeout timer resets each time the authenticated user performs a network action.",
												},
											},
										},
									},
									"tracking": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Action to take in the log when a user successfully authenticates. Allowed values: none, log, alert.",
									},
									"unlimited_sessions": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "When true, a single user may have any number of concurrent authenticated sessions. When false, the number of concurrent sessions is capped at sessions...",
									},
								},
							},
						},
						"enable_identity_captive_portal": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "N/A",
						},
						"limit": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
					},
				},
			},
			"content": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of processed file types that this rule applies on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"content_direction": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "On which direction the file types processing is applied.",
			},
			"content_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if negate is set for data.",
			},
			"custom_fields": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Custom fields.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"field_1": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "First custom field.",
						},
						"field_2": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Second custom field.",
						},
						"field_3": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Third custom field.",
						},
					},
				},
			},
			"destination": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"destination_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if negate is set for destination.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Enable/Disable the rule.",
			},
			"inline_layer": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Inline Layer identified by the name or UID. Relevant only if \"Action\" was set to \"Apply Layer\".",
			},
			"install_on": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Which Gateways identified by the name or UID to install the policy on.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if negate is set for service.",
			},
			"source": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of Network objects identified by the name or UID.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"source_negate": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True if negate is set for source.",
			},
			"time": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "List of time objects. For example: \"Weekend\", \"Off-Work\", \"Every-Day\".",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"track": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Track Settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"accounting": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Turns accounting for track on and off.",
						},
						"alert": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of alert for the track.",
						},
						"enable_firewall_session": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determine whether to generate session log to firewall only connections.",
						},
						"per_connection": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether to perform the log per connection.",
						},
						"per_session": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Determines whether to perform the log per session.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "\"Log\", \"Extended Log\", \"Detailed Log\", \"None\".",
						},
					},
				},
			},
			"user_check": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "User check settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"confirm": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"custom_frequency": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "N/A",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"every": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "N/A",
									},
									"unit": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
						"frequency": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"interaction": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
					},
				},
			},
			"vpn": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Valid values \"Any\", \"All_GwToGw\" or VPN community name",
			},
			"vpn_communities": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Collection of VPN communities identified by name",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vpn_directional": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Collection of VPN directional",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"from": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "From VPN community",
						},
						"to": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "To VPN community",
						},
					},
				},
			},
			"comments": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Comments string.",
			},
			"fields_with_uid_identifier": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of resource fields that will use object UIDs as object identifiers. Default is object name.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"service_resource": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource of the service.",
			},
			"expiration_settings": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Displays the expiration date settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expiration_date": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Expiration date.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"posix": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: `N/A`,
									},
									"iso_8601": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: `N/A`,
									},
								},
							},
						},
						"expired": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Expired rule.",
						},
						"has_expiration_date": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Rule has expiration date.",
						},
					},
				},
			},
			"destination_ranges": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Displays the destination as ranges of IP addresses, in case show-as-ranges is set to true.<br />In this case, 'destination' and 'destination-negate' p...",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"excluded_others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as IP addresses and are negated in the given rule - for example if negate is set for the source or destination of th...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv4": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of IPv4 addresses that match in the given rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"start": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
						"ipv6": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of IPv6 addresses that match in the given rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"start": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
						"others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as IP addresses and match the given rule. The details-level parameter of the request determines whether they are dis...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"hits": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Hits count object.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"first_date": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "N/A",
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
						"last_date": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "N/A",
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
						"level": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"percentage": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "N/A",
						},
						"value": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "N/A",
						},
					},
				},
			},
			"service_ranges": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Displays the services and applications as ranges of port numbers, in case show-as-ranges is set to true.<br />In this case, 'service' and 'service-neg...",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"excluded_others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as port numbers and are negated in the given rule - for example if negate is set for the service of this rule. The d...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as port numbers and match the given rule. The details-level parameter of the request determines whether they are dis...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"tcp": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of TCP ports that match in the given rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"start": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
						"udp": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of UDP ports that match in the given rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"start": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
					},
				},
			},
			"source_ranges": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Displays the source as ranges of IP addresses, in case show-as-ranges is set to true.<br />In this case, 'source' and 'source-negate' parameters are o...",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"excluded_others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as IP addresses and are negated in the given rule - for example if negate is set for the source or destination of th...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv4": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of IPv4 addresses that match in the given rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"start": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
						"ipv6": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Range of IPv6 addresses that match in the given rule.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"end": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
									"start": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "N/A",
									},
								},
							},
						},
						"others": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Objects which are not represented as IP addresses and match the given rule. The details-level parameter of the request determines whether they are dis...",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementAccessRuleRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := map[string]interface{}{
		"layer": d.Get("layer"),
	}

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showAccessRuleRes, err := client.ApiCall("show-access-rule", payload, client.GetSessionID(), true, client.IsProxyUsed())
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showAccessRuleRes.Success {
		return fmt.Errorf("%s", showAccessRuleRes.ErrorMsg)
	}

	accessRule := showAccessRuleRes.GetData()

	log.Println("Read Access Rule - Show JSON = ", accessRule)

	if v := accessRule["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}
	if v := accessRule["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := accessRule["action"]; v != nil {
		actionId := resolveObjectIdentifier("action", accessRule["action"], d)
		if actionId == "Inner Layer" {
			actionId = "Apply Layer"
		}
		_ = d.Set("action", actionId)
	}

	if accessRule["action-settings"] != nil {

		actionSettingsMap := accessRule["action-settings"].(map[string]interface{})

		actionSettingsMapToReturn := make(map[string]interface{})

		if v := actionSettingsMap["user-auth-settings"]; v != nil {
			userAuthSettingsShow := v.(map[string]interface{})
			userAuthSettingsState := make(map[string]interface{})
			if v := userAuthSettingsShow["allowed-http-servers"]; v != nil {
				userAuthSettingsState["allowed_http_servers"] = v
			}
			if v := userAuthSettingsShow["destination"]; v != nil {
				userAuthSettingsState["destination"] = v
			}
			if v := userAuthSettingsShow["source"]; v != nil {
				userAuthSettingsState["source"] = v
			}
			actionSettingsMapToReturn["user_auth_settings"] = []interface{}{userAuthSettingsState}
		}
		if v := actionSettingsMap["client-auth-settings"]; v != nil {
			clientAuthSettingsShow := v.(map[string]interface{})
			clientAuthSettingsState := make(map[string]interface{})
			if v := clientAuthSettingsShow["destination"]; v != nil {
				clientAuthSettingsState["destination"] = v
			}
			if v := clientAuthSettingsShow["require-desktop-config-verification"]; v != nil {
				clientAuthSettingsState["require_desktop_config_verification"] = v
			}
			if v := clientAuthSettingsShow["sessions-limit"]; v != nil {
				clientAuthSettingsState["sessions_limit"] = v
			}
			if v := clientAuthSettingsShow["sign-on-method"]; v != nil {
				clientAuthSettingsState["sign_on_method"] = v
			}
			if v := clientAuthSettingsShow["sign-on-type"]; v != nil {
				clientAuthSettingsState["sign_on_type"] = v
			}
			if v := clientAuthSettingsShow["source"]; v != nil {
				clientAuthSettingsState["source"] = v
			}
			if v := clientAuthSettingsShow["timeout"]; v != nil {
				timeoutShow := v.(map[string]interface{})
				timeoutState := make(map[string]interface{})
				if v := timeoutShow["enable"]; v != nil {
					timeoutState["enable"] = v
				}
				if v := timeoutShow["minutes"]; v != nil {
					timeoutState["minutes"] = v
				}
				if v := timeoutShow["refreshable"]; v != nil {
					timeoutState["refreshable"] = v
				}
				clientAuthSettingsState["timeout"] = []interface{}{timeoutState}
			}
			if v := clientAuthSettingsShow["tracking"]; v != nil {
				clientAuthSettingsState["tracking"] = v
			}
			if v := clientAuthSettingsShow["unlimited-sessions"]; v != nil {
				clientAuthSettingsState["unlimited_sessions"] = v
			}
			actionSettingsMapToReturn["client_auth_settings"] = []interface{}{clientAuthSettingsState}
		}
		if v, _ := actionSettingsMap["enable-identity-captive-portal"]; v != nil {
			actionSettingsMapToReturn["enable_identity_captive_portal"] = v.(bool)
		}

		if v, _ := actionSettingsMap["limit"]; v != nil {
			actionSettingsMapToReturn["limit"] = v.(map[string]interface{})["name"].(string)
		}

		_ = d.Set("action_settings", []interface{}{actionSettingsMapToReturn})
	} else {
		_ = d.Set("action_settings", nil)
	}

	if accessRule["content"] != nil {
		contentIds := resolveListOfIdentifiers("content", accessRule["content"], d)
		_ = d.Set("content", contentIds)
	} else {
		_ = d.Set("content", nil)
	}

	if v := accessRule["content-direction"]; v != nil {
		_ = d.Set("content_direction", v)
	}

	if v := accessRule["content-negate"]; v != nil {
		_ = d.Set("content_negate", v)
	}

	if accessRule["custom-fields"] != nil {

		customFieldsMap := accessRule["custom-fields"].(map[string]interface{})

		customFieldsMapToReturn := make(map[string]interface{})

		if v, _ := customFieldsMap["field-1"]; v != nil {
			customFieldsMapToReturn["field_1"] = v
		}

		if v, _ := customFieldsMap["field-2"]; v != nil {
			customFieldsMapToReturn["field_2"] = v
		}

		if v, _ := customFieldsMap["field-3"]; v != nil {
			customFieldsMapToReturn["field_3"] = v
		}
		_ = d.Set("custom_fields", []interface{}{customFieldsMapToReturn})
	} else {
		_ = d.Set("custom_fields", nil)
	}

	if accessRule["destination"] != nil {
		destinationIds := resolveListOfIdentifiers("destination", accessRule["destination"], d)
		_ = d.Set("destination", destinationIds)
	}

	if v := accessRule["destination-negate"]; v != nil {
		_ = d.Set("destination_negate", v)
	}

	if v := accessRule["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if v := accessRule["inline-layer"]; v != nil {
		_ = d.Set("inline_layer", v)
	}

	if accessRule["install-on"] != nil {
		installOnIds := resolveListOfIdentifiers("install-on", accessRule["install-on"], d)
		_ = d.Set("install_on", installOnIds)
	}

	if accessRule["service"] != nil {
		serviceIds := resolveListOfIdentifiers("service", accessRule["service"], d)
		_ = d.Set("service", serviceIds)
	}

	if v := accessRule["service-negate"]; v != nil {
		_ = d.Set("service_negate", v)
	}

	if accessRule["source"] != nil {
		sourceIds := resolveListOfIdentifiers("source", accessRule["source"], d)
		_ = d.Set("source", sourceIds)
	}

	if v := accessRule["source-negate"]; v != nil {
		_ = d.Set("source_negate", v)
	}

	if accessRule["time"] != nil {
		timeIds := resolveListOfIdentifiers("time", accessRule["time"], d)
		_ = d.Set("time", timeIds)
	}
	if accessRule["track"] != nil {

		trackMap := accessRule["track"].(map[string]interface{})

		trackMapToReturn := make(map[string]interface{})
		if v := trackMap["accounting"]; v != nil {
			trackMapToReturn["accounting"] = v.(bool)
		}

		if v, _ := trackMap["alert"]; v != nil {
			trackMapToReturn["alert"] = v.(string)
		}

		if v := trackMap["enable-firewall-session"]; v != nil {
			trackMapToReturn["enable_firewall_session"] = v.(bool)
		}

		if v := trackMap["per-connection"]; v != nil {
			trackMapToReturn["per_connection"] = v.(bool)
		}

		if v := trackMap["per-session"]; v != nil {
			trackMapToReturn["per_session"] = v.(bool)
		}

		if v, _ := trackMap["type"]; v != nil {
			trackMapToReturn["type"] = v.(map[string]interface{})["name"].(string)
		}
		err = d.Set("track", []interface{}{trackMapToReturn})

	} else {
		_ = d.Set("track", nil)
	}

	if accessRule["user-check"] != nil {

		userCheckMap := accessRule["user-check"].(map[string]interface{})

		userCheckMapToReturn := make(map[string]interface{})

		if v, _ := userCheckMap["confirm"]; v != nil {
			userCheckMapToReturn["confirm"] = v
		}

		if v, ok := userCheckMap["custom-frequency"]; ok {

			userCheckConfigMap := v.(map[string]interface{})
			userCheckConfigMapToReturn := make(map[string]interface{})

			if v, _ := userCheckConfigMap["every"]; v != nil {
				userCheckConfigMapToReturn["every"] = v
			}

			if v, _ := userCheckConfigMap["unit"]; v != nil {
				userCheckConfigMapToReturn["unit"] = v
			}
			userCheckMapToReturn["custom_frequency"] = []interface{}{userCheckConfigMapToReturn}
		}

		if v, _ := userCheckMap["frequency"]; v != nil {
			userCheckMapToReturn["frequency"] = v
		}

		if v, _ := userCheckMap["interaction"]; v != nil {
			userCheckMapToReturn["interaction"] = v.(map[string]interface{})["name"]
		}

		_ = d.Set("user_check", []interface{}{userCheckMapToReturn})
	} else {
		_ = d.Set("user_check", nil)
	}

	if v := accessRule["vpn"]; v != nil {
		vpnList := v.([]interface{})
		if len(vpnList) > 0 {
			vpnType := vpnList[0].(map[string]interface{})["type"].(string)
			if len(vpnList) == 1 && vpnType != "VpnDirectionalElement" { // BC
				vpnId := resolveObjectIdentifier("vpn", v.([]interface{})[0], d)
				_ = d.Set("vpn", vpnId)
				_ = d.Set("vpn_communities", nil)
				_ = d.Set("vpn_directional", nil)
			} else if vpnType != "VpnDirectionalElement" {
				vpnIds := resolveListOfIdentifiers("vpn", vpnList, d)
				_ = d.Set("vpn_communities", vpnIds)
				_ = d.Set("vpn", nil)
				_ = d.Set("vpn_directional", nil)
			} else if vpnType == "VpnDirectionalElement" {
				var vpnDirectionalListState []map[string]interface{}
				for i := range vpnList {
					vpnDirectionalObj := vpnList[i].(map[string]interface{})
					if v, _ := vpnDirectionalObj["name"]; v != nil {
						vpnDirectionalNames := strings.Split(v.(string), "->")
						vpnDirectionalState := make(map[string]interface{})
						vpnDirectionalState["from"] = vpnDirectionalNames[0]
						vpnDirectionalState["to"] = vpnDirectionalNames[1]
						vpnDirectionalListState = append(vpnDirectionalListState, vpnDirectionalState)
					}
				}
				_ = d.Set("vpn_directional", vpnDirectionalListState)
				_ = d.Set("vpn_communities", nil)
				_ = d.Set("vpn", nil)
			} else {
				return fmt.Errorf("Cannot read invalid VPN type [%s]", vpnType)
			}
		}
	}

	if v := accessRule["comments"]; v != nil {
		_ = d.Set("comments", v)
	}
	if v := accessRule["service-resource"]; v != nil {
		_ = d.Set("service_resource", v.(map[string]interface{})["name"])
	}

	if v := accessRule["expiration-settings"]; v != nil {
		expirationSettingsShow := v.(map[string]interface{})
		expirationSettingsState := make(map[string]interface{})
		if v, ok := expirationSettingsShow["expiration-date"].(map[string]interface{}); ok {
			dateState := make(map[string]interface{})
			if p, ok := v["posix"].(float64); ok {
				dateState["posix"] = int(p)
			}
			if iso := v["iso-8601"]; iso != nil {
				dateState["iso_8601"] = iso
			}
			expirationSettingsState["expiration_date"] = []interface{}{dateState}
		}
		if v := expirationSettingsShow["expired"]; v != nil {
			expirationSettingsState["expired"] = v
		}
		if v := expirationSettingsShow["has-expiration-date"]; v != nil {
			expirationSettingsState["has_expiration_date"] = v
		}
		_ = d.Set("expiration_settings", []interface{}{expirationSettingsState})
	}

	if v := accessRule["destination-ranges"]; v != nil {
		destinationRangesShow := v.(map[string]interface{})
		destinationRangesState := make(map[string]interface{})
		if v := destinationRangesShow["excluded-others"]; v != nil {
			excludedOthersIdsList := v.([]interface{})
			var excludedOthersIds = make([]string, 0)
			for _, item := range excludedOthersIdsList {
				excludedOthersIds = append(excludedOthersIds, item.(map[string]interface{})["name"].(string))
			}
			destinationRangesState["excluded_others"] = excludedOthersIds
		}
		if v := destinationRangesShow["ipv4"]; v != nil {
			ipv4List := v.([]interface{})
			var ipv4ListState []map[string]interface{}
			for i := range ipv4List {
				ipv4Show := ipv4List[i].(map[string]interface{})
				ipv4State := make(map[string]interface{})
				if v := ipv4Show["end"]; v != nil {
					ipv4State["end"] = v
				}
				if v := ipv4Show["start"]; v != nil {
					ipv4State["start"] = v
				}
				ipv4ListState = append(ipv4ListState, ipv4State)
			}
			destinationRangesState["ipv4"] = ipv4ListState
		}
		if v := destinationRangesShow["ipv6"]; v != nil {
			ipv6List := v.([]interface{})
			var ipv6ListState []map[string]interface{}
			for i := range ipv6List {
				ipv6Show := ipv6List[i].(map[string]interface{})
				ipv6State := make(map[string]interface{})
				if v := ipv6Show["end"]; v != nil {
					ipv6State["end"] = v
				}
				if v := ipv6Show["start"]; v != nil {
					ipv6State["start"] = v
				}
				ipv6ListState = append(ipv6ListState, ipv6State)
			}
			destinationRangesState["ipv6"] = ipv6ListState
		}
		if v := destinationRangesShow["others"]; v != nil {
			othersIdsList := v.([]interface{})
			var othersIds = make([]string, 0)
			for _, item := range othersIdsList {
				othersIds = append(othersIds, item.(map[string]interface{})["name"].(string))
			}
			destinationRangesState["others"] = othersIds
		}
		_ = d.Set("destination_ranges", []interface{}{destinationRangesState})
	}

	if v := accessRule["hits"]; v != nil {
		hitsShow := v.(map[string]interface{})
		hitsState := make(map[string]interface{})
		if v := hitsShow["first-date"]; v != nil {
			firstDateShow := v.(map[string]interface{})
			firstDateState := make(map[string]interface{})
			if v := firstDateShow["iso-8601"]; v != nil {
				firstDateState["iso_8601"] = v
			}
			if v := firstDateShow["posix"]; v != nil {
				firstDateState["posix"] = v
			}
			hitsState["first_date"] = []interface{}{firstDateState}
		}
		if v := hitsShow["last-date"]; v != nil {
			lastDateShow := v.(map[string]interface{})
			lastDateState := make(map[string]interface{})
			if v := lastDateShow["iso-8601"]; v != nil {
				lastDateState["iso_8601"] = v
			}
			if v := lastDateShow["posix"]; v != nil {
				lastDateState["posix"] = v
			}
			hitsState["last_date"] = []interface{}{lastDateState}
		}
		if v := hitsShow["level"]; v != nil {
			hitsState["level"] = v
		}
		if v := hitsShow["percentage"]; v != nil {
			hitsState["percentage"] = v
		}
		if v := hitsShow["value"]; v != nil {
			hitsState["value"] = v
		}
		_ = d.Set("hits", []interface{}{hitsState})
	}

	if v := accessRule["service-ranges"]; v != nil {
		serviceRangesShow := v.(map[string]interface{})
		serviceRangesState := make(map[string]interface{})
		if v := serviceRangesShow["excluded-others"]; v != nil {
			excludedOthersIdsList := v.([]interface{})
			var excludedOthersIds = make([]string, 0)
			for _, item := range excludedOthersIdsList {
				excludedOthersIds = append(excludedOthersIds, item.(map[string]interface{})["name"].(string))
			}
			serviceRangesState["excluded_others"] = excludedOthersIds
		}
		if v := serviceRangesShow["others"]; v != nil {
			othersIdsList := v.([]interface{})
			var othersIds = make([]string, 0)
			for _, item := range othersIdsList {
				othersIds = append(othersIds, item.(map[string]interface{})["name"].(string))
			}
			serviceRangesState["others"] = othersIds
		}
		if v := serviceRangesShow["tcp"]; v != nil {
			tcpList := v.([]interface{})
			var tcpListState []map[string]interface{}
			for i := range tcpList {
				tcpShow := tcpList[i].(map[string]interface{})
				tcpState := make(map[string]interface{})
				if v := tcpShow["end"]; v != nil {
					tcpState["end"] = v
				}
				if v := tcpShow["start"]; v != nil {
					tcpState["start"] = v
				}
				tcpListState = append(tcpListState, tcpState)
			}
			serviceRangesState["tcp"] = tcpListState
		}
		if v := serviceRangesShow["udp"]; v != nil {
			udpList := v.([]interface{})
			var udpListState []map[string]interface{}
			for i := range udpList {
				udpShow := udpList[i].(map[string]interface{})
				udpState := make(map[string]interface{})
				if v := udpShow["end"]; v != nil {
					udpState["end"] = v
				}
				if v := udpShow["start"]; v != nil {
					udpState["start"] = v
				}
				udpListState = append(udpListState, udpState)
			}
			serviceRangesState["udp"] = udpListState
		}
		_ = d.Set("service_ranges", []interface{}{serviceRangesState})
	}

	if v := accessRule["source-ranges"]; v != nil {
		sourceRangesShow := v.(map[string]interface{})
		sourceRangesState := make(map[string]interface{})
		if v := sourceRangesShow["excluded-others"]; v != nil {
			excludedOthersIdsList := v.([]interface{})
			var excludedOthersIds = make([]string, 0)
			for _, item := range excludedOthersIdsList {
				excludedOthersIds = append(excludedOthersIds, item.(map[string]interface{})["name"].(string))
			}
			sourceRangesState["excluded_others"] = excludedOthersIds
		}
		if v := sourceRangesShow["ipv4"]; v != nil {
			ipv4List := v.([]interface{})
			var ipv4ListState []map[string]interface{}
			for i := range ipv4List {
				ipv4Show := ipv4List[i].(map[string]interface{})
				ipv4State := make(map[string]interface{})
				if v := ipv4Show["end"]; v != nil {
					ipv4State["end"] = v
				}
				if v := ipv4Show["start"]; v != nil {
					ipv4State["start"] = v
				}
				ipv4ListState = append(ipv4ListState, ipv4State)
			}
			sourceRangesState["ipv4"] = ipv4ListState
		}
		if v := sourceRangesShow["ipv6"]; v != nil {
			ipv6List := v.([]interface{})
			var ipv6ListState []map[string]interface{}
			for i := range ipv6List {
				ipv6Show := ipv6List[i].(map[string]interface{})
				ipv6State := make(map[string]interface{})
				if v := ipv6Show["end"]; v != nil {
					ipv6State["end"] = v
				}
				if v := ipv6Show["start"]; v != nil {
					ipv6State["start"] = v
				}
				ipv6ListState = append(ipv6ListState, ipv6State)
			}
			sourceRangesState["ipv6"] = ipv6ListState
		}
		if v := sourceRangesShow["others"]; v != nil {
			othersIdsList := v.([]interface{})
			var othersIds = make([]string, 0)
			for _, item := range othersIdsList {
				othersIds = append(othersIds, item.(map[string]interface{})["name"].(string))
			}
			sourceRangesState["others"] = othersIds
		}
		_ = d.Set("source_ranges", []interface{}{sourceRangesState})
	}

	return nil
}
