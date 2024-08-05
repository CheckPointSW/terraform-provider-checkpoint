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

## Examples of usage
To use Check Point provider, copy and paste this code into your Terraform configuration, update provider configuration and run `terraform init`.

### Terraform 0.12 and earlier:
```hcl
# Configure Check Point Provider for Management API
provider "checkpoint" {
  server   = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  context  = "web_api"
  session_name = "Terraform Session"
}

# Create network object
resource "checkpoint_management_network" "network" {
  name         = "My network"
  subnet4      = "192.0.2.0"
  mask_length4 = "24"
  # ...   
}
```
### Terraform 0.13 and later:
```hcl
terraform {
  required_providers {
    checkpoint = {
      source = "CheckPointSW/checkpoint"
      version = "X.Y.Z"
    }
  }
}

# Configure Check Point Provider for Management API
provider "checkpoint" {
  server   = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  context  = "web_api"
  session_name = "Terraform Session"
}

# Create network object
resource "checkpoint_management_network" "network" {
  name         = "My network"
  subnet4      = "192.0.2.0"
  mask_length4 = "24"
  # ...   
}
```

```hcl
# Configure Check Point Provider for GAIA API
provider "checkpoint" {
  server   = "192.0.2.1"
  username = "gaia_user"
  password = "gaia_password"
  context  = "gaia_api"
}

# Set machine hostname
resource "checkpoint_hostname" "hostname" {
  name = "terraform_host"
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
* `api_key` - (Optional) Check Point Management admin API key. It must be provided, but can also be defined via
  the `CHECKPOINT_API_KEY` environment variable.
* `domain` - (Optional) Login to specific domain. Domain can be identified by name or UID. This can also be defined via
  the `CHECKPOINT_DOMAIN` environment variable.
* `context` - (Optional) Check Point access context - `web_api` or `gaia_api`. This can also be defined via
  the `CHECKPOINT_CONTEXT` environment variable. Default value is `web_api`.
* `port` - (Optional) Port used for connection with the API server. This can also be defined via the `CHECKPOINT_PORT`
  environment variable. Default value is `443`.
* `proxy_host` - (Optional) Proxy host used for proxy connections. This can also be defined via
  the `CHECKPOINT_PROXY_HOST` environment variable.
* `proxy_port` - (Optional) Proxy port used for proxy connections. This can also be defined via
  the `CHECKPOINT_PROXY_PORT` environment variable.
* `session_name` - (Optional) Session unique name. This can also be defined via
  the `CHECKPOINT_SESSION_NAME` environment variable.
* `session_description` - (Optional) Session purpose description. This can also be defined via the `CHECKPOINT_SESSION_DESCRIPTION` environment variable.
* `session_file_name` - (Optional) Session file name used to store the current session id. This can also be defined via
  the `CHECKPOINT_SESSION_FILE_NAME` environment variable. default value is `sid.json`.
* `session_timeout` - (Optional) Timeout in seconds for the session established in Check Point. This can also be defined via
  the `CHECKPOINT_SESSION_TIMEOUT` environment variable. The default for the value is `600`. The timeout can be `10` - `3600`.
* `timeout` - (Optional) Timeout in seconds for the Go SDK to complete a transaction. This can also be defined via
  the `CHECKPOINT_TIMEOUT` environment variable. Default value is `120` seconds.
* `cloud_mgmt_id` - (Optional) Smart-1 Cloud management UID. This can also be defined via
  the `CHECKPOINT_CLOUD_MGMT_ID` environment variable.
* `auto_publish_batch_size` - (Optional) Number of batch size to automatically run publish. This can also be defined via the `CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE` environment variable.

## Authentication

Check Point Provider offers providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials

Usage with username and password:

```hcl
provider "checkpoint" {
  server   = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  domain   = "Domain Name"
  context  = "web_api"
}
```

Usage with API key:
```hcl
provider "checkpoint" {
  server   = "192.0.2.1"
  api_key  = "tBdloE9eOYzzSQicNxS7mA=="
  domain   = "Domain Name"
  context  = "web_api"
}
```

Usage for Smart-1 Cloud:
```hcl
provider "checkpoint" {
  server   = "chkp-vmnc6s4y.maas.checkpoint.com"
  api_key  = "tBdloE9eOYzzSQicNxS7mA=="
  cloud_mgmt_id = "de9a9b08-c7c7-436e-a64a-a54136301701"
  context  = "web_api"
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
$ export CHECKPOINT_SESSION_NAME="Terraform session name"
$ export CHECKPOINT_SESSION_DESCRIPTION="Terraform session description"
$ export CHECKPOINT_SESSION_FILE_NAME="sid.json"
$ export CHECKPOINT_SESSION_TIMEOUT=600
$ export CHECKPOINT_PROXY_HOST="1.2.3.4"
$ export CHECKPOINT_PROXY_PORT="123"
$ export CHECKPOINT_CLOUD_MGMT_ID="de9a9b08-c7c7-436e-a64a-a54136301701"
$ export CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE=100
 ```

Usage with api key:

```bash
$ export CHECKPOINT_SERVER="192.0.2.1"
$ export CHECKPOINT_API_KEY="tBdloE9eOYzzSQicNxS7mA=="
$ export CHECKPOINT_CONTEXT="web_api"
$ export CHECKPOINT_DOMAIN="Domain Name"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
$ export CHECKPOINT_SESSION_NAME="Terraform session name"
$ export CHECKPOINT_SESSION_DESCRIPTION="Terraform session description"
$ export CHECKPOINT_SESSION_FILE_NAME="sid.json"
$ export CHECKPOINT_SESSION_TIMEOUT=600
$ export CHECKPOINT_PROXY_HOST="1.2.3.4"
$ export CHECKPOINT_PROXY_PORT="123"
$ export CHECKPOINT_CLOUD_MGMT_ID="de9a9b08-c7c7-436e-a64a-a54136301701"
$ export CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE=100
 ```

Then configure the Check Point Provider as following:

```hcl
# Configure Check Point Provider via environment variables
provider "checkpoint" {}

# Create network object
resource "checkpoint_management_network" "network" {
  name         = "My network"
  subnet4      = "192.0.2.0"
  mask_length4 = "24"
  # ...
}
```

Or for GAIA API:

```bash
$ export CHECKPOINT_SERVER="192.0.2.1"
$ export CHECKPOINT_USERNAME="gaia_user"
$ export CHECKPOINT_PASSWORD="gaia_password"
$ export CHECKPOINT_CONTEXT="gaia_api"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
```

Then configure the Check Point Provider as following:

```hcl
# Configure Check Point Provider via environment variables
provider "checkpoint" {}

# Set machine hostname
resource "checkpoint_hostname" "hostname" {
  name = "terraform_host"
}
```

## Post Apply / Destroy scripts

As of right now, Terraform does not provide native support for publish and install-policy, so both of them and more post apply actions are handled
out-of-band.

In order to use post Apply / Destroy commands, the authentication method must be via environment variables.

### Publish

Please use the following script for Publish:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/publish
$ go build publish.go
$ mv publish $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && publish
```

Another option is to use `auto_publish_batch_size` provider argument which automatically runs publish.

### Install Policy

The following arguments are supported:

* `policy-package` - (Required) The name of the Policy Package to be installed.
* `target` - (Required) On what targets to execute this command. Targets may be identified by their name or object unique identifier. Multiple targets can be added.

Please use the following script for Install Policy:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/install_policy
$ go build install_policy.go
$ mv install_policy $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && install_policy -policy-package <package name> -target <target name or uid>
```

### Logout

Please use the following script for Logout:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/logout
$ go build logout.go
$ mv logout $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/logout_from_session
$ terraform apply && publish && logout_from_session
```

#### Example of usage

Run terraform then Publish & Install Policy & Logout from session

```bash
$ terraform apply && publish && install_policy -policy-package "standard" -target "corporate-gateway" && logout_from_session
```

### Discard

Please use the following script for Discard:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/discard
$ go build discard.go
$ mv discard $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ discard
```

### Approve Session

Please use the following script for Approve Session:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/approve_session
$ go build approve_session.go
$ mv approve_session $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ approve_session "SESSION_UID"
```

### Reject Session

Please use the following script for Reject Session:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/reject_session
$ go build reject_session.go
$ mv reject_session $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ reject_session "SESSION_UID" "REJECT_REASON"
```

### Submit Session

Please use the following script for Submit Session:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/submit_session
$ go build submit_session.go
$ mv submit_session $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ submit_session "SESSION_UID"
```

if no `session_uid` is provided it will submit the current session.

### Verify Policy

The following arguments are supported:

* `policy-package` - (Required) Policy package identified by the name or UID to be verified.

Please use the following script for Verify Policy:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/verify_policy
$ go build verify_policy.go
$ mv verify_policy $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && verify_policy -policy-package <package name>
```

## Import Resources

In order to import resource, use the `terraform import` command with object unique identifier.

Example:

For existing Host object with UID `9423d36f-2d66-4754-b9e2-e7f4493756d4`

Use the following resource configuration block:

```hcl
resource "checkpoint_management_host" "host" {
  name         = "myhost"
  ipv4_address = "1.1.1.1"
}
```

Run the following command:

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
* Use post apply scripts (e.g. publish, install policy, logout) to run actions after apply your changes. Terraform runs in parallel and because of that we can't predict the order of when changes will execute, running post apply scripts will ensure to run last after all changes submitted successfully.
* Create implicit / explicit dependencies between resources or modules. Terraform uses this dependency information to determine the correct order in which to create the different resources. To do so, it creates a dependency graph of all of the resources defined by the configuration. For more information, please refer [here](https://developer.hashicorp.com/terraform/tutorials/configuration-language/dependencies#dependencies).
