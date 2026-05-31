---
layout: "checkpoint"
page_title: "checkpoint_gaia_maestro_port"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-maestro-port"
description: |-
This resource allows you to execute Check Point Maestro Port.
---

# checkpoint_gaia_maestro_port

This resource allows you to execute Check Point Maestro Port.

## Example Usage


```hcl
resource "checkpoint_gaia_maestro_port" "example" {
  type = "uplink"
  enabled = true
  mtu = 777
  qsfp_mode = "40G"
  auto_negotiation = true
}
```

## Argument Reference

The following arguments are supported:

* `resource_id` - (Optional) Port ID (e.g. '1/13/1') 
* `interface_name` - (Optional) Interface name in case this port is an Uplink or MGMT interface (e.g. 'eth1-25') 
* `enabled` - (Optional) Setting this to false will disable this port, setting to true will enable it. AKA 'admin state' 
* `mtu` - (Optional) MTU of this port 
* `auto_negotiation` - (Optional) If true, Auto Negotiation will be turned on, and vice versa 
* `qsfp_mode` - (Optional) Port QSFP mode. Valid values are: '4x10G', '4x25G', '25G', '40G', '100G' 
* `type` - (Optional) Port type. Valid values are: 'downlink', 'uplink', 'site_sync', 'ssm_sync', 'mgmt' 
