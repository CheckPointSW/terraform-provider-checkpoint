---
layout: "checkpoint"
page_title: "checkpoint_management_command_get_interfaces"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-get-interfaces"
description: |-
This resource allows you to execute Check Point Get Interfaces.
---

# Resource: checkpoint_management_command_get_interfaces

This resource allows you to execute Check Point Get Interfaces.

```hcl
resource "checkpoint_management_command_get_interfaces" "get_interfaces" {
  target_uid = "2220d9ad-a251-5555-9a0a-4772a6511111"
}
```

## Argument Reference

The following arguments are supported:

* `target_uid` - (Required) Target unique identifier.
* `target_name` - (Required) Target name.
* `group_interfaces_by_subnet` - (Optional) Specify whether to group the cluster interfaces by a subnet. Otherwise, group the cluster interfaces by their names.
* `use_defined_by_routes` - (Optional) Specify whether to configure the topology "Defined by Routes" where applicable. Otherwise, configure the topology to "This Network" as default for internal interfaces.
* `with_topology` - (Optioanl) Specify whether to fetch the interfaces with their topology. Otherwise, the Management Server fetches the interfaces without their topology.
* `task_id` - The UID of the "get-interfaces" task. Use the "show-task" command to check the progress of the "get-interfaces" task.

