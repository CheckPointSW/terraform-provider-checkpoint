---
layout: "checkpoint"
page_title: "checkpoint_management_cme_api_versions"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-api-versions"
description: |- Use this data source to get information on existing Check Point CME API versions.
---

# Data Source: checkpoint_management_cme_api_versions

Use this data source to get information on existing Check Point CME API versions.

Available in:

- Check Point Security Management/Multi-Domain Security Management Server R81.10 and higher.
- CME Take 255 and higher.

## Example Usage

```hcl
data "checkpoint_management_cme_api_versions" "api_versions" {
}
```

## Argument Reference

These arguments are supported:

* `current_version` - Current CME API version.
* `supported_versions` - CME supported versions.
