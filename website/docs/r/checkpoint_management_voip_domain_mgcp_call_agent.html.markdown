---
layout: "checkpoint"
page_title: "checkpoint_management_voip_domain_mgcp_call_agent"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-voip-domain-mgcp-call-agent"
description: |-
 This resource allows you to execute Check Point Voip Domain Mgcp Call Agent.
---

# checkpoint_management_voip_domain_mgcp_call_agent

This resource allows you to execute Check Point Voip Domain Mgcp Call Agent.

## Example Usage


```hcl
resource "checkpoint_management_voip_domain_mgcp_call_agent" "example" {
  name = "mgcp1"
  endpoints_domain = "new_group"
  installed_at = "test_host"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `endpoints_domain` - (Required) The related endpoints domain to which the VoIP domain will connect. 
Identified by name or UID. 
* `installed_at` - (Required) The machine the VoIP is installed at. 
Identified by name or UID. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
