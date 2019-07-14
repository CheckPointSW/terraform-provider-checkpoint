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
		},
		ResourcesMap: map[string]*schema.Resource{
			"chkp_network": resourceNetwork(),
			"chkp_publish": resourcePublish(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(data *schema.ResourceData) (interface{}, error) {
	log.Printf("Enter providerConfigure...")
	server := data.Get("server").(string)
	username := data.Get("username").(string)
	password := data.Get("password").(string)

	if server == "" || username == "" || password == "" {
		return nil, fmt.Errorf("chkp-provider missing parameters to initialize (server, username, password)")
	}

	var client *chkp.ApiClient
	doLogin := true

	session, err := GetSid()
	if err != nil {
		return nil,err
	}
	client = chkp.APIClient(chkp.ApiClientArgs{
												Port: chkp.DefaultPort,
												Fingerprint: "",
												Sid: session.Sid,
												Server: server,
												ProxyHost: "",
												ProxyPort: -1,
												ApiVersion: "",
												IgnoreServerCertificate: false,
												AcceptServerCertificate: false,
												DebugFile: "deb.txt",
												Context: chkp.WebContext,
												Timeout: chkp.TimeOut,
												Sleep: chkp.SleepTime,
											})
	if CheckSession(client, session.Uid) {
		log.Printf("Client connected with last session (SID = %s)", session.Sid)
		doLogin = false
	}

	// Session not available. Creating new session...
	if doLogin {
		if err = preformLogin(client,username,password); err != nil {
			return nil,err
		}
	}

	return client,nil
}


func preformLogin(client *chkp.ApiClient, username string, pwd string) error {
	log.Printf("Preform login")
	loginRes, err := client.Login(username,pwd,false,"",false,"")
	if err != nil {
		log.Println("Failed to preform login")
		return err
	}
	s := Session {
		Sid: client.GetSessionID(),
		Uid: loginRes.GetData()["uid"].(string),
	}
	if err := s.Save(); err != nil {
		return err
	}
	log.Printf("Client connected with new session (SID = %s)", s.Sid)
	return nil
}