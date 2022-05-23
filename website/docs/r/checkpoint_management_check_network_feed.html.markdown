---
layout: "checkpoint"
page_title: "checkpoint_management_check_network_feed"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-check-network-feed"
description: |- This resource allows you to execute Check Point Check Network Feed.
---

# checkpoint_management_check_network_feed

This resource allows you to execute Check Point Check Network Feed.

## Example Usage

```hcl
resource "checkpoint_management_check_network_feed" "example" {
  network_feed = {
    name = "existing_feed"
  }
  targets = ["corporate-gateway"]
}
```

## Argument Reference

The following arguments are supported:

* `network_feed` - (Required) network feed parameters.network_feed blocks are documented below.
* `targets` - (Required) On what targets to execute this command. Targets may be identified by their name, or object
  unique identifier.targets blocks are documented below.

`network_feed` supports the following:

* `name` - (Optional) Object name.
* `feed_url` - (Optional) URL of the feed. URL should be written as http or https.
* `certificate_id` - (Optional) Certificate SHA-1 fingerprint to access the feed.
* `feed_format` - (Optional) Feed file format.
* `feed_type` - (Optional) Feed type to be enforced.
* `password` - (Optional) password for authenticating with the URL.
* `username` - (Optional) username for authenticating with the URL.
* `custom_header` - (Optional) Headers to allow different authentication methods with the URL.custom_header blocks are
  documented below.
* `update_interval` - (Optional) Interval in minutes for updating the feed on the Security Gateway.
* `data_column` - (Optional) Number of the column that contains the feed's data.
* `fields_delimiter` - (Optional) The delimiter that separates between the columns in the feed.
* `ignore_lines_that_start_with` - (Optional) A prefix that will determine which lines to ignore.
* `json_query` - (Optional) JQ query to be parsed.
* `use_gateway_proxy` - (Optional) Use the gateway's proxy for retrieving the feed.
* `domains_to_process` - (Optional) Indicates which domains to process the commands on. It cannot be used with the
  details-level full, must be run from the System Domain only and with ignore-warnings true. Valid values are:
  CURRENT_DOMAIN, ALL_DOMAINS_ON_THIS_SERVER.domains_to_process blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If
  ignore-warnings flag was omitted - warnings will also be ignored.

`custom_header` supports the following:

* `header_name` - (Optional) The name of the HTTP header we wish to add.
* `header_value` - (Optional) The name of the HTTP value we wish to add.

## How To Use

Make sure this command will be executed in the right execution order. note: terraform execution is not sequential.  

