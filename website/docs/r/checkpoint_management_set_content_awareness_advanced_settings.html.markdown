---
layout: "checkpoint"
page_title: "checkpoint_management_set_content_awareness_advanced_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-content-awareness-advanced-settings"
description: |-
This resource allows you to execute Check Point Set Content Awareness Advanced Settings.
---

# checkpoint_management_set_content_awareness_advanced_settings

This resource allows you to execute Check Point Set Content Awareness Advanced Settings.

## Example Usage


```hcl
resource "checkpoint_management_set_content_awareness_advanced_settings" "example" {
  internal_error_fail_mode = "block connections"
  supported_services = ["https","http","ftp"]
  httpi_non_standard_ports = false
  inspect_archives = false
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Computed) Object unique identifier.
* `internal_error_fail_mode` - (Optional) In case of internal system error, allow or block all connections. 
* `supported_services` - (Optional) Specify the services that Content Awareness inspects.supported_services blocks are documented below.
* `httpi_non_standard_ports` - (Optional) Servers usually send HTTP traffic on TCP port 80. Some servers send HTTP traffic on other ports also. By default, this setting is enabled and Content Awareness inspects HTTP traffic on non-standard ports. You can disable this setting and configure Content Awareness to inspect HTTP traffic only on port 80. 
* `inspect_archives` - (Optional) Examine the content of archive files. For example, files with the extension .zip, .gz, .tgz, .tar.Z, .tar, .lzma, .tlz, 7z, .rar. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

