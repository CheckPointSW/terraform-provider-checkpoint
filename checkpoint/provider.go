package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"strings"
	"time"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"server": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_SERVER", nil),
				Description: "Check Point Management server IP",
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_USERNAME", nil),
				Description: "Check Point Management admin name",
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_PASSWORD", nil),
				Description: "Check Point Management admin password",
			},
			"context": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_CONTEXT", checkpoint.WebContext),
				Description: "Check Point access context - gaia_api or web_api",
			},
			"domain": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_DOMAIN", nil),
				Description: "login to specific domain. Domain can be identified by name or UID",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_TIMEOUT", 120),
				Description: "Timeout in seconds for the Go SDK to complete a transaction",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_PORT", checkpoint.DefaultPort),
				Description: "Port used for connection to the API server",
			},
			"session_file_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_SESSION_FILE_NAME", DefaultSessionFilename),
				Description: "File name used to store the current session id",
			},
			"proxy_host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_PROXY_HOST", checkpoint.DefaultProxyHost),
				Description: "HTTP proxy server address (without \"http://\")",
			},
			"proxy_port": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_PROXY_PORT", checkpoint.DefaultProxyPort),
				Description: "HTTP proxy port",
			},
			"api_key": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_API_KEY", ""),
				Description: "Administrator API key",
			},
			"session_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_SESSION_NAME", ""),
				Description: "Session unique name",
			},
			"session_description": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_SESSION_DESCRIPTION", ""),
				Description: "Session purpose description",
			},
			"session_timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_SESSION_TIMEOUT", -1),
				Description: "Timeout for the Check Point session in seconds. Can be 10-3600",
			},
			"cloud_mgmt_id": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_CLOUD_MGMT_ID", ""),
				Description: "Smart-1 Cloud management UID",
			},
			"auto_publish_batch_size": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_AUTO_PUBLISH_BATCH_SIZE", -1),
				Description: "Number of batch size to automatically run publish",
			},
			"ignore_server_certificate": {
				Type:        schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_IGNORE_SERVER_CERTIFICATE", false),
				Description: "Indicates that the client should not check the server's certificate",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"checkpoint_management_outbound_inspection_certificate":                resourceManagementOutboundInspectionCertificate(),
			"checkpoint_management_run_trusted_ca_update":                          resourceManagementRunTrustedCaUpdate(),
			"checkpoint_management_delete_custom_trusted_ca_certificate":           resourceManagementDeleteCustomTrustedCaCertificate(),
			"checkpoint_management_add_custom_trusted_ca_certificate":              resourceManagementAddCustomTrustedCaCertificate(),
			"checkpoint_management_set_https_advanced_settings":                    resourceManagementSetHttpsAdvancedSettings(),
			"checkpoint_management_set_gateway_global_use":                         resourceManagementSetGatewayGlobalUse(),
			"checkpoint_management_vsx_provisioning_tool":                          resourceManagementVsxProvisioningTool(),
			"checkpoint_management_delete_infinity_idp_object":                     resourceManagementDeleteInfinityIdpObject(),
			"checkpoint_management_delete_infinity_idp":                            resourceManagementDeleteInfinityIdp(),
			"checkpoint_management_mobile_access_section":                          resourceManagementMobileAccessSection(),
			"checkpoint_management_mobile_access_rule":                             resourceManagementMobileAccessRule(),
			"checkpoint_management_mobile_access_profile_section":                  resourceManagementMobileAccessProfileSection(),
			"checkpoint_management_mobile_access_profile_rule":                     resourceManagementMobileAccessProfileRule(),
			"checkpoint_management_network_probe":                                  resourceManagementNetworkProbe(),
			"checkpoint_management_override_categorization":                        resourceManagementOverrideCategorization(),
			"checkpoint_management_interface":                                      resourceManagementInterface(),
			"checkpoint_management_resource_cifs":                                  resourceManagementResourceCifs(),
			"checkpoint_management_resource_ftp":                                   resourceManagementResourceFtp(),
			"checkpoint_management_resource_smtp":                                  resourceManagementResourceSmtp(),
			"checkpoint_management_resource_uri":                                   resourceManagementResourceUri(),
			"checkpoint_management_mobile_profile":                                 resourceManagementMobileProfile(),
			"checkpoint_management_passcode_profile":                               resourceManagementPasscodeProfile(),
			"checkpoint_management_command_set_trusted_ca_settings":                resourceManagementSetTrustedCaSettings(),
			"checkpoint_management_command_set_cp_trusted_ca_certificate":          resourceManagementSetCpTrustedCaCertificate(),
			"checkpoint_management_command_set_internal_trusted_ca":                resourceManagementSetInternalTrustedCa(),
			"checkpoint_management_external_trusted_ca":                            resourceManagementExternalTrustedCa(),
			"checkpoint_management_opsec_trusted_ca":                               resourceManagementOpsecTrustedCa(),
			"checkpoint_management_multiple_key_exchanges":                         resourceManagementMultipleKeyExchanges(),
			"checkpoint_management_app_control_advanced_settings":                  resourceManagementSetAppControlAdvancedSettings(),
			"checkpoint_management_content_awareness_advanced_settings":            resourceManagementSetContentAwarenessAdvancedSettings(),
			"checkpoint_management_limit":                                          resourceManagementLimit(),
			"checkpoint_management_data_type_compound_group":                       resourceManagementDataTypeCompoundGroup(),
			"checkpoint_management_data_type_traditional_group":                    resourceManagementDataTypeTraditionalGroup(),
			"checkpoint_management_data_type_group":                                resourceManagementDataTypeGroup(),
			"checkpoint_management_data_type_patterns":                             resourceManagementDataTypePatterns(),
			"checkpoint_management_data_type_file_attributes":                      resourceManagementDataTypeFileAttributes(),
			"checkpoint_management_data_type_weighted_keywords":                    resourceManagementDataTypeWeightedKeywords(),
			"checkpoint_management_data_type_keywords":                             resourceManagementDataTypeKeywords(),
			"checkpoint_management_data_center_object":                             resourceDataCenterObject(),
			"checkpoint_management_lsm_cluster":                                    resourceManagementLsmCluster(),
			"checkpoint_management_lsm_gateway":                                    resourceManagementLsmGateway(),
			"checkpoint_management_service_gtp":                                    resourceManagementServiceGtp(),
			"checkpoint_management_smart_task":                                     resourceManagementSmartTask(),
			"checkpoint_management_server_certificate":                             resourceManagementServerCertificate(),
			"checkpoint_management_host":                                           resourceManagementHost(),
			"checkpoint_management_network":                                        resourceManagementNetwork(),
			"checkpoint_management_wildcard":                                       resourceManagementWildcard(),
			"checkpoint_management_group":                                          resourceManagementGroup(),
			"checkpoint_management_address_range":                                  resourceManagementAddressRange(),
			"checkpoint_management_multicast_address_range":                        resourceManagementMulticastAddressRange(),
			"checkpoint_management_group_with_exclusion":                           resourceManagementGroupWithExclusion(),
			"checkpoint_management_security_zone":                                  resourceManagementSecurityZone(),
			"checkpoint_management_time_group":                                     resourceManagementTimeGroup(),
			"checkpoint_management_access_role":                                    resourceManagementAccessRole(),
			"checkpoint_management_dynamic_object":                                 resourceManagementDynamicObject(),
			"checkpoint_management_dns_domain":                                     resourceManagementDnsDomain(),
			"checkpoint_management_opsec_application":                              resourceManagementOpsecApplication(),
			"checkpoint_management_service_tcp":                                    resourceManagementServiceTcp(),
			"checkpoint_management_service_udp":                                    resourceManagementServiceUdp(),
			"checkpoint_management_service_icmp":                                   resourceManagementServiceIcmp(),
			"checkpoint_management_service_icmp6":                                  resourceManagementServiceIcmp6(),
			"checkpoint_management_service_sctp":                                   resourceManagementServiceSctp(),
			"checkpoint_management_service_other":                                  resourceManagementServiceOther(),
			"checkpoint_management_service_group":                                  resourceManagementServiceGroup(),
			"checkpoint_management_application_site":                               resourceManagementApplicationSite(),
			"checkpoint_management_application_site_category":                      resourceManagementApplicationSiteCategory(),
			"checkpoint_management_application_site_group":                         resourceManagementApplicationSiteGroup(),
			"checkpoint_management_service_dce_rpc":                                resourceManagementServiceDceRpc(),
			"checkpoint_management_service_rpc":                                    resourceManagementServiceRpc(),
			"checkpoint_management_access_rule":                                    resourceManagementAccessRule(),
			"checkpoint_management_access_section":                                 resourceManagementAccessSection(),
			"checkpoint_management_access_layer":                                   resourceManagementAccessLayer(),
			"checkpoint_management_vpn_community_meshed":                           resourceManagementVpnCommunityMeshed(),
			"checkpoint_management_vpn_community_star":                             resourceManagementVpnCommunityStar(),
			"checkpoint_management_exception_group":                                resourceManagementExceptionGroup(),
			"checkpoint_management_threat_indicator":                               resourceManagementThreatIndicator(),
			"checkpoint_management_https_rule":                                     resourceManagementHttpsRule(),
			"checkpoint_management_https_section":                                  resourceManagementHttpsSection(),
			"checkpoint_management_https_layer":                                    resourceManagementHttpsLayer(),
			"checkpoint_management_discard":                                        resourceManagementDiscard(),
			"checkpoint_management_disconnect":                                     resourceManagementDisconnect(),
			"checkpoint_management_keepalive":                                      resourceManagementKeepalive(),
			"checkpoint_management_revert_to_revision":                             resourceManagementRevertToRevision(),
			"checkpoint_management_verify_revert":                                  resourceManagementVerifyRevert(),
			"checkpoint_management_set_login_message":                              resourceManagementSetLoginMessage(),
			"checkpoint_management_add_data_center_object":                         resourceManagementAddDataCenterObject(),
			"checkpoint_management_delete_data_center_object":                      resourceManagementDeleteDataCenterObject(),
			"checkpoint_management_update_updatable_objects_repository_content":    resourceManagementUpdateUpdatableObjectsRepositoryContent(),
			"checkpoint_management_add_updatable_object":                           resourceManagementAddUpdatableObject(),
			"checkpoint_management_delete_updatable_object":                        resourceManagementDeleteUpdatableObject(),
			"checkpoint_management_set_ips_update_schedule":                        resourceManagementSetIpsUpdateSchedule(),
			"checkpoint_management_run_threat_emulation_file_types_offline_update": resourceManagementRunThreatEmulationFileTypesOfflineUpdate(),
			"checkpoint_management_verify_policy":                                  resourceManagementVerifyPolicy(),
			"checkpoint_management_set_global_domain":                              resourceManagementSetGlobalDomain(),
			"checkpoint_management_assign_global_assignment":                       resourceManagementAssignGlobalAssignment(),
			"checkpoint_management_restore_domain":                                 resourceManagementRestoreDomain(),
			"checkpoint_management_migrate_import_domain":                          resourceManagementMigrateImportDomain(),
			"checkpoint_management_backup_domain":                                  resourceManagementBackupDomain(),
			"checkpoint_management_migrate_export_domain":                          resourceManagementMigrateExportDomain(),
			"checkpoint_management_uninstall_software_package":                     resourceManagementUninstallSoftwarePackage(),
			"checkpoint_management_package":                                        resourceManagementPackage(),
			"checkpoint_management_verify_software_package":                        resourceManagementVerifySoftwarePackage(),
			"checkpoint_management_install_software_package":                       resourceManagementInstallSoftwarePackage(),
			"checkpoint_management_unlock_administrator":                           resourceManagementUnlockAdministrator(),
			"checkpoint_management_add_api_key":                                    resourceManagementAddApiKey(),
			"checkpoint_management_delete_api_key":                                 resourceManagementDeleteApiKey(),
			"checkpoint_management_set_api_settings":                               resourceManagementSetApiSettings(),
			"checkpoint_management_export":                                         resourceManagementExport(),
			"checkpoint_management_put_file":                                       resourceManagementPutFile(),
			"checkpoint_management_where_used":                                     resourceManagementWhereUsed(),
			"checkpoint_management_run_script":                                     resourceManagementRunScript(),
			"checkpoint_management_install_database":                               resourceManagementInstallDatabase(),
			"checkpoint_management_set_threat_protection":                          resourceManagementSetThreatProtection(),
			"checkpoint_management_add_threat_protections":                         resourceManagementAddThreatProtections(),
			"checkpoint_management_delete_threat_protections":                      resourceManagementDeleteThreatProtections(),
			"checkpoint_hostname":                                                  resourceHostname(),
			"checkpoint_put_file":                                                  resourcePutFile(),
			"checkpoint_physical_interface":                                        resourcePhysicalInterface(),
			"checkpoint_management_login":                                          resourceManagementLogin(),
			"checkpoint_management_logout":                                         resourceManagementLogout(),
			"checkpoint_management_publish":                                        resourceManagementPublish(),
			"checkpoint_management_install_policy":                                 resourceManagementInstallPolicy(),
			"checkpoint_management_run_ips_update":                                 resourceManagementRunIpsUpdate(),
			"checkpoint_management_access_point_name":                              resourceManagementAccessPointName(),
			"checkpoint_management_gsn_handover_group":                             resourceManagementGsnHandoverGroup(),
			"checkpoint_management_identity_tag":                                   resourceManagementIdentityTag(),
			"checkpoint_management_service_citrix_tcp":                             resourceManagementServiceCitrixTcp(),
			"checkpoint_management_service_compound_tcp":                           resourceManagementServiceCompoundTcp(),
			"checkpoint_management_user_group":                                     resourceManagementUserGroup(),
			"checkpoint_management_user_template":                                  resourceManagementUserTemplate(),
			"checkpoint_management_user":                                           resourceManagementUser(),
			"checkpoint_management_mds":                                            resourceManagementMds(),
			"checkpoint_management_vpn_community_remote_access":                    resourceManagementVpnCommunityRemoteAccess(),
			"checkpoint_management_ha_full_sync":                                   resourceManagementHaFullSync(),
			"checkpoint_management_set_automatic_purge":                            resourceManagementSetAutomaticPurge(),
			"checkpoint_management_set_ha_state":                                   resourceManagementSetHaState(),
			"checkpoint_management_checkpoint_host":                                resourceManagementCheckpointHost(),
			"checkpoint_management_get_attachment":                                 resourceManagementGetAttachment(),
			"checkpoint_management_nat_section":                                    resourceManagementNatSection(),
			"checkpoint_management_nat_rule":                                       resourceManagementNatRule(),
			"checkpoint_management_threat_rule":                                    resourceManagementThreatRule(),
			"checkpoint_management_threat_exception":                               resourceManagementThreatException(),
			"checkpoint_management_simple_gateway":                                 resourceManagementSimpleGateway(),
			"checkpoint_management_simple_cluster":                                 resourceManagementSimpleCluster(),
			"checkpoint_management_threat_profile":                                 resourceManagementThreatProfile(),
			"checkpoint_management_generic_data_center_server":                     resourceManagementGenericDataCenterServer(),
			"checkpoint_management_vmware_data_center_server":                      resourceManagementVMwareDataCenterServer(),
			"checkpoint_management_aci_data_center_server":                         resourceManagementAciDataCenterServer(),
			"checkpoint_management_ise_data_center_server":                         resourceManagementIseDataCenterServer(),
			"checkpoint_management_aws_data_center_server":                         resourceManagementAwsDataCenterServer(),
			"checkpoint_management_gcp_data_center_server":                         resourceManagementGcpDataCenterServer(),
			"checkpoint_management_azure_data_center_server":                       resourceManagementAzureDataCenterServer(),
			"checkpoint_management_nuage_data_center_server":                       resourceManagementNuageDataCenterServer(),
			"checkpoint_management_openstack_data_center_server":                   resourceManagementOpenStackDataCenterServer(),
			"checkpoint_management_kubernetes_data_center_server":                  resourceManagementKubernetesDataCenterServer(),
			"checkpoint_management_data_center_query":                              resourceManagementDataCenterQuery(),
			"checkpoint_management_threat_ioc_feed":                                resourceManagementThreatIocFeed(),
			"checkpoint_management_smtp_server":                                    resourceManagementSmtpServer(),
			"checkpoint_management_network_feed":                                   resourceManagementNetworkFeed(),
			"checkpoint_management_md_permissions_profile":                         resourceManagementMdPermissionsProfile(),
			"checkpoint_management_interoperable_device":                           resourceManagementInteroperableDevice(),
			"checkpoint_management_domain_permissions_profile":                     resourceManagementDomainPermissionsProfile(),
			"checkpoint_management_idp_administrator_group":                        resourceManagementIdpAdministratorGroup(),
			"checkpoint_management_check_threat_ioc_feed":                          resourceManagementCheckThreatIocFeed(),
			"checkpoint_management_check_network_feed":                             resourceManagementCheckNetworkFeed(),
			"checkpoint_management_get_platform":                                   resourceManagementGetPlatform(),
			"checkpoint_management_install_lsm_policy":                             resourceManagementInstallLsmPolicy(),
			"checkpoint_management_install_lsm_settings":                           resourceManagementInstallLsmSettings(),
			"checkpoint_management_lsm_run_script":                                 resourceManagementLsmRunScript(),
			"checkpoint_management_reset_sic":                                      resourceManagementResetSic(),
			"checkpoint_management_test_sic_status":                                resourceManagementTestSicStatus(),
			"checkpoint_management_update_provisioned_satellites":                  resourceManagementUpdateProvisionedSatellites(),
			"checkpoint_management_repository_script":                              resourceManagementRepositoryScript(),
			"checkpoint_management_set_idp_default_assignment":                     resourceManagementSetIdpDefaultAssignment(),
			"checkpoint_management_set_idp_to_domain_assignment":                   resourceManagementSetIdpToDomainAssignment(),
			"checkpoint_management_domain":                                         resourceManagementDomain(),
			"checkpoint_management_add_repository_package":                         resourceManagementAddRepositoryPackage(),
			"checkpoint_management_delete_repository_package":                      resourceManagementDeleteRepositoryPackage(),
			"checkpoint_management_time":                                           resourceManagementTime(),
			"checkpoint_management_trusted_client":                                 resourceManagementTrustedClient(),
			"checkpoint_management_connect_cloud_services":                         resourceManagementConnectCloudServices(),
			"checkpoint_management_disconnect_cloud_services":                      resourceManagementDisconnectCloudServices(),
			"checkpoint_management_tacacs_group":                                   resourceManagementTacacsGroup(),
			"checkpoint_management_tacacs_server":                                  resourceManagementTacacsServer(),
			"checkpoint_management_tag":                                            resourceManagementTag(),
			"checkpoint_management_command_unlock_object":                          resourceManagementUnlockObject(),
			"checkpoint_management_threat_layer":                                   resourceManagementThreatLayer(),
			"checkpoint_management_radius_server":                                  resourceManagementRadiusServer(),
			"checkpoint_management_radius_group":                                   resourceManagementRadiusGroup(),
			"checkpoint_management_command_set_policy_settings":                    resourceManagementSetPolicySettings(),
			"checkpoint_management_command_set_threat_advanced_settings":           resourceManagementSetThreatAdvancedSettings(),
			"checkpoint_management_lsv_profile":                                    resourceManagementLsvProfile(),
			"checkpoint_management_command_login_to_domain":                        resourceManagementLoginToDomain(),
			"checkpoint_management_command_lock_object":                            resourceManagementLockObject(),
			"checkpoint_management_command_gaia_api":                               resourceManagementGaiaApi(),
			"checkpoint_management_command_import_smart_task":                      resourceManagementImportSmartTask(),
			"checkpoint_management_command_get_interfaces":                         resourceManagementGetInterfaces(),
			"checkpoint_management_command_abort_get_interfaces":                   resourceManagementAbortGetInterfaces(),
			"checkpoint_management_command_export_management":                      resourceManagementExportManagement(),
			"checkpoint_management_command_export_smart_task":                      resourceManagementExportSmartTask(),
			"checkpoint_management_command_import_management":                      resourceManagementImportManagement(),
			"checkpoint_management_command_set_global_properties":                  resourceManagementSetGlobalProperties(),
			"checkpoint_management_administrator":                                  resourceManagementAdministrator(),
			"checkpoint_management_oracle_cloud_data_center_server":                resourceManagementOracleCloudDataCenterServer(),
			"checkpoint_management_proxmox_data_center_server":                     resourceManagementProxmoxDataCenterServer(),
			"checkpoint_management_nutanix_data_center_server":                     resourceManagementNutanixDataCenterServer(),
			"checkpoint_management_azure_ad":                                       resourceManagementAzureAd(),
			"checkpoint_management_gaia_best_practice":                             resourceManagementGaiaBestPractice(),
			"checkpoint_management_dynamic_global_network_object":                  resourceManagementDynamicGlobalNetworkObject(),
			"checkpoint_management_global_assignment":                              resourceManagementGlobalAssignment(),
			"checkpoint_management_cme_delay_cycle":                                resourceManagementCMEDelayCycle(),
			"checkpoint_management_cme_management":                                 resourceManagementCMEManagement(),
			"checkpoint_management_cme_accounts_azure":                             resourceManagementCMEAccountsAzure(),
			"checkpoint_management_cme_accounts_gcp":                               resourceManagementCMEAccountsGCP(),
			"checkpoint_management_cme_accounts_aws":                               resourceManagementCMEAccountsAWS(),
			"checkpoint_management_cme_gw_configurations_aws":                      resourceManagementCMEGWConfigurationsAWS(),
			"checkpoint_management_cme_gw_configurations_azure":                    resourceManagementCMEGWConfigurationsAzure(),
			"checkpoint_management_cme_gw_configurations_gcp":                      resourceManagementCMEGWConfigurationsGCP(),
			"checkpoint_generic_api":                                               resourceManagementGenericApi(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"checkpoint_management_outbound_inspection_certificate":           dataSourceManagementOutboundInspectionCertificate(),
			"checkpoint_management_custom_trusted_ca_certificate":             dataSourceManagementCustomTrustedCaCertificate(),
			"checkpoint_management_https_advanced_settings":                   dataSourceManagementSetHttpsAdvancedSettings(),
			"checkpoint_management_gateway_global_use":                        dataSourceManagementSetGatewayGlobalUse(),
			"checkpoint_management_gateway_capabilities":                      dataSourceManagementSetGatewayCapabilities(),
			"checkpoint_management_infinity_idp_object":                       dataSourceManagementInfinityIdpObject(),
			"checkpoint_management_infinity_idp":                              dataSourceManagementInfinityIdp(),
			"checkpoint_management_mobile_access_section":                     dataSourceManagementMobileAccessSection(),
			"checkpoint_management_mobile_access_rule":                        dataSourceManagementMobileAccessRule(),
			"checkpoint_management_mobile_access_profile_section":             dataSourceManagementMobileAccessProfileSection(),
			"checkpoint_management_mobile_access_profile_rule":                dataSourceManagementMobileAccessProfileRule(),
			"checkpoint_management_network_probe":                             dataSourceManagementNetworkProbe(),
			"checkpoint_management_override_categorization":                   dataSourceManagementOverrideCategorization(),
			"checkpoint_management_interface":                                 dataSourceManagementInterface(),
			"checkpoint_management_resource_cifs":                             dataSourceManagementResourceCifs(),
			"checkpoint_management_resource_ftp":                              dataSourceManagementResourceFtp(),
			"checkpoint_management_resource_smtp":                             dataSourceManagementResourceSmtp(),
			"checkpoint_management_resource_uri":                              dataDourceManagementResourceUri(),
			"checkpoint_management_mobile_profile":                            dataSourceManagementMobileProfile(),
			"checkpoint_management_passcode_profile":                          dataSourceManagementPasscodeProfile(),
			"checkpoint_management_trusted_ca_settings":                       dataSourceManagementSetTrustedCaSettings(),
			"checkpoint_management_internal_trusted_ca":                       dataSourceManagementSetInternalTrustedCa(),
			"checkpoint_management_external_trusted_ca":                       dataSourceManagementExternalTrustedCa(),
			"checkpoint_management_cp_trusted_ca_certificate":                 dataSourceManagementCpTrustedCaCertificate(),
			"checkpoint_management_opsec_trusted_ca":                          dataSourceManagementOpsecTrustedCa(),
			"checkpoint_management_multiple_key_exchanges":                    dataSourceManagementMultipleKeyExchanges(),
			"checkpoint_management_app_control_advanced_settings":             dataSourceManagementAppControlAdvancedSettings(),
			"checkpoint_management_content_awareness_advanced_settings":       dataSourceManagementContentAwarenessAdvancedSettings(),
			"checkpoint_management_limit":                                     dataSourceManagementLimit(),
			"checkpoint_management_data_type_compound_group":                  dataSourceManagementDataTypeCompoundGroup(),
			"checkpoint_management_data_type_traditional_group":               dataSourceManagementDataTypeTraditionalGroup(),
			"checkpoint_management_data_type_group":                           dataResourceManagementDataTypeGroup(),
			"checkpoint_management_data_type_patterns":                        dataSourceManagementDataTypePatterns(),
			"checkpoint_management_data_type_file_attributes":                 dataSourceManagementDataTypeFileAttributes(),
			"checkpoint_management_data_type_weighted_keywords":               dataSourceManagementDataTypeWeightedKeywords(),
			"checkpoint_management_data_type_keywords":                        dataSourceManagementDataTypeKeywords(),
			"checkpoint_management_data_center_object":                        dataSourceManagementDataCenterObject(),
			"checkpoint_management_updatable_object":                          dataSourceManagementShowUpdatableObject(),
			"checkpoint_management_lsm_cluster":                               dataSourceManagementLsmCluster(),
			"checkpoint_management_lsm_gateway":                               dataSourceManagementLsmGateway(),
			"checkpoint_management_threat_rule_exception_rulebase":            dataSourceManagementShowThreatRuleExceptionRuleBase(),
			"checkpoint_management_smart_task":                                dataSourceManagementSmartTask(),
			"checkpoint_management_service_gtp":                               dataSourceManagementServiceGtp(),
			"checkpoint_management_server_certificate":                        dataSourceManagementServerCertificate(),
			"checkpoint_management_data_host":                                 dataSourceManagementHost(),
			"checkpoint_management_data_network":                              dataSourceManagementNetwork(),
			"checkpoint_management_data_group":                                dataSourceManagementGroup(),
			"checkpoint_management_data_group_with_exclusion":                 dataSourceManagementGroupWithExclusion(),
			"checkpoint_management_data_access_layer":                         dataSourceManagementAccessLayer(),
			"checkpoint_management_data_access_role":                          dataSourceManagementAccessRole(),
			"checkpoint_management_data_access_rule":                          dataSourceManagementAccessRule(),
			"checkpoint_management_data_access_section":                       dataSourceManagementAccessSection(),
			"checkpoint_management_data_address_range":                        dataSourceManagementAddressRange(),
			"checkpoint_management_data_application_site":                     dataSourceManagementApplicationSite(),
			"checkpoint_management_data_application_site_category":            dataSourceManagementApplicationSiteCategory(),
			"checkpoint_management_data_application_site_group":               dataSourceManagementApplicationSiteGroup(),
			"checkpoint_management_data_dns_domain":                           dataSourceManagementDnsDomain(),
			"checkpoint_management_data_dynamic_object":                       dataSourceManagementDynamicObject(),
			"checkpoint_management_data_exception_group":                      dataSourceManagementExceptionGroup(),
			"checkpoint_management_data_https_layer":                          dataSourceManagementHttpsLayer(),
			"checkpoint_management_data_https_rule":                           dataSourceManagementHttpsRule(),
			"checkpoint_management_data_https_section":                        dataSourceManagementHttpsSection(),
			"checkpoint_management_data_multicast_address_range":              dataSourceManagementMulticastAddressRange(),
			"checkpoint_management_data_opsec_application":                    dataSourceManagementOpsecApplication(),
			"checkpoint_management_data_package":                              dataSourceManagementPackage(),
			"checkpoint_management_data_security_zone":                        dataSourceManagementSecurityZone(),
			"checkpoint_management_data_service_dce_rpc":                      dataSourceManagementServiceDceRpc(),
			"checkpoint_management_data_service_group":                        dataSourceManagementServiceGroup(),
			"checkpoint_management_data_service_icmp":                         dataSourceManagementServiceIcmp(),
			"checkpoint_management_data_service_icmp6":                        dataSourceManagementServiceIcmp6(),
			"checkpoint_management_data_service_other":                        dataSourceManagementServiceOther(),
			"checkpoint_management_data_service_rpc":                          dataSourceManagementServiceRpc(),
			"checkpoint_management_data_service_sctp":                         dataSourceManagementServiceSctp(),
			"checkpoint_management_data_service_tcp":                          dataSourceManagementServiceTcp(),
			"checkpoint_management_data_service_udp":                          dataSourceManagementServiceUdp(),
			"checkpoint_management_data_threat_indicator":                     dataSourceManagementThreatIndicator(),
			"checkpoint_management_data_time_group":                           dataSourceManagementTimeGroup(),
			"checkpoint_management_data_vpn_community_star":                   dataSourceManagementVpnCommunityStar(),
			"checkpoint_management_data_vpn_community_meshed":                 dataSourceManagementVpnCommunityMeshed(),
			"checkpoint_management_data_wildcard":                             dataSourceManagementWildcard(),
			"checkpoint_management_access_point_name":                         dataSourceManagementAccessPointName(),
			"checkpoint_management_gsn_handover_group":                        dataSourceManagementGsnHandoverGroup(),
			"checkpoint_management_identity_tag":                              dataSourceManagementIdentityTag(),
			"checkpoint_management_service_citrix_tcp":                        dataSourceManagementServiceCitrixTcp(),
			"checkpoint_management_service_compound_tcp":                      dataSourceManagementServiceCompoundTcp(),
			"checkpoint_management_user":                                      dataSourceManagementUser(),
			"checkpoint_management_user_group":                                dataSourceManagementUserGroup(),
			"checkpoint_management_user_template":                             dataSourceManagementUserTemplate(),
			"checkpoint_management_vpn_community_remote_access":               dataSourceManagementVpnCommunityRemoteAccess(),
			"checkpoint_management_checkpoint_host":                           dataSourceManagementCheckpointHost(),
			"checkpoint_management_mds":                                       dataSourceManagementMds(),
			"checkpoint_management_show_objects":                              dataSourceManagementShowObjects(),
			"checkpoint_management_show_updatable_objects_repository_content": dataSourceManagementShowUpdatableObjectsRepositoryContent(),
			"checkpoint_management_nat_rule":                                  dataSourceManagementNatRule(),
			"checkpoint_management_nat_section":                               dataSourceManagementNatSection(),
			"checkpoint_management_threat_rule":                               dataSourceManagementThreatRule(),
			"checkpoint_management_threat_exception":                          dataSourceManagementThreatException(),
			"checkpoint_management_simple_cluster":                            dataSourceManagementSimpleCluster(),
			"checkpoint_management_simple_gateway":                            dataSourceManagementSimpleGateway(),
			"checkpoint_management_threat_profile":                            dataSourceManagementThreatProfile(),
			"checkpoint_management_generic_data_center_server":                dataSourceManagementGenericDataCenterServer(),
			"checkpoint_management_vmware_data_center_server":                 dataSourceManagementVMwareDataCenterServer(),
			"checkpoint_management_aci_data_center_server":                    dataSourceManagementAciDataCenterServer(),
			"checkpoint_management_ise_data_center_server":                    dataSourceManagementIseDataCenterServer(),
			"checkpoint_management_aws_data_center_server":                    dataSourceManagementAwsDataCenterServer(),
			"checkpoint_management_gcp_data_center_server":                    dataSourceManagementGcpDataCenterServer(),
			"checkpoint_management_azure_data_center_server":                  dataSourceManagementAzureDataCenterServer(),
			"checkpoint_management_nuage_data_center_server":                  dataSourceManagementNuageDataCenterServer(),
			"checkpoint_management_openstack_data_center_server":              dataSourceManagementOpenStackDataCenterServer(),
			"checkpoint_management_kubernetes_data_center_server":             dataSourceManagementKubernetesDataCenterServer(),
			"checkpoint_management_data_center_query":                         dataSourceManagementDataCenterQuery(),
			"checkpoint_management_data_center_content":                       dataSourceManagementDataCenterContent(),
			"checkpoint_management_access_rulebase":                           dataSourceManagementAccessRuleBase(),
			"checkpoint_management_nat_rulebase":                              dataSourceManagementNatRuleBase(),
			"checkpoint_management_threat_rulebase":                           dataSourceManagementThreatRuleBase(),
			"checkpoint_management_https_rulebase":                            dataSourceManagementHttpsRuleBase(),
			"checkpoint_management_threat_ioc_feed":                           dataSourceManagementThreatIocFeed(),
			"checkpoint_management_smtp_server":                               dataSourceManagementSmtpServer(),
			"checkpoint_management_network_feed":                              dataSourceManagementNetworkFeed(),
			"checkpoint_management_interoperable_device":                      dataSourceManagementInteroperableDevice(),
			"checkpoint_management_idp_administrator_group":                   dataSourceManagementIdpAdministratorGroup(),
			"checkpoint_management_md_permissions_profile":                    dataSourceManagementMdPermissionsProfile(),
			"checkpoint_management_domain_permissions_profile":                dataSourceManagementDomainPermissionsProfile(),
			"checkpoint_management_repository_script":                         dataSourceManagementRepositoryScript(),
			"checkpoint_management_idp_default_assignment":                    dataSourceManagementIdpDefaultAssignment(),
			"checkpoint_management_lsm_cluster_profile":                       dataSourceManagementLsmClusterProfile(),
			"checkpoint_management_lsm_gateway_profile":                       dataSourceManagementLsmGatewayProfile(),
			"checkpoint_management_cluster_member":                            dataSourceManagementClusterMember(),
			"checkpoint_management_provisioning_profile":                      dataSourceManagementProvisioningProfile(),
			"checkpoint_management_idp_to_domain_assignment":                  dataSourceManagementIdpToDomainAssignment(),
			"checkpoint_management_domain":                                    dataSourceManagementDomain(),
			"checkpoint_management_repository_package":                        dataSourceManagementRepositoryPackage(),
			"checkpoint_management_time":                                      dataSourceManagementTime(),
			"checkpoint_management_trusted_client":                            dataSourceManagementTrustedClient(),
			"checkpoint_management_cloud_services":                            dataSourceManagementCloudServices(),
			"checkpoint_management_tacacs_group":                              dataSourceManagementTacacsGroup(),
			"checkpoint_management_tag":                                       dataSourceManagementTag(),
			"checkpoint_management_threat_layer":                              dataSourceManagementThreatLayer(),
			"checkpoint_management_radius_server":                             dataSourceManagementRadiusServer(),
			"checkpoint_management_radius_group":                              dataSourceManagementRadiusGroup(),
			"checkpoint_management_automatic_purge":                           dataSourceManagementAutomaticPurge(),
			"checkpoint_management_api_settings":                              dataSourceManagementApiSettings(),
			"checkpoint_management_ips_update_schedule":                       dataSourceManagementIpsUpdateSchedule(),
			"checkpoint_management_task":                                      dataSourceManagementTask(),
			"checkpoint_management_objects":                                   dataSourceManagementObjects(),
			"checkpoint_management_login_message":                             dataSourceManagementLoginMessage(),
			"checkpoint_management_policy_settings":                           dataSourceManagementPolicySettings(),
			"checkpoint_management_threat_advanced_settings":                  dataSourceManagementThreatAdvancedSettings(),
			"checkpoint_management_smart_task_trigger":                        dataSourceManagementSmartTaskTrigger(),
			"checkpoint_management_lsv_profile":                               dataSourceManagementLsvProfile(),
			"checkpoint_management_ips_protection_extended_attribute":         dataSourceManagementIpsProtectionExtendedAttribute(),
			"checkpoint_management_global_domain":                             dataSourceManagementGlobalDomain(),
			"checkpoint_management_tacacs_server":                             dataSourceManagementTacacsServer(),
			"checkpoint_management_administrator":                             dataSourceManagementAdministrator(),
			"checkpoint_management_proxmox_data_center_server":                dataSourceManagementProxmoxDataCenterServer(),
			"checkpoint_management_nutanix_data_center_server":                dataSourceManagementNutanixDataCenterServer(),
			"checkpoint_management_oracle_cloud_data_center_server":           dataSourceManagementOracleCloudDataCenterServer(),
			"checkpoint_management_azure_ad_content":                          dataSourceManagementAzureAdContent(),
			"checkpoint_management_azure_ad":                                  dataSourceManagementAzureAd(),
			"checkpoint_management_gaia_best_practice":                        dataSourceManagementGaiaBestPractice(),
			"checkpoint_management_dynamic_global_network_object":             dataSourceManagementDynamicGlobalNetworkObject(),
			"checkpoint_management_global_assignment":                         dataSourceManagementGlobalAssignment(),
			"checkpoint_management_cme_version":                               dataSourceManagementCMEVersion(),
			"checkpoint_management_cme_api_versions":                          dataSourceManagementCMEAPIVersions(),
			"checkpoint_management_cme_delay_cycle":                           dataSourceManagementCMEDelayCycle(),
			"checkpoint_management_cme_management":                            dataSourceManagementCMEManagement(),
			"checkpoint_management_cme_accounts":                              dataSourceManagementCMEAccounts(),
			"checkpoint_management_cme_accounts_aws":                          dataSourceManagementCMEAccountsAWS(),
			"checkpoint_management_cme_accounts_azure":                        dataSourceManagementCMEAccountsAzure(),
			"checkpoint_management_cme_accounts_gcp":                          dataSourceManagementCMEAccountsGCP(),
			"checkpoint_management_cme_gw_configurations":                     dataSourceManagementCMEGWConfigurations(),
			"checkpoint_management_cme_gw_configurations_aws":                 dataSourceManagementCMEGWConfigurationsAWS(),
			"checkpoint_management_cme_gw_configurations_azure":               dataSourceManagementCMEGWConfigurationsAzure(),
			"checkpoint_management_cme_gw_configurations_gcp":                 dataSourceManagementCMEGWConfigurationsGCP(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	server := data.Get("server").(string)
	username := data.Get("username").(string)
	password := data.Get("password").(string)
	context := data.Get("context").(string)
	domain := data.Get("domain").(string)
	port := data.Get("port").(int)
	timeout := data.Get("timeout").(int)
	sessionFileName := data.Get("session_file_name").(string)
	proxyHost := data.Get("proxy_host").(string)
	proxyPort := data.Get("proxy_port").(int)
	apiKey := data.Get("api_key").(string)
	sessionName := data.Get("session_name").(string)
	sessionDescription := data.Get("session_description").(string)
	sessionTimeout := data.Get("session_timeout").(int)
	cloudMgmtId := data.Get("cloud_mgmt_id").(string)
	autoPublishBatchSize := data.Get("auto_publish_batch_size").(int)
	ignoreServerCertificate := data.Get("ignore_server_certificate").(bool)

	if server == "" || ((username == "" || password == "") && apiKey == "") {
		return nil, fmt.Errorf("checkpoint-provider missing parameters to initialize (server, (username and password) OR api_key)")
	}

	args := checkpoint.ApiClientArgs{
		Port:                    port,
		Fingerprint:             "",
		Sid:                     "",
		Server:                  server,
		ProxyHost:               proxyHost,
		ProxyPort:               proxyPort,
		ApiVersion:              "",
		IgnoreServerCertificate: ignoreServerCertificate,
		AcceptServerCertificate: false,
		DebugFile:               "deb.txt",
		Context:                 context,
		Timeout:                 time.Duration(timeout),
		Sleep:                   checkpoint.SleepTime,
		UserAgent:               "Terraform",
		CloudMgmtId:             cloudMgmtId,
		AutoPublishBatchSize:    autoPublishBatchSize,
	}

	switch context {
	case checkpoint.WebContext:
		var s Session
		var err error
		s, err = GetSession(sessionFileName)
		if err != nil {
			return nil, err
		}
		if s.Sid != "" {
			args.Sid = s.Sid
		}
		mgmt := checkpoint.APIClient(args)
		if ok := CheckSession(mgmt, s.Uid); !ok {
			// session is not valid, need to perform login
			s, err = login(mgmt, username, password, apiKey, domain, sessionName, sessionDescription, sessionTimeout)
			if err != nil {
				log.Println("Failed to perform login")
				return nil, err
			}
			if err := s.Save(sessionFileName); err != nil {
				return nil, err
			}
		}
		log.Printf("Check Point provider connected with session uid [%s]", s.Uid)
		return mgmt, nil
	case checkpoint.GaiaContext:
		gaia := checkpoint.APIClient(args)
		_, err := login(gaia, username, password, "", "", "", "", sessionTimeout)
		if err != nil {
			log.Println("Failed to perform login")
			return nil, err
		}
		return gaia, nil
	default:
		return nil, fmt.Errorf("Invalid access context. Use 'web_api' or 'gaia_api'")
	}
}

func login(client *checkpoint.ApiClient, username string, pwd string, apiKey string, domain string, sessionName string, sessionDescription string, sessionTimeout int) (Session, error) {
	log.Printf("Perform login")
	var loginRes checkpoint.APIResponse
	var err error

	payload := make(map[string]interface{})
	if sessionName != "" {
		payload["session-name"] = sessionName
	}

	if sessionDescription != "" {
		payload["session-description"] = sessionDescription
	}

	if sessionTimeout != -1 {
		payload["session-timeout"] = sessionTimeout
	}

	if apiKey != "" {
		loginRes, err = client.ApiLoginWithApiKey(apiKey, false, domain, false, payload)
	} else {
		loginRes, err = client.ApiLogin(username, pwd, false, domain, false, payload)
	}

	if err != nil {
		localRequestsError := "invalid character '<' looking for beginning of value"
		if strings.Contains(err.Error(), localRequestsError) {
			return Session{}, fmt.Errorf("Login to management server failed: API server needs to be configured to accept requests from all IP addresses")
		}
		return Session{}, err
	}
	if !loginRes.Success {
		return Session{}, fmt.Errorf(loginRes.ErrorMsg)
	}

	uid := ""
	if val, ok := loginRes.GetData()["uid"]; ok {
		uid = val.(string)
	}

	s := Session{
		Sid: client.GetSessionID(),
		Uid: uid,
	}

	return s, nil
}
