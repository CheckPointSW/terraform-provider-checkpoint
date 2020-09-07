---
layout: "checkpoint"
page_title: "Provider: Check Point"
sidebar_current: "docs-checkpoint-index"
description: |-
  The Check Point provider can be used to automate security responses to threats, provision both physical and virtualized next-generation firewalls and automate routine Security Management configuration tasks, saving time and reducing configuration errors. With the Check Point provider, DevOps teams can automate their security and transform it into DevSecOps workflows.
---

# Check Point Provider

The Check Point provider can be used to automate security responses to threats, provision both physical and virtualized next-generation firewalls and automate routine Security Management configuration tasks, saving time and reducing configuration errors. With the Check Point provider, DevOps teams can automate their security and transform it into DevSecOps workflows.

## Examples usage
```hcl
# Configure the Check Point Provider
provider "checkpoint" {
    server = "192.0.2.1"
    username = "aa"
    password = "aaaa"
    context = "web_api"
}

# Create network object
resource "checkpoint_management_network" "network" {
    name = "network"
    subnet4 = "192.0.2.0"	
    mask_length4 = "24"
    # ...   
}
```
```hcl
# Configure the Check Point Provider for GAIA API
provider "checkpoint" {
    server = "192.0.2.1"
    username = "gaia_user"
    password = "gaia_password"
    context = "gaia_api"
}

# Set machine hostname
resource "checkpoint_hostname" "hostname" {
    name = "terrahost"
}
```
## Argument Reference

The following arguments are supported:

* `server` - (Optional) Check Point Management server IP. It must be provided, but can also be defined via the `CHECKPOINT_SERVER` environment variable.
* `username` - (Optional) Check Point Management admin name. It must be provided, but can also be defined via the `CHECKPOINT_USERNAME` environment variable.
* `password` - (Optional) Check Point Management admin password. It must be provided, but can also be defined via the `CHECKPOINT_PASSWORD` environment variable.
* `context` - (Optional) Check Point access context - `web_api` or `gaia_api`. This can also be defined via the `CHECKPOINT_CONTEXT` environment variable. Default value is `web_api`.
* `domain` - (Optional) Login to specific domain. Domain can be identified by name or UID. This can also be defined via the `CHECKPOINT_DOMAIN` environment variable.
* `timeout` - (Optional) Timeout in seconds for the Go SDK to complete a transaction. This can also be defined via the `CHECKPOINT_TIMEOUT` environment variable. Default value is `10` seconds.
* `port` - (Optional) Port used for connection to the API server. This can also be defined via the `CHECKPOINT_PORT` environment variable. Default value is `443`.

## Authentication

The Check Point provider offers providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials

Usage:

```hcl
provider "checkpoint" {
    server = "192.0.2.1"
    username = "aa"
    password = "aaaa"
    context = "web_api"
    domain = "Domain Name"
}
```

Or for GAIA API:

```hcl
provider "checkpoint" {
    server = "192.0.2.1"
    username = "gaia_user"
    password = "gaia_password"
    context = "gaia_api"
}
```

### Environment variables
You can provide your credentials via environment variables. Note that setting your Check Point credentials using static credentials will override the environment variables.

Usage:

```bash
$ export CHECKPOINT_SERVER=192.0.2.1
$ export CHECKPOINT_USERNAME="aa"
$ export CHECKPOINT_PASSWORD="aaaa"
$ export CHECKPOINT_CONTEXT="web_api"
$ export CHECKPOINT_DOMAIN="Domain Name"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
 ```
 
Then configure the Check Point Provider as following:

```hcl
# Configure the Check Point Provider
provider "checkpoint" { }

# Create network object
resource "checkpoint_management_network" "network" {
    name = "network"
    subnet4 = "192.0.2.0"	
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
provider "checkpoint" { }

# Set machine hostname
resource "checkpoint_hostname" "hostname" {
    name = "terrahost"
}
```

## Post Apply/Destroy commands

As of right now, Terraform does not provide native support for publish and install-policy, so both of them are handled out-of-band. 

In order to use post Apply/Destroy commands, the authentication method must be via environment variables.

### Publish

Please use the following for publish:
 
```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/publish
$ go build publish.go
$ mv publish $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && publish
```

### Install Policy

The following arguments are supported:

* `policy-package` - (Required) The name of the Policy Package to be installed.
* `target` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier. Multiple targets can be added.

Please use the following for install policy:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/install_policy
$ go build install_policy.go
$ mv install_policy $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && install_policy -policy-package <package name> -target <target name or uid>
```

### Example usage

```bash
$ terraform apply && publish && install_policy -policy-package "standard" -target "corporate-gateway"
```

## Import Resources

In order to import resource, use the `terraform import` command with object unique identifier.

Example:

Host object with UID `9423d36f-2d66-4754-b9e2-e7f4493756d4`

Write resource configuration block

```hcl
resource "checkpoint_management_host" "host" {
    name = "myhost"
    ipv4_address = "1.1.1.1"
}
```

Run the following command

```bash
$ terraform import checkpoint_management_host.host 9423d36f-2d66-4754-b9e2-e7f4493756d4
```

For more information about `terraform import` command, please refer [here](https://www.terraform.io/docs/import/usage.html).