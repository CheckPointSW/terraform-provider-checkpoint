---
layout: "checkpoint"
page_title: "checkpoint_gaia_virtual_gateway"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-virtual-gateway"
description: |-
This resource allows you to execute Check Point Virtual Gateway.
---

# checkpoint_gaia_virtual_gateway

This resource allows you to execute Check Point Virtual Gateway.

## Example Usage


```hcl
resource "checkpoint_gaia_virtual_gateway" "example" {
  resource_id = 1
}
```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Optional) Virtual gateway identifier can be an integer or the next avaliable id (auto) 
* `one_time_password` - (Optional) One time password, used for the secure internal communication with the Management object 
* `interfaces` - (Optional) Network interface(s) to be attached interfaces blocks are documented below.
* `virtual_switches` - (Optional) Virtual switche(s) to be connected, mgmt-switch (id 500) is set as default virtual_switches blocks are documented below.
* `resources` - (Optional) Additional resources resources blocks are documented below.
* `mgmt_connection` - (Optional) Management connection configurations mgmt_connection blocks are documented below.
* `set_if_exist` - (Optional) If another virtual gateway with the same identifier already exists, it will be updated. The command behaviour will be the same as if originally a set command was called. Pay attention that original virtual gateway's fields will be overwritten by the fields provided in the request payload! 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `action` - (Computed) Computed field, returned in the response. 
* `status` - (Computed) Computed field, returned in the response. 
* `message` - (Computed) Computed field, returned in the response. 
* `vsxd_task_id` - (Computed) Computed field, returned in the response. 
* `vs_id` - (Computed) Computed field, returned in the response. 


`resources` supports the following:

* `firewall_ipv4_instances` - (Optional) CoreXL IPv4 instances amount. Must be between 1 and the greater of 32 and the number of CPU cores. 
* `firewall_ipv6_instances` - (Optional) CoreXL IPv6 instances amount. Must be between 0 and the greater of 32 and the number of CPU cores. Must not exceed the number of IPv4 CoreXL instances. 


`mgmt_connection` supports the following:

* `mgmt_connection_identifier` - (Optional) Management connection identifier according to the connection type (interface or virtual-switch (id or name)) 
* `mgmt_connection_type` - (Optional) Management connection type - interface or virtual link connected to virtual-switch (identified by name or id) 
* `mgmt_ipv4_configuration` - (Optional) Management IPv4 configuration mgmt_ipv4_configuration blocks are documented below.
* `mgmt_ipv6_configuration` - (Optional) Management IPv6 configuration mgmt_ipv6_configuration blocks are documented below.


`mgmt_ipv4_configuration` supports the following:

* `ipv4_address` - (Optional) IPv4 address 
* `ipv4_mask` - (Optional) IPv4 mask 
* `ipv4_default_gateway` - (Optional) IPv4 default gateway 


`mgmt_ipv6_configuration` supports the following:

* `ipv6_address` - (Optional) IPv6 address 
* `ipv6_mask` - (Optional) IPv6 mask length 
* `ipv6_default_gateway` - (Optional) IPv6 default gateway 
