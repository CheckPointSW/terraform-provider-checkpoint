---
layout: "checkpoint"
page_title: "checkpoint_gaia_syslog"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-syslog"
description: |-
This resource allows you to execute Check Point Syslog.
---

# checkpoint_gaia_syslog

This resource allows you to execute Check Point Syslog.

## Example Usage


```hcl
resource "checkpoint_gaia_syslog" "example" {
  audit_log = true
  cp_logs = false
  filename = "/var/log/messages"
  send_to_mgmt = true
  forwarded_logs_files {
    tag = "conf_point2"
    path = "/var/log/config_point_server.log"
  }
  forwarded_logs_files {
    path = "/var/log/gaia_api_server.log"
  }
  forwarded_logs_files {
    path = "/var/log/gaia_init_config.log"
  }
}
```

## Argument Reference

The following arguments are supported:

* `audit_log` - (Optional) syslog auditlog permanent 
* `cp_logs` - (Optional) syslog auditlog permanent 
* `send_to_mgmt` - (Optional) sending logs to Management server 
* `filename` - (Optional) syslog output filename 
* `tls_configuration` - (Optional) system TLS configuration in order to enable sending encrtyped syslog messages to remote host, Supported starting from R82 tls_configuration blocks are documented below.
* `forwarded_logs_files` - (Optional) Custom log files List. Supported starting from Gaia version R82.10 forwarded_logs_files blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`tls_configuration` supports the following:

* `ca_certification` - (Optional) Certificate file path of the certification authority CA. Supported starting from Gaia version R82 
* `public_key` - (Optional) Public key file path signed by the CA. Supported starting from Gaia version R82 
* `private_key` - (Optional) Private key file path. Supported starting from Gaia version R82 


`forwarded_logs_files` supports the following:

* `path` - (Optional) Path to the forwarded log file. 
* `tag` - (Optional) Tag for the forwarded log file. 
