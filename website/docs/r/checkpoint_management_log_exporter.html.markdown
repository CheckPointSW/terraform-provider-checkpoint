---
layout: "checkpoint"
page_title: "checkpoint_management_log_exporter"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-log-exporter"
description: |-
This resource allows you to execute Check Point Log Exporter.
---

# checkpoint_management_log_exporter

This resource allows you to execute Check Point Log Exporter.

## Example Usage


```hcl
resource "checkpoint_management_log_exporter" "example" {
  name = "newLogExporter"
  target_server = "1.2.3.4"
  target_port = 1234
  protocol = "tcp"
  attachments {
    add_link_to_log_attachment = true
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `target_server` - (Required) Target server port to which logs are exported. 
* `target_port` - (Required) Port number of the target server. 
* `protocol` - (Optional) Protocol used to send logs to the target server. 
* `enabled` - (Optional) Indicates whether to enable export. 
* `attachments` - (Optional) Log exporter attachments. attachments blocks are documented below.
* `data_manipulation` - (Optional) Log exporter data manipulation. data_manipulation blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 


`attachments` supports the following:

* `add_link_to_log_attachment` - (Optional) Indicates whether to add link to log attachment in SmartView. 
* `add_link_to_log_details` - (Optional) Indicates whether to add link to log details in SmartView. 
* `add_log_attachment_id` - (Optional) Indicates whether to add log attachment ID. 


`data_manipulation` supports the following:

* `aggregate_log_updates` - (Optional) Indicates whether to aggregate log updates. 
* `format` - (Optional) Logs format. 
