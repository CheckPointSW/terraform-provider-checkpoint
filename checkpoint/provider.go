package checkpoint

import (
	"fmt"
	checkpoint "github.com/Checkpoint/api_go_sdk/APIFiles"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"log"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider {
		Schema: map[string]*schema.Schema{
			"server":{
				Type: schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_SERVER", nil),
				Description: "Check Point Management server IP",
			},
			"username": {
				Type: schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_USERNAME", nil),
				Description: "Check Point Management admin name",
			},
			"password": {
				Type: schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_PASSWORD", nil),
				Description: "Check Point Management admin password",
			},
			"context": {
				Type: schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_CONTEXT", checkpoint.WebContext),
				Description: "Check Point access context - gaia_api or web_api",
			},
			"auto_publish": {
				Type: schema.TypeBool,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHECKPOINT_AUTO_PUBLISH", nil),
				Description: "publish on each change",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"checkpoint_management_network": resourceManagementNetwork(),
			"checkpoint_management_host": resourceManagementHost(),
			"checkpoint_management_publish": resourceManagementPublish(),
			"checkpoint_hostname": resourceHostname(),
			"checkpoint_physical_interface": resourcePhysicalInterface(),
			"checkpoint_put_file": resourcePutFile(),
			"checkpoint_management_install_policy": resourceManagementInstallPolicy(),
			"checkpoint_management_run_ips_update": resourceManagementRunIpsUpdate(),
			"checkpoint_management_address_range": resourceManagementAddressRange(),
			"checkpoint_management_group": resourceManagementGroup(),
			"checkpoint_management_service_group": resourceManagementServiceGroup(),
			"checkpoint_management_service_tcp": resourceManagementServiceTcp(),
			"checkpoint_management_service_udp": resourceManagementServiceUdp(),
			"checkpoint_management_package": resourceManagementPackage(),
			"checkpoint_management_access_rule": resourceManagementAccessRule(),
			"checkpoint_management_login": resourceManagementLogin(),
			"checkpoint_management_logout": resourceManagementLogout(),
			"checkpoint_management_threat_indicator": resourceManagementThreatIndicator(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {

	server := data.Get("server").(string)
	username := data.Get("username").(string)
	password := data.Get("password").(string)
	context := data.Get("context").(string)
	autoPublish := data.Get("auto_publish").(bool)

	if server == "" || username == "" || password == "" {
		return nil, fmt.Errorf("checkpoint-provider missing parameters to initialize (server, username, password)")
	}

	args := checkpoint.ApiClientArgs {
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
		AutoPublish: 			 autoPublish,
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
				s, err := login(mgmt, username, password)
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
			_, err := login(gaia, username, password)
			if err != nil {
				return nil, err
			}
			return gaia, nil
		default:
			return nil, fmt.Errorf("Unsupported access context - gaia_api or web_api")
	}
}

// Perform login. Creating new session...
func login(client *checkpoint.ApiClient, username string, pwd string) (Session, error) {
	log.Printf("Perform login")
	loginRes, err := client.Login(username, pwd,false,"",false,"")
	if err != nil {
		log.Println("Failed to perform login")
		return Session{}, err
	}
	uid := ""
	if val, ok := loginRes.GetData()["uid"]; ok {
		uid = val.(string)
	}

	s := Session {
		Sid: client.GetSessionID(),
		Uid: uid,
	}
	log.Printf("Client connected with new session (SID = %s)", s.Sid)
	return s, nil
}