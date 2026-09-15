package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func dataSourceManagementLogExporter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceManagementLogExporterRead,
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
			"target_server": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target server port to which logs are exported.",
			},
			"target_port": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Port number of the target server.",
			},
			"protocol": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Protocol used to send logs to the target server.",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Indicates whether to enable export.",
			},
			"attachments": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log exporter attachments.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"add_link_to_log_attachment": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to add link to log attachment in SmartView.",
						},
						"add_link_to_log_details": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to add link to log details in SmartView.",
						},
						"add_log_attachment_id": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to add log attachment ID.",
						},
					},
				},
			},
			"data_manipulation": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log exporter data manipulation.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"aggregate_log_updates": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates whether to aggregate log updates.",
						},
						"format": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Logs format.",
						},
						"add_custom_log_header": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Custom header for log export.",
						},
						"aggregation_mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Aggregation mode for aggregated logs.<br> Determines how log updates are processed when aggregate-log-updates is true or not specified:<br> - semi-uni...",
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
			"advanced": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log exporter advanced settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"time_in_milliseconds": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Toggle export logs' time with ms resolution.",
						},
					},
				},
			},
			"fields": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log exporter fields settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"exclusion_list": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "List of field names excluded or included based on the export mode.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"export": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Field export mode. <br><b>all-except</b> (default): exports all fields except those in the exclusion list. <br><b>nothing-except</b>: exports nothing ...",
						},
						"names_mapping": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Field name mappings applied during export. Format: 'orgName1:dstName1,orgName2:dstName2,...'.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"filters": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log exporter filters.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"expressions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of filter expressions. Each expression targets a specific log field.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"conditions": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Filter conditions for this field. Present when 'operator' is 'and' or 'or'.<br>Array of condition objects, each with 'operator' and 'value'.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"operator": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Comparison operator.",
												},
												"value": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The value to compare against.",
												},
											},
										},
									},
									"field_name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Log field name to filter on.",
									},
									"operator": {
										Type:        schema.TypeString,
										Computed:    true,
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
							Computed:    true,
							Description: "Logical operator for combining all filter expressions.",
						},
						"filter_out_connection_logs": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "When set to true, connection logs will be excluded from export.",
						},
					},
				},
			},
			"target_cloud": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Cloud target configuration for log export.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"s3_bucket_auth": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "S3 bucket authentication configuration details.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"complete_configuration": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "S3 bucket ARNs configuration details.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"exporter_enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Indicates whether the log exporter is now enabled after ARN configuration.",
												},
												"message": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Status message indicating ARN configuration success and exporter enablement.",
												},
												"profile_arn": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Configured AWS IAM Roles Anywhere profile ARN.",
												},
												"role_arn": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Configured AWS IAM role ARN.",
												},
												"trust_anchor_arn": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Configured AWS IAM Roles Anywhere trust anchor ARN.",
												},
											},
										},
									},
									"create_certificate": {
										Type:        schema.TypeList,
										Computed:    true,
										Description: "S3 bucket configuration with auto-generated certificates details.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"bucket": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Configured S3 bucket name.",
												},
												"ca_certificate": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "CA certificate that signed the S3 client certificate. Supplied as base64 encoded PEM string. Decode this certificate and upload it to AWS IAM Roles An...",
												},
												"message": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Status message indicating certificate generation success and next steps.",
												},
												"prefix": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Configured S3 prefix (if provided).",
												},
												"region": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Configured AWS region.",
												},
											},
										},
									},
									"operation": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Current S3 authentication operation phase: 'create-certificate' or 'complete-configuration'.",
									},
								},
							},
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud target type. Currently only 's3-bucket' is supported.",
						},
					},
				},
			},
			"tls": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Log exporter TLS settings.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_generate_client_certificates": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "TLS mode where user provides server CA certificate only. Check Point Management automatically generates and signs client certificates. This is the sim...",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_ca_cert": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client generated CA certificate that signed the client certificate. The server should trust this CA. Supplied as <b>base64 encoded string</b>.",
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
							Computed:    true,
							Description: "Indicates whether to use encrypted connection. Can be true only when protocol is TCP.",
						},
						"generate_and_sign_client_csr": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Advanced TLS mode with two separate phases: 'auto-generate-client-csr' to initiate CSR generation (ADD/SET) and 'provide-signed-client-cert' to provid...",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"client_csr": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Client certificate signing request (CSR) in PEM format as <b>base64 encoded string</b>. Present when phase is 'auto-generate-client-csr'.",
									},
									"phase": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Advanced mode phase type: 'auto-generate-client-csr' indicates CSR generation, or 'provide-signed-client-cert' indicates signed certificate provided.",
									},
								},
							},
						},
						"mode": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "TLS mode. Must match the specific mode configuration provided.",
						},
					},
				},
			},
		},
	}
}

