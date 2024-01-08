---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_aws"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-accounts-aws"
description: |- This resource allows you to add/update/delete Check Point CME AWS Account.
---

# checkpoint_management_cme_accounts_aws

This resource allows you to add/update/delete Check Point CME AWS Account.

## Example Usage

```hcl
resource "checkpoint_management_cme_accounts_aws" "aws_account" {
  name                = "awsAccount"
  regions             = ["eu-north-1"]
  credentials_file    = "IAM"
  scan_load_balancers = true
  sts_role            = "arn:aws:iam::123412341234:role/EXAMPLE"
  sts_external_id     = "12345"
  sub_accounts {
    name             = "sub_account_a"
    credentials_file = "IAM"
  }
  sub_accounts {
    name       = "sub_account_b"
    access_key = "AKIAIOSFODNN7EXAMPLE"
    secret_key = "wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Unique account name for identification without spaces.
* `regions` - (Required) Comma-separated list of AWS regions, in which the gateways are being deployed.
* `credentials_file` - (Optional) One of the following options:
    1. The name of a text file containing AWS credentials located in $FWDIR/conf/ directory for a Management Server or
       $MDSDIR/conf/ directory for a Multi-Domain Management Server.
    2. “IAM” to use an IAM role profile
* `deletion_tolerance` - (Optional) The number of CME cycles to wait when the cloud provider does not return a GW until
  its deletion.
* `access_key` - (Optional) AWS access key.
* `secret_key` - (Optional) AWS secret key.
* `sts_role` - (Optional) AWS sts role.
* `sts_external_id` - (Optional) AWS sts external id, must exist with sts role.
* `scan_gateways` - (Optional) Set true in order to scan gateways with AWS TGW.
* `scan_vpn` - (Optional) Set true in order to scan vpn with AWS TGW.
* `scan_load_balancers` - (Optional) Set true in order to scan load balancers access and NAT rules with AWS TGW.
* `scan_subnets` - (Optional) Set true in order to scan subnets with AWS GWLB.
* `communities` - (Optional) Comma-separated list of communities, which are allowed for VPN connections fow AWS TGW that
  are discovered by this account.
* `sub_accounts` - (Optional) AWS sub accounts. supports the following:
    * `name` - (Required) Sub account name.
    * `credentials_file` - (Optional) Sub account credentials file.
    * `access_key` - (Optional) Sub account access key.
    * `secret_key` - (Optional) Sub account secret key.
    * `sts_role` - (Optional) Sub account sts role.
    * `sts_external_id` - (Optional) Sub account sts external id, must exist with sts role.
* `domain` - (Optional) The account's domain name in MDS environment.

## Limitations

`secret_key` attribute can be managed only through the created resources in terraform.
