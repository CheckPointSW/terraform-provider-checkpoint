---
layout: "checkpoint"
page_title: "checkpoint_gaia_open_telemetry"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-open-telemetry"
description: |-
This resource allows you to execute Check Point Open Telemetry.
---

# checkpoint_gaia_open_telemetry

This resource allows you to execute Check Point Open Telemetry.

## Example Usage


```hcl
resource "checkpoint_gaia_open_telemetry" "example" {
  enabled = true
  export_targets {
    type    = "otlp"
    enabled = true
    url     = "http://otel-collector:4317"
    name    = "my-exporter"
  }
}
```

## Argument Reference

The following arguments are supported:

* `export_targets` - (Required) Settings of OpenTelemetry export targets export_targets blocks are documented below.
* `enabled` - (Required) State of OpenTelemetry (enabled or disabled). 
* `metrics` - (Optional) Settings to include or exclude which metrics are exporterd from this machine. metrics blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`export_targets` supports the following:

* `type` - (Optional) Specifies the type of this OpenTelemetry Exporter. 
* `enabled` - (Optional) Specifies the state of this OpenTelemetry Exporter. 
* `url` - (Optional) Specifies the URL of this OpenTelemetry Exporter. 
* `name` - (Optional) Optional. Specifies the user-defined name for this OpenTelemetry Exporter.

 You can use this name parameter later to remove or to change settings of this OpenTelemetry Exporter. 
* `client_auth` - (Optional) Specifies the client authentication method for this OpenTelemetry Exporter. client_auth blocks are documented below.
* `server_auth` - (Optional) Specifies the server authentication method. server_auth blocks are documented below.


`metrics` supports the following:

* `include` - (Optional) Determines the default behavior for metric inclusion. 
* `except` - (Optional) Defines specific metrics to be either excluded (if include is 'all') or included (if include is 'none'). except blocks are documented below.


`client_auth` supports the following:

* `basic` - (Optional) Specifies the client authentication credentials for this OpenTelemetry Exporter.

 If this OpenTelemetry Exporter does not require client authentication, then enter the value "N/A" for the username and for the password. basic blocks are documented below.
* `token` - (Optional) Specifies the client authentication bearer token for this OpenTelemetry Exporter.

 If this OpenTelemetry Exporter does not require client authentication, then enter the value "N/A" for the bearer token. token blocks are documented below.


`server_auth` supports the following:

* `ca_public_key` - (Optional) Specifies the Public key and its type for the Certificate Authority (CA) that this OpenTelemetry Exporter uses. ca_public_key blocks are documented below.


`basic` supports the following:

* `username` - (Optional) The username used for basic HTTP/HTTPS basic authentication on this OpenTelemetry Exporter. 
* `password` - (Optional) The password used for basic HTTP/HTTPS basic authentication on this OpenTelemetry Exporter.

 We recommend to use a bcrypt hash on the password. 


`token` supports the following:

* `header_bearer_token` - (Optional) Specifies the JWT or other similar protocol bearer token. 
* `custom_header` - (Optional) A custom header for authentication purpose. custom_header blocks are documented below.


`ca_public_key` supports the following:

* `type` - (Optional) The CA Public Key Type for this OpenTelemetry Exporter - PEM-X509 or the default CA bundle. 
* `value` - (Optional) The CA Public Key string for this OpenTelemetry Exporter.
If you specified the type 'Default', then enter the value 'N/A'. 


`custom_header` supports the following:

* `key` - (Optional) The key name of the custom header used for authentication. 
* `value` - (Optional) The value of the custom header used for authentication. 