func dataSourceManagementLogExporterRead(d *schema.ResourceData, m interface{}) error {

	client := m.(*checkpoint.ApiClient)

	name := d.Get("name").(string)
	uid := d.Get("uid").(string)

	payload := make(map[string]interface{})

	if name != "" {
		payload["name"] = name
	} else if uid != "" {
		payload["uid"] = uid
	}

	showLogExporterRes, err := client.ApiCallSimple("show-log-exporter", payload)
	if err != nil {
		return fmt.Errorf("%s", err.Error())
	}
	if !showLogExporterRes.Success {
		return fmt.Errorf("%s", showLogExporterRes.ErrorMsg)
	}

	logExporter := showLogExporterRes.GetData()

	log.Println("Read LogExporter - Show JSON = ", logExporter)

	if v := logExporter["uid"]; v != nil {
		_ = d.Set("uid", v)
		d.SetId(v.(string))
	}

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
			attachmentsMapToReturn["add_link_to_log_attachment"] = v.(bool)
		}
		if v, _ := attachmentsMap["add-link-to-log-details"]; v != nil {
			attachmentsMapToReturn["add_link_to_log_details"] = v.(bool)
		}
		if v, _ := attachmentsMap["add-log-attachment-id"]; v != nil {
			attachmentsMapToReturn["add_log_attachment_id"] = v.(bool)
		}

		_ = d.Set("attachments", []interface{}{attachmentsMapToReturn})
	} else {
		_ = d.Set("attachments", nil)
	}

	if logExporter["data-manipulation"] != nil {

		dataManipulationMap := logExporter["data-manipulation"].(map[string]interface{})

		dataManipulationMapToReturn := make(map[string]interface{})

		if v, _ := dataManipulationMap["aggregate-log-updates"]; v != nil {
			dataManipulationMapToReturn["aggregate_log_updates"] = v.(bool)
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

	if v := logExporter["advanced"]; v != nil {
		advancedShow := v.(map[string]interface{})
		advancedState := make(map[string]interface{})
		if v := advancedShow["time-in-milliseconds"]; v != nil {
			advancedState["time_in_milliseconds"] = v
		}
		_ = d.Set("advanced", []interface{}{advancedState})
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
					for i := range conditionsList {
						conditionsShow := conditionsList[i].(map[string]interface{})
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

	if v := logExporter["target-cloud"]; v != nil {
		targetCloudShow := v.(map[string]interface{})
		targetCloudState := make(map[string]interface{})
		if v := targetCloudShow["s3-bucket-auth"]; v != nil {
			s3BucketAuthShow := v.(map[string]interface{})
			s3BucketAuthState := make(map[string]interface{})
			if v := s3BucketAuthShow["complete-configuration"]; v != nil {
				completeConfigurationShow := v.(map[string]interface{})
				completeConfigurationState := make(map[string]interface{})
				if v := completeConfigurationShow["exporter-enabled"]; v != nil {
					completeConfigurationState["exporter_enabled"] = v
				}
				if v := completeConfigurationShow["message"]; v != nil {
					completeConfigurationState["message"] = v
				}
				if v := completeConfigurationShow["profile-arn"]; v != nil {
					completeConfigurationState["profile_arn"] = v
				}
				if v := completeConfigurationShow["role-arn"]; v != nil {
					completeConfigurationState["role_arn"] = v
				}
				if v := completeConfigurationShow["trust-anchor-arn"]; v != nil {
					completeConfigurationState["trust_anchor_arn"] = v
				}
				s3BucketAuthState["complete_configuration"] = []interface{}{completeConfigurationState}
			}
			if v := s3BucketAuthShow["create-certificate"]; v != nil {
				createCertificateShow := v.(map[string]interface{})
				createCertificateState := make(map[string]interface{})
				if v := createCertificateShow["bucket"]; v != nil {
					createCertificateState["bucket"] = v
				}
				if v := createCertificateShow["ca-certificate"]; v != nil {
					createCertificateState["ca_certificate"] = v
				}
				if v := createCertificateShow["message"]; v != nil {
					createCertificateState["message"] = v
				}
				if v := createCertificateShow["prefix"]; v != nil {
					createCertificateState["prefix"] = v
				}
				if v := createCertificateShow["region"]; v != nil {
					createCertificateState["region"] = v
				}
				s3BucketAuthState["create_certificate"] = []interface{}{createCertificateState}
			}
			if v := s3BucketAuthShow["operation"]; v != nil {
				s3BucketAuthState["operation"] = v
			}
			targetCloudState["s3_bucket_auth"] = []interface{}{s3BucketAuthState}
		}
		if v := targetCloudShow["type"]; v != nil {
			targetCloudState["type"] = v
		}
		_ = d.Set("target_cloud", []interface{}{targetCloudState})
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

	return nil
}
