---
layout: "checkpoint"
page_title: "checkpoint_management_install_policy "
sidebar_current: "docs-checkpoint-resource-checkpoint-management-install-policy"
description: |-
  Install the published policy.
---

# checkpoint_management_install_policy

Install the published policy.

## Example Usage

```hcl
resource "checkpoint_management_install_policy" "example" {
  policy_package = "standard"
  targets = ["corporate-gateway"]
}
```

## Argument Reference

The following arguments are supported:

* `policy_package` - (Required) The name of the Policy Package to be installed.
* `targets` - (Required) On what targets to execute this command. Targets may be identified by their name, or object unique identifier.
* `access` - (Optional) Set to be true in order to install the Access Control policy. By default, the value is true if Access Control policy is enabled on the input policy package, otherwise false.
* `desktop_security` - (Optional) Set to be true in order to install the Desktop Security policy. By default, the value is true if desktop security policy is enabled on the input policy package, otherwise false.
* `qos` - (Optional) Set to be true in order to install the QoS policy. By default, the value is true if Quality-of-Service policy is enabled on the input policy package, otherwise false.
* `threat_prevention` - (Optional) Set to be true in order to install the Threat Prevention policy. By default, the value is true if Threat Prevention policy is enabled on the input policy package, otherwise false.
* `install_on_all_cluster_members_or_fail` - (Optional) Relevant for the gateway clusters. If true, the policy is installed on all the cluster members. If the installation on a cluster member fails, don't install on that cluster.
* `prepare_only` - (Optional) If true, prepares the policy for the installation, but doesn't install it on an installation target.
* `revision` - (Optional) The UID of the revision of the policy to install.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  



