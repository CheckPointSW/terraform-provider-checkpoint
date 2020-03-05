---
layout: "checkpoint"
page_title: "checkpoint_management_migrate_export_domain"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-migrate-export-domain"
description: |-
This resource allows you to execute Check Point Migrate Export Domain.
---

# checkpoint_management_migrate_export_domain

This resource allows you to execute Check Point Migrate Export Domain.

## Example Usage


```hcl
resource "checkpoint_management_migrate_export_domain" "example" {
  domain = "domain1"
  file_path = "/var/log/domain1_exported.tgz"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) Domain can be identified by name or UID.<br><font color="red">Required only for</font> exporting domain from Multi-Domain Server. 
* `file_path` - (Optional) Path in which the exported domain data will be saved. <br>Should be the directory path or the full file path with ".tgz" <br>If no path was inserted the default will be: "/var/log/&lt;domain name&gt;_&lt;date&gt;.tgz". 
* `include_logs` - (Optional) Export logs. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

