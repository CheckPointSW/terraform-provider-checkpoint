---
layout: "checkpoint"
page_title: "checkpoint_management_log_exporter"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-log-exporter"
description: |-
This resource allows you to execute Check Point Log Exporter.
---

# checkpoint_management_log_exporter

This resource allows you to execute Check Point Log Exporter.

## Example Usage


```hcl
resource "checkpoint_management_log_exporter" "example" {
  name = "newLogExporter"
  target_server = "1.2.3.4"
  target_port = 1234
  protocol = "tcp"
  attachments {
    add_link_to_log_attachment = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `target_server` - (Required) Target server port to which logs are exported. 
* `target_port` - (Required) Port number of the target server. 
* `protocol` - (Optional) Protocol used to send logs to the target server. 
* `enabled` - (Optional) Indicates whether to enable export. 
* `attachments` - (Optional) Log exporter attachments. attachments blocks are documented below.
* `data_manipulation` - (Optional) Log exporter data manipulation. data_manipulation blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `advanced` - (Optional) Advanced settings for log exporter.advanced blocks are documented below.
* `fields` - (Optional) Log exporter fields mapping settings.fields blocks are documented below.
* `filters` - (Optional) Log exporter filters.filters blocks are documented below.
* `tls` - (Optional) TLS settings for log exporter.tls blocks are documented below.


`attachments` supports the following:

* `add_link_to_log_attachment` - (Optional) Indicates whether to add link to log attachment in SmartView. 
* `add_link_to_log_details` - (Optional) Indicates whether to add link to log details in SmartView. 
* `add_log_attachment_id` - (Optional) Indicates whether to add log attachment ID. 


`data_manipulation` supports the following:

* `aggregate_log_updates` - (Optional) Indicates whether to aggregate log updates. 
* `format` - (Optional) Logs format. 
* `add_custom_log_header` - (Optional) Custom header added to every exported log entry.<br> This header can be used to identify the source or add metadata to logs.<br> The custom header is prepended.
* `aggregation_mode` - (Optional) Aggregation mode for aggregated logs.<br> Determines how log updates are processed when aggregate-log-updates is true or not specified:<br> - semi-unified (defa.


`advanced` supports the following:

* `time_in_milliseconds` - (Optional) Toggle export logs' time with ms resolution.


`fields` supports the following:

* `exclusion_list` - (Optional) List of field names to exclude or include based on the export mode. <br> When export is <b>all-except</b>: fields in this list are excluded from expor...
* `export` - (Optional) Field export mode. <br><b>all-except</b> (default): exports all fields except those in the exclusion list. <br><b>nothing-except</b>: exports nothing ...
* `names_mapping` - (Optional) Field name mappings to rename fields during export. <br> Specify mappings in the format: 'orgName:dstName' (e.g., 'src:source_ip,dst:destination_ip')....


`filters` supports the following:

* `expressions` - (Optional) List of filter expressions. Each expression targets a specific log field.expressions blocks are documented below.
* `expressions_operator` - (Optional) Logical operator for combining all filter expressions.
* `filter_out_connection_logs` - (Optional) When set to true, connection logs will be excluded from export.


`expressions` supports the following:

* `conditions` - (Optional) Filter conditions for this field. Required when 'operator' is 'and' or 'or'.<br>Array of condition objects, each with 'operator' and 'value'.conditions blocks are documented below.
* `field_name` - (Optional) Log field name to filter on.
* `operator` - (Optional) Operator for this expression.<br>Use 'and'/'or' with 'conditions' (array of condition objects).<br>Use 'in'/'not-in' with 'values' (flat array of stri...
* `values` - (Computed) Filter values for this field. Present when 'operator' is 'in' or 'not-in'.<br>Flat array of string values.


`conditions` supports the following:

* `operator` - (Optional) Comparison operator.
* `value` - (Optional) The value to compare against.


`tls` supports the following:

* `auto_generate_client_certificates` - (Optional) TLS mode where user provides server CA certificate only. Check Point Management automatically generates and signs client certificates. This is the sim...auto_generate_client_certificates blocks are documented below.
* `client_cert_expiration_date` - (Computed) Client certificate expiration date extracted from the stored PKCS#12 certificate. Available for all TLS modes when client certificate exists.
* `enabled` - (Optional) Indicates whether to use encrypted connection. Can be true only when protocol is TCP.
* `generate_and_sign_client_csr` - (Optional) Advanced TLS mode with two separate phases: 'auto-generate-client-csr' to initiate CSR generation (ADD/SET) and 'provide-signed-client-cert' to provid...generate_and_sign_client_csr blocks are documented below.
* `mode` - (Optional) TLS mode. Must match the specific mode configuration provided.
* `user_provided_certificates` - (Optional) TLS mode where user provides server CA certificate, client certificate (PKCS#12 format), and passphrase. All certificates provided by the user, Check ...user_provided_certificates blocks are documented below.


`auto_generate_client_certificates` supports the following:

* `client_ca_cert` - (Computed) Client generated CA certificate that signed the client certificate. The server should trust this CA. Supplied as <b>base64 encoded string</b>.
* `server_ca_cert` - (Optional) Server certificate authority certificate, supplied as <b>base64 encoded string</b>.


`generate_and_sign_client_csr` supports the following:

* `client_cert` - (Optional) Client certificate in PEM format as <b>base64 encoded string</b>, signed by external CA. Can only be provided when phase is 'provide-signed-client-cert'.
* `client_csr` - (Computed) Client certificate signing request (CSR) in PEM format as <b>base64 encoded string</b>. Present when phase is 'auto-generate-client-csr'.
* `phase` - (Optional) Mode phase type: 'auto-generate-client-csr' to initiate CSR generation.
* `server_ca_cert` - (Optional) Server CA certificate in PEM format as <b>base64 encoded string</b> to validate the destination server's identity. Required when phase is 'auto-genera...


`user_provided_certificates` supports the following:

* `client_cert` - (Optional) Client certificate, supplied as <b>base64 encoded string</b>.
* `client_secret` - (Optional) Client secret for TLS connection.
* `server_ca_cert` - (Optional) Server certificate authority certificate, supplied as <b>base64 encoded string</b>.
