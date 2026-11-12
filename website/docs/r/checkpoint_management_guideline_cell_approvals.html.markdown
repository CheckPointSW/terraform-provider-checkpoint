---
layout: "checkpoint"
page_title: "checkpoint_management_guideline_cell_approvals"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-guideline-cell-approvals"
description: |-
This resource allows you to execute Check Point Guideline Cell Approvals.
---

# checkpoint_management_guideline_cell_approvals

This resource allows you to execute Check Point Guideline Cell Approvals.

## Example Usage


```hcl
resource "checkpoint_management_access_rule" "rule1" {
  name  = "test-rule-for-approval"
  layer = "Network"
  position {
    top = "top"
  }
}

resource "checkpoint_management_group" "example" {
  name = "approvalTestGroup"
}

resource "checkpoint_management_guideline" "guideline1" {
  name = "Corporate policy"
  access_layers {
    access_layer = "Network"
  }
  guideline_groups {
    name = checkpoint_management_group.example.name
    position {
      top = "top"
    }
  }
}

resource "checkpoint_management_guideline_cell_approvals" "example" {
  guideline = checkpoint_management_guideline.guideline1.name
  comment   = "This is approved for all segments"
  from      = "any"
  to        = "any"
  approvals {
    rules {
      layer = "Network"
      rule  = checkpoint_management_access_rule.rule1.name
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `guideline` - (Required) The guideline (identified by UID or name) in which we approve the violation.
* `approvals` - (Required) List of approved rules.approvals blocks are documented below.
* `from` - (Optional) "from" segment (identified by UID or name), or 'any' to approved the rule across all cells (possible only if "to" is also 'any'). This field is mandatory if "from-type" is 'Network Group'. 
* `to` - (Optional) "to" segment (identified by UID or name), or 'any' to approved the rule across all cells (possible only if "from" is also 'any'). This field is mandatory if "to-type" is 'Network Group'. 
* `comment` - (Required) Comment on the approvals. The same comment to all the requested approvals.
* `from_type` - (Optional) The type of the segment in the 'from' axis. 
* `to_type` - (Optional) The type of the segment in the 'to' axis. 
* `policy_package` - (Optional) The policy package (identified by UID or name) in which we approve the violation. This field is mandatory only if the ordered-access-layer (first layer in path) is from a global domain with AGP. 
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
* `delete_scope` - (Optional) Indicates whether to delete all the approval scope, or only remove the requested cell from the scope. Relevant only for guideline approvals. This field is only relevant on delete - it is sent when the resource is destroyed and has no effect on create or update. To change its value, update it in the configuration and run `terraform apply` (so the value is stored in state) before destroying.


`approvals` supports the following:

* `rules` - (Optional) The full paths (pairs of layer and rule) of the approved rules.rules blocks are documented below.


`rules` supports the following:

* `layer` - (Optional) The Layer identifier (name or UID). 
* `rule` - (Optional) The rule identifier (name if unique, rule position number in rule-base or UID). 
