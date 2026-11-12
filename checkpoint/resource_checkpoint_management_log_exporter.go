package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
	"strconv"
)

func resourceManagementLogExporter() *schema.Resource {
	return &schema.Resource{
		Create: createManagementLogExporter,
		Read:   readManagementLogExporter,
		Update: updateManagementLogExporter,
		Delete: deleteManagementLogExporter,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Object name.",
			},
			"target_server": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Target server port to which logs are exported.",
			},
			"target_port": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Port number of the target server.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Protocol used to send logs to the target server.",
				Default:     "udp",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Indicates whether to enable export.",
				Default:     true,
			},
			"attachments": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Log exporter attachments.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_link_to_log_attachment": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to add link to log attachment in SmartView.",
							Default:     false,
						},
						"add_link_to_log_details": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to add link to log details in SmartView.",
							Default:     false,
						},
						"add_log_attachment_id": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to add log attachment ID.",
							Default:     false,
						},
					},
				},
			},
			"data_manipulation": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Log exporter data manipulation.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregate_log_updates": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to aggregate log updates.",
						},
						"format": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Logs format.",
						},
						"add_custom_log_header": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Custom header added to every exported log entry.<br> This header can be used to identify the source or add metadata to logs.<br> The custom header is prepended",
						},
						"aggregation_mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Aggregation mode for aggregated logs.<br> Determines how log updates are processed when aggregate-log-updates is true or not specified:<br> - semi-unified (defa",
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
			"advanced": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Advanced settings for log exporter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_in_milliseconds": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Toggle export logs' time with ms resolution.",
						},
					},
				},
			},
			"tls": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "TLS settings for log exporter.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_generate_client_certificates": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "TLS mode where user provides server CA certificate only. Check Point Management automatically generates and signs client certificates. This is the sim...",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_ca_cert": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client generated CA certificate that signed the client certificate. The server should trust this CA. Supplied as <b>base64 encoded string</b>.",
									},
									"server_ca_cert": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Server certificate authority certificate, supplied as <b>base64 encoded string</b>.",
									},
								},
							},
						},
						"client_cert_expiration_date": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Client certificate expiration date extracted from the stored PKCS#12 certificate. Available for all TLS modes when client certificate exists.",
						},
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Indicates whether to use encrypted connection. Can be true only when protocol is TCP.",
						},
						"generate_and_sign_client_csr": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Advanced TLS mode with two separate phases: 'auto-generate-client-csr' to initiate CSR generation (ADD/SET) and 'provide-signed-client-cert' to provid...",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_cert": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Client certificate in PEM format as <b>base64 encoded string</b>, signed by external CA. Can only be provided when phase is 'provide-signed-client-cert'.",
									},
									"client_csr": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client certificate signing request (CSR) in PEM format as <b>base64 encoded string</b>. Present when phase is 'auto-generate-client-csr'.",
									},
									"phase": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Mode phase type: 'auto-generate-client-csr' to initiate CSR generation.",
									},
									"server_ca_cert": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Server CA certificate in PEM format as <b>base64 encoded string</b> to validate the destination server's identity. Required when phase is 'auto-genera...",
									},
								},
							},
						},
						"mode": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "TLS mode. Must match the specific mode configuration provided.",
						},
						"user_provided_certificates": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "TLS mode where user provides server CA certificate, client certificate (PKCS#12 format), and passphrase. All certificates provided by the user, Check ...",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_cert": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Client certificate, supplied as <b>base64 encoded string</b>.",
									},
									"client_secret": {
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
										Description: "Client secret for TLS connection.",
									},
									"server_ca_cert": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Server certificate authority certificate, supplied as <b>base64 encoded string</b>.",
									},
								},
							},
						},
					},
				},
			},
			"filters": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Log exporter filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expressions": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of filter expressions. Each expression targets a specific log field.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"conditions": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Filter conditions for this field. Required when 'operator' is 'and' or 'or'.<br>Array of condition objects, each with 'operator' and 'value'.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"operator": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "Comparison operator.",
												},
												"value": {
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The value to compare against.",
												},
											},
										},
									},
									"field_name": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Log field name to filter on.",
									},
									"operator": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Operator for this expression.<br>Use 'and'/'or' with 'conditions' (array of condition objects).<br>Use 'in'/'not-in' with 'values' (flat array of stri...",
									},
									"values": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Filter values for this field. Present when 'operator' is 'in' or 'not-in'.<br>Flat array of string values.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"expressions_operator": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Logical operator for combining all filter expressions.",
						},
						"filter_out_connection_logs": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "When set to true, connection logs will be excluded from export.",
						},
					},
				},
			},
			"fields": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				Description: "Log exporter fields mapping settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclusion_list": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "List of field names to exclude or include based on the export mode. <br> When export is <b>all-except</b>: fields in this list are excluded from expor...",
						},
						"export": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Field export mode. <br><b>all-except</b> (default): exports all fields except those in the exclusion list. <br><b>nothing-except</b>: exports nothing ...",
						},
						"names_mapping": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Field name mappings to rename fields during export. <br> Specify mappings in the format: 'orgName:dstName' (e.g., 'src:source_ip,dst:destination_ip')....",
						},
					},
				},
			},
		},
	}
}

