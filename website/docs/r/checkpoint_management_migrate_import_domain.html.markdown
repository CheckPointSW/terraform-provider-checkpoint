---
layout: "checkpoint"
page_title: "checkpoint_management_migrate_import_domain"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-migrate-import-domain"
description: |-
This resource allows you to execute Check Point Migrate Import Domain.
---

# checkpoint_management_migrate_import_domain

This resource allows you to execute Check Point Migrate Import Domain.

## Example Usage


```hcl
resource "checkpoint_management_migrate_import_domain" "example" {
  file_path = "/var/log/domain1_exported.tgz"
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Required) Path to the exported file to be imported. <br>Should be the full file path (example, "/var/log/domain1_exported.tgz"). 
* `domain_ip_address` - (Required) IPv4 address.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server. 
* `domain_name` - (Required) Domain name. Should be unique in the MDS.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server. 
* `domain_server_name` - (Required) Multi Domain server name.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server. 
* `include_logs` - (Optional) Import logs from the input package. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

