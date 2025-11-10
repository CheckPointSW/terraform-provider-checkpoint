---
layout: "checkpoint"
page_title: "checkpoint_management_voip_domain_h323_gateway"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-voip-domain-h323-gateway"
description: |-
Use this data source to get information on an existing Check Point Voip Domain H323 Gateway.
---

# Data Source: checkpoint_management_voip_domain_h323_gateway

Use this data source to get information on an existing Check Point Voip Domain H323 Gateway.

## Example Usage
```hcl
data "checkpoint_management_voip_domain_h323_gateway" "data_test" {
    name = "h323_gateway1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `endpoints_domain` - The related endpoints domain to which the VoIP domain will connect. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `installed_at` - The machine the VoIP is installed at. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `routing_mode` - The routing mode of the VoIP Domain H323 gateway.routing_mode blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `icon` - Object icon.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.

`routing_mode` supports the following:

* `call_setup` - Indicates whether the routing mode includes call setup (Q.931).
* `call_setup_and_call_control` - Indicates whether the routing mode includes both call setup (Q.931) and call control (H.245).
