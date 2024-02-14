---
layout: "checkpoint"
page_title: "checkpoint_management_data_center_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-center-object"
description: |- This resource allows you to add/update/delete Check Point Data Center Object.
---

# checkpoint_management_data_center_object

This resource allows you to add/update/delete Check Point Data Center Object.

## Example Usage

```hcl
resource "checkpoint_management_data_center_object" "dco1"{
  data_center_name = "myAws1"
  uri = "/Region - EU (Frankfurt)/VPCs/vpc-0e5983c1d08b53e75/VPCEndpoints"
  name = "my_data_object_center"
}
```

## Argument Reference

The following arguments are supported:

* `data_center_name` - (Optional) Name of the Data Center Server the object is in.
* `data_center_uid` - (Optional) Unique identifier of the Data Center Server the object is in.
* `uri` - (Optional) URI of the object in the Data Center Server.
* `uid_in_data_center` - (Optional) Unique identifier of the object in the Data Center Server.
* `name` - (Optional) Override default name on data-center.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `tags` - (Optional) Collection of tag identifiers.
* `groups` - (Optional)Collection of group identifiers.
* `uid_in_data_center` -  Unique identifier of the object in the Data Center.
* `data_center` - The Data Center the object is on.
* `updated_on_data_center` - Last update time in the Data Center
* `deleted` - Indicates if the object is inaccessible or deleted on Data Center Server.
* `type_in_data_center` - Object type in Data Center.
* `additional_properties` - Additional properties on the object.



`data_center` supports the following:
* `name` - Object name. Should be unique in the domain.
* `uid` - Object unique identifier.
* `automatic_refresh` - Object unique identifier.
* `data_center_type` -  Data Center type.
* `properties` - List of Data Center properties.

`properties` supports the following:
* `name` - property name
* `value` - property value

`updated_on_data_center` supports the following:
* `iso_8601` - Time format.
* `posix` - Time format.

`additional_properties` supports the following:
* `name` - property name
* `value` - property value










