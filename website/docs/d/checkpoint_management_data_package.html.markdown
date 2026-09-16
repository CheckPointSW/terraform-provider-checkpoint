---
layout: "checkpoint"
page_title: "checkpoint_management_data_package "
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-package"
description: |-
  Use this data source to get information on an existing Check Point Package Object.
---

# Data Source: checkpoint_management_data_package

Use this data source to get information on an existing Check Point Package Object.

## Example Usage

```hcl
resource "checkpoint_management_package" "package" {
    name = "My Package"
}

data "checkpoint_management_data_package" "data_package" {
    name = "${checkpoint_management_package.package.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier. 
* `show_installation_targets` - (Optional) Indicates whether to calculate and show "installation-targets" field in reply.
* `access` - True - enables, False - disables access & NAT policies, empty - nothing is changed.
* `desktop_security` - True - enables, False - disables Desktop security policy, empty - nothing is changed.
* `qos` - True - enables, False - disables QoS policy, empty - nothing is changed.
* `qos_policy_type` - QoS policy type.
* `threat_prevention` - True - enables, False - disables Threat policy, empty - nothing is changed.
* `vpn_traditional_mode` - True - enables, False - disables VPN traditional mode, empty - nothing is changed.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.
* `autonomous_threat_policy` - UID of the Autonomous Threat Prevention policy.
* `https_inspection_policy` - True - enables, False - disables HTTPS Inspection policy, empty - nothing is changed.
* `installation_targets_revision` - List of installation targets and revisions on which this policy package was installed.installation_targets_revision blocks are documented below.
* `nat_layer` - UID of the NAT policy.
* `nat_policy` - True - enables, False - disables NAT policy, empty - nothing is changed.
* `sd_wan` - True - enables, False - disables SD-WAN policy, empty - nothing is changed.
* `sd_wan_layer` - SD-WAN policy layer. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Stan...


`installation_targets_revision` supports the following:

* `cluster_members_revision` - If this target is a cluster, this list shows a revision which was installed on each cluster member.cluster_members_revision blocks are documented below.
* `revision` - The revision installed on this target. Level of details in the output corresponds to the number of details for search. This table shows the level of d...
* `target_name` - The name of the installation target.
* `target_uid` - Installation target unique identifier.


`cluster_members_revision` supports the following:

* `revision` - The revision installed on this target. Level of details in the output corresponds to the number of details for search. This table shows the level of d...
* `target_name` - The name of the installation target.
* `target_uid` - Installation target unique identifier.
