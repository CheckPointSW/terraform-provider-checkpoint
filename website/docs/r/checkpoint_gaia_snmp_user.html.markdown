---
layout: "checkpoint"
page_title: "checkpoint_gaia_snmp_user"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-snmp-user"
description: |-
This resource allows you to execute Check Point Snmp User.
---

# checkpoint_gaia_snmp_user

This resource allows you to execute Check Point Snmp User.

## Example Usage


```hcl
resource "checkpoint_gaia_snmp_user" "example" {
  name = "test3"

  authentication {
    protocol = "SHA256"
    password = "Mypass123!"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) SNMPv3 USM user 
* `authentication` - (Required) Authentication details authentication blocks are documented below.
* `permission` - (Optional) User permission 
* `allowed_virtual_systems` - (Optional) Configured Virtual Devices allowed for the USM user - vsid range: 0-512 allowed_virtual_systems blocks are documented below.
* `privacy` - (Optional) Privacy details. If provided, data privacy (encryption) is enabled privacy blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
* `data_privacy` - (Computed) Computed field, returned in the response. 


`authentication` supports the following:

* `protocol` - (Optional) Authentication protocol, MD5 and SHA1 are not supported starting from R81 
* `password` - (Optional) Authentication Password - (8 or more printable characters) Each SNMPv3 USM user must have an authentication pass phrase. This will be used by the SNMPv3 agent to verify the identity of the user before granting access. 


`privacy` supports the following:

* `protocol` - (Optional) Privacy protocol 
* `password` - (Optional) Privacy Password - (8 or more printable characters) An SNMPv3 USM user with a privacy security level must have a privacy pass phrase. This will be used by the SNMPv3 agent to keep other parties from eavesdropping on the SNMP interaction.  
