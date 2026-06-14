---
layout: "checkpoint"
page_title: "checkpoint_gaia_ssh_server_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-ssh-server-settings"
description: |-
This resource allows you to execute Check Point Ssh Server Settings.
---

# checkpoint_gaia_ssh_server_settings

This resource allows you to execute Check Point Ssh Server Settings.

## Example Usage


```hcl
resource "checkpoint_gaia_ssh_server_settings" "example" {
  password_authentication = true
  permit_root_login = true
  use_dns = false
  client_alive_interval = 0
  login_grace_time = 120
}
```

## Argument Reference

The following arguments are supported:

* `enabled_ciphers` - (Optional) Specifies the SSH ciphers that are enabled. Ciphers are encryption algorithms used to secure SSH connections. enabled_ciphers blocks are documented below.
* `enabled_mac_algorithms` - (Optional) Specifies the SSH MAC (Message Authentication Code) algorithms that are enabled. These algorithms ensure data integrity and authenticity during SSH communication. enabled_mac_algorithms blocks are documented below.
* `enabled_kex_algorithms` - (Optional) Specifies the SSH key exchange (KEX) algorithms that are enabled. These algorithms are used to securely exchange cryptographic keys between the client and server. enabled_kex_algorithms blocks are documented below.
* `enabled_public_key_algorithms` - (Optional) Specifies the SSH public key algorithms that are enabled. These algorithms are used for authenticating the client to the server using public key cryptography. enabled_public_key_algorithms blocks are documented below.
* `password_authentication` - (Optional) Enables or disables password authentication. When enabled, users can authenticate using a password. 
* `permit_root_login` - (Optional) Enables or disables root login. When enabled, the root user is allowed to log in directly. 
* `use_dns` - (Optional) Enables or disables DNS usage. When enabled, the server performs a reverse DNS lookup to resolve the client's IP to a hostname. 
* `client_alive_interval` - (Optional) Sets the interval (in seconds) for sending alive messages to the client. This helps in keeping the connection active and detecting unresponsive clients. 
* `login_grace_time` - (Optional) Sets the time (in seconds) allowed for a user to successfully log in. If the user fails to log in within this time, the server disconnects the session. 
* `include_disabled_values` - (Computed) Include disabled algorithms 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`enabled_ciphers` supports the following:



`enabled_mac_algorithms` supports the following:



`enabled_kex_algorithms` supports the following:



`enabled_public_key_algorithms` supports the following:

