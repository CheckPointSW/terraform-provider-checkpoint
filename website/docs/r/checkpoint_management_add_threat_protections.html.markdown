---
layout: "checkpoint"
page_title: "checkpoint_management_add_threat_protections"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-add-threat-protections"
description: |-
This resource allows you to execute Check Point Add Threat Protections.
---

# checkpoint_management_add_threat_protections

This resource allows you to execute Check Point Add Threat Protections.

## Example Usage


```hcl
resource "checkpoint_management_add_threat_protections" "example" {
  package_path = "/path/to/community.rules"
  package_format = "snort"
}
```

## Argument Reference

The following arguments are supported:

* `package_format` - (Optional) Protections package format. 
* `package_path` - (Optional) Protections package path. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

