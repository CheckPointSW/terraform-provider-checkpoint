---
layout: "checkpoint"
page_title: "checkpoint_management_literal"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-literal"
description: |-
  This resource allows you to add Check Point objects from a literal value.
---

# Resource: checkpoint_management_literal

This resource allows you to add Check Point objects directly from a literal string value. Based on the format of the `literal`, the matching object is created (or reused if it already exists). The object name is derived deterministically from the `literal` value, which is what lets the resource detect and reuse an existing object instead of creating a duplicate:

* IPv4 / IPv6 address - creates a `host` object named `host_<ip>` (e.g. the literal `192.0.2.1` creates a host named `host_192.0.2.1`).
* Network CIDR (e.g. `192.0.2.0/24`) - creates a `network` object named `network_<subnet>/<mask-length>` (e.g. `network_192.0.2.0/24`).
* `tcp/<port number>` or `udp/<port number>` - creates a `service-tcp` / `service-udp` object named `service_<protocol>_<port>` (e.g. `service_tcp_8080`).
* DNS Domain - a literal that starts with a `.` (e.g. `.www.example.com`) - creates a `dns-domain` object named after the domain itself (e.g. `.www.example.com`). The domain is created with `is-sub-domain` set to `false` by default.

## Example Usage


```hcl
# Host literal (IPv4)
resource "checkpoint_management_literal" "host_example" {
  literal = "192.0.2.1" # host_192.0.2.1
}

# Network literal (CIDR)
resource "checkpoint_management_literal" "network_example" {
  literal = "192.0.2.0/24" # network_192.0.2.0/24
}

# Service literal (tcp/udp)
resource "checkpoint_management_literal" "service_example" {
  literal = "tcp/8080" # service_tcp_8080
}

# DNS Domain literal (starts with '.')
resource "checkpoint_management_literal" "dns_domain_example" {
  literal = ".www.example.com"
}
```

## Argument Reference

The following arguments are supported:

* `literal` - (Optional) Literal represents IPv4 or IPv6 or Network CIDR or Service (tcp/<port number> or udp/<port number>) or DNS Domain (a name starting with '.').

## Attributes Reference

In addition to the arguments above, the following attributes are exported:

* `uid` - Object unique identifier.
* `name` - Object name.
* `type` - Object type (`host`, `network`, `service-tcp`, `service-udp` or `dns-domain`).

## Important Notes

* **The underlying object is not deleted.** Removing the resource from the configuration (or running `terraform destroy`) only removes it from the Terraform state - it does **not** delete the created object from the Check Point management server.
* Changing the `literal` after creation creates a new object instead of editing the existing one.
