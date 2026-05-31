---
layout: "checkpoint"
page_title: "checkpoint_gaia_remote_syslog"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-remote-syslog"
description: |-
This resource allows you to execute Check Point Remote Syslog.
---

# checkpoint_gaia_remote_syslog

This resource allows you to execute Check Point Remote Syslog.

## Example Usage


```hcl
resource "checkpoint_gaia_remote_syslog" "example" {
  level = "debug"
  server_ip = "10::130"
}
```

## Argument Reference

The following arguments are supported:

* `server_ip` - (Required) Remote server address, IPv6 and Hostname supported from R82. 
* `level` - (Required)  
* `port` - (Optional) Log port. Supported starting from Gaia version R81.20 
* `protocol` - (Optional) Log protocol. Supported starting from Gaia version R81.20 
* `queuing_mechanism` - (Optional) Log queuing mechanism state. Supported starting from Gaia version R82 
* `tls_encryption` - (Optional) TLS encryption status. Supported starting from Gaia version R82 tls_encryption blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`tls_encryption` supports the following:

* `enabled` - (Optional) TLS Encryption state. Supported starting from Gaia version R82 
* `auth_mode` - (Optional) Mode used for TLS authentication. supported modes:
          name - certificate validation and subject name authentication. Most secure mode.
          fingerprint - fingerprint of the server's certificate (which can be self-signed).
          certvalid - server's certificate validation only.
            It validates the remote peers certificate, but does not check the subject name.
            It is recommended NOT to use this mode.
          anon - anonymous authentication, this mode is vulnerable to man in the middle attacks as well as unauthorized access.
            It is recommended NOT to use this mode.
        Note:
          When setting anon or certvalid modes, permitted-peers of the remote host will removed.
          When setting name or fingerprint modes, it is mandatory to set permitted-peers for the remote host. Supported starting from Gaia version R82 
* `permitted_peers` - (Optional) Common name CN or fingerprint of the permitted peer.
         In case of fingerprint, Accepted SHA1.
         To specify multiple remote peers separate each one with a comma.
         This parameter is mandatory when using auth-mode of 'name' or 'fingerprint'.

         Note that usually a single remote peer should be all you need.
         Support for multiple peers is primarily included in support of load balancing scenarios.
         If the connection goes to a specific server, only one specific peer is ever expected. Supported starting from Gaia version R82 permitted_peers blocks are documented below.
