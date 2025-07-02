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
  server = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  context = "web_api"
  session_name = "Terraform session"
}

# Create network object
resource "checkpoint_management_network" "network" {
  name = "My network"
  subnet4 = "192.0.2.0"
  mask_length4 = "24"
  # ...   
}
```

```hcl
# Configure Check Point Provider for Management API for specific domain
provider "checkpoint" {
  server = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  context = "web_api"
  domain = "MyDomain"
  session_file_name = "mydomain.json"
  session_name = "Terraform session"
}

# Create network object
resource "checkpoint_management_network" "network" {
  name = "My network"
  subnet4 = "192.0.2.0"
  mask_length4 = "24"
  # ...   
}
```

```hcl
# Configure Check Point Provider for GAIA API
provider "checkpoint" {
  server = "192.0.2.1"
  username = "gaia_user"
  password = "gaia_password"
  context = "gaia_api"
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
* `auto_publish_batch_size` - (Optional) Number of batch size to automatically run publish. This can also be defined via
  the `CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE` environment variable.
* `ignore_server_certificate` - (Optional) Indicates that the client should not check the server's certificate. This can also be defined via
  the `CHECKPOINT_IGNORE_SERVER_CERTIFICATE` environment variable.

## Authentication

Check Point Provider offers providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables

### Static credentials

Usage with username and password:

```hcl
provider "checkpoint" {
  server = "192.0.2.1"
  username = "aa"
  password = "aaaa"
  domain = "MyDomain"
  session_file_name = "mydomain.json"
  context = "web_api"
}
```

Usage with API key:
```hcl
provider "checkpoint" {
  server = "192.0.2.1"
  api_key = "tBdloE9eOYzzSQicNxS7mA=="
  domain = "MyDomain"
  session_file_name = "mydomain.json"
  context = "web_api"
}
```

Usage for Smart-1 Cloud:
```hcl
provider "checkpoint" {
  server = "chkp-vmnc6s4y.maas.checkpoint.com"
  api_key = "tBdloE9eOYzzSQicNxS7mA=="
  cloud_mgmt_id = "de9a9b08-c7c7-436e-a64a-a54136301701"
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

You can provide your credentials via environment variables. Note that setting your Check Point credentials using static
credentials will override the environment variables.

Usage:

```bash
$ export CHECKPOINT_SERVER="192.0.2.1"
$ export CHECKPOINT_USERNAME="aa"
$ export CHECKPOINT_PASSWORD="aaaa"
$ export CHECKPOINT_CONTEXT="web_api"
$ export CHECKPOINT_DOMAIN="MyDomain"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
$ export CHECKPOINT_SESSION_NAME="Terraform session name"
$ export CHECKPOINT_SESSION_DESCRIPTION="Terraform session description"
$ export CHECKPOINT_SESSION_FILE_NAME="mydomain.json"
$ export CHECKPOINT_SESSION_TIMEOUT=600
$ export CHECKPOINT_PROXY_HOST="1.2.3.4"
$ export CHECKPOINT_PROXY_PORT="123"
$ export CHECKPOINT_CLOUD_MGMT_ID="de9a9b08-c7c7-436e-a64a-a54136301701"
$ export CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE=100
$ export CHECKPOINT_IGNORE_SERVER_CERTIFICATE=false
 ```

Usage with api key:

```bash
$ export CHECKPOINT_SERVER="192.0.2.1"
$ export CHECKPOINT_API_KEY="tBdloE9eOYzzSQicNxS7mA=="
$ export CHECKPOINT_CONTEXT="web_api"
$ export CHECKPOINT_DOMAIN="MyDomain"
$ export CHECKPOINT_TIMEOUT=10
$ export CHECKPOINT_PORT=443
$ export CHECKPOINT_SESSION_NAME="Terraform session name"
$ export CHECKPOINT_SESSION_DESCRIPTION="Terraform session description"
$ export CHECKPOINT_SESSION_FILE_NAME="mydomain.json"
$ export CHECKPOINT_SESSION_TIMEOUT=600
$ export CHECKPOINT_PROXY_HOST="1.2.3.4"
$ export CHECKPOINT_PROXY_PORT="123"
$ export CHECKPOINT_CLOUD_MGMT_ID="de9a9b08-c7c7-436e-a64a-a54136301701"
$ export CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE=100
$ export CHECKPOINT_IGNORE_SERVER_CERTIFICATE=false
 ```

Then configure the Check Point Provider as following:

```hcl
# Configure Check Point Provider via environment variables
provider "checkpoint" {}

# Create network object
resource "checkpoint_management_network" "network" {
  name = "My network"
  subnet4 = "192.0.2.0"
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

There are actions that can run out-of-band Terraform using dedicated scripts for publish, install-policy and more.

In order to use post apply or post destroy commands, the authentication method must be via environment variables.

### Publish

Please use the following script for Publish:

```bash
$ cd $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint/commands/publish
$ go build publish.go
$ mv publish $GOPATH/src/github.com/terraform-providers/terraform-provider-checkpoint
$ terraform apply && publish
```

### Install Policy

The following arguments are supported:

* `policy-package` - (Required) The name of the Policy Package to be installed.
* `target` - (Optional) On what targets to execute this command. Targets may be identified by their name or object unique identifier. Multiple targets can be added.
* `access` - (Optional) Set to be true in order to install the Access Control policy. By default, the value is true if Access Control policy is enabled on the input policy package, otherwise false.
* `desktop-security` - (Optional) Set to be true in order to install the Desktop Security policy. By default, the value is true if desktop security policy is enabled on the input policy package, otherwise false.
* `qos` - (Optional) Set to be true in order to install the QoS policy. By default, the value is true if Quality-of-Service policy is enabled on the input policy package, otherwise false.
* `threat-prevention` - (Optional) Set to be true in order to install the Threat Prevention policy. By default, the value is true if Threat Prevention policy is enabled on the input policy package, otherwise false.
* `install-on-all-cluster-members-or-fail` - (Optional) Relevant for the gateway clusters. If true, the policy is installed on all the cluster members. If the installation on a cluster member fails, don't install on that cluster.
* `prepare-only` - (Optional) If true, prepares the policy for the installation, but doesn't install it on an installation target.
* `revision` - (Optional) The UID of the revision of the policy to install.
* `ignore-warnings` - (Optional) Install policy ignoring policy mismatch warnings.

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

## Compatibility with Management
Check Point Provider supports Management server from version R80 and above.
However, some Terraform resources or specific fields in Terraform resource might not be available because they are not supported in your Management API version.
<br>You can check the Management API [versions list](https://sc1.checkpoint.com/documents/latest/APIs/index.html#api_versions) to see what is supported by your Management server.

## Compatibility with CME
Check Point Provider supports configuring objects in CME configuration file starting from Security Management/Multi-Domain Security Management Server version R81.10 and higher.

The table below shows the compatibility between the Terraform Release version and the CME API version:

| Terraform Release version | CME API version | CME Take       |
|---------------------------|-----------------|----------------|
| v2.11.0                   | v1.3.1          | 309 and higher |
| v2.9.0                    | v1.2.2          | 289 and higher |
| v2.8.0                    | v1.2            | 279 and higher |
| v2.7.0                    | v1.1            | 255 and higher |

<br>
-> **Note:** When you install or upgrade the Terraform Release version, make sure to also upgrade CME to the corresponding CME Take to properly configure CME resources.

For details about upgrading CME, please refer to the documentation [here](https://sc1.checkpoint.com/documents/IaaS/WebAdminGuides/EN/CP_CME/Content/Topics-CME/Installing_and_Updating_CME.htm?tocpath=_____4).

## Import Resources

In order to import resource, use the `terraform import` command with object unique identifier.

Example:

For existing Host object with UID `9423d36f-2d66-4754-b9e2-e7f4493756d4`

Use the following resource configuration block:

```hcl
resource "checkpoint_management_host" "host" {
  name = "myhost"
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
* Keep on unique `session_file_name` when configure more than one provider for authentication purposes.
* Resources and Data Sources that start with `checkpoint_management_*` using Management API and require set context to `web_api`. For GAIA API resources set context to `gaia_api`.
* When configure provider context to `gaia_api` you can run only GAIA resources. Management resources will not be supported.
* Provider state policy is to capture all resource attributes into Terraform state. All attributes defined in the resource schema are recorded and kept up-to-date in the state. For more information, please refer [here](https://developer.hashicorp.com/terraform/plugin/sdkv2/best-practices/detecting-drift#capture-all-state-in-read).

### Publish best options and practices

#### Trigger field
From version 1.2 the provider was enhanced where a `triggers` field for resource `install-policy`, `publish` and `logout` was added for re-execution if there are any changes in the configuration files.
```hcl
# Put the Check Point configuration in a sub folder and refer to is as a module
module "chkp_policy" {
  source = "./chkp_policy"
}

# Activate the trigger if there is a change in the configuration files in the folder chkp_policy
locals {
  publish_triggers = [for filename in fileset(path.module, "chkp_policy/*.tf"): filesha256(filename)]
}

# Make the publish resources dependent of the module and **trigger** it if there is a change in the configuration files
resource "checkpoint_management_publish" "publish" {
  depends_on = [ module.policy ]
  triggers = local.publish_triggers
}
```

#### Avoid large bulk publishes
From version 2.5.0 the provider was enhanced with support to auto publish mode using `auto_publish_batch_size` or via the `CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE` environment variable to configure the number of batch size to automatically run publish.
<br>Note: To make sure all changes are published need to do publish explicitly at the end of the execution.
```hcl
# Configure the Check Point Provider
provider "checkpoint" {
  server = "chkp-mgmt-srv.local"
  api_key = "admin_api_key"
  context = "web_api"
  auto_publish_batch_size = "100"
}
```

#### Control publish post destroy
From version 2.6.0 the provider was enhanced where a new flag was added `run_publish_on_destroy` to `checkpoint_management_publish` which indicates whether to run publish on destroy.
```hcl
# Make the publish resources dependent of the module and trigger it if there is a change in the configuration files
resource "checkpoint_management_publish" "publish" {
  depends_on = [ module.policy ]
  triggers = local.publish_triggers
  run_publish_on_destroy = true
}
```