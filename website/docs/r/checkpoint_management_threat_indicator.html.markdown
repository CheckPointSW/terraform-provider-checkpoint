---
layout: "checkpoint"
page_title: "checkpoint_management_threat_indicator"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-threat_indicator"
description: |-
  This resource allows you to add/update/delete Check Point Threat Indicator.
---

# checkpoint_management_threat_indicator

This resource allows you to add/update/delete Check Point Threat Indicator.

## Example Usage


```hcl
resource "checkpoint_management_threat_indicator" "example" {
    name = "My_Indicator"
    observables {
      name = "My_Observable"
      mail_to = "someone@somewhere.com"
      confidence = "medium"
      severity = "low"
      product = "AV"
    } 
    action = "ask"
    profile_overrides {
      profile = "My_Profile"
      action = "detect"
    } 
    ignore_warnings = true
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `observables` - (Optional) The indicator's observables. Indicator's observables blocks are documented below.
* `action` - (Optional) The indicator's action.
* `profile_overrides` - (Optional) Profiles in which to override the indicator's default action. Profile Overrides blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `tags` - (Optional) Collection of tag identifiers.


`observables` supports the following:

* `name` - (Required) Object name. Should be unique in the domain.
* `md5` - (Optional) A valid MD5 sequence.
* `url` - (Optional) A valid URL.
* `ip_address` - (Optional) A valid IP-Address.
* `ip_address_first` - (Optional) A valid IP-Address, the beginning of the range. If you configure this parameter with a value, you must also configure the value of the 'ip-address-last' parameter.
* `ip_address_last` - (Optional) A valid IP-Address, the end of the range. If you configure this parameter with a value, you must also configure the value of the 'ip-address-first' parameter.
* `domain` - (Optional) The name of a domain.
* `mail_to` - (Optional) A valid E-Mail address, recipient filed.
* `mail_from` - (Optional) A valid E-Mail address, sender field.
* `mail_cc` - (Optional) A valid E-Mail address, cc field.
* `mail_reply_to` - (Optional) A valid E-Mail address, reply-to field.
* `mail_subject` - (Optional) Subject of E-Mail.
* `confidence` - (Optional) The confidence level the indicator has that a real threat has been uncovered.
* `product` - (Optional) The software blade that processes the observable: AV - AntiVirus, AB - AntiBot.
* `severity` - (Optional) The severity level of the threat.

`profile_overrides` supports the following:

* `action` - (Optional) The indicator's action in this profile.
* `profile` - (Optional) The profile in which to override the indicator's action.



