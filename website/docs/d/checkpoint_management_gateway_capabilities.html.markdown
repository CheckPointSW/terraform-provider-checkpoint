---
layout: "checkpoint"
page_title: "checkpoint_management_gateway_capabilities"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-gateway_capabilities"
description: |-
Use this data source to get information on an existing Check Point Gateway Capabilities.
---

# checkpoint_management_gateway_capabilities

Use this data source to get information on an existing Check Point Gateway Capabilities.

## Example Usage


```hcl
data "checkpoint_management_gateway_capabilities" "data" {
  hardware = "CloudGuard IaaS"
  platform = "other"
  version = "R82"
}
```

## Argument Reference

The following arguments are supported:

* `hardware` - (Optional) Check Point hardware.
* `platform` - (Optional) Check Point gateway platform.
* `version` - (Optional) Gateway platform version.
* `restrictions` - Set of restrictions.
* `supported_blades` - Supported blades according to restrictions.
* `supported_firmware_platforms` - Supported firmware platforms according to restrictions.
* `supported_hardware` - Supported hardware according to restrictions.
* `supported_platforms` - Supported platforms according to restrictions.
* `supported_versions` - Supported versions according to restrictions.

`restrictions` supports the following: 

* `hardware` -  Check Point hardware.
* `platform` - Check Point gateway platform.
* `version` - Gateway platform version.

`supported_blades` supports the following:

* `management` - Management blades.
* `network_security` - Network Security blades.
* `threat_prevention` - Threat Prevention blades.

`management` supports the following:

* `default` - N/A
* `name` - N/A
* `readonly` - N/A

`network_security` supports the following:

* `default` - N/A
* `name` - N/A
* `readonly` - N/A

`threat_prevention` supports the following:

* `autonomous` - N/A
* `custom` - N/A

`autonomous` supports the following:

* `default` - N/A
* `name` - N/A
* `readonly` - N/A

`custom` supports the following:

* `default` - N/A
* `name` - N/A
* `readonly` - N/A
* `readonly` - N/A

`supported_firmware_platforms` supports the following:

* `default` - Default gateway firmware platform.
* `firmwarePlatforms` - List of gateway firmware platforms.

`supported_hardware` supports the following:

* `default` - Default hardware.
* `hardware`- List of Check Point hardware.

`supported_platforms` supports the following:

* `default` - Default platform.
* `platforms`- List of Check Point gateway platforms.

`supported_versions` supports the following:

* `default` - Default gateway platform version.
* `versions`- List of gateway platform versions.

## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

