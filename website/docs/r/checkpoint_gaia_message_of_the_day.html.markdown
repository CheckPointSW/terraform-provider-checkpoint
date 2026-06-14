---
layout: "checkpoint"
page_title: "checkpoint_gaia_message_of_the_day"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-message-of-the-day"
description: |-
This resource allows you to execute Check Point Message Of The Day.
---

# checkpoint_gaia_message_of_the_day

This resource allows you to execute Check Point Message Of The Day.

## Example Usage


```hcl
resource "checkpoint_gaia_message_of_the_day" "example" {
  enabled = true
  message = "Hello today"
}
```

## Argument Reference

The following arguments are supported:

* `message` - (Optional) Message of the day for web, ssh and serial login. Empty string returns to default 
* `enabled` - (Optional) Message of the day enabled (true/false) 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 
