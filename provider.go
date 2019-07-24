package main

import (
	chkp "api_go_sdk/APIFiles"
	"fmt"
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
				DefaultFunc: schema.EnvDefaultFunc("CHKP_SERVER", nil),
				Description: "Check Point Management server IP",
			},
			"username": {
				Type: schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHKP_USERNAME", nil),
				Description: "Check Point Management admin name",
			},
			"password": {
				Type: schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHKP_PASSWORD", nil),
				Description: "Check Point Management admin password",
			},
			"context": {
				Type: schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("CHKP_CONTEXT", chkp.WebContext),
				Description: "Check Point access context - gaia_api or web_api",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"chkp_network": resourceNetwork(),
			"chkp_publish": resourcePublish(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	server := data.Get("server").(string)
	username := data.Get("username").(string)
	password := data.Get("password").(string)
	context := data.Get("context").(string)

	if server == "" || username == "" || password == "" {
		return nil, fmt.Errorf("chkp-provider missing parameters to initialize (server, username, password)")
	}

	args := chkp.ApiClientArgs {
		Port:                    chkp.DefaultPort,
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
		Timeout:                 chkp.TimeOut,
		Sleep:                   chkp.SleepTime,
	}

	switch context {
		case chkp.WebContext:
			s, err := GetSession()
			if err != nil {
				return nil, err
			}
			if s.Sid != "" {
				args.Sid = s.Sid
			}
			mgmt := chkp.APIClient(args)
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
		case chkp.GaiaContext:
			gaia := chkp.APIClient(args)
			_, err := login(gaia, username, password)
			if err != nil {
				return nil, err
			}
			return gaia, nil
		default:
			return nil, fmt.Errorf("Unsupported access context - gaia_api or web_api")
	}
}

// Preform login. Creating new session...
func login(client *chkp.ApiClient, username string, pwd string) (Session, error) {
	log.Printf("Preform login")
	loginRes, err := client.Login(username, pwd,false,"",false,"")
	if err != nil {
		log.Println("Failed to preform login")
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