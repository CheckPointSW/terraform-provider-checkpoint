---
layout: "checkpoint"
page_title: "checkpoint_management_set_threat_protection"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-threat-protection"
description: |-
This resource allows you to execute Check Point Set Threat Protection.
---

# checkpoint_management_set_threat_protection

This resource allows you to execute Check Point Set Threat Protection.

## Example Usage


```hcl
resource "checkpoint_management_set_threat_protection" "example" {
  name = "FTP Commands"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `comments` - (Optional) Protection comments. 
* `follow_up` - (Optional) Tag the protection with pre-defined follow-up flag. 
* `overrides` - (Optional) Overrides per profile for this protection<br> Note: Remove override for Core protections removes only the actions override. Remove override for Threat Cloud protections removes the action, track and packet captures.overrides blocks are documented below.


`overrides` supports the following:

* `action` - (Optional) Protection action. 
* `profile` - (Optional) Profile name. 
* `capture_packets` - (Optional) Capture packets. 
* `track` - (Optional) Tracking method for protection. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

