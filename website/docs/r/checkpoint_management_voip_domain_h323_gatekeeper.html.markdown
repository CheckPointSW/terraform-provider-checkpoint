---
layout: "checkpoint"
page_title: "checkpoint_management_voip_domain_h323_gatekeeper"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-voip-domain-h323-gatekeeper"
description: |-
 This resource allows you to execute Check Point Voip Domain H323 Gatekeeper.
---

# checkpoint_management_voip_domain_h323_gatekeeper

This resource allows you to execute Check Point Voip Domain H323 Gatekeeper.

## Example Usage


```hcl
resource "checkpoint_management_voip_domain_h323_gatekeeper" "example" {
  name = "vdhg1"
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
* `routing_mode` - (Optional) The routing mode of the VoIP Domain H323 gatekeeper. routing_mode blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`routing_mode` supports the following:

* `direct` - (Optional) Indicates whether the routing mode is direct. 
* `call_setup` - (Optional) Indicates whether the routing mode includes call setup (Q.931). 
* `call_setup_and_call_control` - (Optional) Indicates whether the routing mode includes both call setup (Q.931) and call control (H.245). 
