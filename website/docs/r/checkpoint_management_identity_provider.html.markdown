---
layout: "checkpoint"
page_title: "checkpoint_management_identity_provider"
sidebar_current: "docs-checkpoint-resource-checkpoint-management-identity-provider"
description: |-
This resource allows you to execute Check Point Identity Provider.
---

# checkpoint_management_identity_provider

This resource allows you to execute Check Point Identity Provider.

## Example Usage


```hcl
resource "checkpoint_management_identity_provider" "example" {
  name = "TestIdp2"
  usage = "managing_administrator_access"
  data_receiving = "manually"
  received_identifier = "https://sts.checkpoint.net/621ac12d-4afb-479c-9c14-13e7b743cd07/"
  login_url = "https://login.checkpointonline.com/621ac12d-4afb-479c-9c14-13e7b743cd07/saml2"
  base64_certificate = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM4RENDQWRpZ0F3SUJBZ0lRUTBWZVpLdVBLb0pQUWhaNGhDaWRzREFOQmdrcWhraUc5dzBCQVFzRkFEQTBNVEl3TUFZRFZRUURFeWxOYVdOeWIzTnZablFnUVhwMWNtVWdSbVZrWlhKaGRHVmtJRk5UVHlCRFpYSjBhV1pwWTJGMFpUQWVGdzB4T0RBME1UVXhNVEl6TXpOYUZ3MHlNVEEwTVRVeE1USXpNek5hTURReE1qQXdCZ05WQkFNVEtVMXBZM0p2YzI5bWRDQkJlblZ5WlNCR1pXUmxjbUYwWldRZ1UxTlBJRU5sY25ScFptbGpZWFJsTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUE0VXVqYUd0OFhaODl2dXZ5a3lRVzVYb24vOFIvaVB0ejRhYjBNM3RNVXZHWHozVXh0V1pTRStUR1hydjN3VHRLMCs4RmtNeXVKYUhGSXBLLzRVREZpRk1yQmxzR0Z1dmtTc1p5VjIzZlNGN3paaXlUWTZUN0EwcCtnczUwNVhEOUdBYjlWYmR3R0cwK0tDVnlpc1ZRZ1YySXdKZ2l5aHF3RUNvY3dCcmFuN251SytURU5EMmwyZjlZcng1b1JNRU56NzB3bHlIMzZPWkJtdDBrNTk4L1doMEhEWUxaZW8wZHlTV3JOd3dlWXZTeEU4L01kbTQzWEV1U3pialR6ZnNNMHZVUndGQlNyVUxFYURPMS9JUDJVcjdCc2dId1JJL3hmb3FJbUsxS2twVXEwQWxjVEFpM3YxdTl6Qy9xTGdQd0F5UUl2dzlVQ3NpcnJQQTBZMFlPaFFJREFRQUJNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUURjam9qZEd6L0FJQ2pqTTBaN21ZbGdQNXpic2FRNWRDMmNqZjRESnFta21zV3VmUnBDNHNic3VoODcwY0NCS2N1dmgrb0dpekJRSHJQbTRUaEl2ZklsS0w4WGpMQVhiRnVSUG9IQWcwOHNMWGR2UFRCVE52REYxTWhvcU5zMmo2ZUZxL2ROdXF2ZUJIcjVENXRLblYyWEJHRUhFOVJFOVdUa1pRT2MwaEhtQ3dNbWNZb3JYRzhCa3l1OXFwNXhyMDZMQ0htMnJLcnI2ZENRVldBV0R0MzhiS2t5STRobTVSNTVCclR5UldSdzI1RS9YaFEwVksva1FJYW9GZ0hvaWo0ekg5bmxlZnZMbmhaZDVPRzROL29OS2pBKy9LbkVqaTdPQXhKWVNaR1FmRjU0R1AwQTE4VnF1NVVGaFBKMUZFQXZ5YjR0QnZtTzM1NFFVUys5UTY2agotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0t"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. 
* `usage` - (Optional) Usage of Identity Provider. 
* `gateway` - (Optional) Gateway for the SAML Identity Provider usage.
Identified by name or UID. <font color="red">Required only when</font> 'usage' is set to 'gateway_policy_and_logs'. 
* `service` - (Optional) Service for the selected gateway. <font color="red">Required only when</font> 'usage' is set to 'gateway_policy_and_logs'. 
* `data_receiving` - (Optional) Data receiving method from the SAML Identity Provider. 
* `received_identifier` - (Optional) Received Identifier (Entity ID) based on the provider data. <font color="red">Required only when</font> 'data-receiving' is set to 'manually'. 
* `login_url` - (Optional) Login URL based on the provider data. <font color="red">Required only when</font> 'data-receiving' is set to 'manually'. 
* `base64_metadata_file` - (Optional) Metadata file encoded in base64 based on the provider data. <font color="red">Required only when</font> 'data-receiving' is set to 'metadata_file'. 
* `base64_certificate` - (Optional) Certificate file encoded in base64 based on provider data. <font color="red">Required only when</font> 'data-receiving' is set to 'manually'. 
* `color` - (Optional) Color of the object. Should be one of existing colors. 
* `comments` - (Optional) Comments string. 
* `tags` - (Optional) Collection of tag identifiers.tags blocks are documented below.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings. 
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored. 
