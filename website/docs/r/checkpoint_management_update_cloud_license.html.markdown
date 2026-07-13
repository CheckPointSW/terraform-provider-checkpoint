---
layout: "checkpoint"
page_title: "checkpoint_management_update_cloud_license"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-update-cloud-license"
description: |-
This resource allows you to execute Check Point Update Cloud License.
---

# checkpoint_management_update_cloud_license

This resource allows you to execute Check Point Update Cloud License.

## Example Usage


```hcl
resource "checkpoint_management_update_cloud_license" "example" {
  license = "192.168.1.2 31Dec2026 dTTTTTT-WWWWWW-SSSSSSS-QQQQQQ CPSG-VE+5 CPBS-BECE CPSB-DFW CPSM-C-2 CPSB-VPN CPSB-NPM CPSB-LOGS CPSB-IA CPSB-ADNC CPSB-SSLVWPN-5 CK-66666666"
}
```

## Argument Reference

The following arguments are supported:

* `license` - (Required) The updated license string received from the User Center - without 'cplic put'.

