---
layout: "checkpoint"
page_title: "checkpoint_management_securemote_dns_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-securemote-dns-server"
description: |- Use this data source to get information on an existing SecuRemote DNS Server.
---


# checkpoint_management_securemote_dns_server

Use this data source to get information on an existing SecurRemote DNS Server.

## Example Usage


```hcl
data "checkpoint_management_securemote_dns_server" "data_securemote_dns_server" {
  name = "TestSecuRemoteDNSSever"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `host` - DNS server for remote clients in the Remote access community. Identified by name or UID.
* `domains` - The DNS domains that the remote clients can access. domains blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.


`domains` supports the following:

* `domain_suffix` - DNS Domain suffix.
* `maximum_prefix_label_count` - Maximum number of matching labels preceding the suffix. 
