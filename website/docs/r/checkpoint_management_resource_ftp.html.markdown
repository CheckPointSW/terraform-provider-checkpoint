---
layout: "checkpoint"
page_title: "checkpoint_management_resource_ftp"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-resource-ftp"
description: |-
This resource allows you to execute Check Point Resource Ftp.
---

# checkpoint_management_resource_ftp

This resource allows you to execute Check Point Resource Ftp.

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
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `exception_track` - (Optional) The UID or Name of the exception track to be used to log actions taken as a result of a match on the resource. 
* `resources_path` - (Optional) Refers to a location on the FTP server. 
* `resource_matching_method` - (Optional) GET allows Downloads from the server to the client. PUT allows Uploads from the client to the server. 
* `cvp` - (Optional) Configure CVP inspection on mail messages.cvp blocks are documented below.
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`cvp` supports the following:

* `enable_cvp` - (Optional) Select to enable the Content Vectoring Protocol. 
* `server` - (Optional) The UID or Name of the CVP server, make sure the CVP server is already be defined as an OPSEC Application. 
* `allowed_to_modify_content` - (Optional) Configures the CVP server to inspect but not modify content. 
* `reply_order` - (Optional) Designates when the CVP server returns data to the Security Gateway security server. 
