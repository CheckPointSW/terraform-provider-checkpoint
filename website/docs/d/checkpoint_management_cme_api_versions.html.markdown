---
layout: "checkpoint"
page_title: "checkpoint_management_cme_api_versions"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-api-versions"
description: |- Use this data source to get information on existing Check Point CME API versions.
---

# Data Source: checkpoint_management_cme_api_versions

Use this data source to get information on existing Check Point CME API versions.

## Example Usage

```hcl
data "checkpoint_management_cme_api_versions" "api_versions" {
}
```

## Argument Reference

The following arguments are supported:

* `current_version` - Current CME API version.
* `supported_versions` - CME supported versions.
