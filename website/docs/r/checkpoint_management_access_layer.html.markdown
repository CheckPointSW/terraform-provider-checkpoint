---
layout: "checkpoint"
page_title: "checkpoint_management_access_layer"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-access-layer"
description: |-
This resource allows you to execute Check Point Access Layer.
---

# checkpoint_management_access_layer

This resource allows you to execute Check Point Access Layer.

## Example Usage


```hcl
resource "checkpoint_management_access_layer" "example" {
  name = "New Layer 1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `applications_and_url_filtering` - (Optional) Whether to enable Applications & URL Filtering blade on the layer. 
* `content_awareness` - (Optional) Whether to enable Content Awareness blade on the layer. 
* `detect_using_x_forward_for` - (Optional) Whether to use X-Forward-For HTTP header, which is added by the  proxy server to keep track of the original source IP. 
* `firewall` - (Optional) Whether to enable Firewall blade on the layer. 
* `implicit_cleanup_action` - (Optional) The default "catch-all" action for traffic that does not match any explicit or implied rules in the layer. 
* `mobile_access` - (Optional) Whether to enable Mobile Access blade on the layer. 
* `shared` - (Optional) Whether this layer is shared. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `add_default_rule` - (Optional) Indicates whether to include a cleanup rule in the new layer. 
