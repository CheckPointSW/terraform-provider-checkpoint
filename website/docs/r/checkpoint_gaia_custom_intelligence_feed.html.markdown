---
layout: "checkpoint"
page_title: "checkpoint_gaia_custom_intelligence_feed"
sidebar_current: "docs-checkpoint-resource-checkpoint-gaia-custom-intelligence-feed"
description: |-
This resource allows you to execute Check Point Custom Intelligence Feed.
---

# checkpoint_gaia_custom_intelligence_feed

This resource allows you to execute Check Point Custom Intelligence Feed.

## Example Usage


```hcl
resource "checkpoint_gaia_custom_intelligence_feed" "example" {
  name     = "tf-example-feed"
  protocol = "http"
  url      = "http://example.com/feeds/iocs.txt"
  action   = "detect"
}
```

## Argument Reference

The following arguments are supported:

* `protocol` - (Required)  
* `url` - (Required) Set the feed URL 
* `name` - (Required)  
* `enabled` - (Optional)  
* `action` - (Optional) Set feed action 
* `account_name` - (Optional)  
* `account_password` - (Optional)  
* `custom_csv_settings` - (Optional) Define custom csv settings - CSV structure, Delimiter and rows to skip custom_csv_settings blocks are documented below.
* `format` - (Optional) STIX: https://stixproject.github.io/. For more info see sk132193 
* `https_sha256_fingerprint` - (Optional) Specify HTTPS SHA-256 fingerprint 
* `member_id` - (Computed) Relevant for commands on Scalable and ElasticXL platforms only. When member-id is provided in the login request, show commands during the session will be executed on the specified member, unless a different member-id is provided in a successive requests Set operations will be performed on all members 


`custom_csv_settings` supports the following:

* `csv_observable_value` - (Optional)  
* `csv_observable_type` - (Optional) Set integer for index location in CSV file, or fixed value for the entire feed. 
* `csv_delimiter` - (Optional)  
* `csv_lines_to_be_skipped` - (Optional)  
* `csv_observable_name` - (Optional) Set integer for index location in CSV file, or fixed value for the entire feed. 
* `csv_observable_description` - (Optional) Set integer for index location in CSV file, or fixed value for the entire feed. 
* `csv_observable_confidence` - (Optional)  
* `csv_observable_severity` - (Optional)  
* `csv_observable_product` - (Optional)  
