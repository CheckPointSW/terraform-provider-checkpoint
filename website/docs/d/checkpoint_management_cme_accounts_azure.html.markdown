---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_azure"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-accounts-azure"
description: |- Use this data source to get information on an existing Check Point CME Azure Account.
---

# Data Source: checkpoint_management_cme_accounts_azure

Use this data source to get information on an existing Check Point CME Azure Account.

Available in:

- Check Point Security Management/Multi-Domain Security Management Server R81.10 and higher.
- CME Take 255 and higher.

## Example Usage

```hcl
data "checkpoint_management_cme_accounts_azure" "azure_account" {
  name = "azureAccount"
}
```

## Argument Reference

These arguments are supported:

* `name` - (Required) Unique account name for identification.
* `subscription` - Azure subscription ID.
* `directory_id` - Azure Active Directory tenant ID.
* `application_id` - The application ID related to the service principal.
* `client_secret` - The service principal's client secret.
* `deletion_tolerance` - The number of CME cycles to wait when the cloud provider does not return a Gateway until its
  deletion.
* `domain` - The account's domain name in Multi-Domain Security Management Server environment.
* `platform` - The platform of the account.
* `gw_configurations` - A list of Gateway configurations attached to the account.
