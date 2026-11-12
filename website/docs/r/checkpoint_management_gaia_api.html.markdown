---
layout: "checkpoint"
page_title: "checkpoint_management_command_gaia_api"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-command-gaia-api"
description: |-
This resource allows you to execute Check Point Gaia Api.
---

# Resource: checkpoint_management_command_gaia_api

This resource allows you to run generic `gaia-api` command from the Management.<br>
See the [GAIA API reference](https://sc1.checkpoint.com/documents/latest/GaiaAPIs/index.html) for a complete list of APIs you can run on your Check Point server.<br>
<b>NOTE:</b> Please add a rule to allow the connection from the management to the targets.<br>

## Example Usage


```hcl
resource "checkpoint_management_command_gaia_api" "example1" {
  target = "my_gateway"
  command_name = "show-hostname"
}

resource "checkpoint_management_command_gaia_api" "example2" {
  target = "my_gateway"
  command_name = "show-interface"
  other_parameter = <<EOT
    {
      "name" : "eth0"
    }
  EOT
}

resource "checkpoint_management_command_gaia_api" "example3" {
  target = "my_gateway"
  command_name = "v1.3/show-diagnostics"
  other_parameter = <<EOT
    {
      "category" : "os",
      "topic" : "disk"
    }
  EOT
}
```

## Argument Reference

The following arguments are supported:

* `target` - (Required) Gateway object name or Gateway IP address or Gateway UID. 
* `command_name` - (Required) GAIA API command name or path.
* `other_parameter` - (Optional) Other input parameters for the request payload in JSON format. You can use [heredoc strings](https://developer.hashicorp.com/terraform/language/expressions/strings#heredoc-strings) to write freestyle JSON. 
* `response_message` - Response message in JSON format.


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

