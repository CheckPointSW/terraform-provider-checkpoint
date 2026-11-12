---
layout: "checkpoint"
page_title: "checkpoint_management_subordinate_ca"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-subordinate-ca"
description: |-
 Use this data source to get information on an existing Check Point Subordinate Ca.
---

# Data Source: checkpoint_management_subordinate_ca

Use this data source to get information on an existing Check Point Subordinate Ca.

## Example Usage
```hcl
data "checkpoint_management_subrodinate_ca" "data_test" {
    name = "subordinate_ca1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name. Should be unique in the domain.
* `uid` - (Optional) Object unique identifier.
* `automatic_enrollment` - Certificate automatic enrollment.automatic_enrollment blocks are documented below.
* `base64_certificate` - Certificate file encoded in base64.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `icon` - Object icon.
* `tags` - Collection of tag objects identified by the name or UID. Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.


`automatic_enrollment` supports the following:

* `automatically_enroll_certificate` - Whether to automatically enroll certificate.
* `protocol` - Protocol that communicates with the certificate authority. Available only if "automatically-enroll-certificate" parameter is set to true.
* `scep_settings` - Scep protocol settings. Available only if "protocol" is set to "scep".scep_settings blocks are documented below.
* `cmpv1_settings` - Cmpv1 protocol settings. Available only if "protocol" is set to "cmpv1".cmpv1_settings blocks are documented below.
* `cmpv2_settings` - Cmpv2 protocol settings. Available only if "protocol" is set to "cmpv2".cmpv2_settings blocks are documented below.


`scep_settings` supports the following:

* `ca_identifier` - Certificate authority identifier.
* `url` - Certificate authority URL.


`cmpv1_settings` supports the following:

* `direct_tcp_settings` - Direct tcp transport layer settings.direct_tcp_settings blocks are documented below.


`cmpv2_settings` supports the following:

* `transport_layer` - Transport layer.
* `direct_tcp_settings` - Direct tcp transport layer settings.direct_tcp_settings blocks are documented below.
* `http_settings` - Http transport layer settings.http_settings blocks are documented below.


`direct_tcp_settings` supports the following:

* `ip_address` - Certificate authority IP address.
* `port` - Port number.


`direct_tcp_settings` supports the following:

* `ip_address` - Certificate authority IP address.
* `port` - Port number.


`http_settings` supports the following:

* `url` - Certificate authority URL.
