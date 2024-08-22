---
layout: "checkpoint"
page_title: "checkpoint_management_resource_ftp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-ftp"
description: |-
Use this data source to get information on an existing Check Point Resource Ftp.
---

# Data Source: checkpoint_management_resource_ftp

Use this data source to get information on an existing Check Point Resource Ftp.

## Example Usage


```hcl
resource "checkpoint_management_resource_ftp" "example" {
  name = "newFtpResource"
  resource_matching_method = "get_and_put"
  exception_track = "exception log"
  resources_path = "path"
  cvp {
    allowed_to_modify_content = true
    enable_cvp =  false
    reply_order = "return_data_before_content_is_approved"
    server = "serverName"
  }
}

data "checkpoint_management_resource_ftp" "data" {
  uid = "${checkpoint_management_resource_ftp.example.id}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name. 
* `exception_track` -  The UID or Name of the exception track to be used to log actions taken as a result of a match on the resource. 
* `resources_path` - Refers to a location on the FTP server. 
* `resource_matching_method` -  GET allows Downloads from the server to the client. PUT allows Uploads from the client to the server. 
* `cvp` -  Configure CVP inspection on mail messages.cvp blocks are documented below.
* `tags` -  Collection of tag identifiers.tags blocks are documented below.
* `color` -  Color of the object. Should be one of existing colors. 
* `comments` -  Comments string. 


`cvp` supports the following:

* `enable_cvp` -  Select to enable the Content Vectoring Protocol. 
* `server` -  The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application. 
* `allowed_to_modify_content` -  Configures the CVP server to inspect but not modify content. 
* `reply_order` -  Designates when the CVP server returns data to the Security Gateway security server. 
