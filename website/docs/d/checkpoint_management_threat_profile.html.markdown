---
layout: "checkpoint"
page_title: "checkpoint_management_threat_profile"
sidebar_current: "docs-checkpoint-data-source-checkpoint-management-threat-profile"
description: |-
  This resource allows you to execute Check Point Threat Profile.
---

# Data Source: checkpoint_management_threat_profile

This resource allows you to execute Check Point Threat Profile.

## Example Usage


```hcl
resource "checkpoint_management_threat_profile" "threat_profile" {
	name = "my threat profile"
	active_protections_performance_impact = "high"
	active_protections_severity	 = "Critical"
}

data "checkpoint_management_threat_profile" "test_threat_profile" {
    name = "${checkpoint_management_threat_profile.threat_profile.name}"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Object name. Should be unique in the domain.
* `active_protections_performance_impact` - (Optional) Protections with this performance impact only will be activated in the profile.
* `active_protections_severity` - (Optional) Protections with this severity only will be activated in the profile.
* `confidence_level_high` - (Optional) Action for protections with high confidence level.
* `confidence_level_medium` - (Optional) Action for protections with medium confidence level.
* `confidence_level_low` - (Optional) Action for protections with low confidence level.
* `indicator_overrides` - (Optional) Indicators whose action will be overridden in this profile. indicator_overrides blocks are documented below.
* `ips_settings` - (Optional) IPS blade settings. ips_settings blocks are documented below.
* `malicious_mail_policy_settings` - (Optional) Malicious Mail Policy for MTA Gateways. malicious_mail_policy_settings blocks are documented below.
* `overrides` - (Optional) Overrides per profile for this protection. overrides blocks are documented below.
* `scan_malicious_links` - (Optional) Scans malicious links (URLs) inside email messages. scan_malicious_links blocks are documented below.
* `use_indicators` - (Optional) Indicates whether the profile should make use of indicators.
* `anti_virus` - (Optional) Is Anti-Virus blade activated.
* `threat_extraction` - (Optional) Is Threat-Extraction blade activated.
* `zero_phishing` - (Optional) Is Zero-Phishing blade activated.
* `anti_bot` - (Optional) Is Anti-Bot blade activated.
* `ips` - (Optional) Is IPS blade activated.
* `threat_emulation` - (Optional) Is Threat Emulation blade activated.
* `use_extended_attributes` - (Optional) Whether to activate/deactivate IPS protections according to the extended attributes.
* `activate_protections_by_extended_attributes` - (Optional) Activate protections by these extended attributes. activate_protections_by_extended_attributes blocks are documented below.
* `deactivate_protections_by_extended_attributes` - (Optional) Deactivate protections by these extended attributes. deactivate_protections_by_extended_attributes blocks are documented below.
* `color` - (Optional) Color of the object. Should be one of existing colors.
* `comments` - (Optional) Comments string.
* `tags` - (Optional) Collection of tag identifiers.
* `ignore_warnings` - (Optional) Apply changes ignoring warnings.
* `ignore_errors` - (Optional) Apply changes ignoring errors. You won't be able to publish such a changes. If ignore-warnings flag was omitted - warnings will also be ignored.

`indicator_overrides` supports the following:

* `action` - (Optional) The indicator's action in this profile.
* `indicator` - (Optional) The indicator whose action is to be overriden.

`ips_settings` supports the following:

* `exclude_protection_with_performance_impact` - (Optional) Whether to exclude protections depending on their level of performance impact.
* `exclude_protection_with_performance_impact_mode` - (Optional) Exclude protections with this level of performance impact.
* `exclude_protection_with_severity` - (Optional) Whether to exclude protections depending on their level of severity.
* `exclude_protection_with_severity_mode` - (Optional) Exclude protections with this level of severity.
* `newly_updated_protections` - (Optional) Activation of newly updated protections.

`malicious_mail_policy_settings` supports the following:

* `add_customized_text_to_email_body` - Add customized text to the malicious email body.
* `add_email_subject_prefix` - Add a prefix to the malicious email subject.
* `add_x_header_to_email` - Add an X-Header to the malicious email.
* `email_action` - Block - block the entire malicious email. Allow - pass the malicious email and apply email changes (like: remove attachments and links, add x-header, etc...).
* `email_body_customized_text` - Customized text for the malicious email body. Available predefined fields: $verdicts$ - the malicious/error attachments/links verdict.
* `email_subject_prefix_text` - Prefix for the malicious email subject.
* `failed_to_scan_attachments_text` - Replace attachments that failed to be scanned with this text. Available predefined fields: $filename$ - the malicious file name. $md5$ - MD5 of the malicious file.
* `malicious_attachments_text` - Replace malicious attachments with this text. Available predefined fields: $filename$ - the malicious file name. $md5$ - MD5 of the malicious file.
* `malicious_links_text` - Replace malicious links with this text. Available predefined fields: $neutralized_url$ - neutralized malicious link.
* `remove_attachments_and_links` - Remove attachments and links from the malicious email.
* `send_copy` - Send a copy of the malicious email to the recipient list.
* `send_copy_list` - Recipient list to send a copy of the malicious email.

`overrides` supports the following:

* `protection` - IPS protection identified by name.
* `action` - Protection action.
* `capture_packets` - Capture packets.
* `track` - Tracking method for protection.
* `default` - Default settings. default blocks are documented below.
* `final` - Final settings. final blocks are documented below.
* `protection_external_info` - Collection of industry reference (CVE).
* `protection_uid` - IPS protection unique identifier.

`scan_malicious_links` supports the following:

* `max_bytes` - Scan links in the first bytes of the mail body.
* `max_links` - Maximum links to scan in mail body.

`activate_protections_by_extended_attributes` supports the following:

* `uid` - IPS tag unique identifier.
* `name` - IPS tag name.
* `category` - IPS tag category name.
* `values` - Collection of IPS protection extended attribute values (name and uid).

`deactivate_protections_by_extended_attributes` supports the following:

* `uid` - IPS tag unique identifier.
* `name` - IPS tag name.
* `category` - IPS tag category name.
* `values` - Collection of IPS protection extended attribute values (name and uid).

`default` supports the following:

* `action` - Protection action.
* `capture_packets` - Capture packets.
* `track` - Tracking method for protection.

`final` supports the following:

* `action` - Protection action.
* `capture_packets` - Capture packets.
* `track` - Tracking method for protection.