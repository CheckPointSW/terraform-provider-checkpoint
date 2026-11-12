---
layout: "checkpoint"
page_title: "checkpoint_management_identity_provider"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-identity-provider"
description: |- Use this data source to get information on an existing Identity Provider.
---


# checkpoint_management_identity_provider

Use this data source to get information on an existing Identity Provider.

## Example Usage


```hcl
resource "checkpoint_management_identity_provider" "example" {
  name = "%s"
  usage = "managing_administrator_access"
  data_receiving = "manually"
  received_identifier = "https://sts.checkpoint.net/621ac12d-4afb-479c-9c14-13e7b743cd07/"
  login_url = "https://login.checkpointonline.com/621ac12d-4afb-479c-9c14-13e7b743cd07/saml2"
  base64_certificate = "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUM4RENDQWRpZ0F3SUJBZ0lRUTBWZVpLdVBLb0pQUWhaNGhDaWRzREFOQmdrcWhraUc5dzBCQVFzRkFEQTBNVEl3TUFZRFZRUURFeWxOYVdOeWIzTnZablFnUVhwMWNtVWdSbVZrWlhKaGRHVmtJRk5UVHlCRFpYSjBhV1pwWTJGMFpUQWVGdzB4T0RBME1UVXhNVEl6TXpOYUZ3MHlNVEEwTVRVeE1USXpNek5hTURReE1qQXdCZ05WQkFNVEtVMXBZM0p2YzI5bWRDQkJlblZ5WlNCR1pXUmxjbUYwWldRZ1UxTlBJRU5sY25ScFptbGpZWFJsTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUE0VXVqYUd0OFhaODl2dXZ5a3lRVzVYb24vOFIvaVB0ejRhYjBNM3RNVXZHWHozVXh0V1pTRStUR1hydjN3VHRLMCs4RmtNeXVKYUhGSXBLLzRVREZpRk1yQmxzR0Z1dmtTc1p5VjIzZlNGN3paaXlUWTZUN0EwcCtnczUwNVhEOUdBYjlWYmR3R0cwK0tDVnlpc1ZRZ1YySXdKZ2l5aHF3RUNvY3dCcmFuN251SytURU5EMmwyZjlZcng1b1JNRU56NzB3bHlIMzZPWkJtdDBrNTk4L1doMEhEWUxaZW8wZHlTV3JOd3dlWXZTeEU4L01kbTQzWEV1U3pialR6ZnNNMHZVUndGQlNyVUxFYURPMS9JUDJVcjdCc2dId1JJL3hmb3FJbUsxS2twVXEwQWxjVEFpM3YxdTl6Qy9xTGdQd0F5UUl2dzlVQ3NpcnJQQTBZMFlPaFFJREFRQUJNQTBHQ1NxR1NJYjNEUUVCQ3dVQUE0SUJBUURjam9qZEd6L0FJQ2pqTTBaN21ZbGdQNXpic2FRNWRDMmNqZjRESnFta21zV3VmUnBDNHNic3VoODcwY0NCS2N1dmgrb0dpekJRSHJQbTRUaEl2ZklsS0w4WGpMQVhiRnVSUG9IQWcwOHNMWGR2UFRCVE52REYxTWhvcU5zMmo2ZUZxL2ROdXF2ZUJIcjVENXRLblYyWEJHRUhFOVJFOVdUa1pRT2MwaEhtQ3dNbWNZb3JYRzhCa3l1OXFwNXhyMDZMQ0htMnJLcnI2ZENRVldBV0R0MzhiS2t5STRobTVSNTVCclR5UldSdzI1RS9YaFEwVksva1FJYW9GZ0hvaWo0ekg5bmxlZnZMbmhaZDVPRzROL29OS2pBKy9LbkVqaTdPQXhKWVNaR1FmRjU0R1AwQTE4VnF1NVVGaFBKMUZFQXZ5YjR0QnZtTzM1NFFVUys5UTY2agotLS0tLUVORCBDRVJUSUZJQ0FURS0tLS0t"
}

data "checkpoint_management_identity_provider" "data_identity_provider" {
  name = "${checkpoint_management_identity_provider.example.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Object name.
* `uid` - (Optional) Object uid.
* `usage` - Usage of Identity Provider.
* `gateway` - Gateway for the SAML Identity Provider usage. Identified by name or UID.
* `service` - Service for the selected gateway.
* `data_receiving` - Data receiving method from the SAML Identity Provider.
* `received_identifier` - Received Identifier (Entity ID) based on the provider data.
* `login_url` - Login URL based on the provider data.
* `base64_metadata_file` - Metadata file encoded in base64 based on the provider data.
* `base64_certificate` - Certificate file encoded in base64 based on provider data.
* `color` - Color of the object. Should be one of existing colors.
* `comments` - Comments string.
* `tags` - Collection of tag identifiers.tags blocks are documented below.
