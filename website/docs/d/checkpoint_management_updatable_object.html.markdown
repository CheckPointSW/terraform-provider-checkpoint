---
layout: "checkpoint"
page_title: "checkpoint_management_updatable_object"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-updatable-object"
description: |-
Use this data source to get information on an existing Check Point  Updatable Object
---

# Data Source: checkpoint_management_updatable_object

Use this data source to get information on an existing Check Point Updatable Object

## Example Usage


```hcl
data "checkpoint_management_updatable_object" "updatable_object"{
  name = "Amazon US East 1 Services"
}
```

## Argument Reference

* `name` - (Optional) Object Name.
* `uid` - (Optional) Object UID.
* `type` - Object Type.
* `name_in_updatable_objects_repository` - Object name in the Updatable Objects Repository.
* `uid_in_updatable_objects_repository` - Unique identifier of the object in the Updatable Objects Repository.
* `additional_properties` - Additional properties on the object. additional_properties blocks are documented below.
* `updatable_object_meta_info` - Meta Info about the Updatable Object.
* `tags` - Collection of tag identifiers. 
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.

`additional_properties` supports the following:

* `description` - Description of retrieved Updatable Object.
* `info_text` - Information about the Updatable Object IP ranges source.
* `info_url` - URL of the Updatable Object IP ranges source.
* `uri` - URI of the Updatable Object under the Updatable Objects Repository.

`updatable_object_meta_info` supports the following: 

* `updated_on_updatable_objects_repository` - Last update time from the Updatable Objects Repository.

`updated_on_updatable_objects_repository` supports the following:

* `iso_8601` - Last update time in iso format.
*  `posix` -  Last update time in posix format.

