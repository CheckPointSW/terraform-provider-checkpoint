## 1.9.1 (June 16, 2022)

BUG FIXES

* Fix certificate error (Use GO SDK v1.5.1).

## 1.9.0 (June 14, 2022)

FEATURES

* **New Resource:** `resource_checkpoint_management_get_platform`
* **New Resource:** `resource_checkpoint_management_reset_sic`
* **New Resource:** `resource_checkpoint_management_test_sic_status`
* **New Resource:** `resource_checkpoint_management_set_idp_default_assignment`
* **New Resource:** `resource_checkpoint_management_set_idp_to_domain_assignment`
* **New Resource:** `resource_checkpoint_management_interoperable_device`
* **New Resource:** `resource_checkpoint_management_install_lsm_police`
* **New Resource:** `resource_checkpoint_management_install_lsm_settings`
* **New Resource:** `resource_checkpoint_management_lsm_run_script`
* **New Resource:** `resource_checkpoint_management_update_provisioned_satellites`
* **New Resource:** `resource_checkpoint_management_repository_script`
* **New Resource:** `resource_checkpoint_management_smtp_server`
* **New Resource:** `resource_checkpoint_management_check_threat_ioc_feed`
* **New Resource:** `resource_checkpoint_management_domain_permissions_profile`
* **New Resource:** `resource_checkpoint_management_idp_administrator_group`
* **New Resource:** `resource_checkpoint_management_md_permissions_profile`
* **New Resource:** `resource_checkpoint_management_network_feed`
* **New Resource:** `resource_checkpoint_management_check_network_feed`
* **New Resource:** `resource_checkpoint_management_connect_cloud_services`
* **New Resource:** `resource_checkpoint_management_disconnect_cloud_services`
* **New Data Source:** `data_source_checkpoint_management_cluster_member`
* **New Data Source:** `data_source_checkpoint_management_domain_permissions_profile`
* **New Data Source:** `data_source_checkpoint_management_idp_default_assignment`
* **New Data Source:** `data_source_checkpoint_management_lsm_cluster_profile`
* **New Data Source:** `data_source_checkpoint_management_lsm_gateway_profile`
* **New Data Source:** `data_source_checkpoint_management_provisioning_profile`
* **New Data Source:** `data_source_checkpoint_management_interoperable_device`
* **New Data Source:** `data_source_checkpoint_management_repository_script`
* **New Data Source:** `data_source_checkpoint_management_smtp_server`
* **New Data Source:** `data_source_checkpoint_management_idp_administrator_group`
* **New Data Source:** `data_source_checkpoint_management_md_permissions_profile`
* **New Data Source:** `data_source_checkpoint_management_network_feed`
* **New Data Source:** `data_source_checkpoint_management_cloud_services`

ENHANCEMENTS

* Add `approve_session`, `submit_session` and `reject_session` post apply scripts.
* Add `session_name` field to provider to specify login session name.
* Add `granular_encryptions` and `tunnel_granularity` fields to VPN resources and data sources.

## 1.8.0 (May 22, 2022)

FEATURES

* **New Resource:** `resource_checkpoint_management_threat_ioc_feed`
* **New Resource:** `resource_checkpoint_management_domain`
* **New Resource:** `resource_checkpoint_management_add_repository_package`
* **New Resource:** `resource_checkpoint_management_delete_repository_package`
* **New Resource:** `resource_checkpoint_management_time`
* **New Resource:** `resource_checkpoint_management_trusted_client`
* **New Data Source:** `data_source_checkpoint_management_threat_ioc_feed`
* **New Data Source:** `data_source_checkpoint_management_domain`
* **New Data Source:** `data_source_checkpoint_management_repository_package`
* **New Data Source:** `data_source_checkpoint_management_time`
* **New Data Source:** `data_source_checkpoint_management_trusted_client`

ENHANCEMENTS

* `data_source_checkpoint_management_simple_gateway` - Add support to `application_control_and_url_filtering_settings` field.
* `resource_checkpoint_management_simple_gateway` - Add support to `application_control_and_url_filtering_settings` field.

BUG FIXES

* `resource_checkpoint_management_checkpoint_host` - Fix bug that the `logs_settings` field forced user to put default values as input. Made name field optional and not required.
* `data_source_checkpoint_management_checkpoint_host` - Fix bug that the `logs_settings` field forced user to put default values as input. Made name field optional and not required.
* `resource_checkpoint_management_aws_data_center_server` - Fix bug that the `enable_sts_assume_role` field treated as string instead of bool.

