package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strconv"
)

func dataDourceManagementResourceUri() *schema.Resource {
	return &schema.Resource{

		Read: dataSourceManagementResourceUriRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"uid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Object name.",
			},
			"use_this_resource_to": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Select the use of the URI resource.",
			},
			"connection_methods": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Connection methods.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"transparent": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The security server is invisible to the client that originates the connection, and to the server. The Transparent connection method is the most secure.",
						},
						"proxy": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The Resource is applied when people specify the Check Point Security Gateway as a proxy in their browser.",
						},
						"tunneling": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The Resource is applied when people specify the Security Gateway as a proxy in their browser, and is used for connections where Security Gateway cannot examine the contents of the packets, not even the URL.",
						},
					},
				},
			},
			"uri_match_specification_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type can be Wild Cards or UFP, where a UFP server holds categories of forbidden web sites.",
			},
			"exception_track": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Configures how to track connections that match this rule but fail the content security checks. An example of an exception is a connection with an unsupported scheme or method.",
			},
			"match_ufp": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Match - UFP settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"server": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UID or Name of the UFP server that is an OPSEC certified third party application that checks URLs against a list of permitted categories.",
						},
						"caching_control": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies if and how caching is to be enabled.",
						},
						"ignore_ufp_server_after_failure": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The UFP server will be ignored after numerous UFP server connections were unsuccessful.",
						},
						"number_of_failures_before_ignore": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Signifies at what point the UFP server should be ignored.",
						},
						"timeout_before_reconnecting": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The amount of time that must pass before a UFP server connection should be attempted.",
						},
					},
				},
			},
			"match_wildcards": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Match - Wildcards settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"schemes": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Select the URI Schemes to which this resource applies.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"http": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Http scheme.",
									},
									"ftp": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Ftp scheme.",
									},
									"gopher": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Gopher scheme.",
									},
									"mailto": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Mailto scheme.",
									},
									"news": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "News scheme.",
									},
									"wais": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Wais scheme.",
									},
									"other": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "You can specify another scheme in the Other field. You can use wildcards.",
									},
								},
							},
						},
						"methods": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Select the URI Schemes to which this resource applies.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"get": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "GET method.",
									},
									"post": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "POST method.",
									},
									"head": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "HEAD method.",
									},
									"put": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "PUT method.",
									},
									"other": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "You can specify another method in the Other field. You can use wildcards.",
									},
								},
							},
						},
						"host": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The functionality of the Host parameter depends on the DNS setup of the addressed server. For the host, only the IP address or the full DNS name should be used.",
						},
						"path": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name matching is based on appending the file name in the request to the current working directory (unless the file name is already a full path name) and comparing the result to the path specified in the Resource definition.",
						},
						"query": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The parameters that are sent to the URI when it is accessed.",
						},
					},
				},
			},
			"action": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Action settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"replacement_uri": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "If the Action in a rule which uses this resource is Drop or Reject, then the Replacement URI is displayed instead of the one requested by the user.",
						},
						"strip_script_tags": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip JAVA scripts.",
						},
						"strip_applet_tags": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip JAVA applets.",
						},
						"strip_activex_tags": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip activeX tags.",
						},
						"strip_ftp_links": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip ftp links.",
						},
						"strip_port_strings": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Strip ports.",
						},
					},
				},
			},
			"cvp": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "CVP settings.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enable_cvp": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select to enable the Content Vectoring Protocol.",
						},
						"server": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application.",
						},
						"allowed_to_modify_content": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Configures the CVP server to inspect but not modify content.",
						},
						"send_http_headers_to_cvp": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Select, if you would like the CVP server to check the HTTP headers of the message packets.",
						},
						"reply_order": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Designates when the CVP server returns data to the Security Gateway security server.",
						},
						"send_http_request_to_cvp": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Used to protect against undesirable content in the HTTP request, for example, when inspecting peer-to-peer connections.",
						},
						"send_only_unsafe_file_types": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Improves the performance of the CVP server. This option does not send to the CVP server traffic that is considered safe.",
						},
					},
				},
			},
			"soap": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "SOAP settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"inspection": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Allow all SOAP Requests, or Allow only SOAP requests specified in the following file-id.",
						},
						"file_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A file containing SOAP requests.",
						},
						"track_connections": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The method of tracking SOAP connections.",
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
		},
	}
}

func dataSourceManagementResourceUriRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
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

	if v := resourceUri["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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

	return nil

}
