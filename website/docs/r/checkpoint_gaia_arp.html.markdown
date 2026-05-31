---
layout: "checkpoint"
page_title: "checkpoint_gaia_arp"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-arp"
description: |-
This resource allows you to execute Check Point Arp.
---

# checkpoint_gaia_arp

This resource allows you to execute Check Point Arp.

## Example Usage


```hcl
resource "checkpoint_gaia_arp" "example" {
  proxy {
    interface    = "eth0"
    ipv4_address = "172.23.22.200"
  }
  proxy {
    ipv4_address = "172.23.22.202"
    mac_address  = "1c:a8:5d:ae:f9:81"
  }
  static {
    ipv4_address = "172.23.22.201"
    mac_address  = "1c:a8:5d:ae:f9:81"
  }
}
```

## Argument Reference

The following arguments are supported:

* `settings` - (Optional) Configure ARP settings settings blocks are documented below.
* `proxy` - (Optional) Add a specified Proxy ARP entries proxy blocks are documented below.
* `static` - (Optional) Add a specified Static ARP entries static blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`settings` supports the following:

* `restriction_level` - (Optional) Define different restriction levels for announcing the local source IP address
         from IP packets in ARP requests sent on interface:
         0 - Any local address
         1 - Use address from the same subnet as the target address
         2 - Prefer primary address 
* `cache_size` - (Optional) Specify the maximum number of entries in the arp cache 
* `validity_timeout` - (Optional) Specify time, in seconds, to keep resolved dynamic ARP entries.         If the entry is not referred to and is not used by traffic before the time elapses, it is deleted.        Otherwise, a request will be sent to verify the MAC address. 
* `auto_cache_size` - (Optional) Update cache size to be automatically changed depending on the current ARP table entries in the system. 
ARP table default size: 4096, Range: 1024-131072, Supported starting from R82.00. Supported starting from Gaia version R82 


`proxy` supports the following:

* `ipv4_address` - (Optional) Define the IP address of a new proxy ARP entry 
* `interface` - (Optional) Define the interface used when forwarding packets to the given IP address 
* `mac_address` - (Optional) Define the hardware address used when forwarding packets to the given IP address 
* `real_ipv4_address` - (Optional) Define the real IP address used when forwarding packets to the given IP address 


`static` supports the following:

* `ipv4_address` - (Optional) Define the IP address of a new static ARP entry 
* `mac_address` - (Optional) Specify the hardware address used when forwarding packets to the given IP address 