## 1.7.0 (February 24, 2022)

ENHANCEMENTS

* `commands/discard/discard.go`: Add support for discard post apply script.

BUG FIXES

* updated `go.sum` to fix usage of packages with security vulnerabilities.
* `resource_checkpoint_management_access_rule` - Fix bug that the `track` field forced user to put default values as input. Made name field optional and not required.
* `data_source_checkpoint_management_access_rule` - Fix bug that the `track` field forced user to put default values as input. Made name field optional and not required.
* `resource_checkpoint_management_simple_gateway` - Fix bug that the `logs_settings` field forced user to put default values as input. Made name field optional and not required.
* `data_source_checkpoint_management_simple_gateway` - Fix bug that the `logs_settings` field forced user to put default values as input. Made name field optional and not required.

## 1.6.0 (November 24, 2021)


FEATURES

* **New Resource:** `resource_checkpoint_management_aws_data_center_server`
* **New Resource:** `resource_checkpoint_management_azure_data_center_server`
* **New Resource:** `resource_checkpoint_management_gcp_data_center_server`
* **New Resource:** `resource_checkpoint_management_vmware_data_center_server`
* **New Resource:** `resource_checkpoint_management_aci_data_center_server`
* **New Resource:** `resource_checkpoint_management_ise_data_center_server`
* **New Resource:** `resource_checkpoint_management_nuage_data_center_server`
* **New Resource:** `resource_checkpoint_management_openstack_data_center_server`
* **New Resource:** `resource_checkpoint_management_kubernetes_data_center_server`
* **New Resource:** `resource_checkpoint_management_data_center_query`
* **New Data Source:** `data_source_checkpoint_management_aws_data_center_server`
* **New Data Source:** `data_source_checkpoint_management_azure_data_center_server`
* **New Data Source:** `data_source_checkpoint_management_gcp_data_center_server`
* **New Data Source:** `data_source_checkpoint_management_vmware_data_center_server`
* **New Data Source:** `data_source_checkpoint_management_aci_data_center_server`
* **New Data Source:** `data_source_checkpoint_management_ise_data_center_server`
* **New Data Source:** `data_source_checkpoint_management_nuage_data_center_server`
* **New Data Source:** `data_source_checkpoint_management_openstack_data_center_server`
* **New Data Source:** `data_source_checkpoint_management_kubernetes_data_center_server`
* **New Data Source:** `data_source_checkpoint_management_data_center_query`
* **New Data Source:** `data_source_checkpoint_management_data_center_content`
* **New Data Source:** `data_source_checkpoint_management_access_rulebase`
* **New Data Source:** `data_source_checkpoint_management_threat_rulebase`
* **New Data Source:** `data_source_checkpoint_management_https_rulebase`

ENHANCEMENTS

* Add support to authenticate management server with api key using `api_key` or via environment variable `CHECKPOINT_API_KEY`.
* Add support to select proxy host using `proxy_host` or via environment variable `CHECKPOINT_PROXY_HOST`.
* Add support to select proxy port using `proxy_port` or via environment variable `CHECKPOINT_PROXY_PORT`.

BUG FIXES

* `resource_checkpoint_management_simple_cluster` - Fix bug that the `members` field did not import properly.
* `data_source_checkpoint_management_nat_rulebase` - Save all relevant fields in read function.

## 1.5.0 (October 28, 2021)

FEATURES

* **New Resource:** `checkpoint_management_generic_data_center_server`
* **New Data Source:** `checkpoint_management_generic_data_center_server`

ENHANCEMENTS

* `commands/logout/logout.go`: Add support for logout post apply script.
* `resource_checkpoint_management_access_role`: Increase timeout on access-role resource.
* Add support to select session file name via environment variable `CHECKPOINT_SESSION_FILE_NAME`, default value remains `sid.json`.

BUG FIXES

* `resource_checkpoint_management_access_role`: Fix bug that caused when updating `comment` field.
* `resource_checkpoint_management_threat_rule`: Fix general bug in read function.
* `resource_checkpoint_management_access_rule`: Fix bug in `track` field.
* `resource_checkpoint_management_access_section`: Add `layer` field in update and read functions.
* `resource_checkpoint_management_access_layer`: Remove `detect_using_x_forward_for` field default value from schema.
* `resource_checkpoint_management_service_tcp`: Save `color` field into state.
* `resource_checkpoint_management_service_udp`: Save `color` field into state.

