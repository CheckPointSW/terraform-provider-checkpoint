---
layout: "checkpoint"
page_title: "Provider: Check Point"
sidebar_current: "docs-checkpoint-index"
description: |- The Check Point provider can be used to automate security responses to threats, provision both physical
and virtualized next-generation firewalls and automate routine Security Management configuration tasks, saving time and
reducing configuration errors. With the Check Point provider, DevOps teams can automate their security and transform it
into DevSecOps workflows.
---

# Check Point Provider

The Check Point provider can be used to automate security responses to threats, provision both physical and virtualized
next-generation firewalls and automate routine Security Management configuration tasks, saving time and reducing
configuration errors. With the Check Point provider, DevOps teams can automate their security and transform it into
DevSecOps workflows.

## Examples usage
## Terraform 0.12 and earlier:

```hcl
# Configure the Check Point Provider
provider "checkpoint" {
  server   = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  context  = "web_api"
}

# Create network object
resource "checkpoint_management_network" "network" {
  name         = "network"
  subnet4      = "192.0.2.0"
  mask_length4 = "24"
  # ...   
}
```
## Terraform 0.13 and later:
```hcl
terraform {
  required_providers {
    checkpoint = {
      source  = "checkpointsw/checkpoint"
      version = "~> 1.6.0"
    }
  }
}

# Configure the Check Point Provider
provider "checkpoint" {
  server   = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  context  = "web_api"
}
```

```hcl
# Configure the Check Point Provider for GAIA API
provider "checkpoint" {
  server   = "192.0.2.1"
  username = "gaia_user"
  password = "gaia_password"
  context  = "gaia_api"
}

# Set machine hostname
resource "checkpoint_hostname" "hostname" {
  name = "terrahost"
}
```

## Argument Reference

The following arguments are supported:

* `server` - (Optional) Check Point Management server IP. It must be provided, but can also be defined via
  the `CHECKPOINT_SERVER` environment variable.
* `username` - (Optional) Check Point Management admin name. It must be provided, but can also be defined via
  the `CHECKPOINT_USERNAME` environment variable.
* `password` - (Optional) Check Point Management admin password. It must be provided, but can also be defined via
  the `CHECKPOINT_PASSWORD` environment variable.
* `api_key` - (Optional) Check Point Management admin api key. this can also be defined via
  the `CHECKPOINT_API_KEY` environment variable.
* `context` - (Optional) Check Point access context - `web_api` or `gaia_api`. This can also be defined via
  the `CHECKPOINT_CONTEXT` environment variable. Default value is `web_api`.
* `domain` - (Optional) Login to specific domain. Domain can be identified by name or UID. This can also be defined via
  the `CHECKPOINT_DOMAIN` environment variable.
* `timeout` - (Optional) Timeout in seconds for the Go SDK to complete a transaction. This can also be defined via
  the `CHECKPOINT_TIMEOUT` environment variable. Default value is `10` seconds.
* `port` - (Optional) Port used for connection to the API server. This can also be defined via the `CHECKPOINT_PORT`
  environment variable. Default value is `443`.
* `proxy_host` - (Optional) Proxy host used for proxy connections. this can also be defined via
  the `CHECKPOINT_PROXY_HOST` environment variable.
* `proxy_port` - (Optional) Proxy port used for proxy connections. this can also be defined via
  the `CHECKPOINT_PROXY_PORT` environment variable.
* `session_name` - (Optional) Session unique name. this can also be defined via
  the `CHECKPOINT_SESSION_NAME` environment variable.
* `session_description` - (Optional) A description of the session's purpose. this can also be defined via the `CHECKPOINT_SESSION_DESCRIPTION` environment variable.
* `session_file_name` - (Optional) Session file name used to store the current session id. this can also be defined via
  the `CHECKPOINT_SESSION_FILE_NAME` environment variable. default value is `sid.json`.
* `session_timeout` - (Optional) The timeout in seconds for the session established in Check Point. This can also be defined via
  the `CHECKPOINT_SESSION_TIMEOUT` environment variable. The default for the value is `600`. The timeout can be `10` - `3600`.
* `cloud_mgmt_id` - (Optional) Smart-1 Cloud management UID. this can also be defined via
  the `CHECKPOINT_CLOUD_MGMT_ID` environment variable.

## Authentication

The Check Point provider offers providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials

Usage with username and password:

```hcl
provider "checkpoint" {
  server   = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  context  = "web_api"
  domain   = "Domain Name"
}
```
Usage with api key:
```hcl
provider "checkpoint" {
  server   = "192.0.2.1"
  api_key  = "tBdloE9eOYzzSQicNxS7mA=="
  context  = "web_api"
  domain   = "Domain Name"
}
```

Smart-1 Cloud:
```hcl
provider "checkpoint" {
  server   = "chkp-vmnc6s4y.maas.checkpoint.com"
  api_key  = "tBdloE9eOYzzSQicNxS7mA=="
  context  = "web_api"
  cloud_mgmt_id = "de9a9b08-c7c7-436e-a64a-a54136301701"
}
```

Or for GAIA API:

```hcl
provider "checkpoint" {
  server   = "192.0.2.1"
  username = "gaia_user"
  password = "gaia_password"
  context  = "gaia_api"
}
```

### Environment variables

You can provide your credentials via environment variables. Note that setting your Check Point credentials using static
credentials will override the environment variables.

Usage:

