---
layout: "checkpoint"
page_title: "checkpoint_management_smtp_server"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-smtp-server"
description: |-
Use this data source to get information on an existing Check Point Smtp Server.
---

# Data Source: checkpoint_management_smtp_server

Use this data source to get information on an existing Check Point Smtp Server.

## Example Usage


```hcl
resource "checkpoint_management_smtp_server" "example" {
  name = "SMTP1"
  server = "smtp.example.com"
  port = 25
}

data "checkpoint_management_smtp_server" "data_smtp_server" {
  name = "${checkpoint_management_smtp_server.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Optional) Object unique identifier.
* `name` - (Optional) Object name. 
