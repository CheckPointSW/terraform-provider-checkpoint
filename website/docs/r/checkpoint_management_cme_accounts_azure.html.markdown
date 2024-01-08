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
  name           = "TestAzureController"
  directory_id   = "12345678-1234-1234-1234-123456789123"
  application_id = "12345678-1234-1234-1234-123456789123"
  client_secret  = "mySecret"
  subscription   = "12345678-1234-1234-1234-123456789123"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique account name for identification.
* `subscription` - (Required) Azure subscription ID.
* `directory_id` - (Required) Azure Active Directory tenant ID.
* `application_id` - (Required) The application ID with which the service principal is associated.
* `client_secret` - (Required) The service principal's client secret.
* `deletion_tolerance` - (Optional) The number of CME cycles to wait when the cloud provider does not return a GW until
  its deletion.
* `domain` - (Optional) The account's domain name in MDS environment.

## Limitations

`client_secret` attribute can be updated only through the created resources in terraform.