func createManagementLogExporter(d *schema.ResourceData, m interface{}) error {
	client := m.(*checkpoint.ApiClient)

	logExporter := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		logExporter["name"] = v.(string)
	}

	if v, ok := d.GetOk("target_server"); ok {
		logExporter["target-server"] = v.(string)
	}

	if v, ok := d.GetOk("target_port"); ok {
		logExporter["target-port"] = v.(int)
	}

	if v, ok := d.GetOk("protocol"); ok {
		logExporter["protocol"] = v.(string)
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		logExporter["enabled"] = v.(bool)
	}

	if _, ok := d.GetOk("attachments"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOkExists("attachments.0.add_link_to_log_attachment"); ok {
			res["add-link-to-log-attachment"] = v.(bool)
		}
		if v, ok := d.GetOkExists("attachments.0.add_link_to_log_details"); ok {
			res["add-link-to-log-details"] = v.(bool)
		}
		if v, ok := d.GetOkExists("attachments.0.add_log_attachment_id"); ok {
			res["add-log-attachment-id"] = v.(bool)
		}

		logExporter["attachments"] = res
	}

	if _, ok := d.GetOk("data_manipulation"); ok {

		res := make(map[string]interface{})

		if v, ok := d.GetOkExists("data_manipulation.0.aggregate_log_updates"); ok {
			res["aggregate-log-updates"] = v.(bool)
		}
		if v, ok := d.GetOk("data_manipulation.0.format"); ok {
			res["format"] = v.(string)
		}

		if v, ok := d.GetOk("data_manipulation.0.add_custom_log_header"); ok {
			res["add-custom-log-header"] = v.(string)
		}
		if v, ok := d.GetOk("data_manipulation.0.aggregation_mode"); ok {
			res["aggregation-mode"] = v.(string)
		}
		logExporter["data-manipulation"] = res
	}

	if v, ok := d.GetOk("color"); ok {
		logExporter["color"] = v.(string)
	}

	if v, ok := d.GetOk("comments"); ok {
		logExporter["comments"] = v.(string)
	}

	if v, ok := d.GetOk("tags"); ok {
		logExporter["tags"] = v.(*schema.Set).List()
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		logExporter["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		logExporter["ignore-errors"] = v.(bool)
	}

	log.Println("Create LogExporter - Map = ", logExporter)

	if _, ok := d.GetOk("tls"); ok {
		tlsPayload := make(map[string]interface{})
		if _, ok := d.GetOk("tls.0.auto_generate_client_certificates"); ok {
			autoGenerateClientCertificatesPayload := make(map[string]interface{})
			if v, ok := d.GetOk("tls.0.auto_generate_client_certificates.0.server_ca_cert"); ok {
				autoGenerateClientCertificatesPayload["server-ca-cert"] = v.(string)
			}
			tlsPayload["auto-generate-client-certificates"] = autoGenerateClientCertificatesPayload
		}
		if v, ok := d.GetOkExists("tls.0.enabled"); ok {
			tlsPayload["enabled"] = v.(bool)
		}
		if _, ok := d.GetOk("tls.0.generate_and_sign_client_csr"); ok {
			generateAndSignClientCsrPayload := make(map[string]interface{})
			if v, ok := d.GetOk("tls.0.generate_and_sign_client_csr.0.phase"); ok {
				generateAndSignClientCsrPayload["phase"] = v.(string)
			}
			if v, ok := d.GetOk("tls.0.generate_and_sign_client_csr.0.server_ca_cert"); ok {
				generateAndSignClientCsrPayload["server-ca-cert"] = v.(string)
			}
			tlsPayload["generate-and-sign-client-csr"] = generateAndSignClientCsrPayload
		}
		if v, ok := d.GetOk("tls.0.mode"); ok {
			tlsPayload["mode"] = v.(string)
		}
		if _, ok := d.GetOk("tls.0.user_provided_certificates"); ok {
			userProvidedCertificatesPayload := make(map[string]interface{})
			if v, ok := d.GetOk("tls.0.user_provided_certificates.0.client_cert"); ok {
				userProvidedCertificatesPayload["client-cert"] = v.(string)
			}
			if v, ok := d.GetOk("tls.0.user_provided_certificates.0.client_secret"); ok {
				userProvidedCertificatesPayload["client-secret"] = v.(string)
			}
			if v, ok := d.GetOk("tls.0.user_provided_certificates.0.server_ca_cert"); ok {
				userProvidedCertificatesPayload["server-ca-cert"] = v.(string)
			}
			tlsPayload["user-provided-certificates"] = userProvidedCertificatesPayload
		}
		logExporter["tls"] = tlsPayload
	}

	if _, ok := d.GetOk("filters"); ok {
		filtersPayload := make(map[string]interface{})
		if v, ok := d.GetOk("filters.0.expressions"); ok {
			expressionsList := v.([]interface{})
			if len(expressionsList) > 0 {
				var expressionsPayload []map[string]interface{}
				for i := range expressionsList {
					expressionsItem := make(map[string]interface{})
					if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".conditions"); ok {
						conditionsList := v.([]interface{})
						if len(conditionsList) > 0 {
							var conditionsPayload []map[string]interface{}
							for j := range conditionsList {
								conditionsItem := make(map[string]interface{})
								if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".conditions." + strconv.Itoa(j) + ".operator"); ok {
									conditionsItem["operator"] = v.(string)
								}
								if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".conditions." + strconv.Itoa(j) + ".value"); ok {
									conditionsItem["value"] = v.(string)
								}
								conditionsPayload = append(conditionsPayload, conditionsItem)
							}
							expressionsItem["conditions"] = conditionsPayload
						}
					}
					if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".field_name"); ok {
						expressionsItem["field-name"] = v.(string)
					}
					if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".operator"); ok {
						expressionsItem["operator"] = v.(string)
					}
					expressionsPayload = append(expressionsPayload, expressionsItem)
				}
				filtersPayload["expressions"] = expressionsPayload
			}
		}
		if v, ok := d.GetOk("filters.0.expressions_operator"); ok {
			filtersPayload["expressions-operator"] = v.(string)
		}
		if v, ok := d.GetOkExists("filters.0.filter_out_connection_logs"); ok {
			filtersPayload["filter-out-connection-logs"] = v.(bool)
		}
		logExporter["filters"] = filtersPayload
	}

	if _, ok := d.GetOk("fields"); ok {
		fieldsPayload := make(map[string]interface{})
		if v, ok := d.GetOk("fields.0.exclusion_list"); ok {
			fieldsPayload["exclusion-list"] = v.(string)
		}
		if v, ok := d.GetOk("fields.0.export"); ok {
			fieldsPayload["export"] = v.(string)
		}
		if v, ok := d.GetOk("fields.0.names_mapping"); ok {
			fieldsPayload["names-mapping"] = v.(string)
		}
		logExporter["fields"] = fieldsPayload
	}

	addLogExporterRes, err := client.ApiCallSimple("add-log-exporter", logExporter)
	if err != nil || !addLogExporterRes.Success {
		if addLogExporterRes.ErrorMsg != "" {
			return fmt.Errorf("%s", addLogExporterRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}

	d.SetId(addLogExporterRes.GetData()["uid"].(string))

	return readManagementLogExporter(d, m)
}

func readManagementLogExporter(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	payload := map[string]interface{}{
		"uid": d.Id(),
	}

	showLogExporterRes, err := client.ApiCallSimple("show-log-exporter", payload)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showLogExporterRes.Success {
		if objectNotFound(showLogExporterRes.GetData()["code"].(string)) {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("%s", showLogExporterRes.ErrorMsg)
	}

	logExporter := showLogExporterRes.GetData()

	log.Println("Read LogExporter - Show JSON = ", logExporter)

	if v := logExporter["name"]; v != nil {
		_ = d.Set("name", v)
	}

	if v := logExporter["target-server"]; v != nil {
		_ = d.Set("target_server", v)
	}

	if v := logExporter["target-port"]; v != nil {
		_ = d.Set("target_port", v)
	}

	if v := logExporter["protocol"]; v != nil {
		_ = d.Set("protocol", v)
	}

	if v := logExporter["enabled"]; v != nil {
		_ = d.Set("enabled", v)
	}

	if logExporter["attachments"] != nil {

		attachmentsMap := logExporter["attachments"].(map[string]interface{})

		attachmentsMapToReturn := make(map[string]interface{})

		if v, _ := attachmentsMap["add-link-to-log-attachment"]; v != nil {
			attachmentsMapToReturn["add_link_to_log_attachment"] = v
		}
		if v, _ := attachmentsMap["add-link-to-log-details"]; v != nil {
			attachmentsMapToReturn["add_link_to_log_details"] = v
		}
		if v, _ := attachmentsMap["add-log-attachment-id"]; v != nil {
			attachmentsMapToReturn["add_log_attachment_id"] = v
		}

		_ = d.Set("attachments", []interface{}{attachmentsMapToReturn})
	} else {
		_ = d.Set("attachments", nil)
	}

	if logExporter["data-manipulation"] != nil {

		dataManipulationMap := logExporter["data-manipulation"].(map[string]interface{})

		dataManipulationMapToReturn := make(map[string]interface{})

		if v, _ := dataManipulationMap["aggregate-log-updates"]; v != nil {
			dataManipulationMapToReturn["aggregate_log_updates"] = v
		}
		if v, _ := dataManipulationMap["format"]; v != nil {
			dataManipulationMapToReturn["format"] = v
		}

		if v := dataManipulationMap["add-custom-log-header"]; v != nil {
			dataManipulationMapToReturn["add_custom_log_header"] = v
		}
		if v := dataManipulationMap["aggregation-mode"]; v != nil {
			dataManipulationMapToReturn["aggregation_mode"] = v
		}
		_ = d.Set("data_manipulation", []interface{}{dataManipulationMapToReturn})
	} else {
		_ = d.Set("data_manipulation", nil)
	}

	if v := logExporter["color"]; v != nil {
		_ = d.Set("color", v)
	}

	if v := logExporter["comments"]; v != nil {
		_ = d.Set("comments", v)
	}

	if logExporter["tags"] != nil {
		tagsJson, ok := logExporter["tags"].([]interface{})
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

	if v := logExporter["ignore-warnings"]; v != nil {
		_ = d.Set("ignore_warnings", v)
	}

	if v := logExporter["ignore-errors"]; v != nil {
		_ = d.Set("ignore_errors", v)
	}

	if v := logExporter["advanced"]; v != nil {
		advancedJson := v.(map[string]interface{})
		advancedState := make(map[string]interface{})
		if v := advancedJson["time-in-milliseconds"]; v != nil {
			advancedState["time_in_milliseconds"] = v
		}
		_ = d.Set("advanced", []interface{}{advancedState})
	}

	if v := logExporter["tls"]; v != nil {
		tlsShow := v.(map[string]interface{})
		tlsState := make(map[string]interface{})
		if v := tlsShow["auto-generate-client-certificates"]; v != nil {
			autoGenerateClientCertificatesShow := v.(map[string]interface{})
			autoGenerateClientCertificatesState := make(map[string]interface{})
			if v := autoGenerateClientCertificatesShow["client-ca-cert"]; v != nil {
				autoGenerateClientCertificatesState["client_ca_cert"] = v
			}
			tlsState["auto_generate_client_certificates"] = []interface{}{autoGenerateClientCertificatesState}
		}
		if v := tlsShow["client-cert-expiration-date"]; v != nil {
			tlsState["client_cert_expiration_date"] = v
		}
		if v := tlsShow["enabled"]; v != nil {
			tlsState["enabled"] = v
		}
		if v := tlsShow["generate-and-sign-client-csr"]; v != nil {
			generateAndSignClientCsrShow := v.(map[string]interface{})
			generateAndSignClientCsrState := make(map[string]interface{})
			if v := generateAndSignClientCsrShow["client-csr"]; v != nil {
				generateAndSignClientCsrState["client_csr"] = v
			}
			if v := generateAndSignClientCsrShow["phase"]; v != nil {
				generateAndSignClientCsrState["phase"] = v
			}
			tlsState["generate_and_sign_client_csr"] = []interface{}{generateAndSignClientCsrState}
		}
		if v := tlsShow["mode"]; v != nil {
			tlsState["mode"] = v
		}
		_ = d.Set("tls", []interface{}{tlsState})
	}

	if v := logExporter["filters"]; v != nil {
		filtersShow := v.(map[string]interface{})
		filtersState := make(map[string]interface{})
		if v := filtersShow["expressions"]; v != nil {
			expressionsList := v.([]interface{})
			var expressionsListState []map[string]interface{}
			for i := range expressionsList {
				expressionsShow := expressionsList[i].(map[string]interface{})
				expressionsState := make(map[string]interface{})
				if v := expressionsShow["conditions"]; v != nil {
					conditionsList := v.([]interface{})
					var conditionsListState []map[string]interface{}
					for j := range conditionsList {
						conditionsShow := conditionsList[j].(map[string]interface{})
						conditionsState := make(map[string]interface{})
						if v := conditionsShow["operator"]; v != nil {
							conditionsState["operator"] = v
						}
						if v := conditionsShow["value"]; v != nil {
							conditionsState["value"] = v
						}
						conditionsListState = append(conditionsListState, conditionsState)
					}
					expressionsState["conditions"] = conditionsListState
				}
				if v := expressionsShow["field-name"]; v != nil {
					expressionsState["field_name"] = v
				}
				if v := expressionsShow["operator"]; v != nil {
					expressionsState["operator"] = v
				}
				if v := expressionsShow["values"]; v != nil {
					expressionsState["values"] = v
				}
				expressionsListState = append(expressionsListState, expressionsState)
			}
			filtersState["expressions"] = expressionsListState
		}
		if v := filtersShow["expressions-operator"]; v != nil {
			filtersState["expressions_operator"] = v
		}
		if v := filtersShow["filter-out-connection-logs"]; v != nil {
			filtersState["filter_out_connection_logs"] = v
		}
		_ = d.Set("filters", []interface{}{filtersState})
	}

	if v := logExporter["fields"]; v != nil {
		fieldsShow := v.(map[string]interface{})
		fieldsState := make(map[string]interface{})
		if v := fieldsShow["exclusion-list"]; v != nil {
			fieldsState["exclusion_list"] = v
		}
		if v := fieldsShow["export"]; v != nil {
			fieldsState["export"] = v
		}
		if v := fieldsShow["names-mapping"]; v != nil {
			fieldsState["names_mapping"] = v
		}
		_ = d.Set("fields", []interface{}{fieldsState})
	}

	return nil

}

func updateManagementLogExporter(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)
	logExporter := make(map[string]interface{})

	if ok := d.HasChange("name"); ok {
		oldName, newName := d.GetChange("name")
		logExporter["name"] = oldName
		logExporter["new-name"] = newName
	} else {
		logExporter["name"] = d.Get("name")
	}

	if ok := d.HasChange("target_server"); ok {
		logExporter["target-server"] = d.Get("target_server")
	}

	if ok := d.HasChange("target_port"); ok {
		logExporter["target-port"] = d.Get("target_port")
	}

	if ok := d.HasChange("protocol"); ok {
		logExporter["protocol"] = d.Get("protocol")
	}

	if v, ok := d.GetOkExists("enabled"); ok {
		logExporter["enabled"] = v.(bool)
	}

	if d.HasChange("attachments") {

		if _, ok := d.GetOk("attachments"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOkExists("attachments.0.add_link_to_log_attachment"); ok {
				res["add-link-to-log-attachment"] = v.(bool)
			}

			if v, ok := d.GetOkExists("attachments.0.add_link_to_log_details"); ok {
				res["add-link-to-log-details"] = v.(bool)
			}

			if v, ok := d.GetOkExists("attachments.0.add_log_attachment_id"); ok {
				res["add-log-attachment-id"] = v.(bool)
			}

			logExporter["attachments"] = res
		}
	}

	if d.HasChange("data_manipulation") {

		if _, ok := d.GetOk("data_manipulation"); ok {

			res := make(map[string]interface{})

			if v, ok := d.GetOkExists("data_manipulation.0.aggregate_log_updates"); ok {
				res["aggregate-log-updates"] = v.(bool)
			}

			if v, ok := d.GetOk("data_manipulation.0.format"); ok {
				res["format"] = v.(string)
			}

			if v, ok := d.GetOk("data_manipulation.0.add_custom_log_header"); ok {
				res["add-custom-log-header"] = v.(string)
			}
			if v, ok := d.GetOk("data_manipulation.0.aggregation_mode"); ok {
				res["aggregation-mode"] = v.(string)
			}
			logExporter["data-manipulation"] = res
		}
	}

	if ok := d.HasChange("color"); ok {
		logExporter["color"] = d.Get("color")
	}

	if ok := d.HasChange("comments"); ok {
		logExporter["comments"] = d.Get("comments")
	}

	if d.HasChange("tags") {
		if v, ok := d.GetOk("tags"); ok {
			logExporter["tags"] = v.(*schema.Set).List()
		}
	}

	if v, ok := d.GetOkExists("ignore_warnings"); ok {
		logExporter["ignore-warnings"] = v.(bool)
	}

	if v, ok := d.GetOkExists("ignore_errors"); ok {
		logExporter["ignore-errors"] = v.(bool)
	}

	log.Println("Update LogExporter - Map = ", logExporter)

	if ok := d.HasChange("tls"); ok {
		if _, ok := d.GetOk("tls"); ok {
			tlsPayload := make(map[string]interface{})
			if _, ok := d.GetOk("tls.0.auto_generate_client_certificates"); ok {
				autoGenerateClientCertificatesPayload := make(map[string]interface{})
				if v, ok := d.GetOk("tls.0.auto_generate_client_certificates.0.server_ca_cert"); ok {
					autoGenerateClientCertificatesPayload["server-ca-cert"] = v.(string)
				}
				tlsPayload["auto-generate-client-certificates"] = autoGenerateClientCertificatesPayload
			}
			if v, ok := d.GetOkExists("tls.0.enabled"); ok {
				tlsPayload["enabled"] = v.(bool)
			}
			if _, ok := d.GetOk("tls.0.generate_and_sign_client_csr"); ok {
				generateAndSignClientCsrPayload := make(map[string]interface{})
				if v, ok := d.GetOk("tls.0.generate_and_sign_client_csr.0.client_cert"); ok {
					generateAndSignClientCsrPayload["client-cert"] = v.(string)
				}
				if v, ok := d.GetOk("tls.0.generate_and_sign_client_csr.0.phase"); ok {
					generateAndSignClientCsrPayload["phase"] = v.(string)
				}
				if v, ok := d.GetOk("tls.0.generate_and_sign_client_csr.0.server_ca_cert"); ok {
					generateAndSignClientCsrPayload["server-ca-cert"] = v.(string)
				}
				tlsPayload["generate-and-sign-client-csr"] = generateAndSignClientCsrPayload
			}
			if v, ok := d.GetOk("tls.0.mode"); ok {
				tlsPayload["mode"] = v.(string)
			}
			if _, ok := d.GetOk("tls.0.user_provided_certificates"); ok {
				userProvidedCertificatesPayload := make(map[string]interface{})
				if v, ok := d.GetOk("tls.0.user_provided_certificates.0.client_cert"); ok {
					userProvidedCertificatesPayload["client-cert"] = v.(string)
				}
				if v, ok := d.GetOk("tls.0.user_provided_certificates.0.client_secret"); ok {
					userProvidedCertificatesPayload["client-secret"] = v.(string)
				}
				if v, ok := d.GetOk("tls.0.user_provided_certificates.0.server_ca_cert"); ok {
					userProvidedCertificatesPayload["server-ca-cert"] = v.(string)
				}
				tlsPayload["user-provided-certificates"] = userProvidedCertificatesPayload
			}
			logExporter["tls"] = tlsPayload
		}
	}

	if ok := d.HasChange("filters"); ok {
		if _, ok := d.GetOk("filters"); ok {
			filtersPayload := make(map[string]interface{})
			if v, ok := d.GetOk("filters.0.expressions"); ok {
				expressionsList := v.([]interface{})
				if len(expressionsList) > 0 {
					var expressionsPayload []map[string]interface{}
					for i := range expressionsList {
						expressionsItem := make(map[string]interface{})
						if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".conditions"); ok {
							conditionsList := v.([]interface{})
							if len(conditionsList) > 0 {
								var conditionsPayload []map[string]interface{}
								for j := range conditionsList {
									conditionsItem := make(map[string]interface{})
									if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".conditions." + strconv.Itoa(j) + ".operator"); ok {
										conditionsItem["operator"] = v.(string)
									}
									if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".conditions." + strconv.Itoa(j) + ".value"); ok {
										conditionsItem["value"] = v.(string)
									}
									conditionsPayload = append(conditionsPayload, conditionsItem)
								}
								expressionsItem["conditions"] = conditionsPayload
							}
						}
						if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".field_name"); ok {
							expressionsItem["field-name"] = v.(string)
						}
						if v, ok := d.GetOk("filters.0.expressions." + strconv.Itoa(i) + ".operator"); ok {
							expressionsItem["operator"] = v.(string)
						}
						expressionsPayload = append(expressionsPayload, expressionsItem)
					}
					filtersPayload["expressions"] = expressionsPayload
				}
			}
			if v, ok := d.GetOk("filters.0.expressions_operator"); ok {
				filtersPayload["expressions-operator"] = v.(string)
			}
			if v, ok := d.GetOkExists("filters.0.filter_out_connection_logs"); ok {
				filtersPayload["filter-out-connection-logs"] = v.(bool)
			}
			logExporter["filters"] = filtersPayload
		}
	}

	if ok := d.HasChange("fields"); ok {
		if _, ok := d.GetOk("fields"); ok {
			fieldsPayload := make(map[string]interface{})
			if v, ok := d.GetOk("fields.0.exclusion_list"); ok {
				fieldsPayload["exclusion-list"] = v.(string)
			}
			if v, ok := d.GetOk("fields.0.export"); ok {
				fieldsPayload["export"] = v.(string)
			}
			if v, ok := d.GetOk("fields.0.names_mapping"); ok {
				fieldsPayload["names-mapping"] = v.(string)
			}
			logExporter["fields"] = fieldsPayload
		}
	}

	updateLogExporterRes, err := client.ApiCallSimple("set-log-exporter", logExporter)
	if err != nil || !updateLogExporterRes.Success {
		if updateLogExporterRes.ErrorMsg != "" {
			return fmt.Errorf("%s", updateLogExporterRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}

	return readManagementLogExporter(d, m)
}

func deleteManagementLogExporter(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	logExporterPayload := map[string]interface{}{
		"uid": d.Id(),
	}

	log.Println("Delete LogExporter")

	deleteLogExporterRes, err := client.ApiCallSimple("delete-log-exporter", logExporterPayload)
	if err != nil || !deleteLogExporterRes.Success {
		if deleteLogExporterRes.ErrorMsg != "" {
			return fmt.Errorf("%s", deleteLogExporterRes.ErrorMsg)
		}
		return fmt.Errorf("%s", err.Error())
	}
	d.SetId("")

	return nil
}
