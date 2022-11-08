---
layout: "checkpoint"
page_title: "checkpoint_management_azure_ad_content"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-azure-ad-content"
description: |-
This resource allows you to execute Check Point Azure Ad Content.
---

# Data Source: checkpoint_management_azure_ad_content

This resource allows you to execute Check Point Azure Ad Content.

## Example Usage


```hcl
data "checkpoint_management_azure_ad_content" "azure_ad_content" {
    azure_ad_name = "my_azureAD"
}
```

## Argument Reference

The following arguments are supported:

* `azure_ad_name` - (Optional) Name of the Azure AD Server where to search for objects.
* `azure_ad_uid` - (Optional) Unique identifier of the Azure AD Server where to search for objects.
* `limit` - (Optional) The maximal number of returned results.
* `offset` - (Optional) Number of the results to initially skip.
* `order` - (Optional) Sorts the results by search criteria. Automatically sorts the results by Name, in the ascending order. order blocks are documented below.
* `uid_in_azure_ad` - (Optional) Return result matching the unique identifier of the object on the Azure AD Server.
* `filter` - (Optional) Return results matching the specified filter. filter blocks are documented below.


`order` supports the following:

* `asc` - (Optional) Sorts results by the given field in ascending order.
* `desc` - (Optional) Sorts results by the given field in descending order.


`filter` supports the following:

* `text` - (Optional) Return results containing the specified text value.
* `uri` - (Optional) Return results under the specified Data Center Object (identified by URI).
* `parent_uid_in_data_center` - (Optional) Return results under the specified Data Center Object (identified by UID).

Output:

* `from` - From which element number the query was done.
* `objects` - Remote objects views. objects blocks are documented below.
* `to` - To which element number the query was done.
* `total` - Total number of elements returned by the query.


`objects` supports the following:

* `name_in_azure_ad` - Object name in the Azure AD.
* `uid_in_azure_ad` - Unique identifier of the object in the Azure AD.
* `azure_ad_object` - The imported management object (if exists). Level of details in the output corresponds to the number of details for search. This table shows the level of details in the Standard level.
* `name` - Object management name.
* `type_in_azure_ad` - Object type in Azure AD.
* `additional_properties` - Additional properties on the object. additional_properties blocks are documented below.


`additional_properties` supports the following:

* `name`
* `value`