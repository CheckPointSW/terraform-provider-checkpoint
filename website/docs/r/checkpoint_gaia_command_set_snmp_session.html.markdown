---
layout: "checkpoint"
page_title: "checkpoint_gaia_command_set_snmp_session"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-command-set-snmp-session"
description: |-
This resource allows you to execute Check Point Set Snmp Session.
---

# checkpoint_gaia_command_set_snmp_session

This resource allows you to execute Check Point Set Snmp Session.

## Example Usage


```hcl
# Step 1: enable the SNMP agent
resource "checkpoint_gaia_snmp" "snmp_setup" {
  enabled    = true
  version    = "any"
  interfaces = "all"
}

# Step 2: configure the SNMP session
resource "checkpoint_gaia_command_set_snmp_session" "example" {
  session_timeout = 100
  v3_object {
    name = "testuser"
    authentication {
      protocol = "SHA256"
      password = "MySecurePass123!"
    }
    data_privacy = false
  }

  depends_on = [checkpoint_gaia_snmp.snmp_setup]
}
```

## Argument Reference

The following arguments are supported:

* `community_string` - (Optional) SNMP v2 community password. 
                <b>required for SNMP v1/v2</b> 
* `v3_object` - (Optional) SNMPv3 USM (User-based Security Model) details 
                      <b>required for SNMP v3</b> 
                      <b>Preferred</b> v3_object blocks are documented below.
* `session_timeout` - (Optional) Session expiration timeout in seconds 


`v3_object` supports the following:

* `name` - (Optional) SNMPv3 USM user 
* `authentication` - (Optional) Authentication details authentication blocks are documented below.
* `data_privacy` - (Optional) Related to AutoPriv/AutnNoPriv in SecurityLevel in the RFC. True: AutoPriv False: AuthNoPriv 


`authentication` supports the following:

* `protocol` - (Optional) Authentication protocol, MD5 and SHA1 are not supported starting from R81 
* `password` - (Optional) Authentication Password - (8 or more printable characters) Each SNMPv3 USM user must have an authentication pass phrase. This will be used by the SNMPv3 agent to verify the identity of the user before granting access. 


## How To Use
Make sure this command will be executed in the right execution order.
note: terraform execution is not sequential.

