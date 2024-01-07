---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-accounts"
description: |- Use this data source to get information on all Check Point CME Accounts.
---

# checkpoint_management_cme_accounts

Use this data source to get information on all Check Point CME Accounts.

## Example Usage

```hcl
data "checkpoint_management_cme_accounts" "accounts" {
}
```

## Argument Reference

The following arguments are supported:

* `result` - List of all accounts, each with the following data:
  * `name` - Unique account name for identification.
  * `platform` - The platform of the account.
  * `gw_configurations` - A list of GW configurations attached to the account.
  * `deletion_tolerance` - The number of CME cycles to wait when the cloud provider does not return a GW until its deletion.
  * `domain` - The account's domain name in MDS environment.

Note: For getting the full data for each account, use the specific data source of the account platform (checkpoint_management_cme_accounts_<aws/azure/gcp>).