---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_azure"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-accounts-azure"
description: |- This resource allows you to add/update/delete Check Point CME Azure Account.
---

# Resource: checkpoint_management_cme_accounts_azure

This resource allows you to add/update/delete Check Point CME Azure Account.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](https://registry.terraform.io/providers/CheckPointSW/checkpoint/latest/docs#compatibility-with-cme).


## Example Usage

```hcl
resource "checkpoint_management_cme_accounts_azure" "azure_account" {
  name           = "azureAccount"
  directory_id   = "abcd1234-ab12-cd34-ef56-abcdef123456"
  application_id = "abcd1234-ab12-cd34-ef56-abcdef123456"
  client_secret  = "mySecret"
  subscription   = "abcd1234-ab12-cd34-ef56-abcdef123456"
  environmnet    = "AzureCloud"
}
```

## Argument Reference

These arguments are supported:

* `name` - (Required) Unique account name for identification without spaces.
* `subscription` - (Required) Azure subscription ID.
* `directory_id` - (Required) Azure Active Directory tenant ID.
* `application_id` - (Required) The application ID with which the service principal is associated.
* `client_secret` - (Required) The service principal's client secret.
* `deletion_tolerance` - (Optional) The number of CME cycles to wait when the cloud provider does not return a Gateway
  until its deletion.
* `domain` - (Optional) The account's domain name in Multi-Domain Security Management Server environment.
* `environment` - (Optional) The Azure environmnet.

## Limitations

`client_secret` attribute can be set only through terraform. If the `client_secret` is set with the autoprov_cfg
command line or CME API, terraform will not recognize the change.
