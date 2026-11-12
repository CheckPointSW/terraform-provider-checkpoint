---
layout: "checkpoint"
page_title: "checkpoint_management_network_feed"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-network-feed"
description: |-
This resource allows you to execute Check Point Network Feed.
---

# checkpoint_management_network_feed

This resource allows you to execute Check Point Network Feed.

## Example Usage


```hcl
resource "checkpoint_management_network_feed" "example" {
  name = "network_feed"
  feed_url = "https://www.feedsresource.com/resource"
  username = "feed_username"
  password = "feed_password"
  feed_format = "Flat List"
  feed_type = "IP Address"
  update_interval = 60
  data_column = 1
  use_gateway_proxy = false
  fields_delimiter = "	"
  ignore_lines_that_start_with = "!"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `feed_url` - (Optional) URL of the feed.
URL should be written as http or https. 
* `certificate_id` - (Optional) Certificate SHA-1 fingerprint to access the feed. 
* `feed_format` - (Optional) Feed file format. 
* `feed_type` - (Optional) Feed type to be enforced. 
* `password` - (Optional) password for authenticating with the URL. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `username` - (Optional) username for authenticating with the URL. 
* `custom_header` - (Optional) Headers to allow different authentication methods with the URL.custom_header blocks are documented below.
* `update_interval` - (Optional) Interval in minutes for updating the feed on the Security Gateway. 
* `data_column` - (Optional) Number of the column that contains the feed's data. 
* `fields_delimiter` - (Optional) The delimiter that separates between the columns in the feed. For feed format 'Flat List' default is '\n'(new line). 
* `ignore_lines_that_start_with` - (Optional) A prefix that will determine which lines to ignore. 
* `json_query` - (Optional) JQ query to be parsed. 
* `use_gateway_proxy` - (Optional) Use the gateway's proxy for retrieving the feed. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are: CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`custom_header` supports the following:

* `header_name` - (Optional) The name of the HTTP header we wish to add. 
* `header_value` - (Optional) The name of the HTTP value we wish to add. 
