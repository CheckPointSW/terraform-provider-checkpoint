---
layout: "checkpoint"
page_title: "checkpoint_management_aws_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-aws-data-center-server"
description: |- This resource allows you to execute Check Point aws data center server.
---

# Resource: checkpoint_management_aws_data_center_server

This resource allows you to execute Check Point AWS Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_aws_data_center_server" "testAws" {
  authenticationMethod = "user-authentication"
  accessKeyId          = "MY-KEY-ID"
  secretAccessKey      = "MY-SECRET-KEY"
  region               = "us-east-1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name.
* `authentication_method` - (Required) user-authentication Uses the Access keys to authenticate. role-authentication Uses the AWS IAM role to authenticate. This option requires the Security Management Server be deployed in AWS and has an IAM Role.
* `access_key_id` - (Required for authentication-method: user-authentication) Access key ID for the AWS account. Required for authentication-method:user-authentication.
* `secret_access_key` - (Required for authentication-method: user-authentication) Secret access key for the AWS account. Required for authentication-method:user-authentication.
* `region` - (Optional)  Select the AWS region.
* `enable_sts_assume_role` - (Optional) Enables the STS Assume Role option. After it is enabled, the sts-role field is mandatory, whereas the sts-external-id is optional.
* `sts_role` - (Required for enable-sts-assume-role: true) Enables the STS Assume Role option. After it is enabled, the sts-role field is mandatory, whereas the sts-external-id is optional.
* `sts_external_id` - (Optional) An optional STS External-Id to use when assuming the role.
* `tags` - (Optional) Collection of tag identifiers. tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
