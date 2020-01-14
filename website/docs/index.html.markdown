---
layout: "checkpoint"
page_title: "Provider: Check Point"
sidebar_current: "docs-checkpoint-index"
description: |-
  The Check Point provider can be used to automate security responses to threats, provision both physical and virtualized next-generation firewalls and automate routine Security Management configuration tasks, saving time and reducing configuration errors. With the Check Point provider, DevOps teams can automate their security and transform it into DevSecOps workflows.
---

# Check Point Provider

The Check Point provider can be used to automate security responses to threats, provision both physical and virtualized next-generation firewalls and automate routine Security Management configuration tasks, saving time and reducing configuration errors. With the Check Point provider, DevOps teams can automate their security and transform it into DevSecOps workflows.

##Example usage
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

```hcl
$ export CHECKPOINT_SERVER=192.0.2.1
$ export CHECKPOINT_USERNAME="aa"
$ export CHECKPOINT_PASSWORD="aaaa"
$ export CHECKPOINT_CONTEXT="web_api"
 ```
Or for GAIA API:
```hcl
$ export CHECKPOINT_SERVER=192.0.2.1
$ export CHECKPOINT_USERNAME="gaia_user"
$ export CHECKPOINT_PASSWORD="gaia_password"
$ export CHECKPOINT_CONTEXT="gaia_api"
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
