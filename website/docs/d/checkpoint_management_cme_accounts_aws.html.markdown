---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_aws"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cme-accounts-aws"
description: |- Use this data source to get information on an existing Check Point CME AWS Account.
---

# Data Source: checkpoint_management_cme_accounts_aws

Use this data source to get information on an existing Check Point CME AWS Account.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](https://registry.terraform.io/providers/CheckPointSW/checkpoint/latest/docs#compatibility-with-cme).


## Example Usage

```hcl
data "checkpoint_management_cme_accounts_aws" "aws_account" {
  name = "awsAccount"
}
```

## Argument Reference

These arguments are supported:

* `name` - (Required) Unique account name for identification.
* `regions` - Comma-separated list of AWS regions where gateways are being deployed.
* `platform` - The platform of the account.
* `gw_configurations` - A list of Gateway configurations attached to the account.
* `credentials_file` - The credentials file.
* `deletion_tolerance` - The number of CME cycles to wait when the cloud provider does not return a Gateway until its
  deletion.
* `access_key` - AWS access key.
* `secret_key` - AWS secret key.
* `sts_role` - AWS STS role.
* `sts_external_id` - AWS STS external id. Must exist with STS role.
* `scan_gateways` - true/false for scan gateways with AWS Transit Gateway.
* `scan_vpn` - true/false for scan VPN with AWS Transit Gateway.
* `scan_load_balancers` - true/false for scan load balancers access and NAT rules with AWS Transit Gateway.
* `scan_subnets` - true/false for scan subnets with AWS Gateway Load Balancer.
* `communities` - Comma-separated list of communities which are allowed for VPN connections of AWS Transit Gateway that
  are discovered by this account.
* `sub_accounts` - AWS sub-accounts. Supports the following:
    * `name` - Sub-account name.
    * `credentials_file` - Sub-account credentials file.
    * `access_key` - Sub-account access key.
    * `secret_key` - Sub-account secret key.
    * `sts_role` - Sub-account STS role.
    * `sts_external_id` - Sub-account STS external id. Must exist with STS role.
* `domain` - The account's domain name in Multi-Domain Security Management Server environment.
