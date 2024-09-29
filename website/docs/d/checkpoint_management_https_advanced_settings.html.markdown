---
layout: "checkpoint"
page_title: "checkpoint_management_https_advanced_settings"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-https-advanced-settings"
description: |-
Use this data source to get information on an existing Check Point Https Advanced Settings.
---

# Data Source:checkpoint_management_https_advanced_settings

Use this data source to get information on an existing Check Point Https Advanced Settings.

## Example Usage


```hcl
resource "checkpoint_management_set_https_advanced_settings" "example" {
  bypass_on_failure = false
  bypass_on_client_failure = false
  site_categorization_allow_mode = "background"
  blocked_certificate_tracking = "popup alert"
  bypass_update_services = true
  certificate_pinned_apps_action = "bypass"
  log_sessions = true
  retrieve_intermediate_ca_certificates = true
  server_certificate_validation_actions = {
    block_expired = true
    block_revoked = false
    block_untrusted = true
    track_errors = "log"
  }
  blocked_certificates {
    name = "BlackListed_A71D5266-7EF0-42CF-AE9C-409CD4093879"
    cert_serial_number = "3e:75:ce:d4:6b:69:30:21:21:88:30:ae:86:a8:2a:71"
    comments = "login.yahoo.com"
  }
  blocked_certificates {
    name = "BlackListed_A2B37A3D-53F9-4A24-AD09-D96272CA1710"
    cert_serial_number = "00:d7:55:8f:da:f5:f1:10:5b:b2:13:28:2b:70:77:29:a3"
    comments = "*.EGO.GOV.TR"
  }
}

data "checkpoint_management_https_advanced_settings" "data" {
  depends_on = [checkpoint_management_set_https_advanced_settings.example]
}
```

## Argument Reference

The following arguments are supported:
* `uid` -  Object unique identifier.
* `bypass_on_client_failure` -  Whether all requests should be bypassed or blocked-in case of client errors (Client closes the connection due to authentication issues during handshake)<br><ul style="list-style-type:square"><li>true - Fail-open (bypass all requests).</li><li>false - Fail-close (block all requests.</li></ul><br>The default value is true. 
* `bypass_on_failure` -  Whether all requests should be bypassed or blocked-in case of server errors (for example validation error during GW-Server authentication)<br><ul style="list-style-type:square"><li>true - Fail-open (bypass all requests).</li><li>false - Fail-close (block all requests.</li></ul><br>The default value is true. 
* `bypass_under_load` -  Bypass the HTTPS Inspection temporarily to improve connectivity during a heavy load on the Security Gateway. The HTTPS Inspection would resume as soon as the load decreases.bypass_under_load blocks are documented below.
* `site_categorization_allow_mode` - Whether all requests should be allowed or blocked until categorization is complete.<br><ul style="list-style-type:square"><li>Background - to allow requests until categorization is complete.</li><li>Hold- to block requests until categorization is complete.</li></ul><br>The default value is hold. 
* `server_certificate_validation_actions` - When a Security Gateway receives an untrusted certificate from a website server, define when to drop the connection and how to track it.server_certificate_validation_actions blocks are documented below.
* `retrieve_intermediate_ca_certificates` -  Configure the value "true" to use the "Certificate Authority Information Access" extension to retrieve certificates that are missing from the certificate chain.<br>The default value is true. 
* `blocked_certificates` -  Collection of certificates objects identified by serial number.<br>Drop traffic from servers using the blocked certificate.blocked_certificates blocks are documented below.
* `blocked_certificate_tracking` -  Controls whether to log and send a notification for dropped traffic.<br><ul style="list-style-type:square"><li>None - Does not record the event.</li><li>Log - Records the event details in SmartView.</li><li>Alert - Logs the event and executes a command.</li><li>Mail - Sends an email to the administrator.</li><li>SNMP Trap - Sends an SNMP alert to the SNMP GU.</li><li>User Defined Alert - Sends customized alerts.</li></ul>. 
* `bypass_update_services` -  Configure the value "true" to bypass traffic to well-known software update services.<br>The default value is true. 
* `certificate_pinned_apps_action` -  Configure the value "bypass" to bypass traffic from certificate-pinned applications approved by Check Point.<br>HTTPS Inspection cannot inspect connections initiated by certificate-pinned applications.<br>Configure the value "detect" to send logs for traffic from certificate-pinned applications approved by Check Point.<br>The default value is bypass. 
* `log_sessions` -  The value "true" configures the Security Gateway to send HTTPS Inspection session logs.<br>The default value is true. 
* `domains_to_process` -  Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` -  Apply changes ignoring warnings. 
* `ignore_errors` -  Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`bypass_under_load` supports the following:

* `track` -  Whether to log and send a notification for the bypass under load:<ul style="list-style-type:square"><li>None - Does not record the event.</li><li>Log - Records the event details. Use SmartConsole or SmartView to see the logs.</li><li>Alert - Logs the event and executes a command you configured.</li><li>Mail - Sends an email to the administrator.</li><li>SNMP Trap - Sends an SNMP alert to the configured SNMP Management Server.</li><li>User Defined Alert - Sends a custom alert.</li></ul>. 


`server_certificate_validation_actions` supports the following:

* `block_expired` - Set to be true in order to drop traffic from servers with expired server certificate. 
* `block_revoked` -  Set to be true in order to drop traffic from servers with revoked server certificate (validate CRL). 
* `block_untrusted` -  Set to be true in order to drop traffic from servers with untrusted server certificate. 
* `track_errors` -  Whether to log and send a notification for the server validation errors:<br><ul style="list-style-type:square"><li>None - Does not record the event.</li><li>Log - Records the event details in SmartView.</li><li>Alert - Logs the event and executes a command.</li><li>Mail - Sends an email to the administrator.</li><li>SNMP Trap - Sends an SNMP alert to the SNMP GU.</li><li>User Defined Alert - Sends customized alerts.</li></ul>. 


`blocked_certificates` supports the following:

* `name` -  Describes the name, cannot be overridden. 
* `cert_serial_number` -  Certificate Serial Number (unique) in hexadecimal format HH:HH. 
* `comments` -  Describes the certificate by default, can be overridden by any text. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

