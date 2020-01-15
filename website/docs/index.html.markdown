---
layout: "checkpoint"
page_title: "Provider: Check Point"
sidebar_current: "docs-checkpoint-index"
description: |-
  The Check Point provider can be used to automate security responses to threats, provision both physical and virtualized next-generation firewalls and automate routine Security Management configuration tasks, saving time and reducing configuration errors. With the Check Point provider, DevOps teams can automate their security and transform it into DevSecOps workflows.
---

# Check Point Provider

The Check Point provider can be used to automate security responses to threats, provision both physical and virtualized next-generation firewalls and automate routine Security Management configuration tasks, saving time and reducing configuration errors. With the Check Point provider, DevOps teams can automate their security and transform it into DevSecOps workflows.

##Examples usage
```hcl
# Configure the Check Point Provider
provider "checkpoint" {
	server = "192.0.2.1"
	username = "aa"
	password = "aaaa"
	context = "web_api"
}

# Create a Network Object
resource "checkpoint_management_network" "test1" {
	name = "network1"
	subnet4 = "192.0.2.0"
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
 ```
 
Then configure the Check Point Provider as following:

```hcl
# Configure the Check Point Provider
provider "checkpoint" { }

# Create a Network Object
resource "checkpoint_management_network" "test1" {
	name = "network1"
	subnet4 = "192.0.2.0"
	# ...
}
```

Or for GAIA API:

```bash
$ export CHECKPOINT_SERVER=192.0.2.1
$ export CHECKPOINT_USERNAME="gaia_user"
$ export CHECKPOINT_PASSWORD="gaia_password"
$ export CHECKPOINT_CONTEXT="gaia_api"
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

### Publish

Please use the following for publish:
 
```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/publish
$ go build publish.go
$ mv publish $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && publish
```

### Install-Policy

The following arguments are supported:

* `policy-package` - (Required) The name of the Policy Package to be installed.
* `target` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier. Multiple targets can be added.

Please use the following for install-policy:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/install_policy
$ go build install_policy.go
$ mv install_policy $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && install_policy -policy-package <package name> -target <target name or uid>
```

### Example usage

```bash
$ terraform apply && publish && install_policy -policy-package "Standard" -target "Firewall-harry-main-take-265"
```