---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_gcp"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-accounts-gcp"
description: |- Use this data source to get information on an existing Check Point CME GCP Account.
---

# Data Source: checkpoint_management_cme_accounts_gcp

Use this data source to get information on an existing Check Point CME GCP Account.

Available in:

- Check Point Security Management/Multi Domain Management Server R81.10 and higher.
- CME take 255 and higher.

## Example Usage

```hcl
data "checkpoint_management_cme_accounts_gcp" "gcp_account" {
  name = "gcpAccount"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique account name for identification.
* `project_id` - GCP project id.
* `credentials_file` - The credentials file.
* `credentials_data` - Base64 encoded string that represents the content of the credentials file.
* `deletion_tolerance` - The number of CME cycles to wait when the cloud provider does not return a GW until its
  deletion.
* `domain` - The account's domain name in MDS environment.
* `platform` - The platform of the account.
* `gw_configurations` - A list of GW configurations attached to the account.
