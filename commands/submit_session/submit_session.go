package main

import (
	"fmt"
	"github.com/CheckPointSW/terraform-provider-checkpoint/commands"
	"os"
)

func main() {
	apiClient, err := commands.InitClient()
	if err != nil {
		fmt.Println("Approve Session error: " + err.Error())
		os.Exit(1)
	}

	payload := make(map[string]interface{})
	if len(os.Args) < 2 {
		payload["uid"] = apiClient.GetSessionID()
	} else {
		payload["uid"] = os.Args[1]
	}

	submitSessionRes, err := apiClient.ApiCall("submit-session", payload, apiClient.GetSessionID(), true, apiClient.IsProxyUsed())
	if err != nil {
		fmt.Println("Submit Session error: " + err.Error())
		os.Exit(1)
	}

	if !submitSessionRes.Success {
		errMsg := fmt.Sprintf("Submit Session failed: %s.", submitSessionRes.ErrorMsg)
		fmt.Println(errMsg)
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("Submit Session finished successfully."))
}
