---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-accounts"
description: |- Use this data source to get information on all Check Point CME Accounts.
---

# Data Source: checkpoint_management_cme_accounts

Use this data source to get information on all Check Point CME Accounts.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](../index.html.markdown#compatibility-with-cme).


## Example Usage

```hcl
data "checkpoint_management_cme_accounts" "accounts" {
}
```

## Argument Reference

These arguments are supported:

* `result` - List of all accounts, each with this data:
    * `name` - Unique account name for identification.
    * `platform` - The platform of the account.
    * `gw_configurations` - A list of Gateway configurations attached to the account.
    * `deletion_tolerance` - The number of CME cycles to wait when the cloud provider does not return a GW until its
      deletion.
    * `domain` - The account's domain name in Multi-Domain Security Management Server environment.

Note: To get the full data for each account, use the specific data source of the account platform (checkpoint_management_cme_accounts_<aws/azure/gcp>).