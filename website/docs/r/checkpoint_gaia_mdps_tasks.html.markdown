---
layout: "checkpoint"
page_title: "checkpoint_gaia_mdps_tasks"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-mdps-tasks"
description: |-
This resource allows you to execute Check Point Mdps Tasks.
---

# checkpoint_gaia_mdps_tasks

This resource allows you to execute Check Point Mdps Tasks.

## Example Usage


```hcl
resource "checkpoint_gaia_mdps_tasks" "example" {
  cp_port_protocol {
    port     = 18191
    protocol = "tcp"
  }
  cp_port_protocol {
    port     = 18192
    protocol = "tcp"
  }
  os_service = ["sshd"]
}
```

## Argument Reference

The following arguments are supported:

* `external_address` - (Optional) External address to communicate with via the Management plane external_address blocks are documented below.
* `os_service` - (Optional) OS Service to run on Management Plane, see 'chkconfig --list' (R82 and below) or 'systemctl list-units --type=service' (R82.10 and above) os_service blocks are documented below.
* `os_process` - (Optional) OS Process to run on Management Plane (see <a href='https://support.checkpoint.com/results/sk/sk97638'>sk97638</a> for more information). os_process blocks are documented below.
* `cp_port_protocol` - (Optional) Check Point Port and Protocol to use on the Management plane (see <a href='https://support.checkpoint.com/results/sk/sk52421'>sk52421</a> for more information). cp_port_protocol blocks are documented below.
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`external_address` supports the following:



`os_service` supports the following:



`os_process` supports the following:



`cp_port_protocol` supports the following:

* `port` - (Optional) Port number 
* `protocol` - (Optional) Protocol type 
