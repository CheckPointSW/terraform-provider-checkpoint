---
layout: "checkpoint"
page_title: "checkpoint_management_add_data_center_object"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-add-data-center-object"
description: |-
This resource allows you to execute Check Point Add Data Center Object.
---

# checkpoint_management_add_data_center_object

This resource allows you to execute Check Point Add Data Center Object.

## Example Usage


```hcl
resource "checkpoint_management_add_data_center_object" "example" {
  name = "VM1 mgmt name"
  uri = "/Datacenters/VMs/My VM1"
  data_center_name = "vCenter 1"
}
```

## Argument Reference

The following arguments are supported:

* `data_center_name` - (Required) Name of the Data Center Server the object is in. 
* `data_center_uid` - (Required) Unique identifier of the Data Center Server the object is in. 
* `uri` - (Required) URI of the object in the Data Center Server. 
* `uid_in_data_center` - (Required) Unique identifier of the object in the Data Center Server. 
* `name` - (Optional) Override default name on data-center. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `groups` - (Optional) Collection of group identifiers.groups blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

