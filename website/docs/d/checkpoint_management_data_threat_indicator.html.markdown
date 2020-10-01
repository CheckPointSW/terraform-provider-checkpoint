---
layout: "checkpoint"
page_title: "checkpoint_management_data_threat_indicator"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat_indicator"
description: |-
  Use this data source to get information on an existing Check Point Threat Indicator.
---

# Data Source: checkpoint_management_data_threat_indicator

Use this data source to get information on an existing Check Point Threat Indicator.

## Example Usage


```hcl
resource "checkpoint_management_threat_indicator" "threat_indicator" {
    name = "threat indicator"
	observables {
    	name = "obs1"
    	ip_address = "5.4.7.1"
  	}
	ignore_warnings = true
}

data "checkpoint_management_data_threat_indicator" "data_threat_indicator" {
    name = "${checkpoint_management_threat_indicator.threat_indicator.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `action` - The indicator's action.
* `profile_overrides` - Profiles in which to override the indicator's default action. Profile Overrides blocks are documented below.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.


`profile_overrides` supports the following:

* `action` - The indicator's action in this profile.
* `profile` - The profile in which to override the indicator's action.