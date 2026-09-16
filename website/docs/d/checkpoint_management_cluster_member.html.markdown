---
layout: "checkpoint"
page_title: "checkpoint_management_cluster_member"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-cluster-member"
description: |-
Use this data source to get information on an existing Check Point Cluster Member.
---

# Data Source: checkpoint_management_cluster_member

Use this data source to get information on an existing Check Point Cluster Member.

## Example Usage


```hcl
data "checkpoint_management_cluster_member" "data_cluster_member" {
  uid = "CLUSTER_MEMBER_UID"
  limit_interfaces = 20
}
```

## Argument Reference

The following arguments are supported:

* `uid` - (Required) Object unique identifier.
* `limit_interfaces` - (Optional) Limit number of cluster member interfaces to show.
* `auto_generate_ip` - Use an automatically generated IP address for the Gateway object (applies only to Smart-1 Cloud).
* `trust_details` - Details for trusted communication.trust_details blocks are documented below.
* `trust_method` - Trust method that was used for establishing communication.


`trust_details` supports the following:

* `authentication_token` - Authentication token to use on the Gateway side to establish the communication between the Gateway and the Management Server (applies only to Smart-1 ...
* `cloud_communication_details` - Details about the communication status with cloud (applies only to Smart-1 Cloud).cloud_communication_details blocks are documented below.
* `gateway_mac_address` - Use the Security Gateway MAC address, relevant for the gateway_mac_address identification-method.
* `identification_method` - How to identify the gateway (relevant for DAIP gateways only).
* `status` - Status of the trusted communication with the Security Gateway.
* `token_expiration_date` - Details about the communication status with cloud (applies only to Smart-1 Cloud).


`cloud_communication_details` supports the following:

* `ip` - IP address used for communication between the Gateway and the Management Server (used when 'auto-generate-ip=true').
* `status` - Status of communication between the Gateway and the Management Server.
