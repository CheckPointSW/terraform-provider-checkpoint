---
layout: "checkpoint"
page_title: "checkpoint_management_resource_smtp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-smtp"
description: |-
Use this data source to get information on an existing Check Point Resource Smtp.
---

# Data Source: checkpoint_management_resource_smtp

Use this data source to get information on an existing Check Point Resource Smtp.

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

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name.
* `mail_delivery_server` -  Specify the server to which mail is forwarded. 
* `deliver_messages_using_dns_mx_records` -  MX record resolving is used to set the destination IP address of the connection. 
* `check_rulebase_with_new_destination` -  The Rule Base will be rechecked with the new resolved IP address for mail delivery. 
* `notify_sender_on_error` -  Enable error mail delivery. 
* `error_mail_delivery_server` - Error mail delivery happens if the SMTP security server is unable to deliver the message within the abandon time, and Notify Sender on Error is checked. 
* `error_deliver_messages_using_dns_mx_records` -  MX record resolving will be used to set the source IP address of the connection used to send the error message. 
* `error_check_rulebase_with_new_destination` -  The Rule Base will be rechecked with the new resolved IP address for error mail delivery. 
* `exception_track` -  Determines if an action specified in the Action 2 and CVP categories taken as a result of a resource definition is logged. 
* `match` -  Set the Match properties for the SMTP resource.match blocks are documented below.
* `action_1` - Use the Rewriting Rules to rewrite Sender and Recipient headers in emails, you can also rewrite other email headers by using the custom header field.action_1 blocks are documented below.
* `action_2` - Use this window to configure mail inspection for the SMTP Resource.action_2 blocks are documented below.
* `cvp` -  Configure CVP inspection on mail messages.cvp blocks are documented below.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` -  Comments string. 



`match` supports the following:

* `sender` -  Set the Match sender property for the SMTP resource. 
* `recipient` -  Set the Match recipient property for the SMTP resource. 


`action_1` supports the following:

* `sender` -  Rewrite Sender header.sender blocks are documented below.
* `recipient` -  Rewrite Recipient header.recipient blocks are documented below.
* `custom_field` -  The name of the header.custom_field blocks are documented below.


`action_2` supports the following:

* `strip_mime_of_type` -  Specifies the MIME type to strip from the message. 
* `strip_file_by_name` -  Strips file attachments of the specified name from the message. 
* `mail_capacity` -  Restrict the size (in kb) of incoming email attachments. 
* `allowed_characters` -  The MIME email headers can consist of 8 or 7 bit characters (7 ASCII and 8 for sending Binary characters) in order to encode mail data. 
* `strip_script_tags` -  Strip JAVA scripts. 
* `strip_applet_tags` -  Strip JAVA applets. 
* `strip_activex_tags` -  Strip activeX tags. 
* `strip_ftp_links` -  Strip ftp links. 
* `strip_port_strings` - Strip ports. 


`cvp` supports the following:

* `enable_cvp` -  Select to enable the Content Vectoring Protocol. 
* `server` -  The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application. 
* `allowed_to_modify_content` -  Configures the CVP server to inspect but not modify content. 
* `reply_order` - Designates when the CVP server returns data to the Security Gateway security server. 


`sender` supports the following:

* `original` -  Original field. 
* `rewritten` -  Replacement field. 


`recipient` supports the following:

* `original` -  Original field. 
* `rewritten` -  Replacement field. 


`custom_field` supports the following:

* `original` -  Original field. 
* `rewritten` -  Replacement field. 
* `field` -  The name of the header. 
