---
layout: "checkpoint"
page_title: "checkpoint_management_resource_smtp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-smtp"
description: |-
This resource allows you to execute Check Point Resource Smtp.
---

# checkpoint_management_resource_smtp

This resource allows you to execute Check Point Resource Smtp.

## Example Usage


```hcl
resource "checkpoint_management_resource_smtp" "smtp" {

    name = "newSmtpResource"
    mail_delivery_server = "deliverServer"
    exception_track = "exception log"
    match = {
      sender = "expr1"
      recipient = "expr2"
    }
    action_1 {
      sender {
        original = "one"
        rewritten = "two"
      }
      recipient {
        original = "three"
        rewritten = "four"
      }
      custom_field{
        field = "field"
        original = "five"
        rewritten = "six"
      } 
    }
  cvp {
    allowed_to_modify_content = true
    enable_cvp =  false
    reply_order = "return_data_after_content_is_approved"
    server = "serverName"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `mail_delivery_server` - (Optional) Specify the server to which mail is forwarded. 
* `deliver_messages_using_dns_mx_records` - (Optional) MX record resolving is used to set the destination IP address of the connection. 
* `check_rulebase_with_new_destination` - (Optional) The Rule Base will be rechecked with the new resolved IP address for mail delivery. 
* `notify_sender_on_error` - (Optional) Enable error mail delivery. 
* `error_mail_delivery_server` - (Optional) Error mail delivery happens if the SMTP security server is unable to deliver the message within the abandon time, and Notify Sender on Error is checked. 
* `error_deliver_messages_using_dns_mx_records` - (Optional) MX record resolving will be used to set the source IP address of the connection used to send the error message. 
* `error_check_rulebase_with_new_destination` - (Optional) The Rule Base will be rechecked with the new resolved IP address for error mail delivery. 
* `exception_track` - (Optional) Determines if an action specified in the Action 2 and CVP categories taken as a result of a resource definition is logged. 
* `match` - (Optional) Set the Match properties for the SMTP resource.match blocks are documented below.
* `action_1` - (Optional) Use the Rewriting Rules to rewrite Sender and Recipient headers in emails, you can also rewrite other email headers by using the custom header field.action_1 blocks are documented below.
* `action_2` - (Optional) Use this window to configure mail inspection for the SMTP Resource.action_2 blocks are documented below.
* `cvp` - (Optional) Configure CVP inspection on mail messages.cvp blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`match` supports the following:

* `sender` - (Optional) Set the Match sender property for the SMTP resource. 
* `recipient` - (Optional) Set the Match recipient property for the SMTP resource. 


`action_1` supports the following:

* `sender` - (Optional) Rewrite Sender header.sender blocks are documented below.
* `recipient` - (Optional) Rewrite Recipient header.recipient blocks are documented below.
* `custom_field` - (Optional) The name of the header.custom_field blocks are documented below.


`action_2` supports the following:

* `strip_mime_of_type` - (Optional) Specifies the MIME type to strip from the message. 
* `strip_file_by_name` - (Optional) Strips file attachments of the specified name from the message. 
* `mail_capacity` - (Optional) Restrict the size (in kb) of incoming email attachments. 
* `allowed_characters` - (Optional) The MIME email headers can consist of 8 or 7 bit characters (7 ASCII and 8 for sending Binary characters) in order to encode mail data. 
* `strip_script_tags` - (Optional) Strip JAVA scripts. 
* `strip_applet_tags` - (Optional) Strip JAVA applets. 
* `strip_activex_tags` - (Optional) Strip activeX tags. 
* `strip_ftp_links` - (Optional) Strip ftp links. 
* `strip_port_strings` - (Optional) Strip ports. 


`cvp` supports the following:

* `enable_cvp` - (Optional) Select to enable the Content Vectoring Protocol. 
* `server` - (Optional) The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application. 
* `allowed_to_modify_content` - (Optional) Configures the CVP server to inspect but not modify content. 
* `reply_order` - (Optional) Designates when the CVP server returns data to the Security Gateway security server. 


`sender` supports the following:

* `original` - (Optional) Original field. 
* `rewritten` - (Optional) Replacement field. 


`recipient` supports the following:

* `original` - (Optional) Original field. 
* `rewritten` - (Optional) Replacement field. 


`custom_field` supports the following:

* `original` - (Optional) Original field. 
* `rewritten` - (Optional) Replacement field. 
* `field` - (Optional) The name of the header. 