```bash
$ export CHECKPOINT_SERVER="192.0.2.1"
$ export CHECKPOINT_USERNAME="aa"
$ export CHECKPOINT_PASSWORD="aaaa"
$ export CHECKPOINT_CONTEXT="web_api"
$ export CHECKPOINT_DOMAIN="Domain Name"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
$ export CHECKPOINT_SESSION_NAME="Terraform session"
$ export CHECKPOINT_SESSION_FILE_NAME="sid.json"
$ export CHECKPOINT_SESSION_TIMEOUT=600
$ export CHECKPOINT_PROXY_HOST="1.2.3.4"
$ export CHECKPOINT_PROXY_PORT="123"
$ export CHECKPOINT_CLOUD_MGMT_ID="de9a9b08-c7c7-436e-a64a-a54136301701"
$ export CHECKPOINT_SESSION_DESCRIPTION="session description"
 ```

Usage with api key:

```bash
$ export CHECKPOINT_SERVER="192.0.2.1"
$ export CHECKPOINT_API_KEY="tBdloE9eOYzzSQicNxS7mA=="
$ export CHECKPOINT_CONTEXT="web_api"
$ export CHECKPOINT_DOMAIN="Domain Name"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
$ export CHECKPOINT_SESSION_NAME="Terraform session"
$ export CHECKPOINT_SESSION_FILE_NAME="sid.json"
$ export CHECKPOINT_SESSION_TIMEOUT=600
$ export CHECKPOINT_PROXY_HOST="1.2.3.4"
$ export CHECKPOINT_PROXY_PORT="123"
$ export CHECKPOINT_CLOUD_MGMT_ID="de9a9b08-c7c7-436e-a64a-a54136301701"
$ export CHECKPOINT_SESSION_DESCRIPTION="session description"
 ```

Then configure the Check Point Provider as following:

```hcl
# Configure the Check Point Provider
provider "checkpoint" {}

# Create network object
resource "checkpoint_management_network" "network" {
  name         = "network"
  subnet4      = "192.0.2.0"
  mask_length4 = "24"
  # ...   
}
```

Or for GAIA API:

```bash
$ export CHECKPOINT_SERVER=192.0.2.1
$ export CHECKPOINT_USERNAME="gaia_user"
$ export CHECKPOINT_PASSWORD="gaia_password"
$ export CHECKPOINT_CONTEXT="gaia_api"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
```

Then configure the Check Point Provider as following:

```hcl
# Configure the Check Point Provider
provider "checkpoint" {}

# Set machine hostname
resource "checkpoint_hostname" "hostname" {
  name = "terrahost"
}
```

## Post Apply/Destroy commands

As of right now, Terraform does not provide native support for publish and install-policy, so both of them are handled
out-of-band.

In order to use post Apply/Destroy commands, the authentication method must be via environment variables.

### Publish

Please use the following for publish:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/publish
$ go build publish.go
$ mv publish $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && publish
```

### Logout

Please use the following for logout:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/logout
$ go build logout.go
$ mv logout $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/logout_from_session
$ terraform apply && publish && logout_from_session
```

### Discard

Please use the following for discard:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/discard
$ go build discard.go
$ mv discard $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ discard
```

### Approve session

Please use the following for approve session:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/approve_session
$ go build approve_session.go
$ mv approve_session $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ approve_session "SESSION_UID"
```

### Reject session

Please use the following for reject session:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/reject_session
$ go build reject_session.go
$ mv reject_session $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ reject_session "SESSION_UID" "REJECT_REASON"
```

### Submit session

Please use the following for submit session:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/submit_session
$ go build submit_session.go
$ mv submit_session $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ submit_session "SESSION_UID"
```

if no `session_uid` is provided it will submit the current session.

### Verify Policy

The following arguments are supported:

* `policy-package` - (Required) The name of the Policy Package to be Verified.

Please use the following for Verify policy:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/verify_policy
$ go build verify_policy.go
$ mv verify_policy $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && verify_policy -policy-package <package name>
```

### Install Policy

The following arguments are supported:

* `policy-package` - (Required) The name of the Policy Package to be installed.
* `target` - (Required) On what targets to execute this command. Targets may be identified by their name, or object
  unique identifier. Multiple targets can be added.

Please use the following for install policy:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/install_policy
$ go build install_policy.go
$ mv install_policy $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && install_policy -policy-package <package name> -target <target name or uid>
```

### Example usage

```bash
$ terraform apply && publish && install_policy -policy-package "standard" -target "corporate-gateway" && logout_from_session
```

## Import Resources

In order to import resource, use the `terraform import` command with object unique identifier.

Example:

Host object with UID `9423d36f-2d66-4754-b9e2-e7f4493756d4`

Write resource configuration block

```hcl
resource "checkpoint_management_host" "host" {
  name         = "myhost"
  ipv4_address = "1.1.1.1"
}
```

Run the following command

```bash
$ terraform import checkpoint_management_host.host 9423d36f-2d66-4754-b9e2-e7f4493756d4
```

For more information about `terraform import` command, please
refer [here](https://www.terraform.io/docs/import/usage.html).

## Tips & Best Practices

This section describes best practices for working with the Check Point provider.

* Use one or more dedicated users for provider operations to make sure minimum permissions are granted.
* Keep on object name uniqueness in your environment.
* Use object name when reference to an object (avoid use of object UID).
* Use post apply scripts (e.g. publish, install policy) to run actions after apply your changes. Terraform runs in parallel and because of that we can't predict the order of when changes will execute so running post apply scripts will ensure to run last after all changes submitted successfully.
