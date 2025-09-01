---
layout: "checkpoint"
page_title: "checkpoint_management_securemote_dns_server"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-securemote-dns-server"
description: |-
This resource allows you to execute Check Point Securemote Dns Server.
---

# checkpoint_management_securemote_dns_server

This resource allows you to execute Check Point Securemote Dns Server.

## Example Usage


```hcl
resource "checkpoint_management_securemote_dns_server" "example" {
  name = "TestSecuRemoteDNSSever"
  host = "TestHost"
  domains {
    domain_suffix = ".com"
    maximum_prefix_label_count = 2
  }
    
  domains {
    domain_suffix = ".local"
    maximum_prefix_label_count = 3
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `host` - (Required) DNS server for remote clients in the Remote access community. 
Identified by name or UID. 
* `domains` - (Optional) The DNS domains that the remote clients can access. domains blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`domains` supports the following:

* `domain_suffix` - (Optional) DNS Domain suffix. 
* `maximum_prefix_label_count` - (Optional) Maximum number of matching labels preceding the suffix. 
