package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/CheckPointSW/terraform-provider-checkpoint/commands"
)

func main() {

	var policyPackage string

	flag.StringVar(&policyPackage, "policy-package", "", "Policy package identified by the name or UID to be verified.")
	flag.Parse()

	apiClient, err := commands.InitClient()
	if err != nil {
		fmt.Println("Verify policy error: " + err.Error())
		os.Exit(1)
	}

	var payload = map[string]interface{}{}
	{
		payload["policy-package"] = policyPackage
	}

	verifyRes, err := apiClient.ApiCall("verify-policy", payload, apiClient.GetSessionID(), true, apiClient.IsProxyUsed())
	if err != nil {
		fmt.Println("Verify policy error: " + err.Error())
		os.Exit(1)
	}

	taskId := commands.ResolveTaskId(verifyRes.GetData())

	if !verifyRes.Success {
		errMsg := fmt.Sprintf("Verify policy failed: %s.", verifyRes.ErrorMsg)
		if taskId != nil {
			errMsg += fmt.Sprintf(" task-id [%s]", taskId)
		}
		fmt.Println(errMsg)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Verify policy finished successfully. task-id [%s]", taskId))
}
