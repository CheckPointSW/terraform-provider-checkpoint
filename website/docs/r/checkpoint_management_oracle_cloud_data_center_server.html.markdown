---
layout: "checkpoint"
page_title: "checkpoint_management_oracle_cloud_data_center_server"
sidebar_current: "docs-checkpoint-Resource-checkpoint-management-oracle-cloud-data-center-server"
description: |- This resource allows you to execute Check Point oracle cloud data center server.
---

# Resource: checkpoint_management_oracle_cloud_data_center_server

This resource allows you to execute Check Point Oracle Cloud Data Center Server.

## Example Usage

```hcl
resource "checkpoint_management_oracle_cloud_data_center_server" "testOracleCloud" {
  name = "MY-ORACLE-CLOUD"
  authentication_method = "key-authentication"
  private_key = "0SLtLS1CRUdJTiBQUklWQVRFIEtS0FWtLS0tDQpNSUlFdkFJQkFEQU5CZ2txaGtpRzl3AAAAUUVGQUFTQ0JLWXdnZ1NpQWdFQUFvSUJBUURUdmVrK1laMmNSekVmDQp1QkNoMkFxS2hzUFcrQUhUajY4dE5VbVl4OUFTRXBsREhnMkF0bCtMRWRRWUFRSUtLMUZ5L1JHRitkK3RkWjUrDQpabmprN0hESTQ5V3Rib0xodWN3YjBpNU4xbEVKWHVhOHhEN0FROTJXQy9PdzhzVktPRlJGNVJhMmxSa0svRS8xDQpxeDhKYnRoMGdXdHg0NHBQaWJwU3crMTB0QUhHR2FTLzVwN3hNUXhzajZTOThwL1hnalg5NzN4VStZZ2dLNUx3DQp6WlkzSDQ3UVREcmpyZzhOVmpDSFU3b3IrcEpCbjdldGF0V3psK3BQcVd4ODZub2tjdG5abUQxcHNnWnkwTEdDDQpRYys5ejdURGhEOFhuVERwckxiRGZXRnZqOTVKSmc3Q1krd29zN05vSENEOG5RWjFZZURVQkJjUkVlZXJVRlhBDQpaZ1I3UGNCN0FnTUJBQUVDZ2dFQUdkUWxCZVFER3ROMXJaTXNERGRUU2RudDU3b2NWdXlac2grNW1PRVVMakF3DQptOXhTOUt3ZnlYeTZRaXlxb3JKQ3REa2trR2c0OHhWOFBrVDN1RTBUTzU0aDF1UmxJMjNMbjcvVmFzOUZnVlFmDQpQS1dLVmdwYjdFMWhtT2gwVFNmRDRwRnpETlh4SzhMaXYycWVxdTJTTlRGWVR1M2RBRWpNL3EyWERmdXJQN2tiDQprZ3FKRFBwd2g4RWRXMVg1VVAyVE9CVWxwQllDTndxUkFJQ1E3eWlzbW5xeFlZS3RKc21MK21IQ3JYM3hNRHVTDQp4NHJCVDUvcXVrdVc4MmwrbGZmU3ZTNGpsb0VhajJ2QmozSk1udy9lYlNucFplU3FENTFjOUZlOCtocU4rU3NoDQozTnc0QXVybE1RRG5vZy9STUF3QUR3KzBRUlIwNVdaWDhMVXllVTBVVVFLQmdRRHd6R2I0b25BWHlNeUNkbHB3DQpRRnFCR0pwQnlsWEJjSkZnMGpCd1JMV2tQL3VjWnNoZlFlbkFWbkRZZS9yQ0FnWWxSdFFOVFRvb3BFSjlGcGgyDQp6TkVzd1EwcnV4WjFrVm41U1hwS2dF4668KalUxT3dGa3R1WFlJcEtBNGk5dFoxT04zb1lqdVRtMUlzb2xWZXVTDQpqK3Mwd1o3ZDAyYTNXcDN1UXJ3TFUwVjdpUUtCZ1FEaEcrc2xsNDYveGxONldWWEs3RVpkOGFEcTlTNEU0aEQvDQpvTmUwS0dVcHhZYngyTnFWN1VLSEwzOE41eG5qNGllWGt2U1BnL0twVUpqUmtLN0xJMnZsNmlndUJkdW01VUR1DQp5dW4rL1dNcVdnb2p4anZBbmxsS2lIa0JRMTJ2UFRqcE9HSGIrY0RqVWxROGVnOThFOEJ0ZktUQjFkRlcxUnBlDQorMXY0aXR3RzR3S0JnQzJLeXpMZExnd2hpeVJsbEFkRTlKa1QrU0RXVHMvT0pZREZZQ25ycE5zU3l0aXl5OVRRDQpWNUJzQ04yNDNSMVNXcTAwTHlqdzRUNE1peEt6Y2xTTnVrWVhvUkVUU2xVa0QzdEpmVnFYMVUrTE1XY0c2T1dPDQpmZndaMWRHUWRkM2dPL3BLQ3Q2NHlvUkt0eWJHa0U1ZzcrQkRlbk9ENXhwb2hoUXBCUDJ6V3lIWkFvR0FURndqDQpGUHBuUXVoc3Nza1JFQ2U3NnV3bkVPeWdjcW1ZNkkzUC9kM2lDeHhsSFM3Wlh4Zy9oQW41aUdiSFlvVDV0ekh6DQpZYWQ1cmpPWDB5YklGRUpzdkc0RXVTL2xoYVNvdFJnQjdpeFg4aXJlMjZuSDVSd1IzL1dSVG50aWtTb3NYdmh3DQpRYVZqNS9pcWVHVlRVVnlGM3QzMEtZaDFYWVltVHVmbkY5VktzODhDZ1lCTTNVN2QwOU9MemhMZTh3cnp1dEpuDQpGdmRGWlhCRnhXRGwyNXdteElWTFdpM0kvaWg2QXN5YlRNNWMzbFpTTUVvcjJYeXJqNnFUNzZ6amQ2eGE2NlN3DQpXMEVyL2lEY3dWK244MHpuU3lPSW5lRThIVkh1SGtNYVpPeHkvVzdVWDFqL0RmUnJPZG1iS1NWN2NBV2dVTlBrDQpnd1V5RkM2OTRKTR41Vko0WXZEZU13PT0NCi0tLS0tRU5EIFBSSVZBVEUgS0VZLS0tLS0="
  key_user = "ocid1.user.oc1..aaaaaaaad6n7rniiwgxehy6coo4ax2ti7pr5yr53cbdxdyp6sx6dhrttcz4a"
  key_tenant = "ocid1.tenancy.oc1..aaaaaaaaft6hqvl367uh4e3pmdxnzmca6cxamwjfaag5lm7bnhuwu6ypajca"
  key_region = "eu-frankfurt-1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (**Required**) Object name. Must be unique in the domain.
* `authentication_method` - (**Required**) key-authentication uses the Service Account private key file to authenticate. vm-instance-authentication uses VM Instance to authenticate. This option requires the Security Management Server deployed in Oracle Cloud, and running in a dynamic group with the required permissions. 
* `private_key` - (**Required**) An Oracle Cloud API key PEM file, encoded in base64. Required for authentication-method: key-authentication.
* `key_user` - (**Required**) An Oracle Cloud user id associated with key. Required for authentication-method: key-authentication.
* `key_tenant` - (**Required**) An Oracle Cloud tenancy id where the key was created. Required for authentication-method: key-authentication.
* `key_region` - (**Required**) An Oracle Cloud region for where to create scanner. Required for authentication-method: key-authentication.
* `tags` - (Optional) Collection of tag identifiers.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.
* `automatic_refresh` - Indicates whether the data center server's content is automatically updated.
* `data_center_type` - Data center type.
* `properties` - Data center properties. properties blocks are documented below.


`properties` supports the following:

* `name`
* `value`