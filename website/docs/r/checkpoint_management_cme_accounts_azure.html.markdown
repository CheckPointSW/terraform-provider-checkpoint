---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_azure"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-accounts-azure"
description: |- This resource allows you to add/update/delete Check Point CME Azure Account.
---

# checkpoint_management_cme_accounts_azure

This resource allows you to add/update/delete Check Point CME Azure Account.

## Example Usage

```hcl
resource "checkpoint_management_cme_accounts_azure" "azure_account" {
  name           = "azureAccount"
  directory_id   = "abcd1234-ab12-cd34-ef56-abcdef123456"
  application_id = "abcd1234-ab12-cd34-ef56-abcdef123456"
  client_secret  = "mySecret"
  subscription   = "abcd1234-ab12-cd34-ef56-abcdef123456"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique account name for identification without spaces.
* `subscription` - (Required) Azure subscription ID.
* `directory_id` - (Required) Azure Active Directory tenant ID.
* `application_id` - (Required) The application ID with which the service principal is associated.
* `client_secret` - (Required) The service principal's client secret.
* `deletion_tolerance` - (Optional) The number of CME cycles to wait when the cloud provider does not return a GW until
  its deletion.
* `domain` - (Optional) The account's domain name in MDS environment.

## Limitations

`client_secret` attribute can be managed only through the created resources in terraform.
