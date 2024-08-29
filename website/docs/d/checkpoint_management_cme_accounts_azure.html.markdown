---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_azure"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-accounts-azure"
description: |- Use this data source to get information on an existing Check Point CME Azure Account.
---

# Data Source: checkpoint_management_cme_accounts_azure

Use this data source to get information on an existing Check Point CME Azure Account.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](../index.html.markdown#compatibility-with-cme).


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
* `environmnet` - The Azure environmnet.
* `platform` - The platform of the account.
* `gw_configurations` - A list of Gateway configurations attached to the account.
