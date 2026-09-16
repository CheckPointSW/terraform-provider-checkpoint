---
layout: "checkpoint"
page_title: "checkpoint_management_log_exporter"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-log-exporter"
description: |- Use this data source to get information on an existing Log Exporter.
---


# checkpoint_management_log_exporter

Use this data source to get information on an existing Log Exporter.

## Example Usage


```hcl
data "checkpoint_management_log_exporter" "data_log_exporter" {
name = "example_log_exporter"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `target_server` - Target server port to which logs are exported.
* `target_port` - Port number of the target server.
* `protocol` - Protocol used to send logs to the target server.
* `enabled` - Indicates whether to enable export.
* `attachments` - Log exporter attachments. attachments blocks are documented below.
* `data_manipulation` - Log exporter data manipulation. data_manipulation blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `advanced` - Log exporter advanced settings.advanced blocks are documented below.
* `fields` - Log exporter fields settings.fields blocks are documented below.
* `filters` - Log exporter filters.filters blocks are documented below.
* `target_cloud` - Cloud target configuration for log export.target_cloud blocks are documented below.
* `tls` - Log exporter TLS settings.tls blocks are documented below.


`attachments` supports the following:

* `add_link_to_log_attachment` - Indicates whether to add link to log attachment in SmartView.
* `add_link_to_log_details` - Indicates whether to add link to log details in SmartView.
* `add_log_attachment_id` - Indicates whether to add log attachment ID.


`data_manipulation` supports the following:

* `aggregate_log_updates` - Indicates whether to aggregate log updates.
* `format` - Logs format. 
* `add_custom_log_header` - Custom header for log export.
* `aggregation_mode` - Aggregation mode for aggregated logs.<br> Determines how log updates are processed when aggregate-log-updates is true or not specified:<br> - semi-uni...


`advanced` supports the following:

* `time_in_milliseconds` - Toggle export logs' time with ms resolution.


`fields` supports the following:

* `exclusion_list` - List of field names excluded or included based on the export mode.
* `export` - Field export mode. <br><b>all-except</b> (default): exports all fields except those in the exclusion list. <br><b>nothing-except</b>: exports nothing ...
* `names_mapping` - Field name mappings applied during export. Format: 'orgName1:dstName1,orgName2:dstName2,...'.


`filters` supports the following:

* `expressions` - List of filter expressions. Each expression targets a specific log field.expressions blocks are documented below.
* `expressions_operator` - Logical operator for combining all filter expressions.
* `filter_out_connection_logs` - When set to true, connection logs will be excluded from export.


`expressions` supports the following:

* `conditions` - Filter conditions for this field. Present when 'operator' is 'and' or 'or'.<br>Array of condition objects, each with 'operator' and 'value'.conditions blocks are documented below.
* `field_name` - Log field name to filter on.
* `operator` - Operator for this expression.<br>Use 'and'/'or' with 'conditions' (array of condition objects).<br>Use 'in'/'not-in' with 'values' (flat array of stri...
* `values` - Filter values for this field. Present when 'operator' is 'in' or 'not-in'.<br>Flat array of string values.


`conditions` supports the following:

* `operator` - Comparison operator.
* `value` - The value to compare against.


`target_cloud` supports the following:

* `s3_bucket_auth` - S3 bucket authentication configuration details.s3_bucket_auth blocks are documented below.
* `type` - Cloud target type. Currently only 's3-bucket' is supported.


`s3_bucket_auth` supports the following:

* `complete_configuration` - S3 bucket ARNs configuration details.complete_configuration blocks are documented below.
* `create_certificate` - S3 bucket configuration with auto-generated certificates details.create_certificate blocks are documented below.
* `operation` - Current S3 authentication operation phase: 'create-certificate' or 'complete-configuration'.


`complete_configuration` supports the following:

* `exporter_enabled` - Indicates whether the log exporter is now enabled after ARN configuration.
* `message` - Status message indicating ARN configuration success and exporter enablement.
* `profile_arn` - Configured AWS IAM Roles Anywhere profile ARN.
* `role_arn` - Configured AWS IAM role ARN.
* `trust_anchor_arn` - Configured AWS IAM Roles Anywhere trust anchor ARN.


`create_certificate` supports the following:

* `bucket` - Configured S3 bucket name.
* `ca_certificate` - CA certificate that signed the S3 client certificate. Supplied as base64 encoded PEM string. Decode this certificate and upload it to AWS IAM Roles An...
* `message` - Status message indicating certificate generation success and next steps.
* `prefix` - Configured S3 prefix (if provided).
* `region` - Configured AWS region.


`tls` supports the following:

* `auto_generate_client_certificates` - TLS mode where user provides server CA certificate only. Check Point Management automatically generates and signs client certificates. This is the sim...auto_generate_client_certificates blocks are documented below.
* `client_cert_expiration_date` - Client certificate expiration date extracted from the stored PKCS#12 certificate. Available for all TLS modes when client certificate exists.
* `enabled` - Indicates whether to use encrypted connection. Can be true only when protocol is TCP.
* `generate_and_sign_client_csr` - Advanced TLS mode with two separate phases: 'auto-generate-client-csr' to initiate CSR generation (ADD/SET) and 'provide-signed-client-cert' to provid...generate_and_sign_client_csr blocks are documented below.
* `mode` - TLS mode. Must match the specific mode configuration provided.


`auto_generate_client_certificates` supports the following:

* `client_ca_cert` - Client generated CA certificate that signed the client certificate. The server should trust this CA. Supplied as <b>base64 encoded string</b>.


`generate_and_sign_client_csr` supports the following:

* `client_csr` - Client certificate signing request (CSR) in PEM format as <b>base64 encoded string</b>. Present when phase is 'auto-generate-client-csr'.
* `phase` - Advanced mode phase type: 'auto-generate-client-csr' indicates CSR generation, or 'provide-signed-client-cert' indicates signed certificate provided.
