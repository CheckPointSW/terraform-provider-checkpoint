---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_aws"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-accounts-aws"
description: |- Use this data source to get information on an existing Check Point CME AWS Account.
---

# checkpoint_management_cme_accounts_aws

Use this data source to get information on an existing Check Point CME AWS Account.

## Example Usage

```hcl
data "checkpoint_management_cme_accounts_aws" "aws_account" {
  name = "aws-controller"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique account name for identification.
* `regions` - Comma-separated list of AWS regions, in which the gateways are being deployed.
* `platform` - The platform of the account.
* `gw_configurations` - A list of GW configurations attached to the account.
* `credentials_file` - The credentials file.
* `deletion_tolerance` - The number of CME cycles to wait when the cloud provider does not return a GW until its deletion.
* `access_key` - AWS access key.
* `secret_key` - AWS secret key.
* `sts_role` - AWS sts role.
* `sts_external_id` - AWS sts external id, must exist with sts role.
* `scan_gateways` - true/false for scan gateways with AWS TGW.
* `scan_vpn` - true/false for scan vpn with AWS TGW.
* `scan_load_balancers` - true/false for scan load balancers access and NAT rules with AWS TGW.
* `scan_subnets` - true/false for scan subnets with AWS GWLB.
* `communities` - Comma-separated list of communities, which are allowed for VPN connections fow AWS TGW that are discovered by this account.
* `sub_accounts` - AWS sub accounts. supports the following:
  * `name` - Sub account name.
  * `credentials_file` - Sub account credentials file.
  * `access_key` - Sub account access key.
  * `secret_key` - Sub account secret key.
  * `sts_role` - Sub account sts role.
  * `sts_external_id` - Sub account sts external id, must exist with sts role.
* `domain` - The account's domain name in MDS environment.