## 1.4.0 (March 22, 2021)

FEATURES

* **New Resource:** `checkpoint_management_threat_profile`
* **New Data Source:** `checkpoint_management_threat_profile`

ENHANCEMENTS

* `checkpoint_management_simple_gateway`: Add support for default logs settings.

BUG FIXES

* `checkpoint_management_access_rule`: Add inline layer to payload if action field has changed.
* `checkpoint_management_simple_cluster`: Change members field to type list.

## 1.3.0 (January 12, 2021)

FEATURES

* **New Resource:** `checkpoint_management_simple_gateway`
* **New Resource:** `checkpoint_management_simple_cluster`
* **New Data Source:** `checkpoint_management_simple_gateway`
* **New Data Source:** `checkpoint_management_simple_cluster`

ENHANCEMENTS

* `checkpoint_management_access_section`: Add support for position below specific section or rule.
* `checkpoint_management_access_layer`: Add `add_default_rule` flag indicates whether to include a cleanup rule in the
  new layer.

BUG FIXES

* `checkpoint_management_nat_rule`: Fix call to wrong read function after update resource.

## 1.2.0 (December 17, 2020)

FEATURES

* **New Resource:** `checkpoint_management_nat_rule`
* **New Resource:** `checkpoint_management_nat_section`
* **New Resource:** `checkpoint_management_threat_exception`
* **New Resource:** `checkpoint_management_threat_rule`
* **New Data Source:** `checkpoint_management_nat_rule`
* **New Data Source:** `checkpoint_management_nat_section`
* **New Data Source:** `checkpoint_management_threat_exception`
* **New Data Source:** `checkpoint_management_threat_rule`
* **New Data Source:** `checkpoint_management_show_objects`
* **New Data Source:** `checkpoint_management_show_updatable_objects_repository_content`

ENHANCEMENTS

* Add `triggers` field to resource `checkpoint_management_install_policy`, `checkpoint_management_publish`
  and `checkpoint_management_logout` for re-execution if there are any changes in this list.
* Print publish / install-policy script output to console include task-id.
* Print error message if API server needs to be configured to accept requests from all IP addresses.

BUG FIXES

* `checkpoint_management_access_rule`: Use object UID in update call instead of name.

## 1.1.0 (October 1, 2020)

FEATURES

* **New Resource:** `checkpoint_management_access_point_name`
* **New Resource:** `checkpoint_management_checkpoint_host`
* **New Resource:** `checkpoint_management_gsn_handover_group`
* **New Resource:** `checkpoint_management_identity_tag`
* **New Resource:** `checkpoint_management_mds`
* **New Resource:** `checkpoint_management_service_citrix_tcp`
* **New Resource:** `checkpoint_management_service_compound_tcp`
* **New Resource:** `checkpoint_management_user`
* **New Resource:** `checkpoint_management_user_group`
* **New Resource:** `checkpoint_management_user_template`
* **New Resource:** `checkpoint_management_vpn_community_remote_access`
* **New Resource:** `checkpoint_management_ha_full_sync`
* **New Resource:** `checkpoint_management_set_automatic_purge`
* **New Resource:** `checkpoint_management_set_ha_state`
* **New Resource:** `checkpoint_management_get_attachment`
* **New Data Source:** `checkpoint_management_access_point_name`
* **New Data Source:** `checkpoint_management_checkpoint_host`
* **New Data Source:** `checkpoint_management_mds`
* **New Data Source:** `checkpoint_management_gsn_handover_group`
* **New Data Source:** `checkpoint_management_identity_tag`
* **New Data Source:** `checkpoint_management_service_citrix_tcp`
* **New Data Source:** `checkpoint_management_service_compound_tcp`
* **New Data Source:** `checkpoint_management_user`
* **New Data Source:** `checkpoint_management_user_group`
* **New Data Source:** `checkpoint_management_user_template`
* **New Data Source:** `checkpoint_management_vpn_community_remote_access`

ENHANCEMENTS

* Resources of type command that returns asynchronous task-id(s), will save task-id(s) in state.

