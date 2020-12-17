package main

import (
	"flag"
	"fmt"
	"github.com/CheckPointSW/terraform-provider-checkpoint/commands"
	"os"
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
		fmt.Println("Install policy error: " + err.Error())
		os.Exit(1)
	}

	payload := map[string]interface{}{
		"policy-package": policyPackage,
		"targets":        targets,
	}

	installPolicyRes, err := apiClient.ApiCall("install-policy", payload, apiClient.GetSessionID(), true, false)
	if err != nil {
		fmt.Println("Install policy error: " + err.Error())
		os.Exit(1)
	}

	taskId := commands.ResolveTaskId(installPolicyRes.GetData())

	if !installPolicyRes.Success {
		errMsg := fmt.Sprintf("Install policy failed: %s.", installPolicyRes.ErrorMsg)
		if taskId != nil {
			errMsg += fmt.Sprintf(" task-id [%s]", taskId)
		}
		fmt.Println(errMsg)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Policy installed successfully. task-id [%s]", taskId))
}
