---
layout: "checkpoint"
page_title: "checkpoint_management_data_data_center_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-data-center-object"
description: |- Use this data source to get information on an existing Check Point Data Center Object.
---

# checkpoint_management_data_data_center_object

Use this data source to get information on an existing Check Point Data Center Object.

## Example Usage

```hcl
resource "checkpoint_management_data_center_object" "dco1"{
  data_center_name = "myAws1"
  uri = "/Region - EU (Frankfurt)/VPCs/vpc-0e5983c1d08b53e75/VPCEndpoints"
  name = "my_data_object_center"
}

data "checkpoint_management_data_center_object" "data_dco" {
  name = "${checkpoint_management_data_center_object.dco1.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `uid_in_data_center` -  Unique identifier of the object in the Data Center.
* `data_center` - The Data Center the object is on.
* `updated_on_data_center` - Last update time in the Data Center
* `deleted` - Indicates if the object is inaccessible or deleted on Data Center Server.
* `groups` - Collection of group identifiers.
* `tags` - Collection of tag identifiers.
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