BUG FIXES

* Resources of type command are execute as part of 'add' method and are one-use only.

## 1.0.5 (September 9, 2020)

FEATURES

* **New Resource:** `checkpoint_management_put_file`

ENHANCEMENTS

* Resource `checkpoint_management_access_rule`: Add rule in position relative to specific section
* Print login error message to console and exit

BUG FIXES

* Fix resource `checkpoint_management_access_role`
* Fix import access rule. Use the following UID format: <LAYER_NAME>;<RULE_UID>

## 1.0.4 (September 3, 2020)

* Release for Terraform Registry

## 1.0.3 (July 21, 2020)

FEATURES

* **New Data Source:** `checkpoint_management_data_wildcard`
* **New Data Source:** `checkpoint_management_data_security_zone`
* **New Data Source:** `checkpoint_management_data_time_group`
* **New Data Source:** `checkpoint_management_data_group`
* **New Data Source:** `checkpoint_management_data_exception_group`
* **New Data Source:** `checkpoint_management_data_group_with_exclusion`
* **New Data Source:** `checkpoint_management_data_dynamic_object`
* **New Data Source:** `checkpoint_management_data_dns_domain`
* **New Data Source:** `checkpoint_management_data_opsec_application`
* **New Data Source:** `checkpoint_management_data_service_icmp`
* **New Data Source:** `checkpoint_management_data_service_icmp6`
* **New Data Source:** `checkpoint_management_data_service_sctp`
* **New Data Source:** `checkpoint_management_data_service_other`
* **New Data Source:** `checkpoint_management_data_service_group`
* **New Data Source:** `checkpoint_management_data_service_tcp`
* **New Data Source:** `checkpoint_management_data_service_udp`
* **New Data Source:** `checkpoint_management_data_service_dce_rpc`
* **New Data Source:** `checkpoint_management_data_service_rpc`
* **New Data Source:** `checkpoint_management_data_application_site`
* **New Data Source:** `checkpoint_management_data_application_site_category`
* **New Data Source:** `checkpoint_management_data_application_site_group`
* **New Data Source:** `checkpoint_management_data_access_section`
* **New Data Source:** `checkpoint_management_data_access_role`
* **New Data Source:** `checkpoint_management_data_access_layer`
* **New Data Source:** `checkpoint_management_data_access_rule`
* **New Data Source:** `checkpoint_management_data_package`
* **New Data Source:** `checkpoint_management_data_vpn_community_meshed`
* **New Data Source:** `checkpoint_management_data_vpn_community_star`
* **New Data Source:** `checkpoint_management_data_https_rule`
* **New Data Source:** `checkpoint_management_data_https_section`
* **New Data Source:** `checkpoint_management_data_https_layer`
* **New Data Source:** `checkpoint_management_data_network`
* **New Data Source:** `checkpoint_management_data_host`
* **New Data Source:** `checkpoint_management_data_address_range`
* **New Data Source:** `checkpoint_management_data_multicast_address_range`
* **New Data Source:** `checkpoint_management_data_threat_indicator`

ENHANCEMENTS

* Use port and timeout via environment variable in publish and install-policy script
* Save publish and install-policy scripts output to dedicated log file
* Add support for import resources

BUG FIXES

* Fix groups circular dependency
* Fix internal test of few resources

## 1.0.2 (May 13, 2020)

FEATURES:

* Add support to configure timeout and port of provider
* Add support to user agent
* Fix resources: `checkpoint_management_application_site`, `checkpoint_management_application_site_group`
  , `checkpoint_management_https_layer`, `checkpoint_management_service_sctp` and `checkpoint_management_service_other`

## 1.0.1 (March 17, 2020)

FEATURES:

