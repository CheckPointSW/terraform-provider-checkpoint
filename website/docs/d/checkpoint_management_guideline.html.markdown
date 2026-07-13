---
layout: "checkpoint"
page_title: "checkpoint_management_guideline"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-guideline"
description: |-
Use this data source to get information on an existing Check Point Guideline.
---

# Data Source: checkpoint_management_guideline

Use this data source to get information on an existing Check Point Guideline.

## Example Usage

```hcl
data "checkpoint_management_guideline" "data_test" {
    name = "guideline1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object unique identifier.
* `show_indexing_status` - (Optional) Control whether to show the indexing status of the guideline.
* `indexing_status_layer` - (Optional) Relevant only when show-indexing-status is true. The access-layer to show the indexing status of (identified by unique id or 'any' for all attached access-layers).
* `dereference_group_members` - (Optional) Indicates whether to dereference "members" field by details level for every object in reply.
* `show_membership` - (Optional) Indicates whether to calculate and show "groups" field for every object in reply.
* `access_layers` - (Computed) The access-layers objects attached to the guideline with their policy-package context. access_layers blocks are documented below.
* `cell_actions_override` - (Computed) All the cells that the user changed the default action in. cell_actions_override blocks are documented below.
* `default_action` - (Computed) The default action for guideline cells with two different groups.
* `default_self_action` - (Computed) The default action for guideline cells with the same group in both axis.
* `guideline_groups` - (Computed) The segments displayed in the guideline matrix in at least one of the axes (from or to). guideline_groups blocks are documented below.
* `indexing_status` - (Computed) Task-id map for the indexing tasks of the guideline. indexing_status blocks are documented below.
* `color` - (Computed) Color of the object. Should be one of existing colors.
* `comments` - (Computed) Comments string.
* `icon` - (Computed) Object icon.
* `tags` - (Computed) Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.


`access_layers` supports the following:

* `access_layer` - (Computed) The access-layer object attached to the guideline identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `policy_package` - (Computed) The policy-package object context for the access-layer (only for global access-layers) identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.


`cell_actions_override` supports the following:

* `from` - (Computed) Unique identifier of the segment of the cell in the 'from' axis.
* `from_type` - (Computed) The type of the segment in the 'from' axis.
* `to` - (Computed) Unique identifier of the segment of the cell in the 'to' axis.
* `to_type` - (Computed) The type of the segment in the 'to' axis.
* `action` - (Computed) The action selected for the cell.


`guideline_groups` supports the following:

* `guideline_group` - (Computed) The network-group object identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `members` - (Computed) Group members.
* `position` - (Computed) The position of the guideline group in the axis.


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
