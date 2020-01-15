package main

import (
	"flag"
	"github.com/terraform-providers/terraform-provider-checkpoint/commands"
	"log"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

var targets arrayFlags


func main() {

	var policyPackage string

	flag.StringVar(&policyPackage, "policy-package", "", "The name of the Policy Package to be installed.")
	flag.Var(&targets, "target", "On what targets to execute this command. Targets may be identified by their name, or object unique identifier.")
	flag.Parse()

	apiClient, err := commands.InitClient()
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	payload := map[string]interface{}{
		"policy-package": policyPackage,
		"targets": targets,
	}

	installPolicyRes, err := apiClient.ApiCall("install-policy", payload, apiClient.GetSessionID(),true,false)
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
	if !installPolicyRes.Success {
		log.Fatalf("error: %s", installPolicyRes.ErrorMsg)
	}

	log.Printf("policy installed successfully")

}
