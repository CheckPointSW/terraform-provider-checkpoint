---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_azure"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-accounts-azure"
description: |- Use this data source to get information on an existing Check Point CME Azure Account.
---

# checkpoint_management_cme_accounts_azure

Use this data source to get information on an existing Check Point CME Azure Account.

## Example Usage

```hcl
data "checkpoint_management_cme_accounts_azure" "azure_account" {
  name = "azure-controller"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique account name for identification.
* `subscription` - Azure subscription ID.
* `directory_id` - Azure Active Directory tenant ID.
* `application_id` - The application ID with which the service principal is associated.
* `client_secret` - The service principal's client secret.
* `deletion_tolerance` - The number of CME cycles to wait when the cloud provider does not return a GW until its deletion.
* `domain` - The account's domain name in MDS environment.
* `platform` - The platform of the account.
* `gw_configurations` - A list of GW configurations attached to the account.
