---
layout: "checkpoint"
page_title: "Provider: Check Point"
sidebar_current: "docs-checkpoint-index"
description: |-
  Check point provider is used to interact with a resources supported by Management and Gaia API’s. The provider allows managing firewall, and threat capabilities. 
---

# Check Point Provider

Check point provider is used to interact with a resources supported by Management and Gaia API’s.
  The provider allows managing firewall, and threat capabilities.

##Example usage
```hcl
# Configure the Check Point Provider
provider "checkpoint" {
	server = "192.168.52.178"
	username = "bob"
	password = "Bob123"
	context = "web_api"
}

# Create a Network Object
resource "checkpoint_management_network" "test1" {
	name = "network1"
	subnet4 = "192.0.2.0"
	# ...
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
	server = "192.168.52.178"
	username = "bob"
	password = "Bob123"
	context = "web_api"
}
```

### Environment variables
You can provide your credentials via environment variables. Note that setting your Check Point credentials using static credentials will override the environment variables.

Usage:

```hcl
$ export CHECKPOINT_SERVER=192.168.52.178
$ export CHECKPOINT_USERNAME="bob"
$ export CHECKPOINT_PASSWORD="Bob123"
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
