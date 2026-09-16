---
layout: "checkpoint"
page_title: "checkpoint_management_data_https_rule"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-https-rule"
description: |-
  Use this data source to get information on an existing Check Point Https Rule.
---

# Data Source: checkpoint_management_data_https_rule

Use this data source to get information on an existing Check Point Https Rule.

## Example Usage


```hcl
resource "checkpoint_management_https_rule" "https_rule" {
    name = "HTTPS Rule"
    position {top = "top"}
    layer = "Default Layer"
    blade = ["IPS"]
    destination = ["Internet"]
    enabled = true
    service = ["HTTPS default services"]
    source = ["DMZNet"]
    install_on = ["Policy HTTPS Targets"]
    site_category = ["Any"]
}

data "checkpoint_management_data_https_rule" "data_https_rule" {
    rule_number = "1"
    layer = "${checkpoint_management_https_rule.https_rule.layer}"
}
```

## Argument Reference

The following arguments are supported:

* `layer` - (Required) Layer that holds the Object. Identified by the Name or UID. 
* `rule_number` - (Optional) Rule number.
* `uid` - (Optional) Object unique identifier. 
* `name` - HTTPS rule name. 
* `destination` - Collection of Network objects identified by Name or UID that represents connection destination.
* `service` - Collection of Network objects identified by Name or UID that represents connection service.
* `source` - Collection of Network objects identified by Name or UID that represents connection source.
* `action` - Rule inspect level. "Bypass" or "Inspect". 
* `blade` - Blades for HTTPS Inspection. Identified by Name or UID to enable the inspection for. "Anti Bot","Anti Virus","Application Control","Data Awareness","DLP","IPS","Threat Emulation","Url Filtering".
* `certificate` - Internal Server Certificate identified by Name or UID. otherwise, "Outbound Certificate" is a default value. 
* `destination_negate` - TRUE if "negate" value is set for Destination. 
* `enabled` - Enable/Disable the rule. 
* `install_on` - Which Gateways identified by the name or UID to install the policy on.
* `service_negate` - TRUE if "negate" value is set for Service. 
* `site_category` - Collection of Site Categories objects identified by the name or UID.
* `site_category_negate` - TRUE if "negate" value is set for Site Category. 
* `source_negate` - TRUE if "negate" value is set for Source. 
* `track` - "None","Log","Alert","Mail","SNMP trap","Mail","User Alert", "User Alert 2", "User Alert 3". 
* `comments` - Comments string.
* `hits` - Hits count object.hits blocks are documented below.


`hits` supports the following:

* `first_date` - N/A.first_date blocks are documented below.
* `last_date` - N/A.last_date blocks are documented below.
* `level` - N/A.
* `percentage` - N/A.
* `value` - N/A.


`first_date` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.


`last_date` supports the following:

* `iso_8601` - Date and time represented in international ISO 8601 format.
* `posix` - Number of milliseconds that have elapsed since 00:00:00, 1 January 1970.