* **New Resource:** `checkpoint_management_wildcard`
* **New Resource:** `checkpoint_management_multicast_address_range`
* **New Resource:** `checkpoint_management_group_with_exclusion`
* **New Resource:** `checkpoint_management_security_zone`
* **New Resource:** `checkpoint_management_time_group`
* **New Resource:** `checkpoint_management_access_role`
* **New Resource:** `checkpoint_management_dynamic_object`
* **New Resource:** `checkpoint_management_dns_domain`
* **New Resource:** `checkpoint_management_opsec_application`
* **New Resource:** `checkpoint_management_service_icmp`
* **New Resource:** `checkpoint_management_service_icmp6`
* **New Resource:** `checkpoint_management_service_sctp`
* **New Resource:** `checkpoint_management_service_other`
* **New Resource:** `checkpoint_management_application_site`
* **New Resource:** `checkpoint_management_application_site_category`
* **New Resource:** `checkpoint_management_application_site_group`
* **New Resource:** `checkpoint_management_service_dce_rpc`
* **New Resource:** `checkpoint_management_service_rpc`
* **New Resource:** `checkpoint_management_access_section`
* **New Resource:** `checkpoint_management_access_layer`
* **New Resource:** `checkpoint_management_vpn_community_meshed`
* **New Resource:** `checkpoint_management_vpn_community_star`
* **New Resource:** `checkpoint_management_exception_group`
* **New Resource:** `checkpoint_management_https_rule`
* **New Resource:** `checkpoint_management_https_section`
* **New Resource:** `checkpoint_management_https_layer`
* **New Resource:** `checkpoint_management_discard`
* **New Resource:** `checkpoint_management_disconnect`
* **New Resource:** `checkpoint_management_keepalive`
* **New Resource:** `checkpoint_management_revert_to_revision`
* **New Resource:** `checkpoint_management_verify_revert`
* **New Resource:** `checkpoint_management_set_login_message`
* **New Resource:** `checkpoint_management_add_data_center_object`
* **New Resource:** `checkpoint_management_delete_data_center_object`
* **New Resource:** `checkpoint_management_update_updatable_objects_repository_content`
* **New Resource:** `checkpoint_management_add_updatable_object`
* **New Resource:** `checkpoint_management_delete_updatable_object`
* **New Resource:** `checkpoint_management_set_ips_update_schedule`
* **New Resource:** `checkpoint_management_run_threat_emulation_file_types_offline_update`
* **New Resource:** `checkpoint_management_verify_policy`
* **New Resource:** `checkpoint_management_set_global_domain`
* **New Resource:** `checkpoint_management_assign_global_assignment`
* **New Resource:** `checkpoint_management_restore_domain`
* **New Resource:** `checkpoint_management_migrate_import_domain`
* **New Resource:** `checkpoint_management_backup_domain`
* **New Resource:** `checkpoint_management_migrate_export_domain`
* **New Resource:** `checkpoint_management_uninstall_software_package`
* **New Resource:** `checkpoint_management_verify_software_package`
* **New Resource:** `checkpoint_management_install_software_package`
* **New Resource:** `checkpoint_management_unlock_administrator`
* **New Resource:** `checkpoint_management_add_api_key`
* **New Resource:** `checkpoint_management_delete_api_key`
* **New Resource:** `checkpoint_management_set_api_settings`
* **New Resource:** `checkpoint_management_export`
* **New Resource:** `checkpoint_management_put_file`
* **New Resource:** `checkpoint_management_where_used`
* **New Resource:** `checkpoint_management_run_script`
* **New Resource:** `checkpoint_management_install_database`
* **New Resource:** `checkpoint_management_set_threat_protection`
* **New Resource:** `checkpoint_management_add_threat_protections`
* **New Resource:** `checkpoint_management_delete_threat_protections`
* **New Feature:** Added multi domain server support
* **New Feature:** Added commands support - publish and install policy after execution

## 1.0.0 (January 13, 2020)

FEATURES:

* **New Resource:** `checkpoint_management_network`
* **New Resource:** `checkpoint_management_host`
* **New Resource:** `checkpoint_management_publish`
* **New Resource:** `checkpoint_hostname`
* **New Resource:** `checkpoint_physical_interface`
* **New Resource:** `checkpoint_put_file`
* **New Resource:** `checkpoint_management_install_policy`
* **New Resource:** `checkpoint_management_run_ips_update`
* **New Resource:** `checkpoint_management_address_range`
* **New Resource:** `checkpoint_management_group`
* **New Resource:** `checkpoint_management_service_group`
* **New Resource:** `checkpoint_management_service_tcp`
* **New Resource:** `checkpoint_management_service_udp`
* **New Resource:** `checkpoint_management_package`
* **New Resource:** `checkpoint_management_access_rule`
* **New Resource:** `checkpoint_management_login`
* **New Resource:** `checkpoint_management_logout`
* **New Resource:** `checkpoint_management_threat_indicator`
