---
layout: "checkpoint"
page_title: "checkpoint_generic_api"
sidebar_current: "docs-checkpoint-resource-checkpoint-generic-api"
description: |-
This resource allows you to execute generic Management API calls.
---

# Resource: checkpoint_generic_api

This resource allows you to execute Check Point generic Management API or GAIA API.<br>
See the [Management API reference](https://sc1.checkpoint.com/documents/latest/APIs/index.html) or [GAIA API reference](https://sc1.checkpoint.com/documents/latest/GaiaAPIs/index.html) for a complete list of APIs you can run on your Check Point server.<br>
<b>NOTE:</b> If you configure the provider [context](https://registry.terraform.io/providers/CheckPointSW/checkpoint/latest/docs#context-1) to `gaia_api` you can run only GAIA API and GAIA resources. Management API or any other resource will not be supported.

## Example Usage


```hcl
# Run generic Management API when provider context is 'web_api'
resource "checkpoint_generic_api" "api1" {
  api_command = "add-host"
  payload = <<EOT
  {
    "name": "host1",
    "ip-address": "1.2.3.4"
  }
  EOT
}

# Run generic Management API when provider context is 'web_api'
resource "checkpoint_generic_api" "api2" {
  api_command = "show-hosts"
}

# Run generic Management API when provider context is 'web_api'
resource "checkpoint_generic_api" "api3" {
  api_command = "gaia-api/show-proxy"
  payload = <<EOT
  {
    "target": "gateway1"
  }
  EOT
}

# Run generic GAIA API when provider context is 'gaia_api'
resource "checkpoint_generic_api" "api4" {
  api_command = "show-proxy"
}
```

## Argument Reference

The following arguments are supported:

* `api_command` - (Required) API command name or path.
* `payload` - (Optional) Request payload in JSON format. You can use [heredoc strings](https://developer.hashicorp.com/terraform/language/expressions/strings#heredoc-strings) to write freestyle JSON.
* `method` - (Optional) HTTP request method. Default is `POST`.
* `response` - Response message in JSON format.