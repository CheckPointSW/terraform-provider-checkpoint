---
layout: "checkpoint"
page_title: "checkpoint_management_gaia_best_practice"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-gaia-best-practice"
description: |-
This resource allows you to execute Check Point Gaia Best Practice.
---

# Resoucre: checkpoint_management_gaia_best_practice

This resource allows you to execute Check Point Gaia Best Practice.

## Example Usage


```hcl
resource "checkpoint_management_gaia_best_practice" "example" {
  name = "Make sure that the network access via Telnet is disabled."
  description = "This Gaia Best Practice makes sure that the network access, via Telnet, is disabled."
  action_item = "Validate that the Telnet settings are disabled on the configuration set on the GAIA OS."
  expected_output_text = "Success"
  practice_script_base64 = "IyEvYmluL2Jhc2gKCnRlbG5ldF9vZmY9JChjbGlzaCAtYyAic2hvdyBjb25maWd1cmF0aW9uIiB8IGdyZXAgInNldCBuZXQtYWNjZXNzIHRlbG5ldCIgfCBncmVwICJvZmYiKQppZiBbICEgLXogIiR0ZWxuZXRfb2ZmIiBdOyB0aGVuCgllY2hvIFN1Y2Nlc3MKZWxzZQoJZWNobyBGYWlsCmZp"
}
```

## Argument Reference

The following arguments are supported:

* `best_practice_id` - (Optional) Best Practice ID. 
* `name` - (Required) Best Practice Name. 
* `action_item` - (Optional) To comply with Best Practice, do this action item. 
* `description` - (Optional) Description of the Best Practice. 
* `expected_output_text` - (Optional) The expected output of the script as plain text. 
* `expected_output_base64` - (Optional) The expected output of the script as Base64. 
* `practice_script_path` - (Optional) The absolute path of the script on the Management Server to run on Gaia Security Gateways during the Compliance scans. 
* `practice_script_base64` - (Optional) The entire content of the script encoded in Base64 to run on Gaia Security Gateways during the Compliance scans. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `best_practice_id` - ID of the Best Practice.
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