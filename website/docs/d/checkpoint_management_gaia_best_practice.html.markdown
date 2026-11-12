---
layout: "checkpoint"
page_title: "checkpoint_management_data_host"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-host"
description: |-
Use this data source to get information on an existing Check Point Host.
---

# Data Source: checkpoint_management_gaia_best_practice

Use this data source to get information on an existing Check Point Host.


## Example Usage


```hcl
resource "checkpoint_management_gaia_best_practice" "gaia_best_practice" {
  name = "Make sure that the network access via Telnet is disabled."
  description = "This Gaia Best Practice makes sure that the network access, via Telnet, is disabled."
  action_item = "Validate that the Telnet settings are disabled on the configuration set on the GAIA OS."
  expected_output_text = "Success"
  practice_script_base64 = "IyEvYmluL2Jhc2gKCnRlbG5ldF9vZmY9JChjbGlzaCAtYyAic2hvdyBjb25maWd1cmF0aW9uIiB8IGdyZXAgInNldCBuZXQtYWNjZXNzIHRlbG5ldCIgfCBncmVwICJvZmYiKQppZiBbICEgLXogIiR0ZWxuZXRfb2ZmIiBdOyB0aGVuCgllY2hvIFN1Y2Nlc3MKZWxzZQoJZWNobyBGYWlsCmZp"
}

data "checkpoint_management_gaia_best_practice" "data_gaia_best_practice" {
    name = "${checkpoint_management_gaia_best_practice.gaia_best_practice.name}"
}
```

## Argument Reference

The following arguments are supported:

* `best_practice_id` - (Optional) Best Practice ID.
* `name` - (Optional) Best Practice Name.
* `uid` - (Optional) Best Practice UID.
* `action_item` - Action item to comply with the Best Practice.
* `description` - Description of the Best Practice.
* `expected_output_base64` - The expected output of the script in Base64. Available only for user-defined best practices.
* `practice_script_base64` - The script to run on Gaia Security Gateways during the Compliance scans in Base64. Available only for user-defined best practices.
* `regulations` - The applicable regulations of the Gaia Best Practice. Appear only when the value of the 'details-level' parameter is set to 'full'. regulations blocks are documented below.
* `relevant_objects` - The applicable objects of the Gaia Best Practice. Appear only when the value of the 'details-level' parameter is set to 'full'. relevant_objects blocks are documented below.
* `status` - The current status of the Best Practice.
* `user_defined` - Determines if the Gaia Best Practice is a user-defined best practice.

`regulations` supports the following:

* `regulation_name` - The name of the regulation.
* `requirement_description` - The description of the requirement.
* `requirement_id` - The id of the requirement.
* `requirement_status` - The status of the requirement.


`relevant_objects` supports the following:

* `enabled` - Determines if the relevant object is enabled or not.
* `name` - The name of the relevant object.
* `status` - The status of the relevant object.
* `uid` - The uid of the relevant object.