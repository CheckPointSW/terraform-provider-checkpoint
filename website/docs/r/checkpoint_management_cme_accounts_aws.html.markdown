---
layout: "checkpoint"
page_title: "checkpoint_management_cme_accounts_aws"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-cme-accounts-aws"
description: |- This resource allows you to add/update/delete Check Point CME AWS Account.
---

# Resource: checkpoint_management_cme_accounts_aws

This resource allows you to add/update/delete Check Point CME AWS Account.

For details about the compatibility between the Terraform Release version and the CME API version, please refer to the section [Compatibility with CME](../index.html.markdown#compatibility-with-cme).


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

These arguments are supported:

* `name` - (Required) Unique account name for identification without spaces.
* `regions` - (Required) Comma-separated list of AWS regions in which the Gateways are being deployed.
* `credentials_file` - (Optional) One of the these options:
    1. The name of a text file containing AWS credentials that is located in $FWDIR/conf/ directory for a Management
       Server or
       $MDSDIR/conf/ directory for a Multi-Domain Management Server.
    2. “IAM” to use an IAM role profile
* `deletion_tolerance` - (Optional) The number of CME cycles to wait when the cloud provider does not return a Gateway
  until its deletion.
* `access_key` - (Optional) AWS access key.
* `secret_key` - (Optional) AWS secret key.
* `sts_role` - (Optional) AWS STS role.
* `sts_external_id` - (Optional) AWS STS external id. Must exist with STS role.
* `scan_gateways` - (Optional) Set true in order to scan gateways with AWS Transit Gateway.
* `scan_vpn` - (Optional) Set true in order to scan VPN with AWS Transit Gateway.
* `scan_load_balancers` - (Optional) Set true in order to scan load balancers access and NAT rules with AWS Transit
  Gateway.
* `scan_subnets` - (Optional) Set true in order to scan subnets with AWS Gateway Load Balancer.
* `communities` - (Optional) Comma-separated list of communities that are allowed for VPN connections for AWS Transit
  Gateways that are discovered by this account.
* `sub_accounts` - (Optional) AWS sub-accounts. Supports these parameters:
    * `name` - (Required) Sub-account name.
    * `credentials_file` - (Optional) Sub-account credentials file.
    * `access_key` - (Optional) Sub-account access key.
    * `secret_key` - (Optional) Sub-account secret key.
    * `sts_role` - (Optional) Sub-account STS role.
    * `sts_external_id` - (Optional) Sub-account STS external id. Must exist with STS role.
* `domain` - (Optional) The account's domain name in Multi-Domain Security Management Server environment.

## Limitations

`secret_key` attribute can be set only through terraform. If the `secret_key` is set with the autoprov_cfg command line
or CME API, terraform will not recognize the change.
