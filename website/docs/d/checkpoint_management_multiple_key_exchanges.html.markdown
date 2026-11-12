---
layout: "checkpoint"
page_title: "checkpoint_management_multiple_key_exchanges"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-multiple-key-exchanges"
description: |-
Use this data source to get information on an existing Check Point Multiple Key Exchanges.
---

# checkpoint_management_multiple_key_exchanges

Use this data source to get information on an existing Check Point Multiple Key Exchanges.

## Example Usage


```hcl
resource "checkpoint_management_multiple_key_exchanges" "example" {
  name = "Multiple Key Exchanges"
  key_exchange_methods = ["group-2"]
  additional_key_exchange_1_methods =  ["kyber-1024"]
}

data "checkpoint_management_multiple_key_exchanges" "data" {
  name = "${checkpoint_management_multiple_key_exchanges.test.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `key_exchange_methods` -  Key-Exchange methods to use. Can contain only Diffie-Hellman groups. 
* `additional_key_exchange_1_methods` - Additional Key-Exchange 1 methods to use. 
* `additional_key_exchange_2_methods` -  Additional Key-Exchange 2 methods to use.
* `additional_key_exchange_3_methods` -  Additional Key-Exchange 3 methods to use.
* `additional_key_exchange_4_methods` -  Additional Key-Exchange 4 methods to use.
* `additional_key_exchange_5_methods` -  Additional Key-Exchange 5 methods to use.
* `additional_key_exchange_6_methods` -  Additional Key-Exchange 6 methods to use.
* `additional_key_exchange_7_methods` -  Additional Key-Exchange 7 methods to use.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` -  Comments string. 
