---
layout: "checkpoint"
page_title: "checkpoint_management_azur_ad"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-azur-ad"
description: |-
Use this data source to get information on an existing Check Point Azure Ad.
---

# Data Source: checkpoint_management_azure_ad

Use this data source to get information on an existing Check Point Azure Ad.

## Example Usage


```hcl
resource "checkpoint_management_azure_ad" "azure_ad" {
  name = "example"
  password = "123"
  user_authentication = "user-authentication"
  username = "example"
  application_id = "a8662b33-306f-42ba-9ffb-a0ac27c8903f"
  application_key = "EjdJ2JcNGpw3[GV8:PMN_s2KH]JhtlpO"
  directory_id = "19c063a8-3bee-4ea5-b984-e344asds37f7"
}

data "checkpoint_management_azure_ad" "data_azure_ad" {
  name = "${checkpoint_management_azure_ad.azure_ad.name}"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `properties` - Azure AD connection properties. properties blocks are documented below.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
