---
layout: "checkpoint"
page_title: "checkpoint_management_get_attachment"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-get-attachment"
description: |-
This resource allows you to execute Check Point Get Attachment.
---

# Resource: checkpoint_management_get_attachment

This command resource allows you to execute Check Point Get Attachment.

## Example Usage


```hcl
resource "checkpoint_management_get_attachment" "example" {
  attachment_id = "MjY5HlNtYXJ0RGVmZW5zZR5jbj1jcF9tZ210LG89aHVnbzEtYmxvYkFwaS1uZXctdGFrZS0yLmNoZWNrcG9pbnQuY29tLnM2MjdvMx57MHg1OTg4"
}
```

## Argument Reference

The following arguments are supported:

* `attachment_id` - (Optional) Attachment identifier from a log record. 
* `uid` - (Optional) Log id from a log record. 
* `task_id` - (Computed) Asynchronous task unique identifier. 


## How To Use
Make sure this command will be executed in the right execution order.