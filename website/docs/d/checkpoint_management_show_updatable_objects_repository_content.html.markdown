---
layout: "checkpoint"
page_title: "checkpoint_management_show_updatable_objects_repository_content"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-show-updatable-objects-repository-content"
description: |-
  This resource allows you to execute Check Point Show Updatable Objects Repository Content.
---

# Data Source: checkpoint_management_show_updatable_objects_repository_content

This resource allows you to execute Check Point Show Updatable Objects Repository Content.

## Example Usage


```hcl
data "checkpoint_management_show_updatable_objects_repository_content" "query" {
    filter = {
        text = "API Gateway"
    }
}
```

## Argument Reference

The following arguments are supported:
* `uid_in_updatable_objects_repository` - (Optional) The object's unique identifier in the Updatable Objects repository.
* `filter` - (Optional) Return results matching the specified filter. filter blocks blocks are documented below.
* `limit` - (Optional) The maximal number of returned results.
* `offset` - (Optional) Number of the results to initially skip.
* `order` - (Optional) Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. order blocks are documented below.
* `from` - From which element number the query was done.
* `to` - To which element number the query was done.
* `total` - Total number of elements returned by the query.
* `objects` - Collection of retrieved Updatable Objects. objects blocks blocks are documented below.

`filter` supports the following:

* `text` - (Optional) Return results containing the specified text value.
* `uri` - (Optional) Return results under the specified uri value.
* `parent_uid_in_updatable_objects_repository` - (Optional) Return results under the specified Updatable Object.

`order` supports the following:

* `asc` - (Optional) Sorts results by the given field in ascending order.
* `desc` - (Optional) Sorts results by the given field in descending order.

`objects` supports the following:

* `name_in_updatable_objects_repository` - Object name in the Updatable Objects Repository.
* `uid_in_updatable_objects_repository` - Unique identifier of the object in the Updatable Objects Repository.
* `additional_properties` - Additional properties on the object. additional_properties blocks are documented below.
* `updatable_object` - The imported management object (if exists). updatable_object blocks are documented below.

`additional_properties` supports the following:

* `description` - Description of retrieved Updatable Object.
* `info_text` - Information about the Updatable Object IP ranges source.
* `info_url` - URL of the Updatable Object IP ranges source.
* `uri` - URI of the Updatable Object under the Updatable Objects Repository.

`updatable_object` supports the following:

* `name` - Object name. Must be unique in the domain.
* `uid` - Object unique identifier.
* `type` - Object type.
* `domain` - Information about the domain that holds the Object. domain blocks are documented below.

`domain` supports the following:

* `name` - Object name. Must be unique in the domain.
* `uid` - Object unique identifier.
* `domain_type` - Domain type.