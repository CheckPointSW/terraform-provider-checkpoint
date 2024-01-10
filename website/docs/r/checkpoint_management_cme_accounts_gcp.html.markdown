---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_gcp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-accounts-gcp"
description: |- This resource allows you to add/update/delete Check Point CME GCP Account.
---

# Resource: checkpoint_management_cme_accounts_gcp

This resource allows you to add/update/delete Check Point CME GCP Account.

Available in:

- Check Point Security Management/Multi Domain Management Server R81.10 and higher.
- CME take 255 and higher.

## Example Usage

```hcl
resource "checkpoint_management_cme_accounts_gcp" "gcp_account" {
  name             = "gcpAccount"
  project_id       = "my-project-1"
  credentials_file = "cred_file.json"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique account name for identification without spaces.
* `project_id` - (Required) GCP project id.
* `credentials_file` - (Optional) The name of a text file containing GCP credentials located in $FWDIR/conf directory
  for a Management Server or $MDSDIR/conf directory for a Multi-Domain Management Server.
* `credentials_data` - (Optional) Base64 encoded string that represents the content of the credentials file.
* `deletion_tolerance` - (Optional) The number of CME cycles to wait when the cloud provider does not return a GW until
  its deletion.
* `domain` - (Optional) The account's domain name in MDS environment.
