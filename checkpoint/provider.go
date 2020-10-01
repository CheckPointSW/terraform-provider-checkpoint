package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
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
				Description: "login to specific domain. Domain can be identified by name or UID.",
			},
			"timeout": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_TIMEOUT", -1),
				Description: "Timeout in seconds for the Go SDK to complete a transaction",
			},
			"port": {
				Type:        schema.TypeInt,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_PORT", checkpoint.DefaultPort),
				Description: "Port used for connection to the API server",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
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
			"checkpoint_management_access_point_name": 								resourceManagementAccessPointName(),
			"checkpoint_management_gsn_handover_group": 							resourceManagementGsnHandoverGroup(),
			"checkpoint_management_identity_tag": 									resourceManagementIdentityTag(),
			"checkpoint_management_service_citrix_tcp": 							resourceManagementServiceCitrixTcp(),
			"checkpoint_management_service_compound_tcp": 							resourceManagementServiceCompoundTcp(),
			"checkpoint_management_user_group": 									resourceManagementUserGroup(),
			"checkpoint_management_user_template": 									resourceManagementUserTemplate(),
			"checkpoint_management_user": 											resourceManagementUser(),
			"checkpoint_management_mds": 											resourceManagementMds(),
			"checkpoint_management_vpn_community_remote_access": 					resourceManagementVpnCommunityRemoteAccess(),
			"checkpoint_management_ha_full_sync": 									resourceManagementHaFullSync(),
			"checkpoint_management_set_automatic_purge": 							resourceManagementSetAutomaticPurge(),
			"checkpoint_management_set_ha_state": 									resourceManagementSetHaState(),
			"checkpoint_management_checkpoint_host": 								resourceManagementCheckpointHost(),
			"checkpoint_management_get_attachment": 								resourceManagementGetAttachment(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"checkpoint_management_data_host":                      dataSourceManagementHost(),
			"checkpoint_management_data_network":                   dataSourceManagementNetwork(),
			"checkpoint_management_data_group":                     dataSourceManagementGroup(),
			"checkpoint_management_data_group_with_exclusion":      dataSourceManagementGroupWithExclusion(),
			"checkpoint_management_data_access_layer":              dataSourceManagementAccessLayer(),
			"checkpoint_management_data_access_role":               dataSourceManagementAccessRole(),
			"checkpoint_management_data_access_rule":               dataSourceManagementAccessRule(),
			"checkpoint_management_data_access_section":            dataSourceManagementAccessSection(),
			"checkpoint_management_data_address_range":             dataSourceManagementAddressRange(),
			"checkpoint_management_data_application_site":          dataSourceManagementApplicationSite(),
			"checkpoint_management_data_application_site_category": dataSourceManagementApplicationSiteCategory(),
			"checkpoint_management_data_application_site_group":    dataSourceManagementApplicationSiteGroup(),
			"checkpoint_management_data_dns_domain":                dataSourceManagementDnsDomain(),
			"checkpoint_management_data_dynamic_object":            dataSourceManagementDynamicObject(),
			"checkpoint_management_data_exception_group":           dataSourceManagementExceptionGroup(),
			"checkpoint_management_data_https_layer":               dataSourceManagementHttpsLayer(),
			"checkpoint_management_data_https_rule":                dataSourceManagementHttpsRule(),
			"checkpoint_management_data_https_section":             dataSourceManagementHttpsSection(),
			"checkpoint_management_data_multicast_address_range":   dataSourceManagementMulticastAddressRange(),
			"checkpoint_management_data_opsec_application":         dataSourceManagementOpsecApplication(),
			"checkpoint_management_data_package":                   dataSourceManagementPackage(),
			"checkpoint_management_data_security_zone":             dataSourceManagementSecurityZone(),
			"checkpoint_management_data_service_dce_rpc":           dataSourceManagementServiceDceRpc(),
			"checkpoint_management_data_service_group":             dataSourceManagementServiceGroup(),
			"checkpoint_management_data_service_icmp":              dataSourceManagementServiceIcmp(),
			"checkpoint_management_data_service_icmp6":             dataSourceManagementServiceIcmp6(),
			"checkpoint_management_data_service_other":             dataSourceManagementServiceOther(),
			"checkpoint_management_data_service_rpc":               dataSourceManagementServiceRpc(),
			"checkpoint_management_data_service_sctp":              dataSourceManagementServiceSctp(),
			"checkpoint_management_data_service_tcp":               dataSourceManagementServiceTcp(),
			"checkpoint_management_data_service_udp":               dataSourceManagementServiceUdp(),
			"checkpoint_management_data_threat_indicator":          dataSourceManagementThreatIndicator(),
			"checkpoint_management_data_time_group":                dataSourceManagementTimeGroup(),
			"checkpoint_management_data_vpn_community_star":        dataSourceManagementVpnCommunityStar(),
			"checkpoint_management_data_vpn_community_meshed":      dataSourceManagementVpnCommunityMeshed(),
			"checkpoint_management_data_wildcard":                  dataSourceManagementWildcard(),
			"checkpoint_management_access_point_name": 				dataSourceManagementAccessPointName(),
			"checkpoint_management_gsn_handover_group": 			dataSourceManagementGsnHandoverGroup(),
			"checkpoint_management_identity_tag": 					dataSourceManagementIdentityTag(),
			"checkpoint_management_service_citrix_tcp": 			dataSourceManagementServiceCitrixTcp(),
			"checkpoint_management_service_compound_tcp": 			dataSourceManagementServiceCompoundTcp(),
			"checkpoint_management_user": 							dataSourceManagementUser(),
			"checkpoint_management_user_group": 					dataSourceManagementUserGroup(),
			"checkpoint_management_user_template": 					dataSourceManagementUserTemplate(),
			"checkpoint_management_vpn_community_remote_access":    dataSourceManagementVpnCommunityRemoteAccess(),
			"checkpoint_management_checkpoint_host":    			dataSourceManagementCheckpointHost(),
			"checkpoint_management_mds":    						dataSourceManagementMds(),
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

	if server == "" || username == "" || password == "" {
		return nil, fmt.Errorf("checkpoint-provider missing parameters to initialize (server, username, password)")
	}

	args := checkpoint.ApiClientArgs{
		Port:                    port,
		Fingerprint:             "",
		Sid:                     "",
		Server:                  server,
		ProxyHost:               "",
		ProxyPort:               -1,
		ApiVersion:              "",
		IgnoreServerCertificate: false,
		AcceptServerCertificate: false,
		DebugFile:               "deb.txt",
		Context:                 context,
		Timeout:                 time.Duration(timeout),
		Sleep:                   checkpoint.SleepTime,
		UserAgent:               "Terraform",
	}

	switch context {
	case checkpoint.WebContext:
		var s Session
		var err error
		s, err = GetSession()
		if err != nil {
			return nil, err
		}
		if s.Sid != "" {
			args.Sid = s.Sid
		}
		mgmt := checkpoint.APIClient(args)
		if ok := CheckSession(mgmt, s.Uid); !ok {
			// session is not valid, need to perform login
			s, err = login(mgmt, username, password, domain)
			if err != nil {
				log.Println("Failed to perform login")
				return nil, err
			}
			if err := s.Save(); err != nil {
				return nil, err
			}
		}
		log.Printf("Check Point provider connected with session uid [%s]", s.Uid)
		return mgmt, nil
	case checkpoint.GaiaContext:
		gaia := checkpoint.APIClient(args)
		_, err := login(gaia, username, password, "")
		if err != nil {
			log.Println("Failed to perform login")
			return nil, err
		}
		return gaia, nil
	default:
		return nil, fmt.Errorf("unsupported access context - gaia_api or web_api")
	}
}

func login(client *checkpoint.ApiClient, username string, pwd string, domain string) (Session, error) {
	log.Printf("Perform login")

	loginRes, err := client.Login(username, pwd, false, domain, false, "")
	if err != nil {
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
