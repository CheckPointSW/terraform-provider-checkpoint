---
layout: "checkpoint"
page_title: "checkpoint_management_smart_task"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-smart-task"
description: |-
Use this data source to get information on an existing Check Point Smart Task.
---

#Data Source checkpoint_management_smart_task

TUse this data source to get information on an existing Check Point Smart Task.

## Example Usage


```hcl
 resource "checkpoint_management_smart_task" "smart_task" {

  name = "smt"
  trigger = "Before Publish"
  description = "my smart task"
  action {

    send_web_request {
      url            = "https://demo.example.com/policy-installation-reports"
      fingerprint    = "8023a5652ba2c8f5b0902363a5314cd2b4fdbc5c"
      override_proxy = true
      proxy_url      = "https://demo.example.com/policy-installation-reports"
      time_out       = 200
      shared_secret  = " secret"
    }
  }
  enabled = true
}

data "checkpoint_management_smart_task" "data_smart_task" {
  
  name = "${checkpoint_management_smart_task.smart_task.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `action` - (Optional) The action to be run when the trigger is fired.action blocks are documented below.
* `trigger` -  Trigger type associated with the SmartTask. 
* `custom_data` -  Per SmartTask custom data in JSON format.<br>When the trigger is fired, the trigger data is converted to JSON. The custom data is then concatenated to the trigger data JSON. 
* `description` -  Description of the SmartTask's functionality and options. 
* `enabled` -  Whether the SmartTask is enabled and will run when triggered. 
* `fail_open` -  If the action fails to execute, whether to treat the execution failure as an error, or continue. 
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` -  Comments string. 
* `ignore_warnings` -  Apply changes ignoring warnings. 
* `ignore_errors` -  Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`action` supports the following (exactly one of them must be defined):

* `send_web_request` -  When the trigger is fired, sends an HTTPS POST web request to the configured URL.<br>The trigger data will be passed along with the SmartTask's custom data in the request's payload.send_web_request blocks are documented below.
* `run_script` -  When the trigger is fired, runs the configured Repository Script on the defined targets.<br>The trigger data is then passed to the script as the first parameter. The parameter is JSON encoded in Base64 format.run_script blocks are documented below.
* `send_mail` - When the trigger is fired, sends the configured email to the defined recipients.send_mail blocks are documented below.


`send_web_request` supports the following:

* `url` -  URL used for the web request. 
* `fingerprint` -  The SHA1 fingerprint of the URL's SSL certificate. Used to trust servers with self-signed SSL certificates. 
* `override_proxy` -  Option to send to the web request via a proxy other than the Management's Server proxy (if defined). 
* `proxy_url` - URL of the proxy used to send the request. 
* `shared_secret` -  Shared secret that can be used by the target server to identify the Management Server.<br>The value will be sent as part of the request in the "X-chkp-shared-secret" header. 
* `time_out` -  Web Request time-out in seconds. 


`run_script` supports the following:

* `repository_script` -  Repository script that is executed when the trigger is fired.,  identified by the name or UID. 
* `targets` -  Targets to execute the script on.targets blocks are documented below.
* `time_out` -  Script execution time-out in seconds. 


`send_mail` supports the following:

* `mail_settings` -  The required settings to send the mail by.mail_settings blocks are documented below.
* `smtp_server` -  The UID or the name a preconfigured SMTP server object. 


`mail_settings` supports the following:

* `recipients` -  A comma separated list of recipient mail addresses. 
* `sender_email` -  An email address to send the mail from. 
* `subject` - The email subject. 
* `body` -  The email body. 
* `attachment` -  What file should be attached to the mail. 
* `bcc_recipients` - A comma separated list of bcc recipient mail addresses. 
* `cc_recipients` -  A comma separated list of cc recipient mail addresses. 

`smtp_server` supports the following :

* `name` - the SMTP server name.
* `port` - The SMTP port to use.
* `server` - The SMTP server address.
* `authentication` - Does the mail server requires authentication.
* `encryption`- Encryption type.
* `username`- A username for the SMTP server.