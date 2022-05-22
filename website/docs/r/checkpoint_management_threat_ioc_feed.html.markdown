---
layout: "checkpoint"
page_title: "checkpoint_management_threat_ioc_feed"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-threat-ioc-feed"
description: |-
This resource allows you to execute Check Point Threat Ioc Feed.
---

# checkpoint_management_threat_ioc_feed

This resource allows you to execute Check Point Threat Ioc Feed.

## Example Usage


```hcl
resource "checkpoint_management_threat_ioc_feed" "example" {
  name = "ioc_feed"
  feed_url = "https://www.feedsresource.com/resource"
  action = "Prevent"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `feed_url` - (Optional) URL of the feed.
URL should be written as http or https. 
* `action` - (Optional) The feed indicator's action. 
* `certificate_id` - (Optional) Certificate SHA-1 fingerprint to access the feed. 
* `custom_comment` - (Optional) Custom IOC feed - the column number of comment. 
* `custom_confidence` - (Optional) Custom IOC feed - the column number of confidence. 
* `custom_header` - (Optional) Custom HTTP headers.custom_header blocks are documented below.
* `custom_name` - (Optional) Custom IOC feed - the column number of name. 
* `custom_severity` - (Optional) Custom IOC feed - the column number of severity. 
* `custom_type` - (Optional) Custom IOC feed - the column number of type in case a specific type is not chosen. 
* `custom_value` - (Optional) Custom IOC feed - the column number of value in case a specific type is chosen. 
* `enabled` - (Optional) Sets whether this indicator feed is enabled. 
* `feed_type` - (Optional) Feed type to be enforced. 
* `password` - (Optional) password for authenticating with the URL. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `use_custom_feed_settings` - (Optional) Set in order to configure a custom indicator feed. 
* `username` - (Optional) username for authenticating with the URL. 
* `fields_delimiter` - (Optional) The delimiter that separates between the columns in the feed. 
* `ignore_lines_that_start_with` - (Optional) A prefix that will determine which lines to ignore. 
* `use_gateway_proxy` - (Optional) Use the gateway's proxy for retrieving the feed. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`custom_header` supports the following:

* `header_name` - (Optional) The name of the HTTP header we wish to add. 
* `header_value` - (Optional) The name of the HTTP value we wish to add. 
