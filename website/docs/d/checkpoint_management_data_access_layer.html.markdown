---
layout: "checkpoint"
page_title: "checkpoint_management_data_access_layer"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-access-layer"
description: |-
  This resource allows you to execute Check Point Access Layer.
---

# Data Source: checkpoint_management_data_access_layer

Use this data source to get information on an existing Check Point Access Layer.

## Example Usage


```hcl
resource "checkpoint_management_access_layer" "access_layer" {
    name = "Access Layer 1"
    detect_using_x_forward_for = false
    applications_and_url_filtering = true
}

data "checkpoint_management_data_access_layer" "data_access_layer" {
    name = "${checkpoint_management_access_layer.access_layer.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. 
* `uid` - (Optional) Object unique identifier.
* `applications_and_url_filtering` - Whether to enable Applications & URL Filtering blade on the layer. 
* `content_awareness` - Whether to enable Content Awareness blade on the layer. 
* `detect_using_x_forward_for` - Whether to use X-Forward-For HTTP header, which is added by the  proxy server to keep track of the original source IP. 
* `firewall` - Whether to enable Firewall blade on the layer. 
* `implicit_cleanup_action` - The default "catch-all" action for traffic that does not match any explicit or implied rules in the layer. 
* `mobile_access` - Whether to enable Mobile Access blade on the layer. 
* `shared` - Whether this layer is shared. 
* `tags` - Collection of tag identifiers.
* `color` - Color of the object. Should be one of existing colors. 
* `comments` - Comments string. 
