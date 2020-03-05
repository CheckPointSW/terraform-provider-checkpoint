package checkpoint

import (
	"fmt"
	checkpoint "github.com/CheckPointSW/cp-mgmt-api-go-sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
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
		},
		ResourcesMap: map[string]*schema.Resource{
            "checkpoint_management_host":          resourceManagementHost(),
            "checkpoint_management_network":          resourceManagementNetwork(),
            "checkpoint_management_wildcard":          resourceManagementWildcard(),
            "checkpoint_management_group":          resourceManagementGroup(),
            "checkpoint_management_address_range":          resourceManagementAddressRange(),
            "checkpoint_management_multicast_address_range":          resourceManagementMulticastAddressRange(),
            "checkpoint_management_group_with_exclusion":          resourceManagementGroupWithExclusion(),
            "checkpoint_management_simple_cluster":          resourceManagementSimpleCluster(),
            "checkpoint_management_security_zone":          resourceManagementSecurityZone(),
            "checkpoint_management_time_group":          resourceManagementTimeGroup(),
            "checkpoint_management_access_role":          resourceManagementAccessRole(),
            "checkpoint_management_dynamic_object":          resourceManagementDynamicObject(),
            "checkpoint_management_tag":          resourceManagementTag(),
            "checkpoint_management_dns_domain":          resourceManagementDnsDomain(),
            "checkpoint_management_opsec_application":          resourceManagementOpsecApplication(),
            "checkpoint_management_service_tcp":          resourceManagementServiceTcp(),
            "checkpoint_management_service_udp":          resourceManagementServiceUdp(),
            "checkpoint_management_service_icmp":          resourceManagementServiceIcmp(),
            "checkpoint_management_service_icmp6":          resourceManagementServiceIcmp6(),
            "checkpoint_management_service_sctp":          resourceManagementServiceSctp(),
            "checkpoint_management_service_other":          resourceManagementServiceOther(),
            "checkpoint_management_service_group":          resourceManagementServiceGroup(),
            "checkpoint_management_application_site":          resourceManagementApplicationSite(),
            "checkpoint_management_application_site_category":          resourceManagementApplicationSiteCategory(),
            "checkpoint_management_application_site_group":          resourceManagementApplicationSiteGroup(),
            "checkpoint_management_service_dce_rpc":          resourceManagementServiceDceRpc(),
            "checkpoint_management_service_rpc":          resourceManagementServiceRpc(),
            "checkpoint_management_access_rule":          resourceManagementAccessRule(),
            "checkpoint_management_access_section":          resourceManagementAccessSection(),
            "checkpoint_management_access_layer":          resourceManagementAccessLayer(),
            "checkpoint_management_vpn_community_meshed":          resourceManagementVpnCommunityMeshed(),
            "checkpoint_management_vpn_community_star":          resourceManagementVpnCommunityStar(),
            "checkpoint_management_exception_group":          resourceManagementExceptionGroup(),
            "checkpoint_management_threat_indicator":          resourceManagementThreatIndicator(),
            "checkpoint_management_https_rule":          resourceManagementHttpsRule(),
            "checkpoint_management_https_section":          resourceManagementHttpsSection(),
            "checkpoint_management_https_layer":          resourceManagementHttpsLayer(),
            "checkpoint_management_server_certificate":          resourceManagementServerCertificate(),
            "checkpoint_management_policy_package":          resourceManagementPolicyPackage(),
            "checkpoint_management_discard":          resourceManagementDiscard(),
            "checkpoint_management_disconnect":          resourceManagementDisconnect(),
            "checkpoint_management_keepalive":          resourceManagementKeepalive(),
            "checkpoint_management_revert_to_revision":          resourceManagementRevertToRevision(),
            "checkpoint_management_verify_revert":          resourceManagementVerifyRevert(),
            "checkpoint_management_set_login_message":          resourceManagementSetLoginMessage(),
            "checkpoint_management_add_data_center_object":          resourceManagementAddDataCenterObject(),
            "checkpoint_management_delete_data_center_object":          resourceManagementDeleteDataCenterObject(),
            "checkpoint_management_update_updatable_objects_repository_content":          resourceManagementUpdateUpdatableObjectsRepositoryContent(),
            "checkpoint_management_add_updatable_object":          resourceManagementAddUpdatableObject(),
            "checkpoint_management_delete_updatable_object":          resourceManagementDeleteUpdatableObject(),
            "checkpoint_management_set_ips_update_schedule":          resourceManagementSetIpsUpdateSchedule(),
            "checkpoint_management_run_threat_emulation_file_types_offline_update":          resourceManagementRunThreatEmulationFileTypesOfflineUpdate(),
            "checkpoint_management_verify_policy":          resourceManagementVerifyPolicy(),
            "checkpoint_management_set_global_domain":          resourceManagementSetGlobalDomain(),
            "checkpoint_management_assign_global_assignment":          resourceManagementAssignGlobalAssignment(),
            "checkpoint_management_restore_domain":          resourceManagementRestoreDomain(),
            "checkpoint_management_migrate_import_domain":          resourceManagementMigrateImportDomain(),
            "checkpoint_management_backup_domain":          resourceManagementBackupDomain(),
            "checkpoint_management_migrate_export_domain":          resourceManagementMigrateExportDomain(),
            "checkpoint_management_uninstall_software_package":          resourceManagementUninstallSoftwarePackage(),
            "checkpoint_management_verify_software_package":          resourceManagementVerifySoftwarePackage(),
            "checkpoint_management_install_software_package":          resourceManagementInstallSoftwarePackage(),
            "checkpoint_management_unlock_administrator":          resourceManagementUnlockAdministrator(),
            "checkpoint_management_add_api_key":          resourceManagementAddApiKey(),
            "checkpoint_management_delete_api_key":          resourceManagementDeleteApiKey(),
            "checkpoint_management_set_api_settings":          resourceManagementSetApiSettings(),
            "checkpoint_management_export":          resourceManagementExport(),
            "checkpoint_management_put_file":          resourceManagementPutFile(),
            "checkpoint_management_where_used":          resourceManagementWhereUsed(),
            "checkpoint_management_run_script":          resourceManagementRunScript(),
            "checkpoint_management_install_database":          resourceManagementInstallDatabase(),
            "checkpoint_management_set_threat_protection":          resourceManagementSetThreatProtection(),
            "checkpoint_management_add_threat_protections":          resourceManagementAddThreatProtections(),
            "checkpoint_management_delete_threat_protections":          resourceManagementDeleteThreatProtections(),
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

	if server == "" || username == "" || password == "" {
		return nil, fmt.Errorf("checkpoint-provider missing parameters to initialize (server, username, password)")
	}

	args := checkpoint.ApiClientArgs{
		Port:                    checkpoint.DefaultPort,
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
		Timeout:                 checkpoint.TimeOut,
		Sleep:                   checkpoint.SleepTime,
	}

	switch context {
	case checkpoint.WebContext:
		s, err := GetSession()
		if err != nil {
			return nil, err
		}
		if s.Sid != "" {
			args.Sid = s.Sid
		}
		mgmt := checkpoint.APIClient(args)
		if CheckSession(mgmt, s.Uid) {
			log.Printf("Client connected with last session (SID = %s)", s.Sid)
		} else {
			s, err := login(mgmt, username, password, domain)
			if err != nil {
				return nil, err
			}
			if err := s.Save(); err != nil {
				return nil, err
			}
		}
		return mgmt, nil
	case checkpoint.GaiaContext:
		gaia := checkpoint.APIClient(args)
		_, err := login(gaia, username, password, "")
		if err != nil {
			return nil, err
		}
		return gaia, nil
	default:
		return nil, fmt.Errorf("Unsupported access context - gaia_api or web_api")
	}
}

// Perform login. Creating new session...
func login(client *checkpoint.ApiClient, username string, pwd string, domain string) (Session, error) {
	log.Printf("Perform login")

	loginRes, err := client.Login(username, pwd, false, domain, false, "")
	if err != nil {
		log.Println("Failed to perform login")
		return Session{}, err
	}
	uid := ""
	if val, ok := loginRes.GetData()["uid"]; ok {
		uid = val.(string)
	}

	s := Session{
		Sid: client.GetSessionID(),
		Uid: uid,
	}
	log.Printf("Client connected with new session (SID = %s)", s.Sid)
	return s, nil
}
