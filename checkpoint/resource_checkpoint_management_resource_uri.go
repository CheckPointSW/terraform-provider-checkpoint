package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func resourceManagementResourceUri() *schema.Resource {
	return &schema.Resource{
		Create: createManagementResourceUri,
		Read:   readManagementResourceUri,
		Update: updateManagementResourceUri,
		Delete: deleteManagementResourceUri,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"use_this_resource_to": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Select the use of the URI resource.",
				Default:     "enforce_uri_capabilities",
			},
			"connection_methods": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "Connection methods.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transparent": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The security server is invisible to the client that originates the connection, and to the server. The Transparent connection method is the most secure.",
							Default:     true,
						},
						"proxy": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The Resource is applied when people specify the Check Point Security Gateway as a proxy in their browser.",
							Default:     true,
						},
						"tunneling": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The Resource is applied when people specify the Security Gateway as a proxy in their browser, and is used for connections where Security Gateway cannot examine the contents of the packets, not even the URL.",
							Default:     false,
						},
					},
				},
			},
			"uri_match_specification_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The type can be Wild Cards or UFP, where a UFP server holds categories of forbidden web sites.",
				Default:     "wildcards",
			},
			"exception_track": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Configures how to track connections that match this rule but fail the content security checks. An example of an exception is a connection with an unsupported scheme or method.",
				Default:     "None",
			},
			"match_ufp": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Match - UFP settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UID or Name of the UFP server that is an OPSEC certified third party application that checks URLs against a list of permitted categories.",
						},
						"caching_control": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies if and how caching is to be enabled.",
							Default:     "no_caching",
						},
						"ignore_ufp_server_after_failure": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "The UFP server will be ignored after numerous UFP server connections were unsuccessful.",
							Default:     false,
						},
						"number_of_failures_before_ignore": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Signifies at what point the UFP server should be ignored.",
							Default:     0,
						},
						"timeout_before_reconnecting": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The amount of time that must pass before a UFP server connection should be attempted.",
							Default:     0,
						},
					},
				},
			},
			"match_wildcards": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Match - Wildcards settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schemes": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Select the URI Schemes to which this resource applies.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Http scheme.",
										Default:     false,
									},
									"ftp": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Ftp scheme.",
										Default:     false,
									},
									"gopher": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Gopher scheme.",
										Default:     false,
									},
									"mailto": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Mailto scheme.",
										Default:     false,
									},
									"news": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "News scheme.",
										Default:     false,
									},
									"wais": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Wais scheme.",
										Default:     false,
									},
									"other": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "You can specify another scheme in the Other field. You can use wildcards.",
									},
								},
							},
						},
						"methods": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Select the URI Schemes to which this resource applies.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"get": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "GET method.",
										Default:     false,
									},
									"post": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "POST method.",
										Default:     false,
									},
									"head": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "HEAD method.",
										Default:     false,
									},
									"put": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "PUT method.",
										Default:     false,
									},
									"other": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "You can specify another method in the Other field. You can use wildcards.",
									},
								},
							},
						},
						"host": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The functionality of the Host parameter depends on the DNS setup of the addressed server. For the host, only the IP address or the full DNS name should be used.",
						},
						"path": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Name matching is based on appending the file name in the request to the current working directory (unless the file name is already a full path name) and comparing the result to the path specified in the Resource definition.",
						},
						"query": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The parameters that are sent to the URI when it is accessed.",
						},
					},
				},
			},
			"action": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Action settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"replacement_uri": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "If the Action in a rule which uses this resource is Drop or Reject, then the Replacement URI is displayed instead of the one requested by the user.",
						},
						"strip_script_tags": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip JAVA scripts.",
							Default:     false,
						},
						"strip_applet_tags": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip JAVA applets.",
							Default:     false,
						},
						"strip_activex_tags": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip activeX tags.",
							Default:     false,
						},
						"strip_ftp_links": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip ftp links.",
							Default:     false,
						},
						"strip_port_strings": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Strip ports.",
							Default:     false,
						},
					},
				},
			},
			"cvp": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "CVP settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cvp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Select to enable the Content Vectoring Protocol.",
							Default:     false,
						},
						"server": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application.",
						},
						"allowed_to_modify_content": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Configures the CVP server to inspect but not modify content.",
							Default:     true,
						},
						"send_http_headers_to_cvp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Select, if you would like the CVP server to check the HTTP headers of the message packets.",
							Default:     false,
						},
						"reply_order": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Designates when the CVP server returns data to the Security Gateway security server.",
							Default:     "return_data_after_content_is_approved",
						},
						"send_http_request_to_cvp": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Used to protect against undesirable content in the HTTP request, for example, when inspecting peer-to-peer connections.",
							Default:     false,
						},
						"send_only_unsafe_file_types": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Improves the performance of the CVP server. This option does not send to the CVP server traffic that is considered safe.",
							Default:     true,
						},
					},
				},
			},
			"soap": {
				Type:        schema.TypeMap,
				Optional:    true,
				Description: "SOAP settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspection": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Allow all SOAP Requests, or Allow only SOAP requests specified in the following file-id.",
							Default:     "allow_all_soap_requests",
						},
						"file_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A file containing SOAP requests.",
						},
						"track_connections": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The method of tracking SOAP connections.",
							Default:     "none",
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

