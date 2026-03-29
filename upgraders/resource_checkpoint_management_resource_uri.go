package upgraders

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// ResourceManagementResourceUriV0 is the V0 schema where connection_methods and soap were TypeMap.
func ResourceManagementResourceUriV0() *schema.Resource {
	return &schema.Resource{
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

// ResourceManagementResourceUriStateUpgradeV0 converts connection_methods and soap from TypeMap to TypeList.
func ResourceManagementResourceUriStateUpgradeV0(_ context.Context, rawState map[string]interface{}, _ interface{}) (map[string]interface{}, error) {
	return UpgradeMapsToLists(rawState, "connection_methods", "soap"), nil
}
