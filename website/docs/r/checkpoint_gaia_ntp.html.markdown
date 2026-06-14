---
layout: "checkpoint"
page_title: "checkpoint_gaia_ntp"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-ntp"
description: |-
This resource allows you to execute Check Point Ntp.
---

# checkpoint_gaia_ntp

This resource allows you to execute Check Point Ntp.

## Example Usage


```hcl
resource "checkpoint_gaia_ntp" "example" {
  enabled = true
  preferred = "2.2.2.2"
  servers {
    address = "4.4.4.4"
    type = "pool"
    version = 4
  }
  servers {
    address = "2.2.2.2"
    type = "server"
    version = 4
  }
}
```

## Argument Reference

The following arguments are supported:

* `enabled` - (Optional) NTP status 
* `servers` - (Optional) Add, set or remove NTP server/pool servers blocks are documented below.
* `preferred` - (Optional) Preferred address. Specify a particular server as preferred above others of similar statistical quality 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`servers` supports the following:

* `address` - (Optional)  
* `type` - (Optional) Address type. Should be server or pool (a dynamic collection of servers).
Relevant only from R82 (V1.8).
primary and secondary options are to support backward compatibility 
* `version` - (Optional)  
