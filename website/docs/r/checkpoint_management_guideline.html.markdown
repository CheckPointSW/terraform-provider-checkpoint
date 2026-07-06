---
layout: "checkpoint"
page_title: "checkpoint_management_guideline"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-guideline"
description: |-
This resource allows you to execute Check Point Guideline.
---

# checkpoint_management_guideline

This resource allows you to execute Check Point Guideline.

## Example Usage


```hcl
resource "checkpoint_management_guideline" "example" {
  name = "Corporate policy"
  access_layers = ["Network",]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `access_layers` - (Required) The access-layers (one or more) that will be attached to the guideline, identified by name or UID.access_layers blocks are documented below.
* `guideline_groups` - (Required) The groups that will be part of the guideline (guideline should have between 2-12 segments, including internet-segment and other-segment). It is recommended to select groups that best represent segments of the network.guideline_groups blocks are documented below.
* `cell_actions_override` - (Optional) Cells that their action will override the default actions of the guideline.cell_actions_override blocks are documented below.
* `dereference_group_members` - (Optional) Indicates whether to dereference "members" field by details level for every object in reply. 
* `show_membership` - (Optional) Indicates whether to calculate and show "groups" field for every object in reply.
* `indexing_status` - (Computed) Task-id map for the indexing tasks of the guideline. indexing_status blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`access_layers` supports the following:

* `access_layer` - (Required) Access-layer attached to guideline identified by the name or UID.if Access-Layer is in the global domain due to Global Assignment Local domain Package is required. 
* `policy_package` - (Optional) Policy package context for the access-layer attached to guideline identified by the name or UID.Package will be ignored if the access-layer is local. 


`guideline_groups` supports the following:

* `name` - (Optional) Network group name.
* `position` - (Optional) Guideline-Group Position in the guideline. If a position is specified for one guideline-group, it is required for all guideline-groups.
* `members` - (Computed) Network group members identified by name.


`cell_actions_override` supports the following:

* `from` - (Optional) The segment identifier (name or UID) of the cell in the 'from' axis. The field is mandatory only if "from-type" is "network group". 
* `from_type` - (Optional) The type of the segment in the 'from' axis. 
* `to` - (Optional) The segment identifier (name or UID) of the cell in the 'to' axis. The field is mandatory only if "to-type" is "network group". 
* `to_type` - (Optional) The type of the segment in the 'to' axis. 
* `action` - (Optional) The action to be applied to the cell. The field is mandatory at add command. 
* `allowed_services` - (Optional) Services (identified by name or UID) that are allowed in the cell. Relevant only if the action in the cell is 'All traffic is not allowed'. To remove allowed-services call update with the same "All traffic is not allowed" action, or remove the cell-action-override.allowed_services blocks are documented below.


`indexing_status` supports the following:

* `access_layer_id` - (Computed) The id of the access-layer that is being indexed.
* `indexing_message` - (Computed) Message which offers more details on The indexing task.
* `indexing_task` - (Computed) The id of the task that is indexing the access-layer. Relevant only if the task is in progress.
* `last_update_time` - (Computed) Last time the indexing status was updated. last_update_time blocks are documented below.
* `policy_package_id` - (Computed) The id of the policy Package that is being indexed.(only used if the layer is global).
* `status` - (Computed) The status of the indexing task.


`last_update_time` supports the following:

* `iso_8601` - (Computed) Date and time represented in international ISO 8601 format.
* `posix` - (Computed) Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
