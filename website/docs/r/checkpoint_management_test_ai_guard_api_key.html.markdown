---
layout: "checkpoint"
page_title: "checkpoint_management_test_ai_guard_api_key"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-test-ai-guard-api-key"
description: |-
This resource allows you to execute Check Point Test Ai Guard Api Key.
---

# checkpoint_management_test_ai_guard_api_key

This resource allows you to execute Check Point Test Ai Guard Api Key.

## Example Usage


```hcl
resource "checkpoint_management_test_ai_guard_api_key" "example" {
}
```

## Argument Reference

The following arguments are supported:

* `project_id` - (Optional) Optional Lakera project ID to validate. If provided, also verifies the project belongs to the API key.
* `message` - (Computed) Validation result message.
* `success` - (Computed) Whether the API key (and optional project) is valid.