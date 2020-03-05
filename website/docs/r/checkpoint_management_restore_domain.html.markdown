---
layout: "checkpoint"
page_title: "checkpoint_management_restore_domain"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-restore-domain"
description: |-
This resource allows you to execute Check Point Restore Domain.
---

# checkpoint_management_restore_domain

This resource allows you to execute Check Point Restore Domain.

## Example Usage


```hcl
resource "checkpoint_management_restore_domain" "example" {
  file_path = "/var/log/domain1_backup.tgz"
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Required) Path to the backup file to be restored. <br>Should be the full file path (example, "/var/log/domain1_backup.tgz"). 
* `domain_ip_address` - (Required) IPv4 address.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server. 
* `domain_name` - (Required) Domain name. Should be unique in the MDS.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server. 
* `domain_server_name` - (Required) Multi Domain server name.<br><font color="red">Required only for</font> importing Security Management Server into Multi-Domain Server. 
* `verify_only` - (Optional) If true, verify that the import operation is valid for this input file and this environment <br>Note: Restore operation will not be executed. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

