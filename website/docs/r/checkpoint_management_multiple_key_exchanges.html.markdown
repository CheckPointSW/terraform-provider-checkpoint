---
layout: "checkpoint"
page_title: "checkpoint_management_multiple_key_exchanges"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-multiple-key-exchanges"
description: |-
This resource allows you to execute Check Point Multiple Key Exchanges.
---

# checkpoint_management_multiple_key_exchanges

This resource allows you to execute Check Point Multiple Key Exchanges.

## Example Usage


```hcl
resource "checkpoint_management_multiple_key_exchanges" "example" {
  name = "Multiple Key Exchanges"
  key_exchange_methods = ["group-2"]
  additional_key_exchange_1_methods =  ["kyber-1024"]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `key_exchange_methods` - (Required) Key-Exchange methods to use. Can contain only Diffie-Hellman groups. Valid values: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24
* `additional_key_exchange_1_methods` - (Optional) Additional Key-Exchange 1 methods to use. Valid values: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24, kyber-512, kyber-768, kyber-1024, none
* `additional_key_exchange_2_methods` - (Optional) Additional Key-Exchange 2 methods to use. Valid values: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24, kyber-512, kyber-768, kyber-1024, none
* `additional_key_exchange_3_methods` - (Optional) Additional Key-Exchange 3 methods to use. Valid values: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24, kyber-512, kyber-768, kyber-1024, none
* `additional_key_exchange_4_methods` - (Optional) Additional Key-Exchange 4 methods to use. Valid values: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24, kyber-512, kyber-768, kyber-1024, none
* `additional_key_exchange_5_methods` - (Optional) Additional Key-Exchange 5 methods to use. Valid values: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24, kyber-512, kyber-768, kyber-1024, none
* `additional_key_exchange_6_methods` - (Optional) Additional Key-Exchange 6 methods to use. Valid values: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24, kyber-512, kyber-768, kyber-1024, none
* `additional_key_exchange_7_methods` - (Optional) Additional Key-Exchange 7 methods to use. Valid values: group-1, group-2, group-5, group-14, group-15, group-16, group-17, group-18, group-19, group-20, group-24, kyber-512, kyber-768, kyber-1024, none
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 