func createManagementResourceUri(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	resourceUri := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		resourceUri["name"] = v.(string)
	}

	if v, ok := d.GetOk("use_this_resource_to"); ok {
		resourceUri["use-this-resource-to"] = v.(string)
	}

	if _, ok := d.GetOk("connection_methods"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("connection_methods.transparent"); ok {
			res["transparent"] = v
		}
		if v, ok := d.GetOk("connection_methods.proxy"); ok {
			res["proxy"] = v
		}
		if v, ok := d.GetOk("connection_methods.tunneling"); ok {
			res["tunneling"] = v
		}
		resourceUri["connection-methods"] = res
	}

	if v, ok := d.GetOk("uri_match_specification_type"); ok {
		resourceUri["uri-match-specification-type"] = v.(string)
	}

	if v, ok := d.GetOk("exception_track"); ok {
		resourceUri["exception-track"] = v.(string)
	}

	if _, ok := d.GetOk("match_ufp"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("match_ufp.0.server"); ok {
			res["server"] = v.(string)
		}
		if v, ok := d.GetOk("match_ufp.0.caching_control"); ok {
			res["caching-control"] = v
		}
		if v, ok := d.GetOk("match_ufp.0.ignore_ufp_server_after_failure"); ok {
			res["ignore-ufp-server-after-failure"] = v
		}
		if v, ok := d.GetOk("match_ufp.0.number_of_failures_before_ignore"); ok {
			res["number-of-failures-before-ignore"] = v
		}
		if v, ok := d.GetOk("match_ufp.0.timeout_before_reconnecting"); ok {
			res["timeout-before-reconnecting"] = v
		}
		resourceUri["match-ufp"] = res
	}

	if v, ok := d.GetOk("match_wildcards"); ok {

		matchWildcardsList := v.([]interface{})

		if len(matchWildcardsList) > 0 {

			matchWildcardsPayload := make(map[string]interface{})

			if _, ok := d.GetOk("match_wildcards.0.schemes"); ok {

				schemesPayload := make(map[string]interface{})

				if v, ok := d.GetOk("match_wildcards.0.schemes.0.http"); ok {
					schemesPayload["http"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.schemes.0.ftp"); ok {
					schemesPayload["ftp"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.schemes.0.gopher"); ok {
					schemesPayload["gopher"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.schemes.0.mailto"); ok {
					schemesPayload["mailto"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.schemes.0.news"); ok {
					schemesPayload["news"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.schemes.0.wais"); ok {
					schemesPayload["wais"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.schemes.0.other"); ok {
					schemesPayload["other"] = v
				}
				matchWildcardsPayload["schemes"] = schemesPayload
			}
			if _, ok := d.GetOk("match_wildcards.0.methods"); ok {

				methodsPayload := make(map[string]interface{})

				if v, ok := d.GetOk("match_wildcards.0.methods.0.get"); ok {
					methodsPayload["get"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.methods.0.post"); ok {
					methodsPayload["post"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.methods.0.head"); ok {
					methodsPayload["head"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.methods.0.put"); ok {
					methodsPayload["put"] = v
				}
				if v, ok := d.GetOk("match_wildcards.0.methods.0.other"); ok {
					methodsPayload["other"] = v
				}
				matchWildcardsPayload["methods"] = methodsPayload
			}
			if v, ok := d.GetOk("match_wildcards.0.host"); ok {
				matchWildcardsPayload["host"] = v
			}
			if v, ok := d.GetOk("match_wildcards.0.path"); ok {
				matchWildcardsPayload["path"] = v
			}
			if v, ok := d.GetOk("match_wildcards.0.query"); ok {
				matchWildcardsPayload["query"] = v
			}
			resourceUri["match-wildcards"] = matchWildcardsPayload
		}
	}
	if _, ok := d.GetOk("action"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("action.0.replacement_uri"); ok {
			res["replacement-uri"] = v
		}
		if v, ok := d.GetOk("action.0.strip_script_tags"); ok {
			res["strip-script-tags"] = v
		}
		if v, ok := d.GetOk("action.0.strip_applet_tags"); ok {
			res["strip-applet-tags"] = v
		}
		if v, ok := d.GetOk("action.0.strip_activex_tags"); ok {
			res["strip-activex-tags"] = v
		}
		if v, ok := d.GetOk("action.0.strip_ftp_links"); ok {
			res["strip-ftp-links"] = v
		}
		if v, ok := d.GetOk("action.0.strip_port_strings"); ok {
			res["strip-port-strings"] = v
		}
		resourceUri["action"] = res
	}

	if _, ok := d.GetOk("cvp"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("cvp.0.enable_cvp"); ok {
			res["enable-cvp"] = v
		}
		if v, ok := d.GetOk("cvp.0.server"); ok {

			if len(v.(string)) > 0 {
				res["server"] = v
			}
		}
		if v, ok := d.GetOk("cvp.0.allowed_to_modify_content"); ok {
			res["allowed-to-modify-content"] = v
		}
		if v, ok := d.GetOk("cvp.0.send_http_headers_to_cvp"); ok {
			res["send-http-headers-to-cvp"] = v
		}
		if v, ok := d.GetOk("cvp.0.reply_order"); ok {
			res["reply-order"] = v
		}
		if v, ok := d.GetOk("cvp.0.send_http_request_to_cvp"); ok {
			res["send-http-request-to-cvp"] = v
		}
		if v, ok := d.GetOk("cvp.0.send_only_unsafe_file_types"); ok {
			res["send-only-unsafe-file-types"] = v
		}
		resourceUri["cvp"] = res
	}

	if _, ok := d.GetOk("soap"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOk("soap.inspection"); ok {
			res["inspection"] = v
		}
		if v, ok := d.GetOk("soap.file_id"); ok {
			res["file-id"] = v
		}
		if v, ok := d.GetOk("soap.track_connections"); ok {
			res["track-connections"] = v
		}

		resourceUri["soap"] = res
	}

	if v, ok := d.GetOk("tags"); ok {
		resourceUri["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOk("color"); ok {
		resourceUri["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		resourceUri["comments"] = v.(string)
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceUri["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceUri["ignore-errors"] = v.(bool)
	}

	log.Println("Create ResourceUri - Map = ", resourceUri)

	addResourceUriRes, err := client.ApiCall("add-resource-uri", resourceUri, client.GetSessionID(), true, false)
	if err != nil || !addResourceUriRes.Success {
		if addResourceUriRes.ErrorMsg != "" {
			return fmt.Errorf(addResourceUriRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	d.SetId(addResourceUriRes.GetData()["uid"].(string))

	return readManagementResourceUri(d, m)
}

func readManagementResourceUri(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showResourceUriRes, err := client.ApiCall("show-resource-uri", payload, client.GetSessionID(), true, false)
	if err != nil {
		return fmt.Errorf(err.Error())
	}
	if !showResourceUriRes.Success {
		if objectNotFound(showResourceUriRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf(showResourceUriRes.ErrorMsg)
	}

	resourceUri := showResourceUriRes.GetData()

	log.Println("Read ResourceUri - Show JSON = ", resourceUri)

	if v := resourceUri["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := resourceUri["use-this-resource-to"]; v != nil {
		_ = d.Set("use_this_resource_to", v)
	}

	if resourceUri["connection-methods"] != nil {

		connectionMethodsMap := resourceUri["connection-methods"].(map[string]interface{})

		connectionMethodsMapToReturn := make(map[string]interface{})

		if v, _ := connectionMethodsMap["transparent"]; v != nil {
			connectionMethodsMapToReturn["transparent"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := connectionMethodsMap["proxy"]; v != nil {
			connectionMethodsMapToReturn["proxy"] = strconv.FormatBool(v.(bool))
		}
		if v, _ := connectionMethodsMap["tunneling"]; v != nil {
			connectionMethodsMapToReturn["tunneling"] = strconv.FormatBool(v.(bool))
		}
		_ = d.Set("connection_methods", connectionMethodsMapToReturn)
	} else {
		_ = d.Set("connection_methods", nil)
	}

	if v := resourceUri["uri-match-specification-type"]; v != nil {
		_ = d.Set("uri_match_specification_type", v)
	}

	if v := resourceUri["exception-track"]; v != nil {

		objMap := v.(map[string]interface{})

		if v := objMap["name"]; v != nil {
			_ = d.Set("exception_track", v)
		}

	}

	if resourceUri["match-ufp"] != nil {

		matchUfpMap := resourceUri["match-ufp"].(map[string]interface{})

		matchUfpMapToReturn := make(map[string]interface{})

		if v, _ := matchUfpMap["server"]; v != nil {
			matchUfpMapToReturn["server"] = v
		}
		if v, _ := matchUfpMap["caching-control"]; v != nil {
			matchUfpMapToReturn["caching_control"] = v
		}
		if v, _ := matchUfpMap["ignore-ufp-server-after-failure"]; v != nil {
			matchUfpMapToReturn["ignore_ufp_server_after_failure"] = v
		}
		if v, _ := matchUfpMap["number-of-failures-before-ignore"]; v != nil {
			matchUfpMapToReturn["number_of_failures_before_ignore"] = v
		}
		if v, _ := matchUfpMap["timeout-before-reconnecting"]; v != nil {
			matchUfpMapToReturn["timeout_before_reconnecting"] = v
		}
		_ = d.Set("match_ufp", []interface{}{matchUfpMapToReturn})
	} else {
		_ = d.Set("match_ufp", nil)
	}

	if resourceUri["match-wildcards"] != nil {

		matchWildcardsMap, ok := resourceUri["match-wildcards"].(map[string]interface{})

		if ok {
			matchWildcardsMapToReturn := make(map[string]interface{})

			if v, ok := matchWildcardsMap["schemes"]; ok {

				schemesMap, ok := v.(map[string]interface{})
				if ok {
					schemesMapToReturn := make(map[string]interface{})

					if v, _ := schemesMap["http"]; v != nil {
						schemesMapToReturn["http"] = v
					}
					if v, _ := schemesMap["ftp"]; v != nil {
						schemesMapToReturn["ftp"] = v
					}
					if v, _ := schemesMap["gopher"]; v != nil {
						schemesMapToReturn["gopher"] = v
					}
					if v, _ := schemesMap["mailto"]; v != nil {
						schemesMapToReturn["mailto"] = v
					}
					if v, _ := schemesMap["news"]; v != nil {
						schemesMapToReturn["news"] = v
					}
					if v, _ := schemesMap["wais"]; v != nil {
						schemesMapToReturn["wais"] = v
					}
					if v, _ := schemesMap["other"]; v != nil {
						schemesMapToReturn["other"] = v
					}
					matchWildcardsMapToReturn["schemes"] = []interface{}{schemesMapToReturn}
				}
			}
			if v, ok := matchWildcardsMap["methods"]; ok {

				methodsMap, ok := v.(map[string]interface{})
				if ok {
					methodsMapToReturn := make(map[string]interface{})

					if v, _ := methodsMap["get"]; v != nil {
						methodsMapToReturn["get"] = v
					}
					if v, _ := methodsMap["post"]; v != nil {
						methodsMapToReturn["post"] = v
					}
					if v, _ := methodsMap["head"]; v != nil {
						methodsMapToReturn["head"] = v
					}
					if v, _ := methodsMap["put"]; v != nil {
						methodsMapToReturn["put"] = v
					}
					if v, _ := methodsMap["other"]; v != nil {
						methodsMapToReturn["other"] = v
					}
					matchWildcardsMapToReturn["methods"] = []interface{}{methodsMapToReturn}
				}
			}
			if v := matchWildcardsMap["host"]; v != nil {
				matchWildcardsMapToReturn["host"] = v
			}
			if v := matchWildcardsMap["path"]; v != nil {
				matchWildcardsMapToReturn["path"] = v
			}
			if v := matchWildcardsMap["query"]; v != nil {
				matchWildcardsMapToReturn["query"] = v
			}
			_ = d.Set("match_wildcards", []interface{}{matchWildcardsMapToReturn})

		}
	} else {
		_ = d.Set("match_wildcards", nil)
	}

	if resourceUri["action"] != nil {

		actionMap := resourceUri["action"].(map[string]interface{})

		actionMapToReturn := make(map[string]interface{})

		if v, _ := actionMap["replacement-uri"]; v != nil {
			actionMapToReturn["replacement_uri"] = v
		}
		if v, _ := actionMap["strip-script-tags"]; v != nil {
			actionMapToReturn["strip_script_tags"] = v
		}
		if v, _ := actionMap["strip-applet-tags"]; v != nil {
			actionMapToReturn["strip_applet_tags"] = v
		}
		if v, _ := actionMap["strip-activex-tags"]; v != nil {
			actionMapToReturn["strip_activex_tags"] = v
		}
		if v, _ := actionMap["strip-ftp-links"]; v != nil {
			actionMapToReturn["strip_ftp_links"] = v
		}
		if v, _ := actionMap["strip-port-strings"]; v != nil {
			actionMapToReturn["strip_port_strings"] = v
		}
		_ = d.Set("action", []interface{}{actionMapToReturn})
	} else {
		_ = d.Set("action", nil)
	}

	if resourceUri["cvp"] != nil {

		cvpMap := resourceUri["cvp"].(map[string]interface{})

		cvpMapToReturn := make(map[string]interface{})

		if v, _ := cvpMap["enable-cvp"]; v != nil {
			cvpMapToReturn["enable_cvp"] = v
		}
		if v, _ := cvpMap["server"]; v != nil {
			cvpMapToReturn["server"] = v
		}
		if v, _ := cvpMap["cvp-server-is-allowed-to-modify-content"]; v != nil {
			cvpMapToReturn["allowed_to_modify_content"] = v
		}
		if v, _ := cvpMap["send-http-headers-to-cvp"]; v != nil {
			cvpMapToReturn["send_http_headers_to_cvp"] = v
		}
		if v, _ := cvpMap["reply-order"]; v != nil {
			cvpMapToReturn["reply_order"] = v
		}
		if v, _ := cvpMap["send-http-request-to-cvp"]; v != nil {
			cvpMapToReturn["send_http_request_to_cvp"] = v
		}
		if v, _ := cvpMap["send-only-unsafe-file-types"]; v != nil {
			cvpMapToReturn["send_only_unsafe_file_types"] = v
		}
		_ = d.Set("cvp", []interface{}{cvpMapToReturn})
	} else {
		_ = d.Set("cvp", nil)
	}

	if resourceUri["soap"] != nil {

		soapMap := resourceUri["soap"].(map[string]interface{})

		soapMapToReturn := make(map[string]interface{})

		if v, _ := soapMap["inspection"]; v != nil {
			soapMapToReturn["inspection"] = v
		}
		if v, _ := soapMap["file-id"]; v != nil {
			soapMapToReturn["file_id"] = v
		}
		if v, _ := soapMap["track-connections"]; v != nil {
			soapMapToReturn["track_connections"] = v
		}
		_ = d.Set("soap", soapMapToReturn)
	} else {
		_ = d.Set("soap", nil)
	}

	if resourceUri["tags"] != nil {
		tagsJson, ok := resourceUri["tags"].([]interface{})
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

	if v := resourceUri["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := resourceUri["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if v := resourceUri["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := resourceUri["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	return nil

}

func updateManagementResourceUri(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	resourceUri := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		resourceUri["name"] = oldName
		resourceUri["new-name"] = newName
	} else {
		resourceUri["name"] = d.Get("name")
	}

	if ok := d.HasChange("use_this_resource_to"); ok {
		resourceUri["use-this-resource-to"] = d.Get("use_this_resource_to")
	}

	if d.HasChange("connection_methods") {

		if _, ok := d.GetOk("connection_methods"); ok {

			res := make(map[string]interface{})

			if d.HasChange("connection_methods.transparent") {
				res["transparent"] = d.Get("connection_methods.transparent")
			}
			if d.HasChange("connection_methods.proxy") {
				res["proxy"] = d.Get("connection_methods.proxy")
			}
			if d.HasChange("connection_methods.tunneling") {
				res["tunneling"] = d.Get("connection_methods.tunneling")
			}
			resourceUri["connection-methods"] = res
		}
	}

	if ok := d.HasChange("uri_match_specification_type"); ok {
		resourceUri["uri-match-specification-type"] = d.Get("uri_match_specification_type")
	}

	if ok := d.HasChange("exception_track"); ok {
		resourceUri["exception-track"] = d.Get("exception_track")
	}

	if d.HasChange("match_ufp") {

		if _, ok := d.GetOk("match_ufp"); ok {

			res := make(map[string]interface{})

			if v := d.Get("match_ufp.0.server"); v != nil {
				res["server"] = v
			}
			if v := d.Get("match_ufp.0.caching_control"); v != nil {
				res["server"] = v
			}
			if v := d.Get("match_ufp.0.ignore_ufp_server_after_failure"); v != nil {
				res["server"] = v
			}
			if v := d.Get("match_ufp.0.number_of_failures_before_ignore"); v != nil {
				res["server"] = v
			}
			if v := d.Get("match_ufp.0.timeout_before_reconnecting"); v != nil {
				res["server"] = v
			}

			resourceUri["match-ufp"] = res
		}
	}

	if d.HasChange("match_wildcards") {

		if v, ok := d.GetOk("match_wildcards"); ok {

			matchWildcardsList := v.([]interface{})

			if len(matchWildcardsList) > 0 {

				matchWildcardsPayload := make(map[string]interface{})

				if d.HasChange("match_wildcards.0.schemes") {

					schemesPayload := make(map[string]interface{})

					if d.HasChange("match_wildcards.0.schemes.0.http") {
						schemesPayload["http"] = d.Get("match_wildcards.0.schemes.0.http")
					}
					if d.HasChange("match_wildcards.0.schemes.0.ftp") {
						schemesPayload["ftp"] = d.Get("match_wildcards.0.schemes.0.ftp")
					}
					if d.HasChange("match_wildcards.0.schemes.0.gopher") {
						schemesPayload["gopher"] = d.Get("match_wildcards.0.schemes.0.gopher")
					}
					if d.HasChange("match_wildcards.0.schemes.0.mailto") {
						schemesPayload["mailto"] = d.Get("match_wildcards.0.schemes.0.mailto")
					}
					if d.HasChange("match_wildcards.0.schemes.0.news") {
						schemesPayload["news"] = d.Get("match_wildcards.0.schemes.0.news")
					}
					if d.HasChange("match_wildcards.0.schemes.0.wais") {
						schemesPayload["wais"] = d.Get("match_wildcards.0.schemes.0.wais")
					}
					if d.HasChange("match_wildcards.0.schemes.0.other") {
						schemesPayload["other"] = d.Get("match_wildcards.0.schemes.0.other").(string)
					}
					matchWildcardsPayload["schemes"] = schemesPayload
				}
				if d.HasChange("match_wildcards.0.methods") {

					methodsPayload := make(map[string]interface{})

					if d.HasChange("match_wildcards.0.methods.0.get") {
						methodsPayload["get"] = d.Get("match_wildcards.0.methods.0.get")
					}
					if d.HasChange("match_wildcards.0.methods.0.post") {
						methodsPayload["post"] = d.Get("match_wildcards.0.methods.0.post")
					}
					if d.HasChange("match_wildcards.0.methods.0.head") {
						methodsPayload["head"] = d.Get("match_wildcards.0.methods.0.head")
					}
					if d.HasChange("match_wildcards.0.methods.0.put") {
						methodsPayload["put"] = d.Get("match_wildcards.0.methods.0.put")
					}
					if d.HasChange("match_wildcards.0.methods.0.other") {
						methodsPayload["other"] = d.Get("match_wildcards.0.methods.0.other").(string)
					}
					matchWildcardsPayload["methods"] = methodsPayload
				}
				if d.HasChange("match_wildcards.0.host") {
					matchWildcardsPayload["host"] = d.Get("match_wildcards.0.host").(string)
				}
				if d.HasChange("match_wildcards.0.path") {
					matchWildcardsPayload["path"] = d.Get("match_wildcards.0.path").(string)
				}
				if d.HasChange("match_wildcards.0.query") {
					matchWildcardsPayload["query"] = d.Get("match_wildcards.0.query").(string)
				}
				resourceUri["match-wildcards"] = matchWildcardsPayload
			}
		}
	}

	if d.HasChange("action") {

		if _, ok := d.GetOk("action"); ok {

			res := make(map[string]interface{})

			if v := d.Get("action.0.replacement_uri"); v != nil {
				res["replacement-uri"] = v
			}
			if v := d.Get("action.0.strip_script_tags"); v != nil {
				res["strip-script-tags"] = v
			}
			if v := d.Get("action.0.strip_applet_tags"); v != nil {
				res["strip-applet-tags"] = v
			}
			if v := d.Get("action.0.strip_activex_tags"); v != nil {
				res["strip-activex-tags"] = v
			}
			if v := d.Get("action.0.strip_ftp_links"); v != nil {
				res["strip-ftp-links"] = v
			}
			if v := d.Get("action.0.strip_port_strings"); v != nil {
				res["strip-port-strings"] = v
			}

			resourceUri["action"] = res
		}
	}

	if d.HasChange("cvp") {

		if _, ok := d.GetOk("cvp"); ok {

			res := make(map[string]interface{})

			if v := d.Get("cvp.0.enable_cvp"); v != nil {
				res["enable-cvp"] = v
			}
			if v := d.Get("cvp.0.server"); v != nil {
				if len(v.(string)) > 0 {
					res["server"] = v
				}
			}
			if v := d.Get("cvp.0.allowed_to_modify_content"); v != nil {
				res["allowed-to-modify-content"] = v
			}

			if v := d.Get("cvp.0.send_http_headers_to_cvp"); v != nil {
				res["send-http-headers-to-cvp"] = v
			}
			if v := d.Get("cvp.0.reply_order"); v != nil {
				res["reply-order"] = v
			}
			if v := d.Get("cvp.0.send_http_request_to_cvp"); v != nil {
				res["send-http-request-to-cvp"] = v
			}
			if v := d.Get("cvp.0.send_only_unsafe_file_types"); v != nil {
				res["send-only-unsafe-file-types"] = v
			}
			resourceUri["cvp"] = res
		}
	}

	if d.HasChange("soap") {

		if _, ok := d.GetOk("soap"); ok {

			res := make(map[string]interface{})

			if d.HasChange("soap.inspection") {
				res["inspection"] = d.Get("soap.inspection")
			}
			if d.HasChange("soap.file_id") {
				res["file-id"] = d.Get("soap.file_id")
			}
			if d.HasChange("soap.track_connections") {
				res["track-connections"] = d.Get("soap.track_connections")
			}
			resourceUri["soap"] = res
		}
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			resourceUri["tags"] = v.(*schema.Set).List()
		} else {
			oldTags, _ := d.GetChange("tags")
			resourceUri["tags"] = map[string]interface{}{"remove": oldTags.(*schema.Set).List()}
		}
	}

	if ok := d.HasChange("color"); ok {
		resourceUri["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		resourceUri["comments"] = d.Get("comments")
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		resourceUri["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		resourceUri["ignore-errors"] = v.(bool)
	}

	log.Println("Update ResourceUri - Map = ", resourceUri)

	updateResourceUriRes, err := client.ApiCall("set-resource-uri", resourceUri, client.GetSessionID(), true, false)
	if err != nil || !updateResourceUriRes.Success {
		if updateResourceUriRes.ErrorMsg != "" {
			return fmt.Errorf(updateResourceUriRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}

	return readManagementResourceUri(d, m)
}

func deleteManagementResourceUri(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	resourceUriPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete ResourceUri")

	deleteResourceUriRes, err := client.ApiCall("delete-resource-uri", resourceUriPayload, client.GetSessionID(), true, false)
	if err != nil || !deleteResourceUriRes.Success {
		if deleteResourceUriRes.ErrorMsg != "" {
			return fmt.Errorf(deleteResourceUriRes.ErrorMsg)
		}
		return fmt.Errorf(err.Error())
	}
	d.SetId("")

	return nil
}
