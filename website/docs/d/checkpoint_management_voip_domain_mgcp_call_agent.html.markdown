---
layout: "checkpoint"
page_title: "checkpoint_management_voip_domain_mgcp_call_agent"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-voip-domain-mgcp-call-agent"
description: |-
Use this data source to get information on an existing Check Point Voip Domain Mgcp Call Agent.
---

# Data Source: checkpoint_management_voip_domain_mgcp_call_agent

Use this data source to get information on an existing Check Point Voip Domain Mgcp Call Agent.

## Example Usage
```hcl
data "checkpoint_management_voip_domain_mgcp_call_agent" "data_test" {
    name = "mgcp_call_agent1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `endpoints_domain` - The related endpoints domain to which the VoIP domain will connect. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `installed_at` - The machine the VoIP is installed at. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `icon` - Object icon.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
