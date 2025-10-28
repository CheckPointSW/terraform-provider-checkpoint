---
layout: "checkpoint"
page_title: "checkpoint_management_set_default_administrator_settings"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-set-default-administrator-settings"
description: |-
 This resource allows you to execute Check Point Set Default Administrator Settings.
---

# checkpoint_management_set_default_administrator_settings

This resource allows you to execute Check Point Set Default Administrator Settings.

## Example Usage


```hcl
resource "checkpoint_management_set_default_administrator_settings" "example" {
  expiration_type = "expiration date"
  expiration_date = "2025-06-23"
  indicate_expiration_in_admin_view = false
  notify_expiration_to_admin = true
  days_to_notify_expiration_to_admin = 5
}
```

## Argument Reference

The following arguments are supported:

* `authentication_method` - (Optional) Authentication method for new administrator. 
* `expiration_type` - (Optional) Expiration type for new administrator. 
* `expiration_date` - (Optional) Expiration date for new administrator in YYYY-MM-DD format. <font color="red">Required only when</font> 'expiration-type' is set to 'expiration date'. 
* `expiration_period` - (Optional) Expiration period for new administrator. <font color="red">Required only when</font> 'expiration-type' is set to 'expiration period'. 
* `expiration_period_time_units` - (Optional) Expiration period time units for new administrator. <font color="red">Required only when</font> 'expiration-type' is set to 'expiration period'. 
* `indicate_expiration_in_admin_view` - (Optional) Indicates whether to notify administrator about expiration. 
* `notify_expiration_to_admin` - (Optional) Indicates whether to show 'about to expire' indication in administrator view. 
* `days_to_indicate_expiration_in_admin_view` - (Optional) Number of days in advanced to show 'about to expire' indication in administrator view. 
* `days_to_notify_expiration_to_admin` - (Optional) Number of days in advanced to notify administrator about expiration. 


## How To Use
Make sure this command will be executed in the right execution order. 
note: terraform execution is not sequential.  

