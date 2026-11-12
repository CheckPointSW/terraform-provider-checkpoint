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


## Import

Use the existing resource's `resource_id` as the import ID.

```shell
terraform import checkpoint_gaia_maestro_port.example '1/13/1'
```

To specify optional read context or filters (`interface_name`), use a JSON object with Terraform attribute names and include all required lookup fields.

Create a matching resource block before importing, and review `terraform plan` afterward. Import reads existing configuration; it does not create the resource or recover write-only values such as passwords.
