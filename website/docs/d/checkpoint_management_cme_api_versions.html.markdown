---
layout: "checkpoint"
page_title: "checkpoint_management_cme_api_versions"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-api-versions"
description: |- Use this data source to get information on existing Check Point CME API versions.
---

# Data Source: checkpoint_management_cme_api_versions

Use this data source to get information on existing Check Point CME API versions.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](../index.html.markdown#compatibility-with-cme).


## Example Usage

```hcl
data "checkpoint_management_cme_api_versions" "api_versions" {
}
```

## Argument Reference

These arguments are supported:

* `current_version` - Current CME API version.
* `supported_versions` - CME supported versions.
