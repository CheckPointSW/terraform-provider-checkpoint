---
layout: "checkpoint"
page_title: "checkpoint_management_backup_domain"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-backup-domain"
description: |-
This resource allows you to execute Check Point Backup Domain.
---

# checkpoint_management_backup_domain

This resource allows you to execute Check Point Backup Domain.

## Example Usage


```hcl
resource "checkpoint_management_backup_domain" "example" {
  domain = "domain1"
  file_path = "/var/log/domain1_backup.tgz"
}
```

## Argument Reference

The following arguments are supported:

* `domain` - (Required) Domain can be identified by name or UID. 
* `file_path` - (Optional) Path in which the backup domain data will be saved. <br>Should be the directory path or the full file path with ".tgz" <br>If no path was inserted the default will be: "/var/log/&lt;domain name&gt;_&lt;date&gt;.tgz". 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

