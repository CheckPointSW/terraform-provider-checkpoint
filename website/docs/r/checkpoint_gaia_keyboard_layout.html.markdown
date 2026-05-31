---
layout: "checkpoint"
page_title: "checkpoint_gaia_keyboard_layout"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-keyboard-layout"
description: |-
This resource allows you to execute Check Point Keyboard Layout.
---

# checkpoint_gaia_keyboard_layout

This resource allows you to execute Check Point Keyboard Layout.

## Example Usage


```hcl
resource "checkpoint_gaia_keyboard_layout" "example" {
  keyboard_layout = "us"
}
```

## Argument Reference

The following arguments are supported:

* `keyboard_layout` - (Required) Available languages: be-latin1 - Belgian, bg - Bulgarian, br-abnt2 - Brazilian, cf - Central African Republic, cz-lat2 - Czechoslovakian, de - German, dvorak - Dvorák, dk - Danish, et - Estonian, fi - Finnish, fr - French, fr_CH - Swiss French, sg - Swiss German, hu - Hungarian, is-latin1 - Icelandic, it - Italian, jp106 - Japanese, no - Norwegian, pl - Polish, pt-latin1 - Portuguese, ru - Russian, es - Spanish, se-latin1 - Swedish, trq - Turkish, uk - Great Britain, us - US  
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
